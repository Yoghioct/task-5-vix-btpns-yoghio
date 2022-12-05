package photocontroller

import (
	"net/http"

	"github.com/jeypc/go-jwt-mux/helper"
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
