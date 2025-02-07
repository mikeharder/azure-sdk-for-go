//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armhardwaresecuritymodules

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
)

// DedicatedHsmClientCreateOrUpdatePoller provides polling facilities until the operation reaches a terminal state.
type DedicatedHsmClientCreateOrUpdatePoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *DedicatedHsmClientCreateOrUpdatePoller) Done() bool {
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
func (p *DedicatedHsmClientCreateOrUpdatePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final DedicatedHsmClientCreateOrUpdateResponse will be returned.
func (p *DedicatedHsmClientCreateOrUpdatePoller) FinalResponse(ctx context.Context) (DedicatedHsmClientCreateOrUpdateResponse, error) {
	respType := DedicatedHsmClientCreateOrUpdateResponse{}
	resp, err := p.pt.FinalResponse(ctx, &respType.DedicatedHsm)
	if err != nil {
		return DedicatedHsmClientCreateOrUpdateResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *DedicatedHsmClientCreateOrUpdatePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// DedicatedHsmClientDeletePoller provides polling facilities until the operation reaches a terminal state.
type DedicatedHsmClientDeletePoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *DedicatedHsmClientDeletePoller) Done() bool {
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
func (p *DedicatedHsmClientDeletePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final DedicatedHsmClientDeleteResponse will be returned.
func (p *DedicatedHsmClientDeletePoller) FinalResponse(ctx context.Context) (DedicatedHsmClientDeleteResponse, error) {
	respType := DedicatedHsmClientDeleteResponse{}
	resp, err := p.pt.FinalResponse(ctx, nil)
	if err != nil {
		return DedicatedHsmClientDeleteResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *DedicatedHsmClientDeletePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// DedicatedHsmClientUpdatePoller provides polling facilities until the operation reaches a terminal state.
type DedicatedHsmClientUpdatePoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *DedicatedHsmClientUpdatePoller) Done() bool {
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
func (p *DedicatedHsmClientUpdatePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final DedicatedHsmClientUpdateResponse will be returned.
func (p *DedicatedHsmClientUpdatePoller) FinalResponse(ctx context.Context) (DedicatedHsmClientUpdateResponse, error) {
	respType := DedicatedHsmClientUpdateResponse{}
	resp, err := p.pt.FinalResponse(ctx, &respType.DedicatedHsm)
	if err != nil {
		return DedicatedHsmClientUpdateResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *DedicatedHsmClientUpdatePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}
