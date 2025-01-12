# \TeamViewerRemoteAdministrationApi

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PreviewRemoteAdministrationConfigurationsTeamViewerConfigurationIdSessionsGet**](TeamViewerRemoteAdministrationApi.md#PreviewRemoteAdministrationConfigurationsTeamViewerConfigurationIdSessionsGet) | **Get** /preview/remote-administration-configurations/team-viewer/{configurationId}/sessions | Get a paginated list of sessions 
[**PreviewRemoteAdministrationConfigurationsTeamViewerConfigurationIdSessionsPost**](TeamViewerRemoteAdministrationApi.md#PreviewRemoteAdministrationConfigurationsTeamViewerConfigurationIdSessionsPost) | **Post** /preview/remote-administration-configurations/team-viewer/{configurationId}/sessions | Create a new session
[**PreviewRemoteAdministrationConfigurationsTeamViewerConfigurationIdSessionsSessionIdClosePost**](TeamViewerRemoteAdministrationApi.md#PreviewRemoteAdministrationConfigurationsTeamViewerConfigurationIdSessionsSessionIdClosePost) | **Post** /preview/remote-administration-configurations/team-viewer/{configurationId}/sessions/{sessionId}/close | Close a session
[**PreviewRemoteAdministrationConfigurationsTeamViewerConfigurationIdSessionsSessionIdGet**](TeamViewerRemoteAdministrationApi.md#PreviewRemoteAdministrationConfigurationsTeamViewerConfigurationIdSessionsSessionIdGet) | **Get** /preview/remote-administration-configurations/team-viewer/{configurationId}/sessions/{sessionId} | Get a session by its ID 
[**PreviewRemoteAdministrationConfigurationsTeamViewerConfigurationIdSessionsSessionIdResendNotificationPost**](TeamViewerRemoteAdministrationApi.md#PreviewRemoteAdministrationConfigurationsTeamViewerConfigurationIdSessionsSessionIdResendNotificationPost) | **Post** /preview/remote-administration-configurations/team-viewer/{configurationId}/sessions/{sessionId}/resend-notification | Resend nofications for a session
[**PreviewRemoteAdministrationConfigurationsTeamViewerConfigurationIdSessionsSessionIdStatusGet**](TeamViewerRemoteAdministrationApi.md#PreviewRemoteAdministrationConfigurationsTeamViewerConfigurationIdSessionsSessionIdStatusGet) | **Get** /preview/remote-administration-configurations/team-viewer/{configurationId}/sessions/{sessionId}/status | Get a session status by its ID 
[**PreviewRemoteAdministrationConfigurationsTeamViewerIdDelete**](TeamViewerRemoteAdministrationApi.md#PreviewRemoteAdministrationConfigurationsTeamViewerIdDelete) | **Delete** /preview/remote-administration-configurations/team-viewer/{id} | Delete Team Viewer Remote Administration connection configuration
[**PreviewRemoteAdministrationConfigurationsTeamViewerIdGet**](TeamViewerRemoteAdministrationApi.md#PreviewRemoteAdministrationConfigurationsTeamViewerIdGet) | **Get** /preview/remote-administration-configurations/team-viewer/{id} | Get Team Viewer Remote Administration connection configuration
[**PreviewRemoteAdministrationConfigurationsTeamViewerIdPatch**](TeamViewerRemoteAdministrationApi.md#PreviewRemoteAdministrationConfigurationsTeamViewerIdPatch) | **Patch** /preview/remote-administration-configurations/team-viewer/{id} | Update Team Viewer Remote Administration connection configuration
[**PreviewRemoteAdministrationConfigurationsTeamViewerIdStatusGet**](TeamViewerRemoteAdministrationApi.md#PreviewRemoteAdministrationConfigurationsTeamViewerIdStatusGet) | **Get** /preview/remote-administration-configurations/team-viewer/{id}/status | Get Team Viewer Remote Administration connection status
[**PreviewRemoteAdministrationConfigurationsTeamViewerPost**](TeamViewerRemoteAdministrationApi.md#PreviewRemoteAdministrationConfigurationsTeamViewerPost) | **Post** /preview/remote-administration-configurations/team-viewer | Create Team Viewer Remote Administration connection configuration



## PreviewRemoteAdministrationConfigurationsTeamViewerConfigurationIdSessionsGet

> SessionDetailsSearchResults PreviewRemoteAdministrationConfigurationsTeamViewerConfigurationIdSessionsGet(ctx, configurationId).Page(page).PageSize(pageSize).Filter(filter).Execute()

Get a paginated list of sessions 



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    configurationId := "configurationId_example" // string | ID of the Team Viewer connection configuration
    page := int32(56) // int32 |  (optional) (default to 0)
    pageSize := int32(56) // int32 |  (optional) (default to 100)
    filter := "deviceId==1 and deviceType=="COMPUTER" and state=="OPEN"" // string | Query in the RSQL format, allowing to filter sessions collection. Default filter is empty query - returning all results for the requested page.  Fields allowed in the query: `deviceId`, `deviceType`, `state`  This param can be combined with paging.  (optional) (default to "")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TeamViewerRemoteAdministrationApi.PreviewRemoteAdministrationConfigurationsTeamViewerConfigurationIdSessionsGet(context.Background(), configurationId).Page(page).PageSize(pageSize).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamViewerRemoteAdministrationApi.PreviewRemoteAdministrationConfigurationsTeamViewerConfigurationIdSessionsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PreviewRemoteAdministrationConfigurationsTeamViewerConfigurationIdSessionsGet`: SessionDetailsSearchResults
    fmt.Fprintf(os.Stdout, "Response from `TeamViewerRemoteAdministrationApi.PreviewRemoteAdministrationConfigurationsTeamViewerConfigurationIdSessionsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configurationId** | **string** | ID of the Team Viewer connection configuration | 

### Other Parameters

Other parameters are passed through a pointer to a apiPreviewRemoteAdministrationConfigurationsTeamViewerConfigurationIdSessionsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** |  | [default to 0]
 **pageSize** | **int32** |  | [default to 100]
 **filter** | **string** | Query in the RSQL format, allowing to filter sessions collection. Default filter is empty query - returning all results for the requested page.  Fields allowed in the query: &#x60;deviceId&#x60;, &#x60;deviceType&#x60;, &#x60;state&#x60;  This param can be combined with paging.  | [default to &quot;&quot;]

### Return type

[**SessionDetailsSearchResults**](SessionDetailsSearchResults.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PreviewRemoteAdministrationConfigurationsTeamViewerConfigurationIdSessionsPost

> HrefResponse PreviewRemoteAdministrationConfigurationsTeamViewerConfigurationIdSessionsPost(ctx, configurationId).SessionCandidateRequest(sessionCandidateRequest).Execute()

Create a new session



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    configurationId := "configurationId_example" // string | ID of the Team Viewer connection configuration
    sessionCandidateRequest := *openapiclient.NewSessionCandidateRequest("1", "COMPUTER", "Customer reported that he cannot install application X") // SessionCandidateRequest | Team Viewer session attributes

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TeamViewerRemoteAdministrationApi.PreviewRemoteAdministrationConfigurationsTeamViewerConfigurationIdSessionsPost(context.Background(), configurationId).SessionCandidateRequest(sessionCandidateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamViewerRemoteAdministrationApi.PreviewRemoteAdministrationConfigurationsTeamViewerConfigurationIdSessionsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PreviewRemoteAdministrationConfigurationsTeamViewerConfigurationIdSessionsPost`: HrefResponse
    fmt.Fprintf(os.Stdout, "Response from `TeamViewerRemoteAdministrationApi.PreviewRemoteAdministrationConfigurationsTeamViewerConfigurationIdSessionsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configurationId** | **string** | ID of the Team Viewer connection configuration | 

### Other Parameters

Other parameters are passed through a pointer to a apiPreviewRemoteAdministrationConfigurationsTeamViewerConfigurationIdSessionsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sessionCandidateRequest** | [**SessionCandidateRequest**](SessionCandidateRequest.md) | Team Viewer session attributes | 

### Return type

[**HrefResponse**](HrefResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PreviewRemoteAdministrationConfigurationsTeamViewerConfigurationIdSessionsSessionIdClosePost

> PreviewRemoteAdministrationConfigurationsTeamViewerConfigurationIdSessionsSessionIdClosePost(ctx, configurationId, sessionId).Execute()

Close a session



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    configurationId := "configurationId_example" // string | ID of the Team Viewer connection configuration
    sessionId := "sessionId_example" // string | ID of the Team Viewer session

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TeamViewerRemoteAdministrationApi.PreviewRemoteAdministrationConfigurationsTeamViewerConfigurationIdSessionsSessionIdClosePost(context.Background(), configurationId, sessionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamViewerRemoteAdministrationApi.PreviewRemoteAdministrationConfigurationsTeamViewerConfigurationIdSessionsSessionIdClosePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configurationId** | **string** | ID of the Team Viewer connection configuration | 
**sessionId** | **string** | ID of the Team Viewer session | 

### Other Parameters

Other parameters are passed through a pointer to a apiPreviewRemoteAdministrationConfigurationsTeamViewerConfigurationIdSessionsSessionIdClosePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PreviewRemoteAdministrationConfigurationsTeamViewerConfigurationIdSessionsSessionIdGet

> SessionDetails PreviewRemoteAdministrationConfigurationsTeamViewerConfigurationIdSessionsSessionIdGet(ctx, configurationId, sessionId).Execute()

Get a session by its ID 



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    configurationId := "configurationId_example" // string | ID of the Team Viewer connection configuration
    sessionId := "sessionId_example" // string | ID of the Team Viewer session

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TeamViewerRemoteAdministrationApi.PreviewRemoteAdministrationConfigurationsTeamViewerConfigurationIdSessionsSessionIdGet(context.Background(), configurationId, sessionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamViewerRemoteAdministrationApi.PreviewRemoteAdministrationConfigurationsTeamViewerConfigurationIdSessionsSessionIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PreviewRemoteAdministrationConfigurationsTeamViewerConfigurationIdSessionsSessionIdGet`: SessionDetails
    fmt.Fprintf(os.Stdout, "Response from `TeamViewerRemoteAdministrationApi.PreviewRemoteAdministrationConfigurationsTeamViewerConfigurationIdSessionsSessionIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configurationId** | **string** | ID of the Team Viewer connection configuration | 
**sessionId** | **string** | ID of the Team Viewer session | 

### Other Parameters

Other parameters are passed through a pointer to a apiPreviewRemoteAdministrationConfigurationsTeamViewerConfigurationIdSessionsSessionIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**SessionDetails**](SessionDetails.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PreviewRemoteAdministrationConfigurationsTeamViewerConfigurationIdSessionsSessionIdResendNotificationPost

> PreviewRemoteAdministrationConfigurationsTeamViewerConfigurationIdSessionsSessionIdResendNotificationPost(ctx, configurationId, sessionId).Execute()

Resend nofications for a session



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    configurationId := "configurationId_example" // string | ID of the Team Viewer connection configuration
    sessionId := "sessionId_example" // string | ID of the Team Viewer session

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TeamViewerRemoteAdministrationApi.PreviewRemoteAdministrationConfigurationsTeamViewerConfigurationIdSessionsSessionIdResendNotificationPost(context.Background(), configurationId, sessionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamViewerRemoteAdministrationApi.PreviewRemoteAdministrationConfigurationsTeamViewerConfigurationIdSessionsSessionIdResendNotificationPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configurationId** | **string** | ID of the Team Viewer connection configuration | 
**sessionId** | **string** | ID of the Team Viewer session | 

### Other Parameters

Other parameters are passed through a pointer to a apiPreviewRemoteAdministrationConfigurationsTeamViewerConfigurationIdSessionsSessionIdResendNotificationPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PreviewRemoteAdministrationConfigurationsTeamViewerConfigurationIdSessionsSessionIdStatusGet

> SessionStatus PreviewRemoteAdministrationConfigurationsTeamViewerConfigurationIdSessionsSessionIdStatusGet(ctx, configurationId, sessionId).Execute()

Get a session status by its ID 



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    configurationId := "configurationId_example" // string | ID of the Team Viewer connection configuration
    sessionId := "sessionId_example" // string | ID of the Team Viewer session

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TeamViewerRemoteAdministrationApi.PreviewRemoteAdministrationConfigurationsTeamViewerConfigurationIdSessionsSessionIdStatusGet(context.Background(), configurationId, sessionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamViewerRemoteAdministrationApi.PreviewRemoteAdministrationConfigurationsTeamViewerConfigurationIdSessionsSessionIdStatusGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PreviewRemoteAdministrationConfigurationsTeamViewerConfigurationIdSessionsSessionIdStatusGet`: SessionStatus
    fmt.Fprintf(os.Stdout, "Response from `TeamViewerRemoteAdministrationApi.PreviewRemoteAdministrationConfigurationsTeamViewerConfigurationIdSessionsSessionIdStatusGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configurationId** | **string** | ID of the Team Viewer connection configuration | 
**sessionId** | **string** | ID of the Team Viewer session | 

### Other Parameters

Other parameters are passed through a pointer to a apiPreviewRemoteAdministrationConfigurationsTeamViewerConfigurationIdSessionsSessionIdStatusGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**SessionStatus**](SessionStatus.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PreviewRemoteAdministrationConfigurationsTeamViewerIdDelete

> PreviewRemoteAdministrationConfigurationsTeamViewerIdDelete(ctx, id).Execute()

Delete Team Viewer Remote Administration connection configuration



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | ID of the Team Viewer connection configuration

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TeamViewerRemoteAdministrationApi.PreviewRemoteAdministrationConfigurationsTeamViewerIdDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamViewerRemoteAdministrationApi.PreviewRemoteAdministrationConfigurationsTeamViewerIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the Team Viewer connection configuration | 

### Other Parameters

Other parameters are passed through a pointer to a apiPreviewRemoteAdministrationConfigurationsTeamViewerIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PreviewRemoteAdministrationConfigurationsTeamViewerIdGet

> ConnectionConfigurationResponse PreviewRemoteAdministrationConfigurationsTeamViewerIdGet(ctx, id).Execute()

Get Team Viewer Remote Administration connection configuration



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | ID of the Team Viewer connection configuration

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TeamViewerRemoteAdministrationApi.PreviewRemoteAdministrationConfigurationsTeamViewerIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamViewerRemoteAdministrationApi.PreviewRemoteAdministrationConfigurationsTeamViewerIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PreviewRemoteAdministrationConfigurationsTeamViewerIdGet`: ConnectionConfigurationResponse
    fmt.Fprintf(os.Stdout, "Response from `TeamViewerRemoteAdministrationApi.PreviewRemoteAdministrationConfigurationsTeamViewerIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the Team Viewer connection configuration | 

### Other Parameters

Other parameters are passed through a pointer to a apiPreviewRemoteAdministrationConfigurationsTeamViewerIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ConnectionConfigurationResponse**](ConnectionConfigurationResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PreviewRemoteAdministrationConfigurationsTeamViewerIdPatch

> ConnectionConfigurationResponse PreviewRemoteAdministrationConfigurationsTeamViewerIdPatch(ctx, id).ConnectionConfigurationUpdateRequest(connectionConfigurationUpdateRequest).Execute()

Update Team Viewer Remote Administration connection configuration



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | ID of the Team Viewer connection configuration
    connectionConfigurationUpdateRequest := *openapiclient.NewConnectionConfigurationUpdateRequest() // ConnectionConfigurationUpdateRequest | Team Viewer connection configuration update request

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TeamViewerRemoteAdministrationApi.PreviewRemoteAdministrationConfigurationsTeamViewerIdPatch(context.Background(), id).ConnectionConfigurationUpdateRequest(connectionConfigurationUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamViewerRemoteAdministrationApi.PreviewRemoteAdministrationConfigurationsTeamViewerIdPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PreviewRemoteAdministrationConfigurationsTeamViewerIdPatch`: ConnectionConfigurationResponse
    fmt.Fprintf(os.Stdout, "Response from `TeamViewerRemoteAdministrationApi.PreviewRemoteAdministrationConfigurationsTeamViewerIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the Team Viewer connection configuration | 

### Other Parameters

Other parameters are passed through a pointer to a apiPreviewRemoteAdministrationConfigurationsTeamViewerIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **connectionConfigurationUpdateRequest** | [**ConnectionConfigurationUpdateRequest**](ConnectionConfigurationUpdateRequest.md) | Team Viewer connection configuration update request | 

### Return type

[**ConnectionConfigurationResponse**](ConnectionConfigurationResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PreviewRemoteAdministrationConfigurationsTeamViewerIdStatusGet

> ConnectionConfigurationStatusResponse PreviewRemoteAdministrationConfigurationsTeamViewerIdStatusGet(ctx, id).Execute()

Get Team Viewer Remote Administration connection status



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | ID of the Team Viewer connection configuration

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TeamViewerRemoteAdministrationApi.PreviewRemoteAdministrationConfigurationsTeamViewerIdStatusGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamViewerRemoteAdministrationApi.PreviewRemoteAdministrationConfigurationsTeamViewerIdStatusGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PreviewRemoteAdministrationConfigurationsTeamViewerIdStatusGet`: ConnectionConfigurationStatusResponse
    fmt.Fprintf(os.Stdout, "Response from `TeamViewerRemoteAdministrationApi.PreviewRemoteAdministrationConfigurationsTeamViewerIdStatusGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the Team Viewer connection configuration | 

### Other Parameters

Other parameters are passed through a pointer to a apiPreviewRemoteAdministrationConfigurationsTeamViewerIdStatusGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ConnectionConfigurationStatusResponse**](ConnectionConfigurationStatusResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PreviewRemoteAdministrationConfigurationsTeamViewerPost

> HrefResponse PreviewRemoteAdministrationConfigurationsTeamViewerPost(ctx).ConnectionConfigurationCandidateRequest(connectionConfigurationCandidateRequest).Execute()

Create Team Viewer Remote Administration connection configuration



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    connectionConfigurationCandidateRequest := *openapiclient.NewConnectionConfigurationCandidateRequest("1", "teamViewerConfiguration", "12847340-nPAX96bsaADH4Gz6K6i2", true, int32(15)) // ConnectionConfigurationCandidateRequest | Team Viewer connection configuration create definition

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TeamViewerRemoteAdministrationApi.PreviewRemoteAdministrationConfigurationsTeamViewerPost(context.Background()).ConnectionConfigurationCandidateRequest(connectionConfigurationCandidateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamViewerRemoteAdministrationApi.PreviewRemoteAdministrationConfigurationsTeamViewerPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PreviewRemoteAdministrationConfigurationsTeamViewerPost`: HrefResponse
    fmt.Fprintf(os.Stdout, "Response from `TeamViewerRemoteAdministrationApi.PreviewRemoteAdministrationConfigurationsTeamViewerPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPreviewRemoteAdministrationConfigurationsTeamViewerPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **connectionConfigurationCandidateRequest** | [**ConnectionConfigurationCandidateRequest**](ConnectionConfigurationCandidateRequest.md) | Team Viewer connection configuration create definition | 

### Return type

[**HrefResponse**](HrefResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

