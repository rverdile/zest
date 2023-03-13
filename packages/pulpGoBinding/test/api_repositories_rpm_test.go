/*
Pulp 3 API

Testing RepositoriesRpmApiService

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

func Test_pulpGoBinding_RepositoriesRpmApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test RepositoriesRpmApiService RepositoriesRpmRpmAddRole", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var rpmRpmRepositoryHref string

		resp, httpRes, err := apiClient.RepositoriesRpmApi.RepositoriesRpmRpmAddRole(context.Background(), rpmRpmRepositoryHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RepositoriesRpmApiService RepositoriesRpmRpmCreate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.RepositoriesRpmApi.RepositoriesRpmRpmCreate(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RepositoriesRpmApiService RepositoriesRpmRpmDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var rpmRpmRepositoryHref string

		resp, httpRes, err := apiClient.RepositoriesRpmApi.RepositoriesRpmRpmDelete(context.Background(), rpmRpmRepositoryHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RepositoriesRpmApiService RepositoriesRpmRpmList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.RepositoriesRpmApi.RepositoriesRpmRpmList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RepositoriesRpmApiService RepositoriesRpmRpmListRoles", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var rpmRpmRepositoryHref string

		resp, httpRes, err := apiClient.RepositoriesRpmApi.RepositoriesRpmRpmListRoles(context.Background(), rpmRpmRepositoryHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RepositoriesRpmApiService RepositoriesRpmRpmModify", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var rpmRpmRepositoryHref string

		resp, httpRes, err := apiClient.RepositoriesRpmApi.RepositoriesRpmRpmModify(context.Background(), rpmRpmRepositoryHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RepositoriesRpmApiService RepositoriesRpmRpmMyPermissions", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var rpmRpmRepositoryHref string

		resp, httpRes, err := apiClient.RepositoriesRpmApi.RepositoriesRpmRpmMyPermissions(context.Background(), rpmRpmRepositoryHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RepositoriesRpmApiService RepositoriesRpmRpmPartialUpdate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var rpmRpmRepositoryHref string

		resp, httpRes, err := apiClient.RepositoriesRpmApi.RepositoriesRpmRpmPartialUpdate(context.Background(), rpmRpmRepositoryHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RepositoriesRpmApiService RepositoriesRpmRpmRead", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var rpmRpmRepositoryHref string

		resp, httpRes, err := apiClient.RepositoriesRpmApi.RepositoriesRpmRpmRead(context.Background(), rpmRpmRepositoryHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RepositoriesRpmApiService RepositoriesRpmRpmRemoveRole", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var rpmRpmRepositoryHref string

		resp, httpRes, err := apiClient.RepositoriesRpmApi.RepositoriesRpmRpmRemoveRole(context.Background(), rpmRpmRepositoryHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RepositoriesRpmApiService RepositoriesRpmRpmSync", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var rpmRpmRepositoryHref string

		resp, httpRes, err := apiClient.RepositoriesRpmApi.RepositoriesRpmRpmSync(context.Background(), rpmRpmRepositoryHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RepositoriesRpmApiService RepositoriesRpmRpmUpdate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var rpmRpmRepositoryHref string

		resp, httpRes, err := apiClient.RepositoriesRpmApi.RepositoriesRpmRpmUpdate(context.Background(), rpmRpmRepositoryHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
