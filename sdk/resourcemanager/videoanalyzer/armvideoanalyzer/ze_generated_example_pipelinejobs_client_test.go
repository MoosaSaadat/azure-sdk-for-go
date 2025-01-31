//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armvideoanalyzer_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/videoanalyzer/armvideoanalyzer"
)

// x-ms-original-file: specification/videoanalyzer/resource-manager/Microsoft.Media/preview/2021-11-01-preview/examples/pipeline-job-list.json
func ExamplePipelineJobsClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armvideoanalyzer.NewPipelineJobsClient("<subscription-id>", cred, nil)
	pager := client.List("<resource-group-name>",
		"<account-name>",
		&armvideoanalyzer.PipelineJobsClientListOptions{Filter: nil,
			Top: to.Int32Ptr(2),
		})
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

// x-ms-original-file: specification/videoanalyzer/resource-manager/Microsoft.Media/preview/2021-11-01-preview/examples/pipeline-job-get-by-name.json
func ExamplePipelineJobsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armvideoanalyzer.NewPipelineJobsClient("<subscription-id>", cred, nil)
	res, err := client.Get(ctx,
		"<resource-group-name>",
		"<account-name>",
		"<pipeline-job-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.PipelineJobsClientGetResult)
}

// x-ms-original-file: specification/videoanalyzer/resource-manager/Microsoft.Media/preview/2021-11-01-preview/examples/pipeline-job-create.json
func ExamplePipelineJobsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armvideoanalyzer.NewPipelineJobsClient("<subscription-id>", cred, nil)
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<account-name>",
		"<pipeline-job-name>",
		armvideoanalyzer.PipelineJob{
			Properties: &armvideoanalyzer.PipelineJobProperties{
				Description: to.StringPtr("<description>"),
				Parameters: []*armvideoanalyzer.ParameterDefinition{
					{
						Name:  to.StringPtr("<name>"),
						Value: to.StringPtr("<value>"),
					},
					{
						Name:  to.StringPtr("<name>"),
						Value: to.StringPtr("<value>"),
					}},
				TopologyName: to.StringPtr("<topology-name>"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.PipelineJobsClientCreateOrUpdateResult)
}

// x-ms-original-file: specification/videoanalyzer/resource-manager/Microsoft.Media/preview/2021-11-01-preview/examples/pipeline-job-delete.json
func ExamplePipelineJobsClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armvideoanalyzer.NewPipelineJobsClient("<subscription-id>", cred, nil)
	_, err = client.Delete(ctx,
		"<resource-group-name>",
		"<account-name>",
		"<pipeline-job-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/videoanalyzer/resource-manager/Microsoft.Media/preview/2021-11-01-preview/examples/pipeline-job-patch.json
func ExamplePipelineJobsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armvideoanalyzer.NewPipelineJobsClient("<subscription-id>", cred, nil)
	res, err := client.Update(ctx,
		"<resource-group-name>",
		"<account-name>",
		"<pipeline-job-name>",
		armvideoanalyzer.PipelineJobUpdate{
			Properties: &armvideoanalyzer.PipelineJobPropertiesUpdate{
				Description: to.StringPtr("<description>"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.PipelineJobsClientUpdateResult)
}

// x-ms-original-file: specification/videoanalyzer/resource-manager/Microsoft.Media/preview/2021-11-01-preview/examples/pipeline-job-cancel.json
func ExamplePipelineJobsClient_BeginCancel() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armvideoanalyzer.NewPipelineJobsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCancel(ctx,
		"<resource-group-name>",
		"<account-name>",
		"<pipeline-job-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
}
