package apimanagement

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

// SubscriptionsClient is the use these REST APIs for performing operations on
// entities like API, Product, and Subscription associated with your Azure API
// Management deployment.
type SubscriptionsClient struct {
	ManagementClient
}

// NewSubscriptionsClient creates an instance of the SubscriptionsClient
// client.
func NewSubscriptionsClient(subscriptionID string) SubscriptionsClient {
	return NewSubscriptionsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewSubscriptionsClientWithBaseURI creates an instance of the
// SubscriptionsClient client.
func NewSubscriptionsClientWithBaseURI(baseURI string, subscriptionID string) SubscriptionsClient {
	return SubscriptionsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates or updates the subscription of specified user to the
// specified product.
//
// resourceGroupName is the name of the resource group. serviceName is the name
// of the API Management service. sid is identifier of the subscription.
// parameters is create parameters.
func (client SubscriptionsClient) CreateOrUpdate(resourceGroupName string, serviceName string, sid string, parameters SubscriptionCreateParameters) (result autorest.Response, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: serviceName,
			Constraints: []validation.Constraint{{Target: "serviceName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "serviceName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "serviceName", Name: validation.Pattern, Rule: `^[a-zA-Z](?:[a-zA-Z0-9-]*[a-zA-Z0-9])?$`, Chain: nil}}},
		{TargetValue: sid,
			Constraints: []validation.Constraint{{Target: "sid", Name: validation.MaxLength, Rule: 256, Chain: nil},
				{Target: "sid", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "sid", Name: validation.Pattern, Rule: `^[^*#&+:<>?]+$`, Chain: nil}}},
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.UserID", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "parameters.ProductID", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "parameters.Name", Name: validation.Null, Rule: true,
					Chain: []validation.Constraint{{Target: "parameters.Name", Name: validation.MaxLength, Rule: 100, Chain: nil},
						{Target: "parameters.Name", Name: validation.MinLength, Rule: 1, Chain: nil},
					}},
				{Target: "parameters.PrimaryKey", Name: validation.Null, Rule: false,
					Chain: []validation.Constraint{{Target: "parameters.PrimaryKey", Name: validation.MaxLength, Rule: 256, Chain: nil},
						{Target: "parameters.PrimaryKey", Name: validation.MinLength, Rule: 1, Chain: nil},
					}},
				{Target: "parameters.SecondaryKey", Name: validation.Null, Rule: false,
					Chain: []validation.Constraint{{Target: "parameters.SecondaryKey", Name: validation.MaxLength, Rule: 256, Chain: nil},
						{Target: "parameters.SecondaryKey", Name: validation.MinLength, Rule: 1, Chain: nil},
					}}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "apimanagement.SubscriptionsClient", "CreateOrUpdate")
	}

	req, err := client.CreateOrUpdatePreparer(resourceGroupName, serviceName, sid, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.SubscriptionsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "apimanagement.SubscriptionsClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.SubscriptionsClient", "CreateOrUpdate", resp, "Failure responding to request")
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client SubscriptionsClient) CreateOrUpdatePreparer(resourceGroupName string, serviceName string, sid string, parameters SubscriptionCreateParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
		"sid":               autorest.Encode("path", sid),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-07-07"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/subscriptions/{sid}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client SubscriptionsClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client SubscriptionsClient) CreateOrUpdateResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Delete deletes the specified subscription.
//
// resourceGroupName is the name of the resource group. serviceName is the name
// of the API Management service. sid is identifier of the subscription.
// ifMatch is eTag of the Subscription Entity. ETag should match the current
// entity state from the header response of the GET request or it should be *
// for unconditional update.
func (client SubscriptionsClient) Delete(resourceGroupName string, serviceName string, sid string, ifMatch string) (result autorest.Response, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: serviceName,
			Constraints: []validation.Constraint{{Target: "serviceName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "serviceName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "serviceName", Name: validation.Pattern, Rule: `^[a-zA-Z](?:[a-zA-Z0-9-]*[a-zA-Z0-9])?$`, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "apimanagement.SubscriptionsClient", "Delete")
	}

	req, err := client.DeletePreparer(resourceGroupName, serviceName, sid, ifMatch)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.SubscriptionsClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "apimanagement.SubscriptionsClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.SubscriptionsClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client SubscriptionsClient) DeletePreparer(resourceGroupName string, serviceName string, sid string, ifMatch string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
		"sid":               autorest.Encode("path", sid),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-07-07"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/subscriptions/{sid}", pathParameters),
		autorest.WithQueryParameters(queryParameters),
		autorest.WithHeader("If-Match", autorest.String(ifMatch)))
	return preparer.Prepare(&http.Request{})
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client SubscriptionsClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client SubscriptionsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets the specified Subscription entity.
//
// resourceGroupName is the name of the resource group. serviceName is the name
// of the API Management service. sid is identifier of the subscription.
func (client SubscriptionsClient) Get(resourceGroupName string, serviceName string, sid string) (result SubscriptionContract, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: serviceName,
			Constraints: []validation.Constraint{{Target: "serviceName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "serviceName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "serviceName", Name: validation.Pattern, Rule: `^[a-zA-Z](?:[a-zA-Z0-9-]*[a-zA-Z0-9])?$`, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "apimanagement.SubscriptionsClient", "Get")
	}

	req, err := client.GetPreparer(resourceGroupName, serviceName, sid)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.SubscriptionsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "apimanagement.SubscriptionsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.SubscriptionsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client SubscriptionsClient) GetPreparer(resourceGroupName string, serviceName string, sid string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
		"sid":               autorest.Encode("path", sid),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-07-07"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/subscriptions/{sid}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client SubscriptionsClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client SubscriptionsClient) GetResponder(resp *http.Response) (result SubscriptionContract, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByService lists all subscriptions of the API Management service
// instance.
//
// resourceGroupName is the name of the resource group. serviceName is the name
// of the API Management service. filter is | Field        | Supported
// operators    | Supported functions                         |
// |--------------|------------------------|---------------------------------------------|
// | id           | ge, le, eq, ne, gt, lt | substringof, contains, startswith,
// endswith |
// | name         | ge, le, eq, ne, gt, lt | substringof, contains, startswith,
// endswith |
// | stateComment | ge, le, eq, ne, gt, lt | substringof, contains, startswith,
// endswith |
// | userId       | ge, le, eq, ne, gt, lt | substringof, contains, startswith,
// endswith |
// | productId    | ge, le, eq, ne, gt, lt | substringof, contains, startswith,
// endswith |
// | state        | eq                     |
// | top is number of records to return. skip is number of records to skip.
func (client SubscriptionsClient) ListByService(resourceGroupName string, serviceName string, filter string, top *int32, skip *int32) (result SubscriptionCollection, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: serviceName,
			Constraints: []validation.Constraint{{Target: "serviceName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "serviceName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "serviceName", Name: validation.Pattern, Rule: `^[a-zA-Z](?:[a-zA-Z0-9-]*[a-zA-Z0-9])?$`, Chain: nil}}},
		{TargetValue: top,
			Constraints: []validation.Constraint{{Target: "top", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "top", Name: validation.InclusiveMinimum, Rule: 1, Chain: nil}}}}},
		{TargetValue: skip,
			Constraints: []validation.Constraint{{Target: "skip", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "skip", Name: validation.InclusiveMinimum, Rule: 0, Chain: nil}}}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "apimanagement.SubscriptionsClient", "ListByService")
	}

	req, err := client.ListByServicePreparer(resourceGroupName, serviceName, filter, top, skip)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.SubscriptionsClient", "ListByService", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByServiceSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "apimanagement.SubscriptionsClient", "ListByService", resp, "Failure sending request")
		return
	}

	result, err = client.ListByServiceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.SubscriptionsClient", "ListByService", resp, "Failure responding to request")
	}

	return
}

// ListByServicePreparer prepares the ListByService request.
func (client SubscriptionsClient) ListByServicePreparer(resourceGroupName string, serviceName string, filter string, top *int32, skip *int32) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-07-07"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}
	if skip != nil {
		queryParameters["$skip"] = autorest.Encode("query", *skip)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/subscriptions", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListByServiceSender sends the ListByService request. The method will close the
// http.Response Body if it receives an error.
func (client SubscriptionsClient) ListByServiceSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListByServiceResponder handles the response to the ListByService request. The method always
// closes the http.Response Body.
func (client SubscriptionsClient) ListByServiceResponder(resp *http.Response) (result SubscriptionCollection, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByServiceNextResults retrieves the next set of results, if any.
func (client SubscriptionsClient) ListByServiceNextResults(lastResults SubscriptionCollection) (result SubscriptionCollection, err error) {
	req, err := lastResults.SubscriptionCollectionPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "apimanagement.SubscriptionsClient", "ListByService", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListByServiceSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "apimanagement.SubscriptionsClient", "ListByService", resp, "Failure sending next results request")
	}

	result, err = client.ListByServiceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.SubscriptionsClient", "ListByService", resp, "Failure responding to next results request")
	}

	return
}

// RegeneratePrimaryKey regenerates primary key of existing subscription of the
// API Management service instance.
//
// resourceGroupName is the name of the resource group. serviceName is the name
// of the API Management service. sid is identifier of the subscription.
func (client SubscriptionsClient) RegeneratePrimaryKey(resourceGroupName string, serviceName string, sid string) (result autorest.Response, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: serviceName,
			Constraints: []validation.Constraint{{Target: "serviceName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "serviceName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "serviceName", Name: validation.Pattern, Rule: `^[a-zA-Z](?:[a-zA-Z0-9-]*[a-zA-Z0-9])?$`, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "apimanagement.SubscriptionsClient", "RegeneratePrimaryKey")
	}

	req, err := client.RegeneratePrimaryKeyPreparer(resourceGroupName, serviceName, sid)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.SubscriptionsClient", "RegeneratePrimaryKey", nil, "Failure preparing request")
		return
	}

	resp, err := client.RegeneratePrimaryKeySender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "apimanagement.SubscriptionsClient", "RegeneratePrimaryKey", resp, "Failure sending request")
		return
	}

	result, err = client.RegeneratePrimaryKeyResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.SubscriptionsClient", "RegeneratePrimaryKey", resp, "Failure responding to request")
	}

	return
}

// RegeneratePrimaryKeyPreparer prepares the RegeneratePrimaryKey request.
func (client SubscriptionsClient) RegeneratePrimaryKeyPreparer(resourceGroupName string, serviceName string, sid string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
		"sid":               autorest.Encode("path", sid),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-07-07"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/subscriptions/{sid}/regeneratePrimaryKey", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// RegeneratePrimaryKeySender sends the RegeneratePrimaryKey request. The method will close the
// http.Response Body if it receives an error.
func (client SubscriptionsClient) RegeneratePrimaryKeySender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// RegeneratePrimaryKeyResponder handles the response to the RegeneratePrimaryKey request. The method always
// closes the http.Response Body.
func (client SubscriptionsClient) RegeneratePrimaryKeyResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// RegenerateSecondaryKey regenerates secondary key of existing subscription of
// the API Management service instance.
//
// resourceGroupName is the name of the resource group. serviceName is the name
// of the API Management service. sid is identifier of the subscription.
func (client SubscriptionsClient) RegenerateSecondaryKey(resourceGroupName string, serviceName string, sid string) (result autorest.Response, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: serviceName,
			Constraints: []validation.Constraint{{Target: "serviceName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "serviceName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "serviceName", Name: validation.Pattern, Rule: `^[a-zA-Z](?:[a-zA-Z0-9-]*[a-zA-Z0-9])?$`, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "apimanagement.SubscriptionsClient", "RegenerateSecondaryKey")
	}

	req, err := client.RegenerateSecondaryKeyPreparer(resourceGroupName, serviceName, sid)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.SubscriptionsClient", "RegenerateSecondaryKey", nil, "Failure preparing request")
		return
	}

	resp, err := client.RegenerateSecondaryKeySender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "apimanagement.SubscriptionsClient", "RegenerateSecondaryKey", resp, "Failure sending request")
		return
	}

	result, err = client.RegenerateSecondaryKeyResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.SubscriptionsClient", "RegenerateSecondaryKey", resp, "Failure responding to request")
	}

	return
}

// RegenerateSecondaryKeyPreparer prepares the RegenerateSecondaryKey request.
func (client SubscriptionsClient) RegenerateSecondaryKeyPreparer(resourceGroupName string, serviceName string, sid string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
		"sid":               autorest.Encode("path", sid),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-07-07"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/subscriptions/{sid}/regenerateSecondaryKey", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// RegenerateSecondaryKeySender sends the RegenerateSecondaryKey request. The method will close the
// http.Response Body if it receives an error.
func (client SubscriptionsClient) RegenerateSecondaryKeySender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// RegenerateSecondaryKeyResponder handles the response to the RegenerateSecondaryKey request. The method always
// closes the http.Response Body.
func (client SubscriptionsClient) RegenerateSecondaryKeyResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Update updates the details of a subscription specificied by its identifier.
//
// resourceGroupName is the name of the resource group. serviceName is the name
// of the API Management service. sid is identifier of the subscription.
// parameters is update parameters. ifMatch is eTag of the Subscription Entity.
// ETag should match the current entity state from the header response of the
// GET request or it should be * for unconditional update.
func (client SubscriptionsClient) Update(resourceGroupName string, serviceName string, sid string, parameters SubscriptionUpdateParameters, ifMatch string) (result autorest.Response, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: serviceName,
			Constraints: []validation.Constraint{{Target: "serviceName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "serviceName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "serviceName", Name: validation.Pattern, Rule: `^[a-zA-Z](?:[a-zA-Z0-9-]*[a-zA-Z0-9])?$`, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "apimanagement.SubscriptionsClient", "Update")
	}

	req, err := client.UpdatePreparer(resourceGroupName, serviceName, sid, parameters, ifMatch)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.SubscriptionsClient", "Update", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "apimanagement.SubscriptionsClient", "Update", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.SubscriptionsClient", "Update", resp, "Failure responding to request")
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client SubscriptionsClient) UpdatePreparer(resourceGroupName string, serviceName string, sid string, parameters SubscriptionUpdateParameters, ifMatch string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
		"sid":               autorest.Encode("path", sid),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-07-07"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/subscriptions/{sid}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters),
		autorest.WithHeader("If-Match", autorest.String(ifMatch)))
	return preparer.Prepare(&http.Request{})
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client SubscriptionsClient) UpdateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client SubscriptionsClient) UpdateResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}
