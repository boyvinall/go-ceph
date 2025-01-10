# \OSDAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiOsdFlagsGet**](OSDAPI.md#ApiOsdFlagsGet) | **Get** /api/osd/flags | Display OSD Flags
[**ApiOsdFlagsIndividualGet**](OSDAPI.md#ApiOsdFlagsIndividualGet) | **Get** /api/osd/flags/individual | Displays individual OSD flags
[**ApiOsdFlagsIndividualPut**](OSDAPI.md#ApiOsdFlagsIndividualPut) | **Put** /api/osd/flags/individual | Sets OSD flags for a subset of individual OSDs.
[**ApiOsdFlagsPut**](OSDAPI.md#ApiOsdFlagsPut) | **Put** /api/osd/flags | Sets OSD flags for the entire cluster.
[**ApiOsdGet**](OSDAPI.md#ApiOsdGet) | **Get** /api/osd | 
[**ApiOsdPost**](OSDAPI.md#ApiOsdPost) | **Post** /api/osd | 
[**ApiOsdSafeToDeleteGet**](OSDAPI.md#ApiOsdSafeToDeleteGet) | **Get** /api/osd/safe_to_delete | 
[**ApiOsdSafeToDestroyGet**](OSDAPI.md#ApiOsdSafeToDestroyGet) | **Get** /api/osd/safe_to_destroy | Check If OSD is Safe to Destroy
[**ApiOsdSettingsGet**](OSDAPI.md#ApiOsdSettingsGet) | **Get** /api/osd/settings | 
[**ApiOsdSvcIdDelete**](OSDAPI.md#ApiOsdSvcIdDelete) | **Delete** /api/osd/{svc_id} | 
[**ApiOsdSvcIdDestroyPost**](OSDAPI.md#ApiOsdSvcIdDestroyPost) | **Post** /api/osd/{svc_id}/destroy | 
[**ApiOsdSvcIdDevicesGet**](OSDAPI.md#ApiOsdSvcIdDevicesGet) | **Get** /api/osd/{svc_id}/devices | 
[**ApiOsdSvcIdGet**](OSDAPI.md#ApiOsdSvcIdGet) | **Get** /api/osd/{svc_id} | 
[**ApiOsdSvcIdHistogramGet**](OSDAPI.md#ApiOsdSvcIdHistogramGet) | **Get** /api/osd/{svc_id}/histogram | 
[**ApiOsdSvcIdMarkPut**](OSDAPI.md#ApiOsdSvcIdMarkPut) | **Put** /api/osd/{svc_id}/mark | Mark OSD flags (out, in, down, lost, ...)
[**ApiOsdSvcIdPurgePost**](OSDAPI.md#ApiOsdSvcIdPurgePost) | **Post** /api/osd/{svc_id}/purge | 
[**ApiOsdSvcIdPut**](OSDAPI.md#ApiOsdSvcIdPut) | **Put** /api/osd/{svc_id} | 
[**ApiOsdSvcIdReweightPost**](OSDAPI.md#ApiOsdSvcIdReweightPost) | **Post** /api/osd/{svc_id}/reweight | 
[**ApiOsdSvcIdScrubPost**](OSDAPI.md#ApiOsdSvcIdScrubPost) | **Post** /api/osd/{svc_id}/scrub | 
[**ApiOsdSvcIdSmartGet**](OSDAPI.md#ApiOsdSvcIdSmartGet) | **Get** /api/osd/{svc_id}/smart | 



## ApiOsdFlagsGet

> ApiOsdFlagsGet200Response ApiOsdFlagsGet(ctx).Execute()

Display OSD Flags

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
	resp, r, err := apiClient.OSDAPI.ApiOsdFlagsGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OSDAPI.ApiOsdFlagsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiOsdFlagsGet`: ApiOsdFlagsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `OSDAPI.ApiOsdFlagsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiOsdFlagsGetRequest struct via the builder pattern


### Return type

[**ApiOsdFlagsGet200Response**](ApiOsdFlagsGet200Response.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.ceph.api.v1.0+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiOsdFlagsIndividualGet

> ApiOsdFlagsIndividualGet200Response ApiOsdFlagsIndividualGet(ctx).Execute()

Displays individual OSD flags

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
	resp, r, err := apiClient.OSDAPI.ApiOsdFlagsIndividualGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OSDAPI.ApiOsdFlagsIndividualGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiOsdFlagsIndividualGet`: ApiOsdFlagsIndividualGet200Response
	fmt.Fprintf(os.Stdout, "Response from `OSDAPI.ApiOsdFlagsIndividualGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiOsdFlagsIndividualGetRequest struct via the builder pattern


### Return type

[**ApiOsdFlagsIndividualGet200Response**](ApiOsdFlagsIndividualGet200Response.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.ceph.api.v1.0+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiOsdFlagsIndividualPut

> ApiOsdFlagsIndividualPut200Response ApiOsdFlagsIndividualPut(ctx).ApiOsdFlagsIndividualPutRequest(apiOsdFlagsIndividualPutRequest).Execute()

Sets OSD flags for a subset of individual OSDs.



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
	apiOsdFlagsIndividualPutRequest := *openapiclient.NewApiOsdFlagsIndividualPutRequest(*openapiclient.NewApiOsdFlagsIndividualPutRequestFlags(), []int32{int32(123)}) // ApiOsdFlagsIndividualPutRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OSDAPI.ApiOsdFlagsIndividualPut(context.Background()).ApiOsdFlagsIndividualPutRequest(apiOsdFlagsIndividualPutRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OSDAPI.ApiOsdFlagsIndividualPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiOsdFlagsIndividualPut`: ApiOsdFlagsIndividualPut200Response
	fmt.Fprintf(os.Stdout, "Response from `OSDAPI.ApiOsdFlagsIndividualPut`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiOsdFlagsIndividualPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiOsdFlagsIndividualPutRequest** | [**ApiOsdFlagsIndividualPutRequest**](ApiOsdFlagsIndividualPutRequest.md) |  | 

### Return type

[**ApiOsdFlagsIndividualPut200Response**](ApiOsdFlagsIndividualPut200Response.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.ceph.api.v1.0+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiOsdFlagsPut

> ApiOsdFlagsGet200Response ApiOsdFlagsPut(ctx).ApiOsdFlagsPutRequest(apiOsdFlagsPutRequest).Execute()

Sets OSD flags for the entire cluster.



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
	apiOsdFlagsPutRequest := *openapiclient.NewApiOsdFlagsPutRequest([]string{"Flags_example"}) // ApiOsdFlagsPutRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OSDAPI.ApiOsdFlagsPut(context.Background()).ApiOsdFlagsPutRequest(apiOsdFlagsPutRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OSDAPI.ApiOsdFlagsPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiOsdFlagsPut`: ApiOsdFlagsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `OSDAPI.ApiOsdFlagsPut`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiOsdFlagsPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiOsdFlagsPutRequest** | [**ApiOsdFlagsPutRequest**](ApiOsdFlagsPutRequest.md) |  | 

### Return type

[**ApiOsdFlagsGet200Response**](ApiOsdFlagsGet200Response.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.ceph.api.v1.0+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiOsdGet

> ApiOsdGet(ctx).Offset(offset).Limit(limit).Search(search).Sort(sort).Execute()



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
	offset := int32(56) // int32 |  (optional)
	limit := int32(56) // int32 |  (optional)
	search := "search_example" // string |  (optional)
	sort := "sort_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OSDAPI.ApiOsdGet(context.Background()).Offset(offset).Limit(limit).Search(search).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OSDAPI.ApiOsdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiOsdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **int32** |  | 
 **limit** | **int32** |  | 
 **search** | **string** |  | 
 **sort** | **string** |  | 

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


## ApiOsdPost

> ApiOsdPost(ctx).ApiOsdPostRequest(apiOsdPostRequest).Execute()



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
	apiOsdPostRequest := *openapiclient.NewApiOsdPostRequest("Data_example", "Method_example", "TrackingId_example") // ApiOsdPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OSDAPI.ApiOsdPost(context.Background()).ApiOsdPostRequest(apiOsdPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OSDAPI.ApiOsdPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiOsdPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiOsdPostRequest** | [**ApiOsdPostRequest**](ApiOsdPostRequest.md) |  | 

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


## ApiOsdSafeToDeleteGet

> ApiOsdSafeToDeleteGet(ctx).SvcIds(svcIds).Execute()





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
	svcIds := "svcIds_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OSDAPI.ApiOsdSafeToDeleteGet(context.Background()).SvcIds(svcIds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OSDAPI.ApiOsdSafeToDeleteGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiOsdSafeToDeleteGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **svcIds** | **string** |  | 

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


## ApiOsdSafeToDestroyGet

> ApiOsdSafeToDestroyGet200Response ApiOsdSafeToDestroyGet(ctx).Ids(ids).Execute()

Check If OSD is Safe to Destroy



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
	ids := "ids_example" // string | OSD Service Identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OSDAPI.ApiOsdSafeToDestroyGet(context.Background()).Ids(ids).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OSDAPI.ApiOsdSafeToDestroyGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiOsdSafeToDestroyGet`: ApiOsdSafeToDestroyGet200Response
	fmt.Fprintf(os.Stdout, "Response from `OSDAPI.ApiOsdSafeToDestroyGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiOsdSafeToDestroyGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | **string** | OSD Service Identifier | 

### Return type

[**ApiOsdSafeToDestroyGet200Response**](ApiOsdSafeToDestroyGet200Response.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.ceph.api.v1.0+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiOsdSettingsGet

> ApiOsdSettingsGet(ctx).Execute()



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
	r, err := apiClient.OSDAPI.ApiOsdSettingsGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OSDAPI.ApiOsdSettingsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiOsdSettingsGetRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.ceph.api.v0.1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiOsdSvcIdDelete

> ApiOsdSvcIdDelete(ctx, svcId).PreserveId(preserveId).Force(force).Execute()



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
	svcId := "svcId_example" // string | 
	preserveId := "preserveId_example" // string |  (optional)
	force := "force_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OSDAPI.ApiOsdSvcIdDelete(context.Background(), svcId).PreserveId(preserveId).Force(force).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OSDAPI.ApiOsdSvcIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**svcId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiOsdSvcIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **preserveId** | **string** |  | 
 **force** | **string** |  | 

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


## ApiOsdSvcIdDestroyPost

> ApiOsdSvcIdDestroyPost(ctx, svcId).Execute()





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
	svcId := "svcId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OSDAPI.ApiOsdSvcIdDestroyPost(context.Background(), svcId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OSDAPI.ApiOsdSvcIdDestroyPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**svcId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiOsdSvcIdDestroyPostRequest struct via the builder pattern


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


## ApiOsdSvcIdDevicesGet

> ApiOsdSvcIdDevicesGet(ctx, svcId).Execute()



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
	svcId := "svcId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OSDAPI.ApiOsdSvcIdDevicesGet(context.Background(), svcId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OSDAPI.ApiOsdSvcIdDevicesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**svcId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiOsdSvcIdDevicesGetRequest struct via the builder pattern


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


## ApiOsdSvcIdGet

> ApiOsdSvcIdGet(ctx, svcId).Execute()





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
	svcId := "svcId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OSDAPI.ApiOsdSvcIdGet(context.Background(), svcId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OSDAPI.ApiOsdSvcIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**svcId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiOsdSvcIdGetRequest struct via the builder pattern


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


## ApiOsdSvcIdHistogramGet

> ApiOsdSvcIdHistogramGet(ctx, svcId).Execute()





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
	svcId := "svcId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OSDAPI.ApiOsdSvcIdHistogramGet(context.Background(), svcId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OSDAPI.ApiOsdSvcIdHistogramGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**svcId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiOsdSvcIdHistogramGetRequest struct via the builder pattern


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


## ApiOsdSvcIdMarkPut

> ApiOsdSvcIdMarkPut(ctx, svcId).ApiOsdSvcIdMarkPutRequest(apiOsdSvcIdMarkPutRequest).Execute()

Mark OSD flags (out, in, down, lost, ...)



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
	svcId := "svcId_example" // string | SVC ID
	apiOsdSvcIdMarkPutRequest := *openapiclient.NewApiOsdSvcIdMarkPutRequest("Action_example") // ApiOsdSvcIdMarkPutRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OSDAPI.ApiOsdSvcIdMarkPut(context.Background(), svcId).ApiOsdSvcIdMarkPutRequest(apiOsdSvcIdMarkPutRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OSDAPI.ApiOsdSvcIdMarkPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**svcId** | **string** | SVC ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiOsdSvcIdMarkPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiOsdSvcIdMarkPutRequest** | [**ApiOsdSvcIdMarkPutRequest**](ApiOsdSvcIdMarkPutRequest.md) |  | 

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


## ApiOsdSvcIdPurgePost

> ApiOsdSvcIdPurgePost(ctx, svcId).Execute()





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
	svcId := "svcId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OSDAPI.ApiOsdSvcIdPurgePost(context.Background(), svcId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OSDAPI.ApiOsdSvcIdPurgePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**svcId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiOsdSvcIdPurgePostRequest struct via the builder pattern


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


## ApiOsdSvcIdPut

> ApiOsdSvcIdPut(ctx, svcId).ApiOsdSvcIdPutRequest(apiOsdSvcIdPutRequest).Execute()



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
	svcId := "svcId_example" // string | 
	apiOsdSvcIdPutRequest := *openapiclient.NewApiOsdSvcIdPutRequest("DeviceClass_example") // ApiOsdSvcIdPutRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OSDAPI.ApiOsdSvcIdPut(context.Background(), svcId).ApiOsdSvcIdPutRequest(apiOsdSvcIdPutRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OSDAPI.ApiOsdSvcIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**svcId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiOsdSvcIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiOsdSvcIdPutRequest** | [**ApiOsdSvcIdPutRequest**](ApiOsdSvcIdPutRequest.md) |  | 

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


## ApiOsdSvcIdReweightPost

> ApiOsdSvcIdReweightPost(ctx, svcId).ApiOsdSvcIdReweightPostRequest(apiOsdSvcIdReweightPostRequest).Execute()





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
	svcId := "svcId_example" // string | 
	apiOsdSvcIdReweightPostRequest := *openapiclient.NewApiOsdSvcIdReweightPostRequest("Weight_example") // ApiOsdSvcIdReweightPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OSDAPI.ApiOsdSvcIdReweightPost(context.Background(), svcId).ApiOsdSvcIdReweightPostRequest(apiOsdSvcIdReweightPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OSDAPI.ApiOsdSvcIdReweightPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**svcId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiOsdSvcIdReweightPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiOsdSvcIdReweightPostRequest** | [**ApiOsdSvcIdReweightPostRequest**](ApiOsdSvcIdReweightPostRequest.md) |  | 

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


## ApiOsdSvcIdScrubPost

> ApiOsdSvcIdScrubPost(ctx, svcId).Deep(deep).ApiOsdSvcIdScrubPostRequest(apiOsdSvcIdScrubPostRequest).Execute()



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
	svcId := "svcId_example" // string | 
	deep := true // bool |  (optional)
	apiOsdSvcIdScrubPostRequest := *openapiclient.NewApiOsdSvcIdScrubPostRequest() // ApiOsdSvcIdScrubPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OSDAPI.ApiOsdSvcIdScrubPost(context.Background(), svcId).Deep(deep).ApiOsdSvcIdScrubPostRequest(apiOsdSvcIdScrubPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OSDAPI.ApiOsdSvcIdScrubPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**svcId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiOsdSvcIdScrubPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deep** | **bool** |  | 
 **apiOsdSvcIdScrubPostRequest** | [**ApiOsdSvcIdScrubPostRequest**](ApiOsdSvcIdScrubPostRequest.md) |  | 

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


## ApiOsdSvcIdSmartGet

> ApiOsdSvcIdSmartGet(ctx, svcId).Execute()



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
	svcId := "svcId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OSDAPI.ApiOsdSvcIdSmartGet(context.Background(), svcId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OSDAPI.ApiOsdSvcIdSmartGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**svcId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiOsdSvcIdSmartGetRequest struct via the builder pattern


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

