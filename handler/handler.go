package handler

import ( "net/http"
	"github.com/Medina7276/make-your-game/model"
)


// GET
func getScoreBoard(w http.ResponseWriter, r *http.Request, model.Scoreboard) {
	if r.Method == http.MethodGet {

	}
}

// POST
func postScoreBoard(w http.ResponseWriter, r *http.Request, model.scoreboard) {
	if r.Method == http.MethodPost {
		
	}
}