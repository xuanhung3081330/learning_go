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

// Step 1: Interface
type UserRepo interface {
	GetUser(id int) string
}

// Step 2: Real implementation
type PostgresUserRepo struct {}

func (r PostgresUserRepo) GetUser(id int) string {
	return fmt.Sprintf("User%d from Postgres", id)
}

// Step 3: Mock implementation (for testing)
type MockUserRepo struct {}

func (m MockUserRepo) GetUser(id int) string {
	return fmt.Sprintf("Mocked user %d", id)
}

// Step 4: Service that depends on the UserRepo interface
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

// Step 5: main (Manual DI)
func main(){
	// In production
	repo := PostgresUserRepo{}
	service := NewUserService(repo) // Inject dependency
	service.PrintUser(2)

	// In testing (swap dependency)
	mockRepo := MockUserRepo{}
	mockService := NewUserService(mockRepo)
	mockService.PrintUser(4)
}
