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
	"net/http"
)

// VirtualNetworkGatewaysClient is the the Microsoft Azure Network management
// API provides a RESTful set of web services that interact with Microsoft
// Azure Networks service to manage your network resrources. The API has
// entities that capture the relationship between an end user and the Microsoft
// Azure Networks service.
type VirtualNetworkGatewaysClient struct {
	ManagementClient
}

// NewVirtualNetworkGatewaysClient creates an instance of the
// VirtualNetworkGatewaysClient client.
func NewVirtualNetworkGatewaysClient(subscriptionID string) VirtualNetworkGatewaysClient {
	return NewVirtualNetworkGatewaysClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewVirtualNetworkGatewaysClientWithBaseURI creates an instance of the
// VirtualNetworkGatewaysClient client.
func NewVirtualNetworkGatewaysClientWithBaseURI(baseURI string, subscriptionID string) VirtualNetworkGatewaysClient {
	return VirtualNetworkGatewaysClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate the Put VirtualNetworkGateway operation creates/updates a
// virtual network gateway in the specified resource group through Network
// resource provider. This method may poll for completion. Polling can be
// canceled by passing the cancel channel argument. The channel will be used to
// cancel polling and any outstanding HTTP requests.
//
// resourceGroupName is the name of the resource group.
// virtualNetworkGatewayName is the name of the virtual network gateway.
// parameters is parameters supplied to the Begin Create or update Virtual
// Network Gateway operation through Network resource provider.
func (client VirtualNetworkGatewaysClient) CreateOrUpdate(resourceGroupName string, virtualNetworkGatewayName string, parameters VirtualNetworkGateway, cancel <-chan struct{}) (<-chan VirtualNetworkGateway, <-chan error) {
	resultChan := make(chan VirtualNetworkGateway, 1)
	errChan := make(chan error, 1)
	go func() {
		var err error
		var result VirtualNetworkGateway
		defer func() {
			resultChan <- result
			errChan <- err
			close(resultChan)
			close(errChan)
		}()
		req, err := client.CreateOrUpdatePreparer(resourceGroupName, virtualNetworkGatewayName, parameters, cancel)
		if err != nil {
			err = autorest.NewErrorWithError(err, "network.VirtualNetworkGatewaysClient", "CreateOrUpdate", nil, "Failure preparing request")
			return
		}

		resp, err := client.CreateOrUpdateSender(req)
		if err != nil {
			result.Response = autorest.Response{Response: resp}
			err = autorest.NewErrorWithError(err, "network.VirtualNetworkGatewaysClient", "CreateOrUpdate", resp, "Failure sending request")
			return
		}

		result, err = client.CreateOrUpdateResponder(resp)
		if err != nil {
			err = autorest.NewErrorWithError(err, "network.VirtualNetworkGatewaysClient", "CreateOrUpdate", resp, "Failure responding to request")
		}
	}()
	return resultChan, errChan
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client VirtualNetworkGatewaysClient) CreateOrUpdatePreparer(resourceGroupName string, virtualNetworkGatewayName string, parameters VirtualNetworkGateway, cancel <-chan struct{}) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":         autorest.Encode("path", resourceGroupName),
		"subscriptionId":            autorest.Encode("path", client.SubscriptionID),
		"virtualNetworkGatewayName": autorest.Encode("path", virtualNetworkGatewayName),
	}

	const APIVersion = "2016-03-30"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworkGateways/{virtualNetworkGatewayName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{Cancel: cancel})
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualNetworkGatewaysClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoPollForAsynchronous(client.PollingDelay))
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client VirtualNetworkGatewaysClient) CreateOrUpdateResponder(resp *http.Response) (result VirtualNetworkGateway, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete the Delete VirtualNetworkGateway operation deletes the specifed
// virtual network Gateway through Network resource provider. This method may
// poll for completion. Polling can be canceled by passing the cancel channel
// argument. The channel will be used to cancel polling and any outstanding
// HTTP requests.
//
// resourceGroupName is the name of the resource group.
// virtualNetworkGatewayName is the name of the virtual network gateway.
func (client VirtualNetworkGatewaysClient) Delete(resourceGroupName string, virtualNetworkGatewayName string, cancel <-chan struct{}) (<-chan autorest.Response, <-chan error) {
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
		req, err := client.DeletePreparer(resourceGroupName, virtualNetworkGatewayName, cancel)
		if err != nil {
			err = autorest.NewErrorWithError(err, "network.VirtualNetworkGatewaysClient", "Delete", nil, "Failure preparing request")
			return
		}

		resp, err := client.DeleteSender(req)
		if err != nil {
			result.Response = resp
			err = autorest.NewErrorWithError(err, "network.VirtualNetworkGatewaysClient", "Delete", resp, "Failure sending request")
			return
		}

		result, err = client.DeleteResponder(resp)
		if err != nil {
			err = autorest.NewErrorWithError(err, "network.VirtualNetworkGatewaysClient", "Delete", resp, "Failure responding to request")
		}
	}()
	return resultChan, errChan
}

// DeletePreparer prepares the Delete request.
func (client VirtualNetworkGatewaysClient) DeletePreparer(resourceGroupName string, virtualNetworkGatewayName string, cancel <-chan struct{}) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":         autorest.Encode("path", resourceGroupName),
		"subscriptionId":            autorest.Encode("path", client.SubscriptionID),
		"virtualNetworkGatewayName": autorest.Encode("path", virtualNetworkGatewayName),
	}

	const APIVersion = "2016-03-30"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworkGateways/{virtualNetworkGatewayName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{Cancel: cancel})
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualNetworkGatewaysClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoPollForAsynchronous(client.PollingDelay))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client VirtualNetworkGatewaysClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusNoContent, http.StatusAccepted, http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Generatevpnclientpackage the Generatevpnclientpackage operation generates
// Vpn client package for P2S client of the virtual network gateway in the
// specified resource group through Network resource provider.
//
// resourceGroupName is the name of the resource group.
// virtualNetworkGatewayName is the name of the virtual network gateway.
// parameters is parameters supplied to the Begin Generating  Virtual Network
// Gateway Vpn client package operation through Network resource provider.
func (client VirtualNetworkGatewaysClient) Generatevpnclientpackage(resourceGroupName string, virtualNetworkGatewayName string, parameters VpnClientParameters) (result String, err error) {
	req, err := client.GeneratevpnclientpackagePreparer(resourceGroupName, virtualNetworkGatewayName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VirtualNetworkGatewaysClient", "Generatevpnclientpackage", nil, "Failure preparing request")
		return
	}

	resp, err := client.GeneratevpnclientpackageSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.VirtualNetworkGatewaysClient", "Generatevpnclientpackage", resp, "Failure sending request")
		return
	}

	result, err = client.GeneratevpnclientpackageResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VirtualNetworkGatewaysClient", "Generatevpnclientpackage", resp, "Failure responding to request")
	}

	return
}

// GeneratevpnclientpackagePreparer prepares the Generatevpnclientpackage request.
func (client VirtualNetworkGatewaysClient) GeneratevpnclientpackagePreparer(resourceGroupName string, virtualNetworkGatewayName string, parameters VpnClientParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":         autorest.Encode("path", resourceGroupName),
		"subscriptionId":            autorest.Encode("path", client.SubscriptionID),
		"virtualNetworkGatewayName": autorest.Encode("path", virtualNetworkGatewayName),
	}

	const APIVersion = "2016-03-30"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworkGateways/{virtualNetworkGatewayName}/generatevpnclientpackage", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GeneratevpnclientpackageSender sends the Generatevpnclientpackage request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualNetworkGatewaysClient) GeneratevpnclientpackageSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GeneratevpnclientpackageResponder handles the response to the Generatevpnclientpackage request. The method always
// closes the http.Response Body.
func (client VirtualNetworkGatewaysClient) GeneratevpnclientpackageResponder(resp *http.Response) (result String, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Get the Get VirtualNetworkGateway operation retrieves information about the
// specified virtual network gateway through Network resource provider.
//
// resourceGroupName is the name of the resource group.
// virtualNetworkGatewayName is the name of the virtual network gateway.
func (client VirtualNetworkGatewaysClient) Get(resourceGroupName string, virtualNetworkGatewayName string) (result VirtualNetworkGateway, err error) {
	req, err := client.GetPreparer(resourceGroupName, virtualNetworkGatewayName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VirtualNetworkGatewaysClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.VirtualNetworkGatewaysClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VirtualNetworkGatewaysClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client VirtualNetworkGatewaysClient) GetPreparer(resourceGroupName string, virtualNetworkGatewayName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":         autorest.Encode("path", resourceGroupName),
		"subscriptionId":            autorest.Encode("path", client.SubscriptionID),
		"virtualNetworkGatewayName": autorest.Encode("path", virtualNetworkGatewayName),
	}

	const APIVersion = "2016-03-30"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworkGateways/{virtualNetworkGatewayName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualNetworkGatewaysClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client VirtualNetworkGatewaysClient) GetResponder(resp *http.Response) (result VirtualNetworkGateway, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List the List VirtualNetworkGateways opertion retrieves all the virtual
// network gateways stored.
//
// resourceGroupName is the name of the resource group.
func (client VirtualNetworkGatewaysClient) List(resourceGroupName string) (result VirtualNetworkGatewayListResult, err error) {
	req, err := client.ListPreparer(resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VirtualNetworkGatewaysClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.VirtualNetworkGatewaysClient", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VirtualNetworkGatewaysClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client VirtualNetworkGatewaysClient) ListPreparer(resourceGroupName string) (*http.Request, error) {
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
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworkGateways", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualNetworkGatewaysClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client VirtualNetworkGatewaysClient) ListResponder(resp *http.Response) (result VirtualNetworkGatewayListResult, err error) {
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
func (client VirtualNetworkGatewaysClient) ListNextResults(lastResults VirtualNetworkGatewayListResult) (result VirtualNetworkGatewayListResult, err error) {
	req, err := lastResults.VirtualNetworkGatewayListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network.VirtualNetworkGatewaysClient", "List", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network.VirtualNetworkGatewaysClient", "List", resp, "Failure sending next results request")
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VirtualNetworkGatewaysClient", "List", resp, "Failure responding to next results request")
	}

	return
}

// Reset the Reset VirtualNetworkGateway operation resets the primary of the
// virtual network gateway in the specified resource group through Network
// resource provider. This method may poll for completion. Polling can be
// canceled by passing the cancel channel argument. The channel will be used to
// cancel polling and any outstanding HTTP requests.
//
// resourceGroupName is the name of the resource group.
// virtualNetworkGatewayName is the name of the virtual network gateway.
// parameters is parameters supplied to the Begin Reset Virtual Network Gateway
// operation through Network resource provider.
func (client VirtualNetworkGatewaysClient) Reset(resourceGroupName string, virtualNetworkGatewayName string, parameters VirtualNetworkGateway, cancel <-chan struct{}) (<-chan VirtualNetworkGateway, <-chan error) {
	resultChan := make(chan VirtualNetworkGateway, 1)
	errChan := make(chan error, 1)
	go func() {
		var err error
		var result VirtualNetworkGateway
		defer func() {
			resultChan <- result
			errChan <- err
			close(resultChan)
			close(errChan)
		}()
		req, err := client.ResetPreparer(resourceGroupName, virtualNetworkGatewayName, parameters, cancel)
		if err != nil {
			err = autorest.NewErrorWithError(err, "network.VirtualNetworkGatewaysClient", "Reset", nil, "Failure preparing request")
			return
		}

		resp, err := client.ResetSender(req)
		if err != nil {
			result.Response = autorest.Response{Response: resp}
			err = autorest.NewErrorWithError(err, "network.VirtualNetworkGatewaysClient", "Reset", resp, "Failure sending request")
			return
		}

		result, err = client.ResetResponder(resp)
		if err != nil {
			err = autorest.NewErrorWithError(err, "network.VirtualNetworkGatewaysClient", "Reset", resp, "Failure responding to request")
		}
	}()
	return resultChan, errChan
}

// ResetPreparer prepares the Reset request.
func (client VirtualNetworkGatewaysClient) ResetPreparer(resourceGroupName string, virtualNetworkGatewayName string, parameters VirtualNetworkGateway, cancel <-chan struct{}) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":         autorest.Encode("path", resourceGroupName),
		"subscriptionId":            autorest.Encode("path", client.SubscriptionID),
		"virtualNetworkGatewayName": autorest.Encode("path", virtualNetworkGatewayName),
	}

	const APIVersion = "2016-03-30"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworkGateways/{virtualNetworkGatewayName}/reset", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{Cancel: cancel})
}

// ResetSender sends the Reset request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualNetworkGatewaysClient) ResetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoPollForAsynchronous(client.PollingDelay))
}

// ResetResponder handles the response to the Reset request. The method always
// closes the http.Response Body.
func (client VirtualNetworkGatewaysClient) ResetResponder(resp *http.Response) (result VirtualNetworkGateway, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusAccepted, http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
