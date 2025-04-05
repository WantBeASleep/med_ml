package download

import (
	"context"
	"io"
	"path/filepath"

	"github.com/google/uuid"
)

func (s *service) GetImage(ctx context.Context, uziID uuid.UUID, imageID uuid.UUID) (io.ReadCloser, error) {
	return s.repo.NewFileRepo().GetFile(
		ctx,
		filepath.Join(
			uziID.String(),
			imageID.String(),
			imageID.String(),
		),
	)
}
