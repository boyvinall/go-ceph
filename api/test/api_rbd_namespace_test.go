/*
Ceph REST API

Testing RbdNamespaceAPIService

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

func Test_ceph_RbdNamespaceAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test RbdNamespaceAPIService ApiBlockPoolPoolNameNamespaceGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var poolName string

		httpRes, err := apiClient.RbdNamespaceAPI.ApiBlockPoolPoolNameNamespaceGet(context.Background(), poolName).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RbdNamespaceAPIService ApiBlockPoolPoolNameNamespaceNamespaceDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var poolName string
		var namespace string

		httpRes, err := apiClient.RbdNamespaceAPI.ApiBlockPoolPoolNameNamespaceNamespaceDelete(context.Background(), poolName, namespace).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RbdNamespaceAPIService ApiBlockPoolPoolNameNamespacePost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var poolName string

		httpRes, err := apiClient.RbdNamespaceAPI.ApiBlockPoolPoolNameNamespacePost(context.Background(), poolName).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
