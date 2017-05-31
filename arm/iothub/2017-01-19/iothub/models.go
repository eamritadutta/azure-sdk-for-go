package iothub

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for
// license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is
// regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/date"
	"github.com/Azure/go-autorest/autorest/to"
	"net/http"
)

// AccessRights enumerates the values for access rights.
type AccessRights string

const (
	// DeviceConnect specifies the device connect state for access rights.
	DeviceConnect AccessRights = "DeviceConnect"
	// RegistryRead specifies the registry read state for access rights.
	RegistryRead AccessRights = "RegistryRead"
	// RegistryReadDeviceConnect specifies the registry read device connect
	// state for access rights.
	RegistryReadDeviceConnect AccessRights = "RegistryRead, DeviceConnect"
	// RegistryReadRegistryWrite specifies the registry read registry write
	// state for access rights.
	RegistryReadRegistryWrite AccessRights = "RegistryRead, RegistryWrite"
	// RegistryReadRegistryWriteDeviceConnect specifies the registry read
	// registry write device connect state for access rights.
	RegistryReadRegistryWriteDeviceConnect AccessRights = "RegistryRead, RegistryWrite, DeviceConnect"
	// RegistryReadRegistryWriteServiceConnect specifies the registry read
	// registry write service connect state for access rights.
	RegistryReadRegistryWriteServiceConnect AccessRights = "RegistryRead, RegistryWrite, ServiceConnect"
	// RegistryReadRegistryWriteServiceConnectDeviceConnect specifies the
	// registry read registry write service connect device connect state for
	// access rights.
	RegistryReadRegistryWriteServiceConnectDeviceConnect AccessRights = "RegistryRead, RegistryWrite, ServiceConnect, DeviceConnect"
	// RegistryReadServiceConnect specifies the registry read service connect
	// state for access rights.
	RegistryReadServiceConnect AccessRights = "RegistryRead, ServiceConnect"
	// RegistryReadServiceConnectDeviceConnect specifies the registry read
	// service connect device connect state for access rights.
	RegistryReadServiceConnectDeviceConnect AccessRights = "RegistryRead, ServiceConnect, DeviceConnect"
	// RegistryWrite specifies the registry write state for access rights.
	RegistryWrite AccessRights = "RegistryWrite"
	// RegistryWriteDeviceConnect specifies the registry write device connect
	// state for access rights.
	RegistryWriteDeviceConnect AccessRights = "RegistryWrite, DeviceConnect"
	// RegistryWriteServiceConnect specifies the registry write service connect
	// state for access rights.
	RegistryWriteServiceConnect AccessRights = "RegistryWrite, ServiceConnect"
	// RegistryWriteServiceConnectDeviceConnect specifies the registry write
	// service connect device connect state for access rights.
	RegistryWriteServiceConnectDeviceConnect AccessRights = "RegistryWrite, ServiceConnect, DeviceConnect"
	// ServiceConnect specifies the service connect state for access rights.
	ServiceConnect AccessRights = "ServiceConnect"
	// ServiceConnectDeviceConnect specifies the service connect device connect
	// state for access rights.
	ServiceConnectDeviceConnect AccessRights = "ServiceConnect, DeviceConnect"
)

// Capabilities enumerates the values for capabilities.
type Capabilities string

const (
	// DeviceManagement specifies the device management state for capabilities.
	DeviceManagement Capabilities = "DeviceManagement"
	// None specifies the none state for capabilities.
	None Capabilities = "None"
)

// IPFilterActionType enumerates the values for ip filter action type.
type IPFilterActionType string

const (
	// Accept specifies the accept state for ip filter action type.
	Accept IPFilterActionType = "Accept"
	// Reject specifies the reject state for ip filter action type.
	Reject IPFilterActionType = "Reject"
)

// JobStatus enumerates the values for job status.
type JobStatus string

const (
	// Cancelled specifies the cancelled state for job status.
	Cancelled JobStatus = "cancelled"
	// Completed specifies the completed state for job status.
	Completed JobStatus = "completed"
	// Enqueued specifies the enqueued state for job status.
	Enqueued JobStatus = "enqueued"
	// Failed specifies the failed state for job status.
	Failed JobStatus = "failed"
	// Running specifies the running state for job status.
	Running JobStatus = "running"
	// Unknown specifies the unknown state for job status.
	Unknown JobStatus = "unknown"
)

// JobType enumerates the values for job type.
type JobType string

const (
	// JobTypeBackup specifies the job type backup state for job type.
	JobTypeBackup JobType = "backup"
	// JobTypeExport specifies the job type export state for job type.
	JobTypeExport JobType = "export"
	// JobTypeFactoryResetDevice specifies the job type factory reset device
	// state for job type.
	JobTypeFactoryResetDevice JobType = "factoryResetDevice"
	// JobTypeFirmwareUpdate specifies the job type firmware update state for
	// job type.
	JobTypeFirmwareUpdate JobType = "firmwareUpdate"
	// JobTypeImport specifies the job type import state for job type.
	JobTypeImport JobType = "import"
	// JobTypeReadDeviceProperties specifies the job type read device
	// properties state for job type.
	JobTypeReadDeviceProperties JobType = "readDeviceProperties"
	// JobTypeRebootDevice specifies the job type reboot device state for job
	// type.
	JobTypeRebootDevice JobType = "rebootDevice"
	// JobTypeUnknown specifies the job type unknown state for job type.
	JobTypeUnknown JobType = "unknown"
	// JobTypeUpdateDeviceConfiguration specifies the job type update device
	// configuration state for job type.
	JobTypeUpdateDeviceConfiguration JobType = "updateDeviceConfiguration"
	// JobTypeWriteDeviceProperties specifies the job type write device
	// properties state for job type.
	JobTypeWriteDeviceProperties JobType = "writeDeviceProperties"
)

// NameUnavailabilityReason enumerates the values for name unavailability
// reason.
type NameUnavailabilityReason string

const (
	// AlreadyExists specifies the already exists state for name unavailability
	// reason.
	AlreadyExists NameUnavailabilityReason = "AlreadyExists"
	// Invalid specifies the invalid state for name unavailability reason.
	Invalid NameUnavailabilityReason = "Invalid"
)

// OperationMonitoringLevel enumerates the values for operation monitoring
// level.
type OperationMonitoringLevel string

const (
	// OperationMonitoringLevelError specifies the operation monitoring level
	// error state for operation monitoring level.
	OperationMonitoringLevelError OperationMonitoringLevel = "Error"
	// OperationMonitoringLevelErrorInformation specifies the operation
	// monitoring level error information state for operation monitoring level.
	OperationMonitoringLevelErrorInformation OperationMonitoringLevel = "Error, Information"
	// OperationMonitoringLevelInformation specifies the operation monitoring
	// level information state for operation monitoring level.
	OperationMonitoringLevelInformation OperationMonitoringLevel = "Information"
	// OperationMonitoringLevelNone specifies the operation monitoring level
	// none state for operation monitoring level.
	OperationMonitoringLevelNone OperationMonitoringLevel = "None"
)

// RoutingSource enumerates the values for routing source.
type RoutingSource string

const (
	// DeviceJobLifecycleEvents specifies the device job lifecycle events state
	// for routing source.
	DeviceJobLifecycleEvents RoutingSource = "DeviceJobLifecycleEvents"
	// DeviceLifecycleEvents specifies the device lifecycle events state for
	// routing source.
	DeviceLifecycleEvents RoutingSource = "DeviceLifecycleEvents"
	// DeviceMessages specifies the device messages state for routing source.
	DeviceMessages RoutingSource = "DeviceMessages"
	// TwinChangeEvents specifies the twin change events state for routing
	// source.
	TwinChangeEvents RoutingSource = "TwinChangeEvents"
)

// ScaleType enumerates the values for scale type.
type ScaleType string

const (
	// ScaleTypeAutomatic specifies the scale type automatic state for scale
	// type.
	ScaleTypeAutomatic ScaleType = "Automatic"
	// ScaleTypeManual specifies the scale type manual state for scale type.
	ScaleTypeManual ScaleType = "Manual"
	// ScaleTypeNone specifies the scale type none state for scale type.
	ScaleTypeNone ScaleType = "None"
)

// Sku enumerates the values for sku.
type Sku string

const (
	// F1 specifies the f1 state for sku.
	F1 Sku = "F1"
	// S1 specifies the s1 state for sku.
	S1 Sku = "S1"
	// S2 specifies the s2 state for sku.
	S2 Sku = "S2"
	// S3 specifies the s3 state for sku.
	S3 Sku = "S3"
)

// SkuTier enumerates the values for sku tier.
type SkuTier string

const (
	// Free specifies the free state for sku tier.
	Free SkuTier = "Free"
	// Standard specifies the standard state for sku tier.
	Standard SkuTier = "Standard"
)

// Capacity is ioT Hub capacity information.
type Capacity struct {
	Minimum   *int64    `json:"minimum,omitempty"`
	Maximum   *int64    `json:"maximum,omitempty"`
	Default   *int64    `json:"default,omitempty"`
	ScaleType ScaleType `json:"scaleType,omitempty"`
}

// CloudToDeviceProperties is the IoT hub cloud-to-device messaging properties.
type CloudToDeviceProperties struct {
	MaxDeliveryCount    *int32              `json:"maxDeliveryCount,omitempty"`
	DefaultTTLAsIso8601 *string             `json:"defaultTtlAsIso8601,omitempty"`
	Feedback            *FeedbackProperties `json:"feedback,omitempty"`
}

// Description is the description of the IoT hub.
type Description struct {
	autorest.Response `json:"-"`
	ID                *string             `json:"id,omitempty"`
	Name              *string             `json:"name,omitempty"`
	Type              *string             `json:"type,omitempty"`
	Location          *string             `json:"location,omitempty"`
	Tags              *map[string]*string `json:"tags,omitempty"`
	Subscriptionid    *string             `json:"subscriptionid,omitempty"`
	Resourcegroup     *string             `json:"resourcegroup,omitempty"`
	Etag              *string             `json:"etag,omitempty"`
	Properties        *Properties         `json:"properties,omitempty"`
	Sku               *SkuInfo            `json:"sku,omitempty"`
}

// DescriptionListResult is the JSON-serialized array of IotHubDescription
// objects with a next link.
type DescriptionListResult struct {
	autorest.Response `json:"-"`
	Value             *[]Description `json:"value,omitempty"`
	NextLink          *string        `json:"nextLink,omitempty"`
}

// DescriptionListResultPreparer prepares a request to retrieve the next set of results. It returns
// nil if no more results exist.
func (client DescriptionListResult) DescriptionListResultPreparer() (*http.Request, error) {
	if client.NextLink == nil || len(to.String(client.NextLink)) <= 0 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(client.NextLink)))
}

// ErrorDetails is error details.
type ErrorDetails struct {
	Code           *string `json:"Code,omitempty"`
	HTTPStatusCode *string `json:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty"`
	Details        *string `json:"Details,omitempty"`
}

// EventHubConsumerGroupInfo is the properties of the EventHubConsumerGroupInfo
// object.
type EventHubConsumerGroupInfo struct {
	autorest.Response `json:"-"`
	Tags              *map[string]*string `json:"tags,omitempty"`
	ID                *string             `json:"id,omitempty"`
	Name              *string             `json:"name,omitempty"`
}

// EventHubConsumerGroupsListResult is the JSON-serialized array of Event
// Hub-compatible consumer group names with a next link.
type EventHubConsumerGroupsListResult struct {
	autorest.Response `json:"-"`
	Value             *[]string `json:"value,omitempty"`
	NextLink          *string   `json:"nextLink,omitempty"`
}

// EventHubConsumerGroupsListResultPreparer prepares a request to retrieve the next set of results. It returns
// nil if no more results exist.
func (client EventHubConsumerGroupsListResult) EventHubConsumerGroupsListResultPreparer() (*http.Request, error) {
	if client.NextLink == nil || len(to.String(client.NextLink)) <= 0 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(client.NextLink)))
}

// EventHubProperties is the properties of the provisioned Event Hub-compatible
// endpoint used by the IoT hub.
type EventHubProperties struct {
	RetentionTimeInDays *int64    `json:"retentionTimeInDays,omitempty"`
	PartitionCount      *int32    `json:"partitionCount,omitempty"`
	PartitionIds        *[]string `json:"partitionIds,omitempty"`
	Path                *string   `json:"path,omitempty"`
	Endpoint            *string   `json:"endpoint,omitempty"`
}

// ExportDevicesRequest is use to provide parameters when requesting an export
// of all devices in the IoT hub.
type ExportDevicesRequest struct {
	ExportBlobContainerURI *string `json:"ExportBlobContainerUri,omitempty"`
	ExcludeKeys            *bool   `json:"ExcludeKeys,omitempty"`
}

// FallbackRouteProperties is the properties related to the fallback route
// based on which the IoT hub routes messages to the fallback endpoint.
type FallbackRouteProperties struct {
	Source        *string   `json:"source,omitempty"`
	Condition     *string   `json:"condition,omitempty"`
	EndpointNames *[]string `json:"endpointNames,omitempty"`
	IsEnabled     *bool     `json:"isEnabled,omitempty"`
}

// FeedbackProperties is the properties of the feedback queue for
// cloud-to-device messages.
type FeedbackProperties struct {
	LockDurationAsIso8601 *string `json:"lockDurationAsIso8601,omitempty"`
	TTLAsIso8601          *string `json:"ttlAsIso8601,omitempty"`
	MaxDeliveryCount      *int32  `json:"maxDeliveryCount,omitempty"`
}

// ImportDevicesRequest is use to provide parameters when requesting an import
// of all devices in the hub.
type ImportDevicesRequest struct {
	InputBlobContainerURI  *string `json:"InputBlobContainerUri,omitempty"`
	OutputBlobContainerURI *string `json:"OutputBlobContainerUri,omitempty"`
}

// IPFilterRule is the IP filter rules for the IoT hub.
type IPFilterRule struct {
	FilterName *string            `json:"filterName,omitempty"`
	Action     IPFilterActionType `json:"action,omitempty"`
	IPMask     *string            `json:"ipMask,omitempty"`
}

// JobResponse is the properties of the Job Response object.
type JobResponse struct {
	autorest.Response `json:"-"`
	JobID             *string           `json:"jobId,omitempty"`
	StartTimeUtc      *date.TimeRFC1123 `json:"startTimeUtc,omitempty"`
	EndTimeUtc        *date.TimeRFC1123 `json:"endTimeUtc,omitempty"`
	Type              JobType           `json:"type,omitempty"`
	Status            JobStatus         `json:"status,omitempty"`
	FailureReason     *string           `json:"failureReason,omitempty"`
	StatusMessage     *string           `json:"statusMessage,omitempty"`
	ParentJobID       *string           `json:"parentJobId,omitempty"`
}

// JobResponseListResult is the JSON-serialized array of JobResponse objects
// with a next link.
type JobResponseListResult struct {
	autorest.Response `json:"-"`
	Value             *[]JobResponse `json:"value,omitempty"`
	NextLink          *string        `json:"nextLink,omitempty"`
}

// JobResponseListResultPreparer prepares a request to retrieve the next set of results. It returns
// nil if no more results exist.
func (client JobResponseListResult) JobResponseListResultPreparer() (*http.Request, error) {
	if client.NextLink == nil || len(to.String(client.NextLink)) <= 0 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(client.NextLink)))
}

// MessagingEndpointProperties is the properties of the messaging endpoints
// used by this IoT hub.
type MessagingEndpointProperties struct {
	LockDurationAsIso8601 *string `json:"lockDurationAsIso8601,omitempty"`
	TTLAsIso8601          *string `json:"ttlAsIso8601,omitempty"`
	MaxDeliveryCount      *int32  `json:"maxDeliveryCount,omitempty"`
}

// NameAvailabilityInfo is the properties indicating whether a given IoT hub
// name is available.
type NameAvailabilityInfo struct {
	autorest.Response `json:"-"`
	NameAvailable     *bool                    `json:"nameAvailable,omitempty"`
	Reason            NameUnavailabilityReason `json:"reason,omitempty"`
	Message           *string                  `json:"message,omitempty"`
}

// OperationInputs is input values.
type OperationInputs struct {
	Name *string `json:"Name,omitempty"`
}

// OperationsMonitoringProperties is the operations monitoring properties for
// the IoT hub. The possible keys to the dictionary are Connections,
// DeviceTelemetry, C2DCommands, DeviceIdentityOperations,
// FileUploadOperations, Routes, D2CTwinOperations, C2DTwinOperations,
// TwinQueries, JobsOperations, DirectMethods.
type OperationsMonitoringProperties struct {
	Events *map[string]*OperationMonitoringLevel `json:"events,omitempty"`
}

// Properties is the properties of an IoT hub.
type Properties struct {
	AuthorizationPolicies          *[]SharedAccessSignatureAuthorizationRule `json:"authorizationPolicies,omitempty"`
	IPFilterRules                  *[]IPFilterRule                           `json:"ipFilterRules,omitempty"`
	ProvisioningState              *string                                   `json:"provisioningState,omitempty"`
	HostName                       *string                                   `json:"hostName,omitempty"`
	EventHubEndpoints              *map[string]*EventHubProperties           `json:"eventHubEndpoints,omitempty"`
	Routing                        *RoutingProperties                        `json:"routing,omitempty"`
	StorageEndpoints               *map[string]*StorageEndpointProperties    `json:"storageEndpoints,omitempty"`
	MessagingEndpoints             *map[string]*MessagingEndpointProperties  `json:"messagingEndpoints,omitempty"`
	EnableFileUploadNotifications  *bool                                     `json:"enableFileUploadNotifications,omitempty"`
	CloudToDevice                  *CloudToDeviceProperties                  `json:"cloudToDevice,omitempty"`
	Comments                       *string                                   `json:"comments,omitempty"`
	OperationsMonitoringProperties *OperationsMonitoringProperties           `json:"operationsMonitoringProperties,omitempty"`
	Features                       Capabilities                              `json:"features,omitempty"`
}

// QuotaMetricInfo is quota metrics properties.
type QuotaMetricInfo struct {
	Name         *string `json:"Name,omitempty"`
	CurrentValue *int64  `json:"CurrentValue,omitempty"`
	MaxValue     *int64  `json:"MaxValue,omitempty"`
}

// QuotaMetricInfoListResult is the JSON-serialized array of
// IotHubQuotaMetricInfo objects with a next link.
type QuotaMetricInfoListResult struct {
	autorest.Response `json:"-"`
	Value             *[]QuotaMetricInfo `json:"value,omitempty"`
	NextLink          *string            `json:"nextLink,omitempty"`
}

// QuotaMetricInfoListResultPreparer prepares a request to retrieve the next set of results. It returns
// nil if no more results exist.
func (client QuotaMetricInfoListResult) QuotaMetricInfoListResultPreparer() (*http.Request, error) {
	if client.NextLink == nil || len(to.String(client.NextLink)) <= 0 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(client.NextLink)))
}

// RegistryStatistics is identity registry statistics.
type RegistryStatistics struct {
	autorest.Response   `json:"-"`
	TotalDeviceCount    *int64 `json:"totalDeviceCount,omitempty"`
	EnabledDeviceCount  *int64 `json:"enabledDeviceCount,omitempty"`
	DisabledDeviceCount *int64 `json:"disabledDeviceCount,omitempty"`
}

// Resource is the common properties of an Azure resource.
type Resource struct {
	ID       *string             `json:"id,omitempty"`
	Name     *string             `json:"name,omitempty"`
	Type     *string             `json:"type,omitempty"`
	Location *string             `json:"location,omitempty"`
	Tags     *map[string]*string `json:"tags,omitempty"`
}

// RouteProperties is the properties of a routing rule that your IoT hub uses
// to route messages to endpoints.
type RouteProperties struct {
	Name          *string       `json:"name,omitempty"`
	Source        RoutingSource `json:"source,omitempty"`
	Condition     *string       `json:"condition,omitempty"`
	EndpointNames *[]string     `json:"endpointNames,omitempty"`
	IsEnabled     *bool         `json:"isEnabled,omitempty"`
}

// RoutingEndpoints is the properties related to the custom endpoints to which
// your IoT hub routes messages based on the routing rules. A maximum of 10
// custom endpoints are allowed across all endpoint types for paid hubs and
// only 1 custom endpoint is allowed across all endpoint types for free hubs.
type RoutingEndpoints struct {
	ServiceBusQueues *[]RoutingServiceBusQueueEndpointProperties `json:"serviceBusQueues,omitempty"`
	ServiceBusTopics *[]RoutingServiceBusTopicEndpointProperties `json:"serviceBusTopics,omitempty"`
	EventHubs        *[]RoutingEventHubProperties                `json:"eventHubs,omitempty"`
}

// RoutingEventHubProperties is the properties related to an event hub
// endpoint.
type RoutingEventHubProperties struct {
	ConnectionString *string `json:"connectionString,omitempty"`
	Name             *string `json:"name,omitempty"`
	SubscriptionID   *string `json:"subscriptionId,omitempty"`
	ResourceGroup    *string `json:"resourceGroup,omitempty"`
}

// RoutingProperties is the routing related properties of the IoT hub. See:
// https://docs.microsoft.com/azure/iot-hub/iot-hub-devguide-messaging
type RoutingProperties struct {
	Endpoints     *RoutingEndpoints        `json:"endpoints,omitempty"`
	Routes        *[]RouteProperties       `json:"routes,omitempty"`
	FallbackRoute *FallbackRouteProperties `json:"fallbackRoute,omitempty"`
}

// RoutingServiceBusQueueEndpointProperties is the properties related to
// service bus queue endpoint types.
type RoutingServiceBusQueueEndpointProperties struct {
	ConnectionString *string `json:"connectionString,omitempty"`
	Name             *string `json:"name,omitempty"`
	SubscriptionID   *string `json:"subscriptionId,omitempty"`
	ResourceGroup    *string `json:"resourceGroup,omitempty"`
}

// RoutingServiceBusTopicEndpointProperties is the properties related to
// service bus topic endpoint types.
type RoutingServiceBusTopicEndpointProperties struct {
	ConnectionString *string `json:"connectionString,omitempty"`
	Name             *string `json:"name,omitempty"`
	SubscriptionID   *string `json:"subscriptionId,omitempty"`
	ResourceGroup    *string `json:"resourceGroup,omitempty"`
}

// SetObject is
type SetObject struct {
	autorest.Response `json:"-"`
	Value             *map[string]interface{} `json:"value,omitempty"`
}

// SharedAccessSignatureAuthorizationRule is the properties of an IoT hub
// shared access policy.
type SharedAccessSignatureAuthorizationRule struct {
	autorest.Response `json:"-"`
	KeyName           *string      `json:"keyName,omitempty"`
	PrimaryKey        *string      `json:"primaryKey,omitempty"`
	SecondaryKey      *string      `json:"secondaryKey,omitempty"`
	Rights            AccessRights `json:"rights,omitempty"`
}

// SharedAccessSignatureAuthorizationRuleListResult is the list of shared
// access policies with a next link.
type SharedAccessSignatureAuthorizationRuleListResult struct {
	autorest.Response `json:"-"`
	Value             *[]SharedAccessSignatureAuthorizationRule `json:"value,omitempty"`
	NextLink          *string                                   `json:"nextLink,omitempty"`
}

// SharedAccessSignatureAuthorizationRuleListResultPreparer prepares a request to retrieve the next set of results. It returns
// nil if no more results exist.
func (client SharedAccessSignatureAuthorizationRuleListResult) SharedAccessSignatureAuthorizationRuleListResultPreparer() (*http.Request, error) {
	if client.NextLink == nil || len(to.String(client.NextLink)) <= 0 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(client.NextLink)))
}

// SkuDescription is sKU properties.
type SkuDescription struct {
	ResourceType *string   `json:"resourceType,omitempty"`
	Sku          *SkuInfo  `json:"sku,omitempty"`
	Capacity     *Capacity `json:"capacity,omitempty"`
}

// SkuDescriptionListResult is the JSON-serialized array of
// IotHubSkuDescription objects with a next link.
type SkuDescriptionListResult struct {
	autorest.Response `json:"-"`
	Value             *[]SkuDescription `json:"value,omitempty"`
	NextLink          *string           `json:"nextLink,omitempty"`
}

// SkuDescriptionListResultPreparer prepares a request to retrieve the next set of results. It returns
// nil if no more results exist.
func (client SkuDescriptionListResult) SkuDescriptionListResultPreparer() (*http.Request, error) {
	if client.NextLink == nil || len(to.String(client.NextLink)) <= 0 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(client.NextLink)))
}

// SkuInfo is information about the SKU of the IoT hub.
type SkuInfo struct {
	Name     Sku     `json:"name,omitempty"`
	Tier     SkuTier `json:"tier,omitempty"`
	Capacity *int64  `json:"capacity,omitempty"`
}

// StorageEndpointProperties is the properties of the Azure Storage endpoint
// for file upload.
type StorageEndpointProperties struct {
	SasTTLAsIso8601  *string `json:"sasTtlAsIso8601,omitempty"`
	ConnectionString *string `json:"connectionString,omitempty"`
	ContainerName    *string `json:"containerName,omitempty"`
}
