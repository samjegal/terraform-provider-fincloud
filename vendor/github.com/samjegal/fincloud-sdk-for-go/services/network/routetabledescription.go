package network

// FINCLOUD_APACHE_NO_VERSION

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// RouteTableDescriptionClient is the network Client
type RouteTableDescriptionClient struct {
	BaseClient
}

// NewRouteTableDescriptionClient creates an instance of the RouteTableDescriptionClient client.
func NewRouteTableDescriptionClient() RouteTableDescriptionClient {
	return NewRouteTableDescriptionClientWithBaseURI(DefaultBaseURI)
}

// NewRouteTableDescriptionClientWithBaseURI creates an instance of the RouteTableDescriptionClient client.
func NewRouteTableDescriptionClientWithBaseURI(baseURI string) RouteTableDescriptionClient {
	return RouteTableDescriptionClient{NewWithBaseURI(baseURI)}
}

// Update route Table 설명 변경
// Parameters:
// routeTableNo - route Table 번호
// parameters - route Table 설명
func (client RouteTableDescriptionClient) Update(ctx context.Context, routeTableNo string, parameters RouteTableDescriptionParameter) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RouteTableDescriptionClient.Update")
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
		err = autorest.NewErrorWithError(err, "network.RouteTableDescriptionClient", "Update", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "network.RouteTableDescriptionClient", "Update", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.RouteTableDescriptionClient", "Update", resp, "Failure responding to request")
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client RouteTableDescriptionClient) UpdatePreparer(ctx context.Context, routeTableNo string, parameters RouteTableDescriptionParameter) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"routeTableNo": autorest.Encode("path", routeTableNo),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json;charset=UTF-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/vpc-network/api/network/v1/route-tables/instances/{routeTableNo}/descriptions", pathParameters),
		autorest.WithJSON(parameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client RouteTableDescriptionClient) UpdateSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client RouteTableDescriptionClient) UpdateResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}
