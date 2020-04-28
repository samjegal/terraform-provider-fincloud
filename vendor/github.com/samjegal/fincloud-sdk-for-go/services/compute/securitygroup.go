package compute

// FINCLOUD_APACHE_NO_VERSION

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// SecurityGroupClient is the compute Client
type SecurityGroupClient struct {
	BaseClient
}

// NewSecurityGroupClient creates an instance of the SecurityGroupClient client.
func NewSecurityGroupClient() SecurityGroupClient {
	return NewSecurityGroupClientWithBaseURI(DefaultBaseURI)
}

// NewSecurityGroupClientWithBaseURI creates an instance of the SecurityGroupClient client.
func NewSecurityGroupClientWithBaseURI(baseURI string) SecurityGroupClient {
	return SecurityGroupClient{NewWithBaseURI(baseURI)}
}

// CreateOrUpdate ACG 생성 및 업데이트
// Parameters:
// parameters - ACG 생성 및 업데이트 정보
func (client SecurityGroupClient) CreateOrUpdate(ctx context.Context, parameters SecurityGroupNameParameter) (result SecurityGroupListParameter, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SecurityGroupClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreateOrUpdatePreparer(ctx, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.SecurityGroupClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "compute.SecurityGroupClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.SecurityGroupClient", "CreateOrUpdate", resp, "Failure responding to request")
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client SecurityGroupClient) CreateOrUpdatePreparer(ctx context.Context, parameters SecurityGroupNameParameter) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json;charset=UTF-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/vpc-compute/api/compute/v1/access-control-groups/instances"),
		autorest.WithJSON(parameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client SecurityGroupClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client SecurityGroupClient) CreateOrUpdateResponder(resp *http.Response) (result SecurityGroupListParameter, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete ACG 삭제
// Parameters:
// parameters - ACG 기본 정보
func (client SecurityGroupClient) Delete(ctx context.Context, parameters SecurityGroupNumberParameter) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SecurityGroupClient.Delete")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.SecurityGroupClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "compute.SecurityGroupClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.SecurityGroupClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client SecurityGroupClient) DeletePreparer(ctx context.Context, parameters SecurityGroupNumberParameter) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json;charset=UTF-8"),
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/vpc-compute/api/compute/v1/access-control-groups/instances"),
		autorest.WithJSON(parameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client SecurityGroupClient) DeleteSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client SecurityGroupClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// List ACG 정보 리스트
func (client SecurityGroupClient) List(ctx context.Context) (result SecurityGroupListParameter, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SecurityGroupClient.List")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ListPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.SecurityGroupClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "compute.SecurityGroupClient", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.SecurityGroupClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client SecurityGroupClient) ListPreparer(ctx context.Context) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/vpc-compute/api/compute/v1/access-control-groups/search"))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client SecurityGroupClient) ListSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client SecurityGroupClient) ListResponder(resp *http.Response) (result SecurityGroupListParameter, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByPrefixName ACG 이름 리스트 정보
// Parameters:
// parameters - ACG prefix 이름 정보
func (client SecurityGroupClient) ListByPrefixName(ctx context.Context, parameters SecurityGroupNameParameter) (result SecurityGroupContentListParameter, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SecurityGroupClient.ListByPrefixName")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ListByPrefixNamePreparer(ctx, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.SecurityGroupClient", "ListByPrefixName", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByPrefixNameSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "compute.SecurityGroupClient", "ListByPrefixName", resp, "Failure sending request")
		return
	}

	result, err = client.ListByPrefixNameResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.SecurityGroupClient", "ListByPrefixName", resp, "Failure responding to request")
	}

	return
}

// ListByPrefixNamePreparer prepares the ListByPrefixName request.
func (client SecurityGroupClient) ListByPrefixNamePreparer(ctx context.Context, parameters SecurityGroupNameParameter) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json;charset=UTF-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/vpc-compute/api/compute/v1/access-control-groups/search-by-prefix-name"),
		autorest.WithJSON(parameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByPrefixNameSender sends the ListByPrefixName request. The method will close the
// http.Response Body if it receives an error.
func (client SecurityGroupClient) ListByPrefixNameSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// ListByPrefixNameResponder handles the response to the ListByPrefixName request. The method always
// closes the http.Response Body.
func (client SecurityGroupClient) ListByPrefixNameResponder(resp *http.Response) (result SecurityGroupContentListParameter, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
