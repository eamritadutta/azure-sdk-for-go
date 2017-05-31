package service

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for
// license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is
// regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"net/http"
)

// ReplicationStorageClassificationMappingsClient is the client for the
// ReplicationStorageClassificationMappings methods of the Service service.
type ReplicationStorageClassificationMappingsClient struct {
	ManagementClient
}

// NewReplicationStorageClassificationMappingsClient creates an instance of the
// ReplicationStorageClassificationMappingsClient client.
func NewReplicationStorageClassificationMappingsClient(subscriptionID string, resourceGroupName string, resourceName string) ReplicationStorageClassificationMappingsClient {
	return NewReplicationStorageClassificationMappingsClientWithBaseURI(DefaultBaseURI, subscriptionID, resourceGroupName, resourceName)
}

// NewReplicationStorageClassificationMappingsClientWithBaseURI creates an
// instance of the ReplicationStorageClassificationMappingsClient client.
func NewReplicationStorageClassificationMappingsClientWithBaseURI(baseURI string, subscriptionID string, resourceGroupName string, resourceName string) ReplicationStorageClassificationMappingsClient {
	return ReplicationStorageClassificationMappingsClient{NewWithBaseURI(baseURI, subscriptionID, resourceGroupName, resourceName)}
}

// Create the operation to create a storage classification mapping. This method
// may poll for completion. Polling can be canceled by passing the cancel
// channel argument. The channel will be used to cancel polling and any
// outstanding HTTP requests.
//
// fabricName is fabric name. storageClassificationName is storage
// classification name. storageClassificationMappingName is storage
// classification mapping name. pairingInput is pairing input.
func (client ReplicationStorageClassificationMappingsClient) Create(fabricName string, storageClassificationName string, storageClassificationMappingName string, pairingInput StorageClassificationMappingInput, cancel <-chan struct{}) (<-chan StorageClassificationMapping, <-chan error) {
	resultChan := make(chan StorageClassificationMapping, 1)
	errChan := make(chan error, 1)
	go func() {
		var err error
		var result StorageClassificationMapping
		defer func() {
			resultChan <- result
			errChan <- err
			close(resultChan)
			close(errChan)
		}()
		req, err := client.CreatePreparer(fabricName, storageClassificationName, storageClassificationMappingName, pairingInput, cancel)
		if err != nil {
			err = autorest.NewErrorWithError(err, "service.ReplicationStorageClassificationMappingsClient", "Create", nil, "Failure preparing request")
			return
		}

		resp, err := client.CreateSender(req)
		if err != nil {
			result.Response = autorest.Response{Response: resp}
			err = autorest.NewErrorWithError(err, "service.ReplicationStorageClassificationMappingsClient", "Create", resp, "Failure sending request")
			return
		}

		result, err = client.CreateResponder(resp)
		if err != nil {
			err = autorest.NewErrorWithError(err, "service.ReplicationStorageClassificationMappingsClient", "Create", resp, "Failure responding to request")
		}
	}()
	return resultChan, errChan
}

// CreatePreparer prepares the Create request.
func (client ReplicationStorageClassificationMappingsClient) CreatePreparer(fabricName string, storageClassificationName string, storageClassificationMappingName string, pairingInput StorageClassificationMappingInput, cancel <-chan struct{}) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"fabricName":                       autorest.Encode("path", fabricName),
		"resourceGroupName":                autorest.Encode("path", client.ResourceGroupName),
		"resourceName":                     autorest.Encode("path", client.ResourceName),
		"storageClassificationMappingName": autorest.Encode("path", storageClassificationMappingName),
		"storageClassificationName":        autorest.Encode("path", storageClassificationName),
		"subscriptionId":                   autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-08-10"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationStorageClassifications/{storageClassificationName}/replicationStorageClassificationMappings/{storageClassificationMappingName}", pathParameters),
		autorest.WithJSON(pairingInput),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{Cancel: cancel})
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client ReplicationStorageClassificationMappingsClient) CreateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoPollForAsynchronous(client.PollingDelay))
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client ReplicationStorageClassificationMappingsClient) CreateResponder(resp *http.Response) (result StorageClassificationMapping, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete the operation to delete a storage classification mapping. This method
// may poll for completion. Polling can be canceled by passing the cancel
// channel argument. The channel will be used to cancel polling and any
// outstanding HTTP requests.
//
// fabricName is fabric name. storageClassificationName is storage
// classification name. storageClassificationMappingName is storage
// classification mapping name.
func (client ReplicationStorageClassificationMappingsClient) Delete(fabricName string, storageClassificationName string, storageClassificationMappingName string, cancel <-chan struct{}) (<-chan autorest.Response, <-chan error) {
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
		req, err := client.DeletePreparer(fabricName, storageClassificationName, storageClassificationMappingName, cancel)
		if err != nil {
			err = autorest.NewErrorWithError(err, "service.ReplicationStorageClassificationMappingsClient", "Delete", nil, "Failure preparing request")
			return
		}

		resp, err := client.DeleteSender(req)
		if err != nil {
			result.Response = resp
			err = autorest.NewErrorWithError(err, "service.ReplicationStorageClassificationMappingsClient", "Delete", resp, "Failure sending request")
			return
		}

		result, err = client.DeleteResponder(resp)
		if err != nil {
			err = autorest.NewErrorWithError(err, "service.ReplicationStorageClassificationMappingsClient", "Delete", resp, "Failure responding to request")
		}
	}()
	return resultChan, errChan
}

// DeletePreparer prepares the Delete request.
func (client ReplicationStorageClassificationMappingsClient) DeletePreparer(fabricName string, storageClassificationName string, storageClassificationMappingName string, cancel <-chan struct{}) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"fabricName":                       autorest.Encode("path", fabricName),
		"resourceGroupName":                autorest.Encode("path", client.ResourceGroupName),
		"resourceName":                     autorest.Encode("path", client.ResourceName),
		"storageClassificationMappingName": autorest.Encode("path", storageClassificationMappingName),
		"storageClassificationName":        autorest.Encode("path", storageClassificationName),
		"subscriptionId":                   autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-08-10"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationStorageClassifications/{storageClassificationName}/replicationStorageClassificationMappings/{storageClassificationMappingName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{Cancel: cancel})
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client ReplicationStorageClassificationMappingsClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoPollForAsynchronous(client.PollingDelay))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client ReplicationStorageClassificationMappingsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets the details of the specified storage classification mapping.
//
// fabricName is fabric name. storageClassificationName is storage
// classification name. storageClassificationMappingName is storage
// classification mapping name.
func (client ReplicationStorageClassificationMappingsClient) Get(fabricName string, storageClassificationName string, storageClassificationMappingName string) (result StorageClassificationMapping, err error) {
	req, err := client.GetPreparer(fabricName, storageClassificationName, storageClassificationMappingName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "service.ReplicationStorageClassificationMappingsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "service.ReplicationStorageClassificationMappingsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "service.ReplicationStorageClassificationMappingsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client ReplicationStorageClassificationMappingsClient) GetPreparer(fabricName string, storageClassificationName string, storageClassificationMappingName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"fabricName":                       autorest.Encode("path", fabricName),
		"resourceGroupName":                autorest.Encode("path", client.ResourceGroupName),
		"resourceName":                     autorest.Encode("path", client.ResourceName),
		"storageClassificationMappingName": autorest.Encode("path", storageClassificationMappingName),
		"storageClassificationName":        autorest.Encode("path", storageClassificationName),
		"subscriptionId":                   autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-08-10"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationStorageClassifications/{storageClassificationName}/replicationStorageClassificationMappings/{storageClassificationMappingName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ReplicationStorageClassificationMappingsClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ReplicationStorageClassificationMappingsClient) GetResponder(resp *http.Response) (result StorageClassificationMapping, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List lists the storage classification mappings in the vault.
func (client ReplicationStorageClassificationMappingsClient) List() (result StorageClassificationMappingCollection, err error) {
	req, err := client.ListPreparer()
	if err != nil {
		err = autorest.NewErrorWithError(err, "service.ReplicationStorageClassificationMappingsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "service.ReplicationStorageClassificationMappingsClient", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "service.ReplicationStorageClassificationMappingsClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client ReplicationStorageClassificationMappingsClient) ListPreparer() (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", client.ResourceGroupName),
		"resourceName":      autorest.Encode("path", client.ResourceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-08-10"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationStorageClassificationMappings", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client ReplicationStorageClassificationMappingsClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client ReplicationStorageClassificationMappingsClient) ListResponder(resp *http.Response) (result StorageClassificationMappingCollection, err error) {
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
func (client ReplicationStorageClassificationMappingsClient) ListNextResults(lastResults StorageClassificationMappingCollection) (result StorageClassificationMappingCollection, err error) {
	req, err := lastResults.StorageClassificationMappingCollectionPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "service.ReplicationStorageClassificationMappingsClient", "List", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "service.ReplicationStorageClassificationMappingsClient", "List", resp, "Failure sending next results request")
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "service.ReplicationStorageClassificationMappingsClient", "List", resp, "Failure responding to next results request")
	}

	return
}

// ListByReplicationStorageClassifications lists the storage classification
// mappings for the fabric.
//
// fabricName is fabric name. storageClassificationName is storage
// classfication name.
func (client ReplicationStorageClassificationMappingsClient) ListByReplicationStorageClassifications(fabricName string, storageClassificationName string) (result StorageClassificationMappingCollection, err error) {
	req, err := client.ListByReplicationStorageClassificationsPreparer(fabricName, storageClassificationName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "service.ReplicationStorageClassificationMappingsClient", "ListByReplicationStorageClassifications", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByReplicationStorageClassificationsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "service.ReplicationStorageClassificationMappingsClient", "ListByReplicationStorageClassifications", resp, "Failure sending request")
		return
	}

	result, err = client.ListByReplicationStorageClassificationsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "service.ReplicationStorageClassificationMappingsClient", "ListByReplicationStorageClassifications", resp, "Failure responding to request")
	}

	return
}

// ListByReplicationStorageClassificationsPreparer prepares the ListByReplicationStorageClassifications request.
func (client ReplicationStorageClassificationMappingsClient) ListByReplicationStorageClassificationsPreparer(fabricName string, storageClassificationName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"fabricName":                autorest.Encode("path", fabricName),
		"resourceGroupName":         autorest.Encode("path", client.ResourceGroupName),
		"resourceName":              autorest.Encode("path", client.ResourceName),
		"storageClassificationName": autorest.Encode("path", storageClassificationName),
		"subscriptionId":            autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-08-10"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationStorageClassifications/{storageClassificationName}/replicationStorageClassificationMappings", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListByReplicationStorageClassificationsSender sends the ListByReplicationStorageClassifications request. The method will close the
// http.Response Body if it receives an error.
func (client ReplicationStorageClassificationMappingsClient) ListByReplicationStorageClassificationsSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListByReplicationStorageClassificationsResponder handles the response to the ListByReplicationStorageClassifications request. The method always
// closes the http.Response Body.
func (client ReplicationStorageClassificationMappingsClient) ListByReplicationStorageClassificationsResponder(resp *http.Response) (result StorageClassificationMappingCollection, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByReplicationStorageClassificationsNextResults retrieves the next set of results, if any.
func (client ReplicationStorageClassificationMappingsClient) ListByReplicationStorageClassificationsNextResults(lastResults StorageClassificationMappingCollection) (result StorageClassificationMappingCollection, err error) {
	req, err := lastResults.StorageClassificationMappingCollectionPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "service.ReplicationStorageClassificationMappingsClient", "ListByReplicationStorageClassifications", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListByReplicationStorageClassificationsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "service.ReplicationStorageClassificationMappingsClient", "ListByReplicationStorageClassifications", resp, "Failure sending next results request")
	}

	result, err = client.ListByReplicationStorageClassificationsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "service.ReplicationStorageClassificationMappingsClient", "ListByReplicationStorageClassifications", resp, "Failure responding to next results request")
	}

	return
}
