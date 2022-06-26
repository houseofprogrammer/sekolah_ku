package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sekolah_ku/controllers"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/students", controllers.StudentList).Methods("GET")
	r.HandleFunc("/students/{id}", controllers.GetStudent).Methods("GET")
	r.HandleFunc("/students/{id}", controllers.UpdateStudent).Methods("PUT")
	r.HandleFunc("/students/{id}", controllers.DeleteStudent).Methods("DELETE")
	r.HandleFunc("/students/create", controllers.InsertStudent).Methods("POST")
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		strRes := map[string]string{
			"message": "Succes",
		}
		res, err := json.Marshal(strRes)
		if err != nil {
			// throwing error dengan http status code
			http.Error(w, "Something broke!", http.StatusInternalServerError)
		}
		w.Write(res)
	})
	http.Handle("/", r)
	port := ":8080"
	fmt.Println("Server running on Port" + port)
	httpErr := http.ListenAndServe(port, nil)
	if httpErr != nil {
		panic(httpErr)
	}
}
