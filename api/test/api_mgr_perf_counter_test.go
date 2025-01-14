/*
Ceph REST API

Testing MgrPerfCounterAPIService

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

func Test_ceph_MgrPerfCounterAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test MgrPerfCounterAPIService ApiPerfCountersMgrServiceIdGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var serviceId string

		httpRes, err := apiClient.MgrPerfCounterAPI.ApiPerfCountersMgrServiceIdGet(context.Background(), serviceId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
