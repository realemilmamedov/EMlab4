package main

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/google/uuid"

	"EMlab4/pkg/currencycalc"
)

func main() {
	// Идентификатор операции
	operationID := uuid.New()

	const (
		amount     = 1000.0 // сумма
		rate       = 92.35  // курс конвертации USD/RUB
		feePercent = 2.5    // % комиссии
		pair       = "USD/RUB"
	)

	color.Cyan("=-=-= Лаб работа №4. Вариант 10 =-=-=")
	fmt.Printf("ID операции: %s\n\n", operationID)

	// F1 конвертация суммы по курсу
	converted, err := currencycalc.Convert(amount, rate)
	if err != nil {
		color.Red("Ошибка конвертации: %v", err)
		os.Exit(1)
	}
	fmt.Printf("Конвертация: %8.2f USD * %.4f = %10.2f RUB\n", amount, rate, converted)

	// Расчёт суммы комиссии от полученного
	fee, err := currencycalc.FeeAmount(converted, feePercent)
	if err != nil {
		color.Red("Ошибка расчёта комиссии: %v", err)
		os.Exit(1)
	}
	fmt.Printf("Комиссия:    %5.2f%% от %10.2f RUB = %10.2f RUB\n", feePercent, converted, fee)

	// F2 комиссия
	finalAmount := converted
	if err := currencycalc.ApplyCommission(&finalAmount, fee); err != nil {
		color.Red("Ошибка применения комиссии: %v", err)
		os.Exit(1)
	}
	fmt.Printf("Итог: %.2f RUB\n\n", finalAmount)

	// F3 отчёт
	report, err := currencycalc.FormatCurrencyReport(pair, amount, finalAmount)
	if err != nil {
		color.Red("Ошибка формирования отчёта: %v", err)
		os.Exit(1)
	}
	color.Green("Отчёт: %s", report)

	// Ошибки
	fmt.Println("\n--- Проверка обработки ошибок ---")
	if _, err := currencycalc.Convert(-100, rate); err != nil {
		color.Yellow("Ожидаемая ошибка (отрицательная сумма): %v", err)
	}
	if _, err := currencycalc.FeeAmount(amount, 150); err != nil {
		color.Yellow("Ожидаемая ошибка (некорректный процент): %v", err)
	}
	badAmount := 10.0
	if err := currencycalc.ApplyCommission(&badAmount, 50); err != nil {
		color.Yellow("Ожидаемая ошибка (комиссия больше суммы): %v", err)
	}
}
