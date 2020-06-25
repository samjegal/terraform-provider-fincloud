package loadbalancer

// FINCLOUD_APACHE_NO_VERSION

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// Client is the loadBalancer Client
type Client struct {
	BaseClient
}

// NewClient creates an instance of the Client client.
func NewClient() Client {
	return NewClientWithBaseURI(DefaultBaseURI)
}

// NewClientWithBaseURI creates an instance of the Client client.
func NewClientWithBaseURI(baseURI string) Client {
	return Client{NewWithBaseURI(baseURI)}
}

// CheckName 로드밸런서 이름 적합성 검사
// Parameters:
// loadBalancerName - VPC 번호
func (client Client) CheckName(ctx context.Context, loadBalancerName string) (result CheckNameParameter, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/Client.CheckName")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CheckNamePreparer(ctx, loadBalancerName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "loadbalancer.Client", "CheckName", nil, "Failure preparing request")
		return
	}

	resp, err := client.CheckNameSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "loadbalancer.Client", "CheckName", resp, "Failure sending request")
		return
	}

	result, err = client.CheckNameResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "loadbalancer.Client", "CheckName", resp, "Failure responding to request")
	}

	return
}

// CheckNamePreparer prepares the CheckName request.
func (client Client) CheckNamePreparer(ctx context.Context, loadBalancerName string) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"loadBalancerName": autorest.Encode("query", loadBalancerName),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/vpc-network/api/network/v1/load-balancers/check-load-balancer-name"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CheckNameSender sends the CheckName request. The method will close the
// http.Response Body if it receives an error.
func (client Client) CheckNameSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// CheckNameResponder handles the response to the CheckName request. The method always
// closes the http.Response Body.
func (client Client) CheckNameResponder(resp *http.Response) (result CheckNameParameter, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Create 로드밸런서 생성
// Parameters:
// parameters - 로드밸런서 생성 데이터
func (client Client) Create(ctx context.Context, parameters InstanceParameter) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/Client.Create")
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
		err = autorest.NewErrorWithError(err, "loadbalancer.Client", "Create", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "loadbalancer.Client", "Create", resp, "Failure sending request")
		return
	}

	result, err = client.CreateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "loadbalancer.Client", "Create", resp, "Failure responding to request")
	}

	return
}

// CreatePreparer prepares the Create request.
func (client Client) CreatePreparer(ctx context.Context, parameters InstanceParameter) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json;charset=UTF-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/vpc-network/api/network/v1/load-balancers/instances"),
		autorest.WithJSON(parameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client Client) CreateSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client Client) CreateResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Delete 로드밸런서 삭제
// Parameters:
// parameters - 로드밸런서 번호 정보
func (client Client) Delete(ctx context.Context, parameters InstanceListParameter) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/Client.Delete")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "loadbalancer.Client", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "loadbalancer.Client", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "loadbalancer.Client", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client Client) DeletePreparer(ctx context.Context, parameters InstanceListParameter) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json;charset=UTF-8"),
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/vpc-network/api/network/v1/load-balancers/instances"),
		autorest.WithJSON(parameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client Client) DeleteSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client Client) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Search 로드밸런서 정보 리스트
// Parameters:
// parameters - 로드밸런서 검색 데이터
func (client Client) Search(ctx context.Context, parameters SearchParameter) (result SearchListParameter, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/Client.Search")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.SearchPreparer(ctx, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "loadbalancer.Client", "Search", nil, "Failure preparing request")
		return
	}

	resp, err := client.SearchSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "loadbalancer.Client", "Search", resp, "Failure sending request")
		return
	}

	result, err = client.SearchResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "loadbalancer.Client", "Search", resp, "Failure responding to request")
	}

	return
}

// SearchPreparer prepares the Search request.
func (client Client) SearchPreparer(ctx context.Context, parameters SearchParameter) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json;charset=UTF-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/vpc-network/api/network/v1/load-balancers/instances/search"),
		autorest.WithJSON(parameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// SearchSender sends the Search request. The method will close the
// http.Response Body if it receives an error.
func (client Client) SearchSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// SearchResponder handles the response to the Search request. The method always
// closes the http.Response Body.
func (client Client) SearchResponder(resp *http.Response) (result SearchListParameter, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ServerInstance 로드밸런서 적용 서버 정보
// Parameters:
// vpcNo - VPC 번호
// layerTypeCode - VPC 번호
func (client Client) ServerInstance(ctx context.Context, vpcNo string, layerTypeCode string) (result ServerInstanceListParameter, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/Client.ServerInstance")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ServerInstancePreparer(ctx, vpcNo, layerTypeCode)
	if err != nil {
		err = autorest.NewErrorWithError(err, "loadbalancer.Client", "ServerInstance", nil, "Failure preparing request")
		return
	}

	resp, err := client.ServerInstanceSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "loadbalancer.Client", "ServerInstance", resp, "Failure sending request")
		return
	}

	result, err = client.ServerInstanceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "loadbalancer.Client", "ServerInstance", resp, "Failure responding to request")
	}

	return
}

// ServerInstancePreparer prepares the ServerInstance request.
func (client Client) ServerInstancePreparer(ctx context.Context, vpcNo string, layerTypeCode string) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"layerTypeCode": autorest.Encode("query", layerTypeCode),
		"vpcNo":         autorest.Encode("query", vpcNo),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/vpc-network/api/network/v1/load-balancers/server-instances"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ServerInstanceSender sends the ServerInstance request. The method will close the
// http.Response Body if it receives an error.
func (client Client) ServerInstanceSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// ServerInstanceResponder handles the response to the ServerInstance request. The method always
// closes the http.Response Body.
func (client Client) ServerInstanceResponder(resp *http.Response) (result ServerInstanceListParameter, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Update 로드밸런서 설정 변경
// Parameters:
// parameters - 로드밸런서 설정 데이터
func (client Client) Update(ctx context.Context, parameters SettingParameter) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/Client.Update")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.UpdatePreparer(ctx, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "loadbalancer.Client", "Update", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "loadbalancer.Client", "Update", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "loadbalancer.Client", "Update", resp, "Failure responding to request")
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client Client) UpdatePreparer(ctx context.Context, parameters SettingParameter) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json;charset=UTF-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/vpc-network/api/network/v1/load-balancers/instances/{lbInstanceNo}"),
		autorest.WithJSON(parameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client Client) UpdateSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client Client) UpdateResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// ZoneSubnet 로드밸런서 금융존 서브넷
// Parameters:
// vpcNo - VPC 번호
func (client Client) ZoneSubnet(ctx context.Context, vpcNo string) (result ZoneSubnetParameter, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/Client.ZoneSubnet")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ZoneSubnetPreparer(ctx, vpcNo)
	if err != nil {
		err = autorest.NewErrorWithError(err, "loadbalancer.Client", "ZoneSubnet", nil, "Failure preparing request")
		return
	}

	resp, err := client.ZoneSubnetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "loadbalancer.Client", "ZoneSubnet", resp, "Failure sending request")
		return
	}

	result, err = client.ZoneSubnetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "loadbalancer.Client", "ZoneSubnet", resp, "Failure responding to request")
	}

	return
}

// ZoneSubnetPreparer prepares the ZoneSubnet request.
func (client Client) ZoneSubnetPreparer(ctx context.Context, vpcNo string) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"vpcNo": autorest.Encode("query", vpcNo),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/vpc-network/api/network/v1/load-balancers/zone-subnets"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ZoneSubnetSender sends the ZoneSubnet request. The method will close the
// http.Response Body if it receives an error.
func (client Client) ZoneSubnetSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// ZoneSubnetResponder handles the response to the ZoneSubnet request. The method always
// closes the http.Response Body.
func (client Client) ZoneSubnetResponder(resp *http.Response) (result ZoneSubnetParameter, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
