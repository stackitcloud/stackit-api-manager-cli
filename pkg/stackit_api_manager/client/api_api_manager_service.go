/*
STACKIT API Management Service

STACKIT API Manager

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// APIManagerServiceApiService APIManagerServiceApi service
type APIManagerServiceApiService service

type ApiAPIManagerServiceFetchAPIRequest struct {
	ctx        context.Context
	ApiService *APIManagerServiceApiService
	projectId  string
	identifier string
}

func (r ApiAPIManagerServiceFetchAPIRequest) Execute() (*FetchAPIResponse, *http.Response, error) {
	return r.ApiService.APIManagerServiceFetchAPIExecute(r)
}

/*
APIManagerServiceFetchAPI Fetch API Endpoint

Fetches an already existing API for a dedicated service by providing its identifier

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param projectId
	@param identifier
	@return ApiAPIManagerServiceFetchAPIRequest
*/
func (a *APIManagerServiceApiService) APIManagerServiceFetchAPI(ctx context.Context, projectId string, identifier string) ApiAPIManagerServiceFetchAPIRequest {
	return ApiAPIManagerServiceFetchAPIRequest{
		ApiService: a,
		ctx:        ctx,
		projectId:  projectId,
		identifier: identifier,
	}
}

// Execute executes the request
//
//	@return FetchAPIResponse
func (a *APIManagerServiceApiService) APIManagerServiceFetchAPIExecute(r ApiAPIManagerServiceFetchAPIRequest) (*FetchAPIResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *FetchAPIResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "APIManagerServiceApiService.APIManagerServiceFetchAPI")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/projects/{projectId}/api/{identifier}"
	localVarPath = strings.Replace(localVarPath, "{"+"projectId"+"}", url.PathEscape(parameterToString(r.projectId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"identifier"+"}", url.PathEscape(parameterToString(r.identifier, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
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

type ApiAPIManagerServiceFetchProjectAPIIdentifiersRequest struct {
	ctx        context.Context
	ApiService *APIManagerServiceApiService
	projectId  string
}

func (r ApiAPIManagerServiceFetchProjectAPIIdentifiersRequest) Execute() (*FetchProjectAPIIdentifiersResponse, *http.Response, error) {
	return r.ApiService.APIManagerServiceFetchProjectAPIIdentifiersExecute(r)
}

/*
APIManagerServiceFetchProjectAPIIdentifiers Fetch Project APIIdentifiers Endpoint

Fetches all API identifiers which belong to a project by providing its projectId

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param projectId
	@return ApiAPIManagerServiceFetchProjectAPIIdentifiersRequest
*/
func (a *APIManagerServiceApiService) APIManagerServiceFetchProjectAPIIdentifiers(ctx context.Context, projectId string) ApiAPIManagerServiceFetchProjectAPIIdentifiersRequest {
	return ApiAPIManagerServiceFetchProjectAPIIdentifiersRequest{
		ApiService: a,
		ctx:        ctx,
		projectId:  projectId,
	}
}

// Execute executes the request
//
//	@return FetchProjectAPIIdentifiersResponse
func (a *APIManagerServiceApiService) APIManagerServiceFetchProjectAPIIdentifiersExecute(r ApiAPIManagerServiceFetchProjectAPIIdentifiersRequest) (*FetchProjectAPIIdentifiersResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *FetchProjectAPIIdentifiersResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "APIManagerServiceApiService.APIManagerServiceFetchProjectAPIIdentifiers")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/projects/{projectId}"
	localVarPath = strings.Replace(localVarPath, "{"+"projectId"+"}", url.PathEscape(parameterToString(r.projectId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
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

type ApiAPIManagerServicePublishRequest struct {
	ctx            context.Context
	ApiService     *APIManagerServiceApiService
	projectId      string
	identifier     string
	publishRequest *PublishRequest
}

func (r ApiAPIManagerServicePublishRequest) PublishRequest(publishRequest PublishRequest) ApiAPIManagerServicePublishRequest {
	r.publishRequest = &publishRequest
	return r
}

func (r ApiAPIManagerServicePublishRequest) Execute() (*PublishResponse, *http.Response, error) {
	return r.ApiService.APIManagerServicePublishExecute(r)
}

/*
APIManagerServicePublish Publish API Endpoint

Publish a new API for a dedicated service by providing the OpenApiSpec for it

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param projectId
	@param identifier
	@return ApiAPIManagerServicePublishRequest
*/
func (a *APIManagerServiceApiService) APIManagerServicePublish(ctx context.Context, projectId string, identifier string) ApiAPIManagerServicePublishRequest {
	return ApiAPIManagerServicePublishRequest{
		ApiService: a,
		ctx:        ctx,
		projectId:  projectId,
		identifier: identifier,
	}
}

// Execute executes the request
//
//	@return PublishResponse
func (a *APIManagerServiceApiService) APIManagerServicePublishExecute(r ApiAPIManagerServicePublishRequest) (*PublishResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PublishResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "APIManagerServiceApiService.APIManagerServicePublish")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/projects/{projectId}/api/{identifier}"
	localVarPath = strings.Replace(localVarPath, "{"+"projectId"+"}", url.PathEscape(parameterToString(r.projectId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"identifier"+"}", url.PathEscape(parameterToString(r.identifier, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.publishRequest == nil {
		return localVarReturnValue, nil, reportError("publishRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.publishRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
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

type ApiAPIManagerServicePublishValidateRequest struct {
	ctx                    context.Context
	ApiService             *APIManagerServiceApiService
	projectId              string
	identifier             string
	publishValidateRequest *PublishValidateRequest
}

func (r ApiAPIManagerServicePublishValidateRequest) PublishValidateRequest(publishValidateRequest PublishValidateRequest) ApiAPIManagerServicePublishValidateRequest {
	r.publishValidateRequest = &publishValidateRequest
	return r
}

func (r ApiAPIManagerServicePublishValidateRequest) Execute() (*PublishValidateResponse, *http.Response, error) {
	return r.ApiService.APIManagerServicePublishValidateExecute(r)
}

/*
APIManagerServicePublishValidate Validate API Endpoint

Validate the OpenApiSpec for an API by providing the OAS for it

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param projectId
	@param identifier
	@return ApiAPIManagerServicePublishValidateRequest
*/
func (a *APIManagerServiceApiService) APIManagerServicePublishValidate(ctx context.Context, projectId string, identifier string) ApiAPIManagerServicePublishValidateRequest {
	return ApiAPIManagerServicePublishValidateRequest{
		ApiService: a,
		ctx:        ctx,
		projectId:  projectId,
		identifier: identifier,
	}
}

// Execute executes the request
//
//	@return PublishValidateResponse
func (a *APIManagerServiceApiService) APIManagerServicePublishValidateExecute(r ApiAPIManagerServicePublishValidateRequest) (*PublishValidateResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PublishValidateResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "APIManagerServiceApiService.APIManagerServicePublishValidate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/projects/{projectId}/api/{identifier}/validate"
	localVarPath = strings.Replace(localVarPath, "{"+"projectId"+"}", url.PathEscape(parameterToString(r.projectId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"identifier"+"}", url.PathEscape(parameterToString(r.identifier, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.publishValidateRequest == nil {
		return localVarReturnValue, nil, reportError("publishValidateRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.publishValidateRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
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

type ApiAPIManagerServiceRetireRequest struct {
	ctx           context.Context
	ApiService    *APIManagerServiceApiService
	projectId     string
	identifier    string
	retireRequest *RetireRequest
}

func (r ApiAPIManagerServiceRetireRequest) RetireRequest(retireRequest RetireRequest) ApiAPIManagerServiceRetireRequest {
	r.retireRequest = &retireRequest
	return r
}

func (r ApiAPIManagerServiceRetireRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.APIManagerServiceRetireExecute(r)
}

/*
APIManagerServiceRetire Retire API Endpoint

Retire an already existing API for a dedicated service by providing its Identifier

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param projectId
	@param identifier
	@return ApiAPIManagerServiceRetireRequest
*/
func (a *APIManagerServiceApiService) APIManagerServiceRetire(ctx context.Context, projectId string, identifier string) ApiAPIManagerServiceRetireRequest {
	return ApiAPIManagerServiceRetireRequest{
		ApiService: a,
		ctx:        ctx,
		projectId:  projectId,
		identifier: identifier,
	}
}

// Execute executes the request
//
//	@return map[string]interface{}
func (a *APIManagerServiceApiService) APIManagerServiceRetireExecute(r ApiAPIManagerServiceRetireRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodDelete
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "APIManagerServiceApiService.APIManagerServiceRetire")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/projects/{projectId}/api/{identifier}"
	localVarPath = strings.Replace(localVarPath, "{"+"projectId"+"}", url.PathEscape(parameterToString(r.projectId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"identifier"+"}", url.PathEscape(parameterToString(r.identifier, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.retireRequest == nil {
		return localVarReturnValue, nil, reportError("retireRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.retireRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
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
