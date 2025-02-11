//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armiotsecurity

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

// LocationsClient contains the methods for the Locations group.
// Don't use this type directly, use NewLocationsClient() instead.
type LocationsClient struct {
	host                string
	subscriptionID      string
	iotDefenderLocation string
	pl                  runtime.Pipeline
}

// NewLocationsClient creates a new instance of LocationsClient with the specified values.
// subscriptionID - The ID of the target subscription.
// iotDefenderLocation - Defender for IoT location
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewLocationsClient(subscriptionID string, iotDefenderLocation string, credential azcore.TokenCredential, options *arm.ClientOptions) *LocationsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Endpoint) == 0 {
		cp.Endpoint = arm.AzurePublicCloud
	}
	client := &LocationsClient{
		subscriptionID:      subscriptionID,
		iotDefenderLocation: iotDefenderLocation,
		host:                string(cp.Endpoint),
		pl:                  armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, &cp),
	}
	return client
}

// Get - Get a Defender for IoT location associated with the given subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// options - LocationsClientGetOptions contains the optional parameters for the LocationsClient.Get method.
func (client *LocationsClient) Get(ctx context.Context, options *LocationsClientGetOptions) (LocationsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, options)
	if err != nil {
		return LocationsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return LocationsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return LocationsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *LocationsClient) getCreateRequest(ctx context.Context, options *LocationsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.IoTSecurity/locations/{iotDefenderLocation}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if client.iotDefenderLocation == "" {
		return nil, errors.New("parameter client.iotDefenderLocation cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{iotDefenderLocation}", url.PathEscape(client.iotDefenderLocation))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-02-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *LocationsClient) getHandleResponse(resp *http.Response) (LocationsClientGetResponse, error) {
	result := LocationsClientGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.LocationModel); err != nil {
		return LocationsClientGetResponse{}, err
	}
	return result, nil
}

// List - Lists Defender for IoT locations associated with the given subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// options - LocationsClientListOptions contains the optional parameters for the LocationsClient.List method.
func (client *LocationsClient) List(options *LocationsClientListOptions) *LocationsClientListPager {
	return &LocationsClientListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp LocationsClientListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.LocationList.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *LocationsClient) listCreateRequest(ctx context.Context, options *LocationsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.IoTSecurity/locations"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-02-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *LocationsClient) listHandleResponse(resp *http.Response) (LocationsClientListResponse, error) {
	result := LocationsClientListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.LocationList); err != nil {
		return LocationsClientListResponse{}, err
	}
	return result, nil
}
