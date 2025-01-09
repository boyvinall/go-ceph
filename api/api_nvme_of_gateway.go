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
)


// NVMeOFGatewayAPIService NVMeOFGatewayAPI service
type NVMeOFGatewayAPIService service

type ApiApiNvmeofGatewayGetRequest struct {
	ctx context.Context
	ApiService *NVMeOFGatewayAPIService
	gwGroup *string
}

func (r ApiApiNvmeofGatewayGetRequest) GwGroup(gwGroup string) ApiApiNvmeofGatewayGetRequest {
	r.gwGroup = &gwGroup
	return r
}

func (r ApiApiNvmeofGatewayGetRequest) Execute() (*http.Response, error) {
	return r.ApiService.ApiNvmeofGatewayGetExecute(r)
}

/*
ApiNvmeofGatewayGet Get information about the NVMeoF gateway

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiApiNvmeofGatewayGetRequest
*/
func (a *NVMeOFGatewayAPIService) ApiNvmeofGatewayGet(ctx context.Context) ApiApiNvmeofGatewayGetRequest {
	return ApiApiNvmeofGatewayGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *NVMeOFGatewayAPIService) ApiNvmeofGatewayGetExecute(r ApiApiNvmeofGatewayGetRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NVMeOFGatewayAPIService.ApiNvmeofGatewayGet")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/nvmeof/gateway"

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

type ApiApiNvmeofGatewayGroupGetRequest struct {
	ctx context.Context
	ApiService *NVMeOFGatewayAPIService
}

func (r ApiApiNvmeofGatewayGroupGetRequest) Execute() (*http.Response, error) {
	return r.ApiService.ApiNvmeofGatewayGroupGetExecute(r)
}

/*
ApiNvmeofGatewayGroupGet Method for ApiNvmeofGatewayGroupGet

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiApiNvmeofGatewayGroupGetRequest
*/
func (a *NVMeOFGatewayAPIService) ApiNvmeofGatewayGroupGet(ctx context.Context) ApiApiNvmeofGatewayGroupGetRequest {
	return ApiApiNvmeofGatewayGroupGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *NVMeOFGatewayAPIService) ApiNvmeofGatewayGroupGetExecute(r ApiApiNvmeofGatewayGroupGetRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NVMeOFGatewayAPIService.ApiNvmeofGatewayGroupGet")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/nvmeof/gateway/group"

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