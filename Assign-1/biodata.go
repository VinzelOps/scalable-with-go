package main

import (
	"fmt"
	"os"
	"strconv"
)

type person struct { // membuat struct person yang di dalamnya terdapat beberapa data yang betipe string dan int
	name, address, job, reason string
}

func main() {
	input, _ := strconv.Atoi(os.Args[1])
	var allStudentClass = []person{
		{
			name:    "Kev",
			address: "Jakarta",
			job:     "Mahasiswa",
			reason:  "Belajar",
		},
		{
			name:    "Vin",
			address: "Depok",
			job:     "Mahasiswa",
			reason:  "Belajar",
		},
		{
			name:    "Zel",
			address: "Jakara",
			job:     "Mahasiswa",
			reason:  "Belajar"},
		{
			name:    "Vinzel",
			address: "Tangerang",
			job:     "Pelajar",
			reason:  "Belajar"},
		{
			name:    "Lezniv",
			address: "Banten",
			job:     "Pelajar",
			reason:  "Belajar",
		},
	}
	for i, student := range allStudentClass {
		num := input - 1
		if num == i {
			fmt.Printf("Nama: %s\nAlamat: %s\njob: %s\nreason: %s", student.name, student.address, student.job, student.reason)
		}
	}
}
