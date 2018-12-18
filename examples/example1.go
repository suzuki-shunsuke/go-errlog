package main

import (
	"fmt"

	"github.com/sirupsen/logrus"

	"github.com/suzuki-shunsuke/go-errlog"
)

func foo() (int, error) {
	age, err := getAge("foo")
	if err != nil {
		return 0, errlog.Wrap(err, nil, "failed to foo")
	}
	return age, err
}

func getAge(name string) (int, error) {
	return 0, errlog.New(logrus.Fields{
		"name": name,
	}, "invalid name", "failed to get an age")
}

func main() {
	logger := errlog.NewLogger(nil)
	logger.WithFields(logrus.Fields{"program": "example1"})
	age, err := foo()
	if err != nil {
		logger.Fatal(err)
	}
	fmt.Printf("age: %d\n", age)
}
