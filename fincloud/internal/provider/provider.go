package provider

import (
	"fmt"
	"log"
	"os"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/samjegal/go-fincloud-helpers/authentication"
	"github.com/samjegal/terraform-provider-fincloud/fincloud/internal/clients"
)

func FincloudProvider() terraform.ResourceProvider {
	return fincloudProvider()
}

func TestFincloudProvider() terraform.ResourceProvider {
	return fincloudProvider()
}

func fincloudProvider() terraform.ResourceProvider {
	var debugLog = func(f string, v ...interface{}) {
		if os.Getenv("TF_LOG") == "" {
			return
		}

		if os.Getenv("TF_ACC") != "" {
			return
		}

		log.Printf(f, v...)
	}

	dataSources := make(map[string]*schema.Resource)
	resources := make(map[string]*schema.Resource)

	for _, service := range SupportedServices() {
		debugLog("[DEBUG] Registering Data Sources for %q..", service.Name())
		for k, v := range service.SupportedDataSources() {
			if existing := dataSources[k]; existing != nil {
				panic(fmt.Sprintf("An existing Data Source exists for %q", k))
			}

			dataSources[k] = v
		}

		debugLog("[DEBUG] Registering Resources for %q..", service.Name())
		for k, v := range service.SupportedResources() {
			if existing := resources[k]; existing != nil {
				panic(fmt.Sprintf("An existing Resource exists for %q", k))
			}

			resources[k] = v
		}
	}

	p := &schema.Provider{
		Schema: map[string]*schema.Schema{
			"cert_token_path": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("FINCLOUD_CERTIFICATE_PATH", ""),
				Description: "금융 클라우드 홈페이지를 로그인 성공했을 때 나오는 ncp 쿠키 값",
			},

			"access_key_id": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("FINCLOUD_ACCESS_KEY_ID", ""),
				Description: "금융 클라우드 플랫폼 홈페이지 또는 부 계정에서 발급받은 Access Key ID",
			},

			"secret_key": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("FINCLOUD_SECRET_KEY_ID", ""),
				Description: "금융 클라우드 플랫폼 홈페이지 또는 부 계정에서 발급받은 Secret Key",
			},

			"subaccount": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("FINCLOUD_SUBACCOUNT_URL", ""),
				Description: "금융 클라우드 플랫폼에 접속하기 위한 부계정 정보, 공백이면 메인계정을 접속",
			},

			"username": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("FINCLOUD_USERNAME", ""),
				Description: "금융 클라우드 플랫폼에서 발급받은 사용자 계정",
			},

			"password": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("FINCLOUD_PASSWORD", ""),
				Description: "금융 클라우드 플랫폼에서 발급받은 사용자 비밀번호",
			},
		},

		DataSourcesMap: dataSources,
		ResourcesMap:   resources,
	}

	p.ConfigureFunc = providerConfigure(p)

	return p
}

func providerConfigure(p *schema.Provider) schema.ConfigureFunc {
	return func(d *schema.ResourceData) (interface{}, error) {
		builder := &authentication.Builder{
			// 인증 토큰 정보의 경로 값 설정
			CertTokenPath: d.Get("cert_token_path").(string),

			// API GW를 사용하기 위한 KEY 값 설정
			AccessKeyId: d.Get("access_key_id").(string),
			SecretKey:   d.Get("secret_key").(string),

			// 계정 정보를 이용해 접속하기 위한 기본 값 설정
			Subaccount: d.Get("subaccount").(string),
			Username:   d.Get("username").(string),
			Password:   d.Get("password").(string),
		}

		config, err := builder.Build()
		if err != nil {
			return nil, fmt.Errorf("Error building Financial Cloud Resource Manager Client: %s", err)
		}

		terraformVersion := p.TerraformVersion
		if terraformVersion == "" {
			terraformVersion = "0.12+compatible"
		}

		clientBuilder := clients.ClientBuilder{
			AuthConfig:       config,
			TerraformVersion: terraformVersion,
		}
		client, err := clients.Build(p.StopContext(), clientBuilder)
		if err != nil {
			return nil, err
		}

		client.StopContext = p.StopContext()

		return client, nil
	}
}
