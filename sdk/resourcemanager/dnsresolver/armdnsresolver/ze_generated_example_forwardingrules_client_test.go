//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdnsresolver_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dnsresolver/armdnsresolver"
)

// x-ms-original-file: specification/dnsresolver/resource-manager/Microsoft.Network/preview/2020-04-01-preview/examples/ForwardingRule_Put.json
func ExampleForwardingRulesClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdnsresolver.NewForwardingRulesClient("<subscription-id>", cred, nil)
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<dns-forwarding-ruleset-name>",
		"<forwarding-rule-name>",
		armdnsresolver.ForwardingRule{
			Properties: &armdnsresolver.ForwardingRuleProperties{
				DomainName:          to.StringPtr("<domain-name>"),
				ForwardingRuleState: armdnsresolver.ForwardingRuleState("Enabled").ToPtr(),
				Metadata: map[string]*string{
					"additionalProp1": to.StringPtr("value1"),
				},
				TargetDNSServers: []*armdnsresolver.TargetDNSServer{
					{
						IPAddress: to.StringPtr("<ipaddress>"),
						Port:      to.Int32Ptr(53),
					},
					{
						IPAddress: to.StringPtr("<ipaddress>"),
						Port:      to.Int32Ptr(53),
					}},
			},
		},
		&armdnsresolver.ForwardingRulesClientCreateOrUpdateOptions{IfMatch: nil,
			IfNoneMatch: nil,
		})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.ForwardingRulesClientCreateOrUpdateResult)
}

// x-ms-original-file: specification/dnsresolver/resource-manager/Microsoft.Network/preview/2020-04-01-preview/examples/ForwardingRule_Patch.json
func ExampleForwardingRulesClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdnsresolver.NewForwardingRulesClient("<subscription-id>", cred, nil)
	res, err := client.Update(ctx,
		"<resource-group-name>",
		"<dns-forwarding-ruleset-name>",
		"<forwarding-rule-name>",
		armdnsresolver.ForwardingRulePatch{
			Properties: &armdnsresolver.ForwardingRulePatchProperties{
				ForwardingRuleState: armdnsresolver.ForwardingRuleState("Disabled").ToPtr(),
				Metadata: map[string]*string{
					"additionalProp2": to.StringPtr("value2"),
				},
			},
		},
		&armdnsresolver.ForwardingRulesClientUpdateOptions{IfMatch: nil})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.ForwardingRulesClientUpdateResult)
}

// x-ms-original-file: specification/dnsresolver/resource-manager/Microsoft.Network/preview/2020-04-01-preview/examples/ForwardingRule_Delete.json
func ExampleForwardingRulesClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdnsresolver.NewForwardingRulesClient("<subscription-id>", cred, nil)
	_, err = client.Delete(ctx,
		"<resource-group-name>",
		"<dns-forwarding-ruleset-name>",
		"<forwarding-rule-name>",
		&armdnsresolver.ForwardingRulesClientDeleteOptions{IfMatch: nil})
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/dnsresolver/resource-manager/Microsoft.Network/preview/2020-04-01-preview/examples/ForwardingRule_Get.json
func ExampleForwardingRulesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdnsresolver.NewForwardingRulesClient("<subscription-id>", cred, nil)
	res, err := client.Get(ctx,
		"<resource-group-name>",
		"<dns-forwarding-ruleset-name>",
		"<forwarding-rule-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.ForwardingRulesClientGetResult)
}

// x-ms-original-file: specification/dnsresolver/resource-manager/Microsoft.Network/preview/2020-04-01-preview/examples/ForwardingRule_List.json
func ExampleForwardingRulesClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdnsresolver.NewForwardingRulesClient("<subscription-id>", cred, nil)
	pager := client.List("<resource-group-name>",
		"<dns-forwarding-ruleset-name>",
		&armdnsresolver.ForwardingRulesClientListOptions{Top: nil})
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
