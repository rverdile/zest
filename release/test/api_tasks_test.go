/*
Pulp 3 API

Testing TasksApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package zest

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/content-services/zest/release"
)

func Test_zest_TasksApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test TasksApiService TasksAddRole", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var taskHref string

		resp, httpRes, err := apiClient.TasksApi.TasksAddRole(context.Background(), taskHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TasksApiService TasksCancel", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var taskHref string

		resp, httpRes, err := apiClient.TasksApi.TasksCancel(context.Background(), taskHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TasksApiService TasksDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var taskHref string

		httpRes, err := apiClient.TasksApi.TasksDelete(context.Background(), taskHref).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TasksApiService TasksList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.TasksApi.TasksList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TasksApiService TasksListRoles", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var taskHref string

		resp, httpRes, err := apiClient.TasksApi.TasksListRoles(context.Background(), taskHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TasksApiService TasksMyPermissions", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var taskHref string

		resp, httpRes, err := apiClient.TasksApi.TasksMyPermissions(context.Background(), taskHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TasksApiService TasksPurge", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.TasksApi.TasksPurge(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TasksApiService TasksRead", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var taskHref string

		resp, httpRes, err := apiClient.TasksApi.TasksRead(context.Background(), taskHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TasksApiService TasksRemoveRole", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var taskHref string

		resp, httpRes, err := apiClient.TasksApi.TasksRemoveRole(context.Background(), taskHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
