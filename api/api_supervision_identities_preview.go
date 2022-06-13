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
	"os"
)


type SupervisionIdentitiesPreviewApi interface {

	/*
	V1SupervisionIdentitiesGet Search for sorted and paged Supervision Identities 

	Search for sorted and paged supervision identities

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiV1SupervisionIdentitiesGetRequest
	*/
	V1SupervisionIdentitiesGet(ctx context.Context) ApiV1SupervisionIdentitiesGetRequest

	// V1SupervisionIdentitiesGetExecute executes the request
	//  @return SupervisionIdentitySearchResults
	V1SupervisionIdentitiesGetExecute(r ApiV1SupervisionIdentitiesGetRequest) (*SupervisionIdentitySearchResults, *http.Response, error)

	/*
	V1SupervisionIdentitiesIdDelete Delete a Supervision Identity with the supplied id 

	Deletes a Supervision Identity with the supplied id

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id Supervision Identity identifier
	@return ApiV1SupervisionIdentitiesIdDeleteRequest
	*/
	V1SupervisionIdentitiesIdDelete(ctx context.Context, id int32) ApiV1SupervisionIdentitiesIdDeleteRequest

	// V1SupervisionIdentitiesIdDeleteExecute executes the request
	V1SupervisionIdentitiesIdDeleteExecute(r ApiV1SupervisionIdentitiesIdDeleteRequest) (*http.Response, error)

	/*
	V1SupervisionIdentitiesIdDownloadGet Download the Supervision Identity .p12 file 

	Download the Supervision Identity .p12 file

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id Supervision Identity identifier
	@return ApiV1SupervisionIdentitiesIdDownloadGetRequest
	*/
	V1SupervisionIdentitiesIdDownloadGet(ctx context.Context, id int32) ApiV1SupervisionIdentitiesIdDownloadGetRequest

	// V1SupervisionIdentitiesIdDownloadGetExecute executes the request
	//  @return *os.File
	V1SupervisionIdentitiesIdDownloadGetExecute(r ApiV1SupervisionIdentitiesIdDownloadGetRequest) (**os.File, *http.Response, error)

	/*
	V1SupervisionIdentitiesIdGet Retrieve a Supervision Identity with the supplied id 

	Retrieves a Supervision Identity with the supplied id

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id Supervision Identity identifier
	@return ApiV1SupervisionIdentitiesIdGetRequest
	*/
	V1SupervisionIdentitiesIdGet(ctx context.Context, id int32) ApiV1SupervisionIdentitiesIdGetRequest

	// V1SupervisionIdentitiesIdGetExecute executes the request
	//  @return SupervisionIdentity
	V1SupervisionIdentitiesIdGetExecute(r ApiV1SupervisionIdentitiesIdGetRequest) (*SupervisionIdentity, *http.Response, error)

	/*
	V1SupervisionIdentitiesIdPut Update a Supervision Identity with the supplied information 

	Updates a Supervision Identity with the supplied information

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id Supervision Identity identifier
	@return ApiV1SupervisionIdentitiesIdPutRequest
	*/
	V1SupervisionIdentitiesIdPut(ctx context.Context, id int32) ApiV1SupervisionIdentitiesIdPutRequest

	// V1SupervisionIdentitiesIdPutExecute executes the request
	//  @return SupervisionIdentity
	V1SupervisionIdentitiesIdPutExecute(r ApiV1SupervisionIdentitiesIdPutRequest) (*SupervisionIdentity, *http.Response, error)

	/*
	V1SupervisionIdentitiesPost Create a Supervision Identity for the supplied information 

	Creates a Supervision Identity for the supplied information

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiV1SupervisionIdentitiesPostRequest
	*/
	V1SupervisionIdentitiesPost(ctx context.Context) ApiV1SupervisionIdentitiesPostRequest

	// V1SupervisionIdentitiesPostExecute executes the request
	//  @return SupervisionIdentity
	V1SupervisionIdentitiesPostExecute(r ApiV1SupervisionIdentitiesPostRequest) (*SupervisionIdentity, *http.Response, error)

	/*
	V1SupervisionIdentitiesUploadPost Upload the Supervision Identity .p12 file 

	Uploads the Supervision Identity .p12 file

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiV1SupervisionIdentitiesUploadPostRequest
	*/
	V1SupervisionIdentitiesUploadPost(ctx context.Context) ApiV1SupervisionIdentitiesUploadPostRequest

	// V1SupervisionIdentitiesUploadPostExecute executes the request
	//  @return SupervisionIdentity
	V1SupervisionIdentitiesUploadPostExecute(r ApiV1SupervisionIdentitiesUploadPostRequest) (*SupervisionIdentity, *http.Response, error)
}

// SupervisionIdentitiesPreviewApiService SupervisionIdentitiesPreviewApi service
type SupervisionIdentitiesPreviewApiService service

type ApiV1SupervisionIdentitiesGetRequest struct {
	ctx context.Context
	ApiService SupervisionIdentitiesPreviewApi
	page *int32
	size *int32
	pagesize *int32
	pageSize *int32
	sort *string
}

func (r ApiV1SupervisionIdentitiesGetRequest) Page(page int32) ApiV1SupervisionIdentitiesGetRequest {
	r.page = &page
	return r
}

// Deprecated
func (r ApiV1SupervisionIdentitiesGetRequest) Size(size int32) ApiV1SupervisionIdentitiesGetRequest {
	r.size = &size
	return r
}

// Deprecated
func (r ApiV1SupervisionIdentitiesGetRequest) Pagesize(pagesize int32) ApiV1SupervisionIdentitiesGetRequest {
	r.pagesize = &pagesize
	return r
}

func (r ApiV1SupervisionIdentitiesGetRequest) PageSize(pageSize int32) ApiV1SupervisionIdentitiesGetRequest {
	r.pageSize = &pageSize
	return r
}

// Sorting criteria in the format: property:asc/desc. Multiple sort criteria are supported and must be separated with a comma. Example: sort&#x3D;date:desc,name:asc 
func (r ApiV1SupervisionIdentitiesGetRequest) Sort(sort string) ApiV1SupervisionIdentitiesGetRequest {
	r.sort = &sort
	return r
}

func (r ApiV1SupervisionIdentitiesGetRequest) Execute() (*SupervisionIdentitySearchResults, *http.Response, error) {
	return r.ApiService.V1SupervisionIdentitiesGetExecute(r)
}

/*
V1SupervisionIdentitiesGet Search for sorted and paged Supervision Identities 

Search for sorted and paged supervision identities

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiV1SupervisionIdentitiesGetRequest
*/
func (a *SupervisionIdentitiesPreviewApiService) V1SupervisionIdentitiesGet(ctx context.Context) ApiV1SupervisionIdentitiesGetRequest {
	return ApiV1SupervisionIdentitiesGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return SupervisionIdentitySearchResults
func (a *SupervisionIdentitiesPreviewApiService) V1SupervisionIdentitiesGetExecute(r ApiV1SupervisionIdentitiesGetRequest) (*SupervisionIdentitySearchResults, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SupervisionIdentitySearchResults
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SupervisionIdentitiesPreviewApiService.V1SupervisionIdentitiesGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/supervision-identities"

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

type ApiV1SupervisionIdentitiesIdDeleteRequest struct {
	ctx context.Context
	ApiService SupervisionIdentitiesPreviewApi
	id int32
}

func (r ApiV1SupervisionIdentitiesIdDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.V1SupervisionIdentitiesIdDeleteExecute(r)
}

/*
V1SupervisionIdentitiesIdDelete Delete a Supervision Identity with the supplied id 

Deletes a Supervision Identity with the supplied id

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id Supervision Identity identifier
 @return ApiV1SupervisionIdentitiesIdDeleteRequest
*/
func (a *SupervisionIdentitiesPreviewApiService) V1SupervisionIdentitiesIdDelete(ctx context.Context, id int32) ApiV1SupervisionIdentitiesIdDeleteRequest {
	return ApiV1SupervisionIdentitiesIdDeleteRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *SupervisionIdentitiesPreviewApiService) V1SupervisionIdentitiesIdDeleteExecute(r ApiV1SupervisionIdentitiesIdDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SupervisionIdentitiesPreviewApiService.V1SupervisionIdentitiesIdDelete")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/supervision-identities/{id}"
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

type ApiV1SupervisionIdentitiesIdDownloadGetRequest struct {
	ctx context.Context
	ApiService SupervisionIdentitiesPreviewApi
	id int32
}

func (r ApiV1SupervisionIdentitiesIdDownloadGetRequest) Execute() (**os.File, *http.Response, error) {
	return r.ApiService.V1SupervisionIdentitiesIdDownloadGetExecute(r)
}

/*
V1SupervisionIdentitiesIdDownloadGet Download the Supervision Identity .p12 file 

Download the Supervision Identity .p12 file

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id Supervision Identity identifier
 @return ApiV1SupervisionIdentitiesIdDownloadGetRequest
*/
func (a *SupervisionIdentitiesPreviewApiService) V1SupervisionIdentitiesIdDownloadGet(ctx context.Context, id int32) ApiV1SupervisionIdentitiesIdDownloadGetRequest {
	return ApiV1SupervisionIdentitiesIdDownloadGetRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return *os.File
func (a *SupervisionIdentitiesPreviewApiService) V1SupervisionIdentitiesIdDownloadGetExecute(r ApiV1SupervisionIdentitiesIdDownloadGetRequest) (**os.File, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  **os.File
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SupervisionIdentitiesPreviewApiService.V1SupervisionIdentitiesIdDownloadGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/supervision-identities/{id}/download"
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
	localVarHTTPHeaderAccepts := []string{"application/octet-stream", "application/json"}

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

type ApiV1SupervisionIdentitiesIdGetRequest struct {
	ctx context.Context
	ApiService SupervisionIdentitiesPreviewApi
	id int32
}

func (r ApiV1SupervisionIdentitiesIdGetRequest) Execute() (*SupervisionIdentity, *http.Response, error) {
	return r.ApiService.V1SupervisionIdentitiesIdGetExecute(r)
}

/*
V1SupervisionIdentitiesIdGet Retrieve a Supervision Identity with the supplied id 

Retrieves a Supervision Identity with the supplied id

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id Supervision Identity identifier
 @return ApiV1SupervisionIdentitiesIdGetRequest
*/
func (a *SupervisionIdentitiesPreviewApiService) V1SupervisionIdentitiesIdGet(ctx context.Context, id int32) ApiV1SupervisionIdentitiesIdGetRequest {
	return ApiV1SupervisionIdentitiesIdGetRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return SupervisionIdentity
func (a *SupervisionIdentitiesPreviewApiService) V1SupervisionIdentitiesIdGetExecute(r ApiV1SupervisionIdentitiesIdGetRequest) (*SupervisionIdentity, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SupervisionIdentity
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SupervisionIdentitiesPreviewApiService.V1SupervisionIdentitiesIdGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/supervision-identities/{id}"
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

type ApiV1SupervisionIdentitiesIdPutRequest struct {
	ctx context.Context
	ApiService SupervisionIdentitiesPreviewApi
	id int32
	supervisionIdentityUpdate *SupervisionIdentityUpdate
}

func (r ApiV1SupervisionIdentitiesIdPutRequest) SupervisionIdentityUpdate(supervisionIdentityUpdate SupervisionIdentityUpdate) ApiV1SupervisionIdentitiesIdPutRequest {
	r.supervisionIdentityUpdate = &supervisionIdentityUpdate
	return r
}

func (r ApiV1SupervisionIdentitiesIdPutRequest) Execute() (*SupervisionIdentity, *http.Response, error) {
	return r.ApiService.V1SupervisionIdentitiesIdPutExecute(r)
}

/*
V1SupervisionIdentitiesIdPut Update a Supervision Identity with the supplied information 

Updates a Supervision Identity with the supplied information

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id Supervision Identity identifier
 @return ApiV1SupervisionIdentitiesIdPutRequest
*/
func (a *SupervisionIdentitiesPreviewApiService) V1SupervisionIdentitiesIdPut(ctx context.Context, id int32) ApiV1SupervisionIdentitiesIdPutRequest {
	return ApiV1SupervisionIdentitiesIdPutRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return SupervisionIdentity
func (a *SupervisionIdentitiesPreviewApiService) V1SupervisionIdentitiesIdPutExecute(r ApiV1SupervisionIdentitiesIdPutRequest) (*SupervisionIdentity, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SupervisionIdentity
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SupervisionIdentitiesPreviewApiService.V1SupervisionIdentitiesIdPut")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/supervision-identities/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.supervisionIdentityUpdate == nil {
		return localVarReturnValue, nil, reportError("supervisionIdentityUpdate is required and must be specified")
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
	localVarPostBody = r.supervisionIdentityUpdate
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
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
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

type ApiV1SupervisionIdentitiesPostRequest struct {
	ctx context.Context
	ApiService SupervisionIdentitiesPreviewApi
	supervisionIdentityCreate *SupervisionIdentityCreate
}

func (r ApiV1SupervisionIdentitiesPostRequest) SupervisionIdentityCreate(supervisionIdentityCreate SupervisionIdentityCreate) ApiV1SupervisionIdentitiesPostRequest {
	r.supervisionIdentityCreate = &supervisionIdentityCreate
	return r
}

func (r ApiV1SupervisionIdentitiesPostRequest) Execute() (*SupervisionIdentity, *http.Response, error) {
	return r.ApiService.V1SupervisionIdentitiesPostExecute(r)
}

/*
V1SupervisionIdentitiesPost Create a Supervision Identity for the supplied information 

Creates a Supervision Identity for the supplied information

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiV1SupervisionIdentitiesPostRequest
*/
func (a *SupervisionIdentitiesPreviewApiService) V1SupervisionIdentitiesPost(ctx context.Context) ApiV1SupervisionIdentitiesPostRequest {
	return ApiV1SupervisionIdentitiesPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return SupervisionIdentity
func (a *SupervisionIdentitiesPreviewApiService) V1SupervisionIdentitiesPostExecute(r ApiV1SupervisionIdentitiesPostRequest) (*SupervisionIdentity, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SupervisionIdentity
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SupervisionIdentitiesPreviewApiService.V1SupervisionIdentitiesPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/supervision-identities"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.supervisionIdentityCreate == nil {
		return localVarReturnValue, nil, reportError("supervisionIdentityCreate is required and must be specified")
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
	localVarPostBody = r.supervisionIdentityCreate
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
		if localVarHTTPResponse.StatusCode == 500 {
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

type ApiV1SupervisionIdentitiesUploadPostRequest struct {
	ctx context.Context
	ApiService SupervisionIdentitiesPreviewApi
	supervisionIdentityCertificateUpload *SupervisionIdentityCertificateUpload
}

// The base 64 encoded .p12 file alone with other needed information
func (r ApiV1SupervisionIdentitiesUploadPostRequest) SupervisionIdentityCertificateUpload(supervisionIdentityCertificateUpload SupervisionIdentityCertificateUpload) ApiV1SupervisionIdentitiesUploadPostRequest {
	r.supervisionIdentityCertificateUpload = &supervisionIdentityCertificateUpload
	return r
}

func (r ApiV1SupervisionIdentitiesUploadPostRequest) Execute() (*SupervisionIdentity, *http.Response, error) {
	return r.ApiService.V1SupervisionIdentitiesUploadPostExecute(r)
}

/*
V1SupervisionIdentitiesUploadPost Upload the Supervision Identity .p12 file 

Uploads the Supervision Identity .p12 file

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiV1SupervisionIdentitiesUploadPostRequest
*/
func (a *SupervisionIdentitiesPreviewApiService) V1SupervisionIdentitiesUploadPost(ctx context.Context) ApiV1SupervisionIdentitiesUploadPostRequest {
	return ApiV1SupervisionIdentitiesUploadPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return SupervisionIdentity
func (a *SupervisionIdentitiesPreviewApiService) V1SupervisionIdentitiesUploadPostExecute(r ApiV1SupervisionIdentitiesUploadPostRequest) (*SupervisionIdentity, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SupervisionIdentity
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SupervisionIdentitiesPreviewApiService.V1SupervisionIdentitiesUploadPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/supervision-identities/upload"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.supervisionIdentityCertificateUpload == nil {
		return localVarReturnValue, nil, reportError("supervisionIdentityCertificateUpload is required and must be specified")
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
	localVarPostBody = r.supervisionIdentityCertificateUpload
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
		if localVarHTTPResponse.StatusCode == 500 {
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