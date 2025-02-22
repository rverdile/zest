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


// ContentPackageenvironmentsAPIService ContentPackageenvironmentsAPI service
type ContentPackageenvironmentsAPIService service

type ContentPackageenvironmentsAPIContentRpmPackageenvironmentsListRequest struct {
	ctx context.Context
	ApiService *ContentPackageenvironmentsAPIService
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
func (r ContentPackageenvironmentsAPIContentRpmPackageenvironmentsListRequest) Limit(limit int32) ContentPackageenvironmentsAPIContentRpmPackageenvironmentsListRequest {
	r.limit = &limit
	return r
}

// The initial index from which to return the results.
func (r ContentPackageenvironmentsAPIContentRpmPackageenvironmentsListRequest) Offset(offset int32) ContentPackageenvironmentsAPIContentRpmPackageenvironmentsListRequest {
	r.offset = &offset
	return r
}

// Ordering  * &#x60;pk&#x60; - Pk * &#x60;-pk&#x60; - Pk (descending)
func (r ContentPackageenvironmentsAPIContentRpmPackageenvironmentsListRequest) Ordering(ordering []string) ContentPackageenvironmentsAPIContentRpmPackageenvironmentsListRequest {
	r.ordering = &ordering
	return r
}

// Multiple values may be separated by commas.
func (r ContentPackageenvironmentsAPIContentRpmPackageenvironmentsListRequest) PulpHrefIn(pulpHrefIn []string) ContentPackageenvironmentsAPIContentRpmPackageenvironmentsListRequest {
	r.pulpHrefIn = &pulpHrefIn
	return r
}

// Multiple values may be separated by commas.
func (r ContentPackageenvironmentsAPIContentRpmPackageenvironmentsListRequest) PulpIdIn(pulpIdIn []string) ContentPackageenvironmentsAPIContentRpmPackageenvironmentsListRequest {
	r.pulpIdIn = &pulpIdIn
	return r
}

// Repository Version referenced by HREF
func (r ContentPackageenvironmentsAPIContentRpmPackageenvironmentsListRequest) RepositoryVersion(repositoryVersion string) ContentPackageenvironmentsAPIContentRpmPackageenvironmentsListRequest {
	r.repositoryVersion = &repositoryVersion
	return r
}

// Repository Version referenced by HREF
func (r ContentPackageenvironmentsAPIContentRpmPackageenvironmentsListRequest) RepositoryVersionAdded(repositoryVersionAdded string) ContentPackageenvironmentsAPIContentRpmPackageenvironmentsListRequest {
	r.repositoryVersionAdded = &repositoryVersionAdded
	return r
}

// Repository Version referenced by HREF
func (r ContentPackageenvironmentsAPIContentRpmPackageenvironmentsListRequest) RepositoryVersionRemoved(repositoryVersionRemoved string) ContentPackageenvironmentsAPIContentRpmPackageenvironmentsListRequest {
	r.repositoryVersionRemoved = &repositoryVersionRemoved
	return r
}

// A list of fields to include in the response.
func (r ContentPackageenvironmentsAPIContentRpmPackageenvironmentsListRequest) Fields(fields []string) ContentPackageenvironmentsAPIContentRpmPackageenvironmentsListRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r ContentPackageenvironmentsAPIContentRpmPackageenvironmentsListRequest) ExcludeFields(excludeFields []string) ContentPackageenvironmentsAPIContentRpmPackageenvironmentsListRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r ContentPackageenvironmentsAPIContentRpmPackageenvironmentsListRequest) Execute() (*PaginatedrpmPackageEnvironmentResponseList, *http.Response, error) {
	return r.ApiService.ContentRpmPackageenvironmentsListExecute(r)
}

/*
ContentRpmPackageenvironmentsList List package environments

PackageEnvironment ViewSet.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pulpDomain
 @return ContentPackageenvironmentsAPIContentRpmPackageenvironmentsListRequest
*/
func (a *ContentPackageenvironmentsAPIService) ContentRpmPackageenvironmentsList(ctx context.Context, pulpDomain string) ContentPackageenvironmentsAPIContentRpmPackageenvironmentsListRequest {
	return ContentPackageenvironmentsAPIContentRpmPackageenvironmentsListRequest{
		ApiService: a,
		ctx: ctx,
		pulpDomain: pulpDomain,
	}
}

// Execute executes the request
//  @return PaginatedrpmPackageEnvironmentResponseList
func (a *ContentPackageenvironmentsAPIService) ContentRpmPackageenvironmentsListExecute(r ContentPackageenvironmentsAPIContentRpmPackageenvironmentsListRequest) (*PaginatedrpmPackageEnvironmentResponseList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedrpmPackageEnvironmentResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentPackageenvironmentsAPIService.ContentRpmPackageenvironmentsList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp/{pulp_domain}/api/v3/content/rpm/packageenvironments/"
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

type ContentPackageenvironmentsAPIContentRpmPackageenvironmentsReadRequest struct {
	ctx context.Context
	ApiService *ContentPackageenvironmentsAPIService
	rpmPackageEnvironmentHref string
	fields *[]string
	excludeFields *[]string
}

// A list of fields to include in the response.
func (r ContentPackageenvironmentsAPIContentRpmPackageenvironmentsReadRequest) Fields(fields []string) ContentPackageenvironmentsAPIContentRpmPackageenvironmentsReadRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r ContentPackageenvironmentsAPIContentRpmPackageenvironmentsReadRequest) ExcludeFields(excludeFields []string) ContentPackageenvironmentsAPIContentRpmPackageenvironmentsReadRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r ContentPackageenvironmentsAPIContentRpmPackageenvironmentsReadRequest) Execute() (*RpmPackageEnvironmentResponse, *http.Response, error) {
	return r.ApiService.ContentRpmPackageenvironmentsReadExecute(r)
}

/*
ContentRpmPackageenvironmentsRead Inspect a package environment

PackageEnvironment ViewSet.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param rpmPackageEnvironmentHref
 @return ContentPackageenvironmentsAPIContentRpmPackageenvironmentsReadRequest
*/
func (a *ContentPackageenvironmentsAPIService) ContentRpmPackageenvironmentsRead(ctx context.Context, rpmPackageEnvironmentHref string) ContentPackageenvironmentsAPIContentRpmPackageenvironmentsReadRequest {
	return ContentPackageenvironmentsAPIContentRpmPackageenvironmentsReadRequest{
		ApiService: a,
		ctx: ctx,
		rpmPackageEnvironmentHref: rpmPackageEnvironmentHref,
	}
}

// Execute executes the request
//  @return RpmPackageEnvironmentResponse
func (a *ContentPackageenvironmentsAPIService) ContentRpmPackageenvironmentsReadExecute(r ContentPackageenvironmentsAPIContentRpmPackageenvironmentsReadRequest) (*RpmPackageEnvironmentResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *RpmPackageEnvironmentResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentPackageenvironmentsAPIService.ContentRpmPackageenvironmentsRead")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{rpm_package_environment_href}"
	localVarPath = strings.Replace(localVarPath, "{"+"rpm_package_environment_href"+"}", url.PathEscape(parameterValueToString(r.rpmPackageEnvironmentHref, "rpmPackageEnvironmentHref")), -1)
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
