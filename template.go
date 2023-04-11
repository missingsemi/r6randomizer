package main

type TemplateData[T any] struct {
	Side     string
	Operator T
}
