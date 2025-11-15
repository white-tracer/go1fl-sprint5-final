package actioninfo

import (
	"fmt"
	"log"
)

type DataParser interface {
	Parse(datastring string) (err error)
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	for _, value := range dataset {
		err := dp.Parse(value)
		if err != nil {
			log.Println(err)
			continue
		}
		trainingInfo, err := dp.ActionInfo()
		if err != nil {
			log.Println("error in trainingInfo function:", err)
			continue
		}
		fmt.Println(trainingInfo)
	}
}
