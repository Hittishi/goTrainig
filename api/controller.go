package api

import (
	"database/sql"
	"encoding/json"
	"goTraining/models"
	"net/http"
)

func OrderHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "method not supported", http.StatusMethodNotAllowed)
		}

		var o models.Order

		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&o); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if err := ProcessOrderLogic(db, o); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(o)

	}

}

// func CreateHandler(db *sql.DB) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		if r.Method != http.MethodPost {
// 			http.Error(w, "method not supported", http.StatusMethodNotAllowed)
// 		}
// 		var book models.Book
// 		decoder := json.NewDecoder(r.Body)
// 		if err := decoder.Decode(&book); err != nil {
// 			http.Error(w, err.Error(), http.StatusBadRequest)
// 			return
// 		}
// 		if err := CreateBookLogic(db, book); err != nil {
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 			return
// 		}
// 		w.WriteHeader(http.StatusCreated)
// 		json.NewEncoder(w).Encode(book)

// 	}
// }
