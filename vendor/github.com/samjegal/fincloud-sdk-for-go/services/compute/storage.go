package compute

// FINCLOUD_APACHE_NO_VERSION

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// StorageClient is the compute Client
type StorageClient struct {
	BaseClient
}

// NewStorageClient creates an instance of the StorageClient client.
func NewStorageClient() StorageClient {
	return NewStorageClientWithBaseURI(DefaultBaseURI)
}

// NewStorageClientWithBaseURI creates an instance of the StorageClient client.
func NewStorageClientWithBaseURI(baseURI string) StorageClient {
	return StorageClient{NewWithBaseURI(baseURI)}
}

// Addable storage 적용 가능 여부
// Parameters:
// diskType2DetailCode - storage 디스크 타입 상세 코드
// zoneNo - zone 번호
// blockStorageInstanceNo - 블럭 Storage 인스턴스 번호
func (client StorageClient) Addable(ctx context.Context, diskType2DetailCode string, zoneNo string, blockStorageInstanceNo string) (result StorageAddableParameter, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/StorageClient.Addable")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.AddablePreparer(ctx, diskType2DetailCode, zoneNo, blockStorageInstanceNo)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.StorageClient", "Addable", nil, "Failure preparing request")
		return
	}

	resp, err := client.AddableSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "compute.StorageClient", "Addable", resp, "Failure sending request")
		return
	}

	result, err = client.AddableResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.StorageClient", "Addable", resp, "Failure responding to request")
	}

	return
}

// AddablePreparer prepares the Addable request.
func (client StorageClient) AddablePreparer(ctx context.Context, diskType2DetailCode string, zoneNo string, blockStorageInstanceNo string) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"blockStorageInstanceNo": autorest.Encode("query", blockStorageInstanceNo),
		"diskType2DetailCode":    autorest.Encode("query", diskType2DetailCode),
		"zoneNo":                 autorest.Encode("query", zoneNo),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/vpc-compute/api/compute/v1/storages/addable-servers"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// AddableSender sends the Addable request. The method will close the
// http.Response Body if it receives an error.
func (client StorageClient) AddableSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// AddableResponder handles the response to the Addable request. The method always
// closes the http.Response Body.
func (client StorageClient) AddableResponder(resp *http.Response) (result StorageAddableParameter, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Create storage 생성
func (client StorageClient) Create(ctx context.Context, parameters StorageParameter) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/StorageClient.Create")
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
		err = autorest.NewErrorWithError(err, "compute.StorageClient", "Create", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "compute.StorageClient", "Create", resp, "Failure sending request")
		return
	}

	result, err = client.CreateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.StorageClient", "Create", resp, "Failure responding to request")
	}

	return
}

// CreatePreparer prepares the Create request.
func (client StorageClient) CreatePreparer(ctx context.Context, parameters StorageParameter) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json;charset=UTF-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/vpc-compute/api/compute/v1/storages/instances"),
		autorest.WithJSON(parameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client StorageClient) CreateSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client StorageClient) CreateResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Delete storage 삭제
func (client StorageClient) Delete(ctx context.Context, parameters StorageDetachAndDeleteParameter) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/StorageClient.Delete")
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
		err = autorest.NewErrorWithError(err, "compute.StorageClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "compute.StorageClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.StorageClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client StorageClient) DeletePreparer(ctx context.Context, parameters StorageDetachAndDeleteParameter) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json;charset=UTF-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/vpc-compute/api/compute/v1/storages/detach-and-delete"),
		autorest.WithJSON(parameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client StorageClient) DeleteSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client StorageClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Detach storage 연결 해제
func (client StorageClient) Detach(ctx context.Context, parameters StorageDetachAndDeleteParameter) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/StorageClient.Detach")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DetachPreparer(ctx, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.StorageClient", "Detach", nil, "Failure preparing request")
		return
	}

	resp, err := client.DetachSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "compute.StorageClient", "Detach", resp, "Failure sending request")
		return
	}

	result, err = client.DetachResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.StorageClient", "Detach", resp, "Failure responding to request")
	}

	return
}

// DetachPreparer prepares the Detach request.
func (client StorageClient) DetachPreparer(ctx context.Context, parameters StorageDetachAndDeleteParameter) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json;charset=UTF-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/vpc-compute/api/compute/v1/storages/detach"),
		autorest.WithJSON(parameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DetachSender sends the Detach request. The method will close the
// http.Response Body if it receives an error.
func (client StorageClient) DetachSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// DetachResponder handles the response to the Detach request. The method always
// closes the http.Response Body.
func (client StorageClient) DetachResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// List storage 정보 리스트
// Parameters:
// pageNo - storage 리스트 페이지 번호
// pageSizeNo - storage 리스트 페이지 크기
// diskType2Code - storage 디스크 타입
// instanceStatusCode - storage 인스턴스 상태 코드
func (client StorageClient) List(ctx context.Context, pageNo string, pageSizeNo string, diskType2Code string, instanceStatusCode string) (result StorageListParameter, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/StorageClient.List")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ListPreparer(ctx, pageNo, pageSizeNo, diskType2Code, instanceStatusCode)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.StorageClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "compute.StorageClient", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.StorageClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client StorageClient) ListPreparer(ctx context.Context, pageNo string, pageSizeNo string, diskType2Code string, instanceStatusCode string) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"diskType2Code":      autorest.Encode("query", diskType2Code),
		"instanceStatusCode": autorest.Encode("query", instanceStatusCode),
		"pageNo":             autorest.Encode("query", pageNo),
		"pageSizeNo":         autorest.Encode("query", pageSizeNo),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/vpc-compute/api/compute/v1/storages/instances"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client StorageClient) ListSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client StorageClient) ListResponder(resp *http.Response) (result StorageListParameter, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
