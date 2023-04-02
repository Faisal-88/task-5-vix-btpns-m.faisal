package photocontroller

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jeypc/go-jwt-mux/helper"
	"github.com/jeypc/go-jwt-mux/models"
)

func Index(w http.ResponseWriter, r *http.Request) {

	data := []map[string]interface{}{
		{
			"ID":       1,
			"Title":    "Gunung",
			"Caption":  "berhasil manampilkan gambar gunung",
			"PhotoUrl": "https://www.btpns.com/gunung.png",
			"UserID":   1,
		},
		{
			"ID":       2,
			"Title":    "Mobil",
			"Caption":  "berhasil manampilkan gambar mobil",
			"PhotoUrl": "https://www.btpns.com/mobil.png",
			"UserID":   1,
		},
		{
			"ID":       3,
			"Title":    "Laptop",
			"Caption":  "berhasil manampilkan gambar laptop",
			"PhotoUrl": "https://www.btpns.com/laptop.jpg",
			"UserID":   1,
		},
	}

	helper.ResponseJSON(w, http.StatusOK, data)

}

func Upload(w http.ResponseWriter, r *http.Request) {
	// parse JSON request body
	var photoUpload models.Photos
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&photoUpload); err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}
	defer r.Body.Close()

	// set user ID from the request
	// (assumes user ID is already set in photoUpload.UserId)
	userId := photoUpload.UserId

	// insert into database
	photoUpload.UserId = userId
	if err := models.DB.Create(&photoUpload).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	// send success response
	response := map[string]string{"message": "Berhasil Create Data"}
	helper.ResponseJSON(w, http.StatusOK, response)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	// ambil id dari path parameter
	vars := mux.Vars(r)
	id := vars["id"]

	// hapus data dari database
	if err := models.DB.Where("id = ?", id).Delete(&models.Photos{}).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := map[string]string{"message": "Berhasil Delete"}
	helper.ResponseJSON(w, http.StatusOK, response)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	// ambil id dari path parameter
	vars := mux.Vars(r)
	id := vars["id"]

	// daftar melalui json
	var photoUpload models.Photos
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&photoUpload); err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}
	defer r.Body.Close()

	// update data di database
	if err := models.DB.Model(&models.Photos{}).Where("id = ?", id).Updates(photoUpload).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := map[string]string{"message": "Berhasil Edit"}
	helper.ResponseJSON(w, http.StatusOK, response)
}
