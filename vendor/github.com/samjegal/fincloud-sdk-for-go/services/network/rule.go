package network

// FINCLOUD_APACHE_NO_VERSION

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// RuleClient is the network Client
type RuleClient struct {
	BaseClient
}

// NewRuleClient creates an instance of the RuleClient client.
func NewRuleClient() RuleClient {
	return NewRuleClientWithBaseURI(DefaultBaseURI)
}

// NewRuleClientWithBaseURI creates an instance of the RuleClient client.
func NewRuleClientWithBaseURI(baseURI string) RuleClient {
	return RuleClient{NewWithBaseURI(baseURI)}
}

// CreateOrUpdate ACL rule 생성 및 업데이트
// Parameters:
// networkACLNo - ACL 번호
// parameters - ACL rule 파라미터
func (client RuleClient) CreateOrUpdate(ctx context.Context, networkACLNo string, parameters ACLRuleListParameter) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RuleClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreateOrUpdatePreparer(ctx, networkACLNo, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.RuleClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "network.RuleClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.RuleClient", "CreateOrUpdate", resp, "Failure responding to request")
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client RuleClient) CreateOrUpdatePreparer(ctx context.Context, networkACLNo string, parameters ACLRuleListParameter) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"networkAclNo": autorest.Encode("path", networkACLNo),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json;charset=UTF-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/vpc-network/api/network/v1/network-acls/{networkAclNo}/rules", pathParameters),
		autorest.WithJSON(parameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client RuleClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client RuleClient) CreateOrUpdateResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}
