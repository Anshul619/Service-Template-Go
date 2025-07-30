package main

import (
	"context"
	"log"
	"net/http"
	"template-service/internal/handler"
	"template-service/internal/model"
	"template-service/internal/service"
)

// Fake repo for demo
type fakeRepo struct{}

func (f *fakeRepo) FindByID(_ context.Context, id string) (*model.User, error) {
	return &model.User{ID: id, Name: "John Doe", Email: "john@example.com"}, nil
}

func main() {
	repo := &fakeRepo{}
	svc := service.NewUserService(repo)
	handler := handler.NewUserHandler(svc)

	// If we want to go mux ("github.com/gorilla/mux"), we can initilize HandleFunc like following
	// 	r := mux.NewRouter()
	// 	r.HandleFunc("/", myHandler)
	http.HandleFunc("/user", handler.GetUser)

	log.Println("Listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
