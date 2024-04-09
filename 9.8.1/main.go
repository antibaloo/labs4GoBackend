package main

import (
	"errors"
	"fmt"
)

type Man struct {
	Name     string
	LastName string
	Age      int
	Gender   string
	Crimes   int
}

func (m Man) String() string {
	return fmt.Sprintf("Имя: %s\nФамилия: %s\nВозраст:%d\nПол: %s\nКол-во преступлений: %d", m.Name, m.LastName, m.Age, m.Gender, m.Crimes)
}

func main() {
	people := map[string]Man{
		"Андрей": {
			Name:     "Андрей",
			LastName: "Иванов",
			Age:      32,
			Gender:   "мужской",
			Crimes:   2,
		},
		"Светлана": {
			Name:     "Свелана",
			LastName: "Петрова",
			Age:      24,
			Gender:   "женский",
			Crimes:   0,
		},
		"Иван": {
			Name:     "Иван",
			LastName: "Сидоров",
			Age:      37,
			Gender:   "мужской",
			Crimes:   6,
		},
		"Людмила": {
			Name:     "Людмила",
			LastName: "Смирнова",
			Age:      31,
			Gender:   "женский",
			Crimes:   1,
		},
		"Михаил": {
			Name:     "Михаил",
			LastName: "Старков",
			Age:      54,
			Gender:   "мужской",
			Crimes:   20,
		},
		"Евгений": {
			Name:     "Евгений",
			LastName: "Кротов",
			Age:      40,
			Gender:   "мужской",
			Crimes:   3,
		},
		"Николай": {
			Name:     "Николай",
			LastName: "Белов",
			Age:      47,
			Gender:   "мужской",
			Crimes:   1,
		},
	}
	suspects := []string{"Игорь", "Константин"}
	fmt.Println("Список подозреваемых:")
	for i, person := range suspects {
		fmt.Println(i+1, person)
	}
	fmt.Println("--------------------")

	most, err := mostWantedPerson(suspects, people)
	if err != nil {
		fmt.Printf("%s", err)
	} else {
		fmt.Printf("Подозреваемый с наибольшим количеством преступлений:\n%s", most)
	}

}

func mostWantedPerson(s []string, p map[string]Man) (Man, error) {
	var res Man = Man{}
	var maxCrimes int
	var found bool
	if len(s) == 0 {
		return res, errors.New("В базе данных нет информации по запрошенным подозреваемым")
	}
	for _, suspect := range s {
		if value, ok := p[suspect]; ok {
			found = true
			if value.Crimes >= maxCrimes {
				res = p[suspect]
				maxCrimes = p[suspect].Crimes
			}
		}
	}
	if !found {
		return res, errors.New("В базе данных нет информации по запрошенным подозреваемым")
	}
	return res, nil
}
