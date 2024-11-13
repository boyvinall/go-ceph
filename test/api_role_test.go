/*
Ceph REST API

Testing RoleAPIService

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

func Test_ceph_RoleAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test RoleAPIService ApiRoleGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.RoleAPI.ApiRoleGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RoleAPIService ApiRoleNameClonePost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var name string

		httpRes, err := apiClient.RoleAPI.ApiRoleNameClonePost(context.Background(), name).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RoleAPIService ApiRoleNameDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var name string

		httpRes, err := apiClient.RoleAPI.ApiRoleNameDelete(context.Background(), name).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RoleAPIService ApiRoleNameGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var name string

		httpRes, err := apiClient.RoleAPI.ApiRoleNameGet(context.Background(), name).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RoleAPIService ApiRoleNamePut", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var name string

		httpRes, err := apiClient.RoleAPI.ApiRoleNamePut(context.Background(), name).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RoleAPIService ApiRolePost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.RoleAPI.ApiRolePost(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}