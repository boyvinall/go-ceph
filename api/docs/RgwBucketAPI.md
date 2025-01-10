# \RgwBucketAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiRgwBucketBucketDelete**](RgwBucketAPI.md#ApiRgwBucketBucketDelete) | **Delete** /api/rgw/bucket/{bucket} | 
[**ApiRgwBucketBucketGet**](RgwBucketAPI.md#ApiRgwBucketBucketGet) | **Get** /api/rgw/bucket/{bucket} | 
[**ApiRgwBucketBucketPut**](RgwBucketAPI.md#ApiRgwBucketBucketPut) | **Put** /api/rgw/bucket/{bucket} | 
[**ApiRgwBucketDeleteEncryptionDelete**](RgwBucketAPI.md#ApiRgwBucketDeleteEncryptionDelete) | **Delete** /api/rgw/bucket/deleteEncryption | 
[**ApiRgwBucketGet**](RgwBucketAPI.md#ApiRgwBucketGet) | **Get** /api/rgw/bucket | 
[**ApiRgwBucketGetEncryptionConfigGet**](RgwBucketAPI.md#ApiRgwBucketGetEncryptionConfigGet) | **Get** /api/rgw/bucket/getEncryptionConfig | 
[**ApiRgwBucketGetEncryptionGet**](RgwBucketAPI.md#ApiRgwBucketGetEncryptionGet) | **Get** /api/rgw/bucket/getEncryption | 
[**ApiRgwBucketPost**](RgwBucketAPI.md#ApiRgwBucketPost) | **Post** /api/rgw/bucket | 
[**ApiRgwBucketSetEncryptionConfigPut**](RgwBucketAPI.md#ApiRgwBucketSetEncryptionConfigPut) | **Put** /api/rgw/bucket/setEncryptionConfig | 



## ApiRgwBucketBucketDelete

> ApiRgwBucketBucketDelete(ctx, bucket).PurgeObjects(purgeObjects).DaemonName(daemonName).Execute()



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
	bucket := "bucket_example" // string | 
	purgeObjects := "purgeObjects_example" // string |  (optional)
	daemonName := "daemonName_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RgwBucketAPI.ApiRgwBucketBucketDelete(context.Background(), bucket).PurgeObjects(purgeObjects).DaemonName(daemonName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RgwBucketAPI.ApiRgwBucketBucketDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucket** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiRgwBucketBucketDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **purgeObjects** | **string** |  | 
 **daemonName** | **string** |  | 

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


## ApiRgwBucketBucketGet

> ApiRgwBucketBucketGet(ctx, bucket).DaemonName(daemonName).Execute()



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
	bucket := "bucket_example" // string | 
	daemonName := "daemonName_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RgwBucketAPI.ApiRgwBucketBucketGet(context.Background(), bucket).DaemonName(daemonName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RgwBucketAPI.ApiRgwBucketBucketGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucket** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiRgwBucketBucketGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **daemonName** | **string** |  | 

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


## ApiRgwBucketBucketPut

> ApiRgwBucketBucketPut(ctx, bucket).ApiRgwBucketBucketPutRequest(apiRgwBucketBucketPutRequest).Execute()



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
	bucket := "bucket_example" // string | 
	apiRgwBucketBucketPutRequest := *openapiclient.NewApiRgwBucketBucketPutRequest("BucketId_example") // ApiRgwBucketBucketPutRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RgwBucketAPI.ApiRgwBucketBucketPut(context.Background(), bucket).ApiRgwBucketBucketPutRequest(apiRgwBucketBucketPutRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RgwBucketAPI.ApiRgwBucketBucketPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucket** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiRgwBucketBucketPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiRgwBucketBucketPutRequest** | [**ApiRgwBucketBucketPutRequest**](ApiRgwBucketBucketPutRequest.md) |  | 

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


## ApiRgwBucketDeleteEncryptionDelete

> ApiRgwBucketDeleteEncryptionDelete(ctx).BucketName(bucketName).DaemonName(daemonName).Owner(owner).Execute()



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
	bucketName := "bucketName_example" // string | 
	daemonName := "daemonName_example" // string |  (optional)
	owner := "owner_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RgwBucketAPI.ApiRgwBucketDeleteEncryptionDelete(context.Background()).BucketName(bucketName).DaemonName(daemonName).Owner(owner).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RgwBucketAPI.ApiRgwBucketDeleteEncryptionDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiRgwBucketDeleteEncryptionDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bucketName** | **string** |  | 
 **daemonName** | **string** |  | 
 **owner** | **string** |  | 

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


## ApiRgwBucketGet

> ApiRgwBucketGet(ctx).Stats(stats).DaemonName(daemonName).Uid(uid).Execute()



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
	stats := true // bool |  (optional)
	daemonName := "daemonName_example" // string |  (optional)
	uid := "uid_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RgwBucketAPI.ApiRgwBucketGet(context.Background()).Stats(stats).DaemonName(daemonName).Uid(uid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RgwBucketAPI.ApiRgwBucketGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiRgwBucketGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **stats** | **bool** |  | 
 **daemonName** | **string** |  | 
 **uid** | **string** |  | 

### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.ceph.api.v1.1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiRgwBucketGetEncryptionConfigGet

> ApiRgwBucketGetEncryptionConfigGet(ctx).DaemonName(daemonName).Owner(owner).Execute()



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
	daemonName := "daemonName_example" // string |  (optional)
	owner := "owner_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RgwBucketAPI.ApiRgwBucketGetEncryptionConfigGet(context.Background()).DaemonName(daemonName).Owner(owner).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RgwBucketAPI.ApiRgwBucketGetEncryptionConfigGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiRgwBucketGetEncryptionConfigGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **daemonName** | **string** |  | 
 **owner** | **string** |  | 

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


## ApiRgwBucketGetEncryptionGet

> ApiRgwBucketGetEncryptionGet(ctx).BucketName(bucketName).DaemonName(daemonName).Owner(owner).Execute()



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
	bucketName := "bucketName_example" // string | 
	daemonName := "daemonName_example" // string |  (optional)
	owner := "owner_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RgwBucketAPI.ApiRgwBucketGetEncryptionGet(context.Background()).BucketName(bucketName).DaemonName(daemonName).Owner(owner).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RgwBucketAPI.ApiRgwBucketGetEncryptionGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiRgwBucketGetEncryptionGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bucketName** | **string** |  | 
 **daemonName** | **string** |  | 
 **owner** | **string** |  | 

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


## ApiRgwBucketPost

> ApiRgwBucketPost(ctx).ApiRgwBucketPostRequest(apiRgwBucketPostRequest).Execute()



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
	apiRgwBucketPostRequest := *openapiclient.NewApiRgwBucketPostRequest("Bucket_example", "Uid_example") // ApiRgwBucketPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RgwBucketAPI.ApiRgwBucketPost(context.Background()).ApiRgwBucketPostRequest(apiRgwBucketPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RgwBucketAPI.ApiRgwBucketPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiRgwBucketPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiRgwBucketPostRequest** | [**ApiRgwBucketPostRequest**](ApiRgwBucketPostRequest.md) |  | 

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


## ApiRgwBucketSetEncryptionConfigPut

> ApiRgwBucketSetEncryptionConfigPut(ctx).ApiRgwBucketSetEncryptionConfigPutRequest(apiRgwBucketSetEncryptionConfigPutRequest).Execute()



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
	apiRgwBucketSetEncryptionConfigPutRequest := *openapiclient.NewApiRgwBucketSetEncryptionConfigPutRequest() // ApiRgwBucketSetEncryptionConfigPutRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RgwBucketAPI.ApiRgwBucketSetEncryptionConfigPut(context.Background()).ApiRgwBucketSetEncryptionConfigPutRequest(apiRgwBucketSetEncryptionConfigPutRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RgwBucketAPI.ApiRgwBucketSetEncryptionConfigPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiRgwBucketSetEncryptionConfigPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiRgwBucketSetEncryptionConfigPutRequest** | [**ApiRgwBucketSetEncryptionConfigPutRequest**](ApiRgwBucketSetEncryptionConfigPutRequest.md) |  | 

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

