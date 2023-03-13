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
	"time"
	"reflect"
)


// PublicationsApiService PublicationsApi service
type PublicationsApiService service

type PublicationsApiPublicationsListRequest struct {
	ctx context.Context
	ApiService *PublicationsApiService
	content *string
	contentIn *string
	limit *int32
	offset *int32
	ordering *[]string
	pulpCreated *time.Time
	pulpCreatedGt *time.Time
	pulpCreatedGte *time.Time
	pulpCreatedLt *time.Time
	pulpCreatedLte *time.Time
	pulpCreatedRange *[]time.Time
	repository *string
	repositoryVersion *string
	fields *[]string
	excludeFields *[]string
}

// Content Unit referenced by HREF
func (r PublicationsApiPublicationsListRequest) Content(content string) PublicationsApiPublicationsListRequest {
	r.content = &content
	return r
}

// Content Unit referenced by HREF
func (r PublicationsApiPublicationsListRequest) ContentIn(contentIn string) PublicationsApiPublicationsListRequest {
	r.contentIn = &contentIn
	return r
}

// Number of results to return per page.
func (r PublicationsApiPublicationsListRequest) Limit(limit int32) PublicationsApiPublicationsListRequest {
	r.limit = &limit
	return r
}

// The initial index from which to return the results.
func (r PublicationsApiPublicationsListRequest) Offset(offset int32) PublicationsApiPublicationsListRequest {
	r.offset = &offset
	return r
}

// Ordering
func (r PublicationsApiPublicationsListRequest) Ordering(ordering []string) PublicationsApiPublicationsListRequest {
	r.ordering = &ordering
	return r
}

// Filter results where pulp_created matches value
func (r PublicationsApiPublicationsListRequest) PulpCreated(pulpCreated time.Time) PublicationsApiPublicationsListRequest {
	r.pulpCreated = &pulpCreated
	return r
}

// Filter results where pulp_created is greater than value
func (r PublicationsApiPublicationsListRequest) PulpCreatedGt(pulpCreatedGt time.Time) PublicationsApiPublicationsListRequest {
	r.pulpCreatedGt = &pulpCreatedGt
	return r
}

// Filter results where pulp_created is greater than or equal to value
func (r PublicationsApiPublicationsListRequest) PulpCreatedGte(pulpCreatedGte time.Time) PublicationsApiPublicationsListRequest {
	r.pulpCreatedGte = &pulpCreatedGte
	return r
}

// Filter results where pulp_created is less than value
func (r PublicationsApiPublicationsListRequest) PulpCreatedLt(pulpCreatedLt time.Time) PublicationsApiPublicationsListRequest {
	r.pulpCreatedLt = &pulpCreatedLt
	return r
}

// Filter results where pulp_created is less than or equal to value
func (r PublicationsApiPublicationsListRequest) PulpCreatedLte(pulpCreatedLte time.Time) PublicationsApiPublicationsListRequest {
	r.pulpCreatedLte = &pulpCreatedLte
	return r
}

// Filter results where pulp_created is between two comma separated values
func (r PublicationsApiPublicationsListRequest) PulpCreatedRange(pulpCreatedRange []time.Time) PublicationsApiPublicationsListRequest {
	r.pulpCreatedRange = &pulpCreatedRange
	return r
}

// Repository referenced by HREF
func (r PublicationsApiPublicationsListRequest) Repository(repository string) PublicationsApiPublicationsListRequest {
	r.repository = &repository
	return r
}

// Repository Version referenced by HREF
func (r PublicationsApiPublicationsListRequest) RepositoryVersion(repositoryVersion string) PublicationsApiPublicationsListRequest {
	r.repositoryVersion = &repositoryVersion
	return r
}

// A list of fields to include in the response.
func (r PublicationsApiPublicationsListRequest) Fields(fields []string) PublicationsApiPublicationsListRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r PublicationsApiPublicationsListRequest) ExcludeFields(excludeFields []string) PublicationsApiPublicationsListRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r PublicationsApiPublicationsListRequest) Execute() (*PaginatedPublicationResponseList, *http.Response, error) {
	return r.ApiService.PublicationsListExecute(r)
}

/*
PublicationsList List publications

A customized named ModelViewSet that knows how to register itself with the Pulp API router.

This viewset is discoverable by its name.
"Normal" Django Models and Master/Detail models are supported by the ``register_with`` method.

Attributes:
    lookup_field (str): The name of the field by which an object should be looked up, in
        addition to any parent lookups if this ViewSet is nested. Defaults to 'pk'
    endpoint_name (str): The name of the final path segment that should identify the ViewSet's
        collection endpoint.
    nest_prefix (str): Optional prefix under which this ViewSet should be nested. This must
        correspond to the "parent_prefix" of a router with rest_framework_nested.NestedMixin.
        None indicates this ViewSet should not be nested.
    parent_lookup_kwargs (dict): Optional mapping of key names that would appear in self.kwargs
        to django model filter expressions that can be used with the corresponding value from
        self.kwargs, used only by a nested ViewSet to filter based on the parent object's
        identity.
    schema (DefaultSchema): The schema class to use by default in a viewset.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return PublicationsApiPublicationsListRequest
*/
func (a *PublicationsApiService) PublicationsList(ctx context.Context) PublicationsApiPublicationsListRequest {
	return PublicationsApiPublicationsListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PaginatedPublicationResponseList
func (a *PublicationsApiService) PublicationsListExecute(r PublicationsApiPublicationsListRequest) (*PaginatedPublicationResponseList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedPublicationResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PublicationsApiService.PublicationsList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp/api/v3/publications/"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.content != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "content", r.content, "")
	}
	if r.contentIn != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "content__in", r.contentIn, "")
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
	if r.pulpCreated != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pulp_created", r.pulpCreated, "")
	}
	if r.pulpCreatedGt != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pulp_created__gt", r.pulpCreatedGt, "")
	}
	if r.pulpCreatedGte != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pulp_created__gte", r.pulpCreatedGte, "")
	}
	if r.pulpCreatedLt != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pulp_created__lt", r.pulpCreatedLt, "")
	}
	if r.pulpCreatedLte != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pulp_created__lte", r.pulpCreatedLte, "")
	}
	if r.pulpCreatedRange != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pulp_created__range", r.pulpCreatedRange, "csv")
	}
	if r.repository != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "repository", r.repository, "")
	}
	if r.repositoryVersion != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "repository_version", r.repositoryVersion, "")
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
