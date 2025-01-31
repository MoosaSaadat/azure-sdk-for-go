//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdeviceprovisioningservices

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"reflect"
)

// IotDpsResourceClientListByResourceGroupPager provides operations for iterating over paged responses.
type IotDpsResourceClientListByResourceGroupPager struct {
	client    *IotDpsResourceClient
	current   IotDpsResourceClientListByResourceGroupResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, IotDpsResourceClientListByResourceGroupResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *IotDpsResourceClientListByResourceGroupPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *IotDpsResourceClientListByResourceGroupPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.ProvisioningServiceDescriptionListResult.NextLink == nil || len(*p.current.ProvisioningServiceDescriptionListResult.NextLink) == 0 {
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
	result, err := p.client.listByResourceGroupHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current IotDpsResourceClientListByResourceGroupResponse page.
func (p *IotDpsResourceClientListByResourceGroupPager) PageResponse() IotDpsResourceClientListByResourceGroupResponse {
	return p.current
}

// IotDpsResourceClientListBySubscriptionPager provides operations for iterating over paged responses.
type IotDpsResourceClientListBySubscriptionPager struct {
	client    *IotDpsResourceClient
	current   IotDpsResourceClientListBySubscriptionResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, IotDpsResourceClientListBySubscriptionResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *IotDpsResourceClientListBySubscriptionPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *IotDpsResourceClientListBySubscriptionPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.ProvisioningServiceDescriptionListResult.NextLink == nil || len(*p.current.ProvisioningServiceDescriptionListResult.NextLink) == 0 {
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
	result, err := p.client.listBySubscriptionHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current IotDpsResourceClientListBySubscriptionResponse page.
func (p *IotDpsResourceClientListBySubscriptionPager) PageResponse() IotDpsResourceClientListBySubscriptionResponse {
	return p.current
}

// IotDpsResourceClientListKeysPager provides operations for iterating over paged responses.
type IotDpsResourceClientListKeysPager struct {
	client    *IotDpsResourceClient
	current   IotDpsResourceClientListKeysResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, IotDpsResourceClientListKeysResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *IotDpsResourceClientListKeysPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *IotDpsResourceClientListKeysPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.SharedAccessSignatureAuthorizationRuleListResult.NextLink == nil || len(*p.current.SharedAccessSignatureAuthorizationRuleListResult.NextLink) == 0 {
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
	result, err := p.client.listKeysHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current IotDpsResourceClientListKeysResponse page.
func (p *IotDpsResourceClientListKeysPager) PageResponse() IotDpsResourceClientListKeysResponse {
	return p.current
}

// IotDpsResourceClientListValidSKUsPager provides operations for iterating over paged responses.
type IotDpsResourceClientListValidSKUsPager struct {
	client    *IotDpsResourceClient
	current   IotDpsResourceClientListValidSKUsResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, IotDpsResourceClientListValidSKUsResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *IotDpsResourceClientListValidSKUsPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *IotDpsResourceClientListValidSKUsPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.IotDpsSKUDefinitionListResult.NextLink == nil || len(*p.current.IotDpsSKUDefinitionListResult.NextLink) == 0 {
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
	result, err := p.client.listValidSKUsHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current IotDpsResourceClientListValidSKUsResponse page.
func (p *IotDpsResourceClientListValidSKUsPager) PageResponse() IotDpsResourceClientListValidSKUsResponse {
	return p.current
}

// OperationsClientListPager provides operations for iterating over paged responses.
type OperationsClientListPager struct {
	client    *OperationsClient
	current   OperationsClientListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, OperationsClientListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *OperationsClientListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *OperationsClientListPager) NextPage(ctx context.Context) bool {
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
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current OperationsClientListResponse page.
func (p *OperationsClientListPager) PageResponse() OperationsClientListResponse {
	return p.current
}
