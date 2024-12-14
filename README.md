# Math Helper

Реализация основных математических алгоритмов на языке Go.

> [!warning]
> Этот проект в данный момент находится в разработке. Далеко не все функции
реализованы, и могут встречаться ошибки.

## Использование

Программа пока что не имеет пользовательского интерфейса. Это библиотека
(моудль) Go. Чтобы воспользоваться программой, необходимо клонировать
репозиторий и написать свою программу в файле `main.go`. Для запуска потребуется
установить [инструментарий языка Go](https://go.dev/dl/).

```sh
git clone https://github.com/wadrodrog/math-helper
cd math-helper
vim main.go  # Редактировать файл
go run .     # Запустить программу
```

## Документация

Используйте инструмент [godoc](https://pkg.go.dev/golang.org/x/tools/cmd/godoc)
для генерации документации.

```sh
go install golang.org/x/tools/cmd/godoc@latest
godoc -http :8080
# Откройте в браузере http://localhost:8080/
```

## Цели

Целью этого проекта является изучение математических алгоритмов и языка
программирования Go.

Целью этого проекта НЕ является эффективная реализация алгоритмов.

## Алгоритмы

- Перестановки
    - Вычисление количества инверсий
    - Определение чётности перестановки
    - Разложение на циклы и транспозиции
    - Сборка из циклов транспозиций
    - Умножение перестановок

## Использованные технологии и возможности

- Unit-тестирование
- Документация (godoc)

## Ресурсы

- [Знак перестановки: транспозиции vs инверсии — Хабр](https://habr.com/ru/articles/762338/)

## Лицензия

[GNU General Public License v3.0 or later](https://www.gnu.org/licenses/gpl-3.0.html)
