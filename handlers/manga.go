package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/batt0s/Ninja-Manga-Api/database"
	"github.com/go-chi/chi"
)

func GetAllMangaHandler(w http.ResponseWriter, r *http.Request) {
	var mangas []database.Manga
	mangas, err := database.GetAllManga()
	if err != nil {
		sendResponse(w, http.StatusInternalServerError, map[string]string{
			"message": "Failed to get mangas from database.",
			"error":   err.Error(),
		})
		return
	}
	sendResponse(w, http.StatusOK, map[string]interface{}{
		"message": "Success.",
		"mangas":  mangas,
	})
}

func GetMangaByIdHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		sendResponse(w, http.StatusInternalServerError, map[string]string{
			"message": "Unknown error. Please look at the 'error' field.",
			"error":   err.Error(),
		})
		return
	}
	manga := new(database.Manga)
	err = manga.GetById(uint(intId))
	if err != nil {
		sendResponse(w, http.StatusInternalServerError, map[string]string{
			"message": "Failed to get manga from database.",
			"error":   err.Error(),
		})
		return
	}
	sendResponse(w, http.StatusOK, map[string]interface{}{
		"message": "Success.",
		"manga":   manga,
	})
}

func GetAllMangaByTagHandler(w http.ResponseWriter, r *http.Request) {
	tag := chi.URLParam(r, "tag")
	mangas, err := database.GetAllMangaByTag(tag)
	if err != nil {
		sendResponse(w, http.StatusInternalServerError, map[string]string{
			"message": "Failed to get mangas from database.",
			"error":   err.Error(),
		})
		return
	}
	if len(mangas) <= 0 {
		sendResponse(w, http.StatusOK, map[string]string{
			"message": "No manga with that tag.",
		})
		return
	}
	sendResponse(w, http.StatusOK, map[string]interface{}{
		"message": "Success.",
		"mangas":  mangas,
	})
}

func CreateMangaHandler(w http.ResponseWriter, r *http.Request) {
	var manga database.Manga
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&manga); err != nil {
		sendResponse(w, http.StatusInternalServerError, map[string]string{
			"message": "Invalid request payload.",
			"error":   err.Error(),
		})
		return
	}
	err := manga.Save()
	if err != nil {
		sendResponse(w, http.StatusInternalServerError, map[string]string{
			"message": "Failed to save manga to database.",
			"error":   err.Error(),
		})
		return
	}
	sendResponse(w, http.StatusOK, map[string]interface{}{
		"message": "Success.",
		"manga":   manga,
	})
}

func SearchMangaHandler(w http.ResponseWriter, r *http.Request) {
	keyword := chi.URLParam(r, "keyword")
	mangas, err := database.SearchManga(keyword)
	if err != nil {
		sendResponse(w, http.StatusInternalServerError, map[string]string{
			"message": "Failed to get mangas.",
			"error":   err.Error(),
		})
		return
	}
	if len(mangas) <= 0 {
		sendResponse(w, http.StatusOK, map[string]string{
			"message": "No manga found.",
		})
		return
	}
	sendResponse(w, http.StatusOK, map[string]interface{}{
		"message": "Success",
		"mangas":  mangas,
	})
}

func DeleteMangaHandler(w http.ResponseWriter, r *http.Request) {
	mangaid, err := strconv.Atoi(chi.URLParam(r, "mangaid"))
	if err != nil {
		sendResponse(w, http.StatusInternalServerError, map[string]string{
			"message": "Unknown error. Please look at the 'error' field.",
			"error":   err.Error(),
		})
		return
	}
	manga := new(database.Manga)
	manga.GetById(uint(mangaid))
	err = manga.Delete()
	if err != nil {
		sendResponse(w, http.StatusInternalServerError, map[string]string{
			"message": "Failed to delete manga.",
			"error":   err.Error(),
		})
		return
	}
	sendResponse(w, http.StatusInternalServerError, map[string]string{
		"message": "Successfully deleted.",
	})
}
