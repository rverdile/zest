/*
Pulp 3 API

Testing RepositoriesMavenVersionsApiService

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

func Test_zest_RepositoriesMavenVersionsApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test RepositoriesMavenVersionsApiService RepositoriesMavenMavenVersionsDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var mavenMavenRepositoryVersionHref string

		resp, httpRes, err := apiClient.RepositoriesMavenVersionsApi.RepositoriesMavenMavenVersionsDelete(context.Background(), mavenMavenRepositoryVersionHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RepositoriesMavenVersionsApiService RepositoriesMavenMavenVersionsList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var mavenMavenRepositoryHref string

		resp, httpRes, err := apiClient.RepositoriesMavenVersionsApi.RepositoriesMavenMavenVersionsList(context.Background(), mavenMavenRepositoryHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RepositoriesMavenVersionsApiService RepositoriesMavenMavenVersionsRead", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var mavenMavenRepositoryVersionHref string

		resp, httpRes, err := apiClient.RepositoriesMavenVersionsApi.RepositoriesMavenMavenVersionsRead(context.Background(), mavenMavenRepositoryVersionHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RepositoriesMavenVersionsApiService RepositoriesMavenMavenVersionsRepair", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var mavenMavenRepositoryVersionHref string

		resp, httpRes, err := apiClient.RepositoriesMavenVersionsApi.RepositoriesMavenMavenVersionsRepair(context.Background(), mavenMavenRepositoryVersionHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
