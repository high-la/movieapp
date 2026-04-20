package metadata

import (
	"context"
	"errors"

	"github.com/high-la/movieapp/metadata/internal/repository"
	"github.com/high-la/movieapp/metadata/pkg/model"
)

// ErrNotFound is returned when a requisted record is not found.
var ErrNotFound = errors.New("note found")

// metadataRepository defines the behavior required by the controller.
// Any type that implements Get(ctx, id) will satisfy this interface implicitly (no explicit declaration needed).
type metadataRepository interface {
	Get(ctx context.Context, id string) (*model.Metadata, error)
	Put(ctx context.Context, id string, metadata *model.Metadata) error
}

// Controller defines a metadata service controller.
type Controller struct {
	repo metadataRepository
}

// New creates a metadata service controller.
func New(repo metadataRepository) *Controller {

	return &Controller{repo}
}

// Get returns movie metadata by id.
func (c *Controller) Get(ctx context.Context, id string) (*model.Metadata, error) {

	res, err := c.repo.Get(ctx, id)
	if err != nil && errors.Is(err, repository.ErrNotFound) {
		return nil, ErrNotFound
	}

	return res, err
}
