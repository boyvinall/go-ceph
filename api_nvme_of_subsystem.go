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


// NVMeOFSubsystemAPIService NVMeOFSubsystemAPI service
type NVMeOFSubsystemAPIService service

type ApiApiNvmeofSubsystemGetRequest struct {
	ctx context.Context
	ApiService *NVMeOFSubsystemAPIService
	gwGroup *string
}

func (r ApiApiNvmeofSubsystemGetRequest) GwGroup(gwGroup string) ApiApiNvmeofSubsystemGetRequest {
	r.gwGroup = &gwGroup
	return r
}

func (r ApiApiNvmeofSubsystemGetRequest) Execute() (*http.Response, error) {
	return r.ApiService.ApiNvmeofSubsystemGetExecute(r)
}

/*
ApiNvmeofSubsystemGet List all NVMeoF subsystems

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiApiNvmeofSubsystemGetRequest
*/
func (a *NVMeOFSubsystemAPIService) ApiNvmeofSubsystemGet(ctx context.Context) ApiApiNvmeofSubsystemGetRequest {
	return ApiApiNvmeofSubsystemGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *NVMeOFSubsystemAPIService) ApiNvmeofSubsystemGetExecute(r ApiApiNvmeofSubsystemGetRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NVMeOFSubsystemAPIService.ApiNvmeofSubsystemGet")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/nvmeof/subsystem"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.gwGroup != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "gw_group", r.gwGroup, "form", "")
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

type ApiApiNvmeofSubsystemNqnDeleteRequest struct {
	ctx context.Context
	ApiService *NVMeOFSubsystemAPIService
	nqn string
	force *bool
	gwGroup *string
}

// Force delete
func (r ApiApiNvmeofSubsystemNqnDeleteRequest) Force(force bool) ApiApiNvmeofSubsystemNqnDeleteRequest {
	r.force = &force
	return r
}

// NVMeoF gateway group
func (r ApiApiNvmeofSubsystemNqnDeleteRequest) GwGroup(gwGroup string) ApiApiNvmeofSubsystemNqnDeleteRequest {
	r.gwGroup = &gwGroup
	return r
}

func (r ApiApiNvmeofSubsystemNqnDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.ApiNvmeofSubsystemNqnDeleteExecute(r)
}

/*
ApiNvmeofSubsystemNqnDelete Delete an existing NVMeoF subsystem

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param nqn NVMeoF subsystem NQN
 @return ApiApiNvmeofSubsystemNqnDeleteRequest
*/
func (a *NVMeOFSubsystemAPIService) ApiNvmeofSubsystemNqnDelete(ctx context.Context, nqn string) ApiApiNvmeofSubsystemNqnDeleteRequest {
	return ApiApiNvmeofSubsystemNqnDeleteRequest{
		ApiService: a,
		ctx: ctx,
		nqn: nqn,
	}
}

// Execute executes the request
func (a *NVMeOFSubsystemAPIService) ApiNvmeofSubsystemNqnDeleteExecute(r ApiApiNvmeofSubsystemNqnDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NVMeOFSubsystemAPIService.ApiNvmeofSubsystemNqnDelete")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/nvmeof/subsystem/{nqn}"
	localVarPath = strings.Replace(localVarPath, "{"+"nqn"+"}", url.PathEscape(parameterValueToString(r.nqn, "nqn")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.force != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "force", r.force, "form", "")
	}
	if r.gwGroup != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "gw_group", r.gwGroup, "form", "")
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

type ApiApiNvmeofSubsystemNqnGetRequest struct {
	ctx context.Context
	ApiService *NVMeOFSubsystemAPIService
	nqn string
	gwGroup *string
}

// NVMeoF gateway group
func (r ApiApiNvmeofSubsystemNqnGetRequest) GwGroup(gwGroup string) ApiApiNvmeofSubsystemNqnGetRequest {
	r.gwGroup = &gwGroup
	return r
}

func (r ApiApiNvmeofSubsystemNqnGetRequest) Execute() (*http.Response, error) {
	return r.ApiService.ApiNvmeofSubsystemNqnGetExecute(r)
}

/*
ApiNvmeofSubsystemNqnGet Get information from a specific NVMeoF subsystem

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param nqn NVMeoF subsystem NQN
 @return ApiApiNvmeofSubsystemNqnGetRequest
*/
func (a *NVMeOFSubsystemAPIService) ApiNvmeofSubsystemNqnGet(ctx context.Context, nqn string) ApiApiNvmeofSubsystemNqnGetRequest {
	return ApiApiNvmeofSubsystemNqnGetRequest{
		ApiService: a,
		ctx: ctx,
		nqn: nqn,
	}
}

// Execute executes the request
func (a *NVMeOFSubsystemAPIService) ApiNvmeofSubsystemNqnGetExecute(r ApiApiNvmeofSubsystemNqnGetRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NVMeOFSubsystemAPIService.ApiNvmeofSubsystemNqnGet")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/nvmeof/subsystem/{nqn}"
	localVarPath = strings.Replace(localVarPath, "{"+"nqn"+"}", url.PathEscape(parameterValueToString(r.nqn, "nqn")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.gwGroup != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "gw_group", r.gwGroup, "form", "")
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

type ApiApiNvmeofSubsystemPostRequest struct {
	ctx context.Context
	ApiService *NVMeOFSubsystemAPIService
	apiNvmeofSubsystemPostRequest *ApiNvmeofSubsystemPostRequest
}

func (r ApiApiNvmeofSubsystemPostRequest) ApiNvmeofSubsystemPostRequest(apiNvmeofSubsystemPostRequest ApiNvmeofSubsystemPostRequest) ApiApiNvmeofSubsystemPostRequest {
	r.apiNvmeofSubsystemPostRequest = &apiNvmeofSubsystemPostRequest
	return r
}

func (r ApiApiNvmeofSubsystemPostRequest) Execute() (*http.Response, error) {
	return r.ApiService.ApiNvmeofSubsystemPostExecute(r)
}

/*
ApiNvmeofSubsystemPost Create a new NVMeoF subsystem

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiApiNvmeofSubsystemPostRequest
*/
func (a *NVMeOFSubsystemAPIService) ApiNvmeofSubsystemPost(ctx context.Context) ApiApiNvmeofSubsystemPostRequest {
	return ApiApiNvmeofSubsystemPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *NVMeOFSubsystemAPIService) ApiNvmeofSubsystemPostExecute(r ApiApiNvmeofSubsystemPostRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NVMeOFSubsystemAPIService.ApiNvmeofSubsystemPost")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/nvmeof/subsystem"

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
	localVarPostBody = r.apiNvmeofSubsystemPostRequest
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
