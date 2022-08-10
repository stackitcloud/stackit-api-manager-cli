# \APIManagerServiceApi

All URIs are relative to *http://api-manager.api.dev.stackit.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**APIManagerServicePublish**](APIManagerServiceApi.md#APIManagerServicePublish) | **Post** /v1/projects/{projectId}/api/{identifier} | Publish API Endpoint
[**APIManagerServiceRetire**](APIManagerServiceApi.md#APIManagerServiceRetire) | **Delete** /v1/projects/{projectId}/api/{identifier} | Retire API Endpoint



## APIManagerServicePublish

> map[string]interface{} APIManagerServicePublish(ctx, projectId, identifier).Body(body).Execute()

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
    projectId := "projectId_example" // string | 
    identifier := "identifier_example" // string | 
    body := *openapiclient.NewAPIManagerServicePublishRequest() // APIManagerServicePublishRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.APIManagerServiceApi.APIManagerServicePublish(context.Background(), projectId, identifier).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `APIManagerServiceApi.APIManagerServicePublish``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `APIManagerServicePublish`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `APIManagerServiceApi.APIManagerServicePublish`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** |  | 
**identifier** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAPIManagerServicePublishRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**APIManagerServicePublishRequest**](APIManagerServicePublishRequest.md) |  | 

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

> map[string]interface{} APIManagerServiceRetire(ctx, projectId, identifier).Body(body).Execute()

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
    projectId := "projectId_example" // string | 
    identifier := "identifier_example" // string | 
    body := *openapiclient.NewAPIManagerServiceRetireRequest() // APIManagerServiceRetireRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.APIManagerServiceApi.APIManagerServiceRetire(context.Background(), projectId, identifier).Body(body).Execute()
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
**projectId** | **string** |  | 
**identifier** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAPIManagerServiceRetireRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**APIManagerServiceRetireRequest**](APIManagerServiceRetireRequest.md) |  | 

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

