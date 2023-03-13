/*
Pulp 3 API

Testing RepositoriesFileVersionsApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package pulpGoBinding

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/content-services/pulpGoBinding/packages/pulpGoBinding"
)

func Test_pulpGoBinding_RepositoriesFileVersionsApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test RepositoriesFileVersionsApiService RepositoriesFileFileVersionsDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var fileFileRepositoryVersionHref string

		resp, httpRes, err := apiClient.RepositoriesFileVersionsApi.RepositoriesFileFileVersionsDelete(context.Background(), fileFileRepositoryVersionHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RepositoriesFileVersionsApiService RepositoriesFileFileVersionsList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var fileFileRepositoryHref string

		resp, httpRes, err := apiClient.RepositoriesFileVersionsApi.RepositoriesFileFileVersionsList(context.Background(), fileFileRepositoryHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RepositoriesFileVersionsApiService RepositoriesFileFileVersionsRead", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var fileFileRepositoryVersionHref string

		resp, httpRes, err := apiClient.RepositoriesFileVersionsApi.RepositoriesFileFileVersionsRead(context.Background(), fileFileRepositoryVersionHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RepositoriesFileVersionsApiService RepositoriesFileFileVersionsRepair", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var fileFileRepositoryVersionHref string

		resp, httpRes, err := apiClient.RepositoriesFileVersionsApi.RepositoriesFileFileVersionsRepair(context.Background(), fileFileRepositoryVersionHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
