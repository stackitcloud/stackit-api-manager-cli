# \APIManagerServiceApi

All URIs are relative to *https://api-manager.api.stackit.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**APIManagerServicePublish**](APIManagerServiceApi.md#APIManagerServicePublish) | **Post** /v1/projects/{metadata.project_id}/api/{metadata.identifier} | Publish API Endpoint
[**APIManagerServicePublishValidate**](APIManagerServiceApi.md#APIManagerServicePublishValidate) | **Post** /v1/projects/{metadata.project_id}/api/{metadata.identifier}/validate | Validate API Endpoint
[**APIManagerServiceRetire**](APIManagerServiceApi.md#APIManagerServiceRetire) | **Delete** /v1/projects/{metadata.project_id}/api/{metadata.identifier} | Retire API Endpoint



## APIManagerServicePublish

> PublishResponse APIManagerServicePublish(ctx, metadataProjectId, metadataIdentifier).PublishRequest(publishRequest).Execute()

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
    metadataProjectId := "metadataProjectId_example" // string | 
    metadataIdentifier := "metadataIdentifier_example" // string | 
    publishRequest := *openapiclient.NewPublishRequest() // PublishRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.APIManagerServiceApi.APIManagerServicePublish(context.Background(), metadataProjectId, metadataIdentifier).PublishRequest(publishRequest).Execute()
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
**metadataProjectId** | **string** |  | 
**metadataIdentifier** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAPIManagerServicePublishRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **publishRequest** | [**PublishRequest**](PublishRequest.md) |  | 

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

> map[string]interface{} APIManagerServicePublishValidate(ctx, metadataProjectId, metadataIdentifier).PublishValidateRequest(publishValidateRequest).Execute()

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
    metadataProjectId := "metadataProjectId_example" // string | 
    metadataIdentifier := "metadataIdentifier_example" // string | 
    publishValidateRequest := *openapiclient.NewPublishValidateRequest() // PublishValidateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.APIManagerServiceApi.APIManagerServicePublishValidate(context.Background(), metadataProjectId, metadataIdentifier).PublishValidateRequest(publishValidateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `APIManagerServiceApi.APIManagerServicePublishValidate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `APIManagerServicePublishValidate`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `APIManagerServiceApi.APIManagerServicePublishValidate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**metadataProjectId** | **string** |  | 
**metadataIdentifier** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAPIManagerServicePublishValidateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **publishValidateRequest** | [**PublishValidateRequest**](PublishValidateRequest.md) |  | 

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


## APIManagerServiceRetire

> map[string]interface{} APIManagerServiceRetire(ctx, metadataProjectId, metadataIdentifier).RetireRequest(retireRequest).Execute()

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
    metadataProjectId := "metadataProjectId_example" // string | 
    metadataIdentifier := "metadataIdentifier_example" // string | 
    retireRequest := *openapiclient.NewRetireRequest() // RetireRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.APIManagerServiceApi.APIManagerServiceRetire(context.Background(), metadataProjectId, metadataIdentifier).RetireRequest(retireRequest).Execute()
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
**metadataProjectId** | **string** |  | 
**metadataIdentifier** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAPIManagerServiceRetireRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **retireRequest** | [**RetireRequest**](RetireRequest.md) |  | 

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

