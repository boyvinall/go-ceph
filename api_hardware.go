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


// HardwareAPIService HardwareAPI service
type HardwareAPIService service

type ApiApiHardwareSummaryGetRequest struct {
	ctx context.Context
	ApiService *HardwareAPIService
	categories *string
	hostname *string
}

func (r ApiApiHardwareSummaryGetRequest) Categories(categories string) ApiApiHardwareSummaryGetRequest {
	r.categories = &categories
	return r
}

func (r ApiApiHardwareSummaryGetRequest) Hostname(hostname string) ApiApiHardwareSummaryGetRequest {
	r.hostname = &hostname
	return r
}

func (r ApiApiHardwareSummaryGetRequest) Execute() (*http.Response, error) {
	return r.ApiService.ApiHardwareSummaryGetExecute(r)
}

/*
ApiHardwareSummaryGet Retrieve a summary of the hardware health status


        Get the health status of as many hardware categories, or all of them if none is given
        :param categories: The hardware type, all of them by default
        :param hostname: The host to retrieve from, all of them by default
        

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiApiHardwareSummaryGetRequest
*/
func (a *HardwareAPIService) ApiHardwareSummaryGet(ctx context.Context) ApiApiHardwareSummaryGetRequest {
	return ApiApiHardwareSummaryGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *HardwareAPIService) ApiHardwareSummaryGetExecute(r ApiApiHardwareSummaryGetRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "HardwareAPIService.ApiHardwareSummaryGet")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/hardware/summary"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.categories != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "categories", r.categories, "form", "")
	}
	if r.hostname != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "hostname", r.hostname, "form", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.ceph.api.v0.1+json"}

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
