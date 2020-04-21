package compute

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/samjegal/fincloud-sdk-for-go/profiles/compute"
	"github.com/samjegal/terraform-provider-fincloud/fincloud/helpers/validate"
	"github.com/samjegal/terraform-provider-fincloud/fincloud/internal/clients"
	"github.com/samjegal/terraform-provider-fincloud/fincloud/internal/timeouts"
	"github.com/samjegal/terraform-provider-fincloud/fincloud/utils"
)

var initScriptResourceName = "fincloud_init_script"

func resourceInitScript() *schema.Resource {
	return &schema.Resource{
		Create: resourceInitScriptCreate,
		Read:   resourceInitScriptRead,
		Update: resourceInitScriptUpdate,
		Delete: resourceInitScriptDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(60 * time.Minute),
			Read:   schema.DefaultTimeout(10 * time.Minute),
			Update: schema.DefaultTimeout(60 * time.Minute),
			Delete: schema.DefaultTimeout(60 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validate.NoEmptyStrings,
			},

			"ostype": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: validation.StringInSlice([]string{
					"linux",
					"windows",
				}, true),
			},

			"content": {
				Type:     schema.TypeString,
				Required: true,
			},

			"description": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validate.NoEmptyStrings,
			},
		},
	}
}

func resourceInitScriptCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client).Compute.InitScriptClient
	ctx, cancel := timeouts.ForCreate(meta.(*clients.Client).StopContext, d)
	defer cancel()

	name := d.Get("name").(string)

	initScript, err := expandInitScriptParameter(d)
	if err != nil {
		return err
	}

	resp, err := client.Create(ctx, *initScript)
	if err != nil {
		return err
	}

	if resp.Response.StatusCode == http.StatusBadRequest {
		return fmt.Errorf("초기화 스크립트를 생성 중 문제가 발생했습니다. InitScript: %q: %+v", name, err)
	}

	id, err := initScriptNumber(ctx, meta, name)
	if err != nil {
		return err
	}

	d.SetId(id)

	return resourceInitScriptRead(d, meta)
}

func resourceInitScriptRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client).Compute.InitScriptClient
	ctx, cancel := timeouts.ForCreate(meta.(*clients.Client).StopContext, d)
	defer cancel()

	_, err := client.Get(ctx, d.Id())
	if err != nil {
		return err
	}

	return nil
}

func resourceInitScriptUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client).Compute.InitScriptClient
	ctx, cancel := timeouts.ForCreate(meta.(*clients.Client).StopContext, d)
	defer cancel()

	initScriptId := d.Id()

	initScript, err := expandInitScriptNumberParameter(d, initScriptId)
	if err != nil {
		return err
	}

	_, err = client.Update(ctx, initScriptId, *initScript)
	if err != nil {
		return err
	}

	return resourceInitScriptRead(d, meta)
}

func resourceInitScriptDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client).Compute.InitScriptClient
	ctx, cancel := timeouts.ForCreate(meta.(*clients.Client).StopContext, d)
	defer cancel()

	_, err := client.Delete(ctx, expandInitScriptNumberListParameter(d, d.Id()))
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}

func initScriptNumber(ctx context.Context, meta interface{}, name string) (string, error) {
	client := meta.(*clients.Client).Compute.InitScriptClient

	resp, err := client.List(ctx, "1", "100")
	if err != nil {
		return "", fmt.Errorf("초기화 스크립트 전체 리스트 정보를 가져오는 중 에러가 발생했습니다. InitScript: %q: %+v", name, err)
	}

	for _, v := range *resp.Content {
		if name == *v.InitConfigurationScriptName {
			return *v.InitConfigurationScriptNo, nil
		}
	}

	return "", fmt.Errorf("초기화 스크립트 정보가 존재하지 않습니다. InitScript: %q: %+v", name, err)
}

func expandInitScriptParameter(d *schema.ResourceData) (*compute.InitScriptParameter, error) {
	param := &compute.InitScriptParameter{
		InitConfigurationScriptName:        utils.String(d.Get("name").(string)),
		InitConfigurationScriptType:        utils.String("MANU"),
		InitConfigurationScriptContent:     utils.String(d.Get("content").(string)),
		InitConfigurationScriptDescription: utils.String(d.Get("description").(string)),
		UseYn:                              utils.String("Y"),
	}

	if v := d.Get("ostype").(string); v == "linux" {
		param.OsTypeCode = utils.String("LNX")
	} else if v == "windows" {
		param.OsTypeCode = utils.String("WND")
	}

	return param, nil
}

func expandInitScriptNumberListParameter(d *schema.ResourceData, initScriptId string) compute.InitScriptNumberListParameter {
	initScriptNumberList := make([]string, 0)
	initScriptNumberList = append(initScriptNumberList, initScriptId)

	return compute.InitScriptNumberListParameter{
		InitConfigurationScriptNoList: &initScriptNumberList,
	}
}

func expandInitScriptNumberParameter(d *schema.ResourceData, initScriptId string) (*compute.InitScriptNumberParameter, error) {
	param := &compute.InitScriptNumberParameter{
		InitConfigurationScriptNo:      utils.String(initScriptId),
		InitConfigurationScriptType:    utils.String("MANU"),
		InitConfigurationScriptContent: utils.String(d.Get("content").(string)),
		UseYn:                          utils.String("Y"),
	}

	if v := d.Get("ostype").(string); v == "linux" {
		param.OsTypeCode = utils.String("LNX")
	} else if v == "windows" {
		param.OsTypeCode = utils.String("WND")
	}

	return param, nil
}
