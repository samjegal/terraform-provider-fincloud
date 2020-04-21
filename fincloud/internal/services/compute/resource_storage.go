package compute

import (
	"fmt"
	"net/http"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/samjegal/fincloud-sdk-for-go/profiles/compute"
	"github.com/samjegal/terraform-provider-fincloud/fincloud/helpers/validate"
	"github.com/samjegal/terraform-provider-fincloud/fincloud/internal/clients"
	"github.com/samjegal/terraform-provider-fincloud/fincloud/internal/timeouts"
	"github.com/samjegal/terraform-provider-fincloud/fincloud/utils"
)

var storageResourceName = "fincloud_storage"

func resourceStorage() *schema.Resource {
	return &schema.Resource{
		Create: resourceStorageCreate,
		Read:   resourceStorageRead,
		Update: resourceStorageUpdate,
		Delete: resourceStorageDelete,

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
			"server_id": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validate.NoEmptyStrings,
			},

			"name": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validate.NoEmptyStrings,
			},

			"type": {
				Type:     schema.TypeString,
				Required: true,
				ValidateFunc: validation.StringInSlice([]string{
					"hdd",
					"ssd",
				}, true),
			},

			"size": {
				Type:     schema.TypeInt,
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

func resourceStorageCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client)
	ctx, cancel := timeouts.ForCreate(meta.(*clients.Client).StopContext, d)
	defer cancel()
	storageClient := client.Compute.StorageClient

	name := d.Get("name").(string)

	resp, err := storageClient.Create(ctx, expandStorageParameter(d))
	if err != nil || resp.Response.StatusCode == http.StatusBadRequest {
		return fmt.Errorf("스토리지 생성 중 에러가 발생했습니다. Storage: %q: %+v", name, err)
	}

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"initialized", "attaching", "attached", "detaching", "creating"},
		Target:     []string{"attached"},
		Refresh:    storageStateRefreshFunc(client, name),
		Timeout:    30 * time.Minute,
		MinTimeout: 15 * time.Second,
	}
	if _, err := stateConf.WaitForState(); err != nil {
		return fmt.Errorf("스토리지 상세정보의 상태 코드를 가져오는 중 문제가 발생했다. Storage: %q: %+v", name, err)
	}

	id, err := storageInstanceNumber(client, name)
	if err != nil {
		return err
	}

	d.SetId(id)

	return resourceStorageRead(d, meta)
}

func resourceStorageRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client).Compute.StorageClient
	ctx, cancel := timeouts.ForRead(meta.(*clients.Client).StopContext, d)
	defer cancel()

	name := d.Get("name").(string)

	_, err := client.List(ctx, "1", "200", "", "")
	if err != nil {
		return fmt.Errorf("전체 스토리지 정보를 가져오는 중 에러가 발생했습니다. Storage: %q: %+v", name, err)
	}

	return nil
}

func resourceStorageUpdate(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceStorageDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client)
	ctx, cancel := timeouts.ForCreate(meta.(*clients.Client).StopContext, d)
	defer cancel()
	storageClient := client.Compute.StorageClient

	name := d.Get("name").(string)
	storageId := d.Id()

	_, err := storageClient.Delete(ctx, expandStorageDetachAndDeleteParameter(storageId))
	if err != nil {
		return err
	}

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"attaching", "detaching", "attached", "initialized", "creating"},
		Target:     []string{"notfound"},
		Refresh:    storageStateRefreshFunc(client, name),
		Timeout:    30 * time.Minute,
		MinTimeout: 15 * time.Second,
	}
	if _, err := stateConf.WaitForState(); err != nil {
		return fmt.Errorf("스토리지 상세정보의 상태 코드를 가져오는 중 문제가 발생했다. Storage: %q: %+v", name, err)
	}

	d.SetId("")

	return nil
}

func expandStorageParameter(d *schema.ResourceData) compute.StorageParameter {
	param := compute.StorageParameter{
		ServerInstanceNo:    utils.String(d.Get("server_id").(string)),
		DiskType2DetailCode: compute.DiskType2DetailCode(d.Get("type").(string)),
		VolumeName:          utils.String(d.Get("name").(string)),
		VolumeSize:          utils.Int32(int32(d.Get("size").(int))),
	}

	if v, ok := d.GetOk("description"); ok {
		param.InstanceDesc = utils.String(v.(string))
	}

	return param
}

func storageStateRefreshFunc(client *clients.Client, name string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		ctx := client.StopContext
		resp, err := client.Compute.StorageClient.List(ctx, "1", "200", "", "")
		if err != nil {
			return nil, "notfound", fmt.Errorf("해당하는 스토리지 정보가 존재하지 않습니다. Storage: %q: %+v", name, err)
		}

		for _, v := range *resp.Content {
			if name == *v.VolumeName {
				return v, string(compute.StorageStatusName(v.InstanceStatusName)), nil
			}
		}

		return resp, "notfound", nil
	}
}

func expandStorageDetachAndDeleteParameter(storageId string) compute.StorageDetachAndDeleteParameter {
	storageIdList := make([]string, 0)
	storageIdList = append(storageIdList, storageId)

	return compute.StorageDetachAndDeleteParameter{
		InstanceNoList: &storageIdList,
	}
}

func storageInstanceNumber(client *clients.Client, name string) (string, error) {
	ctx := client.StopContext
	resp, err := client.Compute.StorageClient.List(ctx, "1", "200", "", "")
	if err != nil {
		return "", fmt.Errorf("해당하는 스토리지 정보가 존재하지 않습니다. Storage: %q: %+v", name, err)
	}

	for _, v := range *resp.Content {
		if name == *v.VolumeName {
			return fmt.Sprintf("%v", *v.InstanceNo), nil
		}
	}

	return "", fmt.Errorf("해당하는 스토리지 정보가 존재하지 않습니다. Storage: %q", name)
}
