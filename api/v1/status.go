package api // import "github.com/nylar/triage/api/v1"
import (
	"encoding/json"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/nylar/triage/triage"
)

type StatusService struct {
	db *sqlx.DB
}

func (ss *StatusService) List() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		statuses := &triage.Statuses{}
		if err := statuses.FindAll(ss.db); err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(statuses)
	})
}
