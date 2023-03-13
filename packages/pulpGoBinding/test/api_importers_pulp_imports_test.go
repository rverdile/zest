/*
Pulp 3 API

Testing ImportersPulpImportsApiService

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

func Test_pulpGoBinding_ImportersPulpImportsApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ImportersPulpImportsApiService ImportersCorePulpImportsCreate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pulpImporterHref string

		resp, httpRes, err := apiClient.ImportersPulpImportsApi.ImportersCorePulpImportsCreate(context.Background(), pulpImporterHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ImportersPulpImportsApiService ImportersCorePulpImportsDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pulpPulpImportHref string

		httpRes, err := apiClient.ImportersPulpImportsApi.ImportersCorePulpImportsDelete(context.Background(), pulpPulpImportHref).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ImportersPulpImportsApiService ImportersCorePulpImportsList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pulpImporterHref string

		resp, httpRes, err := apiClient.ImportersPulpImportsApi.ImportersCorePulpImportsList(context.Background(), pulpImporterHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ImportersPulpImportsApiService ImportersCorePulpImportsRead", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pulpPulpImportHref string

		resp, httpRes, err := apiClient.ImportersPulpImportsApi.ImportersCorePulpImportsRead(context.Background(), pulpPulpImportHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
