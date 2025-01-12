/*
Jamf Pro API

## Overview The Jamf Pro API is a RESTful API for Jamf Pro built to enable consistent and efficient programmatic access to Jamf Pro.<br/><br/> The swagger schema can be found [here](/api/schema/). 

API version: production
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
)


type JamfProInitializationPreviewApi interface {

	/*
	SystemInitializeDatabaseConnectionPost Provide Database Password during startup 

	Provide database password during startup. Endpoint is accessible when database password was not configured and Jamf Pro server has not been initialized yet.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiSystemInitializeDatabaseConnectionPostRequest

	Deprecated
	*/
	SystemInitializeDatabaseConnectionPost(ctx context.Context) ApiSystemInitializeDatabaseConnectionPostRequest

	// SystemInitializeDatabaseConnectionPostExecute executes the request
	// Deprecated
	SystemInitializeDatabaseConnectionPostExecute(r ApiSystemInitializeDatabaseConnectionPostRequest) (*http.Response, error)

	/*
	SystemInitializePost Set up fresh installed Jamf Pro Server 

	Set up fresh installed Jamf Pro Server


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiSystemInitializePostRequest

	Deprecated
	*/
	SystemInitializePost(ctx context.Context) ApiSystemInitializePostRequest

	// SystemInitializePostExecute executes the request
	// Deprecated
	SystemInitializePostExecute(r ApiSystemInitializePostRequest) (*http.Response, error)
}

// JamfProInitializationPreviewApiService JamfProInitializationPreviewApi service
type JamfProInitializationPreviewApiService service

type ApiSystemInitializeDatabaseConnectionPostRequest struct {
	ctx context.Context
	ApiService JamfProInitializationPreviewApi
	databasePassword *DatabasePassword
}

func (r ApiSystemInitializeDatabaseConnectionPostRequest) DatabasePassword(databasePassword DatabasePassword) ApiSystemInitializeDatabaseConnectionPostRequest {
	r.databasePassword = &databasePassword
	return r
}

func (r ApiSystemInitializeDatabaseConnectionPostRequest) Execute() (*http.Response, error) {
	return r.ApiService.SystemInitializeDatabaseConnectionPostExecute(r)
}

/*
SystemInitializeDatabaseConnectionPost Provide Database Password during startup 

Provide database password during startup. Endpoint is accessible when database password was not configured and Jamf Pro server has not been initialized yet.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSystemInitializeDatabaseConnectionPostRequest

Deprecated
*/
func (a *JamfProInitializationPreviewApiService) SystemInitializeDatabaseConnectionPost(ctx context.Context) ApiSystemInitializeDatabaseConnectionPostRequest {
	return ApiSystemInitializeDatabaseConnectionPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
// Deprecated
func (a *JamfProInitializationPreviewApiService) SystemInitializeDatabaseConnectionPostExecute(r ApiSystemInitializeDatabaseConnectionPostRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "JamfProInitializationPreviewApiService.SystemInitializeDatabaseConnectionPost")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/system/initialize-database-connection"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.databasePassword == nil {
		return nil, reportError("databasePassword is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.databasePassword
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ApiError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v ApiError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiSystemInitializePostRequest struct {
	ctx context.Context
	ApiService JamfProInitializationPreviewApi
	initialize *Initialize
}

func (r ApiSystemInitializePostRequest) Initialize(initialize Initialize) ApiSystemInitializePostRequest {
	r.initialize = &initialize
	return r
}

func (r ApiSystemInitializePostRequest) Execute() (*http.Response, error) {
	return r.ApiService.SystemInitializePostExecute(r)
}

/*
SystemInitializePost Set up fresh installed Jamf Pro Server 

Set up fresh installed Jamf Pro Server


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSystemInitializePostRequest

Deprecated
*/
func (a *JamfProInitializationPreviewApiService) SystemInitializePost(ctx context.Context) ApiSystemInitializePostRequest {
	return ApiSystemInitializePostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
// Deprecated
func (a *JamfProInitializationPreviewApiService) SystemInitializePostExecute(r ApiSystemInitializePostRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "JamfProInitializationPreviewApiService.SystemInitializePost")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/system/initialize"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.initialize == nil {
		return nil, reportError("initialize is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.initialize
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ApiError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}
