# \RgwRealmAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiRgwRealmGet**](RgwRealmAPI.md#ApiRgwRealmGet) | **Get** /api/rgw/realm | 
[**ApiRgwRealmGetAllRealmsInfoGet**](RgwRealmAPI.md#ApiRgwRealmGetAllRealmsInfoGet) | **Get** /api/rgw/realm/get_all_realms_info | 
[**ApiRgwRealmGetRealmTokensGet**](RgwRealmAPI.md#ApiRgwRealmGetRealmTokensGet) | **Get** /api/rgw/realm/get_realm_tokens | 
[**ApiRgwRealmImportRealmTokenPost**](RgwRealmAPI.md#ApiRgwRealmImportRealmTokenPost) | **Post** /api/rgw/realm/import_realm_token | 
[**ApiRgwRealmPost**](RgwRealmAPI.md#ApiRgwRealmPost) | **Post** /api/rgw/realm | 
[**ApiRgwRealmRealmNameDelete**](RgwRealmAPI.md#ApiRgwRealmRealmNameDelete) | **Delete** /api/rgw/realm/{realm_name} | 
[**ApiRgwRealmRealmNameGet**](RgwRealmAPI.md#ApiRgwRealmRealmNameGet) | **Get** /api/rgw/realm/{realm_name} | 
[**ApiRgwRealmRealmNamePut**](RgwRealmAPI.md#ApiRgwRealmRealmNamePut) | **Put** /api/rgw/realm/{realm_name} | 



## ApiRgwRealmGet

> ApiRgwRealmGet(ctx).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/boyvinall/go-ceph/api"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RgwRealmAPI.ApiRgwRealmGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RgwRealmAPI.ApiRgwRealmGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiRgwRealmGetRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.ceph.api.v1.0+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiRgwRealmGetAllRealmsInfoGet

> ApiRgwRealmGetAllRealmsInfoGet(ctx).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/boyvinall/go-ceph/api"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RgwRealmAPI.ApiRgwRealmGetAllRealmsInfoGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RgwRealmAPI.ApiRgwRealmGetAllRealmsInfoGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiRgwRealmGetAllRealmsInfoGetRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.ceph.api.v1.0+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiRgwRealmGetRealmTokensGet

> ApiRgwRealmGetRealmTokensGet(ctx).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/boyvinall/go-ceph/api"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RgwRealmAPI.ApiRgwRealmGetRealmTokensGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RgwRealmAPI.ApiRgwRealmGetRealmTokensGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiRgwRealmGetRealmTokensGetRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.ceph.api.v1.0+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiRgwRealmImportRealmTokenPost

> ApiRgwRealmImportRealmTokenPost(ctx).ApiRgwRealmImportRealmTokenPostRequest(apiRgwRealmImportRealmTokenPostRequest).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/boyvinall/go-ceph/api"
)

func main() {
	apiRgwRealmImportRealmTokenPostRequest := *openapiclient.NewApiRgwRealmImportRealmTokenPostRequest("Port_example", "RealmToken_example", "ZoneName_example") // ApiRgwRealmImportRealmTokenPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RgwRealmAPI.ApiRgwRealmImportRealmTokenPost(context.Background()).ApiRgwRealmImportRealmTokenPostRequest(apiRgwRealmImportRealmTokenPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RgwRealmAPI.ApiRgwRealmImportRealmTokenPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiRgwRealmImportRealmTokenPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiRgwRealmImportRealmTokenPostRequest** | [**ApiRgwRealmImportRealmTokenPostRequest**](ApiRgwRealmImportRealmTokenPostRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.ceph.api.v1.0+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiRgwRealmPost

> ApiRgwRealmPost(ctx).ApiRgwRealmPostRequest(apiRgwRealmPostRequest).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/boyvinall/go-ceph/api"
)

func main() {
	apiRgwRealmPostRequest := *openapiclient.NewApiRgwRealmPostRequest("Default_example", "RealmName_example") // ApiRgwRealmPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RgwRealmAPI.ApiRgwRealmPost(context.Background()).ApiRgwRealmPostRequest(apiRgwRealmPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RgwRealmAPI.ApiRgwRealmPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiRgwRealmPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiRgwRealmPostRequest** | [**ApiRgwRealmPostRequest**](ApiRgwRealmPostRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.ceph.api.v1.0+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiRgwRealmRealmNameDelete

> ApiRgwRealmRealmNameDelete(ctx, realmName).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/boyvinall/go-ceph/api"
)

func main() {
	realmName := "realmName_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RgwRealmAPI.ApiRgwRealmRealmNameDelete(context.Background(), realmName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RgwRealmAPI.ApiRgwRealmRealmNameDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realmName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiRgwRealmRealmNameDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.ceph.api.v1.0+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiRgwRealmRealmNameGet

> ApiRgwRealmRealmNameGet(ctx, realmName).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/boyvinall/go-ceph/api"
)

func main() {
	realmName := "realmName_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RgwRealmAPI.ApiRgwRealmRealmNameGet(context.Background(), realmName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RgwRealmAPI.ApiRgwRealmRealmNameGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realmName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiRgwRealmRealmNameGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.ceph.api.v1.0+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiRgwRealmRealmNamePut

> ApiRgwRealmRealmNamePut(ctx, realmName).ApiRgwRealmRealmNamePutRequest(apiRgwRealmRealmNamePutRequest).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/boyvinall/go-ceph/api"
)

func main() {
	realmName := "realmName_example" // string | 
	apiRgwRealmRealmNamePutRequest := *openapiclient.NewApiRgwRealmRealmNamePutRequest("NewRealmName_example") // ApiRgwRealmRealmNamePutRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RgwRealmAPI.ApiRgwRealmRealmNamePut(context.Background(), realmName).ApiRgwRealmRealmNamePutRequest(apiRgwRealmRealmNamePutRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RgwRealmAPI.ApiRgwRealmRealmNamePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realmName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiRgwRealmRealmNamePutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiRgwRealmRealmNamePutRequest** | [**ApiRgwRealmRealmNamePutRequest**](ApiRgwRealmRealmNamePutRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.ceph.api.v1.0+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

