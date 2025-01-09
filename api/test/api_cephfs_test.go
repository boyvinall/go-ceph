/*
Ceph REST API

Testing CephfsAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package ceph

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/boyvinall/go-ceph"
)

func Test_ceph_CephfsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test CephfsAPIService ApiCephfsAuthPut", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.CephfsAPI.ApiCephfsAuthPut(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CephfsAPIService ApiCephfsFsIdClientClientIdDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var fsId string
		var clientId string

		httpRes, err := apiClient.CephfsAPI.ApiCephfsFsIdClientClientIdDelete(context.Background(), fsId, clientId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CephfsAPIService ApiCephfsFsIdClientsGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var fsId string

		httpRes, err := apiClient.CephfsAPI.ApiCephfsFsIdClientsGet(context.Background(), fsId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CephfsAPIService ApiCephfsFsIdGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var fsId string

		httpRes, err := apiClient.CephfsAPI.ApiCephfsFsIdGet(context.Background(), fsId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CephfsAPIService ApiCephfsFsIdGetRootDirectoryGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var fsId string

		httpRes, err := apiClient.CephfsAPI.ApiCephfsFsIdGetRootDirectoryGet(context.Background(), fsId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CephfsAPIService ApiCephfsFsIdLsDirGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var fsId string

		httpRes, err := apiClient.CephfsAPI.ApiCephfsFsIdLsDirGet(context.Background(), fsId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CephfsAPIService ApiCephfsFsIdMdsCountersGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var fsId string

		httpRes, err := apiClient.CephfsAPI.ApiCephfsFsIdMdsCountersGet(context.Background(), fsId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CephfsAPIService ApiCephfsFsIdQuotaGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var fsId string

		resp, httpRes, err := apiClient.CephfsAPI.ApiCephfsFsIdQuotaGet(context.Background(), fsId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CephfsAPIService ApiCephfsFsIdQuotaPut", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var fsId string

		httpRes, err := apiClient.CephfsAPI.ApiCephfsFsIdQuotaPut(context.Background(), fsId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CephfsAPIService ApiCephfsFsIdRenamePathPut", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var fsId string

		httpRes, err := apiClient.CephfsAPI.ApiCephfsFsIdRenamePathPut(context.Background(), fsId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CephfsAPIService ApiCephfsFsIdSnapshotDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var fsId string

		httpRes, err := apiClient.CephfsAPI.ApiCephfsFsIdSnapshotDelete(context.Background(), fsId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CephfsAPIService ApiCephfsFsIdSnapshotPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var fsId string

		httpRes, err := apiClient.CephfsAPI.ApiCephfsFsIdSnapshotPost(context.Background(), fsId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CephfsAPIService ApiCephfsFsIdStatfsGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var fsId string

		resp, httpRes, err := apiClient.CephfsAPI.ApiCephfsFsIdStatfsGet(context.Background(), fsId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CephfsAPIService ApiCephfsFsIdTreeDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var fsId string

		httpRes, err := apiClient.CephfsAPI.ApiCephfsFsIdTreeDelete(context.Background(), fsId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CephfsAPIService ApiCephfsFsIdTreePost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var fsId string

		httpRes, err := apiClient.CephfsAPI.ApiCephfsFsIdTreePost(context.Background(), fsId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CephfsAPIService ApiCephfsFsIdUnlinkDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var fsId string

		httpRes, err := apiClient.CephfsAPI.ApiCephfsFsIdUnlinkDelete(context.Background(), fsId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CephfsAPIService ApiCephfsFsIdWriteToFilePost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var fsId string

		httpRes, err := apiClient.CephfsAPI.ApiCephfsFsIdWriteToFilePost(context.Background(), fsId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CephfsAPIService ApiCephfsGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.CephfsAPI.ApiCephfsGet(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CephfsAPIService ApiCephfsPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.CephfsAPI.ApiCephfsPost(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CephfsAPIService ApiCephfsRemoveNameDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var name string

		httpRes, err := apiClient.CephfsAPI.ApiCephfsRemoveNameDelete(context.Background(), name).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CephfsAPIService ApiCephfsRenamePut", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.CephfsAPI.ApiCephfsRenamePut(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}