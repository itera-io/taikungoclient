/*
Taikun - WebApi

Testing StripeAPIService

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

func Test_taikunshowback_StripeAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test StripeAPIService StripeSubscriptionItemList", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var subscriptionId string

		resp, httpRes, err := apiClient.StripeAPI.StripeSubscriptionItemList(context.Background(), subscriptionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
