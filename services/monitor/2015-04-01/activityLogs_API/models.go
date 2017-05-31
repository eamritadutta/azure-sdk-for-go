package activitylogsapi

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
	"github.com/Azure/go-autorest/autorest/date"
	"github.com/Azure/go-autorest/autorest/to"
	"net/http"
)

// EventLevel enumerates the values for event level.
type EventLevel string

const (
	// Critical specifies the critical state for event level.
	Critical EventLevel = "Critical"
	// Error specifies the error state for event level.
	Error EventLevel = "Error"
	// Informational specifies the informational state for event level.
	Informational EventLevel = "Informational"
	// Verbose specifies the verbose state for event level.
	Verbose EventLevel = "Verbose"
	// Warning specifies the warning state for event level.
	Warning EventLevel = "Warning"
)

// ErrorResponse is describes the format of Error response.
type ErrorResponse struct {
	Code    *string `json:"code,omitempty"`
	Message *string `json:"message,omitempty"`
}

// EventData is the Azure event log entries are of type EventData
type EventData struct {
	Authorization        *SenderAuthorization `json:"authorization,omitempty"`
	Claims               *map[string]*string  `json:"claims,omitempty"`
	Caller               *string              `json:"caller,omitempty"`
	Description          *string              `json:"description,omitempty"`
	ID                   *string              `json:"id,omitempty"`
	EventDataID          *string              `json:"eventDataId,omitempty"`
	CorrelationID        *string              `json:"correlationId,omitempty"`
	EventName            *LocalizableString   `json:"eventName,omitempty"`
	Category             *LocalizableString   `json:"category,omitempty"`
	HTTPRequest          *HTTPRequestInfo     `json:"httpRequest,omitempty"`
	Level                EventLevel           `json:"level,omitempty"`
	ResourceGroupName    *string              `json:"resourceGroupName,omitempty"`
	ResourceProviderName *LocalizableString   `json:"resourceProviderName,omitempty"`
	ResourceID           *string              `json:"resourceId,omitempty"`
	ResourceType         *LocalizableString   `json:"resourceType,omitempty"`
	OperationID          *string              `json:"operationId,omitempty"`
	OperationName        *LocalizableString   `json:"operationName,omitempty"`
	Properties           *map[string]*string  `json:"properties,omitempty"`
	Status               *LocalizableString   `json:"status,omitempty"`
	SubStatus            *LocalizableString   `json:"subStatus,omitempty"`
	EventTimestamp       *date.Time           `json:"eventTimestamp,omitempty"`
	SubmissionTimestamp  *date.Time           `json:"submissionTimestamp,omitempty"`
	SubscriptionID       *string              `json:"subscriptionId,omitempty"`
	TenantID             *string              `json:"tenantId,omitempty"`
}

// EventDataCollection is represents collection of events.
type EventDataCollection struct {
	autorest.Response `json:"-"`
	Value             *[]EventData `json:"value,omitempty"`
	NextLink          *string      `json:"nextLink,omitempty"`
}

// EventDataCollectionPreparer prepares a request to retrieve the next set of results. It returns
// nil if no more results exist.
func (client EventDataCollection) EventDataCollectionPreparer() (*http.Request, error) {
	if client.NextLink == nil || len(to.String(client.NextLink)) <= 0 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(client.NextLink)))
}

// HTTPRequestInfo is the Http request info.
type HTTPRequestInfo struct {
	ClientRequestID *string `json:"clientRequestId,omitempty"`
	ClientIPAddress *string `json:"clientIpAddress,omitempty"`
	Method          *string `json:"method,omitempty"`
	URI             *string `json:"uri,omitempty"`
}

// LocalizableString is the localizable string class.
type LocalizableString struct {
	Value          *string `json:"value,omitempty"`
	LocalizedValue *string `json:"localizedValue,omitempty"`
}

// SenderAuthorization is the authorization used by the user who has performed
// the operation that led to this event. This captures the RBAC properties of
// the event. These usually include the 'action', 'role' and the 'scope'
type SenderAuthorization struct {
	Action *string `json:"action,omitempty"`
	Role   *string `json:"role,omitempty"`
	Scope  *string `json:"scope,omitempty"`
}
