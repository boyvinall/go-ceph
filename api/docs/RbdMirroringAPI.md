# \RbdMirroringAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiBlockMirroringSiteNameGet**](RbdMirroringAPI.md#ApiBlockMirroringSiteNameGet) | **Get** /api/block/mirroring/site_name | Display Rbd Mirroring sitename
[**ApiBlockMirroringSiteNamePut**](RbdMirroringAPI.md#ApiBlockMirroringSiteNamePut) | **Put** /api/block/mirroring/site_name | 



## ApiBlockMirroringSiteNameGet

> ApiBlockMirroringSiteNameGet200Response ApiBlockMirroringSiteNameGet(ctx).Execute()

Display Rbd Mirroring sitename

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
	resp, r, err := apiClient.RbdMirroringAPI.ApiBlockMirroringSiteNameGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RbdMirroringAPI.ApiBlockMirroringSiteNameGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiBlockMirroringSiteNameGet`: ApiBlockMirroringSiteNameGet200Response
	fmt.Fprintf(os.Stdout, "Response from `RbdMirroringAPI.ApiBlockMirroringSiteNameGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiBlockMirroringSiteNameGetRequest struct via the builder pattern


### Return type

[**ApiBlockMirroringSiteNameGet200Response**](ApiBlockMirroringSiteNameGet200Response.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.ceph.api.v1.0+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiBlockMirroringSiteNamePut

> ApiBlockMirroringSiteNamePut(ctx).ApiBlockMirroringSiteNamePutRequest(apiBlockMirroringSiteNamePutRequest).Execute()



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
	apiBlockMirroringSiteNamePutRequest := *openapiclient.NewApiBlockMirroringSiteNamePutRequest("SiteName_example") // ApiBlockMirroringSiteNamePutRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RbdMirroringAPI.ApiBlockMirroringSiteNamePut(context.Background()).ApiBlockMirroringSiteNamePutRequest(apiBlockMirroringSiteNamePutRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RbdMirroringAPI.ApiBlockMirroringSiteNamePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiBlockMirroringSiteNamePutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiBlockMirroringSiteNamePutRequest** | [**ApiBlockMirroringSiteNamePutRequest**](ApiBlockMirroringSiteNamePutRequest.md) |  | 

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

