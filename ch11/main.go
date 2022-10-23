package main

import (
	"fmt"
	"go-detailed-web-app-development/ch11/myapp"
)

type ServiceImpl struct{}

func (s *ServiceImpl) Apply(id int) error {
	return nil
}

type OrderService interface {
	Apply(int) error
}

type Application struct {
	os OrderService
}

func (a *Application) Apply(id int) error {
	return a.os.Apply(id)
}

func NewApplication(os OrderService) *Application {
	return &Application{os: os}
}

func main() {
	fmt.Println("ch11 start!")
	application := NewApplication(&ServiceImpl{})
	err := application.Apply(19)
	if err != nil {
		fmt.Errorf("error : %v", err)
	}

	myApplication := NewApplication(&myapp.MyApp{})
	err = myApplication.Apply(20)
	if err != nil {
		err := fmt.Errorf("error : %v", err)
		fmt.Println(err.Error())
	}
}
