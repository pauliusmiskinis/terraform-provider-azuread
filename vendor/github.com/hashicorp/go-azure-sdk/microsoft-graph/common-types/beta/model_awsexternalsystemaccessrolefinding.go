package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Finding = AwsExternalSystemAccessRoleFinding{}

type AwsExternalSystemAccessRoleFinding struct {
	// The IDs of the accounts that this role is able to access.
	AccessibleSystemIds *[]string `json:"accessibleSystemIds,omitempty"`

	PermissionsCreepIndex *PermissionsCreepIndex `json:"permissionsCreepIndex,omitempty"`
	Role                  *AwsRole               `json:"role,omitempty"`

	// Fields inherited from Finding

	// Defines when the finding was created.
	CreatedDateTime *string `json:"createdDateTime,omitempty"`

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s AwsExternalSystemAccessRoleFinding) Finding() BaseFindingImpl {
	return BaseFindingImpl{
		CreatedDateTime: s.CreatedDateTime,
		Id:              s.Id,
		ODataId:         s.ODataId,
		ODataType:       s.ODataType,
	}
}

func (s AwsExternalSystemAccessRoleFinding) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = AwsExternalSystemAccessRoleFinding{}

func (s AwsExternalSystemAccessRoleFinding) MarshalJSON() ([]byte, error) {
	type wrapper AwsExternalSystemAccessRoleFinding
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AwsExternalSystemAccessRoleFinding: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AwsExternalSystemAccessRoleFinding: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.awsExternalSystemAccessRoleFinding"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AwsExternalSystemAccessRoleFinding: %+v", err)
	}

	return encoded, nil
}