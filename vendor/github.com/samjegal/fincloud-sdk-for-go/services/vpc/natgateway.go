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

// NatGatewayClient is the VPC Client
type NatGatewayClient struct {
	BaseClient
}

// NewNatGatewayClient creates an instance of the NatGatewayClient client.
func NewNatGatewayClient() NatGatewayClient {
	return NewNatGatewayClientWithBaseURI(DefaultBaseURI)
}

// NewNatGatewayClientWithBaseURI creates an instance of the NatGatewayClient client using a custom endpoint.  Use this
// when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewNatGatewayClientWithBaseURI(baseURI string) NatGatewayClient {
	return NatGatewayClient{NewWithBaseURI(baseURI)}
}

// Create NAT Gateway 인스턴스를 생성
// Parameters:
// vpcNo - VPC 번호
// zoneCode - ZONE 코드
// regionCode - REGION 코드
// natGatewayName - NAT Gateway 이름
// natGatewayDescription - NAT Gateway 설명
func (client NatGatewayClient) Create(ctx context.Context, vpcNo string, zoneCode string, regionCode string, natGatewayName string, natGatewayDescription string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/NatGatewayClient.Create")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreatePreparer(ctx, vpcNo, zoneCode, regionCode, natGatewayName, natGatewayDescription)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vpc.NatGatewayClient", "Create", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "vpc.NatGatewayClient", "Create", resp, "Failure sending request")
		return
	}

	result, err = client.CreateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vpc.NatGatewayClient", "Create", resp, "Failure responding to request")
	}

	return
}

// CreatePreparer prepares the Create request.
func (client NatGatewayClient) CreatePreparer(ctx context.Context, vpcNo string, zoneCode string, regionCode string, natGatewayName string, natGatewayDescription string) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"responseFormatType": autorest.Encode("query", "json"),
		"vpcNo":              autorest.Encode("query", vpcNo),
		"zoneCode":           autorest.Encode("query", zoneCode),
	}
	if len(regionCode) > 0 {
		queryParameters["regionCode"] = autorest.Encode("query", regionCode)
	} else {
		queryParameters["regionCode"] = autorest.Encode("query", "FKR")
	}
	if len(natGatewayName) > 0 {
		queryParameters["natGatewayName"] = autorest.Encode("query", natGatewayName)
	}
	if len(natGatewayDescription) > 0 {
		queryParameters["natGatewayDescription"] = autorest.Encode("query", natGatewayDescription)
	}

	timestamp := strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)
	sec := security.NewSignature(client.Client.Secretkey, crypto.SHA256)
	signature, err := sec.Signature("POST", autorest.GetPath(DefaultBaseURI, "/createNatGatewayInstance")+"?"+autorest.GetQuery(queryParameters), client.Client.AccessKey, timestamp)
	if err != nil {
		return nil, err
	}
	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/createNatGatewayInstance"),
		autorest.WithQueryParameters(queryParameters),
		autorest.WithHeader("x-ncp-apigw-timestamp", timestamp),
		autorest.WithHeader("x-ncp-iam-access-key", client.Client.AccessKey),
		autorest.WithHeader("x-ncp-apigw-signature-v2", signature))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client NatGatewayClient) CreateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client NatGatewayClient) CreateResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Delete NAT Gateway 인스턴스를 삭제
// Parameters:
// natGatewayInstanceNo - NAT Gateway 인스턴스 번호
// regionCode - REGION 코드
func (client NatGatewayClient) Delete(ctx context.Context, natGatewayInstanceNo string, regionCode string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/NatGatewayClient.Delete")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, natGatewayInstanceNo, regionCode)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vpc.NatGatewayClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "vpc.NatGatewayClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vpc.NatGatewayClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client NatGatewayClient) DeletePreparer(ctx context.Context, natGatewayInstanceNo string, regionCode string) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"natGatewayInstanceNo": autorest.Encode("query", natGatewayInstanceNo),
		"responseFormatType":   autorest.Encode("query", "json"),
	}
	if len(regionCode) > 0 {
		queryParameters["regionCode"] = autorest.Encode("query", regionCode)
	} else {
		queryParameters["regionCode"] = autorest.Encode("query", "FKR")
	}

	timestamp := strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)
	sec := security.NewSignature(client.Client.Secretkey, crypto.SHA256)
	signature, err := sec.Signature("POST", autorest.GetPath(DefaultBaseURI, "/deleteNatGatewayInstance")+"?"+autorest.GetQuery(queryParameters), client.Client.AccessKey, timestamp)
	if err != nil {
		return nil, err
	}
	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/deleteNatGatewayInstance"),
		autorest.WithQueryParameters(queryParameters),
		autorest.WithHeader("x-ncp-apigw-timestamp", timestamp),
		autorest.WithHeader("x-ncp-iam-access-key", client.Client.AccessKey),
		autorest.WithHeader("x-ncp-apigw-signature-v2", signature))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client NatGatewayClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client NatGatewayClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// GetDetail NAT Gateway 인스턴스 상세 정보를 조회
// Parameters:
// natGatewayInstanceNo - NAT Gateway 인스턴스 번호
// regionCode - REGION 코드
func (client NatGatewayClient) GetDetail(ctx context.Context, natGatewayInstanceNo string, regionCode string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/NatGatewayClient.GetDetail")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetDetailPreparer(ctx, natGatewayInstanceNo, regionCode)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vpc.NatGatewayClient", "GetDetail", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetDetailSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "vpc.NatGatewayClient", "GetDetail", resp, "Failure sending request")
		return
	}

	result, err = client.GetDetailResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vpc.NatGatewayClient", "GetDetail", resp, "Failure responding to request")
	}

	return
}

// GetDetailPreparer prepares the GetDetail request.
func (client NatGatewayClient) GetDetailPreparer(ctx context.Context, natGatewayInstanceNo string, regionCode string) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"natGatewayInstanceNo": autorest.Encode("query", natGatewayInstanceNo),
		"responseFormatType":   autorest.Encode("query", "json"),
	}
	if len(regionCode) > 0 {
		queryParameters["regionCode"] = autorest.Encode("query", regionCode)
	} else {
		queryParameters["regionCode"] = autorest.Encode("query", "FKR")
	}

	timestamp := strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)
	sec := security.NewSignature(client.Client.Secretkey, crypto.SHA256)
	signature, err := sec.Signature("GET", autorest.GetPath(DefaultBaseURI, "/getNatGatewayInstanceDetail")+"?"+autorest.GetQuery(queryParameters), client.Client.AccessKey, timestamp)
	if err != nil {
		return nil, err
	}
	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/getNatGatewayInstanceDetail"),
		autorest.WithQueryParameters(queryParameters),
		autorest.WithHeader("x-ncp-apigw-timestamp", timestamp),
		autorest.WithHeader("x-ncp-iam-access-key", client.Client.AccessKey),
		autorest.WithHeader("x-ncp-apigw-signature-v2", signature))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetDetailSender sends the GetDetail request. The method will close the
// http.Response Body if it receives an error.
func (client NatGatewayClient) GetDetailSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetDetailResponder handles the response to the GetDetail request. The method always
// closes the http.Response Body.
func (client NatGatewayClient) GetDetailResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// GetList NAT Gateway 인스턴스 리스트를 조회
// Parameters:
// regionCode - REGION 코드
// natGatewayInstanceNoListN - NAT Gateway 인스턴스 번호 리스트
// publicIP - 공인 IP 주소
// vpcName - VPC 이름
// natGatewayName - NAT Gateway 이름
// natGatewayInstanceStatusCode - NAT Gateway 인스턴스 상태 코드
// pageNo - 페이지 번호
// pageSize - 페이지 사이즈
func (client NatGatewayClient) GetList(ctx context.Context, regionCode string, natGatewayInstanceNoListN string, publicIP string, vpcName string, natGatewayName string, natGatewayInstanceStatusCode NatGatewayInstanceStatusCode, pageNo string, pageSize string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/NatGatewayClient.GetList")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetListPreparer(ctx, regionCode, natGatewayInstanceNoListN, publicIP, vpcName, natGatewayName, natGatewayInstanceStatusCode, pageNo, pageSize)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vpc.NatGatewayClient", "GetList", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetListSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "vpc.NatGatewayClient", "GetList", resp, "Failure sending request")
		return
	}

	result, err = client.GetListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vpc.NatGatewayClient", "GetList", resp, "Failure responding to request")
	}

	return
}

// GetListPreparer prepares the GetList request.
func (client NatGatewayClient) GetListPreparer(ctx context.Context, regionCode string, natGatewayInstanceNoListN string, publicIP string, vpcName string, natGatewayName string, natGatewayInstanceStatusCode NatGatewayInstanceStatusCode, pageNo string, pageSize string) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"responseFormatType": autorest.Encode("query", "json"),
	}
	if len(regionCode) > 0 {
		queryParameters["regionCode"] = autorest.Encode("query", regionCode)
	} else {
		queryParameters["regionCode"] = autorest.Encode("query", "FKR")
	}
	if len(natGatewayInstanceNoListN) > 0 {
		queryParameters["natGatewayInstanceNoList.N"] = autorest.Encode("query", natGatewayInstanceNoListN)
	}
	if len(publicIP) > 0 {
		queryParameters["publicIp"] = autorest.Encode("query", publicIP)
	}
	if len(vpcName) > 0 {
		queryParameters["vpcName"] = autorest.Encode("query", vpcName)
	}
	if len(natGatewayName) > 0 {
		queryParameters["natGatewayName"] = autorest.Encode("query", natGatewayName)
	}
	if len(string(natGatewayInstanceStatusCode)) > 0 {
		queryParameters["natGatewayInstanceStatusCode"] = autorest.Encode("query", natGatewayInstanceStatusCode)
	}
	if len(pageNo) > 0 {
		queryParameters["pageNo"] = autorest.Encode("query", pageNo)
	}
	if len(pageSize) > 0 {
		queryParameters["pageSize"] = autorest.Encode("query", pageSize)
	}

	timestamp := strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)
	sec := security.NewSignature(client.Client.Secretkey, crypto.SHA256)
	signature, err := sec.Signature("GET", autorest.GetPath(DefaultBaseURI, "/getNatGatewayInstanceList")+"?"+autorest.GetQuery(queryParameters), client.Client.AccessKey, timestamp)
	if err != nil {
		return nil, err
	}
	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/getNatGatewayInstanceList"),
		autorest.WithQueryParameters(queryParameters),
		autorest.WithHeader("x-ncp-apigw-timestamp", timestamp),
		autorest.WithHeader("x-ncp-iam-access-key", client.Client.AccessKey),
		autorest.WithHeader("x-ncp-apigw-signature-v2", signature))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetListSender sends the GetList request. The method will close the
// http.Response Body if it receives an error.
func (client NatGatewayClient) GetListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetListResponder handles the response to the GetList request. The method always
// closes the http.Response Body.
func (client NatGatewayClient) GetListResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}
