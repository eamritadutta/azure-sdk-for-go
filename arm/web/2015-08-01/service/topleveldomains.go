package service

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

// TopLevelDomainsClient is the use these APIs to manage Azure Websites
// resources through the Azure Resource Manager. All task operations conform to
// the HTTP/1.1 protocol specification and each operation returns an
// x-ms-request-id header that can be used to obtain information about the
// request. You must make sure that requests made to these resources are
// secure. For more information, see
// https://msdn.microsoft.com/en-us/library/azure/dn790557.aspx.
type TopLevelDomainsClient struct {
	ManagementClient
}

// NewTopLevelDomainsClient creates an instance of the TopLevelDomainsClient
// client.
func NewTopLevelDomainsClient(subscriptionID string) TopLevelDomainsClient {
	return NewTopLevelDomainsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewTopLevelDomainsClientWithBaseURI creates an instance of the
// TopLevelDomainsClient client.
func NewTopLevelDomainsClientWithBaseURI(baseURI string, subscriptionID string) TopLevelDomainsClient {
	return TopLevelDomainsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// GetGetTopLevelDomains sends the get get top level domains request.
func (client TopLevelDomainsClient) GetGetTopLevelDomains() (result TopLevelDomainCollection, err error) {
	req, err := client.GetGetTopLevelDomainsPreparer()
	if err != nil {
		err = autorest.NewErrorWithError(err, "service.TopLevelDomainsClient", "GetGetTopLevelDomains", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetGetTopLevelDomainsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "service.TopLevelDomainsClient", "GetGetTopLevelDomains", resp, "Failure sending request")
		return
	}

	result, err = client.GetGetTopLevelDomainsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "service.TopLevelDomainsClient", "GetGetTopLevelDomains", resp, "Failure responding to request")
	}

	return
}

// GetGetTopLevelDomainsPreparer prepares the GetGetTopLevelDomains request.
func (client TopLevelDomainsClient) GetGetTopLevelDomainsPreparer() (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-08-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.DomainRegistration/topLevelDomains", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetGetTopLevelDomainsSender sends the GetGetTopLevelDomains request. The method will close the
// http.Response Body if it receives an error.
func (client TopLevelDomainsClient) GetGetTopLevelDomainsSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetGetTopLevelDomainsResponder handles the response to the GetGetTopLevelDomains request. The method always
// closes the http.Response Body.
func (client TopLevelDomainsClient) GetGetTopLevelDomainsResponder(resp *http.Response) (result TopLevelDomainCollection, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetGetTopLevelDomainsNextResults retrieves the next set of results, if any.
func (client TopLevelDomainsClient) GetGetTopLevelDomainsNextResults(lastResults TopLevelDomainCollection) (result TopLevelDomainCollection, err error) {
	req, err := lastResults.TopLevelDomainCollectionPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "service.TopLevelDomainsClient", "GetGetTopLevelDomains", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}

	resp, err := client.GetGetTopLevelDomainsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "service.TopLevelDomainsClient", "GetGetTopLevelDomains", resp, "Failure sending next results request")
	}

	result, err = client.GetGetTopLevelDomainsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "service.TopLevelDomainsClient", "GetGetTopLevelDomains", resp, "Failure responding to next results request")
	}

	return
}

// GetTopLevelDomain sends the get top level domain request.
//
// name is name of the top level domain
func (client TopLevelDomainsClient) GetTopLevelDomain(name string) (result TopLevelDomain, err error) {
	req, err := client.GetTopLevelDomainPreparer(name)
	if err != nil {
		err = autorest.NewErrorWithError(err, "service.TopLevelDomainsClient", "GetTopLevelDomain", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetTopLevelDomainSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "service.TopLevelDomainsClient", "GetTopLevelDomain", resp, "Failure sending request")
		return
	}

	result, err = client.GetTopLevelDomainResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "service.TopLevelDomainsClient", "GetTopLevelDomain", resp, "Failure responding to request")
	}

	return
}

// GetTopLevelDomainPreparer prepares the GetTopLevelDomain request.
func (client TopLevelDomainsClient) GetTopLevelDomainPreparer(name string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"name":           autorest.Encode("path", name),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-08-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.DomainRegistration/topLevelDomains/{name}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetTopLevelDomainSender sends the GetTopLevelDomain request. The method will close the
// http.Response Body if it receives an error.
func (client TopLevelDomainsClient) GetTopLevelDomainSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetTopLevelDomainResponder handles the response to the GetTopLevelDomain request. The method always
// closes the http.Response Body.
func (client TopLevelDomainsClient) GetTopLevelDomainResponder(resp *http.Response) (result TopLevelDomain, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListTopLevelDomainAgreements sends the list top level domain agreements
// request.
//
// name is name of the top level domain agreementOption is domain agreement
// options
func (client TopLevelDomainsClient) ListTopLevelDomainAgreements(name string, agreementOption TopLevelDomainAgreementOption) (result TldLegalAgreementCollection, err error) {
	req, err := client.ListTopLevelDomainAgreementsPreparer(name, agreementOption)
	if err != nil {
		err = autorest.NewErrorWithError(err, "service.TopLevelDomainsClient", "ListTopLevelDomainAgreements", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListTopLevelDomainAgreementsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "service.TopLevelDomainsClient", "ListTopLevelDomainAgreements", resp, "Failure sending request")
		return
	}

	result, err = client.ListTopLevelDomainAgreementsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "service.TopLevelDomainsClient", "ListTopLevelDomainAgreements", resp, "Failure responding to request")
	}

	return
}

// ListTopLevelDomainAgreementsPreparer prepares the ListTopLevelDomainAgreements request.
func (client TopLevelDomainsClient) ListTopLevelDomainAgreementsPreparer(name string, agreementOption TopLevelDomainAgreementOption) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"name":           autorest.Encode("path", name),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-08-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.DomainRegistration/topLevelDomains/{name}/listAgreements", pathParameters),
		autorest.WithJSON(agreementOption),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListTopLevelDomainAgreementsSender sends the ListTopLevelDomainAgreements request. The method will close the
// http.Response Body if it receives an error.
func (client TopLevelDomainsClient) ListTopLevelDomainAgreementsSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListTopLevelDomainAgreementsResponder handles the response to the ListTopLevelDomainAgreements request. The method always
// closes the http.Response Body.
func (client TopLevelDomainsClient) ListTopLevelDomainAgreementsResponder(resp *http.Response) (result TldLegalAgreementCollection, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListTopLevelDomainAgreementsNextResults retrieves the next set of results, if any.
func (client TopLevelDomainsClient) ListTopLevelDomainAgreementsNextResults(lastResults TldLegalAgreementCollection) (result TldLegalAgreementCollection, err error) {
	req, err := lastResults.TldLegalAgreementCollectionPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "service.TopLevelDomainsClient", "ListTopLevelDomainAgreements", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListTopLevelDomainAgreementsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "service.TopLevelDomainsClient", "ListTopLevelDomainAgreements", resp, "Failure sending next results request")
	}

	result, err = client.ListTopLevelDomainAgreementsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "service.TopLevelDomainsClient", "ListTopLevelDomainAgreements", resp, "Failure responding to next results request")
	}

	return
}
