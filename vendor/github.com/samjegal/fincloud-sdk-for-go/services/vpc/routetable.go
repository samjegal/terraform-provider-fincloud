package vpc

// FINCLOUD_APACHE_NO_VERSION

import (
	"context"
	"crypto"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"github.com/samjegal/fincloud-sdk-for-go/common"
	"github.com/samjegal/go-fincloud-helpers/security"
	"net/http"
	"strconv"
	"time"
)

// RouteTableClient is the VPC Client
type RouteTableClient struct {
	BaseClient
}

// NewRouteTableClient creates an instance of the RouteTableClient client.
func NewRouteTableClient() RouteTableClient {
	return NewRouteTableClientWithBaseURI(DefaultBaseURI)
}

func NewRouteTableClientWithKey(accessKey string, secretKey string) RouteTableClient {
	return NewRouteTableClientWithBaseURIWithKey(DefaultBaseURI, accessKey, secretKey)
}

// NewRouteTableClientWithBaseURI creates an instance of the RouteTableClient client using a custom endpoint.  Use this
// when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewRouteTableClientWithBaseURI(baseURI string) RouteTableClient {
	return RouteTableClient{NewWithBaseURI(baseURI)}
}

func NewRouteTableClientWithBaseURIWithKey(baseURI string, accessKey string, secretKey string) RouteTableClient {
	return RouteTableClient{NewWithBaseURIWithKey(baseURI, accessKey, secretKey)}
}

// Create 라우트 테이블을 생성
// Parameters:
// vpcNo - VPC 번호
// routeTableName - 라우트 테이블 이름
// supportedSubnetTypeCode - 지원하는 서브넷 유형 코드
// routeTableDescription - 라우트 테이블 설명
func (client RouteTableClient) Create(ctx context.Context, vpcNo string, routeTableName string, supportedSubnetTypeCode SupportedSubnetTypeCode, routeTableDescription string) (result RouteTableCreateResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RouteTableClient.Create")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreatePreparer(ctx, vpcNo, routeTableName, supportedSubnetTypeCode, routeTableDescription)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vpc.RouteTableClient", "Create", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "vpc.RouteTableClient", "Create", resp, "Failure sending request")
		return
	}

	result, err = client.CreateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vpc.RouteTableClient", "Create", resp, "Failure responding to request")
	}

	return
}

// CreatePreparer prepares the Create request.
func (client RouteTableClient) CreatePreparer(ctx context.Context, vpcNo string, routeTableName string, supportedSubnetTypeCode SupportedSubnetTypeCode, routeTableDescription string) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"responseFormatType": autorest.Encode("query", "json"),
		"vpcNo":              autorest.Encode("query", vpcNo),
	}

	queryParameters["regionCode"] = autorest.Encode("query", "FKR")

	if len(routeTableName) > 0 {
		queryParameters["routeTableName"] = autorest.Encode("query", routeTableName)
	}
	if len(string(supportedSubnetTypeCode)) > 0 {
		queryParameters["supportedSubnetTypeCode"] = autorest.Encode("query", supportedSubnetTypeCode)
	}
	if len(routeTableDescription) > 0 {
		queryParameters["routeTableDescription"] = autorest.Encode("query", routeTableDescription)
	}

	timestamp := strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)
	sec := security.NewSignature(client.Secretkey, crypto.SHA256)
	signature, err := sec.Signature("POST", common.GetPath(DefaultBaseURI, "/createRouteTable")+"?"+common.GetQuery(queryParameters), client.AccessKey, timestamp)
	if err != nil {
		return nil, err
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/createRouteTable"),
		autorest.WithQueryParameters(queryParameters),
		autorest.WithHeader("x-ncp-apigw-timestamp", timestamp),
		autorest.WithHeader("x-ncp-iam-access-key", client.AccessKey),
		autorest.WithHeader("x-ncp-apigw-signature-v2", signature))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client RouteTableClient) CreateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client RouteTableClient) CreateResponder(resp *http.Response) (result RouteTableCreateResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete 라우트 테이블을 삭제
// Parameters:
// routeTableNo - 라우트 테이블 번호
func (client RouteTableClient) Delete(ctx context.Context, routeTableNo string) (result RouteTableDeleteResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RouteTableClient.Delete")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, routeTableNo)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vpc.RouteTableClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "vpc.RouteTableClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vpc.RouteTableClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client RouteTableClient) DeletePreparer(ctx context.Context, routeTableNo string) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"responseFormatType": autorest.Encode("query", "json"),
		"routeTableNo":       autorest.Encode("query", routeTableNo),
	}

	queryParameters["regionCode"] = autorest.Encode("query", "FKR")

	timestamp := strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)
	sec := security.NewSignature(client.Secretkey, crypto.SHA256)
	signature, err := sec.Signature("POST", common.GetPath(DefaultBaseURI, "/deleteRouteTable")+"?"+common.GetQuery(queryParameters), client.AccessKey, timestamp)
	if err != nil {
		return nil, err
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/deleteRouteTable"),
		autorest.WithQueryParameters(queryParameters),
		autorest.WithHeader("x-ncp-apigw-timestamp", timestamp),
		autorest.WithHeader("x-ncp-iam-access-key", client.AccessKey),
		autorest.WithHeader("x-ncp-apigw-signature-v2", signature))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client RouteTableClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client RouteTableClient) DeleteResponder(resp *http.Response) (result RouteTableDeleteResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetDetail 라우트 테이블 상세 정보를 조회
// Parameters:
// routeTableNo - 라우트 테이블 번호
func (client RouteTableClient) GetDetail(ctx context.Context, routeTableNo string) (result RouteTableDetailResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RouteTableClient.GetDetail")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetDetailPreparer(ctx, routeTableNo)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vpc.RouteTableClient", "GetDetail", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetDetailSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "vpc.RouteTableClient", "GetDetail", resp, "Failure sending request")
		return
	}

	result, err = client.GetDetailResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vpc.RouteTableClient", "GetDetail", resp, "Failure responding to request")
	}

	return
}

// GetDetailPreparer prepares the GetDetail request.
func (client RouteTableClient) GetDetailPreparer(ctx context.Context, routeTableNo string) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"responseFormatType": autorest.Encode("query", "json"),
		"routeTableNo":       autorest.Encode("query", routeTableNo),
	}

	queryParameters["regionCode"] = autorest.Encode("query", "FKR")

	timestamp := strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)
	sec := security.NewSignature(client.Secretkey, crypto.SHA256)
	signature, err := sec.Signature("GET", common.GetPath(DefaultBaseURI, "/getRouteTableDetail")+"?"+common.GetQuery(queryParameters), client.AccessKey, timestamp)
	if err != nil {
		return nil, err
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/getRouteTableDetail"),
		autorest.WithQueryParameters(queryParameters),
		autorest.WithHeader("x-ncp-apigw-timestamp", timestamp),
		autorest.WithHeader("x-ncp-iam-access-key", client.AccessKey),
		autorest.WithHeader("x-ncp-apigw-signature-v2", signature))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetDetailSender sends the GetDetail request. The method will close the
// http.Response Body if it receives an error.
func (client RouteTableClient) GetDetailSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetDetailResponder handles the response to the GetDetail request. The method always
// closes the http.Response Body.
func (client RouteTableClient) GetDetailResponder(resp *http.Response) (result RouteTableDetailResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetList 라우트 테이블 리스트를 조회
// Parameters:
// routeTableNoListN - 라우트 테이블 번호 리스트
// routeTableName - 라우트 테이블 이름
// supportedSubnetTypeCode - 지원하는 서브넷 유형 코드
// pageNo - 페이지 번호
// pageSize - 페이지 사이즈
// sortedBy - 정렬 대상
// sortingOrder - 정렬 순서
// vpcNo - VPC 번호
func (client RouteTableClient) GetList(ctx context.Context, routeTableNoListN string, routeTableName string, supportedSubnetTypeCode SupportedSubnetTypeCode, pageNo string, pageSize string, sortedBy string, sortingOrder SortingOrder, vpcNo string) (result RouteTableListResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RouteTableClient.GetList")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetListPreparer(ctx, routeTableNoListN, routeTableName, supportedSubnetTypeCode, pageNo, pageSize, sortedBy, sortingOrder, vpcNo)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vpc.RouteTableClient", "GetList", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "vpc.RouteTableClient", "GetList", resp, "Failure sending request")
		return
	}

	result, err = client.GetListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vpc.RouteTableClient", "GetList", resp, "Failure responding to request")
	}

	return
}

// GetListPreparer prepares the GetList request.
func (client RouteTableClient) GetListPreparer(ctx context.Context, routeTableNoListN string, routeTableName string, supportedSubnetTypeCode SupportedSubnetTypeCode, pageNo string, pageSize string, sortedBy string, sortingOrder SortingOrder, vpcNo string) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"responseFormatType": autorest.Encode("query", "json"),
	}

	queryParameters["regionCode"] = autorest.Encode("query", "FKR")

	if len(routeTableNoListN) > 0 {
		queryParameters["routeTableNoList.N"] = autorest.Encode("query", routeTableNoListN)
	}
	if len(routeTableName) > 0 {
		queryParameters["routeTableName"] = autorest.Encode("query", routeTableName)
	}
	if len(string(supportedSubnetTypeCode)) > 0 {
		queryParameters["supportedSubnetTypeCode"] = autorest.Encode("query", supportedSubnetTypeCode)
	}
	if len(pageNo) > 0 {
		queryParameters["pageNo"] = autorest.Encode("query", pageNo)
	}
	if len(pageSize) > 0 {
		queryParameters["pageSize"] = autorest.Encode("query", pageSize)
	}
	if len(sortedBy) > 0 {
		queryParameters["sortedBy"] = autorest.Encode("query", sortedBy)
	}
	if len(string(sortingOrder)) > 0 {
		queryParameters["sortingOrder"] = autorest.Encode("query", sortingOrder)
	}
	if len(vpcNo) > 0 {
		queryParameters["vpcNo"] = autorest.Encode("query", vpcNo)
	}

	timestamp := strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)
	sec := security.NewSignature(client.Secretkey, crypto.SHA256)
	signature, err := sec.Signature("GET", common.GetPath(DefaultBaseURI, "/getRouteTableList")+"?"+common.GetQuery(queryParameters), client.AccessKey, timestamp)
	if err != nil {
		return nil, err
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/getRouteTableList"),
		autorest.WithQueryParameters(queryParameters),
		autorest.WithHeader("x-ncp-apigw-timestamp", timestamp),
		autorest.WithHeader("x-ncp-iam-access-key", client.AccessKey),
		autorest.WithHeader("x-ncp-apigw-signature-v2", signature))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetListSender sends the GetList request. The method will close the
// http.Response Body if it receives an error.
func (client RouteTableClient) GetListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetListResponder handles the response to the GetList request. The method always
// closes the http.Response Body.
func (client RouteTableClient) GetListResponder(resp *http.Response) (result RouteTableListResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
