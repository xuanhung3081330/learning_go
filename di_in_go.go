package main

import "fmt"

//type Service interface {
//	DoSomething()
//}

//type RealService struct{}

//func (s RealService) DoSomething() {
//	fmt.Println("Doing something real")
//}

//type Controller struct {
//	service Service
//}

//func NewController(s Service) *Controller {
//	return &Controller{
//		service: s,
//	}
//}

// Interface
type UserRepo interface {
	GetUser(id int) string
}

// Real implementation
type PostgresUserRepo struct {}

func (r PostgresUserRepo) GetUser(id int) string {
	return fmt.Sprintf("User%d from Postgres", id)
}

// Service
type UserService struct {
	repo UserRepo
}

func NewUserService(r UserRepo) *UserService {
	return &UserService{
		repo: r,
	}
}

func (s *UserService) PrintUser(id int) {
	fmt.Println(s.repo.GetUser(id))
}

func main(){
	repo := PostgresUserRepo{}
	service := NewUserService(repo) // Inject dependency
	service.PrintUser(2)
}
