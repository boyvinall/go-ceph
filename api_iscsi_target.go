/*
Ceph REST API

This is the official Ceph REST API

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ceph

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


// IscsiTargetAPIService IscsiTargetAPI service
type IscsiTargetAPIService service

type ApiApiIscsiTargetGetRequest struct {
	ctx context.Context
	ApiService *IscsiTargetAPIService
}

func (r ApiApiIscsiTargetGetRequest) Execute() (*http.Response, error) {
	return r.ApiService.ApiIscsiTargetGetExecute(r)
}

/*
ApiIscsiTargetGet Method for ApiIscsiTargetGet

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiApiIscsiTargetGetRequest
*/
func (a *IscsiTargetAPIService) ApiIscsiTargetGet(ctx context.Context) ApiApiIscsiTargetGetRequest {
	return ApiApiIscsiTargetGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *IscsiTargetAPIService) ApiIscsiTargetGetExecute(r ApiApiIscsiTargetGetRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IscsiTargetAPIService.ApiIscsiTargetGet")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/iscsi/target"

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
	localVarHTTPHeaderAccepts := []string{"application/vnd.ceph.api.v1.0+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiApiIscsiTargetPostRequest struct {
	ctx context.Context
	ApiService *IscsiTargetAPIService
	apiIscsiTargetPostRequest *ApiIscsiTargetPostRequest
}

func (r ApiApiIscsiTargetPostRequest) ApiIscsiTargetPostRequest(apiIscsiTargetPostRequest ApiIscsiTargetPostRequest) ApiApiIscsiTargetPostRequest {
	r.apiIscsiTargetPostRequest = &apiIscsiTargetPostRequest
	return r
}

func (r ApiApiIscsiTargetPostRequest) Execute() (*http.Response, error) {
	return r.ApiService.ApiIscsiTargetPostExecute(r)
}

/*
ApiIscsiTargetPost Method for ApiIscsiTargetPost

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiApiIscsiTargetPostRequest
*/
func (a *IscsiTargetAPIService) ApiIscsiTargetPost(ctx context.Context) ApiApiIscsiTargetPostRequest {
	return ApiApiIscsiTargetPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *IscsiTargetAPIService) ApiIscsiTargetPostExecute(r ApiApiIscsiTargetPostRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IscsiTargetAPIService.ApiIscsiTargetPost")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/iscsi/target"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.ceph.api.v1.0+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.apiIscsiTargetPostRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiApiIscsiTargetTargetIqnDeleteRequest struct {
	ctx context.Context
	ApiService *IscsiTargetAPIService
	targetIqn string
}

func (r ApiApiIscsiTargetTargetIqnDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.ApiIscsiTargetTargetIqnDeleteExecute(r)
}

/*
ApiIscsiTargetTargetIqnDelete Method for ApiIscsiTargetTargetIqnDelete

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param targetIqn
 @return ApiApiIscsiTargetTargetIqnDeleteRequest
*/
func (a *IscsiTargetAPIService) ApiIscsiTargetTargetIqnDelete(ctx context.Context, targetIqn string) ApiApiIscsiTargetTargetIqnDeleteRequest {
	return ApiApiIscsiTargetTargetIqnDeleteRequest{
		ApiService: a,
		ctx: ctx,
		targetIqn: targetIqn,
	}
}

// Execute executes the request
func (a *IscsiTargetAPIService) ApiIscsiTargetTargetIqnDeleteExecute(r ApiApiIscsiTargetTargetIqnDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IscsiTargetAPIService.ApiIscsiTargetTargetIqnDelete")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/iscsi/target/{target_iqn}"
	localVarPath = strings.Replace(localVarPath, "{"+"target_iqn"+"}", url.PathEscape(parameterValueToString(r.targetIqn, "targetIqn")), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/vnd.ceph.api.v1.0+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiApiIscsiTargetTargetIqnGetRequest struct {
	ctx context.Context
	ApiService *IscsiTargetAPIService
	targetIqn string
}

func (r ApiApiIscsiTargetTargetIqnGetRequest) Execute() (*http.Response, error) {
	return r.ApiService.ApiIscsiTargetTargetIqnGetExecute(r)
}

/*
ApiIscsiTargetTargetIqnGet Method for ApiIscsiTargetTargetIqnGet

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param targetIqn
 @return ApiApiIscsiTargetTargetIqnGetRequest
*/
func (a *IscsiTargetAPIService) ApiIscsiTargetTargetIqnGet(ctx context.Context, targetIqn string) ApiApiIscsiTargetTargetIqnGetRequest {
	return ApiApiIscsiTargetTargetIqnGetRequest{
		ApiService: a,
		ctx: ctx,
		targetIqn: targetIqn,
	}
}

// Execute executes the request
func (a *IscsiTargetAPIService) ApiIscsiTargetTargetIqnGetExecute(r ApiApiIscsiTargetTargetIqnGetRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IscsiTargetAPIService.ApiIscsiTargetTargetIqnGet")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/iscsi/target/{target_iqn}"
	localVarPath = strings.Replace(localVarPath, "{"+"target_iqn"+"}", url.PathEscape(parameterValueToString(r.targetIqn, "targetIqn")), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/vnd.ceph.api.v1.0+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiApiIscsiTargetTargetIqnPutRequest struct {
	ctx context.Context
	ApiService *IscsiTargetAPIService
	targetIqn string
	apiIscsiTargetTargetIqnPutRequest *ApiIscsiTargetTargetIqnPutRequest
}

func (r ApiApiIscsiTargetTargetIqnPutRequest) ApiIscsiTargetTargetIqnPutRequest(apiIscsiTargetTargetIqnPutRequest ApiIscsiTargetTargetIqnPutRequest) ApiApiIscsiTargetTargetIqnPutRequest {
	r.apiIscsiTargetTargetIqnPutRequest = &apiIscsiTargetTargetIqnPutRequest
	return r
}

func (r ApiApiIscsiTargetTargetIqnPutRequest) Execute() (*http.Response, error) {
	return r.ApiService.ApiIscsiTargetTargetIqnPutExecute(r)
}

/*
ApiIscsiTargetTargetIqnPut Method for ApiIscsiTargetTargetIqnPut

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param targetIqn
 @return ApiApiIscsiTargetTargetIqnPutRequest
*/
func (a *IscsiTargetAPIService) ApiIscsiTargetTargetIqnPut(ctx context.Context, targetIqn string) ApiApiIscsiTargetTargetIqnPutRequest {
	return ApiApiIscsiTargetTargetIqnPutRequest{
		ApiService: a,
		ctx: ctx,
		targetIqn: targetIqn,
	}
}

// Execute executes the request
func (a *IscsiTargetAPIService) ApiIscsiTargetTargetIqnPutExecute(r ApiApiIscsiTargetTargetIqnPutRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IscsiTargetAPIService.ApiIscsiTargetTargetIqnPut")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/iscsi/target/{target_iqn}"
	localVarPath = strings.Replace(localVarPath, "{"+"target_iqn"+"}", url.PathEscape(parameterValueToString(r.targetIqn, "targetIqn")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.ceph.api.v1.0+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.apiIscsiTargetTargetIqnPutRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}
