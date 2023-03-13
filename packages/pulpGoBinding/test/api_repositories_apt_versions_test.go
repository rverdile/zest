/*
Pulp 3 API

Testing RepositoriesAptVersionsApiService

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

func Test_pulpGoBinding_RepositoriesAptVersionsApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test RepositoriesAptVersionsApiService RepositoriesDebAptVersionsDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var debAptRepositoryVersionHref string

		resp, httpRes, err := apiClient.RepositoriesAptVersionsApi.RepositoriesDebAptVersionsDelete(context.Background(), debAptRepositoryVersionHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RepositoriesAptVersionsApiService RepositoriesDebAptVersionsList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var debAptRepositoryHref string

		resp, httpRes, err := apiClient.RepositoriesAptVersionsApi.RepositoriesDebAptVersionsList(context.Background(), debAptRepositoryHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RepositoriesAptVersionsApiService RepositoriesDebAptVersionsRead", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var debAptRepositoryVersionHref string

		resp, httpRes, err := apiClient.RepositoriesAptVersionsApi.RepositoriesDebAptVersionsRead(context.Background(), debAptRepositoryVersionHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RepositoriesAptVersionsApiService RepositoriesDebAptVersionsRepair", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var debAptRepositoryVersionHref string

		resp, httpRes, err := apiClient.RepositoriesAptVersionsApi.RepositoriesDebAptVersionsRepair(context.Background(), debAptRepositoryVersionHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
