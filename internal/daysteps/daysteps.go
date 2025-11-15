package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	PersonalData "github.com/Yandex-Practicum/tracker/internal/personaldata"
	SpentEnergy "github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type DaySteps struct {
	PersonalData.Personal

	Steps    int
	Duration time.Duration
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	parsedData := strings.Split(datastring, ",")
	if len(parsedData) != 2 {
		return errors.New("Длина слайса не равна двум")
	}

	ds.Steps, err = strconv.Atoi(parsedData[0])
	if err != nil {
		err := errors.New("Ошибка при получении количества шагов")
		return err
	}
	if ds.Steps <= 0 {
		err := errors.New("Шагов 0 или меньше")
		return err
	}

	parsedTime, err := time.ParseDuration(parsedData[1])
	if err != nil {
		return err
	}
	if parsedTime <= 0 {
		return errors.New("Длительность указана некорректно")
	}
	ds.Duration = parsedTime
	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	distanse := SpentEnergy.Distance(ds.Steps, float64(ds.Personal.Height))
	spentCalories, err := SpentEnergy.WalkingSpentCalories(ds.Steps, ds.Personal.Weight, ds.Personal.Height, ds.Duration)
	if err != nil {
		err := errors.New("Ошибка вычисления потраченных каллорий")
		return "", err
	}

	actionInfo := fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", ds.Steps, distanse, spentCalories)
	fmt.Println("Это строка экшн инфо:", actionInfo)
	return actionInfo, nil
}
