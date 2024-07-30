/*
Taikun - WebApi

Testing StandaloneProfileAPIService

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

func Test_taikuncore_StandaloneProfileAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test StandaloneProfileAPIService StandaloneprofileCreate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.StandaloneProfileAPI.StandaloneprofileCreate(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test StandaloneProfileAPIService StandaloneprofileDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.StandaloneProfileAPI.StandaloneprofileDelete(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test StandaloneProfileAPIService StandaloneprofileDropdown", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.StandaloneProfileAPI.StandaloneprofileDropdown(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test StandaloneProfileAPIService StandaloneprofileEdit", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.StandaloneProfileAPI.StandaloneprofileEdit(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test StandaloneProfileAPIService StandaloneprofileList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.StandaloneProfileAPI.StandaloneprofileList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test StandaloneProfileAPIService StandaloneprofileLockManagement", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.StandaloneProfileAPI.StandaloneprofileLockManagement(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
