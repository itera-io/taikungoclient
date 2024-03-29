/*
Taikun - WebApi

Testing AccessProfilesAPIService

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

func Test_taikuncore_AccessProfilesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test AccessProfilesAPIService AccessprofilesCreate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.AccessProfilesAPI.AccessprofilesCreate(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AccessProfilesAPIService AccessprofilesDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id int32

		httpRes, err := apiClient.AccessProfilesAPI.AccessprofilesDelete(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AccessProfilesAPIService AccessprofilesDropdown", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.AccessProfilesAPI.AccessprofilesDropdown(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AccessProfilesAPIService AccessprofilesList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.AccessProfilesAPI.AccessprofilesList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AccessProfilesAPIService AccessprofilesLockManager", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.AccessProfilesAPI.AccessprofilesLockManager(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AccessProfilesAPIService AccessprofilesUpdate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id int32

		httpRes, err := apiClient.AccessProfilesAPI.AccessprofilesUpdate(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
