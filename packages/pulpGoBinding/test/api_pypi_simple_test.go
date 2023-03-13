/*
Pulp 3 API

Testing PypiSimpleApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package zest

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/content-services/zest/packages/pulpGoBinding"
)

func Test_zest_PypiSimpleApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test PypiSimpleApiService PypiSimpleCreate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var path string

		resp, httpRes, err := apiClient.PypiSimpleApi.PypiSimpleCreate(context.Background(), path).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PypiSimpleApiService PypiSimpleRead", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var path string

		httpRes, err := apiClient.PypiSimpleApi.PypiSimpleRead(context.Background(), path).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PypiSimpleApiService PypiSimpleReadPackage", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var package_ string
		var path string

		httpRes, err := apiClient.PypiSimpleApi.PypiSimpleReadPackage(context.Background(), package_, path).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
