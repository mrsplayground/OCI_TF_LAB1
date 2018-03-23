// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package core

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// GetVolumeBackupPolicyRequest wrapper for the GetVolumeBackupPolicy operation
type GetVolumeBackupPolicyRequest struct {

	// The OCID of the volume backup policy.
	PolicyId *string `mandatory:"true" contributesTo:"path" name:"policyId"`
}

func (request GetVolumeBackupPolicyRequest) String() string {
	return common.PointerString(request)
}

// GetVolumeBackupPolicyResponse wrapper for the GetVolumeBackupPolicy operation
type GetVolumeBackupPolicyResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The VolumeBackupPolicy instance
	VolumeBackupPolicy `presentIn:"body"`

	// For optimistic concurrency control. See `if-match`.
	Etag *string `presentIn:"header" name:"etag"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetVolumeBackupPolicyResponse) String() string {
	return common.PointerString(response)
}