package memory

import (
	"context"
	"sync"

	"github.com/high-la/movieapp/metadata/internal/repository"
	model "github.com/high-la/movieapp/metadata/pkg/model"
)

// Repository is an in-memory implementation of the metadataRepository interface.
// It satisfies the interface implicitly by implementing the Get method.
type Repository struct {
	sync.RWMutex
	data map[string]*model.Metadata
}

// New creates a new in-memory repository.
// The returned *Repository can be passed wherever metadataRepository is required.
func New() *Repository {
	return &Repository{data: map[string]*model.Metadata{}}
}

// Get retrieves movie metadata by movie id.
// This method makes *Repository satisfy the metadataRepository interface.
func (r *Repository) Get(_ context.Context, id string) (*model.Metadata, error) {

	r.RLock()
	defer r.RUnlock()

	m, ok := r.data[id]
	if !ok {
		return nil, repository.ErrNotFound
	}

	return m, nil
}

// Put adds movie metadata for a given movie id.
func (r *Repository) Put(_ context.Context, id string, metadata *model.Metadata) error {

	r.Lock()
	defer r.Unlock()

	r.data[id] = metadata

	return nil
}
