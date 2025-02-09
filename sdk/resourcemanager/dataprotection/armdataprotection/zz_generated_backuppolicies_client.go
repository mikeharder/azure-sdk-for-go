//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdataprotection

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// BackupPoliciesClient contains the methods for the BackupPolicies group.
// Don't use this type directly, use NewBackupPoliciesClient() instead.
type BackupPoliciesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewBackupPoliciesClient creates a new instance of BackupPoliciesClient with the specified values.
// subscriptionID - The subscription Id.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewBackupPoliciesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *BackupPoliciesClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Endpoint) == 0 {
		cp.Endpoint = arm.AzurePublicCloud
	}
	client := &BackupPoliciesClient{
		subscriptionID: subscriptionID,
		host:           string(cp.Endpoint),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, &cp),
	}
	return client
}

// CreateOrUpdate - Creates or Updates a backup policy belonging to a backup vault
// If the operation fails it returns an *azcore.ResponseError type.
// vaultName - The name of the backup vault.
// resourceGroupName - The name of the resource group where the backup vault is present.
// backupPolicyName - Name of the policy
// parameters - Request body for operation
// options - BackupPoliciesClientCreateOrUpdateOptions contains the optional parameters for the BackupPoliciesClient.CreateOrUpdate
// method.
func (client *BackupPoliciesClient) CreateOrUpdate(ctx context.Context, vaultName string, resourceGroupName string, backupPolicyName string, parameters BaseBackupPolicyResource, options *BackupPoliciesClientCreateOrUpdateOptions) (BackupPoliciesClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, vaultName, resourceGroupName, backupPolicyName, parameters, options)
	if err != nil {
		return BackupPoliciesClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return BackupPoliciesClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return BackupPoliciesClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *BackupPoliciesClient) createOrUpdateCreateRequest(ctx context.Context, vaultName string, resourceGroupName string, backupPolicyName string, parameters BaseBackupPolicyResource, options *BackupPoliciesClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataProtection/backupVaults/{vaultName}/backupPolicies/{backupPolicyName}"
	if vaultName == "" {
		return nil, errors.New("parameter vaultName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vaultName}", url.PathEscape(vaultName))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if backupPolicyName == "" {
		return nil, errors.New("parameter backupPolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{backupPolicyName}", url.PathEscape(backupPolicyName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *BackupPoliciesClient) createOrUpdateHandleResponse(resp *http.Response) (BackupPoliciesClientCreateOrUpdateResponse, error) {
	result := BackupPoliciesClientCreateOrUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.BaseBackupPolicyResource); err != nil {
		return BackupPoliciesClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes a backup policy belonging to a backup vault
// If the operation fails it returns an *azcore.ResponseError type.
// vaultName - The name of the backup vault.
// resourceGroupName - The name of the resource group where the backup vault is present.
// options - BackupPoliciesClientDeleteOptions contains the optional parameters for the BackupPoliciesClient.Delete method.
func (client *BackupPoliciesClient) Delete(ctx context.Context, vaultName string, resourceGroupName string, backupPolicyName string, options *BackupPoliciesClientDeleteOptions) (BackupPoliciesClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, vaultName, resourceGroupName, backupPolicyName, options)
	if err != nil {
		return BackupPoliciesClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return BackupPoliciesClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return BackupPoliciesClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return BackupPoliciesClientDeleteResponse{RawResponse: resp}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *BackupPoliciesClient) deleteCreateRequest(ctx context.Context, vaultName string, resourceGroupName string, backupPolicyName string, options *BackupPoliciesClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataProtection/backupVaults/{vaultName}/backupPolicies/{backupPolicyName}"
	if vaultName == "" {
		return nil, errors.New("parameter vaultName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vaultName}", url.PathEscape(vaultName))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if backupPolicyName == "" {
		return nil, errors.New("parameter backupPolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{backupPolicyName}", url.PathEscape(backupPolicyName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Gets a backup policy belonging to a backup vault
// If the operation fails it returns an *azcore.ResponseError type.
// vaultName - The name of the backup vault.
// resourceGroupName - The name of the resource group where the backup vault is present.
// options - BackupPoliciesClientGetOptions contains the optional parameters for the BackupPoliciesClient.Get method.
func (client *BackupPoliciesClient) Get(ctx context.Context, vaultName string, resourceGroupName string, backupPolicyName string, options *BackupPoliciesClientGetOptions) (BackupPoliciesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, vaultName, resourceGroupName, backupPolicyName, options)
	if err != nil {
		return BackupPoliciesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return BackupPoliciesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return BackupPoliciesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *BackupPoliciesClient) getCreateRequest(ctx context.Context, vaultName string, resourceGroupName string, backupPolicyName string, options *BackupPoliciesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataProtection/backupVaults/{vaultName}/backupPolicies/{backupPolicyName}"
	if vaultName == "" {
		return nil, errors.New("parameter vaultName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vaultName}", url.PathEscape(vaultName))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if backupPolicyName == "" {
		return nil, errors.New("parameter backupPolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{backupPolicyName}", url.PathEscape(backupPolicyName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *BackupPoliciesClient) getHandleResponse(resp *http.Response) (BackupPoliciesClientGetResponse, error) {
	result := BackupPoliciesClientGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.BaseBackupPolicyResource); err != nil {
		return BackupPoliciesClientGetResponse{}, err
	}
	return result, nil
}

// List - Returns list of backup policies belonging to a backup vault
// If the operation fails it returns an *azcore.ResponseError type.
// vaultName - The name of the backup vault.
// resourceGroupName - The name of the resource group where the backup vault is present.
// options - BackupPoliciesClientListOptions contains the optional parameters for the BackupPoliciesClient.List method.
func (client *BackupPoliciesClient) List(vaultName string, resourceGroupName string, options *BackupPoliciesClientListOptions) *BackupPoliciesClientListPager {
	return &BackupPoliciesClientListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, vaultName, resourceGroupName, options)
		},
		advancer: func(ctx context.Context, resp BackupPoliciesClientListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.BaseBackupPolicyResourceList.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *BackupPoliciesClient) listCreateRequest(ctx context.Context, vaultName string, resourceGroupName string, options *BackupPoliciesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataProtection/backupVaults/{vaultName}/backupPolicies"
	if vaultName == "" {
		return nil, errors.New("parameter vaultName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vaultName}", url.PathEscape(vaultName))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *BackupPoliciesClient) listHandleResponse(resp *http.Response) (BackupPoliciesClientListResponse, error) {
	result := BackupPoliciesClientListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.BaseBackupPolicyResourceList); err != nil {
		return BackupPoliciesClientListResponse{}, err
	}
	return result, nil
}
