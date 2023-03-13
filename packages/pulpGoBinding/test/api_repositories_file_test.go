/*
Pulp 3 API

Testing RepositoriesFileApiService

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

func Test_zest_RepositoriesFileApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test RepositoriesFileApiService RepositoriesFileFileAddRole", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var fileFileRepositoryHref string

		resp, httpRes, err := apiClient.RepositoriesFileApi.RepositoriesFileFileAddRole(context.Background(), fileFileRepositoryHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RepositoriesFileApiService RepositoriesFileFileCreate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.RepositoriesFileApi.RepositoriesFileFileCreate(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RepositoriesFileApiService RepositoriesFileFileDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var fileFileRepositoryHref string

		resp, httpRes, err := apiClient.RepositoriesFileApi.RepositoriesFileFileDelete(context.Background(), fileFileRepositoryHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RepositoriesFileApiService RepositoriesFileFileList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.RepositoriesFileApi.RepositoriesFileFileList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RepositoriesFileApiService RepositoriesFileFileListRoles", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var fileFileRepositoryHref string

		resp, httpRes, err := apiClient.RepositoriesFileApi.RepositoriesFileFileListRoles(context.Background(), fileFileRepositoryHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RepositoriesFileApiService RepositoriesFileFileModify", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var fileFileRepositoryHref string

		resp, httpRes, err := apiClient.RepositoriesFileApi.RepositoriesFileFileModify(context.Background(), fileFileRepositoryHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RepositoriesFileApiService RepositoriesFileFileMyPermissions", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var fileFileRepositoryHref string

		resp, httpRes, err := apiClient.RepositoriesFileApi.RepositoriesFileFileMyPermissions(context.Background(), fileFileRepositoryHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RepositoriesFileApiService RepositoriesFileFilePartialUpdate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var fileFileRepositoryHref string

		resp, httpRes, err := apiClient.RepositoriesFileApi.RepositoriesFileFilePartialUpdate(context.Background(), fileFileRepositoryHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RepositoriesFileApiService RepositoriesFileFileRead", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var fileFileRepositoryHref string

		resp, httpRes, err := apiClient.RepositoriesFileApi.RepositoriesFileFileRead(context.Background(), fileFileRepositoryHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RepositoriesFileApiService RepositoriesFileFileRemoveRole", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var fileFileRepositoryHref string

		resp, httpRes, err := apiClient.RepositoriesFileApi.RepositoriesFileFileRemoveRole(context.Background(), fileFileRepositoryHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RepositoriesFileApiService RepositoriesFileFileSync", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var fileFileRepositoryHref string

		resp, httpRes, err := apiClient.RepositoriesFileApi.RepositoriesFileFileSync(context.Background(), fileFileRepositoryHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RepositoriesFileApiService RepositoriesFileFileUpdate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var fileFileRepositoryHref string

		resp, httpRes, err := apiClient.RepositoriesFileApi.RepositoriesFileFileUpdate(context.Background(), fileFileRepositoryHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
