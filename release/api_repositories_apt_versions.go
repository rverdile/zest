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
	"time"
	"reflect"
)


// RepositoriesAptVersionsApiService RepositoriesAptVersionsApi service
type RepositoriesAptVersionsApiService service

type RepositoriesAptVersionsApiRepositoriesDebAptVersionsDeleteRequest struct {
	ctx context.Context
	ApiService *RepositoriesAptVersionsApiService
	debAptRepositoryVersionHref string
}

func (r RepositoriesAptVersionsApiRepositoriesDebAptVersionsDeleteRequest) Execute() (*AsyncOperationResponse, *http.Response, error) {
	return r.ApiService.RepositoriesDebAptVersionsDeleteExecute(r)
}

/*
RepositoriesDebAptVersionsDelete Delete a repository version

Trigger an asynchronous task to delete a repository version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param debAptRepositoryVersionHref
 @return RepositoriesAptVersionsApiRepositoriesDebAptVersionsDeleteRequest
*/
func (a *RepositoriesAptVersionsApiService) RepositoriesDebAptVersionsDelete(ctx context.Context, debAptRepositoryVersionHref string) RepositoriesAptVersionsApiRepositoriesDebAptVersionsDeleteRequest {
	return RepositoriesAptVersionsApiRepositoriesDebAptVersionsDeleteRequest{
		ApiService: a,
		ctx: ctx,
		debAptRepositoryVersionHref: debAptRepositoryVersionHref,
	}
}

// Execute executes the request
//  @return AsyncOperationResponse
func (a *RepositoriesAptVersionsApiService) RepositoriesDebAptVersionsDeleteExecute(r RepositoriesAptVersionsApiRepositoriesDebAptVersionsDeleteRequest) (*AsyncOperationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AsyncOperationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RepositoriesAptVersionsApiService.RepositoriesDebAptVersionsDelete")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{deb_apt_repository_version_href}"
	localVarPath = strings.Replace(localVarPath, "{"+"deb_apt_repository_version_href"+"}", url.PathEscape(parameterValueToString(r.debAptRepositoryVersionHref, "debAptRepositoryVersionHref")), -1)
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

type RepositoriesAptVersionsApiRepositoriesDebAptVersionsListRequest struct {
	ctx context.Context
	ApiService *RepositoriesAptVersionsApiService
	debAptRepositoryHref string
	content *string
	contentIn *string
	limit *int32
	number *int32
	numberGt *int32
	numberGte *int32
	numberLt *int32
	numberLte *int32
	numberRange *[]int32
	offset *int32
	ordering *[]string
	pulpCreated *time.Time
	pulpCreatedGt *time.Time
	pulpCreatedGte *time.Time
	pulpCreatedLt *time.Time
	pulpCreatedLte *time.Time
	pulpCreatedRange *[]time.Time
	pulpHrefIn *[]string
	fields *[]string
	excludeFields *[]string
}

// Content Unit referenced by HREF
func (r RepositoriesAptVersionsApiRepositoriesDebAptVersionsListRequest) Content(content string) RepositoriesAptVersionsApiRepositoriesDebAptVersionsListRequest {
	r.content = &content
	return r
}

// Content Unit referenced by HREF
func (r RepositoriesAptVersionsApiRepositoriesDebAptVersionsListRequest) ContentIn(contentIn string) RepositoriesAptVersionsApiRepositoriesDebAptVersionsListRequest {
	r.contentIn = &contentIn
	return r
}

// Number of results to return per page.
func (r RepositoriesAptVersionsApiRepositoriesDebAptVersionsListRequest) Limit(limit int32) RepositoriesAptVersionsApiRepositoriesDebAptVersionsListRequest {
	r.limit = &limit
	return r
}

// Filter results where number matches value
func (r RepositoriesAptVersionsApiRepositoriesDebAptVersionsListRequest) Number(number int32) RepositoriesAptVersionsApiRepositoriesDebAptVersionsListRequest {
	r.number = &number
	return r
}

// Filter results where number is greater than value
func (r RepositoriesAptVersionsApiRepositoriesDebAptVersionsListRequest) NumberGt(numberGt int32) RepositoriesAptVersionsApiRepositoriesDebAptVersionsListRequest {
	r.numberGt = &numberGt
	return r
}

// Filter results where number is greater than or equal to value
func (r RepositoriesAptVersionsApiRepositoriesDebAptVersionsListRequest) NumberGte(numberGte int32) RepositoriesAptVersionsApiRepositoriesDebAptVersionsListRequest {
	r.numberGte = &numberGte
	return r
}

// Filter results where number is less than value
func (r RepositoriesAptVersionsApiRepositoriesDebAptVersionsListRequest) NumberLt(numberLt int32) RepositoriesAptVersionsApiRepositoriesDebAptVersionsListRequest {
	r.numberLt = &numberLt
	return r
}

// Filter results where number is less than or equal to value
func (r RepositoriesAptVersionsApiRepositoriesDebAptVersionsListRequest) NumberLte(numberLte int32) RepositoriesAptVersionsApiRepositoriesDebAptVersionsListRequest {
	r.numberLte = &numberLte
	return r
}

// Filter results where number is between two comma separated values
func (r RepositoriesAptVersionsApiRepositoriesDebAptVersionsListRequest) NumberRange(numberRange []int32) RepositoriesAptVersionsApiRepositoriesDebAptVersionsListRequest {
	r.numberRange = &numberRange
	return r
}

// The initial index from which to return the results.
func (r RepositoriesAptVersionsApiRepositoriesDebAptVersionsListRequest) Offset(offset int32) RepositoriesAptVersionsApiRepositoriesDebAptVersionsListRequest {
	r.offset = &offset
	return r
}

// Ordering  * &#x60;pulp_id&#x60; - Pulp id * &#x60;-pulp_id&#x60; - Pulp id (descending) * &#x60;pulp_created&#x60; - Pulp created * &#x60;-pulp_created&#x60; - Pulp created (descending) * &#x60;pulp_last_updated&#x60; - Pulp last updated * &#x60;-pulp_last_updated&#x60; - Pulp last updated (descending) * &#x60;number&#x60; - Number * &#x60;-number&#x60; - Number (descending) * &#x60;complete&#x60; - Complete * &#x60;-complete&#x60; - Complete (descending) * &#x60;info&#x60; - Info * &#x60;-info&#x60; - Info (descending) * &#x60;pk&#x60; - Pk * &#x60;-pk&#x60; - Pk (descending)
func (r RepositoriesAptVersionsApiRepositoriesDebAptVersionsListRequest) Ordering(ordering []string) RepositoriesAptVersionsApiRepositoriesDebAptVersionsListRequest {
	r.ordering = &ordering
	return r
}

// Filter results where pulp_created matches value
func (r RepositoriesAptVersionsApiRepositoriesDebAptVersionsListRequest) PulpCreated(pulpCreated time.Time) RepositoriesAptVersionsApiRepositoriesDebAptVersionsListRequest {
	r.pulpCreated = &pulpCreated
	return r
}

// Filter results where pulp_created is greater than value
func (r RepositoriesAptVersionsApiRepositoriesDebAptVersionsListRequest) PulpCreatedGt(pulpCreatedGt time.Time) RepositoriesAptVersionsApiRepositoriesDebAptVersionsListRequest {
	r.pulpCreatedGt = &pulpCreatedGt
	return r
}

// Filter results where pulp_created is greater than or equal to value
func (r RepositoriesAptVersionsApiRepositoriesDebAptVersionsListRequest) PulpCreatedGte(pulpCreatedGte time.Time) RepositoriesAptVersionsApiRepositoriesDebAptVersionsListRequest {
	r.pulpCreatedGte = &pulpCreatedGte
	return r
}

// Filter results where pulp_created is less than value
func (r RepositoriesAptVersionsApiRepositoriesDebAptVersionsListRequest) PulpCreatedLt(pulpCreatedLt time.Time) RepositoriesAptVersionsApiRepositoriesDebAptVersionsListRequest {
	r.pulpCreatedLt = &pulpCreatedLt
	return r
}

// Filter results where pulp_created is less than or equal to value
func (r RepositoriesAptVersionsApiRepositoriesDebAptVersionsListRequest) PulpCreatedLte(pulpCreatedLte time.Time) RepositoriesAptVersionsApiRepositoriesDebAptVersionsListRequest {
	r.pulpCreatedLte = &pulpCreatedLte
	return r
}

// Filter results where pulp_created is between two comma separated values
func (r RepositoriesAptVersionsApiRepositoriesDebAptVersionsListRequest) PulpCreatedRange(pulpCreatedRange []time.Time) RepositoriesAptVersionsApiRepositoriesDebAptVersionsListRequest {
	r.pulpCreatedRange = &pulpCreatedRange
	return r
}

// Multiple values may be separated by commas.
func (r RepositoriesAptVersionsApiRepositoriesDebAptVersionsListRequest) PulpHrefIn(pulpHrefIn []string) RepositoriesAptVersionsApiRepositoriesDebAptVersionsListRequest {
	r.pulpHrefIn = &pulpHrefIn
	return r
}

// A list of fields to include in the response.
func (r RepositoriesAptVersionsApiRepositoriesDebAptVersionsListRequest) Fields(fields []string) RepositoriesAptVersionsApiRepositoriesDebAptVersionsListRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r RepositoriesAptVersionsApiRepositoriesDebAptVersionsListRequest) ExcludeFields(excludeFields []string) RepositoriesAptVersionsApiRepositoriesDebAptVersionsListRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r RepositoriesAptVersionsApiRepositoriesDebAptVersionsListRequest) Execute() (*PaginatedRepositoryVersionResponseList, *http.Response, error) {
	return r.ApiService.RepositoriesDebAptVersionsListExecute(r)
}

/*
RepositoriesDebAptVersionsList List repository versions

An AptRepositoryVersion represents a single APT repository version as stored by Pulp.

It may be used as the basis for the creation of Pulp distributions in order to actually serve
the content contained within the repository version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param debAptRepositoryHref
 @return RepositoriesAptVersionsApiRepositoriesDebAptVersionsListRequest
*/
func (a *RepositoriesAptVersionsApiService) RepositoriesDebAptVersionsList(ctx context.Context, debAptRepositoryHref string) RepositoriesAptVersionsApiRepositoriesDebAptVersionsListRequest {
	return RepositoriesAptVersionsApiRepositoriesDebAptVersionsListRequest{
		ApiService: a,
		ctx: ctx,
		debAptRepositoryHref: debAptRepositoryHref,
	}
}

// Execute executes the request
//  @return PaginatedRepositoryVersionResponseList
func (a *RepositoriesAptVersionsApiService) RepositoriesDebAptVersionsListExecute(r RepositoriesAptVersionsApiRepositoriesDebAptVersionsListRequest) (*PaginatedRepositoryVersionResponseList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedRepositoryVersionResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RepositoriesAptVersionsApiService.RepositoriesDebAptVersionsList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{deb_apt_repository_href}versions/"
	localVarPath = strings.Replace(localVarPath, "{"+"deb_apt_repository_href"+"}", url.PathEscape(parameterValueToString(r.debAptRepositoryHref, "debAptRepositoryHref")), -1)
        localVarPath = strings.Replace(localVarPath, "/%2F", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.content != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "content", r.content, "")
	}
	if r.contentIn != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "content__in", r.contentIn, "")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	}
	if r.number != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "number", r.number, "")
	}
	if r.numberGt != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "number__gt", r.numberGt, "")
	}
	if r.numberGte != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "number__gte", r.numberGte, "")
	}
	if r.numberLt != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "number__lt", r.numberLt, "")
	}
	if r.numberLte != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "number__lte", r.numberLte, "")
	}
	if r.numberRange != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "number__range", r.numberRange, "csv")
	}
	if r.offset != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "offset", r.offset, "")
	}
	if r.ordering != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "ordering", r.ordering, "csv")
	}
	if r.pulpCreated != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pulp_created", r.pulpCreated, "")
	}
	if r.pulpCreatedGt != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pulp_created__gt", r.pulpCreatedGt, "")
	}
	if r.pulpCreatedGte != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pulp_created__gte", r.pulpCreatedGte, "")
	}
	if r.pulpCreatedLt != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pulp_created__lt", r.pulpCreatedLt, "")
	}
	if r.pulpCreatedLte != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pulp_created__lte", r.pulpCreatedLte, "")
	}
	if r.pulpCreatedRange != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pulp_created__range", r.pulpCreatedRange, "csv")
	}
	if r.pulpHrefIn != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pulp_href__in", r.pulpHrefIn, "csv")
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

type RepositoriesAptVersionsApiRepositoriesDebAptVersionsReadRequest struct {
	ctx context.Context
	ApiService *RepositoriesAptVersionsApiService
	debAptRepositoryVersionHref string
	fields *[]string
	excludeFields *[]string
}

// A list of fields to include in the response.
func (r RepositoriesAptVersionsApiRepositoriesDebAptVersionsReadRequest) Fields(fields []string) RepositoriesAptVersionsApiRepositoriesDebAptVersionsReadRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r RepositoriesAptVersionsApiRepositoriesDebAptVersionsReadRequest) ExcludeFields(excludeFields []string) RepositoriesAptVersionsApiRepositoriesDebAptVersionsReadRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r RepositoriesAptVersionsApiRepositoriesDebAptVersionsReadRequest) Execute() (*RepositoryVersionResponse, *http.Response, error) {
	return r.ApiService.RepositoriesDebAptVersionsReadExecute(r)
}

/*
RepositoriesDebAptVersionsRead Inspect a repository version

An AptRepositoryVersion represents a single APT repository version as stored by Pulp.

It may be used as the basis for the creation of Pulp distributions in order to actually serve
the content contained within the repository version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param debAptRepositoryVersionHref
 @return RepositoriesAptVersionsApiRepositoriesDebAptVersionsReadRequest
*/
func (a *RepositoriesAptVersionsApiService) RepositoriesDebAptVersionsRead(ctx context.Context, debAptRepositoryVersionHref string) RepositoriesAptVersionsApiRepositoriesDebAptVersionsReadRequest {
	return RepositoriesAptVersionsApiRepositoriesDebAptVersionsReadRequest{
		ApiService: a,
		ctx: ctx,
		debAptRepositoryVersionHref: debAptRepositoryVersionHref,
	}
}

// Execute executes the request
//  @return RepositoryVersionResponse
func (a *RepositoriesAptVersionsApiService) RepositoriesDebAptVersionsReadExecute(r RepositoriesAptVersionsApiRepositoriesDebAptVersionsReadRequest) (*RepositoryVersionResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *RepositoryVersionResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RepositoriesAptVersionsApiService.RepositoriesDebAptVersionsRead")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{deb_apt_repository_version_href}"
	localVarPath = strings.Replace(localVarPath, "{"+"deb_apt_repository_version_href"+"}", url.PathEscape(parameterValueToString(r.debAptRepositoryVersionHref, "debAptRepositoryVersionHref")), -1)
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

type RepositoriesAptVersionsApiRepositoriesDebAptVersionsRepairRequest struct {
	ctx context.Context
	ApiService *RepositoriesAptVersionsApiService
	debAptRepositoryVersionHref string
	repair *Repair
}

func (r RepositoriesAptVersionsApiRepositoriesDebAptVersionsRepairRequest) Repair(repair Repair) RepositoriesAptVersionsApiRepositoriesDebAptVersionsRepairRequest {
	r.repair = &repair
	return r
}

func (r RepositoriesAptVersionsApiRepositoriesDebAptVersionsRepairRequest) Execute() (*AsyncOperationResponse, *http.Response, error) {
	return r.ApiService.RepositoriesDebAptVersionsRepairExecute(r)
}

/*
RepositoriesDebAptVersionsRepair Method for RepositoriesDebAptVersionsRepair

Trigger an asynchronous task to repair a repository version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param debAptRepositoryVersionHref
 @return RepositoriesAptVersionsApiRepositoriesDebAptVersionsRepairRequest
*/
func (a *RepositoriesAptVersionsApiService) RepositoriesDebAptVersionsRepair(ctx context.Context, debAptRepositoryVersionHref string) RepositoriesAptVersionsApiRepositoriesDebAptVersionsRepairRequest {
	return RepositoriesAptVersionsApiRepositoriesDebAptVersionsRepairRequest{
		ApiService: a,
		ctx: ctx,
		debAptRepositoryVersionHref: debAptRepositoryVersionHref,
	}
}

// Execute executes the request
//  @return AsyncOperationResponse
func (a *RepositoriesAptVersionsApiService) RepositoriesDebAptVersionsRepairExecute(r RepositoriesAptVersionsApiRepositoriesDebAptVersionsRepairRequest) (*AsyncOperationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AsyncOperationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RepositoriesAptVersionsApiService.RepositoriesDebAptVersionsRepair")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{deb_apt_repository_version_href}repair/"
	localVarPath = strings.Replace(localVarPath, "{"+"deb_apt_repository_version_href"+"}", url.PathEscape(parameterValueToString(r.debAptRepositoryVersionHref, "debAptRepositoryVersionHref")), -1)
        localVarPath = strings.Replace(localVarPath, "/%2F", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.repair == nil {
		return localVarReturnValue, nil, reportError("repair is required and must be specified")
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
	localVarPostBody = r.repair
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
