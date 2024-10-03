/*
Taikun - WebApi

This Api will be responsible for overall data distribution and authorization.

API version: v1
Contact: noreply@taikun.cloud
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package taikuncore

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)


// GenericKubernetesCloudCredentialAPIService GenericKubernetesCloudCredentialAPI service
type GenericKubernetesCloudCredentialAPIService service

type ApiGenericKubernetesListRequest struct {
	ctx context.Context
	ApiService *GenericKubernetesCloudCredentialAPIService
	limit *int32
	offset *int32
	organizationId *int32
	sortBy *string
	sortDirection *string
	search *string
	searchId *string
	id *int32
}

func (r ApiGenericKubernetesListRequest) Limit(limit int32) ApiGenericKubernetesListRequest {
	r.limit = &limit
	return r
}

func (r ApiGenericKubernetesListRequest) Offset(offset int32) ApiGenericKubernetesListRequest {
	r.offset = &offset
	return r
}

func (r ApiGenericKubernetesListRequest) OrganizationId(organizationId int32) ApiGenericKubernetesListRequest {
	r.organizationId = &organizationId
	return r
}

func (r ApiGenericKubernetesListRequest) SortBy(sortBy string) ApiGenericKubernetesListRequest {
	r.sortBy = &sortBy
	return r
}

func (r ApiGenericKubernetesListRequest) SortDirection(sortDirection string) ApiGenericKubernetesListRequest {
	r.sortDirection = &sortDirection
	return r
}

func (r ApiGenericKubernetesListRequest) Search(search string) ApiGenericKubernetesListRequest {
	r.search = &search
	return r
}

func (r ApiGenericKubernetesListRequest) SearchId(searchId string) ApiGenericKubernetesListRequest {
	r.searchId = &searchId
	return r
}

func (r ApiGenericKubernetesListRequest) Id(id int32) ApiGenericKubernetesListRequest {
	r.id = &id
	return r
}

func (r ApiGenericKubernetesListRequest) Execute() (*GenericKubernetesList, *http.Response, error) {
	return r.ApiService.GenericKubernetesListExecute(r)
}

/*
GenericKubernetesList Retrieve list of generic kubernetes cloud credentials

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGenericKubernetesListRequest
*/
func (a *GenericKubernetesCloudCredentialAPIService) GenericKubernetesList(ctx context.Context) ApiGenericKubernetesListRequest {
	return ApiGenericKubernetesListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GenericKubernetesList
func (a *GenericKubernetesCloudCredentialAPIService) GenericKubernetesListExecute(r ApiGenericKubernetesListRequest) (*GenericKubernetesList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GenericKubernetesList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GenericKubernetesCloudCredentialAPIService.GenericKubernetesList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/generic-kubernetes/list"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "Limit", r.limit, "form", "")
	}
	if r.offset != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "Offset", r.offset, "form", "")
	}
	if r.organizationId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "OrganizationId", r.organizationId, "form", "")
	}
	if r.sortBy != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "SortBy", r.sortBy, "form", "")
	}
	if r.sortDirection != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "SortDirection", r.sortDirection, "form", "")
	}
	if r.search != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "Search", r.search, "form", "")
	}
	if r.searchId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "SearchId", r.searchId, "form", "")
	}
	if r.id != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "Id", r.id, "form", "")
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
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["Bearer"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
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
		if localVarHTTPResponse.StatusCode == 400 {
			var v ProblemDetails
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ProblemDetails
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ProblemDetails
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ProblemDetails
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ProblemDetails
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
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
