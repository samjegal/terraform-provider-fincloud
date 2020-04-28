package compute

// FINCLOUD_APACHE_NO_VERSION

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// NetworkInterfaceClient is the compute Client
type NetworkInterfaceClient struct {
	BaseClient
}

// NewNetworkInterfaceClient creates an instance of the NetworkInterfaceClient client.
func NewNetworkInterfaceClient() NetworkInterfaceClient {
	return NewNetworkInterfaceClientWithBaseURI(DefaultBaseURI)
}

// NewNetworkInterfaceClientWithBaseURI creates an instance of the NetworkInterfaceClient client.
func NewNetworkInterfaceClientWithBaseURI(baseURI string) NetworkInterfaceClient {
	return NetworkInterfaceClient{NewWithBaseURI(baseURI)}
}

// Attach 서버 네트워크 인터페이스 할당
// Parameters:
// instanceNo - server instance 번호
// networkInterfaceNo - network interface 번호
func (client NetworkInterfaceClient) Attach(ctx context.Context, instanceNo string, networkInterfaceNo string, parameter NetworkInterfaceSubnetParameter) (result NetworkInterfaceAttachableContentParameter, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/NetworkInterfaceClient.Attach")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.AttachPreparer(ctx, instanceNo, networkInterfaceNo, parameter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.NetworkInterfaceClient", "Attach", nil, "Failure preparing request")
		return
	}

	resp, err := client.AttachSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "compute.NetworkInterfaceClient", "Attach", resp, "Failure sending request")
		return
	}

	result, err = client.AttachResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.NetworkInterfaceClient", "Attach", resp, "Failure responding to request")
	}

	return
}

// AttachPreparer prepares the Attach request.
func (client NetworkInterfaceClient) AttachPreparer(ctx context.Context, instanceNo string, networkInterfaceNo string, parameter NetworkInterfaceSubnetParameter) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"instanceNo":         autorest.Encode("path", instanceNo),
		"networkInterfaceNo": autorest.Encode("path", networkInterfaceNo),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json;charset=UTF-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/vpc-compute/api/compute/v1/vpc-servers/instances/{instanceNo}/network-interfaces/{networkInterfaceNo}/attach", pathParameters),
		autorest.WithJSON(parameter))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// AttachSender sends the Attach request. The method will close the
// http.Response Body if it receives an error.
func (client NetworkInterfaceClient) AttachSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// AttachResponder handles the response to the Attach request. The method always
// closes the http.Response Body.
func (client NetworkInterfaceClient) AttachResponder(resp *http.Response) (result NetworkInterfaceAttachableContentParameter, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// AttachableList server network interface 할당 가능한 정보 리스트
// Parameters:
// subnetNo - server network interface 할당 가능 정보
func (client NetworkInterfaceClient) AttachableList(ctx context.Context, subnetNo int32) (result NetworkInterfaceAttachableListParameter, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/NetworkInterfaceClient.AttachableList")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.AttachableListPreparer(ctx, subnetNo)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.NetworkInterfaceClient", "AttachableList", nil, "Failure preparing request")
		return
	}

	resp, err := client.AttachableListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "compute.NetworkInterfaceClient", "AttachableList", resp, "Failure sending request")
		return
	}

	result, err = client.AttachableListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.NetworkInterfaceClient", "AttachableList", resp, "Failure responding to request")
	}

	return
}

// AttachableListPreparer prepares the AttachableList request.
func (client NetworkInterfaceClient) AttachableListPreparer(ctx context.Context, subnetNo int32) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"subnetNo": autorest.Encode("query", subnetNo),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/vpc-compute/api/compute/v1/network-interfaces/attachable"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// AttachableListSender sends the AttachableList request. The method will close the
// http.Response Body if it receives an error.
func (client NetworkInterfaceClient) AttachableListSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// AttachableListResponder handles the response to the AttachableList request. The method always
// closes the http.Response Body.
func (client NetworkInterfaceClient) AttachableListResponder(resp *http.Response) (result NetworkInterfaceAttachableListParameter, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Create network interface 생성
// Parameters:
// parameters - network interface 생성 정보
func (client NetworkInterfaceClient) Create(ctx context.Context, parameters NetworkInterfaceParameter) (result ErrorMessageParameter, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/NetworkInterfaceClient.Create")
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
			Constraints: []validation.Constraint{{Target: "parameters.AccessControlGroups", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "parameters.AccessControlGroupNoList", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "parameters.NetworkInterfaceName", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "parameters.SubnetNo", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "parameters.VpcNo", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("compute.NetworkInterfaceClient", "Create", err.Error())
	}

	req, err := client.CreatePreparer(ctx, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.NetworkInterfaceClient", "Create", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "compute.NetworkInterfaceClient", "Create", resp, "Failure sending request")
		return
	}

	result, err = client.CreateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.NetworkInterfaceClient", "Create", resp, "Failure responding to request")
	}

	return
}

// CreatePreparer prepares the Create request.
func (client NetworkInterfaceClient) CreatePreparer(ctx context.Context, parameters NetworkInterfaceParameter) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json;charset=UTF-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/vpc-compute/api/compute/v1/network-interfaces/instances"),
		autorest.WithJSON(parameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client NetworkInterfaceClient) CreateSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client NetworkInterfaceClient) CreateResponder(resp *http.Response) (result ErrorMessageParameter, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated, http.StatusBadRequest),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete network Interface 삭제
// Parameters:
// networkInterfaceNo - 삭제할 network interface 번호
func (client NetworkInterfaceClient) Delete(ctx context.Context, networkInterfaceNo string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/NetworkInterfaceClient.Delete")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, networkInterfaceNo)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.NetworkInterfaceClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "compute.NetworkInterfaceClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.NetworkInterfaceClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client NetworkInterfaceClient) DeletePreparer(ctx context.Context, networkInterfaceNo string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"networkInterfaceNo": autorest.Encode("path", networkInterfaceNo),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/vpc-compute/api/compute/v1/network-interfaces/instances/{networkInterfaceNo}", pathParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client NetworkInterfaceClient) DeleteSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client NetworkInterfaceClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Detach 서버 네트워크 인터페이스 해제
// Parameters:
// instanceNo - server instance 번호
// networkInterfaceNo - network interface 번호
func (client NetworkInterfaceClient) Detach(ctx context.Context, instanceNo string, networkInterfaceNo string, parameter NetworkInterfaceSubnetParameter) (result NetworkInterfaceContentParameter, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/NetworkInterfaceClient.Detach")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DetachPreparer(ctx, instanceNo, networkInterfaceNo, parameter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.NetworkInterfaceClient", "Detach", nil, "Failure preparing request")
		return
	}

	resp, err := client.DetachSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "compute.NetworkInterfaceClient", "Detach", resp, "Failure sending request")
		return
	}

	result, err = client.DetachResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.NetworkInterfaceClient", "Detach", resp, "Failure responding to request")
	}

	return
}

// DetachPreparer prepares the Detach request.
func (client NetworkInterfaceClient) DetachPreparer(ctx context.Context, instanceNo string, networkInterfaceNo string, parameter NetworkInterfaceSubnetParameter) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"instanceNo":         autorest.Encode("path", instanceNo),
		"networkInterfaceNo": autorest.Encode("path", networkInterfaceNo),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json;charset=UTF-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/vpc-compute/api/compute/v1/vpc-servers/instances/{instanceNo}/network-interfaces/{networkInterfaceNo}/detach", pathParameters),
		autorest.WithJSON(parameter))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DetachSender sends the Detach request. The method will close the
// http.Response Body if it receives an error.
func (client NetworkInterfaceClient) DetachSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// DetachResponder handles the response to the Detach request. The method always
// closes the http.Response Body.
func (client NetworkInterfaceClient) DetachResponder(resp *http.Response) (result NetworkInterfaceContentParameter, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List network interface 정보 리스트
// Parameters:
// parameters - network interface 검색 조건부 데이터
func (client NetworkInterfaceClient) List(ctx context.Context, parameters NetworkInterfaceSearchParameter) (result NetworkInterfaceListParameter, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/NetworkInterfaceClient.List")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ListPreparer(ctx, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.NetworkInterfaceClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "compute.NetworkInterfaceClient", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.NetworkInterfaceClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client NetworkInterfaceClient) ListPreparer(ctx context.Context, parameters NetworkInterfaceSearchParameter) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json;charset=UTF-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/vpc-compute/api/compute/v1/network-interfaces/search"),
		autorest.WithJSON(parameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client NetworkInterfaceClient) ListSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client NetworkInterfaceClient) ListResponder(resp *http.Response) (result NetworkInterfaceListParameter, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ValidNetworkInterfaceIP network Interface IP 할당 가능 여부
// Parameters:
// overlayIP - netowrk interface에 할당할 IP 주소
// subnetNo - subnet 번호
func (client NetworkInterfaceClient) ValidNetworkInterfaceIP(ctx context.Context, overlayIP string, subnetNo string) (result NetworkInterfaceValidationParameter, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/NetworkInterfaceClient.ValidNetworkInterfaceIP")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ValidNetworkInterfaceIPPreparer(ctx, overlayIP, subnetNo)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.NetworkInterfaceClient", "ValidNetworkInterfaceIP", nil, "Failure preparing request")
		return
	}

	resp, err := client.ValidNetworkInterfaceIPSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "compute.NetworkInterfaceClient", "ValidNetworkInterfaceIP", resp, "Failure sending request")
		return
	}

	result, err = client.ValidNetworkInterfaceIPResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.NetworkInterfaceClient", "ValidNetworkInterfaceIP", resp, "Failure responding to request")
	}

	return
}

// ValidNetworkInterfaceIPPreparer prepares the ValidNetworkInterfaceIP request.
func (client NetworkInterfaceClient) ValidNetworkInterfaceIPPreparer(ctx context.Context, overlayIP string, subnetNo string) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"overlayIp": autorest.Encode("query", overlayIP),
		"subnetNo":  autorest.Encode("query", subnetNo),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/vpc-compute/api/compute/v1/network-interfaces/valid-network-interface-ip"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ValidNetworkInterfaceIPSender sends the ValidNetworkInterfaceIP request. The method will close the
// http.Response Body if it receives an error.
func (client NetworkInterfaceClient) ValidNetworkInterfaceIPSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// ValidNetworkInterfaceIPResponder handles the response to the ValidNetworkInterfaceIP request. The method always
// closes the http.Response Body.
func (client NetworkInterfaceClient) ValidNetworkInterfaceIPResponder(resp *http.Response) (result NetworkInterfaceValidationParameter, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
