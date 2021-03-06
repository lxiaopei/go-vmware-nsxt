/* Copyright © 2017 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause

   Generated by: https://github.com/swagger-api/swagger-codegen.git */

package manager

type LogicalSwitchState struct {

	// Array of configuration state of various sub systems
	Details []ConfigurationStateElement `json:"details,omitempty"`

	// Error code
	FailureCode int64 `json:"failure_code,omitempty"`

	// Error message in case of failure
	FailureMessage string `json:"failure_message,omitempty"`

	// Gives details of state of desired configuration
	State string `json:"state,omitempty"`

	// Id of the logical switch
	LogicalSwitchId string `json:"logical_switch_id,omitempty"`
}
