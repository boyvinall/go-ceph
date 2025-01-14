/*
Ceph REST API

Testing CephfsSubvolumeSnapshotAPIService

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

func Test_ceph_CephfsSubvolumeSnapshotAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test CephfsSubvolumeSnapshotAPIService ApiCephfsSubvolumeSnapshotPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.CephfsSubvolumeSnapshotAPI.ApiCephfsSubvolumeSnapshotPost(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CephfsSubvolumeSnapshotAPIService ApiCephfsSubvolumeSnapshotVolNameSubvolNameDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var volName string
		var subvolName string

		httpRes, err := apiClient.CephfsSubvolumeSnapshotAPI.ApiCephfsSubvolumeSnapshotVolNameSubvolNameDelete(context.Background(), volName, subvolName).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CephfsSubvolumeSnapshotAPIService ApiCephfsSubvolumeSnapshotVolNameSubvolNameGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var volName string
		var subvolName string

		httpRes, err := apiClient.CephfsSubvolumeSnapshotAPI.ApiCephfsSubvolumeSnapshotVolNameSubvolNameGet(context.Background(), volName, subvolName).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CephfsSubvolumeSnapshotAPIService ApiCephfsSubvolumeSnapshotVolNameSubvolNameInfoGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var volName string
		var subvolName string

		httpRes, err := apiClient.CephfsSubvolumeSnapshotAPI.ApiCephfsSubvolumeSnapshotVolNameSubvolNameInfoGet(context.Background(), volName, subvolName).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
