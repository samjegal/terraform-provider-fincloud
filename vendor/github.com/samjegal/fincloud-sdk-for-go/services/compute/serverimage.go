package compute

// FINCLOUD_APACHE_NO_VERSION

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// ServerImageClient is the compute Client
type ServerImageClient struct {
	BaseClient
}

// NewServerImageClient creates an instance of the ServerImageClient client.
func NewServerImageClient() ServerImageClient {
	return NewServerImageClientWithBaseURI(DefaultBaseURI)
}

// NewServerImageClientWithBaseURI creates an instance of the ServerImageClient client.
func NewServerImageClientWithBaseURI(baseURI string) ServerImageClient {
	return ServerImageClient{NewWithBaseURI(baseURI)}
}

// Create server 이미지 생성
// Parameters:
// instanceNo - server instance 번호
func (client ServerImageClient) Create(ctx context.Context, instanceNo string, infraResourceTypeCode string, instanceName string, blockStorageUsageTypeCode string, parameter ServerImageParameter) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ServerImageClient.Create")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreatePreparer(ctx, instanceNo, infraResourceTypeCode, instanceName, blockStorageUsageTypeCode, parameter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.ServerImageClient", "Create", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "compute.ServerImageClient", "Create", resp, "Failure sending request")
		return
	}

	result, err = client.CreateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.ServerImageClient", "Create", resp, "Failure responding to request")
	}

	return
}

// CreatePreparer prepares the Create request.
func (client ServerImageClient) CreatePreparer(ctx context.Context, instanceNo string, infraResourceTypeCode string, instanceName string, blockStorageUsageTypeCode string, parameter ServerImageParameter) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"instanceNo": autorest.Encode("path", instanceNo),
	}

	queryParameters := map[string]interface{}{
		"blockStorageUsageTypeCode": autorest.Encode("query", blockStorageUsageTypeCode),
		"infraResourceTypeCode":     autorest.Encode("query", infraResourceTypeCode),
		"instanceName":              autorest.Encode("query", instanceName),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json;charset=UTF-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/vpc-compute/api/compute/v1/vpc-servers/instances/{instanceNo}/server-images", pathParameters),
		autorest.WithJSON(parameter),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client ServerImageClient) CreateSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client ServerImageClient) CreateResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByClosing())
	result.Response = resp
	return
}
