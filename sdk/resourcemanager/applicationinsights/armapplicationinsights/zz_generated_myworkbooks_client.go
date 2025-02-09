//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armapplicationinsights

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
	"strconv"
	"strings"
)

// MyWorkbooksClient contains the methods for the MyWorkbooks group.
// Don't use this type directly, use NewMyWorkbooksClient() instead.
type MyWorkbooksClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewMyWorkbooksClient creates a new instance of MyWorkbooksClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewMyWorkbooksClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *MyWorkbooksClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Endpoint) == 0 {
		cp.Endpoint = arm.AzurePublicCloud
	}
	client := &MyWorkbooksClient{
		subscriptionID: subscriptionID,
		host:           string(cp.Endpoint),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, &cp),
	}
	return client
}

// CreateOrUpdate - Create a new private workbook.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// resourceName - The name of the Application Insights component resource.
// workbookProperties - Properties that need to be specified to create a new private workbook.
// options - MyWorkbooksClientCreateOrUpdateOptions contains the optional parameters for the MyWorkbooksClient.CreateOrUpdate
// method.
func (client *MyWorkbooksClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, resourceName string, workbookProperties MyWorkbook, options *MyWorkbooksClientCreateOrUpdateOptions) (MyWorkbooksClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, resourceName, workbookProperties, options)
	if err != nil {
		return MyWorkbooksClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return MyWorkbooksClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return MyWorkbooksClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *MyWorkbooksClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, workbookProperties MyWorkbook, options *MyWorkbooksClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/myWorkbooks/{resourceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.SourceID != nil {
		reqQP.Set("sourceId", *options.SourceID)
	}
	reqQP.Set("api-version", "2021-03-08")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, workbookProperties)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *MyWorkbooksClient) createOrUpdateHandleResponse(resp *http.Response) (MyWorkbooksClientCreateOrUpdateResponse, error) {
	result := MyWorkbooksClientCreateOrUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.MyWorkbook); err != nil {
		return MyWorkbooksClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Delete a private workbook.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// resourceName - The name of the Application Insights component resource.
// options - MyWorkbooksClientDeleteOptions contains the optional parameters for the MyWorkbooksClient.Delete method.
func (client *MyWorkbooksClient) Delete(ctx context.Context, resourceGroupName string, resourceName string, options *MyWorkbooksClientDeleteOptions) (MyWorkbooksClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, resourceName, options)
	if err != nil {
		return MyWorkbooksClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return MyWorkbooksClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return MyWorkbooksClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return MyWorkbooksClientDeleteResponse{RawResponse: resp}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *MyWorkbooksClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, options *MyWorkbooksClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/myWorkbooks/{resourceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-08")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Get a single private workbook by its resourceName.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// resourceName - The name of the Application Insights component resource.
// options - MyWorkbooksClientGetOptions contains the optional parameters for the MyWorkbooksClient.Get method.
func (client *MyWorkbooksClient) Get(ctx context.Context, resourceGroupName string, resourceName string, options *MyWorkbooksClientGetOptions) (MyWorkbooksClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, resourceName, options)
	if err != nil {
		return MyWorkbooksClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return MyWorkbooksClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return MyWorkbooksClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *MyWorkbooksClient) getCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, options *MyWorkbooksClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/myWorkbooks/{resourceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-08")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *MyWorkbooksClient) getHandleResponse(resp *http.Response) (MyWorkbooksClientGetResponse, error) {
	result := MyWorkbooksClientGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.MyWorkbook); err != nil {
		return MyWorkbooksClientGetResponse{}, err
	}
	return result, nil
}

// ListByResourceGroup - Get all private workbooks defined within a specified resource group and category.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// category - Category of workbook to return.
// options - MyWorkbooksClientListByResourceGroupOptions contains the optional parameters for the MyWorkbooksClient.ListByResourceGroup
// method.
func (client *MyWorkbooksClient) ListByResourceGroup(resourceGroupName string, category CategoryType, options *MyWorkbooksClientListByResourceGroupOptions) *MyWorkbooksClientListByResourceGroupPager {
	return &MyWorkbooksClientListByResourceGroupPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, category, options)
		},
		advancer: func(ctx context.Context, resp MyWorkbooksClientListByResourceGroupResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.MyWorkbooksListResult.NextLink)
		},
	}
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *MyWorkbooksClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, category CategoryType, options *MyWorkbooksClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/myWorkbooks"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("category", string(category))
	if options != nil && options.Tags != nil {
		reqQP.Set("tags", strings.Join(options.Tags, ","))
	}
	if options != nil && options.SourceID != nil {
		reqQP.Set("sourceId", *options.SourceID)
	}
	if options != nil && options.CanFetchContent != nil {
		reqQP.Set("canFetchContent", strconv.FormatBool(*options.CanFetchContent))
	}
	reqQP.Set("api-version", "2021-03-08")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *MyWorkbooksClient) listByResourceGroupHandleResponse(resp *http.Response) (MyWorkbooksClientListByResourceGroupResponse, error) {
	result := MyWorkbooksClientListByResourceGroupResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.MyWorkbooksListResult); err != nil {
		return MyWorkbooksClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// ListBySubscription - Get all private workbooks defined within a specified subscription and category.
// If the operation fails it returns an *azcore.ResponseError type.
// category - Category of workbook to return.
// options - MyWorkbooksClientListBySubscriptionOptions contains the optional parameters for the MyWorkbooksClient.ListBySubscription
// method.
func (client *MyWorkbooksClient) ListBySubscription(category CategoryType, options *MyWorkbooksClientListBySubscriptionOptions) *MyWorkbooksClientListBySubscriptionPager {
	return &MyWorkbooksClientListBySubscriptionPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listBySubscriptionCreateRequest(ctx, category, options)
		},
		advancer: func(ctx context.Context, resp MyWorkbooksClientListBySubscriptionResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.MyWorkbooksListResult.NextLink)
		},
	}
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *MyWorkbooksClient) listBySubscriptionCreateRequest(ctx context.Context, category CategoryType, options *MyWorkbooksClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Insights/myWorkbooks"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("category", string(category))
	if options != nil && options.Tags != nil {
		reqQP.Set("tags", strings.Join(options.Tags, ","))
	}
	if options != nil && options.CanFetchContent != nil {
		reqQP.Set("canFetchContent", strconv.FormatBool(*options.CanFetchContent))
	}
	reqQP.Set("api-version", "2021-03-08")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *MyWorkbooksClient) listBySubscriptionHandleResponse(resp *http.Response) (MyWorkbooksClientListBySubscriptionResponse, error) {
	result := MyWorkbooksClientListBySubscriptionResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.MyWorkbooksListResult); err != nil {
		return MyWorkbooksClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// Update - Updates a private workbook that has already been added.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// resourceName - The name of the Application Insights component resource.
// workbookProperties - Properties that need to be specified to create a new private workbook.
// options - MyWorkbooksClientUpdateOptions contains the optional parameters for the MyWorkbooksClient.Update method.
func (client *MyWorkbooksClient) Update(ctx context.Context, resourceGroupName string, resourceName string, workbookProperties MyWorkbook, options *MyWorkbooksClientUpdateOptions) (MyWorkbooksClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, resourceName, workbookProperties, options)
	if err != nil {
		return MyWorkbooksClientUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return MyWorkbooksClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusCreated) {
		return MyWorkbooksClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *MyWorkbooksClient) updateCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, workbookProperties MyWorkbook, options *MyWorkbooksClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/myWorkbooks/{resourceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.SourceID != nil {
		reqQP.Set("sourceId", *options.SourceID)
	}
	reqQP.Set("api-version", "2021-03-08")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, workbookProperties)
}

// updateHandleResponse handles the Update response.
func (client *MyWorkbooksClient) updateHandleResponse(resp *http.Response) (MyWorkbooksClientUpdateResponse, error) {
	result := MyWorkbooksClientUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.MyWorkbook); err != nil {
		return MyWorkbooksClientUpdateResponse{}, err
	}
	return result, nil
}
