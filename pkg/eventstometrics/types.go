// Code generated by tutone: DO NOT EDIT
package eventstometrics

import (
	"github.com/cegeka-everesst/newrelic-client-go/pkg/nrtime"
)

// EventsToMetricsErrorReason - General error categories.
type EventsToMetricsErrorReason string

var EventsToMetricsErrorReasonTypes = struct {
	// Other errors.
	GENERAL EventsToMetricsErrorReason
	// Indicates some part of your submission was invalid.
	INVALID_INPUT EventsToMetricsErrorReason
	// The user attempting to submit this rule is not authorized to do so.
	USER_NOT_AUTHORIZED EventsToMetricsErrorReason
}{
	// Other errors.
	GENERAL: "GENERAL",
	// Indicates some part of your submission was invalid.
	INVALID_INPUT: "INVALID_INPUT",
	// The user attempting to submit this rule is not authorized to do so.
	USER_NOT_AUTHORIZED: "USER_NOT_AUTHORIZED",
}

// EventsToMetricsAccountStitchedFields - Account stitched fields to enable autostitching in NerdGraph
type EventsToMetricsAccountStitchedFields struct {
	// List all rules for your account.
	AllRules EventsToMetricsListRuleResult `json:"allRules,omitempty"`
	// List rules for your account by id.
	RulesById EventsToMetricsListRuleResult `json:"rulesById,omitempty"`
}

// EventsToMetricsCreateRuleFailure - Error details about the events to metrics rule that failed to be created and why.
type EventsToMetricsCreateRuleFailure struct {
	// Information about why the create failed.
	Errors []EventsToMetricsError `json:"errors,omitempty"`
	// Input information about a submitted rule that was unable to be created.
	Submitted EventsToMetricsCreateRuleSubmission `json:"submitted,omitempty"`
}

// EventsToMetricsCreateRuleInput - Details needed to create an events to metrics conversion rule.
type EventsToMetricsCreateRuleInput struct {
	// The account where the events exist and the metrics will be put.
	AccountID int `json:"accountId"`
	// Provides additional information about the rule.
	Description string `json:"description,omitempty"`
	// Explains how to create one or more metrics from events.
	NRQL string `json:"nrql"`
	// The name of the rule. This must be unique within a given account.
	Name string `json:"name"`
}

// EventsToMetricsCreateRuleResult - The result of which submitted events to metrics rules were successfully and unsuccessfully created
type EventsToMetricsCreateRuleResult struct {
	// Rules that were not created and why.
	Failures []EventsToMetricsCreateRuleFailure `json:"failures,omitempty"`
	// Rules that were successfully created.
	Successes []EventsToMetricsRule `json:"successes,omitempty"`
}

// EventsToMetricsCreateRuleSubmission - The details that were submitted when creating an events to metrics conversion rule.
type EventsToMetricsCreateRuleSubmission struct {
	// The account where the events exist and the metrics will be put.
	AccountID int `json:"accountId"`
	// Provides additional information about the rule.
	Description string `json:"description,omitempty"`
	// Explains how to create one or more metrics from events.
	NRQL string `json:"nrql"`
	// The name of the rule. This must be unique within a given account.
	Name string `json:"name"`
}

// EventsToMetricsDeleteRuleFailure - Error details about the events to metrics rule that failed to be deleted and why.
type EventsToMetricsDeleteRuleFailure struct {
	// Information about why the delete failed.
	Errors []EventsToMetricsError `json:"errors,omitempty"`
	// Input information about a submitted rule that was unable to be deleted.
	Submitted EventsToMetricsDeleteRuleSubmission `json:"submitted,omitempty"`
}

// EventsToMetricsDeleteRuleInput - Identifying information about the events to metrics rule you want to delete.
type EventsToMetricsDeleteRuleInput struct {
	// A submitted account id.
	AccountID int `json:"accountId"`
	// A submitted rule id.
	RuleId string `json:"ruleId"`
}

// EventsToMetricsDeleteRuleResult - The result of which submitted events to metrics rules were successfully and unsuccessfully deleted.
type EventsToMetricsDeleteRuleResult struct {
	// Information about the rules that could not be deleted.
	Failures []EventsToMetricsDeleteRuleFailure `json:"failures,omitempty"`
	// Rules that were successfully deleted.
	Successes []EventsToMetricsRule `json:"successes,omitempty"`
}

// EventsToMetricsDeleteRuleSubmission - The details that were submitted when deleteing an events to metrics conversion rule.
type EventsToMetricsDeleteRuleSubmission struct {
	// A submitted account id.
	AccountID int `json:"accountId"`
	// A submitted rule id.
	RuleId string `json:"ruleId"`
}

// EventsToMetricsError - Error details when processing events to metrics rule requests.
type EventsToMetricsError struct {
	// A detailed error message.
	Description string `json:"description,omitempty"`
	// The category of error that occurred.
	Reason EventsToMetricsErrorReason `json:"reason,omitempty"`
}

// EventsToMetricsListRuleResult - A list of rule details to be returned.
type EventsToMetricsListRuleResult struct {
	// Event-to-metric rules to be returned.
	Rules []EventsToMetricsRule `json:"rules,omitempty"`
}

// EventsToMetricsRule - Information about an event-to-metric rule which creates metrics from events.
type EventsToMetricsRule struct {
	// Account with the event and where the metrics will be placed.
	AccountID int `json:"accountId"`
	// The time at which the rule was created
	CreatedAt nrtime.DateTime `json:"createdAt"`
	// Additional information about the rule.
	Description string `json:"description,omitempty"`
	// True means this rule is enabled. False means the rule is currently not creating metrics.
	Enabled bool `json:"enabled"`
	// The id, uniquely identifying the rule.
	ID string `json:"id"`
	// Explains how to create metrics from events.
	NRQL string `json:"nrql"`
	// The name of the rule. This must be unique within an account.
	Name string `json:"name"`
	// The time at which the rule was updated
	UpdatedAt nrtime.DateTime `json:"updatedAt"`
}

// EventsToMetricsUpdateRuleFailure - Error details about the events to metrics rule that failed to be updated and why.
type EventsToMetricsUpdateRuleFailure struct {
	// Information about why the update failed.
	Errors []EventsToMetricsError `json:"errors,omitempty"`
	// Input information about a failed update.
	Submitted EventsToMetricsUpdateRuleSubmission `json:"submitted,omitempty"`
}

// EventsToMetricsUpdateRuleInput - Identifying information about the events to metrics rule you want to update.
type EventsToMetricsUpdateRuleInput struct {
	// A submitted account id.
	AccountID int `json:"accountId"`
	// Changes the state of the rule as being enabled or disabled.
	Enabled bool `json:"enabled"`
	// A submitted rule id.
	RuleId string `json:"ruleId"`
}

// EventsToMetricsUpdateRuleResult - The result of which submitted events to metrics rules were successfully and unsuccessfully update.
type EventsToMetricsUpdateRuleResult struct {
	// Rules that failed to get updated.
	Failures []EventsToMetricsUpdateRuleFailure `json:"failures,omitempty"`
	// Rules that were successfully enabled or disabled.
	Successes []EventsToMetricsRule `json:"successes,omitempty"`
}

// EventsToMetricsUpdateRuleSubmission - The details that were submitted when updating an events to metrics conversion rule.
type EventsToMetricsUpdateRuleSubmission struct {
	// A submitted account id.
	AccountID int `json:"accountId"`
	// Changes the state of the rule as being enabled or disabled.
	Enabled bool `json:"enabled"`
	// A submitted rule id.
	RuleId string `json:"ruleId"`
}
