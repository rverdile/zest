# \PulpAnsibleDefaultApiV3NamespacesApi

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PulpAnsibleGalaxyDefaultApiV3NamespacesList**](PulpAnsibleDefaultApiV3NamespacesApi.md#PulpAnsibleGalaxyDefaultApiV3NamespacesList) | **Get** /pulp_ansible/galaxy/default/api/v3/namespaces/ | 
[**PulpAnsibleGalaxyDefaultApiV3NamespacesRead**](PulpAnsibleDefaultApiV3NamespacesApi.md#PulpAnsibleGalaxyDefaultApiV3NamespacesRead) | **Get** /pulp_ansible/galaxy/default/api/v3/namespaces/{name}/ | 



## PulpAnsibleGalaxyDefaultApiV3NamespacesList

> PaginatedansibleAnsibleNamespaceMetadataResponseList PulpAnsibleGalaxyDefaultApiV3NamespacesList(ctx).Company(company).CompanyContains(companyContains).CompanyIcontains(companyIcontains).CompanyIn(companyIn).CompanyStartswith(companyStartswith).Limit(limit).MetadataSha256(metadataSha256).MetadataSha256In(metadataSha256In).Name(name).NameContains(nameContains).NameIcontains(nameIcontains).NameIn(nameIn).NameStartswith(nameStartswith).Offset(offset).Ordering(ordering).Fields(fields).ExcludeFields(excludeFields).Execute()





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
    resp, r, err := apiClient.PulpAnsibleDefaultApiV3NamespacesApi.PulpAnsibleGalaxyDefaultApiV3NamespacesList(context.Background()).Company(company).CompanyContains(companyContains).CompanyIcontains(companyIcontains).CompanyIn(companyIn).CompanyStartswith(companyStartswith).Limit(limit).MetadataSha256(metadataSha256).MetadataSha256In(metadataSha256In).Name(name).NameContains(nameContains).NameIcontains(nameIcontains).NameIn(nameIn).NameStartswith(nameStartswith).Offset(offset).Ordering(ordering).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PulpAnsibleDefaultApiV3NamespacesApi.PulpAnsibleGalaxyDefaultApiV3NamespacesList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PulpAnsibleGalaxyDefaultApiV3NamespacesList`: PaginatedansibleAnsibleNamespaceMetadataResponseList
    fmt.Fprintf(os.Stdout, "Response from `PulpAnsibleDefaultApiV3NamespacesApi.PulpAnsibleGalaxyDefaultApiV3NamespacesList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPulpAnsibleGalaxyDefaultApiV3NamespacesListRequest struct via the builder pattern


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


## PulpAnsibleGalaxyDefaultApiV3NamespacesRead

> AnsibleAnsibleNamespaceMetadataResponse PulpAnsibleGalaxyDefaultApiV3NamespacesRead(ctx, name).Fields(fields).ExcludeFields(excludeFields).Execute()





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
    name := "name_example" // string | 
    fields := []string{"Inner_example"} // []string | A list of fields to include in the response. (optional)
    excludeFields := []string{"Inner_example"} // []string | A list of fields to exclude from the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PulpAnsibleDefaultApiV3NamespacesApi.PulpAnsibleGalaxyDefaultApiV3NamespacesRead(context.Background(), name).Fields(fields).ExcludeFields(excludeFields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PulpAnsibleDefaultApiV3NamespacesApi.PulpAnsibleGalaxyDefaultApiV3NamespacesRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PulpAnsibleGalaxyDefaultApiV3NamespacesRead`: AnsibleAnsibleNamespaceMetadataResponse
    fmt.Fprintf(os.Stdout, "Response from `PulpAnsibleDefaultApiV3NamespacesApi.PulpAnsibleGalaxyDefaultApiV3NamespacesRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPulpAnsibleGalaxyDefaultApiV3NamespacesReadRequest struct via the builder pattern


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

