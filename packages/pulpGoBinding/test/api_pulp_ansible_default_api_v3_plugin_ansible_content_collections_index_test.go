/*
Pulp 3 API

Testing PulpAnsibleDefaultApiV3PluginAnsibleContentCollectionsIndexApiService

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

func Test_zest_PulpAnsibleDefaultApiV3PluginAnsibleContentCollectionsIndexApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test PulpAnsibleDefaultApiV3PluginAnsibleContentCollectionsIndexApiService PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsIndexDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var distroBasePath string
		var name string
		var namespace string

		resp, httpRes, err := apiClient.PulpAnsibleDefaultApiV3PluginAnsibleContentCollectionsIndexApi.PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsIndexDelete(context.Background(), distroBasePath, name, namespace).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PulpAnsibleDefaultApiV3PluginAnsibleContentCollectionsIndexApiService PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsIndexList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var distroBasePath string

		resp, httpRes, err := apiClient.PulpAnsibleDefaultApiV3PluginAnsibleContentCollectionsIndexApi.PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsIndexList(context.Background(), distroBasePath).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PulpAnsibleDefaultApiV3PluginAnsibleContentCollectionsIndexApiService PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsIndexRead", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var distroBasePath string
		var name string
		var namespace string

		resp, httpRes, err := apiClient.PulpAnsibleDefaultApiV3PluginAnsibleContentCollectionsIndexApi.PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsIndexRead(context.Background(), distroBasePath, name, namespace).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PulpAnsibleDefaultApiV3PluginAnsibleContentCollectionsIndexApiService PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsIndexUpdate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var distroBasePath string
		var name string
		var namespace string

		resp, httpRes, err := apiClient.PulpAnsibleDefaultApiV3PluginAnsibleContentCollectionsIndexApi.PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsIndexUpdate(context.Background(), distroBasePath, name, namespace).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
