package compute

// FINCLOUD_APACHE_NO_VERSION

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// OutboundRuleClient is the compute Client
type OutboundRuleClient struct {
	BaseClient
}

// NewOutboundRuleClient creates an instance of the OutboundRuleClient client.
func NewOutboundRuleClient() OutboundRuleClient {
	return NewOutboundRuleClientWithBaseURI(DefaultBaseURI)
}

// NewOutboundRuleClientWithBaseURI creates an instance of the OutboundRuleClient client.
func NewOutboundRuleClientWithBaseURI(baseURI string) OutboundRuleClient {
	return OutboundRuleClient{NewWithBaseURI(baseURI)}
}

// Get ACG outbound rule 정보
// Parameters:
// networkAcgNo - ACG 번호
func (client OutboundRuleClient) Get(ctx context.Context, networkAcgNo string) (result SecurityGroupRuleContentParameter, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/OutboundRuleClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, networkAcgNo)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.OutboundRuleClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "compute.OutboundRuleClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.OutboundRuleClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client OutboundRuleClient) GetPreparer(ctx context.Context, networkAcgNo string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"networkAcgNo": autorest.Encode("path", networkAcgNo),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/vpc-compute/api/compute/v1/access-control-groups/edit/instances/{networkAcgNo}/outbound-rules", pathParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client OutboundRuleClient) GetSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client OutboundRuleClient) GetResponder(resp *http.Response) (result SecurityGroupRuleContentParameter, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
