// Code generated by tutone: DO NOT EDIT
package ai

// AiWorkflowsConfiguration - Workflows configuration interface
type AiWorkflowsConfiguration struct {
	// NRQL enrichment query
	Query string `json:"query,omitempty"`
}

func (x *AiWorkflowsConfiguration) ImplementsAiWorkflowsConfiguration() {}

// AiWorkflowsConfigurationInterface - Enrichment configuration object
type AiWorkflowsConfigurationInterface interface {
	ImplementsAiWorkflowsConfiguration()
}

// AiNotificationsError - Notifications error interface
type AiNotificationsError struct {
	// Error details
	Details string `json:"details,omitempty"`
	// Error description
	Description string `json:"description,omitempty"`
	// SuggestionError type
	Type AiNotificationsErrorType `json:"type,omitempty"`
	// List of invalid fields
	Fields []AiNotificationsFieldError `json:"fields,omitempty"`
	// Names of other constraints this constraint is dependent on
	Dependencies []string `json:"dependencies,omitempty"`
	// Name of the missing constraint
	Name string `json:"name,omitempty"`
}

func (x *AiNotificationsError) ImplementsAiNotificationsError() {}

// AiNotificationsErrorInterface - Notifications error interface
type AiNotificationsErrorInterface interface {
	ImplementsAiNotificationsError()
}
