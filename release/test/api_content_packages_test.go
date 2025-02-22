/*
Pulp 3 API

Testing ContentPackagesAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package zest

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/content-services/zest/release/v2023"
)

func Test_zest_ContentPackagesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ContentPackagesAPIService ContentRpmPackagesCreate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pulpDomain string

		resp, httpRes, err := apiClient.ContentPackagesAPI.ContentRpmPackagesCreate(context.Background(), pulpDomain).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContentPackagesAPIService ContentRpmPackagesList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pulpDomain string

		resp, httpRes, err := apiClient.ContentPackagesAPI.ContentRpmPackagesList(context.Background(), pulpDomain).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContentPackagesAPIService ContentRpmPackagesRead", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var rpmPackageHref string

		resp, httpRes, err := apiClient.ContentPackagesAPI.ContentRpmPackagesRead(context.Background(), rpmPackageHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
