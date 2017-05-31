package dtl

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

// LabsClient is the the DevTest Labs Client.
type LabsClient struct {
	ManagementClient
}

// NewLabsClient creates an instance of the LabsClient client.
func NewLabsClient(subscriptionID string) LabsClient {
	return NewLabsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewLabsClientWithBaseURI creates an instance of the LabsClient client.
func NewLabsClientWithBaseURI(baseURI string, subscriptionID string) LabsClient {
	return LabsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// ClaimAnyVM claim a random claimable virtual machine in the lab. This
// operation can take a while to complete. This method may poll for completion.
// Polling can be canceled by passing the cancel channel argument. The channel
// will be used to cancel polling and any outstanding HTTP requests.
//
// resourceGroupName is the name of the resource group. name is the name of the
// lab.
func (client LabsClient) ClaimAnyVM(resourceGroupName string, name string, cancel <-chan struct{}) (<-chan autorest.Response, <-chan error) {
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
		req, err := client.ClaimAnyVMPreparer(resourceGroupName, name, cancel)
		if err != nil {
			err = autorest.NewErrorWithError(err, "dtl.LabsClient", "ClaimAnyVM", nil, "Failure preparing request")
			return
		}

		resp, err := client.ClaimAnyVMSender(req)
		if err != nil {
			result.Response = resp
			err = autorest.NewErrorWithError(err, "dtl.LabsClient", "ClaimAnyVM", resp, "Failure sending request")
			return
		}

		result, err = client.ClaimAnyVMResponder(resp)
		if err != nil {
			err = autorest.NewErrorWithError(err, "dtl.LabsClient", "ClaimAnyVM", resp, "Failure responding to request")
		}
	}()
	return resultChan, errChan
}

// ClaimAnyVMPreparer prepares the ClaimAnyVM request.
func (client LabsClient) ClaimAnyVMPreparer(resourceGroupName string, name string, cancel <-chan struct{}) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"name":              autorest.Encode("path", name),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-05-15"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{name}/claimAnyVm", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{Cancel: cancel})
}

// ClaimAnyVMSender sends the ClaimAnyVM request. The method will close the
// http.Response Body if it receives an error.
func (client LabsClient) ClaimAnyVMSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoPollForAsynchronous(client.PollingDelay))
}

// ClaimAnyVMResponder handles the response to the ClaimAnyVM request. The method always
// closes the http.Response Body.
func (client LabsClient) ClaimAnyVMResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}

// CreateEnvironment create virtual machines in a lab. This operation can take
// a while to complete. This method may poll for completion. Polling can be
// canceled by passing the cancel channel argument. The channel will be used to
// cancel polling and any outstanding HTTP requests.
//
// resourceGroupName is the name of the resource group. name is the name of the
// lab. labVirtualMachineCreationParameter is properties for creating a virtual
// machine.
func (client LabsClient) CreateEnvironment(resourceGroupName string, name string, labVirtualMachineCreationParameter LabVirtualMachineCreationParameter, cancel <-chan struct{}) (<-chan autorest.Response, <-chan error) {
	resultChan := make(chan autorest.Response, 1)
	errChan := make(chan error, 1)
	if err := validation.Validate([]validation.Validation{
		{TargetValue: labVirtualMachineCreationParameter,
			Constraints: []validation.Constraint{{Target: "labVirtualMachineCreationParameter.LabVirtualMachineCreationParameterProperties", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "labVirtualMachineCreationParameter.LabVirtualMachineCreationParameterProperties.ApplicableSchedule", Name: validation.Null, Rule: false,
					Chain: []validation.Constraint{{Target: "labVirtualMachineCreationParameter.LabVirtualMachineCreationParameterProperties.ApplicableSchedule.ApplicableScheduleProperties", Name: validation.Null, Rule: true,
						Chain: []validation.Constraint{{Target: "labVirtualMachineCreationParameter.LabVirtualMachineCreationParameterProperties.ApplicableSchedule.ApplicableScheduleProperties.LabVmsShutdown", Name: validation.Null, Rule: false,
							Chain: []validation.Constraint{{Target: "labVirtualMachineCreationParameter.LabVirtualMachineCreationParameterProperties.ApplicableSchedule.ApplicableScheduleProperties.LabVmsShutdown.ScheduleProperties", Name: validation.Null, Rule: true, Chain: nil}}},
							{Target: "labVirtualMachineCreationParameter.LabVirtualMachineCreationParameterProperties.ApplicableSchedule.ApplicableScheduleProperties.LabVmsStartup", Name: validation.Null, Rule: false,
								Chain: []validation.Constraint{{Target: "labVirtualMachineCreationParameter.LabVirtualMachineCreationParameterProperties.ApplicableSchedule.ApplicableScheduleProperties.LabVmsStartup.ScheduleProperties", Name: validation.Null, Rule: true, Chain: nil}}},
						}},
					}},
				}}}}}); err != nil {
		errChan <- validation.NewErrorWithValidationError(err, "dtl.LabsClient", "CreateEnvironment")
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
		req, err := client.CreateEnvironmentPreparer(resourceGroupName, name, labVirtualMachineCreationParameter, cancel)
		if err != nil {
			err = autorest.NewErrorWithError(err, "dtl.LabsClient", "CreateEnvironment", nil, "Failure preparing request")
			return
		}

		resp, err := client.CreateEnvironmentSender(req)
		if err != nil {
			result.Response = resp
			err = autorest.NewErrorWithError(err, "dtl.LabsClient", "CreateEnvironment", resp, "Failure sending request")
			return
		}

		result, err = client.CreateEnvironmentResponder(resp)
		if err != nil {
			err = autorest.NewErrorWithError(err, "dtl.LabsClient", "CreateEnvironment", resp, "Failure responding to request")
		}
	}()
	return resultChan, errChan
}

// CreateEnvironmentPreparer prepares the CreateEnvironment request.
func (client LabsClient) CreateEnvironmentPreparer(resourceGroupName string, name string, labVirtualMachineCreationParameter LabVirtualMachineCreationParameter, cancel <-chan struct{}) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"name":              autorest.Encode("path", name),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-05-15"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{name}/createEnvironment", pathParameters),
		autorest.WithJSON(labVirtualMachineCreationParameter),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{Cancel: cancel})
}

// CreateEnvironmentSender sends the CreateEnvironment request. The method will close the
// http.Response Body if it receives an error.
func (client LabsClient) CreateEnvironmentSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoPollForAsynchronous(client.PollingDelay))
}

// CreateEnvironmentResponder handles the response to the CreateEnvironment request. The method always
// closes the http.Response Body.
func (client LabsClient) CreateEnvironmentResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}

// CreateOrUpdate create or replace an existing lab. This operation can take a
// while to complete. This method may poll for completion. Polling can be
// canceled by passing the cancel channel argument. The channel will be used to
// cancel polling and any outstanding HTTP requests.
//
// resourceGroupName is the name of the resource group. name is the name of the
// lab. lab is a lab.
func (client LabsClient) CreateOrUpdate(resourceGroupName string, name string, lab Lab, cancel <-chan struct{}) (<-chan Lab, <-chan error) {
	resultChan := make(chan Lab, 1)
	errChan := make(chan error, 1)
	go func() {
		var err error
		var result Lab
		defer func() {
			resultChan <- result
			errChan <- err
			close(resultChan)
			close(errChan)
		}()
		req, err := client.CreateOrUpdatePreparer(resourceGroupName, name, lab, cancel)
		if err != nil {
			err = autorest.NewErrorWithError(err, "dtl.LabsClient", "CreateOrUpdate", nil, "Failure preparing request")
			return
		}

		resp, err := client.CreateOrUpdateSender(req)
		if err != nil {
			result.Response = autorest.Response{Response: resp}
			err = autorest.NewErrorWithError(err, "dtl.LabsClient", "CreateOrUpdate", resp, "Failure sending request")
			return
		}

		result, err = client.CreateOrUpdateResponder(resp)
		if err != nil {
			err = autorest.NewErrorWithError(err, "dtl.LabsClient", "CreateOrUpdate", resp, "Failure responding to request")
		}
	}()
	return resultChan, errChan
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client LabsClient) CreateOrUpdatePreparer(resourceGroupName string, name string, lab Lab, cancel <-chan struct{}) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"name":              autorest.Encode("path", name),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-05-15"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{name}", pathParameters),
		autorest.WithJSON(lab),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{Cancel: cancel})
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client LabsClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoPollForAsynchronous(client.PollingDelay))
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client LabsClient) CreateOrUpdateResponder(resp *http.Response) (result Lab, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete delete lab. This operation can take a while to complete. This method
// may poll for completion. Polling can be canceled by passing the cancel
// channel argument. The channel will be used to cancel polling and any
// outstanding HTTP requests.
//
// resourceGroupName is the name of the resource group. name is the name of the
// lab.
func (client LabsClient) Delete(resourceGroupName string, name string, cancel <-chan struct{}) (<-chan autorest.Response, <-chan error) {
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
		req, err := client.DeletePreparer(resourceGroupName, name, cancel)
		if err != nil {
			err = autorest.NewErrorWithError(err, "dtl.LabsClient", "Delete", nil, "Failure preparing request")
			return
		}

		resp, err := client.DeleteSender(req)
		if err != nil {
			result.Response = resp
			err = autorest.NewErrorWithError(err, "dtl.LabsClient", "Delete", resp, "Failure sending request")
			return
		}

		result, err = client.DeleteResponder(resp)
		if err != nil {
			err = autorest.NewErrorWithError(err, "dtl.LabsClient", "Delete", resp, "Failure responding to request")
		}
	}()
	return resultChan, errChan
}

// DeletePreparer prepares the Delete request.
func (client LabsClient) DeletePreparer(resourceGroupName string, name string, cancel <-chan struct{}) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"name":              autorest.Encode("path", name),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-05-15"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{name}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{Cancel: cancel})
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client LabsClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoPollForAsynchronous(client.PollingDelay))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client LabsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// ExportResourceUsage exports the lab resource usage into a storage account
// This operation can take a while to complete. This method may poll for
// completion. Polling can be canceled by passing the cancel channel argument.
// The channel will be used to cancel polling and any outstanding HTTP
// requests.
//
// resourceGroupName is the name of the resource group. name is the name of the
// lab. exportResourceUsageParameters is the parameters of the export
// operation.
func (client LabsClient) ExportResourceUsage(resourceGroupName string, name string, exportResourceUsageParameters ExportResourceUsageParameters, cancel <-chan struct{}) (<-chan autorest.Response, <-chan error) {
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
		req, err := client.ExportResourceUsagePreparer(resourceGroupName, name, exportResourceUsageParameters, cancel)
		if err != nil {
			err = autorest.NewErrorWithError(err, "dtl.LabsClient", "ExportResourceUsage", nil, "Failure preparing request")
			return
		}

		resp, err := client.ExportResourceUsageSender(req)
		if err != nil {
			result.Response = resp
			err = autorest.NewErrorWithError(err, "dtl.LabsClient", "ExportResourceUsage", resp, "Failure sending request")
			return
		}

		result, err = client.ExportResourceUsageResponder(resp)
		if err != nil {
			err = autorest.NewErrorWithError(err, "dtl.LabsClient", "ExportResourceUsage", resp, "Failure responding to request")
		}
	}()
	return resultChan, errChan
}

// ExportResourceUsagePreparer prepares the ExportResourceUsage request.
func (client LabsClient) ExportResourceUsagePreparer(resourceGroupName string, name string, exportResourceUsageParameters ExportResourceUsageParameters, cancel <-chan struct{}) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"name":              autorest.Encode("path", name),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-05-15"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{name}/exportResourceUsage", pathParameters),
		autorest.WithJSON(exportResourceUsageParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{Cancel: cancel})
}

// ExportResourceUsageSender sends the ExportResourceUsage request. The method will close the
// http.Response Body if it receives an error.
func (client LabsClient) ExportResourceUsageSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoPollForAsynchronous(client.PollingDelay))
}

// ExportResourceUsageResponder handles the response to the ExportResourceUsage request. The method always
// closes the http.Response Body.
func (client LabsClient) ExportResourceUsageResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}

// GenerateUploadURI generate a URI for uploading custom disk images to a Lab.
//
// resourceGroupName is the name of the resource group. name is the name of the
// lab. generateUploadURIParameter is properties for generating an upload URI.
func (client LabsClient) GenerateUploadURI(resourceGroupName string, name string, generateUploadURIParameter GenerateUploadURIParameter) (result GenerateUploadURIResponse, err error) {
	req, err := client.GenerateUploadURIPreparer(resourceGroupName, name, generateUploadURIParameter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dtl.LabsClient", "GenerateUploadURI", nil, "Failure preparing request")
		return
	}

	resp, err := client.GenerateUploadURISender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "dtl.LabsClient", "GenerateUploadURI", resp, "Failure sending request")
		return
	}

	result, err = client.GenerateUploadURIResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dtl.LabsClient", "GenerateUploadURI", resp, "Failure responding to request")
	}

	return
}

// GenerateUploadURIPreparer prepares the GenerateUploadURI request.
func (client LabsClient) GenerateUploadURIPreparer(resourceGroupName string, name string, generateUploadURIParameter GenerateUploadURIParameter) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"name":              autorest.Encode("path", name),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-05-15"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{name}/generateUploadUri", pathParameters),
		autorest.WithJSON(generateUploadURIParameter),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GenerateUploadURISender sends the GenerateUploadURI request. The method will close the
// http.Response Body if it receives an error.
func (client LabsClient) GenerateUploadURISender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GenerateUploadURIResponder handles the response to the GenerateUploadURI request. The method always
// closes the http.Response Body.
func (client LabsClient) GenerateUploadURIResponder(resp *http.Response) (result GenerateUploadURIResponse, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Get get lab.
//
// resourceGroupName is the name of the resource group. name is the name of the
// lab. expand is specify the $expand query. Example:
// 'properties($select=defaultStorageAccount)'
func (client LabsClient) Get(resourceGroupName string, name string, expand string) (result Lab, err error) {
	req, err := client.GetPreparer(resourceGroupName, name, expand)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dtl.LabsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "dtl.LabsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dtl.LabsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client LabsClient) GetPreparer(resourceGroupName string, name string, expand string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"name":              autorest.Encode("path", name),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-05-15"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(expand) > 0 {
		queryParameters["$expand"] = autorest.Encode("query", expand)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{name}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client LabsClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client LabsClient) GetResponder(resp *http.Response) (result Lab, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByResourceGroup list labs in a resource group.
//
// resourceGroupName is the name of the resource group. expand is specify the
// $expand query. Example: 'properties($select=defaultStorageAccount)' filter
// is the filter to apply to the operation. top is the maximum number of
// resources to return from the operation. orderby is the ordering expression
// for the results, using OData notation.
func (client LabsClient) ListByResourceGroup(resourceGroupName string, expand string, filter string, top *int32, orderby string) (result ResponseWithContinuationLab, err error) {
	req, err := client.ListByResourceGroupPreparer(resourceGroupName, expand, filter, top, orderby)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dtl.LabsClient", "ListByResourceGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "dtl.LabsClient", "ListByResourceGroup", resp, "Failure sending request")
		return
	}

	result, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dtl.LabsClient", "ListByResourceGroup", resp, "Failure responding to request")
	}

	return
}

// ListByResourceGroupPreparer prepares the ListByResourceGroup request.
func (client LabsClient) ListByResourceGroupPreparer(resourceGroupName string, expand string, filter string, top *int32, orderby string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-05-15"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(expand) > 0 {
		queryParameters["$expand"] = autorest.Encode("query", expand)
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}
	if len(orderby) > 0 {
		queryParameters["$orderby"] = autorest.Encode("query", orderby)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListByResourceGroupSender sends the ListByResourceGroup request. The method will close the
// http.Response Body if it receives an error.
func (client LabsClient) ListByResourceGroupSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListByResourceGroupResponder handles the response to the ListByResourceGroup request. The method always
// closes the http.Response Body.
func (client LabsClient) ListByResourceGroupResponder(resp *http.Response) (result ResponseWithContinuationLab, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByResourceGroupNextResults retrieves the next set of results, if any.
func (client LabsClient) ListByResourceGroupNextResults(lastResults ResponseWithContinuationLab) (result ResponseWithContinuationLab, err error) {
	req, err := lastResults.ResponseWithContinuationLabPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "dtl.LabsClient", "ListByResourceGroup", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "dtl.LabsClient", "ListByResourceGroup", resp, "Failure sending next results request")
	}

	result, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dtl.LabsClient", "ListByResourceGroup", resp, "Failure responding to next results request")
	}

	return
}

// ListBySubscription list labs in a subscription.
//
// expand is specify the $expand query. Example:
// 'properties($select=defaultStorageAccount)' filter is the filter to apply to
// the operation. top is the maximum number of resources to return from the
// operation. orderby is the ordering expression for the results, using OData
// notation.
func (client LabsClient) ListBySubscription(expand string, filter string, top *int32, orderby string) (result ResponseWithContinuationLab, err error) {
	req, err := client.ListBySubscriptionPreparer(expand, filter, top, orderby)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dtl.LabsClient", "ListBySubscription", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListBySubscriptionSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "dtl.LabsClient", "ListBySubscription", resp, "Failure sending request")
		return
	}

	result, err = client.ListBySubscriptionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dtl.LabsClient", "ListBySubscription", resp, "Failure responding to request")
	}

	return
}

// ListBySubscriptionPreparer prepares the ListBySubscription request.
func (client LabsClient) ListBySubscriptionPreparer(expand string, filter string, top *int32, orderby string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-05-15"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(expand) > 0 {
		queryParameters["$expand"] = autorest.Encode("query", expand)
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}
	if len(orderby) > 0 {
		queryParameters["$orderby"] = autorest.Encode("query", orderby)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.DevTestLab/labs", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListBySubscriptionSender sends the ListBySubscription request. The method will close the
// http.Response Body if it receives an error.
func (client LabsClient) ListBySubscriptionSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListBySubscriptionResponder handles the response to the ListBySubscription request. The method always
// closes the http.Response Body.
func (client LabsClient) ListBySubscriptionResponder(resp *http.Response) (result ResponseWithContinuationLab, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListBySubscriptionNextResults retrieves the next set of results, if any.
func (client LabsClient) ListBySubscriptionNextResults(lastResults ResponseWithContinuationLab) (result ResponseWithContinuationLab, err error) {
	req, err := lastResults.ResponseWithContinuationLabPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "dtl.LabsClient", "ListBySubscription", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListBySubscriptionSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "dtl.LabsClient", "ListBySubscription", resp, "Failure sending next results request")
	}

	result, err = client.ListBySubscriptionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dtl.LabsClient", "ListBySubscription", resp, "Failure responding to next results request")
	}

	return
}

// ListVhds list disk images available for custom image creation.
//
// resourceGroupName is the name of the resource group. name is the name of the
// lab.
func (client LabsClient) ListVhds(resourceGroupName string, name string) (result ResponseWithContinuationLabVhd, err error) {
	req, err := client.ListVhdsPreparer(resourceGroupName, name)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dtl.LabsClient", "ListVhds", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListVhdsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "dtl.LabsClient", "ListVhds", resp, "Failure sending request")
		return
	}

	result, err = client.ListVhdsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dtl.LabsClient", "ListVhds", resp, "Failure responding to request")
	}

	return
}

// ListVhdsPreparer prepares the ListVhds request.
func (client LabsClient) ListVhdsPreparer(resourceGroupName string, name string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"name":              autorest.Encode("path", name),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-05-15"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{name}/listVhds", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListVhdsSender sends the ListVhds request. The method will close the
// http.Response Body if it receives an error.
func (client LabsClient) ListVhdsSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListVhdsResponder handles the response to the ListVhds request. The method always
// closes the http.Response Body.
func (client LabsClient) ListVhdsResponder(resp *http.Response) (result ResponseWithContinuationLabVhd, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListVhdsNextResults retrieves the next set of results, if any.
func (client LabsClient) ListVhdsNextResults(lastResults ResponseWithContinuationLabVhd) (result ResponseWithContinuationLabVhd, err error) {
	req, err := lastResults.ResponseWithContinuationLabVhdPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "dtl.LabsClient", "ListVhds", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListVhdsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "dtl.LabsClient", "ListVhds", resp, "Failure sending next results request")
	}

	result, err = client.ListVhdsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dtl.LabsClient", "ListVhds", resp, "Failure responding to next results request")
	}

	return
}

// Update modify properties of labs.
//
// resourceGroupName is the name of the resource group. name is the name of the
// lab. lab is a lab.
func (client LabsClient) Update(resourceGroupName string, name string, lab LabFragment) (result Lab, err error) {
	req, err := client.UpdatePreparer(resourceGroupName, name, lab)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dtl.LabsClient", "Update", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "dtl.LabsClient", "Update", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dtl.LabsClient", "Update", resp, "Failure responding to request")
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client LabsClient) UpdatePreparer(resourceGroupName string, name string, lab LabFragment) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"name":              autorest.Encode("path", name),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-05-15"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{name}", pathParameters),
		autorest.WithJSON(lab),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client LabsClient) UpdateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client LabsClient) UpdateResponder(resp *http.Response) (result Lab, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
