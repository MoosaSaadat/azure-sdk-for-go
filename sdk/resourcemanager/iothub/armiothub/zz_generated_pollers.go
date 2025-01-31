//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armiothub

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
)

// ClientManualFailoverPoller provides polling facilities until the operation reaches a terminal state.
type ClientManualFailoverPoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *ClientManualFailoverPoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *ClientManualFailoverPoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final ClientManualFailoverResponse will be returned.
func (p *ClientManualFailoverPoller) FinalResponse(ctx context.Context) (ClientManualFailoverResponse, error) {
	respType := ClientManualFailoverResponse{}
	resp, err := p.pt.FinalResponse(ctx, nil)
	if err != nil {
		return ClientManualFailoverResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *ClientManualFailoverPoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// PrivateEndpointConnectionsClientDeletePoller provides polling facilities until the operation reaches a terminal state.
type PrivateEndpointConnectionsClientDeletePoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *PrivateEndpointConnectionsClientDeletePoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *PrivateEndpointConnectionsClientDeletePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final PrivateEndpointConnectionsClientDeleteResponse will be returned.
func (p *PrivateEndpointConnectionsClientDeletePoller) FinalResponse(ctx context.Context) (PrivateEndpointConnectionsClientDeleteResponse, error) {
	respType := PrivateEndpointConnectionsClientDeleteResponse{}
	resp, err := p.pt.FinalResponse(ctx, &respType.PrivateEndpointConnection)
	if err != nil {
		return PrivateEndpointConnectionsClientDeleteResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *PrivateEndpointConnectionsClientDeletePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// PrivateEndpointConnectionsClientUpdatePoller provides polling facilities until the operation reaches a terminal state.
type PrivateEndpointConnectionsClientUpdatePoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *PrivateEndpointConnectionsClientUpdatePoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *PrivateEndpointConnectionsClientUpdatePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final PrivateEndpointConnectionsClientUpdateResponse will be returned.
func (p *PrivateEndpointConnectionsClientUpdatePoller) FinalResponse(ctx context.Context) (PrivateEndpointConnectionsClientUpdateResponse, error) {
	respType := PrivateEndpointConnectionsClientUpdateResponse{}
	resp, err := p.pt.FinalResponse(ctx, &respType.PrivateEndpointConnection)
	if err != nil {
		return PrivateEndpointConnectionsClientUpdateResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *PrivateEndpointConnectionsClientUpdatePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// ResourceClientCreateOrUpdatePoller provides polling facilities until the operation reaches a terminal state.
type ResourceClientCreateOrUpdatePoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *ResourceClientCreateOrUpdatePoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *ResourceClientCreateOrUpdatePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final ResourceClientCreateOrUpdateResponse will be returned.
func (p *ResourceClientCreateOrUpdatePoller) FinalResponse(ctx context.Context) (ResourceClientCreateOrUpdateResponse, error) {
	respType := ResourceClientCreateOrUpdateResponse{}
	resp, err := p.pt.FinalResponse(ctx, &respType.Description)
	if err != nil {
		return ResourceClientCreateOrUpdateResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *ResourceClientCreateOrUpdatePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// ResourceClientDeletePoller provides polling facilities until the operation reaches a terminal state.
type ResourceClientDeletePoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *ResourceClientDeletePoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *ResourceClientDeletePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final ResourceClientDeleteResponse will be returned.
func (p *ResourceClientDeletePoller) FinalResponse(ctx context.Context) (ResourceClientDeleteResponse, error) {
	respType := ResourceClientDeleteResponse{}
	resp, err := p.pt.FinalResponse(ctx, &respType.Description)
	if err != nil {
		return ResourceClientDeleteResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *ResourceClientDeletePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// ResourceClientUpdatePoller provides polling facilities until the operation reaches a terminal state.
type ResourceClientUpdatePoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *ResourceClientUpdatePoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *ResourceClientUpdatePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final ResourceClientUpdateResponse will be returned.
func (p *ResourceClientUpdatePoller) FinalResponse(ctx context.Context) (ResourceClientUpdateResponse, error) {
	respType := ResourceClientUpdateResponse{}
	resp, err := p.pt.FinalResponse(ctx, &respType.Description)
	if err != nil {
		return ResourceClientUpdateResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *ResourceClientUpdatePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}
