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


// ProjectQuotasAPIService ProjectQuotasAPI service
type ProjectQuotasAPIService service

type ApiProjectquotasListRequest struct {
	ctx context.Context
	ApiService *ProjectQuotasAPIService
	sortBy *string
	sortDirection *string
	search *string
	limit *int32
	offset *int32
	startRam *float64
	endRam *float64
	startDiskSize *int64
	endDiskSize *int64
	startCpu *int32
	endCpu *int32
	organizationId *int32
	id *int32
}

func (r ApiProjectquotasListRequest) SortBy(sortBy string) ApiProjectquotasListRequest {
	r.sortBy = &sortBy
	return r
}

func (r ApiProjectquotasListRequest) SortDirection(sortDirection string) ApiProjectquotasListRequest {
	r.sortDirection = &sortDirection
	return r
}

func (r ApiProjectquotasListRequest) Search(search string) ApiProjectquotasListRequest {
	r.search = &search
	return r
}

func (r ApiProjectquotasListRequest) Limit(limit int32) ApiProjectquotasListRequest {
	r.limit = &limit
	return r
}

func (r ApiProjectquotasListRequest) Offset(offset int32) ApiProjectquotasListRequest {
	r.offset = &offset
	return r
}

func (r ApiProjectquotasListRequest) StartRam(startRam float64) ApiProjectquotasListRequest {
	r.startRam = &startRam
	return r
}

func (r ApiProjectquotasListRequest) EndRam(endRam float64) ApiProjectquotasListRequest {
	r.endRam = &endRam
	return r
}

func (r ApiProjectquotasListRequest) StartDiskSize(startDiskSize int64) ApiProjectquotasListRequest {
	r.startDiskSize = &startDiskSize
	return r
}

func (r ApiProjectquotasListRequest) EndDiskSize(endDiskSize int64) ApiProjectquotasListRequest {
	r.endDiskSize = &endDiskSize
	return r
}

func (r ApiProjectquotasListRequest) StartCpu(startCpu int32) ApiProjectquotasListRequest {
	r.startCpu = &startCpu
	return r
}

func (r ApiProjectquotasListRequest) EndCpu(endCpu int32) ApiProjectquotasListRequest {
	r.endCpu = &endCpu
	return r
}

func (r ApiProjectquotasListRequest) OrganizationId(organizationId int32) ApiProjectquotasListRequest {
	r.organizationId = &organizationId
	return r
}

func (r ApiProjectquotasListRequest) Id(id int32) ApiProjectquotasListRequest {
	r.id = &id
	return r
}

func (r ApiProjectquotasListRequest) Execute() (*ProjectQuotaList, *http.Response, error) {
	return r.ApiService.ProjectquotasListExecute(r)
}

/*
ProjectquotasList Retrieve all project quotas

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiProjectquotasListRequest
*/
func (a *ProjectQuotasAPIService) ProjectquotasList(ctx context.Context) ApiProjectquotasListRequest {
	return ApiProjectquotasListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ProjectQuotaList
func (a *ProjectQuotasAPIService) ProjectquotasListExecute(r ApiProjectquotasListRequest) (*ProjectQuotaList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ProjectQuotaList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProjectQuotasAPIService.ProjectquotasList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/projectquotas"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.sortBy == nil {
		return localVarReturnValue, nil, reportError("sortBy is required and must be specified")
	}
	if r.sortDirection == nil {
		return localVarReturnValue, nil, reportError("sortDirection is required and must be specified")
	}
	if r.search == nil {
		return localVarReturnValue, nil, reportError("search is required and must be specified")
	}

	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "Limit", r.limit, "form", "")
	}
	if r.offset != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "Offset", r.offset, "form", "")
	}
	parameterAddToHeaderOrQuery(localVarQueryParams, "SortBy", r.sortBy, "form", "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "SortDirection", r.sortDirection, "form", "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "Search", r.search, "form", "")
	if r.startRam != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "StartRam", r.startRam, "form", "")
	}
	if r.endRam != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "EndRam", r.endRam, "form", "")
	}
	if r.startDiskSize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "StartDiskSize", r.startDiskSize, "form", "")
	}
	if r.endDiskSize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "EndDiskSize", r.endDiskSize, "form", "")
	}
	if r.startCpu != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "StartCpu", r.startCpu, "form", "")
	}
	if r.endCpu != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "EndCpu", r.endCpu, "form", "")
	}
	if r.organizationId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "OrganizationId", r.organizationId, "form", "")
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

type ApiProjectquotasUpdateRequest struct {
	ctx context.Context
	ApiService *ProjectQuotasAPIService
	updateQuotaCommand *UpdateQuotaCommand
}

func (r ApiProjectquotasUpdateRequest) UpdateQuotaCommand(updateQuotaCommand UpdateQuotaCommand) ApiProjectquotasUpdateRequest {
	r.updateQuotaCommand = &updateQuotaCommand
	return r
}

func (r ApiProjectquotasUpdateRequest) Execute() (*http.Response, error) {
	return r.ApiService.ProjectquotasUpdateExecute(r)
}

/*
ProjectquotasUpdate Edit project quota

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiProjectquotasUpdateRequest
*/
func (a *ProjectQuotasAPIService) ProjectquotasUpdate(ctx context.Context) ApiProjectquotasUpdateRequest {
	return ApiProjectquotasUpdateRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *ProjectQuotasAPIService) ProjectquotasUpdateExecute(r ApiProjectquotasUpdateRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProjectQuotasAPIService.ProjectquotasUpdate")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/projectquotas/update"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.updateQuotaCommand == nil {
		return nil, reportError("updateQuotaCommand is required and must be specified")
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
	localVarPostBody = r.updateQuotaCommand
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
		if localVarHTTPResponse.StatusCode == 400 {
			var v ProblemDetails
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ProblemDetails
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ProblemDetails
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ProblemDetails
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ProblemDetails
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}
