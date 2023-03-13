/*
Pulp 3 API

Testing ContentPackageReleaseComponentsApiService

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

func Test_pulpGoBinding_ContentPackageReleaseComponentsApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ContentPackageReleaseComponentsApiService ContentDebPackageReleaseComponentsCreate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ContentPackageReleaseComponentsApi.ContentDebPackageReleaseComponentsCreate(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContentPackageReleaseComponentsApiService ContentDebPackageReleaseComponentsList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ContentPackageReleaseComponentsApi.ContentDebPackageReleaseComponentsList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContentPackageReleaseComponentsApiService ContentDebPackageReleaseComponentsRead", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var debPackageReleaseComponentHref string

		resp, httpRes, err := apiClient.ContentPackageReleaseComponentsApi.ContentDebPackageReleaseComponentsRead(context.Background(), debPackageReleaseComponentHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
