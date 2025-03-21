/*
Taikun - WebApi

Testing ProjectAppParamsAPIService

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

func Test_taikuncore_ProjectAppParamsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ProjectAppParamsAPIService ProjectappparamEdit", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var projectAppId int32

		resp, httpRes, err := apiClient.ProjectAppParamsAPI.ProjectappparamEdit(context.Background(), projectAppId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
