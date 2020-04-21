package acceptance

import "os"

type ClientData struct {
	CertTokenPath string
}

func (td TestData) Client() ClientData {
	return ClientData{
		CertTokenPath: os.Getenv("FINCLOUD_CERTIFICATE_PATH"),
	}
}
