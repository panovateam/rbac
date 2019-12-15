package authen

import (
	"github.com/onskycloud/go-redis"
	"github.com/onskycloud/errors"
	"github.com/onskycloud/resource-helper"
	"github.com/onskycloud/rbac/authen/model"
	"github.com/mitchellh/mapstructure"
)


// ServiceName Seperate Char
const ServiceName = "rbac"

const UserCacheKey = ""
func User(c echo.Context) *model.AuthUser {
	id := c.Get("id").(int)
	customerID := c.Get("customer_id").(int)
	customerNumber := c.Get("customer_number").(string)
	user := c.Get("username").(string)
	role := c.Get("role").(int8)
	userUuid := c.Get("user_uuid").(string)

	//check cache
	result := model.AuthUser{
		ID:             id,
		Username:       user,
		CustomerNumber: customerNumber,
		CustomerID:     customerID,
		UserUuid:       userUuid,
		Role:           model.AccessRole(role),
	}
	err := enforceRole(role,model.CustomerUserRole)
	if err!=nil{
		return &result
	}
		customer := helper.GetCustomerCache(customerNumber, self.redis)
		if customer != nil && customer.Clients != nil {
			cliTemp := funk.Map(customer.Clients, func(item model.Customer) string {
				return item.AccNumber
			})
			if cliTemp != nil {
				result.Clients = cliTemp.([]string)
			}
		}
	return &result
}
//only have permission when meet requirement policy and resouce
// params:
// - serviceName string
// - action string
// - resourceType string
// - resourceValue string
// - accountnumber string
// - patrition string
// - region string
// EnforcePolicy enfore policy
func EnforcePolicy(db *redis.Redis,cacheKey string, action string, compareResources ...string) ([]string, error) {
	var policies []model.Policy
	var cacheItem interface{}
	var assignedResources []string

	if action == "" || db == nil{
		return nil ,errors.Forbidden(ServiceName, "enforcePolicy:invalidParams")
	}
	actionCache := &ItemActionCache{}

		var temp interface{}
		err := db.GetObject(cacheKey, action, &temp)
		if err != nil {
			if err.Error()=="redis: nil"{
			// TODO: update cache by get from db
			return nil,errors.Forbidden(ServiceName, "enforcePolicy:notFound")
			}
			return nil, err
		}
		err = mapstructure.Decode(temp.(map[string]interface{}), &actionCache)
			if err != nil {
				return nil,err
			}
	if actionCache.Name == "" {
		return nil,errors.Forbidden(ServiceName, "enforcePolicy:invalidAction")
	}

	auth := User(c)
	if auth == nil {
		return nil, errors.Forbidden(ServiceName, "enforcePolicy:invalidCustomer")
	}
	//---if role is superadmin,admin , bypass check policy (fullacess)
	err = enforceRole(role,model.AdminRole)
	if err == nil{
		var results []string
		for _, v := range compareResources {
			rItem := resourceHelper.GetResourceFormat(v, actionCache, auth.CustomerNumber)
			if rItem == nil {
				return nil, errors.Forbidden(ServiceName, "enforcePolicy:invalidCustomer")
			}
			results = append(results, rItem.ToString())
		}
		return results, nil
	}
	/*------------------------------------------------------------------------------------------------------------*/
	//check resource, if empty set value "*"
	if compareResources == nil || len(compareResources) == 0 {
		compareResources = []string{"*"}
	}
	//-------------get policies of user-------------
		err = db.GetObject(UserCacheKey, auth.UserUuid, &cacheItem)
		if err != nil {
			if err.Error()=="redis: nil"{
				// TODO: update cache by get from db
			// err = self.micro.GetPolicies(auth.UserUuid, &policies)
			return nil,errors.Forbidden(ServiceName, "enforcePolicy:user:notFound")
			}
			return nil,err

		}
			var raw = cacheItem.(map[string]interface{})
			if (raw == nil || raw["policies"] == nil){
				// TODO: get from db
				// err = self.micro.GetPolicies(auth.UserUuid, &policies)
			}
			if raw != nil && raw["policies"] != nil {
				originalPolicies := raw["policies"].([]interface{})
				err := mapstructure.Decode(originalPolicies, &policies)
				if err != nil {
					return nil, errors.Forbidden(ServiceName, "enforcePolicy:policy:wrongData")
				}
			}
			if (raw == nil || raw["policies"] == nil){
				return nil, errors.Forbidden(ServiceName, "enforcePolicy:policy:notFound")
			}

			if policies != nil && len(policies) > 0 {

		// fmt.Println("EnforcePolicy 8.1 - policies ok ")
		funk.ForEach(policies, func(p model.Policy) {
			funk.ForEach(p.ResourceTypes, func(r model.ResourceType) {
				isMatchAction := funk.Contains(r.Actions, action) || funk.Contains(r.Actions, "*")
				if isMatchAction {
					//--refine to make sure that any item does not contain each other.
					resourceArray := funk.Filter(r.Resources, func(item string) bool {
						temp := funk.Find(assignedResources, func(x string) bool {
							isOk, _, _ := helper.ParseResource(item, x, actionCache, auth.CustomerNumber)
							return isOk
						})
						return temp == nil

					}).([]string)

					if resourceArray != nil && len(resourceArray) > 0 {
						funk.ForEach(resourceArray, func(item string) {
							//filter items not match new resource item
							//make sure that any item does not contain each other.
							assignedResources = funk.Filter(assignedResources, func(ors string) bool {
								isOk, _, _ := helper.ParseResource(ors, item, actionCache, auth.CustomerNumber)
								fmt.Printf("EnforcePolicy 8.1 - isOk:%v\n", isOk)
								return !isOk
							}).([]string)
							// fmt.Printf("item:%v\n", item)

							// fmt.Printf("assignedResources:%v\n", assignedResources)
							assignedResources = append(assignedResources, item)

						})

					}
				}
			})
		})
	}
	// fmt.Printf("assignedResources:%v\n", assignedResources)
	// fmt.Printf("compareResources:%v\n", compareResources)
	/*------------------------------------------------------------------------------------------------------------*/
	//check formart of parseResourceStr must follow `orn:%s:%s:%s:%s:%s/%s`
	// fmt.Println("EnforcePolicy 9 - check formart of parseResourceStr must follow orn.....")
	// fmt.Printf("EnforcePolicy 9 - compareResources: %+v\n, condition: %+v\n", compareResources, compareResources == nil || len(compareResources) == 0)
	invalidResource := funk.Find(compareResources, func(item string) bool {
		err := helper.CheckResourceFormat(item, actionCache, auth.CustomerNumber)
		return err != nil
	})

	if invalidResource != nil {
		// fmt.Printf("EnforcePolicy 9.1 - invalid resource != nil: %+v\n", invalidResource)
		return nil, model.ErrUnauthorized
	}
	//
	var resultResources []string
	/*------------------------------------------------------------------------------------------------------------*/
	findResource := func(parseResourceStr string) ([]string, error) {
		//---------------------
		//compareWith parseResourceStr, find resource item  match the request
		// fmt.Println("EnforcePolicy 10 - func findResource - compareWith parseResourceStr, find resource item  match the request")
		for _, element := range assignedResources {
			isOk, dataSourceObject, _ := helper.ParseResource(parseResourceStr, element, actionCache, auth.CustomerNumber)
			// fmt.Printf("ParseResource:%v - element:%v - dataSourceObject :%v\n", parseResourceStr, element, dataSourceObject)
			if !isOk {
				// is all-item
				// only happend when LIST and param *, if the origin resource is not *, ParseResource will be failed
				// fmt.Printf("ParseResource:%v - element:%v\n", dataSourceObject.Resource, actionCache)

				if dataSourceObject.Resource == "*" && actionCache.Type == "List" {
					//if type action is LIST
					//collect all resources which is assigned to user
					resultResources = append(resultResources, assignedResources...)
					break
				}
			} else {
				resultResources = append(resultResources, dataSourceObject.ToString())
				break
			}
		}
		// fmt.Printf("resultResources:%v\n", resultResources)
		// fmt.Printf("resultResources:%v\n", resultResources)
		if resultResources != nil && len(resultResources) > 0 {
			// fmt.Println("EnforcePolicy 10.2 - resultResource > 0 ___ ok")
			return resultResources, nil
		}
		// fmt.Println("EnforcePolicy 10.3 - resultResource is nil or len == 0 ___ not ok")
		return nil, model.ErrUnauthorized

	}
	/*------------------------------------------------------------------------------------------------------------*/
	invalidItem := funk.Find(compareResources, func(resourceItem string) bool {
		// fmt.Printf("resourceItem:%v\n", resourceItem)
		_, err := findResource(resourceItem)
		// if err != nil {
		// 	fmt.Printf("err:%v\n", err)
		// }
		return err != nil
	})
	// on invalid item found
	// fmt.Printf("invalidItem:%v\n", invalidItem)
	if invalidItem == nil {
		// fmt.Println("EnforcePolicy 10.4 - invalid item ok")
		return resultResources, nil
	}
	// fmt.Printf("EnforcePolicy 10.5 - invalid item not ok: %+v\n", invalidItem)
	return nil, model.ErrUnauthorized
}

//--
func enforceCustomer(role int) error {
	return enforceRole(role, model.CustomerUserRole)
}

//--
func enforceAdmin(role int) error {
	return enforceRole(role, model.AdminRole)
}

//-- Client
func enforceClient(role int) error {
	return enforceRole(role, model.ClientAdminRole)
}
func enforceUserClient(role int) error {
	return enforceRole(role, model.ClientUserRole)
}
// EnforceRole authorizes request by AccessRole
func enforceRole(role int, roleBase model.AccessRole) error {
	if role<= roleBase{
		return nil
	}
	return errors.Forbidden(ServiceName, "enforcePolicy:forbidden")
}

// ItemActionCache item cache for action
type ItemActionCache struct {
	Name         string `json:"name,omitempty"`
	ResourceType string `json:"resourceType,omitempty"`
	Type         string `json:"type,omitempty"`
	Service      string `json:"service,omitempty"`
}