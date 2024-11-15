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


// RgwZoneAPIService RgwZoneAPI service
type RgwZoneAPIService service

type ApiApiRgwZoneCreateSystemUserPutRequest struct {
	ctx context.Context
	ApiService *RgwZoneAPIService
	apiRgwZoneCreateSystemUserPutRequest *ApiRgwZoneCreateSystemUserPutRequest
}

func (r ApiApiRgwZoneCreateSystemUserPutRequest) ApiRgwZoneCreateSystemUserPutRequest(apiRgwZoneCreateSystemUserPutRequest ApiRgwZoneCreateSystemUserPutRequest) ApiApiRgwZoneCreateSystemUserPutRequest {
	r.apiRgwZoneCreateSystemUserPutRequest = &apiRgwZoneCreateSystemUserPutRequest
	return r
}

func (r ApiApiRgwZoneCreateSystemUserPutRequest) Execute() (*http.Response, error) {
	return r.ApiService.ApiRgwZoneCreateSystemUserPutExecute(r)
}

/*
ApiRgwZoneCreateSystemUserPut Method for ApiRgwZoneCreateSystemUserPut

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiApiRgwZoneCreateSystemUserPutRequest
*/
func (a *RgwZoneAPIService) ApiRgwZoneCreateSystemUserPut(ctx context.Context) ApiApiRgwZoneCreateSystemUserPutRequest {
	return ApiApiRgwZoneCreateSystemUserPutRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *RgwZoneAPIService) ApiRgwZoneCreateSystemUserPutExecute(r ApiApiRgwZoneCreateSystemUserPutRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RgwZoneAPIService.ApiRgwZoneCreateSystemUserPut")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/rgw/zone/create_system_user"

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
	localVarPostBody = r.apiRgwZoneCreateSystemUserPutRequest
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

type ApiApiRgwZoneGetRequest struct {
	ctx context.Context
	ApiService *RgwZoneAPIService
}

func (r ApiApiRgwZoneGetRequest) Execute() (*http.Response, error) {
	return r.ApiService.ApiRgwZoneGetExecute(r)
}

/*
ApiRgwZoneGet Method for ApiRgwZoneGet

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiApiRgwZoneGetRequest
*/
func (a *RgwZoneAPIService) ApiRgwZoneGet(ctx context.Context) ApiApiRgwZoneGetRequest {
	return ApiApiRgwZoneGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *RgwZoneAPIService) ApiRgwZoneGetExecute(r ApiApiRgwZoneGetRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RgwZoneAPIService.ApiRgwZoneGet")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/rgw/zone"

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

type ApiApiRgwZoneGetAllZonesInfoGetRequest struct {
	ctx context.Context
	ApiService *RgwZoneAPIService
}

func (r ApiApiRgwZoneGetAllZonesInfoGetRequest) Execute() (*http.Response, error) {
	return r.ApiService.ApiRgwZoneGetAllZonesInfoGetExecute(r)
}

/*
ApiRgwZoneGetAllZonesInfoGet Method for ApiRgwZoneGetAllZonesInfoGet

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiApiRgwZoneGetAllZonesInfoGetRequest
*/
func (a *RgwZoneAPIService) ApiRgwZoneGetAllZonesInfoGet(ctx context.Context) ApiApiRgwZoneGetAllZonesInfoGetRequest {
	return ApiApiRgwZoneGetAllZonesInfoGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *RgwZoneAPIService) ApiRgwZoneGetAllZonesInfoGetExecute(r ApiApiRgwZoneGetAllZonesInfoGetRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RgwZoneAPIService.ApiRgwZoneGetAllZonesInfoGet")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/rgw/zone/get_all_zones_info"

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

type ApiApiRgwZoneGetPoolNamesGetRequest struct {
	ctx context.Context
	ApiService *RgwZoneAPIService
}

func (r ApiApiRgwZoneGetPoolNamesGetRequest) Execute() (*http.Response, error) {
	return r.ApiService.ApiRgwZoneGetPoolNamesGetExecute(r)
}

/*
ApiRgwZoneGetPoolNamesGet Method for ApiRgwZoneGetPoolNamesGet

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiApiRgwZoneGetPoolNamesGetRequest
*/
func (a *RgwZoneAPIService) ApiRgwZoneGetPoolNamesGet(ctx context.Context) ApiApiRgwZoneGetPoolNamesGetRequest {
	return ApiApiRgwZoneGetPoolNamesGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *RgwZoneAPIService) ApiRgwZoneGetPoolNamesGetExecute(r ApiApiRgwZoneGetPoolNamesGetRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RgwZoneAPIService.ApiRgwZoneGetPoolNamesGet")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/rgw/zone/get_pool_names"

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

type ApiApiRgwZoneGetUserListGetRequest struct {
	ctx context.Context
	ApiService *RgwZoneAPIService
	zoneName *string
}

func (r ApiApiRgwZoneGetUserListGetRequest) ZoneName(zoneName string) ApiApiRgwZoneGetUserListGetRequest {
	r.zoneName = &zoneName
	return r
}

func (r ApiApiRgwZoneGetUserListGetRequest) Execute() (*http.Response, error) {
	return r.ApiService.ApiRgwZoneGetUserListGetExecute(r)
}

/*
ApiRgwZoneGetUserListGet Method for ApiRgwZoneGetUserListGet

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiApiRgwZoneGetUserListGetRequest
*/
func (a *RgwZoneAPIService) ApiRgwZoneGetUserListGet(ctx context.Context) ApiApiRgwZoneGetUserListGetRequest {
	return ApiApiRgwZoneGetUserListGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *RgwZoneAPIService) ApiRgwZoneGetUserListGetExecute(r ApiApiRgwZoneGetUserListGetRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RgwZoneAPIService.ApiRgwZoneGetUserListGet")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/rgw/zone/get_user_list"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.zoneName != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "zoneName", r.zoneName, "form", "")
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

type ApiApiRgwZonePostRequest struct {
	ctx context.Context
	ApiService *RgwZoneAPIService
	apiRgwZonePostRequest *ApiRgwZonePostRequest
}

func (r ApiApiRgwZonePostRequest) ApiRgwZonePostRequest(apiRgwZonePostRequest ApiRgwZonePostRequest) ApiApiRgwZonePostRequest {
	r.apiRgwZonePostRequest = &apiRgwZonePostRequest
	return r
}

func (r ApiApiRgwZonePostRequest) Execute() (*http.Response, error) {
	return r.ApiService.ApiRgwZonePostExecute(r)
}

/*
ApiRgwZonePost Method for ApiRgwZonePost

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiApiRgwZonePostRequest
*/
func (a *RgwZoneAPIService) ApiRgwZonePost(ctx context.Context) ApiApiRgwZonePostRequest {
	return ApiApiRgwZonePostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *RgwZoneAPIService) ApiRgwZonePostExecute(r ApiApiRgwZonePostRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RgwZoneAPIService.ApiRgwZonePost")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/rgw/zone"

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
	localVarPostBody = r.apiRgwZonePostRequest
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

type ApiApiRgwZoneZoneNameDeleteRequest struct {
	ctx context.Context
	ApiService *RgwZoneAPIService
	zoneName string
	deletePools *string
	pools *string
	zonegroupName *string
}

func (r ApiApiRgwZoneZoneNameDeleteRequest) DeletePools(deletePools string) ApiApiRgwZoneZoneNameDeleteRequest {
	r.deletePools = &deletePools
	return r
}

func (r ApiApiRgwZoneZoneNameDeleteRequest) Pools(pools string) ApiApiRgwZoneZoneNameDeleteRequest {
	r.pools = &pools
	return r
}

func (r ApiApiRgwZoneZoneNameDeleteRequest) ZonegroupName(zonegroupName string) ApiApiRgwZoneZoneNameDeleteRequest {
	r.zonegroupName = &zonegroupName
	return r
}

func (r ApiApiRgwZoneZoneNameDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.ApiRgwZoneZoneNameDeleteExecute(r)
}

/*
ApiRgwZoneZoneNameDelete Method for ApiRgwZoneZoneNameDelete

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param zoneName
 @return ApiApiRgwZoneZoneNameDeleteRequest
*/
func (a *RgwZoneAPIService) ApiRgwZoneZoneNameDelete(ctx context.Context, zoneName string) ApiApiRgwZoneZoneNameDeleteRequest {
	return ApiApiRgwZoneZoneNameDeleteRequest{
		ApiService: a,
		ctx: ctx,
		zoneName: zoneName,
	}
}

// Execute executes the request
func (a *RgwZoneAPIService) ApiRgwZoneZoneNameDeleteExecute(r ApiApiRgwZoneZoneNameDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RgwZoneAPIService.ApiRgwZoneZoneNameDelete")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/rgw/zone/{zone_name}"
	localVarPath = strings.Replace(localVarPath, "{"+"zone_name"+"}", url.PathEscape(parameterValueToString(r.zoneName, "zoneName")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.deletePools == nil {
		return nil, reportError("deletePools is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "delete_pools", r.deletePools, "form", "")
	if r.pools != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pools", r.pools, "form", "")
	}
	if r.zonegroupName != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "zonegroup_name", r.zonegroupName, "form", "")
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

type ApiApiRgwZoneZoneNameGetRequest struct {
	ctx context.Context
	ApiService *RgwZoneAPIService
	zoneName string
}

func (r ApiApiRgwZoneZoneNameGetRequest) Execute() (*http.Response, error) {
	return r.ApiService.ApiRgwZoneZoneNameGetExecute(r)
}

/*
ApiRgwZoneZoneNameGet Method for ApiRgwZoneZoneNameGet

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param zoneName
 @return ApiApiRgwZoneZoneNameGetRequest
*/
func (a *RgwZoneAPIService) ApiRgwZoneZoneNameGet(ctx context.Context, zoneName string) ApiApiRgwZoneZoneNameGetRequest {
	return ApiApiRgwZoneZoneNameGetRequest{
		ApiService: a,
		ctx: ctx,
		zoneName: zoneName,
	}
}

// Execute executes the request
func (a *RgwZoneAPIService) ApiRgwZoneZoneNameGetExecute(r ApiApiRgwZoneZoneNameGetRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RgwZoneAPIService.ApiRgwZoneZoneNameGet")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/rgw/zone/{zone_name}"
	localVarPath = strings.Replace(localVarPath, "{"+"zone_name"+"}", url.PathEscape(parameterValueToString(r.zoneName, "zoneName")), -1)

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

type ApiApiRgwZoneZoneNamePutRequest struct {
	ctx context.Context
	ApiService *RgwZoneAPIService
	zoneName string
	apiRgwZoneZoneNamePutRequest *ApiRgwZoneZoneNamePutRequest
}

func (r ApiApiRgwZoneZoneNamePutRequest) ApiRgwZoneZoneNamePutRequest(apiRgwZoneZoneNamePutRequest ApiRgwZoneZoneNamePutRequest) ApiApiRgwZoneZoneNamePutRequest {
	r.apiRgwZoneZoneNamePutRequest = &apiRgwZoneZoneNamePutRequest
	return r
}

func (r ApiApiRgwZoneZoneNamePutRequest) Execute() (*http.Response, error) {
	return r.ApiService.ApiRgwZoneZoneNamePutExecute(r)
}

/*
ApiRgwZoneZoneNamePut Method for ApiRgwZoneZoneNamePut

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param zoneName
 @return ApiApiRgwZoneZoneNamePutRequest
*/
func (a *RgwZoneAPIService) ApiRgwZoneZoneNamePut(ctx context.Context, zoneName string) ApiApiRgwZoneZoneNamePutRequest {
	return ApiApiRgwZoneZoneNamePutRequest{
		ApiService: a,
		ctx: ctx,
		zoneName: zoneName,
	}
}

// Execute executes the request
func (a *RgwZoneAPIService) ApiRgwZoneZoneNamePutExecute(r ApiApiRgwZoneZoneNamePutRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RgwZoneAPIService.ApiRgwZoneZoneNamePut")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/rgw/zone/{zone_name}"
	localVarPath = strings.Replace(localVarPath, "{"+"zone_name"+"}", url.PathEscape(parameterValueToString(r.zoneName, "zoneName")), -1)

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
	localVarPostBody = r.apiRgwZoneZoneNamePutRequest
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
