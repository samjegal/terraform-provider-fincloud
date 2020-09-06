package authentication

import (
	"github.com/Azure/go-autorest/autorest"
)

type serviceClientAccessAuth struct {
	accessKeyId string
}

func (s serviceClientAccessAuth) isApplicable(b Builder) bool {
	return b.AccessKey != ""
}

func (s serviceClientAccessAuth) build(b Builder) (authMethod, error) {
	method := serviceClientAccessAuth{
		accessKeyId: b.AccessKey,
	}
	return method, nil
}

func (s serviceClientAccessAuth) getAuthorizationToken(sender autorest.Sender, endpoint string) (autorest.Authorizer, error) {
	headers := make(map[string]interface{})

	// x-ncp-iam-access-key 네이버 클라우드 플랫폼 홈페이지 또는 sub account에서 발급받은 Access Key ID
	headers["x-ncp-iam-access-key"] = s.accessKeyId

	auth := autorest.NewAPIKeyAuthorizer(headers, nil)
	return auth, nil
}

func (s serviceClientAccessAuth) name() string {
	return "Service Client Access Certificate"
}
