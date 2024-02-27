/*
Daytona Server API

Daytona Server API

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package serverapiclient

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


// PluginAPIService PluginAPI service
type PluginAPIService service

type ApiInstallAgentServicePluginRequest struct {
	ctx context.Context
	ApiService *PluginAPIService
	plugin *InstallPluginRequest
}

// Plugin to install
func (r ApiInstallAgentServicePluginRequest) Plugin(plugin InstallPluginRequest) ApiInstallAgentServicePluginRequest {
	r.plugin = &plugin
	return r
}

func (r ApiInstallAgentServicePluginRequest) Execute() (*http.Response, error) {
	return r.ApiService.InstallAgentServicePluginExecute(r)
}

/*
InstallAgentServicePlugin Install an agent service plugin

Install an agent service plugin

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiInstallAgentServicePluginRequest
*/
func (a *PluginAPIService) InstallAgentServicePlugin(ctx context.Context) ApiInstallAgentServicePluginRequest {
	return ApiInstallAgentServicePluginRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *PluginAPIService) InstallAgentServicePluginExecute(r ApiInstallAgentServicePluginRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PluginAPIService.InstallAgentServicePlugin")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/plugin/agent-service/install"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.plugin == nil {
		return nil, reportError("plugin is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.plugin
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
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiInstallProvisionerPluginRequest struct {
	ctx context.Context
	ApiService *PluginAPIService
	plugin *InstallPluginRequest
}

// Plugin to install
func (r ApiInstallProvisionerPluginRequest) Plugin(plugin InstallPluginRequest) ApiInstallProvisionerPluginRequest {
	r.plugin = &plugin
	return r
}

func (r ApiInstallProvisionerPluginRequest) Execute() (*http.Response, error) {
	return r.ApiService.InstallProvisionerPluginExecute(r)
}

/*
InstallProvisionerPlugin Install a provisioner plugin

Install a provisioner plugin

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiInstallProvisionerPluginRequest
*/
func (a *PluginAPIService) InstallProvisionerPlugin(ctx context.Context) ApiInstallProvisionerPluginRequest {
	return ApiInstallProvisionerPluginRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *PluginAPIService) InstallProvisionerPluginExecute(r ApiInstallProvisionerPluginRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PluginAPIService.InstallProvisionerPlugin")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/plugin/provisioner/install"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.plugin == nil {
		return nil, reportError("plugin is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.plugin
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
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiListAgentServicePluginsRequest struct {
	ctx context.Context
	ApiService *PluginAPIService
}

func (r ApiListAgentServicePluginsRequest) Execute() ([]AgentServicePlugin, *http.Response, error) {
	return r.ApiService.ListAgentServicePluginsExecute(r)
}

/*
ListAgentServicePlugins List agent service plugins

List agent service plugins

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListAgentServicePluginsRequest
*/
func (a *PluginAPIService) ListAgentServicePlugins(ctx context.Context) ApiListAgentServicePluginsRequest {
	return ApiListAgentServicePluginsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []AgentServicePlugin
func (a *PluginAPIService) ListAgentServicePluginsExecute(r ApiListAgentServicePluginsRequest) ([]AgentServicePlugin, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []AgentServicePlugin
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PluginAPIService.ListAgentServicePlugins")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/plugin/agent-service"

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

type ApiListProvisionerPluginsRequest struct {
	ctx context.Context
	ApiService *PluginAPIService
}

func (r ApiListProvisionerPluginsRequest) Execute() ([]ProvisionerPlugin, *http.Response, error) {
	return r.ApiService.ListProvisionerPluginsExecute(r)
}

/*
ListProvisionerPlugins List provisioner plugins

List provisioner plugins

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListProvisionerPluginsRequest
*/
func (a *PluginAPIService) ListProvisionerPlugins(ctx context.Context) ApiListProvisionerPluginsRequest {
	return ApiListProvisionerPluginsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []ProvisionerPlugin
func (a *PluginAPIService) ListProvisionerPluginsExecute(r ApiListProvisionerPluginsRequest) ([]ProvisionerPlugin, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []ProvisionerPlugin
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PluginAPIService.ListProvisionerPlugins")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/plugin/provisioner"

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

type ApiUninstallAgentServicePluginRequest struct {
	ctx context.Context
	ApiService *PluginAPIService
	agentService string
}

func (r ApiUninstallAgentServicePluginRequest) Execute() (*http.Response, error) {
	return r.ApiService.UninstallAgentServicePluginExecute(r)
}

/*
UninstallAgentServicePlugin Uninstall an agent service plugin

Uninstall an agent service plugin

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param agentService Agent Service to uninstall
 @return ApiUninstallAgentServicePluginRequest
*/
func (a *PluginAPIService) UninstallAgentServicePlugin(ctx context.Context, agentService string) ApiUninstallAgentServicePluginRequest {
	return ApiUninstallAgentServicePluginRequest{
		ApiService: a,
		ctx: ctx,
		agentService: agentService,
	}
}

// Execute executes the request
func (a *PluginAPIService) UninstallAgentServicePluginExecute(r ApiUninstallAgentServicePluginRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PluginAPIService.UninstallAgentServicePlugin")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/plugin/agent-service/uninstall"
	localVarPath = strings.Replace(localVarPath, "{"+"agent-service"+"}", url.PathEscape(parameterValueToString(r.agentService, "agentService")), -1)

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
	localVarHTTPHeaderAccepts := []string{}

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
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiUninstallProvisionerPluginRequest struct {
	ctx context.Context
	ApiService *PluginAPIService
	provisioner string
}

func (r ApiUninstallProvisionerPluginRequest) Execute() (*http.Response, error) {
	return r.ApiService.UninstallProvisionerPluginExecute(r)
}

/*
UninstallProvisionerPlugin Uninstall a provisioner plugin

Uninstall a provisioner plugin

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param provisioner Provisioner to uninstall
 @return ApiUninstallProvisionerPluginRequest
*/
func (a *PluginAPIService) UninstallProvisionerPlugin(ctx context.Context, provisioner string) ApiUninstallProvisionerPluginRequest {
	return ApiUninstallProvisionerPluginRequest{
		ApiService: a,
		ctx: ctx,
		provisioner: provisioner,
	}
}

// Execute executes the request
func (a *PluginAPIService) UninstallProvisionerPluginExecute(r ApiUninstallProvisionerPluginRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PluginAPIService.UninstallProvisionerPlugin")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/plugin/provisioner/{provisioner}/uninstall"
	localVarPath = strings.Replace(localVarPath, "{"+"provisioner"+"}", url.PathEscape(parameterValueToString(r.provisioner, "provisioner")), -1)

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
	localVarHTTPHeaderAccepts := []string{}

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
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}