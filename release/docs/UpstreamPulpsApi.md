# \UpstreamPulpsApi

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UpstreamPulpsCreate**](UpstreamPulpsApi.md#UpstreamPulpsCreate) | **Post** /pulp/api/v3/upstream-pulps/ | Create an upstream pulp
[**UpstreamPulpsDelete**](UpstreamPulpsApi.md#UpstreamPulpsDelete) | **Delete** /{upstream_pulp_href} | Delete an upstream pulp
[**UpstreamPulpsList**](UpstreamPulpsApi.md#UpstreamPulpsList) | **Get** /pulp/api/v3/upstream-pulps/ | List upstream pulps
[**UpstreamPulpsPartialUpdate**](UpstreamPulpsApi.md#UpstreamPulpsPartialUpdate) | **Patch** /{upstream_pulp_href} | Update an upstream pulp
[**UpstreamPulpsRead**](UpstreamPulpsApi.md#UpstreamPulpsRead) | **Get** /{upstream_pulp_href} | Inspect an upstream pulp
[**UpstreamPulpsReplicate**](UpstreamPulpsApi.md#UpstreamPulpsReplicate) | **Post** /{upstream_pulp_href}replicate/ | Replicate
[**UpstreamPulpsUpdate**](UpstreamPulpsApi.md#UpstreamPulpsUpdate) | **Put** /{upstream_pulp_href} | Update an upstream pulp



## UpstreamPulpsCreate

> UpstreamPulpResponse UpstreamPulpsCreate(ctx).UpstreamPulp(upstreamPulp).Execute()

Create an upstream pulp



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/content-services/zest/v3/release--package-name&#x3D;zest"
)

func main() {
    upstreamPulp := *openapiclient.NewUpstreamPulp("Name_example", "BaseUrl_example", "ApiRoot_example") // UpstreamPulp | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UpstreamPulpsApi.UpstreamPulpsCreate(context.Background()).UpstreamPulp(upstreamPulp).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UpstreamPulpsApi.UpstreamPulpsCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpstreamPulpsCreate`: UpstreamPulpResponse
    fmt.Fprintf(os.Stdout, "Response from `UpstreamPulpsApi.UpstreamPulpsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpstreamPulpsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **upstreamPulp** | [**UpstreamPulp**](UpstreamPulp.md) |  | 

### Return type

[**UpstreamPulpResponse**](UpstreamPulpResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpstreamPulpsDelete

> UpstreamPulpsDelete(ctx, upstreamPulpHref).Execute()

Delete an upstream pulp



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/content-services/zest/v3/release--package-name&#x3D;zest"
)

func main() {
    upstreamPulpHref := "upstreamPulpHref_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.UpstreamPulpsApi.UpstreamPulpsDelete(context.Background(), upstreamPulpHref).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UpstreamPulpsApi.UpstreamPulpsDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**upstreamPulpHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpstreamPulpsDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpstreamPulpsList

> PaginatedUpstreamPulpResponseList UpstreamPulpsList(ctx).Limit(limit).Offset(offset).Fields(fields).ExcludeFields(excludeFields).Execute()

List upstream pulps



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/content-services/zest/v3/release--package-name&#x3D;zest"
)

func main() {
    limit := int32(56) // int32 | Number of results to return per page. (optional)
    offset := int32(56) // int32 | The initial index from which to return the results. (optional)
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UpstreamPulpsApi.UpstreamPulpsList(context.Background()).Limit(limit).Offset(offset).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UpstreamPulpsApi.UpstreamPulpsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpstreamPulpsList`: PaginatedUpstreamPulpResponseList
    fmt.Fprintf(os.Stdout, "Response from `UpstreamPulpsApi.UpstreamPulpsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpstreamPulpsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Number of results to return per page. | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**PaginatedUpstreamPulpResponseList**](PaginatedUpstreamPulpResponseList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpstreamPulpsPartialUpdate

> UpstreamPulpResponse UpstreamPulpsPartialUpdate(ctx, upstreamPulpHref).PatchedUpstreamPulp(patchedUpstreamPulp).Execute()

Update an upstream pulp



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/content-services/zest/v3/release--package-name&#x3D;zest"
)

func main() {
    upstreamPulpHref := "upstreamPulpHref_example" // string | 
    patchedUpstreamPulp := *openapiclient.NewPatchedUpstreamPulp() // PatchedUpstreamPulp | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UpstreamPulpsApi.UpstreamPulpsPartialUpdate(context.Background(), upstreamPulpHref).PatchedUpstreamPulp(patchedUpstreamPulp).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UpstreamPulpsApi.UpstreamPulpsPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpstreamPulpsPartialUpdate`: UpstreamPulpResponse
    fmt.Fprintf(os.Stdout, "Response from `UpstreamPulpsApi.UpstreamPulpsPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**upstreamPulpHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpstreamPulpsPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedUpstreamPulp** | [**PatchedUpstreamPulp**](PatchedUpstreamPulp.md) |  | 

### Return type

[**UpstreamPulpResponse**](UpstreamPulpResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpstreamPulpsRead

> UpstreamPulpResponse UpstreamPulpsRead(ctx, upstreamPulpHref).Fields(fields).ExcludeFields(excludeFields).Execute()

Inspect an upstream pulp



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/content-services/zest/v3/release--package-name&#x3D;zest"
)

func main() {
    upstreamPulpHref := "upstreamPulpHref_example" // string | 
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UpstreamPulpsApi.UpstreamPulpsRead(context.Background(), upstreamPulpHref).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UpstreamPulpsApi.UpstreamPulpsRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpstreamPulpsRead`: UpstreamPulpResponse
    fmt.Fprintf(os.Stdout, "Response from `UpstreamPulpsApi.UpstreamPulpsRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**upstreamPulpHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpstreamPulpsReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**UpstreamPulpResponse**](UpstreamPulpResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpstreamPulpsReplicate

> AsyncOperationResponse UpstreamPulpsReplicate(ctx, upstreamPulpHref).UpstreamPulp(upstreamPulp).Execute()

Replicate



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/content-services/zest/v3/release--package-name&#x3D;zest"
)

func main() {
    upstreamPulpHref := "upstreamPulpHref_example" // string | 
    upstreamPulp := *openapiclient.NewUpstreamPulp("Name_example", "BaseUrl_example", "ApiRoot_example") // UpstreamPulp | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UpstreamPulpsApi.UpstreamPulpsReplicate(context.Background(), upstreamPulpHref).UpstreamPulp(upstreamPulp).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UpstreamPulpsApi.UpstreamPulpsReplicate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpstreamPulpsReplicate`: AsyncOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `UpstreamPulpsApi.UpstreamPulpsReplicate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**upstreamPulpHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpstreamPulpsReplicateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **upstreamPulp** | [**UpstreamPulp**](UpstreamPulp.md) |  | 

### Return type

[**AsyncOperationResponse**](AsyncOperationResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpstreamPulpsUpdate

> UpstreamPulpResponse UpstreamPulpsUpdate(ctx, upstreamPulpHref).UpstreamPulp(upstreamPulp).Execute()

Update an upstream pulp



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/content-services/zest/v3/release--package-name&#x3D;zest"
)

func main() {
    upstreamPulpHref := "upstreamPulpHref_example" // string | 
    upstreamPulp := *openapiclient.NewUpstreamPulp("Name_example", "BaseUrl_example", "ApiRoot_example") // UpstreamPulp | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UpstreamPulpsApi.UpstreamPulpsUpdate(context.Background(), upstreamPulpHref).UpstreamPulp(upstreamPulp).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UpstreamPulpsApi.UpstreamPulpsUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpstreamPulpsUpdate`: UpstreamPulpResponse
    fmt.Fprintf(os.Stdout, "Response from `UpstreamPulpsApi.UpstreamPulpsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**upstreamPulpHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpstreamPulpsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **upstreamPulp** | [**UpstreamPulp**](UpstreamPulp.md) |  | 

### Return type

[**UpstreamPulpResponse**](UpstreamPulpResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

