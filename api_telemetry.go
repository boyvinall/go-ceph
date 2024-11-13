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


// TelemetryAPIService TelemetryAPI service
type TelemetryAPIService service

type ApiApiTelemetryPutRequest struct {
	ctx context.Context
	ApiService *TelemetryAPIService
	apiTelemetryPutRequest *ApiTelemetryPutRequest
}

func (r ApiApiTelemetryPutRequest) ApiTelemetryPutRequest(apiTelemetryPutRequest ApiTelemetryPutRequest) ApiApiTelemetryPutRequest {
	r.apiTelemetryPutRequest = &apiTelemetryPutRequest
	return r
}

func (r ApiApiTelemetryPutRequest) Execute() (*http.Response, error) {
	return r.ApiService.ApiTelemetryPutExecute(r)
}

/*
ApiTelemetryPut Method for ApiTelemetryPut


        Enables or disables sending data collected by the Telemetry
        module.
        :param enable: Enable or disable sending data
        :type enable: bool
        :param license_name: License string e.g. 'sharing-1-0' to
        make sure the user is aware of and accepts the license
        for sharing Telemetry data.
        :type license_name: string
        

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiApiTelemetryPutRequest
*/
func (a *TelemetryAPIService) ApiTelemetryPut(ctx context.Context) ApiApiTelemetryPutRequest {
	return ApiApiTelemetryPutRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *TelemetryAPIService) ApiTelemetryPutExecute(r ApiApiTelemetryPutRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TelemetryAPIService.ApiTelemetryPut")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/telemetry"

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
	localVarPostBody = r.apiTelemetryPutRequest
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

type ApiApiTelemetryReportGetRequest struct {
	ctx context.Context
	ApiService *TelemetryAPIService
}

func (r ApiApiTelemetryReportGetRequest) Execute() (*ApiTelemetryReportGet200Response, *http.Response, error) {
	return r.ApiService.ApiTelemetryReportGetExecute(r)
}

/*
ApiTelemetryReportGet Get Detailed Telemetry report


        Get Ceph and device report data
        :return: Ceph and device report data
        :rtype: dict
        

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiApiTelemetryReportGetRequest
*/
func (a *TelemetryAPIService) ApiTelemetryReportGet(ctx context.Context) ApiApiTelemetryReportGetRequest {
	return ApiApiTelemetryReportGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ApiTelemetryReportGet200Response
func (a *TelemetryAPIService) ApiTelemetryReportGetExecute(r ApiApiTelemetryReportGetRequest) (*ApiTelemetryReportGet200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ApiTelemetryReportGet200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TelemetryAPIService.ApiTelemetryReportGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/telemetry/report"

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