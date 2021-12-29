package reportsummary

import "github.com/armosec/opa-utils/reporthandling/apis"

type ICounters interface {
	Excluded() int
	Passed() int
	Skipped() int
	Failed() int
	All() int

	Increase(status apis.IStatus)
}

// =================================== Counters ============================================

// NumberOfExcluded get the number of excluded resources
func (resourceCounters *ResourceCounters) Excluded() int {
	return resourceCounters.ExcludedResources
}

// NumberOfPassed get the number of passed resources
func (resourceCounters *ResourceCounters) Passed() int {
	return resourceCounters.PassedResources
}

// NumberOfSkipped get the number of skipped resources
func (resourceCounters *ResourceCounters) Skipped() int {
	return resourceCounters.SkippedResources
}

// NumberOfFailed get the number of failed resources
func (resourceCounters *ResourceCounters) Failed() int {
	return resourceCounters.FailedResources
}

// NumberOfAll get the number of all resources
func (resourceCounters *ResourceCounters) All() int {
	return resourceCounters.Excluded() + resourceCounters.Failed() + resourceCounters.Skipped() + resourceCounters.Passed()
}

// =================================== Setters ============================================

// Increase increases the counter based on the status
func (resourceCounters *ResourceCounters) Increase(status apis.IStatus) {
	switch status.Status() {
	case apis.StatusExcluded:
		resourceCounters.FailedResources++
	case apis.StatusFailed:
		resourceCounters.FailedResources++
	case apis.StatusSkipped:
		resourceCounters.SkippedResources++
	case apis.StatusPassed:
		resourceCounters.PassedResources++
	}
}
