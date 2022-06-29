package repository

import (
	"context"
	"github.com/bhatti/GSSI/api/v1/types"
	"github.com/bhatti/GSSI/internal/config"
	"sync"
)

// credentialsRepositoryMemoryRepository - provides in-memory persistence for tuple space
type credentialsRepositoryMemoryRepository struct {
	lock    sync.RWMutex
	storage map[string]*types.VerifiableCredential
}

// NewCredentialsRepositoryMemory creates repository
func NewCredentialsRepositoryMemory(config *config.Config) (CredentialsRepository, error) {
	return &credentialsRepositoryMemoryRepository{
			storage: make(map[string]*types.VerifiableCredential),
		},
		nil
}

func (r *credentialsRepositoryMemoryRepository) Get(
	_ context.Context,
	id string) (tuple *types.VerifiableCredential, err error) {
	r.lock.RLock()
	defer r.lock.RUnlock()
	return r.storage[id], nil
}

// Save - adds a tuple
func (r *credentialsRepositoryMemoryRepository) Save(
	_ context.Context,
	credential *types.VerifiableCredential) (err error) {
	r.lock.Lock()
	defer r.lock.Unlock()
	r.storage[credential.Id] = credential
	return nil
}

// Query - queries tuple space
func (r *credentialsRepositoryMemoryRepository) Query(
	_ context.Context,
	_ string,
	_ int64,
	_ int32) (res []*types.VerifiableCredential, err error) {
	r.lock.RLock()
	defer r.lock.RUnlock()
	for _, next := range r.storage {
		res = append(res, next)
	}
	return
}

func (r *credentialsRepositoryMemoryRepository) Remove(
	ctx context.Context,
	id string) (credential *types.VerifiableCredential, err error) {
	r.lock.Lock()
	defer r.lock.Unlock()
	credential = r.storage[id]
	if credential != nil {
		delete(r.storage, id)
	}
	return
}
