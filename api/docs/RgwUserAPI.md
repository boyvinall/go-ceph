# \RgwUserAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiRgwUserGet**](RgwUserAPI.md#ApiRgwUserGet) | **Get** /api/rgw/user | Display RGW Users
[**ApiRgwUserGetEmailsGet**](RgwUserAPI.md#ApiRgwUserGetEmailsGet) | **Get** /api/rgw/user/get_emails | 
[**ApiRgwUserPost**](RgwUserAPI.md#ApiRgwUserPost) | **Post** /api/rgw/user | 
[**ApiRgwUserUidCapabilityDelete**](RgwUserAPI.md#ApiRgwUserUidCapabilityDelete) | **Delete** /api/rgw/user/{uid}/capability | 
[**ApiRgwUserUidCapabilityPost**](RgwUserAPI.md#ApiRgwUserUidCapabilityPost) | **Post** /api/rgw/user/{uid}/capability | 
[**ApiRgwUserUidDelete**](RgwUserAPI.md#ApiRgwUserUidDelete) | **Delete** /api/rgw/user/{uid} | 
[**ApiRgwUserUidGet**](RgwUserAPI.md#ApiRgwUserUidGet) | **Get** /api/rgw/user/{uid} | 
[**ApiRgwUserUidKeyDelete**](RgwUserAPI.md#ApiRgwUserUidKeyDelete) | **Delete** /api/rgw/user/{uid}/key | 
[**ApiRgwUserUidKeyPost**](RgwUserAPI.md#ApiRgwUserUidKeyPost) | **Post** /api/rgw/user/{uid}/key | 
[**ApiRgwUserUidPut**](RgwUserAPI.md#ApiRgwUserUidPut) | **Put** /api/rgw/user/{uid} | 
[**ApiRgwUserUidQuotaGet**](RgwUserAPI.md#ApiRgwUserUidQuotaGet) | **Get** /api/rgw/user/{uid}/quota | 
[**ApiRgwUserUidQuotaPut**](RgwUserAPI.md#ApiRgwUserUidQuotaPut) | **Put** /api/rgw/user/{uid}/quota | 
[**ApiRgwUserUidSubuserPost**](RgwUserAPI.md#ApiRgwUserUidSubuserPost) | **Post** /api/rgw/user/{uid}/subuser | 
[**ApiRgwUserUidSubuserSubuserDelete**](RgwUserAPI.md#ApiRgwUserUidSubuserSubuserDelete) | **Delete** /api/rgw/user/{uid}/subuser/{subuser} | 



## ApiRgwUserGet

> ApiRgwUserGet200Response ApiRgwUserGet(ctx).DaemonName(daemonName).Execute()

Display RGW Users

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
	daemonName := "daemonName_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RgwUserAPI.ApiRgwUserGet(context.Background()).DaemonName(daemonName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RgwUserAPI.ApiRgwUserGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiRgwUserGet`: ApiRgwUserGet200Response
	fmt.Fprintf(os.Stdout, "Response from `RgwUserAPI.ApiRgwUserGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiRgwUserGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **daemonName** | **string** |  | 

### Return type

[**ApiRgwUserGet200Response**](ApiRgwUserGet200Response.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.ceph.api.v1.0+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiRgwUserGetEmailsGet

> ApiRgwUserGetEmailsGet(ctx).DaemonName(daemonName).Execute()



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
	daemonName := "daemonName_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RgwUserAPI.ApiRgwUserGetEmailsGet(context.Background()).DaemonName(daemonName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RgwUserAPI.ApiRgwUserGetEmailsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiRgwUserGetEmailsGetRequest struct via the builder pattern


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


## ApiRgwUserPost

> ApiRgwUserPost(ctx).ApiRgwUserPostRequest(apiRgwUserPostRequest).Execute()



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
	apiRgwUserPostRequest := *openapiclient.NewApiRgwUserPostRequest("DisplayName_example", "Uid_example") // ApiRgwUserPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RgwUserAPI.ApiRgwUserPost(context.Background()).ApiRgwUserPostRequest(apiRgwUserPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RgwUserAPI.ApiRgwUserPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiRgwUserPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiRgwUserPostRequest** | [**ApiRgwUserPostRequest**](ApiRgwUserPostRequest.md) |  | 

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


## ApiRgwUserUidCapabilityDelete

> ApiRgwUserUidCapabilityDelete(ctx, uid).Type_(type_).Perm(perm).DaemonName(daemonName).Execute()



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
	uid := "uid_example" // string | 
	type_ := "type__example" // string | 
	perm := "perm_example" // string | 
	daemonName := "daemonName_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RgwUserAPI.ApiRgwUserUidCapabilityDelete(context.Background(), uid).Type_(type_).Perm(perm).DaemonName(daemonName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RgwUserAPI.ApiRgwUserUidCapabilityDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiRgwUserUidCapabilityDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | **string** |  | 
 **perm** | **string** |  | 
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


## ApiRgwUserUidCapabilityPost

> ApiRgwUserUidCapabilityPost(ctx, uid).ApiRgwUserUidCapabilityPostRequest(apiRgwUserUidCapabilityPostRequest).Execute()



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
	uid := "uid_example" // string | 
	apiRgwUserUidCapabilityPostRequest := *openapiclient.NewApiRgwUserUidCapabilityPostRequest("Perm_example", "Type_example") // ApiRgwUserUidCapabilityPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RgwUserAPI.ApiRgwUserUidCapabilityPost(context.Background(), uid).ApiRgwUserUidCapabilityPostRequest(apiRgwUserUidCapabilityPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RgwUserAPI.ApiRgwUserUidCapabilityPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiRgwUserUidCapabilityPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiRgwUserUidCapabilityPostRequest** | [**ApiRgwUserUidCapabilityPostRequest**](ApiRgwUserUidCapabilityPostRequest.md) |  | 

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


## ApiRgwUserUidDelete

> ApiRgwUserUidDelete(ctx, uid).DaemonName(daemonName).Execute()



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
	uid := "uid_example" // string | 
	daemonName := "daemonName_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RgwUserAPI.ApiRgwUserUidDelete(context.Background(), uid).DaemonName(daemonName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RgwUserAPI.ApiRgwUserUidDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiRgwUserUidDeleteRequest struct via the builder pattern


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


## ApiRgwUserUidGet

> ApiRgwUserUidGet(ctx, uid).DaemonName(daemonName).Stats(stats).Execute()



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
	uid := "uid_example" // string | 
	daemonName := "daemonName_example" // string |  (optional)
	stats := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RgwUserAPI.ApiRgwUserUidGet(context.Background(), uid).DaemonName(daemonName).Stats(stats).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RgwUserAPI.ApiRgwUserUidGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiRgwUserUidGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **daemonName** | **string** |  | 
 **stats** | **bool** |  | 

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


## ApiRgwUserUidKeyDelete

> ApiRgwUserUidKeyDelete(ctx, uid).KeyType(keyType).Subuser(subuser).AccessKey(accessKey).DaemonName(daemonName).Execute()



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
	uid := "uid_example" // string | 
	keyType := "keyType_example" // string |  (optional)
	subuser := "subuser_example" // string |  (optional)
	accessKey := "accessKey_example" // string |  (optional)
	daemonName := "daemonName_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RgwUserAPI.ApiRgwUserUidKeyDelete(context.Background(), uid).KeyType(keyType).Subuser(subuser).AccessKey(accessKey).DaemonName(daemonName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RgwUserAPI.ApiRgwUserUidKeyDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiRgwUserUidKeyDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **keyType** | **string** |  | 
 **subuser** | **string** |  | 
 **accessKey** | **string** |  | 
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


## ApiRgwUserUidKeyPost

> ApiRgwUserUidKeyPost(ctx, uid).ApiRgwUserUidKeyPostRequest(apiRgwUserUidKeyPostRequest).Execute()



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
	uid := "uid_example" // string | 
	apiRgwUserUidKeyPostRequest := *openapiclient.NewApiRgwUserUidKeyPostRequest() // ApiRgwUserUidKeyPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RgwUserAPI.ApiRgwUserUidKeyPost(context.Background(), uid).ApiRgwUserUidKeyPostRequest(apiRgwUserUidKeyPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RgwUserAPI.ApiRgwUserUidKeyPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiRgwUserUidKeyPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiRgwUserUidKeyPostRequest** | [**ApiRgwUserUidKeyPostRequest**](ApiRgwUserUidKeyPostRequest.md) |  | 

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


## ApiRgwUserUidPut

> ApiRgwUserUidPut(ctx, uid).ApiRgwUserUidPutRequest(apiRgwUserUidPutRequest).Execute()



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
	uid := "uid_example" // string | 
	apiRgwUserUidPutRequest := *openapiclient.NewApiRgwUserUidPutRequest() // ApiRgwUserUidPutRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RgwUserAPI.ApiRgwUserUidPut(context.Background(), uid).ApiRgwUserUidPutRequest(apiRgwUserUidPutRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RgwUserAPI.ApiRgwUserUidPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiRgwUserUidPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiRgwUserUidPutRequest** | [**ApiRgwUserUidPutRequest**](ApiRgwUserUidPutRequest.md) |  | 

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


## ApiRgwUserUidQuotaGet

> ApiRgwUserUidQuotaGet(ctx, uid).DaemonName(daemonName).Execute()



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
	uid := "uid_example" // string | 
	daemonName := "daemonName_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RgwUserAPI.ApiRgwUserUidQuotaGet(context.Background(), uid).DaemonName(daemonName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RgwUserAPI.ApiRgwUserUidQuotaGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiRgwUserUidQuotaGetRequest struct via the builder pattern


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


## ApiRgwUserUidQuotaPut

> ApiRgwUserUidQuotaPut(ctx, uid).ApiRgwUserUidQuotaPutRequest(apiRgwUserUidQuotaPutRequest).Execute()



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
	uid := "uid_example" // string | 
	apiRgwUserUidQuotaPutRequest := *openapiclient.NewApiRgwUserUidQuotaPutRequest("Enabled_example", "MaxObjects_example", int32(123), "QuotaType_example") // ApiRgwUserUidQuotaPutRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RgwUserAPI.ApiRgwUserUidQuotaPut(context.Background(), uid).ApiRgwUserUidQuotaPutRequest(apiRgwUserUidQuotaPutRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RgwUserAPI.ApiRgwUserUidQuotaPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiRgwUserUidQuotaPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiRgwUserUidQuotaPutRequest** | [**ApiRgwUserUidQuotaPutRequest**](ApiRgwUserUidQuotaPutRequest.md) |  | 

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


## ApiRgwUserUidSubuserPost

> ApiRgwUserUidSubuserPost(ctx, uid).ApiRgwUserUidSubuserPostRequest(apiRgwUserUidSubuserPostRequest).Execute()



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
	uid := "uid_example" // string | 
	apiRgwUserUidSubuserPostRequest := *openapiclient.NewApiRgwUserUidSubuserPostRequest("Access_example", "Subuser_example") // ApiRgwUserUidSubuserPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RgwUserAPI.ApiRgwUserUidSubuserPost(context.Background(), uid).ApiRgwUserUidSubuserPostRequest(apiRgwUserUidSubuserPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RgwUserAPI.ApiRgwUserUidSubuserPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiRgwUserUidSubuserPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiRgwUserUidSubuserPostRequest** | [**ApiRgwUserUidSubuserPostRequest**](ApiRgwUserUidSubuserPostRequest.md) |  | 

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


## ApiRgwUserUidSubuserSubuserDelete

> ApiRgwUserUidSubuserSubuserDelete(ctx, uid, subuser).PurgeKeys(purgeKeys).DaemonName(daemonName).Execute()





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
	uid := "uid_example" // string | 
	subuser := "subuser_example" // string | 
	purgeKeys := "purgeKeys_example" // string |  (optional)
	daemonName := "daemonName_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RgwUserAPI.ApiRgwUserUidSubuserSubuserDelete(context.Background(), uid, subuser).PurgeKeys(purgeKeys).DaemonName(daemonName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RgwUserAPI.ApiRgwUserUidSubuserSubuserDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **string** |  | 
**subuser** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiRgwUserUidSubuserSubuserDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **purgeKeys** | **string** |  | 
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

