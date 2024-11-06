/*
Taikun - WebApi

Testing ImportedClusterAPIService

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

func Test_taikuncore_ImportedClusterAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ImportedClusterAPIService ImportedClusterAsCloudCredential", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var projectId int32

		resp, httpRes, err := apiClient.ImportedClusterAPI.ImportedClusterAsCloudCredential(context.Background(), projectId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ImportedClusterAPIService ImportedClusterAsFullyManaged", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var projectId int32

		resp, httpRes, err := apiClient.ImportedClusterAPI.ImportedClusterAsFullyManaged(context.Background(), projectId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ImportedClusterAPIService ImportedClusterAsReadOnly", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var projectId int32

		resp, httpRes, err := apiClient.ImportedClusterAPI.ImportedClusterAsReadOnly(context.Background(), projectId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ImportedClusterAPIService ImportedClusterDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.ImportedClusterAPI.ImportedClusterDelete(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ImportedClusterAPIService ImportedClusterDetails", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id int32

		resp, httpRes, err := apiClient.ImportedClusterAPI.ImportedClusterDetails(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
