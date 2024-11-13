# \RgwZoneAPI

All URIs are relative to *https://raw.githubusercontent.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiRgwZoneCreateSystemUserPut**](RgwZoneAPI.md#ApiRgwZoneCreateSystemUserPut) | **Put** /api/rgw/zone/create_system_user | 
[**ApiRgwZoneGet**](RgwZoneAPI.md#ApiRgwZoneGet) | **Get** /api/rgw/zone | 
[**ApiRgwZoneGetAllZonesInfoGet**](RgwZoneAPI.md#ApiRgwZoneGetAllZonesInfoGet) | **Get** /api/rgw/zone/get_all_zones_info | 
[**ApiRgwZoneGetPoolNamesGet**](RgwZoneAPI.md#ApiRgwZoneGetPoolNamesGet) | **Get** /api/rgw/zone/get_pool_names | 
[**ApiRgwZoneGetUserListGet**](RgwZoneAPI.md#ApiRgwZoneGetUserListGet) | **Get** /api/rgw/zone/get_user_list | 
[**ApiRgwZonePost**](RgwZoneAPI.md#ApiRgwZonePost) | **Post** /api/rgw/zone | 
[**ApiRgwZoneZoneNameDelete**](RgwZoneAPI.md#ApiRgwZoneZoneNameDelete) | **Delete** /api/rgw/zone/{zone_name} | 
[**ApiRgwZoneZoneNameGet**](RgwZoneAPI.md#ApiRgwZoneZoneNameGet) | **Get** /api/rgw/zone/{zone_name} | 
[**ApiRgwZoneZoneNamePut**](RgwZoneAPI.md#ApiRgwZoneZoneNamePut) | **Put** /api/rgw/zone/{zone_name} | 



## ApiRgwZoneCreateSystemUserPut

> ApiRgwZoneCreateSystemUserPut(ctx).ApiRgwZoneCreateSystemUserPutRequest(apiRgwZoneCreateSystemUserPutRequest).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/boyvinall/go-ceph"
)

func main() {
	apiRgwZoneCreateSystemUserPutRequest := *openapiclient.NewApiRgwZoneCreateSystemUserPutRequest("UserName_example", "ZoneName_example") // ApiRgwZoneCreateSystemUserPutRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RgwZoneAPI.ApiRgwZoneCreateSystemUserPut(context.Background()).ApiRgwZoneCreateSystemUserPutRequest(apiRgwZoneCreateSystemUserPutRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RgwZoneAPI.ApiRgwZoneCreateSystemUserPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiRgwZoneCreateSystemUserPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiRgwZoneCreateSystemUserPutRequest** | [**ApiRgwZoneCreateSystemUserPutRequest**](ApiRgwZoneCreateSystemUserPutRequest.md) |  | 

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


## ApiRgwZoneGet

> ApiRgwZoneGet(ctx).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/boyvinall/go-ceph"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RgwZoneAPI.ApiRgwZoneGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RgwZoneAPI.ApiRgwZoneGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiRgwZoneGetRequest struct via the builder pattern


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


## ApiRgwZoneGetAllZonesInfoGet

> ApiRgwZoneGetAllZonesInfoGet(ctx).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/boyvinall/go-ceph"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RgwZoneAPI.ApiRgwZoneGetAllZonesInfoGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RgwZoneAPI.ApiRgwZoneGetAllZonesInfoGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiRgwZoneGetAllZonesInfoGetRequest struct via the builder pattern


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


## ApiRgwZoneGetPoolNamesGet

> ApiRgwZoneGetPoolNamesGet(ctx).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/boyvinall/go-ceph"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RgwZoneAPI.ApiRgwZoneGetPoolNamesGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RgwZoneAPI.ApiRgwZoneGetPoolNamesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiRgwZoneGetPoolNamesGetRequest struct via the builder pattern


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


## ApiRgwZoneGetUserListGet

> ApiRgwZoneGetUserListGet(ctx).ZoneName(zoneName).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/boyvinall/go-ceph"
)

func main() {
	zoneName := "zoneName_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RgwZoneAPI.ApiRgwZoneGetUserListGet(context.Background()).ZoneName(zoneName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RgwZoneAPI.ApiRgwZoneGetUserListGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiRgwZoneGetUserListGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **zoneName** | **string** |  | 

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


## ApiRgwZonePost

> ApiRgwZonePost(ctx).ApiRgwZonePostRequest(apiRgwZonePostRequest).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/boyvinall/go-ceph"
)

func main() {
	apiRgwZonePostRequest := *openapiclient.NewApiRgwZonePostRequest("ZoneName_example") // ApiRgwZonePostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RgwZoneAPI.ApiRgwZonePost(context.Background()).ApiRgwZonePostRequest(apiRgwZonePostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RgwZoneAPI.ApiRgwZonePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiRgwZonePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiRgwZonePostRequest** | [**ApiRgwZonePostRequest**](ApiRgwZonePostRequest.md) |  | 

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


## ApiRgwZoneZoneNameDelete

> ApiRgwZoneZoneNameDelete(ctx, zoneName).DeletePools(deletePools).Pools(pools).ZonegroupName(zonegroupName).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/boyvinall/go-ceph"
)

func main() {
	zoneName := "zoneName_example" // string | 
	deletePools := "deletePools_example" // string | 
	pools := "pools_example" // string |  (optional)
	zonegroupName := "zonegroupName_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RgwZoneAPI.ApiRgwZoneZoneNameDelete(context.Background(), zoneName).DeletePools(deletePools).Pools(pools).ZonegroupName(zonegroupName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RgwZoneAPI.ApiRgwZoneZoneNameDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**zoneName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiRgwZoneZoneNameDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deletePools** | **string** |  | 
 **pools** | **string** |  | 
 **zonegroupName** | **string** |  | 

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


## ApiRgwZoneZoneNameGet

> ApiRgwZoneZoneNameGet(ctx, zoneName).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/boyvinall/go-ceph"
)

func main() {
	zoneName := "zoneName_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RgwZoneAPI.ApiRgwZoneZoneNameGet(context.Background(), zoneName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RgwZoneAPI.ApiRgwZoneZoneNameGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**zoneName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiRgwZoneZoneNameGetRequest struct via the builder pattern


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


## ApiRgwZoneZoneNamePut

> ApiRgwZoneZoneNamePut(ctx, zoneName).ApiRgwZoneZoneNamePutRequest(apiRgwZoneZoneNamePutRequest).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/boyvinall/go-ceph"
)

func main() {
	zoneName := "zoneName_example" // string | 
	apiRgwZoneZoneNamePutRequest := *openapiclient.NewApiRgwZoneZoneNamePutRequest("NewZoneName_example", "ZonegroupName_example") // ApiRgwZoneZoneNamePutRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RgwZoneAPI.ApiRgwZoneZoneNamePut(context.Background(), zoneName).ApiRgwZoneZoneNamePutRequest(apiRgwZoneZoneNamePutRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RgwZoneAPI.ApiRgwZoneZoneNamePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**zoneName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiRgwZoneZoneNamePutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiRgwZoneZoneNamePutRequest** | [**ApiRgwZoneZoneNamePutRequest**](ApiRgwZoneZoneNamePutRequest.md) |  | 

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

