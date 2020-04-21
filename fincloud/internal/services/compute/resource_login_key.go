package compute

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/samjegal/fincloud-sdk-for-go/services/compute"
	"github.com/samjegal/terraform-provider-fincloud/fincloud/helpers/validate"
	"github.com/samjegal/terraform-provider-fincloud/fincloud/internal/clients"
	"github.com/samjegal/terraform-provider-fincloud/fincloud/internal/timeouts"
)

var loginKeyResourceName = "fincloud_login_key"

func resourceLoginKey() *schema.Resource {
	return &schema.Resource{
		Create: resourceLoginKeyCreate,
		Read:   resourceLoginKeyRead,
		Update: resourceLoginKeyUpdate,
		Delete: resourceLoginKeyDelete,

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

			"path": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validate.NoEmptyStrings,
			},

			"key": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceLoginKeyCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client).Compute.LoginKeyClient
	ctx, cancel := timeouts.ForCreate(meta.(*clients.Client).StopContext, d)
	defer cancel()

	name := d.Get("name").(string)

	resp, err := client.Create(ctx, name)
	if err != nil || resp.Response.StatusCode == http.StatusBadRequest {
		return fmt.Errorf("로그인 키 생성 중 에러가 발생했습니다. LoginKey: %q: %+v", name, err)
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)

	d.Set("key", buf.String())

	if v, ok := d.GetOk("path"); ok {
		pem, err := os.Create(v.(string) + name + ".pem")
		defer pem.Close()
		if err != nil {
			return fmt.Errorf("로그인 키 파일 생성 중 에러가 발생했습니다. %q: %+v", name, err)
		}

		_, err = pem.WriteString(buf.String())
		pem.Sync()
	}

	d.SetId(name)

	return nil
}

func resourceLoginKeyRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceLoginKeyUpdate(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceLoginKeyDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client).Compute.LoginKeyClient
	ctx, cancel := timeouts.ForCreate(meta.(*clients.Client).StopContext, d)
	defer cancel()

	name := d.Get("name").(string)

	_, err := client.Delete(ctx, expandLoginKeyListParameter(name))
	if err != nil {
		return err
	}

	if v, ok := d.GetOk("path"); ok {
		err = os.Remove(v.(string) + name + ".pem")
		if err != nil {
			return fmt.Errorf("로그인 키 파일 삭제 중 에러가 발생했습니다. %q: %+v", name, err)
		}
	}

	d.SetId("")

	return nil
}

func expandLoginKeyListParameter(name string) compute.LoginKeyListParameter {
	loginKeyNameList := make([]string, 0)
	loginKeyNameList = append(loginKeyNameList, name)

	return compute.LoginKeyListParameter{
		KeyNameList: &loginKeyNameList,
	}
}
