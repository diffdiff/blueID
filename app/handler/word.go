package handler

import (
	"encoding/json"
	"net/http"

	"strconv"

	"github.com/diffdiff/blueID/app/model"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

// GetAllWord find translation
func GetAllWord(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	Word := []model.Word{}
	db.Preload("Lang").Find(&Word)
	respondJSON(w, http.StatusOK, Word)
}

// CreateWord create translation
func CreateWord(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	Word := model.Word{}
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

// GetWord get Word
func GetWord(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	word := vars["word"]
	Word := getWordOr404(db, word, w, r)
	if Word == nil {
		return
	}
	respondJSON(w, http.StatusOK, Word)
}

// UpdateWord update Word
func UpdateWord(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	word := vars["word"]
	Word := getWordOr404(db, word, w, r)
	if Word == nil {
		return
	}

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
	respondJSON(w, http.StatusOK, Word)
}

// DeleteWord delete Word
func DeleteWord(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	word := vars["word"]
	Word := getWordOr404(db, word, w, r)
	if Word == nil {
		return
	}
	if err := db.Delete(&Word).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusNoContent, nil)
}

// getWordOr404 gets a Word instance if exists, or respond the 404 error otherwise
func getWordOr404(db *gorm.DB, id string, w http.ResponseWriter, r *http.Request) *model.Word {
	Word := model.Word{}
	wordID, err := strconv.Atoi(id)

	if err != nil {
		respondError(w, http.StatusBadRequest, "invalid word id")
		return nil
	}

	if err := db.Preload("Lang").First(&Word, wordID).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return nil
	}

	return &Word
}
