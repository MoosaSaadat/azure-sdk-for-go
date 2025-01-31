//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmanagedservices

import (
	"context"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"net/http"
	"time"
)

// MarketplaceRegistrationDefinitionsClientGetResponse contains the response from method MarketplaceRegistrationDefinitionsClient.Get.
type MarketplaceRegistrationDefinitionsClientGetResponse struct {
	MarketplaceRegistrationDefinitionsClientGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// MarketplaceRegistrationDefinitionsClientGetResult contains the result from method MarketplaceRegistrationDefinitionsClient.Get.
type MarketplaceRegistrationDefinitionsClientGetResult struct {
	MarketplaceRegistrationDefinition
}

// MarketplaceRegistrationDefinitionsClientListResponse contains the response from method MarketplaceRegistrationDefinitionsClient.List.
type MarketplaceRegistrationDefinitionsClientListResponse struct {
	MarketplaceRegistrationDefinitionsClientListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// MarketplaceRegistrationDefinitionsClientListResult contains the result from method MarketplaceRegistrationDefinitionsClient.List.
type MarketplaceRegistrationDefinitionsClientListResult struct {
	MarketplaceRegistrationDefinitionList
}

// MarketplaceRegistrationDefinitionsWithoutScopeClientGetResponse contains the response from method MarketplaceRegistrationDefinitionsWithoutScopeClient.Get.
type MarketplaceRegistrationDefinitionsWithoutScopeClientGetResponse struct {
	MarketplaceRegistrationDefinitionsWithoutScopeClientGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// MarketplaceRegistrationDefinitionsWithoutScopeClientGetResult contains the result from method MarketplaceRegistrationDefinitionsWithoutScopeClient.Get.
type MarketplaceRegistrationDefinitionsWithoutScopeClientGetResult struct {
	MarketplaceRegistrationDefinition
}

// MarketplaceRegistrationDefinitionsWithoutScopeClientListResponse contains the response from method MarketplaceRegistrationDefinitionsWithoutScopeClient.List.
type MarketplaceRegistrationDefinitionsWithoutScopeClientListResponse struct {
	MarketplaceRegistrationDefinitionsWithoutScopeClientListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// MarketplaceRegistrationDefinitionsWithoutScopeClientListResult contains the result from method MarketplaceRegistrationDefinitionsWithoutScopeClient.List.
type MarketplaceRegistrationDefinitionsWithoutScopeClientListResult struct {
	MarketplaceRegistrationDefinitionList
}

// OperationsClientListResponse contains the response from method OperationsClient.List.
type OperationsClientListResponse struct {
	OperationsClientListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// OperationsClientListResult contains the result from method OperationsClient.List.
type OperationsClientListResult struct {
	OperationList
}

// RegistrationAssignmentsClientCreateOrUpdatePollerResponse contains the response from method RegistrationAssignmentsClient.CreateOrUpdate.
type RegistrationAssignmentsClientCreateOrUpdatePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *RegistrationAssignmentsClientCreateOrUpdatePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l RegistrationAssignmentsClientCreateOrUpdatePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (RegistrationAssignmentsClientCreateOrUpdateResponse, error) {
	respType := RegistrationAssignmentsClientCreateOrUpdateResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, &respType.RegistrationAssignment)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a RegistrationAssignmentsClientCreateOrUpdatePollerResponse from the provided client and resume token.
func (l *RegistrationAssignmentsClientCreateOrUpdatePollerResponse) Resume(ctx context.Context, client *RegistrationAssignmentsClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("RegistrationAssignmentsClient.CreateOrUpdate", token, client.pl)
	if err != nil {
		return err
	}
	poller := &RegistrationAssignmentsClientCreateOrUpdatePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// RegistrationAssignmentsClientCreateOrUpdateResponse contains the response from method RegistrationAssignmentsClient.CreateOrUpdate.
type RegistrationAssignmentsClientCreateOrUpdateResponse struct {
	RegistrationAssignmentsClientCreateOrUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// RegistrationAssignmentsClientCreateOrUpdateResult contains the result from method RegistrationAssignmentsClient.CreateOrUpdate.
type RegistrationAssignmentsClientCreateOrUpdateResult struct {
	RegistrationAssignment
}

// RegistrationAssignmentsClientDeletePollerResponse contains the response from method RegistrationAssignmentsClient.Delete.
type RegistrationAssignmentsClientDeletePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *RegistrationAssignmentsClientDeletePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l RegistrationAssignmentsClientDeletePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (RegistrationAssignmentsClientDeleteResponse, error) {
	respType := RegistrationAssignmentsClientDeleteResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, nil)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a RegistrationAssignmentsClientDeletePollerResponse from the provided client and resume token.
func (l *RegistrationAssignmentsClientDeletePollerResponse) Resume(ctx context.Context, client *RegistrationAssignmentsClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("RegistrationAssignmentsClient.Delete", token, client.pl)
	if err != nil {
		return err
	}
	poller := &RegistrationAssignmentsClientDeletePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// RegistrationAssignmentsClientDeleteResponse contains the response from method RegistrationAssignmentsClient.Delete.
type RegistrationAssignmentsClientDeleteResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// RegistrationAssignmentsClientGetResponse contains the response from method RegistrationAssignmentsClient.Get.
type RegistrationAssignmentsClientGetResponse struct {
	RegistrationAssignmentsClientGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// RegistrationAssignmentsClientGetResult contains the result from method RegistrationAssignmentsClient.Get.
type RegistrationAssignmentsClientGetResult struct {
	RegistrationAssignment
}

// RegistrationAssignmentsClientListResponse contains the response from method RegistrationAssignmentsClient.List.
type RegistrationAssignmentsClientListResponse struct {
	RegistrationAssignmentsClientListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// RegistrationAssignmentsClientListResult contains the result from method RegistrationAssignmentsClient.List.
type RegistrationAssignmentsClientListResult struct {
	RegistrationAssignmentList
}

// RegistrationDefinitionsClientCreateOrUpdatePollerResponse contains the response from method RegistrationDefinitionsClient.CreateOrUpdate.
type RegistrationDefinitionsClientCreateOrUpdatePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *RegistrationDefinitionsClientCreateOrUpdatePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l RegistrationDefinitionsClientCreateOrUpdatePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (RegistrationDefinitionsClientCreateOrUpdateResponse, error) {
	respType := RegistrationDefinitionsClientCreateOrUpdateResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, &respType.RegistrationDefinition)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a RegistrationDefinitionsClientCreateOrUpdatePollerResponse from the provided client and resume token.
func (l *RegistrationDefinitionsClientCreateOrUpdatePollerResponse) Resume(ctx context.Context, client *RegistrationDefinitionsClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("RegistrationDefinitionsClient.CreateOrUpdate", token, client.pl)
	if err != nil {
		return err
	}
	poller := &RegistrationDefinitionsClientCreateOrUpdatePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// RegistrationDefinitionsClientCreateOrUpdateResponse contains the response from method RegistrationDefinitionsClient.CreateOrUpdate.
type RegistrationDefinitionsClientCreateOrUpdateResponse struct {
	RegistrationDefinitionsClientCreateOrUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// RegistrationDefinitionsClientCreateOrUpdateResult contains the result from method RegistrationDefinitionsClient.CreateOrUpdate.
type RegistrationDefinitionsClientCreateOrUpdateResult struct {
	RegistrationDefinition
}

// RegistrationDefinitionsClientDeleteResponse contains the response from method RegistrationDefinitionsClient.Delete.
type RegistrationDefinitionsClientDeleteResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// RegistrationDefinitionsClientGetResponse contains the response from method RegistrationDefinitionsClient.Get.
type RegistrationDefinitionsClientGetResponse struct {
	RegistrationDefinitionsClientGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// RegistrationDefinitionsClientGetResult contains the result from method RegistrationDefinitionsClient.Get.
type RegistrationDefinitionsClientGetResult struct {
	RegistrationDefinition
}

// RegistrationDefinitionsClientListResponse contains the response from method RegistrationDefinitionsClient.List.
type RegistrationDefinitionsClientListResponse struct {
	RegistrationDefinitionsClientListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// RegistrationDefinitionsClientListResult contains the result from method RegistrationDefinitionsClient.List.
type RegistrationDefinitionsClientListResult struct {
	RegistrationDefinitionList
}
