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


// RgwSiteAPIService RgwSiteAPI service
type RgwSiteAPIService service

type ApiApiRgwSiteGetRequest struct {
	ctx context.Context
	ApiService *RgwSiteAPIService
	query *string
	daemonName *string
}

func (r ApiApiRgwSiteGetRequest) Query(query string) ApiApiRgwSiteGetRequest {
	r.query = &query
	return r
}

func (r ApiApiRgwSiteGetRequest) DaemonName(daemonName string) ApiApiRgwSiteGetRequest {
	r.daemonName = &daemonName
	return r
}

func (r ApiApiRgwSiteGetRequest) Execute() (*http.Response, error) {
	return r.ApiService.ApiRgwSiteGetExecute(r)
}

/*
ApiRgwSiteGet Method for ApiRgwSiteGet

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiApiRgwSiteGetRequest
*/
func (a *RgwSiteAPIService) ApiRgwSiteGet(ctx context.Context) ApiApiRgwSiteGetRequest {
	return ApiApiRgwSiteGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *RgwSiteAPIService) ApiRgwSiteGetExecute(r ApiApiRgwSiteGetRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RgwSiteAPIService.ApiRgwSiteGet")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/rgw/site"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.query != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "query", r.query, "form", "")
	}
	if r.daemonName != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "daemon_name", r.daemonName, "form", "")
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
