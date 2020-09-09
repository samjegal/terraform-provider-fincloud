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

// RouteClient is the VPC Client
type RouteClient struct {
	BaseClient
}

// NewRouteClient creates an instance of the RouteClient client.
func NewRouteClient() RouteClient {
	return NewRouteClientWithBaseURI(DefaultBaseURI)
}

// NewRouteClientWithBaseURI creates an instance of the RouteClient client using a custom endpoint.  Use this when
// interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewRouteClientWithBaseURI(baseURI string) RouteClient {
	return RouteClient{NewWithBaseURI(baseURI)}
}

// Add 라우트를 추가
// Parameters:
// vpcNo - VPC 번호
// routeTableNo - 라우트 테이블 번호
// routeListNdestinationCidrBlock - 목적지 CIDR 블록
// routeListNtargetTypeCode - 목적지 유형 코드
// routeListNtargetNo - 목적지 번호
// routeListNtargetName - 목적지 이름
func (client RouteClient) Add(ctx context.Context, vpcNo string, routeTableNo string, routeListNdestinationCidrBlock string, routeListNtargetTypeCode TargetTypeCode, routeListNtargetNo string, routeListNtargetName string) (result RouteAddResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RouteClient.Add")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.AddPreparer(ctx, vpcNo, routeTableNo, routeListNdestinationCidrBlock, routeListNtargetTypeCode, routeListNtargetNo, routeListNtargetName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vpc.RouteClient", "Add", nil, "Failure preparing request")
		return
	}

	resp, err := client.AddSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "vpc.RouteClient", "Add", resp, "Failure sending request")
		return
	}

	result, err = client.AddResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vpc.RouteClient", "Add", resp, "Failure responding to request")
	}

	return
}

// AddPreparer prepares the Add request.
func (client RouteClient) AddPreparer(ctx context.Context, vpcNo string, routeTableNo string, routeListNdestinationCidrBlock string, routeListNtargetTypeCode TargetTypeCode, routeListNtargetNo string, routeListNtargetName string) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"responseFormatType":               autorest.Encode("query", "json"),
		"routeList.N.destinationCidrBlock": autorest.Encode("query", routeListNdestinationCidrBlock),
		"routeList.N.targetName":           autorest.Encode("query", routeListNtargetName),
		"routeList.N.targetNo":             autorest.Encode("query", routeListNtargetNo),
		"routeList.N.targetTypeCode":       autorest.Encode("query", routeListNtargetTypeCode),
		"routeTableNo":                     autorest.Encode("query", routeTableNo),
		"vpcNo":                            autorest.Encode("query", vpcNo),
	}

	queryParameters["regionCode"] = autorest.Encode("query", "FKR")

	timestamp := strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)
	sec := security.NewSignature(client.Client.Secretkey, crypto.SHA256)
	signature, err := sec.Signature("POST", autorest.GetPath(DefaultBaseURI, "/addRoute")+"?"+autorest.GetQuery(queryParameters), client.Client.AccessKey, timestamp)
	if err != nil {
		return nil, err
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/addRoute"),
		autorest.WithQueryParameters(queryParameters),
		autorest.WithHeader("x-ncp-apigw-timestamp", timestamp),
		autorest.WithHeader("x-ncp-iam-access-key", client.Client.AccessKey),
		autorest.WithHeader("x-ncp-apigw-signature-v2", signature))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// AddSender sends the Add request. The method will close the
// http.Response Body if it receives an error.
func (client RouteClient) AddSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// AddResponder handles the response to the Add request. The method always
// closes the http.Response Body.
func (client RouteClient) AddResponder(resp *http.Response) (result RouteAddResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetList 라우트 리스트를 조회
// Parameters:
// routeTableNo - 라우트 테이블 번호
// vpcNo - VPC 번호
func (client RouteClient) GetList(ctx context.Context, routeTableNo string, vpcNo string) (result RouteListResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RouteClient.GetList")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetListPreparer(ctx, routeTableNo, vpcNo)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vpc.RouteClient", "GetList", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "vpc.RouteClient", "GetList", resp, "Failure sending request")
		return
	}

	result, err = client.GetListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vpc.RouteClient", "GetList", resp, "Failure responding to request")
	}

	return
}

// GetListPreparer prepares the GetList request.
func (client RouteClient) GetListPreparer(ctx context.Context, routeTableNo string, vpcNo string) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"responseFormatType": autorest.Encode("query", "json"),
		"routeTableNo":       autorest.Encode("query", routeTableNo),
		"vpcNo":              autorest.Encode("query", vpcNo),
	}

	queryParameters["regionCode"] = autorest.Encode("query", "FKR")

	timestamp := strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)
	sec := security.NewSignature(client.Client.Secretkey, crypto.SHA256)
	signature, err := sec.Signature("GET", autorest.GetPath(DefaultBaseURI, "/getRouteList")+"?"+autorest.GetQuery(queryParameters), client.Client.AccessKey, timestamp)
	if err != nil {
		return nil, err
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/getRouteList"),
		autorest.WithQueryParameters(queryParameters),
		autorest.WithHeader("x-ncp-apigw-timestamp", timestamp),
		autorest.WithHeader("x-ncp-iam-access-key", client.Client.AccessKey),
		autorest.WithHeader("x-ncp-apigw-signature-v2", signature))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetListSender sends the GetList request. The method will close the
// http.Response Body if it receives an error.
func (client RouteClient) GetListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetListResponder handles the response to the GetList request. The method always
// closes the http.Response Body.
func (client RouteClient) GetListResponder(resp *http.Response) (result RouteListResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Remove 라우트를 제거
// Parameters:
// vpcNo - VPC 번호
// routeTableNo - 라우트 테이블 번호
// routeListNdestinationCidrBlock - 목적지 CIDR 블록
// routeListNtargetTypeCode - 목적지 유형 코드
// routeListNtargetNo - 목적지 번호
// routeListNtargetName - 목적지 이름
func (client RouteClient) Remove(ctx context.Context, vpcNo string, routeTableNo string, routeListNdestinationCidrBlock string, routeListNtargetTypeCode TargetTypeCode, routeListNtargetNo string, routeListNtargetName string) (result RouteRemoveResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RouteClient.Remove")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.RemovePreparer(ctx, vpcNo, routeTableNo, routeListNdestinationCidrBlock, routeListNtargetTypeCode, routeListNtargetNo, routeListNtargetName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vpc.RouteClient", "Remove", nil, "Failure preparing request")
		return
	}

	resp, err := client.RemoveSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "vpc.RouteClient", "Remove", resp, "Failure sending request")
		return
	}

	result, err = client.RemoveResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vpc.RouteClient", "Remove", resp, "Failure responding to request")
	}

	return
}

// RemovePreparer prepares the Remove request.
func (client RouteClient) RemovePreparer(ctx context.Context, vpcNo string, routeTableNo string, routeListNdestinationCidrBlock string, routeListNtargetTypeCode TargetTypeCode, routeListNtargetNo string, routeListNtargetName string) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"responseFormatType":               autorest.Encode("query", "json"),
		"routeList.N.destinationCidrBlock": autorest.Encode("query", routeListNdestinationCidrBlock),
		"routeList.N.targetName":           autorest.Encode("query", routeListNtargetName),
		"routeList.N.targetNo":             autorest.Encode("query", routeListNtargetNo),
		"routeList.N.targetTypeCode":       autorest.Encode("query", routeListNtargetTypeCode),
		"routeTableNo":                     autorest.Encode("query", routeTableNo),
		"vpcNo":                            autorest.Encode("query", vpcNo),
	}

	queryParameters["regionCode"] = autorest.Encode("query", "FKR")

	timestamp := strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)
	sec := security.NewSignature(client.Client.Secretkey, crypto.SHA256)
	signature, err := sec.Signature("POST", autorest.GetPath(DefaultBaseURI, "/removeRoute")+"?"+autorest.GetQuery(queryParameters), client.Client.AccessKey, timestamp)
	if err != nil {
		return nil, err
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/removeRoute"),
		autorest.WithQueryParameters(queryParameters),
		autorest.WithHeader("x-ncp-apigw-timestamp", timestamp),
		autorest.WithHeader("x-ncp-iam-access-key", client.Client.AccessKey),
		autorest.WithHeader("x-ncp-apigw-signature-v2", signature))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// RemoveSender sends the Remove request. The method will close the
// http.Response Body if it receives an error.
func (client RouteClient) RemoveSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// RemoveResponder handles the response to the Remove request. The method always
// closes the http.Response Body.
func (client RouteClient) RemoveResponder(resp *http.Response) (result RouteRemoveResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
