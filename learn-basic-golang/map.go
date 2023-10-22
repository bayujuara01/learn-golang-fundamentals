package main

import "fmt"

type Person_ struct {
	name string
	age  int8
}

func main() {
	by := Person_{
		age:  1,
		name: "lalal",
	}

	rex := map[string]Person_{
		"hey": by,
	}

	rex["xex"] = Person_{
		age:  2,
		name: "lalal2",
	}
	fmt.Println(rex["xex"])

	mapWithMake := make(map[string]Person_)

	mapWithMake["YAY"] = Person_{age: 1, name: "KL"}

	tempVar, ok := rex["xex"]

	if ok {
		fmt.Println(tempVar.name)
	}

	var numKey map[int8]string = map[int8]string{
		1: "lalal",
	}

	fmt.Println(numKey[1])

	for key, p := range rex {
		fmt.Println(key, p.name)
	}

	person := map[string]string{
		"name":    "Eko",
		"address": "Subang",
	}

	person["title"] = "Programmer"

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	var book map[string]string = make(map[string]string)
	book["title"] = "Belajar Go-Lang"
	book["author"] = "Eko"
	book["ups"] = "Salah"
	fmt.Println(book)
	fmt.Println(len(book))

	delete(book, "ups")

	fmt.Println(book)
	fmt.Println(len(book))
}
