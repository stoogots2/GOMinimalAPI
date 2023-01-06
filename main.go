package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func main() {

	var persons []Person = GetPeople()

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/api/person", func(w http.ResponseWriter, r *http.Request) {
		b, _ := json.Marshal(persons)

		w.Write(b)

	})
	r.Get("/api/person/{index}", func(w http.ResponseWriter, r *http.Request) {
		indexStr := chi.URLParam(r, "index")
		index, err := strconv.Atoi(indexStr)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if index > len(persons) {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		b, _ := json.Marshal(persons[index-1])

		w.Write(b)

	})
	r.Post("/api/person", func(w http.ResponseWriter, r *http.Request) {

		var person Person
		err := json.NewDecoder(r.Body).Decode(&person)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		persons = append(persons, person)
	})

	http.ListenAndServe(":8080", r)

	//_ = r
}

type Person struct {
	Id               int    `json:"Id"`
	WustleduId       string `json:"WWWeduid,omitempty"`
	EmployeeId       string `json:"EmployeeId,omitempty"`
	FirstName        string `json:"FirstName,omitempty"`
	LastName         string `json:"LastName,omitempty"`
	BusinessTitle    string `json:"BusinessTitle,omitempty"`
	EmploymentStatus string `json:"EmploymentStatus,omitempty"`
	WorkerType       string `json:"WorkerType,omitempty"`
	HiddenField      string `json:"-"`
}

func GetPeople() []Person {
	var persons = []Person{
		Person{
			Id:               1,
			FirstName:        "Rick",
			LastName:         "Rondis",
			WWWeduId:       strings.Repeat("0", 12),
			EmployeeId:       strings.Repeat("0", 6),
			BusinessTitle:    "",
			EmploymentStatus: "",
			WorkerType:       "",
		},
		Person{
			Id:               2,
			FirstName:        "Bott",
			LastName:         "Qughes",
			WWWeduId:       strings.Repeat("1", 12),
			EmployeeId:       strings.Repeat("1", 6),
			BusinessTitle:    "",
			EmploymentStatus: "",
			WorkerType:       "",
		},
		Person{
			Id:               3,
			FirstName:        "Natt",
			LastName:         "Jouse",
			WWWeduId:       strings.Repeat("2", 12),
			EmployeeId:       strings.Repeat("2", 6),
			BusinessTitle:    "",
			EmploymentStatus: "",
			WorkerType:       "",
		},
		Person{
			Id:               4,
			FirstName:        "Telly",
			LastName:         "Snavick",
			WWWeduId:       strings.Repeat("3", 12),
			EmployeeId:       strings.Repeat("3", 6),
			BusinessTitle:    "",
			EmploymentStatus: "",
			WorkerType:       "",
		},
		Person{
			Id:               5,
			FirstName:        "Sarah",
			LastName:         "Bith",
			WWWeduId:       strings.Repeat("4", 12),
			EmployeeId:       strings.Repeat("4", 6),
			BusinessTitle:    "",
			EmploymentStatus: "",
			WorkerType:       "",
		},
		Person{
			Id:               6,
			FirstName:        "Tarah",
			LastName:         "Pveress",
			WWWeduId:       strings.Repeat("5", 12),
			EmployeeId:       strings.Repeat("5", 6),
			BusinessTitle:    "",
			EmploymentStatus: "",
			WorkerType:       "",
		},
		Person{
			Id:               7,
			FirstName:        "Bollin",
			LastName:         "Bolchaus",
			WWWeduId:       strings.Repeat("6", 12),
			EmployeeId:       strings.Repeat("6", 6),
			BusinessTitle:    "",
			EmploymentStatus: "",
			WorkerType:       "",
		},
		Person{
			Id:               8,
			FirstName:        "Mennis",
			LastName:         "Malmay",
			WWWeduId:       strings.Repeat("7", 12),
			EmployeeId:       strings.Repeat("7", 6),
			BusinessTitle:    "",
			EmploymentStatus: "",
			WorkerType:       "",
		},
		Person{
			Id:               9,
			FirstName:        "",
			LastName:         "",
			WWWeduId:       strings.Repeat("8", 12),
			EmployeeId:       strings.Repeat("8", 6),
			BusinessTitle:    "",
			EmploymentStatus: "",
			WorkerType:       "",
		},
		Person{
			Id:               10,
			FirstName:        "",
			LastName:         "",
			WWWeduId:       strings.Repeat("9", 12),
			EmployeeId:       strings.Repeat("9", 6),
			BusinessTitle:    "",
			EmploymentStatus: "",
			WorkerType:       "",
		},
	}
	return persons

}
