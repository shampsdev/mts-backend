# Бекенд репозиторий
Посмотреть результат тут: https://api.mts.shamps.dev

## Команда "Шампиньоны"
 - Александр Дьяконов [DevOps-инженер, backend-разработчик]
 - Мишель де Джофрой [frontend-разработчик]
 - Виктория Кулешова [UX/UI дизайнер]
 - Иван Тарасов [backend-разработчик, DevOps-инженер]
 - Анастасия Богданова [бизнес-аналитик]

## Решение
В качестве основного языка был выбрал `Go` версии 1.23. Проект интегрируется с `Kubernetes` и `Helm` для удобного разворачивания и управления инфраструктурой.

Мы решили полностью отказаться от хранения данных у себя в целях сохранения консистентности данных. Поэтому для того, чтобы подключить внешние базы данных или АПИ достаточно написать `adapter`, чтобы наш сервис смог их отдать в обработанном виде фронтенду. Интерфейс для реализации находится в `external`. Сейчас там дублируется код (модель `domain.Person` совпадает с `json` моделью), но это сделано в качестве примера реализации адаптера.

Для сборки образа бекенда используется `scratch` чтобы максимально облегчить его.

> Мы пришли к этому далеко не сразу. Сначала писали на питоне, хотели подключить эластик и возможно как-то обойтись без бекенда, но оказалось, что можно поступить еще лучше... Написать на Go!

## Технологии
- **Go** (версия 1.23)
- **Bleve** — библиотека для полнотекстового поиска
- **Kubernetes** — для развертывания и управления
- **Helm** — для упрощения процесса развертывания в Kubernetes

## Шаги для запуска

Возможен запуск отдельно сервиса без Docker'a и кубера

### Я выбираю docker! (или кубер)
Для этого мы сделали отдельный [деплой репозиторий](https://github.com/shampsdev/mts-deploy)

 - https://github.com/shampsdev/mts-deploy

В нем подробно описан запуск

### Путь без докера
Для установки зависимостей проекта, выполните команду:

```bash
go mod tidy
```

Чтобы запустить проект, используйте команду:
```
go run cmd/main.go
```