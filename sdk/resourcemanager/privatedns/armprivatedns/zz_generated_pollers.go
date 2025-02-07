//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armprivatedns

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
)

// PrivateZonesClientCreateOrUpdatePoller provides polling facilities until the operation reaches a terminal state.
type PrivateZonesClientCreateOrUpdatePoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *PrivateZonesClientCreateOrUpdatePoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *PrivateZonesClientCreateOrUpdatePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final PrivateZonesClientCreateOrUpdateResponse will be returned.
func (p *PrivateZonesClientCreateOrUpdatePoller) FinalResponse(ctx context.Context) (PrivateZonesClientCreateOrUpdateResponse, error) {
	respType := PrivateZonesClientCreateOrUpdateResponse{}
	resp, err := p.pt.FinalResponse(ctx, &respType.PrivateZone)
	if err != nil {
		return PrivateZonesClientCreateOrUpdateResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *PrivateZonesClientCreateOrUpdatePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// PrivateZonesClientDeletePoller provides polling facilities until the operation reaches a terminal state.
type PrivateZonesClientDeletePoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *PrivateZonesClientDeletePoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *PrivateZonesClientDeletePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final PrivateZonesClientDeleteResponse will be returned.
func (p *PrivateZonesClientDeletePoller) FinalResponse(ctx context.Context) (PrivateZonesClientDeleteResponse, error) {
	respType := PrivateZonesClientDeleteResponse{}
	resp, err := p.pt.FinalResponse(ctx, nil)
	if err != nil {
		return PrivateZonesClientDeleteResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *PrivateZonesClientDeletePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// PrivateZonesClientUpdatePoller provides polling facilities until the operation reaches a terminal state.
type PrivateZonesClientUpdatePoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *PrivateZonesClientUpdatePoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *PrivateZonesClientUpdatePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final PrivateZonesClientUpdateResponse will be returned.
func (p *PrivateZonesClientUpdatePoller) FinalResponse(ctx context.Context) (PrivateZonesClientUpdateResponse, error) {
	respType := PrivateZonesClientUpdateResponse{}
	resp, err := p.pt.FinalResponse(ctx, &respType.PrivateZone)
	if err != nil {
		return PrivateZonesClientUpdateResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *PrivateZonesClientUpdatePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// VirtualNetworkLinksClientCreateOrUpdatePoller provides polling facilities until the operation reaches a terminal state.
type VirtualNetworkLinksClientCreateOrUpdatePoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *VirtualNetworkLinksClientCreateOrUpdatePoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *VirtualNetworkLinksClientCreateOrUpdatePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final VirtualNetworkLinksClientCreateOrUpdateResponse will be returned.
func (p *VirtualNetworkLinksClientCreateOrUpdatePoller) FinalResponse(ctx context.Context) (VirtualNetworkLinksClientCreateOrUpdateResponse, error) {
	respType := VirtualNetworkLinksClientCreateOrUpdateResponse{}
	resp, err := p.pt.FinalResponse(ctx, &respType.VirtualNetworkLink)
	if err != nil {
		return VirtualNetworkLinksClientCreateOrUpdateResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *VirtualNetworkLinksClientCreateOrUpdatePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// VirtualNetworkLinksClientDeletePoller provides polling facilities until the operation reaches a terminal state.
type VirtualNetworkLinksClientDeletePoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *VirtualNetworkLinksClientDeletePoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *VirtualNetworkLinksClientDeletePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final VirtualNetworkLinksClientDeleteResponse will be returned.
func (p *VirtualNetworkLinksClientDeletePoller) FinalResponse(ctx context.Context) (VirtualNetworkLinksClientDeleteResponse, error) {
	respType := VirtualNetworkLinksClientDeleteResponse{}
	resp, err := p.pt.FinalResponse(ctx, nil)
	if err != nil {
		return VirtualNetworkLinksClientDeleteResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *VirtualNetworkLinksClientDeletePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// VirtualNetworkLinksClientUpdatePoller provides polling facilities until the operation reaches a terminal state.
type VirtualNetworkLinksClientUpdatePoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *VirtualNetworkLinksClientUpdatePoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *VirtualNetworkLinksClientUpdatePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final VirtualNetworkLinksClientUpdateResponse will be returned.
func (p *VirtualNetworkLinksClientUpdatePoller) FinalResponse(ctx context.Context) (VirtualNetworkLinksClientUpdateResponse, error) {
	respType := VirtualNetworkLinksClientUpdateResponse{}
	resp, err := p.pt.FinalResponse(ctx, &respType.VirtualNetworkLink)
	if err != nil {
		return VirtualNetworkLinksClientUpdateResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *VirtualNetworkLinksClientUpdatePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}
