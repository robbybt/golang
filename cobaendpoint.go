package main

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/mholt/binding"
	"log"
	"net/http"
	"time"
)

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/todos", TodoIndex)
	router.HandleFunc("/todos/{todoId}&{nama}", TodoShow)

	log.Fatal(http.ListenAndServe(":8090", router))
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func TodoIndex(w http.ResponseWriter, r *http.Request) {
	a := r.FormValue("id")
	fmt.Fprintln(w, a)
	request := requestjson{}
	err := binding.Bind(r, &request)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Fprintf(w, "%+v", request)
}

type requestjson struct {
	id   int
	nama string
}

func (f *requestjson) FieldMap(r *http.Request) binding.FieldMap {
	return binding.FieldMap{
		&f.id:   "id",
		&f.nama: "nama",
	}
}

func TodoShow(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	ctx = context.WithValue(ctx, "start_time", time.Now())
	ctx = context.WithValue(ctx, "cacat", "asd")
	ctx = context.WithValue(ctx, "tole", "asd2")
	fmt.Fprintf(w, "%+v\n", ctx)
	fmt.Fprintln(w, ctx.Value("tole"))
}
