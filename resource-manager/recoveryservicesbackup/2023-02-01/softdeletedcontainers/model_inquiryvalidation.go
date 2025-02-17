package softdeletedcontainers

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InquiryValidation struct {
	AdditionalDetail *string      `json:"additionalDetail,omitempty"`
	ErrorDetail      *ErrorDetail `json:"errorDetail,omitempty"`
	Status           *string      `json:"status,omitempty"`
}
