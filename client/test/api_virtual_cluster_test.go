/*
Taikun - WebApi

Testing VirtualClusterAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package taikuncore

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/itera-io/taikungoclient/client"
)

func Test_taikuncore_VirtualClusterAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test VirtualClusterAPIService VirtualClusterCreate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.VirtualClusterAPI.VirtualClusterCreate(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VirtualClusterAPIService VirtualClusterDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.VirtualClusterAPI.VirtualClusterDelete(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VirtualClusterAPIService VirtualClusterUpdate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.VirtualClusterAPI.VirtualClusterUpdate(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}