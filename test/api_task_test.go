/*
Ceph REST API

Testing TaskAPIService

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

func Test_ceph_TaskAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test TaskAPIService ApiTaskGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.TaskAPI.ApiTaskGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
