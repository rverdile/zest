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
	"os"
)


// AnsibleCollectionsApiService AnsibleCollectionsApi service
type AnsibleCollectionsApiService service

type AnsibleCollectionsApiAnsibleCollectionsAddRoleRequest struct {
	ctx context.Context
	ApiService *AnsibleCollectionsApiService
	ansibleCollectionHref string
	nestedRole *NestedRole
}

func (r AnsibleCollectionsApiAnsibleCollectionsAddRoleRequest) NestedRole(nestedRole NestedRole) AnsibleCollectionsApiAnsibleCollectionsAddRoleRequest {
	r.nestedRole = &nestedRole
	return r
}

func (r AnsibleCollectionsApiAnsibleCollectionsAddRoleRequest) Execute() (*NestedRoleResponse, *http.Response, error) {
	return r.ApiService.AnsibleCollectionsAddRoleExecute(r)
}

/*
AnsibleCollectionsAddRole Method for AnsibleCollectionsAddRole

Add a role for this object to users/groups.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ansibleCollectionHref
 @return AnsibleCollectionsApiAnsibleCollectionsAddRoleRequest
*/
func (a *AnsibleCollectionsApiService) AnsibleCollectionsAddRole(ctx context.Context, ansibleCollectionHref string) AnsibleCollectionsApiAnsibleCollectionsAddRoleRequest {
	return AnsibleCollectionsApiAnsibleCollectionsAddRoleRequest{
		ApiService: a,
		ctx: ctx,
		ansibleCollectionHref: ansibleCollectionHref,
	}
}

// Execute executes the request
//  @return NestedRoleResponse
func (a *AnsibleCollectionsApiService) AnsibleCollectionsAddRoleExecute(r AnsibleCollectionsApiAnsibleCollectionsAddRoleRequest) (*NestedRoleResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *NestedRoleResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AnsibleCollectionsApiService.AnsibleCollectionsAddRole")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{ansible_collection_href}add_role/"
	localVarPath = strings.Replace(localVarPath, "{"+"ansible_collection_href"+"}", url.PathEscape(parameterValueToString(r.ansibleCollectionHref, "ansibleCollectionHref")), -1)
        localVarPath = strings.Replace(localVarPath, "/%2F", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.nestedRole == nil {
		return localVarReturnValue, nil, reportError("nestedRole is required and must be specified")
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
	localVarPostBody = r.nestedRole
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

type AnsibleCollectionsApiAnsibleCollectionsListRequest struct {
	ctx context.Context
	ApiService *AnsibleCollectionsApiService
	limit *int32
	name *string
	namespace *string
	offset *int32
	ordering *[]string
	pulpHrefIn *[]string
	pulpIdIn *[]string
	fields *[]string
	excludeFields *[]string
}

// Number of results to return per page.
func (r AnsibleCollectionsApiAnsibleCollectionsListRequest) Limit(limit int32) AnsibleCollectionsApiAnsibleCollectionsListRequest {
	r.limit = &limit
	return r
}

func (r AnsibleCollectionsApiAnsibleCollectionsListRequest) Name(name string) AnsibleCollectionsApiAnsibleCollectionsListRequest {
	r.name = &name
	return r
}

func (r AnsibleCollectionsApiAnsibleCollectionsListRequest) Namespace(namespace string) AnsibleCollectionsApiAnsibleCollectionsListRequest {
	r.namespace = &namespace
	return r
}

// The initial index from which to return the results.
func (r AnsibleCollectionsApiAnsibleCollectionsListRequest) Offset(offset int32) AnsibleCollectionsApiAnsibleCollectionsListRequest {
	r.offset = &offset
	return r
}

// Ordering  * &#x60;pulp_id&#x60; - Pulp id * &#x60;-pulp_id&#x60; - Pulp id (descending) * &#x60;pulp_created&#x60; - Pulp created * &#x60;-pulp_created&#x60; - Pulp created (descending) * &#x60;pulp_last_updated&#x60; - Pulp last updated * &#x60;-pulp_last_updated&#x60; - Pulp last updated (descending) * &#x60;namespace&#x60; - Namespace * &#x60;-namespace&#x60; - Namespace (descending) * &#x60;name&#x60; - Name * &#x60;-name&#x60; - Name (descending) * &#x60;pk&#x60; - Pk * &#x60;-pk&#x60; - Pk (descending)
func (r AnsibleCollectionsApiAnsibleCollectionsListRequest) Ordering(ordering []string) AnsibleCollectionsApiAnsibleCollectionsListRequest {
	r.ordering = &ordering
	return r
}

// Multiple values may be separated by commas.
func (r AnsibleCollectionsApiAnsibleCollectionsListRequest) PulpHrefIn(pulpHrefIn []string) AnsibleCollectionsApiAnsibleCollectionsListRequest {
	r.pulpHrefIn = &pulpHrefIn
	return r
}

// Multiple values may be separated by commas.
func (r AnsibleCollectionsApiAnsibleCollectionsListRequest) PulpIdIn(pulpIdIn []string) AnsibleCollectionsApiAnsibleCollectionsListRequest {
	r.pulpIdIn = &pulpIdIn
	return r
}

// A list of fields to include in the response.
func (r AnsibleCollectionsApiAnsibleCollectionsListRequest) Fields(fields []string) AnsibleCollectionsApiAnsibleCollectionsListRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r AnsibleCollectionsApiAnsibleCollectionsListRequest) ExcludeFields(excludeFields []string) AnsibleCollectionsApiAnsibleCollectionsListRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r AnsibleCollectionsApiAnsibleCollectionsListRequest) Execute() (*PaginatedansibleCollectionResponseList, *http.Response, error) {
	return r.ApiService.AnsibleCollectionsListExecute(r)
}

/*
AnsibleCollectionsList List collections

Viewset for Ansible Collections.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return AnsibleCollectionsApiAnsibleCollectionsListRequest
*/
func (a *AnsibleCollectionsApiService) AnsibleCollectionsList(ctx context.Context) AnsibleCollectionsApiAnsibleCollectionsListRequest {
	return AnsibleCollectionsApiAnsibleCollectionsListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PaginatedansibleCollectionResponseList
func (a *AnsibleCollectionsApiService) AnsibleCollectionsListExecute(r AnsibleCollectionsApiAnsibleCollectionsListRequest) (*PaginatedansibleCollectionResponseList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedansibleCollectionResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AnsibleCollectionsApiService.AnsibleCollectionsList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp/api/v3/ansible/collections/"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	}
	if r.name != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "name", r.name, "")
	}
	if r.namespace != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "namespace", r.namespace, "")
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

type AnsibleCollectionsApiAnsibleCollectionsListRolesRequest struct {
	ctx context.Context
	ApiService *AnsibleCollectionsApiService
	ansibleCollectionHref string
	fields *[]string
	excludeFields *[]string
}

// A list of fields to include in the response.
func (r AnsibleCollectionsApiAnsibleCollectionsListRolesRequest) Fields(fields []string) AnsibleCollectionsApiAnsibleCollectionsListRolesRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r AnsibleCollectionsApiAnsibleCollectionsListRolesRequest) ExcludeFields(excludeFields []string) AnsibleCollectionsApiAnsibleCollectionsListRolesRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r AnsibleCollectionsApiAnsibleCollectionsListRolesRequest) Execute() (*ObjectRolesResponse, *http.Response, error) {
	return r.ApiService.AnsibleCollectionsListRolesExecute(r)
}

/*
AnsibleCollectionsListRoles Method for AnsibleCollectionsListRoles

List roles assigned to this object.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ansibleCollectionHref
 @return AnsibleCollectionsApiAnsibleCollectionsListRolesRequest
*/
func (a *AnsibleCollectionsApiService) AnsibleCollectionsListRoles(ctx context.Context, ansibleCollectionHref string) AnsibleCollectionsApiAnsibleCollectionsListRolesRequest {
	return AnsibleCollectionsApiAnsibleCollectionsListRolesRequest{
		ApiService: a,
		ctx: ctx,
		ansibleCollectionHref: ansibleCollectionHref,
	}
}

// Execute executes the request
//  @return ObjectRolesResponse
func (a *AnsibleCollectionsApiService) AnsibleCollectionsListRolesExecute(r AnsibleCollectionsApiAnsibleCollectionsListRolesRequest) (*ObjectRolesResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ObjectRolesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AnsibleCollectionsApiService.AnsibleCollectionsListRoles")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{ansible_collection_href}list_roles/"
	localVarPath = strings.Replace(localVarPath, "{"+"ansible_collection_href"+"}", url.PathEscape(parameterValueToString(r.ansibleCollectionHref, "ansibleCollectionHref")), -1)
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

type AnsibleCollectionsApiAnsibleCollectionsMyPermissionsRequest struct {
	ctx context.Context
	ApiService *AnsibleCollectionsApiService
	ansibleCollectionHref string
	fields *[]string
	excludeFields *[]string
}

// A list of fields to include in the response.
func (r AnsibleCollectionsApiAnsibleCollectionsMyPermissionsRequest) Fields(fields []string) AnsibleCollectionsApiAnsibleCollectionsMyPermissionsRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r AnsibleCollectionsApiAnsibleCollectionsMyPermissionsRequest) ExcludeFields(excludeFields []string) AnsibleCollectionsApiAnsibleCollectionsMyPermissionsRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r AnsibleCollectionsApiAnsibleCollectionsMyPermissionsRequest) Execute() (*MyPermissionsResponse, *http.Response, error) {
	return r.ApiService.AnsibleCollectionsMyPermissionsExecute(r)
}

/*
AnsibleCollectionsMyPermissions Method for AnsibleCollectionsMyPermissions

List permissions available to the current user on this object.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ansibleCollectionHref
 @return AnsibleCollectionsApiAnsibleCollectionsMyPermissionsRequest
*/
func (a *AnsibleCollectionsApiService) AnsibleCollectionsMyPermissions(ctx context.Context, ansibleCollectionHref string) AnsibleCollectionsApiAnsibleCollectionsMyPermissionsRequest {
	return AnsibleCollectionsApiAnsibleCollectionsMyPermissionsRequest{
		ApiService: a,
		ctx: ctx,
		ansibleCollectionHref: ansibleCollectionHref,
	}
}

// Execute executes the request
//  @return MyPermissionsResponse
func (a *AnsibleCollectionsApiService) AnsibleCollectionsMyPermissionsExecute(r AnsibleCollectionsApiAnsibleCollectionsMyPermissionsRequest) (*MyPermissionsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MyPermissionsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AnsibleCollectionsApiService.AnsibleCollectionsMyPermissions")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{ansible_collection_href}my_permissions/"
	localVarPath = strings.Replace(localVarPath, "{"+"ansible_collection_href"+"}", url.PathEscape(parameterValueToString(r.ansibleCollectionHref, "ansibleCollectionHref")), -1)
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

type AnsibleCollectionsApiAnsibleCollectionsRemoveRoleRequest struct {
	ctx context.Context
	ApiService *AnsibleCollectionsApiService
	ansibleCollectionHref string
	nestedRole *NestedRole
}

func (r AnsibleCollectionsApiAnsibleCollectionsRemoveRoleRequest) NestedRole(nestedRole NestedRole) AnsibleCollectionsApiAnsibleCollectionsRemoveRoleRequest {
	r.nestedRole = &nestedRole
	return r
}

func (r AnsibleCollectionsApiAnsibleCollectionsRemoveRoleRequest) Execute() (*NestedRoleResponse, *http.Response, error) {
	return r.ApiService.AnsibleCollectionsRemoveRoleExecute(r)
}

/*
AnsibleCollectionsRemoveRole Method for AnsibleCollectionsRemoveRole

Remove a role for this object from users/groups.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ansibleCollectionHref
 @return AnsibleCollectionsApiAnsibleCollectionsRemoveRoleRequest
*/
func (a *AnsibleCollectionsApiService) AnsibleCollectionsRemoveRole(ctx context.Context, ansibleCollectionHref string) AnsibleCollectionsApiAnsibleCollectionsRemoveRoleRequest {
	return AnsibleCollectionsApiAnsibleCollectionsRemoveRoleRequest{
		ApiService: a,
		ctx: ctx,
		ansibleCollectionHref: ansibleCollectionHref,
	}
}

// Execute executes the request
//  @return NestedRoleResponse
func (a *AnsibleCollectionsApiService) AnsibleCollectionsRemoveRoleExecute(r AnsibleCollectionsApiAnsibleCollectionsRemoveRoleRequest) (*NestedRoleResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *NestedRoleResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AnsibleCollectionsApiService.AnsibleCollectionsRemoveRole")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{ansible_collection_href}remove_role/"
	localVarPath = strings.Replace(localVarPath, "{"+"ansible_collection_href"+"}", url.PathEscape(parameterValueToString(r.ansibleCollectionHref, "ansibleCollectionHref")), -1)
        localVarPath = strings.Replace(localVarPath, "/%2F", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.nestedRole == nil {
		return localVarReturnValue, nil, reportError("nestedRole is required and must be specified")
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
	localVarPostBody = r.nestedRole
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

type AnsibleCollectionsApiUploadCollectionRequest struct {
	ctx context.Context
	ApiService *AnsibleCollectionsApiService
	file *os.File
	sha256 *string
	expectedNamespace *string
	expectedName *string
	expectedVersion *string
}

// The Collection tarball.
func (r AnsibleCollectionsApiUploadCollectionRequest) File(file *os.File) AnsibleCollectionsApiUploadCollectionRequest {
	r.file = file
	return r
}

// An optional sha256 checksum of the uploaded file.
func (r AnsibleCollectionsApiUploadCollectionRequest) Sha256(sha256 string) AnsibleCollectionsApiUploadCollectionRequest {
	r.sha256 = &sha256
	return r
}

// The expected &#39;namespace&#39; of the Collection to be verified against the metadata during import.
func (r AnsibleCollectionsApiUploadCollectionRequest) ExpectedNamespace(expectedNamespace string) AnsibleCollectionsApiUploadCollectionRequest {
	r.expectedNamespace = &expectedNamespace
	return r
}

// The expected &#39;name&#39; of the Collection to be verified against the metadata during import.
func (r AnsibleCollectionsApiUploadCollectionRequest) ExpectedName(expectedName string) AnsibleCollectionsApiUploadCollectionRequest {
	r.expectedName = &expectedName
	return r
}

// The expected version of the Collection to be verified against the metadata during import.
func (r AnsibleCollectionsApiUploadCollectionRequest) ExpectedVersion(expectedVersion string) AnsibleCollectionsApiUploadCollectionRequest {
	r.expectedVersion = &expectedVersion
	return r
}

func (r AnsibleCollectionsApiUploadCollectionRequest) Execute() (*AsyncOperationResponse, *http.Response, error) {
	return r.ApiService.UploadCollectionExecute(r)
}

/*
UploadCollection Upload a collection

Create an artifact and trigger an asynchronous task to create Collection content from it.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return AnsibleCollectionsApiUploadCollectionRequest

Deprecated
*/
func (a *AnsibleCollectionsApiService) UploadCollection(ctx context.Context) AnsibleCollectionsApiUploadCollectionRequest {
	return AnsibleCollectionsApiUploadCollectionRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AsyncOperationResponse
// Deprecated
func (a *AnsibleCollectionsApiService) UploadCollectionExecute(r AnsibleCollectionsApiUploadCollectionRequest) (*AsyncOperationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AsyncOperationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AnsibleCollectionsApiService.UploadCollection")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/ansible/collections/"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.file == nil {
		return localVarReturnValue, nil, reportError("file is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"multipart/form-data", "application/x-www-form-urlencoded"}

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
	var fileLocalVarFormFileName string
	var fileLocalVarFileName     string
	var fileLocalVarFileBytes    []byte

	fileLocalVarFormFileName = "file"


	fileLocalVarFile := r.file

	if fileLocalVarFile != nil {
		fbs, _ := io.ReadAll(fileLocalVarFile)

		fileLocalVarFileBytes = fbs
		fileLocalVarFileName = fileLocalVarFile.Name()
		fileLocalVarFile.Close()
		formFiles = append(formFiles, formFile{fileBytes: fileLocalVarFileBytes, fileName: fileLocalVarFileName, formFileName: fileLocalVarFormFileName})
	}
	if r.sha256 != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "sha256", r.sha256, "")
	}
	if r.expectedNamespace != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "expected_namespace", r.expectedNamespace, "")
	}
	if r.expectedName != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "expected_name", r.expectedName, "")
	}
	if r.expectedVersion != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "expected_version", r.expectedVersion, "")
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
