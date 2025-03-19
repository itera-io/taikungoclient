/*
Taikun - WebApi

Testing KubernetesAPIService

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

func Test_taikuncore_KubernetesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test KubernetesAPIService KubernetesAlertList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var projectId int32

		resp, httpRes, err := apiClient.KubernetesAPI.KubernetesAlertList(context.Background(), projectId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KubernetesAPIService KubernetesConfigMapList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var projectId int32

		resp, httpRes, err := apiClient.KubernetesAPI.KubernetesConfigMapList(context.Background(), projectId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KubernetesAPIService KubernetesCrdList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var projectId int32

		resp, httpRes, err := apiClient.KubernetesAPI.KubernetesCrdList(context.Background(), projectId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KubernetesAPIService KubernetesCreateResource", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.KubernetesAPI.KubernetesCreateResource(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KubernetesAPIService KubernetesCronJobList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var projectId int32

		resp, httpRes, err := apiClient.KubernetesAPI.KubernetesCronJobList(context.Background(), projectId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KubernetesAPIService KubernetesCronjobActions", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.KubernetesAPI.KubernetesCronjobActions(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KubernetesAPIService KubernetesDaemonsetActions", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.KubernetesAPI.KubernetesDaemonsetActions(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KubernetesAPIService KubernetesDashboardList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var projectId int32

		resp, httpRes, err := apiClient.KubernetesAPI.KubernetesDashboardList(context.Background(), projectId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KubernetesAPIService KubernetesDeleteResource", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.KubernetesAPI.KubernetesDeleteResource(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KubernetesAPIService KubernetesDeploymentActions", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.KubernetesAPI.KubernetesDeploymentActions(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KubernetesAPIService KubernetesDeploymentList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var projectId int32

		resp, httpRes, err := apiClient.KubernetesAPI.KubernetesDeploymentList(context.Background(), projectId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KubernetesAPIService KubernetesDescribeResource", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.KubernetesAPI.KubernetesDescribeResource(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KubernetesAPIService KubernetesDownloadManifest", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.KubernetesAPI.KubernetesDownloadManifest(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KubernetesAPIService KubernetesExport", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.KubernetesAPI.KubernetesExport(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KubernetesAPIService KubernetesForbiddenNamespaces", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var projectId int32

		resp, httpRes, err := apiClient.KubernetesAPI.KubernetesForbiddenNamespaces(context.Background(), projectId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KubernetesAPIService KubernetesGetSupportedList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.KubernetesAPI.KubernetesGetSupportedList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KubernetesAPIService KubernetesHelmReleaseList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var projectId int32

		resp, httpRes, err := apiClient.KubernetesAPI.KubernetesHelmReleaseList(context.Background(), projectId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KubernetesAPIService KubernetesIngressClasses", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.KubernetesAPI.KubernetesIngressClasses(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KubernetesAPIService KubernetesIngressList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var projectId int32

		resp, httpRes, err := apiClient.KubernetesAPI.KubernetesIngressList(context.Background(), projectId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KubernetesAPIService KubernetesInteractiveShell", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.KubernetesAPI.KubernetesInteractiveShell(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KubernetesAPIService KubernetesJobsList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var projectId int32

		resp, httpRes, err := apiClient.KubernetesAPI.KubernetesJobsList(context.Background(), projectId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KubernetesAPIService KubernetesKillPod", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var projectId int32
		var metadataName string
		var namespace string

		resp, httpRes, err := apiClient.KubernetesAPI.KubernetesKillPod(context.Background(), projectId, metadataName, namespace).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KubernetesAPIService KubernetesNetworkPolicyList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var projectId int32

		resp, httpRes, err := apiClient.KubernetesAPI.KubernetesNetworkPolicyList(context.Background(), projectId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KubernetesAPIService KubernetesNodeActions", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.KubernetesAPI.KubernetesNodeActions(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KubernetesAPIService KubernetesNodeList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var projectId int32

		resp, httpRes, err := apiClient.KubernetesAPI.KubernetesNodeList(context.Background(), projectId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KubernetesAPIService KubernetesOverview", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.KubernetesAPI.KubernetesOverview(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KubernetesAPIService KubernetesPatchResource", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.KubernetesAPI.KubernetesPatchResource(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KubernetesAPIService KubernetesPdbList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var projectId int32

		resp, httpRes, err := apiClient.KubernetesAPI.KubernetesPdbList(context.Background(), projectId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KubernetesAPIService KubernetesPodList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var projectId int32

		resp, httpRes, err := apiClient.KubernetesAPI.KubernetesPodList(context.Background(), projectId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KubernetesAPIService KubernetesPodLogs", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.KubernetesAPI.KubernetesPodLogs(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KubernetesAPIService KubernetesPvcList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var projectId int32

		resp, httpRes, err := apiClient.KubernetesAPI.KubernetesPvcList(context.Background(), projectId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KubernetesAPIService KubernetesQuota", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var projectId int32

		resp, httpRes, err := apiClient.KubernetesAPI.KubernetesQuota(context.Background(), projectId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KubernetesAPIService KubernetesRemovealerts", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.KubernetesAPI.KubernetesRemovealerts(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KubernetesAPIService KubernetesResources", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var projectId int32

		resp, httpRes, err := apiClient.KubernetesAPI.KubernetesResources(context.Background(), projectId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KubernetesAPIService KubernetesRestartDaemonSet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.KubernetesAPI.KubernetesRestartDaemonSet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KubernetesAPIService KubernetesRestartDeployment", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.KubernetesAPI.KubernetesRestartDeployment(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KubernetesAPIService KubernetesRestartSts", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.KubernetesAPI.KubernetesRestartSts(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KubernetesAPIService KubernetesSecretList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var projectId int32

		resp, httpRes, err := apiClient.KubernetesAPI.KubernetesSecretList(context.Background(), projectId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KubernetesAPIService KubernetesServiceList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var projectId int32

		resp, httpRes, err := apiClient.KubernetesAPI.KubernetesServiceList(context.Background(), projectId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KubernetesAPIService KubernetesSilenceManager", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.KubernetesAPI.KubernetesSilenceManager(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KubernetesAPIService KubernetesStorageClassList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var projectId int32

		resp, httpRes, err := apiClient.KubernetesAPI.KubernetesStorageClassList(context.Background(), projectId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KubernetesAPIService KubernetesStsActions", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.KubernetesAPI.KubernetesStsActions(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KubernetesAPIService KubernetesStsList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var projectId int32

		resp, httpRes, err := apiClient.KubernetesAPI.KubernetesStsList(context.Background(), projectId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
