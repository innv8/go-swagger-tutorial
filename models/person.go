package models

import "go-swagger-tutorial/entities"

var people = []entities.Person{}

func AddPerson(name string, age uint) (id int) {
	id = len(people) + 1
	people = append(people, entities.Person{
		Id:   id,
		Name: name,
		Age:  age,
	})
	return id
}

func GetPersons() []entities.Person {
	return people
}
