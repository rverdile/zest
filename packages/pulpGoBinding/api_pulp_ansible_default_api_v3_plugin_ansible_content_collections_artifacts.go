/*
Pulp 3 API

Fetch, Upload, Organize, and Distribute Software Packages

API version: v3
Contact: pulp-list@redhat.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pulpGoBinding

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
	"reflect"
)


// PulpAnsibleDefaultApiV3PluginAnsibleContentCollectionsArtifactsApiService PulpAnsibleDefaultApiV3PluginAnsibleContentCollectionsArtifactsApi service
type PulpAnsibleDefaultApiV3PluginAnsibleContentCollectionsArtifactsApiService service

type PulpAnsibleDefaultApiV3PluginAnsibleContentCollectionsArtifactsApiPulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsArtifactsDownloadRequest struct {
	ctx context.Context
	ApiService *PulpAnsibleDefaultApiV3PluginAnsibleContentCollectionsArtifactsApiService
	distroBasePath string
	filename string
	fields *[]string
	excludeFields *[]string
}

// A list of fields to include in the response.
func (r PulpAnsibleDefaultApiV3PluginAnsibleContentCollectionsArtifactsApiPulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsArtifactsDownloadRequest) Fields(fields []string) PulpAnsibleDefaultApiV3PluginAnsibleContentCollectionsArtifactsApiPulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsArtifactsDownloadRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r PulpAnsibleDefaultApiV3PluginAnsibleContentCollectionsArtifactsApiPulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsArtifactsDownloadRequest) ExcludeFields(excludeFields []string) PulpAnsibleDefaultApiV3PluginAnsibleContentCollectionsArtifactsApiPulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsArtifactsDownloadRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r PulpAnsibleDefaultApiV3PluginAnsibleContentCollectionsArtifactsApiPulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsArtifactsDownloadRequest) Execute() (*http.Response, error) {
	return r.ApiService.PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsArtifactsDownloadExecute(r)
}

/*
PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsArtifactsDownload Method for PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsArtifactsDownload

Collection download endpoint.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param distroBasePath
 @param filename
 @return PulpAnsibleDefaultApiV3PluginAnsibleContentCollectionsArtifactsApiPulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsArtifactsDownloadRequest
*/
func (a *PulpAnsibleDefaultApiV3PluginAnsibleContentCollectionsArtifactsApiService) PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsArtifactsDownload(ctx context.Context, distroBasePath string, filename string) PulpAnsibleDefaultApiV3PluginAnsibleContentCollectionsArtifactsApiPulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsArtifactsDownloadRequest {
	return PulpAnsibleDefaultApiV3PluginAnsibleContentCollectionsArtifactsApiPulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsArtifactsDownloadRequest{
		ApiService: a,
		ctx: ctx,
		distroBasePath: distroBasePath,
		filename: filename,
	}
}

// Execute executes the request
func (a *PulpAnsibleDefaultApiV3PluginAnsibleContentCollectionsArtifactsApiService) PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsArtifactsDownloadExecute(r PulpAnsibleDefaultApiV3PluginAnsibleContentCollectionsArtifactsApiPulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsArtifactsDownloadRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PulpAnsibleDefaultApiV3PluginAnsibleContentCollectionsArtifactsApiService.PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentCollectionsArtifactsDownload")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp_ansible/galaxy/default/api/v3/plugin/ansible/content/{distro_base_path}/collections/artifacts/{filename}"
	localVarPath = strings.Replace(localVarPath, "{"+"distro_base_path"+"}", url.PathEscape(parameterValueToString(r.distroBasePath, "distroBasePath")), -1)
        localVarPath = strings.Replace(localVarPath, "/%2F", "/", -1)

	localVarPath = strings.Replace(localVarPath, "{"+"filename"+"}", url.PathEscape(parameterValueToString(r.filename, "filename")), -1)
        localVarPath = strings.Replace(localVarPath, "/%2F", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.fields != nil {
		t := *r.fields
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "fields", s.Index(i), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "fields", t, "multi")
		}
	}
	if r.excludeFields != nil {
		t := *r.excludeFields
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "exclude_fields", s.Index(i), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "exclude_fields", t, "multi")
		}
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}
