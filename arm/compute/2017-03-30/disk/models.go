package disk

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

// AccessLevel enumerates the values for access level.
type AccessLevel string

const (
	// None specifies the none state for access level.
	None AccessLevel = "None"
	// Read specifies the read state for access level.
	Read AccessLevel = "Read"
)

// CreateOption enumerates the values for create option.
type CreateOption string

const (
	// Attach specifies the attach state for create option.
	Attach CreateOption = "Attach"
	// Copy specifies the copy state for create option.
	Copy CreateOption = "Copy"
	// Empty specifies the empty state for create option.
	Empty CreateOption = "Empty"
	// FromImage specifies the from image state for create option.
	FromImage CreateOption = "FromImage"
	// Import specifies the import state for create option.
	Import CreateOption = "Import"
)

// OperatingSystemTypes enumerates the values for operating system types.
type OperatingSystemTypes string

const (
	// Linux specifies the linux state for operating system types.
	Linux OperatingSystemTypes = "Linux"
	// Windows specifies the windows state for operating system types.
	Windows OperatingSystemTypes = "Windows"
)

// StorageAccountTypes enumerates the values for storage account types.
type StorageAccountTypes string

const (
	// PremiumLRS specifies the premium lrs state for storage account types.
	PremiumLRS StorageAccountTypes = "Premium_LRS"
	// StandardLRS specifies the standard lrs state for storage account types.
	StandardLRS StorageAccountTypes = "Standard_LRS"
)

// AccessURI is a disk access SAS uri.
type AccessURI struct {
	autorest.Response `json:"-"`
	*AccessURIOutput  `json:"properties,omitempty"`
}

// AccessURIOutput is azure properties, including output.
type AccessURIOutput struct {
	*AccessURIRaw `json:"output,omitempty"`
}

// AccessURIRaw is this object gets 'bubbled up' through flattening.
type AccessURIRaw struct {
	AccessSAS *string `json:"accessSAS,omitempty"`
}

// APIError is api error.
type APIError struct {
	Details    *[]APIErrorBase `json:"details,omitempty"`
	Innererror *InnerError     `json:"innererror,omitempty"`
	Code       *string         `json:"code,omitempty"`
	Target     *string         `json:"target,omitempty"`
	Message    *string         `json:"message,omitempty"`
}

// APIErrorBase is api error base.
type APIErrorBase struct {
	Code    *string `json:"code,omitempty"`
	Target  *string `json:"target,omitempty"`
	Message *string `json:"message,omitempty"`
}

// CreationData is data used when creating a disk.
type CreationData struct {
	CreateOption     CreateOption        `json:"createOption,omitempty"`
	StorageAccountID *string             `json:"storageAccountId,omitempty"`
	ImageReference   *ImageDiskReference `json:"imageReference,omitempty"`
	SourceURI        *string             `json:"sourceUri,omitempty"`
	SourceResourceID *string             `json:"sourceResourceId,omitempty"`
}

// EncryptionSettings is encryption settings for disk or snapshot
type EncryptionSettings struct {
	Enabled           *bool                       `json:"enabled,omitempty"`
	DiskEncryptionKey *KeyVaultAndSecretReference `json:"diskEncryptionKey,omitempty"`
	KeyEncryptionKey  *KeyVaultAndKeyReference    `json:"keyEncryptionKey,omitempty"`
}

// GrantAccessData is data used for requesting a SAS.
type GrantAccessData struct {
	Access            AccessLevel `json:"access,omitempty"`
	DurationInSeconds *int32      `json:"durationInSeconds,omitempty"`
}

// ImageDiskReference is the source image used for creating the disk.
type ImageDiskReference struct {
	ID  *string `json:"id,omitempty"`
	Lun *int32  `json:"lun,omitempty"`
}

// InnerError is inner error details.
type InnerError struct {
	Exceptiontype *string `json:"exceptiontype,omitempty"`
	Errordetail   *string `json:"errordetail,omitempty"`
}

// KeyVaultAndKeyReference is key Vault Key Url and vault id of KeK, KeK is
// optional and when provided is used to unwrap the encryptionKey
type KeyVaultAndKeyReference struct {
	SourceVault *SourceVault `json:"sourceVault,omitempty"`
	KeyURL      *string      `json:"keyUrl,omitempty"`
}

// KeyVaultAndSecretReference is key Vault Secret Url and vault id of the
// encryption key
type KeyVaultAndSecretReference struct {
	SourceVault *SourceVault `json:"sourceVault,omitempty"`
	SecretURL   *string      `json:"secretUrl,omitempty"`
}

// ListType is the List Disks operation response.
type ListType struct {
	autorest.Response `json:"-"`
	Value             *[]Model `json:"value,omitempty"`
	NextLink          *string  `json:"nextLink,omitempty"`
}

// ListTypePreparer prepares a request to retrieve the next set of results. It returns
// nil if no more results exist.
func (client ListType) ListTypePreparer() (*http.Request, error) {
	if client.NextLink == nil || len(to.String(client.NextLink)) <= 0 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(client.NextLink)))
}

// Model is disk resource.
type Model struct {
	autorest.Response `json:"-"`
	ID                *string             `json:"id,omitempty"`
	Name              *string             `json:"name,omitempty"`
	Type              *string             `json:"type,omitempty"`
	Location          *string             `json:"location,omitempty"`
	Tags              *map[string]*string `json:"tags,omitempty"`
	ManagedBy         *string             `json:"managedBy,omitempty"`
	Sku               *Sku                `json:"sku,omitempty"`
	*Properties       `json:"properties,omitempty"`
}

// OperationStatusResponse is operation status response
type OperationStatusResponse struct {
	autorest.Response `json:"-"`
	Name              *string    `json:"name,omitempty"`
	Status            *string    `json:"status,omitempty"`
	StartTime         *date.Time `json:"startTime,omitempty"`
	EndTime           *date.Time `json:"endTime,omitempty"`
	Error             *APIError  `json:"error,omitempty"`
}

// Properties is disk resource properties.
type Properties struct {
	TimeCreated        *date.Time           `json:"timeCreated,omitempty"`
	OsType             OperatingSystemTypes `json:"osType,omitempty"`
	CreationData       *CreationData        `json:"creationData,omitempty"`
	DiskSizeGB         *int32               `json:"diskSizeGB,omitempty"`
	EncryptionSettings *EncryptionSettings  `json:"encryptionSettings,omitempty"`
	ProvisioningState  *string              `json:"provisioningState,omitempty"`
}

// Resource is the Resource model definition.
type Resource struct {
	ID       *string             `json:"id,omitempty"`
	Name     *string             `json:"name,omitempty"`
	Type     *string             `json:"type,omitempty"`
	Location *string             `json:"location,omitempty"`
	Tags     *map[string]*string `json:"tags,omitempty"`
}

// ResourceUpdate is the Resource model definition.
type ResourceUpdate struct {
	Tags *map[string]*string `json:"tags,omitempty"`
	Sku  *Sku                `json:"sku,omitempty"`
}

// Sku is the disks and snapshots sku name. Can be Standard_LRS or Premium_LRS.
type Sku struct {
	Name StorageAccountTypes `json:"name,omitempty"`
	Tier *string             `json:"tier,omitempty"`
}

// Snapshot is snapshot resource.
type Snapshot struct {
	autorest.Response `json:"-"`
	ID                *string             `json:"id,omitempty"`
	Name              *string             `json:"name,omitempty"`
	Type              *string             `json:"type,omitempty"`
	Location          *string             `json:"location,omitempty"`
	Tags              *map[string]*string `json:"tags,omitempty"`
	Sku               *Sku                `json:"sku,omitempty"`
	*Properties       `json:"properties,omitempty"`
}

// SnapshotList is the List Snapshots operation response.
type SnapshotList struct {
	autorest.Response `json:"-"`
	Value             *[]Snapshot `json:"value,omitempty"`
	NextLink          *string     `json:"nextLink,omitempty"`
}

// SnapshotListPreparer prepares a request to retrieve the next set of results. It returns
// nil if no more results exist.
func (client SnapshotList) SnapshotListPreparer() (*http.Request, error) {
	if client.NextLink == nil || len(to.String(client.NextLink)) <= 0 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(client.NextLink)))
}

// SnapshotUpdate is snapshot update resource.
type SnapshotUpdate struct {
	Tags              *map[string]*string `json:"tags,omitempty"`
	Sku               *Sku                `json:"sku,omitempty"`
	*UpdateProperties `json:"properties,omitempty"`
}

// SourceVault is the vault id is an Azure Resource Manager Resoure id in the
// form
// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.KeyVault/vaults/{vaultName}
type SourceVault struct {
	ID *string `json:"id,omitempty"`
}

// UpdateProperties is disk resource update properties.
type UpdateProperties struct {
	OsType             OperatingSystemTypes `json:"osType,omitempty"`
	DiskSizeGB         *int32               `json:"diskSizeGB,omitempty"`
	EncryptionSettings *EncryptionSettings  `json:"encryptionSettings,omitempty"`
}

// UpdateType is disk update resource.
type UpdateType struct {
	Tags              *map[string]*string `json:"tags,omitempty"`
	Sku               *Sku                `json:"sku,omitempty"`
	*UpdateProperties `json:"properties,omitempty"`
}
