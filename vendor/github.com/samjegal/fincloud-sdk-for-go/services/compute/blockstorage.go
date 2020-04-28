package compute

// FINCLOUD_APACHE_NO_VERSION

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// BlockStorageClient is the compute Client
type BlockStorageClient struct {
	BaseClient
}

// NewBlockStorageClient creates an instance of the BlockStorageClient client.
func NewBlockStorageClient() BlockStorageClient {
	return NewBlockStorageClientWithBaseURI(DefaultBaseURI)
}

// NewBlockStorageClientWithBaseURI creates an instance of the BlockStorageClient client.
func NewBlockStorageClientWithBaseURI(baseURI string) BlockStorageClient {
	return BlockStorageClient{NewWithBaseURI(baseURI)}
}

// Get server 블럭 스토리지 정보
// Parameters:
// instanceNo - server instance 번호
func (client BlockStorageClient) Get(ctx context.Context, instanceNo string) (result BlockStorageContentParameter, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BlockStorageClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, instanceNo)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.BlockStorageClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "compute.BlockStorageClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.BlockStorageClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client BlockStorageClient) GetPreparer(ctx context.Context, instanceNo string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"instanceNo": autorest.Encode("path", instanceNo),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/vpc-compute/api/compute/v1/vpc-servers/instances/{instanceNo}/block-storages", pathParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client BlockStorageClient) GetSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client BlockStorageClient) GetResponder(resp *http.Response) (result BlockStorageContentParameter, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
