/*
Warp Builds API Docs

This is the docs for warp builds api for argonaut

API version: 0.4.0
Contact: support@swagger.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package warpbuild

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)


<<<<<<< HEAD
type V1OrganizationAPI interface {
=======
type V1OrganizationApi interface {
>>>>>>> prajjwal-warp-323

	/*
	CreateOrganization Adds a new organisation for a current user

	User can manage multiple tenanats from one account, this api provides user a functionality to add a new tenant/org to the user.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiCreateOrganizationRequest
	*/
	CreateOrganization(ctx context.Context) ApiCreateOrganizationRequest

	// CreateOrganizationExecute executes the request
	//  @return SwitchOrganizationResponse
	CreateOrganizationExecute(r ApiCreateOrganizationRequest) (*SwitchOrganizationResponse, *http.Response, error)

	/*
	GetOrganization Get organization details for the current organization. Current organization is figured from the authorization token

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetOrganizationRequest
	*/
	GetOrganization(ctx context.Context) ApiGetOrganizationRequest

	// GetOrganizationExecute executes the request
	//  @return CommonsOrganization
	GetOrganizationExecute(r ApiGetOrganizationRequest) (*CommonsOrganization, *http.Response, error)

	/*
	ListOrgUsers ListOrgUsers list the users for the current organization

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiListOrgUsersRequest
	*/
	ListOrgUsers(ctx context.Context) ApiListOrgUsersRequest

	// ListOrgUsersExecute executes the request
	//  @return []V1ListUsersForOrganizationResult
	ListOrgUsersExecute(r ApiListOrgUsersRequest) ([]V1ListUsersForOrganizationResult, *http.Response, error)

	/*
	ListUserOrganizations ListUserOrganizations lists all the organization user has access to.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiListUserOrganizationsRequest
	*/
	ListUserOrganizations(ctx context.Context) ApiListUserOrganizationsRequest

	// ListUserOrganizationsExecute executes the request
	//  @return []V1Organization
	ListUserOrganizationsExecute(r ApiListUserOrganizationsRequest) ([]V1Organization, *http.Response, error)

	/*
	UpdateOrganization Updates existing organization based on the fields provided.

	Organization is figured out from the auth token since tokens are specific to organizations.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiUpdateOrganizationRequest
	*/
	UpdateOrganization(ctx context.Context) ApiUpdateOrganizationRequest

	// UpdateOrganizationExecute executes the request
	//  @return CommonsOrganization
	UpdateOrganizationExecute(r ApiUpdateOrganizationRequest) (*CommonsOrganization, *http.Response, error)
}

<<<<<<< HEAD
// V1OrganizationAPIService V1OrganizationAPI service
type V1OrganizationAPIService service

type ApiCreateOrganizationRequest struct {
	ctx context.Context
	ApiService V1OrganizationAPI
=======
// V1OrganizationApiService V1OrganizationApi service
type V1OrganizationApiService service

type ApiCreateOrganizationRequest struct {
	ctx context.Context
	ApiService V1OrganizationApi
>>>>>>> prajjwal-warp-323
}

func (r ApiCreateOrganizationRequest) Execute() (*SwitchOrganizationResponse, *http.Response, error) {
	return r.ApiService.CreateOrganizationExecute(r)
}

/*
CreateOrganization Adds a new organisation for a current user

User can manage multiple tenanats from one account, this api provides user a functionality to add a new tenant/org to the user.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateOrganizationRequest
*/
<<<<<<< HEAD
func (a *V1OrganizationAPIService) CreateOrganization(ctx context.Context) ApiCreateOrganizationRequest {
=======
func (a *V1OrganizationApiService) CreateOrganization(ctx context.Context) ApiCreateOrganizationRequest {
>>>>>>> prajjwal-warp-323
	return ApiCreateOrganizationRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return SwitchOrganizationResponse
<<<<<<< HEAD
func (a *V1OrganizationAPIService) CreateOrganizationExecute(r ApiCreateOrganizationRequest) (*SwitchOrganizationResponse, *http.Response, error) {
=======
func (a *V1OrganizationApiService) CreateOrganizationExecute(r ApiCreateOrganizationRequest) (*SwitchOrganizationResponse, *http.Response, error) {
>>>>>>> prajjwal-warp-323
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SwitchOrganizationResponse
	)

<<<<<<< HEAD
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "V1OrganizationAPIService.CreateOrganization")
=======
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "V1OrganizationApiService.CreateOrganization")
>>>>>>> prajjwal-warp-323
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organization"

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
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["JWTKeyAuth"]; ok {
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
			var v WarpBuildAPIError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
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

type ApiGetOrganizationRequest struct {
	ctx context.Context
<<<<<<< HEAD
	ApiService V1OrganizationAPI
=======
	ApiService V1OrganizationApi
>>>>>>> prajjwal-warp-323
}

func (r ApiGetOrganizationRequest) Execute() (*CommonsOrganization, *http.Response, error) {
	return r.ApiService.GetOrganizationExecute(r)
}

/*
GetOrganization Get organization details for the current organization. Current organization is figured from the authorization token

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetOrganizationRequest
*/
<<<<<<< HEAD
func (a *V1OrganizationAPIService) GetOrganization(ctx context.Context) ApiGetOrganizationRequest {
=======
func (a *V1OrganizationApiService) GetOrganization(ctx context.Context) ApiGetOrganizationRequest {
>>>>>>> prajjwal-warp-323
	return ApiGetOrganizationRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CommonsOrganization
<<<<<<< HEAD
func (a *V1OrganizationAPIService) GetOrganizationExecute(r ApiGetOrganizationRequest) (*CommonsOrganization, *http.Response, error) {
=======
func (a *V1OrganizationApiService) GetOrganizationExecute(r ApiGetOrganizationRequest) (*CommonsOrganization, *http.Response, error) {
>>>>>>> prajjwal-warp-323
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CommonsOrganization
	)

<<<<<<< HEAD
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "V1OrganizationAPIService.GetOrganization")
=======
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "V1OrganizationApiService.GetOrganization")
>>>>>>> prajjwal-warp-323
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organization"

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
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["JWTKeyAuth"]; ok {
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
			var v WarpBuildAPIError
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
			var v WarpBuildAPIError
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

type ApiListOrgUsersRequest struct {
	ctx context.Context
<<<<<<< HEAD
	ApiService V1OrganizationAPI
=======
	ApiService V1OrganizationApi
>>>>>>> prajjwal-warp-323
}

func (r ApiListOrgUsersRequest) Execute() ([]V1ListUsersForOrganizationResult, *http.Response, error) {
	return r.ApiService.ListOrgUsersExecute(r)
}

/*
ListOrgUsers ListOrgUsers list the users for the current organization

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListOrgUsersRequest
*/
<<<<<<< HEAD
func (a *V1OrganizationAPIService) ListOrgUsers(ctx context.Context) ApiListOrgUsersRequest {
=======
func (a *V1OrganizationApiService) ListOrgUsers(ctx context.Context) ApiListOrgUsersRequest {
>>>>>>> prajjwal-warp-323
	return ApiListOrgUsersRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []V1ListUsersForOrganizationResult
<<<<<<< HEAD
func (a *V1OrganizationAPIService) ListOrgUsersExecute(r ApiListOrgUsersRequest) ([]V1ListUsersForOrganizationResult, *http.Response, error) {
=======
func (a *V1OrganizationApiService) ListOrgUsersExecute(r ApiListOrgUsersRequest) ([]V1ListUsersForOrganizationResult, *http.Response, error) {
>>>>>>> prajjwal-warp-323
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []V1ListUsersForOrganizationResult
	)

<<<<<<< HEAD
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "V1OrganizationAPIService.ListOrgUsers")
=======
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "V1OrganizationApiService.ListOrgUsers")
>>>>>>> prajjwal-warp-323
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organization/users"

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
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["JWTKeyAuth"]; ok {
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
			var v WarpBuildAPIError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
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

type ApiListUserOrganizationsRequest struct {
	ctx context.Context
<<<<<<< HEAD
	ApiService V1OrganizationAPI
=======
	ApiService V1OrganizationApi
>>>>>>> prajjwal-warp-323
}

func (r ApiListUserOrganizationsRequest) Execute() ([]V1Organization, *http.Response, error) {
	return r.ApiService.ListUserOrganizationsExecute(r)
}

/*
ListUserOrganizations ListUserOrganizations lists all the organization user has access to.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListUserOrganizationsRequest
*/
<<<<<<< HEAD
func (a *V1OrganizationAPIService) ListUserOrganizations(ctx context.Context) ApiListUserOrganizationsRequest {
=======
func (a *V1OrganizationApiService) ListUserOrganizations(ctx context.Context) ApiListUserOrganizationsRequest {
>>>>>>> prajjwal-warp-323
	return ApiListUserOrganizationsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []V1Organization
<<<<<<< HEAD
func (a *V1OrganizationAPIService) ListUserOrganizationsExecute(r ApiListUserOrganizationsRequest) ([]V1Organization, *http.Response, error) {
=======
func (a *V1OrganizationApiService) ListUserOrganizationsExecute(r ApiListUserOrganizationsRequest) ([]V1Organization, *http.Response, error) {
>>>>>>> prajjwal-warp-323
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []V1Organization
	)

<<<<<<< HEAD
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "V1OrganizationAPIService.ListUserOrganizations")
=======
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "V1OrganizationApiService.ListUserOrganizations")
>>>>>>> prajjwal-warp-323
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations"

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
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["JWTKeyAuth"]; ok {
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
			var v WarpBuildAPIError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
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

type ApiUpdateOrganizationRequest struct {
	ctx context.Context
<<<<<<< HEAD
	ApiService V1OrganizationAPI
=======
	ApiService V1OrganizationApi
>>>>>>> prajjwal-warp-323
	body *UpdateOrganizationRequest
}

// Update existing organization body
func (r ApiUpdateOrganizationRequest) Body(body UpdateOrganizationRequest) ApiUpdateOrganizationRequest {
	r.body = &body
	return r
}

func (r ApiUpdateOrganizationRequest) Execute() (*CommonsOrganization, *http.Response, error) {
	return r.ApiService.UpdateOrganizationExecute(r)
}

/*
UpdateOrganization Updates existing organization based on the fields provided.

Organization is figured out from the auth token since tokens are specific to organizations.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiUpdateOrganizationRequest
*/
<<<<<<< HEAD
func (a *V1OrganizationAPIService) UpdateOrganization(ctx context.Context) ApiUpdateOrganizationRequest {
=======
func (a *V1OrganizationApiService) UpdateOrganization(ctx context.Context) ApiUpdateOrganizationRequest {
>>>>>>> prajjwal-warp-323
	return ApiUpdateOrganizationRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CommonsOrganization
<<<<<<< HEAD
func (a *V1OrganizationAPIService) UpdateOrganizationExecute(r ApiUpdateOrganizationRequest) (*CommonsOrganization, *http.Response, error) {
=======
func (a *V1OrganizationApiService) UpdateOrganizationExecute(r ApiUpdateOrganizationRequest) (*CommonsOrganization, *http.Response, error) {
>>>>>>> prajjwal-warp-323
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CommonsOrganization
	)

<<<<<<< HEAD
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "V1OrganizationAPIService.UpdateOrganization")
=======
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "V1OrganizationApiService.UpdateOrganization")
>>>>>>> prajjwal-warp-323
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organization"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.body == nil {
		return localVarReturnValue, nil, reportError("body is required and must be specified")
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
	// body params
	localVarPostBody = r.body
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["JWTKeyAuth"]; ok {
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
			var v WarpBuildAPIError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
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
