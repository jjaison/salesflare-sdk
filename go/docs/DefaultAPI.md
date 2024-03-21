# \DefaultAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MessengerWebMetricsPost**](DefaultAPI.md#MessengerWebMetricsPost) | **Post** /messenger/web/metrics | /messenger/web/metrics
[**TasksGet**](DefaultAPI.md#TasksGet) | **Get** /tasks | /tasks
[**TasksPost**](DefaultAPI.md#TasksPost) | **Post** /tasks | /tasks
[**V1MPost**](DefaultAPI.md#V1MPost) | **Post** /v1/m | /v1/m



## MessengerWebMetricsPost

> MessengerWebMetricsPost(ctx).AppId(appId).V(v).G(g).S(s).R(r).Platform(platform).InstallationType(installationType).IdempotencyKey(idempotencyKey).Internal(internal).IsIntersectionBooted(isIntersectionBooted).PageTitle(pageTitle).UserActiveCompanyId(userActiveCompanyId).Metrics(metrics).Logs(logs).OpMetrics(opMetrics).HcMetrics(hcMetrics).Referer(referer).AnonymousSession(anonymousSession).DeviceIdentifier(deviceIdentifier).Execute()

/messenger/web/metrics



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	appId := "appId_example" // string | 
	v := "v_example" // string | 
	g := "g_example" // string | 
	s := "s_example" // string | 
	r := "r_example" // string | 
	platform := "platform_example" // string | 
	installationType := "installationType_example" // string | 
	idempotencyKey := "idempotencyKey_example" // string | 
	internal := "internal_example" // string | 
	isIntersectionBooted := "isIntersectionBooted_example" // string | 
	pageTitle := "pageTitle_example" // string | 
	userActiveCompanyId := "userActiveCompanyId_example" // string | 
	metrics := "metrics_example" // string | 
	logs := "logs_example" // string | 
	opMetrics := "opMetrics_example" // string | 
	hcMetrics := "hcMetrics_example" // string | 
	referer := "referer_example" // string | 
	anonymousSession := "anonymousSession_example" // string | 
	deviceIdentifier := "deviceIdentifier_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.MessengerWebMetricsPost(context.Background()).AppId(appId).V(v).G(g).S(s).R(r).Platform(platform).InstallationType(installationType).IdempotencyKey(idempotencyKey).Internal(internal).IsIntersectionBooted(isIntersectionBooted).PageTitle(pageTitle).UserActiveCompanyId(userActiveCompanyId).Metrics(metrics).Logs(logs).OpMetrics(opMetrics).HcMetrics(hcMetrics).Referer(referer).AnonymousSession(anonymousSession).DeviceIdentifier(deviceIdentifier).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.MessengerWebMetricsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMessengerWebMetricsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **appId** | **string** |  | 
 **v** | **string** |  | 
 **g** | **string** |  | 
 **s** | **string** |  | 
 **r** | **string** |  | 
 **platform** | **string** |  | 
 **installationType** | **string** |  | 
 **idempotencyKey** | **string** |  | 
 **internal** | **string** |  | 
 **isIntersectionBooted** | **string** |  | 
 **pageTitle** | **string** |  | 
 **userActiveCompanyId** | **string** |  | 
 **metrics** | **string** |  | 
 **logs** | **string** |  | 
 **opMetrics** | **string** |  | 
 **hcMetrics** | **string** |  | 
 **referer** | **string** |  | 
 **anonymousSession** | **string** |  | 
 **deviceIdentifier** | **string** |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TasksGet

> []TasksGet200ResponseInner TasksGet(ctx).Limit(limit).Q(q).Execute()

/tasks



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	limit := "limit_example" // string |  (optional)
	q := "q_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.TasksGet(context.Background()).Limit(limit).Q(q).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.TasksGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TasksGet`: []TasksGet200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.TasksGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTasksGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **string** |  | 
 **q** | **string** |  | 

### Return type

[**[]TasksGet200ResponseInner**](TasksGet200ResponseInner.md)

### Authorization

[apikey cookie  hjsessionuser 374224](../README.md#apikey cookie  hjsessionuser 374224), [apikey cookie  gid](../README.md#apikey cookie  gid), [apikey cookie  hjsession 374224](../README.md#apikey cookie  hjsession 374224), [apikey cookie ajs user id](../README.md#apikey cookie ajs user id), [apikey cookie intercom-id-nhqftro2](../README.md#apikey cookie intercom-id-nhqftro2), [apikey cookie intercom-session-nhqftro2](../README.md#apikey cookie intercom-session-nhqftro2), [apikey cookie intercom-device-id-nhqftro2](../README.md#apikey cookie intercom-device-id-nhqftro2), [apikey cookie salesflare-session](../README.md#apikey cookie salesflare-session), [apikey cookie ajs anonymous id](../README.md#apikey cookie ajs anonymous id), [apikey header cookie](../README.md#apikey header cookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TasksPost

> TasksPost200Response TasksPost(ctx).Execute()

/tasks



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.TasksPost(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.TasksPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TasksPost`: TasksPost200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.TasksPost`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiTasksPostRequest struct via the builder pattern


### Return type

[**TasksPost200Response**](TasksPost200Response.md)

### Authorization

[apikey cookie  hjsessionuser 374224](../README.md#apikey cookie  hjsessionuser 374224), [apikey cookie  gid](../README.md#apikey cookie  gid), [apikey cookie  hjsession 374224](../README.md#apikey cookie  hjsession 374224), [apikey cookie ajs user id](../README.md#apikey cookie ajs user id), [apikey cookie intercom-id-nhqftro2](../README.md#apikey cookie intercom-id-nhqftro2), [apikey cookie intercom-session-nhqftro2](../README.md#apikey cookie intercom-session-nhqftro2), [apikey cookie intercom-device-id-nhqftro2](../README.md#apikey cookie intercom-device-id-nhqftro2), [apikey cookie salesflare-session](../README.md#apikey cookie salesflare-session), [apikey cookie ajs anonymous id](../README.md#apikey cookie ajs anonymous id), [apikey header cookie](../README.md#apikey header cookie)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1MPost

> V1MPost200Response V1MPost(ctx).Execute()

/v1/m



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.V1MPost(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.V1MPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1MPost`: V1MPost200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.V1MPost`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1MPostRequest struct via the builder pattern


### Return type

[**V1MPost200Response**](V1MPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: text/plain
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

