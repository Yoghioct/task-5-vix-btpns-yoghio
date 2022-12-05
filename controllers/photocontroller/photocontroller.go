package photocontroller

import (
	"encoding/json"
	"net/http"

	"github.com/jeypc/go-jwt-mux/helper"
	"github.com/jeypc/go-jwt-mux/models"
)

func Index(w http.ResponseWriter, r *http.Request) {

	data := []map[string]interface{}{
		{
			"ID":       1,
			"Title":    "photo example 1",
			"Caption":  "lorem ipsum dolor sit amet",
			"PhotoUrl": "https://www.btpn.com/header/logo.jpg",
			"UserID":   1,
		},
		{
			"ID":       2,
			"Title":    "photo example 2",
			"Caption":  "lorem ipsum dolor sit amet",
			"PhotoUrl": "https://www.btpn.com/header/logo.jpg",
			"UserID":   1,
		},
		{
			"ID":       3,
			"Title":    "photo example 3",
			"Caption":  "lorem ipsum dolor sit amet",
			"PhotoUrl": "https://www.btpn.com/header/logo.jpg",
			"UserID":   1,
		},
		{
			"ID":       4,
			"Title":    "photo example 4",
			"Caption":  "lorem ipsum dolor sit amet",
			"PhotoUrl": "https://www.btpn.com/header/logo.jpg",
			"UserID":   1,
		},
		{
			"ID":       5,
			"Title":    "photo example 5",
			"Caption":  "lorem ipsum dolor sit amet",
			"PhotoUrl": "https://www.btpn.com/header/logo.jpg",
			"UserID":   1,
		},
	}

	helper.ResponseJSON(w, http.StatusOK, data)

}

func Upload(w http.ResponseWriter, r *http.Request) {
	// daftar melalui json
	var photoUpload models.Photos
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&photoUpload); err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}
	defer r.Body.Close()

	// insert ke database
	if err := models.DB.Create(&photoUpload).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}
	response := map[string]string{"message": "success"}
	helper.ResponseJSON(w, http.StatusOK, response)
}
