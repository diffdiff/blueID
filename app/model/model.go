package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Lang
type Lang struct {
	gorm.Model
	LangName string `json:"lang"`
}

// Word
type Word struct {
	gorm.Model
	WordText    string `json:"word"`
	Lang        Lang   `gorm:"foreignkey:LangID"`
	LangID      uint
	Example     string `json:"example"`
	Description string `json:"description"`
}

// Translation
type Translation struct {
	gorm.Model
	Word1   Word `gorm:"foreignkey:WordID1"`
	WordID1 uint
	Word2   Word `gorm:"foreignkey:WordID2"`
	WordID2 uint
}

// DBMigrate
func DBMigrate(db *gorm.DB) *gorm.DB {
	// db.DropTableIfExists(&Lang{}, &Word{}, &Translation{})
	db.AutoMigrate(&Lang{}, &Word{}, &Translation{})
	// db.Model(&Word{}).AddForeignKey("lang_id", "langs(id)", "SET NULL", "SET NULL")
	// db.Model(&Translation{}).AddForeignKey("word_id1", "words(id)", "SET NULL", "SET NULL")
	// db.Model(&Translation{}).AddForeignKey("word_id2", "words(id)", "SET NULL", "SET NULL")
	return db
}
