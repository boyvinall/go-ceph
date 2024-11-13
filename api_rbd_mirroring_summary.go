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


// RbdMirroringSummaryAPIService RbdMirroringSummaryAPI service
type RbdMirroringSummaryAPIService service

type ApiApiBlockMirroringSummaryGetRequest struct {
	ctx context.Context
	ApiService *RbdMirroringSummaryAPIService
}

func (r ApiApiBlockMirroringSummaryGetRequest) Execute() (*ApiBlockMirroringSummaryGet200Response, *http.Response, error) {
	return r.ApiService.ApiBlockMirroringSummaryGetExecute(r)
}

/*
ApiBlockMirroringSummaryGet Display Rbd Mirroring Summary

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiApiBlockMirroringSummaryGetRequest
*/
func (a *RbdMirroringSummaryAPIService) ApiBlockMirroringSummaryGet(ctx context.Context) ApiApiBlockMirroringSummaryGetRequest {
	return ApiApiBlockMirroringSummaryGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ApiBlockMirroringSummaryGet200Response
func (a *RbdMirroringSummaryAPIService) ApiBlockMirroringSummaryGetExecute(r ApiApiBlockMirroringSummaryGetRequest) (*ApiBlockMirroringSummaryGet200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ApiBlockMirroringSummaryGet200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RbdMirroringSummaryAPIService.ApiBlockMirroringSummaryGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/block/mirroring/summary"

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
