//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmachinelearningservices_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/machinelearningservices/armmachinelearningservices"
)

// x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2021-07-01/examples/Workspace/get.json
func ExampleWorkspacesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmachinelearningservices.NewWorkspacesClient("<subscription-id>", cred, nil)
	res, err := client.Get(ctx,
		"<resource-group-name>",
		"<workspace-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.WorkspacesClientGetResult)
}

// x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2021-07-01/examples/Workspace/create.json
func ExampleWorkspacesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmachinelearningservices.NewWorkspacesClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<workspace-name>",
		armmachinelearningservices.Workspace{
			Identity: &armmachinelearningservices.Identity{
				Type: armmachinelearningservices.ResourceIdentityTypeSystemAssignedUserAssigned.ToPtr(),
				UserAssignedIdentities: map[string]*armmachinelearningservices.UserAssignedIdentity{
					"/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/workspace-1234/providers/Microsoft.ManagedIdentity/userAssignedIdentities/testuai": {},
				},
			},
			Location: to.StringPtr("<location>"),
			Properties: &armmachinelearningservices.WorkspaceProperties{
				Description:         to.StringPtr("<description>"),
				ApplicationInsights: to.StringPtr("<application-insights>"),
				ContainerRegistry:   to.StringPtr("<container-registry>"),
				Encryption: &armmachinelearningservices.EncryptionProperty{
					Identity: &armmachinelearningservices.IdentityForCmk{
						UserAssignedIdentity: to.StringPtr("<user-assigned-identity>"),
					},
					KeyVaultProperties: &armmachinelearningservices.KeyVaultProperties{
						IdentityClientID: to.StringPtr("<identity-client-id>"),
						KeyIdentifier:    to.StringPtr("<key-identifier>"),
						KeyVaultArmID:    to.StringPtr("<key-vault-arm-id>"),
					},
					Status: armmachinelearningservices.EncryptionStatus("Enabled").ToPtr(),
				},
				FriendlyName: to.StringPtr("<friendly-name>"),
				HbiWorkspace: to.BoolPtr(false),
				KeyVault:     to.StringPtr("<key-vault>"),
				SharedPrivateLinkResources: []*armmachinelearningservices.SharedPrivateLinkResource{
					{
						Name: to.StringPtr("<name>"),
						Properties: &armmachinelearningservices.SharedPrivateLinkResourceProperty{
							GroupID:               to.StringPtr("<group-id>"),
							PrivateLinkResourceID: to.StringPtr("<private-link-resource-id>"),
							RequestMessage:        to.StringPtr("<request-message>"),
							Status:                armmachinelearningservices.PrivateEndpointServiceConnectionStatus("Approved").ToPtr(),
						},
					}},
				StorageAccount: to.StringPtr("<storage-account>"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.WorkspacesClientCreateOrUpdateResult)
}

// x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2021-07-01/examples/Workspace/delete.json
func ExampleWorkspacesClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmachinelearningservices.NewWorkspacesClient("<subscription-id>", cred, nil)
	poller, err := client.BeginDelete(ctx,
		"<resource-group-name>",
		"<workspace-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2021-07-01/examples/Workspace/update.json
func ExampleWorkspacesClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmachinelearningservices.NewWorkspacesClient("<subscription-id>", cred, nil)
	res, err := client.Update(ctx,
		"<resource-group-name>",
		"<workspace-name>",
		armmachinelearningservices.WorkspaceUpdateParameters{
			Properties: &armmachinelearningservices.WorkspacePropertiesUpdateParameters{
				Description:         to.StringPtr("<description>"),
				FriendlyName:        to.StringPtr("<friendly-name>"),
				PublicNetworkAccess: armmachinelearningservices.PublicNetworkAccess("Disabled").ToPtr(),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.WorkspacesClientUpdateResult)
}

// x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2021-07-01/examples/Workspace/listByResourceGroup.json
func ExampleWorkspacesClient_ListByResourceGroup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmachinelearningservices.NewWorkspacesClient("<subscription-id>", cred, nil)
	pager := client.ListByResourceGroup("<resource-group-name>",
		&armmachinelearningservices.WorkspacesClientListByResourceGroupOptions{Skip: nil})
	for {
		nextResult := pager.NextPage(ctx)
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		if !nextResult {
			break
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("Pager result: %#v\n", v)
		}
	}
}

// x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2021-07-01/examples/Workspace/diagnose.json
func ExampleWorkspacesClient_BeginDiagnose() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmachinelearningservices.NewWorkspacesClient("<subscription-id>", cred, nil)
	poller, err := client.BeginDiagnose(ctx,
		"<resource-group-name>",
		"<workspace-name>",
		&armmachinelearningservices.WorkspacesClientBeginDiagnoseOptions{Parameters: &armmachinelearningservices.DiagnoseWorkspaceParameters{
			Value: &armmachinelearningservices.DiagnoseRequestProperties{
				ApplicationInsights: map[string]map[string]interface{}{},
				ContainerRegistry:   map[string]map[string]interface{}{},
				DNSResolution:       map[string]map[string]interface{}{},
				KeyVault:            map[string]map[string]interface{}{},
				Nsg:                 map[string]map[string]interface{}{},
				Others:              map[string]map[string]interface{}{},
				ResourceLock:        map[string]map[string]interface{}{},
				StorageAccount:      map[string]map[string]interface{}{},
				Udr:                 map[string]map[string]interface{}{},
			},
		},
		})
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.WorkspacesClientDiagnoseResult)
}

// x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2021-07-01/examples/Workspace/listKeys.json
func ExampleWorkspacesClient_ListKeys() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmachinelearningservices.NewWorkspacesClient("<subscription-id>", cred, nil)
	res, err := client.ListKeys(ctx,
		"<resource-group-name>",
		"<workspace-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.WorkspacesClientListKeysResult)
}

// x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2021-07-01/examples/Workspace/resyncKeys.json
func ExampleWorkspacesClient_BeginResyncKeys() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmachinelearningservices.NewWorkspacesClient("<subscription-id>", cred, nil)
	poller, err := client.BeginResyncKeys(ctx,
		"<resource-group-name>",
		"<workspace-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2021-07-01/examples/Workspace/listBySubscription.json
func ExampleWorkspacesClient_ListBySubscription() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmachinelearningservices.NewWorkspacesClient("<subscription-id>", cred, nil)
	pager := client.ListBySubscription(&armmachinelearningservices.WorkspacesClientListBySubscriptionOptions{Skip: nil})
	for {
		nextResult := pager.NextPage(ctx)
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		if !nextResult {
			break
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("Pager result: %#v\n", v)
		}
	}
}

// x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2021-07-01/examples/Workspace/listNotebookAccessToken.json
func ExampleWorkspacesClient_ListNotebookAccessToken() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmachinelearningservices.NewWorkspacesClient("<subscription-id>", cred, nil)
	res, err := client.ListNotebookAccessToken(ctx,
		"<resource-group-name>",
		"<workspace-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.WorkspacesClientListNotebookAccessTokenResult)
}

// x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2021-07-01/examples/Notebook/prepare.json
func ExampleWorkspacesClient_BeginPrepareNotebook() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmachinelearningservices.NewWorkspacesClient("<subscription-id>", cred, nil)
	poller, err := client.BeginPrepareNotebook(ctx,
		"<resource-group-name>",
		"<workspace-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.WorkspacesClientPrepareNotebookResult)
}

// x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2021-07-01/examples/Workspace/listStorageAccountKeys.json
func ExampleWorkspacesClient_ListStorageAccountKeys() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmachinelearningservices.NewWorkspacesClient("<subscription-id>", cred, nil)
	res, err := client.ListStorageAccountKeys(ctx,
		"<resource-group-name>",
		"<workspace-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.WorkspacesClientListStorageAccountKeysResult)
}

// x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2021-07-01/examples/Notebook/listKeys.json
func ExampleWorkspacesClient_ListNotebookKeys() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmachinelearningservices.NewWorkspacesClient("<subscription-id>", cred, nil)
	res, err := client.ListNotebookKeys(ctx,
		"<resource-group-name>",
		"<workspace-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.WorkspacesClientListNotebookKeysResult)
}

// x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2021-07-01/examples/ExternalFQDN/get.json
func ExampleWorkspacesClient_ListOutboundNetworkDependenciesEndpoints() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmachinelearningservices.NewWorkspacesClient("<subscription-id>", cred, nil)
	res, err := client.ListOutboundNetworkDependenciesEndpoints(ctx,
		"<resource-group-name>",
		"<workspace-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.WorkspacesClientListOutboundNetworkDependenciesEndpointsResult)
}
