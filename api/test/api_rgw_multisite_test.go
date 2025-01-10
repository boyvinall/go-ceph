/*
Ceph REST API

Testing RgwMultisiteAPIService

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

func Test_ceph_RgwMultisiteAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test RgwMultisiteAPIService ApiRgwMultisiteSyncFlowFlowIdFlowTypeGroupIdDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var flowId string
		var flowType string
		var groupId string

		httpRes, err := apiClient.RgwMultisiteAPI.ApiRgwMultisiteSyncFlowFlowIdFlowTypeGroupIdDelete(context.Background(), flowId, flowType, groupId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RgwMultisiteAPIService ApiRgwMultisiteSyncFlowPut", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.RgwMultisiteAPI.ApiRgwMultisiteSyncFlowPut(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RgwMultisiteAPIService ApiRgwMultisiteSyncPipeGroupIdPipeIdDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var groupId string
		var pipeId string

		httpRes, err := apiClient.RgwMultisiteAPI.ApiRgwMultisiteSyncPipeGroupIdPipeIdDelete(context.Background(), groupId, pipeId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RgwMultisiteAPIService ApiRgwMultisiteSyncPipePut", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.RgwMultisiteAPI.ApiRgwMultisiteSyncPipePut(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RgwMultisiteAPIService ApiRgwMultisiteSyncPolicyGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.RgwMultisiteAPI.ApiRgwMultisiteSyncPolicyGet(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RgwMultisiteAPIService ApiRgwMultisiteSyncPolicyGroupGroupIdDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var groupId string

		httpRes, err := apiClient.RgwMultisiteAPI.ApiRgwMultisiteSyncPolicyGroupGroupIdDelete(context.Background(), groupId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RgwMultisiteAPIService ApiRgwMultisiteSyncPolicyGroupGroupIdGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var groupId string

		httpRes, err := apiClient.RgwMultisiteAPI.ApiRgwMultisiteSyncPolicyGroupGroupIdGet(context.Background(), groupId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RgwMultisiteAPIService ApiRgwMultisiteSyncPolicyGroupPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.RgwMultisiteAPI.ApiRgwMultisiteSyncPolicyGroupPost(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RgwMultisiteAPIService ApiRgwMultisiteSyncPolicyGroupPut", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.RgwMultisiteAPI.ApiRgwMultisiteSyncPolicyGroupPut(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RgwMultisiteAPIService ApiRgwMultisiteSyncStatusGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.RgwMultisiteAPI.ApiRgwMultisiteSyncStatusGet(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
