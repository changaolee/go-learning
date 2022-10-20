package main

import "github.com/changaolee/go-learning/log/cuslog"

func main() {
	l := cuslog.New(cuslog.WithLevel(cuslog.InfoLevel))
	l.Info("custom log with json formatter")
}
