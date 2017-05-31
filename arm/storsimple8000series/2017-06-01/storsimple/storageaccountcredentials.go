package storsimple

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

// StorageAccountCredentialsClient is the client for the
// StorageAccountCredentials methods of the Storsimple service.
type StorageAccountCredentialsClient struct {
	ManagementClient
}

// NewStorageAccountCredentialsClient creates an instance of the
// StorageAccountCredentialsClient client.
func NewStorageAccountCredentialsClient(deviceName string, subscriptionID string) StorageAccountCredentialsClient {
	return NewStorageAccountCredentialsClientWithBaseURI(DefaultBaseURI, deviceName, subscriptionID)
}

// NewStorageAccountCredentialsClientWithBaseURI creates an instance of the
// StorageAccountCredentialsClient client.
func NewStorageAccountCredentialsClientWithBaseURI(baseURI string, deviceName string, subscriptionID string) StorageAccountCredentialsClient {
	return StorageAccountCredentialsClient{NewWithBaseURI(baseURI, deviceName, subscriptionID)}
}

// CreateOrUpdate creates or updates the storage account credential. This
// method may poll for completion. Polling can be canceled by passing the
// cancel channel argument. The channel will be used to cancel polling and any
// outstanding HTTP requests.
//
// storageAccountCredentialName is the storage account credential name.
// parameters is the storage account credential to be added or updated.
// resourceGroupName is the resource group name managerName is the manager name
func (client StorageAccountCredentialsClient) CreateOrUpdate(storageAccountCredentialName string, parameters StorageAccountCredential, resourceGroupName string, managerName string, cancel <-chan struct{}) (<-chan StorageAccountCredential, <-chan error) {
	resultChan := make(chan StorageAccountCredential, 1)
	errChan := make(chan error, 1)
	if err := validation.Validate([]validation.Validation{
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.StorageAccountCredentialProperties", Name: validation.Null, Rule: true,
				Chain: []validation.Constraint{{Target: "parameters.StorageAccountCredentialProperties.EndPoint", Name: validation.Null, Rule: true, Chain: nil},
					{Target: "parameters.StorageAccountCredentialProperties.AccessKey", Name: validation.Null, Rule: false,
						Chain: []validation.Constraint{{Target: "parameters.StorageAccountCredentialProperties.AccessKey.Value", Name: validation.Null, Rule: true, Chain: nil}}},
				}}}},
		{TargetValue: managerName,
			Constraints: []validation.Constraint{{Target: "managerName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "managerName", Name: validation.MinLength, Rule: 2, Chain: nil}}}}); err != nil {
		errChan <- validation.NewErrorWithValidationError(err, "storsimple.StorageAccountCredentialsClient", "CreateOrUpdate")
		close(errChan)
		close(resultChan)
		return resultChan, errChan
	}

	go func() {
		var err error
		var result StorageAccountCredential
		defer func() {
			resultChan <- result
			errChan <- err
			close(resultChan)
			close(errChan)
		}()
		req, err := client.CreateOrUpdatePreparer(storageAccountCredentialName, parameters, resourceGroupName, managerName, cancel)
		if err != nil {
			err = autorest.NewErrorWithError(err, "storsimple.StorageAccountCredentialsClient", "CreateOrUpdate", nil, "Failure preparing request")
			return
		}

		resp, err := client.CreateOrUpdateSender(req)
		if err != nil {
			result.Response = autorest.Response{Response: resp}
			err = autorest.NewErrorWithError(err, "storsimple.StorageAccountCredentialsClient", "CreateOrUpdate", resp, "Failure sending request")
			return
		}

		result, err = client.CreateOrUpdateResponder(resp)
		if err != nil {
			err = autorest.NewErrorWithError(err, "storsimple.StorageAccountCredentialsClient", "CreateOrUpdate", resp, "Failure responding to request")
		}
	}()
	return resultChan, errChan
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client StorageAccountCredentialsClient) CreateOrUpdatePreparer(storageAccountCredentialName string, parameters StorageAccountCredential, resourceGroupName string, managerName string, cancel <-chan struct{}) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"managerName":                  managerName,
		"resourceGroupName":            resourceGroupName,
		"storageAccountCredentialName": storageAccountCredentialName,
		"subscriptionId":               client.SubscriptionID,
	}

	const APIVersion = "2017-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/storageAccountCredentials/{storageAccountCredentialName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{Cancel: cancel})
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client StorageAccountCredentialsClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoPollForAsynchronous(client.PollingDelay))
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client StorageAccountCredentialsClient) CreateOrUpdateResponder(resp *http.Response) (result StorageAccountCredential, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes the storage account credential. This method may poll for
// completion. Polling can be canceled by passing the cancel channel argument.
// The channel will be used to cancel polling and any outstanding HTTP
// requests.
//
// storageAccountCredentialName is the name of the storage account credential.
// resourceGroupName is the resource group name managerName is the manager name
func (client StorageAccountCredentialsClient) Delete(storageAccountCredentialName string, resourceGroupName string, managerName string, cancel <-chan struct{}) (<-chan autorest.Response, <-chan error) {
	resultChan := make(chan autorest.Response, 1)
	errChan := make(chan error, 1)
	if err := validation.Validate([]validation.Validation{
		{TargetValue: managerName,
			Constraints: []validation.Constraint{{Target: "managerName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "managerName", Name: validation.MinLength, Rule: 2, Chain: nil}}}}); err != nil {
		errChan <- validation.NewErrorWithValidationError(err, "storsimple.StorageAccountCredentialsClient", "Delete")
		close(errChan)
		close(resultChan)
		return resultChan, errChan
	}

	go func() {
		var err error
		var result autorest.Response
		defer func() {
			resultChan <- result
			errChan <- err
			close(resultChan)
			close(errChan)
		}()
		req, err := client.DeletePreparer(storageAccountCredentialName, resourceGroupName, managerName, cancel)
		if err != nil {
			err = autorest.NewErrorWithError(err, "storsimple.StorageAccountCredentialsClient", "Delete", nil, "Failure preparing request")
			return
		}

		resp, err := client.DeleteSender(req)
		if err != nil {
			result.Response = resp
			err = autorest.NewErrorWithError(err, "storsimple.StorageAccountCredentialsClient", "Delete", resp, "Failure sending request")
			return
		}

		result, err = client.DeleteResponder(resp)
		if err != nil {
			err = autorest.NewErrorWithError(err, "storsimple.StorageAccountCredentialsClient", "Delete", resp, "Failure responding to request")
		}
	}()
	return resultChan, errChan
}

// DeletePreparer prepares the Delete request.
func (client StorageAccountCredentialsClient) DeletePreparer(storageAccountCredentialName string, resourceGroupName string, managerName string, cancel <-chan struct{}) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"managerName":                  managerName,
		"resourceGroupName":            resourceGroupName,
		"storageAccountCredentialName": storageAccountCredentialName,
		"subscriptionId":               client.SubscriptionID,
	}

	const APIVersion = "2017-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/storageAccountCredentials/{storageAccountCredentialName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{Cancel: cancel})
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client StorageAccountCredentialsClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoPollForAsynchronous(client.PollingDelay))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client StorageAccountCredentialsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets the properties of the specified storage account credential name.
//
// storageAccountCredentialName is the name of storage account credential to be
// fetched. resourceGroupName is the resource group name managerName is the
// manager name
func (client StorageAccountCredentialsClient) Get(storageAccountCredentialName string, resourceGroupName string, managerName string) (result StorageAccountCredential, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: managerName,
			Constraints: []validation.Constraint{{Target: "managerName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "managerName", Name: validation.MinLength, Rule: 2, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "storsimple.StorageAccountCredentialsClient", "Get")
	}

	req, err := client.GetPreparer(storageAccountCredentialName, resourceGroupName, managerName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storsimple.StorageAccountCredentialsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "storsimple.StorageAccountCredentialsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storsimple.StorageAccountCredentialsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client StorageAccountCredentialsClient) GetPreparer(storageAccountCredentialName string, resourceGroupName string, managerName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"managerName":                  managerName,
		"resourceGroupName":            resourceGroupName,
		"storageAccountCredentialName": storageAccountCredentialName,
		"subscriptionId":               client.SubscriptionID,
	}

	const APIVersion = "2017-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/storageAccountCredentials/{storageAccountCredentialName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client StorageAccountCredentialsClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client StorageAccountCredentialsClient) GetResponder(resp *http.Response) (result StorageAccountCredential, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListAccessKey returns the encrypted storage account access key.
//
// storageAccountCredentialName is name of the storage account credential.
// resourceGroupName is the resource group name managerName is the manager name
func (client StorageAccountCredentialsClient) ListAccessKey(storageAccountCredentialName string, resourceGroupName string, managerName string) (result AsymmetricEncryptedSecret, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: managerName,
			Constraints: []validation.Constraint{{Target: "managerName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "managerName", Name: validation.MinLength, Rule: 2, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "storsimple.StorageAccountCredentialsClient", "ListAccessKey")
	}

	req, err := client.ListAccessKeyPreparer(storageAccountCredentialName, resourceGroupName, managerName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storsimple.StorageAccountCredentialsClient", "ListAccessKey", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListAccessKeySender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "storsimple.StorageAccountCredentialsClient", "ListAccessKey", resp, "Failure sending request")
		return
	}

	result, err = client.ListAccessKeyResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storsimple.StorageAccountCredentialsClient", "ListAccessKey", resp, "Failure responding to request")
	}

	return
}

// ListAccessKeyPreparer prepares the ListAccessKey request.
func (client StorageAccountCredentialsClient) ListAccessKeyPreparer(storageAccountCredentialName string, resourceGroupName string, managerName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"managerName":                  managerName,
		"resourceGroupName":            resourceGroupName,
		"storageAccountCredentialName": storageAccountCredentialName,
		"subscriptionId":               client.SubscriptionID,
	}

	const APIVersion = "2017-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/storageAccountCredentials/{storageAccountCredentialName}/listAccessKey", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListAccessKeySender sends the ListAccessKey request. The method will close the
// http.Response Body if it receives an error.
func (client StorageAccountCredentialsClient) ListAccessKeySender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListAccessKeyResponder handles the response to the ListAccessKey request. The method always
// closes the http.Response Body.
func (client StorageAccountCredentialsClient) ListAccessKeyResponder(resp *http.Response) (result AsymmetricEncryptedSecret, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByManager gets all the storage account credentials in a manager.
//
// resourceGroupName is the resource group name managerName is the manager name
func (client StorageAccountCredentialsClient) ListByManager(resourceGroupName string, managerName string) (result StorageAccountCredentialList, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: managerName,
			Constraints: []validation.Constraint{{Target: "managerName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "managerName", Name: validation.MinLength, Rule: 2, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "storsimple.StorageAccountCredentialsClient", "ListByManager")
	}

	req, err := client.ListByManagerPreparer(resourceGroupName, managerName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storsimple.StorageAccountCredentialsClient", "ListByManager", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByManagerSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "storsimple.StorageAccountCredentialsClient", "ListByManager", resp, "Failure sending request")
		return
	}

	result, err = client.ListByManagerResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storsimple.StorageAccountCredentialsClient", "ListByManager", resp, "Failure responding to request")
	}

	return
}

// ListByManagerPreparer prepares the ListByManager request.
func (client StorageAccountCredentialsClient) ListByManagerPreparer(resourceGroupName string, managerName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"managerName":       managerName,
		"resourceGroupName": resourceGroupName,
		"subscriptionId":    client.SubscriptionID,
	}

	const APIVersion = "2017-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/storageAccountCredentials", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListByManagerSender sends the ListByManager request. The method will close the
// http.Response Body if it receives an error.
func (client StorageAccountCredentialsClient) ListByManagerSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListByManagerResponder handles the response to the ListByManager request. The method always
// closes the http.Response Body.
func (client StorageAccountCredentialsClient) ListByManagerResponder(resp *http.Response) (result StorageAccountCredentialList, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
