//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsqlvirtualmachine

import (
	"context"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"net/http"
	"time"
)

// AvailabilityGroupListenersClientCreateOrUpdatePollerResponse contains the response from method AvailabilityGroupListenersClient.CreateOrUpdate.
type AvailabilityGroupListenersClientCreateOrUpdatePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *AvailabilityGroupListenersClientCreateOrUpdatePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l AvailabilityGroupListenersClientCreateOrUpdatePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (AvailabilityGroupListenersClientCreateOrUpdateResponse, error) {
	respType := AvailabilityGroupListenersClientCreateOrUpdateResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, &respType.AvailabilityGroupListener)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a AvailabilityGroupListenersClientCreateOrUpdatePollerResponse from the provided client and resume token.
func (l *AvailabilityGroupListenersClientCreateOrUpdatePollerResponse) Resume(ctx context.Context, client *AvailabilityGroupListenersClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("AvailabilityGroupListenersClient.CreateOrUpdate", token, client.pl)
	if err != nil {
		return err
	}
	poller := &AvailabilityGroupListenersClientCreateOrUpdatePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// AvailabilityGroupListenersClientCreateOrUpdateResponse contains the response from method AvailabilityGroupListenersClient.CreateOrUpdate.
type AvailabilityGroupListenersClientCreateOrUpdateResponse struct {
	AvailabilityGroupListenersClientCreateOrUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// AvailabilityGroupListenersClientCreateOrUpdateResult contains the result from method AvailabilityGroupListenersClient.CreateOrUpdate.
type AvailabilityGroupListenersClientCreateOrUpdateResult struct {
	AvailabilityGroupListener
}

// AvailabilityGroupListenersClientDeletePollerResponse contains the response from method AvailabilityGroupListenersClient.Delete.
type AvailabilityGroupListenersClientDeletePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *AvailabilityGroupListenersClientDeletePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l AvailabilityGroupListenersClientDeletePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (AvailabilityGroupListenersClientDeleteResponse, error) {
	respType := AvailabilityGroupListenersClientDeleteResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, nil)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a AvailabilityGroupListenersClientDeletePollerResponse from the provided client and resume token.
func (l *AvailabilityGroupListenersClientDeletePollerResponse) Resume(ctx context.Context, client *AvailabilityGroupListenersClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("AvailabilityGroupListenersClient.Delete", token, client.pl)
	if err != nil {
		return err
	}
	poller := &AvailabilityGroupListenersClientDeletePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// AvailabilityGroupListenersClientDeleteResponse contains the response from method AvailabilityGroupListenersClient.Delete.
type AvailabilityGroupListenersClientDeleteResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// AvailabilityGroupListenersClientGetResponse contains the response from method AvailabilityGroupListenersClient.Get.
type AvailabilityGroupListenersClientGetResponse struct {
	AvailabilityGroupListenersClientGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// AvailabilityGroupListenersClientGetResult contains the result from method AvailabilityGroupListenersClient.Get.
type AvailabilityGroupListenersClientGetResult struct {
	AvailabilityGroupListener
}

// AvailabilityGroupListenersClientListByGroupResponse contains the response from method AvailabilityGroupListenersClient.ListByGroup.
type AvailabilityGroupListenersClientListByGroupResponse struct {
	AvailabilityGroupListenersClientListByGroupResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// AvailabilityGroupListenersClientListByGroupResult contains the result from method AvailabilityGroupListenersClient.ListByGroup.
type AvailabilityGroupListenersClientListByGroupResult struct {
	AvailabilityGroupListenerListResult
}

// GroupsClientCreateOrUpdatePollerResponse contains the response from method GroupsClient.CreateOrUpdate.
type GroupsClientCreateOrUpdatePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *GroupsClientCreateOrUpdatePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l GroupsClientCreateOrUpdatePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (GroupsClientCreateOrUpdateResponse, error) {
	respType := GroupsClientCreateOrUpdateResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, &respType.Group)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a GroupsClientCreateOrUpdatePollerResponse from the provided client and resume token.
func (l *GroupsClientCreateOrUpdatePollerResponse) Resume(ctx context.Context, client *GroupsClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("GroupsClient.CreateOrUpdate", token, client.pl)
	if err != nil {
		return err
	}
	poller := &GroupsClientCreateOrUpdatePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// GroupsClientCreateOrUpdateResponse contains the response from method GroupsClient.CreateOrUpdate.
type GroupsClientCreateOrUpdateResponse struct {
	GroupsClientCreateOrUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// GroupsClientCreateOrUpdateResult contains the result from method GroupsClient.CreateOrUpdate.
type GroupsClientCreateOrUpdateResult struct {
	Group
}

// GroupsClientDeletePollerResponse contains the response from method GroupsClient.Delete.
type GroupsClientDeletePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *GroupsClientDeletePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l GroupsClientDeletePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (GroupsClientDeleteResponse, error) {
	respType := GroupsClientDeleteResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, nil)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a GroupsClientDeletePollerResponse from the provided client and resume token.
func (l *GroupsClientDeletePollerResponse) Resume(ctx context.Context, client *GroupsClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("GroupsClient.Delete", token, client.pl)
	if err != nil {
		return err
	}
	poller := &GroupsClientDeletePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// GroupsClientDeleteResponse contains the response from method GroupsClient.Delete.
type GroupsClientDeleteResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// GroupsClientGetResponse contains the response from method GroupsClient.Get.
type GroupsClientGetResponse struct {
	GroupsClientGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// GroupsClientGetResult contains the result from method GroupsClient.Get.
type GroupsClientGetResult struct {
	Group
}

// GroupsClientListByResourceGroupResponse contains the response from method GroupsClient.ListByResourceGroup.
type GroupsClientListByResourceGroupResponse struct {
	GroupsClientListByResourceGroupResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// GroupsClientListByResourceGroupResult contains the result from method GroupsClient.ListByResourceGroup.
type GroupsClientListByResourceGroupResult struct {
	GroupListResult
}

// GroupsClientListResponse contains the response from method GroupsClient.List.
type GroupsClientListResponse struct {
	GroupsClientListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// GroupsClientListResult contains the result from method GroupsClient.List.
type GroupsClientListResult struct {
	GroupListResult
}

// GroupsClientUpdatePollerResponse contains the response from method GroupsClient.Update.
type GroupsClientUpdatePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *GroupsClientUpdatePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l GroupsClientUpdatePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (GroupsClientUpdateResponse, error) {
	respType := GroupsClientUpdateResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, &respType.Group)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a GroupsClientUpdatePollerResponse from the provided client and resume token.
func (l *GroupsClientUpdatePollerResponse) Resume(ctx context.Context, client *GroupsClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("GroupsClient.Update", token, client.pl)
	if err != nil {
		return err
	}
	poller := &GroupsClientUpdatePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// GroupsClientUpdateResponse contains the response from method GroupsClient.Update.
type GroupsClientUpdateResponse struct {
	GroupsClientUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// GroupsClientUpdateResult contains the result from method GroupsClient.Update.
type GroupsClientUpdateResult struct {
	Group
}

// OperationsClientListResponse contains the response from method OperationsClient.List.
type OperationsClientListResponse struct {
	OperationsClientListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// OperationsClientListResult contains the result from method OperationsClient.List.
type OperationsClientListResult struct {
	OperationListResult
}

// SQLVirtualMachinesClientCreateOrUpdatePollerResponse contains the response from method SQLVirtualMachinesClient.CreateOrUpdate.
type SQLVirtualMachinesClientCreateOrUpdatePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *SQLVirtualMachinesClientCreateOrUpdatePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l SQLVirtualMachinesClientCreateOrUpdatePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (SQLVirtualMachinesClientCreateOrUpdateResponse, error) {
	respType := SQLVirtualMachinesClientCreateOrUpdateResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, &respType.SQLVirtualMachine)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a SQLVirtualMachinesClientCreateOrUpdatePollerResponse from the provided client and resume token.
func (l *SQLVirtualMachinesClientCreateOrUpdatePollerResponse) Resume(ctx context.Context, client *SQLVirtualMachinesClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("SQLVirtualMachinesClient.CreateOrUpdate", token, client.pl)
	if err != nil {
		return err
	}
	poller := &SQLVirtualMachinesClientCreateOrUpdatePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// SQLVirtualMachinesClientCreateOrUpdateResponse contains the response from method SQLVirtualMachinesClient.CreateOrUpdate.
type SQLVirtualMachinesClientCreateOrUpdateResponse struct {
	SQLVirtualMachinesClientCreateOrUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// SQLVirtualMachinesClientCreateOrUpdateResult contains the result from method SQLVirtualMachinesClient.CreateOrUpdate.
type SQLVirtualMachinesClientCreateOrUpdateResult struct {
	SQLVirtualMachine
}

// SQLVirtualMachinesClientDeletePollerResponse contains the response from method SQLVirtualMachinesClient.Delete.
type SQLVirtualMachinesClientDeletePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *SQLVirtualMachinesClientDeletePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l SQLVirtualMachinesClientDeletePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (SQLVirtualMachinesClientDeleteResponse, error) {
	respType := SQLVirtualMachinesClientDeleteResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, nil)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a SQLVirtualMachinesClientDeletePollerResponse from the provided client and resume token.
func (l *SQLVirtualMachinesClientDeletePollerResponse) Resume(ctx context.Context, client *SQLVirtualMachinesClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("SQLVirtualMachinesClient.Delete", token, client.pl)
	if err != nil {
		return err
	}
	poller := &SQLVirtualMachinesClientDeletePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// SQLVirtualMachinesClientDeleteResponse contains the response from method SQLVirtualMachinesClient.Delete.
type SQLVirtualMachinesClientDeleteResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// SQLVirtualMachinesClientGetResponse contains the response from method SQLVirtualMachinesClient.Get.
type SQLVirtualMachinesClientGetResponse struct {
	SQLVirtualMachinesClientGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// SQLVirtualMachinesClientGetResult contains the result from method SQLVirtualMachinesClient.Get.
type SQLVirtualMachinesClientGetResult struct {
	SQLVirtualMachine
}

// SQLVirtualMachinesClientListByResourceGroupResponse contains the response from method SQLVirtualMachinesClient.ListByResourceGroup.
type SQLVirtualMachinesClientListByResourceGroupResponse struct {
	SQLVirtualMachinesClientListByResourceGroupResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// SQLVirtualMachinesClientListByResourceGroupResult contains the result from method SQLVirtualMachinesClient.ListByResourceGroup.
type SQLVirtualMachinesClientListByResourceGroupResult struct {
	ListResult
}

// SQLVirtualMachinesClientListBySQLVMGroupResponse contains the response from method SQLVirtualMachinesClient.ListBySQLVMGroup.
type SQLVirtualMachinesClientListBySQLVMGroupResponse struct {
	SQLVirtualMachinesClientListBySQLVMGroupResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// SQLVirtualMachinesClientListBySQLVMGroupResult contains the result from method SQLVirtualMachinesClient.ListBySQLVMGroup.
type SQLVirtualMachinesClientListBySQLVMGroupResult struct {
	ListResult
}

// SQLVirtualMachinesClientListResponse contains the response from method SQLVirtualMachinesClient.List.
type SQLVirtualMachinesClientListResponse struct {
	SQLVirtualMachinesClientListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// SQLVirtualMachinesClientListResult contains the result from method SQLVirtualMachinesClient.List.
type SQLVirtualMachinesClientListResult struct {
	ListResult
}

// SQLVirtualMachinesClientUpdatePollerResponse contains the response from method SQLVirtualMachinesClient.Update.
type SQLVirtualMachinesClientUpdatePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *SQLVirtualMachinesClientUpdatePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l SQLVirtualMachinesClientUpdatePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (SQLVirtualMachinesClientUpdateResponse, error) {
	respType := SQLVirtualMachinesClientUpdateResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, &respType.SQLVirtualMachine)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a SQLVirtualMachinesClientUpdatePollerResponse from the provided client and resume token.
func (l *SQLVirtualMachinesClientUpdatePollerResponse) Resume(ctx context.Context, client *SQLVirtualMachinesClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("SQLVirtualMachinesClient.Update", token, client.pl)
	if err != nil {
		return err
	}
	poller := &SQLVirtualMachinesClientUpdatePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// SQLVirtualMachinesClientUpdateResponse contains the response from method SQLVirtualMachinesClient.Update.
type SQLVirtualMachinesClientUpdateResponse struct {
	SQLVirtualMachinesClientUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// SQLVirtualMachinesClientUpdateResult contains the result from method SQLVirtualMachinesClient.Update.
type SQLVirtualMachinesClientUpdateResult struct {
	SQLVirtualMachine
}
