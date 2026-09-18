// Package currencycalc Набор функций для конвертации валют, расчёта комиссии и формирования отчётов
package currencycalc

import (
	"fmt"
)

// Convert конвертирует amount по заданному rate
func Convert(amount, rate float64) (float64, error) {
	if amount < 0 {
		return 0, fmt.Errorf("currencycalc.Convert: сумма не может быть отрицательной: %.2f", amount)
	}
	if rate <= 0 {
		return 0, fmt.Errorf("currencycalc.Convert: курс должен быть положительным числом: %.4f", rate)
	}

	result := amount * rate
	return result, nil
}

// FeeAmount вычисляет размер комиссии от amount по feePercent.
func FeeAmount(amount, feePercent float64) (float64, error) {
	if amount < 0 {
		return 0, fmt.Errorf("currencycalc.FeeAmount: сумма не может быть отрицательной: %.2f", amount)
	}
	if feePercent < 0 || feePercent > 100 {
		return 0, fmt.Errorf("currencycalc.FeeAmount: некорректный процент комиссии: %.2f (допустимо 0..100)", feePercent)
	}

	fee := amount * feePercent / 100
	return fee, nil
}

// ApplyCommission изменяет значение, на которое указывает amount, вычитая из него fee.
func ApplyCommission(amount *float64, fee float64) error {
	if amount == nil {
		return fmt.Errorf("currencycalc.ApplyCommission: передан nil-указатель на сумму")
	}
	if fee < 0 {
		return fmt.Errorf("currencycalc.ApplyCommission: комиссия не может быть отрицательной: %.2f", fee)
	}
	if *amount-fee < 0 {
		return fmt.Errorf("currencycalc.ApplyCommission: комиссия (%.2f) превышает сумму (%.2f)", fee, *amount)
	}

	*amount -= fee
	return nil
}

// FormatCurrencyReport формирует текстовый отчёт по валютной операции в формате "<исходная сумма> <pair> -> <итоговая сумма>".
func FormatCurrencyReport(pair string, src, dst float64) (string, error) {
	if pair == "" {
		return "", fmt.Errorf("currencycalc.FormatCurrencyReport: не указана валютная пара")
	}
	if src < 0 || dst < 0 {
		return "", fmt.Errorf("currencycalc.FormatCurrencyReport: суммы не могут быть отрицательными (src=%.2f, dst=%.2f)", src, dst)
	}

	report := fmt.Sprintf("Операция [%s]: %10.2f  ->  %10.2f", pair, src, dst)
	return report, nil
}
