/*
Pulp 3 API

Testing RepositoriesContainerApiService

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

func Test_zest_RepositoriesContainerApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test RepositoriesContainerApiService RepositoriesContainerContainerAdd", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var containerContainerRepositoryHref string

		resp, httpRes, err := apiClient.RepositoriesContainerApi.RepositoriesContainerContainerAdd(context.Background(), containerContainerRepositoryHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RepositoriesContainerApiService RepositoriesContainerContainerAddRole", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var containerContainerRepositoryHref string

		resp, httpRes, err := apiClient.RepositoriesContainerApi.RepositoriesContainerContainerAddRole(context.Background(), containerContainerRepositoryHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RepositoriesContainerApiService RepositoriesContainerContainerBuildImage", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var containerContainerRepositoryHref string

		resp, httpRes, err := apiClient.RepositoriesContainerApi.RepositoriesContainerContainerBuildImage(context.Background(), containerContainerRepositoryHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RepositoriesContainerApiService RepositoriesContainerContainerCopyManifests", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var containerContainerRepositoryHref string

		resp, httpRes, err := apiClient.RepositoriesContainerApi.RepositoriesContainerContainerCopyManifests(context.Background(), containerContainerRepositoryHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RepositoriesContainerApiService RepositoriesContainerContainerCopyTags", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var containerContainerRepositoryHref string

		resp, httpRes, err := apiClient.RepositoriesContainerApi.RepositoriesContainerContainerCopyTags(context.Background(), containerContainerRepositoryHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RepositoriesContainerApiService RepositoriesContainerContainerCreate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.RepositoriesContainerApi.RepositoriesContainerContainerCreate(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RepositoriesContainerApiService RepositoriesContainerContainerDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var containerContainerRepositoryHref string

		resp, httpRes, err := apiClient.RepositoriesContainerApi.RepositoriesContainerContainerDelete(context.Background(), containerContainerRepositoryHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RepositoriesContainerApiService RepositoriesContainerContainerList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.RepositoriesContainerApi.RepositoriesContainerContainerList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RepositoriesContainerApiService RepositoriesContainerContainerListRoles", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var containerContainerRepositoryHref string

		resp, httpRes, err := apiClient.RepositoriesContainerApi.RepositoriesContainerContainerListRoles(context.Background(), containerContainerRepositoryHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RepositoriesContainerApiService RepositoriesContainerContainerMyPermissions", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var containerContainerRepositoryHref string

		resp, httpRes, err := apiClient.RepositoriesContainerApi.RepositoriesContainerContainerMyPermissions(context.Background(), containerContainerRepositoryHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RepositoriesContainerApiService RepositoriesContainerContainerPartialUpdate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var containerContainerRepositoryHref string

		resp, httpRes, err := apiClient.RepositoriesContainerApi.RepositoriesContainerContainerPartialUpdate(context.Background(), containerContainerRepositoryHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RepositoriesContainerApiService RepositoriesContainerContainerRead", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var containerContainerRepositoryHref string

		resp, httpRes, err := apiClient.RepositoriesContainerApi.RepositoriesContainerContainerRead(context.Background(), containerContainerRepositoryHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RepositoriesContainerApiService RepositoriesContainerContainerRemove", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var containerContainerRepositoryHref string

		resp, httpRes, err := apiClient.RepositoriesContainerApi.RepositoriesContainerContainerRemove(context.Background(), containerContainerRepositoryHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RepositoriesContainerApiService RepositoriesContainerContainerRemoveRole", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var containerContainerRepositoryHref string

		resp, httpRes, err := apiClient.RepositoriesContainerApi.RepositoriesContainerContainerRemoveRole(context.Background(), containerContainerRepositoryHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RepositoriesContainerApiService RepositoriesContainerContainerSign", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var containerContainerRepositoryHref string

		resp, httpRes, err := apiClient.RepositoriesContainerApi.RepositoriesContainerContainerSign(context.Background(), containerContainerRepositoryHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RepositoriesContainerApiService RepositoriesContainerContainerSync", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var containerContainerRepositoryHref string

		resp, httpRes, err := apiClient.RepositoriesContainerApi.RepositoriesContainerContainerSync(context.Background(), containerContainerRepositoryHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RepositoriesContainerApiService RepositoriesContainerContainerTag", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var containerContainerRepositoryHref string

		resp, httpRes, err := apiClient.RepositoriesContainerApi.RepositoriesContainerContainerTag(context.Background(), containerContainerRepositoryHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RepositoriesContainerApiService RepositoriesContainerContainerUntag", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var containerContainerRepositoryHref string

		resp, httpRes, err := apiClient.RepositoriesContainerApi.RepositoriesContainerContainerUntag(context.Background(), containerContainerRepositoryHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RepositoriesContainerApiService RepositoriesContainerContainerUpdate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var containerContainerRepositoryHref string

		resp, httpRes, err := apiClient.RepositoriesContainerApi.RepositoriesContainerContainerUpdate(context.Background(), containerContainerRepositoryHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
