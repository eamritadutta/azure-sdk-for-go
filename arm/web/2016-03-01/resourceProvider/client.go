// Package resourceprovider implements the Azure ARM Resourceprovider service
// API version 2016-03-01.
//
//
package resourceprovider

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

const (
	// DefaultBaseURI is the default URI used for the service Resourceprovider
	DefaultBaseURI = "https://management.azure.com"
)

// ManagementClient is the base client for Resourceprovider.
type ManagementClient struct {
	autorest.Client
	BaseURI        string
	SubscriptionID string
}

// New creates an instance of the ManagementClient client.
func New(subscriptionID string) ManagementClient {
	return NewWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewWithBaseURI creates an instance of the ManagementClient client.
func NewWithBaseURI(baseURI string, subscriptionID string) ManagementClient {
	return ManagementClient{
		Client:         autorest.NewClientWithUserAgent(UserAgent()),
		BaseURI:        baseURI,
		SubscriptionID: subscriptionID,
	}
}

// CheckNameAvailability check if a resource name is available.
//
// request is name availability request.
func (client ManagementClient) CheckNameAvailability(request ResourceNameAvailabilityRequest) (result ResourceNameAvailability, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: request,
			Constraints: []validation.Constraint{{Target: "request.Name", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "resourceprovider.ManagementClient", "CheckNameAvailability")
	}

	req, err := client.CheckNameAvailabilityPreparer(request)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resourceprovider.ManagementClient", "CheckNameAvailability", nil, "Failure preparing request")
		return
	}

	resp, err := client.CheckNameAvailabilitySender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "resourceprovider.ManagementClient", "CheckNameAvailability", resp, "Failure sending request")
		return
	}

	result, err = client.CheckNameAvailabilityResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resourceprovider.ManagementClient", "CheckNameAvailability", resp, "Failure responding to request")
	}

	return
}

// CheckNameAvailabilityPreparer prepares the CheckNameAvailability request.
func (client ManagementClient) CheckNameAvailabilityPreparer(request ResourceNameAvailabilityRequest) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Web/checknameavailability", pathParameters),
		autorest.WithJSON(request),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// CheckNameAvailabilitySender sends the CheckNameAvailability request. The method will close the
// http.Response Body if it receives an error.
func (client ManagementClient) CheckNameAvailabilitySender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// CheckNameAvailabilityResponder handles the response to the CheckNameAvailability request. The method always
// closes the http.Response Body.
func (client ManagementClient) CheckNameAvailabilityResponder(resp *http.Response) (result ResourceNameAvailability, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetPublishingUser gets publishing user
func (client ManagementClient) GetPublishingUser() (result User, err error) {
	req, err := client.GetPublishingUserPreparer()
	if err != nil {
		err = autorest.NewErrorWithError(err, "resourceprovider.ManagementClient", "GetPublishingUser", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetPublishingUserSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "resourceprovider.ManagementClient", "GetPublishingUser", resp, "Failure sending request")
		return
	}

	result, err = client.GetPublishingUserResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resourceprovider.ManagementClient", "GetPublishingUser", resp, "Failure responding to request")
	}

	return
}

// GetPublishingUserPreparer prepares the GetPublishingUser request.
func (client ManagementClient) GetPublishingUserPreparer() (*http.Request, error) {
	const APIVersion = "2016-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/providers/Microsoft.Web/publishingUsers/web"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetPublishingUserSender sends the GetPublishingUser request. The method will close the
// http.Response Body if it receives an error.
func (client ManagementClient) GetPublishingUserSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetPublishingUserResponder handles the response to the GetPublishingUser request. The method always
// closes the http.Response Body.
func (client ManagementClient) GetPublishingUserResponder(resp *http.Response) (result User, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListGeoRegions get a list of available geographical regions.
//
// sku is name of SKU used to filter the regions. linuxWorkersEnabled is
// specify <code>true</code> if you want to filter to only regions that support
// Linux workers.
func (client ManagementClient) ListGeoRegions(sku SkuName, linuxWorkersEnabled *bool) (result GeoRegionCollection, err error) {
	req, err := client.ListGeoRegionsPreparer(sku, linuxWorkersEnabled)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resourceprovider.ManagementClient", "ListGeoRegions", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListGeoRegionsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "resourceprovider.ManagementClient", "ListGeoRegions", resp, "Failure sending request")
		return
	}

	result, err = client.ListGeoRegionsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resourceprovider.ManagementClient", "ListGeoRegions", resp, "Failure responding to request")
	}

	return
}

// ListGeoRegionsPreparer prepares the ListGeoRegions request.
func (client ManagementClient) ListGeoRegionsPreparer(sku SkuName, linuxWorkersEnabled *bool) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(string(sku)) > 0 {
		queryParameters["sku"] = autorest.Encode("query", sku)
	}
	if linuxWorkersEnabled != nil {
		queryParameters["linuxWorkersEnabled"] = autorest.Encode("query", *linuxWorkersEnabled)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Web/geoRegions", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListGeoRegionsSender sends the ListGeoRegions request. The method will close the
// http.Response Body if it receives an error.
func (client ManagementClient) ListGeoRegionsSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListGeoRegionsResponder handles the response to the ListGeoRegions request. The method always
// closes the http.Response Body.
func (client ManagementClient) ListGeoRegionsResponder(resp *http.Response) (result GeoRegionCollection, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListGeoRegionsNextResults retrieves the next set of results, if any.
func (client ManagementClient) ListGeoRegionsNextResults(lastResults GeoRegionCollection) (result GeoRegionCollection, err error) {
	req, err := lastResults.GeoRegionCollectionPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "resourceprovider.ManagementClient", "ListGeoRegions", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListGeoRegionsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "resourceprovider.ManagementClient", "ListGeoRegions", resp, "Failure sending next results request")
	}

	result, err = client.ListGeoRegionsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resourceprovider.ManagementClient", "ListGeoRegions", resp, "Failure responding to next results request")
	}

	return
}

// ListPremierAddOnOffers list all premier add-on offers.
func (client ManagementClient) ListPremierAddOnOffers() (result PremierAddOnOfferCollection, err error) {
	req, err := client.ListPremierAddOnOffersPreparer()
	if err != nil {
		err = autorest.NewErrorWithError(err, "resourceprovider.ManagementClient", "ListPremierAddOnOffers", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListPremierAddOnOffersSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "resourceprovider.ManagementClient", "ListPremierAddOnOffers", resp, "Failure sending request")
		return
	}

	result, err = client.ListPremierAddOnOffersResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resourceprovider.ManagementClient", "ListPremierAddOnOffers", resp, "Failure responding to request")
	}

	return
}

// ListPremierAddOnOffersPreparer prepares the ListPremierAddOnOffers request.
func (client ManagementClient) ListPremierAddOnOffersPreparer() (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Web/premieraddonoffers", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListPremierAddOnOffersSender sends the ListPremierAddOnOffers request. The method will close the
// http.Response Body if it receives an error.
func (client ManagementClient) ListPremierAddOnOffersSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListPremierAddOnOffersResponder handles the response to the ListPremierAddOnOffers request. The method always
// closes the http.Response Body.
func (client ManagementClient) ListPremierAddOnOffersResponder(resp *http.Response) (result PremierAddOnOfferCollection, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListPremierAddOnOffersNextResults retrieves the next set of results, if any.
func (client ManagementClient) ListPremierAddOnOffersNextResults(lastResults PremierAddOnOfferCollection) (result PremierAddOnOfferCollection, err error) {
	req, err := lastResults.PremierAddOnOfferCollectionPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "resourceprovider.ManagementClient", "ListPremierAddOnOffers", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListPremierAddOnOffersSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "resourceprovider.ManagementClient", "ListPremierAddOnOffers", resp, "Failure sending next results request")
	}

	result, err = client.ListPremierAddOnOffersResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resourceprovider.ManagementClient", "ListPremierAddOnOffers", resp, "Failure responding to next results request")
	}

	return
}

// ListSkus list all SKUs.
func (client ManagementClient) ListSkus() (result SkuInfos, err error) {
	req, err := client.ListSkusPreparer()
	if err != nil {
		err = autorest.NewErrorWithError(err, "resourceprovider.ManagementClient", "ListSkus", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSkusSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "resourceprovider.ManagementClient", "ListSkus", resp, "Failure sending request")
		return
	}

	result, err = client.ListSkusResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resourceprovider.ManagementClient", "ListSkus", resp, "Failure responding to request")
	}

	return
}

// ListSkusPreparer prepares the ListSkus request.
func (client ManagementClient) ListSkusPreparer() (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Web/skus", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListSkusSender sends the ListSkus request. The method will close the
// http.Response Body if it receives an error.
func (client ManagementClient) ListSkusSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListSkusResponder handles the response to the ListSkus request. The method always
// closes the http.Response Body.
func (client ManagementClient) ListSkusResponder(resp *http.Response) (result SkuInfos, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListSourceControls gets the source controls available for Azure websites.
func (client ManagementClient) ListSourceControls() (result SourceControlCollection, err error) {
	req, err := client.ListSourceControlsPreparer()
	if err != nil {
		err = autorest.NewErrorWithError(err, "resourceprovider.ManagementClient", "ListSourceControls", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSourceControlsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "resourceprovider.ManagementClient", "ListSourceControls", resp, "Failure sending request")
		return
	}

	result, err = client.ListSourceControlsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resourceprovider.ManagementClient", "ListSourceControls", resp, "Failure responding to request")
	}

	return
}

// ListSourceControlsPreparer prepares the ListSourceControls request.
func (client ManagementClient) ListSourceControlsPreparer() (*http.Request, error) {
	const APIVersion = "2016-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/providers/Microsoft.Web/sourcecontrols"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListSourceControlsSender sends the ListSourceControls request. The method will close the
// http.Response Body if it receives an error.
func (client ManagementClient) ListSourceControlsSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListSourceControlsResponder handles the response to the ListSourceControls request. The method always
// closes the http.Response Body.
func (client ManagementClient) ListSourceControlsResponder(resp *http.Response) (result SourceControlCollection, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListSourceControlsNextResults retrieves the next set of results, if any.
func (client ManagementClient) ListSourceControlsNextResults(lastResults SourceControlCollection) (result SourceControlCollection, err error) {
	req, err := lastResults.SourceControlCollectionPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "resourceprovider.ManagementClient", "ListSourceControls", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListSourceControlsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "resourceprovider.ManagementClient", "ListSourceControls", resp, "Failure sending next results request")
	}

	result, err = client.ListSourceControlsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resourceprovider.ManagementClient", "ListSourceControls", resp, "Failure responding to next results request")
	}

	return
}

// Move move resources between resource groups.
//
// resourceGroupName is name of the resource group to which the resource
// belongs. moveResourceEnvelope is object that represents the resource to
// move.
func (client ManagementClient) Move(resourceGroupName string, moveResourceEnvelope CsmMoveResourceEnvelope) (result autorest.Response, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+[^\.]$`, Chain: nil}}},
		{TargetValue: moveResourceEnvelope,
			Constraints: []validation.Constraint{{Target: "moveResourceEnvelope.TargetResourceGroup", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "moveResourceEnvelope.TargetResourceGroup", Name: validation.MaxLength, Rule: 90, Chain: nil},
					{Target: "moveResourceEnvelope.TargetResourceGroup", Name: validation.MinLength, Rule: 1, Chain: nil},
					{Target: "moveResourceEnvelope.TargetResourceGroup", Name: validation.Pattern, Rule: ` ^[-\w\._\(\)]+[^\.]$`, Chain: nil},
				}}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "resourceprovider.ManagementClient", "Move")
	}

	req, err := client.MovePreparer(resourceGroupName, moveResourceEnvelope)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resourceprovider.ManagementClient", "Move", nil, "Failure preparing request")
		return
	}

	resp, err := client.MoveSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "resourceprovider.ManagementClient", "Move", resp, "Failure sending request")
		return
	}

	result, err = client.MoveResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resourceprovider.ManagementClient", "Move", resp, "Failure responding to request")
	}

	return
}

// MovePreparer prepares the Move request.
func (client ManagementClient) MovePreparer(resourceGroupName string, moveResourceEnvelope CsmMoveResourceEnvelope) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/moveResources", pathParameters),
		autorest.WithJSON(moveResourceEnvelope),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// MoveSender sends the Move request. The method will close the
// http.Response Body if it receives an error.
func (client ManagementClient) MoveSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// MoveResponder handles the response to the Move request. The method always
// closes the http.Response Body.
func (client ManagementClient) MoveResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// UpdatePublishingUser updates publishing user
//
// userDetails is details of publishing user
func (client ManagementClient) UpdatePublishingUser(userDetails User) (result User, err error) {
	req, err := client.UpdatePublishingUserPreparer(userDetails)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resourceprovider.ManagementClient", "UpdatePublishingUser", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdatePublishingUserSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "resourceprovider.ManagementClient", "UpdatePublishingUser", resp, "Failure sending request")
		return
	}

	result, err = client.UpdatePublishingUserResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resourceprovider.ManagementClient", "UpdatePublishingUser", resp, "Failure responding to request")
	}

	return
}

// UpdatePublishingUserPreparer prepares the UpdatePublishingUser request.
func (client ManagementClient) UpdatePublishingUserPreparer(userDetails User) (*http.Request, error) {
	const APIVersion = "2016-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/providers/Microsoft.Web/publishingUsers/web"),
		autorest.WithJSON(userDetails),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// UpdatePublishingUserSender sends the UpdatePublishingUser request. The method will close the
// http.Response Body if it receives an error.
func (client ManagementClient) UpdatePublishingUserSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// UpdatePublishingUserResponder handles the response to the UpdatePublishingUser request. The method always
// closes the http.Response Body.
func (client ManagementClient) UpdatePublishingUserResponder(resp *http.Response) (result User, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// UpdateSourceControl updates source control token
//
// sourceControlType is type of source control requestMessage is source control
// token information
func (client ManagementClient) UpdateSourceControl(sourceControlType string, requestMessage SourceControl) (result SourceControl, err error) {
	req, err := client.UpdateSourceControlPreparer(sourceControlType, requestMessage)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resourceprovider.ManagementClient", "UpdateSourceControl", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateSourceControlSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "resourceprovider.ManagementClient", "UpdateSourceControl", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateSourceControlResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resourceprovider.ManagementClient", "UpdateSourceControl", resp, "Failure responding to request")
	}

	return
}

// UpdateSourceControlPreparer prepares the UpdateSourceControl request.
func (client ManagementClient) UpdateSourceControlPreparer(sourceControlType string, requestMessage SourceControl) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"sourceControlType": autorest.Encode("path", sourceControlType),
	}

	const APIVersion = "2016-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Web/sourcecontrols/{sourceControlType}", pathParameters),
		autorest.WithJSON(requestMessage),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// UpdateSourceControlSender sends the UpdateSourceControl request. The method will close the
// http.Response Body if it receives an error.
func (client ManagementClient) UpdateSourceControlSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// UpdateSourceControlResponder handles the response to the UpdateSourceControl request. The method always
// closes the http.Response Body.
func (client ManagementClient) UpdateSourceControlResponder(resp *http.Response) (result SourceControl, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Validate validate if a resource can be created.
//
// resourceGroupName is name of the resource group to which the resource
// belongs. validateRequest is request with the resources to validate.
func (client ManagementClient) Validate(resourceGroupName string, validateRequest ValidateRequest) (result ValidateResponse, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+[^\.]$`, Chain: nil}}},
		{TargetValue: validateRequest,
			Constraints: []validation.Constraint{{Target: "validateRequest.Name", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "validateRequest.Location", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "validateRequest.ValidateProperties", Name: validation.Null, Rule: true,
					Chain: []validation.Constraint{{Target: "validateRequest.ValidateProperties.Capacity", Name: validation.Null, Rule: false,
						Chain: []validation.Constraint{{Target: "validateRequest.ValidateProperties.Capacity", Name: validation.InclusiveMinimum, Rule: 1, Chain: nil}}},
					}}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "resourceprovider.ManagementClient", "Validate")
	}

	req, err := client.ValidatePreparer(resourceGroupName, validateRequest)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resourceprovider.ManagementClient", "Validate", nil, "Failure preparing request")
		return
	}

	resp, err := client.ValidateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "resourceprovider.ManagementClient", "Validate", resp, "Failure sending request")
		return
	}

	result, err = client.ValidateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resourceprovider.ManagementClient", "Validate", resp, "Failure responding to request")
	}

	return
}

// ValidatePreparer prepares the Validate request.
func (client ManagementClient) ValidatePreparer(resourceGroupName string, validateRequest ValidateRequest) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/validate", pathParameters),
		autorest.WithJSON(validateRequest),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ValidateSender sends the Validate request. The method will close the
// http.Response Body if it receives an error.
func (client ManagementClient) ValidateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ValidateResponder handles the response to the Validate request. The method always
// closes the http.Response Body.
func (client ManagementClient) ValidateResponder(resp *http.Response) (result ValidateResponse, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ValidateMove validate whether a resource can be moved.
//
// resourceGroupName is name of the resource group to which the resource
// belongs. moveResourceEnvelope is object that represents the resource to
// move.
func (client ManagementClient) ValidateMove(resourceGroupName string, moveResourceEnvelope CsmMoveResourceEnvelope) (result autorest.Response, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+[^\.]$`, Chain: nil}}},
		{TargetValue: moveResourceEnvelope,
			Constraints: []validation.Constraint{{Target: "moveResourceEnvelope.TargetResourceGroup", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "moveResourceEnvelope.TargetResourceGroup", Name: validation.MaxLength, Rule: 90, Chain: nil},
					{Target: "moveResourceEnvelope.TargetResourceGroup", Name: validation.MinLength, Rule: 1, Chain: nil},
					{Target: "moveResourceEnvelope.TargetResourceGroup", Name: validation.Pattern, Rule: ` ^[-\w\._\(\)]+[^\.]$`, Chain: nil},
				}}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "resourceprovider.ManagementClient", "ValidateMove")
	}

	req, err := client.ValidateMovePreparer(resourceGroupName, moveResourceEnvelope)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resourceprovider.ManagementClient", "ValidateMove", nil, "Failure preparing request")
		return
	}

	resp, err := client.ValidateMoveSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "resourceprovider.ManagementClient", "ValidateMove", resp, "Failure sending request")
		return
	}

	result, err = client.ValidateMoveResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resourceprovider.ManagementClient", "ValidateMove", resp, "Failure responding to request")
	}

	return
}

// ValidateMovePreparer prepares the ValidateMove request.
func (client ManagementClient) ValidateMovePreparer(resourceGroupName string, moveResourceEnvelope CsmMoveResourceEnvelope) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/validateMoveResources", pathParameters),
		autorest.WithJSON(moveResourceEnvelope),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ValidateMoveSender sends the ValidateMove request. The method will close the
// http.Response Body if it receives an error.
func (client ManagementClient) ValidateMoveSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ValidateMoveResponder handles the response to the ValidateMove request. The method always
// closes the http.Response Body.
func (client ManagementClient) ValidateMoveResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}
