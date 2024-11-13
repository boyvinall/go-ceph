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


// ReportAPIService ReportAPI service
type ReportAPIService service

type ApiApiFeedbackApiKeyDeleteRequest struct {
	ctx context.Context
	ApiService *ReportAPIService
}

func (r ApiApiFeedbackApiKeyDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.ApiFeedbackApiKeyDeleteExecute(r)
}

/*
ApiFeedbackApiKeyDelete Method for ApiFeedbackApiKeyDelete


        Deletes Ceph tracker API key.
        

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiApiFeedbackApiKeyDeleteRequest
*/
func (a *ReportAPIService) ApiFeedbackApiKeyDelete(ctx context.Context) ApiApiFeedbackApiKeyDeleteRequest {
	return ApiApiFeedbackApiKeyDeleteRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *ReportAPIService) ApiFeedbackApiKeyDeleteExecute(r ApiApiFeedbackApiKeyDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReportAPIService.ApiFeedbackApiKeyDelete")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/feedback/api_key"

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

type ApiApiFeedbackApiKeyGetRequest struct {
	ctx context.Context
	ApiService *ReportAPIService
}

func (r ApiApiFeedbackApiKeyGetRequest) Execute() (*http.Response, error) {
	return r.ApiService.ApiFeedbackApiKeyGetExecute(r)
}

/*
ApiFeedbackApiKeyGet Method for ApiFeedbackApiKeyGet


        Returns Ceph tracker API key.
        

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiApiFeedbackApiKeyGetRequest
*/
func (a *ReportAPIService) ApiFeedbackApiKeyGet(ctx context.Context) ApiApiFeedbackApiKeyGetRequest {
	return ApiApiFeedbackApiKeyGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *ReportAPIService) ApiFeedbackApiKeyGetExecute(r ApiApiFeedbackApiKeyGetRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReportAPIService.ApiFeedbackApiKeyGet")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/feedback/api_key"

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

type ApiApiFeedbackApiKeyPostRequest struct {
	ctx context.Context
	ApiService *ReportAPIService
	apiFeedbackApiKeyPostRequest *ApiFeedbackApiKeyPostRequest
}

func (r ApiApiFeedbackApiKeyPostRequest) ApiFeedbackApiKeyPostRequest(apiFeedbackApiKeyPostRequest ApiFeedbackApiKeyPostRequest) ApiApiFeedbackApiKeyPostRequest {
	r.apiFeedbackApiKeyPostRequest = &apiFeedbackApiKeyPostRequest
	return r
}

func (r ApiApiFeedbackApiKeyPostRequest) Execute() (*http.Response, error) {
	return r.ApiService.ApiFeedbackApiKeyPostExecute(r)
}

/*
ApiFeedbackApiKeyPost Method for ApiFeedbackApiKeyPost


        Sets Ceph tracker API key.
        :param api_key: The Ceph tracker API key.
        

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiApiFeedbackApiKeyPostRequest
*/
func (a *ReportAPIService) ApiFeedbackApiKeyPost(ctx context.Context) ApiApiFeedbackApiKeyPostRequest {
	return ApiApiFeedbackApiKeyPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *ReportAPIService) ApiFeedbackApiKeyPostExecute(r ApiApiFeedbackApiKeyPostRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReportAPIService.ApiFeedbackApiKeyPost")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/feedback/api_key"

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
	localVarHTTPHeaderAccepts := []string{"application/vnd.ceph.api.v0.1+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.apiFeedbackApiKeyPostRequest
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

type ApiApiFeedbackGetRequest struct {
	ctx context.Context
	ApiService *ReportAPIService
}

func (r ApiApiFeedbackGetRequest) Execute() (*http.Response, error) {
	return r.ApiService.ApiFeedbackGetExecute(r)
}

/*
ApiFeedbackGet Method for ApiFeedbackGet


        List all issues details.
        

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiApiFeedbackGetRequest
*/
func (a *ReportAPIService) ApiFeedbackGet(ctx context.Context) ApiApiFeedbackGetRequest {
	return ApiApiFeedbackGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *ReportAPIService) ApiFeedbackGetExecute(r ApiApiFeedbackGetRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReportAPIService.ApiFeedbackGet")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/feedback"

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

type ApiApiFeedbackPostRequest struct {
	ctx context.Context
	ApiService *ReportAPIService
	apiFeedbackPostRequest *ApiFeedbackPostRequest
}

func (r ApiApiFeedbackPostRequest) ApiFeedbackPostRequest(apiFeedbackPostRequest ApiFeedbackPostRequest) ApiApiFeedbackPostRequest {
	r.apiFeedbackPostRequest = &apiFeedbackPostRequest
	return r
}

func (r ApiApiFeedbackPostRequest) Execute() (*http.Response, error) {
	return r.ApiService.ApiFeedbackPostExecute(r)
}

/*
ApiFeedbackPost Method for ApiFeedbackPost


        Create an issue.
        :param project: The affected ceph component.
        :param tracker: The tracker type.
        :param subject: The title of the issue.
        :param description: The description of the issue.
        :param api_key: Ceph tracker api key.
        

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiApiFeedbackPostRequest
*/
func (a *ReportAPIService) ApiFeedbackPost(ctx context.Context) ApiApiFeedbackPostRequest {
	return ApiApiFeedbackPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *ReportAPIService) ApiFeedbackPostExecute(r ApiApiFeedbackPostRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReportAPIService.ApiFeedbackPost")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/feedback"

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
	localVarHTTPHeaderAccepts := []string{"application/vnd.ceph.api.v0.1+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.apiFeedbackPostRequest
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
