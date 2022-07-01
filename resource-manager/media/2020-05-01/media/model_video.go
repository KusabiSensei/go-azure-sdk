package media

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Codec = Video{}

type Video struct {
	KeyFrameInterval *string        `json:"keyFrameInterval,omitempty"`
	StretchMode      *StretchMode   `json:"stretchMode,omitempty"`
	SyncMode         *VideoSyncMode `json:"syncMode,omitempty"`

	// Fields inherited from Codec
	Label *string `json:"label,omitempty"`
}

var _ json.Marshaler = Video{}

func (s Video) MarshalJSON() ([]byte, error) {
	type wrapper Video
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling Video: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling Video: %+v", err)
	}
	decoded["@odata.type"] = "#Microsoft.Media.Video"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling Video: %+v", err)
	}

	return encoded, nil
}
