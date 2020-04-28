package network

// FINCLOUD_APACHE_NO_VERSION

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// RouteTableClient is the network Client
type RouteTableClient struct {
	BaseClient
}

// NewRouteTableClient creates an instance of the RouteTableClient client.
func NewRouteTableClient() RouteTableClient {
	return NewRouteTableClientWithBaseURI(DefaultBaseURI)
}

// NewRouteTableClientWithBaseURI creates an instance of the RouteTableClient client.
func NewRouteTableClientWithBaseURI(baseURI string) RouteTableClient {
	return RouteTableClient{NewWithBaseURI(baseURI)}
}

// Create route Table 생성
// Parameters:
// parameters - route Table 생성 데이터
func (client RouteTableClient) Create(ctx context.Context, parameters RouteTableParameter) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RouteTableClient.Create")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreatePreparer(ctx, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.RouteTableClient", "Create", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "network.RouteTableClient", "Create", resp, "Failure sending request")
		return
	}

	result, err = client.CreateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.RouteTableClient", "Create", resp, "Failure responding to request")
	}

	return
}

// CreatePreparer prepares the Create request.
func (client RouteTableClient) CreatePreparer(ctx context.Context, parameters RouteTableParameter) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json;charset=UTF-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/vpc-network/api/network/v1/route-tables/instances"),
		autorest.WithJSON(parameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client RouteTableClient) CreateSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client RouteTableClient) CreateResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Delete route Table 삭제
// Parameters:
// routeTableNo - route Table 번호
func (client RouteTableClient) Delete(ctx context.Context, routeTableNo string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RouteTableClient.Delete")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, routeTableNo)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.RouteTableClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "network.RouteTableClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.RouteTableClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client RouteTableClient) DeletePreparer(ctx context.Context, routeTableNo string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"routeTableNo": autorest.Encode("path", routeTableNo),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/vpc-network/api/network/v1/route-tables/instances/{routeTableNo}", pathParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client RouteTableClient) DeleteSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client RouteTableClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get route Table 룰 정보
func (client RouteTableClient) Get(ctx context.Context) (result RouteTableRuleSearchParameter, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RouteTableClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.RouteTableClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.RouteTableClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.RouteTableClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client RouteTableClient) GetPreparer(ctx context.Context) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/vpc-network/api/network/v1/route-tables/routes"))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client RouteTableClient) GetSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client RouteTableClient) GetResponder(resp *http.Response) (result RouteTableRuleSearchParameter, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetEndpoints route Table 엔드포인트 정보
func (client RouteTableClient) GetEndpoints(ctx context.Context) (result RouteTableEndpointsParameter, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RouteTableClient.GetEndpoints")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetEndpointsPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.RouteTableClient", "GetEndpoints", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetEndpointsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.RouteTableClient", "GetEndpoints", resp, "Failure sending request")
		return
	}

	result, err = client.GetEndpointsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.RouteTableClient", "GetEndpoints", resp, "Failure responding to request")
	}

	return
}

// GetEndpointsPreparer prepares the GetEndpoints request.
func (client RouteTableClient) GetEndpointsPreparer(ctx context.Context) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/vpc-network/api/network/v1/route-tables/endpoints"))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetEndpointsSender sends the GetEndpoints request. The method will close the
// http.Response Body if it receives an error.
func (client RouteTableClient) GetEndpointsSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// GetEndpointsResponder handles the response to the GetEndpoints request. The method always
// closes the http.Response Body.
func (client RouteTableClient) GetEndpointsResponder(resp *http.Response) (result RouteTableEndpointsParameter, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetEndpointTypes route Table 엔드포인트 타입 정보
func (client RouteTableClient) GetEndpointTypes(ctx context.Context) (result RouteTableEndpointTypesParameter, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RouteTableClient.GetEndpointTypes")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetEndpointTypesPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.RouteTableClient", "GetEndpointTypes", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetEndpointTypesSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.RouteTableClient", "GetEndpointTypes", resp, "Failure sending request")
		return
	}

	result, err = client.GetEndpointTypesResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.RouteTableClient", "GetEndpointTypes", resp, "Failure responding to request")
	}

	return
}

// GetEndpointTypesPreparer prepares the GetEndpointTypes request.
func (client RouteTableClient) GetEndpointTypesPreparer(ctx context.Context) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/vpc-network/api/network/v1/route-tables/endpoint-types"))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetEndpointTypesSender sends the GetEndpointTypes request. The method will close the
// http.Response Body if it receives an error.
func (client RouteTableClient) GetEndpointTypesSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// GetEndpointTypesResponder handles the response to the GetEndpointTypes request. The method always
// closes the http.Response Body.
func (client RouteTableClient) GetEndpointTypesResponder(resp *http.Response) (result RouteTableEndpointTypesParameter, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// IsSetted route Table 설정 확인
func (client RouteTableClient) IsSetted(ctx context.Context) (result RouteTableIsSettedParameter, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RouteTableClient.IsSetted")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.IsSettedPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.RouteTableClient", "IsSetted", nil, "Failure preparing request")
		return
	}

	resp, err := client.IsSettedSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.RouteTableClient", "IsSetted", resp, "Failure sending request")
		return
	}

	result, err = client.IsSettedResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.RouteTableClient", "IsSetted", resp, "Failure responding to request")
	}

	return
}

// IsSettedPreparer prepares the IsSetted request.
func (client RouteTableClient) IsSettedPreparer(ctx context.Context) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/vpc-network/api/network/v1/route-tables/is-setted"))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// IsSettedSender sends the IsSetted request. The method will close the
// http.Response Body if it receives an error.
func (client RouteTableClient) IsSettedSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// IsSettedResponder handles the response to the IsSetted request. The method always
// closes the http.Response Body.
func (client RouteTableClient) IsSettedResponder(resp *http.Response) (result RouteTableIsSettedParameter, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List route Table 정보 리스트
// Parameters:
// parameters - route Table 정보 리스트 데이터
func (client RouteTableClient) List(ctx context.Context, parameters RouteTableSearchParameter) (result RouteTableSearchListParameter, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RouteTableClient.List")
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
		err = autorest.NewErrorWithError(err, "network.RouteTableClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.RouteTableClient", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.RouteTableClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client RouteTableClient) ListPreparer(ctx context.Context, parameters RouteTableSearchParameter) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json;charset=UTF-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/vpc-network/api/network/v1/route-tables/search"),
		autorest.WithJSON(parameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client RouteTableClient) ListSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client RouteTableClient) ListResponder(resp *http.Response) (result RouteTableSearchListParameter, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Update route Table 업데이트
// Parameters:
// routeTableNo - route Table 번호
// parameters - route Table 업데이트
func (client RouteTableClient) Update(ctx context.Context, routeTableNo string, parameters RouteTableRuleParameter) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RouteTableClient.Update")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.UpdatePreparer(ctx, routeTableNo, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.RouteTableClient", "Update", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "network.RouteTableClient", "Update", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.RouteTableClient", "Update", resp, "Failure responding to request")
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client RouteTableClient) UpdatePreparer(ctx context.Context, routeTableNo string, parameters RouteTableRuleParameter) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"routeTableNo": autorest.Encode("path", routeTableNo),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json;charset=UTF-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/vpc-network/api/network/v1/route-tables/instances/{routeTableNo}", pathParameters),
		autorest.WithJSON(parameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client RouteTableClient) UpdateSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client RouteTableClient) UpdateResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}
