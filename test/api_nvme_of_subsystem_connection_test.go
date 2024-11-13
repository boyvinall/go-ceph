/*
Ceph REST API

Testing NVMeOFSubsystemConnectionAPIService

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

func Test_ceph_NVMeOFSubsystemConnectionAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test NVMeOFSubsystemConnectionAPIService ApiNvmeofSubsystemNqnConnectionGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var nqn string

		httpRes, err := apiClient.NVMeOFSubsystemConnectionAPI.ApiNvmeofSubsystemNqnConnectionGet(context.Background(), nqn).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
