package read_service

import (
	"context"
	"pastebin/sqlc_db"
)

type ReadService struct {
	querier *sqlc_db.Queries
}

// return an object that is pointer to ReadService.
func New(querier *sqlc_db.Queries) *ReadService {
	return &ReadService{querier}
}

func (rs *ReadService) PerformRead(ctx context.Context, shortlink string) sqlc_db.GetPasteRow {
	pasteDetails, err := rs.querier.GetPaste(ctx, shortlink)

	if err != nil {
		panic(err)
	}

	return pasteDetails
}
