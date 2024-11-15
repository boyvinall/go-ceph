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


// OsdPerfCounterAPIService OsdPerfCounterAPI service
type OsdPerfCounterAPIService service

type ApiApiPerfCountersOsdServiceIdGetRequest struct {
	ctx context.Context
	ApiService *OsdPerfCounterAPIService
	serviceId string
}

func (r ApiApiPerfCountersOsdServiceIdGetRequest) Execute() (*http.Response, error) {
	return r.ApiService.ApiPerfCountersOsdServiceIdGetExecute(r)
}

/*
ApiPerfCountersOsdServiceIdGet Method for ApiPerfCountersOsdServiceIdGet

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceId
 @return ApiApiPerfCountersOsdServiceIdGetRequest
*/
func (a *OsdPerfCounterAPIService) ApiPerfCountersOsdServiceIdGet(ctx context.Context, serviceId string) ApiApiPerfCountersOsdServiceIdGetRequest {
	return ApiApiPerfCountersOsdServiceIdGetRequest{
		ApiService: a,
		ctx: ctx,
		serviceId: serviceId,
	}
}

// Execute executes the request
func (a *OsdPerfCounterAPIService) ApiPerfCountersOsdServiceIdGetExecute(r ApiApiPerfCountersOsdServiceIdGetRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OsdPerfCounterAPIService.ApiPerfCountersOsdServiceIdGet")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/perf_counters/osd/{service_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"service_id"+"}", url.PathEscape(parameterValueToString(r.serviceId, "serviceId")), -1)

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
