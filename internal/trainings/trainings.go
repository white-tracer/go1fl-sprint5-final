package trainings

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	PersonalData "github.com/Yandex-Practicum/tracker/internal/personaldata"
	SpentEnergy "github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type Training struct {
	PersonalData.Personal

	Steps        int
	TrainingType string
	Duration     time.Duration
}

func (t *Training) Parse(datastring string) (err error) {
	parsedData := strings.Split(datastring, ",")
	if len(parsedData) != 3 {
		return errors.New("error in slice length")
	}

	t.Steps, err = strconv.Atoi(parsedData[0])
	if err != nil {
		return fmt.Errorf("conversion error: %w", err)
	}
	if t.Steps <= 0 {
		err := errors.New("error in steps count: zero or below")
		return err
	}

	t.TrainingType = parsedData[1]
	fmt.Println(t.TrainingType) // УБРАТЬ ебажное

	t.Duration, err = time.ParseDuration(parsedData[2])
	if err != nil {
		return fmt.Errorf("duration error: %w", err)
	}
	if t.Duration <= 0 {
		err := errors.New("incorrect duration")
		return err
	}
	return nil
}

func (t Training) ActionInfo() (string, error) {
	var spentCalories float64
	var err error
	distanse := SpentEnergy.Distance(t.Steps, float64(t.Personal.Height))
	meanSpeed := SpentEnergy.MeanSpeed(t.Steps, float64(t.Personal.Height), t.Duration)
	durationForPrint := fmt.Sprintf("%.2f", t.Duration.Hours())

	if t.TrainingType == "Бег" {
		spentCaloriesForRunTraining, err := SpentEnergy.RunningSpentCalories(t.Steps, t.Personal.Weight, t.Personal.Height, t.Duration)
		if err != nil {
			return "", err
		}
		spentCalories = spentCaloriesForRunTraining
	}
	if t.TrainingType == "Ходьба" {
		spentCaloriesForWalkTraining, err := SpentEnergy.WalkingSpentCalories(t.Steps, t.Personal.Weight, t.Personal.Height, t.Duration)
		if err != nil {
			return "", err
		}
		spentCalories = spentCaloriesForWalkTraining
	}
	if t.TrainingType != "Бег" && t.TrainingType != "Ходьба" {
		err := errors.New("error: unknown training type")
		return "", err
	}

	actionInfo := fmt.Sprintf("Тип тренировки: %s\nДлительность: %v ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", t.TrainingType, durationForPrint, distanse, meanSpeed, spentCalories)
	return actionInfo, err
}
