# DefaultApi

All URIs are relative to *http://localhost*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**messengerWebMetricsPost**](DefaultApi.md#messengerWebMetricsPost) | **POST** /messenger/web/metrics | /messenger/web/metrics |
| [**tasksGet**](DefaultApi.md#tasksGet) | **GET** /tasks | /tasks |
| [**tasksPost**](DefaultApi.md#tasksPost) | **POST** /tasks | /tasks |
| [**v1MPost**](DefaultApi.md#v1MPost) | **POST** /v1/m | /v1/m |


<a id="messengerWebMetricsPost"></a>
# **messengerWebMetricsPost**
> messengerWebMetricsPost(appId, v, g, s, r, platform, installationType, idempotencyKey, internal, isIntersectionBooted, pageTitle, userActiveCompanyId, metrics, logs, opMetrics, hcMetrics, referer, anonymousSession, deviceIdentifier)

/messenger/web/metrics

**Host**: http://api-iam.intercom.io

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.models.*;
import org.openapitools.client.api.DefaultApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("http://localhost");

    DefaultApi apiInstance = new DefaultApi(defaultClient);
    String appId = "appId_example"; // String | 
    String v = "v_example"; // String | 
    String g = "g_example"; // String | 
    String s = "s_example"; // String | 
    String r = "r_example"; // String | 
    String platform = "platform_example"; // String | 
    String installationType = "installationType_example"; // String | 
    String idempotencyKey = "idempotencyKey_example"; // String | 
    String internal = "internal_example"; // String | 
    String isIntersectionBooted = "isIntersectionBooted_example"; // String | 
    String pageTitle = "pageTitle_example"; // String | 
    String userActiveCompanyId = "userActiveCompanyId_example"; // String | 
    String metrics = "metrics_example"; // String | 
    String logs = "logs_example"; // String | 
    String opMetrics = "opMetrics_example"; // String | 
    String hcMetrics = "hcMetrics_example"; // String | 
    String referer = "referer_example"; // String | 
    String anonymousSession = "anonymousSession_example"; // String | 
    String deviceIdentifier = "deviceIdentifier_example"; // String | 
    try {
      apiInstance.messengerWebMetricsPost(appId, v, g, s, r, platform, installationType, idempotencyKey, internal, isIntersectionBooted, pageTitle, userActiveCompanyId, metrics, logs, opMetrics, hcMetrics, referer, anonymousSession, deviceIdentifier);
    } catch (ApiException e) {
      System.err.println("Exception when calling DefaultApi#messengerWebMetricsPost");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **appId** | **String**|  | |
| **v** | **String**|  | |
| **g** | **String**|  | |
| **s** | **String**|  | |
| **r** | **String**|  | |
| **platform** | **String**|  | |
| **installationType** | **String**|  | |
| **idempotencyKey** | **String**|  | |
| **internal** | **String**|  | |
| **isIntersectionBooted** | **String**|  | |
| **pageTitle** | **String**|  | |
| **userActiveCompanyId** | **String**|  | |
| **metrics** | **String**|  | |
| **logs** | **String**|  | |
| **opMetrics** | **String**|  | |
| **hcMetrics** | **String**|  | |
| **referer** | **String**|  | |
| **anonymousSession** | **String**|  | |
| **deviceIdentifier** | **String**|  | |

### Return type

null (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: text/html

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** |  |  * status -  <br>  * x-ami-version -  <br>  * x-content-type-options -  <br>  * x-intercom-version -  <br>  * x-request-queueing -  <br>  * x-runtime -  <br>  |

<a id="tasksGet"></a>
# **tasksGet**
> List&lt;TasksGet200ResponseInner&gt; tasksGet(limit, q)

/tasks

**Host**: http://api.salesflare.com

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.DefaultApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("http://localhost");
    
    // Configure API key authorization: apikey cookie  hjsessionuser 374224
    ApiKeyAuth apikey cookie  hjsessionuser 374224 = (ApiKeyAuth) defaultClient.getAuthentication("apikey cookie  hjsessionuser 374224");
    apikey cookie  hjsessionuser 374224.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //apikey cookie  hjsessionuser 374224.setApiKeyPrefix("Token");

    // Configure API key authorization: apikey cookie  gid
    ApiKeyAuth apikey cookie  gid = (ApiKeyAuth) defaultClient.getAuthentication("apikey cookie  gid");
    apikey cookie  gid.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //apikey cookie  gid.setApiKeyPrefix("Token");

    // Configure API key authorization: apikey cookie  hjsession 374224
    ApiKeyAuth apikey cookie  hjsession 374224 = (ApiKeyAuth) defaultClient.getAuthentication("apikey cookie  hjsession 374224");
    apikey cookie  hjsession 374224.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //apikey cookie  hjsession 374224.setApiKeyPrefix("Token");

    // Configure API key authorization: apikey cookie ajs user id
    ApiKeyAuth apikey cookie ajs user id = (ApiKeyAuth) defaultClient.getAuthentication("apikey cookie ajs user id");
    apikey cookie ajs user id.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //apikey cookie ajs user id.setApiKeyPrefix("Token");

    // Configure API key authorization: apikey cookie intercom-id-nhqftro2
    ApiKeyAuth apikey cookie intercom-id-nhqftro2 = (ApiKeyAuth) defaultClient.getAuthentication("apikey cookie intercom-id-nhqftro2");
    apikey cookie intercom-id-nhqftro2.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //apikey cookie intercom-id-nhqftro2.setApiKeyPrefix("Token");

    // Configure API key authorization: apikey cookie intercom-session-nhqftro2
    ApiKeyAuth apikey cookie intercom-session-nhqftro2 = (ApiKeyAuth) defaultClient.getAuthentication("apikey cookie intercom-session-nhqftro2");
    apikey cookie intercom-session-nhqftro2.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //apikey cookie intercom-session-nhqftro2.setApiKeyPrefix("Token");

    // Configure API key authorization: apikey cookie intercom-device-id-nhqftro2
    ApiKeyAuth apikey cookie intercom-device-id-nhqftro2 = (ApiKeyAuth) defaultClient.getAuthentication("apikey cookie intercom-device-id-nhqftro2");
    apikey cookie intercom-device-id-nhqftro2.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //apikey cookie intercom-device-id-nhqftro2.setApiKeyPrefix("Token");

    // Configure API key authorization: apikey cookie salesflare-session
    ApiKeyAuth apikey cookie salesflare-session = (ApiKeyAuth) defaultClient.getAuthentication("apikey cookie salesflare-session");
    apikey cookie salesflare-session.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //apikey cookie salesflare-session.setApiKeyPrefix("Token");

    // Configure API key authorization: apikey cookie ajs anonymous id
    ApiKeyAuth apikey cookie ajs anonymous id = (ApiKeyAuth) defaultClient.getAuthentication("apikey cookie ajs anonymous id");
    apikey cookie ajs anonymous id.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //apikey cookie ajs anonymous id.setApiKeyPrefix("Token");

    // Configure API key authorization: apikey header cookie
    ApiKeyAuth apikey header cookie = (ApiKeyAuth) defaultClient.getAuthentication("apikey header cookie");
    apikey header cookie.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //apikey header cookie.setApiKeyPrefix("Token");

    DefaultApi apiInstance = new DefaultApi(defaultClient);
    String limit = "limit_example"; // String | 
    String q = "q_example"; // String | 
    try {
      List<TasksGet200ResponseInner> result = apiInstance.tasksGet(limit, q);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling DefaultApi#tasksGet");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **limit** | **String**|  | [optional] |
| **q** | **String**|  | [optional] |

### Return type

[**List&lt;TasksGet200ResponseInner&gt;**](TasksGet200ResponseInner.md)

### Authorization

[apikey cookie  hjsessionuser 374224](../README.md#apikey cookie  hjsessionuser 374224), [apikey cookie  gid](../README.md#apikey cookie  gid), [apikey cookie  hjsession 374224](../README.md#apikey cookie  hjsession 374224), [apikey cookie ajs user id](../README.md#apikey cookie ajs user id), [apikey cookie intercom-id-nhqftro2](../README.md#apikey cookie intercom-id-nhqftro2), [apikey cookie intercom-session-nhqftro2](../README.md#apikey cookie intercom-session-nhqftro2), [apikey cookie intercom-device-id-nhqftro2](../README.md#apikey cookie intercom-device-id-nhqftro2), [apikey cookie salesflare-session](../README.md#apikey cookie salesflare-session), [apikey cookie ajs anonymous id](../README.md#apikey cookie ajs anonymous id), [apikey header cookie](../README.md#apikey header cookie)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** |  |  -  |

<a id="tasksPost"></a>
# **tasksPost**
> TasksPost200Response tasksPost()

/tasks

**Host**: http://api.salesflare.com

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.auth.*;
import org.openapitools.client.models.*;
import org.openapitools.client.api.DefaultApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("http://localhost");
    
    // Configure API key authorization: apikey cookie  hjsessionuser 374224
    ApiKeyAuth apikey cookie  hjsessionuser 374224 = (ApiKeyAuth) defaultClient.getAuthentication("apikey cookie  hjsessionuser 374224");
    apikey cookie  hjsessionuser 374224.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //apikey cookie  hjsessionuser 374224.setApiKeyPrefix("Token");

    // Configure API key authorization: apikey cookie  gid
    ApiKeyAuth apikey cookie  gid = (ApiKeyAuth) defaultClient.getAuthentication("apikey cookie  gid");
    apikey cookie  gid.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //apikey cookie  gid.setApiKeyPrefix("Token");

    // Configure API key authorization: apikey cookie  hjsession 374224
    ApiKeyAuth apikey cookie  hjsession 374224 = (ApiKeyAuth) defaultClient.getAuthentication("apikey cookie  hjsession 374224");
    apikey cookie  hjsession 374224.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //apikey cookie  hjsession 374224.setApiKeyPrefix("Token");

    // Configure API key authorization: apikey cookie ajs user id
    ApiKeyAuth apikey cookie ajs user id = (ApiKeyAuth) defaultClient.getAuthentication("apikey cookie ajs user id");
    apikey cookie ajs user id.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //apikey cookie ajs user id.setApiKeyPrefix("Token");

    // Configure API key authorization: apikey cookie intercom-id-nhqftro2
    ApiKeyAuth apikey cookie intercom-id-nhqftro2 = (ApiKeyAuth) defaultClient.getAuthentication("apikey cookie intercom-id-nhqftro2");
    apikey cookie intercom-id-nhqftro2.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //apikey cookie intercom-id-nhqftro2.setApiKeyPrefix("Token");

    // Configure API key authorization: apikey cookie intercom-session-nhqftro2
    ApiKeyAuth apikey cookie intercom-session-nhqftro2 = (ApiKeyAuth) defaultClient.getAuthentication("apikey cookie intercom-session-nhqftro2");
    apikey cookie intercom-session-nhqftro2.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //apikey cookie intercom-session-nhqftro2.setApiKeyPrefix("Token");

    // Configure API key authorization: apikey cookie intercom-device-id-nhqftro2
    ApiKeyAuth apikey cookie intercom-device-id-nhqftro2 = (ApiKeyAuth) defaultClient.getAuthentication("apikey cookie intercom-device-id-nhqftro2");
    apikey cookie intercom-device-id-nhqftro2.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //apikey cookie intercom-device-id-nhqftro2.setApiKeyPrefix("Token");

    // Configure API key authorization: apikey cookie salesflare-session
    ApiKeyAuth apikey cookie salesflare-session = (ApiKeyAuth) defaultClient.getAuthentication("apikey cookie salesflare-session");
    apikey cookie salesflare-session.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //apikey cookie salesflare-session.setApiKeyPrefix("Token");

    // Configure API key authorization: apikey cookie ajs anonymous id
    ApiKeyAuth apikey cookie ajs anonymous id = (ApiKeyAuth) defaultClient.getAuthentication("apikey cookie ajs anonymous id");
    apikey cookie ajs anonymous id.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //apikey cookie ajs anonymous id.setApiKeyPrefix("Token");

    // Configure API key authorization: apikey header cookie
    ApiKeyAuth apikey header cookie = (ApiKeyAuth) defaultClient.getAuthentication("apikey header cookie");
    apikey header cookie.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //apikey header cookie.setApiKeyPrefix("Token");

    DefaultApi apiInstance = new DefaultApi(defaultClient);
    try {
      TasksPost200Response result = apiInstance.tasksPost();
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling DefaultApi#tasksPost");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
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

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** |  |  * x-cloud-trace-context -  <br>  |

<a id="v1MPost"></a>
# **v1MPost**
> V1MPost200Response v1MPost()

/v1/m

**Host**: http://api.segment.io

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.models.*;
import org.openapitools.client.api.DefaultApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("http://localhost");

    DefaultApi apiInstance = new DefaultApi(defaultClient);
    try {
      V1MPost200Response result = apiInstance.v1MPost();
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling DefaultApi#v1MPost");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
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

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** |  |  -  |

