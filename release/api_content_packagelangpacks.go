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


// ContentPackagelangpacksApiService ContentPackagelangpacksApi service
type ContentPackagelangpacksApiService service

type ContentPackagelangpacksApiContentRpmPackagelangpacksListRequest struct {
	ctx context.Context
	ApiService *ContentPackagelangpacksApiService
	limit *int32
	offset *int32
	ordering *[]string
	pulpHrefIn *[]string
	pulpIdIn *[]string
	repositoryVersion *string
	repositoryVersionAdded *string
	repositoryVersionRemoved *string
	fields *[]string
	excludeFields *[]string
}

// Number of results to return per page.
func (r ContentPackagelangpacksApiContentRpmPackagelangpacksListRequest) Limit(limit int32) ContentPackagelangpacksApiContentRpmPackagelangpacksListRequest {
	r.limit = &limit
	return r
}

// The initial index from which to return the results.
func (r ContentPackagelangpacksApiContentRpmPackagelangpacksListRequest) Offset(offset int32) ContentPackagelangpacksApiContentRpmPackagelangpacksListRequest {
	r.offset = &offset
	return r
}

// Ordering  * &#x60;pk&#x60; - Pk * &#x60;-pk&#x60; - Pk (descending)
func (r ContentPackagelangpacksApiContentRpmPackagelangpacksListRequest) Ordering(ordering []string) ContentPackagelangpacksApiContentRpmPackagelangpacksListRequest {
	r.ordering = &ordering
	return r
}

// Multiple values may be separated by commas.
func (r ContentPackagelangpacksApiContentRpmPackagelangpacksListRequest) PulpHrefIn(pulpHrefIn []string) ContentPackagelangpacksApiContentRpmPackagelangpacksListRequest {
	r.pulpHrefIn = &pulpHrefIn
	return r
}

// Multiple values may be separated by commas.
func (r ContentPackagelangpacksApiContentRpmPackagelangpacksListRequest) PulpIdIn(pulpIdIn []string) ContentPackagelangpacksApiContentRpmPackagelangpacksListRequest {
	r.pulpIdIn = &pulpIdIn
	return r
}

// Repository Version referenced by HREF
func (r ContentPackagelangpacksApiContentRpmPackagelangpacksListRequest) RepositoryVersion(repositoryVersion string) ContentPackagelangpacksApiContentRpmPackagelangpacksListRequest {
	r.repositoryVersion = &repositoryVersion
	return r
}

// Repository Version referenced by HREF
func (r ContentPackagelangpacksApiContentRpmPackagelangpacksListRequest) RepositoryVersionAdded(repositoryVersionAdded string) ContentPackagelangpacksApiContentRpmPackagelangpacksListRequest {
	r.repositoryVersionAdded = &repositoryVersionAdded
	return r
}

// Repository Version referenced by HREF
func (r ContentPackagelangpacksApiContentRpmPackagelangpacksListRequest) RepositoryVersionRemoved(repositoryVersionRemoved string) ContentPackagelangpacksApiContentRpmPackagelangpacksListRequest {
	r.repositoryVersionRemoved = &repositoryVersionRemoved
	return r
}

// A list of fields to include in the response.
func (r ContentPackagelangpacksApiContentRpmPackagelangpacksListRequest) Fields(fields []string) ContentPackagelangpacksApiContentRpmPackagelangpacksListRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r ContentPackagelangpacksApiContentRpmPackagelangpacksListRequest) ExcludeFields(excludeFields []string) ContentPackagelangpacksApiContentRpmPackagelangpacksListRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r ContentPackagelangpacksApiContentRpmPackagelangpacksListRequest) Execute() (*PaginatedrpmPackageLangpacksResponseList, *http.Response, error) {
	return r.ApiService.ContentRpmPackagelangpacksListExecute(r)
}

/*
ContentRpmPackagelangpacksList List package langpackss

PackageLangpacks ViewSet.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ContentPackagelangpacksApiContentRpmPackagelangpacksListRequest
*/
func (a *ContentPackagelangpacksApiService) ContentRpmPackagelangpacksList(ctx context.Context) ContentPackagelangpacksApiContentRpmPackagelangpacksListRequest {
	return ContentPackagelangpacksApiContentRpmPackagelangpacksListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PaginatedrpmPackageLangpacksResponseList
func (a *ContentPackagelangpacksApiService) ContentRpmPackagelangpacksListExecute(r ContentPackagelangpacksApiContentRpmPackagelangpacksListRequest) (*PaginatedrpmPackageLangpacksResponseList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedrpmPackageLangpacksResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentPackagelangpacksApiService.ContentRpmPackagelangpacksList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp/api/v3/content/rpm/packagelangpacks/"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
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
	if r.repositoryVersion != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "repository_version", r.repositoryVersion, "")
	}
	if r.repositoryVersionAdded != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "repository_version_added", r.repositoryVersionAdded, "")
	}
	if r.repositoryVersionRemoved != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "repository_version_removed", r.repositoryVersionRemoved, "")
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

type ContentPackagelangpacksApiContentRpmPackagelangpacksReadRequest struct {
	ctx context.Context
	ApiService *ContentPackagelangpacksApiService
	rpmPackageLangpacksHref string
	fields *[]string
	excludeFields *[]string
}

// A list of fields to include in the response.
func (r ContentPackagelangpacksApiContentRpmPackagelangpacksReadRequest) Fields(fields []string) ContentPackagelangpacksApiContentRpmPackagelangpacksReadRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r ContentPackagelangpacksApiContentRpmPackagelangpacksReadRequest) ExcludeFields(excludeFields []string) ContentPackagelangpacksApiContentRpmPackagelangpacksReadRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r ContentPackagelangpacksApiContentRpmPackagelangpacksReadRequest) Execute() (*RpmPackageLangpacksResponse, *http.Response, error) {
	return r.ApiService.ContentRpmPackagelangpacksReadExecute(r)
}

/*
ContentRpmPackagelangpacksRead Inspect a package langpacks

PackageLangpacks ViewSet.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param rpmPackageLangpacksHref
 @return ContentPackagelangpacksApiContentRpmPackagelangpacksReadRequest
*/
func (a *ContentPackagelangpacksApiService) ContentRpmPackagelangpacksRead(ctx context.Context, rpmPackageLangpacksHref string) ContentPackagelangpacksApiContentRpmPackagelangpacksReadRequest {
	return ContentPackagelangpacksApiContentRpmPackagelangpacksReadRequest{
		ApiService: a,
		ctx: ctx,
		rpmPackageLangpacksHref: rpmPackageLangpacksHref,
	}
}

// Execute executes the request
//  @return RpmPackageLangpacksResponse
func (a *ContentPackagelangpacksApiService) ContentRpmPackagelangpacksReadExecute(r ContentPackagelangpacksApiContentRpmPackagelangpacksReadRequest) (*RpmPackageLangpacksResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *RpmPackageLangpacksResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentPackagelangpacksApiService.ContentRpmPackagelangpacksRead")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{rpm_package_langpacks_href}"
	localVarPath = strings.Replace(localVarPath, "{"+"rpm_package_langpacks_href"+"}", url.PathEscape(parameterValueToString(r.rpmPackageLangpacksHref, "rpmPackageLangpacksHref")), -1)
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
