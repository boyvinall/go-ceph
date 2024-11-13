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


// RbdSnapshotAPIService RbdSnapshotAPI service
type RbdSnapshotAPIService service

type ApiApiBlockImageImageSpecSnapPostRequest struct {
	ctx context.Context
	ApiService *RbdSnapshotAPIService
	imageSpec string
	apiBlockImageImageSpecSnapPostRequest *ApiBlockImageImageSpecSnapPostRequest
}

func (r ApiApiBlockImageImageSpecSnapPostRequest) ApiBlockImageImageSpecSnapPostRequest(apiBlockImageImageSpecSnapPostRequest ApiBlockImageImageSpecSnapPostRequest) ApiApiBlockImageImageSpecSnapPostRequest {
	r.apiBlockImageImageSpecSnapPostRequest = &apiBlockImageImageSpecSnapPostRequest
	return r
}

func (r ApiApiBlockImageImageSpecSnapPostRequest) Execute() (*http.Response, error) {
	return r.ApiService.ApiBlockImageImageSpecSnapPostExecute(r)
}

/*
ApiBlockImageImageSpecSnapPost Method for ApiBlockImageImageSpecSnapPost

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param imageSpec
 @return ApiApiBlockImageImageSpecSnapPostRequest
*/
func (a *RbdSnapshotAPIService) ApiBlockImageImageSpecSnapPost(ctx context.Context, imageSpec string) ApiApiBlockImageImageSpecSnapPostRequest {
	return ApiApiBlockImageImageSpecSnapPostRequest{
		ApiService: a,
		ctx: ctx,
		imageSpec: imageSpec,
	}
}

// Execute executes the request
func (a *RbdSnapshotAPIService) ApiBlockImageImageSpecSnapPostExecute(r ApiApiBlockImageImageSpecSnapPostRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RbdSnapshotAPIService.ApiBlockImageImageSpecSnapPost")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/block/image/{image_spec}/snap"
	localVarPath = strings.Replace(localVarPath, "{"+"image_spec"+"}", url.PathEscape(parameterValueToString(r.imageSpec, "imageSpec")), -1)

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
	localVarPostBody = r.apiBlockImageImageSpecSnapPostRequest
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

type ApiApiBlockImageImageSpecSnapSnapshotNameClonePostRequest struct {
	ctx context.Context
	ApiService *RbdSnapshotAPIService
	imageSpec string
	snapshotName string
	apiBlockImageImageSpecSnapSnapshotNameClonePostRequest *ApiBlockImageImageSpecSnapSnapshotNameClonePostRequest
}

func (r ApiApiBlockImageImageSpecSnapSnapshotNameClonePostRequest) ApiBlockImageImageSpecSnapSnapshotNameClonePostRequest(apiBlockImageImageSpecSnapSnapshotNameClonePostRequest ApiBlockImageImageSpecSnapSnapshotNameClonePostRequest) ApiApiBlockImageImageSpecSnapSnapshotNameClonePostRequest {
	r.apiBlockImageImageSpecSnapSnapshotNameClonePostRequest = &apiBlockImageImageSpecSnapSnapshotNameClonePostRequest
	return r
}

func (r ApiApiBlockImageImageSpecSnapSnapshotNameClonePostRequest) Execute() (*http.Response, error) {
	return r.ApiService.ApiBlockImageImageSpecSnapSnapshotNameClonePostExecute(r)
}

/*
ApiBlockImageImageSpecSnapSnapshotNameClonePost Method for ApiBlockImageImageSpecSnapSnapshotNameClonePost


        Clones a snapshot to an image
        

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param imageSpec
 @param snapshotName
 @return ApiApiBlockImageImageSpecSnapSnapshotNameClonePostRequest
*/
func (a *RbdSnapshotAPIService) ApiBlockImageImageSpecSnapSnapshotNameClonePost(ctx context.Context, imageSpec string, snapshotName string) ApiApiBlockImageImageSpecSnapSnapshotNameClonePostRequest {
	return ApiApiBlockImageImageSpecSnapSnapshotNameClonePostRequest{
		ApiService: a,
		ctx: ctx,
		imageSpec: imageSpec,
		snapshotName: snapshotName,
	}
}

// Execute executes the request
func (a *RbdSnapshotAPIService) ApiBlockImageImageSpecSnapSnapshotNameClonePostExecute(r ApiApiBlockImageImageSpecSnapSnapshotNameClonePostRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RbdSnapshotAPIService.ApiBlockImageImageSpecSnapSnapshotNameClonePost")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/block/image/{image_spec}/snap/{snapshot_name}/clone"
	localVarPath = strings.Replace(localVarPath, "{"+"image_spec"+"}", url.PathEscape(parameterValueToString(r.imageSpec, "imageSpec")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"snapshot_name"+"}", url.PathEscape(parameterValueToString(r.snapshotName, "snapshotName")), -1)

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
	localVarPostBody = r.apiBlockImageImageSpecSnapSnapshotNameClonePostRequest
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

type ApiApiBlockImageImageSpecSnapSnapshotNameDeleteRequest struct {
	ctx context.Context
	ApiService *RbdSnapshotAPIService
	imageSpec string
	snapshotName string
}

func (r ApiApiBlockImageImageSpecSnapSnapshotNameDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.ApiBlockImageImageSpecSnapSnapshotNameDeleteExecute(r)
}

/*
ApiBlockImageImageSpecSnapSnapshotNameDelete Method for ApiBlockImageImageSpecSnapSnapshotNameDelete

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param imageSpec
 @param snapshotName
 @return ApiApiBlockImageImageSpecSnapSnapshotNameDeleteRequest
*/
func (a *RbdSnapshotAPIService) ApiBlockImageImageSpecSnapSnapshotNameDelete(ctx context.Context, imageSpec string, snapshotName string) ApiApiBlockImageImageSpecSnapSnapshotNameDeleteRequest {
	return ApiApiBlockImageImageSpecSnapSnapshotNameDeleteRequest{
		ApiService: a,
		ctx: ctx,
		imageSpec: imageSpec,
		snapshotName: snapshotName,
	}
}

// Execute executes the request
func (a *RbdSnapshotAPIService) ApiBlockImageImageSpecSnapSnapshotNameDeleteExecute(r ApiApiBlockImageImageSpecSnapSnapshotNameDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RbdSnapshotAPIService.ApiBlockImageImageSpecSnapSnapshotNameDelete")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/block/image/{image_spec}/snap/{snapshot_name}"
	localVarPath = strings.Replace(localVarPath, "{"+"image_spec"+"}", url.PathEscape(parameterValueToString(r.imageSpec, "imageSpec")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"snapshot_name"+"}", url.PathEscape(parameterValueToString(r.snapshotName, "snapshotName")), -1)

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

type ApiApiBlockImageImageSpecSnapSnapshotNamePutRequest struct {
	ctx context.Context
	ApiService *RbdSnapshotAPIService
	imageSpec string
	snapshotName string
	apiBlockImageImageSpecSnapSnapshotNamePutRequest *ApiBlockImageImageSpecSnapSnapshotNamePutRequest
}

func (r ApiApiBlockImageImageSpecSnapSnapshotNamePutRequest) ApiBlockImageImageSpecSnapSnapshotNamePutRequest(apiBlockImageImageSpecSnapSnapshotNamePutRequest ApiBlockImageImageSpecSnapSnapshotNamePutRequest) ApiApiBlockImageImageSpecSnapSnapshotNamePutRequest {
	r.apiBlockImageImageSpecSnapSnapshotNamePutRequest = &apiBlockImageImageSpecSnapSnapshotNamePutRequest
	return r
}

func (r ApiApiBlockImageImageSpecSnapSnapshotNamePutRequest) Execute() (*http.Response, error) {
	return r.ApiService.ApiBlockImageImageSpecSnapSnapshotNamePutExecute(r)
}

/*
ApiBlockImageImageSpecSnapSnapshotNamePut Method for ApiBlockImageImageSpecSnapSnapshotNamePut

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param imageSpec
 @param snapshotName
 @return ApiApiBlockImageImageSpecSnapSnapshotNamePutRequest
*/
func (a *RbdSnapshotAPIService) ApiBlockImageImageSpecSnapSnapshotNamePut(ctx context.Context, imageSpec string, snapshotName string) ApiApiBlockImageImageSpecSnapSnapshotNamePutRequest {
	return ApiApiBlockImageImageSpecSnapSnapshotNamePutRequest{
		ApiService: a,
		ctx: ctx,
		imageSpec: imageSpec,
		snapshotName: snapshotName,
	}
}

// Execute executes the request
func (a *RbdSnapshotAPIService) ApiBlockImageImageSpecSnapSnapshotNamePutExecute(r ApiApiBlockImageImageSpecSnapSnapshotNamePutRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RbdSnapshotAPIService.ApiBlockImageImageSpecSnapSnapshotNamePut")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/block/image/{image_spec}/snap/{snapshot_name}"
	localVarPath = strings.Replace(localVarPath, "{"+"image_spec"+"}", url.PathEscape(parameterValueToString(r.imageSpec, "imageSpec")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"snapshot_name"+"}", url.PathEscape(parameterValueToString(r.snapshotName, "snapshotName")), -1)

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
	localVarPostBody = r.apiBlockImageImageSpecSnapSnapshotNamePutRequest
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

type ApiApiBlockImageImageSpecSnapSnapshotNameRollbackPostRequest struct {
	ctx context.Context
	ApiService *RbdSnapshotAPIService
	imageSpec string
	snapshotName string
}

func (r ApiApiBlockImageImageSpecSnapSnapshotNameRollbackPostRequest) Execute() (*http.Response, error) {
	return r.ApiService.ApiBlockImageImageSpecSnapSnapshotNameRollbackPostExecute(r)
}

/*
ApiBlockImageImageSpecSnapSnapshotNameRollbackPost Method for ApiBlockImageImageSpecSnapSnapshotNameRollbackPost

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param imageSpec
 @param snapshotName
 @return ApiApiBlockImageImageSpecSnapSnapshotNameRollbackPostRequest
*/
func (a *RbdSnapshotAPIService) ApiBlockImageImageSpecSnapSnapshotNameRollbackPost(ctx context.Context, imageSpec string, snapshotName string) ApiApiBlockImageImageSpecSnapSnapshotNameRollbackPostRequest {
	return ApiApiBlockImageImageSpecSnapSnapshotNameRollbackPostRequest{
		ApiService: a,
		ctx: ctx,
		imageSpec: imageSpec,
		snapshotName: snapshotName,
	}
}

// Execute executes the request
func (a *RbdSnapshotAPIService) ApiBlockImageImageSpecSnapSnapshotNameRollbackPostExecute(r ApiApiBlockImageImageSpecSnapSnapshotNameRollbackPostRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RbdSnapshotAPIService.ApiBlockImageImageSpecSnapSnapshotNameRollbackPost")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/block/image/{image_spec}/snap/{snapshot_name}/rollback"
	localVarPath = strings.Replace(localVarPath, "{"+"image_spec"+"}", url.PathEscape(parameterValueToString(r.imageSpec, "imageSpec")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"snapshot_name"+"}", url.PathEscape(parameterValueToString(r.snapshotName, "snapshotName")), -1)

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
