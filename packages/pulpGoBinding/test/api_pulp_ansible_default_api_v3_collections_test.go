/*
Pulp 3 API

Testing PulpAnsibleDefaultApiV3CollectionsApiService

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

func Test_pulpGoBinding_PulpAnsibleDefaultApiV3CollectionsApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test PulpAnsibleDefaultApiV3CollectionsApiService PulpAnsibleGalaxyDefaultApiV3CollectionsDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var name string
		var namespace string

		resp, httpRes, err := apiClient.PulpAnsibleDefaultApiV3CollectionsApi.PulpAnsibleGalaxyDefaultApiV3CollectionsDelete(context.Background(), name, namespace).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PulpAnsibleDefaultApiV3CollectionsApiService PulpAnsibleGalaxyDefaultApiV3CollectionsList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.PulpAnsibleDefaultApiV3CollectionsApi.PulpAnsibleGalaxyDefaultApiV3CollectionsList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PulpAnsibleDefaultApiV3CollectionsApiService PulpAnsibleGalaxyDefaultApiV3CollectionsRead", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var name string
		var namespace string

		resp, httpRes, err := apiClient.PulpAnsibleDefaultApiV3CollectionsApi.PulpAnsibleGalaxyDefaultApiV3CollectionsRead(context.Background(), name, namespace).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PulpAnsibleDefaultApiV3CollectionsApiService PulpAnsibleGalaxyDefaultApiV3CollectionsUpdate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var name string
		var namespace string

		resp, httpRes, err := apiClient.PulpAnsibleDefaultApiV3CollectionsApi.PulpAnsibleGalaxyDefaultApiV3CollectionsUpdate(context.Background(), name, namespace).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
