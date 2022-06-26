package controllers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sekolah_ku/configs"
	"sekolah_ku/models"

	"github.com/gorilla/mux"
)

func StudentList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var student models.Student
	var arr_student []models.Student
	db := configs.Connect()
	defer db.Close()
	query := "Select id, nama_depan, nama_belakang, no_hp, gender, jenjang, hobi, alamat from siswa"
	rows, err := db.Query(query)
	if err != nil {
		log.Print(err)
		http.Error(w, "Something broke!", http.StatusInternalServerError)
	}

	for rows.Next() {
		if err := rows.Scan(&student.Id, &student.FirstName, &student.LastName, &student.PhoneNumber, &student.Gender, &student.Grade, &student.Hobbies, &student.
			Address); err != nil {
			log.Fatal(err.Error())
			http.Error(w, "Something broke!", http.StatusInternalServerError)
		} else {
			arr_student = append(arr_student, student)
		}
	}

	res := models.ResponseList{}
	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = arr_student
	fmt.Println(res.Data)
	// result, err := JsonStringify(res)
	result, err := res.Marshal()
	if err != nil {
		// throwing error dengan http status code
		http.Error(w, "Something broke!", http.StatusInternalServerError)
		return
	}
	w.Write(result)
}

func InsertStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db := configs.Connect()
	defer db.Close()

	body, errBody := ioutil.ReadAll(r.Body)
	if errBody != nil {
		// throwing error dengan http status code
		http.Error(w, "Something broke!", http.StatusInternalServerError)
		return
	}
	st, err2 := models.JsonParse(body)
	if err2 != nil {
		// throwing error dengan http status code
		http.Error(w, "Something broke!", http.StatusInternalServerError)
		return
	}
	query := "INSERT INTO siswa (nama_depan, nama_belakang, no_hp, gender, jenjang, hobi, alamat) values (?,?,?,?,?,?,?)"
	_, errDb := db.Exec(query,
		st.FirstName,
		st.LastName,
		st.PhoneNumber,
		st.Gender,
		st.Grade,
		st.Hobbies,
		st.Address,
	)
	if errDb != nil {
		// throwing error dengan http status code
		http.Error(w, "Something wrong!", http.StatusBadRequest)
		return
	}
	res := models.ResponseList{}
	res.Status = http.StatusCreated
	res.Message = "Berhasil membuat data siswa"
	response, err := res.Marshal()
	if err != nil {
		// throwing error dengan http status code
		http.Error(w, "Something broke!", http.StatusInternalServerError)
		return
	}
	w.Write(response)
}

func GetStudent(w http.ResponseWriter, r *http.Request) {
	var student models.Student
	var arr_student []models.Student
	db := configs.Connect()
	defer db.Close()
	params := mux.Vars(r)
	id := params["id"]
	query := "Select id, nama_depan, nama_belakang, no_hp, gender, jenjang, hobi, alamat from siswa where id = ?"
	rows, err := db.Query(query, id)

	if err != nil {
		log.Print(err)
		http.Error(w, "Something broke!", http.StatusInternalServerError)
	}

	for rows.Next() {
		if err := rows.Scan(&student.Id, &student.FirstName, &student.LastName, &student.PhoneNumber, &student.Gender, &student.Grade, &student.Hobbies, &student.
			Address); err != nil {
			log.Fatal(err.Error())
			http.Error(w, "Something broke!", http.StatusInternalServerError)
		} else {
			arr_student = append(arr_student, student)
		}
	}

	res := models.Response{}
	if len(arr_student) > 0 {
		res.Status = http.StatusOK
		res.Data = arr_student[0]
		res.Message = "Success"
	} else {
		res.Status = http.StatusNotFound
		res.Message = "Data tidak ditemukan"
	}
	result, err := res.Marshal()
	if err != nil {
		// throwing error dengan http status code
		http.Error(w, "Something broke!", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(result)
}

func UpdateStudent(w http.ResponseWriter, r *http.Request) {
	db := configs.Connect()
	defer db.Close()
	params := mux.Vars(r)
	id := params["id"]
	body, errBody := ioutil.ReadAll(r.Body)
	if errBody != nil {
		// throwing error dengan http status code
		log.Print(errBody)
		http.Error(w, "Something broke!", http.StatusInternalServerError)
		return
	}
	st, err2 := models.JsonParse(body)
	if err2 != nil {
		// throwing error dengan http status code
		log.Print(err2)
		http.Error(w, "Something broke!", http.StatusInternalServerError)
		return
	}
	query := "UPDATE siswa SET nama_depan=?, nama_belakang=?, no_hp=?, gender=?, jenjang=?, hobi=?, alamat=? where id = ?"
	_, errDb := db.Query(query,
		st.FirstName,
		st.LastName,
		st.PhoneNumber,
		st.Gender,
		st.Grade,
		st.Hobbies,
		st.Address,
		id,
	)

	if errDb != nil {
		log.Print(errDb)
		http.Error(w, "Something broke!", http.StatusBadRequest)
		return
	}
	res := models.ResponseList{}
	res.Status = http.StatusOK
	res.Message = "Berhasil membuat data siswa"
	response, err := res.Marshal()
	if err != nil {
		// throwing error dengan http status code

		http.Error(w, "Something broke!", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func DeleteStudent(w http.ResponseWriter, r *http.Request) {
	db := configs.Connect()
	defer db.Close()
	params := mux.Vars(r)
	id := params["id"]
	query := "DELETE from siswa where id = ?"
	_, errDb := db.Query(query, id)

	if errDb != nil {
		log.Print(errDb)
		http.Error(w, "Something broke!", http.StatusBadRequest)
		return
	}
	res := models.ResponseList{}
	res.Status = http.StatusOK
	res.Message = "Berhasil menghapus data siswa"
	response, err := res.Marshal()
	if err != nil {
		// throwing error dengan http status code
		log.Print(err)
		http.Error(w, "Something broke!", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}
