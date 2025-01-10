/*
Ceph REST API

Testing RbdMirroringPoolPeerAPIService

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

func Test_ceph_RbdMirroringPoolPeerAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test RbdMirroringPoolPeerAPIService ApiBlockMirroringPoolPoolNamePeerGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var poolName string

		httpRes, err := apiClient.RbdMirroringPoolPeerAPI.ApiBlockMirroringPoolPoolNamePeerGet(context.Background(), poolName).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RbdMirroringPoolPeerAPIService ApiBlockMirroringPoolPoolNamePeerPeerUuidDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var poolName string
		var peerUuid string

		httpRes, err := apiClient.RbdMirroringPoolPeerAPI.ApiBlockMirroringPoolPoolNamePeerPeerUuidDelete(context.Background(), poolName, peerUuid).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RbdMirroringPoolPeerAPIService ApiBlockMirroringPoolPoolNamePeerPeerUuidGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var poolName string
		var peerUuid string

		httpRes, err := apiClient.RbdMirroringPoolPeerAPI.ApiBlockMirroringPoolPoolNamePeerPeerUuidGet(context.Background(), poolName, peerUuid).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RbdMirroringPoolPeerAPIService ApiBlockMirroringPoolPoolNamePeerPeerUuidPut", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var poolName string
		var peerUuid string

		httpRes, err := apiClient.RbdMirroringPoolPeerAPI.ApiBlockMirroringPoolPoolNamePeerPeerUuidPut(context.Background(), poolName, peerUuid).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RbdMirroringPoolPeerAPIService ApiBlockMirroringPoolPoolNamePeerPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var poolName string

		httpRes, err := apiClient.RbdMirroringPoolPeerAPI.ApiBlockMirroringPoolPoolNamePeerPost(context.Background(), poolName).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
