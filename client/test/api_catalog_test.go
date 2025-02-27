/*
Taikun - WebApi

Testing CatalogAPIService

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

func Test_taikuncore_CatalogAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test CatalogAPIService CatalogAddProject", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id int32

		httpRes, err := apiClient.CatalogAPI.CatalogAddProject(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CatalogAPIService CatalogCreate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.CatalogAPI.CatalogCreate(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CatalogAPIService CatalogDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id int32

		httpRes, err := apiClient.CatalogAPI.CatalogDelete(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CatalogAPIService CatalogDeleteProject", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id int32

		httpRes, err := apiClient.CatalogAPI.CatalogDeleteProject(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CatalogAPIService CatalogDetails", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id int32

		resp, httpRes, err := apiClient.CatalogAPI.CatalogDetails(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CatalogAPIService CatalogDropdown", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.CatalogAPI.CatalogDropdown(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CatalogAPIService CatalogEdit", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.CatalogAPI.CatalogEdit(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CatalogAPIService CatalogList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.CatalogAPI.CatalogList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CatalogAPIService CatalogLock", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.CatalogAPI.CatalogLock(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CatalogAPIService CatalogMakeDefault", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.CatalogAPI.CatalogMakeDefault(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
