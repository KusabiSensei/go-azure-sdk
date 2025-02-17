package inventoryitems

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InventoryItemProperties interface {
}

func unmarshalInventoryItemPropertiesImplementation(input []byte) (InventoryItemProperties, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling InventoryItemProperties into map[string]interface: %+v", err)
	}

	value, ok := temp["inventoryType"].(string)
	if !ok {
		return nil, nil
	}

	if strings.EqualFold(value, "Cluster") {
		var out ClusterInventoryItem
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ClusterInventoryItem: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "Datastore") {
		var out DatastoreInventoryItem
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DatastoreInventoryItem: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "Host") {
		var out HostInventoryItem
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into HostInventoryItem: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "ResourcePool") {
		var out ResourcePoolInventoryItem
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ResourcePoolInventoryItem: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "VirtualMachine") {
		var out VirtualMachineInventoryItem
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into VirtualMachineInventoryItem: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "VirtualMachineTemplate") {
		var out VirtualMachineTemplateInventoryItem
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into VirtualMachineTemplateInventoryItem: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "VirtualNetwork") {
		var out VirtualNetworkInventoryItem
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into VirtualNetworkInventoryItem: %+v", err)
		}
		return out, nil
	}

	type RawInventoryItemPropertiesImpl struct {
		Type   string                 `json:"-"`
		Values map[string]interface{} `json:"-"`
	}
	out := RawInventoryItemPropertiesImpl{
		Type:   value,
		Values: temp,
	}
	return out, nil

}
