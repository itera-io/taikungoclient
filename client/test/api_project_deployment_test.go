/*
Taikun - WebApi

Testing ProjectDeploymentAPIService

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

func Test_taikuncore_ProjectDeploymentAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ProjectDeploymentAPIService ProjectDeploymentCommit", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.ProjectDeploymentAPI.ProjectDeploymentCommit(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectDeploymentAPIService ProjectDeploymentCommitVm", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.ProjectDeploymentAPI.ProjectDeploymentCommitVm(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectDeploymentAPIService ProjectDeploymentCompleted", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.ProjectDeploymentAPI.ProjectDeploymentCompleted(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectDeploymentAPIService ProjectDeploymentDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.ProjectDeploymentAPI.ProjectDeploymentDelete(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectDeploymentAPIService ProjectDeploymentDeleteVmDisks", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.ProjectDeploymentAPI.ProjectDeploymentDeleteVmDisks(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectDeploymentAPIService ProjectDeploymentDeleteVms", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.ProjectDeploymentAPI.ProjectDeploymentDeleteVms(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectDeploymentAPIService ProjectDeploymentDisableAi", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.ProjectDeploymentAPI.ProjectDeploymentDisableAi(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectDeploymentAPIService ProjectDeploymentDisableBackup", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.ProjectDeploymentAPI.ProjectDeploymentDisableBackup(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectDeploymentAPIService ProjectDeploymentDisableMonitoring", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.ProjectDeploymentAPI.ProjectDeploymentDisableMonitoring(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectDeploymentAPIService ProjectDeploymentDisableOpa", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.ProjectDeploymentAPI.ProjectDeploymentDisableOpa(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectDeploymentAPIService ProjectDeploymentEnableAi", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.ProjectDeploymentAPI.ProjectDeploymentEnableAi(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectDeploymentAPIService ProjectDeploymentEnableBackup", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.ProjectDeploymentAPI.ProjectDeploymentEnableBackup(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectDeploymentAPIService ProjectDeploymentEnableMonitoring", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.ProjectDeploymentAPI.ProjectDeploymentEnableMonitoring(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectDeploymentAPIService ProjectDeploymentEnableOpa", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.ProjectDeploymentAPI.ProjectDeploymentEnableOpa(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectDeploymentAPIService ProjectDeploymentImportCluster", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.ProjectDeploymentAPI.ProjectDeploymentImportCluster(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectDeploymentAPIService ProjectDeploymentRepair", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.ProjectDeploymentAPI.ProjectDeploymentRepair(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectDeploymentAPIService ProjectDeploymentRepairVm", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.ProjectDeploymentAPI.ProjectDeploymentRepairVm(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectDeploymentAPIService ProjectDeploymentTofuMigrate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.ProjectDeploymentAPI.ProjectDeploymentTofuMigrate(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectDeploymentAPIService ProjectDeploymentUpdate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.ProjectDeploymentAPI.ProjectDeploymentUpdate(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectDeploymentAPIService ProjectDeploymentUpgrade", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var projectId int32

		httpRes, err := apiClient.ProjectDeploymentAPI.ProjectDeploymentUpgrade(context.Background(), projectId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
