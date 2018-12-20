package main

import (
	"fmt"

	"github.com/sirupsen/logrus"

	"github.com/suzuki-shunsuke/go-errlog"
)

func foo() (int, error) {
	age, err := getAge("foo")
	return age, errlog.Wrap(err, nil, "failed to foo")
}

func getAge(name string) (int, error) {
	return 0, errlog.New(logrus.Fields{
		"name": name,
	}, "invalid name", "failed to get an age")
}

func main() {
	logger := errlog.NewLogger(nil).
		WithFields(logrus.Fields{"program": "example1"})
	age, err := foo()
	logger.Fatal(err) // you don't have to check err is nil or not.
	fmt.Printf("age: %d\n", age)
}
