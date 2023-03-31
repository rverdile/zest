# \PulpAnsibleDefaultApiV3PluginAnsibleContentNamespacesApi

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentNamespacesCreate**](PulpAnsibleDefaultApiV3PluginAnsibleContentNamespacesApi.md#PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentNamespacesCreate) | **Post** /pulp_ansible/galaxy/default/api/v3/plugin/ansible/content/{distro_base_path}/namespaces/ | 
[**PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentNamespacesDelete**](PulpAnsibleDefaultApiV3PluginAnsibleContentNamespacesApi.md#PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentNamespacesDelete) | **Delete** /pulp_ansible/galaxy/default/api/v3/plugin/ansible/content/{distro_base_path}/namespaces/{name}/ | 
[**PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentNamespacesList**](PulpAnsibleDefaultApiV3PluginAnsibleContentNamespacesApi.md#PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentNamespacesList) | **Get** /pulp_ansible/galaxy/default/api/v3/plugin/ansible/content/{distro_base_path}/namespaces/ | 
[**PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentNamespacesPartialUpdate**](PulpAnsibleDefaultApiV3PluginAnsibleContentNamespacesApi.md#PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentNamespacesPartialUpdate) | **Patch** /pulp_ansible/galaxy/default/api/v3/plugin/ansible/content/{distro_base_path}/namespaces/{name}/ | 
[**PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentNamespacesRead**](PulpAnsibleDefaultApiV3PluginAnsibleContentNamespacesApi.md#PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentNamespacesRead) | **Get** /pulp_ansible/galaxy/default/api/v3/plugin/ansible/content/{distro_base_path}/namespaces/{name}/ | 



## PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentNamespacesCreate

> AsyncOperationResponse PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentNamespacesCreate(ctx, distroBasePath).Name(name).Company(company).Email(email).Description(description).Resources(resources).Links(links).Avatar(avatar).Execute()





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
    distroBasePath := "distroBasePath_example" // string | 
    name := "name_example" // string | Required named, only accepts lowercase, numbers and underscores.
    company := "company_example" // string | Optional namespace company owner. (optional)
    email := "email_example" // string | Optional namespace contact email. (optional)
    description := "description_example" // string | Optional short description. (optional)
    resources := "resources_example" // string | Optional resource page in markdown format. (optional)
    links := []openapiclient.NamespaceLink{*openapiclient.NewNamespaceLink("Url_example", "Name_example")} // []NamespaceLink | Labeled related links. (optional)
    avatar := os.NewFile(1234, "some_file") // *os.File | Optional avatar image for Namespace (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PulpAnsibleDefaultApiV3PluginAnsibleContentNamespacesApi.PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentNamespacesCreate(context.Background(), distroBasePath).Name(name).Company(company).Email(email).Description(description).Resources(resources).Links(links).Avatar(avatar).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PulpAnsibleDefaultApiV3PluginAnsibleContentNamespacesApi.PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentNamespacesCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentNamespacesCreate`: AsyncOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `PulpAnsibleDefaultApiV3PluginAnsibleContentNamespacesApi.PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentNamespacesCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**distroBasePath** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentNamespacesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **string** | Required named, only accepts lowercase, numbers and underscores. | 
 **company** | **string** | Optional namespace company owner. | 
 **email** | **string** | Optional namespace contact email. | 
 **description** | **string** | Optional short description. | 
 **resources** | **string** | Optional resource page in markdown format. | 
 **links** | [**[]NamespaceLink**](NamespaceLink.md) | Labeled related links. | 
 **avatar** | ***os.File** | Optional avatar image for Namespace | 

### Return type

[**AsyncOperationResponse**](AsyncOperationResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentNamespacesDelete

> AsyncOperationResponse PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentNamespacesDelete(ctx, distroBasePath, name).Execute()





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
    distroBasePath := "distroBasePath_example" // string | 
    name := "name_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PulpAnsibleDefaultApiV3PluginAnsibleContentNamespacesApi.PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentNamespacesDelete(context.Background(), distroBasePath, name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PulpAnsibleDefaultApiV3PluginAnsibleContentNamespacesApi.PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentNamespacesDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentNamespacesDelete`: AsyncOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `PulpAnsibleDefaultApiV3PluginAnsibleContentNamespacesApi.PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentNamespacesDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**distroBasePath** | **string** |  | 
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentNamespacesDeleteRequest struct via the builder pattern


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


## PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentNamespacesList

> PaginatedansibleAnsibleNamespaceMetadataResponseList PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentNamespacesList(ctx, distroBasePath).Company(company).CompanyContains(companyContains).CompanyIcontains(companyIcontains).CompanyIn(companyIn).CompanyStartswith(companyStartswith).Limit(limit).MetadataSha256(metadataSha256).MetadataSha256In(metadataSha256In).Name(name).NameContains(nameContains).NameIcontains(nameIcontains).NameIn(nameIn).NameStartswith(nameStartswith).Offset(offset).Ordering(ordering).Fields(fields).ExcludeFields(excludeFields).Execute()





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
    distroBasePath := "distroBasePath_example" // string | 
    company := "company_example" // string | Filter results where company matches value (optional)
    companyContains := "companyContains_example" // string | Filter results where company contains value (optional)
    companyIcontains := "companyIcontains_example" // string | Filter results where company contains value (optional)
    companyIn := []string{"Inner_example"} // []string | Filter results where company is in a comma-separated list of values (optional)
    companyStartswith := "companyStartswith_example" // string | Filter results where company starts with value (optional)
    limit := int32(56) // int32 | Number of results to return per page. (optional)
    metadataSha256 := "metadataSha256_example" // string | Filter results where metadata_sha256 matches value (optional)
    metadataSha256In := []string{"Inner_example"} // []string | Filter results where metadata_sha256 is in a comma-separated list of values (optional)
    name := "name_example" // string | Filter results where name matches value (optional)
    nameContains := "nameContains_example" // string | Filter results where name contains value (optional)
    nameIcontains := "nameIcontains_example" // string | Filter results where name contains value (optional)
    nameIn := []string{"Inner_example"} // []string | Filter results where name is in a comma-separated list of values (optional)
    nameStartswith := "nameStartswith_example" // string | Filter results where name starts with value (optional)
    offset := int32(56) // int32 | The initial index from which to return the results. (optional)
    ordering := []string{"Ordering_example"} // []string | Ordering (optional)
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PulpAnsibleDefaultApiV3PluginAnsibleContentNamespacesApi.PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentNamespacesList(context.Background(), distroBasePath).Company(company).CompanyContains(companyContains).CompanyIcontains(companyIcontains).CompanyIn(companyIn).CompanyStartswith(companyStartswith).Limit(limit).MetadataSha256(metadataSha256).MetadataSha256In(metadataSha256In).Name(name).NameContains(nameContains).NameIcontains(nameIcontains).NameIn(nameIn).NameStartswith(nameStartswith).Offset(offset).Ordering(ordering).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PulpAnsibleDefaultApiV3PluginAnsibleContentNamespacesApi.PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentNamespacesList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentNamespacesList`: PaginatedansibleAnsibleNamespaceMetadataResponseList
    fmt.Fprintf(os.Stdout, "Response from `PulpAnsibleDefaultApiV3PluginAnsibleContentNamespacesApi.PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentNamespacesList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**distroBasePath** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentNamespacesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **company** | **string** | Filter results where company matches value | 
 **companyContains** | **string** | Filter results where company contains value | 
 **companyIcontains** | **string** | Filter results where company contains value | 
 **companyIn** | **[]string** | Filter results where company is in a comma-separated list of values | 
 **companyStartswith** | **string** | Filter results where company starts with value | 
 **limit** | **int32** | Number of results to return per page. | 
 **metadataSha256** | **string** | Filter results where metadata_sha256 matches value | 
 **metadataSha256In** | **[]string** | Filter results where metadata_sha256 is in a comma-separated list of values | 
 **name** | **string** | Filter results where name matches value | 
 **nameContains** | **string** | Filter results where name contains value | 
 **nameIcontains** | **string** | Filter results where name contains value | 
 **nameIn** | **[]string** | Filter results where name is in a comma-separated list of values | 
 **nameStartswith** | **string** | Filter results where name starts with value | 
 **offset** | **int32** | The initial index from which to return the results. | 
 **ordering** | **[]string** | Ordering | 
 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**PaginatedansibleAnsibleNamespaceMetadataResponseList**](PaginatedansibleAnsibleNamespaceMetadataResponseList.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentNamespacesPartialUpdate

> AsyncOperationResponse PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentNamespacesPartialUpdate(ctx, distroBasePath, name).Name2(name2).Company(company).Email(email).Description(description).Resources(resources).Links(links).Avatar(avatar).Execute()





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
    distroBasePath := "distroBasePath_example" // string | 
    name := "name_example" // string | 
    name2 := "name_example" // string | Required named, only accepts lowercase, numbers and underscores. (optional)
    company := "company_example" // string | Optional namespace company owner. (optional)
    email := "email_example" // string | Optional namespace contact email. (optional)
    description := "description_example" // string | Optional short description. (optional)
    resources := "resources_example" // string | Optional resource page in markdown format. (optional)
    links := []openapiclient.NamespaceLink{*openapiclient.NewNamespaceLink("Url_example", "Name_example")} // []NamespaceLink | Labeled related links. (optional)
    avatar := os.NewFile(1234, "some_file") // *os.File | Optional avatar image for Namespace (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PulpAnsibleDefaultApiV3PluginAnsibleContentNamespacesApi.PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentNamespacesPartialUpdate(context.Background(), distroBasePath, name).Name2(name2).Company(company).Email(email).Description(description).Resources(resources).Links(links).Avatar(avatar).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PulpAnsibleDefaultApiV3PluginAnsibleContentNamespacesApi.PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentNamespacesPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentNamespacesPartialUpdate`: AsyncOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `PulpAnsibleDefaultApiV3PluginAnsibleContentNamespacesApi.PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentNamespacesPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**distroBasePath** | **string** |  | 
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentNamespacesPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **name2** | **string** | Required named, only accepts lowercase, numbers and underscores. | 
 **company** | **string** | Optional namespace company owner. | 
 **email** | **string** | Optional namespace contact email. | 
 **description** | **string** | Optional short description. | 
 **resources** | **string** | Optional resource page in markdown format. | 
 **links** | [**[]NamespaceLink**](NamespaceLink.md) | Labeled related links. | 
 **avatar** | ***os.File** | Optional avatar image for Namespace | 

### Return type

[**AsyncOperationResponse**](AsyncOperationResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentNamespacesRead

> AnsibleAnsibleNamespaceMetadataResponse PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentNamespacesRead(ctx, distroBasePath, name).Fields(fields).ExcludeFields(excludeFields).Execute()





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
    distroBasePath := "distroBasePath_example" // string | 
    name := "name_example" // string | 
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PulpAnsibleDefaultApiV3PluginAnsibleContentNamespacesApi.PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentNamespacesRead(context.Background(), distroBasePath, name).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PulpAnsibleDefaultApiV3PluginAnsibleContentNamespacesApi.PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentNamespacesRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentNamespacesRead`: AnsibleAnsibleNamespaceMetadataResponse
    fmt.Fprintf(os.Stdout, "Response from `PulpAnsibleDefaultApiV3PluginAnsibleContentNamespacesApi.PulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentNamespacesRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**distroBasePath** | **string** |  | 
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPulpAnsibleGalaxyDefaultApiV3PluginAnsibleContentNamespacesReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **[]string** | A list of fields to include in the response. | 
 **excludeFields** | **[]string** | A list of fields to exclude from the response. | 

### Return type

[**AnsibleAnsibleNamespaceMetadataResponse**](AnsibleAnsibleNamespaceMetadataResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

