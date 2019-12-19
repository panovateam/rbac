package authen

import (
	"testing"

	"github.com/onskycloud/go-redis"
	"github.com/onskycloud/rbac/model"
	resourcehelper "github.com/onskycloud/resource-helper"
)

// UserCacheKey cache key for user
const UserCacheKey = "1"

// ActionCacheKey cache key for user
const ActionCacheKey = "2"

// CustomerCacheKey cache key for customer
const CustomerCacheKey = "3"
// Key cache key for customer
const Key = "4"
// Algo cache key for customer
const Algo = "5"
func TestContainer(t *testing.T) {
	role := 3
	customerNumber := "customer_number"
	userUUID := "user_id"
	action := "iot:editThing"
	compareResources := []string{"*"}
	options := Redis{
		Addr: "192.168.0.52:30333",
	}
	userModel := model.User{
		Uuid: "user_id",
		Policies: []model.Policy{
			{
				ID:          1,
				PolicyId:    "8428a75a",
				Name:        "AdminClient",
				Description: "admin acess IOT service for client",
				Type:        "system",
				Version:     "16-11-2019",
				ResourceTypes: []model.ResourceType{
					{
						Name:        "things",
						Effect: "Allow",
						ResouceType: "things",
						Resources:     []string{"*"},
						Actions:   []string{"iot:listThing",
						"iot:editThing",
						"iot:readThing",
						"iot:controlThing"},
					},
				},
			},
		},
	}
	actionCache := resourcehelper.ItemActionCache{
		Name:         "things",
		ResourceType: "things",
		Type:         "system",
		Service:      "iot",
	}

	client := redis.NewRedis(&options)
	client.SetObject("onsky:authen:users", userUUID, userModel)
	client.SetObject("onsky:authen:actions", action, actionCache)


	rbac:= Init(client,UserCacheKey,ActionCacheKey,CustomerCacheKey,Key,Algo)

	testEnforcePolicy(t, rbac, uint8(role), customerNumber, userUUID, action, compareResources...)

}
func testEnforcePolicy(t *testing.T, rbac *RBAC, role uint8, customerNumber string, userUUID string, action string, compareResources ...string) {
	resources, err := rbac.EnforcePolicy( role, customerNumber, userUUID, action, compareResources...)
	if resources == nil || err != nil {
		t.Fatalf("policy reject: %+v\n", err)
	}
}
// Redis configuration
type Redis struct {
	Addr     string `yaml:"addr,omitempty"`
	Password string `yaml:"password,omitempty"`
}
