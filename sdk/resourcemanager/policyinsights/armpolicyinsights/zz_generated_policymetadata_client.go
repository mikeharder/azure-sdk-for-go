//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armpolicyinsights

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"strconv"
	"strings"
)

// PolicyMetadataClient contains the methods for the PolicyMetadata group.
// Don't use this type directly, use NewPolicyMetadataClient() instead.
type PolicyMetadataClient struct {
	host string
	pl   runtime.Pipeline
}

// NewPolicyMetadataClient creates a new instance of PolicyMetadataClient with the specified values.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewPolicyMetadataClient(credential azcore.TokenCredential, options *arm.ClientOptions) *PolicyMetadataClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &PolicyMetadataClient{
		host: string(ep),
		pl:   armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// GetResource - Get policy metadata resource.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceName - The name of the policy metadata resource.
// options - PolicyMetadataClientGetResourceOptions contains the optional parameters for the PolicyMetadataClient.GetResource
// method.
func (client *PolicyMetadataClient) GetResource(ctx context.Context, resourceName string, options *PolicyMetadataClientGetResourceOptions) (PolicyMetadataClientGetResourceResponse, error) {
	req, err := client.getResourceCreateRequest(ctx, resourceName, options)
	if err != nil {
		return PolicyMetadataClientGetResourceResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PolicyMetadataClientGetResourceResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PolicyMetadataClientGetResourceResponse{}, runtime.NewResponseError(resp)
	}
	return client.getResourceHandleResponse(resp)
}

// getResourceCreateRequest creates the GetResource request.
func (client *PolicyMetadataClient) getResourceCreateRequest(ctx context.Context, resourceName string, options *PolicyMetadataClientGetResourceOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.PolicyInsights/policyMetadata/{resourceName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", resourceName)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getResourceHandleResponse handles the GetResource response.
func (client *PolicyMetadataClient) getResourceHandleResponse(resp *http.Response) (PolicyMetadataClientGetResourceResponse, error) {
	result := PolicyMetadataClientGetResourceResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.PolicyMetadata); err != nil {
		return PolicyMetadataClientGetResourceResponse{}, err
	}
	return result, nil
}

// List - Get a list of the policy metadata resources.
// If the operation fails it returns an *azcore.ResponseError type.
// options - QueryOptions contains a group of parameters for the PolicyTrackedResourcesClient.ListQueryResultsForManagementGroup
// method.
func (client *PolicyMetadataClient) List(options *QueryOptions) *PolicyMetadataClientListPager {
	return &PolicyMetadataClientListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp PolicyMetadataClientListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.PolicyMetadataCollection.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *PolicyMetadataClient) listCreateRequest(ctx context.Context, options *QueryOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.PolicyInsights/policyMetadata"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-10-01")
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *PolicyMetadataClient) listHandleResponse(resp *http.Response) (PolicyMetadataClientListResponse, error) {
	result := PolicyMetadataClientListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.PolicyMetadataCollection); err != nil {
		return PolicyMetadataClientListResponse{}, err
	}
	return result, nil
}
