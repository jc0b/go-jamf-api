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


type ReEnrollmentPreviewApi interface {

	/*
	V1ReenrollmentGet Get Re-enrollment object 

	Gets Re-enrollment object


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiV1ReenrollmentGetRequest
	*/
	V1ReenrollmentGet(ctx context.Context) ApiV1ReenrollmentGetRequest

	// V1ReenrollmentGetExecute executes the request
	//  @return Reenrollment
	V1ReenrollmentGetExecute(r ApiV1ReenrollmentGetRequest) (*Reenrollment, *http.Response, error)

	/*
	V1ReenrollmentHistoryGet Get Re-enrollment history object 

	Gets Re-enrollment history object


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiV1ReenrollmentHistoryGetRequest
	*/
	V1ReenrollmentHistoryGet(ctx context.Context) ApiV1ReenrollmentHistoryGetRequest

	// V1ReenrollmentHistoryGetExecute executes the request
	//  @return HistorySearchResults
	V1ReenrollmentHistoryGetExecute(r ApiV1ReenrollmentHistoryGetRequest) (*HistorySearchResults, *http.Response, error)

	/*
	V1ReenrollmentHistoryPost Add specified Re-enrollment history object notes 

	Adds specified Re-enrollment history object notes


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiV1ReenrollmentHistoryPostRequest
	*/
	V1ReenrollmentHistoryPost(ctx context.Context) ApiV1ReenrollmentHistoryPostRequest

	// V1ReenrollmentHistoryPostExecute executes the request
	//  @return ObjectHistory
	V1ReenrollmentHistoryPostExecute(r ApiV1ReenrollmentHistoryPostRequest) (*ObjectHistory, *http.Response, error)

	/*
	V1ReenrollmentPut Update the Re-enrollment object 

	Update the Re-enrollment object


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiV1ReenrollmentPutRequest
	*/
	V1ReenrollmentPut(ctx context.Context) ApiV1ReenrollmentPutRequest

	// V1ReenrollmentPutExecute executes the request
	//  @return Reenrollment
	V1ReenrollmentPutExecute(r ApiV1ReenrollmentPutRequest) (*Reenrollment, *http.Response, error)
}

// ReEnrollmentPreviewApiService ReEnrollmentPreviewApi service
type ReEnrollmentPreviewApiService service

type ApiV1ReenrollmentGetRequest struct {
	ctx context.Context
	ApiService ReEnrollmentPreviewApi
}

func (r ApiV1ReenrollmentGetRequest) Execute() (*Reenrollment, *http.Response, error) {
	return r.ApiService.V1ReenrollmentGetExecute(r)
}

/*
V1ReenrollmentGet Get Re-enrollment object 

Gets Re-enrollment object


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiV1ReenrollmentGetRequest
*/
func (a *ReEnrollmentPreviewApiService) V1ReenrollmentGet(ctx context.Context) ApiV1ReenrollmentGetRequest {
	return ApiV1ReenrollmentGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return Reenrollment
func (a *ReEnrollmentPreviewApiService) V1ReenrollmentGetExecute(r ApiV1ReenrollmentGetRequest) (*Reenrollment, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Reenrollment
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReEnrollmentPreviewApiService.V1ReenrollmentGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/reenrollment"

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

type ApiV1ReenrollmentHistoryGetRequest struct {
	ctx context.Context
	ApiService ReEnrollmentPreviewApi
	page *int32
	size *int32
	pagesize *int32
	pageSize *int32
	sort *string
}

func (r ApiV1ReenrollmentHistoryGetRequest) Page(page int32) ApiV1ReenrollmentHistoryGetRequest {
	r.page = &page
	return r
}

// Deprecated
func (r ApiV1ReenrollmentHistoryGetRequest) Size(size int32) ApiV1ReenrollmentHistoryGetRequest {
	r.size = &size
	return r
}

// Deprecated
func (r ApiV1ReenrollmentHistoryGetRequest) Pagesize(pagesize int32) ApiV1ReenrollmentHistoryGetRequest {
	r.pagesize = &pagesize
	return r
}

func (r ApiV1ReenrollmentHistoryGetRequest) PageSize(pageSize int32) ApiV1ReenrollmentHistoryGetRequest {
	r.pageSize = &pageSize
	return r
}

// Sorting criteria in the format: property:asc/desc. Default sort is date:desc. Multiple sort criteria are supported and must be separated with a comma. Example: sort&#x3D;date:desc,name:asc 
func (r ApiV1ReenrollmentHistoryGetRequest) Sort(sort string) ApiV1ReenrollmentHistoryGetRequest {
	r.sort = &sort
	return r
}

func (r ApiV1ReenrollmentHistoryGetRequest) Execute() (*HistorySearchResults, *http.Response, error) {
	return r.ApiService.V1ReenrollmentHistoryGetExecute(r)
}

/*
V1ReenrollmentHistoryGet Get Re-enrollment history object 

Gets Re-enrollment history object


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiV1ReenrollmentHistoryGetRequest
*/
func (a *ReEnrollmentPreviewApiService) V1ReenrollmentHistoryGet(ctx context.Context) ApiV1ReenrollmentHistoryGetRequest {
	return ApiV1ReenrollmentHistoryGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return HistorySearchResults
func (a *ReEnrollmentPreviewApiService) V1ReenrollmentHistoryGetExecute(r ApiV1ReenrollmentHistoryGetRequest) (*HistorySearchResults, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *HistorySearchResults
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReEnrollmentPreviewApiService.V1ReenrollmentHistoryGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/reenrollment/history"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.page != nil {
		localVarQueryParams.Add("page", parameterToString(*r.page, ""))
	}
	if r.size != nil {
		localVarQueryParams.Add("size", parameterToString(*r.size, ""))
	}
	if r.pagesize != nil {
		localVarQueryParams.Add("pagesize", parameterToString(*r.pagesize, ""))
	}
	if r.pageSize != nil {
		localVarQueryParams.Add("page-size", parameterToString(*r.pageSize, ""))
	}
	if r.sort != nil {
		localVarQueryParams.Add("sort", parameterToString(*r.sort, ""))
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

type ApiV1ReenrollmentHistoryPostRequest struct {
	ctx context.Context
	ApiService ReEnrollmentPreviewApi
	objectHistoryNote *ObjectHistoryNote
}

// history notes to create
func (r ApiV1ReenrollmentHistoryPostRequest) ObjectHistoryNote(objectHistoryNote ObjectHistoryNote) ApiV1ReenrollmentHistoryPostRequest {
	r.objectHistoryNote = &objectHistoryNote
	return r
}

func (r ApiV1ReenrollmentHistoryPostRequest) Execute() (*ObjectHistory, *http.Response, error) {
	return r.ApiService.V1ReenrollmentHistoryPostExecute(r)
}

/*
V1ReenrollmentHistoryPost Add specified Re-enrollment history object notes 

Adds specified Re-enrollment history object notes


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiV1ReenrollmentHistoryPostRequest
*/
func (a *ReEnrollmentPreviewApiService) V1ReenrollmentHistoryPost(ctx context.Context) ApiV1ReenrollmentHistoryPostRequest {
	return ApiV1ReenrollmentHistoryPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ObjectHistory
func (a *ReEnrollmentPreviewApiService) V1ReenrollmentHistoryPostExecute(r ApiV1ReenrollmentHistoryPostRequest) (*ObjectHistory, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ObjectHistory
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReEnrollmentPreviewApiService.V1ReenrollmentHistoryPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/reenrollment/history"

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

type ApiV1ReenrollmentPutRequest struct {
	ctx context.Context
	ApiService ReEnrollmentPreviewApi
	reenrollment *Reenrollment
}

// Re-enrollment object to update
func (r ApiV1ReenrollmentPutRequest) Reenrollment(reenrollment Reenrollment) ApiV1ReenrollmentPutRequest {
	r.reenrollment = &reenrollment
	return r
}

func (r ApiV1ReenrollmentPutRequest) Execute() (*Reenrollment, *http.Response, error) {
	return r.ApiService.V1ReenrollmentPutExecute(r)
}

/*
V1ReenrollmentPut Update the Re-enrollment object 

Update the Re-enrollment object


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiV1ReenrollmentPutRequest
*/
func (a *ReEnrollmentPreviewApiService) V1ReenrollmentPut(ctx context.Context) ApiV1ReenrollmentPutRequest {
	return ApiV1ReenrollmentPutRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return Reenrollment
func (a *ReEnrollmentPreviewApiService) V1ReenrollmentPutExecute(r ApiV1ReenrollmentPutRequest) (*Reenrollment, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Reenrollment
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReEnrollmentPreviewApiService.V1ReenrollmentPut")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/reenrollment"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.reenrollment == nil {
		return localVarReturnValue, nil, reportError("reenrollment is required and must be specified")
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
	localVarPostBody = r.reenrollment
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
