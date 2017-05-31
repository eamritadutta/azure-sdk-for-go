package network

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

// PublicIPAddressesClient is the the Microsoft Azure Network management API
// provides a RESTful set of web services that interact with Microsoft Azure
// Networks service to manage your network resrources. The API has entities
// that capture the relationship between an end user and the Microsoft Azure
// Networks service.
type PublicIPAddressesClient struct {
	ManagementClient
}

// NewPublicIPAddressesClient creates an instance of the
// PublicIPAddressesClient client.
func NewPublicIPAddressesClient(subscriptionID string) PublicIPAddressesClient {
	return NewPublicIPAddressesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewPublicIPAddressesClientWithBaseURI creates an instance of the
// PublicIPAddressesClient client.
func NewPublicIPAddressesClientWithBaseURI(baseURI string, subscriptionID string) PublicIPAddressesClient {
	return PublicIPAddressesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate the Put PublicIPAddress operation creates/updates a
// stable/dynamic PublicIP address This method may poll for completion. Polling
// can be canceled by passing the cancel channel argument. The channel will be
// used to cancel polling and any outstanding HTTP requests.
//
// resourceGroupName is the name of the resource group. publicIPAddressName is
// the name of the publicIpAddress. parameters is parameters supplied to the
// create/update PublicIPAddress operation
func (client PublicIPAddressesClient) CreateOrUpdate(resourceGroupName string, publicIPAddressName string, parameters PublicIPAddress, cancel <-chan struct{}) (<-chan PublicIPAddress, <-chan error) {
	resultChan := make(chan PublicIPAddress, 1)
	errChan := make(chan error, 1)
	if err := validation.Validate([]validation.Validation{
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.PublicIPAddressPropertiesFormat", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "parameters.PublicIPAddressPropertiesFormat.IPConfiguration", Name: validation.Null, Rule: false,
					Chain: []validation.Constraint{{Target: "parameters.PublicIPAddressPropertiesFormat.IPConfiguration.IPConfigurationPropertiesFormat", Name: validation.Null, Rule: false,
						Chain: []validation.Constraint{{Target: "parameters.PublicIPAddressPropertiesFormat.IPConfiguration.IPConfigurationPropertiesFormat.PublicIPAddress", Name: validation.Null, Rule: false, Chain: nil}}},
					}},
				}}}}}); err != nil {
		errChan <- validation.NewErrorWithValidationError(err, "network.PublicIPAddressesClient", "CreateOrUpdate")
		close(errChan)
		close(resultChan)
		return resultChan, errChan
	}

	go func() {
		var err error
		var result PublicIPAddress
		defer func() {
			resultChan <- result
			errChan <- err
			close(resultChan)
			close(errChan)
		}()
		req, err := client.CreateOrUpdatePreparer(resourceGroupName, publicIPAddressName, parameters, cancel)
		if err != nil {
			err = autorest.NewErrorWithError(err, "network.PublicIPAddressesClient", "CreateOrUpdate", nil, "Failure preparing request")
			return
		}

		resp, err := client.CreateOrUpdateSender(req)
		if err != nil {
			result.Response = autorest.Response{Response: resp}
			err = autorest.NewErrorWithError(err, "network.PublicIPAddressesClient", "CreateOrUpdate", resp, "Failure sending request")
			return
		}

		result, err = client.CreateOrUpdateResponder(resp)
		if err != nil {
			err = autorest.NewErrorWithError(err, "network.PublicIPAddressesClient", "CreateOrUpdate", resp, "Failure responding to request")
		}
	}()
	return resultChan, errChan
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client PublicIPAddressesClient) CreateOrUpdatePreparer(resourceGroupName string, publicIPAddressName string, parameters PublicIPAddress, cancel <-chan struct{}) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"publicIpAddressName": autorest.Encode("path", publicIPAddressName),
		"resourceGroupName":   autorest.Encode("path", resourceGroupName),
		"subscriptionId":      autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-03-30"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/publicIPAddresses/{publicIpAddressName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{Cancel: cancel})
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client PublicIPAddressesClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoPollForAsynchronous(client.PollingDelay))
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client PublicIPAddressesClient) CreateOrUpdateResponder(resp *http.Response) (result PublicIPAddress, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusCreated, http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete the delete publicIpAddress operation deletes the specified
// publicIpAddress. This method may poll for completion. Polling can be
// canceled by passing the cancel channel argument. The channel will be used to
// cancel polling and any outstanding HTTP requests.
//
// resourceGroupName is the name of the resource group. publicIPAddressName is
// the name of the subnet.
func (client PublicIPAddressesClient) Delete(resourceGroupName string, publicIPAddressName string, cancel <-chan struct{}) (<-chan autorest.Response, <-chan error) {
	resultChan := make(chan autorest.Response, 1)
	errChan := make(chan error, 1)
	go func() {
		var err error
		var result autorest.Response
		defer func() {
			resultChan <- result
			errChan <- err
			close(resultChan)
			close(errChan)
		}()
		req, err := client.DeletePreparer(resourceGroupName, publicIPAddressName, cancel)
		if err != nil {
			err = autorest.NewErrorWithError(err, "network.PublicIPAddressesClient", "Delete", nil, "Failure preparing request")
			return
		}

		resp, err := client.DeleteSender(req)
		if err != nil {
			result.Response = resp
			err = autorest.NewErrorWithError(err, "network.PublicIPAddressesClient", "Delete", resp, "Failure sending request")
			return
		}

		result, err = client.DeleteResponder(resp)
		if err != nil {
			err = autorest.NewErrorWithError(err, "network.PublicIPAddressesClient", "Delete", resp, "Failure responding to request")
		}
	}()
	return resultChan, errChan
}

// DeletePreparer prepares the Delete request.
func (client PublicIPAddressesClient) DeletePreparer(resourceGroupName string, publicIPAddressName string, cancel <-chan struct{}) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"publicIpAddressName": autorest.Encode("path", publicIPAddressName),
		"resourceGroupName":   autorest.Encode("path", resourceGroupName),
		"subscriptionId":      autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-03-30"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/publicIPAddresses/{publicIpAddressName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{Cancel: cancel})
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client PublicIPAddressesClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoPollForAsynchronous(client.PollingDelay))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client PublicIPAddressesClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusNoContent, http.StatusAccepted, http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get the Get publicIpAddress operation retreives information about the
// specified pubicIpAddress
//
// resourceGroupName is the name of the resource group. publicIPAddressName is
// the name of the subnet. expand is expand references resources.
func (client PublicIPAddressesClient) Get(resourceGroupName string, publicIPAddressName string, expand string) (result PublicIPAddress, err error) {
	req, err := client.GetPreparer(resourceGroupName, publicIPAddressName, expand)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.PublicIPAddressesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.PublicIPAddressesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.PublicIPAddressesClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client PublicIPAddressesClient) GetPreparer(resourceGroupName string, publicIPAddressName string, expand string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"publicIpAddressName": autorest.Encode("path", publicIPAddressName),
		"resourceGroupName":   autorest.Encode("path", resourceGroupName),
		"subscriptionId":      autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-03-30"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(expand) > 0 {
		queryParameters["$expand"] = autorest.Encode("query", expand)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/publicIPAddresses/{publicIpAddressName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client PublicIPAddressesClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client PublicIPAddressesClient) GetResponder(resp *http.Response) (result PublicIPAddress, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List the List publicIpAddress opertion retrieves all the publicIpAddresses
// in a resource group.
//
// resourceGroupName is the name of the resource group.
func (client PublicIPAddressesClient) List(resourceGroupName string) (result PublicIPAddressListResult, err error) {
	req, err := client.ListPreparer(resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.PublicIPAddressesClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.PublicIPAddressesClient", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.PublicIPAddressesClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client PublicIPAddressesClient) ListPreparer(resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-03-30"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/publicIPAddresses", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client PublicIPAddressesClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client PublicIPAddressesClient) ListResponder(resp *http.Response) (result PublicIPAddressListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListNextResults retrieves the next set of results, if any.
func (client PublicIPAddressesClient) ListNextResults(lastResults PublicIPAddressListResult) (result PublicIPAddressListResult, err error) {
	req, err := lastResults.PublicIPAddressListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network.PublicIPAddressesClient", "List", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network.PublicIPAddressesClient", "List", resp, "Failure sending next results request")
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.PublicIPAddressesClient", "List", resp, "Failure responding to next results request")
	}

	return
}

// ListAll the List publicIpAddress opertion retrieves all the
// publicIpAddresses in a subscription.
func (client PublicIPAddressesClient) ListAll() (result PublicIPAddressListResult, err error) {
	req, err := client.ListAllPreparer()
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.PublicIPAddressesClient", "ListAll", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListAllSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.PublicIPAddressesClient", "ListAll", resp, "Failure sending request")
		return
	}

	result, err = client.ListAllResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.PublicIPAddressesClient", "ListAll", resp, "Failure responding to request")
	}

	return
}

// ListAllPreparer prepares the ListAll request.
func (client PublicIPAddressesClient) ListAllPreparer() (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-03-30"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Network/publicIPAddresses", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListAllSender sends the ListAll request. The method will close the
// http.Response Body if it receives an error.
func (client PublicIPAddressesClient) ListAllSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListAllResponder handles the response to the ListAll request. The method always
// closes the http.Response Body.
func (client PublicIPAddressesClient) ListAllResponder(resp *http.Response) (result PublicIPAddressListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListAllNextResults retrieves the next set of results, if any.
func (client PublicIPAddressesClient) ListAllNextResults(lastResults PublicIPAddressListResult) (result PublicIPAddressListResult, err error) {
	req, err := lastResults.PublicIPAddressListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network.PublicIPAddressesClient", "ListAll", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListAllSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network.PublicIPAddressesClient", "ListAll", resp, "Failure sending next results request")
	}

	result, err = client.ListAllResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.PublicIPAddressesClient", "ListAll", resp, "Failure responding to next results request")
	}

	return
}
