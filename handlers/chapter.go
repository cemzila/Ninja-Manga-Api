package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/batt0s/Ninja-Manga-Api/database"
	"github.com/go-chi/chi"
)

func GetChapterByIdHandler(w http.ResponseWriter, r *http.Request) {
	sid := chi.URLParam(r, "id")
	id, err := strconv.Atoi(sid)
	if err != nil {
		sendResponse(w, http.StatusInternalServerError, map[string]string{
			"message": "Unknown error. Please look at the 'error' field.",
			"error":   err.Error(),
		})
		return
	}
	chapter := new(database.Chapter)
	err = chapter.GetById(uint(id))
	if err != nil {
		sendResponse(w, http.StatusInternalServerError, map[string]string{
			"message": "Failed to get chapter from database.",
			"error":   err.Error(),
		})
		return
	}
	sendResponse(w, http.StatusOK, map[string]interface{}{
		"message": "Success.",
		"chapter": chapter,
	})
}

func GetChaptersByMangaHandler(w http.ResponseWriter, r *http.Request) {
	mangaid, err := strconv.Atoi(chi.URLParam(r, "mangaid"))
	if err != nil {
		sendResponse(w, http.StatusInternalServerError, map[string]string{
			"message": "Unknown error. Please look at the 'error' field.",
			"error":   err.Error(),
		})
		return
	}
	chapters, err := database.GetChaptersByManga(uint(mangaid))
	if err != nil {
		sendResponse(w, http.StatusInternalServerError, map[string]string{
			"message": "Failed to get chapters by manga.",
			"error":   err.Error(),
		})
		return
	}
	if len(chapters) <= 0 {
		sendResponse(w, http.StatusOK, map[string]string{
			"message": "No chapter for that manga.",
		})
		return
	}
	sendResponse(w, http.StatusOK, map[string]interface{}{
		"message":  "Success.",
		"chapters": chapters,
	})
}

func CreateChapterHandler(w http.ResponseWriter, r *http.Request) {
	var chapter database.Chapter
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&chapter); err != nil {
		sendResponse(w, http.StatusInternalServerError, map[string]string{
			"message": "Invalid request payload.",
			"error":   err.Error(),
		})
		return
	}
	err := chapter.Save()
	if err != nil {
		sendResponse(w, http.StatusInternalServerError, map[string]string{
			"message": "Failed to save chapter to database.",
			"error":   err.Error(),
		})
		return
	}
	sendResponse(w, http.StatusOK, map[string]interface{}{
		"message": "Success.",
		"chapter": chapter,
	})
}

func DeleteChapterHandler(w http.ResponseWriter, r *http.Request) {
	chapterid, err := strconv.Atoi(chi.URLParam(r, "chapterid"))
	if err != nil {
		sendResponse(w, http.StatusInternalServerError, map[string]string{
			"message": "Unknown error. Please look at the 'error' field.",
			"error":   err.Error(),
		})
		return
	}
	chapter := new(database.Chapter)
	chapter.GetById(uint(chapterid))
	err = chapter.Delete()
	if err != nil {
		sendResponse(w, http.StatusInternalServerError, map[string]string{
			"message": "Failed to delete chapter.",
			"error":   err.Error(),
		})
		return
	}
	sendResponse(w, http.StatusInternalServerError, map[string]string{
		"message": "Successfully deleted.",
	})
}
