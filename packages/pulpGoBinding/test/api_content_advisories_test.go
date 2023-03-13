/*
Pulp 3 API

Testing ContentAdvisoriesApiService

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

func Test_pulpGoBinding_ContentAdvisoriesApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ContentAdvisoriesApiService ContentRpmAdvisoriesCreate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ContentAdvisoriesApi.ContentRpmAdvisoriesCreate(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContentAdvisoriesApiService ContentRpmAdvisoriesList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ContentAdvisoriesApi.ContentRpmAdvisoriesList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ContentAdvisoriesApiService ContentRpmAdvisoriesRead", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var rpmUpdateRecordHref string

		resp, httpRes, err := apiClient.ContentAdvisoriesApi.ContentRpmAdvisoriesRead(context.Background(), rpmUpdateRecordHref).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
