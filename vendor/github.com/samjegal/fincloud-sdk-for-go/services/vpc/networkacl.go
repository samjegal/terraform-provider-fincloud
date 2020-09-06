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

// NetworkACLClient is the VPC Client
type NetworkACLClient struct {
	BaseClient
}

// NewNetworkACLClient creates an instance of the NetworkACLClient client.
func NewNetworkACLClient() NetworkACLClient {
	return NewNetworkACLClientWithBaseURI(DefaultBaseURI)
}

// NewNetworkACLClientWithBaseURI creates an instance of the NetworkACLClient client using a custom endpoint.  Use this
// when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewNetworkACLClientWithBaseURI(baseURI string) NetworkACLClient {
	return NetworkACLClient{NewWithBaseURI(baseURI)}
}

// AddInboundRule 네트워크 ACL의 Inbound Rule을 추가
// Parameters:
// networkACLNo - 네트워크 ACL 번호
// networkACLRuleListNpriority - 우선순위
// networkACLRuleListNprotocolTypeCode - 네트워크 ACL Rule 유형 코드
// networkACLRuleListNipBlock - IP 블록
// networkACLRuleListNruleActionCode - rule 액션 코드
// regionCode - REGION 코드
// networkACLRuleListNportRange - 포트 범위
// networkACLRuleListNnetworkACLRuleDescription - 네트워크 ACL Rule 설명
func (client NetworkACLClient) AddInboundRule(ctx context.Context, networkACLNo string, networkACLRuleListNpriority string, networkACLRuleListNprotocolTypeCode ProtocolTypeCode, networkACLRuleListNipBlock string, networkACLRuleListNruleActionCode RuleActionCode, regionCode string, networkACLRuleListNportRange string, networkACLRuleListNnetworkACLRuleDescription string) (result NetworkACLInboundRuleResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/NetworkACLClient.AddInboundRule")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.AddInboundRulePreparer(ctx, networkACLNo, networkACLRuleListNpriority, networkACLRuleListNprotocolTypeCode, networkACLRuleListNipBlock, networkACLRuleListNruleActionCode, regionCode, networkACLRuleListNportRange, networkACLRuleListNnetworkACLRuleDescription)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vpc.NetworkACLClient", "AddInboundRule", nil, "Failure preparing request")
		return
	}

	resp, err := client.AddInboundRuleSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "vpc.NetworkACLClient", "AddInboundRule", resp, "Failure sending request")
		return
	}

	result, err = client.AddInboundRuleResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vpc.NetworkACLClient", "AddInboundRule", resp, "Failure responding to request")
	}

	return
}

// AddInboundRulePreparer prepares the AddInboundRule request.
func (client NetworkACLClient) AddInboundRulePreparer(ctx context.Context, networkACLNo string, networkACLRuleListNpriority string, networkACLRuleListNprotocolTypeCode ProtocolTypeCode, networkACLRuleListNipBlock string, networkACLRuleListNruleActionCode RuleActionCode, regionCode string, networkACLRuleListNportRange string, networkACLRuleListNnetworkACLRuleDescription string) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"networkAclNo":                          autorest.Encode("query", networkACLNo),
		"networkAclRuleList.N.ipBlock":          autorest.Encode("query", networkACLRuleListNipBlock),
		"networkAclRuleList.N.priority":         autorest.Encode("query", networkACLRuleListNpriority),
		"networkAclRuleList.N.protocolTypeCode": autorest.Encode("query", networkACLRuleListNprotocolTypeCode),
		"networkAclRuleList.N.ruleActionCode":   autorest.Encode("query", networkACLRuleListNruleActionCode),
		"responseFormatType":                    autorest.Encode("query", "json"),
	}
	if len(regionCode) > 0 {
		queryParameters["regionCode"] = autorest.Encode("query", regionCode)
	} else {
		queryParameters["regionCode"] = autorest.Encode("query", "FKR")
	}
	if len(networkACLRuleListNportRange) > 0 {
		queryParameters["networkAclRuleList.N.portRange"] = autorest.Encode("query", networkACLRuleListNportRange)
	}
	if len(networkACLRuleListNnetworkACLRuleDescription) > 0 {
		queryParameters["networkAclRuleList.N.networkAclRuleDescription"] = autorest.Encode("query", networkACLRuleListNnetworkACLRuleDescription)
	}

	timestamp := strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)
	sec := security.NewSignature(client.Client.Secretkey, crypto.SHA256)
	signature, err := sec.Signature("POST", autorest.GetPath(DefaultBaseURI, "/addNetworkAclInboundRule")+"?"+autorest.GetQuery(queryParameters), client.Client.AccessKey, timestamp)
	if err != nil {
		return nil, err
	}
	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/addNetworkAclInboundRule"),
		autorest.WithQueryParameters(queryParameters),
		autorest.WithHeader("x-ncp-apigw-timestamp", timestamp),
		autorest.WithHeader("x-ncp-iam-access-key", client.Client.AccessKey),
		autorest.WithHeader("x-ncp-apigw-signature-v2", signature))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// AddInboundRuleSender sends the AddInboundRule request. The method will close the
// http.Response Body if it receives an error.
func (client NetworkACLClient) AddInboundRuleSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// AddInboundRuleResponder handles the response to the AddInboundRule request. The method always
// closes the http.Response Body.
func (client NetworkACLClient) AddInboundRuleResponder(resp *http.Response) (result NetworkACLInboundRuleResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// AddOutboundRule 네트워크 ACL의 Outbound Rule을 추가
// Parameters:
// networkACLNo - 네트워크 ACL 번호
// networkACLRuleListNpriority - 우선순위
// networkACLRuleListNprotocolTypeCode - 네트워크 ACL Rule 유형 코드
// networkACLRuleListNipBlock - IP 블록
// networkACLRuleListNruleActionCode - rule 액션 코드
// regionCode - REGION 코드
// networkACLRuleListNportRange - 포트 범위
// networkACLRuleListNnetworkACLRuleDescription - 네트워크 ACL Rule 설명
func (client NetworkACLClient) AddOutboundRule(ctx context.Context, networkACLNo string, networkACLRuleListNpriority string, networkACLRuleListNprotocolTypeCode ProtocolTypeCode, networkACLRuleListNipBlock string, networkACLRuleListNruleActionCode RuleActionCode, regionCode string, networkACLRuleListNportRange string, networkACLRuleListNnetworkACLRuleDescription string) (result NetworkACLOutboundRuleResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/NetworkACLClient.AddOutboundRule")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.AddOutboundRulePreparer(ctx, networkACLNo, networkACLRuleListNpriority, networkACLRuleListNprotocolTypeCode, networkACLRuleListNipBlock, networkACLRuleListNruleActionCode, regionCode, networkACLRuleListNportRange, networkACLRuleListNnetworkACLRuleDescription)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vpc.NetworkACLClient", "AddOutboundRule", nil, "Failure preparing request")
		return
	}

	resp, err := client.AddOutboundRuleSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "vpc.NetworkACLClient", "AddOutboundRule", resp, "Failure sending request")
		return
	}

	result, err = client.AddOutboundRuleResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vpc.NetworkACLClient", "AddOutboundRule", resp, "Failure responding to request")
	}

	return
}

// AddOutboundRulePreparer prepares the AddOutboundRule request.
func (client NetworkACLClient) AddOutboundRulePreparer(ctx context.Context, networkACLNo string, networkACLRuleListNpriority string, networkACLRuleListNprotocolTypeCode ProtocolTypeCode, networkACLRuleListNipBlock string, networkACLRuleListNruleActionCode RuleActionCode, regionCode string, networkACLRuleListNportRange string, networkACLRuleListNnetworkACLRuleDescription string) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"networkAclNo":                          autorest.Encode("query", networkACLNo),
		"networkAclRuleList.N.ipBlock":          autorest.Encode("query", networkACLRuleListNipBlock),
		"networkAclRuleList.N.priority":         autorest.Encode("query", networkACLRuleListNpriority),
		"networkAclRuleList.N.protocolTypeCode": autorest.Encode("query", networkACLRuleListNprotocolTypeCode),
		"networkAclRuleList.N.ruleActionCode":   autorest.Encode("query", networkACLRuleListNruleActionCode),
		"responseFormatType":                    autorest.Encode("query", "json"),
	}
	if len(regionCode) > 0 {
		queryParameters["regionCode"] = autorest.Encode("query", regionCode)
	} else {
		queryParameters["regionCode"] = autorest.Encode("query", "FKR")
	}
	if len(networkACLRuleListNportRange) > 0 {
		queryParameters["networkAclRuleList.N.portRange"] = autorest.Encode("query", networkACLRuleListNportRange)
	}
	if len(networkACLRuleListNnetworkACLRuleDescription) > 0 {
		queryParameters["networkAclRuleList.N.networkAclRuleDescription"] = autorest.Encode("query", networkACLRuleListNnetworkACLRuleDescription)
	}

	timestamp := strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)
	sec := security.NewSignature(client.Client.Secretkey, crypto.SHA256)
	signature, err := sec.Signature("POST", autorest.GetPath(DefaultBaseURI, "/addNetworkAclOutboundRule")+"?"+autorest.GetQuery(queryParameters), client.Client.AccessKey, timestamp)
	if err != nil {
		return nil, err
	}
	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/addNetworkAclOutboundRule"),
		autorest.WithQueryParameters(queryParameters),
		autorest.WithHeader("x-ncp-apigw-timestamp", timestamp),
		autorest.WithHeader("x-ncp-iam-access-key", client.Client.AccessKey),
		autorest.WithHeader("x-ncp-apigw-signature-v2", signature))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// AddOutboundRuleSender sends the AddOutboundRule request. The method will close the
// http.Response Body if it receives an error.
func (client NetworkACLClient) AddOutboundRuleSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// AddOutboundRuleResponder handles the response to the AddOutboundRule request. The method always
// closes the http.Response Body.
func (client NetworkACLClient) AddOutboundRuleResponder(resp *http.Response) (result NetworkACLOutboundRuleResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Create 네트워크 ACL을 생성
// Parameters:
// vpcNo - VPC 번호
// regionCode - REGION 코드
// networkACLName - 네트워크 ACL 이름
// networkACLDescription - 네트워크 ACL 설명
func (client NetworkACLClient) Create(ctx context.Context, vpcNo string, regionCode string, networkACLName string, networkACLDescription string) (result NetworkACLResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/NetworkACLClient.Create")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreatePreparer(ctx, vpcNo, regionCode, networkACLName, networkACLDescription)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vpc.NetworkACLClient", "Create", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "vpc.NetworkACLClient", "Create", resp, "Failure sending request")
		return
	}

	result, err = client.CreateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vpc.NetworkACLClient", "Create", resp, "Failure responding to request")
	}

	return
}

// CreatePreparer prepares the Create request.
func (client NetworkACLClient) CreatePreparer(ctx context.Context, vpcNo string, regionCode string, networkACLName string, networkACLDescription string) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"responseFormatType": autorest.Encode("query", "json"),
		"vpcNo":              autorest.Encode("query", vpcNo),
	}
	if len(regionCode) > 0 {
		queryParameters["regionCode"] = autorest.Encode("query", regionCode)
	} else {
		queryParameters["regionCode"] = autorest.Encode("query", "FKR")
	}
	if len(networkACLName) > 0 {
		queryParameters["networkAclName"] = autorest.Encode("query", networkACLName)
	}
	if len(networkACLDescription) > 0 {
		queryParameters["networkAclDescription"] = autorest.Encode("query", networkACLDescription)
	}

	timestamp := strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)
	sec := security.NewSignature(client.Client.Secretkey, crypto.SHA256)
	signature, err := sec.Signature("POST", autorest.GetPath(DefaultBaseURI, "/createNetworkAcl")+"?"+autorest.GetQuery(queryParameters), client.Client.AccessKey, timestamp)
	if err != nil {
		return nil, err
	}
	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/createNetworkAcl"),
		autorest.WithQueryParameters(queryParameters),
		autorest.WithHeader("x-ncp-apigw-timestamp", timestamp),
		autorest.WithHeader("x-ncp-iam-access-key", client.Client.AccessKey),
		autorest.WithHeader("x-ncp-apigw-signature-v2", signature))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client NetworkACLClient) CreateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client NetworkACLClient) CreateResponder(resp *http.Response) (result NetworkACLResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete 네트워크 ACL을 삭제
// Parameters:
// networkACLNo - 네트워크 ACL 번호
// regionCode - REGION 코드
func (client NetworkACLClient) Delete(ctx context.Context, networkACLNo string, regionCode string) (result NetworkACLResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/NetworkACLClient.Delete")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, networkACLNo, regionCode)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vpc.NetworkACLClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "vpc.NetworkACLClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vpc.NetworkACLClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client NetworkACLClient) DeletePreparer(ctx context.Context, networkACLNo string, regionCode string) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"networkAclNo":       autorest.Encode("query", networkACLNo),
		"responseFormatType": autorest.Encode("query", "json"),
	}
	if len(regionCode) > 0 {
		queryParameters["regionCode"] = autorest.Encode("query", regionCode)
	} else {
		queryParameters["regionCode"] = autorest.Encode("query", "FKR")
	}

	timestamp := strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)
	sec := security.NewSignature(client.Client.Secretkey, crypto.SHA256)
	signature, err := sec.Signature("POST", autorest.GetPath(DefaultBaseURI, "/deleteNetworkAcl")+"?"+autorest.GetQuery(queryParameters), client.Client.AccessKey, timestamp)
	if err != nil {
		return nil, err
	}
	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/deleteNetworkAcl"),
		autorest.WithQueryParameters(queryParameters),
		autorest.WithHeader("x-ncp-apigw-timestamp", timestamp),
		autorest.WithHeader("x-ncp-iam-access-key", client.Client.AccessKey),
		autorest.WithHeader("x-ncp-apigw-signature-v2", signature))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client NetworkACLClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client NetworkACLClient) DeleteResponder(resp *http.Response) (result NetworkACLResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetDetail 네트워크 ACL 상세 정보를 조회
// Parameters:
// networkACLNo - 네트워크 ACL 번호
// regionCode - REGION 코드
func (client NetworkACLClient) GetDetail(ctx context.Context, networkACLNo string, regionCode string) (result NetworkACLDetailResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/NetworkACLClient.GetDetail")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetDetailPreparer(ctx, networkACLNo, regionCode)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vpc.NetworkACLClient", "GetDetail", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetDetailSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "vpc.NetworkACLClient", "GetDetail", resp, "Failure sending request")
		return
	}

	result, err = client.GetDetailResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vpc.NetworkACLClient", "GetDetail", resp, "Failure responding to request")
	}

	return
}

// GetDetailPreparer prepares the GetDetail request.
func (client NetworkACLClient) GetDetailPreparer(ctx context.Context, networkACLNo string, regionCode string) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"networkAclNo":       autorest.Encode("query", networkACLNo),
		"responseFormatType": autorest.Encode("query", "json"),
	}
	if len(regionCode) > 0 {
		queryParameters["regionCode"] = autorest.Encode("query", regionCode)
	} else {
		queryParameters["regionCode"] = autorest.Encode("query", "FKR")
	}

	timestamp := strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)
	sec := security.NewSignature(client.Client.Secretkey, crypto.SHA256)
	signature, err := sec.Signature("GET", autorest.GetPath(DefaultBaseURI, "/getNetworkAclDetail")+"?"+autorest.GetQuery(queryParameters), client.Client.AccessKey, timestamp)
	if err != nil {
		return nil, err
	}
	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/getNetworkAclDetail"),
		autorest.WithQueryParameters(queryParameters),
		autorest.WithHeader("x-ncp-apigw-timestamp", timestamp),
		autorest.WithHeader("x-ncp-iam-access-key", client.Client.AccessKey),
		autorest.WithHeader("x-ncp-apigw-signature-v2", signature))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetDetailSender sends the GetDetail request. The method will close the
// http.Response Body if it receives an error.
func (client NetworkACLClient) GetDetailSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetDetailResponder handles the response to the GetDetail request. The method always
// closes the http.Response Body.
func (client NetworkACLClient) GetDetailResponder(resp *http.Response) (result NetworkACLDetailResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetList 네트워크 ACL 리스트를 조회
// Parameters:
// regionCode - REGION 코드
// networkACLName - 네트워크 ACL 이름
// networkACLStatusCode - 네트워크 ACL 상태 코드
// networkACLNoListN - 네트워크 ACL 번호 리스트
// pageNo - 페이지 번호
// pageSize - 페이지 사이즈
// vpcNo - VPC 번호
func (client NetworkACLClient) GetList(ctx context.Context, regionCode string, networkACLName string, networkACLStatusCode NetworkACLStatusCode, networkACLNoListN string, pageNo string, pageSize string, vpcNo string) (result NetworkACLListResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/NetworkACLClient.GetList")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetListPreparer(ctx, regionCode, networkACLName, networkACLStatusCode, networkACLNoListN, pageNo, pageSize, vpcNo)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vpc.NetworkACLClient", "GetList", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "vpc.NetworkACLClient", "GetList", resp, "Failure sending request")
		return
	}

	result, err = client.GetListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vpc.NetworkACLClient", "GetList", resp, "Failure responding to request")
	}

	return
}

// GetListPreparer prepares the GetList request.
func (client NetworkACLClient) GetListPreparer(ctx context.Context, regionCode string, networkACLName string, networkACLStatusCode NetworkACLStatusCode, networkACLNoListN string, pageNo string, pageSize string, vpcNo string) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"responseFormatType": autorest.Encode("query", "json"),
	}
	if len(regionCode) > 0 {
		queryParameters["regionCode"] = autorest.Encode("query", regionCode)
	} else {
		queryParameters["regionCode"] = autorest.Encode("query", "FKR")
	}
	if len(networkACLName) > 0 {
		queryParameters["networkAclName"] = autorest.Encode("query", networkACLName)
	}
	if len(string(networkACLStatusCode)) > 0 {
		queryParameters["networkAclStatusCode"] = autorest.Encode("query", networkACLStatusCode)
	}
	if len(networkACLNoListN) > 0 {
		queryParameters["networkAclNoList.N"] = autorest.Encode("query", networkACLNoListN)
	}
	if len(pageNo) > 0 {
		queryParameters["pageNo"] = autorest.Encode("query", pageNo)
	}
	if len(pageSize) > 0 {
		queryParameters["pageSize"] = autorest.Encode("query", pageSize)
	}
	if len(vpcNo) > 0 {
		queryParameters["vpcNo"] = autorest.Encode("query", vpcNo)
	}

	timestamp := strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)
	sec := security.NewSignature(client.Client.Secretkey, crypto.SHA256)
	signature, err := sec.Signature("GET", autorest.GetPath(DefaultBaseURI, "/getNetworkAclList")+"?"+autorest.GetQuery(queryParameters), client.Client.AccessKey, timestamp)
	if err != nil {
		return nil, err
	}
	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/getNetworkAclList"),
		autorest.WithQueryParameters(queryParameters),
		autorest.WithHeader("x-ncp-apigw-timestamp", timestamp),
		autorest.WithHeader("x-ncp-iam-access-key", client.Client.AccessKey),
		autorest.WithHeader("x-ncp-apigw-signature-v2", signature))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetListSender sends the GetList request. The method will close the
// http.Response Body if it receives an error.
func (client NetworkACLClient) GetListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetListResponder handles the response to the GetList request. The method always
// closes the http.Response Body.
func (client NetworkACLClient) GetListResponder(resp *http.Response) (result NetworkACLListResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetRuleList 네트워크 ACL의 Rule 리스트를 조회
// Parameters:
// networkACLNo - 네트워크 ACL 번호
// regionCode - REGION 코드
// networkACLRuleTypeCode - 네트워크 ACL Rule 유형 코드
func (client NetworkACLClient) GetRuleList(ctx context.Context, networkACLNo string, regionCode string, networkACLRuleTypeCode NetworkACLRuleTypeCode) (result NetworkACLRuleListResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/NetworkACLClient.GetRuleList")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetRuleListPreparer(ctx, networkACLNo, regionCode, networkACLRuleTypeCode)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vpc.NetworkACLClient", "GetRuleList", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetRuleListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "vpc.NetworkACLClient", "GetRuleList", resp, "Failure sending request")
		return
	}

	result, err = client.GetRuleListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vpc.NetworkACLClient", "GetRuleList", resp, "Failure responding to request")
	}

	return
}

// GetRuleListPreparer prepares the GetRuleList request.
func (client NetworkACLClient) GetRuleListPreparer(ctx context.Context, networkACLNo string, regionCode string, networkACLRuleTypeCode NetworkACLRuleTypeCode) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"networkAclNo":       autorest.Encode("query", networkACLNo),
		"responseFormatType": autorest.Encode("query", "json"),
	}
	if len(regionCode) > 0 {
		queryParameters["regionCode"] = autorest.Encode("query", regionCode)
	} else {
		queryParameters["regionCode"] = autorest.Encode("query", "FKR")
	}
	if len(string(networkACLRuleTypeCode)) > 0 {
		queryParameters["networkAclRuleTypeCode"] = autorest.Encode("query", networkACLRuleTypeCode)
	}

	timestamp := strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)
	sec := security.NewSignature(client.Client.Secretkey, crypto.SHA256)
	signature, err := sec.Signature("GET", autorest.GetPath(DefaultBaseURI, "/getNetworkAclRuleList")+"?"+autorest.GetQuery(queryParameters), client.Client.AccessKey, timestamp)
	if err != nil {
		return nil, err
	}
	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/getNetworkAclRuleList"),
		autorest.WithQueryParameters(queryParameters),
		autorest.WithHeader("x-ncp-apigw-timestamp", timestamp),
		autorest.WithHeader("x-ncp-iam-access-key", client.Client.AccessKey),
		autorest.WithHeader("x-ncp-apigw-signature-v2", signature))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetRuleListSender sends the GetRuleList request. The method will close the
// http.Response Body if it receives an error.
func (client NetworkACLClient) GetRuleListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetRuleListResponder handles the response to the GetRuleList request. The method always
// closes the http.Response Body.
func (client NetworkACLClient) GetRuleListResponder(resp *http.Response) (result NetworkACLRuleListResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// RemoveInboundRule 네트워크 ACL의 Inbound Rule을 제거
// Parameters:
// networkACLNo - 네트워크 ACL 번호
// networkACLRuleListNpriority - 우선순위
// networkACLRuleListNprotocolTypeCode - 네트워크 ACL Rule 유형 코드
// networkACLRuleListNipBlock - IP 블록
// networkACLRuleListNruleActionCode - rule 액션 코드
// regionCode - REGION 코드
// networkACLRuleListNportRange - 포트 범위
func (client NetworkACLClient) RemoveInboundRule(ctx context.Context, networkACLNo string, networkACLRuleListNpriority string, networkACLRuleListNprotocolTypeCode ProtocolTypeCode, networkACLRuleListNipBlock string, networkACLRuleListNruleActionCode RuleActionCode, regionCode string, networkACLRuleListNportRange string) (result NetworkACLInboundRuleResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/NetworkACLClient.RemoveInboundRule")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.RemoveInboundRulePreparer(ctx, networkACLNo, networkACLRuleListNpriority, networkACLRuleListNprotocolTypeCode, networkACLRuleListNipBlock, networkACLRuleListNruleActionCode, regionCode, networkACLRuleListNportRange)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vpc.NetworkACLClient", "RemoveInboundRule", nil, "Failure preparing request")
		return
	}

	resp, err := client.RemoveInboundRuleSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "vpc.NetworkACLClient", "RemoveInboundRule", resp, "Failure sending request")
		return
	}

	result, err = client.RemoveInboundRuleResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vpc.NetworkACLClient", "RemoveInboundRule", resp, "Failure responding to request")
	}

	return
}

// RemoveInboundRulePreparer prepares the RemoveInboundRule request.
func (client NetworkACLClient) RemoveInboundRulePreparer(ctx context.Context, networkACLNo string, networkACLRuleListNpriority string, networkACLRuleListNprotocolTypeCode ProtocolTypeCode, networkACLRuleListNipBlock string, networkACLRuleListNruleActionCode RuleActionCode, regionCode string, networkACLRuleListNportRange string) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"networkAclNo":                          autorest.Encode("query", networkACLNo),
		"networkAclRuleList.N.ipBlock":          autorest.Encode("query", networkACLRuleListNipBlock),
		"networkAclRuleList.N.priority":         autorest.Encode("query", networkACLRuleListNpriority),
		"networkAclRuleList.N.protocolTypeCode": autorest.Encode("query", networkACLRuleListNprotocolTypeCode),
		"networkAclRuleList.N.ruleActionCode":   autorest.Encode("query", networkACLRuleListNruleActionCode),
		"responseFormatType":                    autorest.Encode("query", "json"),
	}
	if len(regionCode) > 0 {
		queryParameters["regionCode"] = autorest.Encode("query", regionCode)
	} else {
		queryParameters["regionCode"] = autorest.Encode("query", "FKR")
	}
	if len(networkACLRuleListNportRange) > 0 {
		queryParameters["networkAclRuleList.N.portRange"] = autorest.Encode("query", networkACLRuleListNportRange)
	}

	timestamp := strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)
	sec := security.NewSignature(client.Client.Secretkey, crypto.SHA256)
	signature, err := sec.Signature("POST", autorest.GetPath(DefaultBaseURI, "/removeNetworkAclInboundRule")+"?"+autorest.GetQuery(queryParameters), client.Client.AccessKey, timestamp)
	if err != nil {
		return nil, err
	}
	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/removeNetworkAclInboundRule"),
		autorest.WithQueryParameters(queryParameters),
		autorest.WithHeader("x-ncp-apigw-timestamp", timestamp),
		autorest.WithHeader("x-ncp-iam-access-key", client.Client.AccessKey),
		autorest.WithHeader("x-ncp-apigw-signature-v2", signature))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// RemoveInboundRuleSender sends the RemoveInboundRule request. The method will close the
// http.Response Body if it receives an error.
func (client NetworkACLClient) RemoveInboundRuleSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// RemoveInboundRuleResponder handles the response to the RemoveInboundRule request. The method always
// closes the http.Response Body.
func (client NetworkACLClient) RemoveInboundRuleResponder(resp *http.Response) (result NetworkACLInboundRuleResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// RemoveOutboundRule 네트워크 ACL의 Outbound Rule을 제거
// Parameters:
// networkACLNo - 네트워크 ACL 번호
// networkACLRuleListNpriority - 우선순위
// networkACLRuleListNprotocolTypeCode - 네트워크 ACL Rule 유형 코드
// networkACLRuleListNipBlock - IP 블록
// networkACLRuleListNruleActionCode - rule 액션 코드
// regionCode - REGION 코드
// networkACLRuleListNportRange - 포트 범위
func (client NetworkACLClient) RemoveOutboundRule(ctx context.Context, networkACLNo string, networkACLRuleListNpriority string, networkACLRuleListNprotocolTypeCode ProtocolTypeCode, networkACLRuleListNipBlock string, networkACLRuleListNruleActionCode RuleActionCode, regionCode string, networkACLRuleListNportRange string) (result NetworkACLOutboundRuleResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/NetworkACLClient.RemoveOutboundRule")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.RemoveOutboundRulePreparer(ctx, networkACLNo, networkACLRuleListNpriority, networkACLRuleListNprotocolTypeCode, networkACLRuleListNipBlock, networkACLRuleListNruleActionCode, regionCode, networkACLRuleListNportRange)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vpc.NetworkACLClient", "RemoveOutboundRule", nil, "Failure preparing request")
		return
	}

	resp, err := client.RemoveOutboundRuleSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "vpc.NetworkACLClient", "RemoveOutboundRule", resp, "Failure sending request")
		return
	}

	result, err = client.RemoveOutboundRuleResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vpc.NetworkACLClient", "RemoveOutboundRule", resp, "Failure responding to request")
	}

	return
}

// RemoveOutboundRulePreparer prepares the RemoveOutboundRule request.
func (client NetworkACLClient) RemoveOutboundRulePreparer(ctx context.Context, networkACLNo string, networkACLRuleListNpriority string, networkACLRuleListNprotocolTypeCode ProtocolTypeCode, networkACLRuleListNipBlock string, networkACLRuleListNruleActionCode RuleActionCode, regionCode string, networkACLRuleListNportRange string) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"networkAclNo":                          autorest.Encode("query", networkACLNo),
		"networkAclRuleList.N.ipBlock":          autorest.Encode("query", networkACLRuleListNipBlock),
		"networkAclRuleList.N.priority":         autorest.Encode("query", networkACLRuleListNpriority),
		"networkAclRuleList.N.protocolTypeCode": autorest.Encode("query", networkACLRuleListNprotocolTypeCode),
		"networkAclRuleList.N.ruleActionCode":   autorest.Encode("query", networkACLRuleListNruleActionCode),
		"responseFormatType":                    autorest.Encode("query", "json"),
	}
	if len(regionCode) > 0 {
		queryParameters["regionCode"] = autorest.Encode("query", regionCode)
	} else {
		queryParameters["regionCode"] = autorest.Encode("query", "FKR")
	}
	if len(networkACLRuleListNportRange) > 0 {
		queryParameters["networkAclRuleList.N.portRange"] = autorest.Encode("query", networkACLRuleListNportRange)
	}

	timestamp := strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)
	sec := security.NewSignature(client.Client.Secretkey, crypto.SHA256)
	signature, err := sec.Signature("POST", autorest.GetPath(DefaultBaseURI, "/removeNetworkAclOutboundRule")+"?"+autorest.GetQuery(queryParameters), client.Client.AccessKey, timestamp)
	if err != nil {
		return nil, err
	}
	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/removeNetworkAclOutboundRule"),
		autorest.WithQueryParameters(queryParameters),
		autorest.WithHeader("x-ncp-apigw-timestamp", timestamp),
		autorest.WithHeader("x-ncp-iam-access-key", client.Client.AccessKey),
		autorest.WithHeader("x-ncp-apigw-signature-v2", signature))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// RemoveOutboundRuleSender sends the RemoveOutboundRule request. The method will close the
// http.Response Body if it receives an error.
func (client NetworkACLClient) RemoveOutboundRuleSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// RemoveOutboundRuleResponder handles the response to the RemoveOutboundRule request. The method always
// closes the http.Response Body.
func (client NetworkACLClient) RemoveOutboundRuleResponder(resp *http.Response) (result NetworkACLOutboundRuleResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// SetSubnet 서브넷의 네트워크 ACL을 설정
// Parameters:
// networkACLNo - 네트워크 ACL 번호
// subnetNo - 서브넷 번호
// regionCode - REGION 코드
func (client NetworkACLClient) SetSubnet(ctx context.Context, networkACLNo string, subnetNo string, regionCode string) (result SubnetNetworkACLResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/NetworkACLClient.SetSubnet")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.SetSubnetPreparer(ctx, networkACLNo, subnetNo, regionCode)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vpc.NetworkACLClient", "SetSubnet", nil, "Failure preparing request")
		return
	}

	resp, err := client.SetSubnetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "vpc.NetworkACLClient", "SetSubnet", resp, "Failure sending request")
		return
	}

	result, err = client.SetSubnetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vpc.NetworkACLClient", "SetSubnet", resp, "Failure responding to request")
	}

	return
}

// SetSubnetPreparer prepares the SetSubnet request.
func (client NetworkACLClient) SetSubnetPreparer(ctx context.Context, networkACLNo string, subnetNo string, regionCode string) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"networkAclNo":       autorest.Encode("query", networkACLNo),
		"responseFormatType": autorest.Encode("query", "json"),
		"subnetNo":           autorest.Encode("query", subnetNo),
	}
	if len(regionCode) > 0 {
		queryParameters["regionCode"] = autorest.Encode("query", regionCode)
	} else {
		queryParameters["regionCode"] = autorest.Encode("query", "FKR")
	}

	timestamp := strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)
	sec := security.NewSignature(client.Client.Secretkey, crypto.SHA256)
	signature, err := sec.Signature("POST", autorest.GetPath(DefaultBaseURI, "/setSubnetNetworkAcl")+"?"+autorest.GetQuery(queryParameters), client.Client.AccessKey, timestamp)
	if err != nil {
		return nil, err
	}
	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/setSubnetNetworkAcl"),
		autorest.WithQueryParameters(queryParameters),
		autorest.WithHeader("x-ncp-apigw-timestamp", timestamp),
		autorest.WithHeader("x-ncp-iam-access-key", client.Client.AccessKey),
		autorest.WithHeader("x-ncp-apigw-signature-v2", signature))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// SetSubnetSender sends the SetSubnet request. The method will close the
// http.Response Body if it receives an error.
func (client NetworkACLClient) SetSubnetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// SetSubnetResponder handles the response to the SetSubnet request. The method always
// closes the http.Response Body.
func (client NetworkACLClient) SetSubnetResponder(resp *http.Response) (result SubnetNetworkACLResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
