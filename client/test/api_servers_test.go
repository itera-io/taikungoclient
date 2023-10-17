/*
Taikun - WebApi

Testing ServersAPIService

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

func Test_taikuncore_ServersAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ServersAPIService ServersConsole", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ServersAPI.ServersConsole(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ServersAPIService ServersCreate", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ServersAPI.ServersCreate(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ServersAPIService ServersDelete", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.ServersAPI.ServersDelete(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ServersAPIService ServersDetails", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var projectId int32

		resp, httpRes, err := apiClient.ServersAPI.ServersDetails(context.Background(), projectId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ServersAPIService ServersList", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ServersAPI.ServersList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ServersAPIService ServersReboot", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.ServersAPI.ServersReboot(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ServersAPIService ServersReset", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.ServersAPI.ServersReset(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ServersAPIService ServersStatus", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var serverId int32

		resp, httpRes, err := apiClient.ServersAPI.ServersStatus(context.Background(), serverId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ServersAPIService ServersUpdate", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.ServersAPI.ServersUpdate(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ServersAPIService ServersUpdateByProjectId", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var projectId int32

		httpRes, err := apiClient.ServersAPI.ServersUpdateByProjectId(context.Background(), projectId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
