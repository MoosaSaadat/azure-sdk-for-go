//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armbaremetalinfrastructure

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"reflect"
)

// AzureBareMetalInstancesClientListByResourceGroupPager provides operations for iterating over paged responses.
type AzureBareMetalInstancesClientListByResourceGroupPager struct {
	client    *AzureBareMetalInstancesClient
	current   AzureBareMetalInstancesClientListByResourceGroupResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, AzureBareMetalInstancesClientListByResourceGroupResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *AzureBareMetalInstancesClientListByResourceGroupPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *AzureBareMetalInstancesClientListByResourceGroupPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.AzureBareMetalInstancesListResult.NextLink == nil || len(*p.current.AzureBareMetalInstancesListResult.NextLink) == 0 {
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

// PageResponse returns the current AzureBareMetalInstancesClientListByResourceGroupResponse page.
func (p *AzureBareMetalInstancesClientListByResourceGroupPager) PageResponse() AzureBareMetalInstancesClientListByResourceGroupResponse {
	return p.current
}

// AzureBareMetalInstancesClientListBySubscriptionPager provides operations for iterating over paged responses.
type AzureBareMetalInstancesClientListBySubscriptionPager struct {
	client    *AzureBareMetalInstancesClient
	current   AzureBareMetalInstancesClientListBySubscriptionResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, AzureBareMetalInstancesClientListBySubscriptionResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *AzureBareMetalInstancesClientListBySubscriptionPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *AzureBareMetalInstancesClientListBySubscriptionPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.AzureBareMetalInstancesListResult.NextLink == nil || len(*p.current.AzureBareMetalInstancesListResult.NextLink) == 0 {
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

// PageResponse returns the current AzureBareMetalInstancesClientListBySubscriptionResponse page.
func (p *AzureBareMetalInstancesClientListBySubscriptionPager) PageResponse() AzureBareMetalInstancesClientListBySubscriptionResponse {
	return p.current
}
