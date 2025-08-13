package main

import "fmt"

type Service interface {
	DoSomething()
}

type RealService struct{}

func (s RealService) DoSomething() {
	fmt.Println("Doing something real")
}

type Controller struct {
	service Service
}

func NewController(s Service) *Controller {
	return &Controller{
		service: s,
	}
}
