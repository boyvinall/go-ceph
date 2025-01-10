/*
Ceph REST API

Testing RbdTrashAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package ceph

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/boyvinall/go-ceph/api"
)

func Test_ceph_RbdTrashAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test RbdTrashAPIService ApiBlockImageTrashGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.RbdTrashAPI.ApiBlockImageTrashGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RbdTrashAPIService ApiBlockImageTrashImageIdSpecDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var imageIdSpec string

		httpRes, err := apiClient.RbdTrashAPI.ApiBlockImageTrashImageIdSpecDelete(context.Background(), imageIdSpec).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RbdTrashAPIService ApiBlockImageTrashImageIdSpecRestorePost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var imageIdSpec string

		httpRes, err := apiClient.RbdTrashAPI.ApiBlockImageTrashImageIdSpecRestorePost(context.Background(), imageIdSpec).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RbdTrashAPIService ApiBlockImageTrashPurgePost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.RbdTrashAPI.ApiBlockImageTrashPurgePost(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
