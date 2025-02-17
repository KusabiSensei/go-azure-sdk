package changedatacapture

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MapperSourceConnectionsInfo struct {
	Connection     *MapperConnection `json:"Connection,omitempty"`
	SourceEntities *[]MapperTable    `json:"SourceEntities,omitempty"`
}
