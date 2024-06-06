package systemerror

import (
	"github.com/jmoiron/sqlx"
)

/*
* if nats is used then we would log message in the file in each server. then we would use a logger microservice to log to database or open serch
 */
type PostgressDBLogWriter struct {
	*sqlx.DB
}

func (yo PostgressDBLogWriter) Write(p []byte) (n int, err error) {
	n = 0
	// id, err := uuid.NewV7()
	// if err != nil {
	// 	return
	// }

	// query := `
	// INSERT INTO system_error_logs
	// 	(
	// 		sel_id,
	// 	 	sel_log
	// 	)
	// 	VALUES ($1,$2,)
	// `

	// _, err = yo.DB.Exec(query, id.String(), string(p))

	// n = len(p)
	// return

	query := `
	INSERT INTO system_error_logs 
		(
		 	sel_log 
		) 
		VALUES ($1) 
	`

	_, err = yo.DB.Exec(query, string(p))

	if err == nil {
		n = len(p)
	}
	return

}

/*
type DBHandler struct {
	// *commonHandler
	*slog.JSONHandler
}

// Enabled reports whether the handler handles records at the given level.
// The handler ignores records whose level is lower.
func (h *DBHandler) Enabled(c context.Context, level slog.Level) bool {
	return h.JSONHandler.Enabled(c, level)
}

// WithAttrs returns a new [JSONHandler] whose attributes consists
// of h's attributes followed by attrs.
func (h *DBHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	jsonHandler := h.JSONHandler.WithAttrs(attrs).(*slog.JSONHandler)
	return &DBHandler{JSONHandler: jsonHandler}
}

func (h *DBHandler) WithGroup(name string) slog.Handler {
	return &DBHandler{JSONHandler: h.JSONHandler.WithGroup(name).(*slog.JSONHandler)}
}

func (h *DBHandler) Handle(_ context.Context, r slog.Record) error {
	// return h.commonHandler.handle(r)
	return nil
}

*/

// type PostgressDBLogTableDetail struct {
// 	UuidPrimaryKey bool
// 	PrimaryKey     string
// 	tableName      string
// }
