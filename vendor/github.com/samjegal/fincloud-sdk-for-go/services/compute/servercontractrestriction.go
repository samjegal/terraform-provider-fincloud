package compute

// FINCLOUD_APACHE_NO_VERSION

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// ServerContractRestrictionClient is the compute Client
type ServerContractRestrictionClient struct {
	BaseClient
}

// NewServerContractRestrictionClient creates an instance of the ServerContractRestrictionClient client.
func NewServerContractRestrictionClient() ServerContractRestrictionClient {
	return NewServerContractRestrictionClientWithBaseURI(DefaultBaseURI)
}

// NewServerContractRestrictionClientWithBaseURI creates an instance of the ServerContractRestrictionClient client.
func NewServerContractRestrictionClientWithBaseURI(baseURI string) ServerContractRestrictionClient {
	return ServerContractRestrictionClient{NewWithBaseURI(baseURI)}
}

// Get 서버 사용 계약 제한
// Parameters:
// productType2Code - server product 타입 2 코드
// singleProductContractTypeCode - 단일 product 계약 타입 코드
func (client ServerContractRestrictionClient) Get(ctx context.Context, productType2Code string, singleProductContractTypeCode string) (result ServerContentContractRestrictionParameter, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ServerContractRestrictionClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, productType2Code, singleProductContractTypeCode)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.ServerContractRestrictionClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "compute.ServerContractRestrictionClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.ServerContractRestrictionClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client ServerContractRestrictionClient) GetPreparer(ctx context.Context, productType2Code string, singleProductContractTypeCode string) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"productType2Code":              autorest.Encode("query", productType2Code),
		"singleProductContractTypeCode": autorest.Encode("query", singleProductContractTypeCode),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/vpc-compute/api/compute/v1/vpc-servers/server-contract-restriction"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ServerContractRestrictionClient) GetSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ServerContractRestrictionClient) GetResponder(resp *http.Response) (result ServerContentContractRestrictionParameter, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
