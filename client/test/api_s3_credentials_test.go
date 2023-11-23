/*
Taikun - WebApi

Testing S3CredentialsAPIService

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

func Test_taikuncore_S3CredentialsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test S3CredentialsAPIService S3credentialsCreate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.S3CredentialsAPI.S3credentialsCreate(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test S3CredentialsAPIService S3credentialsDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id int32

		httpRes, err := apiClient.S3CredentialsAPI.S3credentialsDelete(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test S3CredentialsAPIService S3credentialsDropdown", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.S3CredentialsAPI.S3credentialsDropdown(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test S3CredentialsAPIService S3credentialsList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.S3CredentialsAPI.S3credentialsList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test S3CredentialsAPIService S3credentialsLockManagement", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.S3CredentialsAPI.S3credentialsLockManagement(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test S3CredentialsAPIService S3credentialsMakeDeafult", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.S3CredentialsAPI.S3credentialsMakeDeafult(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test S3CredentialsAPIService S3credentialsUpdate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.S3CredentialsAPI.S3credentialsUpdate(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
