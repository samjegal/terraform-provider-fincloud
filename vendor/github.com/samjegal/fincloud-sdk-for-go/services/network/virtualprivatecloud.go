package network

// FINCLOUD_APACHE_NO_VERSION

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// VirtualPrivateCloudClient is the network Client
type VirtualPrivateCloudClient struct {
	BaseClient
}

// NewVirtualPrivateCloudClient creates an instance of the VirtualPrivateCloudClient client.
func NewVirtualPrivateCloudClient() VirtualPrivateCloudClient {
	return NewVirtualPrivateCloudClientWithBaseURI(DefaultBaseURI)
}

// NewVirtualPrivateCloudClientWithBaseURI creates an instance of the VirtualPrivateCloudClient client.
func NewVirtualPrivateCloudClientWithBaseURI(baseURI string) VirtualPrivateCloudClient {
	return VirtualPrivateCloudClient{NewWithBaseURI(baseURI)}
}

// CreateOrUpdate VPC 생성
// Parameters:
// parameters - VPC 생성 정보
func (client VirtualPrivateCloudClient) CreateOrUpdate(ctx context.Context, parameters VirtualPrivateCloudParameter) (result VirtualPrivateCloudListParameter, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VirtualPrivateCloudClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.VpcName", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("network.VirtualPrivateCloudClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VirtualPrivateCloudClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.VirtualPrivateCloudClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VirtualPrivateCloudClient", "CreateOrUpdate", resp, "Failure responding to request")
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client VirtualPrivateCloudClient) CreateOrUpdatePreparer(ctx context.Context, parameters VirtualPrivateCloudParameter) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json;charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/vpc-network/api/network/v1/vpcs/instances"),
		autorest.WithJSON(parameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualPrivateCloudClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client VirtualPrivateCloudClient) CreateOrUpdateResponder(resp *http.Response) (result VirtualPrivateCloudListParameter, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete VPC 삭제
// Parameters:
// vpcNo - 삭제할 VPC 번호
// parameters - VPC 생성 정보
func (client VirtualPrivateCloudClient) Delete(ctx context.Context, vpcNo string, parameters VirtualPrivateCloudParameter) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VirtualPrivateCloudClient.Delete")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.VpcName", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("network.VirtualPrivateCloudClient", "Delete", err.Error())
	}

	req, err := client.DeletePreparer(ctx, vpcNo, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VirtualPrivateCloudClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "network.VirtualPrivateCloudClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VirtualPrivateCloudClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client VirtualPrivateCloudClient) DeletePreparer(ctx context.Context, vpcNo string, parameters VirtualPrivateCloudParameter) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"vpcNo": autorest.Encode("path", vpcNo),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json;charset=utf-8"),
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/vpc-network/api/network/v1/vpcs/instances/{vpcNo}", pathParameters),
		autorest.WithJSON(parameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualPrivateCloudClient) DeleteSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client VirtualPrivateCloudClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// List VPC 정보 리스트
func (client VirtualPrivateCloudClient) List(ctx context.Context) (result VirtualPrivateCloudListParameter, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VirtualPrivateCloudClient.List")
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
		err = autorest.NewErrorWithError(err, "network.VirtualPrivateCloudClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.VirtualPrivateCloudClient", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VirtualPrivateCloudClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client VirtualPrivateCloudClient) ListPreparer(ctx context.Context) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/vpc-network/api/network/v1/vpcs/instances"))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualPrivateCloudClient) ListSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client VirtualPrivateCloudClient) ListResponder(resp *http.Response) (result VirtualPrivateCloudListParameter, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
