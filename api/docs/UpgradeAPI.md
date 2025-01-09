# \UpgradeAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiClusterUpgradeGet**](UpgradeAPI.md#ApiClusterUpgradeGet) | **Get** /api/cluster/upgrade | Get the available versions to upgrade
[**ApiClusterUpgradePausePut**](UpgradeAPI.md#ApiClusterUpgradePausePut) | **Put** /api/cluster/upgrade/pause | Pause the cluster upgrade
[**ApiClusterUpgradeResumePut**](UpgradeAPI.md#ApiClusterUpgradeResumePut) | **Put** /api/cluster/upgrade/resume | Resume the cluster upgrade
[**ApiClusterUpgradeStartPost**](UpgradeAPI.md#ApiClusterUpgradeStartPost) | **Post** /api/cluster/upgrade/start | Start the cluster upgrade
[**ApiClusterUpgradeStatusGet**](UpgradeAPI.md#ApiClusterUpgradeStatusGet) | **Get** /api/cluster/upgrade/status | Get the cluster upgrade status
[**ApiClusterUpgradeStopPut**](UpgradeAPI.md#ApiClusterUpgradeStopPut) | **Put** /api/cluster/upgrade/stop | Stop the cluster upgrade



## ApiClusterUpgradeGet

> ApiClusterUpgradeGet(ctx).Tags(tags).Image(image).ShowAllVersions(showAllVersions).Execute()

Get the available versions to upgrade

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
	tags := true // bool | Show all image tags (optional)
	image := "image_example" // string | Ceph Image (optional)
	showAllVersions := true // bool | Show all available versions (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UpgradeAPI.ApiClusterUpgradeGet(context.Background()).Tags(tags).Image(image).ShowAllVersions(showAllVersions).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UpgradeAPI.ApiClusterUpgradeGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiClusterUpgradeGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tags** | **bool** | Show all image tags | 
 **image** | **string** | Ceph Image | 
 **showAllVersions** | **bool** | Show all available versions | 

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


## ApiClusterUpgradePausePut

> ApiClusterUpgradePausePut(ctx).Execute()

Pause the cluster upgrade

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
	r, err := apiClient.UpgradeAPI.ApiClusterUpgradePausePut(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UpgradeAPI.ApiClusterUpgradePausePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiClusterUpgradePausePutRequest struct via the builder pattern


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


## ApiClusterUpgradeResumePut

> ApiClusterUpgradeResumePut(ctx).Execute()

Resume the cluster upgrade

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
	r, err := apiClient.UpgradeAPI.ApiClusterUpgradeResumePut(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UpgradeAPI.ApiClusterUpgradeResumePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiClusterUpgradeResumePutRequest struct via the builder pattern


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


## ApiClusterUpgradeStartPost

> ApiClusterUpgradeStartPost(ctx).ApiClusterUpgradeStartPostRequest(apiClusterUpgradeStartPostRequest).Execute()

Start the cluster upgrade

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
	apiClusterUpgradeStartPostRequest := *openapiclient.NewApiClusterUpgradeStartPostRequest() // ApiClusterUpgradeStartPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UpgradeAPI.ApiClusterUpgradeStartPost(context.Background()).ApiClusterUpgradeStartPostRequest(apiClusterUpgradeStartPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UpgradeAPI.ApiClusterUpgradeStartPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiClusterUpgradeStartPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiClusterUpgradeStartPostRequest** | [**ApiClusterUpgradeStartPostRequest**](ApiClusterUpgradeStartPostRequest.md) |  | 

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


## ApiClusterUpgradeStatusGet

> ApiClusterUpgradeStatusGet(ctx).Execute()

Get the cluster upgrade status

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
	r, err := apiClient.UpgradeAPI.ApiClusterUpgradeStatusGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UpgradeAPI.ApiClusterUpgradeStatusGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiClusterUpgradeStatusGetRequest struct via the builder pattern


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


## ApiClusterUpgradeStopPut

> ApiClusterUpgradeStopPut(ctx).Execute()

Stop the cluster upgrade

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
	r, err := apiClient.UpgradeAPI.ApiClusterUpgradeStopPut(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UpgradeAPI.ApiClusterUpgradeStopPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiClusterUpgradeStopPutRequest struct via the builder pattern


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

