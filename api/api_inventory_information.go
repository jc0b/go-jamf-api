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


type InventoryInformationApi interface {

	/*
	V1InventoryInformationGet Get statistics about managed/unmanaged devices and computers in the inventory 

	Gets statistics about managed/unmanaged devices and computers in the inventory.


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiV1InventoryInformationGetRequest
	*/
	V1InventoryInformationGet(ctx context.Context) ApiV1InventoryInformationGetRequest

	// V1InventoryInformationGetExecute executes the request
	//  @return InventoryInformation
	V1InventoryInformationGetExecute(r ApiV1InventoryInformationGetRequest) (*InventoryInformation, *http.Response, error)
}

// InventoryInformationApiService InventoryInformationApi service
type InventoryInformationApiService service

type ApiV1InventoryInformationGetRequest struct {
	ctx context.Context
	ApiService InventoryInformationApi
}

func (r ApiV1InventoryInformationGetRequest) Execute() (*InventoryInformation, *http.Response, error) {
	return r.ApiService.V1InventoryInformationGetExecute(r)
}

/*
V1InventoryInformationGet Get statistics about managed/unmanaged devices and computers in the inventory 

Gets statistics about managed/unmanaged devices and computers in the inventory.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiV1InventoryInformationGetRequest
*/
func (a *InventoryInformationApiService) V1InventoryInformationGet(ctx context.Context) ApiV1InventoryInformationGetRequest {
	return ApiV1InventoryInformationGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return InventoryInformation
func (a *InventoryInformationApiService) V1InventoryInformationGetExecute(r ApiV1InventoryInformationGetRequest) (*InventoryInformation, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *InventoryInformation
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InventoryInformationApiService.V1InventoryInformationGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/inventory-information"

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
