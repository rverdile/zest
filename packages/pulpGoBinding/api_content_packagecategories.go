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


// ContentPackagecategoriesApiService ContentPackagecategoriesApi service
type ContentPackagecategoriesApiService service

type ContentPackagecategoriesApiContentRpmPackagecategoriesListRequest struct {
	ctx context.Context
	ApiService *ContentPackagecategoriesApiService
	limit *int32
	offset *int32
	ordering *[]string
	repositoryVersion *string
	repositoryVersionAdded *string
	repositoryVersionRemoved *string
	fields *[]string
	excludeFields *[]string
}

// Number of results to return per page.
func (r ContentPackagecategoriesApiContentRpmPackagecategoriesListRequest) Limit(limit int32) ContentPackagecategoriesApiContentRpmPackagecategoriesListRequest {
	r.limit = &limit
	return r
}

// The initial index from which to return the results.
func (r ContentPackagecategoriesApiContentRpmPackagecategoriesListRequest) Offset(offset int32) ContentPackagecategoriesApiContentRpmPackagecategoriesListRequest {
	r.offset = &offset
	return r
}

// Ordering
func (r ContentPackagecategoriesApiContentRpmPackagecategoriesListRequest) Ordering(ordering []string) ContentPackagecategoriesApiContentRpmPackagecategoriesListRequest {
	r.ordering = &ordering
	return r
}

// Repository Version referenced by HREF
func (r ContentPackagecategoriesApiContentRpmPackagecategoriesListRequest) RepositoryVersion(repositoryVersion string) ContentPackagecategoriesApiContentRpmPackagecategoriesListRequest {
	r.repositoryVersion = &repositoryVersion
	return r
}

// Repository Version referenced by HREF
func (r ContentPackagecategoriesApiContentRpmPackagecategoriesListRequest) RepositoryVersionAdded(repositoryVersionAdded string) ContentPackagecategoriesApiContentRpmPackagecategoriesListRequest {
	r.repositoryVersionAdded = &repositoryVersionAdded
	return r
}

// Repository Version referenced by HREF
func (r ContentPackagecategoriesApiContentRpmPackagecategoriesListRequest) RepositoryVersionRemoved(repositoryVersionRemoved string) ContentPackagecategoriesApiContentRpmPackagecategoriesListRequest {
	r.repositoryVersionRemoved = &repositoryVersionRemoved
	return r
}

// A list of fields to include in the response.
func (r ContentPackagecategoriesApiContentRpmPackagecategoriesListRequest) Fields(fields []string) ContentPackagecategoriesApiContentRpmPackagecategoriesListRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r ContentPackagecategoriesApiContentRpmPackagecategoriesListRequest) ExcludeFields(excludeFields []string) ContentPackagecategoriesApiContentRpmPackagecategoriesListRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r ContentPackagecategoriesApiContentRpmPackagecategoriesListRequest) Execute() (*PaginatedrpmPackageCategoryResponseList, *http.Response, error) {
	return r.ApiService.ContentRpmPackagecategoriesListExecute(r)
}

/*
ContentRpmPackagecategoriesList List package categorys

PackageCategory ViewSet.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ContentPackagecategoriesApiContentRpmPackagecategoriesListRequest
*/
func (a *ContentPackagecategoriesApiService) ContentRpmPackagecategoriesList(ctx context.Context) ContentPackagecategoriesApiContentRpmPackagecategoriesListRequest {
	return ContentPackagecategoriesApiContentRpmPackagecategoriesListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PaginatedrpmPackageCategoryResponseList
func (a *ContentPackagecategoriesApiService) ContentRpmPackagecategoriesListExecute(r ContentPackagecategoriesApiContentRpmPackagecategoriesListRequest) (*PaginatedrpmPackageCategoryResponseList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedrpmPackageCategoryResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentPackagecategoriesApiService.ContentRpmPackagecategoriesList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp/api/v3/content/rpm/packagecategories/"
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

type ContentPackagecategoriesApiContentRpmPackagecategoriesReadRequest struct {
	ctx context.Context
	ApiService *ContentPackagecategoriesApiService
	rpmPackageCategoryHref string
	fields *[]string
	excludeFields *[]string
}

// A list of fields to include in the response.
func (r ContentPackagecategoriesApiContentRpmPackagecategoriesReadRequest) Fields(fields []string) ContentPackagecategoriesApiContentRpmPackagecategoriesReadRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r ContentPackagecategoriesApiContentRpmPackagecategoriesReadRequest) ExcludeFields(excludeFields []string) ContentPackagecategoriesApiContentRpmPackagecategoriesReadRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r ContentPackagecategoriesApiContentRpmPackagecategoriesReadRequest) Execute() (*RpmPackageCategoryResponse, *http.Response, error) {
	return r.ApiService.ContentRpmPackagecategoriesReadExecute(r)
}

/*
ContentRpmPackagecategoriesRead Inspect a package category

PackageCategory ViewSet.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param rpmPackageCategoryHref
 @return ContentPackagecategoriesApiContentRpmPackagecategoriesReadRequest
*/
func (a *ContentPackagecategoriesApiService) ContentRpmPackagecategoriesRead(ctx context.Context, rpmPackageCategoryHref string) ContentPackagecategoriesApiContentRpmPackagecategoriesReadRequest {
	return ContentPackagecategoriesApiContentRpmPackagecategoriesReadRequest{
		ApiService: a,
		ctx: ctx,
		rpmPackageCategoryHref: rpmPackageCategoryHref,
	}
}

// Execute executes the request
//  @return RpmPackageCategoryResponse
func (a *ContentPackagecategoriesApiService) ContentRpmPackagecategoriesReadExecute(r ContentPackagecategoriesApiContentRpmPackagecategoriesReadRequest) (*RpmPackageCategoryResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *RpmPackageCategoryResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentPackagecategoriesApiService.ContentRpmPackagecategoriesRead")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{rpm_package_category_href}"
	localVarPath = strings.Replace(localVarPath, "{"+"rpm_package_category_href"+"}", url.PathEscape(parameterValueToString(r.rpmPackageCategoryHref, "rpmPackageCategoryHref")), -1)
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
