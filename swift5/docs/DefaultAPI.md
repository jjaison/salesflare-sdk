# DefaultAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**messengerWebMetricsPost**](DefaultAPI.md#messengerwebmetricspost) | **POST** /messenger/web/metrics | /messenger/web/metrics
[**tasksGet**](DefaultAPI.md#tasksget) | **GET** /tasks | /tasks
[**tasksPost**](DefaultAPI.md#taskspost) | **POST** /tasks | /tasks
[**v1MPost**](DefaultAPI.md#v1mpost) | **POST** /v1/m | /v1/m


# **messengerWebMetricsPost**
```swift
    open class func messengerWebMetricsPost(appId: String, v: String, g: String, s: String, r: String, platform: String, installationType: String, idempotencyKey: String, _internal: String, isIntersectionBooted: String, pageTitle: String, userActiveCompanyId: String, metrics: String, logs: String, opMetrics: String, hcMetrics: String, referer: String, anonymousSession: String, deviceIdentifier: String, completion: @escaping (_ data: Void?, _ error: Error?) -> Void)
```

/messenger/web/metrics

**Host**: http://api-iam.intercom.io

### Example
```swift
// The following code samples are still beta. For any issue, please report via http://github.com/OpenAPITools/openapi-generator/issues/new
import OpenAPIClient

let appId = "appId_example" // String | 
let v = "v_example" // String | 
let g = "g_example" // String | 
let s = "s_example" // String | 
let r = "r_example" // String | 
let platform = "platform_example" // String | 
let installationType = "installationType_example" // String | 
let idempotencyKey = "idempotencyKey_example" // String | 
let _internal = "_internal_example" // String | 
let isIntersectionBooted = "isIntersectionBooted_example" // String | 
let pageTitle = "pageTitle_example" // String | 
let userActiveCompanyId = "userActiveCompanyId_example" // String | 
let metrics = "metrics_example" // String | 
let logs = "logs_example" // String | 
let opMetrics = "opMetrics_example" // String | 
let hcMetrics = "hcMetrics_example" // String | 
let referer = "referer_example" // String | 
let anonymousSession = "anonymousSession_example" // String | 
let deviceIdentifier = "deviceIdentifier_example" // String | 

// /messenger/web/metrics
DefaultAPI.messengerWebMetricsPost(appId: appId, v: v, g: g, s: s, r: r, platform: platform, installationType: installationType, idempotencyKey: idempotencyKey, _internal: _internal, isIntersectionBooted: isIntersectionBooted, pageTitle: pageTitle, userActiveCompanyId: userActiveCompanyId, metrics: metrics, logs: logs, opMetrics: opMetrics, hcMetrics: hcMetrics, referer: referer, anonymousSession: anonymousSession, deviceIdentifier: deviceIdentifier) { (response, error) in
    guard error == nil else {
        print(error)
        return
    }

    if (response) {
        dump(response)
    }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **appId** | **String** |  | 
 **v** | **String** |  | 
 **g** | **String** |  | 
 **s** | **String** |  | 
 **r** | **String** |  | 
 **platform** | **String** |  | 
 **installationType** | **String** |  | 
 **idempotencyKey** | **String** |  | 
 **_internal** | **String** |  | 
 **isIntersectionBooted** | **String** |  | 
 **pageTitle** | **String** |  | 
 **userActiveCompanyId** | **String** |  | 
 **metrics** | **String** |  | 
 **logs** | **String** |  | 
 **opMetrics** | **String** |  | 
 **hcMetrics** | **String** |  | 
 **referer** | **String** |  | 
 **anonymousSession** | **String** |  | 
 **deviceIdentifier** | **String** |  | 

### Return type

Void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **tasksGet**
```swift
    open class func tasksGet(limit: String? = nil, q: String? = nil, completion: @escaping (_ data: [TasksGet200ResponseInner]?, _ error: Error?) -> Void)
```

/tasks

**Host**: http://api.salesflare.com

### Example
```swift
// The following code samples are still beta. For any issue, please report via http://github.com/OpenAPITools/openapi-generator/issues/new
import OpenAPIClient

let limit = "limit_example" // String |  (optional)
let q = "q_example" // String |  (optional)

// /tasks
DefaultAPI.tasksGet(limit: limit, q: q) { (response, error) in
    guard error == nil else {
        print(error)
        return
    }

    if (response) {
        dump(response)
    }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **String** |  | [optional] 
 **q** | **String** |  | [optional] 

### Return type

[**[TasksGet200ResponseInner]**](TasksGet200ResponseInner.md)

### Authorization

[apikey cookie  hjsessionuser 374224](../README.md#apikey cookie  hjsessionuser 374224), [apikey cookie  gid](../README.md#apikey cookie  gid), [apikey cookie  hjsession 374224](../README.md#apikey cookie  hjsession 374224), [apikey cookie ajs user id](../README.md#apikey cookie ajs user id), [apikey cookie intercom-id-nhqftro2](../README.md#apikey cookie intercom-id-nhqftro2), [apikey cookie intercom-session-nhqftro2](../README.md#apikey cookie intercom-session-nhqftro2), [apikey cookie intercom-device-id-nhqftro2](../README.md#apikey cookie intercom-device-id-nhqftro2), [apikey cookie salesflare-session](../README.md#apikey cookie salesflare-session), [apikey cookie ajs anonymous id](../README.md#apikey cookie ajs anonymous id), [apikey header cookie](../README.md#apikey header cookie)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **tasksPost**
```swift
    open class func tasksPost(completion: @escaping (_ data: TasksPost200Response?, _ error: Error?) -> Void)
```

/tasks

**Host**: http://api.salesflare.com

### Example
```swift
// The following code samples are still beta. For any issue, please report via http://github.com/OpenAPITools/openapi-generator/issues/new
import OpenAPIClient


// /tasks
DefaultAPI.tasksPost() { (response, error) in
    guard error == nil else {
        print(error)
        return
    }

    if (response) {
        dump(response)
    }
}
```

### Parameters
This endpoint does not need any parameter.

### Return type

[**TasksPost200Response**](TasksPost200Response.md)

### Authorization

[apikey cookie  hjsessionuser 374224](../README.md#apikey cookie  hjsessionuser 374224), [apikey cookie  gid](../README.md#apikey cookie  gid), [apikey cookie  hjsession 374224](../README.md#apikey cookie  hjsession 374224), [apikey cookie ajs user id](../README.md#apikey cookie ajs user id), [apikey cookie intercom-id-nhqftro2](../README.md#apikey cookie intercom-id-nhqftro2), [apikey cookie intercom-session-nhqftro2](../README.md#apikey cookie intercom-session-nhqftro2), [apikey cookie intercom-device-id-nhqftro2](../README.md#apikey cookie intercom-device-id-nhqftro2), [apikey cookie salesflare-session](../README.md#apikey cookie salesflare-session), [apikey cookie ajs anonymous id](../README.md#apikey cookie ajs anonymous id), [apikey header cookie](../README.md#apikey header cookie)

### HTTP request headers

 - **Content-Type**: application/json;charset=UTF-8
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **v1MPost**
```swift
    open class func v1MPost(completion: @escaping (_ data: V1MPost200Response?, _ error: Error?) -> Void)
```

/v1/m

**Host**: http://api.segment.io

### Example
```swift
// The following code samples are still beta. For any issue, please report via http://github.com/OpenAPITools/openapi-generator/issues/new
import OpenAPIClient


// /v1/m
DefaultAPI.v1MPost() { (response, error) in
    guard error == nil else {
        print(error)
        return
    }

    if (response) {
        dump(response)
    }
}
```

### Parameters
This endpoint does not need any parameter.

### Return type

[**V1MPost200Response**](V1MPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

