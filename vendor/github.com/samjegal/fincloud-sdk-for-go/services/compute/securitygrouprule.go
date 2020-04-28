package compute

// FINCLOUD_APACHE_NO_VERSION

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// SecurityGroupRuleClient is the compute Client
type SecurityGroupRuleClient struct {
	BaseClient
}

// NewSecurityGroupRuleClient creates an instance of the SecurityGroupRuleClient client.
func NewSecurityGroupRuleClient() SecurityGroupRuleClient {
	return NewSecurityGroupRuleClientWithBaseURI(DefaultBaseURI)
}

// NewSecurityGroupRuleClientWithBaseURI creates an instance of the SecurityGroupRuleClient client.
func NewSecurityGroupRuleClientWithBaseURI(baseURI string) SecurityGroupRuleClient {
	return SecurityGroupRuleClient{NewWithBaseURI(baseURI)}
}

// CreateOrUpdate ACG rule 생성 및 업데이트
// Parameters:
// networkAcgNo - ACG 번호
// parameters - ACG rule 파라미터
func (client SecurityGroupRuleClient) CreateOrUpdate(ctx context.Context, networkAcgNo string, parameters SecurityGroupRulesProperties) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SecurityGroupRuleClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreateOrUpdatePreparer(ctx, networkAcgNo, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.SecurityGroupRuleClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "compute.SecurityGroupRuleClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.SecurityGroupRuleClient", "CreateOrUpdate", resp, "Failure responding to request")
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client SecurityGroupRuleClient) CreateOrUpdatePreparer(ctx context.Context, networkAcgNo string, parameters SecurityGroupRulesProperties) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"networkAcgNo": autorest.Encode("path", networkAcgNo),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json;charset=UTF-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/vpc-compute/api/compute/v1/access-control-groups/instances/{networkAcgNo}/rules", pathParameters),
		autorest.WithJSON(parameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client SecurityGroupRuleClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client SecurityGroupRuleClient) CreateOrUpdateResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}
