package alertrulesapis

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RuleAction interface {
}

func unmarshalRuleActionImplementation(input []byte) (RuleAction, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling RuleAction into map[string]interface: %+v", err)
	}

	value, ok := temp["odata.type"].(string)
	if !ok {
		return nil, nil
	}

	if strings.EqualFold(value, "Microsoft.Azure.Management.Insights.Models.RuleEmailAction") {
		var out RuleEmailAction
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into RuleEmailAction: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "Microsoft.Azure.Management.Insights.Models.RuleWebhookAction") {
		var out RuleWebhookAction
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into RuleWebhookAction: %+v", err)
		}
		return out, nil
	}

	type RawRuleActionImpl struct {
		Type   string                 `json:"-"`
		Values map[string]interface{} `json:"-"`
	}
	out := RawRuleActionImpl{
		Type:   value,
		Values: temp,
	}
	return out, nil

}
