//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armautomation_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/automation/armautomation"
)

// x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/deleteConnection.json
func ExampleConnectionClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armautomation.NewConnectionClient("<subscription-id>", cred, nil)
	_, err = client.Delete(ctx,
		"<resource-group-name>",
		"<automation-account-name>",
		"<connection-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/getConnection.json
func ExampleConnectionClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armautomation.NewConnectionClient("<subscription-id>", cred, nil)
	res, err := client.Get(ctx,
		"<resource-group-name>",
		"<automation-account-name>",
		"<connection-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.ConnectionClientGetResult)
}

// x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/createOrUpdateConnection.json
func ExampleConnectionClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armautomation.NewConnectionClient("<subscription-id>", cred, nil)
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<automation-account-name>",
		"<connection-name>",
		armautomation.ConnectionCreateOrUpdateParameters{
			Name: to.StringPtr("<name>"),
			Properties: &armautomation.ConnectionCreateOrUpdateProperties{
				Description: to.StringPtr("<description>"),
				ConnectionType: &armautomation.ConnectionTypeAssociationProperty{
					Name: to.StringPtr("<name>"),
				},
				FieldDefinitionValues: map[string]*string{
					"AutomationCertificateName": to.StringPtr("mysCertificateName"),
					"SubscriptionID":            to.StringPtr("subid"),
				},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.ConnectionClientCreateOrUpdateResult)
}

// x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/updateConnection.json
func ExampleConnectionClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armautomation.NewConnectionClient("<subscription-id>", cred, nil)
	res, err := client.Update(ctx,
		"<resource-group-name>",
		"<automation-account-name>",
		"<connection-name>",
		armautomation.ConnectionUpdateParameters{
			Name: to.StringPtr("<name>"),
			Properties: &armautomation.ConnectionUpdateProperties{
				Description: to.StringPtr("<description>"),
				FieldDefinitionValues: map[string]*string{
					"AutomationCertificateName": to.StringPtr("myCertificateName"),
					"SubscriptionID":            to.StringPtr("b5e4748c-f69a-467c-8749-e2f9c8cd3009"),
				},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.ConnectionClientUpdateResult)
}

// x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/listConnectionsByAutomationAccount_First100.json
func ExampleConnectionClient_ListByAutomationAccount() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armautomation.NewConnectionClient("<subscription-id>", cred, nil)
	pager := client.ListByAutomationAccount("<resource-group-name>",
		"<automation-account-name>",
		nil)
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
