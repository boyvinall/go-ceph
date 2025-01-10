/*
Ceph REST API

Testing RgwRealmAPIService

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

func Test_ceph_RgwRealmAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test RgwRealmAPIService ApiRgwRealmGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.RgwRealmAPI.ApiRgwRealmGet(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RgwRealmAPIService ApiRgwRealmGetAllRealmsInfoGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.RgwRealmAPI.ApiRgwRealmGetAllRealmsInfoGet(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RgwRealmAPIService ApiRgwRealmGetRealmTokensGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.RgwRealmAPI.ApiRgwRealmGetRealmTokensGet(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RgwRealmAPIService ApiRgwRealmImportRealmTokenPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.RgwRealmAPI.ApiRgwRealmImportRealmTokenPost(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RgwRealmAPIService ApiRgwRealmPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.RgwRealmAPI.ApiRgwRealmPost(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RgwRealmAPIService ApiRgwRealmRealmNameDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var realmName string

		httpRes, err := apiClient.RgwRealmAPI.ApiRgwRealmRealmNameDelete(context.Background(), realmName).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RgwRealmAPIService ApiRgwRealmRealmNameGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var realmName string

		httpRes, err := apiClient.RgwRealmAPI.ApiRgwRealmRealmNameGet(context.Background(), realmName).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RgwRealmAPIService ApiRgwRealmRealmNamePut", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var realmName string

		httpRes, err := apiClient.RgwRealmAPI.ApiRgwRealmRealmNamePut(context.Background(), realmName).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
