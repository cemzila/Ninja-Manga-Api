package database

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Chapter struct {
	gorm.Model
	Manga uint           `json:"manga"`
	Title string         `json:"title"`
	Pages pq.StringArray `json:"pages" gorm:"type:text[]"`
}

func (c *Chapter) Save() error {
	result := DB.Create(&c)
	return result.Error
}

func GetChaptersByManga(mangaid uint) ([]Chapter, error) {
	var chapters []Chapter
	result := DB.Where("manga = ?", mangaid).Find(&chapters)
	return chapters, result.Error
}

func (c *Chapter) GetById(id uint) error {
	result := DB.First(&c, id)
	return result.Error
}

func (c *Chapter) Delete() error {
	result := DB.Delete(&c)
	return result.Error
}
