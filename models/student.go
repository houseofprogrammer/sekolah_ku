package models

import (
	"encoding/json"
)

type Student struct {
	Id          int    `db:"id" json:"id"`
	FirstName   string `db:"nama_depan" json:"nama_depan"`
	LastName    string `db:"nama_belakang" json:"nama_belakang"`
	PhoneNumber string `db:"no_hp" json:"no_hp"`
	Gender      string `db:"gender" json:"gender"`
	Grade       string `db:"jenjang" json:"jenjang"`
	Hobbies     string `db:"hobi" json:"hobi"`
	Address     string `db:"alamat" json:"alamat"`
}

type ResponseList struct {
	Status  int32     `json:"status"`
	Message string    `json:"message"`
	Data    []Student `json:"data"`
}

type Response struct {
	Status  int32   `json:"status"`
	Message string  `json:"message"`
	Data    Student `json:"data"`
}

func (r *ResponseList) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func (r *Response) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func JsonParse(jsonData []byte) (Student, error) {
	var StudentData Student
	err := json.Unmarshal(jsonData, &StudentData)
	return StudentData, err
}

func JsonStringify(val any) ([]byte, error) {
	return json.Marshal(val)
}
