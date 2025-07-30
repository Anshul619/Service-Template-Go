package v1

import (
	"context"
	"net/http"
	"template-service/internal/model"
	"template-service/internal/service"
)

// Fake repo for demo
type fakeRepo struct{}

func (f *fakeRepo) FindByID(_ context.Context, id string) (*model.User, error) {
	return &model.User{ID: id, Name: "John Doe", Email: "john@example.com"}, nil
}

func InitRoutes() {
	repo := &fakeRepo{}
	svc := service.NewUserService(repo)
	handler := NewUserHandler(svc)

	// Mux - If we want to go mux ("github.com/gorilla/mux"), we can initilize HandleFunc like following
	// 	r := mux.NewRouter()
	// 	r.HandleFunc("/", myHandler)

	// GIN - r := gin.Default()
	// v1Group := r.Group("/api/v1")
	// {
	//     v1Group.GET("/users", v1.GetUsers)
	//     v1Group.POST("/users", v1.CreateUser)
	// }
	// return r

	// API v1 routes
	http.HandleFunc("/api/v1/users", handler.GetUser)
}
