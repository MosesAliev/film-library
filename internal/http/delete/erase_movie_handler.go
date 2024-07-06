package delete

import (
	"bytes"
	"encoding/json"
	"film-library/internal/database"
	"film-library/internal/models"
	"net/http"
)

// http-обработчик для удаления информации о фильме
func EraseMovieHandler(w http.ResponseWriter, r *http.Request) {
	if w.Header().Get("role") != "admin" {
		w.WriteHeader(http.StatusForbidden)
		w.Header().Set("Content-Type", "text/Text")
		w.Write([]byte("нет доступа"))
		return
	}

	var buf bytes.Buffer
	buf.ReadFrom(r.Body)
	deletedMovie := models.Movie{}
	json.Unmarshal(buf.Bytes(), &deletedMovie)
	result := database.DB.Db.Delete(&deletedMovie) // запрос в БД для удаления информации о фильме
	if result.Error != nil {
		w.Header().Set("Content-Type", "applictaion/json")
		w.Write([]byte("Фильм не найден"))
		return
	}

	w.Write([]byte("deleted"))
}