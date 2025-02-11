//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcognitiveservices_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cognitiveservices/armcognitiveservices"
)

// x-ms-original-file: specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2021-10-01/examples/CreateAccount.json
func ExampleAccountsClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armcognitiveservices.NewAccountsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreate(ctx,
		"<resource-group-name>",
		"<account-name>",
		armcognitiveservices.Account{
			Identity: &armcognitiveservices.Identity{
				Type: armcognitiveservices.ResourceIdentityTypeSystemAssigned.ToPtr(),
			},
			Kind:     to.StringPtr("<kind>"),
			Location: to.StringPtr("<location>"),
			Properties: &armcognitiveservices.AccountProperties{
				Encryption: &armcognitiveservices.Encryption{
					KeySource: armcognitiveservices.KeySource("Microsoft.KeyVault").ToPtr(),
					KeyVaultProperties: &armcognitiveservices.KeyVaultProperties{
						KeyName:     to.StringPtr("<key-name>"),
						KeyVaultURI: to.StringPtr("<key-vault-uri>"),
						KeyVersion:  to.StringPtr("<key-version>"),
					},
				},
				UserOwnedStorage: []*armcognitiveservices.UserOwnedStorage{
					{
						ResourceID: to.StringPtr("<resource-id>"),
					}},
			},
			SKU: &armcognitiveservices.SKU{
				Name: to.StringPtr("<name>"),
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
	log.Printf("Response result: %#v\n", res.AccountsClientCreateResult)
}

// x-ms-original-file: specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2021-10-01/examples/UpdateAccount.json
func ExampleAccountsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armcognitiveservices.NewAccountsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginUpdate(ctx,
		"<resource-group-name>",
		"<account-name>",
		armcognitiveservices.Account{
			Location: to.StringPtr("<location>"),
			SKU: &armcognitiveservices.SKU{
				Name: to.StringPtr("<name>"),
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
	log.Printf("Response result: %#v\n", res.AccountsClientUpdateResult)
}

// x-ms-original-file: specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2021-10-01/examples/DeleteAccount.json
func ExampleAccountsClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armcognitiveservices.NewAccountsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginDelete(ctx,
		"<resource-group-name>",
		"<account-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2021-10-01/examples/GetAccount.json
func ExampleAccountsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armcognitiveservices.NewAccountsClient("<subscription-id>", cred, nil)
	res, err := client.Get(ctx,
		"<resource-group-name>",
		"<account-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.AccountsClientGetResult)
}

// x-ms-original-file: specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2021-10-01/examples/ListAccountsByResourceGroup.json
func ExampleAccountsClient_ListByResourceGroup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armcognitiveservices.NewAccountsClient("<subscription-id>", cred, nil)
	pager := client.ListByResourceGroup("<resource-group-name>",
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

// x-ms-original-file: specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2021-10-01/examples/ListAccountsBySubscription.json
func ExampleAccountsClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armcognitiveservices.NewAccountsClient("<subscription-id>", cred, nil)
	pager := client.List(nil)
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

// x-ms-original-file: specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2021-10-01/examples/ListKeys.json
func ExampleAccountsClient_ListKeys() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armcognitiveservices.NewAccountsClient("<subscription-id>", cred, nil)
	res, err := client.ListKeys(ctx,
		"<resource-group-name>",
		"<account-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.AccountsClientListKeysResult)
}

// x-ms-original-file: specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2021-10-01/examples/RegenerateKey.json
func ExampleAccountsClient_RegenerateKey() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armcognitiveservices.NewAccountsClient("<subscription-id>", cred, nil)
	res, err := client.RegenerateKey(ctx,
		"<resource-group-name>",
		"<account-name>",
		armcognitiveservices.RegenerateKeyParameters{
			KeyName: armcognitiveservices.KeyNameKey2.ToPtr(),
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.AccountsClientRegenerateKeyResult)
}

// x-ms-original-file: specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2021-10-01/examples/GetUsages.json
func ExampleAccountsClient_ListUsages() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armcognitiveservices.NewAccountsClient("<subscription-id>", cred, nil)
	res, err := client.ListUsages(ctx,
		"<resource-group-name>",
		"<account-name>",
		&armcognitiveservices.AccountsClientListUsagesOptions{Filter: nil})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.AccountsClientListUsagesResult)
}
