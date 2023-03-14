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


// UsersApiService UsersApi service
type UsersApiService service

type UsersApiUsersCreateRequest struct {
	ctx context.Context
	ApiService *UsersApiService
	user *User
}

func (r UsersApiUsersCreateRequest) User(user User) UsersApiUsersCreateRequest {
	r.user = &user
	return r
}

func (r UsersApiUsersCreateRequest) Execute() (*UserResponse, *http.Response, error) {
	return r.ApiService.UsersCreateExecute(r)
}

/*
UsersCreate Create an user

ViewSet for User.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return UsersApiUsersCreateRequest
*/
func (a *UsersApiService) UsersCreate(ctx context.Context) UsersApiUsersCreateRequest {
	return UsersApiUsersCreateRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return UserResponse
func (a *UsersApiService) UsersCreateExecute(r UsersApiUsersCreateRequest) (*UserResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *UserResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsersApiService.UsersCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp/api/v3/users/"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.user == nil {
		return localVarReturnValue, nil, reportError("user is required and must be specified")
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
	localVarPostBody = r.user
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

type UsersApiUsersDeleteRequest struct {
	ctx context.Context
	ApiService *UsersApiService
	authUserHref string
}

func (r UsersApiUsersDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.UsersDeleteExecute(r)
}

/*
UsersDelete Delete an user

ViewSet for User.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param authUserHref
 @return UsersApiUsersDeleteRequest
*/
func (a *UsersApiService) UsersDelete(ctx context.Context, authUserHref string) UsersApiUsersDeleteRequest {
	return UsersApiUsersDeleteRequest{
		ApiService: a,
		ctx: ctx,
		authUserHref: authUserHref,
	}
}

// Execute executes the request
func (a *UsersApiService) UsersDeleteExecute(r UsersApiUsersDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsersApiService.UsersDelete")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{auth_user_href}"
	localVarPath = strings.Replace(localVarPath, "{"+"auth_user_href"+"}", url.PathEscape(parameterValueToString(r.authUserHref, "authUserHref")), -1)
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

type UsersApiUsersListRequest struct {
	ctx context.Context
	ApiService *UsersApiService
	email *string
	emailContains *string
	emailIcontains *string
	emailIexact *string
	emailIn *[]string
	firstName *string
	firstNameContains *string
	firstNameIcontains *string
	firstNameIexact *string
	firstNameIn *[]string
	isActive *bool
	isStaff *bool
	lastName *string
	lastNameContains *string
	lastNameIcontains *string
	lastNameIexact *string
	lastNameIn *[]string
	limit *int32
	offset *int32
	ordering *[]string
	username *string
	usernameContains *string
	usernameIcontains *string
	usernameIexact *string
	usernameIn *[]string
	fields *[]string
	excludeFields *[]string
}

// Filter results where email matches value
func (r UsersApiUsersListRequest) Email(email string) UsersApiUsersListRequest {
	r.email = &email
	return r
}

// Filter results where email contains value
func (r UsersApiUsersListRequest) EmailContains(emailContains string) UsersApiUsersListRequest {
	r.emailContains = &emailContains
	return r
}

// Filter results where email contains value
func (r UsersApiUsersListRequest) EmailIcontains(emailIcontains string) UsersApiUsersListRequest {
	r.emailIcontains = &emailIcontains
	return r
}

// Filter results where email matches value
func (r UsersApiUsersListRequest) EmailIexact(emailIexact string) UsersApiUsersListRequest {
	r.emailIexact = &emailIexact
	return r
}

// Filter results where email is in a comma-separated list of values
func (r UsersApiUsersListRequest) EmailIn(emailIn []string) UsersApiUsersListRequest {
	r.emailIn = &emailIn
	return r
}

// Filter results where first_name matches value
func (r UsersApiUsersListRequest) FirstName(firstName string) UsersApiUsersListRequest {
	r.firstName = &firstName
	return r
}

// Filter results where first_name contains value
func (r UsersApiUsersListRequest) FirstNameContains(firstNameContains string) UsersApiUsersListRequest {
	r.firstNameContains = &firstNameContains
	return r
}

// Filter results where first_name contains value
func (r UsersApiUsersListRequest) FirstNameIcontains(firstNameIcontains string) UsersApiUsersListRequest {
	r.firstNameIcontains = &firstNameIcontains
	return r
}

// Filter results where first_name matches value
func (r UsersApiUsersListRequest) FirstNameIexact(firstNameIexact string) UsersApiUsersListRequest {
	r.firstNameIexact = &firstNameIexact
	return r
}

// Filter results where first_name is in a comma-separated list of values
func (r UsersApiUsersListRequest) FirstNameIn(firstNameIn []string) UsersApiUsersListRequest {
	r.firstNameIn = &firstNameIn
	return r
}

// Filter results where is_active matches value
func (r UsersApiUsersListRequest) IsActive(isActive bool) UsersApiUsersListRequest {
	r.isActive = &isActive
	return r
}

// Filter results where is_staff matches value
func (r UsersApiUsersListRequest) IsStaff(isStaff bool) UsersApiUsersListRequest {
	r.isStaff = &isStaff
	return r
}

// Filter results where last_name matches value
func (r UsersApiUsersListRequest) LastName(lastName string) UsersApiUsersListRequest {
	r.lastName = &lastName
	return r
}

// Filter results where last_name contains value
func (r UsersApiUsersListRequest) LastNameContains(lastNameContains string) UsersApiUsersListRequest {
	r.lastNameContains = &lastNameContains
	return r
}

// Filter results where last_name contains value
func (r UsersApiUsersListRequest) LastNameIcontains(lastNameIcontains string) UsersApiUsersListRequest {
	r.lastNameIcontains = &lastNameIcontains
	return r
}

// Filter results where last_name matches value
func (r UsersApiUsersListRequest) LastNameIexact(lastNameIexact string) UsersApiUsersListRequest {
	r.lastNameIexact = &lastNameIexact
	return r
}

// Filter results where last_name is in a comma-separated list of values
func (r UsersApiUsersListRequest) LastNameIn(lastNameIn []string) UsersApiUsersListRequest {
	r.lastNameIn = &lastNameIn
	return r
}

// Number of results to return per page.
func (r UsersApiUsersListRequest) Limit(limit int32) UsersApiUsersListRequest {
	r.limit = &limit
	return r
}

// The initial index from which to return the results.
func (r UsersApiUsersListRequest) Offset(offset int32) UsersApiUsersListRequest {
	r.offset = &offset
	return r
}

// Ordering
func (r UsersApiUsersListRequest) Ordering(ordering []string) UsersApiUsersListRequest {
	r.ordering = &ordering
	return r
}

// Filter results where username matches value
func (r UsersApiUsersListRequest) Username(username string) UsersApiUsersListRequest {
	r.username = &username
	return r
}

// Filter results where username contains value
func (r UsersApiUsersListRequest) UsernameContains(usernameContains string) UsersApiUsersListRequest {
	r.usernameContains = &usernameContains
	return r
}

// Filter results where username contains value
func (r UsersApiUsersListRequest) UsernameIcontains(usernameIcontains string) UsersApiUsersListRequest {
	r.usernameIcontains = &usernameIcontains
	return r
}

// Filter results where username matches value
func (r UsersApiUsersListRequest) UsernameIexact(usernameIexact string) UsersApiUsersListRequest {
	r.usernameIexact = &usernameIexact
	return r
}

// Filter results where username is in a comma-separated list of values
func (r UsersApiUsersListRequest) UsernameIn(usernameIn []string) UsersApiUsersListRequest {
	r.usernameIn = &usernameIn
	return r
}

// A list of fields to include in the response.
func (r UsersApiUsersListRequest) Fields(fields []string) UsersApiUsersListRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r UsersApiUsersListRequest) ExcludeFields(excludeFields []string) UsersApiUsersListRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r UsersApiUsersListRequest) Execute() (*PaginatedUserResponseList, *http.Response, error) {
	return r.ApiService.UsersListExecute(r)
}

/*
UsersList List users

ViewSet for User.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return UsersApiUsersListRequest
*/
func (a *UsersApiService) UsersList(ctx context.Context) UsersApiUsersListRequest {
	return UsersApiUsersListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PaginatedUserResponseList
func (a *UsersApiService) UsersListExecute(r UsersApiUsersListRequest) (*PaginatedUserResponseList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedUserResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsersApiService.UsersList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp/api/v3/users/"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.email != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "email", r.email, "")
	}
	if r.emailContains != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "email__contains", r.emailContains, "")
	}
	if r.emailIcontains != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "email__icontains", r.emailIcontains, "")
	}
	if r.emailIexact != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "email__iexact", r.emailIexact, "")
	}
	if r.emailIn != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "email__in", r.emailIn, "csv")
	}
	if r.firstName != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "first_name", r.firstName, "")
	}
	if r.firstNameContains != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "first_name__contains", r.firstNameContains, "")
	}
	if r.firstNameIcontains != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "first_name__icontains", r.firstNameIcontains, "")
	}
	if r.firstNameIexact != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "first_name__iexact", r.firstNameIexact, "")
	}
	if r.firstNameIn != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "first_name__in", r.firstNameIn, "csv")
	}
	if r.isActive != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "is_active", r.isActive, "")
	}
	if r.isStaff != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "is_staff", r.isStaff, "")
	}
	if r.lastName != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "last_name", r.lastName, "")
	}
	if r.lastNameContains != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "last_name__contains", r.lastNameContains, "")
	}
	if r.lastNameIcontains != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "last_name__icontains", r.lastNameIcontains, "")
	}
	if r.lastNameIexact != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "last_name__iexact", r.lastNameIexact, "")
	}
	if r.lastNameIn != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "last_name__in", r.lastNameIn, "csv")
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
	if r.username != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "username", r.username, "")
	}
	if r.usernameContains != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "username__contains", r.usernameContains, "")
	}
	if r.usernameIcontains != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "username__icontains", r.usernameIcontains, "")
	}
	if r.usernameIexact != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "username__iexact", r.usernameIexact, "")
	}
	if r.usernameIn != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "username__in", r.usernameIn, "csv")
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

type UsersApiUsersPartialUpdateRequest struct {
	ctx context.Context
	ApiService *UsersApiService
	authUserHref string
	patchedUser *PatchedUser
}

func (r UsersApiUsersPartialUpdateRequest) PatchedUser(patchedUser PatchedUser) UsersApiUsersPartialUpdateRequest {
	r.patchedUser = &patchedUser
	return r
}

func (r UsersApiUsersPartialUpdateRequest) Execute() (*UserResponse, *http.Response, error) {
	return r.ApiService.UsersPartialUpdateExecute(r)
}

/*
UsersPartialUpdate Update an user

ViewSet for User.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param authUserHref
 @return UsersApiUsersPartialUpdateRequest
*/
func (a *UsersApiService) UsersPartialUpdate(ctx context.Context, authUserHref string) UsersApiUsersPartialUpdateRequest {
	return UsersApiUsersPartialUpdateRequest{
		ApiService: a,
		ctx: ctx,
		authUserHref: authUserHref,
	}
}

// Execute executes the request
//  @return UserResponse
func (a *UsersApiService) UsersPartialUpdateExecute(r UsersApiUsersPartialUpdateRequest) (*UserResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *UserResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsersApiService.UsersPartialUpdate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{auth_user_href}"
	localVarPath = strings.Replace(localVarPath, "{"+"auth_user_href"+"}", url.PathEscape(parameterValueToString(r.authUserHref, "authUserHref")), -1)
        localVarPath = strings.Replace(localVarPath, "/%2F", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.patchedUser == nil {
		return localVarReturnValue, nil, reportError("patchedUser is required and must be specified")
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
	localVarPostBody = r.patchedUser
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

type UsersApiUsersReadRequest struct {
	ctx context.Context
	ApiService *UsersApiService
	authUserHref string
	fields *[]string
	excludeFields *[]string
}

// A list of fields to include in the response.
func (r UsersApiUsersReadRequest) Fields(fields []string) UsersApiUsersReadRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r UsersApiUsersReadRequest) ExcludeFields(excludeFields []string) UsersApiUsersReadRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r UsersApiUsersReadRequest) Execute() (*UserResponse, *http.Response, error) {
	return r.ApiService.UsersReadExecute(r)
}

/*
UsersRead Inspect an user

ViewSet for User.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param authUserHref
 @return UsersApiUsersReadRequest
*/
func (a *UsersApiService) UsersRead(ctx context.Context, authUserHref string) UsersApiUsersReadRequest {
	return UsersApiUsersReadRequest{
		ApiService: a,
		ctx: ctx,
		authUserHref: authUserHref,
	}
}

// Execute executes the request
//  @return UserResponse
func (a *UsersApiService) UsersReadExecute(r UsersApiUsersReadRequest) (*UserResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *UserResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsersApiService.UsersRead")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{auth_user_href}"
	localVarPath = strings.Replace(localVarPath, "{"+"auth_user_href"+"}", url.PathEscape(parameterValueToString(r.authUserHref, "authUserHref")), -1)
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

type UsersApiUsersUpdateRequest struct {
	ctx context.Context
	ApiService *UsersApiService
	authUserHref string
	user *User
}

func (r UsersApiUsersUpdateRequest) User(user User) UsersApiUsersUpdateRequest {
	r.user = &user
	return r
}

func (r UsersApiUsersUpdateRequest) Execute() (*UserResponse, *http.Response, error) {
	return r.ApiService.UsersUpdateExecute(r)
}

/*
UsersUpdate Update an user

ViewSet for User.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param authUserHref
 @return UsersApiUsersUpdateRequest
*/
func (a *UsersApiService) UsersUpdate(ctx context.Context, authUserHref string) UsersApiUsersUpdateRequest {
	return UsersApiUsersUpdateRequest{
		ApiService: a,
		ctx: ctx,
		authUserHref: authUserHref,
	}
}

// Execute executes the request
//  @return UserResponse
func (a *UsersApiService) UsersUpdateExecute(r UsersApiUsersUpdateRequest) (*UserResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *UserResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsersApiService.UsersUpdate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{auth_user_href}"
	localVarPath = strings.Replace(localVarPath, "{"+"auth_user_href"+"}", url.PathEscape(parameterValueToString(r.authUserHref, "authUserHref")), -1)
        localVarPath = strings.Replace(localVarPath, "/%2F", "/", -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.user == nil {
		return localVarReturnValue, nil, reportError("user is required and must be specified")
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
	localVarPostBody = r.user
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
