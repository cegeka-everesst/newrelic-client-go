package workloads

import (
	"context"

	"github.com/cegeka-everesst/newrelic-client-go/internal/serialization"
	"github.com/cegeka-everesst/newrelic-client-go/pkg/errors"
	"github.com/cegeka-everesst/newrelic-client-go/pkg/nerdgraph"
	"github.com/cegeka-everesst/newrelic-client-go/pkg/users"
)

// Workload represents a New Relic One workload.
//
// Deprecated: use WorkloadCollection instead
type Workload struct {
	Account             nerdgraph.AccountReference `json:"account,omitempty"`
	CreatedAt           serialization.EpochTime    `json:"createdAt,omitempty"`
	CreatedBy           users.UserReference        `json:"createdBy,omitempty"`
	Entities            []EntityRef                `json:"entities,omitempty"`
	EntitySearchQueries []EntitySearchQuery        `json:"entitySearchQueries,omitempty"`
	EntitySearchQuery   string                     `json:"entitySearchQuery,omitempty"`
	GUID                string                     `json:"guid,omitempty"`
	ID                  int                        `json:"id,omitempty"`
	Name                string                     `json:"name,omitempty"`
	Permalink           string                     `json:"permalink,omitempty"`
	ScopeAccounts       ScopeAccounts              `json:"scopeAccounts,omitempty"`
	UpdatedAt           *serialization.EpochTime   `json:"updatedAt,omitempty"`
}

// EntityRef represents an entity referenced by this workload.
//
// Deprecated: use WorkloadEntityRef instead
type EntityRef struct {
	GUID string `json:"guid,omitempty"`
}

// EntitySearchQuery represents an entity search used by this workload.
//
// Deprecated: use WorkloadEntitySearchQuery instead
type EntitySearchQuery struct {
	CreatedAt serialization.EpochTime  `json:"createdAt,omitempty"`
	CreatedBy users.UserReference      `json:"createdBy,omitempty"`
	ID        int                      `json:"id,omitempty"`
	Query     string                   `json:"query,omitempty"`
	UpdatedAt *serialization.EpochTime `json:"updatedAt,omitempty"`
}

// ScopeAccounts represents the accounts used to scope this workload.
//
// Deprecated: use WorkloadScopeAccounts instead
type ScopeAccounts struct {
	AccountIDs []int `json:"accountIds,omitempty"`
}

// CreateInput represents the input parameters used for creating or updating a workload.
//
// Deprecated: use WorkloadCreateInput instead
type CreateInput struct {
	EntityGUIDs         []string                 `json:"entityGuids,omitempty"`
	EntitySearchQueries []EntitySearchQueryInput `json:"entitySearchQueries,omitempty"`
	Name                string                   `json:"name,omitempty"`
	ScopeAccountsInput  *ScopeAccountsInput      `json:"scopeAccounts,omitempty"`
}

// EntitySearchQueryInput represents an entity search query for creating or updating a workload.
//
// Deprecated: use WorkloadEntitySearchQueryInput instead
type EntitySearchQueryInput struct {
	Query string `json:"query,omitempty"`
}

// UpdateCollectionEntitySearchQueryInput represents an entity search query for creating or updating a workload.
//
// Deprecated: use WorkloadUpdateCollectionEntitySearchQueryInput instead
type UpdateCollectionEntitySearchQueryInput struct {
	ID int `json:"id,omitempty"`
	EntitySearchQueryInput
}

// ScopeAccountsInput is the input object containing accounts that will be used to get entities from.
//
// Deprecated: use WorkloadScopeAccountsInput instead
type ScopeAccountsInput struct {
	AccountIDs []int `json:"accountIds,omitempty"`
}

// DuplicateInput represents the input object used to identify the workload to be duplicated.
//
// Deprecated: use WorkloadDuplicateInput instead
type DuplicateInput struct {
	Name string `json:"name,omitempty"`
}

// UpdateInput represents the input object used to identify the workload to be updated and its new changes.
//
// Deprecated: use WorkloadUpdateInput instead
type UpdateInput struct {
	EntityGUIDs         []string                 `json:"entityGuids,omitempty"`
	EntitySearchQueries []EntitySearchQueryInput `json:"entitySearchQueries,omitempty"`
	Name                string                   `json:"name,omitempty"`
	ScopeAccountsInput  *ScopeAccountsInput      `json:"scopeAccounts,omitempty"`
}

// ListWorkloads retrieves a set of New Relic One workloads by their account ID.
//
// Deprecated: use entities.GetEntities instead
func (e *Workloads) ListWorkloads(accountID int) ([]*Workload, error) {
	return e.ListWorkloadsWithContext(context.Background(), accountID)
}

// ListWorkloadsWithContext retrieves a set of New Relic One workloads by their account ID.
//
// Deprecated: use entities.GetEntitiesWithContext instead
func (e *Workloads) ListWorkloadsWithContext(ctx context.Context, accountID int) ([]*Workload, error) {
	resp := workloadsResponse{}
	vars := map[string]interface{}{
		"accountId": accountID,
	}

	if err := e.client.NerdGraphQueryWithContext(ctx, listWorkloadsQuery, vars, &resp); err != nil {
		return nil, err
	}

	if len(resp.Actor.Account.Workload.Collections) == 0 {
		return nil, errors.NewNotFound("")
	}

	return resp.Actor.Account.Workload.Collections, nil
}

// GetWorkload retrieves a New Relic One workload by its GUID.
//
// Deprecated: use entities.GetEntity instead
func (e *Workloads) GetWorkload(accountID int, workloadGUID string) (*Workload, error) {
	return e.GetWorkloadWithContext(context.Background(), accountID, workloadGUID)
}

// GetWorkloadWithContext retrieves a New Relic One workload by its GUID.
//
// Deprecated: use entities.GetEntityWithContext instead
func (e *Workloads) GetWorkloadWithContext(ctx context.Context, accountID int, workloadGUID string) (*Workload, error) {
	resp := workloadResponse{}
	vars := map[string]interface{}{
		"accountId": accountID,
		"guid":      workloadGUID,
	}

	if err := e.client.NerdGraphQueryWithContext(ctx, getWorkloadQuery, vars, &resp); err != nil {
		return nil, err
	}

	return &resp.Actor.Account.Workload.Collection, nil
}

// CreateWorkload creates a New Relic One workload.
//
// Deprecated: use WorkloadCreate instead
func (e *Workloads) CreateWorkload(accountID int, workload CreateInput) (*Workload, error) {
	return e.CreateWorkloadWithContext(context.Background(), accountID, workload)
}

// CreateWorkloadWithContext creates a New Relic One workload.
//
// Deprecated: use WorkloadCreateWithContext instead
func (e *Workloads) CreateWorkloadWithContext(ctx context.Context, accountID int, workload CreateInput) (*Workload, error) {
	resp := workloadCreateResponse{}
	vars := map[string]interface{}{
		"accountId": accountID,
		"workload":  workload,
	}

	if err := e.client.NerdGraphQueryWithContext(ctx, createWorkloadMutation, vars, &resp); err != nil {
		return nil, err
	}

	return &resp.WorkloadCreate, nil
}

// DeleteWorkload deletes a New Relic One workload.
//
// Deprecated: use WorkloadDelete instead
func (e *Workloads) DeleteWorkload(guid string) (*Workload, error) {
	return e.DeleteWorkloadWithContext(context.Background(), guid)
}

// DeleteWorkloadWithContext deletes a New Relic One workload.
//
// Deprecated: use WorkloadDeleteWithContext instead
func (e *Workloads) DeleteWorkloadWithContext(ctx context.Context, guid string) (*Workload, error) {
	resp := workloadDeleteResponse{}
	vars := map[string]interface{}{
		"guid": guid,
	}

	if err := e.client.NerdGraphQueryWithContext(ctx, deleteWorkloadMutation, vars, &resp); err != nil {
		return nil, err
	}

	return &resp.WorkloadDelete, nil
}

// DuplicateWorkload duplicates a New Relic One workload.
//
// Deprecated: use WorkloadDuplicate instead
func (e *Workloads) DuplicateWorkload(accountID int, sourceGUID string, workload *DuplicateInput) (*Workload, error) {
	return e.DuplicateWorkloadWithContext(context.Background(), accountID, sourceGUID, workload)
}

// DuplicateWorkloadWithContext duplicates a New Relic One workload.
//
// Deprecated: use WorkloadDuplicateWithContext instead
func (e *Workloads) DuplicateWorkloadWithContext(ctx context.Context, accountID int, sourceGUID string, workload *DuplicateInput) (*Workload, error) {
	resp := workloadDuplicateResponse{}
	vars := map[string]interface{}{
		"accountId":  accountID,
		"sourceGuid": sourceGUID,
		"workload":   workload,
	}

	if err := e.client.NerdGraphQueryWithContext(ctx, duplicateWorkloadMutation, vars, &resp); err != nil {
		return nil, err
	}

	return &resp.WorkloadDuplicate, nil
}

// UpdateWorkload updates a New Relic One workload.
//
// Deprecated: use WorkloadUpdate instead
func (e *Workloads) UpdateWorkload(guid string, workload UpdateInput) (*Workload, error) {
	return e.UpdateWorkloadWithContext(context.Background(), guid, workload)
}

// UpdateWorkloadWithContext updates a New Relic One workload.
//
// Deprecated: use WorkloadUpdateWithContext instead
func (e *Workloads) UpdateWorkloadWithContext(ctx context.Context, guid string, workload UpdateInput) (*Workload, error) {
	resp := workloadUpdateResponse{}
	vars := map[string]interface{}{
		"guid":     guid,
		"workload": workload,
	}

	if err := e.client.NerdGraphQueryWithContext(ctx, updateWorkloadMutation, vars, &resp); err != nil {
		return nil, err
	}

	return &resp.WorkloadUpdate, nil
}

const (
	// graphqlWorkloadStructFields is the set of fields that we want returned on workload queries,
	// and should map back directly to the Workload struct
	graphqlWorkloadStructFields = `
			account {
				id
				name
			}
			createdBy {
				email
				gravatar
				id
				name
			}
			createdAt
			entities {
				guid
			}
			entitySearchQueries {
				createdAt
				createdBy {
					email
					gravatar
					id
					name
				}
				id
				query
				updatedAt
			}
			entitySearchQuery
			guid
			id
			permalink
			name
			scopeAccounts {
				accountIds
			}
			updatedAt
`

	getWorkloadQuery = `query($guid: EntityGuid!, $accountId: Int!) { actor { account(id: $accountId) { workload { collection(guid: $guid) {` +
		graphqlWorkloadStructFields +
		` } } } } }`

	listWorkloadsQuery = `query($accountId: Int!) { actor { account(id: $accountId) { workload { collections {` +
		graphqlWorkloadStructFields +
		` } } } } }`

	createWorkloadMutation = `
		mutation($accountId: Int!, $workload: WorkloadCreateInput!) {
			workloadCreate(accountId: $accountId, workload: $workload) {` +
		graphqlWorkloadStructFields +
		` } }`

	deleteWorkloadMutation = `
		mutation($guid: EntityGuid!) {
			workloadDelete(guid: $guid) {` +
		graphqlWorkloadStructFields +
		` } }`

	duplicateWorkloadMutation = `
		mutation($accountId: Int!, $sourceGuid: EntityGuid!, $workload: WorkloadDuplicateInput) {
			workloadDuplicate(accountId: $accountId, sourceGuid: $sourceGuid, workload: $workload) {` +
		graphqlWorkloadStructFields +
		` } }`

	updateWorkloadMutation = `
		mutation($guid: EntityGuid!, $workload: WorkloadUpdateInput!) {
			workloadUpdate(guid: $guid, workload: $workload) {` +
		graphqlWorkloadStructFields +
		` } }`
)

type workloadsResponse struct {
	Actor struct {
		Account struct {
			Workload struct {
				Collections []*Workload
			}
		}
	}
}

type workloadResponse struct {
	Actor struct {
		Account struct {
			Workload struct {
				Collection Workload
			}
		}
	}
}

type workloadCreateResponse struct {
	WorkloadCreate Workload
}

type workloadDeleteResponse struct {
	WorkloadDelete Workload
}

type workloadDuplicateResponse struct {
	WorkloadDuplicate Workload
}

type workloadUpdateResponse struct {
	WorkloadUpdate Workload
}
