package compute

// FINCLOUD_APACHE_NO_VERSION

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"io"
	"net/http"
)

// RootPasswordClient is the compute Client
type RootPasswordClient struct {
	BaseClient
}

// NewRootPasswordClient creates an instance of the RootPasswordClient client.
func NewRootPasswordClient() RootPasswordClient {
	return NewRootPasswordClientWithBaseURI(DefaultBaseURI)
}

// NewRootPasswordClientWithBaseURI creates an instance of the RootPasswordClient client.
func NewRootPasswordClientWithBaseURI(baseURI string) RootPasswordClient {
	return RootPasswordClient{NewWithBaseURI(baseURI)}
}

// Get 서버 관리자 패스워드 확인
// Parameters:
// instanceNo - server instance 번호
// privateKeyFile - private key 파일
func (client RootPasswordClient) Get(ctx context.Context, instanceNo string, privateKeyFile io.ReadCloser) (result RootPasswordContentParameter, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RootPasswordClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, instanceNo, privateKeyFile)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.RootPasswordClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "compute.RootPasswordClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.RootPasswordClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client RootPasswordClient) GetPreparer(ctx context.Context, instanceNo string, privateKeyFile io.ReadCloser) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"instanceNo": autorest.Encode("path", instanceNo),
	}

	formDataParameters := map[string]interface{}{
		"private-key-file": privateKeyFile,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/vpc-compute/api/compute/v1/vpc-servers/instances/{instanceNo}/root-password", pathParameters),
		autorest.WithMultiPartFormData(formDataParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client RootPasswordClient) GetSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client RootPasswordClient) GetResponder(resp *http.Response) (result RootPasswordContentParameter, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
