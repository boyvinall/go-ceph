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


// IscsiAPIService IscsiAPI service
type IscsiAPIService service

type ApiApiIscsiDiscoveryauthGetRequest struct {
	ctx context.Context
	ApiService *IscsiAPIService
}

func (r ApiApiIscsiDiscoveryauthGetRequest) Execute() ([]ApiIscsiDiscoveryauthGet200ResponseInner, *http.Response, error) {
	return r.ApiService.ApiIscsiDiscoveryauthGetExecute(r)
}

/*
ApiIscsiDiscoveryauthGet Get Iscsi discoveryauth Details

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiApiIscsiDiscoveryauthGetRequest
*/
func (a *IscsiAPIService) ApiIscsiDiscoveryauthGet(ctx context.Context) ApiApiIscsiDiscoveryauthGetRequest {
	return ApiApiIscsiDiscoveryauthGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []ApiIscsiDiscoveryauthGet200ResponseInner
func (a *IscsiAPIService) ApiIscsiDiscoveryauthGetExecute(r ApiApiIscsiDiscoveryauthGetRequest) ([]ApiIscsiDiscoveryauthGet200ResponseInner, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []ApiIscsiDiscoveryauthGet200ResponseInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IscsiAPIService.ApiIscsiDiscoveryauthGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/iscsi/discoveryauth"

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

type ApiApiIscsiDiscoveryauthPutRequest struct {
	ctx context.Context
	ApiService *IscsiAPIService
	user *string
	password *string
	mutualUser *string
	mutualPassword *string
	apiIscsiDiscoveryauthPutRequest *ApiIscsiDiscoveryauthPutRequest
}

// Username
func (r ApiApiIscsiDiscoveryauthPutRequest) User(user string) ApiApiIscsiDiscoveryauthPutRequest {
	r.user = &user
	return r
}

// Password
func (r ApiApiIscsiDiscoveryauthPutRequest) Password(password string) ApiApiIscsiDiscoveryauthPutRequest {
	r.password = &password
	return r
}

// Mutual UserName
func (r ApiApiIscsiDiscoveryauthPutRequest) MutualUser(mutualUser string) ApiApiIscsiDiscoveryauthPutRequest {
	r.mutualUser = &mutualUser
	return r
}

// Mutual Password
func (r ApiApiIscsiDiscoveryauthPutRequest) MutualPassword(mutualPassword string) ApiApiIscsiDiscoveryauthPutRequest {
	r.mutualPassword = &mutualPassword
	return r
}

func (r ApiApiIscsiDiscoveryauthPutRequest) ApiIscsiDiscoveryauthPutRequest(apiIscsiDiscoveryauthPutRequest ApiIscsiDiscoveryauthPutRequest) ApiApiIscsiDiscoveryauthPutRequest {
	r.apiIscsiDiscoveryauthPutRequest = &apiIscsiDiscoveryauthPutRequest
	return r
}

func (r ApiApiIscsiDiscoveryauthPutRequest) Execute() (*http.Response, error) {
	return r.ApiService.ApiIscsiDiscoveryauthPutExecute(r)
}

/*
ApiIscsiDiscoveryauthPut Set Iscsi discoveryauth

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiApiIscsiDiscoveryauthPutRequest
*/
func (a *IscsiAPIService) ApiIscsiDiscoveryauthPut(ctx context.Context) ApiApiIscsiDiscoveryauthPutRequest {
	return ApiApiIscsiDiscoveryauthPutRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *IscsiAPIService) ApiIscsiDiscoveryauthPutExecute(r ApiApiIscsiDiscoveryauthPutRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IscsiAPIService.ApiIscsiDiscoveryauthPut")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/iscsi/discoveryauth"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.user == nil {
		return nil, reportError("user is required and must be specified")
	}
	if r.password == nil {
		return nil, reportError("password is required and must be specified")
	}
	if r.mutualUser == nil {
		return nil, reportError("mutualUser is required and must be specified")
	}
	if r.mutualPassword == nil {
		return nil, reportError("mutualPassword is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "user", r.user, "form", "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "password", r.password, "form", "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "mutual_user", r.mutualUser, "form", "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "mutual_password", r.mutualPassword, "form", "")
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
	localVarPostBody = r.apiIscsiDiscoveryauthPutRequest
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
