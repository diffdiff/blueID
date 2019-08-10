package handler

import (
	"encoding/json"
	"net/http"

	"github.com/diffdiff/blueID/app/model"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

// GetAllTranslate find translation
func GetAllTranslate(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	translate := []model.Translation{}
	db.Preload("Word1").Preload("Word2").Preload("Word1.Lang").Preload("Word2.Lang").Find(&translate)
	respondJSON(w, http.StatusOK, translate)
}

// CreateTranslate create translation
func CreateTranslate(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	Word := model.Translation{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&Word); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&Word).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondJSON(w, http.StatusCreated, Word)
}

// GetTranslate get translate
func GetTranslate(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	word := vars["word"]

	Translate := getTranslateOr404(db, word, w, r)
	if Translate == nil {
		return
	}
	respondJSON(w, http.StatusOK, Translate)
}

// UpdateTranslate update translate
func UpdateTranslate(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]
	Translate := getTranslateOr404(db, id, w, r)
	if Translate == nil {
		return
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&Translate); err != nil {	
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&Translate).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, Translate)
}

// DeleteTranslate delete translate
func DeleteTranslate(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]
	Translate := getTranslateOr404(db, id, w, r)
	if Translate == nil {
		return
	}
	if err := db.Delete(&Translate).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusNoContent, nil)
}

// getTranslateOr404 gets a Translate instance if exists, or respond the 404 error otherwise
func getTranslateOr404(db *gorm.DB, word string, w http.ResponseWriter, r *http.Request) *model.Translation {
	Word := model.Word{}
	if err := db.First(&Word, model.Word{WordText: word}).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return nil
	}

	Translation := model.Translation{}

	if err := db.Preload("Word1").Preload("Word2").Preload("Word1.Lang").Preload("Word2.Lang").Where("word_id1 = ?", Word.ID).Or("word_id2 = ?", Word.ID).Find(&Translation).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return nil
	}

	return &Translation
}
