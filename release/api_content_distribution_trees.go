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


// ContentDistributionTreesAPIService ContentDistributionTreesAPI service
type ContentDistributionTreesAPIService service

type ContentDistributionTreesAPIContentRpmDistributionTreesListRequest struct {
	ctx context.Context
	ApiService *ContentDistributionTreesAPIService
	pulpDomain string
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
func (r ContentDistributionTreesAPIContentRpmDistributionTreesListRequest) Limit(limit int32) ContentDistributionTreesAPIContentRpmDistributionTreesListRequest {
	r.limit = &limit
	return r
}

// The initial index from which to return the results.
func (r ContentDistributionTreesAPIContentRpmDistributionTreesListRequest) Offset(offset int32) ContentDistributionTreesAPIContentRpmDistributionTreesListRequest {
	r.offset = &offset
	return r
}

// Ordering  * &#x60;pk&#x60; - Pk * &#x60;-pk&#x60; - Pk (descending)
func (r ContentDistributionTreesAPIContentRpmDistributionTreesListRequest) Ordering(ordering []string) ContentDistributionTreesAPIContentRpmDistributionTreesListRequest {
	r.ordering = &ordering
	return r
}

// Multiple values may be separated by commas.
func (r ContentDistributionTreesAPIContentRpmDistributionTreesListRequest) PulpHrefIn(pulpHrefIn []string) ContentDistributionTreesAPIContentRpmDistributionTreesListRequest {
	r.pulpHrefIn = &pulpHrefIn
	return r
}

// Multiple values may be separated by commas.
func (r ContentDistributionTreesAPIContentRpmDistributionTreesListRequest) PulpIdIn(pulpIdIn []string) ContentDistributionTreesAPIContentRpmDistributionTreesListRequest {
	r.pulpIdIn = &pulpIdIn
	return r
}

// Repository Version referenced by HREF
func (r ContentDistributionTreesAPIContentRpmDistributionTreesListRequest) RepositoryVersion(repositoryVersion string) ContentDistributionTreesAPIContentRpmDistributionTreesListRequest {
	r.repositoryVersion = &repositoryVersion
	return r
}

// Repository Version referenced by HREF
func (r ContentDistributionTreesAPIContentRpmDistributionTreesListRequest) RepositoryVersionAdded(repositoryVersionAdded string) ContentDistributionTreesAPIContentRpmDistributionTreesListRequest {
	r.repositoryVersionAdded = &repositoryVersionAdded
	return r
}

// Repository Version referenced by HREF
func (r ContentDistributionTreesAPIContentRpmDistributionTreesListRequest) RepositoryVersionRemoved(repositoryVersionRemoved string) ContentDistributionTreesAPIContentRpmDistributionTreesListRequest {
	r.repositoryVersionRemoved = &repositoryVersionRemoved
	return r
}

// A list of fields to include in the response.
func (r ContentDistributionTreesAPIContentRpmDistributionTreesListRequest) Fields(fields []string) ContentDistributionTreesAPIContentRpmDistributionTreesListRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r ContentDistributionTreesAPIContentRpmDistributionTreesListRequest) ExcludeFields(excludeFields []string) ContentDistributionTreesAPIContentRpmDistributionTreesListRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r ContentDistributionTreesAPIContentRpmDistributionTreesListRequest) Execute() (*PaginatedrpmDistributionTreeResponseList, *http.Response, error) {
	return r.ApiService.ContentRpmDistributionTreesListExecute(r)
}

/*
ContentRpmDistributionTreesList List distribution trees

Distribution Tree Viewset.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pulpDomain
 @return ContentDistributionTreesAPIContentRpmDistributionTreesListRequest
*/
func (a *ContentDistributionTreesAPIService) ContentRpmDistributionTreesList(ctx context.Context, pulpDomain string) ContentDistributionTreesAPIContentRpmDistributionTreesListRequest {
	return ContentDistributionTreesAPIContentRpmDistributionTreesListRequest{
		ApiService: a,
		ctx: ctx,
		pulpDomain: pulpDomain,
	}
}

// Execute executes the request
//  @return PaginatedrpmDistributionTreeResponseList
func (a *ContentDistributionTreesAPIService) ContentRpmDistributionTreesListExecute(r ContentDistributionTreesAPIContentRpmDistributionTreesListRequest) (*PaginatedrpmDistributionTreeResponseList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedrpmDistributionTreeResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentDistributionTreesAPIService.ContentRpmDistributionTreesList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp/{pulp_domain}/api/v3/content/rpm/distribution_trees/"
	localVarPath = strings.Replace(localVarPath, "{"+"pulp_domain"+"}", url.PathEscape(parameterValueToString(r.pulpDomain, "pulpDomain")), -1)
        localVarPath = strings.Replace(localVarPath, "/%2F", "/", -1)

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

type ContentDistributionTreesAPIContentRpmDistributionTreesReadRequest struct {
	ctx context.Context
	ApiService *ContentDistributionTreesAPIService
	rpmDistributionTreeHref string
	fields *[]string
	excludeFields *[]string
}

// A list of fields to include in the response.
func (r ContentDistributionTreesAPIContentRpmDistributionTreesReadRequest) Fields(fields []string) ContentDistributionTreesAPIContentRpmDistributionTreesReadRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r ContentDistributionTreesAPIContentRpmDistributionTreesReadRequest) ExcludeFields(excludeFields []string) ContentDistributionTreesAPIContentRpmDistributionTreesReadRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r ContentDistributionTreesAPIContentRpmDistributionTreesReadRequest) Execute() (*RpmDistributionTreeResponse, *http.Response, error) {
	return r.ApiService.ContentRpmDistributionTreesReadExecute(r)
}

/*
ContentRpmDistributionTreesRead Inspect a distribution tree

Distribution Tree Viewset.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param rpmDistributionTreeHref
 @return ContentDistributionTreesAPIContentRpmDistributionTreesReadRequest
*/
func (a *ContentDistributionTreesAPIService) ContentRpmDistributionTreesRead(ctx context.Context, rpmDistributionTreeHref string) ContentDistributionTreesAPIContentRpmDistributionTreesReadRequest {
	return ContentDistributionTreesAPIContentRpmDistributionTreesReadRequest{
		ApiService: a,
		ctx: ctx,
		rpmDistributionTreeHref: rpmDistributionTreeHref,
	}
}

// Execute executes the request
//  @return RpmDistributionTreeResponse
func (a *ContentDistributionTreesAPIService) ContentRpmDistributionTreesReadExecute(r ContentDistributionTreesAPIContentRpmDistributionTreesReadRequest) (*RpmDistributionTreeResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *RpmDistributionTreeResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentDistributionTreesAPIService.ContentRpmDistributionTreesRead")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{rpm_distribution_tree_href}"
	localVarPath = strings.Replace(localVarPath, "{"+"rpm_distribution_tree_href"+"}", url.PathEscape(parameterValueToString(r.rpmDistributionTreeHref, "rpmDistributionTreeHref")), -1)
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
