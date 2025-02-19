/*
app.salesflare.com

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package salesflare

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)


// DefaultAPIService DefaultAPI service
type DefaultAPIService service

type ApiMessengerWebMetricsPostRequest struct {
	ctx context.Context
	ApiService *DefaultAPIService
	appId *string
	v *string
	g *string
	s *string
	r *string
	platform *string
	installationType *string
	idempotencyKey *string
	internal *string
	isIntersectionBooted *string
	pageTitle *string
	userActiveCompanyId *string
	metrics *string
	logs *string
	opMetrics *string
	hcMetrics *string
	referer *string
	anonymousSession *string
	deviceIdentifier *string
}

func (r ApiMessengerWebMetricsPostRequest) AppId(appId string) ApiMessengerWebMetricsPostRequest {
	r.appId = &appId
	return r
}

func (r ApiMessengerWebMetricsPostRequest) V(v string) ApiMessengerWebMetricsPostRequest {
	r.v = &v
	return r
}

func (r ApiMessengerWebMetricsPostRequest) G(g string) ApiMessengerWebMetricsPostRequest {
	r.g = &g
	return r
}

func (r ApiMessengerWebMetricsPostRequest) S(s string) ApiMessengerWebMetricsPostRequest {
	r.s = &s
	return r
}

func (r ApiMessengerWebMetricsPostRequest) R(r string) ApiMessengerWebMetricsPostRequest {
	r.r = &r
	return r
}

func (r ApiMessengerWebMetricsPostRequest) Platform(platform string) ApiMessengerWebMetricsPostRequest {
	r.platform = &platform
	return r
}

func (r ApiMessengerWebMetricsPostRequest) InstallationType(installationType string) ApiMessengerWebMetricsPostRequest {
	r.installationType = &installationType
	return r
}

func (r ApiMessengerWebMetricsPostRequest) IdempotencyKey(idempotencyKey string) ApiMessengerWebMetricsPostRequest {
	r.idempotencyKey = &idempotencyKey
	return r
}

func (r ApiMessengerWebMetricsPostRequest) Internal(internal string) ApiMessengerWebMetricsPostRequest {
	r.internal = &internal
	return r
}

func (r ApiMessengerWebMetricsPostRequest) IsIntersectionBooted(isIntersectionBooted string) ApiMessengerWebMetricsPostRequest {
	r.isIntersectionBooted = &isIntersectionBooted
	return r
}

func (r ApiMessengerWebMetricsPostRequest) PageTitle(pageTitle string) ApiMessengerWebMetricsPostRequest {
	r.pageTitle = &pageTitle
	return r
}

func (r ApiMessengerWebMetricsPostRequest) UserActiveCompanyId(userActiveCompanyId string) ApiMessengerWebMetricsPostRequest {
	r.userActiveCompanyId = &userActiveCompanyId
	return r
}

func (r ApiMessengerWebMetricsPostRequest) Metrics(metrics string) ApiMessengerWebMetricsPostRequest {
	r.metrics = &metrics
	return r
}

func (r ApiMessengerWebMetricsPostRequest) Logs(logs string) ApiMessengerWebMetricsPostRequest {
	r.logs = &logs
	return r
}

func (r ApiMessengerWebMetricsPostRequest) OpMetrics(opMetrics string) ApiMessengerWebMetricsPostRequest {
	r.opMetrics = &opMetrics
	return r
}

func (r ApiMessengerWebMetricsPostRequest) HcMetrics(hcMetrics string) ApiMessengerWebMetricsPostRequest {
	r.hcMetrics = &hcMetrics
	return r
}

func (r ApiMessengerWebMetricsPostRequest) Referer(referer string) ApiMessengerWebMetricsPostRequest {
	r.referer = &referer
	return r
}

func (r ApiMessengerWebMetricsPostRequest) AnonymousSession(anonymousSession string) ApiMessengerWebMetricsPostRequest {
	r.anonymousSession = &anonymousSession
	return r
}

func (r ApiMessengerWebMetricsPostRequest) DeviceIdentifier(deviceIdentifier string) ApiMessengerWebMetricsPostRequest {
	r.deviceIdentifier = &deviceIdentifier
	return r
}

func (r ApiMessengerWebMetricsPostRequest) Execute() (*http.Response, error) {
	return r.ApiService.MessengerWebMetricsPostExecute(r)
}

/*
MessengerWebMetricsPost /messenger/web/metrics

**Host**: http://api-iam.intercom.io

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiMessengerWebMetricsPostRequest
*/
func (a *DefaultAPIService) MessengerWebMetricsPost(ctx context.Context) ApiMessengerWebMetricsPostRequest {
	return ApiMessengerWebMetricsPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *DefaultAPIService) MessengerWebMetricsPostExecute(r ApiMessengerWebMetricsPostRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultAPIService.MessengerWebMetricsPost")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/messenger/web/metrics"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.appId == nil {
		return nil, reportError("appId is required and must be specified")
	}
	if r.v == nil {
		return nil, reportError("v is required and must be specified")
	}
	if r.g == nil {
		return nil, reportError("g is required and must be specified")
	}
	if r.s == nil {
		return nil, reportError("s is required and must be specified")
	}
	if r.r == nil {
		return nil, reportError("r is required and must be specified")
	}
	if r.platform == nil {
		return nil, reportError("platform is required and must be specified")
	}
	if r.installationType == nil {
		return nil, reportError("installationType is required and must be specified")
	}
	if r.idempotencyKey == nil {
		return nil, reportError("idempotencyKey is required and must be specified")
	}
	if r.internal == nil {
		return nil, reportError("internal is required and must be specified")
	}
	if r.isIntersectionBooted == nil {
		return nil, reportError("isIntersectionBooted is required and must be specified")
	}
	if r.pageTitle == nil {
		return nil, reportError("pageTitle is required and must be specified")
	}
	if r.userActiveCompanyId == nil {
		return nil, reportError("userActiveCompanyId is required and must be specified")
	}
	if r.metrics == nil {
		return nil, reportError("metrics is required and must be specified")
	}
	if r.logs == nil {
		return nil, reportError("logs is required and must be specified")
	}
	if r.opMetrics == nil {
		return nil, reportError("opMetrics is required and must be specified")
	}
	if r.hcMetrics == nil {
		return nil, reportError("hcMetrics is required and must be specified")
	}
	if r.referer == nil {
		return nil, reportError("referer is required and must be specified")
	}
	if r.anonymousSession == nil {
		return nil, reportError("anonymousSession is required and must be specified")
	}
	if r.deviceIdentifier == nil {
		return nil, reportError("deviceIdentifier is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/x-www-form-urlencoded"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"text/html"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarFormParams, "app_id", r.appId, "")
	parameterAddToHeaderOrQuery(localVarFormParams, "v", r.v, "")
	parameterAddToHeaderOrQuery(localVarFormParams, "g", r.g, "")
	parameterAddToHeaderOrQuery(localVarFormParams, "s", r.s, "")
	parameterAddToHeaderOrQuery(localVarFormParams, "r", r.r, "")
	parameterAddToHeaderOrQuery(localVarFormParams, "platform", r.platform, "")
	parameterAddToHeaderOrQuery(localVarFormParams, "installation_type", r.installationType, "")
	parameterAddToHeaderOrQuery(localVarFormParams, "Idempotency-Key", r.idempotencyKey, "")
	parameterAddToHeaderOrQuery(localVarFormParams, "internal", r.internal, "")
	parameterAddToHeaderOrQuery(localVarFormParams, "is_intersection_booted", r.isIntersectionBooted, "")
	parameterAddToHeaderOrQuery(localVarFormParams, "page_title", r.pageTitle, "")
	parameterAddToHeaderOrQuery(localVarFormParams, "user_active_company_id", r.userActiveCompanyId, "")
	parameterAddToHeaderOrQuery(localVarFormParams, "metrics", r.metrics, "")
	parameterAddToHeaderOrQuery(localVarFormParams, "logs", r.logs, "")
	parameterAddToHeaderOrQuery(localVarFormParams, "op_metrics", r.opMetrics, "")
	parameterAddToHeaderOrQuery(localVarFormParams, "hc_metrics", r.hcMetrics, "")
	parameterAddToHeaderOrQuery(localVarFormParams, "referer", r.referer, "")
	parameterAddToHeaderOrQuery(localVarFormParams, "anonymous_session", r.anonymousSession, "")
	parameterAddToHeaderOrQuery(localVarFormParams, "device_identifier", r.deviceIdentifier, "")
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiTasksGetRequest struct {
	ctx context.Context
	ApiService *DefaultAPIService
	limit *string
	q *string
}

func (r ApiTasksGetRequest) Limit(limit string) ApiTasksGetRequest {
	r.limit = &limit
	return r
}

func (r ApiTasksGetRequest) Q(q string) ApiTasksGetRequest {
	r.q = &q
	return r
}

func (r ApiTasksGetRequest) Execute() ([]TasksGet200ResponseInner, *http.Response, error) {
	return r.ApiService.TasksGetExecute(r)
}

/*
TasksGet /tasks

**Host**: http://api.salesflare.com

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiTasksGetRequest
*/
func (a *DefaultAPIService) TasksGet(ctx context.Context) ApiTasksGetRequest {
	return ApiTasksGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []TasksGet200ResponseInner
func (a *DefaultAPIService) TasksGetExecute(r ApiTasksGetRequest) ([]TasksGet200ResponseInner, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []TasksGet200ResponseInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultAPIService.TasksGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/tasks"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	}
	if r.q != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "q", r.q, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apikey header cookie"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["COOKIE"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiTasksPostRequest struct {
	ctx context.Context
	ApiService *DefaultAPIService
}

func (r ApiTasksPostRequest) Execute() (*TasksPost200Response, *http.Response, error) {
	return r.ApiService.TasksPostExecute(r)
}

/*
TasksPost /tasks

**Host**: http://api.salesflare.com

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiTasksPostRequest
*/
func (a *DefaultAPIService) TasksPost(ctx context.Context) ApiTasksPostRequest {
	return ApiTasksPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return TasksPost200Response
func (a *DefaultAPIService) TasksPostExecute(r ApiTasksPostRequest) (*TasksPost200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *TasksPost200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultAPIService.TasksPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/tasks"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json;charset=UTF-8"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apikey header cookie"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["COOKIE"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiV1MPostRequest struct {
	ctx context.Context
	ApiService *DefaultAPIService
}

func (r ApiV1MPostRequest) Execute() (*V1MPost200Response, *http.Response, error) {
	return r.ApiService.V1MPostExecute(r)
}

/*
V1MPost /v1/m

**Host**: http://api.segment.io

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiV1MPostRequest
*/
func (a *DefaultAPIService) V1MPost(ctx context.Context) ApiV1MPostRequest {
	return ApiV1MPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return V1MPost200Response
func (a *DefaultAPIService) V1MPostExecute(r ApiV1MPostRequest) (*V1MPost200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *V1MPost200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultAPIService.V1MPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/m"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"text/plain"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
