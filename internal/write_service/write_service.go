package write_service

import (
	"context"
	"pastebin/sqlc_db"
)

type WriteService struct {
	querier *sqlc_db.Queries
}

// return an object that is pointer to WriteService.
func New(querier *sqlc_db.Queries) *WriteService {
	return &WriteService{querier}
}

func (ws *WriteService) PerformWrite(ctx context.Context, args sqlc_db.CreatePasteParams) string {
	// store paste in a filesystem.
	// get the pastepath.
	// map the pastepath with hased shortlink.
	// insert in db.
	ws.querier.CreatePaste(ctx, args)

	return args.Shortlink
}

func (ws *WriteService) CreateShortlink(ctx context.Context, pasteContent string) string {
	// obtain pastePath

	// map this to a short link

	return "shortlink"
}
