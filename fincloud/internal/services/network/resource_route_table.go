package network

// import (
// 	"time"

// 	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
// )

// var routeTableResourceName = "fincloud_route_table"

// func resourceRouteTable() *schema.Resource {
// 	return &schema.Resource{
// 		Create: resourceRouteTableCreate,
// 		Read:   resourceRouteTableRead,
// 		Update: resourceRouteTableUpdate,
// 		Delete: resourceRouteTableDelete,

// 		Importer: &schema.ResourceImporter{
// 			State: schema.ImportStatePassthrough,
// 		},

// 		Timeouts: &schema.ResourceTimeout{
// 			Create: schema.DefaultTimeout(60 * time.Minute),
// 			Read:   schema.DefaultTimeout(10 * time.Minute),
// 			Update: schema.DefaultTimeout(60 * time.Minute),
// 			Delete: schema.DefaultTimeout(60 * time.Minute),
// 		},

// 		Schema: map[string]*schema.Schema{
// 			"name": {},

// 			"vpc_id": {},

// 			"usage": {},

// 			"subnet": {
// 				Type:     schema.TypeList,
// 				Optional: true,
// 				Elem: &schema.Resource{
// 					Schema: map[string]*schema.Schema{
// 						"id": {},
// 					},
// 				},
// 			},

// 			"route": {
// 				Type:     schema.TypeSet,
// 				Optional: true,
// 				Elem: &schema.Resource{
// 					Schema: map[string]*schema.Schema{
// 						"id": {},

// 						"cidr_block": {},

// 						"endpoint": {},
// 					},
// 				},
// 			},

// 			"description": {},
// 		},
// 	}
// }

// func resourceRouteTableCreate(d *schema.ResourceData, meta interface{}) error {
// 	return nil
// }

// func resourceRouteTableRead(d *schema.ResourceData, meta interface{}) error {
// 	return nil
// }

// func resourceRouteTableUpdate(d *schema.ResourceData, meta interface{}) error {
// 	return nil
// }

// func resourceRouteTableDelete(d *schema.ResourceData, meta interface{}) error {
// 	return nil
// }
