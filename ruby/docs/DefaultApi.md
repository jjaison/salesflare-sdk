# OpenapiClient::DefaultApi

All URIs are relative to *http://localhost*

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**messenger_web_metrics_post**](DefaultApi.md#messenger_web_metrics_post) | **POST** /messenger/web/metrics | /messenger/web/metrics |
| [**tasks_get**](DefaultApi.md#tasks_get) | **GET** /tasks | /tasks |
| [**tasks_post**](DefaultApi.md#tasks_post) | **POST** /tasks | /tasks |
| [**v1_m_post**](DefaultApi.md#v1_m_post) | **POST** /v1/m | /v1/m |


## messenger_web_metrics_post

> messenger_web_metrics_post(app_id, v, g, s, r, platform, installation_type, idempotency_key, internal, is_intersection_booted, page_title, user_active_company_id, metrics, logs, op_metrics, hc_metrics, referer, anonymous_session, device_identifier)

/messenger/web/metrics

**Host**: http://api-iam.intercom.io

### Examples

```ruby
require 'time'
require 'openapi_client'

api_instance = OpenapiClient::DefaultApi.new
app_id = 'app_id_example' # String | 
v = 'v_example' # String | 
g = 'g_example' # String | 
s = 's_example' # String | 
r = 'r_example' # String | 
platform = 'platform_example' # String | 
installation_type = 'installation_type_example' # String | 
idempotency_key = 'idempotency_key_example' # String | 
internal = 'internal_example' # String | 
is_intersection_booted = 'is_intersection_booted_example' # String | 
page_title = 'page_title_example' # String | 
user_active_company_id = 'user_active_company_id_example' # String | 
metrics = 'metrics_example' # String | 
logs = 'logs_example' # String | 
op_metrics = 'op_metrics_example' # String | 
hc_metrics = 'hc_metrics_example' # String | 
referer = 'referer_example' # String | 
anonymous_session = 'anonymous_session_example' # String | 
device_identifier = 'device_identifier_example' # String | 

begin
  # /messenger/web/metrics
  api_instance.messenger_web_metrics_post(app_id, v, g, s, r, platform, installation_type, idempotency_key, internal, is_intersection_booted, page_title, user_active_company_id, metrics, logs, op_metrics, hc_metrics, referer, anonymous_session, device_identifier)
rescue OpenapiClient::ApiError => e
  puts "Error when calling DefaultApi->messenger_web_metrics_post: #{e}"
end
```

#### Using the messenger_web_metrics_post_with_http_info variant

This returns an Array which contains the response data (`nil` in this case), status code and headers.

> <Array(nil, Integer, Hash)> messenger_web_metrics_post_with_http_info(app_id, v, g, s, r, platform, installation_type, idempotency_key, internal, is_intersection_booted, page_title, user_active_company_id, metrics, logs, op_metrics, hc_metrics, referer, anonymous_session, device_identifier)

```ruby
begin
  # /messenger/web/metrics
  data, status_code, headers = api_instance.messenger_web_metrics_post_with_http_info(app_id, v, g, s, r, platform, installation_type, idempotency_key, internal, is_intersection_booted, page_title, user_active_company_id, metrics, logs, op_metrics, hc_metrics, referer, anonymous_session, device_identifier)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => nil
rescue OpenapiClient::ApiError => e
  puts "Error when calling DefaultApi->messenger_web_metrics_post_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **app_id** | **String** |  |  |
| **v** | **String** |  |  |
| **g** | **String** |  |  |
| **s** | **String** |  |  |
| **r** | **String** |  |  |
| **platform** | **String** |  |  |
| **installation_type** | **String** |  |  |
| **idempotency_key** | **String** |  |  |
| **internal** | **String** |  |  |
| **is_intersection_booted** | **String** |  |  |
| **page_title** | **String** |  |  |
| **user_active_company_id** | **String** |  |  |
| **metrics** | **String** |  |  |
| **logs** | **String** |  |  |
| **op_metrics** | **String** |  |  |
| **hc_metrics** | **String** |  |  |
| **referer** | **String** |  |  |
| **anonymous_session** | **String** |  |  |
| **device_identifier** | **String** |  |  |

### Return type

nil (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: text/html


## tasks_get

> <Array<TasksGet200ResponseInner>> tasks_get(opts)

/tasks

**Host**: http://api.salesflare.com

### Examples

```ruby
require 'time'
require 'openapi_client'
# setup authorization
OpenapiClient.configure do |config|
  # Configure API key authorization: apikey cookie  hjsessionuser 374224
  config.api_key['apikey cookie  hjsessionuser 374224'] = 'YOUR API KEY'
  # Uncomment the following line to set a prefix for the API key, e.g. 'Bearer' (defaults to nil)
  # config.api_key_prefix['apikey cookie  hjsessionuser 374224'] = 'Bearer'

  # Configure API key authorization: apikey cookie  gid
  config.api_key['apikey cookie  gid'] = 'YOUR API KEY'
  # Uncomment the following line to set a prefix for the API key, e.g. 'Bearer' (defaults to nil)
  # config.api_key_prefix['apikey cookie  gid'] = 'Bearer'

  # Configure API key authorization: apikey cookie  hjsession 374224
  config.api_key['apikey cookie  hjsession 374224'] = 'YOUR API KEY'
  # Uncomment the following line to set a prefix for the API key, e.g. 'Bearer' (defaults to nil)
  # config.api_key_prefix['apikey cookie  hjsession 374224'] = 'Bearer'

  # Configure API key authorization: apikey cookie ajs user id
  config.api_key['apikey cookie ajs user id'] = 'YOUR API KEY'
  # Uncomment the following line to set a prefix for the API key, e.g. 'Bearer' (defaults to nil)
  # config.api_key_prefix['apikey cookie ajs user id'] = 'Bearer'

  # Configure API key authorization: apikey cookie intercom-id-nhqftro2
  config.api_key['apikey cookie intercom-id-nhqftro2'] = 'YOUR API KEY'
  # Uncomment the following line to set a prefix for the API key, e.g. 'Bearer' (defaults to nil)
  # config.api_key_prefix['apikey cookie intercom-id-nhqftro2'] = 'Bearer'

  # Configure API key authorization: apikey cookie intercom-session-nhqftro2
  config.api_key['apikey cookie intercom-session-nhqftro2'] = 'YOUR API KEY'
  # Uncomment the following line to set a prefix for the API key, e.g. 'Bearer' (defaults to nil)
  # config.api_key_prefix['apikey cookie intercom-session-nhqftro2'] = 'Bearer'

  # Configure API key authorization: apikey cookie intercom-device-id-nhqftro2
  config.api_key['apikey cookie intercom-device-id-nhqftro2'] = 'YOUR API KEY'
  # Uncomment the following line to set a prefix for the API key, e.g. 'Bearer' (defaults to nil)
  # config.api_key_prefix['apikey cookie intercom-device-id-nhqftro2'] = 'Bearer'

  # Configure API key authorization: apikey cookie salesflare-session
  config.api_key['apikey cookie salesflare-session'] = 'YOUR API KEY'
  # Uncomment the following line to set a prefix for the API key, e.g. 'Bearer' (defaults to nil)
  # config.api_key_prefix['apikey cookie salesflare-session'] = 'Bearer'

  # Configure API key authorization: apikey cookie ajs anonymous id
  config.api_key['apikey cookie ajs anonymous id'] = 'YOUR API KEY'
  # Uncomment the following line to set a prefix for the API key, e.g. 'Bearer' (defaults to nil)
  # config.api_key_prefix['apikey cookie ajs anonymous id'] = 'Bearer'

  # Configure API key authorization: apikey header cookie
  config.api_key['apikey header cookie'] = 'YOUR API KEY'
  # Uncomment the following line to set a prefix for the API key, e.g. 'Bearer' (defaults to nil)
  # config.api_key_prefix['apikey header cookie'] = 'Bearer'
end

api_instance = OpenapiClient::DefaultApi.new
opts = {
  limit: 'limit_example', # String | 
  q: 'q_example' # String | 
}

begin
  # /tasks
  result = api_instance.tasks_get(opts)
  p result
rescue OpenapiClient::ApiError => e
  puts "Error when calling DefaultApi->tasks_get: #{e}"
end
```

#### Using the tasks_get_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<Array<TasksGet200ResponseInner>>, Integer, Hash)> tasks_get_with_http_info(opts)

```ruby
begin
  # /tasks
  data, status_code, headers = api_instance.tasks_get_with_http_info(opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <Array<TasksGet200ResponseInner>>
rescue OpenapiClient::ApiError => e
  puts "Error when calling DefaultApi->tasks_get_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **limit** | **String** |  | [optional] |
| **q** | **String** |  | [optional] |

### Return type

[**Array&lt;TasksGet200ResponseInner&gt;**](TasksGet200ResponseInner.md)

### Authorization

[apikey cookie  hjsessionuser 374224](../README.md#apikey cookie  hjsessionuser 374224), [apikey cookie  gid](../README.md#apikey cookie  gid), [apikey cookie  hjsession 374224](../README.md#apikey cookie  hjsession 374224), [apikey cookie ajs user id](../README.md#apikey cookie ajs user id), [apikey cookie intercom-id-nhqftro2](../README.md#apikey cookie intercom-id-nhqftro2), [apikey cookie intercom-session-nhqftro2](../README.md#apikey cookie intercom-session-nhqftro2), [apikey cookie intercom-device-id-nhqftro2](../README.md#apikey cookie intercom-device-id-nhqftro2), [apikey cookie salesflare-session](../README.md#apikey cookie salesflare-session), [apikey cookie ajs anonymous id](../README.md#apikey cookie ajs anonymous id), [apikey header cookie](../README.md#apikey header cookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## tasks_post

> <TasksPost200Response> tasks_post

/tasks

**Host**: http://api.salesflare.com

### Examples

```ruby
require 'time'
require 'openapi_client'
# setup authorization
OpenapiClient.configure do |config|
  # Configure API key authorization: apikey cookie  hjsessionuser 374224
  config.api_key['apikey cookie  hjsessionuser 374224'] = 'YOUR API KEY'
  # Uncomment the following line to set a prefix for the API key, e.g. 'Bearer' (defaults to nil)
  # config.api_key_prefix['apikey cookie  hjsessionuser 374224'] = 'Bearer'

  # Configure API key authorization: apikey cookie  gid
  config.api_key['apikey cookie  gid'] = 'YOUR API KEY'
  # Uncomment the following line to set a prefix for the API key, e.g. 'Bearer' (defaults to nil)
  # config.api_key_prefix['apikey cookie  gid'] = 'Bearer'

  # Configure API key authorization: apikey cookie  hjsession 374224
  config.api_key['apikey cookie  hjsession 374224'] = 'YOUR API KEY'
  # Uncomment the following line to set a prefix for the API key, e.g. 'Bearer' (defaults to nil)
  # config.api_key_prefix['apikey cookie  hjsession 374224'] = 'Bearer'

  # Configure API key authorization: apikey cookie ajs user id
  config.api_key['apikey cookie ajs user id'] = 'YOUR API KEY'
  # Uncomment the following line to set a prefix for the API key, e.g. 'Bearer' (defaults to nil)
  # config.api_key_prefix['apikey cookie ajs user id'] = 'Bearer'

  # Configure API key authorization: apikey cookie intercom-id-nhqftro2
  config.api_key['apikey cookie intercom-id-nhqftro2'] = 'YOUR API KEY'
  # Uncomment the following line to set a prefix for the API key, e.g. 'Bearer' (defaults to nil)
  # config.api_key_prefix['apikey cookie intercom-id-nhqftro2'] = 'Bearer'

  # Configure API key authorization: apikey cookie intercom-session-nhqftro2
  config.api_key['apikey cookie intercom-session-nhqftro2'] = 'YOUR API KEY'
  # Uncomment the following line to set a prefix for the API key, e.g. 'Bearer' (defaults to nil)
  # config.api_key_prefix['apikey cookie intercom-session-nhqftro2'] = 'Bearer'

  # Configure API key authorization: apikey cookie intercom-device-id-nhqftro2
  config.api_key['apikey cookie intercom-device-id-nhqftro2'] = 'YOUR API KEY'
  # Uncomment the following line to set a prefix for the API key, e.g. 'Bearer' (defaults to nil)
  # config.api_key_prefix['apikey cookie intercom-device-id-nhqftro2'] = 'Bearer'

  # Configure API key authorization: apikey cookie salesflare-session
  config.api_key['apikey cookie salesflare-session'] = 'YOUR API KEY'
  # Uncomment the following line to set a prefix for the API key, e.g. 'Bearer' (defaults to nil)
  # config.api_key_prefix['apikey cookie salesflare-session'] = 'Bearer'

  # Configure API key authorization: apikey cookie ajs anonymous id
  config.api_key['apikey cookie ajs anonymous id'] = 'YOUR API KEY'
  # Uncomment the following line to set a prefix for the API key, e.g. 'Bearer' (defaults to nil)
  # config.api_key_prefix['apikey cookie ajs anonymous id'] = 'Bearer'

  # Configure API key authorization: apikey header cookie
  config.api_key['apikey header cookie'] = 'YOUR API KEY'
  # Uncomment the following line to set a prefix for the API key, e.g. 'Bearer' (defaults to nil)
  # config.api_key_prefix['apikey header cookie'] = 'Bearer'
end

api_instance = OpenapiClient::DefaultApi.new

begin
  # /tasks
  result = api_instance.tasks_post
  p result
rescue OpenapiClient::ApiError => e
  puts "Error when calling DefaultApi->tasks_post: #{e}"
end
```

#### Using the tasks_post_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<TasksPost200Response>, Integer, Hash)> tasks_post_with_http_info

```ruby
begin
  # /tasks
  data, status_code, headers = api_instance.tasks_post_with_http_info
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <TasksPost200Response>
rescue OpenapiClient::ApiError => e
  puts "Error when calling DefaultApi->tasks_post_with_http_info: #{e}"
end
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


## v1_m_post

> <V1MPost200Response> v1_m_post

/v1/m

**Host**: http://api.segment.io

### Examples

```ruby
require 'time'
require 'openapi_client'

api_instance = OpenapiClient::DefaultApi.new

begin
  # /v1/m
  result = api_instance.v1_m_post
  p result
rescue OpenapiClient::ApiError => e
  puts "Error when calling DefaultApi->v1_m_post: #{e}"
end
```

#### Using the v1_m_post_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<V1MPost200Response>, Integer, Hash)> v1_m_post_with_http_info

```ruby
begin
  # /v1/m
  data, status_code, headers = api_instance.v1_m_post_with_http_info
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <V1MPost200Response>
rescue OpenapiClient::ApiError => e
  puts "Error when calling DefaultApi->v1_m_post_with_http_info: #{e}"
end
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

