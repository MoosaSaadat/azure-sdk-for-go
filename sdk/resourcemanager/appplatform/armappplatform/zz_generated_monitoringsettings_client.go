//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armappplatform

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// MonitoringSettingsClient contains the methods for the MonitoringSettings group.
// Don't use this type directly, use NewMonitoringSettingsClient() instead.
type MonitoringSettingsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewMonitoringSettingsClient creates a new instance of MonitoringSettingsClient with the specified values.
// subscriptionID - Gets subscription ID which uniquely identify the Microsoft Azure subscription. The subscription ID forms
// part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewMonitoringSettingsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *MonitoringSettingsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Endpoint) == 0 {
		cp.Endpoint = arm.AzurePublicCloud
	}
	client := &MonitoringSettingsClient{
		subscriptionID: subscriptionID,
		host:           string(cp.Endpoint),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, &cp),
	}
	return client
}

// Get - Get the Monitoring Setting and its properties.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// serviceName - The name of the Service resource.
// options - MonitoringSettingsClientGetOptions contains the optional parameters for the MonitoringSettingsClient.Get method.
func (client *MonitoringSettingsClient) Get(ctx context.Context, resourceGroupName string, serviceName string, options *MonitoringSettingsClientGetOptions) (MonitoringSettingsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, serviceName, options)
	if err != nil {
		return MonitoringSettingsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return MonitoringSettingsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return MonitoringSettingsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *MonitoringSettingsClient) getCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, options *MonitoringSettingsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/monitoringSettings/default"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *MonitoringSettingsClient) getHandleResponse(resp *http.Response) (MonitoringSettingsClientGetResponse, error) {
	result := MonitoringSettingsClientGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.MonitoringSettingResource); err != nil {
		return MonitoringSettingsClientGetResponse{}, err
	}
	return result, nil
}

// BeginUpdatePatch - Update the Monitoring Setting.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// serviceName - The name of the Service resource.
// monitoringSettingResource - Parameters for the update operation
// options - MonitoringSettingsClientBeginUpdatePatchOptions contains the optional parameters for the MonitoringSettingsClient.BeginUpdatePatch
// method.
func (client *MonitoringSettingsClient) BeginUpdatePatch(ctx context.Context, resourceGroupName string, serviceName string, monitoringSettingResource MonitoringSettingResource, options *MonitoringSettingsClientBeginUpdatePatchOptions) (MonitoringSettingsClientUpdatePatchPollerResponse, error) {
	resp, err := client.updatePatch(ctx, resourceGroupName, serviceName, monitoringSettingResource, options)
	if err != nil {
		return MonitoringSettingsClientUpdatePatchPollerResponse{}, err
	}
	result := MonitoringSettingsClientUpdatePatchPollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("MonitoringSettingsClient.UpdatePatch", "azure-async-operation", resp, client.pl)
	if err != nil {
		return MonitoringSettingsClientUpdatePatchPollerResponse{}, err
	}
	result.Poller = &MonitoringSettingsClientUpdatePatchPoller{
		pt: pt,
	}
	return result, nil
}

// UpdatePatch - Update the Monitoring Setting.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *MonitoringSettingsClient) updatePatch(ctx context.Context, resourceGroupName string, serviceName string, monitoringSettingResource MonitoringSettingResource, options *MonitoringSettingsClientBeginUpdatePatchOptions) (*http.Response, error) {
	req, err := client.updatePatchCreateRequest(ctx, resourceGroupName, serviceName, monitoringSettingResource, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// updatePatchCreateRequest creates the UpdatePatch request.
func (client *MonitoringSettingsClient) updatePatchCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, monitoringSettingResource MonitoringSettingResource, options *MonitoringSettingsClientBeginUpdatePatchOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/monitoringSettings/default"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, monitoringSettingResource)
}

// BeginUpdatePut - Update the Monitoring Setting.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// serviceName - The name of the Service resource.
// monitoringSettingResource - Parameters for the update operation
// options - MonitoringSettingsClientBeginUpdatePutOptions contains the optional parameters for the MonitoringSettingsClient.BeginUpdatePut
// method.
func (client *MonitoringSettingsClient) BeginUpdatePut(ctx context.Context, resourceGroupName string, serviceName string, monitoringSettingResource MonitoringSettingResource, options *MonitoringSettingsClientBeginUpdatePutOptions) (MonitoringSettingsClientUpdatePutPollerResponse, error) {
	resp, err := client.updatePut(ctx, resourceGroupName, serviceName, monitoringSettingResource, options)
	if err != nil {
		return MonitoringSettingsClientUpdatePutPollerResponse{}, err
	}
	result := MonitoringSettingsClientUpdatePutPollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("MonitoringSettingsClient.UpdatePut", "azure-async-operation", resp, client.pl)
	if err != nil {
		return MonitoringSettingsClientUpdatePutPollerResponse{}, err
	}
	result.Poller = &MonitoringSettingsClientUpdatePutPoller{
		pt: pt,
	}
	return result, nil
}

// UpdatePut - Update the Monitoring Setting.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *MonitoringSettingsClient) updatePut(ctx context.Context, resourceGroupName string, serviceName string, monitoringSettingResource MonitoringSettingResource, options *MonitoringSettingsClientBeginUpdatePutOptions) (*http.Response, error) {
	req, err := client.updatePutCreateRequest(ctx, resourceGroupName, serviceName, monitoringSettingResource, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// updatePutCreateRequest creates the UpdatePut request.
func (client *MonitoringSettingsClient) updatePutCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, monitoringSettingResource MonitoringSettingResource, options *MonitoringSettingsClientBeginUpdatePutOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/monitoringSettings/default"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, monitoringSettingResource)
}
