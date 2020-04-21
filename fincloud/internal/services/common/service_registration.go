package common

import "github.com/hashicorp/terraform-plugin-sdk/helper/schema"

type ServiceRegistration interface {
	// Name 서비스 이름
	Name() string

	// SupportedDataSources 테라폼 데이터 소스 서비스를 정의
	SupportedDataSources() map[string]*schema.Resource

	// SupportedResources 테라폼 리소스의 서비스를 정의
	SupportedResources() map[string]*schema.Resource
}
