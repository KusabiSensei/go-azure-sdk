package wcfrelays

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessRights string

const (
	AccessRightsListen AccessRights = "Listen"
	AccessRightsManage AccessRights = "Manage"
	AccessRightsSend   AccessRights = "Send"
)

func PossibleValuesForAccessRights() []string {
	return []string{
		string(AccessRightsListen),
		string(AccessRightsManage),
		string(AccessRightsSend),
	}
}

func parseAccessRights(input string) (*AccessRights, error) {
	vals := map[string]AccessRights{
		"listen": AccessRightsListen,
		"manage": AccessRightsManage,
		"send":   AccessRightsSend,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AccessRights(input)
	return &out, nil
}

type KeyType string

const (
	KeyTypePrimaryKey   KeyType = "PrimaryKey"
	KeyTypeSecondaryKey KeyType = "SecondaryKey"
)

func PossibleValuesForKeyType() []string {
	return []string{
		string(KeyTypePrimaryKey),
		string(KeyTypeSecondaryKey),
	}
}

func parseKeyType(input string) (*KeyType, error) {
	vals := map[string]KeyType{
		"primarykey":   KeyTypePrimaryKey,
		"secondarykey": KeyTypeSecondaryKey,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := KeyType(input)
	return &out, nil
}

type Relaytype string

const (
	RelaytypeHTTP   Relaytype = "Http"
	RelaytypeNetTcp Relaytype = "NetTcp"
)

func PossibleValuesForRelaytype() []string {
	return []string{
		string(RelaytypeHTTP),
		string(RelaytypeNetTcp),
	}
}

func parseRelaytype(input string) (*Relaytype, error) {
	vals := map[string]Relaytype{
		"http":   RelaytypeHTTP,
		"nettcp": RelaytypeNetTcp,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Relaytype(input)
	return &out, nil
}
