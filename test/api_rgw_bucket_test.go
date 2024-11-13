/*
Ceph REST API

Testing RgwBucketAPIService

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

func Test_ceph_RgwBucketAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test RgwBucketAPIService ApiRgwBucketBucketDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var bucket string

		httpRes, err := apiClient.RgwBucketAPI.ApiRgwBucketBucketDelete(context.Background(), bucket).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RgwBucketAPIService ApiRgwBucketBucketGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var bucket string

		httpRes, err := apiClient.RgwBucketAPI.ApiRgwBucketBucketGet(context.Background(), bucket).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RgwBucketAPIService ApiRgwBucketBucketPut", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var bucket string

		httpRes, err := apiClient.RgwBucketAPI.ApiRgwBucketBucketPut(context.Background(), bucket).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RgwBucketAPIService ApiRgwBucketDeleteEncryptionDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.RgwBucketAPI.ApiRgwBucketDeleteEncryptionDelete(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RgwBucketAPIService ApiRgwBucketGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.RgwBucketAPI.ApiRgwBucketGet(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RgwBucketAPIService ApiRgwBucketGetEncryptionConfigGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.RgwBucketAPI.ApiRgwBucketGetEncryptionConfigGet(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RgwBucketAPIService ApiRgwBucketGetEncryptionGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.RgwBucketAPI.ApiRgwBucketGetEncryptionGet(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RgwBucketAPIService ApiRgwBucketPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.RgwBucketAPI.ApiRgwBucketPost(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RgwBucketAPIService ApiRgwBucketSetEncryptionConfigPut", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.RgwBucketAPI.ApiRgwBucketSetEncryptionConfigPut(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}