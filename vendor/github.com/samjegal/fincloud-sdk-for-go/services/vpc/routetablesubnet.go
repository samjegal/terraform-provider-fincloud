package vpc

// FINCLOUD_APACHE_NO_VERSION

import (
	"context"
	"crypto"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"github.com/samjegal/go-fincloud-helpers/security"
	"net/http"
	"strconv"
	"time"
)

// RouteTableSubnetClient is the VPC Client
type RouteTableSubnetClient struct {
	BaseClient
}

// NewRouteTableSubnetClient creates an instance of the RouteTableSubnetClient client.
func NewRouteTableSubnetClient() RouteTableSubnetClient {
	return NewRouteTableSubnetClientWithBaseURI(DefaultBaseURI)
}

// NewRouteTableSubnetClientWithBaseURI creates an instance of the RouteTableSubnetClient client using a custom
// endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure
// stack).
func NewRouteTableSubnetClientWithBaseURI(baseURI string) RouteTableSubnetClient {
	return RouteTableSubnetClient{NewWithBaseURI(baseURI)}
}

// Add 라우트 테이블의 연관 서브넷을 추가
// Parameters:
// vpcNo - VPC 번호
// routeTableNo - 라우트 테이블 번호
// subnetNoListN - 서브넷 번호 리스트
func (client RouteTableSubnetClient) Add(ctx context.Context, vpcNo string, routeTableNo string, subnetNoListN string) (result RouteTableSubnetAddResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RouteTableSubnetClient.Add")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.AddPreparer(ctx, vpcNo, routeTableNo, subnetNoListN)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vpc.RouteTableSubnetClient", "Add", nil, "Failure preparing request")
		return
	}

	resp, err := client.AddSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "vpc.RouteTableSubnetClient", "Add", resp, "Failure sending request")
		return
	}

	result, err = client.AddResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vpc.RouteTableSubnetClient", "Add", resp, "Failure responding to request")
	}

	return
}

// AddPreparer prepares the Add request.
func (client RouteTableSubnetClient) AddPreparer(ctx context.Context, vpcNo string, routeTableNo string, subnetNoListN string) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"responseFormatType": autorest.Encode("query", "json"),
		"routeTableNo":       autorest.Encode("query", routeTableNo),
		"subnetNoList.N":     autorest.Encode("query", subnetNoListN),
		"vpcNo":              autorest.Encode("query", vpcNo),
	}

	queryParameters["regionCode"] = autorest.Encode("query", "FKR")

	timestamp := strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)
	sec := security.NewSignature(client.Client.Secretkey, crypto.SHA256)
	signature, err := sec.Signature("POST", autorest.GetPath(DefaultBaseURI, "/addRouteTableSubnet")+"?"+autorest.GetQuery(queryParameters), client.Client.AccessKey, timestamp)
	if err != nil {
		return nil, err
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/addRouteTableSubnet"),
		autorest.WithQueryParameters(queryParameters),
		autorest.WithHeader("x-ncp-apigw-timestamp", timestamp),
		autorest.WithHeader("x-ncp-iam-access-key", client.Client.AccessKey),
		autorest.WithHeader("x-ncp-apigw-signature-v2", signature))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// AddSender sends the Add request. The method will close the
// http.Response Body if it receives an error.
func (client RouteTableSubnetClient) AddSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// AddResponder handles the response to the Add request. The method always
// closes the http.Response Body.
func (client RouteTableSubnetClient) AddResponder(resp *http.Response) (result RouteTableSubnetAddResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetList 라우트 테이블에 연관된 서브넷 리스트를 조회
// Parameters:
// routeTableNo - 라우트 테이블 번호
func (client RouteTableSubnetClient) GetList(ctx context.Context, routeTableNo string) (result RouteTableSubnetListResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RouteTableSubnetClient.GetList")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetListPreparer(ctx, routeTableNo)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vpc.RouteTableSubnetClient", "GetList", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "vpc.RouteTableSubnetClient", "GetList", resp, "Failure sending request")
		return
	}

	result, err = client.GetListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vpc.RouteTableSubnetClient", "GetList", resp, "Failure responding to request")
	}

	return
}

// GetListPreparer prepares the GetList request.
func (client RouteTableSubnetClient) GetListPreparer(ctx context.Context, routeTableNo string) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"responseFormatType": autorest.Encode("query", "json"),
		"routeTableNo":       autorest.Encode("query", routeTableNo),
	}

	queryParameters["regionCode"] = autorest.Encode("query", "FKR")

	timestamp := strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)
	sec := security.NewSignature(client.Client.Secretkey, crypto.SHA256)
	signature, err := sec.Signature("GET", autorest.GetPath(DefaultBaseURI, "/getRouteTableSubnetList")+"?"+autorest.GetQuery(queryParameters), client.Client.AccessKey, timestamp)
	if err != nil {
		return nil, err
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/getRouteTableSubnetList"),
		autorest.WithQueryParameters(queryParameters),
		autorest.WithHeader("x-ncp-apigw-timestamp", timestamp),
		autorest.WithHeader("x-ncp-iam-access-key", client.Client.AccessKey),
		autorest.WithHeader("x-ncp-apigw-signature-v2", signature))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetListSender sends the GetList request. The method will close the
// http.Response Body if it receives an error.
func (client RouteTableSubnetClient) GetListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetListResponder handles the response to the GetList request. The method always
// closes the http.Response Body.
func (client RouteTableSubnetClient) GetListResponder(resp *http.Response) (result RouteTableSubnetListResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Remove 라우트 테이블의 연관 서브넷을 제거
// Parameters:
// vpcNo - VPC 번호
// routeTableNo - 라우트 테이블 번호
// subnetNoListN - 서브넷 번호 리스트
func (client RouteTableSubnetClient) Remove(ctx context.Context, vpcNo string, routeTableNo string, subnetNoListN string) (result RouteTableSubnetRemoveResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RouteTableSubnetClient.Remove")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.RemovePreparer(ctx, vpcNo, routeTableNo, subnetNoListN)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vpc.RouteTableSubnetClient", "Remove", nil, "Failure preparing request")
		return
	}

	resp, err := client.RemoveSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "vpc.RouteTableSubnetClient", "Remove", resp, "Failure sending request")
		return
	}

	result, err = client.RemoveResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vpc.RouteTableSubnetClient", "Remove", resp, "Failure responding to request")
	}

	return
}

// RemovePreparer prepares the Remove request.
func (client RouteTableSubnetClient) RemovePreparer(ctx context.Context, vpcNo string, routeTableNo string, subnetNoListN string) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"responseFormatType": autorest.Encode("query", "json"),
		"routeTableNo":       autorest.Encode("query", routeTableNo),
		"subnetNoList.N":     autorest.Encode("query", subnetNoListN),
		"vpcNo":              autorest.Encode("query", vpcNo),
	}

	queryParameters["regionCode"] = autorest.Encode("query", "FKR")

	timestamp := strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)
	sec := security.NewSignature(client.Client.Secretkey, crypto.SHA256)
	signature, err := sec.Signature("POST", autorest.GetPath(DefaultBaseURI, "/removeRouteTableSubnet")+"?"+autorest.GetQuery(queryParameters), client.Client.AccessKey, timestamp)
	if err != nil {
		return nil, err
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/removeRouteTableSubnet"),
		autorest.WithQueryParameters(queryParameters),
		autorest.WithHeader("x-ncp-apigw-timestamp", timestamp),
		autorest.WithHeader("x-ncp-iam-access-key", client.Client.AccessKey),
		autorest.WithHeader("x-ncp-apigw-signature-v2", signature))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// RemoveSender sends the Remove request. The method will close the
// http.Response Body if it receives an error.
func (client RouteTableSubnetClient) RemoveSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// RemoveResponder handles the response to the Remove request. The method always
// closes the http.Response Body.
func (client RouteTableSubnetClient) RemoveResponder(resp *http.Response) (result RouteTableSubnetRemoveResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
