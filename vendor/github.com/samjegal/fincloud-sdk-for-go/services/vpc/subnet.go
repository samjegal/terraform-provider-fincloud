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

// SubnetClient is the VPC Client
type SubnetClient struct {
	BaseClient
}

// NewSubnetClient creates an instance of the SubnetClient client.
func NewSubnetClient() SubnetClient {
	return NewSubnetClientWithBaseURI(DefaultBaseURI)
}

// NewSubnetClientWithBaseURI creates an instance of the SubnetClient client using a custom endpoint.  Use this when
// interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewSubnetClientWithBaseURI(baseURI string) SubnetClient {
	return SubnetClient{NewWithBaseURI(baseURI)}
}

// Create subnet을 생성
// Parameters:
// zoneCode - ZONE 코드
// vpcNo - VPC 번호
// subnet - 서브넷
// networkACLNo - 네트워크 ACL 번호
// subnetTypeCode - 서브넷 유형 코드
// subnetName - 서브넷 이름
// usageTypeCode - 용도 유형 코드
func (client SubnetClient) Create(ctx context.Context, zoneCode string, vpcNo string, subnet string, networkACLNo string, subnetTypeCode SubnetTypeCode, subnetName string, usageTypeCode UsageTypeCode) (result SubnetCreateResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SubnetClient.Create")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreatePreparer(ctx, zoneCode, vpcNo, subnet, networkACLNo, subnetTypeCode, subnetName, usageTypeCode)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vpc.SubnetClient", "Create", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "vpc.SubnetClient", "Create", resp, "Failure sending request")
		return
	}

	result, err = client.CreateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vpc.SubnetClient", "Create", resp, "Failure responding to request")
	}

	return
}

// CreatePreparer prepares the Create request.
func (client SubnetClient) CreatePreparer(ctx context.Context, zoneCode string, vpcNo string, subnet string, networkACLNo string, subnetTypeCode SubnetTypeCode, subnetName string, usageTypeCode UsageTypeCode) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"networkAclNo":       autorest.Encode("query", networkACLNo),
		"responseFormatType": autorest.Encode("query", "json"),
		"subnet":             autorest.Encode("query", subnet),
		"subnetTypeCode":     autorest.Encode("query", subnetTypeCode),
		"vpcNo":              autorest.Encode("query", vpcNo),
		"zoneCode":           autorest.Encode("query", zoneCode),
	}

	queryParameters["regionCode"] = autorest.Encode("query", "FKR")

	if len(subnetName) > 0 {
		queryParameters["subnetName"] = autorest.Encode("query", subnetName)
	}
	if len(string(usageTypeCode)) > 0 {
		queryParameters["usageTypeCode"] = autorest.Encode("query", usageTypeCode)
	}

	timestamp := strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)
	sec := security.NewSignature(client.Client.Secretkey, crypto.SHA256)
	signature, err := sec.Signature("POST", autorest.GetPath(DefaultBaseURI, "/createSubnet")+"?"+autorest.GetQuery(queryParameters), client.Client.AccessKey, timestamp)
	if err != nil {
		return nil, err
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/createSubnet"),
		autorest.WithQueryParameters(queryParameters),
		autorest.WithHeader("x-ncp-apigw-timestamp", timestamp),
		autorest.WithHeader("x-ncp-iam-access-key", client.Client.AccessKey),
		autorest.WithHeader("x-ncp-apigw-signature-v2", signature))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client SubnetClient) CreateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client SubnetClient) CreateResponder(resp *http.Response) (result SubnetCreateResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete subnet을 삭제
// Parameters:
// subnetNo - 서브넷 번호
func (client SubnetClient) Delete(ctx context.Context, subnetNo string) (result SubnetDeleteResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SubnetClient.Delete")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, subnetNo)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vpc.SubnetClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "vpc.SubnetClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vpc.SubnetClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client SubnetClient) DeletePreparer(ctx context.Context, subnetNo string) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"responseFormatType": autorest.Encode("query", "json"),
		"subnetNo":           autorest.Encode("query", subnetNo),
	}

	queryParameters["regionCode"] = autorest.Encode("query", "FKR")

	timestamp := strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)
	sec := security.NewSignature(client.Client.Secretkey, crypto.SHA256)
	signature, err := sec.Signature("POST", autorest.GetPath(DefaultBaseURI, "/deleteSubnet")+"?"+autorest.GetQuery(queryParameters), client.Client.AccessKey, timestamp)
	if err != nil {
		return nil, err
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/deleteSubnet"),
		autorest.WithQueryParameters(queryParameters),
		autorest.WithHeader("x-ncp-apigw-timestamp", timestamp),
		autorest.WithHeader("x-ncp-iam-access-key", client.Client.AccessKey),
		autorest.WithHeader("x-ncp-apigw-signature-v2", signature))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client SubnetClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client SubnetClient) DeleteResponder(resp *http.Response) (result SubnetDeleteResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetDetail subnet 상세 정보를 조회
// Parameters:
// subnetNo - 서브넷 번호
func (client SubnetClient) GetDetail(ctx context.Context, subnetNo string) (result SubnetDetailResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SubnetClient.GetDetail")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetDetailPreparer(ctx, subnetNo)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vpc.SubnetClient", "GetDetail", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetDetailSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "vpc.SubnetClient", "GetDetail", resp, "Failure sending request")
		return
	}

	result, err = client.GetDetailResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vpc.SubnetClient", "GetDetail", resp, "Failure responding to request")
	}

	return
}

// GetDetailPreparer prepares the GetDetail request.
func (client SubnetClient) GetDetailPreparer(ctx context.Context, subnetNo string) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"responseFormatType": autorest.Encode("query", "json"),
		"subnetNo":           autorest.Encode("query", subnetNo),
	}

	queryParameters["regionCode"] = autorest.Encode("query", "FKR")

	timestamp := strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)
	sec := security.NewSignature(client.Client.Secretkey, crypto.SHA256)
	signature, err := sec.Signature("GET", autorest.GetPath(DefaultBaseURI, "/getSubnetDetail")+"?"+autorest.GetQuery(queryParameters), client.Client.AccessKey, timestamp)
	if err != nil {
		return nil, err
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/getSubnetDetail"),
		autorest.WithQueryParameters(queryParameters),
		autorest.WithHeader("x-ncp-apigw-timestamp", timestamp),
		autorest.WithHeader("x-ncp-iam-access-key", client.Client.AccessKey),
		autorest.WithHeader("x-ncp-apigw-signature-v2", signature))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetDetailSender sends the GetDetail request. The method will close the
// http.Response Body if it receives an error.
func (client SubnetClient) GetDetailSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetDetailResponder handles the response to the GetDetail request. The method always
// closes the http.Response Body.
func (client SubnetClient) GetDetailResponder(resp *http.Response) (result SubnetDetailResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetList subnet 리스트를 조회
// Parameters:
// subnetNoListN - 서브넷 번호 리스트
// subnetName - 서브넷 이름
// subnet - 서브넷
// subnetTypeCode - 서브넷 유형 코드
// usageTypeCode - 용도 유형 코드
// networkACLNo - 네트워크 ACL 번호
// pageNo - 페이지 번호
// pageSize - 페이지 사이즈
// subnetStatusCode - 서브넷 상태 코드
// vpcNo - VPC 번호
// zoneCode - ZONE 코드
func (client SubnetClient) GetList(ctx context.Context, subnetNoListN string, subnetName string, subnet string, subnetTypeCode SubnetTypeCode, usageTypeCode UsageTypeCode, networkACLNo string, pageNo string, pageSize string, subnetStatusCode SubnetStatusCode, vpcNo string, zoneCode string) (result SubnetListResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SubnetClient.GetList")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetListPreparer(ctx, subnetNoListN, subnetName, subnet, subnetTypeCode, usageTypeCode, networkACLNo, pageNo, pageSize, subnetStatusCode, vpcNo, zoneCode)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vpc.SubnetClient", "GetList", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "vpc.SubnetClient", "GetList", resp, "Failure sending request")
		return
	}

	result, err = client.GetListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vpc.SubnetClient", "GetList", resp, "Failure responding to request")
	}

	return
}

// GetListPreparer prepares the GetList request.
func (client SubnetClient) GetListPreparer(ctx context.Context, subnetNoListN string, subnetName string, subnet string, subnetTypeCode SubnetTypeCode, usageTypeCode UsageTypeCode, networkACLNo string, pageNo string, pageSize string, subnetStatusCode SubnetStatusCode, vpcNo string, zoneCode string) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"responseFormatType": autorest.Encode("query", "json"),
	}

	queryParameters["regionCode"] = autorest.Encode("query", "FKR")

	if len(subnetNoListN) > 0 {
		queryParameters["subnetNoList.N"] = autorest.Encode("query", subnetNoListN)
	}
	if len(subnetName) > 0 {
		queryParameters["subnetName"] = autorest.Encode("query", subnetName)
	}
	if len(subnet) > 0 {
		queryParameters["subnet"] = autorest.Encode("query", subnet)
	}
	if len(string(subnetTypeCode)) > 0 {
		queryParameters["subnetTypeCode"] = autorest.Encode("query", subnetTypeCode)
	}
	if len(string(usageTypeCode)) > 0 {
		queryParameters["usageTypeCode"] = autorest.Encode("query", usageTypeCode)
	}
	if len(networkACLNo) > 0 {
		queryParameters["networkAclNo"] = autorest.Encode("query", networkACLNo)
	}
	if len(pageNo) > 0 {
		queryParameters["pageNo"] = autorest.Encode("query", pageNo)
	}
	if len(pageSize) > 0 {
		queryParameters["pageSize"] = autorest.Encode("query", pageSize)
	}
	if len(string(subnetStatusCode)) > 0 {
		queryParameters["subnetStatusCode"] = autorest.Encode("query", subnetStatusCode)
	}
	if len(vpcNo) > 0 {
		queryParameters["vpcNo"] = autorest.Encode("query", vpcNo)
	}
	if len(zoneCode) > 0 {
		queryParameters["zoneCode"] = autorest.Encode("query", zoneCode)
	}

	timestamp := strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)
	sec := security.NewSignature(client.Client.Secretkey, crypto.SHA256)
	signature, err := sec.Signature("GET", autorest.GetPath(DefaultBaseURI, "/getSubnetList")+"?"+autorest.GetQuery(queryParameters), client.Client.AccessKey, timestamp)
	if err != nil {
		return nil, err
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/getSubnetList"),
		autorest.WithQueryParameters(queryParameters),
		autorest.WithHeader("x-ncp-apigw-timestamp", timestamp),
		autorest.WithHeader("x-ncp-iam-access-key", client.Client.AccessKey),
		autorest.WithHeader("x-ncp-apigw-signature-v2", signature))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetListSender sends the GetList request. The method will close the
// http.Response Body if it receives an error.
func (client SubnetClient) GetListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetListResponder handles the response to the GetList request. The method always
// closes the http.Response Body.
func (client SubnetClient) GetListResponder(resp *http.Response) (result SubnetListResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
