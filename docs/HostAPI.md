# \HostAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiHostGet**](HostAPI.md#ApiHostGet) | **Get** /api/host | List Host Specifications
[**ApiHostHostnameDaemonsGet**](HostAPI.md#ApiHostHostnameDaemonsGet) | **Get** /api/host/{hostname}/daemons | 
[**ApiHostHostnameDelete**](HostAPI.md#ApiHostHostnameDelete) | **Delete** /api/host/{hostname} | 
[**ApiHostHostnameDevicesGet**](HostAPI.md#ApiHostHostnameDevicesGet) | **Get** /api/host/{hostname}/devices | 
[**ApiHostHostnameGet**](HostAPI.md#ApiHostHostnameGet) | **Get** /api/host/{hostname} | 
[**ApiHostHostnameIdentifyDevicePost**](HostAPI.md#ApiHostHostnameIdentifyDevicePost) | **Post** /api/host/{hostname}/identify_device | 
[**ApiHostHostnameInventoryGet**](HostAPI.md#ApiHostHostnameInventoryGet) | **Get** /api/host/{hostname}/inventory | Get inventory of a host
[**ApiHostHostnamePut**](HostAPI.md#ApiHostHostnamePut) | **Put** /api/host/{hostname} | 
[**ApiHostHostnameSmartGet**](HostAPI.md#ApiHostHostnameSmartGet) | **Get** /api/host/{hostname}/smart | 
[**ApiHostPost**](HostAPI.md#ApiHostPost) | **Post** /api/host | 



## ApiHostGet

> ApiHostGet200Response ApiHostGet(ctx).Sources(sources).Facts(facts).Offset(offset).Limit(limit).Search(search).Sort(sort).Execute()

List Host Specifications

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
	sources := "sources_example" // string | Host Sources (optional)
	facts := true // bool | Host Facts (optional)
	offset := int32(56) // int32 |  (optional)
	limit := int32(56) // int32 |  (optional)
	search := "search_example" // string |  (optional)
	sort := "sort_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HostAPI.ApiHostGet(context.Background()).Sources(sources).Facts(facts).Offset(offset).Limit(limit).Search(search).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostAPI.ApiHostGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiHostGet`: ApiHostGet200Response
	fmt.Fprintf(os.Stdout, "Response from `HostAPI.ApiHostGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiHostGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sources** | **string** | Host Sources | 
 **facts** | **bool** | Host Facts | 
 **offset** | **int32** |  | 
 **limit** | **int32** |  | 
 **search** | **string** |  | 
 **sort** | **string** |  | 

### Return type

[**ApiHostGet200Response**](ApiHostGet200Response.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.ceph.api.v1.3+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiHostHostnameDaemonsGet

> ApiHostHostnameDaemonsGet(ctx, hostname).Execute()



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
	hostname := "hostname_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.HostAPI.ApiHostHostnameDaemonsGet(context.Background(), hostname).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostAPI.ApiHostHostnameDaemonsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostname** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiHostHostnameDaemonsGetRequest struct via the builder pattern


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


## ApiHostHostnameDelete

> ApiHostHostnameDelete(ctx, hostname).Execute()



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
	hostname := "hostname_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.HostAPI.ApiHostHostnameDelete(context.Background(), hostname).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostAPI.ApiHostHostnameDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostname** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiHostHostnameDeleteRequest struct via the builder pattern


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


## ApiHostHostnameDevicesGet

> ApiHostHostnameDevicesGet(ctx, hostname).Execute()



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
	hostname := "hostname_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.HostAPI.ApiHostHostnameDevicesGet(context.Background(), hostname).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostAPI.ApiHostHostnameDevicesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostname** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiHostHostnameDevicesGetRequest struct via the builder pattern


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


## ApiHostHostnameGet

> ApiHostHostnameGet(ctx, hostname).Execute()





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
	hostname := "hostname_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.HostAPI.ApiHostHostnameGet(context.Background(), hostname).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostAPI.ApiHostHostnameGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostname** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiHostHostnameGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.ceph.api.v1.2+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiHostHostnameIdentifyDevicePost

> ApiHostHostnameIdentifyDevicePost(ctx, hostname).ApiHostHostnameIdentifyDevicePostRequest(apiHostHostnameIdentifyDevicePostRequest).Execute()





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
	hostname := "hostname_example" // string | 
	apiHostHostnameIdentifyDevicePostRequest := *openapiclient.NewApiHostHostnameIdentifyDevicePostRequest("Device_example", "Duration_example") // ApiHostHostnameIdentifyDevicePostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.HostAPI.ApiHostHostnameIdentifyDevicePost(context.Background(), hostname).ApiHostHostnameIdentifyDevicePostRequest(apiHostHostnameIdentifyDevicePostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostAPI.ApiHostHostnameIdentifyDevicePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostname** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiHostHostnameIdentifyDevicePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiHostHostnameIdentifyDevicePostRequest** | [**ApiHostHostnameIdentifyDevicePostRequest**](ApiHostHostnameIdentifyDevicePostRequest.md) |  | 

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


## ApiHostHostnameInventoryGet

> ApiHostHostnameInventoryGet200Response ApiHostHostnameInventoryGet(ctx, hostname).Refresh(refresh).Execute()

Get inventory of a host

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
	hostname := "hostname_example" // string | Hostname
	refresh := "refresh_example" // string | Trigger asynchronous refresh (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HostAPI.ApiHostHostnameInventoryGet(context.Background(), hostname).Refresh(refresh).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostAPI.ApiHostHostnameInventoryGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiHostHostnameInventoryGet`: ApiHostHostnameInventoryGet200Response
	fmt.Fprintf(os.Stdout, "Response from `HostAPI.ApiHostHostnameInventoryGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostname** | **string** | Hostname | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiHostHostnameInventoryGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **refresh** | **string** | Trigger asynchronous refresh | 

### Return type

[**ApiHostHostnameInventoryGet200Response**](ApiHostHostnameInventoryGet200Response.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.ceph.api.v1.0+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiHostHostnamePut

> map[string]interface{} ApiHostHostnamePut(ctx, hostname).ApiHostHostnamePutRequest(apiHostHostnamePutRequest).Execute()





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
	hostname := "hostname_example" // string | Hostname
	apiHostHostnamePutRequest := *openapiclient.NewApiHostHostnamePutRequest() // ApiHostHostnamePutRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HostAPI.ApiHostHostnamePut(context.Background(), hostname).ApiHostHostnamePutRequest(apiHostHostnamePutRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostAPI.ApiHostHostnamePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiHostHostnamePut`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `HostAPI.ApiHostHostnamePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostname** | **string** | Hostname | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiHostHostnamePutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiHostHostnamePutRequest** | [**ApiHostHostnamePutRequest**](ApiHostHostnamePutRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.ceph.api.v0.1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiHostHostnameSmartGet

> ApiHostHostnameSmartGet(ctx, hostname).Execute()



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
	hostname := "hostname_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.HostAPI.ApiHostHostnameSmartGet(context.Background(), hostname).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostAPI.ApiHostHostnameSmartGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostname** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiHostHostnameSmartGetRequest struct via the builder pattern


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


## ApiHostPost

> ApiHostPost(ctx).ApiHostPostRequest(apiHostPostRequest).Execute()



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
	apiHostPostRequest := *openapiclient.NewApiHostPostRequest("Hostname_example") // ApiHostPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.HostAPI.ApiHostPost(context.Background()).ApiHostPostRequest(apiHostPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostAPI.ApiHostPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiHostPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiHostPostRequest** | [**ApiHostPostRequest**](ApiHostPostRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.ceph.api.v0.1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

