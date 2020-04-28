package authentication

import (
	"fmt"
	"io/ioutil"

	"github.com/Azure/go-autorest/autorest"
)

type serviceClientCertificateAuth struct {
	certTokenPath string
}

func (s serviceClientCertificateAuth) build(b Builder) (authMethod, error) {
	method := serviceClientCertificateAuth{
		certTokenPath: b.CertTokenPath,
	}
	return method, nil
}

func (s serviceClientCertificateAuth) isApplicable(b Builder) bool {
	return b.CertTokenPath != ""
}

func (s serviceClientCertificateAuth) getAuthorizationToken(sender autorest.Sender, endpoint string) (autorest.Authorizer, error) {
	certificateData, err := ioutil.ReadFile(s.certTokenPath)
	if err != nil {
		return nil, fmt.Errorf("Error reading Client Certificate %q: %v", s.certTokenPath, err)
	}

	staticAuth := "ncp_lang=ko-KR; ncp_locale=ko_KR; ncp_version=v2; ncp_region=FKR; ncp=" + string(certificateData)

	headers := make(map[string]interface{})
	headers["Cookie"] = staticAuth

	auth := autorest.NewAPIKeyAuthorizer(headers, nil)
	return auth, nil
}

func (s serviceClientCertificateAuth) name() string {
	return "Service Client Certificate"
}
