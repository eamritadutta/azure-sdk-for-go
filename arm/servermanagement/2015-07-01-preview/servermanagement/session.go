package servermanagement

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator 1.0.1.0
// Changes may cause incorrect behavior and will be lost if the code is
// regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"net/http"
)

// SessionClient is the rEST API for Azure Server Management Service.
type SessionClient struct {
	ManagementClient
}

// NewSessionClient creates an instance of the SessionClient client.
func NewSessionClient(subscriptionID string) SessionClient {
	return NewSessionClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewSessionClientWithBaseURI creates an instance of the SessionClient client.
func NewSessionClientWithBaseURI(baseURI string, subscriptionID string) SessionClient {
	return SessionClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Create creates a session for a node. This method may poll for completion.
// Polling can be canceled by passing the cancel channel argument. The channel
// will be used to cancel polling and any outstanding HTTP requests.
//
// resourceGroupName is the resource group name uniquely identifies the
// resource group within the user subscriptionId. nodeName is the node name
// (256 characters maximum). session is the sessionId from the user.
// sessionParameters is parameters supplied to the CreateOrUpdate operation.
func (client SessionClient) Create(resourceGroupName string, nodeName string, session string, sessionParameters SessionParameters, cancel <-chan struct{}) (<-chan SessionResource, <-chan error) {
	resultChan := make(chan SessionResource, 1)
	errChan := make(chan error, 1)
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `[a-zA-Z0-9]+`, Chain: nil}}},
		{TargetValue: nodeName,
			Constraints: []validation.Constraint{{Target: "nodeName", Name: validation.MaxLength, Rule: 256, Chain: nil},
				{Target: "nodeName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "nodeName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9][a-zA-Z0-9_.-]*$`, Chain: nil}}}}); err != nil {
		errChan <- validation.NewErrorWithValidationError(err, "servermanagement.SessionClient", "Create")
		close(errChan)
		close(resultChan)
		return resultChan, errChan
	}

	go func() {
		var err error
		var result SessionResource
		defer func() {
			resultChan <- result
			errChan <- err
			close(resultChan)
			close(errChan)
		}()
		req, err := client.CreatePreparer(resourceGroupName, nodeName, session, sessionParameters, cancel)
		if err != nil {
			err = autorest.NewErrorWithError(err, "servermanagement.SessionClient", "Create", nil, "Failure preparing request")
			return
		}

		resp, err := client.CreateSender(req)
		if err != nil {
			result.Response = autorest.Response{Response: resp}
			err = autorest.NewErrorWithError(err, "servermanagement.SessionClient", "Create", resp, "Failure sending request")
			return
		}

		result, err = client.CreateResponder(resp)
		if err != nil {
			err = autorest.NewErrorWithError(err, "servermanagement.SessionClient", "Create", resp, "Failure responding to request")
		}
	}()
	return resultChan, errChan
}

// CreatePreparer prepares the Create request.
func (client SessionClient) CreatePreparer(resourceGroupName string, nodeName string, session string, sessionParameters SessionParameters, cancel <-chan struct{}) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"nodeName":          autorest.Encode("path", nodeName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"session":           autorest.Encode("path", session),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-07-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServerManagement/nodes/{nodeName}/sessions/{session}", pathParameters),
		autorest.WithJSON(sessionParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{Cancel: cancel})
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client SessionClient) CreateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoPollForAsynchronous(client.PollingDelay))
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client SessionClient) CreateResponder(resp *http.Response) (result SessionResource, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes a session for a node.
//
// resourceGroupName is the resource group name uniquely identifies the
// resource group within the user subscriptionId. nodeName is the node name
// (256 characters maximum). session is the sessionId from the user.
func (client SessionClient) Delete(resourceGroupName string, nodeName string, session string) (result autorest.Response, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `[a-zA-Z0-9]+`, Chain: nil}}},
		{TargetValue: nodeName,
			Constraints: []validation.Constraint{{Target: "nodeName", Name: validation.MaxLength, Rule: 256, Chain: nil},
				{Target: "nodeName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "nodeName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9][a-zA-Z0-9_.-]*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "servermanagement.SessionClient", "Delete")
	}

	req, err := client.DeletePreparer(resourceGroupName, nodeName, session)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servermanagement.SessionClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "servermanagement.SessionClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servermanagement.SessionClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client SessionClient) DeletePreparer(resourceGroupName string, nodeName string, session string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"nodeName":          autorest.Encode("path", nodeName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"session":           autorest.Encode("path", session),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-07-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServerManagement/nodes/{nodeName}/sessions/{session}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client SessionClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client SessionClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets a session for a node.
//
// resourceGroupName is the resource group name uniquely identifies the
// resource group within the user subscriptionId. nodeName is the node name
// (256 characters maximum). session is the sessionId from the user.
func (client SessionClient) Get(resourceGroupName string, nodeName string, session string) (result SessionResource, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `[a-zA-Z0-9]+`, Chain: nil}}},
		{TargetValue: nodeName,
			Constraints: []validation.Constraint{{Target: "nodeName", Name: validation.MaxLength, Rule: 256, Chain: nil},
				{Target: "nodeName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "nodeName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9][a-zA-Z0-9_.-]*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "servermanagement.SessionClient", "Get")
	}

	req, err := client.GetPreparer(resourceGroupName, nodeName, session)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servermanagement.SessionClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "servermanagement.SessionClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servermanagement.SessionClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client SessionClient) GetPreparer(resourceGroupName string, nodeName string, session string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"nodeName":          autorest.Encode("path", nodeName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"session":           autorest.Encode("path", session),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-07-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServerManagement/nodes/{nodeName}/sessions/{session}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client SessionClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client SessionClient) GetResponder(resp *http.Response) (result SessionResource, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
