// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------
package google

import (
	"fmt"
)

type VPCAccessOperationWaiter struct {
	Config  *Config
	Project string
	CommonOperationWaiter
}

func (w *VPCAccessOperationWaiter) QueryOp() (interface{}, error) {
	if w == nil {
		return nil, fmt.Errorf("Cannot query operation, it's unset or nil.")
	}
	// Returns the proper get.
	url := fmt.Sprintf("https://vpcaccess.googleapis.com/v1/%s", w.CommonOperationWaiter.Op.Name)
	return sendRequest(w.Config, "GET", w.Project, url, nil)
}

func vpcAccessOperationWaitTime(config *Config, op map[string]interface{}, project, activity string, timeoutMinutes int) error {
	if val, ok := op["name"]; !ok || val == "" {
		// This was a synchronous call - there is no operation to wait for.
		return nil
	}
	w := &VPCAccessOperationWaiter{
		Config:  config,
		Project: project,
	}
	if err := w.CommonOperationWaiter.SetOp(op); err != nil {
		return err
	}
	return OperationWait(w, activity, timeoutMinutes)
}
