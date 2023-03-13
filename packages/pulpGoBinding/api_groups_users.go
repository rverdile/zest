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


// GroupsUsersApiService GroupsUsersApi service
type GroupsUsersApiService service

type GroupsUsersApiGroupsUsersCreateRequest struct {
	ctx context.Context
	ApiService *GroupsUsersApiService
	groupHref string
	groupUser *GroupUser
}

func (r GroupsUsersApiGroupsUsersCreateRequest) GroupUser(groupUser GroupUser) GroupsUsersApiGroupsUsersCreateRequest {
	r.groupUser = &groupUser
	return r
}

func (r GroupsUsersApiGroupsUsersCreateRequest) Execute() (*GroupUserResponse, *http.Response, error) {
	return r.ApiService.GroupsUsersCreateExecute(r)
}

/*
GroupsUsersCreate Create an user

Add a user to a group.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param groupHref
 @return GroupsUsersApiGroupsUsersCreateRequest
*/
func (a *GroupsUsersApiService) GroupsUsersCreate(ctx context.Context, groupHref string) GroupsUsersApiGroupsUsersCreateRequest {
	return GroupsUsersApiGroupsUsersCreateRequest{
		ApiService: a,
		ctx: ctx,
		groupHref: groupHref,
	}
}

// Execute executes the request
//  @return GroupUserResponse
func (a *GroupsUsersApiService) GroupsUsersCreateExecute(r GroupsUsersApiGroupsUsersCreateRequest) (*GroupUserResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GroupUserResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GroupsUsersApiService.GroupsUsersCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{group_href}users/"
	localVarPath = strings.Replace(localVarPath, "{"+"group_href"+"}", url.PathEscape(parameterValueToString(r.groupHref, "groupHref")), -1)
        localVarPath = strings.Replace(localVarPath, "/%2F", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.groupUser == nil {
		return localVarReturnValue, nil, reportError("groupUser is required and must be specified")
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
	localVarPostBody = r.groupUser
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

type GroupsUsersApiGroupsUsersDeleteRequest struct {
	ctx context.Context
	ApiService *GroupsUsersApiService
	groupsUserHref string
}

func (r GroupsUsersApiGroupsUsersDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.GroupsUsersDeleteExecute(r)
}

/*
GroupsUsersDelete Delete an user

Remove a user from a group.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param groupsUserHref
 @return GroupsUsersApiGroupsUsersDeleteRequest
*/
func (a *GroupsUsersApiService) GroupsUsersDelete(ctx context.Context, groupsUserHref string) GroupsUsersApiGroupsUsersDeleteRequest {
	return GroupsUsersApiGroupsUsersDeleteRequest{
		ApiService: a,
		ctx: ctx,
		groupsUserHref: groupsUserHref,
	}
}

// Execute executes the request
func (a *GroupsUsersApiService) GroupsUsersDeleteExecute(r GroupsUsersApiGroupsUsersDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GroupsUsersApiService.GroupsUsersDelete")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{groups_user_href}"
	localVarPath = strings.Replace(localVarPath, "{"+"groups_user_href"+"}", url.PathEscape(parameterValueToString(r.groupsUserHref, "groupsUserHref")), -1)
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

type GroupsUsersApiGroupsUsersListRequest struct {
	ctx context.Context
	ApiService *GroupsUsersApiService
	groupHref string
	limit *int32
	offset *int32
	fields *[]string
	excludeFields *[]string
}

// Number of results to return per page.
func (r GroupsUsersApiGroupsUsersListRequest) Limit(limit int32) GroupsUsersApiGroupsUsersListRequest {
	r.limit = &limit
	return r
}

// The initial index from which to return the results.
func (r GroupsUsersApiGroupsUsersListRequest) Offset(offset int32) GroupsUsersApiGroupsUsersListRequest {
	r.offset = &offset
	return r
}

// A list of fields to include in the response.
func (r GroupsUsersApiGroupsUsersListRequest) Fields(fields []string) GroupsUsersApiGroupsUsersListRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r GroupsUsersApiGroupsUsersListRequest) ExcludeFields(excludeFields []string) GroupsUsersApiGroupsUsersListRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r GroupsUsersApiGroupsUsersListRequest) Execute() (*PaginatedGroupUserResponseList, *http.Response, error) {
	return r.ApiService.GroupsUsersListExecute(r)
}

/*
GroupsUsersList List users

List group users.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param groupHref
 @return GroupsUsersApiGroupsUsersListRequest
*/
func (a *GroupsUsersApiService) GroupsUsersList(ctx context.Context, groupHref string) GroupsUsersApiGroupsUsersListRequest {
	return GroupsUsersApiGroupsUsersListRequest{
		ApiService: a,
		ctx: ctx,
		groupHref: groupHref,
	}
}

// Execute executes the request
//  @return PaginatedGroupUserResponseList
func (a *GroupsUsersApiService) GroupsUsersListExecute(r GroupsUsersApiGroupsUsersListRequest) (*PaginatedGroupUserResponseList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedGroupUserResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GroupsUsersApiService.GroupsUsersList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{group_href}users/"
	localVarPath = strings.Replace(localVarPath, "{"+"group_href"+"}", url.PathEscape(parameterValueToString(r.groupHref, "groupHref")), -1)
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
