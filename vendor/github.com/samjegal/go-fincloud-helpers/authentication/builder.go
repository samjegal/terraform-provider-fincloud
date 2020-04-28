package authentication

import (
	"fmt"
	"log"
)

type Builder struct {
	CertTokenPath string
	Environment   string

	AccessKeyId string
	SecretKey   string

	Subaccount string
	Username   string
	Password   string

	HttpMethod string
	RequestURL string
}

func (b Builder) Build() (*Config, error) {
	config := Config{
		Environment: "FINCLOUD", //b.Environment,
	}

	supportAuthenticationMethods := []authMethod{
		serviceClientCertificateAuth{},
		serviceNsaCertkeyAuth{},
		serviceApiGWSignatureAuth{},
	}

	for _, method := range supportAuthenticationMethods {
		name := method.name()
		log.Printf("Testing if %s is applicable for Authentication..", name)

		if !method.isApplicable(b) {
			continue
		}

		log.Printf("Using %s for Authentication", name)
		auth, err := method.build(b)
		if err != nil {
			return nil, err
		}

		config.authMethod = auth

		return &config, nil
	}

	return nil, fmt.Errorf("No support authenticattion methods were found")
}
