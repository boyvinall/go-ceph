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


// RbdTrashAPIService RbdTrashAPI service
type RbdTrashAPIService service

type ApiApiBlockImageTrashGetRequest struct {
	ctx context.Context
	ApiService *RbdTrashAPIService
	poolName *string
}

// Name of the pool
func (r ApiApiBlockImageTrashGetRequest) PoolName(poolName string) ApiApiBlockImageTrashGetRequest {
	r.poolName = &poolName
	return r
}

func (r ApiApiBlockImageTrashGetRequest) Execute() ([]ApiBlockImageTrashGet200ResponseInner, *http.Response, error) {
	return r.ApiService.ApiBlockImageTrashGetExecute(r)
}

/*
ApiBlockImageTrashGet Get RBD Trash Details by pool name

List all entries from trash.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiApiBlockImageTrashGetRequest
*/
func (a *RbdTrashAPIService) ApiBlockImageTrashGet(ctx context.Context) ApiApiBlockImageTrashGetRequest {
	return ApiApiBlockImageTrashGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []ApiBlockImageTrashGet200ResponseInner
func (a *RbdTrashAPIService) ApiBlockImageTrashGetExecute(r ApiApiBlockImageTrashGetRequest) ([]ApiBlockImageTrashGet200ResponseInner, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []ApiBlockImageTrashGet200ResponseInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RbdTrashAPIService.ApiBlockImageTrashGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/block/image/trash"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.poolName != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pool_name", r.poolName, "form", "")
	}
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
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
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

type ApiApiBlockImageTrashImageIdSpecDeleteRequest struct {
	ctx context.Context
	ApiService *RbdTrashAPIService
	imageIdSpec string
	force *bool
}

func (r ApiApiBlockImageTrashImageIdSpecDeleteRequest) Force(force bool) ApiApiBlockImageTrashImageIdSpecDeleteRequest {
	r.force = &force
	return r
}

func (r ApiApiBlockImageTrashImageIdSpecDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.ApiBlockImageTrashImageIdSpecDeleteExecute(r)
}

/*
ApiBlockImageTrashImageIdSpecDelete Method for ApiBlockImageTrashImageIdSpecDelete

Delete an image from trash.
        If image deferment time has not expired you can not removed it unless use force.
        But an actively in-use by clones or has snapshots can not be removed.
        

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param imageIdSpec
 @return ApiApiBlockImageTrashImageIdSpecDeleteRequest
*/
func (a *RbdTrashAPIService) ApiBlockImageTrashImageIdSpecDelete(ctx context.Context, imageIdSpec string) ApiApiBlockImageTrashImageIdSpecDeleteRequest {
	return ApiApiBlockImageTrashImageIdSpecDeleteRequest{
		ApiService: a,
		ctx: ctx,
		imageIdSpec: imageIdSpec,
	}
}

// Execute executes the request
func (a *RbdTrashAPIService) ApiBlockImageTrashImageIdSpecDeleteExecute(r ApiApiBlockImageTrashImageIdSpecDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RbdTrashAPIService.ApiBlockImageTrashImageIdSpecDelete")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/block/image/trash/{image_id_spec}"
	localVarPath = strings.Replace(localVarPath, "{"+"image_id_spec"+"}", url.PathEscape(parameterValueToString(r.imageIdSpec, "imageIdSpec")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.force != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "force", r.force, "form", "")
	}
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

type ApiApiBlockImageTrashImageIdSpecRestorePostRequest struct {
	ctx context.Context
	ApiService *RbdTrashAPIService
	imageIdSpec string
	apiBlockImageTrashImageIdSpecRestorePostRequest *ApiBlockImageTrashImageIdSpecRestorePostRequest
}

func (r ApiApiBlockImageTrashImageIdSpecRestorePostRequest) ApiBlockImageTrashImageIdSpecRestorePostRequest(apiBlockImageTrashImageIdSpecRestorePostRequest ApiBlockImageTrashImageIdSpecRestorePostRequest) ApiApiBlockImageTrashImageIdSpecRestorePostRequest {
	r.apiBlockImageTrashImageIdSpecRestorePostRequest = &apiBlockImageTrashImageIdSpecRestorePostRequest
	return r
}

func (r ApiApiBlockImageTrashImageIdSpecRestorePostRequest) Execute() (*http.Response, error) {
	return r.ApiService.ApiBlockImageTrashImageIdSpecRestorePostExecute(r)
}

/*
ApiBlockImageTrashImageIdSpecRestorePost Method for ApiBlockImageTrashImageIdSpecRestorePost

Restore an image from trash.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param imageIdSpec
 @return ApiApiBlockImageTrashImageIdSpecRestorePostRequest
*/
func (a *RbdTrashAPIService) ApiBlockImageTrashImageIdSpecRestorePost(ctx context.Context, imageIdSpec string) ApiApiBlockImageTrashImageIdSpecRestorePostRequest {
	return ApiApiBlockImageTrashImageIdSpecRestorePostRequest{
		ApiService: a,
		ctx: ctx,
		imageIdSpec: imageIdSpec,
	}
}

// Execute executes the request
func (a *RbdTrashAPIService) ApiBlockImageTrashImageIdSpecRestorePostExecute(r ApiApiBlockImageTrashImageIdSpecRestorePostRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RbdTrashAPIService.ApiBlockImageTrashImageIdSpecRestorePost")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/block/image/trash/{image_id_spec}/restore"
	localVarPath = strings.Replace(localVarPath, "{"+"image_id_spec"+"}", url.PathEscape(parameterValueToString(r.imageIdSpec, "imageIdSpec")), -1)

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
	localVarPostBody = r.apiBlockImageTrashImageIdSpecRestorePostRequest
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

type ApiApiBlockImageTrashPurgePostRequest struct {
	ctx context.Context
	ApiService *RbdTrashAPIService
	poolName *string
	apiBlockImageTrashPurgePostRequest *ApiBlockImageTrashPurgePostRequest
}

func (r ApiApiBlockImageTrashPurgePostRequest) PoolName(poolName string) ApiApiBlockImageTrashPurgePostRequest {
	r.poolName = &poolName
	return r
}

func (r ApiApiBlockImageTrashPurgePostRequest) ApiBlockImageTrashPurgePostRequest(apiBlockImageTrashPurgePostRequest ApiBlockImageTrashPurgePostRequest) ApiApiBlockImageTrashPurgePostRequest {
	r.apiBlockImageTrashPurgePostRequest = &apiBlockImageTrashPurgePostRequest
	return r
}

func (r ApiApiBlockImageTrashPurgePostRequest) Execute() (*http.Response, error) {
	return r.ApiService.ApiBlockImageTrashPurgePostExecute(r)
}

/*
ApiBlockImageTrashPurgePost Method for ApiBlockImageTrashPurgePost

Remove all expired images from trash.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiApiBlockImageTrashPurgePostRequest
*/
func (a *RbdTrashAPIService) ApiBlockImageTrashPurgePost(ctx context.Context) ApiApiBlockImageTrashPurgePostRequest {
	return ApiApiBlockImageTrashPurgePostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *RbdTrashAPIService) ApiBlockImageTrashPurgePostExecute(r ApiApiBlockImageTrashPurgePostRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RbdTrashAPIService.ApiBlockImageTrashPurgePost")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/block/image/trash/purge"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.poolName != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pool_name", r.poolName, "form", "")
	}
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
	localVarPostBody = r.apiBlockImageTrashPurgePostRequest
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
