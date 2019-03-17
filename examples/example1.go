package main

import (
	"fmt"

	"github.com/sirupsen/logrus"

	"github.com/suzuki-shunsuke/go-errlog"
)

func foo() (int, error) {
	age, err := getAge("foo")
	return age, errlog.Wrap(err, logrus.Fields{
		"name": "foo",
	}, "failed to get age")
}

func getAge(name string) (int, error) {
	return 0, fmt.Errorf("invalid name")
}

func main() {
	logger := errlog.NewLogger(nil).
		WithFields(logrus.Fields{"program": "example1"})
	age, err := foo()
	logger.Fatal(err, nil, "function foo is failure") // you don't have to check err is nil or not.
	fmt.Printf("age: %d\n", age)
}
