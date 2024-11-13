# \RgwZonegroupAPI

All URIs are relative to *https://raw.githubusercontent.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiRgwZonegroupGet**](RgwZonegroupAPI.md#ApiRgwZonegroupGet) | **Get** /api/rgw/zonegroup | 
[**ApiRgwZonegroupGetAllZonegroupsInfoGet**](RgwZonegroupAPI.md#ApiRgwZonegroupGetAllZonegroupsInfoGet) | **Get** /api/rgw/zonegroup/get_all_zonegroups_info | 
[**ApiRgwZonegroupPost**](RgwZonegroupAPI.md#ApiRgwZonegroupPost) | **Post** /api/rgw/zonegroup | 
[**ApiRgwZonegroupZonegroupNameDelete**](RgwZonegroupAPI.md#ApiRgwZonegroupZonegroupNameDelete) | **Delete** /api/rgw/zonegroup/{zonegroup_name} | 
[**ApiRgwZonegroupZonegroupNameGet**](RgwZonegroupAPI.md#ApiRgwZonegroupZonegroupNameGet) | **Get** /api/rgw/zonegroup/{zonegroup_name} | 
[**ApiRgwZonegroupZonegroupNamePut**](RgwZonegroupAPI.md#ApiRgwZonegroupZonegroupNamePut) | **Put** /api/rgw/zonegroup/{zonegroup_name} | 



## ApiRgwZonegroupGet

> ApiRgwZonegroupGet(ctx).Execute()



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
	r, err := apiClient.RgwZonegroupAPI.ApiRgwZonegroupGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RgwZonegroupAPI.ApiRgwZonegroupGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiRgwZonegroupGetRequest struct via the builder pattern


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


## ApiRgwZonegroupGetAllZonegroupsInfoGet

> ApiRgwZonegroupGetAllZonegroupsInfoGet(ctx).Execute()



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
	r, err := apiClient.RgwZonegroupAPI.ApiRgwZonegroupGetAllZonegroupsInfoGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RgwZonegroupAPI.ApiRgwZonegroupGetAllZonegroupsInfoGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiRgwZonegroupGetAllZonegroupsInfoGetRequest struct via the builder pattern


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


## ApiRgwZonegroupPost

> ApiRgwZonegroupPost(ctx).ApiRgwZonegroupPostRequest(apiRgwZonegroupPostRequest).Execute()



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
	apiRgwZonegroupPostRequest := *openapiclient.NewApiRgwZonegroupPostRequest("RealmName_example", "ZonegroupName_example") // ApiRgwZonegroupPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RgwZonegroupAPI.ApiRgwZonegroupPost(context.Background()).ApiRgwZonegroupPostRequest(apiRgwZonegroupPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RgwZonegroupAPI.ApiRgwZonegroupPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiRgwZonegroupPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiRgwZonegroupPostRequest** | [**ApiRgwZonegroupPostRequest**](ApiRgwZonegroupPostRequest.md) |  | 

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


## ApiRgwZonegroupZonegroupNameDelete

> ApiRgwZonegroupZonegroupNameDelete(ctx, zonegroupName).DeletePools(deletePools).Pools(pools).Execute()



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
	zonegroupName := "zonegroupName_example" // string | 
	deletePools := "deletePools_example" // string | 
	pools := "pools_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RgwZonegroupAPI.ApiRgwZonegroupZonegroupNameDelete(context.Background(), zonegroupName).DeletePools(deletePools).Pools(pools).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RgwZonegroupAPI.ApiRgwZonegroupZonegroupNameDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**zonegroupName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiRgwZonegroupZonegroupNameDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deletePools** | **string** |  | 
 **pools** | **string** |  | 

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


## ApiRgwZonegroupZonegroupNameGet

> ApiRgwZonegroupZonegroupNameGet(ctx, zonegroupName).Execute()



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
	zonegroupName := "zonegroupName_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RgwZonegroupAPI.ApiRgwZonegroupZonegroupNameGet(context.Background(), zonegroupName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RgwZonegroupAPI.ApiRgwZonegroupZonegroupNameGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**zonegroupName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiRgwZonegroupZonegroupNameGetRequest struct via the builder pattern


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


## ApiRgwZonegroupZonegroupNamePut

> ApiRgwZonegroupZonegroupNamePut(ctx, zonegroupName).ApiRgwZonegroupZonegroupNamePutRequest(apiRgwZonegroupZonegroupNamePutRequest).Execute()



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
	zonegroupName := "zonegroupName_example" // string | 
	apiRgwZonegroupZonegroupNamePutRequest := *openapiclient.NewApiRgwZonegroupZonegroupNamePutRequest("NewZonegroupName_example", "RealmName_example") // ApiRgwZonegroupZonegroupNamePutRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RgwZonegroupAPI.ApiRgwZonegroupZonegroupNamePut(context.Background(), zonegroupName).ApiRgwZonegroupZonegroupNamePutRequest(apiRgwZonegroupZonegroupNamePutRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RgwZonegroupAPI.ApiRgwZonegroupZonegroupNamePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**zonegroupName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiRgwZonegroupZonegroupNamePutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiRgwZonegroupZonegroupNamePutRequest** | [**ApiRgwZonegroupZonegroupNamePutRequest**](ApiRgwZonegroupZonegroupNamePutRequest.md) |  | 

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

