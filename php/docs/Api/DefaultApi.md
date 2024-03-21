# OpenAPI\Client\DefaultApi

All URIs are relative to http://localhost, except if the operation defines another base path.

| Method | HTTP request | Description |
| ------------- | ------------- | ------------- |
| [**messengerWebMetricsPost()**](DefaultApi.md#messengerWebMetricsPost) | **POST** /messenger/web/metrics | /messenger/web/metrics |
| [**tasksGet()**](DefaultApi.md#tasksGet) | **GET** /tasks | /tasks |
| [**tasksPost()**](DefaultApi.md#tasksPost) | **POST** /tasks | /tasks |
| [**v1MPost()**](DefaultApi.md#v1MPost) | **POST** /v1/m | /v1/m |


## `messengerWebMetricsPost()`

```php
messengerWebMetricsPost($app_id, $v, $g, $s, $r, $platform, $installation_type, $idempotency_key, $internal, $is_intersection_booted, $page_title, $user_active_company_id, $metrics, $logs, $op_metrics, $hc_metrics, $referer, $anonymous_session, $device_identifier)
```

/messenger/web/metrics

**Host**: http://api-iam.intercom.io

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');



$apiInstance = new OpenAPI\Client\Api\DefaultApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client()
);
$app_id = 'app_id_example'; // string
$v = 'v_example'; // string
$g = 'g_example'; // string
$s = 's_example'; // string
$r = 'r_example'; // string
$platform = 'platform_example'; // string
$installation_type = 'installation_type_example'; // string
$idempotency_key = 'idempotency_key_example'; // string
$internal = 'internal_example'; // string
$is_intersection_booted = 'is_intersection_booted_example'; // string
$page_title = 'page_title_example'; // string
$user_active_company_id = 'user_active_company_id_example'; // string
$metrics = 'metrics_example'; // string
$logs = 'logs_example'; // string
$op_metrics = 'op_metrics_example'; // string
$hc_metrics = 'hc_metrics_example'; // string
$referer = 'referer_example'; // string
$anonymous_session = 'anonymous_session_example'; // string
$device_identifier = 'device_identifier_example'; // string

try {
    $apiInstance->messengerWebMetricsPost($app_id, $v, $g, $s, $r, $platform, $installation_type, $idempotency_key, $internal, $is_intersection_booted, $page_title, $user_active_company_id, $metrics, $logs, $op_metrics, $hc_metrics, $referer, $anonymous_session, $device_identifier);
} catch (Exception $e) {
    echo 'Exception when calling DefaultApi->messengerWebMetricsPost: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

| Name | Type | Description  | Notes |
| ------------- | ------------- | ------------- | ------------- |
| **app_id** | **string**|  | |
| **v** | **string**|  | |
| **g** | **string**|  | |
| **s** | **string**|  | |
| **r** | **string**|  | |
| **platform** | **string**|  | |
| **installation_type** | **string**|  | |
| **idempotency_key** | **string**|  | |
| **internal** | **string**|  | |
| **is_intersection_booted** | **string**|  | |
| **page_title** | **string**|  | |
| **user_active_company_id** | **string**|  | |
| **metrics** | **string**|  | |
| **logs** | **string**|  | |
| **op_metrics** | **string**|  | |
| **hc_metrics** | **string**|  | |
| **referer** | **string**|  | |
| **anonymous_session** | **string**|  | |
| **device_identifier** | **string**|  | |

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: `application/x-www-form-urlencoded`
- **Accept**: `text/html`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `tasksGet()`

```php
tasksGet($limit, $q): \OpenAPI\Client\Model\TasksGet200ResponseInner[]
```

/tasks

**Host**: http://api.salesflare.com

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure API key authorization: apikey cookie  hjsessionuser 374224
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKey('_hjsessionuser_374224', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// $config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('_hjsessionuser_374224', 'Bearer');

// Configure API key authorization: apikey cookie  gid
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKey('_gid', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// $config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('_gid', 'Bearer');

// Configure API key authorization: apikey cookie  hjsession 374224
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKey('_hjsession_374224', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// $config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('_hjsession_374224', 'Bearer');

// Configure API key authorization: apikey cookie ajs user id
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKey('ajs_user_id', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// $config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('ajs_user_id', 'Bearer');

// Configure API key authorization: apikey cookie intercom-id-nhqftro2
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKey('intercom-id-nhqftro2', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// $config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('intercom-id-nhqftro2', 'Bearer');

// Configure API key authorization: apikey cookie intercom-session-nhqftro2
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKey('intercom-session-nhqftro2', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// $config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('intercom-session-nhqftro2', 'Bearer');

// Configure API key authorization: apikey cookie intercom-device-id-nhqftro2
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKey('intercom-device-id-nhqftro2', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// $config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('intercom-device-id-nhqftro2', 'Bearer');

// Configure API key authorization: apikey cookie salesflare-session
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKey('salesflare-session', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// $config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('salesflare-session', 'Bearer');

// Configure API key authorization: apikey cookie ajs anonymous id
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKey('ajs_anonymous_id', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// $config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('ajs_anonymous_id', 'Bearer');

// Configure API key authorization: apikey header cookie
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKey('COOKIE', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// $config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('COOKIE', 'Bearer');


$apiInstance = new OpenAPI\Client\Api\DefaultApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);
$limit = 'limit_example'; // string
$q = 'q_example'; // string

try {
    $result = $apiInstance->tasksGet($limit, $q);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling DefaultApi->tasksGet: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

| Name | Type | Description  | Notes |
| ------------- | ------------- | ------------- | ------------- |
| **limit** | **string**|  | [optional] |
| **q** | **string**|  | [optional] |

### Return type

[**\OpenAPI\Client\Model\TasksGet200ResponseInner[]**](../Model/TasksGet200ResponseInner.md)

### Authorization

[apikey cookie  hjsessionuser 374224](../../README.md#apikey cookie  hjsessionuser 374224), [apikey cookie  gid](../../README.md#apikey cookie  gid), [apikey cookie  hjsession 374224](../../README.md#apikey cookie  hjsession 374224), [apikey cookie ajs user id](../../README.md#apikey cookie ajs user id), [apikey cookie intercom-id-nhqftro2](../../README.md#apikey cookie intercom-id-nhqftro2), [apikey cookie intercom-session-nhqftro2](../../README.md#apikey cookie intercom-session-nhqftro2), [apikey cookie intercom-device-id-nhqftro2](../../README.md#apikey cookie intercom-device-id-nhqftro2), [apikey cookie salesflare-session](../../README.md#apikey cookie salesflare-session), [apikey cookie ajs anonymous id](../../README.md#apikey cookie ajs anonymous id), [apikey header cookie](../../README.md#apikey header cookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `tasksPost()`

```php
tasksPost(): \OpenAPI\Client\Model\TasksPost200Response
```

/tasks

**Host**: http://api.salesflare.com

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');


// Configure API key authorization: apikey cookie  hjsessionuser 374224
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKey('_hjsessionuser_374224', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// $config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('_hjsessionuser_374224', 'Bearer');

// Configure API key authorization: apikey cookie  gid
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKey('_gid', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// $config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('_gid', 'Bearer');

// Configure API key authorization: apikey cookie  hjsession 374224
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKey('_hjsession_374224', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// $config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('_hjsession_374224', 'Bearer');

// Configure API key authorization: apikey cookie ajs user id
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKey('ajs_user_id', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// $config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('ajs_user_id', 'Bearer');

// Configure API key authorization: apikey cookie intercom-id-nhqftro2
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKey('intercom-id-nhqftro2', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// $config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('intercom-id-nhqftro2', 'Bearer');

// Configure API key authorization: apikey cookie intercom-session-nhqftro2
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKey('intercom-session-nhqftro2', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// $config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('intercom-session-nhqftro2', 'Bearer');

// Configure API key authorization: apikey cookie intercom-device-id-nhqftro2
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKey('intercom-device-id-nhqftro2', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// $config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('intercom-device-id-nhqftro2', 'Bearer');

// Configure API key authorization: apikey cookie salesflare-session
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKey('salesflare-session', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// $config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('salesflare-session', 'Bearer');

// Configure API key authorization: apikey cookie ajs anonymous id
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKey('ajs_anonymous_id', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// $config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('ajs_anonymous_id', 'Bearer');

// Configure API key authorization: apikey header cookie
$config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKey('COOKIE', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// $config = OpenAPI\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('COOKIE', 'Bearer');


$apiInstance = new OpenAPI\Client\Api\DefaultApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client(),
    $config
);

try {
    $result = $apiInstance->tasksPost();
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling DefaultApi->tasksPost: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

This endpoint does not need any parameter.

### Return type

[**\OpenAPI\Client\Model\TasksPost200Response**](../Model/TasksPost200Response.md)

### Authorization

[apikey cookie  hjsessionuser 374224](../../README.md#apikey cookie  hjsessionuser 374224), [apikey cookie  gid](../../README.md#apikey cookie  gid), [apikey cookie  hjsession 374224](../../README.md#apikey cookie  hjsession 374224), [apikey cookie ajs user id](../../README.md#apikey cookie ajs user id), [apikey cookie intercom-id-nhqftro2](../../README.md#apikey cookie intercom-id-nhqftro2), [apikey cookie intercom-session-nhqftro2](../../README.md#apikey cookie intercom-session-nhqftro2), [apikey cookie intercom-device-id-nhqftro2](../../README.md#apikey cookie intercom-device-id-nhqftro2), [apikey cookie salesflare-session](../../README.md#apikey cookie salesflare-session), [apikey cookie ajs anonymous id](../../README.md#apikey cookie ajs anonymous id), [apikey header cookie](../../README.md#apikey header cookie)

### HTTP request headers

- **Content-Type**: `application/json;charset=UTF-8`
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)

## `v1MPost()`

```php
v1MPost(): \OpenAPI\Client\Model\V1MPost200Response
```

/v1/m

**Host**: http://api.segment.io

### Example

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');



$apiInstance = new OpenAPI\Client\Api\DefaultApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client()
);

try {
    $result = $apiInstance->v1MPost();
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling DefaultApi->v1MPost: ', $e->getMessage(), PHP_EOL;
}
```

### Parameters

This endpoint does not need any parameter.

### Return type

[**\OpenAPI\Client\Model\V1MPost200Response**](../Model/V1MPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: `text/plain`
- **Accept**: `application/json`

[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)
[[Back to Model list]](../../README.md#models)
[[Back to README]](../../README.md)
