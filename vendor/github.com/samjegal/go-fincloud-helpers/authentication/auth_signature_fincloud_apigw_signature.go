package authentication

import (
	"crypto"
	"strconv"

	"github.com/Azure/go-autorest/autorest"
	"github.com/samjegal/go-fincloud-helpers/security"
)

type serviceApiGWSignatureAuth struct {
	AccessKeyId string
	secretKey   string

	httpMethod string
	requestURL string
}

func (s serviceApiGWSignatureAuth) build(b Builder) (authMethod, error) {
	method := serviceApiGWSignatureAuth{
		AccessKeyId: b.AccessKeyId,
		secretKey:   b.SecretKey,
		httpMethod:  b.HttpMethod,
		requestURL:  b.RequestURL,
	}
	return method, nil
}

func (s serviceApiGWSignatureAuth) isApplicable(b Builder) bool {
	return b.AccessKeyId != "" && b.SecretKey != ""
}

func (s serviceApiGWSignatureAuth) getAuthorizationToken(sender autorest.Sender, endpoint string) (autorest.Authorizer, error) {
	headers := make(map[string]interface{})
	timestamp := strconv.FormatInt(makeTimestamp(), 10)

	sec := security.NewSignature(s.secretKey, crypto.SHA256)
	signature, err := sec.Signature(s.httpMethod, s.requestURL, s.AccessKeyId, timestamp)
	if err != nil {
		return nil, err
	}

	// 1970년 1월 1일 00:00:00 협정 세계시(UTC)부터의 경과 시간을 밀리초(Millisecond)
	headers["x-ncp-apigw-timestamp"] = timestamp

	// 네이버 클라우드 플랫폼 홈페이지 또는 sub account에서 발급받은 Access Key ID
	headers["x-ncp-iam-access-key"] = s.AccessKeyId

	// Body를 Access Key ID와 맵핑되는 Secret Key로 암호화한 서명
	headers["x-ncp-apigw-signature-v2"] = signature

	auth := autorest.NewAPIKeyAuthorizer(headers, nil)
	return auth, nil
}

func (s serviceApiGWSignatureAuth) name() string {
	return "Service Api Gateway Client Signature Certificate"
}
