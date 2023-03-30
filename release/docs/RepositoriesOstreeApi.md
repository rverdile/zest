# \RepositoriesOstreeApi

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RepositoriesOstreeOstreeCreate**](RepositoriesOstreeApi.md#RepositoriesOstreeOstreeCreate) | **Post** /pulp/api/v3/repositories/ostree/ostree/ | Create an ostree repository
[**RepositoriesOstreeOstreeDelete**](RepositoriesOstreeApi.md#RepositoriesOstreeOstreeDelete) | **Delete** /{ostree_ostree_repository_href} | Delete an ostree repository
[**RepositoriesOstreeOstreeImportAll**](RepositoriesOstreeApi.md#RepositoriesOstreeOstreeImportAll) | **Post** /{ostree_ostree_repository_href}import_all/ | Import refs and commits to a repository
[**RepositoriesOstreeOstreeImportCommits**](RepositoriesOstreeApi.md#RepositoriesOstreeOstreeImportCommits) | **Post** /{ostree_ostree_repository_href}import_commits/ | Append child commits to a repository
[**RepositoriesOstreeOstreeList**](RepositoriesOstreeApi.md#RepositoriesOstreeOstreeList) | **Get** /pulp/api/v3/repositories/ostree/ostree/ | List ostree repositorys
[**RepositoriesOstreeOstreeModify**](RepositoriesOstreeApi.md#RepositoriesOstreeOstreeModify) | **Post** /{ostree_ostree_repository_href}modify/ | Modify repository
[**RepositoriesOstreeOstreePartialUpdate**](RepositoriesOstreeApi.md#RepositoriesOstreeOstreePartialUpdate) | **Patch** /{ostree_ostree_repository_href} | Update an ostree repository
[**RepositoriesOstreeOstreeRead**](RepositoriesOstreeApi.md#RepositoriesOstreeOstreeRead) | **Get** /{ostree_ostree_repository_href} | Inspect an ostree repository
[**RepositoriesOstreeOstreeSync**](RepositoriesOstreeApi.md#RepositoriesOstreeOstreeSync) | **Post** /{ostree_ostree_repository_href}sync/ | Sync from remote
[**RepositoriesOstreeOstreeUpdate**](RepositoriesOstreeApi.md#RepositoriesOstreeOstreeUpdate) | **Put** /{ostree_ostree_repository_href} | Update an ostree repository



## RepositoriesOstreeOstreeCreate

> OstreeOstreeRepositoryResponse RepositoriesOstreeOstreeCreate(ctx).OstreeOstreeRepository(ostreeOstreeRepository).Execute()

Create an ostree repository



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/content-services/zest/release/v3"
)

func main() {
    ostreeOstreeRepository := *openapiclient.NewOstreeOstreeRepository("Name_example") // OstreeOstreeRepository | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RepositoriesOstreeApi.RepositoriesOstreeOstreeCreate(context.Background()).OstreeOstreeRepository(ostreeOstreeRepository).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesOstreeApi.RepositoriesOstreeOstreeCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesOstreeOstreeCreate`: OstreeOstreeRepositoryResponse
    fmt.Fprintf(os.Stdout, "Response from `RepositoriesOstreeApi.RepositoriesOstreeOstreeCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesOstreeOstreeCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ostreeOstreeRepository** | [**OstreeOstreeRepository**](OstreeOstreeRepository.md) |  | 

### Return type

[**OstreeOstreeRepositoryResponse**](OstreeOstreeRepositoryResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RepositoriesOstreeOstreeDelete

> AsyncOperationResponse RepositoriesOstreeOstreeDelete(ctx, ostreeOstreeRepositoryHref).Execute()

Delete an ostree repository



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/content-services/zest/release/v3"
)

func main() {
    ostreeOstreeRepositoryHref := "ostreeOstreeRepositoryHref_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RepositoriesOstreeApi.RepositoriesOstreeOstreeDelete(context.Background(), ostreeOstreeRepositoryHref).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesOstreeApi.RepositoriesOstreeOstreeDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesOstreeOstreeDelete`: AsyncOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `RepositoriesOstreeApi.RepositoriesOstreeOstreeDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ostreeOstreeRepositoryHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesOstreeOstreeDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AsyncOperationResponse**](AsyncOperationResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RepositoriesOstreeOstreeImportAll

> AsyncOperationResponse RepositoriesOstreeOstreeImportAll(ctx, ostreeOstreeRepositoryHref).OstreeImportAll(ostreeImportAll).Execute()

Import refs and commits to a repository



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/content-services/zest/release/v3"
)

func main() {
    ostreeOstreeRepositoryHref := "ostreeOstreeRepositoryHref_example" // string | 
    ostreeImportAll := *openapiclient.NewOstreeImportAll("Artifact_example", "RepositoryName_example") // OstreeImportAll | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RepositoriesOstreeApi.RepositoriesOstreeOstreeImportAll(context.Background(), ostreeOstreeRepositoryHref).OstreeImportAll(ostreeImportAll).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesOstreeApi.RepositoriesOstreeOstreeImportAll``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesOstreeOstreeImportAll`: AsyncOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `RepositoriesOstreeApi.RepositoriesOstreeOstreeImportAll`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ostreeOstreeRepositoryHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesOstreeOstreeImportAllRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ostreeImportAll** | [**OstreeImportAll**](OstreeImportAll.md) |  | 

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


## RepositoriesOstreeOstreeImportCommits

> AsyncOperationResponse RepositoriesOstreeOstreeImportCommits(ctx, ostreeOstreeRepositoryHref).OstreeImportCommitsToRef(ostreeImportCommitsToRef).Execute()

Append child commits to a repository



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/content-services/zest/release/v3"
)

func main() {
    ostreeOstreeRepositoryHref := "ostreeOstreeRepositoryHref_example" // string | 
    ostreeImportCommitsToRef := *openapiclient.NewOstreeImportCommitsToRef("Artifact_example", "RepositoryName_example", "Ref_example") // OstreeImportCommitsToRef | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RepositoriesOstreeApi.RepositoriesOstreeOstreeImportCommits(context.Background(), ostreeOstreeRepositoryHref).OstreeImportCommitsToRef(ostreeImportCommitsToRef).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesOstreeApi.RepositoriesOstreeOstreeImportCommits``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesOstreeOstreeImportCommits`: AsyncOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `RepositoriesOstreeApi.RepositoriesOstreeOstreeImportCommits`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ostreeOstreeRepositoryHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesOstreeOstreeImportCommitsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ostreeImportCommitsToRef** | [**OstreeImportCommitsToRef**](OstreeImportCommitsToRef.md) |  | 

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


## RepositoriesOstreeOstreeList

> PaginatedostreeOstreeRepositoryResponseList RepositoriesOstreeOstreeList(ctx).Limit(limit).Name(name).NameContains(nameContains).NameIcontains(nameIcontains).NameIn(nameIn).NameStartswith(nameStartswith).Offset(offset).Ordering(ordering).PulpLabelSelect(pulpLabelSelect).Remote(remote).RetainRepoVersions(retainRepoVersions).RetainRepoVersionsGt(retainRepoVersionsGt).RetainRepoVersionsGte(retainRepoVersionsGte).RetainRepoVersionsIsnull(retainRepoVersionsIsnull).RetainRepoVersionsLt(retainRepoVersionsLt).RetainRepoVersionsLte(retainRepoVersionsLte).RetainRepoVersionsNe(retainRepoVersionsNe).RetainRepoVersionsRange(retainRepoVersionsRange).Fields(fields).ExcludeFields(excludeFields).Execute()

List ostree repositorys



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/content-services/zest/release/v3"
)

func main() {
    limit := int32(56) // int32 | Number of results to return per page. (optional)
    name := "name_example" // string | Filter results where name matches value (optional)
    nameContains := "nameContains_example" // string | Filter results where name contains value (optional)
    nameIcontains := "nameIcontains_example" // string | Filter results where name contains value (optional)
    nameIn := []string{"Inner_example"} // []string | Filter results where name is in a comma-separated list of values (optional)
    nameStartswith := "nameStartswith_example" // string | Filter results where name starts with value (optional)
    offset := int32(56) // int32 | The initial index from which to return the results. (optional)
    ordering := []string{"Ordering_example"} // []string | Ordering (optional)
    pulpLabelSelect := "pulpLabelSelect_example" // string | Filter labels by search string (optional)
    remote := "remote_example" // string | Foreign Key referenced by HREF (optional)
    retainRepoVersions := int32(56) // int32 | Filter results where retain_repo_versions matches value (optional)
    retainRepoVersionsGt := int32(56) // int32 | Filter results where retain_repo_versions is greater than value (optional)
    retainRepoVersionsGte := int32(56) // int32 | Filter results where retain_repo_versions is greater than or equal to value (optional)
    retainRepoVersionsIsnull := true // bool | Filter results where retain_repo_versions has a null value (optional)
    retainRepoVersionsLt := int32(56) // int32 | Filter results where retain_repo_versions is less than value (optional)
    retainRepoVersionsLte := int32(56) // int32 | Filter results where retain_repo_versions is less than or equal to value (optional)
    retainRepoVersionsNe := int32(56) // int32 | Filter results where retain_repo_versions not equal to value (optional)
    retainRepoVersionsRange := []int32{int32(123)} // []int32 | Filter results where retain_repo_versions is between two comma separated values (optional)
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RepositoriesOstreeApi.RepositoriesOstreeOstreeList(context.Background()).Limit(limit).Name(name).NameContains(nameContains).NameIcontains(nameIcontains).NameIn(nameIn).NameStartswith(nameStartswith).Offset(offset).Ordering(ordering).PulpLabelSelect(pulpLabelSelect).Remote(remote).RetainRepoVersions(retainRepoVersions).RetainRepoVersionsGt(retainRepoVersionsGt).RetainRepoVersionsGte(retainRepoVersionsGte).RetainRepoVersionsIsnull(retainRepoVersionsIsnull).RetainRepoVersionsLt(retainRepoVersionsLt).RetainRepoVersionsLte(retainRepoVersionsLte).RetainRepoVersionsNe(retainRepoVersionsNe).RetainRepoVersionsRange(retainRepoVersionsRange).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesOstreeApi.RepositoriesOstreeOstreeList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesOstreeOstreeList`: PaginatedostreeOstreeRepositoryResponseList
    fmt.Fprintf(os.Stdout, "Response from `RepositoriesOstreeApi.RepositoriesOstreeOstreeList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesOstreeOstreeListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Number of results to return per page. | 
 **name** | **string** | Filter results where name matches value | 
 **nameContains** | **string** | Filter results where name contains value | 
 **nameIcontains** | **string** | Filter results where name contains value | 
 **nameIn** | **[]string** | Filter results where name is in a comma-separated list of values | 
 **nameStartswith** | **string** | Filter results where name starts with value | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **ordering** | **[]string** | Ordering | 
 **pulpLabelSelect** | **string** | Filter labels by search string | 
 **remote** | **string** | Foreign Key referenced by HREF | 
 **retainRepoVersions** | **int32** | Filter results where retain_repo_versions matches value | 
 **retainRepoVersionsGt** | **int32** | Filter results where retain_repo_versions is greater than value | 
 **retainRepoVersionsGte** | **int32** | Filter results where retain_repo_versions is greater than or equal to value | 
 **retainRepoVersionsIsnull** | **bool** | Filter results where retain_repo_versions has a null value | 
 **retainRepoVersionsLt** | **int32** | Filter results where retain_repo_versions is less than value | 
 **retainRepoVersionsLte** | **int32** | Filter results where retain_repo_versions is less than or equal to value | 
 **retainRepoVersionsNe** | **int32** | Filter results where retain_repo_versions not equal to value | 
 **retainRepoVersionsRange** | **[]int32** | Filter results where retain_repo_versions is between two comma separated values | 
 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**PaginatedostreeOstreeRepositoryResponseList**](PaginatedostreeOstreeRepositoryResponseList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RepositoriesOstreeOstreeModify

> AsyncOperationResponse RepositoriesOstreeOstreeModify(ctx, ostreeOstreeRepositoryHref).RepositoryAddRemoveContent(repositoryAddRemoveContent).Execute()

Modify repository



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/content-services/zest/release/v3"
)

func main() {
    ostreeOstreeRepositoryHref := "ostreeOstreeRepositoryHref_example" // string | 
    repositoryAddRemoveContent := *openapiclient.NewRepositoryAddRemoveContent() // RepositoryAddRemoveContent | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RepositoriesOstreeApi.RepositoriesOstreeOstreeModify(context.Background(), ostreeOstreeRepositoryHref).RepositoryAddRemoveContent(repositoryAddRemoveContent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesOstreeApi.RepositoriesOstreeOstreeModify``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesOstreeOstreeModify`: AsyncOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `RepositoriesOstreeApi.RepositoriesOstreeOstreeModify`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ostreeOstreeRepositoryHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesOstreeOstreeModifyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **repositoryAddRemoveContent** | [**RepositoryAddRemoveContent**](RepositoryAddRemoveContent.md) |  | 

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


## RepositoriesOstreeOstreePartialUpdate

> AsyncOperationResponse RepositoriesOstreeOstreePartialUpdate(ctx, ostreeOstreeRepositoryHref).PatchedostreeOstreeRepository(patchedostreeOstreeRepository).Execute()

Update an ostree repository



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/content-services/zest/release/v3"
)

func main() {
    ostreeOstreeRepositoryHref := "ostreeOstreeRepositoryHref_example" // string | 
    patchedostreeOstreeRepository := *openapiclient.NewPatchedostreeOstreeRepository() // PatchedostreeOstreeRepository | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RepositoriesOstreeApi.RepositoriesOstreeOstreePartialUpdate(context.Background(), ostreeOstreeRepositoryHref).PatchedostreeOstreeRepository(patchedostreeOstreeRepository).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesOstreeApi.RepositoriesOstreeOstreePartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesOstreeOstreePartialUpdate`: AsyncOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `RepositoriesOstreeApi.RepositoriesOstreeOstreePartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ostreeOstreeRepositoryHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesOstreeOstreePartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedostreeOstreeRepository** | [**PatchedostreeOstreeRepository**](PatchedostreeOstreeRepository.md) |  | 

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


## RepositoriesOstreeOstreeRead

> OstreeOstreeRepositoryResponse RepositoriesOstreeOstreeRead(ctx, ostreeOstreeRepositoryHref).Fields(fields).ExcludeFields(excludeFields).Execute()

Inspect an ostree repository



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/content-services/zest/release/v3"
)

func main() {
    ostreeOstreeRepositoryHref := "ostreeOstreeRepositoryHref_example" // string | 
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RepositoriesOstreeApi.RepositoriesOstreeOstreeRead(context.Background(), ostreeOstreeRepositoryHref).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesOstreeApi.RepositoriesOstreeOstreeRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesOstreeOstreeRead`: OstreeOstreeRepositoryResponse
    fmt.Fprintf(os.Stdout, "Response from `RepositoriesOstreeApi.RepositoriesOstreeOstreeRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ostreeOstreeRepositoryHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesOstreeOstreeReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**OstreeOstreeRepositoryResponse**](OstreeOstreeRepositoryResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RepositoriesOstreeOstreeSync

> AsyncOperationResponse RepositoriesOstreeOstreeSync(ctx, ostreeOstreeRepositoryHref).RepositorySyncURL(repositorySyncURL).Execute()

Sync from remote



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/content-services/zest/release/v3"
)

func main() {
    ostreeOstreeRepositoryHref := "ostreeOstreeRepositoryHref_example" // string | 
    repositorySyncURL := *openapiclient.NewRepositorySyncURL() // RepositorySyncURL | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RepositoriesOstreeApi.RepositoriesOstreeOstreeSync(context.Background(), ostreeOstreeRepositoryHref).RepositorySyncURL(repositorySyncURL).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesOstreeApi.RepositoriesOstreeOstreeSync``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesOstreeOstreeSync`: AsyncOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `RepositoriesOstreeApi.RepositoriesOstreeOstreeSync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ostreeOstreeRepositoryHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesOstreeOstreeSyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **repositorySyncURL** | [**RepositorySyncURL**](RepositorySyncURL.md) |  | 

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


## RepositoriesOstreeOstreeUpdate

> AsyncOperationResponse RepositoriesOstreeOstreeUpdate(ctx, ostreeOstreeRepositoryHref).OstreeOstreeRepository(ostreeOstreeRepository).Execute()

Update an ostree repository



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/content-services/zest/release/v3"
)

func main() {
    ostreeOstreeRepositoryHref := "ostreeOstreeRepositoryHref_example" // string | 
    ostreeOstreeRepository := *openapiclient.NewOstreeOstreeRepository("Name_example") // OstreeOstreeRepository | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RepositoriesOstreeApi.RepositoriesOstreeOstreeUpdate(context.Background(), ostreeOstreeRepositoryHref).OstreeOstreeRepository(ostreeOstreeRepository).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesOstreeApi.RepositoriesOstreeOstreeUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RepositoriesOstreeOstreeUpdate`: AsyncOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `RepositoriesOstreeApi.RepositoriesOstreeOstreeUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ostreeOstreeRepositoryHref** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesOstreeOstreeUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ostreeOstreeRepository** | [**OstreeOstreeRepository**](OstreeOstreeRepository.md) |  | 

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

