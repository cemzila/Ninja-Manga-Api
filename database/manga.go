package database

import (
	"errors"

	"gorm.io/gorm"
)

type Manga struct {
	gorm.Model
	Title       string `json:"title"`
	Cover       string `json:"cover"`
	Artist      string `json:"artist"`
	Description string `json:"description"`
	Tags        string `json:"tags"`
}

func (m *Manga) Save() error {
	mang := Manga{}
	if result := DB.Where("title = ?", m.Title).First(&mang); result.Error == nil {
		return errors.New("manga already in database")
	}
	if result := DB.Create(&m); result.Error != nil {
		return result.Error
	}
	return nil
}

func (m *Manga) GetById(id uint) error {
	result := DB.First(&m, id)
	return result.Error
}

func GetAllManga() ([]Manga, error) {
	var mangas []Manga
	result := DB.Find(&mangas)
	return mangas, result.Error
}

func GetAllMangaByTag(tag string) ([]Manga, error) {
	var mangas []Manga
	result := DB.Where("tags LIKE ?", "%"+tag+"%").Find(&mangas)
	return mangas, result.Error
}

func SearchManga(keyword string) ([]Manga, error) {
	var mangas []Manga
	result := DB.Where("title LIKE ?", "%"+keyword+"%").Find(&mangas)
	return mangas, result.Error
}

func (m *Manga) Delete() error {
	result := DB.Delete(&m)
	return result.Error
}
