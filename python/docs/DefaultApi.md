# salesflare.DefaultApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**messenger_web_metrics_post**](DefaultApi.md#messenger_web_metrics_post) | **POST** /messenger/web/metrics | /messenger/web/metrics
[**tasks_get**](DefaultApi.md#tasks_get) | **GET** /tasks | /tasks
[**tasks_post**](DefaultApi.md#tasks_post) | **POST** /tasks | /tasks
[**v1_m_post**](DefaultApi.md#v1_m_post) | **POST** /v1/m | /v1/m


# **messenger_web_metrics_post**
> messenger_web_metrics_post(app_id, v, g, s, r, platform, installation_type, idempotency_key, internal, is_intersection_booted, page_title, user_active_company_id, metrics, logs, op_metrics, hc_metrics, referer, anonymous_session, device_identifier)

/messenger/web/metrics

**Host**: http://api-iam.intercom.io

### Example


```python
import salesflare
from salesflare.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to http://localhost
# See configuration.py for a list of all supported configuration parameters.
configuration = salesflare.Configuration(
    host = "http://localhost"
)


# Enter a context with an instance of the API client
with salesflare.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = salesflare.DefaultApi(api_client)
    app_id = 'app_id_example' # str | 
    v = 'v_example' # str | 
    g = 'g_example' # str | 
    s = 's_example' # str | 
    r = 'r_example' # str | 
    platform = 'platform_example' # str | 
    installation_type = 'installation_type_example' # str | 
    idempotency_key = 'idempotency_key_example' # str | 
    internal = 'internal_example' # str | 
    is_intersection_booted = 'is_intersection_booted_example' # str | 
    page_title = 'page_title_example' # str | 
    user_active_company_id = 'user_active_company_id_example' # str | 
    metrics = 'metrics_example' # str | 
    logs = 'logs_example' # str | 
    op_metrics = 'op_metrics_example' # str | 
    hc_metrics = 'hc_metrics_example' # str | 
    referer = 'referer_example' # str | 
    anonymous_session = 'anonymous_session_example' # str | 
    device_identifier = 'device_identifier_example' # str | 

    try:
        # /messenger/web/metrics
        api_instance.messenger_web_metrics_post(app_id, v, g, s, r, platform, installation_type, idempotency_key, internal, is_intersection_booted, page_title, user_active_company_id, metrics, logs, op_metrics, hc_metrics, referer, anonymous_session, device_identifier)
    except Exception as e:
        print("Exception when calling DefaultApi->messenger_web_metrics_post: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **app_id** | **str**|  | 
 **v** | **str**|  | 
 **g** | **str**|  | 
 **s** | **str**|  | 
 **r** | **str**|  | 
 **platform** | **str**|  | 
 **installation_type** | **str**|  | 
 **idempotency_key** | **str**|  | 
 **internal** | **str**|  | 
 **is_intersection_booted** | **str**|  | 
 **page_title** | **str**|  | 
 **user_active_company_id** | **str**|  | 
 **metrics** | **str**|  | 
 **logs** | **str**|  | 
 **op_metrics** | **str**|  | 
 **hc_metrics** | **str**|  | 
 **referer** | **str**|  | 
 **anonymous_session** | **str**|  | 
 **device_identifier** | **str**|  | 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: text/html

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** |  |  * status -  <br>  * x-ami-version -  <br>  * x-content-type-options -  <br>  * x-intercom-version -  <br>  * x-request-queueing -  <br>  * x-runtime -  <br>  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **tasks_get**
> List[TasksGet200ResponseInner] tasks_get(limit=limit, q=q)

/tasks

**Host**: http://api.salesflare.com

### Example

* Api Key Authentication (apikey cookie  hjsessionuser 374224):
* Api Key Authentication (apikey cookie  gid):
* Api Key Authentication (apikey cookie  hjsession 374224):
* Api Key Authentication (apikey cookie ajs user id):
* Api Key Authentication (apikey cookie intercom-id-nhqftro2):
* Api Key Authentication (apikey cookie intercom-session-nhqftro2):
* Api Key Authentication (apikey cookie intercom-device-id-nhqftro2):
* Api Key Authentication (apikey cookie salesflare-session):
* Api Key Authentication (apikey cookie ajs anonymous id):
* Api Key Authentication (apikey header cookie):

```python
import salesflare
from salesflare.models.tasks_get200_response_inner import TasksGet200ResponseInner
from salesflare.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to http://localhost
# See configuration.py for a list of all supported configuration parameters.
configuration = salesflare.Configuration(
    host = "http://localhost"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure API key authorization: apikey cookie  hjsessionuser 374224
configuration.api_key['apikey cookie  hjsessionuser 374224'] = os.environ["API_KEY"]

# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['apikey cookie  hjsessionuser 374224'] = 'Bearer'

# Configure API key authorization: apikey cookie  gid
configuration.api_key['apikey cookie  gid'] = os.environ["API_KEY"]

# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['apikey cookie  gid'] = 'Bearer'

# Configure API key authorization: apikey cookie  hjsession 374224
configuration.api_key['apikey cookie  hjsession 374224'] = os.environ["API_KEY"]

# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['apikey cookie  hjsession 374224'] = 'Bearer'

# Configure API key authorization: apikey cookie ajs user id
configuration.api_key['apikey cookie ajs user id'] = os.environ["API_KEY"]

# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['apikey cookie ajs user id'] = 'Bearer'

# Configure API key authorization: apikey cookie intercom-id-nhqftro2
configuration.api_key['apikey cookie intercom-id-nhqftro2'] = os.environ["API_KEY"]

# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['apikey cookie intercom-id-nhqftro2'] = 'Bearer'

# Configure API key authorization: apikey cookie intercom-session-nhqftro2
configuration.api_key['apikey cookie intercom-session-nhqftro2'] = os.environ["API_KEY"]

# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['apikey cookie intercom-session-nhqftro2'] = 'Bearer'

# Configure API key authorization: apikey cookie intercom-device-id-nhqftro2
configuration.api_key['apikey cookie intercom-device-id-nhqftro2'] = os.environ["API_KEY"]

# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['apikey cookie intercom-device-id-nhqftro2'] = 'Bearer'

# Configure API key authorization: apikey cookie salesflare-session
configuration.api_key['apikey cookie salesflare-session'] = os.environ["API_KEY"]

# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['apikey cookie salesflare-session'] = 'Bearer'

# Configure API key authorization: apikey cookie ajs anonymous id
configuration.api_key['apikey cookie ajs anonymous id'] = os.environ["API_KEY"]

# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['apikey cookie ajs anonymous id'] = 'Bearer'

# Configure API key authorization: apikey header cookie
configuration.api_key['apikey header cookie'] = os.environ["API_KEY"]

# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['apikey header cookie'] = 'Bearer'

# Enter a context with an instance of the API client
with salesflare.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = salesflare.DefaultApi(api_client)
    limit = 'limit_example' # str |  (optional)
    q = 'q_example' # str |  (optional)

    try:
        # /tasks
        api_response = api_instance.tasks_get(limit=limit, q=q)
        print("The response of DefaultApi->tasks_get:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling DefaultApi->tasks_get: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **str**|  | [optional] 
 **q** | **str**|  | [optional] 

### Return type

[**List[TasksGet200ResponseInner]**](TasksGet200ResponseInner.md)

### Authorization

[apikey cookie  hjsessionuser 374224](../README.md#apikey cookie  hjsessionuser 374224), [apikey cookie  gid](../README.md#apikey cookie  gid), [apikey cookie  hjsession 374224](../README.md#apikey cookie  hjsession 374224), [apikey cookie ajs user id](../README.md#apikey cookie ajs user id), [apikey cookie intercom-id-nhqftro2](../README.md#apikey cookie intercom-id-nhqftro2), [apikey cookie intercom-session-nhqftro2](../README.md#apikey cookie intercom-session-nhqftro2), [apikey cookie intercom-device-id-nhqftro2](../README.md#apikey cookie intercom-device-id-nhqftro2), [apikey cookie salesflare-session](../README.md#apikey cookie salesflare-session), [apikey cookie ajs anonymous id](../README.md#apikey cookie ajs anonymous id), [apikey header cookie](../README.md#apikey header cookie)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** |  |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **tasks_post**
> TasksPost200Response tasks_post()

/tasks

**Host**: http://api.salesflare.com

### Example

* Api Key Authentication (apikey cookie  hjsessionuser 374224):
* Api Key Authentication (apikey cookie  gid):
* Api Key Authentication (apikey cookie  hjsession 374224):
* Api Key Authentication (apikey cookie ajs user id):
* Api Key Authentication (apikey cookie intercom-id-nhqftro2):
* Api Key Authentication (apikey cookie intercom-session-nhqftro2):
* Api Key Authentication (apikey cookie intercom-device-id-nhqftro2):
* Api Key Authentication (apikey cookie salesflare-session):
* Api Key Authentication (apikey cookie ajs anonymous id):
* Api Key Authentication (apikey header cookie):

```python
import salesflare
from salesflare.models.tasks_post200_response import TasksPost200Response
from salesflare.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to http://localhost
# See configuration.py for a list of all supported configuration parameters.
configuration = salesflare.Configuration(
    host = "http://localhost"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure API key authorization: apikey cookie  hjsessionuser 374224
configuration.api_key['apikey cookie  hjsessionuser 374224'] = os.environ["API_KEY"]

# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['apikey cookie  hjsessionuser 374224'] = 'Bearer'

# Configure API key authorization: apikey cookie  gid
configuration.api_key['apikey cookie  gid'] = os.environ["API_KEY"]

# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['apikey cookie  gid'] = 'Bearer'

# Configure API key authorization: apikey cookie  hjsession 374224
configuration.api_key['apikey cookie  hjsession 374224'] = os.environ["API_KEY"]

# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['apikey cookie  hjsession 374224'] = 'Bearer'

# Configure API key authorization: apikey cookie ajs user id
configuration.api_key['apikey cookie ajs user id'] = os.environ["API_KEY"]

# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['apikey cookie ajs user id'] = 'Bearer'

# Configure API key authorization: apikey cookie intercom-id-nhqftro2
configuration.api_key['apikey cookie intercom-id-nhqftro2'] = os.environ["API_KEY"]

# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['apikey cookie intercom-id-nhqftro2'] = 'Bearer'

# Configure API key authorization: apikey cookie intercom-session-nhqftro2
configuration.api_key['apikey cookie intercom-session-nhqftro2'] = os.environ["API_KEY"]

# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['apikey cookie intercom-session-nhqftro2'] = 'Bearer'

# Configure API key authorization: apikey cookie intercom-device-id-nhqftro2
configuration.api_key['apikey cookie intercom-device-id-nhqftro2'] = os.environ["API_KEY"]

# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['apikey cookie intercom-device-id-nhqftro2'] = 'Bearer'

# Configure API key authorization: apikey cookie salesflare-session
configuration.api_key['apikey cookie salesflare-session'] = os.environ["API_KEY"]

# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['apikey cookie salesflare-session'] = 'Bearer'

# Configure API key authorization: apikey cookie ajs anonymous id
configuration.api_key['apikey cookie ajs anonymous id'] = os.environ["API_KEY"]

# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['apikey cookie ajs anonymous id'] = 'Bearer'

# Configure API key authorization: apikey header cookie
configuration.api_key['apikey header cookie'] = os.environ["API_KEY"]

# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['apikey header cookie'] = 'Bearer'

# Enter a context with an instance of the API client
with salesflare.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = salesflare.DefaultApi(api_client)

    try:
        # /tasks
        api_response = api_instance.tasks_post()
        print("The response of DefaultApi->tasks_post:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling DefaultApi->tasks_post: %s\n" % e)
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
**200** |  |  * x-cloud-trace-context -  <br>  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **v1_m_post**
> V1MPost200Response v1_m_post()

/v1/m

**Host**: http://api.segment.io

### Example


```python
import salesflare
from salesflare.models.v1_m_post200_response import V1MPost200Response
from salesflare.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to http://localhost
# See configuration.py for a list of all supported configuration parameters.
configuration = salesflare.Configuration(
    host = "http://localhost"
)


# Enter a context with an instance of the API client
with salesflare.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = salesflare.DefaultApi(api_client)

    try:
        # /v1/m
        api_response = api_instance.v1_m_post()
        print("The response of DefaultApi->v1_m_post:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling DefaultApi->v1_m_post: %s\n" % e)
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
**200** |  |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

