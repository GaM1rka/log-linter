# Log Linter

Кастомный линтер для проверки лог-сообщений в Go-коде.

Реализован на базе `go/analysis` и интегрирован как module plugin для **golangci-lint v2**.

## Возможности

Линтер проверяет лог-сообщения на соответствие правилам:

✅ Сообщение начинается со строчной буквы  
✅ Используются только английские буквы  
✅ Отсутствуют emoji и специальные символы  
✅ Отсутствуют потенциально чувствительные данные (password, token, apiKey и т.д.)

## Поддерживаемые логгеры

- `log/slog`
- `go.uber.org/zap`

## Использование

### 1️⃣ Запуск как отдельный CLI

**Сборка:**
```bash
go build -o loglinter ./cmd/loglinter
```
**Запуск**
```bash
./loglinter ./...
```

### Интеграция с golangci-lint (v2)

Линтер реализован как **module plugin** для golangci-lint v2.

### 1 Создать `.custom-gcl.yml`

```yaml
version: v2.9.0

plugins:
  - module: github.com/GaM1rka/log-linter
    import: github.com/GaM1rka/log-linter/plugin/golangci
    version: v0.1.4
```
### 2 Собрать кастомный бинарник
```bash
golangci-lint custom -v
```
### 3 Настроить .golangci.yml
```yaml
version: "2"

linters:
  default: none
  enable:
    - loglinter

  settings:
    custom:
      loglinter:
        type: module
        settings:
          enableLowercase: true
          enableEnglishOnly: true
          enableNoSpecial: true
          enableNoSensitive: true
```
### 4 Запуск линтера
```bash
./custom-gcl run ./...
```
# Пример вывода:
<img width="846" height="208" alt="image" src="https://github.com/user-attachments/assets/536a9ce7-dff4-4421-af39-0533a4bbb0d1" />

## Конфигурация правил

Доступны следующие настройки:

| Параметр          | Описание                          |
|-------------------|-----------------------------------|
| `enableLowercase` | Проверка на строчную первую букву |
| `enableEnglishOnly` | Разрешены только английские буквы |
| `enableNoSpecial` | Запрет emoji и спецсимволов      |
| `enableNoSensitive` | Проверка на чувствительные данные |

## Тестирование

**Запуск тестов:**
```bash
go test ./... -v
```
В проекте реализованы:
- Unit-тесты для правил
- Тестирование анализатора через analysistest
- Тестовые файлы в testdata/

## CI

Проект использует **GitHub Actions**:

- Проверка линтера
- Запуск unit-тестов
- Проверка сборки

## Структура проекта
```swift
cmd/loglinter/        CLI входная точка
pkg/analyzer/         Основная логика анализатора
plugin/golangci/      Плагин для golangci-lint v2
```

## Версия Go

**Go 1.22+**

## Дальнейшее развитие

Возможные улучшения:

- Поддержка **SuggestedFix** (автоисправление)
- Расширяемые паттерны чувствительных данных
- Улучшенная система конфигурации (+ исправление текущих ошибок)

