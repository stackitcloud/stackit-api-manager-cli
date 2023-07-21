# \APIManagerServiceApi

All URIs are relative to *https://api-manager.api.stackit.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**APIManagerServiceFetchAPI**](APIManagerServiceApi.md#APIManagerServiceFetchAPI) | **Get** /v1/projects/{projectId}/api/{identifier} | Fetch API Endpoint
[**APIManagerServiceFetchProjectAPIIdentifiers**](APIManagerServiceApi.md#APIManagerServiceFetchProjectAPIIdentifiers) | **Get** /v1/projects/{projectId} | Fetch Project APIIdentifiers Endpoint
[**APIManagerServicePublish**](APIManagerServiceApi.md#APIManagerServicePublish) | **Post** /v1/projects/{projectId}/api/{identifier} | Publish API Endpoint
[**APIManagerServicePublishValidate**](APIManagerServiceApi.md#APIManagerServicePublishValidate) | **Post** /v1/projects/{projectId}/api/{identifier}/validate | Validate API Endpoint
[**APIManagerServiceRetire**](APIManagerServiceApi.md#APIManagerServiceRetire) | **Delete** /v1/projects/{projectId}/api/{identifier} | Retire API Endpoint
[**APIManagerServiceRetireVersion**](APIManagerServiceApi.md#APIManagerServiceRetireVersion) | **Delete** /v1/projects/{projectId}/api/{identifier}/version/{version} | Retire a specific API Version



## APIManagerServiceFetchAPI

> FetchAPIResponse APIManagerServiceFetchAPI(ctx, projectId, identifier).ApiVersion(apiVersion).Execute()

Fetch API Endpoint



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    projectId := "5s239152-24ky-5924-1077-m29ad542f6s" // string | Project ID for API to be fetched
    identifier := "api-identifier" // string | Identifier of API to be fetched
    apiVersion := "apiVersion_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.APIManagerServiceApi.APIManagerServiceFetchAPI(context.Background(), projectId, identifier).ApiVersion(apiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `APIManagerServiceApi.APIManagerServiceFetchAPI``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `APIManagerServiceFetchAPI`: FetchAPIResponse
    fmt.Fprintf(os.Stdout, "Response from `APIManagerServiceApi.APIManagerServiceFetchAPI`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID for API to be fetched | 
**identifier** | **string** | Identifier of API to be fetched | 

### Other Parameters

Other parameters are passed through a pointer to a apiAPIManagerServiceFetchAPIRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **apiVersion** | **string** |  | 

### Return type

[**FetchAPIResponse**](FetchAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## APIManagerServiceFetchProjectAPIIdentifiers

> FetchProjectAPIIdentifiersResponse APIManagerServiceFetchProjectAPIIdentifiers(ctx, projectId).Execute()

Fetch Project APIIdentifiers Endpoint



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    projectId := "5s239152-24ky-5924-1077-m29ad542f6s" // string | Project ID for which APIs are to be fetched

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.APIManagerServiceApi.APIManagerServiceFetchProjectAPIIdentifiers(context.Background(), projectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `APIManagerServiceApi.APIManagerServiceFetchProjectAPIIdentifiers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `APIManagerServiceFetchProjectAPIIdentifiers`: FetchProjectAPIIdentifiersResponse
    fmt.Fprintf(os.Stdout, "Response from `APIManagerServiceApi.APIManagerServiceFetchProjectAPIIdentifiers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID for which APIs are to be fetched | 

### Other Parameters

Other parameters are passed through a pointer to a apiAPIManagerServiceFetchProjectAPIIdentifiersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FetchProjectAPIIdentifiersResponse**](FetchProjectAPIIdentifiersResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## APIManagerServicePublish

> PublishResponse APIManagerServicePublish(ctx, projectId, identifier).PublishRequest(publishRequest).Execute()

Publish API Endpoint



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    projectId := "5s239152-24ky-5924-1077-m29ad542f6s" // string | Project ID for API to be published
    identifier := "api-identifier" // string | Identifier of API to be published
    publishRequest := *openapiclient.NewPublishRequest() // PublishRequest | Request body for the Publish request containing the resources to publish an API

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.APIManagerServiceApi.APIManagerServicePublish(context.Background(), projectId, identifier).PublishRequest(publishRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `APIManagerServiceApi.APIManagerServicePublish``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `APIManagerServicePublish`: PublishResponse
    fmt.Fprintf(os.Stdout, "Response from `APIManagerServiceApi.APIManagerServicePublish`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID for API to be published | 
**identifier** | **string** | Identifier of API to be published | 

### Other Parameters

Other parameters are passed through a pointer to a apiAPIManagerServicePublishRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **publishRequest** | [**PublishRequest**](PublishRequest.md) | Request body for the Publish request containing the resources to publish an API | 

### Return type

[**PublishResponse**](PublishResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## APIManagerServicePublishValidate

> PublishValidateResponse APIManagerServicePublishValidate(ctx, projectId, identifier).PublishValidateRequest(publishValidateRequest).Execute()

Validate API Endpoint



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    projectId := "5s239152-24ky-5924-1077-m29ad542f6s" // string | Project ID for API to be validated
    identifier := "api-identifier" // string | Identifier of API to be validated
    publishValidateRequest := *openapiclient.NewPublishValidateRequest() // PublishValidateRequest | Request body for the PublishValidate request containing the resources to publish an API

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.APIManagerServiceApi.APIManagerServicePublishValidate(context.Background(), projectId, identifier).PublishValidateRequest(publishValidateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `APIManagerServiceApi.APIManagerServicePublishValidate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `APIManagerServicePublishValidate`: PublishValidateResponse
    fmt.Fprintf(os.Stdout, "Response from `APIManagerServiceApi.APIManagerServicePublishValidate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID for API to be validated | 
**identifier** | **string** | Identifier of API to be validated | 

### Other Parameters

Other parameters are passed through a pointer to a apiAPIManagerServicePublishValidateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **publishValidateRequest** | [**PublishValidateRequest**](PublishValidateRequest.md) | Request body for the PublishValidate request containing the resources to publish an API | 

### Return type

[**PublishValidateResponse**](PublishValidateResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## APIManagerServiceRetire

> map[string]interface{} APIManagerServiceRetire(ctx, projectId, identifier).RetireRequest(retireRequest).Execute()

Retire API Endpoint



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    projectId := "5s239152-24ky-5924-1077-m29ad542f6s" // string | Project ID for API to be retired
    identifier := "api-identifier" // string | Identifier of API to be retired
    retireRequest := *openapiclient.NewRetireRequest() // RetireRequest | Request body for the Retire request containing the resources to retire an API

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.APIManagerServiceApi.APIManagerServiceRetire(context.Background(), projectId, identifier).RetireRequest(retireRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `APIManagerServiceApi.APIManagerServiceRetire``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `APIManagerServiceRetire`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `APIManagerServiceApi.APIManagerServiceRetire`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID for API to be retired | 
**identifier** | **string** | Identifier of API to be retired | 

### Other Parameters

Other parameters are passed through a pointer to a apiAPIManagerServiceRetireRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **retireRequest** | [**RetireRequest**](RetireRequest.md) | Request body for the Retire request containing the resources to retire an API | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## APIManagerServiceRetireVersion

> map[string]interface{} APIManagerServiceRetireVersion(ctx, projectId, identifier, version).RetireVersionRequest(retireVersionRequest).Execute()

Retire a specific API Version



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    projectId := "5s239152-24ky-5924-1077-m29ad542f6s" // string | Project ID for API to be retired
    identifier := "api-identifier" // string | Identifier of API to be retired
    version := "v1" // string | version of the API to be retired
    retireVersionRequest := *openapiclient.NewRetireVersionRequest() // RetireVersionRequest | Request body for the Retire Version request containing the resources to retire an API Version

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.APIManagerServiceApi.APIManagerServiceRetireVersion(context.Background(), projectId, identifier, version).RetireVersionRequest(retireVersionRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `APIManagerServiceApi.APIManagerServiceRetireVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `APIManagerServiceRetireVersion`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `APIManagerServiceApi.APIManagerServiceRetireVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Project ID for API to be retired | 
**identifier** | **string** | Identifier of API to be retired | 
**version** | **string** | version of the API to be retired | 

### Other Parameters

Other parameters are passed through a pointer to a apiAPIManagerServiceRetireVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **retireVersionRequest** | [**RetireVersionRequest**](RetireVersionRequest.md) | Request body for the Retire Version request containing the resources to retire an API Version | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

