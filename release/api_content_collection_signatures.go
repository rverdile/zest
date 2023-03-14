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


// ContentCollectionSignaturesApiService ContentCollectionSignaturesApi service
type ContentCollectionSignaturesApiService service

type ContentCollectionSignaturesApiContentAnsibleCollectionSignaturesCreateRequest struct {
	ctx context.Context
	ApiService *ContentCollectionSignaturesApiService
	file *os.File
	signedCollection *string
	repository *string
}

// An uploaded file that may be turned into the artifact of the content unit.
func (r ContentCollectionSignaturesApiContentAnsibleCollectionSignaturesCreateRequest) File(file *os.File) ContentCollectionSignaturesApiContentAnsibleCollectionSignaturesCreateRequest {
	r.file = file
	return r
}

// The content this signature is pointing to.
func (r ContentCollectionSignaturesApiContentAnsibleCollectionSignaturesCreateRequest) SignedCollection(signedCollection string) ContentCollectionSignaturesApiContentAnsibleCollectionSignaturesCreateRequest {
	r.signedCollection = &signedCollection
	return r
}

// A URI of a repository the new content unit should be associated with.
func (r ContentCollectionSignaturesApiContentAnsibleCollectionSignaturesCreateRequest) Repository(repository string) ContentCollectionSignaturesApiContentAnsibleCollectionSignaturesCreateRequest {
	r.repository = &repository
	return r
}

func (r ContentCollectionSignaturesApiContentAnsibleCollectionSignaturesCreateRequest) Execute() (*AsyncOperationResponse, *http.Response, error) {
	return r.ApiService.ContentAnsibleCollectionSignaturesCreateExecute(r)
}

/*
ContentAnsibleCollectionSignaturesCreate Create a collection version signature

Trigger an asynchronous task to create content,optionally create new repository version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ContentCollectionSignaturesApiContentAnsibleCollectionSignaturesCreateRequest
*/
func (a *ContentCollectionSignaturesApiService) ContentAnsibleCollectionSignaturesCreate(ctx context.Context) ContentCollectionSignaturesApiContentAnsibleCollectionSignaturesCreateRequest {
	return ContentCollectionSignaturesApiContentAnsibleCollectionSignaturesCreateRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AsyncOperationResponse
func (a *ContentCollectionSignaturesApiService) ContentAnsibleCollectionSignaturesCreateExecute(r ContentCollectionSignaturesApiContentAnsibleCollectionSignaturesCreateRequest) (*AsyncOperationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AsyncOperationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentCollectionSignaturesApiService.ContentAnsibleCollectionSignaturesCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp/api/v3/content/ansible/collection_signatures/"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.file == nil {
		return localVarReturnValue, nil, reportError("file is required and must be specified")
	}
	if r.signedCollection == nil {
		return localVarReturnValue, nil, reportError("signedCollection is required and must be specified")
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
	if r.repository != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "repository", r.repository, "")
	}
	parameterAddToHeaderOrQuery(localVarFormParams, "signed_collection", r.signedCollection, "")
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

type ContentCollectionSignaturesApiContentAnsibleCollectionSignaturesListRequest struct {
	ctx context.Context
	ApiService *ContentCollectionSignaturesApiService
	limit *int32
	offset *int32
	ordering *[]string
	pubkeyFingerprint *string
	pubkeyFingerprintIn *[]string
	repositoryVersion *string
	repositoryVersionAdded *string
	repositoryVersionRemoved *string
	signedCollection *string
	signingService *string
	fields *[]string
	excludeFields *[]string
}

// Number of results to return per page.
func (r ContentCollectionSignaturesApiContentAnsibleCollectionSignaturesListRequest) Limit(limit int32) ContentCollectionSignaturesApiContentAnsibleCollectionSignaturesListRequest {
	r.limit = &limit
	return r
}

// The initial index from which to return the results.
func (r ContentCollectionSignaturesApiContentAnsibleCollectionSignaturesListRequest) Offset(offset int32) ContentCollectionSignaturesApiContentAnsibleCollectionSignaturesListRequest {
	r.offset = &offset
	return r
}

// Ordering
func (r ContentCollectionSignaturesApiContentAnsibleCollectionSignaturesListRequest) Ordering(ordering []string) ContentCollectionSignaturesApiContentAnsibleCollectionSignaturesListRequest {
	r.ordering = &ordering
	return r
}

// Filter results where pubkey_fingerprint matches value
func (r ContentCollectionSignaturesApiContentAnsibleCollectionSignaturesListRequest) PubkeyFingerprint(pubkeyFingerprint string) ContentCollectionSignaturesApiContentAnsibleCollectionSignaturesListRequest {
	r.pubkeyFingerprint = &pubkeyFingerprint
	return r
}

// Filter results where pubkey_fingerprint is in a comma-separated list of values
func (r ContentCollectionSignaturesApiContentAnsibleCollectionSignaturesListRequest) PubkeyFingerprintIn(pubkeyFingerprintIn []string) ContentCollectionSignaturesApiContentAnsibleCollectionSignaturesListRequest {
	r.pubkeyFingerprintIn = &pubkeyFingerprintIn
	return r
}

// Repository Version referenced by HREF
func (r ContentCollectionSignaturesApiContentAnsibleCollectionSignaturesListRequest) RepositoryVersion(repositoryVersion string) ContentCollectionSignaturesApiContentAnsibleCollectionSignaturesListRequest {
	r.repositoryVersion = &repositoryVersion
	return r
}

// Repository Version referenced by HREF
func (r ContentCollectionSignaturesApiContentAnsibleCollectionSignaturesListRequest) RepositoryVersionAdded(repositoryVersionAdded string) ContentCollectionSignaturesApiContentAnsibleCollectionSignaturesListRequest {
	r.repositoryVersionAdded = &repositoryVersionAdded
	return r
}

// Repository Version referenced by HREF
func (r ContentCollectionSignaturesApiContentAnsibleCollectionSignaturesListRequest) RepositoryVersionRemoved(repositoryVersionRemoved string) ContentCollectionSignaturesApiContentAnsibleCollectionSignaturesListRequest {
	r.repositoryVersionRemoved = &repositoryVersionRemoved
	return r
}

// Filter signatures for collection version
func (r ContentCollectionSignaturesApiContentAnsibleCollectionSignaturesListRequest) SignedCollection(signedCollection string) ContentCollectionSignaturesApiContentAnsibleCollectionSignaturesListRequest {
	r.signedCollection = &signedCollection
	return r
}

// Filter signatures produced by signature service
func (r ContentCollectionSignaturesApiContentAnsibleCollectionSignaturesListRequest) SigningService(signingService string) ContentCollectionSignaturesApiContentAnsibleCollectionSignaturesListRequest {
	r.signingService = &signingService
	return r
}

// A list of fields to include in the response.
func (r ContentCollectionSignaturesApiContentAnsibleCollectionSignaturesListRequest) Fields(fields []string) ContentCollectionSignaturesApiContentAnsibleCollectionSignaturesListRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r ContentCollectionSignaturesApiContentAnsibleCollectionSignaturesListRequest) ExcludeFields(excludeFields []string) ContentCollectionSignaturesApiContentAnsibleCollectionSignaturesListRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r ContentCollectionSignaturesApiContentAnsibleCollectionSignaturesListRequest) Execute() (*PaginatedansibleCollectionVersionSignatureResponseList, *http.Response, error) {
	return r.ApiService.ContentAnsibleCollectionSignaturesListExecute(r)
}

/*
ContentAnsibleCollectionSignaturesList List collection version signatures

ViewSet for looking at signature objects for CollectionVersion content.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ContentCollectionSignaturesApiContentAnsibleCollectionSignaturesListRequest
*/
func (a *ContentCollectionSignaturesApiService) ContentAnsibleCollectionSignaturesList(ctx context.Context) ContentCollectionSignaturesApiContentAnsibleCollectionSignaturesListRequest {
	return ContentCollectionSignaturesApiContentAnsibleCollectionSignaturesListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PaginatedansibleCollectionVersionSignatureResponseList
func (a *ContentCollectionSignaturesApiService) ContentAnsibleCollectionSignaturesListExecute(r ContentCollectionSignaturesApiContentAnsibleCollectionSignaturesListRequest) (*PaginatedansibleCollectionVersionSignatureResponseList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedansibleCollectionVersionSignatureResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentCollectionSignaturesApiService.ContentAnsibleCollectionSignaturesList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp/api/v3/content/ansible/collection_signatures/"
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
	if r.pubkeyFingerprint != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pubkey_fingerprint", r.pubkeyFingerprint, "")
	}
	if r.pubkeyFingerprintIn != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pubkey_fingerprint__in", r.pubkeyFingerprintIn, "csv")
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
	if r.signedCollection != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "signed_collection", r.signedCollection, "")
	}
	if r.signingService != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "signing_service", r.signingService, "")
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

type ContentCollectionSignaturesApiContentAnsibleCollectionSignaturesReadRequest struct {
	ctx context.Context
	ApiService *ContentCollectionSignaturesApiService
	ansibleCollectionVersionSignatureHref string
	fields *[]string
	excludeFields *[]string
}

// A list of fields to include in the response.
func (r ContentCollectionSignaturesApiContentAnsibleCollectionSignaturesReadRequest) Fields(fields []string) ContentCollectionSignaturesApiContentAnsibleCollectionSignaturesReadRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r ContentCollectionSignaturesApiContentAnsibleCollectionSignaturesReadRequest) ExcludeFields(excludeFields []string) ContentCollectionSignaturesApiContentAnsibleCollectionSignaturesReadRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r ContentCollectionSignaturesApiContentAnsibleCollectionSignaturesReadRequest) Execute() (*AnsibleCollectionVersionSignatureResponse, *http.Response, error) {
	return r.ApiService.ContentAnsibleCollectionSignaturesReadExecute(r)
}

/*
ContentAnsibleCollectionSignaturesRead Inspect a collection version signature

ViewSet for looking at signature objects for CollectionVersion content.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ansibleCollectionVersionSignatureHref
 @return ContentCollectionSignaturesApiContentAnsibleCollectionSignaturesReadRequest
*/
func (a *ContentCollectionSignaturesApiService) ContentAnsibleCollectionSignaturesRead(ctx context.Context, ansibleCollectionVersionSignatureHref string) ContentCollectionSignaturesApiContentAnsibleCollectionSignaturesReadRequest {
	return ContentCollectionSignaturesApiContentAnsibleCollectionSignaturesReadRequest{
		ApiService: a,
		ctx: ctx,
		ansibleCollectionVersionSignatureHref: ansibleCollectionVersionSignatureHref,
	}
}

// Execute executes the request
//  @return AnsibleCollectionVersionSignatureResponse
func (a *ContentCollectionSignaturesApiService) ContentAnsibleCollectionSignaturesReadExecute(r ContentCollectionSignaturesApiContentAnsibleCollectionSignaturesReadRequest) (*AnsibleCollectionVersionSignatureResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AnsibleCollectionVersionSignatureResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentCollectionSignaturesApiService.ContentAnsibleCollectionSignaturesRead")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{ansible_collection_version_signature_href}"
	localVarPath = strings.Replace(localVarPath, "{"+"ansible_collection_version_signature_href"+"}", url.PathEscape(parameterValueToString(r.ansibleCollectionVersionSignatureHref, "ansibleCollectionVersionSignatureHref")), -1)
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
