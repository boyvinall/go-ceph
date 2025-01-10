# \RgwSiteAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiRgwSiteGet**](RgwSiteAPI.md#ApiRgwSiteGet) | **Get** /api/rgw/site | 



## ApiRgwSiteGet

> ApiRgwSiteGet(ctx).Query(query).DaemonName(daemonName).Execute()



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
	query := "query_example" // string |  (optional)
	daemonName := "daemonName_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RgwSiteAPI.ApiRgwSiteGet(context.Background()).Query(query).DaemonName(daemonName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RgwSiteAPI.ApiRgwSiteGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiRgwSiteGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **string** |  | 
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

