package main

import (
	"fmt"
	"reflect"
)

type Sample struct {
	Nama string `ada:"true"`
	Age  int    `max:"30"`
	Nim  int
}

func isValid(value interface{}) bool {
	v := reflect.TypeOf(value)
	for a := 0; a < v.NumField(); a++ {
		field := v.Field(a)
		if field.Tag.Get("ada") == "true" {
			penanda2 := reflect.ValueOf(value).Field(a).Interface()
			if penanda2 == "" {
				return false
			}
		}
	}
	return true
	for i := 0; i < v.NumField(); i++ {
		t := v.Field(i)
		if t.Tag.Get("max") == "30" {
			tM := reflect.ValueOf(value).Field(i).Interface()
			if tM == "" {
				return false
			}
		}
	}
	return true
}

func main() {
	dS := Sample{
		"Rizki",
		20,
		20200801041,
	}

	var sampleDs reflect.Type = reflect.TypeOf(dS)
	fmt.Println(sampleDs.NumField())
	fmt.Println(sampleDs.Field(0).Name)
	fmt.Println(sampleDs.Field(1).Name)
	fmt.Println(sampleDs.Field(2).Name)
	fmt.Println(sampleDs.Field(0).Tag.Get("ada"))
	fmt.Println(sampleDs.Field(1).Tag.Get("max"))

	fmt.Println(isValid(dS))
	fmt.Println(isValid(dS))
	dS.Nama = ""
	fmt.Println(isValid(dS))
	dS.Age = 10
	fmt.Println(isValid(dS))
}
