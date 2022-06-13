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


type AppDynamicsConfigurationPreviewApi interface {

	/*
	V1AppDynamicsScriptConfigurationGet Get Application Dynamics Config object 

	Gets `AppDynamicsConfig` object.
Details for using the response as script configuration are available in the official documentation - https://docs.appdynamics.com/display/PRO45/Configure+the+JavaScript+Agent


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiV1AppDynamicsScriptConfigurationGetRequest
	*/
	V1AppDynamicsScriptConfigurationGet(ctx context.Context) ApiV1AppDynamicsScriptConfigurationGetRequest

	// V1AppDynamicsScriptConfigurationGetExecute executes the request
	//  @return AppDynamicsConfig
	V1AppDynamicsScriptConfigurationGetExecute(r ApiV1AppDynamicsScriptConfigurationGetRequest) (*AppDynamicsConfig, *http.Response, error)
}

// AppDynamicsConfigurationPreviewApiService AppDynamicsConfigurationPreviewApi service
type AppDynamicsConfigurationPreviewApiService service

type ApiV1AppDynamicsScriptConfigurationGetRequest struct {
	ctx context.Context
	ApiService AppDynamicsConfigurationPreviewApi
}

func (r ApiV1AppDynamicsScriptConfigurationGetRequest) Execute() (*AppDynamicsConfig, *http.Response, error) {
	return r.ApiService.V1AppDynamicsScriptConfigurationGetExecute(r)
}

/*
V1AppDynamicsScriptConfigurationGet Get Application Dynamics Config object 

Gets `AppDynamicsConfig` object.
Details for using the response as script configuration are available in the official documentation - https://docs.appdynamics.com/display/PRO45/Configure+the+JavaScript+Agent


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiV1AppDynamicsScriptConfigurationGetRequest
*/
func (a *AppDynamicsConfigurationPreviewApiService) V1AppDynamicsScriptConfigurationGet(ctx context.Context) ApiV1AppDynamicsScriptConfigurationGetRequest {
	return ApiV1AppDynamicsScriptConfigurationGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AppDynamicsConfig
func (a *AppDynamicsConfigurationPreviewApiService) V1AppDynamicsScriptConfigurationGetExecute(r ApiV1AppDynamicsScriptConfigurationGetRequest) (*AppDynamicsConfig, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AppDynamicsConfig
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AppDynamicsConfigurationPreviewApiService.V1AppDynamicsScriptConfigurationGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/app-dynamics/script-configuration"

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
	localVarHTTPHeaderAccepts := []string{"application/json"}

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

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
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
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
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