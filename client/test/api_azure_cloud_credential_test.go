/*
Taikun - WebApi

Testing AzureCloudCredentialAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package taikuncore

import (
	"context"
	openapiclient "github.com/itera-io/taikungoclient/client"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_taikuncore_AzureCloudCredentialAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test AzureCloudCredentialAPIService AzureCreate", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.AzureCloudCredentialAPI.AzureCreate(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AzureCloudCredentialAPIService AzureDashboard", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.AzureCloudCredentialAPI.AzureDashboard(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AzureCloudCredentialAPIService AzureList", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.AzureCloudCredentialAPI.AzureList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AzureCloudCredentialAPIService AzureLocations", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.AzureCloudCredentialAPI.AzureLocations(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AzureCloudCredentialAPIService AzureOffers", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var cloudId int32
		var publisher string

		resp, httpRes, err := apiClient.AzureCloudCredentialAPI.AzureOffers(context.Background(), cloudId, publisher).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AzureCloudCredentialAPIService AzurePublishers", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var cloudId int32

		resp, httpRes, err := apiClient.AzureCloudCredentialAPI.AzurePublishers(context.Background(), cloudId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AzureCloudCredentialAPIService AzureSkus", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var cloudId int32
		var publisher string
		var offer string

		resp, httpRes, err := apiClient.AzureCloudCredentialAPI.AzureSkus(context.Background(), cloudId, publisher, offer).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AzureCloudCredentialAPIService AzureSubscriptions", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.AzureCloudCredentialAPI.AzureSubscriptions(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AzureCloudCredentialAPIService AzureUpdate", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.AzureCloudCredentialAPI.AzureUpdate(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AzureCloudCredentialAPIService AzureZones", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.AzureCloudCredentialAPI.AzureZones(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
