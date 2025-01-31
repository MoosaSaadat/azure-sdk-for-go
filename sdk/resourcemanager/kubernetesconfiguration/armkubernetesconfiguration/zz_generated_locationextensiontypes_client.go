//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armkubernetesconfiguration

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

// LocationExtensionTypesClient contains the methods for the LocationExtensionTypes group.
// Don't use this type directly, use NewLocationExtensionTypesClient() instead.
type LocationExtensionTypesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewLocationExtensionTypesClient creates a new instance of LocationExtensionTypesClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewLocationExtensionTypesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *LocationExtensionTypesClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &LocationExtensionTypesClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// List - List all Extension Types
// If the operation fails it returns an *azcore.ResponseError type.
// location - extension location
// options - LocationExtensionTypesClientListOptions contains the optional parameters for the LocationExtensionTypesClient.List
// method.
func (client *LocationExtensionTypesClient) List(location string, options *LocationExtensionTypesClientListOptions) *LocationExtensionTypesClientListPager {
	return &LocationExtensionTypesClientListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, location, options)
		},
		advancer: func(ctx context.Context, resp LocationExtensionTypesClientListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ExtensionTypeList.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *LocationExtensionTypesClient) listCreateRequest(ctx context.Context, location string, options *LocationExtensionTypesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.KubernetesConfiguration/locations/{location}/extensionTypes"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
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

// listHandleResponse handles the List response.
func (client *LocationExtensionTypesClient) listHandleResponse(resp *http.Response) (LocationExtensionTypesClientListResponse, error) {
	result := LocationExtensionTypesClientListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ExtensionTypeList); err != nil {
		return LocationExtensionTypesClientListResponse{}, err
	}
	return result, nil
}
