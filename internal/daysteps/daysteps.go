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
		return errors.New("error in slice length: not equal 2")
	}

	ds.Steps, err = strconv.Atoi(parsedData[0])
	if err != nil {
		return fmt.Errorf("error in steps count: %w", err)
	}
	if ds.Steps <= 0 {
		err := errors.New("error in steps count: zero or below")
		return err
	}

	parsedTime, err := time.ParseDuration(parsedData[1])
	if err != nil {
		return err
	}
	if parsedTime <= 0 {
		return errors.New("error in duration: zero or below")
	}
	ds.Duration = parsedTime
	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	distanse := SpentEnergy.Distance(ds.Steps, float64(ds.Personal.Height))
	spentCalories, err := SpentEnergy.WalkingSpentCalories(ds.Steps, ds.Personal.Weight, ds.Personal.Height, ds.Duration)
	if err != nil {
		return "", fmt.Errorf("error in calories calculation: %w", err)
	}

	actionInfo := fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", ds.Steps, distanse, spentCalories)
	return actionInfo, nil
}
