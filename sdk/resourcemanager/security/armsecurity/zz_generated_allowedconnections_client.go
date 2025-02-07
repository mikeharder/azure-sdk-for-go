//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsecurity

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

// AllowedConnectionsClient contains the methods for the AllowedConnections group.
// Don't use this type directly, use NewAllowedConnectionsClient() instead.
type AllowedConnectionsClient struct {
	host           string
	subscriptionID string
	ascLocation    string
	pl             runtime.Pipeline
}

// NewAllowedConnectionsClient creates a new instance of AllowedConnectionsClient with the specified values.
// subscriptionID - Azure subscription ID
// ascLocation - The location where ASC stores the data of the subscription. can be retrieved from Get locations
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewAllowedConnectionsClient(subscriptionID string, ascLocation string, credential azcore.TokenCredential, options *arm.ClientOptions) *AllowedConnectionsClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &AllowedConnectionsClient{
		subscriptionID: subscriptionID,
		ascLocation:    ascLocation,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// Get - Gets the list of all possible traffic between resources for the subscription and location, based on connection type.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group within the user's subscription. The name is case insensitive.
// connectionType - The type of allowed connections (Internal, External)
// options - AllowedConnectionsClientGetOptions contains the optional parameters for the AllowedConnectionsClient.Get method.
func (client *AllowedConnectionsClient) Get(ctx context.Context, resourceGroupName string, connectionType ConnectionType, options *AllowedConnectionsClientGetOptions) (AllowedConnectionsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, connectionType, options)
	if err != nil {
		return AllowedConnectionsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AllowedConnectionsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AllowedConnectionsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *AllowedConnectionsClient) getCreateRequest(ctx context.Context, resourceGroupName string, connectionType ConnectionType, options *AllowedConnectionsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Security/locations/{ascLocation}/allowedConnections/{connectionType}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.ascLocation == "" {
		return nil, errors.New("parameter client.ascLocation cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ascLocation}", url.PathEscape(client.ascLocation))
	if connectionType == "" {
		return nil, errors.New("parameter connectionType cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{connectionType}", url.PathEscape(string(connectionType)))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *AllowedConnectionsClient) getHandleResponse(resp *http.Response) (AllowedConnectionsClientGetResponse, error) {
	result := AllowedConnectionsClientGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.AllowedConnectionsResource); err != nil {
		return AllowedConnectionsClientGetResponse{}, err
	}
	return result, nil
}

// List - Gets the list of all possible traffic between resources for the subscription
// If the operation fails it returns an *azcore.ResponseError type.
// options - AllowedConnectionsClientListOptions contains the optional parameters for the AllowedConnectionsClient.List method.
func (client *AllowedConnectionsClient) List(options *AllowedConnectionsClientListOptions) *AllowedConnectionsClientListPager {
	return &AllowedConnectionsClientListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp AllowedConnectionsClientListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.AllowedConnectionsList.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *AllowedConnectionsClient) listCreateRequest(ctx context.Context, options *AllowedConnectionsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Security/allowedConnections"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *AllowedConnectionsClient) listHandleResponse(resp *http.Response) (AllowedConnectionsClientListResponse, error) {
	result := AllowedConnectionsClientListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.AllowedConnectionsList); err != nil {
		return AllowedConnectionsClientListResponse{}, err
	}
	return result, nil
}

// ListByHomeRegion - Gets the list of all possible traffic between resources for the subscription and location.
// If the operation fails it returns an *azcore.ResponseError type.
// options - AllowedConnectionsClientListByHomeRegionOptions contains the optional parameters for the AllowedConnectionsClient.ListByHomeRegion
// method.
func (client *AllowedConnectionsClient) ListByHomeRegion(options *AllowedConnectionsClientListByHomeRegionOptions) *AllowedConnectionsClientListByHomeRegionPager {
	return &AllowedConnectionsClientListByHomeRegionPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByHomeRegionCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp AllowedConnectionsClientListByHomeRegionResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.AllowedConnectionsList.NextLink)
		},
	}
}

// listByHomeRegionCreateRequest creates the ListByHomeRegion request.
func (client *AllowedConnectionsClient) listByHomeRegionCreateRequest(ctx context.Context, options *AllowedConnectionsClientListByHomeRegionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Security/locations/{ascLocation}/allowedConnections"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if client.ascLocation == "" {
		return nil, errors.New("parameter client.ascLocation cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ascLocation}", url.PathEscape(client.ascLocation))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByHomeRegionHandleResponse handles the ListByHomeRegion response.
func (client *AllowedConnectionsClient) listByHomeRegionHandleResponse(resp *http.Response) (AllowedConnectionsClientListByHomeRegionResponse, error) {
	result := AllowedConnectionsClientListByHomeRegionResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.AllowedConnectionsList); err != nil {
		return AllowedConnectionsClientListByHomeRegionResponse{}, err
	}
	return result, nil
}
