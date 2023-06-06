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
	"os"
	"reflect"
)


// ContentCollectionVersionsAPIService ContentCollectionVersionsAPI service
type ContentCollectionVersionsAPIService service

type ContentCollectionVersionsAPIContentAnsibleCollectionVersionsCreateRequest struct {
	ctx context.Context
	ApiService *ContentCollectionVersionsAPIService
	repository *string
	upload *string
	artifact *string
	file *os.File
	expectedName *string
	expectedNamespace *string
	expectedVersion *string
}

// A URI of a repository the new content unit should be associated with.
func (r ContentCollectionVersionsAPIContentAnsibleCollectionVersionsCreateRequest) Repository(repository string) ContentCollectionVersionsAPIContentAnsibleCollectionVersionsCreateRequest {
	r.repository = &repository
	return r
}

// An uncommitted upload that may be turned into the artifact of the content unit.
func (r ContentCollectionVersionsAPIContentAnsibleCollectionVersionsCreateRequest) Upload(upload string) ContentCollectionVersionsAPIContentAnsibleCollectionVersionsCreateRequest {
	r.upload = &upload
	return r
}

// Artifact file representing the physical content
func (r ContentCollectionVersionsAPIContentAnsibleCollectionVersionsCreateRequest) Artifact(artifact string) ContentCollectionVersionsAPIContentAnsibleCollectionVersionsCreateRequest {
	r.artifact = &artifact
	return r
}

// An uploaded file that may be turned into the artifact of the content unit.
func (r ContentCollectionVersionsAPIContentAnsibleCollectionVersionsCreateRequest) File(file *os.File) ContentCollectionVersionsAPIContentAnsibleCollectionVersionsCreateRequest {
	r.file = file
	return r
}

// The name of the collection.
func (r ContentCollectionVersionsAPIContentAnsibleCollectionVersionsCreateRequest) ExpectedName(expectedName string) ContentCollectionVersionsAPIContentAnsibleCollectionVersionsCreateRequest {
	r.expectedName = &expectedName
	return r
}

// The namespace of the collection.
func (r ContentCollectionVersionsAPIContentAnsibleCollectionVersionsCreateRequest) ExpectedNamespace(expectedNamespace string) ContentCollectionVersionsAPIContentAnsibleCollectionVersionsCreateRequest {
	r.expectedNamespace = &expectedNamespace
	return r
}

// The version of the collection.
func (r ContentCollectionVersionsAPIContentAnsibleCollectionVersionsCreateRequest) ExpectedVersion(expectedVersion string) ContentCollectionVersionsAPIContentAnsibleCollectionVersionsCreateRequest {
	r.expectedVersion = &expectedVersion
	return r
}

func (r ContentCollectionVersionsAPIContentAnsibleCollectionVersionsCreateRequest) Execute() (*AsyncOperationResponse, *http.Response, error) {
	return r.ApiService.ContentAnsibleCollectionVersionsCreateExecute(r)
}

/*
ContentAnsibleCollectionVersionsCreate Create a collection version

Trigger an asynchronous task to create content,optionally create new repository version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ContentCollectionVersionsAPIContentAnsibleCollectionVersionsCreateRequest
*/
func (a *ContentCollectionVersionsAPIService) ContentAnsibleCollectionVersionsCreate(ctx context.Context) ContentCollectionVersionsAPIContentAnsibleCollectionVersionsCreateRequest {
	return ContentCollectionVersionsAPIContentAnsibleCollectionVersionsCreateRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AsyncOperationResponse
func (a *ContentCollectionVersionsAPIService) ContentAnsibleCollectionVersionsCreateExecute(r ContentCollectionVersionsAPIContentAnsibleCollectionVersionsCreateRequest) (*AsyncOperationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AsyncOperationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentCollectionVersionsAPIService.ContentAnsibleCollectionVersionsCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp/api/v3/content/ansible/collection_versions/"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
	if r.repository != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "repository", r.repository, "")
	}
	if r.upload != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "upload", r.upload, "")
	}
	if r.artifact != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "artifact", r.artifact, "")
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
	if r.expectedName != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "expected_name", r.expectedName, "")
	}
	if r.expectedNamespace != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "expected_namespace", r.expectedNamespace, "")
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

type ContentCollectionVersionsAPIContentAnsibleCollectionVersionsListRequest struct {
	ctx context.Context
	ApiService *ContentCollectionVersionsAPIService
	isHighest *bool
	limit *int32
	name *string
	namespace *string
	offset *int32
	ordering *[]string
	pulpHrefIn *[]string
	pulpIdIn *[]string
	q *string
	repositoryVersion *string
	repositoryVersionAdded *string
	repositoryVersionRemoved *string
	tags *string
	version *string
	fields *[]string
	excludeFields *[]string
}

func (r ContentCollectionVersionsAPIContentAnsibleCollectionVersionsListRequest) IsHighest(isHighest bool) ContentCollectionVersionsAPIContentAnsibleCollectionVersionsListRequest {
	r.isHighest = &isHighest
	return r
}

// Number of results to return per page.
func (r ContentCollectionVersionsAPIContentAnsibleCollectionVersionsListRequest) Limit(limit int32) ContentCollectionVersionsAPIContentAnsibleCollectionVersionsListRequest {
	r.limit = &limit
	return r
}

func (r ContentCollectionVersionsAPIContentAnsibleCollectionVersionsListRequest) Name(name string) ContentCollectionVersionsAPIContentAnsibleCollectionVersionsListRequest {
	r.name = &name
	return r
}

func (r ContentCollectionVersionsAPIContentAnsibleCollectionVersionsListRequest) Namespace(namespace string) ContentCollectionVersionsAPIContentAnsibleCollectionVersionsListRequest {
	r.namespace = &namespace
	return r
}

// The initial index from which to return the results.
func (r ContentCollectionVersionsAPIContentAnsibleCollectionVersionsListRequest) Offset(offset int32) ContentCollectionVersionsAPIContentAnsibleCollectionVersionsListRequest {
	r.offset = &offset
	return r
}

// Ordering  * &#x60;pulp_id&#x60; - Pulp id * &#x60;-pulp_id&#x60; - Pulp id (descending) * &#x60;pulp_created&#x60; - Pulp created * &#x60;-pulp_created&#x60; - Pulp created (descending) * &#x60;pulp_last_updated&#x60; - Pulp last updated * &#x60;-pulp_last_updated&#x60; - Pulp last updated (descending) * &#x60;pulp_type&#x60; - Pulp type * &#x60;-pulp_type&#x60; - Pulp type (descending) * &#x60;upstream_id&#x60; - Upstream id * &#x60;-upstream_id&#x60; - Upstream id (descending) * &#x60;timestamp_of_interest&#x60; - Timestamp of interest * &#x60;-timestamp_of_interest&#x60; - Timestamp of interest (descending) * &#x60;authors&#x60; - Authors * &#x60;-authors&#x60; - Authors (descending) * &#x60;contents&#x60; - Contents * &#x60;-contents&#x60; - Contents (descending) * &#x60;dependencies&#x60; - Dependencies * &#x60;-dependencies&#x60; - Dependencies (descending) * &#x60;description&#x60; - Description * &#x60;-description&#x60; - Description (descending) * &#x60;docs_blob&#x60; - Docs blob * &#x60;-docs_blob&#x60; - Docs blob (descending) * &#x60;manifest&#x60; - Manifest * &#x60;-manifest&#x60; - Manifest (descending) * &#x60;files&#x60; - Files * &#x60;-files&#x60; - Files (descending) * &#x60;documentation&#x60; - Documentation * &#x60;-documentation&#x60; - Documentation (descending) * &#x60;homepage&#x60; - Homepage * &#x60;-homepage&#x60; - Homepage (descending) * &#x60;issues&#x60; - Issues * &#x60;-issues&#x60; - Issues (descending) * &#x60;license&#x60; - License * &#x60;-license&#x60; - License (descending) * &#x60;name&#x60; - Name * &#x60;-name&#x60; - Name (descending) * &#x60;namespace&#x60; - Namespace * &#x60;-namespace&#x60; - Namespace (descending) * &#x60;repository&#x60; - Repository * &#x60;-repository&#x60; - Repository (descending) * &#x60;version&#x60; - Version * &#x60;-version&#x60; - Version (descending) * &#x60;requires_ansible&#x60; - Requires ansible * &#x60;-requires_ansible&#x60; - Requires ansible (descending) * &#x60;is_highest&#x60; - Is highest * &#x60;-is_highest&#x60; - Is highest (descending) * &#x60;search_vector&#x60; - Search vector * &#x60;-search_vector&#x60; - Search vector (descending) * &#x60;pk&#x60; - Pk * &#x60;-pk&#x60; - Pk (descending)
func (r ContentCollectionVersionsAPIContentAnsibleCollectionVersionsListRequest) Ordering(ordering []string) ContentCollectionVersionsAPIContentAnsibleCollectionVersionsListRequest {
	r.ordering = &ordering
	return r
}

// Multiple values may be separated by commas.
func (r ContentCollectionVersionsAPIContentAnsibleCollectionVersionsListRequest) PulpHrefIn(pulpHrefIn []string) ContentCollectionVersionsAPIContentAnsibleCollectionVersionsListRequest {
	r.pulpHrefIn = &pulpHrefIn
	return r
}

// Multiple values may be separated by commas.
func (r ContentCollectionVersionsAPIContentAnsibleCollectionVersionsListRequest) PulpIdIn(pulpIdIn []string) ContentCollectionVersionsAPIContentAnsibleCollectionVersionsListRequest {
	r.pulpIdIn = &pulpIdIn
	return r
}

func (r ContentCollectionVersionsAPIContentAnsibleCollectionVersionsListRequest) Q(q string) ContentCollectionVersionsAPIContentAnsibleCollectionVersionsListRequest {
	r.q = &q
	return r
}

// Repository Version referenced by HREF
func (r ContentCollectionVersionsAPIContentAnsibleCollectionVersionsListRequest) RepositoryVersion(repositoryVersion string) ContentCollectionVersionsAPIContentAnsibleCollectionVersionsListRequest {
	r.repositoryVersion = &repositoryVersion
	return r
}

// Repository Version referenced by HREF
func (r ContentCollectionVersionsAPIContentAnsibleCollectionVersionsListRequest) RepositoryVersionAdded(repositoryVersionAdded string) ContentCollectionVersionsAPIContentAnsibleCollectionVersionsListRequest {
	r.repositoryVersionAdded = &repositoryVersionAdded
	return r
}

// Repository Version referenced by HREF
func (r ContentCollectionVersionsAPIContentAnsibleCollectionVersionsListRequest) RepositoryVersionRemoved(repositoryVersionRemoved string) ContentCollectionVersionsAPIContentAnsibleCollectionVersionsListRequest {
	r.repositoryVersionRemoved = &repositoryVersionRemoved
	return r
}

// Filter by comma separate list of tags that must all be matched
func (r ContentCollectionVersionsAPIContentAnsibleCollectionVersionsListRequest) Tags(tags string) ContentCollectionVersionsAPIContentAnsibleCollectionVersionsListRequest {
	r.tags = &tags
	return r
}

// Filter results where version matches value
func (r ContentCollectionVersionsAPIContentAnsibleCollectionVersionsListRequest) Version(version string) ContentCollectionVersionsAPIContentAnsibleCollectionVersionsListRequest {
	r.version = &version
	return r
}

// A list of fields to include in the response.
func (r ContentCollectionVersionsAPIContentAnsibleCollectionVersionsListRequest) Fields(fields []string) ContentCollectionVersionsAPIContentAnsibleCollectionVersionsListRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r ContentCollectionVersionsAPIContentAnsibleCollectionVersionsListRequest) ExcludeFields(excludeFields []string) ContentCollectionVersionsAPIContentAnsibleCollectionVersionsListRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r ContentCollectionVersionsAPIContentAnsibleCollectionVersionsListRequest) Execute() (*PaginatedansibleCollectionVersionResponseList, *http.Response, error) {
	return r.ApiService.ContentAnsibleCollectionVersionsListExecute(r)
}

/*
ContentAnsibleCollectionVersionsList List collection versions

ViewSet for Ansible Collection.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ContentCollectionVersionsAPIContentAnsibleCollectionVersionsListRequest
*/
func (a *ContentCollectionVersionsAPIService) ContentAnsibleCollectionVersionsList(ctx context.Context) ContentCollectionVersionsAPIContentAnsibleCollectionVersionsListRequest {
	return ContentCollectionVersionsAPIContentAnsibleCollectionVersionsListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PaginatedansibleCollectionVersionResponseList
func (a *ContentCollectionVersionsAPIService) ContentAnsibleCollectionVersionsListExecute(r ContentCollectionVersionsAPIContentAnsibleCollectionVersionsListRequest) (*PaginatedansibleCollectionVersionResponseList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedansibleCollectionVersionResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentCollectionVersionsAPIService.ContentAnsibleCollectionVersionsList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp/api/v3/content/ansible/collection_versions/"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.isHighest != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "is_highest", r.isHighest, "")
	}
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
	if r.q != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "q", r.q, "")
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
	if r.tags != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "tags", r.tags, "")
	}
	if r.version != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "version", r.version, "")
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

type ContentCollectionVersionsAPIContentAnsibleCollectionVersionsReadRequest struct {
	ctx context.Context
	ApiService *ContentCollectionVersionsAPIService
	ansibleCollectionVersionHref string
	fields *[]string
	excludeFields *[]string
}

// A list of fields to include in the response.
func (r ContentCollectionVersionsAPIContentAnsibleCollectionVersionsReadRequest) Fields(fields []string) ContentCollectionVersionsAPIContentAnsibleCollectionVersionsReadRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r ContentCollectionVersionsAPIContentAnsibleCollectionVersionsReadRequest) ExcludeFields(excludeFields []string) ContentCollectionVersionsAPIContentAnsibleCollectionVersionsReadRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r ContentCollectionVersionsAPIContentAnsibleCollectionVersionsReadRequest) Execute() (*AnsibleCollectionVersionResponse, *http.Response, error) {
	return r.ApiService.ContentAnsibleCollectionVersionsReadExecute(r)
}

/*
ContentAnsibleCollectionVersionsRead Inspect a collection version

ViewSet for Ansible Collection.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ansibleCollectionVersionHref
 @return ContentCollectionVersionsAPIContentAnsibleCollectionVersionsReadRequest
*/
func (a *ContentCollectionVersionsAPIService) ContentAnsibleCollectionVersionsRead(ctx context.Context, ansibleCollectionVersionHref string) ContentCollectionVersionsAPIContentAnsibleCollectionVersionsReadRequest {
	return ContentCollectionVersionsAPIContentAnsibleCollectionVersionsReadRequest{
		ApiService: a,
		ctx: ctx,
		ansibleCollectionVersionHref: ansibleCollectionVersionHref,
	}
}

// Execute executes the request
//  @return AnsibleCollectionVersionResponse
func (a *ContentCollectionVersionsAPIService) ContentAnsibleCollectionVersionsReadExecute(r ContentCollectionVersionsAPIContentAnsibleCollectionVersionsReadRequest) (*AnsibleCollectionVersionResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AnsibleCollectionVersionResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentCollectionVersionsAPIService.ContentAnsibleCollectionVersionsRead")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{ansible_collection_version_href}"
	localVarPath = strings.Replace(localVarPath, "{"+"ansible_collection_version_href"+"}", url.PathEscape(parameterValueToString(r.ansibleCollectionVersionHref, "ansibleCollectionVersionHref")), -1)
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
