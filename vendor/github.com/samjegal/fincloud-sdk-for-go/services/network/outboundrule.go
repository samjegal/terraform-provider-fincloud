package network

// FINCLOUD_APACHE_NO_VERSION

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// OutboundRuleClient is the network Client
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

// Get ACL outbound rule 정보
// Parameters:
// networkACLNo - ACL 번호
func (client OutboundRuleClient) Get(ctx context.Context, networkACLNo string) (result ACLRuleContentParameter, err error) {
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
	req, err := client.GetPreparer(ctx, networkACLNo)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.OutboundRuleClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.OutboundRuleClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.OutboundRuleClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client OutboundRuleClient) GetPreparer(ctx context.Context, networkACLNo string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"networkAclNo": autorest.Encode("path", networkACLNo),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/vpc-network/api/network/v1/network-acls/{networkAclNo}/outbound-rules", pathParameters))
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
func (client OutboundRuleClient) GetResponder(resp *http.Response) (result ACLRuleContentParameter, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
