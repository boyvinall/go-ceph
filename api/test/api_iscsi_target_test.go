/*
Ceph REST API

Testing IscsiTargetAPIService

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

func Test_ceph_IscsiTargetAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test IscsiTargetAPIService ApiIscsiTargetGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.IscsiTargetAPI.ApiIscsiTargetGet(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IscsiTargetAPIService ApiIscsiTargetPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.IscsiTargetAPI.ApiIscsiTargetPost(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IscsiTargetAPIService ApiIscsiTargetTargetIqnDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var targetIqn string

		httpRes, err := apiClient.IscsiTargetAPI.ApiIscsiTargetTargetIqnDelete(context.Background(), targetIqn).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IscsiTargetAPIService ApiIscsiTargetTargetIqnGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var targetIqn string

		httpRes, err := apiClient.IscsiTargetAPI.ApiIscsiTargetTargetIqnGet(context.Background(), targetIqn).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IscsiTargetAPIService ApiIscsiTargetTargetIqnPut", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var targetIqn string

		httpRes, err := apiClient.IscsiTargetAPI.ApiIscsiTargetTargetIqnPut(context.Background(), targetIqn).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}