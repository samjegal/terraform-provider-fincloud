package compute

// FINCLOUD_APACHE_NO_VERSION

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// PublicIPAddressClient is the compute Client
type PublicIPAddressClient struct {
	BaseClient
}

// NewPublicIPAddressClient creates an instance of the PublicIPAddressClient client.
func NewPublicIPAddressClient() PublicIPAddressClient {
	return NewPublicIPAddressClientWithBaseURI(DefaultBaseURI)
}

// NewPublicIPAddressClientWithBaseURI creates an instance of the PublicIPAddressClient client.
func NewPublicIPAddressClientWithBaseURI(baseURI string) PublicIPAddressClient {
	return PublicIPAddressClient{NewWithBaseURI(baseURI)}
}

// Assign 공인 IP 주소 서버 할당
// Parameters:
// instanceNo - 공인 IP 인스턴스 번호
func (client PublicIPAddressClient) Assign(ctx context.Context, instanceNo string, parameters PublicIPAddressServerInstanceParameter) (result PublicIPAddressSearchParameter, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PublicIPAddressClient.Assign")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.AssignPreparer(ctx, instanceNo, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.PublicIPAddressClient", "Assign", nil, "Failure preparing request")
		return
	}

	resp, err := client.AssignSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "compute.PublicIPAddressClient", "Assign", resp, "Failure sending request")
		return
	}

	result, err = client.AssignResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.PublicIPAddressClient", "Assign", resp, "Failure responding to request")
	}

	return
}

// AssignPreparer prepares the Assign request.
func (client PublicIPAddressClient) AssignPreparer(ctx context.Context, instanceNo string, parameters PublicIPAddressServerInstanceParameter) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"instanceNo": autorest.Encode("path", instanceNo),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json;charset=UTF-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/vpc-compute/api/compute/v1/public-ips/instances/{instanceNo}/associate", pathParameters),
		autorest.WithJSON(parameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// AssignSender sends the Assign request. The method will close the
// http.Response Body if it receives an error.
func (client PublicIPAddressClient) AssignSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// AssignResponder handles the response to the Assign request. The method always
// closes the http.Response Body.
func (client PublicIPAddressClient) AssignResponder(resp *http.Response) (result PublicIPAddressSearchParameter, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Create 공인 IP 주소 서버 생성
func (client PublicIPAddressClient) Create(ctx context.Context, parameters PublicIPAddressServerInstanceParameter) (result PublicIPAddressSearchParameter, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PublicIPAddressClient.Create")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreatePreparer(ctx, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.PublicIPAddressClient", "Create", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "compute.PublicIPAddressClient", "Create", resp, "Failure sending request")
		return
	}

	result, err = client.CreateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.PublicIPAddressClient", "Create", resp, "Failure responding to request")
	}

	return
}

// CreatePreparer prepares the Create request.
func (client PublicIPAddressClient) CreatePreparer(ctx context.Context, parameters PublicIPAddressServerInstanceParameter) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json;charset=UTF-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/vpc-compute/api/compute/v1/public-ips/instances"),
		autorest.WithJSON(parameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client PublicIPAddressClient) CreateSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client PublicIPAddressClient) CreateResponder(resp *http.Response) (result PublicIPAddressSearchParameter, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete 공인 IP 주소 반납
// Parameters:
// instanceNo - 공인 IP 인스턴스 번호
func (client PublicIPAddressClient) Delete(ctx context.Context, instanceNo string, parameters PublicIPAddressServerInstanceParameter) (result PublicIPAddressSearchParameter, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PublicIPAddressClient.Delete")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, instanceNo, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.PublicIPAddressClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "compute.PublicIPAddressClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.PublicIPAddressClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client PublicIPAddressClient) DeletePreparer(ctx context.Context, instanceNo string, parameters PublicIPAddressServerInstanceParameter) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"instanceNo": autorest.Encode("path", instanceNo),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json;charset=UTF-8"),
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/vpc-compute/api/compute/v1/public-ips/instances/{instanceNo}", pathParameters),
		autorest.WithJSON(parameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client PublicIPAddressClient) DeleteSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client PublicIPAddressClient) DeleteResponder(resp *http.Response) (result PublicIPAddressSearchParameter, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List 공인 IP 주소 정보 리스트
// Parameters:
// parameters - 공인 IP 주소 검색 조건부 데이터
func (client PublicIPAddressClient) List(ctx context.Context, parameters PublicIPAddressSearchFilterParameter) (result PublicIPAddressSearchParameter, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PublicIPAddressClient.List")
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
		err = autorest.NewErrorWithError(err, "compute.PublicIPAddressClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "compute.PublicIPAddressClient", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.PublicIPAddressClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client PublicIPAddressClient) ListPreparer(ctx context.Context, parameters PublicIPAddressSearchFilterParameter) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json;charset=UTF-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/vpc-compute/api/compute/v1/public-ips/search"),
		autorest.WithJSON(parameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client PublicIPAddressClient) ListSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client PublicIPAddressClient) ListResponder(resp *http.Response) (result PublicIPAddressSearchParameter, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Remove 공인 IP 주소 서버 해제
// Parameters:
// instanceNo - 공인 IP 인스턴스 번호
func (client PublicIPAddressClient) Remove(ctx context.Context, instanceNo string, parameters PublicIPAddressServerInstanceParameter) (result PublicIPAddressSearchParameter, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PublicIPAddressClient.Remove")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.RemovePreparer(ctx, instanceNo, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.PublicIPAddressClient", "Remove", nil, "Failure preparing request")
		return
	}

	resp, err := client.RemoveSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "compute.PublicIPAddressClient", "Remove", resp, "Failure sending request")
		return
	}

	result, err = client.RemoveResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.PublicIPAddressClient", "Remove", resp, "Failure responding to request")
	}

	return
}

// RemovePreparer prepares the Remove request.
func (client PublicIPAddressClient) RemovePreparer(ctx context.Context, instanceNo string, parameters PublicIPAddressServerInstanceParameter) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"instanceNo": autorest.Encode("path", instanceNo),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json;charset=UTF-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/vpc-compute/api/compute/v1/public-ips/instances/{instanceNo}/dis-associate", pathParameters),
		autorest.WithJSON(parameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// RemoveSender sends the Remove request. The method will close the
// http.Response Body if it receives an error.
func (client PublicIPAddressClient) RemoveSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// RemoveResponder handles the response to the Remove request. The method always
// closes the http.Response Body.
func (client PublicIPAddressClient) RemoveResponder(resp *http.Response) (result PublicIPAddressSearchParameter, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ServerList 공인 IP 주소 할당 가능 리스트
func (client PublicIPAddressClient) ServerList(ctx context.Context) (result PublicIPAddressServerListParameter, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PublicIPAddressClient.ServerList")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ServerListPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.PublicIPAddressClient", "ServerList", nil, "Failure preparing request")
		return
	}

	resp, err := client.ServerListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "compute.PublicIPAddressClient", "ServerList", resp, "Failure sending request")
		return
	}

	result, err = client.ServerListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.PublicIPAddressClient", "ServerList", resp, "Failure responding to request")
	}

	return
}

// ServerListPreparer prepares the ServerList request.
func (client PublicIPAddressClient) ServerListPreparer(ctx context.Context) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/vpc-compute/api/compute/v1/public-ips/servers"))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ServerListSender sends the ServerList request. The method will close the
// http.Response Body if it receives an error.
func (client PublicIPAddressClient) ServerListSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// ServerListResponder handles the response to the ServerList request. The method always
// closes the http.Response Body.
func (client PublicIPAddressClient) ServerListResponder(resp *http.Response) (result PublicIPAddressServerListParameter, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Summary 공인 IP 주소 요약 정보
func (client PublicIPAddressClient) Summary(ctx context.Context) (result PublicIPAddressSummaryParameter, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PublicIPAddressClient.Summary")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.SummaryPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.PublicIPAddressClient", "Summary", nil, "Failure preparing request")
		return
	}

	resp, err := client.SummarySender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "compute.PublicIPAddressClient", "Summary", resp, "Failure sending request")
		return
	}

	result, err = client.SummaryResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.PublicIPAddressClient", "Summary", resp, "Failure responding to request")
	}

	return
}

// SummaryPreparer prepares the Summary request.
func (client PublicIPAddressClient) SummaryPreparer(ctx context.Context) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/vpc-compute/api/compute/v1/public-ips/summary"))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// SummarySender sends the Summary request. The method will close the
// http.Response Body if it receives an error.
func (client PublicIPAddressClient) SummarySender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// SummaryResponder handles the response to the Summary request. The method always
// closes the http.Response Body.
func (client PublicIPAddressClient) SummaryResponder(resp *http.Response) (result PublicIPAddressSummaryParameter, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
