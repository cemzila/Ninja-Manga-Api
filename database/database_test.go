package database_test

import (
	"testing"

	"github.com/batt0s/Ninja-Manga-Api/database"
)

func TestInitDb(t *testing.T) {
	err := database.InitDB("dev")
	if err != nil {
		t.Errorf("ERROR: %v", err)
	}
}

var (
	id          uint
	title       string = "Test Manga"
	cover       string = "https://preview.redd.it/n1wssjtq1u191.jpg?width=1080&crop=smart&auto=webp&s=990ab08075cbcdb69b7d7dab4e1e1cc57b5a9a74"
	artist      string = "cemzilla"
	description string = "Bu bir test mangasıdır. Okumak için değildir."
	tags        string = "shounen,falan,filan"
)

func TestSaveManga(t *testing.T) {
	manga := database.Manga{
		Title:       title,
		Cover:       cover,
		Artist:      artist,
		Description: description,
		Tags:        tags,
	}
	if err := manga.Save(); err != nil {
		t.Errorf("ERROR: %v", err)
	}
	id = manga.ID
}

func TestGetMangaById(t *testing.T) {
	manga := database.Manga{}
	manga.GetById(id)
	if manga.Title != title {
		t.Errorf("Expected %s, got %s.", title, manga.Title)
	}
}

func TestGetAllManga(t *testing.T) {
	var mangas []database.Manga
	mangas, err := database.GetAllManga()
	if err != nil {
		t.Errorf("ERROR: %v", err)
	}
	if len(mangas) != 1 {
		t.Errorf("Excepted 1, got %d.", len(mangas))
	}
	if mangas[0].Title != title {
		t.Errorf("Expected %s, got %s.", title, mangas[0].Title)
	}
}

func TestGetAllMangaByTag(t *testing.T) {
	mangas, err := database.GetAllMangaByTag("shounen")
	if err != nil {
		t.Errorf("ERROR: %v", err)
	}
	if len(mangas) != 1 {
		t.Errorf("Excepted 1, got %d.", len(mangas))
	}
	if mangas[0].Title != title {
		t.Errorf("Expected %s, got %s.", title, mangas[0].Title)
	}
}

func TestSearchManga(t *testing.T) {
	mangas, err := database.SearchManga("Test")
	if err != nil {
		t.Errorf("ERROR: %v", err)
	}
	if len(mangas) != 1 {
		t.Errorf("Excepted 1, got %d.", len(mangas))
	}
	if mangas[0].Title != title {
		t.Errorf("Expected %s, got %s.", title, mangas[0].Title)
	}
}

var (
	chpId    uint
	chpTitle string = "Test Chapter"
	pages           = []string{
		"https://preview.redd.it/v4wvyhtq1u191.jpg?width=1080&crop=smart&auto=webp&s=334a62eca13aad21deebcd308a664d45c95e313b",
		"https://preview.redd.it/n1wssjtq1u191.jpg?width=1080&crop=smart&auto=webp&s=990ab08075cbcdb69b7d7dab4e1e1cc57b5a9a74",
		"https://preview.redd.it/86n0ehtq1u191.jpg?width=1080&crop=smart&auto=webp&s=459b0cae1b8cecae89c59a00e370c14f121eec53",
	}
)

func TestSaveChapter(t *testing.T) {
	chp := database.Chapter{
		Manga: id,
		Title: chpTitle,
		Pages: pages,
	}
	if err := chp.Save(); err != nil {
		t.Errorf("ERROR: %v", err)
	}
	chpId = chp.ID
}

func TestGetChapterById(t *testing.T) {
	chp := database.Chapter{}
	err := chp.GetById(chpId)
	if err != nil {
		t.Errorf("ERROR: %v", err)
	}
	if chp.Title != chpTitle {
		t.Errorf("Excepted %s, got %s.", chpTitle, chp.Title)
	}
}

func TestGetChaptersByManga(t *testing.T) {
	chps, err := database.GetChaptersByManga(id)
	if err != nil {
		t.Errorf("ERROR: %v", err)
	}
	if len(chps) != 1 {
		t.Errorf("Excepted 1, got %d.", len(chps))
	}
	if chps[0].Title != chpTitle {
		t.Errorf("Excepted %s, got %s.", chpTitle, chps[0].Title)
	}
}

func TestDeleteChapter(t *testing.T) {
	chp := new(database.Chapter)
	chp.GetById(chpId)
	err := chp.Delete()
	if err != nil {
		t.Errorf("ERROR: %v", err)
	}
}

func TestDeleteManga(t *testing.T) {
	mng := new(database.Manga)
	mng.GetById(id)
	err := mng.Delete()
	if err != nil {
		t.Errorf("ERROR: %v", err)
	}
}
