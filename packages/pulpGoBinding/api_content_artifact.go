/*
Pulp 3 API

Fetch, Upload, Organize, and Distribute Software Packages

API version: v3
Contact: pulp-list@redhat.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package zest

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
	"reflect"
)


// ContentArtifactApiService ContentArtifactApi service
type ContentArtifactApiService service

type ContentArtifactApiContentMavenArtifactCreateRequest struct {
	ctx context.Context
	ApiService *ContentArtifactApiService
	mavenMavenArtifact *MavenMavenArtifact
}

func (r ContentArtifactApiContentMavenArtifactCreateRequest) MavenMavenArtifact(mavenMavenArtifact MavenMavenArtifact) ContentArtifactApiContentMavenArtifactCreateRequest {
	r.mavenMavenArtifact = &mavenMavenArtifact
	return r
}

func (r ContentArtifactApiContentMavenArtifactCreateRequest) Execute() (*MavenMavenArtifactResponse, *http.Response, error) {
	return r.ApiService.ContentMavenArtifactCreateExecute(r)
}

/*
ContentMavenArtifactCreate Create a maven artifact

A ViewSet for MavenArtifact.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ContentArtifactApiContentMavenArtifactCreateRequest
*/
func (a *ContentArtifactApiService) ContentMavenArtifactCreate(ctx context.Context) ContentArtifactApiContentMavenArtifactCreateRequest {
	return ContentArtifactApiContentMavenArtifactCreateRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return MavenMavenArtifactResponse
func (a *ContentArtifactApiService) ContentMavenArtifactCreateExecute(r ContentArtifactApiContentMavenArtifactCreateRequest) (*MavenMavenArtifactResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MavenMavenArtifactResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentArtifactApiService.ContentMavenArtifactCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp/api/v3/content/maven/artifact/"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.mavenMavenArtifact == nil {
		return localVarReturnValue, nil, reportError("mavenMavenArtifact is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json", "application/x-www-form-urlencoded", "multipart/form-data"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.mavenMavenArtifact
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ContentArtifactApiContentMavenArtifactListRequest struct {
	ctx context.Context
	ApiService *ContentArtifactApiService
	artifactId *string
	filename *string
	groupId *string
	limit *int32
	offset *int32
	ordering *[]string
	repositoryVersion *string
	repositoryVersionAdded *string
	repositoryVersionRemoved *string
	version *string
	fields *[]string
	excludeFields *[]string
}

// Filter results where artifact_id matches value
func (r ContentArtifactApiContentMavenArtifactListRequest) ArtifactId(artifactId string) ContentArtifactApiContentMavenArtifactListRequest {
	r.artifactId = &artifactId
	return r
}

// Filter results where filename matches value
func (r ContentArtifactApiContentMavenArtifactListRequest) Filename(filename string) ContentArtifactApiContentMavenArtifactListRequest {
	r.filename = &filename
	return r
}

// Filter results where group_id matches value
func (r ContentArtifactApiContentMavenArtifactListRequest) GroupId(groupId string) ContentArtifactApiContentMavenArtifactListRequest {
	r.groupId = &groupId
	return r
}

// Number of results to return per page.
func (r ContentArtifactApiContentMavenArtifactListRequest) Limit(limit int32) ContentArtifactApiContentMavenArtifactListRequest {
	r.limit = &limit
	return r
}

// The initial index from which to return the results.
func (r ContentArtifactApiContentMavenArtifactListRequest) Offset(offset int32) ContentArtifactApiContentMavenArtifactListRequest {
	r.offset = &offset
	return r
}

// Ordering
func (r ContentArtifactApiContentMavenArtifactListRequest) Ordering(ordering []string) ContentArtifactApiContentMavenArtifactListRequest {
	r.ordering = &ordering
	return r
}

// Repository Version referenced by HREF
func (r ContentArtifactApiContentMavenArtifactListRequest) RepositoryVersion(repositoryVersion string) ContentArtifactApiContentMavenArtifactListRequest {
	r.repositoryVersion = &repositoryVersion
	return r
}

// Repository Version referenced by HREF
func (r ContentArtifactApiContentMavenArtifactListRequest) RepositoryVersionAdded(repositoryVersionAdded string) ContentArtifactApiContentMavenArtifactListRequest {
	r.repositoryVersionAdded = &repositoryVersionAdded
	return r
}

// Repository Version referenced by HREF
func (r ContentArtifactApiContentMavenArtifactListRequest) RepositoryVersionRemoved(repositoryVersionRemoved string) ContentArtifactApiContentMavenArtifactListRequest {
	r.repositoryVersionRemoved = &repositoryVersionRemoved
	return r
}

// Filter results where version matches value
func (r ContentArtifactApiContentMavenArtifactListRequest) Version(version string) ContentArtifactApiContentMavenArtifactListRequest {
	r.version = &version
	return r
}

// A list of fields to include in the response.
func (r ContentArtifactApiContentMavenArtifactListRequest) Fields(fields []string) ContentArtifactApiContentMavenArtifactListRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r ContentArtifactApiContentMavenArtifactListRequest) ExcludeFields(excludeFields []string) ContentArtifactApiContentMavenArtifactListRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r ContentArtifactApiContentMavenArtifactListRequest) Execute() (*PaginatedmavenMavenArtifactResponseList, *http.Response, error) {
	return r.ApiService.ContentMavenArtifactListExecute(r)
}

/*
ContentMavenArtifactList List maven artifacts

A ViewSet for MavenArtifact.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ContentArtifactApiContentMavenArtifactListRequest
*/
func (a *ContentArtifactApiService) ContentMavenArtifactList(ctx context.Context) ContentArtifactApiContentMavenArtifactListRequest {
	return ContentArtifactApiContentMavenArtifactListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PaginatedmavenMavenArtifactResponseList
func (a *ContentArtifactApiService) ContentMavenArtifactListExecute(r ContentArtifactApiContentMavenArtifactListRequest) (*PaginatedmavenMavenArtifactResponseList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedmavenMavenArtifactResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentArtifactApiService.ContentMavenArtifactList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp/api/v3/content/maven/artifact/"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.artifactId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "artifact_id", r.artifactId, "")
	}
	if r.filename != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filename", r.filename, "")
	}
	if r.groupId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "group_id", r.groupId, "")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	}
	if r.offset != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "offset", r.offset, "")
	}
	if r.ordering != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "ordering", r.ordering, "csv")
	}
	if r.repositoryVersion != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "repository_version", r.repositoryVersion, "")
	}
	if r.repositoryVersionAdded != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "repository_version_added", r.repositoryVersionAdded, "")
	}
	if r.repositoryVersionRemoved != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "repository_version_removed", r.repositoryVersionRemoved, "")
	}
	if r.version != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "version", r.version, "")
	}
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
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ContentArtifactApiContentMavenArtifactReadRequest struct {
	ctx context.Context
	ApiService *ContentArtifactApiService
	mavenMavenArtifactHref string
	fields *[]string
	excludeFields *[]string
}

// A list of fields to include in the response.
func (r ContentArtifactApiContentMavenArtifactReadRequest) Fields(fields []string) ContentArtifactApiContentMavenArtifactReadRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r ContentArtifactApiContentMavenArtifactReadRequest) ExcludeFields(excludeFields []string) ContentArtifactApiContentMavenArtifactReadRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r ContentArtifactApiContentMavenArtifactReadRequest) Execute() (*MavenMavenArtifactResponse, *http.Response, error) {
	return r.ApiService.ContentMavenArtifactReadExecute(r)
}

/*
ContentMavenArtifactRead Inspect a maven artifact

A ViewSet for MavenArtifact.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param mavenMavenArtifactHref
 @return ContentArtifactApiContentMavenArtifactReadRequest
*/
func (a *ContentArtifactApiService) ContentMavenArtifactRead(ctx context.Context, mavenMavenArtifactHref string) ContentArtifactApiContentMavenArtifactReadRequest {
	return ContentArtifactApiContentMavenArtifactReadRequest{
		ApiService: a,
		ctx: ctx,
		mavenMavenArtifactHref: mavenMavenArtifactHref,
	}
}

// Execute executes the request
//  @return MavenMavenArtifactResponse
func (a *ContentArtifactApiService) ContentMavenArtifactReadExecute(r ContentArtifactApiContentMavenArtifactReadRequest) (*MavenMavenArtifactResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MavenMavenArtifactResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentArtifactApiService.ContentMavenArtifactRead")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{maven_maven_artifact_href}"
	localVarPath = strings.Replace(localVarPath, "{"+"maven_maven_artifact_href"+"}", url.PathEscape(parameterValueToString(r.mavenMavenArtifactHref, "mavenMavenArtifactHref")), -1)
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
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
