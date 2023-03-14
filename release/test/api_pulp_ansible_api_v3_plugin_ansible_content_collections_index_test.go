/*
Pulp 3 API

Testing PulpAnsibleApiV3PluginAnsibleContentCollectionsIndexApiService

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

func Test_zest_PulpAnsibleApiV3PluginAnsibleContentCollectionsIndexApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test PulpAnsibleApiV3PluginAnsibleContentCollectionsIndexApiService PulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsIndexDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var distroBasePath string
		var name string
		var namespace string
		var path string

		resp, httpRes, err := apiClient.PulpAnsibleApiV3PluginAnsibleContentCollectionsIndexApi.PulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsIndexDelete(context.Background(), distroBasePath, name, namespace, path).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PulpAnsibleApiV3PluginAnsibleContentCollectionsIndexApiService PulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsIndexList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var distroBasePath string
		var path string

		resp, httpRes, err := apiClient.PulpAnsibleApiV3PluginAnsibleContentCollectionsIndexApi.PulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsIndexList(context.Background(), distroBasePath, path).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PulpAnsibleApiV3PluginAnsibleContentCollectionsIndexApiService PulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsIndexRead", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var distroBasePath string
		var name string
		var namespace string
		var path string

		resp, httpRes, err := apiClient.PulpAnsibleApiV3PluginAnsibleContentCollectionsIndexApi.PulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsIndexRead(context.Background(), distroBasePath, name, namespace, path).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PulpAnsibleApiV3PluginAnsibleContentCollectionsIndexApiService PulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsIndexUpdate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var distroBasePath string
		var name string
		var namespace string
		var path string

		resp, httpRes, err := apiClient.PulpAnsibleApiV3PluginAnsibleContentCollectionsIndexApi.PulpAnsibleGalaxyApiV3PluginAnsibleContentCollectionsIndexUpdate(context.Background(), distroBasePath, name, namespace, path).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
