# EMlab4 — currencycalc (конвертация валют)

Лабораторная работа №4, вариант 10. Пакет `currencycalc` реализует
конвертацию валют, расчёт и применение комиссии, а также формирование
текстового отчёта по операции.

## Структура проекта

```
EMlab4/
 go.mod
 cmd/app/main.go
 pkg/currencycalc/currencycalc.go
 README.md
```

## Функции пакета `currencycalc`

| Функция | Сигнатура | Роль |
|---|---|---|
| `Convert` | `func Convert(amount, rate float64) (float64, error)` | F1 — вычислительная функция: конвертация суммы по курсу |
| `FeeAmount` | `func FeeAmount(amount, feePercent float64) (float64, error)` | вычисление размера комиссии |
| `ApplyCommission` | `func ApplyCommission(amount *float64, fee float64) error` | F2 — функция с указателем: списывает комиссию из суммы "на месте" |
| `FormatCurrencyReport` | `func FormatCurrencyReport(pair string, src, dst float64) (string, error)` | F3 — формирование строки отчёта через `fmt.Sprintf` |

Все экспортируемые функции:
- названы в PascalCase;
- снабжены doc-комментариями;
- проверяют входные данные и возвращают ошибки через `fmt.Errorf`.

## Внешние зависимости

Проект подключает два внешних пакета:

1. **github.com/fatih/color** — цветной вывод в консоль (заголовки, ошибки,
   успешные отчёты, предупреждения).
2. **github.com/google/uuid** — публичный пакет, генерирующий уникальный
   идентификатор операции (`uuid.New()`), используемый как "чужой готовый
   пакет" по требованиям задания.

## Установка зависимостей и запуск

```bash
go mod tidy
go run ./cmd/app
```

## Пример вывода

```
=== Лабораторная работа №4. Вариант 10: currencycalc ===
ID операции: 5f2c1e2a-....

Конвертация:  1000.00 USD * 92.3500 =   92350.00 RUB
Комиссия:     2.50% от   92350.00 RUB =    2308.75 RUB
Итог после комиссии: 90041.25 RUB

Отчёт: Операция [USD/RUB]:    1000.00  ->   90041.25

--- Проверка обработки ошибок ---
Ожидаемая ошибка (отрицательная сумма): currencycalc.Convert: сумма не может быть отрицательной: -100.00
Ожидаемая ошибка (некорректный процент): currencycalc.FeeAmount: некорректный процент комиссии: 150.00 (допустимо 0..100)
Ожидаемая ошибка (комиссия больше суммы): currencycalc.ApplyCommission: комиссия (50.00) превышает сумму (10.00)
```

## Документация пакета

Просмотр документации в терминале:

```bash
go doc ./...
go doc ./pkg/currencycalc
```

Просмотр документации в браузере:

```bash
go install golang.org/x/tools/cmd/godoc@latest
godoc -http=:6060
# затем открыть http://localhost:6060/pkg/EMlab4/pkg/currencycalc/
```

## Публикация в GitHub

```bash
git init
git add .
git commit -m "lab4 variant10: currencycalc package"
git branch -M main
git remote add origin https://github.com/<username>/EMlab4.git
git push -u origin main
```

## Подключение чужого/публичного пакета

В данном проекте в качестве "чужого готового пакета" используется
публичный пакет `github.com/google/uuid` для генерации идентификатора
операции. При необходимости замените его на пакет одногруппника,
указав его репозиторий в `go.mod` и импортировав соответствующую функцию
в `cmd/app/main.go`.
