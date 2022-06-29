package repository

import (
	"context"
	"github.com/bhatti/GSSI/api/v1/types"
)

type CredentialsRepository interface {
	// Get - finds a credential
	Get(ctx context.Context, id string) (*types.VerifiableCredential, error)
	// Save - adds a tuple to the credential
	Save(ctx context.Context, credential *types.VerifiableCredential) error
	// Query - queries credentials
	Query(ctx context.Context, query string, offset int64, limit int32) ([]*types.VerifiableCredential, error)
	// Remove - remove a credential
	Remove(ctx context.Context, id string) (*types.VerifiableCredential, error)
}
