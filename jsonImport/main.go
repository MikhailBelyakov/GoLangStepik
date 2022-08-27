package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type userStruct struct {
	GlobalId int `json:"global_id"`
}

/**
Данная задача ориентирована на реальную работу с данными в формате JSON. Для решения мы будем использовать справочник ОКВЭД (Общероссийский классификатор видов экономической деятельности), который был размещен на web-портале data.gov.ru.

Необходимая вам информация о структуре данных содержится в файле struct-20190514T0000.json, а сами данные, которые вам потребуется декодировать, содержатся в файле data-20190514T0100.json. Файлы размещены в нашем репозитории на github.com.

Для того, чтобы показать, что вы действительно смогли декодировать документ вам необходимо в качестве ответа записать сумму полей global_id всех элементов, закодированных в файле.
*/
func main() {
	dataFromFile, _ := ioutil.ReadFile("jsonImport/data.json")

	var inputS []userStruct
	var total int

	err := json.Unmarshal(dataFromFile, &inputS)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, val := range inputS {
		total += val.GlobalId
	}

	fmt.Println(total)
}
