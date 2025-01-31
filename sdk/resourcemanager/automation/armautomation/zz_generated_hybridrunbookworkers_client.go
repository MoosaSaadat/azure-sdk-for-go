//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armautomation

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

// HybridRunbookWorkersClient contains the methods for the HybridRunbookWorkers group.
// Don't use this type directly, use NewHybridRunbookWorkersClient() instead.
type HybridRunbookWorkersClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewHybridRunbookWorkersClient creates a new instance of HybridRunbookWorkersClient with the specified values.
// subscriptionID - Gets subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID
// forms part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewHybridRunbookWorkersClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *HybridRunbookWorkersClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Endpoint) == 0 {
		cp.Endpoint = arm.AzurePublicCloud
	}
	client := &HybridRunbookWorkersClient{
		subscriptionID: subscriptionID,
		host:           string(cp.Endpoint),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, &cp),
	}
	return client
}

// Create - Create a hybrid runbook worker.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - Name of an Azure Resource group.
// automationAccountName - The name of the automation account.
// hybridRunbookWorkerGroupName - The hybrid runbook worker group name
// hybridRunbookWorkerID - The hybrid runbook worker id
// hybridRunbookWorkerCreationParameters - The create or update parameters for hybrid runbook worker.
// options - HybridRunbookWorkersClientCreateOptions contains the optional parameters for the HybridRunbookWorkersClient.Create
// method.
func (client *HybridRunbookWorkersClient) Create(ctx context.Context, resourceGroupName string, automationAccountName string, hybridRunbookWorkerGroupName string, hybridRunbookWorkerID string, hybridRunbookWorkerCreationParameters HybridRunbookWorkerCreateParameters, options *HybridRunbookWorkersClientCreateOptions) (HybridRunbookWorkersClientCreateResponse, error) {
	req, err := client.createCreateRequest(ctx, resourceGroupName, automationAccountName, hybridRunbookWorkerGroupName, hybridRunbookWorkerID, hybridRunbookWorkerCreationParameters, options)
	if err != nil {
		return HybridRunbookWorkersClientCreateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return HybridRunbookWorkersClientCreateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return HybridRunbookWorkersClientCreateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createHandleResponse(resp)
}

// createCreateRequest creates the Create request.
func (client *HybridRunbookWorkersClient) createCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, hybridRunbookWorkerGroupName string, hybridRunbookWorkerID string, hybridRunbookWorkerCreationParameters HybridRunbookWorkerCreateParameters, options *HybridRunbookWorkersClientCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/hybridRunbookWorkerGroups/{hybridRunbookWorkerGroupName}/hybridRunbookWorkers/{hybridRunbookWorkerId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	if hybridRunbookWorkerGroupName == "" {
		return nil, errors.New("parameter hybridRunbookWorkerGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hybridRunbookWorkerGroupName}", url.PathEscape(hybridRunbookWorkerGroupName))
	if hybridRunbookWorkerID == "" {
		return nil, errors.New("parameter hybridRunbookWorkerID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hybridRunbookWorkerId}", url.PathEscape(hybridRunbookWorkerID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-22")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, hybridRunbookWorkerCreationParameters)
}

// createHandleResponse handles the Create response.
func (client *HybridRunbookWorkersClient) createHandleResponse(resp *http.Response) (HybridRunbookWorkersClientCreateResponse, error) {
	result := HybridRunbookWorkersClientCreateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.HybridRunbookWorker); err != nil {
		return HybridRunbookWorkersClientCreateResponse{}, err
	}
	return result, nil
}

// Delete - Delete a hybrid runbook worker.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - Name of an Azure Resource group.
// automationAccountName - The name of the automation account.
// hybridRunbookWorkerGroupName - The hybrid runbook worker group name
// hybridRunbookWorkerID - The hybrid runbook worker id
// options - HybridRunbookWorkersClientDeleteOptions contains the optional parameters for the HybridRunbookWorkersClient.Delete
// method.
func (client *HybridRunbookWorkersClient) Delete(ctx context.Context, resourceGroupName string, automationAccountName string, hybridRunbookWorkerGroupName string, hybridRunbookWorkerID string, options *HybridRunbookWorkersClientDeleteOptions) (HybridRunbookWorkersClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, automationAccountName, hybridRunbookWorkerGroupName, hybridRunbookWorkerID, options)
	if err != nil {
		return HybridRunbookWorkersClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return HybridRunbookWorkersClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return HybridRunbookWorkersClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return HybridRunbookWorkersClientDeleteResponse{RawResponse: resp}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *HybridRunbookWorkersClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, hybridRunbookWorkerGroupName string, hybridRunbookWorkerID string, options *HybridRunbookWorkersClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/hybridRunbookWorkerGroups/{hybridRunbookWorkerGroupName}/hybridRunbookWorkers/{hybridRunbookWorkerId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	if hybridRunbookWorkerGroupName == "" {
		return nil, errors.New("parameter hybridRunbookWorkerGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hybridRunbookWorkerGroupName}", url.PathEscape(hybridRunbookWorkerGroupName))
	if hybridRunbookWorkerID == "" {
		return nil, errors.New("parameter hybridRunbookWorkerID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hybridRunbookWorkerId}", url.PathEscape(hybridRunbookWorkerID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-22")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Retrieve a hybrid runbook worker.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - Name of an Azure Resource group.
// automationAccountName - The name of the automation account.
// hybridRunbookWorkerGroupName - The hybrid runbook worker group name
// hybridRunbookWorkerID - The hybrid runbook worker id
// options - HybridRunbookWorkersClientGetOptions contains the optional parameters for the HybridRunbookWorkersClient.Get
// method.
func (client *HybridRunbookWorkersClient) Get(ctx context.Context, resourceGroupName string, automationAccountName string, hybridRunbookWorkerGroupName string, hybridRunbookWorkerID string, options *HybridRunbookWorkersClientGetOptions) (HybridRunbookWorkersClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, automationAccountName, hybridRunbookWorkerGroupName, hybridRunbookWorkerID, options)
	if err != nil {
		return HybridRunbookWorkersClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return HybridRunbookWorkersClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return HybridRunbookWorkersClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *HybridRunbookWorkersClient) getCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, hybridRunbookWorkerGroupName string, hybridRunbookWorkerID string, options *HybridRunbookWorkersClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/hybridRunbookWorkerGroups/{hybridRunbookWorkerGroupName}/hybridRunbookWorkers/{hybridRunbookWorkerId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	if hybridRunbookWorkerGroupName == "" {
		return nil, errors.New("parameter hybridRunbookWorkerGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hybridRunbookWorkerGroupName}", url.PathEscape(hybridRunbookWorkerGroupName))
	if hybridRunbookWorkerID == "" {
		return nil, errors.New("parameter hybridRunbookWorkerID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hybridRunbookWorkerId}", url.PathEscape(hybridRunbookWorkerID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-22")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *HybridRunbookWorkersClient) getHandleResponse(resp *http.Response) (HybridRunbookWorkersClientGetResponse, error) {
	result := HybridRunbookWorkersClientGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.HybridRunbookWorker); err != nil {
		return HybridRunbookWorkersClientGetResponse{}, err
	}
	return result, nil
}

// ListByHybridRunbookWorkerGroup - Retrieve a list of hybrid runbook workers.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - Name of an Azure Resource group.
// automationAccountName - The name of the automation account.
// hybridRunbookWorkerGroupName - The hybrid runbook worker group name
// options - HybridRunbookWorkersClientListByHybridRunbookWorkerGroupOptions contains the optional parameters for the HybridRunbookWorkersClient.ListByHybridRunbookWorkerGroup
// method.
func (client *HybridRunbookWorkersClient) ListByHybridRunbookWorkerGroup(resourceGroupName string, automationAccountName string, hybridRunbookWorkerGroupName string, options *HybridRunbookWorkersClientListByHybridRunbookWorkerGroupOptions) *HybridRunbookWorkersClientListByHybridRunbookWorkerGroupPager {
	return &HybridRunbookWorkersClientListByHybridRunbookWorkerGroupPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByHybridRunbookWorkerGroupCreateRequest(ctx, resourceGroupName, automationAccountName, hybridRunbookWorkerGroupName, options)
		},
		advancer: func(ctx context.Context, resp HybridRunbookWorkersClientListByHybridRunbookWorkerGroupResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.HybridRunbookWorkersListResult.NextLink)
		},
	}
}

// listByHybridRunbookWorkerGroupCreateRequest creates the ListByHybridRunbookWorkerGroup request.
func (client *HybridRunbookWorkersClient) listByHybridRunbookWorkerGroupCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, hybridRunbookWorkerGroupName string, options *HybridRunbookWorkersClientListByHybridRunbookWorkerGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/hybridRunbookWorkerGroups/{hybridRunbookWorkerGroupName}/hybridRunbookWorkers"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	if hybridRunbookWorkerGroupName == "" {
		return nil, errors.New("parameter hybridRunbookWorkerGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hybridRunbookWorkerGroupName}", url.PathEscape(hybridRunbookWorkerGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	reqQP.Set("api-version", "2021-06-22")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByHybridRunbookWorkerGroupHandleResponse handles the ListByHybridRunbookWorkerGroup response.
func (client *HybridRunbookWorkersClient) listByHybridRunbookWorkerGroupHandleResponse(resp *http.Response) (HybridRunbookWorkersClientListByHybridRunbookWorkerGroupResponse, error) {
	result := HybridRunbookWorkersClientListByHybridRunbookWorkerGroupResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.HybridRunbookWorkersListResult); err != nil {
		return HybridRunbookWorkersClientListByHybridRunbookWorkerGroupResponse{}, err
	}
	return result, nil
}

// Move - Move a hybrid worker to a different group.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - Name of an Azure Resource group.
// automationAccountName - The name of the automation account.
// hybridRunbookWorkerGroupName - The hybrid runbook worker group name
// hybridRunbookWorkerID - The hybrid runbook worker id
// hybridRunbookWorkerMoveParameters - The hybrid runbook worker move parameters
// options - HybridRunbookWorkersClientMoveOptions contains the optional parameters for the HybridRunbookWorkersClient.Move
// method.
func (client *HybridRunbookWorkersClient) Move(ctx context.Context, resourceGroupName string, automationAccountName string, hybridRunbookWorkerGroupName string, hybridRunbookWorkerID string, hybridRunbookWorkerMoveParameters HybridRunbookWorkerMoveParameters, options *HybridRunbookWorkersClientMoveOptions) (HybridRunbookWorkersClientMoveResponse, error) {
	req, err := client.moveCreateRequest(ctx, resourceGroupName, automationAccountName, hybridRunbookWorkerGroupName, hybridRunbookWorkerID, hybridRunbookWorkerMoveParameters, options)
	if err != nil {
		return HybridRunbookWorkersClientMoveResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return HybridRunbookWorkersClientMoveResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return HybridRunbookWorkersClientMoveResponse{}, runtime.NewResponseError(resp)
	}
	return HybridRunbookWorkersClientMoveResponse{RawResponse: resp}, nil
}

// moveCreateRequest creates the Move request.
func (client *HybridRunbookWorkersClient) moveCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, hybridRunbookWorkerGroupName string, hybridRunbookWorkerID string, hybridRunbookWorkerMoveParameters HybridRunbookWorkerMoveParameters, options *HybridRunbookWorkersClientMoveOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/hybridRunbookWorkerGroups/{hybridRunbookWorkerGroupName}/hybridRunbookWorkers/{hybridRunbookWorkerId}/move"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	if hybridRunbookWorkerGroupName == "" {
		return nil, errors.New("parameter hybridRunbookWorkerGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hybridRunbookWorkerGroupName}", url.PathEscape(hybridRunbookWorkerGroupName))
	if hybridRunbookWorkerID == "" {
		return nil, errors.New("parameter hybridRunbookWorkerID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hybridRunbookWorkerId}", url.PathEscape(hybridRunbookWorkerID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-22")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, hybridRunbookWorkerMoveParameters)
}
