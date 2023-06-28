/*
Pulp 3 API

Testing ImportersPulpImportsAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package zest

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/content-services/zest/release/v3"
)

func Test_zest_ImportersPulpImportsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ImportersPulpImportsAPIService ImportersCorePulpImportsCreate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pulpImporterHref string

		resp, httpRes, err := apiClient.ImportersPulpImportsAPI.ImportersCorePulpImportsCreate(context.Background(), pulpImporterHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ImportersPulpImportsAPIService ImportersCorePulpImportsDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pulpPulpImportHref string

		httpRes, err := apiClient.ImportersPulpImportsAPI.ImportersCorePulpImportsDelete(context.Background(), pulpPulpImportHref).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ImportersPulpImportsAPIService ImportersCorePulpImportsList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pulpImporterHref string

		resp, httpRes, err := apiClient.ImportersPulpImportsAPI.ImportersCorePulpImportsList(context.Background(), pulpImporterHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ImportersPulpImportsAPIService ImportersCorePulpImportsRead", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pulpPulpImportHref string

		resp, httpRes, err := apiClient.ImportersPulpImportsAPI.ImportersCorePulpImportsRead(context.Background(), pulpPulpImportHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
