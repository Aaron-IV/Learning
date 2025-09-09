package array

import (
	"fmt"

	"github.com/k0kubun/pp"
)

type User struct {
	Name    string
	Rating  float64
	Premium bool
}

func main() {
	userArray := [3]User{
		{
			Name:    "Azelit",
			Rating:  3.9,
			Premium: true,
		},
		{
			Name:    "Миф",
			Rating:  2.5,
			Premium: false,
		},
		{
			Name:    "Мистер Пропер",
			Rating:  3.3,
			Premium: true,
		},
	}

	fmt.Println("Before:")
	fmt.Println("<-------->")
	for index, user := range userArray {
		pp.Println(index, user)
	}
	fmt.Println("")

	for i := 0; i < len(userArray); i++ {
		if userArray[i].Premium {
			userArray[i].Rating += 1.0
		}
	}

	fmt.Println("After:")
	fmt.Println("<-------->")
	for i := 0; i < len(userArray); i++ {
		pp.Println(userArray[i])
	}
	fmt.Println("")
}
