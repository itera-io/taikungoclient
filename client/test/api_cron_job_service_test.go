/*
Taikun - WebApi

Testing CronJobServiceAPIService

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

func Test_taikuncore_CronJobServiceAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test CronJobServiceAPIService CronjobAutoUpgradeProjects", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.CronJobServiceAPI.CronjobAutoUpgradeProjects(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CronJobServiceAPIService CronjobBlockOrganization", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.CronJobServiceAPI.CronjobBlockOrganization(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CronJobServiceAPIService CronjobCancelExpiredSubscriptions", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.CronJobServiceAPI.CronjobCancelExpiredSubscriptions(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CronJobServiceAPIService CronjobCreateKeyPool", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.CronJobServiceAPI.CronjobCreateKeyPool(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CronJobServiceAPIService CronjobDeleteExpiredAlerts", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.CronJobServiceAPI.CronjobDeleteExpiredAlerts(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CronJobServiceAPIService CronjobDeleteExpiredHistoryLogs", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.CronJobServiceAPI.CronjobDeleteExpiredHistoryLogs(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CronJobServiceAPIService CronjobDeleteExpiredOrgs", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.CronJobServiceAPI.CronjobDeleteExpiredOrgs(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CronJobServiceAPIService CronjobDeleteExpiredRefreshTokens", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.CronJobServiceAPI.CronjobDeleteExpiredRefreshTokens(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CronJobServiceAPIService CronjobDeleteExpiredServers", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.CronJobServiceAPI.CronjobDeleteExpiredServers(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CronJobServiceAPIService CronjobDeleteImportedBackupLocation", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.CronJobServiceAPI.CronjobDeleteImportedBackupLocation(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CronJobServiceAPIService CronjobDeleteKubeConfigs", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.CronJobServiceAPI.CronjobDeleteKubeConfigs(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CronJobServiceAPIService CronjobDeleteRemovedSpotInstances", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.CronJobServiceAPI.CronjobDeleteRemovedSpotInstances(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CronJobServiceAPIService CronjobDeleteUselessProjectActions", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.CronJobServiceAPI.CronjobDeleteUselessProjectActions(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CronJobServiceAPIService CronjobEmailForProjectExpiration", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.CronJobServiceAPI.CronjobEmailForProjectExpiration(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CronJobServiceAPIService CronjobFetchArtifactOrganizations", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.CronJobServiceAPI.CronjobFetchArtifactOrganizations(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CronJobServiceAPIService CronjobFetchAzureFlavorPrices", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.CronJobServiceAPI.CronjobFetchAzureFlavorPrices(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CronJobServiceAPIService CronjobFetchK8sAlertData", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.CronJobServiceAPI.CronjobFetchK8sAlertData(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CronJobServiceAPIService CronjobFetchK8sOverviewData", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.CronJobServiceAPI.CronjobFetchK8sOverviewData(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CronJobServiceAPIService CronjobFetchOrganizationDetails", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.CronJobServiceAPI.CronjobFetchOrganizationDetails(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CronJobServiceAPIService CronjobPurgeExpiredProjects", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.CronJobServiceAPI.CronjobPurgeExpiredProjects(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CronJobServiceAPIService CronjobRemindUsersByAlertingProfile", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.CronJobServiceAPI.CronjobRemindUsersByAlertingProfile(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CronJobServiceAPIService CronjobSyncBackupCredentials", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.CronJobServiceAPI.CronjobSyncBackupCredentials(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CronJobServiceAPIService CronjobSyncOpaProfiles", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.CronJobServiceAPI.CronjobSyncOpaProfiles(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CronJobServiceAPIService CronjobSyncOrganizations", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.CronJobServiceAPI.CronjobSyncOrganizations(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CronJobServiceAPIService CronjobSyncProjectApps", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.CronJobServiceAPI.CronjobSyncProjectApps(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CronJobServiceAPIService CronjobSyncProjects", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.CronJobServiceAPI.CronjobSyncProjects(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CronJobServiceAPIService CronjobTektonPipelines", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.CronJobServiceAPI.CronjobTektonPipelines(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CronJobServiceAPIService CronjobUpdateProjectAppStatus", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.CronJobServiceAPI.CronjobUpdateProjectAppStatus(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CronJobServiceAPIService CronjobUpdateProjectQuotaMessage", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.CronJobServiceAPI.CronjobUpdateProjectQuotaMessage(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
