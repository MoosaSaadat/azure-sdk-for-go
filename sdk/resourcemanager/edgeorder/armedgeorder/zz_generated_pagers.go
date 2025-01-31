//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armedgeorder

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"reflect"
)

// ManagementClientListAddressesAtResourceGroupLevelPager provides operations for iterating over paged responses.
type ManagementClientListAddressesAtResourceGroupLevelPager struct {
	client    *ManagementClient
	current   ManagementClientListAddressesAtResourceGroupLevelResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, ManagementClientListAddressesAtResourceGroupLevelResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *ManagementClientListAddressesAtResourceGroupLevelPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *ManagementClientListAddressesAtResourceGroupLevelPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.AddressResourceList.NextLink == nil || len(*p.current.AddressResourceList.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = runtime.NewResponseError(resp)
		return false
	}
	result, err := p.client.listAddressesAtResourceGroupLevelHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current ManagementClientListAddressesAtResourceGroupLevelResponse page.
func (p *ManagementClientListAddressesAtResourceGroupLevelPager) PageResponse() ManagementClientListAddressesAtResourceGroupLevelResponse {
	return p.current
}

// ManagementClientListAddressesAtSubscriptionLevelPager provides operations for iterating over paged responses.
type ManagementClientListAddressesAtSubscriptionLevelPager struct {
	client    *ManagementClient
	current   ManagementClientListAddressesAtSubscriptionLevelResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, ManagementClientListAddressesAtSubscriptionLevelResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *ManagementClientListAddressesAtSubscriptionLevelPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *ManagementClientListAddressesAtSubscriptionLevelPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.AddressResourceList.NextLink == nil || len(*p.current.AddressResourceList.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = runtime.NewResponseError(resp)
		return false
	}
	result, err := p.client.listAddressesAtSubscriptionLevelHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current ManagementClientListAddressesAtSubscriptionLevelResponse page.
func (p *ManagementClientListAddressesAtSubscriptionLevelPager) PageResponse() ManagementClientListAddressesAtSubscriptionLevelResponse {
	return p.current
}

// ManagementClientListConfigurationsPager provides operations for iterating over paged responses.
type ManagementClientListConfigurationsPager struct {
	client    *ManagementClient
	current   ManagementClientListConfigurationsResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, ManagementClientListConfigurationsResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *ManagementClientListConfigurationsPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *ManagementClientListConfigurationsPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.Configurations.NextLink == nil || len(*p.current.Configurations.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = runtime.NewResponseError(resp)
		return false
	}
	result, err := p.client.listConfigurationsHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current ManagementClientListConfigurationsResponse page.
func (p *ManagementClientListConfigurationsPager) PageResponse() ManagementClientListConfigurationsResponse {
	return p.current
}

// ManagementClientListOperationsPager provides operations for iterating over paged responses.
type ManagementClientListOperationsPager struct {
	client    *ManagementClient
	current   ManagementClientListOperationsResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, ManagementClientListOperationsResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *ManagementClientListOperationsPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *ManagementClientListOperationsPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.OperationListResult.NextLink == nil || len(*p.current.OperationListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = runtime.NewResponseError(resp)
		return false
	}
	result, err := p.client.listOperationsHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current ManagementClientListOperationsResponse page.
func (p *ManagementClientListOperationsPager) PageResponse() ManagementClientListOperationsResponse {
	return p.current
}

// ManagementClientListOrderAtResourceGroupLevelPager provides operations for iterating over paged responses.
type ManagementClientListOrderAtResourceGroupLevelPager struct {
	client    *ManagementClient
	current   ManagementClientListOrderAtResourceGroupLevelResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, ManagementClientListOrderAtResourceGroupLevelResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *ManagementClientListOrderAtResourceGroupLevelPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *ManagementClientListOrderAtResourceGroupLevelPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.OrderResourceList.NextLink == nil || len(*p.current.OrderResourceList.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = runtime.NewResponseError(resp)
		return false
	}
	result, err := p.client.listOrderAtResourceGroupLevelHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current ManagementClientListOrderAtResourceGroupLevelResponse page.
func (p *ManagementClientListOrderAtResourceGroupLevelPager) PageResponse() ManagementClientListOrderAtResourceGroupLevelResponse {
	return p.current
}

// ManagementClientListOrderAtSubscriptionLevelPager provides operations for iterating over paged responses.
type ManagementClientListOrderAtSubscriptionLevelPager struct {
	client    *ManagementClient
	current   ManagementClientListOrderAtSubscriptionLevelResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, ManagementClientListOrderAtSubscriptionLevelResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *ManagementClientListOrderAtSubscriptionLevelPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *ManagementClientListOrderAtSubscriptionLevelPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.OrderResourceList.NextLink == nil || len(*p.current.OrderResourceList.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = runtime.NewResponseError(resp)
		return false
	}
	result, err := p.client.listOrderAtSubscriptionLevelHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current ManagementClientListOrderAtSubscriptionLevelResponse page.
func (p *ManagementClientListOrderAtSubscriptionLevelPager) PageResponse() ManagementClientListOrderAtSubscriptionLevelResponse {
	return p.current
}

// ManagementClientListOrderItemsAtResourceGroupLevelPager provides operations for iterating over paged responses.
type ManagementClientListOrderItemsAtResourceGroupLevelPager struct {
	client    *ManagementClient
	current   ManagementClientListOrderItemsAtResourceGroupLevelResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, ManagementClientListOrderItemsAtResourceGroupLevelResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *ManagementClientListOrderItemsAtResourceGroupLevelPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *ManagementClientListOrderItemsAtResourceGroupLevelPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.OrderItemResourceList.NextLink == nil || len(*p.current.OrderItemResourceList.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = runtime.NewResponseError(resp)
		return false
	}
	result, err := p.client.listOrderItemsAtResourceGroupLevelHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current ManagementClientListOrderItemsAtResourceGroupLevelResponse page.
func (p *ManagementClientListOrderItemsAtResourceGroupLevelPager) PageResponse() ManagementClientListOrderItemsAtResourceGroupLevelResponse {
	return p.current
}

// ManagementClientListOrderItemsAtSubscriptionLevelPager provides operations for iterating over paged responses.
type ManagementClientListOrderItemsAtSubscriptionLevelPager struct {
	client    *ManagementClient
	current   ManagementClientListOrderItemsAtSubscriptionLevelResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, ManagementClientListOrderItemsAtSubscriptionLevelResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *ManagementClientListOrderItemsAtSubscriptionLevelPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *ManagementClientListOrderItemsAtSubscriptionLevelPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.OrderItemResourceList.NextLink == nil || len(*p.current.OrderItemResourceList.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = runtime.NewResponseError(resp)
		return false
	}
	result, err := p.client.listOrderItemsAtSubscriptionLevelHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current ManagementClientListOrderItemsAtSubscriptionLevelResponse page.
func (p *ManagementClientListOrderItemsAtSubscriptionLevelPager) PageResponse() ManagementClientListOrderItemsAtSubscriptionLevelResponse {
	return p.current
}

// ManagementClientListProductFamiliesMetadataPager provides operations for iterating over paged responses.
type ManagementClientListProductFamiliesMetadataPager struct {
	client    *ManagementClient
	current   ManagementClientListProductFamiliesMetadataResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, ManagementClientListProductFamiliesMetadataResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *ManagementClientListProductFamiliesMetadataPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *ManagementClientListProductFamiliesMetadataPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.ProductFamiliesMetadata.NextLink == nil || len(*p.current.ProductFamiliesMetadata.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = runtime.NewResponseError(resp)
		return false
	}
	result, err := p.client.listProductFamiliesMetadataHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current ManagementClientListProductFamiliesMetadataResponse page.
func (p *ManagementClientListProductFamiliesMetadataPager) PageResponse() ManagementClientListProductFamiliesMetadataResponse {
	return p.current
}

// ManagementClientListProductFamiliesPager provides operations for iterating over paged responses.
type ManagementClientListProductFamiliesPager struct {
	client    *ManagementClient
	current   ManagementClientListProductFamiliesResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, ManagementClientListProductFamiliesResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *ManagementClientListProductFamiliesPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *ManagementClientListProductFamiliesPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.ProductFamilies.NextLink == nil || len(*p.current.ProductFamilies.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = runtime.NewResponseError(resp)
		return false
	}
	result, err := p.client.listProductFamiliesHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current ManagementClientListProductFamiliesResponse page.
func (p *ManagementClientListProductFamiliesPager) PageResponse() ManagementClientListProductFamiliesResponse {
	return p.current
}
