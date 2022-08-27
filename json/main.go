package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type student struct {
	LastName   string
	FirstName  string
	MiddleName string
	Birthday   string
	Address    string
	Phone      string
	Rating     []float32
}

type groupStruct struct {
	Id       int
	Number   string
	Year     int
	Students []student
}

type outputStruct struct {
	Average float32
}

/**
На стандартный ввод подаются данные о студентах университетской группы в формате JSON:
{
    "ID":134,
    "Number":"ИЛМ-1274",
    "Year":2,
    "Students":[
        {
            "LastName":"Вещий",
            "FirstName":"Лифон",
            "MiddleName":"Вениаминович",
            "Birthday":"4апреля1970года",
            "Address":"632432,г.Тобольск,ул.Киевская,дом6,квартира23",
            "Phone":"+7(948)709-47-24",
            "Rating":[1,2,3]
        },
        {
            "LastName":"Вещий",
            "FirstName":"Лифон",
            "MiddleName":"Вениаминович",
            "Birthday":"4апреля1970года",
            "Address":"632432,г.Тобольск,ул.Киевская,дом6,квартира23",
            "Phone":"+7(948)709-47-24",
            "Rating":[1,2,3,4,4,5]
        },
		{
            "LastName":"Вещий",
            "FirstName":"Лифон",
            "MiddleName":"Вениаминович",
            "Birthday":"4апреля1970года",
            "Address":"632432,г.Тобольск,ул.Киевская,дом6,квартира23",
            "Phone":"+7(948)709-47-24",
            "Rating":[1,3]
        }
    ]
}

*/
func main() {
	// для проверки подставляем. data, _ := ioutil.ReadAll(os.Stdin)

	data := []byte("{\n    \"ID\":134,\n    \"Number\":\"ИЛМ-1274\",\n    \"Year\":2,\n    \"Students\":[\n        {\n            \"LastName\":\"Вещий\",\n            \"FirstName\":\"Лифон\",\n            \"MiddleName\":\"Вениаминович\",\n            \"Birthday\":\"4апреля1970года\",\n            \"Address\":\"632432,г.Тобольск,ул.Киевская,дом6,квартира23\",\n            \"Phone\":\"+7(948)709-47-24\",\n            \"Rating\":[1,2,3]\n        },\n        {\n            \"LastName\":\"Вещий\",\n            \"FirstName\":\"Лифон\",\n            \"MiddleName\":\"Вениаминович\",\n            \"Birthday\":\"4апреля1970года\",\n            \"Address\":\"632432,г.Тобольск,ул.Киевская,дом6,квартира23\",\n            \"Phone\":\"+7(948)709-47-24\",\n            \"Rating\":[1,2,3,4,4,5]\n        },\n\t\t{\n            \"LastName\":\"Вещий\",\n            \"FirstName\":\"Лифон\",\n            \"MiddleName\":\"Вениаминович\",\n            \"Birthday\":\"4апреля1970года\",\n            \"Address\":\"632432,г.Тобольск,ул.Киевская,дом6,квартира23\",\n            \"Phone\":\"+7(948)709-47-24\",\n            \"Rating\":[1,3]\n        }\n    ]\n}")

	var inputGroup groupStruct

	err := json.Unmarshal(data, &inputGroup)

	if err != nil {
		return
	}

	var countRating float32 = 0
	var countStudent float32 = 0

	for _, student := range inputGroup.Students {
		for range student.Rating {
			countRating += 1
		}
		countStudent += 1
	}

	result := outputStruct{}
	result.Average = countRating / countStudent

	output, _ := json.MarshalIndent(result, "", "    ")

	io.WriteString(os.Stdout, fmt.Sprintf("%s", output))
}
