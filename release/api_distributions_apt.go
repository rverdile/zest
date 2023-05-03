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


// DistributionsAptApiService DistributionsAptApi service
type DistributionsAptApiService service

type DistributionsAptApiDistributionsDebAptCreateRequest struct {
	ctx context.Context
	ApiService *DistributionsAptApiService
	debAptDistribution *DebAptDistribution
}

func (r DistributionsAptApiDistributionsDebAptCreateRequest) DebAptDistribution(debAptDistribution DebAptDistribution) DistributionsAptApiDistributionsDebAptCreateRequest {
	r.debAptDistribution = &debAptDistribution
	return r
}

func (r DistributionsAptApiDistributionsDebAptCreateRequest) Execute() (*AsyncOperationResponse, *http.Response, error) {
	return r.ApiService.DistributionsDebAptCreateExecute(r)
}

/*
DistributionsDebAptCreate Create an apt distribution

Trigger an asynchronous create task

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return DistributionsAptApiDistributionsDebAptCreateRequest
*/
func (a *DistributionsAptApiService) DistributionsDebAptCreate(ctx context.Context) DistributionsAptApiDistributionsDebAptCreateRequest {
	return DistributionsAptApiDistributionsDebAptCreateRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AsyncOperationResponse
func (a *DistributionsAptApiService) DistributionsDebAptCreateExecute(r DistributionsAptApiDistributionsDebAptCreateRequest) (*AsyncOperationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AsyncOperationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DistributionsAptApiService.DistributionsDebAptCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp/api/v3/distributions/deb/apt/"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.debAptDistribution == nil {
		return localVarReturnValue, nil, reportError("debAptDistribution is required and must be specified")
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
	localVarPostBody = r.debAptDistribution
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

type DistributionsAptApiDistributionsDebAptDeleteRequest struct {
	ctx context.Context
	ApiService *DistributionsAptApiService
	debAptDistributionHref string
}

func (r DistributionsAptApiDistributionsDebAptDeleteRequest) Execute() (*AsyncOperationResponse, *http.Response, error) {
	return r.ApiService.DistributionsDebAptDeleteExecute(r)
}

/*
DistributionsDebAptDelete Delete an apt distribution

Trigger an asynchronous delete task

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param debAptDistributionHref
 @return DistributionsAptApiDistributionsDebAptDeleteRequest
*/
func (a *DistributionsAptApiService) DistributionsDebAptDelete(ctx context.Context, debAptDistributionHref string) DistributionsAptApiDistributionsDebAptDeleteRequest {
	return DistributionsAptApiDistributionsDebAptDeleteRequest{
		ApiService: a,
		ctx: ctx,
		debAptDistributionHref: debAptDistributionHref,
	}
}

// Execute executes the request
//  @return AsyncOperationResponse
func (a *DistributionsAptApiService) DistributionsDebAptDeleteExecute(r DistributionsAptApiDistributionsDebAptDeleteRequest) (*AsyncOperationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AsyncOperationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DistributionsAptApiService.DistributionsDebAptDelete")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{deb_apt_distribution_href}"
	localVarPath = strings.Replace(localVarPath, "{"+"deb_apt_distribution_href"+"}", url.PathEscape(parameterValueToString(r.debAptDistributionHref, "debAptDistributionHref")), -1)
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

type DistributionsAptApiDistributionsDebAptListRequest struct {
	ctx context.Context
	ApiService *DistributionsAptApiService
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
func (r DistributionsAptApiDistributionsDebAptListRequest) BasePath(basePath string) DistributionsAptApiDistributionsDebAptListRequest {
	r.basePath = &basePath
	return r
}

// Filter results where base_path contains value
func (r DistributionsAptApiDistributionsDebAptListRequest) BasePathContains(basePathContains string) DistributionsAptApiDistributionsDebAptListRequest {
	r.basePathContains = &basePathContains
	return r
}

// Filter results where base_path contains value
func (r DistributionsAptApiDistributionsDebAptListRequest) BasePathIcontains(basePathIcontains string) DistributionsAptApiDistributionsDebAptListRequest {
	r.basePathIcontains = &basePathIcontains
	return r
}

// Filter results where base_path is in a comma-separated list of values
func (r DistributionsAptApiDistributionsDebAptListRequest) BasePathIn(basePathIn []string) DistributionsAptApiDistributionsDebAptListRequest {
	r.basePathIn = &basePathIn
	return r
}

// Number of results to return per page.
func (r DistributionsAptApiDistributionsDebAptListRequest) Limit(limit int32) DistributionsAptApiDistributionsDebAptListRequest {
	r.limit = &limit
	return r
}

// Filter results where name matches value
func (r DistributionsAptApiDistributionsDebAptListRequest) Name(name string) DistributionsAptApiDistributionsDebAptListRequest {
	r.name = &name
	return r
}

// Filter results where name contains value
func (r DistributionsAptApiDistributionsDebAptListRequest) NameContains(nameContains string) DistributionsAptApiDistributionsDebAptListRequest {
	r.nameContains = &nameContains
	return r
}

// Filter results where name contains value
func (r DistributionsAptApiDistributionsDebAptListRequest) NameIcontains(nameIcontains string) DistributionsAptApiDistributionsDebAptListRequest {
	r.nameIcontains = &nameIcontains
	return r
}

// Filter results where name is in a comma-separated list of values
func (r DistributionsAptApiDistributionsDebAptListRequest) NameIn(nameIn []string) DistributionsAptApiDistributionsDebAptListRequest {
	r.nameIn = &nameIn
	return r
}

// Filter results where name starts with value
func (r DistributionsAptApiDistributionsDebAptListRequest) NameStartswith(nameStartswith string) DistributionsAptApiDistributionsDebAptListRequest {
	r.nameStartswith = &nameStartswith
	return r
}

// The initial index from which to return the results.
func (r DistributionsAptApiDistributionsDebAptListRequest) Offset(offset int32) DistributionsAptApiDistributionsDebAptListRequest {
	r.offset = &offset
	return r
}

// Ordering  * &#x60;pulp_id&#x60; - Pulp id * &#x60;-pulp_id&#x60; - Pulp id (descending) * &#x60;pulp_created&#x60; - Pulp created * &#x60;-pulp_created&#x60; - Pulp created (descending) * &#x60;pulp_last_updated&#x60; - Pulp last updated * &#x60;-pulp_last_updated&#x60; - Pulp last updated (descending) * &#x60;pulp_type&#x60; - Pulp type * &#x60;-pulp_type&#x60; - Pulp type (descending) * &#x60;name&#x60; - Name * &#x60;-name&#x60; - Name (descending) * &#x60;pulp_labels&#x60; - Pulp labels * &#x60;-pulp_labels&#x60; - Pulp labels (descending) * &#x60;base_path&#x60; - Base path * &#x60;-base_path&#x60; - Base path (descending) * &#x60;pk&#x60; - Pk * &#x60;-pk&#x60; - Pk (descending)
func (r DistributionsAptApiDistributionsDebAptListRequest) Ordering(ordering []string) DistributionsAptApiDistributionsDebAptListRequest {
	r.ordering = &ordering
	return r
}

// Multiple values may be separated by commas.
func (r DistributionsAptApiDistributionsDebAptListRequest) PulpHrefIn(pulpHrefIn []string) DistributionsAptApiDistributionsDebAptListRequest {
	r.pulpHrefIn = &pulpHrefIn
	return r
}

// Multiple values may be separated by commas.
func (r DistributionsAptApiDistributionsDebAptListRequest) PulpIdIn(pulpIdIn []string) DistributionsAptApiDistributionsDebAptListRequest {
	r.pulpIdIn = &pulpIdIn
	return r
}

// Filter labels by search string
func (r DistributionsAptApiDistributionsDebAptListRequest) PulpLabelSelect(pulpLabelSelect string) DistributionsAptApiDistributionsDebAptListRequest {
	r.pulpLabelSelect = &pulpLabelSelect
	return r
}

// Filter results where repository matches value
func (r DistributionsAptApiDistributionsDebAptListRequest) Repository(repository string) DistributionsAptApiDistributionsDebAptListRequest {
	r.repository = &repository
	return r
}

// Filter results where repository is in a comma-separated list of values
func (r DistributionsAptApiDistributionsDebAptListRequest) RepositoryIn(repositoryIn []string) DistributionsAptApiDistributionsDebAptListRequest {
	r.repositoryIn = &repositoryIn
	return r
}

// Filter distributions based on the content served by them
func (r DistributionsAptApiDistributionsDebAptListRequest) WithContent(withContent string) DistributionsAptApiDistributionsDebAptListRequest {
	r.withContent = &withContent
	return r
}

// A list of fields to include in the response.
func (r DistributionsAptApiDistributionsDebAptListRequest) Fields(fields []string) DistributionsAptApiDistributionsDebAptListRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r DistributionsAptApiDistributionsDebAptListRequest) ExcludeFields(excludeFields []string) DistributionsAptApiDistributionsDebAptListRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r DistributionsAptApiDistributionsDebAptListRequest) Execute() (*PaginateddebAptDistributionResponseList, *http.Response, error) {
	return r.ApiService.DistributionsDebAptListExecute(r)
}

/*
DistributionsDebAptList List apt distributions

An AptDistribution is just an AptPublication made available via the content app.

Creating an AptDistribution is a comparatively quick action. This way Pulp users may take as
much time as is needed to prepare a VerbatimPublication or AptPublication, and then control the
exact moment when that publication is made available.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return DistributionsAptApiDistributionsDebAptListRequest
*/
func (a *DistributionsAptApiService) DistributionsDebAptList(ctx context.Context) DistributionsAptApiDistributionsDebAptListRequest {
	return DistributionsAptApiDistributionsDebAptListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PaginateddebAptDistributionResponseList
func (a *DistributionsAptApiService) DistributionsDebAptListExecute(r DistributionsAptApiDistributionsDebAptListRequest) (*PaginateddebAptDistributionResponseList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginateddebAptDistributionResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DistributionsAptApiService.DistributionsDebAptList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp/api/v3/distributions/deb/apt/"
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

type DistributionsAptApiDistributionsDebAptPartialUpdateRequest struct {
	ctx context.Context
	ApiService *DistributionsAptApiService
	debAptDistributionHref string
	patcheddebAptDistribution *PatcheddebAptDistribution
}

func (r DistributionsAptApiDistributionsDebAptPartialUpdateRequest) PatcheddebAptDistribution(patcheddebAptDistribution PatcheddebAptDistribution) DistributionsAptApiDistributionsDebAptPartialUpdateRequest {
	r.patcheddebAptDistribution = &patcheddebAptDistribution
	return r
}

func (r DistributionsAptApiDistributionsDebAptPartialUpdateRequest) Execute() (*AsyncOperationResponse, *http.Response, error) {
	return r.ApiService.DistributionsDebAptPartialUpdateExecute(r)
}

/*
DistributionsDebAptPartialUpdate Update an apt distribution

Trigger an asynchronous partial update task

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param debAptDistributionHref
 @return DistributionsAptApiDistributionsDebAptPartialUpdateRequest
*/
func (a *DistributionsAptApiService) DistributionsDebAptPartialUpdate(ctx context.Context, debAptDistributionHref string) DistributionsAptApiDistributionsDebAptPartialUpdateRequest {
	return DistributionsAptApiDistributionsDebAptPartialUpdateRequest{
		ApiService: a,
		ctx: ctx,
		debAptDistributionHref: debAptDistributionHref,
	}
}

// Execute executes the request
//  @return AsyncOperationResponse
func (a *DistributionsAptApiService) DistributionsDebAptPartialUpdateExecute(r DistributionsAptApiDistributionsDebAptPartialUpdateRequest) (*AsyncOperationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AsyncOperationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DistributionsAptApiService.DistributionsDebAptPartialUpdate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{deb_apt_distribution_href}"
	localVarPath = strings.Replace(localVarPath, "{"+"deb_apt_distribution_href"+"}", url.PathEscape(parameterValueToString(r.debAptDistributionHref, "debAptDistributionHref")), -1)
        localVarPath = strings.Replace(localVarPath, "/%2F", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.patcheddebAptDistribution == nil {
		return localVarReturnValue, nil, reportError("patcheddebAptDistribution is required and must be specified")
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
	localVarPostBody = r.patcheddebAptDistribution
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

type DistributionsAptApiDistributionsDebAptReadRequest struct {
	ctx context.Context
	ApiService *DistributionsAptApiService
	debAptDistributionHref string
	fields *[]string
	excludeFields *[]string
}

// A list of fields to include in the response.
func (r DistributionsAptApiDistributionsDebAptReadRequest) Fields(fields []string) DistributionsAptApiDistributionsDebAptReadRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r DistributionsAptApiDistributionsDebAptReadRequest) ExcludeFields(excludeFields []string) DistributionsAptApiDistributionsDebAptReadRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r DistributionsAptApiDistributionsDebAptReadRequest) Execute() (*DebAptDistributionResponse, *http.Response, error) {
	return r.ApiService.DistributionsDebAptReadExecute(r)
}

/*
DistributionsDebAptRead Inspect an apt distribution

An AptDistribution is just an AptPublication made available via the content app.

Creating an AptDistribution is a comparatively quick action. This way Pulp users may take as
much time as is needed to prepare a VerbatimPublication or AptPublication, and then control the
exact moment when that publication is made available.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param debAptDistributionHref
 @return DistributionsAptApiDistributionsDebAptReadRequest
*/
func (a *DistributionsAptApiService) DistributionsDebAptRead(ctx context.Context, debAptDistributionHref string) DistributionsAptApiDistributionsDebAptReadRequest {
	return DistributionsAptApiDistributionsDebAptReadRequest{
		ApiService: a,
		ctx: ctx,
		debAptDistributionHref: debAptDistributionHref,
	}
}

// Execute executes the request
//  @return DebAptDistributionResponse
func (a *DistributionsAptApiService) DistributionsDebAptReadExecute(r DistributionsAptApiDistributionsDebAptReadRequest) (*DebAptDistributionResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DebAptDistributionResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DistributionsAptApiService.DistributionsDebAptRead")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{deb_apt_distribution_href}"
	localVarPath = strings.Replace(localVarPath, "{"+"deb_apt_distribution_href"+"}", url.PathEscape(parameterValueToString(r.debAptDistributionHref, "debAptDistributionHref")), -1)
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

type DistributionsAptApiDistributionsDebAptUpdateRequest struct {
	ctx context.Context
	ApiService *DistributionsAptApiService
	debAptDistributionHref string
	debAptDistribution *DebAptDistribution
}

func (r DistributionsAptApiDistributionsDebAptUpdateRequest) DebAptDistribution(debAptDistribution DebAptDistribution) DistributionsAptApiDistributionsDebAptUpdateRequest {
	r.debAptDistribution = &debAptDistribution
	return r
}

func (r DistributionsAptApiDistributionsDebAptUpdateRequest) Execute() (*AsyncOperationResponse, *http.Response, error) {
	return r.ApiService.DistributionsDebAptUpdateExecute(r)
}

/*
DistributionsDebAptUpdate Update an apt distribution

Trigger an asynchronous update task

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param debAptDistributionHref
 @return DistributionsAptApiDistributionsDebAptUpdateRequest
*/
func (a *DistributionsAptApiService) DistributionsDebAptUpdate(ctx context.Context, debAptDistributionHref string) DistributionsAptApiDistributionsDebAptUpdateRequest {
	return DistributionsAptApiDistributionsDebAptUpdateRequest{
		ApiService: a,
		ctx: ctx,
		debAptDistributionHref: debAptDistributionHref,
	}
}

// Execute executes the request
//  @return AsyncOperationResponse
func (a *DistributionsAptApiService) DistributionsDebAptUpdateExecute(r DistributionsAptApiDistributionsDebAptUpdateRequest) (*AsyncOperationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AsyncOperationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DistributionsAptApiService.DistributionsDebAptUpdate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{deb_apt_distribution_href}"
	localVarPath = strings.Replace(localVarPath, "{"+"deb_apt_distribution_href"+"}", url.PathEscape(parameterValueToString(r.debAptDistributionHref, "debAptDistributionHref")), -1)
        localVarPath = strings.Replace(localVarPath, "/%2F", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.debAptDistribution == nil {
		return localVarReturnValue, nil, reportError("debAptDistribution is required and must be specified")
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
	localVarPostBody = r.debAptDistribution
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
