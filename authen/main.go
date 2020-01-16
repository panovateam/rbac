package authen

import (
	"context"
	"encoding/json"

	"github.com/labstack/echo"
	"github.com/micro/go-micro/metadata"
	"github.com/onskycloud/errors"
	"github.com/onskycloud/go-redis"
	"github.com/onskycloud/rbac/model"
	"github.com/onskycloud/rbac/utl"
	resourcehelper "github.com/onskycloud/resource-helper"
)

// ServiceName Seperate Char
const ServiceName = "rbac"

// ResourceAll resource key
const ResourceAll = "*"

// ResourceList resource key
const ResourceList = "List"

// CallbackType type of callback func
type CallbackType uint8

const (
	//CUSTOMER callbacktype is customer
	CUSTOMER = iota + 1
	//POLICY callbacktype is policy
	POLICY
)

// Callback func
type Callback func(CallbackType, map[string]string) error

// RBAC model representation
type RBAC struct {
	DB               *redis.Redis
	UserCacheKey     string
	ActionCacheKey   string
	CustomerCacheKey string
	Key              string
	Algo             string
	Callback         Callback
	Cfg              map[string]string
}

// Init init redis receiver
func Init(db *redis.Redis, userCacheKey string, actionCacheKey string, customerCacheKey string, key string, algo string, callback Callback) *RBAC {
	return &RBAC{
		DB:               db,
		UserCacheKey:     userCacheKey,
		ActionCacheKey:   actionCacheKey,
		CustomerCacheKey: customerCacheKey,
		Key:              key,
		Algo:             algo,
		Callback:         callback,
	}
}

// UserFromContext instance
func (r *RBAC) UserFromContext(c echo.Context) *model.AuthUser {
	customer := new(model.Customer)

	id := c.Get("id").(int)
	customerID := c.Get("customer_id").(int)
	customerNumber := c.Get("customer_number").(string)
	user := c.Get("username").(string)
	role := c.Get("role").(uint8)
	userUUID := c.Get("user_uuid").(string)

	//check cache
	result := model.AuthUser{
		ID:             id,
		Username:       user,
		CustomerNumber: customerNumber,
		CustomerID:     customerID,
		UserUuid:       userUUID,
		Role:           model.AccessRole(role),
	}

	err := enforceRole(role, model.CustomerUserRole)
	if err != nil {
		return &result
	}

	err = r.GetDataFromCache(r.CustomerCacheKey, customerNumber, customer)
	if err != nil || customer == nil || customer.Clients == nil {
		return &result
	}

	for i := 0; i < len(customer.Clients); i++ {
		result.Clients = append(result.Clients, customer.Clients[i].AccNumber)
	}

	return &result
}

// UserFromMetadata instance
func (r *RBAC) UserFromMetadata(ctx context.Context) (*model.AuthUser, error) {
	customer := new(model.Customer)

	meta, ok := metadata.FromContext(ctx)
	if !ok {
		return nil, errors.Forbidden(ServiceName, "rbac:authen:UserFromMetadata:invalidParams")
	}

	token, ok := meta["authorization"]
	// support UPERCASE for rpc
	if !ok {
		token, ok = meta["Authorization"]
	}
	if !ok {
		return nil, errors.Forbidden(ServiceName, "rbac:authen:UserFromMetadata:missingToken")
	}
	result, err := utl.ParseToken(r.Key, r.Algo, token)
	if err != nil {
		return nil, err
	}

	err = enforceRole(uint8(result.Role), model.CustomerUserRole)
	if err != nil {
		return result, nil
	}

	err = r.GetDataFromCache(r.CustomerCacheKey, result.CustomerNumber, customer)
	if err != nil || customer == nil || customer.Clients == nil {
		return result, nil
	}

	for i := 0; i < len(customer.Clients); i++ {
		result.Clients = append(result.Clients, customer.Clients[i].AccNumber)
	}

	return result, nil
}

// EnforcePolicy only have permission when meet requirement policy and resouce
// params:
// - serviceName string
// - action string
// - resourceType string
// - resourceValue string
// - accountnumber string
// - patrition string
// - region string
// EnforcePolicy enfore policy
func (r *RBAC) EnforcePolicy(role uint8, customerNumber string, userUUID string, action string, compareResources ...string) ([]string, error) {
	user := &model.User{}
	actionCache := &resourcehelper.ItemActionCache{}

	if action == "" || r == nil || role == 0 || customerNumber == "" || userUUID == "" {
		return nil, errors.Forbidden(ServiceName, "enforcePolicy:invalidParams")
	}

	err := r.GetDataFromCache(r.ActionCacheKey, action, actionCache)
	if err != nil {
		return nil, err
	}
	if actionCache.Name == "" {
		return nil, errors.Forbidden(ServiceName, "enforcePolicy:invalidAction")
	}

	results, err := checkAdmin(role, compareResources, actionCache, customerNumber)
	if err != nil {
		return nil, err
	}
	//---if role is superadmin,admin , bypass check policy (fullacess)
	if results != nil {
		return results, nil
	}

	//check resource, if empty set value "*"
	if compareResources == nil || len(compareResources) == 0 {
		compareResources = []string{ResourceAll}
	}
	//-------------get policies of user-------------
	err = r.GetDataFromCache(r.UserCacheKey, userUUID, user)
	if err != nil {
		return nil, err
	}

	assignedResources := getAssignedResources(action, actionCache, customerNumber, user.Policies)
	err = checkInvalidResource(compareResources, actionCache, customerNumber)
	if err != nil {
		return nil, err
	}

	result := getResources(compareResources, assignedResources, actionCache, customerNumber)
	if result == nil || len(result) == 0 {
		return nil, errors.Forbidden(ServiceName, "enforcePolicy:restricted")
	}
	return result, nil
}

// EnforceRole authorizes request by AccessRole
func enforceRole(role uint8, roleBase model.AccessRole) error {
	if role <= uint8(roleBase) {
		return nil
	}
	return errors.Forbidden(ServiceName, "enforcePolicy:forbidden")
}
func getAssignedResources(action string, actionCache *resourcehelper.ItemActionCache, customerNumber string, policies []model.Policy) []string {
	result := []string{}
	resources := []string{}
	if policies == nil || len(policies) == 0 {
		return result
	}
	for i := 0; i < len(policies); i++ {
		for j := 0; j < len(policies[i].ResourceTypes); j++ {
			resourceType := policies[i].ResourceTypes[j]
			isMatch := resourcehelper.Contains(resourceType.Actions, action) || resourcehelper.Contains(resourceType.Actions, "*")
			if !isMatch {
				continue
			}
			//--refine to make sure that any item does not contain each other.
			for k := 0; k < len(resourceType.Resources); k++ {
				temp := ""
				for m := 0; m < len(result); m++ {
					ok, _, _ := resourcehelper.ParseResource(resourceType.Resources[k], result[m], actionCache, customerNumber)
					if !ok {
						continue
					}
					temp = result[m]
				}
				if temp != "" {
					continue
				}
				resources = append(resources, resourceType.Resources[k])
			}
			if len(resources) == 0 {
				continue
			}

			//filter items not match new resource item
			//make sure that any item does not contain each other.
			for k := 0; k < len(resources); k++ {
				assignedResources := []string{}
				for m := 0; m < len(result); m++ {
					ok, _, _ := resourcehelper.ParseResource(result[m], resources[k], actionCache, customerNumber)
					if ok {
						continue
					}
					assignedResources = append(assignedResources, result[m])
				}
				result = append(result, resources[k])
			}
		}
	}
	return result
}
func checkInvalidResource(compareResources []string, actionCache *resourcehelper.ItemActionCache, customerNumber string) error {
	for i := 0; i < len(compareResources); i++ {
		result := resourcehelper.GetResourceFormat(compareResources[i], actionCache, customerNumber)
		if result == nil {
			return errors.Forbidden(ServiceName, "enforcePolicy:invalidResource")
		}
	}
	return nil
}
func getResources(compareResources []string, assignedResources []string, actionCache *resourcehelper.ItemActionCache, customerNumber string) []string {
	result := []string{}

	for i := 0; i < len(compareResources); i++ {
		for j := 0; j < len(assignedResources); j++ {
			ok, dataSourceObject, _ := resourcehelper.ParseResource(compareResources[i], assignedResources[j], actionCache, customerNumber)
			if ok {
				result = append(result, dataSourceObject.ToString())
			}
			if dataSourceObject.Resource == ResourceAll && actionCache.Type == ResourceList {
				result = append(result, assignedResources...)
			}
		}
	}
	return result
}

// GetDataFromCache get data from cache
func (r *RBAC) GetDataFromCache(key string, field string, result interface{}) error {
	var temp interface{}
	err := r.DB.GetObject(key, field, &temp)
	if err.Error() == errors.RedisEmpty {
		//  update cache by callback
		input := map[string]string{
			"userUUID": key,
		}
		err = r.Callback(POLICY, input)
		if err != nil {
			return err
		}
		err = r.DB.GetObject(key, field, &temp)
		if err.Error() == errors.RedisEmpty {
			return errors.NotFound(ServiceName, "enforcePolicy:notFound")
		}
	}
	if err != nil {
		return err
	}
	b, err := json.Marshal(temp)
	if err != nil {
		return errors.InternalServerError(ServiceName, "enforcePolicy:marshalProblem")
	}
	json.Unmarshal(b, result)

	if result == nil {
		return errors.InternalServerError(ServiceName, "enforcePolicy:wrongData")
	}
	return nil
}

func checkAdmin(role uint8, compareResources []string, actionCache *resourcehelper.ItemActionCache, customerNumber string) ([]string, error) {
	var results []string
	err := enforceRole(role, model.AdminRole)
	if err != nil {
		return nil, nil
	}
	for i := 0; i < len(compareResources); i++ {
		resource := resourcehelper.GetResourceFormat(compareResources[i], actionCache, customerNumber)
		if resource == nil {
			return nil, errors.InternalServerError(ServiceName, "enforcePolicy:invalidCustomer")
		}
		results = append(results, resource.ToString())
	}

	return results, nil
}
