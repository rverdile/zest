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


// TaskSchedulesApiService TaskSchedulesApi service
type TaskSchedulesApiService service

type TaskSchedulesApiTaskSchedulesAddRoleRequest struct {
	ctx context.Context
	ApiService *TaskSchedulesApiService
	taskScheduleHref string
	nestedRole *NestedRole
}

func (r TaskSchedulesApiTaskSchedulesAddRoleRequest) NestedRole(nestedRole NestedRole) TaskSchedulesApiTaskSchedulesAddRoleRequest {
	r.nestedRole = &nestedRole
	return r
}

func (r TaskSchedulesApiTaskSchedulesAddRoleRequest) Execute() (*NestedRoleResponse, *http.Response, error) {
	return r.ApiService.TaskSchedulesAddRoleExecute(r)
}

/*
TaskSchedulesAddRole Method for TaskSchedulesAddRole

Add a role for this object to users/groups.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param taskScheduleHref
 @return TaskSchedulesApiTaskSchedulesAddRoleRequest
*/
func (a *TaskSchedulesApiService) TaskSchedulesAddRole(ctx context.Context, taskScheduleHref string) TaskSchedulesApiTaskSchedulesAddRoleRequest {
	return TaskSchedulesApiTaskSchedulesAddRoleRequest{
		ApiService: a,
		ctx: ctx,
		taskScheduleHref: taskScheduleHref,
	}
}

// Execute executes the request
//  @return NestedRoleResponse
func (a *TaskSchedulesApiService) TaskSchedulesAddRoleExecute(r TaskSchedulesApiTaskSchedulesAddRoleRequest) (*NestedRoleResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *NestedRoleResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TaskSchedulesApiService.TaskSchedulesAddRole")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{task_schedule_href}add_role/"
	localVarPath = strings.Replace(localVarPath, "{"+"task_schedule_href"+"}", url.PathEscape(parameterValueToString(r.taskScheduleHref, "taskScheduleHref")), -1)
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

type TaskSchedulesApiTaskSchedulesListRequest struct {
	ctx context.Context
	ApiService *TaskSchedulesApiService
	limit *int32
	name *string
	nameContains *string
	offset *int32
	ordering *[]string
	pulpHrefIn *[]string
	pulpIdIn *[]string
	taskName *string
	taskNameContains *string
	fields *[]string
	excludeFields *[]string
}

// Number of results to return per page.
func (r TaskSchedulesApiTaskSchedulesListRequest) Limit(limit int32) TaskSchedulesApiTaskSchedulesListRequest {
	r.limit = &limit
	return r
}

// Filter results where name matches value
func (r TaskSchedulesApiTaskSchedulesListRequest) Name(name string) TaskSchedulesApiTaskSchedulesListRequest {
	r.name = &name
	return r
}

// Filter results where name contains value
func (r TaskSchedulesApiTaskSchedulesListRequest) NameContains(nameContains string) TaskSchedulesApiTaskSchedulesListRequest {
	r.nameContains = &nameContains
	return r
}

// The initial index from which to return the results.
func (r TaskSchedulesApiTaskSchedulesListRequest) Offset(offset int32) TaskSchedulesApiTaskSchedulesListRequest {
	r.offset = &offset
	return r
}

// Ordering  * &#x60;pulp_id&#x60; - Pulp id * &#x60;-pulp_id&#x60; - Pulp id (descending) * &#x60;pulp_created&#x60; - Pulp created * &#x60;-pulp_created&#x60; - Pulp created (descending) * &#x60;pulp_last_updated&#x60; - Pulp last updated * &#x60;-pulp_last_updated&#x60; - Pulp last updated (descending) * &#x60;name&#x60; - Name * &#x60;-name&#x60; - Name (descending) * &#x60;next_dispatch&#x60; - Next dispatch * &#x60;-next_dispatch&#x60; - Next dispatch (descending) * &#x60;dispatch_interval&#x60; - Dispatch interval * &#x60;-dispatch_interval&#x60; - Dispatch interval (descending) * &#x60;task_name&#x60; - Task name * &#x60;-task_name&#x60; - Task name (descending) * &#x60;pk&#x60; - Pk * &#x60;-pk&#x60; - Pk (descending)
func (r TaskSchedulesApiTaskSchedulesListRequest) Ordering(ordering []string) TaskSchedulesApiTaskSchedulesListRequest {
	r.ordering = &ordering
	return r
}

// Multiple values may be separated by commas.
func (r TaskSchedulesApiTaskSchedulesListRequest) PulpHrefIn(pulpHrefIn []string) TaskSchedulesApiTaskSchedulesListRequest {
	r.pulpHrefIn = &pulpHrefIn
	return r
}

// Multiple values may be separated by commas.
func (r TaskSchedulesApiTaskSchedulesListRequest) PulpIdIn(pulpIdIn []string) TaskSchedulesApiTaskSchedulesListRequest {
	r.pulpIdIn = &pulpIdIn
	return r
}

// Filter results where task_name matches value
func (r TaskSchedulesApiTaskSchedulesListRequest) TaskName(taskName string) TaskSchedulesApiTaskSchedulesListRequest {
	r.taskName = &taskName
	return r
}

// Filter results where task_name contains value
func (r TaskSchedulesApiTaskSchedulesListRequest) TaskNameContains(taskNameContains string) TaskSchedulesApiTaskSchedulesListRequest {
	r.taskNameContains = &taskNameContains
	return r
}

// A list of fields to include in the response.
func (r TaskSchedulesApiTaskSchedulesListRequest) Fields(fields []string) TaskSchedulesApiTaskSchedulesListRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r TaskSchedulesApiTaskSchedulesListRequest) ExcludeFields(excludeFields []string) TaskSchedulesApiTaskSchedulesListRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r TaskSchedulesApiTaskSchedulesListRequest) Execute() (*PaginatedTaskScheduleResponseList, *http.Response, error) {
	return r.ApiService.TaskSchedulesListExecute(r)
}

/*
TaskSchedulesList List task schedules

ViewSet to monitor task schedules.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return TaskSchedulesApiTaskSchedulesListRequest
*/
func (a *TaskSchedulesApiService) TaskSchedulesList(ctx context.Context) TaskSchedulesApiTaskSchedulesListRequest {
	return TaskSchedulesApiTaskSchedulesListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PaginatedTaskScheduleResponseList
func (a *TaskSchedulesApiService) TaskSchedulesListExecute(r TaskSchedulesApiTaskSchedulesListRequest) (*PaginatedTaskScheduleResponseList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedTaskScheduleResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TaskSchedulesApiService.TaskSchedulesList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp/api/v3/task-schedules/"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	}
	if r.name != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "name", r.name, "")
	}
	if r.nameContains != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "name__contains", r.nameContains, "")
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
	if r.taskName != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "task_name", r.taskName, "")
	}
	if r.taskNameContains != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "task_name__contains", r.taskNameContains, "")
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

type TaskSchedulesApiTaskSchedulesListRolesRequest struct {
	ctx context.Context
	ApiService *TaskSchedulesApiService
	taskScheduleHref string
	fields *[]string
	excludeFields *[]string
}

// A list of fields to include in the response.
func (r TaskSchedulesApiTaskSchedulesListRolesRequest) Fields(fields []string) TaskSchedulesApiTaskSchedulesListRolesRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r TaskSchedulesApiTaskSchedulesListRolesRequest) ExcludeFields(excludeFields []string) TaskSchedulesApiTaskSchedulesListRolesRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r TaskSchedulesApiTaskSchedulesListRolesRequest) Execute() (*ObjectRolesResponse, *http.Response, error) {
	return r.ApiService.TaskSchedulesListRolesExecute(r)
}

/*
TaskSchedulesListRoles Method for TaskSchedulesListRoles

List roles assigned to this object.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param taskScheduleHref
 @return TaskSchedulesApiTaskSchedulesListRolesRequest
*/
func (a *TaskSchedulesApiService) TaskSchedulesListRoles(ctx context.Context, taskScheduleHref string) TaskSchedulesApiTaskSchedulesListRolesRequest {
	return TaskSchedulesApiTaskSchedulesListRolesRequest{
		ApiService: a,
		ctx: ctx,
		taskScheduleHref: taskScheduleHref,
	}
}

// Execute executes the request
//  @return ObjectRolesResponse
func (a *TaskSchedulesApiService) TaskSchedulesListRolesExecute(r TaskSchedulesApiTaskSchedulesListRolesRequest) (*ObjectRolesResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ObjectRolesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TaskSchedulesApiService.TaskSchedulesListRoles")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{task_schedule_href}list_roles/"
	localVarPath = strings.Replace(localVarPath, "{"+"task_schedule_href"+"}", url.PathEscape(parameterValueToString(r.taskScheduleHref, "taskScheduleHref")), -1)
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

type TaskSchedulesApiTaskSchedulesMyPermissionsRequest struct {
	ctx context.Context
	ApiService *TaskSchedulesApiService
	taskScheduleHref string
	fields *[]string
	excludeFields *[]string
}

// A list of fields to include in the response.
func (r TaskSchedulesApiTaskSchedulesMyPermissionsRequest) Fields(fields []string) TaskSchedulesApiTaskSchedulesMyPermissionsRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r TaskSchedulesApiTaskSchedulesMyPermissionsRequest) ExcludeFields(excludeFields []string) TaskSchedulesApiTaskSchedulesMyPermissionsRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r TaskSchedulesApiTaskSchedulesMyPermissionsRequest) Execute() (*MyPermissionsResponse, *http.Response, error) {
	return r.ApiService.TaskSchedulesMyPermissionsExecute(r)
}

/*
TaskSchedulesMyPermissions Method for TaskSchedulesMyPermissions

List permissions available to the current user on this object.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param taskScheduleHref
 @return TaskSchedulesApiTaskSchedulesMyPermissionsRequest
*/
func (a *TaskSchedulesApiService) TaskSchedulesMyPermissions(ctx context.Context, taskScheduleHref string) TaskSchedulesApiTaskSchedulesMyPermissionsRequest {
	return TaskSchedulesApiTaskSchedulesMyPermissionsRequest{
		ApiService: a,
		ctx: ctx,
		taskScheduleHref: taskScheduleHref,
	}
}

// Execute executes the request
//  @return MyPermissionsResponse
func (a *TaskSchedulesApiService) TaskSchedulesMyPermissionsExecute(r TaskSchedulesApiTaskSchedulesMyPermissionsRequest) (*MyPermissionsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MyPermissionsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TaskSchedulesApiService.TaskSchedulesMyPermissions")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{task_schedule_href}my_permissions/"
	localVarPath = strings.Replace(localVarPath, "{"+"task_schedule_href"+"}", url.PathEscape(parameterValueToString(r.taskScheduleHref, "taskScheduleHref")), -1)
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

type TaskSchedulesApiTaskSchedulesReadRequest struct {
	ctx context.Context
	ApiService *TaskSchedulesApiService
	taskScheduleHref string
	fields *[]string
	excludeFields *[]string
}

// A list of fields to include in the response.
func (r TaskSchedulesApiTaskSchedulesReadRequest) Fields(fields []string) TaskSchedulesApiTaskSchedulesReadRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r TaskSchedulesApiTaskSchedulesReadRequest) ExcludeFields(excludeFields []string) TaskSchedulesApiTaskSchedulesReadRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r TaskSchedulesApiTaskSchedulesReadRequest) Execute() (*TaskScheduleResponse, *http.Response, error) {
	return r.ApiService.TaskSchedulesReadExecute(r)
}

/*
TaskSchedulesRead Inspect a task schedule

ViewSet to monitor task schedules.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param taskScheduleHref
 @return TaskSchedulesApiTaskSchedulesReadRequest
*/
func (a *TaskSchedulesApiService) TaskSchedulesRead(ctx context.Context, taskScheduleHref string) TaskSchedulesApiTaskSchedulesReadRequest {
	return TaskSchedulesApiTaskSchedulesReadRequest{
		ApiService: a,
		ctx: ctx,
		taskScheduleHref: taskScheduleHref,
	}
}

// Execute executes the request
//  @return TaskScheduleResponse
func (a *TaskSchedulesApiService) TaskSchedulesReadExecute(r TaskSchedulesApiTaskSchedulesReadRequest) (*TaskScheduleResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *TaskScheduleResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TaskSchedulesApiService.TaskSchedulesRead")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{task_schedule_href}"
	localVarPath = strings.Replace(localVarPath, "{"+"task_schedule_href"+"}", url.PathEscape(parameterValueToString(r.taskScheduleHref, "taskScheduleHref")), -1)
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

type TaskSchedulesApiTaskSchedulesRemoveRoleRequest struct {
	ctx context.Context
	ApiService *TaskSchedulesApiService
	taskScheduleHref string
	nestedRole *NestedRole
}

func (r TaskSchedulesApiTaskSchedulesRemoveRoleRequest) NestedRole(nestedRole NestedRole) TaskSchedulesApiTaskSchedulesRemoveRoleRequest {
	r.nestedRole = &nestedRole
	return r
}

func (r TaskSchedulesApiTaskSchedulesRemoveRoleRequest) Execute() (*NestedRoleResponse, *http.Response, error) {
	return r.ApiService.TaskSchedulesRemoveRoleExecute(r)
}

/*
TaskSchedulesRemoveRole Method for TaskSchedulesRemoveRole

Remove a role for this object from users/groups.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param taskScheduleHref
 @return TaskSchedulesApiTaskSchedulesRemoveRoleRequest
*/
func (a *TaskSchedulesApiService) TaskSchedulesRemoveRole(ctx context.Context, taskScheduleHref string) TaskSchedulesApiTaskSchedulesRemoveRoleRequest {
	return TaskSchedulesApiTaskSchedulesRemoveRoleRequest{
		ApiService: a,
		ctx: ctx,
		taskScheduleHref: taskScheduleHref,
	}
}

// Execute executes the request
//  @return NestedRoleResponse
func (a *TaskSchedulesApiService) TaskSchedulesRemoveRoleExecute(r TaskSchedulesApiTaskSchedulesRemoveRoleRequest) (*NestedRoleResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *NestedRoleResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TaskSchedulesApiService.TaskSchedulesRemoveRole")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{task_schedule_href}remove_role/"
	localVarPath = strings.Replace(localVarPath, "{"+"task_schedule_href"+"}", url.PathEscape(parameterValueToString(r.taskScheduleHref, "taskScheduleHref")), -1)
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
