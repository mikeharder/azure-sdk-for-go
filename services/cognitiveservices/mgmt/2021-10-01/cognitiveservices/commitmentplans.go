package cognitiveservices

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// CommitmentPlansClient is the cognitive Services Management Client
type CommitmentPlansClient struct {
	BaseClient
}

// NewCommitmentPlansClient creates an instance of the CommitmentPlansClient client.
func NewCommitmentPlansClient(subscriptionID string) CommitmentPlansClient {
	return NewCommitmentPlansClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewCommitmentPlansClientWithBaseURI creates an instance of the CommitmentPlansClient client using a custom endpoint.
// Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewCommitmentPlansClientWithBaseURI(baseURI string, subscriptionID string) CommitmentPlansClient {
	return CommitmentPlansClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate update the state of specified commitmentPlans associated with the Cognitive Services account.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// accountName - the name of Cognitive Services account.
// commitmentPlanName - the name of the commitmentPlan associated with the Cognitive Services Account
// commitmentPlan - the commitmentPlan properties.
func (client CommitmentPlansClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, accountName string, commitmentPlanName string, commitmentPlan CommitmentPlan) (result CommitmentPlan, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/CommitmentPlansClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: accountName,
			Constraints: []validation.Constraint{{Target: "accountName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "accountName", Name: validation.MinLength, Rule: 2, Chain: nil},
				{Target: "accountName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9][a-zA-Z0-9_.-]*$`, Chain: nil}}},
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("cognitiveservices.CommitmentPlansClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, accountName, commitmentPlanName, commitmentPlan)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cognitiveservices.CommitmentPlansClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "cognitiveservices.CommitmentPlansClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cognitiveservices.CommitmentPlansClient", "CreateOrUpdate", resp, "Failure responding to request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client CommitmentPlansClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, accountName string, commitmentPlanName string, commitmentPlan CommitmentPlan) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":        autorest.Encode("path", accountName),
		"commitmentPlanName": autorest.Encode("path", commitmentPlanName),
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-10-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	commitmentPlan.SystemData = nil
	commitmentPlan.Etag = nil
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CognitiveServices/accounts/{accountName}/commitmentPlans/{commitmentPlanName}", pathParameters),
		autorest.WithJSON(commitmentPlan),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client CommitmentPlansClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client CommitmentPlansClient) CreateOrUpdateResponder(resp *http.Response) (result CommitmentPlan, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes the specified commitmentPlan associated with the Cognitive Services account.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// accountName - the name of Cognitive Services account.
// commitmentPlanName - the name of the commitmentPlan associated with the Cognitive Services Account
func (client CommitmentPlansClient) Delete(ctx context.Context, resourceGroupName string, accountName string, commitmentPlanName string) (result CommitmentPlansDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/CommitmentPlansClient.Delete")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: accountName,
			Constraints: []validation.Constraint{{Target: "accountName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "accountName", Name: validation.MinLength, Rule: 2, Chain: nil},
				{Target: "accountName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9][a-zA-Z0-9_.-]*$`, Chain: nil}}},
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("cognitiveservices.CommitmentPlansClient", "Delete", err.Error())
	}

	req, err := client.DeletePreparer(ctx, resourceGroupName, accountName, commitmentPlanName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cognitiveservices.CommitmentPlansClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cognitiveservices.CommitmentPlansClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client CommitmentPlansClient) DeletePreparer(ctx context.Context, resourceGroupName string, accountName string, commitmentPlanName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":        autorest.Encode("path", accountName),
		"commitmentPlanName": autorest.Encode("path", commitmentPlanName),
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-10-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CognitiveServices/accounts/{accountName}/commitmentPlans/{commitmentPlanName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client CommitmentPlansClient) DeleteSender(req *http.Request) (future CommitmentPlansDeleteFuture, err error) {
	var resp *http.Response
	future.FutureAPI = &azure.Future{}
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = future.result
	return
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client CommitmentPlansClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets the specified commitmentPlans associated with the Cognitive Services account.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// accountName - the name of Cognitive Services account.
// commitmentPlanName - the name of the commitmentPlan associated with the Cognitive Services Account
func (client CommitmentPlansClient) Get(ctx context.Context, resourceGroupName string, accountName string, commitmentPlanName string) (result CommitmentPlan, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/CommitmentPlansClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: accountName,
			Constraints: []validation.Constraint{{Target: "accountName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "accountName", Name: validation.MinLength, Rule: 2, Chain: nil},
				{Target: "accountName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9][a-zA-Z0-9_.-]*$`, Chain: nil}}},
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("cognitiveservices.CommitmentPlansClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, accountName, commitmentPlanName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cognitiveservices.CommitmentPlansClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "cognitiveservices.CommitmentPlansClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cognitiveservices.CommitmentPlansClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client CommitmentPlansClient) GetPreparer(ctx context.Context, resourceGroupName string, accountName string, commitmentPlanName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":        autorest.Encode("path", accountName),
		"commitmentPlanName": autorest.Encode("path", commitmentPlanName),
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-10-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CognitiveServices/accounts/{accountName}/commitmentPlans/{commitmentPlanName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client CommitmentPlansClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client CommitmentPlansClient) GetResponder(resp *http.Response) (result CommitmentPlan, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List gets the commitmentPlans associated with the Cognitive Services account.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// accountName - the name of Cognitive Services account.
func (client CommitmentPlansClient) List(ctx context.Context, resourceGroupName string, accountName string) (result CommitmentPlanListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/CommitmentPlansClient.List")
		defer func() {
			sc := -1
			if result.cplr.Response.Response != nil {
				sc = result.cplr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: accountName,
			Constraints: []validation.Constraint{{Target: "accountName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "accountName", Name: validation.MinLength, Rule: 2, Chain: nil},
				{Target: "accountName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9][a-zA-Z0-9_.-]*$`, Chain: nil}}},
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("cognitiveservices.CommitmentPlansClient", "List", err.Error())
	}

	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, resourceGroupName, accountName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cognitiveservices.CommitmentPlansClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.cplr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "cognitiveservices.CommitmentPlansClient", "List", resp, "Failure sending request")
		return
	}

	result.cplr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cognitiveservices.CommitmentPlansClient", "List", resp, "Failure responding to request")
		return
	}
	if result.cplr.hasNextLink() && result.cplr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client CommitmentPlansClient) ListPreparer(ctx context.Context, resourceGroupName string, accountName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":       autorest.Encode("path", accountName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-10-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CognitiveServices/accounts/{accountName}/commitmentPlans", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client CommitmentPlansClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client CommitmentPlansClient) ListResponder(resp *http.Response) (result CommitmentPlanListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client CommitmentPlansClient) listNextResults(ctx context.Context, lastResults CommitmentPlanListResult) (result CommitmentPlanListResult, err error) {
	req, err := lastResults.commitmentPlanListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "cognitiveservices.CommitmentPlansClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "cognitiveservices.CommitmentPlansClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cognitiveservices.CommitmentPlansClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client CommitmentPlansClient) ListComplete(ctx context.Context, resourceGroupName string, accountName string) (result CommitmentPlanListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/CommitmentPlansClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, resourceGroupName, accountName)
	return
}
