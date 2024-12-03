# go-ca-template-generator

Генератор структуры проекта — это инструмент командной строки, написанный на Go, который помогает быстро настроить новую структуру проекта на основе принципов чистой архитектуры. Этот инструмент автоматизирует создание директорий, файлов и необходимых конфигурационных файлов, облегчая начало нового проекта с хорошо организованной структурой.

## Возможности
Инициализация проекта: Автоматически создает необходимые директории и файлы для нового проекта.
Интеграция с Git: Опционально инициализирует репозиторий Git.
Конфигурационные Файлы: Генерирует файлы go.mod, .gitignore, .env, Dockerfile, Makefile и README.md.
Настраиваемая Структура: Позволяет указать сущности, сервисы, use case, интерфейсы, инфраструктуру, приложения и репозитории.

## Начало работы
### Предварительные условия
- Установлен Go на вашем компьютере. Вы можете скачать его с golang.org.
- Установлен Git на вашем компьютере. Вы можете скачать его с git-scm.com.

## Установка

Клонируйте репозиторий:
```sh
git clone https://github.com/andrewovchinnikov/go-ca-template-generator.git
```
Перейдите в директорию проекта:
```sh
cd go-ca-template-generator
```
Запустите генераторт:
```sh
go run main.go
```
Следуйте подсказкам для указания названия проекта, сущностей, сервисов, use case, интерфейсов, инфраструктуры, приложений, репозиториев и других конфигурационных опций.

## Пример
Использование инструмента:
```sh
Введите название проекта: testApp
Выполнить git init? (да/нет): да
Введите названия сущностей (через запятую): User,Product
Введите названия сервисов (через запятую): UserService,ProductService,Notification
Будем указывать названия usecase? (да/нет): да
Введите названия usecase (через запятую): CreateUser,CreateProduct
Будем указывать названия interface? (да/нет): да
Введите названия interface (через запятую): notification
Будем указывать названия infrastructure? (да/нет): да
Введите названия infrastructure (через запятую): db,cache
Будем указывать названия application? (да/нет): да
Введите названия application (через запятую): UserApp,ProductApp
Будем указывать названия repositories? (да/нет): да
Введите названия repositories (через запятую): UserRepository,ProductRepository
Создать файл go.mod? (да/нет): да
Выберите версию Go (например, 1.18): 1.18
Создать файл .gitignore? (да/нет): да
Создать файл .env? (да/нет): да
Создать Dockerfile? (да/нет): да
Создать Makefile? (да/нет): да
Создать README.md? (да/нет): да
Структура проекта создана успешно!
```

### Структура директорий
```sh
generated/
├── testApp/
│   ├── cmd/
│   │   └── main.go
│   ├── internal/
│   │   ├── domain/
│   │   │   ├── User/
│   │   │   │   └── User.go
│   │   │   └── Product/
│   │   │       └── Product.go
│   │   ├── usecase/
│   │   │   ├── UserService/
│   │   │   │   └── UserService.go
│   │   │   ├── ProductService/
│   │   │   │   └── ProductService.go
│   │   │   ├── Notification/
│   │   │   │   └── Notification.go
│   │   │   ├── CreateUser/
│   │   │   │   └── CreateUser.go
│   │   │   └── CreateProduct/
│   │   │       └── CreateProduct.go
│   │   ├── interface/
│   │   │   └── notification/
│   │   │       └── notification.go
│   │   ├── infrastructure/
│   │   │   ├── db/
│   │   │   │   └── db.go
│   │   │   ├── cache/
│   │   │   │   └── cache.go
│   │   │   └── repositories/
│   │   │       ├── UserRepository/
│   │   │       │   └── UserRepository.go
│   │   │       └── ProductRepository/
│   │   │           └── ProductRepository.go
│   │   └── application/
│   │       ├── UserApp/
│   │       │   └── UserApp.go
│   │       └── ProductApp/
│   │           └── ProductApp.go
│   ├── pkg/
│   ├── api/
│   ├── config/
│   ├── go.mod
│   ├── .gitignore
│   ├── .env
│   ├── Dockerfile
│   ├── Makefile
│   └── README.md

```