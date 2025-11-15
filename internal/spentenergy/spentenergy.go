package spentenergy

import (
	"errors"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {

	if steps <= 0 {
		err := errors.New("Ошибка функции 'WalkingSpentCalories' - Шаги указаны некорректно")
		return 0, err
	}
	if weight <= 0 {
		err := errors.New("Ошибка функции 'WalkingSpentCalories' - Вес указан некорректно")
		return 0, err
	}

	if height <= 0 {
		err := errors.New("Ошибка функции 'WalkingSpentCalories' - Рост указан некорректно")
		return 0, err
	}

	if duration <= 0 {
		err := errors.New("Ошибка функции 'WalkingSpentCalories' - Длительность тренировка указана некорректно")
		return 0, err
	}

	walkingSpentCalories := (weight * MeanSpeed(steps, height, duration) * float64(duration.Minutes())) / float64(minInH) * walkingCaloriesCoefficient
	if walkingSpentCalories <= 0 {
		err := errors.New("Ошибка функции 'WalkingSpentCalories' - Неверный подсчет каллорий")
		return 0, err
	}
	return walkingSpentCalories, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {

	if steps <= 0 {
		err := errors.New("Ошибка функции 'RunningSpentCalories' - Шаги указаны некорректно")
		return 0, err
	}
	if weight <= 0 {
		err := errors.New("Ошибка функции 'RunningSpentCalories' - Вес указан некорректно")
		return 0, err
	}

	if height <= 0 {
		err := errors.New("Ошибка функции 'RunningSpentCalories' - Рост указан некорректно")
		return 0, err
	}

	if duration <= 0 {
		err := errors.New("Ошибка функции 'RunningSpentCalories' - Длительность тренировка указана некорректно")
		return 0, err
	}

	spentCalories := (weight * MeanSpeed(steps, height, duration) * float64(duration.Minutes())) / float64(minInH)
	if spentCalories <= 0 {
		err := errors.New("Ошибка функции 'RunningSpentCalories' - Неверный подсчет каллорий")
		return 0, err
	}
	return spentCalories, nil
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if steps < 0 || duration <= 0 {
		return 0
	}
	meanSpeed := Distance(steps, height) / float64(duration.Hours())
	return meanSpeed
}

func Distance(steps int, height float64) float64 {
	stepLength := height * stepLengthCoefficient
	distance := float64(stepLength) * float64(steps) / float64(mInKm)
	return distance
}
