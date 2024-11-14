# \RbdAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiBlockImageCloneFormatVersionGet**](RbdAPI.md#ApiBlockImageCloneFormatVersionGet) | **Get** /api/block/image/clone_format_version | 
[**ApiBlockImageDefaultFeaturesGet**](RbdAPI.md#ApiBlockImageDefaultFeaturesGet) | **Get** /api/block/image/default_features | 
[**ApiBlockImageGet**](RbdAPI.md#ApiBlockImageGet) | **Get** /api/block/image | Display Rbd Images
[**ApiBlockImageImageSpecCopyPost**](RbdAPI.md#ApiBlockImageImageSpecCopyPost) | **Post** /api/block/image/{image_spec}/copy | 
[**ApiBlockImageImageSpecDelete**](RbdAPI.md#ApiBlockImageImageSpecDelete) | **Delete** /api/block/image/{image_spec} | 
[**ApiBlockImageImageSpecFlattenPost**](RbdAPI.md#ApiBlockImageImageSpecFlattenPost) | **Post** /api/block/image/{image_spec}/flatten | 
[**ApiBlockImageImageSpecGet**](RbdAPI.md#ApiBlockImageImageSpecGet) | **Get** /api/block/image/{image_spec} | Get Rbd Image Info
[**ApiBlockImageImageSpecMoveTrashPost**](RbdAPI.md#ApiBlockImageImageSpecMoveTrashPost) | **Post** /api/block/image/{image_spec}/move_trash | 
[**ApiBlockImageImageSpecPut**](RbdAPI.md#ApiBlockImageImageSpecPut) | **Put** /api/block/image/{image_spec} | 
[**ApiBlockImagePost**](RbdAPI.md#ApiBlockImagePost) | **Post** /api/block/image | 



## ApiBlockImageCloneFormatVersionGet

> ApiBlockImageCloneFormatVersionGet(ctx).Execute()





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
	r, err := apiClient.RbdAPI.ApiBlockImageCloneFormatVersionGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RbdAPI.ApiBlockImageCloneFormatVersionGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiBlockImageCloneFormatVersionGetRequest struct via the builder pattern


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


## ApiBlockImageDefaultFeaturesGet

> ApiBlockImageDefaultFeaturesGet(ctx).Execute()



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
	r, err := apiClient.RbdAPI.ApiBlockImageDefaultFeaturesGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RbdAPI.ApiBlockImageDefaultFeaturesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiBlockImageDefaultFeaturesGetRequest struct via the builder pattern


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


## ApiBlockImageGet

> []ApiBlockImageGet200ResponseInner ApiBlockImageGet(ctx).PoolName(poolName).Offset(offset).Limit(limit).Search(search).Sort(sort).Execute()

Display Rbd Images

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
	poolName := "poolName_example" // string | Pool Name (optional)
	offset := int32(56) // int32 | offset (optional)
	limit := int32(56) // int32 | limit (optional)
	search := "search_example" // string |  (optional)
	sort := "sort_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RbdAPI.ApiBlockImageGet(context.Background()).PoolName(poolName).Offset(offset).Limit(limit).Search(search).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RbdAPI.ApiBlockImageGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiBlockImageGet`: []ApiBlockImageGet200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `RbdAPI.ApiBlockImageGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiBlockImageGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **poolName** | **string** | Pool Name | 
 **offset** | **int32** | offset | 
 **limit** | **int32** | limit | 
 **search** | **string** |  | 
 **sort** | **string** |  | 

### Return type

[**[]ApiBlockImageGet200ResponseInner**](ApiBlockImageGet200ResponseInner.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.ceph.api.v2.0+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiBlockImageImageSpecCopyPost

> ApiBlockImageImageSpecCopyPost(ctx, imageSpec).ApiBlockImageImageSpecCopyPostRequest(apiBlockImageImageSpecCopyPostRequest).Execute()



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
	imageSpec := "imageSpec_example" // string | 
	apiBlockImageImageSpecCopyPostRequest := *openapiclient.NewApiBlockImageImageSpecCopyPostRequest("DestImageName_example", "DestNamespace_example", "DestPoolName_example") // ApiBlockImageImageSpecCopyPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RbdAPI.ApiBlockImageImageSpecCopyPost(context.Background(), imageSpec).ApiBlockImageImageSpecCopyPostRequest(apiBlockImageImageSpecCopyPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RbdAPI.ApiBlockImageImageSpecCopyPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imageSpec** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiBlockImageImageSpecCopyPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiBlockImageImageSpecCopyPostRequest** | [**ApiBlockImageImageSpecCopyPostRequest**](ApiBlockImageImageSpecCopyPostRequest.md) |  | 

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


## ApiBlockImageImageSpecDelete

> ApiBlockImageImageSpecDelete(ctx, imageSpec).Execute()



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
	imageSpec := "imageSpec_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RbdAPI.ApiBlockImageImageSpecDelete(context.Background(), imageSpec).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RbdAPI.ApiBlockImageImageSpecDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imageSpec** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiBlockImageImageSpecDeleteRequest struct via the builder pattern


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


## ApiBlockImageImageSpecFlattenPost

> ApiBlockImageImageSpecFlattenPost(ctx, imageSpec).Execute()



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
	imageSpec := "imageSpec_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RbdAPI.ApiBlockImageImageSpecFlattenPost(context.Background(), imageSpec).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RbdAPI.ApiBlockImageImageSpecFlattenPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imageSpec** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiBlockImageImageSpecFlattenPostRequest struct via the builder pattern


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


## ApiBlockImageImageSpecGet

> ApiBlockImageImageSpecGet200Response ApiBlockImageImageSpecGet(ctx, imageSpec).OmitUsage(omitUsage).Execute()

Get Rbd Image Info

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
	imageSpec := "imageSpec_example" // string | URL-encoded \"pool/rbd_name\". e.g. \"rbd%2Ffoo\"
	omitUsage := true // bool | When true, usage information is not returned (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RbdAPI.ApiBlockImageImageSpecGet(context.Background(), imageSpec).OmitUsage(omitUsage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RbdAPI.ApiBlockImageImageSpecGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiBlockImageImageSpecGet`: ApiBlockImageImageSpecGet200Response
	fmt.Fprintf(os.Stdout, "Response from `RbdAPI.ApiBlockImageImageSpecGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imageSpec** | **string** | URL-encoded \&quot;pool/rbd_name\&quot;. e.g. \&quot;rbd%2Ffoo\&quot; | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiBlockImageImageSpecGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **omitUsage** | **bool** | When true, usage information is not returned | 

### Return type

[**ApiBlockImageImageSpecGet200Response**](ApiBlockImageImageSpecGet200Response.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.ceph.api.v1.0+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiBlockImageImageSpecMoveTrashPost

> ApiBlockImageImageSpecMoveTrashPost(ctx, imageSpec).ApiBlockImageImageSpecMoveTrashPostRequest(apiBlockImageImageSpecMoveTrashPostRequest).Execute()





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
	imageSpec := "imageSpec_example" // string | 
	apiBlockImageImageSpecMoveTrashPostRequest := *openapiclient.NewApiBlockImageImageSpecMoveTrashPostRequest() // ApiBlockImageImageSpecMoveTrashPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RbdAPI.ApiBlockImageImageSpecMoveTrashPost(context.Background(), imageSpec).ApiBlockImageImageSpecMoveTrashPostRequest(apiBlockImageImageSpecMoveTrashPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RbdAPI.ApiBlockImageImageSpecMoveTrashPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imageSpec** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiBlockImageImageSpecMoveTrashPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiBlockImageImageSpecMoveTrashPostRequest** | [**ApiBlockImageImageSpecMoveTrashPostRequest**](ApiBlockImageImageSpecMoveTrashPostRequest.md) |  | 

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


## ApiBlockImageImageSpecPut

> ApiBlockImageImageSpecPut(ctx, imageSpec).ApiBlockImageImageSpecPutRequest(apiBlockImageImageSpecPutRequest).Execute()



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
	imageSpec := "imageSpec_example" // string | 
	apiBlockImageImageSpecPutRequest := *openapiclient.NewApiBlockImageImageSpecPutRequest() // ApiBlockImageImageSpecPutRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RbdAPI.ApiBlockImageImageSpecPut(context.Background(), imageSpec).ApiBlockImageImageSpecPutRequest(apiBlockImageImageSpecPutRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RbdAPI.ApiBlockImageImageSpecPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imageSpec** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiBlockImageImageSpecPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiBlockImageImageSpecPutRequest** | [**ApiBlockImageImageSpecPutRequest**](ApiBlockImageImageSpecPutRequest.md) |  | 

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


## ApiBlockImagePost

> ApiBlockImagePost(ctx).ApiBlockImagePostRequest(apiBlockImagePostRequest).Execute()



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
	apiBlockImagePostRequest := *openapiclient.NewApiBlockImagePostRequest("Name_example", "PoolName_example", int32(123)) // ApiBlockImagePostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RbdAPI.ApiBlockImagePost(context.Background()).ApiBlockImagePostRequest(apiBlockImagePostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RbdAPI.ApiBlockImagePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiBlockImagePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiBlockImagePostRequest** | [**ApiBlockImagePostRequest**](ApiBlockImagePostRequest.md) |  | 

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

