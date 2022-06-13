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
	"strings"
	"reflect"
)


type SelfServiceBrandingMacosApi interface {

	/*
	V1SelfServiceBrandingMacosGet Search for sorted and paged macOS branding configurations 

	Search for sorted and paged macOS branding configurations

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiV1SelfServiceBrandingMacosGetRequest
	*/
	V1SelfServiceBrandingMacosGet(ctx context.Context) ApiV1SelfServiceBrandingMacosGetRequest

	// V1SelfServiceBrandingMacosGetExecute executes the request
	//  @return MacOsBrandingSearchResults
	V1SelfServiceBrandingMacosGetExecute(r ApiV1SelfServiceBrandingMacosGetRequest) (*MacOsBrandingSearchResults, *http.Response, error)

	/*
	V1SelfServiceBrandingMacosIdDelete Delete the Self Service macOS branding configuration indicated by the provided id 

	Delete the Self Service macOS branding configuration indicated by the provided id.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id id of macOS branding configuration
	@return ApiV1SelfServiceBrandingMacosIdDeleteRequest
	*/
	V1SelfServiceBrandingMacosIdDelete(ctx context.Context, id string) ApiV1SelfServiceBrandingMacosIdDeleteRequest

	// V1SelfServiceBrandingMacosIdDeleteExecute executes the request
	V1SelfServiceBrandingMacosIdDeleteExecute(r ApiV1SelfServiceBrandingMacosIdDeleteRequest) (*http.Response, error)

	/*
	V1SelfServiceBrandingMacosIdGet Read a single Self Service macOS branding configuration indicated by the provided id 

	Read a single Self Service macOS branding configuration indicated by the provided id.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id id of macOS branding configuration
	@return ApiV1SelfServiceBrandingMacosIdGetRequest
	*/
	V1SelfServiceBrandingMacosIdGet(ctx context.Context, id string) ApiV1SelfServiceBrandingMacosIdGetRequest

	// V1SelfServiceBrandingMacosIdGetExecute executes the request
	//  @return MacOsBrandingConfiguration
	V1SelfServiceBrandingMacosIdGetExecute(r ApiV1SelfServiceBrandingMacosIdGetRequest) (*MacOsBrandingConfiguration, *http.Response, error)

	/*
	V1SelfServiceBrandingMacosIdPut Update a Self Service macOS branding configuration with the supplied details 

	Update a Self Service macOS branding configuration with the supplied details

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id id of macOS branding configuration
	@return ApiV1SelfServiceBrandingMacosIdPutRequest
	*/
	V1SelfServiceBrandingMacosIdPut(ctx context.Context, id string) ApiV1SelfServiceBrandingMacosIdPutRequest

	// V1SelfServiceBrandingMacosIdPutExecute executes the request
	//  @return MacOsBrandingConfiguration
	V1SelfServiceBrandingMacosIdPutExecute(r ApiV1SelfServiceBrandingMacosIdPutRequest) (*MacOsBrandingConfiguration, *http.Response, error)

	/*
	V1SelfServiceBrandingMacosPost Create a Self Service macOS branding configuration with the supplied 

	Create a Self Service macOS branding configuration with the supplied details

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiV1SelfServiceBrandingMacosPostRequest
	*/
	V1SelfServiceBrandingMacosPost(ctx context.Context) ApiV1SelfServiceBrandingMacosPostRequest

	// V1SelfServiceBrandingMacosPostExecute executes the request
	//  @return HrefResponse
	V1SelfServiceBrandingMacosPostExecute(r ApiV1SelfServiceBrandingMacosPostRequest) (*HrefResponse, *http.Response, error)
}

// SelfServiceBrandingMacosApiService SelfServiceBrandingMacosApi service
type SelfServiceBrandingMacosApiService service

type ApiV1SelfServiceBrandingMacosGetRequest struct {
	ctx context.Context
	ApiService SelfServiceBrandingMacosApi
	page *int32
	pageSize *int32
	sort *[]string
}

func (r ApiV1SelfServiceBrandingMacosGetRequest) Page(page int32) ApiV1SelfServiceBrandingMacosGetRequest {
	r.page = &page
	return r
}

func (r ApiV1SelfServiceBrandingMacosGetRequest) PageSize(pageSize int32) ApiV1SelfServiceBrandingMacosGetRequest {
	r.pageSize = &pageSize
	return r
}

// Sorting criteria in the format: property:asc/desc. Default sort is id:asc. Multiple sort criteria are supported and must be separated with a comma. Example: sort&#x3D;id:desc,brandingName:asc 
func (r ApiV1SelfServiceBrandingMacosGetRequest) Sort(sort []string) ApiV1SelfServiceBrandingMacosGetRequest {
	r.sort = &sort
	return r
}

func (r ApiV1SelfServiceBrandingMacosGetRequest) Execute() (*MacOsBrandingSearchResults, *http.Response, error) {
	return r.ApiService.V1SelfServiceBrandingMacosGetExecute(r)
}

/*
V1SelfServiceBrandingMacosGet Search for sorted and paged macOS branding configurations 

Search for sorted and paged macOS branding configurations

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiV1SelfServiceBrandingMacosGetRequest
*/
func (a *SelfServiceBrandingMacosApiService) V1SelfServiceBrandingMacosGet(ctx context.Context) ApiV1SelfServiceBrandingMacosGetRequest {
	return ApiV1SelfServiceBrandingMacosGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return MacOsBrandingSearchResults
func (a *SelfServiceBrandingMacosApiService) V1SelfServiceBrandingMacosGetExecute(r ApiV1SelfServiceBrandingMacosGetRequest) (*MacOsBrandingSearchResults, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MacOsBrandingSearchResults
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SelfServiceBrandingMacosApiService.V1SelfServiceBrandingMacosGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/self-service/branding/macos"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.page != nil {
		localVarQueryParams.Add("page", parameterToString(*r.page, ""))
	}
	if r.pageSize != nil {
		localVarQueryParams.Add("page-size", parameterToString(*r.pageSize, ""))
	}
	if r.sort != nil {
		t := *r.sort
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("sort", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("sort", parameterToString(t, "multi"))
		}
	}
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

type ApiV1SelfServiceBrandingMacosIdDeleteRequest struct {
	ctx context.Context
	ApiService SelfServiceBrandingMacosApi
	id string
}

func (r ApiV1SelfServiceBrandingMacosIdDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.V1SelfServiceBrandingMacosIdDeleteExecute(r)
}

/*
V1SelfServiceBrandingMacosIdDelete Delete the Self Service macOS branding configuration indicated by the provided id 

Delete the Self Service macOS branding configuration indicated by the provided id.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id id of macOS branding configuration
 @return ApiV1SelfServiceBrandingMacosIdDeleteRequest
*/
func (a *SelfServiceBrandingMacosApiService) V1SelfServiceBrandingMacosIdDelete(ctx context.Context, id string) ApiV1SelfServiceBrandingMacosIdDeleteRequest {
	return ApiV1SelfServiceBrandingMacosIdDeleteRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *SelfServiceBrandingMacosApiService) V1SelfServiceBrandingMacosIdDeleteExecute(r ApiV1SelfServiceBrandingMacosIdDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SelfServiceBrandingMacosApiService.V1SelfServiceBrandingMacosIdDelete")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/self-service/branding/macos/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterToString(r.id, "")), -1)

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
		if localVarHTTPResponse.StatusCode == 404 {
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

type ApiV1SelfServiceBrandingMacosIdGetRequest struct {
	ctx context.Context
	ApiService SelfServiceBrandingMacosApi
	id string
}

func (r ApiV1SelfServiceBrandingMacosIdGetRequest) Execute() (*MacOsBrandingConfiguration, *http.Response, error) {
	return r.ApiService.V1SelfServiceBrandingMacosIdGetExecute(r)
}

/*
V1SelfServiceBrandingMacosIdGet Read a single Self Service macOS branding configuration indicated by the provided id 

Read a single Self Service macOS branding configuration indicated by the provided id.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id id of macOS branding configuration
 @return ApiV1SelfServiceBrandingMacosIdGetRequest
*/
func (a *SelfServiceBrandingMacosApiService) V1SelfServiceBrandingMacosIdGet(ctx context.Context, id string) ApiV1SelfServiceBrandingMacosIdGetRequest {
	return ApiV1SelfServiceBrandingMacosIdGetRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return MacOsBrandingConfiguration
func (a *SelfServiceBrandingMacosApiService) V1SelfServiceBrandingMacosIdGetExecute(r ApiV1SelfServiceBrandingMacosIdGetRequest) (*MacOsBrandingConfiguration, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MacOsBrandingConfiguration
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SelfServiceBrandingMacosApiService.V1SelfServiceBrandingMacosIdGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/self-service/branding/macos/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterToString(r.id, "")), -1)

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

type ApiV1SelfServiceBrandingMacosIdPutRequest struct {
	ctx context.Context
	ApiService SelfServiceBrandingMacosApi
	id string
	macOsBrandingConfiguration *MacOsBrandingConfiguration
}

// The macOS branding configuration values to update
func (r ApiV1SelfServiceBrandingMacosIdPutRequest) MacOsBrandingConfiguration(macOsBrandingConfiguration MacOsBrandingConfiguration) ApiV1SelfServiceBrandingMacosIdPutRequest {
	r.macOsBrandingConfiguration = &macOsBrandingConfiguration
	return r
}

func (r ApiV1SelfServiceBrandingMacosIdPutRequest) Execute() (*MacOsBrandingConfiguration, *http.Response, error) {
	return r.ApiService.V1SelfServiceBrandingMacosIdPutExecute(r)
}

/*
V1SelfServiceBrandingMacosIdPut Update a Self Service macOS branding configuration with the supplied details 

Update a Self Service macOS branding configuration with the supplied details

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id id of macOS branding configuration
 @return ApiV1SelfServiceBrandingMacosIdPutRequest
*/
func (a *SelfServiceBrandingMacosApiService) V1SelfServiceBrandingMacosIdPut(ctx context.Context, id string) ApiV1SelfServiceBrandingMacosIdPutRequest {
	return ApiV1SelfServiceBrandingMacosIdPutRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return MacOsBrandingConfiguration
func (a *SelfServiceBrandingMacosApiService) V1SelfServiceBrandingMacosIdPutExecute(r ApiV1SelfServiceBrandingMacosIdPutRequest) (*MacOsBrandingConfiguration, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MacOsBrandingConfiguration
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SelfServiceBrandingMacosApiService.V1SelfServiceBrandingMacosIdPut")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/self-service/branding/macos/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterToString(r.id, "")), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.macOsBrandingConfiguration
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
		if localVarHTTPResponse.StatusCode == 404 {
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

type ApiV1SelfServiceBrandingMacosPostRequest struct {
	ctx context.Context
	ApiService SelfServiceBrandingMacosApi
	macOsBrandingConfiguration *MacOsBrandingConfiguration
}

// The macOS branding configuration to create
func (r ApiV1SelfServiceBrandingMacosPostRequest) MacOsBrandingConfiguration(macOsBrandingConfiguration MacOsBrandingConfiguration) ApiV1SelfServiceBrandingMacosPostRequest {
	r.macOsBrandingConfiguration = &macOsBrandingConfiguration
	return r
}

func (r ApiV1SelfServiceBrandingMacosPostRequest) Execute() (*HrefResponse, *http.Response, error) {
	return r.ApiService.V1SelfServiceBrandingMacosPostExecute(r)
}

/*
V1SelfServiceBrandingMacosPost Create a Self Service macOS branding configuration with the supplied 

Create a Self Service macOS branding configuration with the supplied details

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiV1SelfServiceBrandingMacosPostRequest
*/
func (a *SelfServiceBrandingMacosApiService) V1SelfServiceBrandingMacosPost(ctx context.Context) ApiV1SelfServiceBrandingMacosPostRequest {
	return ApiV1SelfServiceBrandingMacosPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return HrefResponse
func (a *SelfServiceBrandingMacosApiService) V1SelfServiceBrandingMacosPostExecute(r ApiV1SelfServiceBrandingMacosPostRequest) (*HrefResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *HrefResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SelfServiceBrandingMacosApiService.V1SelfServiceBrandingMacosPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/self-service/branding/macos"

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
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.macOsBrandingConfiguration
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