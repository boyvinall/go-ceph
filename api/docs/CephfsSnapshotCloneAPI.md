# \CephfsSnapshotCloneAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiCephfsSubvolumeSnapshotClonePost**](CephfsSnapshotCloneAPI.md#ApiCephfsSubvolumeSnapshotClonePost) | **Post** /api/cephfs/subvolume/snapshot/clone | Create a clone of a subvolume snapshot



## ApiCephfsSubvolumeSnapshotClonePost

> ApiCephfsSubvolumeSnapshotClonePost(ctx).ApiCephfsSubvolumeSnapshotClonePostRequest(apiCephfsSubvolumeSnapshotClonePostRequest).Execute()

Create a clone of a subvolume snapshot

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
	apiCephfsSubvolumeSnapshotClonePostRequest := *openapiclient.NewApiCephfsSubvolumeSnapshotClonePostRequest("CloneName_example", "SnapName_example", "SubvolName_example", "VolName_example") // ApiCephfsSubvolumeSnapshotClonePostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CephfsSnapshotCloneAPI.ApiCephfsSubvolumeSnapshotClonePost(context.Background()).ApiCephfsSubvolumeSnapshotClonePostRequest(apiCephfsSubvolumeSnapshotClonePostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CephfsSnapshotCloneAPI.ApiCephfsSubvolumeSnapshotClonePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiCephfsSubvolumeSnapshotClonePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiCephfsSubvolumeSnapshotClonePostRequest** | [**ApiCephfsSubvolumeSnapshotClonePostRequest**](ApiCephfsSubvolumeSnapshotClonePostRequest.md) |  | 

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

