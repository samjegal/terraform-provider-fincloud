package authentication

import (
	"crypto"
	"strconv"

	"github.com/Azure/go-autorest/autorest"
	"github.com/samjegal/go-fincloud-helpers/security"
)

type serviceApiGWSignatureAuth struct {
	// accessKeyId
	accessKey string

	// secretKey
	secretKey string

	// apiGatewayKey API Gateway의 product 생성시 발급되는 키 정보
	apiGatewayKey string

	httpMethod string
	requestURL string
}

func (s serviceApiGWSignatureAuth) build(b Builder) (authMethod, error) {
	method := serviceApiGWSignatureAuth{
		apiGatewayKey: b.ApiGatewayKey,
		accessKey:     b.AccessKey,
		secretKey:     b.SecretKey,
		httpMethod:    b.HttpMethod,
		requestURL:    b.RequestURL,
	}
	return method, nil
}

func (s serviceApiGWSignatureAuth) isApplicable(b Builder) bool {
	return b.AccessKey != "" && b.SecretKey != ""
}

func (s serviceApiGWSignatureAuth) getAuthorizationToken(sender autorest.Sender, endpoint string) (autorest.Authorizer, error) {
	headers := make(map[string]interface{})
	timestamp := strconv.FormatInt(makeTimestamp(), 10)

	sec := security.NewSignature(s.secretKey, crypto.SHA256)
	signature, err := sec.Signature(s.httpMethod, s.requestURL, s.accessKey, timestamp)
	if err != nil {
		return nil, err
	}

	if s.apiGatewayKey != "" {
		headers["x-ncp-apigw-api-key"] = s.apiGatewayKey
		headers["x-ncp-dmn_cd"] = "FIN"
		headers["x-ncp-region_code"] = "FKR"
	}

	// 1970년 1월 1일 00:00:00 협정 세계시(UTC)부터의 경과 시간을 밀리초(Millisecond)
	headers["x-ncp-apigw-timestamp"] = timestamp

	// 네이버 클라우드 플랫폼 홈페이지 또는 sub account에서 발급받은 Access Key ID
	headers["x-ncp-iam-access-key"] = s.accessKey

	// Body를 Access Key ID와 맵핑되는 Secret Key로 암호화한 서명
	headers["x-ncp-apigw-signature-v2"] = signature

	auth := autorest.NewAPIKeyAuthorizer(headers, nil)
	return auth, nil
}

func (s serviceApiGWSignatureAuth) name() string {
	return "Service Api Gateway Client Signature Certificate"
}
