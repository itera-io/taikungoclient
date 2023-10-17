/*
Taikun - WebApi

Testing DnsServersAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package taikunshowback

import (
	"context"
	openapiclient "github.com/itera-io/taikungoclient/showbackclient"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_taikunshowback_DnsServersAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test DnsServersAPIService DnsserversCreate", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.DnsServersAPI.DnsserversCreate(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DnsServersAPIService DnsserversDelete", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id int32

		httpRes, err := apiClient.DnsServersAPI.DnsserversDelete(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DnsServersAPIService DnsserversEdit", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id int32

		httpRes, err := apiClient.DnsServersAPI.DnsserversEdit(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DnsServersAPIService DnsserversList", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var accessProfileId int32

		resp, httpRes, err := apiClient.DnsServersAPI.DnsserversList(context.Background(), accessProfileId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
