package network

// FINCLOUD_APACHE_NO_VERSION

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// RouteTableSubnetClient is the network Client
type RouteTableSubnetClient struct {
	BaseClient
}

// NewRouteTableSubnetClient creates an instance of the RouteTableSubnetClient client.
func NewRouteTableSubnetClient() RouteTableSubnetClient {
	return NewRouteTableSubnetClientWithBaseURI(DefaultBaseURI)
}

// NewRouteTableSubnetClientWithBaseURI creates an instance of the RouteTableSubnetClient client.
func NewRouteTableSubnetClientWithBaseURI(baseURI string) RouteTableSubnetClient {
	return RouteTableSubnetClient{NewWithBaseURI(baseURI)}
}

// List routeTable 연관 서브넷 정보 리스트
// Parameters:
// routeTableNo - route Table 번호
func (client RouteTableSubnetClient) List(ctx context.Context, routeTableNo string) (result RouteTableSubnetListParameter, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RouteTableSubnetClient.List")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ListPreparer(ctx, routeTableNo)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.RouteTableSubnetClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.RouteTableSubnetClient", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.RouteTableSubnetClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client RouteTableSubnetClient) ListPreparer(ctx context.Context, routeTableNo string) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"routeTableNo": autorest.Encode("query", routeTableNo),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/vpc-network/api/network/v1/route-tables/subnets"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client RouteTableSubnetClient) ListSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client RouteTableSubnetClient) ListResponder(resp *http.Response) (result RouteTableSubnetListParameter, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Update route Table 연관 서브넷 변경
// Parameters:
// routeTableNo - route Table 번호
// parameters - route Table 연관 서브넷 정보
func (client RouteTableSubnetClient) Update(ctx context.Context, routeTableNo string, parameters RouteTableSubnetParameter) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RouteTableSubnetClient.Update")
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
		err = autorest.NewErrorWithError(err, "network.RouteTableSubnetClient", "Update", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "network.RouteTableSubnetClient", "Update", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.RouteTableSubnetClient", "Update", resp, "Failure responding to request")
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client RouteTableSubnetClient) UpdatePreparer(ctx context.Context, routeTableNo string, parameters RouteTableSubnetParameter) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"routeTableNo": autorest.Encode("path", routeTableNo),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json;charset=UTF-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/vpc-network/api/network/v1/route-tables/instances/{routeTableNo}/subnets", pathParameters),
		autorest.WithJSON(parameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client RouteTableSubnetClient) UpdateSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client RouteTableSubnetClient) UpdateResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}
