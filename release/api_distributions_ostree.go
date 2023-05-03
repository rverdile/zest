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


// DistributionsOstreeApiService DistributionsOstreeApi service
type DistributionsOstreeApiService service

type DistributionsOstreeApiDistributionsOstreeOstreeCreateRequest struct {
	ctx context.Context
	ApiService *DistributionsOstreeApiService
	ostreeOstreeDistribution *OstreeOstreeDistribution
}

func (r DistributionsOstreeApiDistributionsOstreeOstreeCreateRequest) OstreeOstreeDistribution(ostreeOstreeDistribution OstreeOstreeDistribution) DistributionsOstreeApiDistributionsOstreeOstreeCreateRequest {
	r.ostreeOstreeDistribution = &ostreeOstreeDistribution
	return r
}

func (r DistributionsOstreeApiDistributionsOstreeOstreeCreateRequest) Execute() (*AsyncOperationResponse, *http.Response, error) {
	return r.ApiService.DistributionsOstreeOstreeCreateExecute(r)
}

/*
DistributionsOstreeOstreeCreate Create an ostree distribution

Trigger an asynchronous create task

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return DistributionsOstreeApiDistributionsOstreeOstreeCreateRequest
*/
func (a *DistributionsOstreeApiService) DistributionsOstreeOstreeCreate(ctx context.Context) DistributionsOstreeApiDistributionsOstreeOstreeCreateRequest {
	return DistributionsOstreeApiDistributionsOstreeOstreeCreateRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AsyncOperationResponse
func (a *DistributionsOstreeApiService) DistributionsOstreeOstreeCreateExecute(r DistributionsOstreeApiDistributionsOstreeOstreeCreateRequest) (*AsyncOperationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AsyncOperationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DistributionsOstreeApiService.DistributionsOstreeOstreeCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp/api/v3/distributions/ostree/ostree/"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.ostreeOstreeDistribution == nil {
		return localVarReturnValue, nil, reportError("ostreeOstreeDistribution is required and must be specified")
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
	localVarPostBody = r.ostreeOstreeDistribution
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

type DistributionsOstreeApiDistributionsOstreeOstreeDeleteRequest struct {
	ctx context.Context
	ApiService *DistributionsOstreeApiService
	ostreeOstreeDistributionHref string
}

func (r DistributionsOstreeApiDistributionsOstreeOstreeDeleteRequest) Execute() (*AsyncOperationResponse, *http.Response, error) {
	return r.ApiService.DistributionsOstreeOstreeDeleteExecute(r)
}

/*
DistributionsOstreeOstreeDelete Delete an ostree distribution

Trigger an asynchronous delete task

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ostreeOstreeDistributionHref
 @return DistributionsOstreeApiDistributionsOstreeOstreeDeleteRequest
*/
func (a *DistributionsOstreeApiService) DistributionsOstreeOstreeDelete(ctx context.Context, ostreeOstreeDistributionHref string) DistributionsOstreeApiDistributionsOstreeOstreeDeleteRequest {
	return DistributionsOstreeApiDistributionsOstreeOstreeDeleteRequest{
		ApiService: a,
		ctx: ctx,
		ostreeOstreeDistributionHref: ostreeOstreeDistributionHref,
	}
}

// Execute executes the request
//  @return AsyncOperationResponse
func (a *DistributionsOstreeApiService) DistributionsOstreeOstreeDeleteExecute(r DistributionsOstreeApiDistributionsOstreeOstreeDeleteRequest) (*AsyncOperationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AsyncOperationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DistributionsOstreeApiService.DistributionsOstreeOstreeDelete")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{ostree_ostree_distribution_href}"
	localVarPath = strings.Replace(localVarPath, "{"+"ostree_ostree_distribution_href"+"}", url.PathEscape(parameterValueToString(r.ostreeOstreeDistributionHref, "ostreeOstreeDistributionHref")), -1)
        localVarPath = strings.Replace(localVarPath, "/%2F", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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

type DistributionsOstreeApiDistributionsOstreeOstreeListRequest struct {
	ctx context.Context
	ApiService *DistributionsOstreeApiService
	basePath *string
	basePathContains *string
	basePathIcontains *string
	basePathIn *[]string
	limit *int32
	name *string
	nameContains *string
	nameIcontains *string
	nameIn *[]string
	nameStartswith *string
	offset *int32
	ordering *[]string
	pulpHrefIn *[]string
	pulpIdIn *[]string
	pulpLabelSelect *string
	repository *string
	repositoryIn *[]string
	withContent *string
	fields *[]string
	excludeFields *[]string
}

// Filter results where base_path matches value
func (r DistributionsOstreeApiDistributionsOstreeOstreeListRequest) BasePath(basePath string) DistributionsOstreeApiDistributionsOstreeOstreeListRequest {
	r.basePath = &basePath
	return r
}

// Filter results where base_path contains value
func (r DistributionsOstreeApiDistributionsOstreeOstreeListRequest) BasePathContains(basePathContains string) DistributionsOstreeApiDistributionsOstreeOstreeListRequest {
	r.basePathContains = &basePathContains
	return r
}

// Filter results where base_path contains value
func (r DistributionsOstreeApiDistributionsOstreeOstreeListRequest) BasePathIcontains(basePathIcontains string) DistributionsOstreeApiDistributionsOstreeOstreeListRequest {
	r.basePathIcontains = &basePathIcontains
	return r
}

// Filter results where base_path is in a comma-separated list of values
func (r DistributionsOstreeApiDistributionsOstreeOstreeListRequest) BasePathIn(basePathIn []string) DistributionsOstreeApiDistributionsOstreeOstreeListRequest {
	r.basePathIn = &basePathIn
	return r
}

// Number of results to return per page.
func (r DistributionsOstreeApiDistributionsOstreeOstreeListRequest) Limit(limit int32) DistributionsOstreeApiDistributionsOstreeOstreeListRequest {
	r.limit = &limit
	return r
}

// Filter results where name matches value
func (r DistributionsOstreeApiDistributionsOstreeOstreeListRequest) Name(name string) DistributionsOstreeApiDistributionsOstreeOstreeListRequest {
	r.name = &name
	return r
}

// Filter results where name contains value
func (r DistributionsOstreeApiDistributionsOstreeOstreeListRequest) NameContains(nameContains string) DistributionsOstreeApiDistributionsOstreeOstreeListRequest {
	r.nameContains = &nameContains
	return r
}

// Filter results where name contains value
func (r DistributionsOstreeApiDistributionsOstreeOstreeListRequest) NameIcontains(nameIcontains string) DistributionsOstreeApiDistributionsOstreeOstreeListRequest {
	r.nameIcontains = &nameIcontains
	return r
}

// Filter results where name is in a comma-separated list of values
func (r DistributionsOstreeApiDistributionsOstreeOstreeListRequest) NameIn(nameIn []string) DistributionsOstreeApiDistributionsOstreeOstreeListRequest {
	r.nameIn = &nameIn
	return r
}

// Filter results where name starts with value
func (r DistributionsOstreeApiDistributionsOstreeOstreeListRequest) NameStartswith(nameStartswith string) DistributionsOstreeApiDistributionsOstreeOstreeListRequest {
	r.nameStartswith = &nameStartswith
	return r
}

// The initial index from which to return the results.
func (r DistributionsOstreeApiDistributionsOstreeOstreeListRequest) Offset(offset int32) DistributionsOstreeApiDistributionsOstreeOstreeListRequest {
	r.offset = &offset
	return r
}

// Ordering  * &#x60;pulp_id&#x60; - Pulp id * &#x60;-pulp_id&#x60; - Pulp id (descending) * &#x60;pulp_created&#x60; - Pulp created * &#x60;-pulp_created&#x60; - Pulp created (descending) * &#x60;pulp_last_updated&#x60; - Pulp last updated * &#x60;-pulp_last_updated&#x60; - Pulp last updated (descending) * &#x60;pulp_type&#x60; - Pulp type * &#x60;-pulp_type&#x60; - Pulp type (descending) * &#x60;name&#x60; - Name * &#x60;-name&#x60; - Name (descending) * &#x60;pulp_labels&#x60; - Pulp labels * &#x60;-pulp_labels&#x60; - Pulp labels (descending) * &#x60;base_path&#x60; - Base path * &#x60;-base_path&#x60; - Base path (descending) * &#x60;pk&#x60; - Pk * &#x60;-pk&#x60; - Pk (descending)
func (r DistributionsOstreeApiDistributionsOstreeOstreeListRequest) Ordering(ordering []string) DistributionsOstreeApiDistributionsOstreeOstreeListRequest {
	r.ordering = &ordering
	return r
}

// Multiple values may be separated by commas.
func (r DistributionsOstreeApiDistributionsOstreeOstreeListRequest) PulpHrefIn(pulpHrefIn []string) DistributionsOstreeApiDistributionsOstreeOstreeListRequest {
	r.pulpHrefIn = &pulpHrefIn
	return r
}

// Multiple values may be separated by commas.
func (r DistributionsOstreeApiDistributionsOstreeOstreeListRequest) PulpIdIn(pulpIdIn []string) DistributionsOstreeApiDistributionsOstreeOstreeListRequest {
	r.pulpIdIn = &pulpIdIn
	return r
}

// Filter labels by search string
func (r DistributionsOstreeApiDistributionsOstreeOstreeListRequest) PulpLabelSelect(pulpLabelSelect string) DistributionsOstreeApiDistributionsOstreeOstreeListRequest {
	r.pulpLabelSelect = &pulpLabelSelect
	return r
}

// Filter results where repository matches value
func (r DistributionsOstreeApiDistributionsOstreeOstreeListRequest) Repository(repository string) DistributionsOstreeApiDistributionsOstreeOstreeListRequest {
	r.repository = &repository
	return r
}

// Filter results where repository is in a comma-separated list of values
func (r DistributionsOstreeApiDistributionsOstreeOstreeListRequest) RepositoryIn(repositoryIn []string) DistributionsOstreeApiDistributionsOstreeOstreeListRequest {
	r.repositoryIn = &repositoryIn
	return r
}

// Filter distributions based on the content served by them
func (r DistributionsOstreeApiDistributionsOstreeOstreeListRequest) WithContent(withContent string) DistributionsOstreeApiDistributionsOstreeOstreeListRequest {
	r.withContent = &withContent
	return r
}

// A list of fields to include in the response.
func (r DistributionsOstreeApiDistributionsOstreeOstreeListRequest) Fields(fields []string) DistributionsOstreeApiDistributionsOstreeOstreeListRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r DistributionsOstreeApiDistributionsOstreeOstreeListRequest) ExcludeFields(excludeFields []string) DistributionsOstreeApiDistributionsOstreeOstreeListRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r DistributionsOstreeApiDistributionsOstreeOstreeListRequest) Execute() (*PaginatedostreeOstreeDistributionResponseList, *http.Response, error) {
	return r.ApiService.DistributionsOstreeOstreeListExecute(r)
}

/*
DistributionsOstreeOstreeList List ostree distributions

A ViewSet class for OSTree distributions.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return DistributionsOstreeApiDistributionsOstreeOstreeListRequest
*/
func (a *DistributionsOstreeApiService) DistributionsOstreeOstreeList(ctx context.Context) DistributionsOstreeApiDistributionsOstreeOstreeListRequest {
	return DistributionsOstreeApiDistributionsOstreeOstreeListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PaginatedostreeOstreeDistributionResponseList
func (a *DistributionsOstreeApiService) DistributionsOstreeOstreeListExecute(r DistributionsOstreeApiDistributionsOstreeOstreeListRequest) (*PaginatedostreeOstreeDistributionResponseList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedostreeOstreeDistributionResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DistributionsOstreeApiService.DistributionsOstreeOstreeList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp/api/v3/distributions/ostree/ostree/"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.basePath != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "base_path", r.basePath, "")
	}
	if r.basePathContains != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "base_path__contains", r.basePathContains, "")
	}
	if r.basePathIcontains != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "base_path__icontains", r.basePathIcontains, "")
	}
	if r.basePathIn != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "base_path__in", r.basePathIn, "csv")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	}
	if r.name != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "name", r.name, "")
	}
	if r.nameContains != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "name__contains", r.nameContains, "")
	}
	if r.nameIcontains != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "name__icontains", r.nameIcontains, "")
	}
	if r.nameIn != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "name__in", r.nameIn, "csv")
	}
	if r.nameStartswith != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "name__startswith", r.nameStartswith, "")
	}
	if r.offset != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "offset", r.offset, "")
	}
	if r.ordering != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "ordering", r.ordering, "csv")
	}
	if r.pulpHrefIn != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pulp_href__in", r.pulpHrefIn, "csv")
	}
	if r.pulpIdIn != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pulp_id__in", r.pulpIdIn, "csv")
	}
	if r.pulpLabelSelect != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pulp_label_select", r.pulpLabelSelect, "")
	}
	if r.repository != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "repository", r.repository, "")
	}
	if r.repositoryIn != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "repository__in", r.repositoryIn, "csv")
	}
	if r.withContent != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "with_content", r.withContent, "")
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

type DistributionsOstreeApiDistributionsOstreeOstreePartialUpdateRequest struct {
	ctx context.Context
	ApiService *DistributionsOstreeApiService
	ostreeOstreeDistributionHref string
	patchedostreeOstreeDistribution *PatchedostreeOstreeDistribution
}

func (r DistributionsOstreeApiDistributionsOstreeOstreePartialUpdateRequest) PatchedostreeOstreeDistribution(patchedostreeOstreeDistribution PatchedostreeOstreeDistribution) DistributionsOstreeApiDistributionsOstreeOstreePartialUpdateRequest {
	r.patchedostreeOstreeDistribution = &patchedostreeOstreeDistribution
	return r
}

func (r DistributionsOstreeApiDistributionsOstreeOstreePartialUpdateRequest) Execute() (*AsyncOperationResponse, *http.Response, error) {
	return r.ApiService.DistributionsOstreeOstreePartialUpdateExecute(r)
}

/*
DistributionsOstreeOstreePartialUpdate Update an ostree distribution

Trigger an asynchronous partial update task

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ostreeOstreeDistributionHref
 @return DistributionsOstreeApiDistributionsOstreeOstreePartialUpdateRequest
*/
func (a *DistributionsOstreeApiService) DistributionsOstreeOstreePartialUpdate(ctx context.Context, ostreeOstreeDistributionHref string) DistributionsOstreeApiDistributionsOstreeOstreePartialUpdateRequest {
	return DistributionsOstreeApiDistributionsOstreeOstreePartialUpdateRequest{
		ApiService: a,
		ctx: ctx,
		ostreeOstreeDistributionHref: ostreeOstreeDistributionHref,
	}
}

// Execute executes the request
//  @return AsyncOperationResponse
func (a *DistributionsOstreeApiService) DistributionsOstreeOstreePartialUpdateExecute(r DistributionsOstreeApiDistributionsOstreeOstreePartialUpdateRequest) (*AsyncOperationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AsyncOperationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DistributionsOstreeApiService.DistributionsOstreeOstreePartialUpdate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{ostree_ostree_distribution_href}"
	localVarPath = strings.Replace(localVarPath, "{"+"ostree_ostree_distribution_href"+"}", url.PathEscape(parameterValueToString(r.ostreeOstreeDistributionHref, "ostreeOstreeDistributionHref")), -1)
        localVarPath = strings.Replace(localVarPath, "/%2F", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.patchedostreeOstreeDistribution == nil {
		return localVarReturnValue, nil, reportError("patchedostreeOstreeDistribution is required and must be specified")
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
	localVarPostBody = r.patchedostreeOstreeDistribution
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

type DistributionsOstreeApiDistributionsOstreeOstreeReadRequest struct {
	ctx context.Context
	ApiService *DistributionsOstreeApiService
	ostreeOstreeDistributionHref string
	fields *[]string
	excludeFields *[]string
}

// A list of fields to include in the response.
func (r DistributionsOstreeApiDistributionsOstreeOstreeReadRequest) Fields(fields []string) DistributionsOstreeApiDistributionsOstreeOstreeReadRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r DistributionsOstreeApiDistributionsOstreeOstreeReadRequest) ExcludeFields(excludeFields []string) DistributionsOstreeApiDistributionsOstreeOstreeReadRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r DistributionsOstreeApiDistributionsOstreeOstreeReadRequest) Execute() (*OstreeOstreeDistributionResponse, *http.Response, error) {
	return r.ApiService.DistributionsOstreeOstreeReadExecute(r)
}

/*
DistributionsOstreeOstreeRead Inspect an ostree distribution

A ViewSet class for OSTree distributions.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ostreeOstreeDistributionHref
 @return DistributionsOstreeApiDistributionsOstreeOstreeReadRequest
*/
func (a *DistributionsOstreeApiService) DistributionsOstreeOstreeRead(ctx context.Context, ostreeOstreeDistributionHref string) DistributionsOstreeApiDistributionsOstreeOstreeReadRequest {
	return DistributionsOstreeApiDistributionsOstreeOstreeReadRequest{
		ApiService: a,
		ctx: ctx,
		ostreeOstreeDistributionHref: ostreeOstreeDistributionHref,
	}
}

// Execute executes the request
//  @return OstreeOstreeDistributionResponse
func (a *DistributionsOstreeApiService) DistributionsOstreeOstreeReadExecute(r DistributionsOstreeApiDistributionsOstreeOstreeReadRequest) (*OstreeOstreeDistributionResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *OstreeOstreeDistributionResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DistributionsOstreeApiService.DistributionsOstreeOstreeRead")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{ostree_ostree_distribution_href}"
	localVarPath = strings.Replace(localVarPath, "{"+"ostree_ostree_distribution_href"+"}", url.PathEscape(parameterValueToString(r.ostreeOstreeDistributionHref, "ostreeOstreeDistributionHref")), -1)
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

type DistributionsOstreeApiDistributionsOstreeOstreeUpdateRequest struct {
	ctx context.Context
	ApiService *DistributionsOstreeApiService
	ostreeOstreeDistributionHref string
	ostreeOstreeDistribution *OstreeOstreeDistribution
}

func (r DistributionsOstreeApiDistributionsOstreeOstreeUpdateRequest) OstreeOstreeDistribution(ostreeOstreeDistribution OstreeOstreeDistribution) DistributionsOstreeApiDistributionsOstreeOstreeUpdateRequest {
	r.ostreeOstreeDistribution = &ostreeOstreeDistribution
	return r
}

func (r DistributionsOstreeApiDistributionsOstreeOstreeUpdateRequest) Execute() (*AsyncOperationResponse, *http.Response, error) {
	return r.ApiService.DistributionsOstreeOstreeUpdateExecute(r)
}

/*
DistributionsOstreeOstreeUpdate Update an ostree distribution

Trigger an asynchronous update task

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ostreeOstreeDistributionHref
 @return DistributionsOstreeApiDistributionsOstreeOstreeUpdateRequest
*/
func (a *DistributionsOstreeApiService) DistributionsOstreeOstreeUpdate(ctx context.Context, ostreeOstreeDistributionHref string) DistributionsOstreeApiDistributionsOstreeOstreeUpdateRequest {
	return DistributionsOstreeApiDistributionsOstreeOstreeUpdateRequest{
		ApiService: a,
		ctx: ctx,
		ostreeOstreeDistributionHref: ostreeOstreeDistributionHref,
	}
}

// Execute executes the request
//  @return AsyncOperationResponse
func (a *DistributionsOstreeApiService) DistributionsOstreeOstreeUpdateExecute(r DistributionsOstreeApiDistributionsOstreeOstreeUpdateRequest) (*AsyncOperationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AsyncOperationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DistributionsOstreeApiService.DistributionsOstreeOstreeUpdate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{ostree_ostree_distribution_href}"
	localVarPath = strings.Replace(localVarPath, "{"+"ostree_ostree_distribution_href"+"}", url.PathEscape(parameterValueToString(r.ostreeOstreeDistributionHref, "ostreeOstreeDistributionHref")), -1)
        localVarPath = strings.Replace(localVarPath, "/%2F", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.ostreeOstreeDistribution == nil {
		return localVarReturnValue, nil, reportError("ostreeOstreeDistribution is required and must be specified")
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
	localVarPostBody = r.ostreeOstreeDistribution
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
