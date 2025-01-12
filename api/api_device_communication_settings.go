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
	"reflect"
)


type DeviceCommunicationSettingsApi interface {

	/*
	V1DeviceCommunicationSettingsGet Retrieves all settings for device communication 

	Retrieves all device communication settings, including automatic renewal of the MDM profile.


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiV1DeviceCommunicationSettingsGetRequest
	*/
	V1DeviceCommunicationSettingsGet(ctx context.Context) ApiV1DeviceCommunicationSettingsGetRequest

	// V1DeviceCommunicationSettingsGetExecute executes the request
	//  @return DeviceCommunicationSettings
	V1DeviceCommunicationSettingsGetExecute(r ApiV1DeviceCommunicationSettingsGetRequest) (*DeviceCommunicationSettings, *http.Response, error)

	/*
	V1DeviceCommunicationSettingsHistoryGet Get Device Communication settings history 

	Gets Device Communication settings history


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiV1DeviceCommunicationSettingsHistoryGetRequest
	*/
	V1DeviceCommunicationSettingsHistoryGet(ctx context.Context) ApiV1DeviceCommunicationSettingsHistoryGetRequest

	// V1DeviceCommunicationSettingsHistoryGetExecute executes the request
	//  @return HistorySearchResults
	V1DeviceCommunicationSettingsHistoryGetExecute(r ApiV1DeviceCommunicationSettingsHistoryGetRequest) (*HistorySearchResults, *http.Response, error)

	/*
	V1DeviceCommunicationSettingsHistoryPost Add Device Communication Settings history notes 

	Adds Device Communication Settings history notes


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiV1DeviceCommunicationSettingsHistoryPostRequest
	*/
	V1DeviceCommunicationSettingsHistoryPost(ctx context.Context) ApiV1DeviceCommunicationSettingsHistoryPostRequest

	// V1DeviceCommunicationSettingsHistoryPostExecute executes the request
	//  @return ObjectHistory
	V1DeviceCommunicationSettingsHistoryPostExecute(r ApiV1DeviceCommunicationSettingsHistoryPostRequest) (*ObjectHistory, *http.Response, error)

	/*
	V1DeviceCommunicationSettingsPut Update device communication settings 

	Update device communication settings


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiV1DeviceCommunicationSettingsPutRequest
	*/
	V1DeviceCommunicationSettingsPut(ctx context.Context) ApiV1DeviceCommunicationSettingsPutRequest

	// V1DeviceCommunicationSettingsPutExecute executes the request
	//  @return DeviceCommunicationSettings
	V1DeviceCommunicationSettingsPutExecute(r ApiV1DeviceCommunicationSettingsPutRequest) (*DeviceCommunicationSettings, *http.Response, error)
}

// DeviceCommunicationSettingsApiService DeviceCommunicationSettingsApi service
type DeviceCommunicationSettingsApiService service

type ApiV1DeviceCommunicationSettingsGetRequest struct {
	ctx context.Context
	ApiService DeviceCommunicationSettingsApi
}

func (r ApiV1DeviceCommunicationSettingsGetRequest) Execute() (*DeviceCommunicationSettings, *http.Response, error) {
	return r.ApiService.V1DeviceCommunicationSettingsGetExecute(r)
}

/*
V1DeviceCommunicationSettingsGet Retrieves all settings for device communication 

Retrieves all device communication settings, including automatic renewal of the MDM profile.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiV1DeviceCommunicationSettingsGetRequest
*/
func (a *DeviceCommunicationSettingsApiService) V1DeviceCommunicationSettingsGet(ctx context.Context) ApiV1DeviceCommunicationSettingsGetRequest {
	return ApiV1DeviceCommunicationSettingsGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return DeviceCommunicationSettings
func (a *DeviceCommunicationSettingsApiService) V1DeviceCommunicationSettingsGetExecute(r ApiV1DeviceCommunicationSettingsGetRequest) (*DeviceCommunicationSettings, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DeviceCommunicationSettings
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DeviceCommunicationSettingsApiService.V1DeviceCommunicationSettingsGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/device-communication-settings"

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

type ApiV1DeviceCommunicationSettingsHistoryGetRequest struct {
	ctx context.Context
	ApiService DeviceCommunicationSettingsApi
	page *int32
	pageSize *int32
	sort *[]string
	filter *string
}

func (r ApiV1DeviceCommunicationSettingsHistoryGetRequest) Page(page int32) ApiV1DeviceCommunicationSettingsHistoryGetRequest {
	r.page = &page
	return r
}

func (r ApiV1DeviceCommunicationSettingsHistoryGetRequest) PageSize(pageSize int32) ApiV1DeviceCommunicationSettingsHistoryGetRequest {
	r.pageSize = &pageSize
	return r
}

// Sorting criteria in the format: property:asc/desc. Default sort is date:desc. Multiple sort criteria are supported and must be separated with a comma. Example: sort&#x3D;date:desc,name:asc 
func (r ApiV1DeviceCommunicationSettingsHistoryGetRequest) Sort(sort []string) ApiV1DeviceCommunicationSettingsHistoryGetRequest {
	r.sort = &sort
	return r
}

// Query in the RSQL format, allowing to filter history notes collection. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: username, date, note, details. This param can be combined with paging and sorting. Example: filter&#x3D;username!&#x3D;admin and details&#x3D;&#x3D;*disabled* and date&lt;2019-12-15
func (r ApiV1DeviceCommunicationSettingsHistoryGetRequest) Filter(filter string) ApiV1DeviceCommunicationSettingsHistoryGetRequest {
	r.filter = &filter
	return r
}

func (r ApiV1DeviceCommunicationSettingsHistoryGetRequest) Execute() (*HistorySearchResults, *http.Response, error) {
	return r.ApiService.V1DeviceCommunicationSettingsHistoryGetExecute(r)
}

/*
V1DeviceCommunicationSettingsHistoryGet Get Device Communication settings history 

Gets Device Communication settings history


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiV1DeviceCommunicationSettingsHistoryGetRequest
*/
func (a *DeviceCommunicationSettingsApiService) V1DeviceCommunicationSettingsHistoryGet(ctx context.Context) ApiV1DeviceCommunicationSettingsHistoryGetRequest {
	return ApiV1DeviceCommunicationSettingsHistoryGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return HistorySearchResults
func (a *DeviceCommunicationSettingsApiService) V1DeviceCommunicationSettingsHistoryGetExecute(r ApiV1DeviceCommunicationSettingsHistoryGetRequest) (*HistorySearchResults, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *HistorySearchResults
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DeviceCommunicationSettingsApiService.V1DeviceCommunicationSettingsHistoryGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/device-communication-settings/history"

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
	if r.filter != nil {
		localVarQueryParams.Add("filter", parameterToString(*r.filter, ""))
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

type ApiV1DeviceCommunicationSettingsHistoryPostRequest struct {
	ctx context.Context
	ApiService DeviceCommunicationSettingsApi
	objectHistoryNote *ObjectHistoryNote
}

// history notes to create
func (r ApiV1DeviceCommunicationSettingsHistoryPostRequest) ObjectHistoryNote(objectHistoryNote ObjectHistoryNote) ApiV1DeviceCommunicationSettingsHistoryPostRequest {
	r.objectHistoryNote = &objectHistoryNote
	return r
}

func (r ApiV1DeviceCommunicationSettingsHistoryPostRequest) Execute() (*ObjectHistory, *http.Response, error) {
	return r.ApiService.V1DeviceCommunicationSettingsHistoryPostExecute(r)
}

/*
V1DeviceCommunicationSettingsHistoryPost Add Device Communication Settings history notes 

Adds Device Communication Settings history notes


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiV1DeviceCommunicationSettingsHistoryPostRequest
*/
func (a *DeviceCommunicationSettingsApiService) V1DeviceCommunicationSettingsHistoryPost(ctx context.Context) ApiV1DeviceCommunicationSettingsHistoryPostRequest {
	return ApiV1DeviceCommunicationSettingsHistoryPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ObjectHistory
func (a *DeviceCommunicationSettingsApiService) V1DeviceCommunicationSettingsHistoryPostExecute(r ApiV1DeviceCommunicationSettingsHistoryPostRequest) (*ObjectHistory, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ObjectHistory
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DeviceCommunicationSettingsApiService.V1DeviceCommunicationSettingsHistoryPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/device-communication-settings/history"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.objectHistoryNote == nil {
		return localVarReturnValue, nil, reportError("objectHistoryNote is required and must be specified")
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
	localVarPostBody = r.objectHistoryNote
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
		if localVarHTTPResponse.StatusCode == 503 {
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

type ApiV1DeviceCommunicationSettingsPutRequest struct {
	ctx context.Context
	ApiService DeviceCommunicationSettingsApi
	deviceCommunicationSettings *DeviceCommunicationSettings
}

func (r ApiV1DeviceCommunicationSettingsPutRequest) DeviceCommunicationSettings(deviceCommunicationSettings DeviceCommunicationSettings) ApiV1DeviceCommunicationSettingsPutRequest {
	r.deviceCommunicationSettings = &deviceCommunicationSettings
	return r
}

func (r ApiV1DeviceCommunicationSettingsPutRequest) Execute() (*DeviceCommunicationSettings, *http.Response, error) {
	return r.ApiService.V1DeviceCommunicationSettingsPutExecute(r)
}

/*
V1DeviceCommunicationSettingsPut Update device communication settings 

Update device communication settings


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiV1DeviceCommunicationSettingsPutRequest
*/
func (a *DeviceCommunicationSettingsApiService) V1DeviceCommunicationSettingsPut(ctx context.Context) ApiV1DeviceCommunicationSettingsPutRequest {
	return ApiV1DeviceCommunicationSettingsPutRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return DeviceCommunicationSettings
func (a *DeviceCommunicationSettingsApiService) V1DeviceCommunicationSettingsPutExecute(r ApiV1DeviceCommunicationSettingsPutRequest) (*DeviceCommunicationSettings, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DeviceCommunicationSettings
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DeviceCommunicationSettingsApiService.V1DeviceCommunicationSettingsPut")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/device-communication-settings"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.deviceCommunicationSettings == nil {
		return localVarReturnValue, nil, reportError("deviceCommunicationSettings is required and must be specified")
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
	localVarPostBody = r.deviceCommunicationSettings
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
