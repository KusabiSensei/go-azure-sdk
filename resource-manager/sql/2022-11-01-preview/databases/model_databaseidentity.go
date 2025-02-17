package databases

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DatabaseIdentity struct {
	TenantId               *string                          `json:"tenantId,omitempty"`
	Type                   *DatabaseIdentityType            `json:"type,omitempty"`
	UserAssignedIdentities *map[string]DatabaseUserIdentity `json:"userAssignedIdentities,omitempty"`
}
