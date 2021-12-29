package reportsummary

import (
	"github.com/armosec/opa-utils/reporthandling/apis"
	helpersv1 "github.com/armosec/opa-utils/reporthandling/helpers/v1"
)

// =================================== Status ============================================

// GetStatus get the control status. returns an apis.ScanningStatus object
func (controlSummary *ControlSummary) GetStatus() apis.IStatus {
	if controlSummary.Status == apis.StatusUnknown {
		controlSummary.CalculateStatus()
	}
	return helpersv1.NewStatus(controlSummary.Status)
}

// CalculateStatus set the control status based on the resource counters
func (controlSummary *ControlSummary) CalculateStatus() {
	controlSummary.Status = calculateStatus(&controlSummary.ResourceCounters)
}

// =================================== Counters ============================================

// =================================== Counters ============================================

// NumberOfExcluded get the number of excluded resources
func (controlSummary *ControlSummary) NumberOf() ICounters {
	return &controlSummary.ResourceCounters
}

// Increase increases the counter based on the status
func (controlSummary *ControlSummary) Increase(status apis.IStatus) {
	controlSummary.ResourceCounters.Increase(status)
}

// =================================== Score ============================================

// GetScore return control score
func (controlSummary *ControlSummary) GetScore() float32 {
	return controlSummary.Score
}

// =================================== Name ============================================

// GetName return control name
func (controlSummary *ControlSummary) GetName() string {
	return controlSummary.Name
}

// GetName return control ID
func (controlSummary *ControlSummary) GetID() string {
	return controlSummary.Name
}