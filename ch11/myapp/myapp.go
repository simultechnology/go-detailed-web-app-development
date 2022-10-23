package myapp

import "errors"

type MyApp struct{}

func (m *MyApp) Apply(id int) error {
	return errors.New("custom error")
}
