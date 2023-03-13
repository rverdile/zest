/*
Pulp 3 API

Testing PulpAnsibleDefaultApiV3PluginAnsibleContentCollectionsAllCollectionsApiService

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

func Test_zest_PulpAnsibleDefaultApiV3PluginAnsibleContentCollectionsAllCollectionsApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test PulpAnsibleDefaultApiV3PluginAnsibleContentCollectionsAllCollectionsApiService PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsAllCollectionsList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var distroBasePath string

		resp, httpRes, err := apiClient.PulpAnsibleDefaultApiV3PluginAnsibleContentCollectionsAllCollectionsApi.PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsAllCollectionsList(context.Background(), distroBasePath).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
