package handler

import (
	"encoding/json"
	"net/http"

	"strconv"

	"github.com/diffdiff/blueID/app/model"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

//CreateLang create language
func CreateLang(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	lang := model.Lang{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&lang); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}

	defer r.Body.Close()

	if err := db.Save(&lang).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondJSON(w, http.StatusCreated, lang)
}

// GetAllLangs
func GetAllLangs(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	lang := []model.Lang{}
	db.Find(&lang)
	respondJSON(w, http.StatusOK, lang)
}

// GetLang get Lang
func GetLang(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]
	Lang := getLangOr404(db, id, w, r)
	if Lang == nil {
		return
	}
	respondJSON(w, http.StatusOK, Lang)
}

// UpdateLang update Lang
func UpdateLang(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]
	Lang := getLangOr404(db, id, w, r)
	if Lang == nil {
		return
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&Lang); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&Lang).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, Lang)
}

// DeleteLang delete Lang
func DeleteLang(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]
	Lang := getLangOr404(db, id, w, r)
	if Lang == nil {
		return
	}
	if err := db.Delete(&Lang).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusNoContent, nil)
}

// getLangOr404 gets a Lang instance if exists, or respond the 404 error otherwise
func getLangOr404(db *gorm.DB, id string, w http.ResponseWriter, r *http.Request) *model.Lang {
	Lang := model.Lang{}

	langID, err := strconv.Atoi(id)

	if err != nil {
		respondError(w, http.StatusBadRequest, "invalid lang id")
		return nil
	}

	if err := db.First(&Lang, langID).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return nil
	}

	return &Lang
}
