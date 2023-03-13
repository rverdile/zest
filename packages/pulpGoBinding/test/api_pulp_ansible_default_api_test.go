/*
Pulp 3 API

Testing PulpAnsibleDefaultApiApiService

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

func Test_zest_PulpAnsibleDefaultApiApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test PulpAnsibleDefaultApiApiService PulpAnsibleGalaxyDefaultApiGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.PulpAnsibleDefaultApiApi.PulpAnsibleGalaxyDefaultApiGet(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
