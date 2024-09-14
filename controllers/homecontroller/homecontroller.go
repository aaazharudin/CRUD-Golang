package homecontroller

import (
	"html/template"
	"net/http"
)

func Welcome(w http.ResponseWriter, r *http.Request) {
	//Parse the tamplate file
	temp, err := template.ParseFiles("view/home/index.html")
	if err != nil {
		http.Error(w, "internal server Error ", http.StatusInternalServerError)
	}

	//eksekusi tamplate & write response writer
	err = temp.Execute(w, nil)
	if err != nil {
		http.Error(w, "Intrnal server Error", http.StatusInternalServerError)
	}
}
