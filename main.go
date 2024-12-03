package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Введите название проекта: ")
	projectName, _ := reader.ReadString('\n')
	projectName = strings.TrimSpace(projectName)

	fmt.Print("Выполнить git init? (да/нет): ")
	gitInitInput, _ := reader.ReadString('\n')
	gitInitInput = strings.TrimSpace(gitInitInput)

	fmt.Print("Введите названия сущностей (через запятую): ")
	entitiesInput, _ := reader.ReadString('\n')
	entities := splitAndTrim(entitiesInput)

	fmt.Print("Введите названия сервисов (через запятую): ")
	servicesInput, _ := reader.ReadString('\n')
	services := splitAndTrim(servicesInput)

	fmt.Print("Будем указывать названия usecase? (да/нет): ")
	usecaseInput, _ := reader.ReadString('\n')
	usecaseInput = strings.TrimSpace(usecaseInput)
	var usecases []string
	if strings.ToLower(usecaseInput) == "да" {
		fmt.Print("Введите названия usecase (через запятую): ")
		usecasesInput, _ := reader.ReadString('\n')
		usecases = splitAndTrim(usecasesInput)
	}

	fmt.Print("Будем указывать названия interface? (да/нет): ")
	interfaceInput, _ := reader.ReadString('\n')
	interfaceInput = strings.TrimSpace(interfaceInput)
	var interfaces []string
	if strings.ToLower(interfaceInput) == "да" {
		fmt.Print("Введите названия interface (через запятую): ")
		interfacesInput, _ := reader.ReadString('\n')
		interfaces = splitAndTrim(interfacesInput)
	}

	fmt.Print("Будем указывать названия infrastructure? (да/нет): ")
	infrastructureInput, _ := reader.ReadString('\n')
	infrastructureInput = strings.TrimSpace(infrastructureInput)
	var infrastructures []string
	if strings.ToLower(infrastructureInput) == "да" {
		fmt.Print("Введите названия infrastructure (через запятую): ")
		infrastructuresInput, _ := reader.ReadString('\n')
		infrastructures = splitAndTrim(infrastructuresInput)
	}

	fmt.Print("Будем указывать названия application? (да/нет): ")
	applicationInput, _ := reader.ReadString('\n')
	applicationInput = strings.TrimSpace(applicationInput)
	var applications []string
	if strings.ToLower(applicationInput) == "да" {
		fmt.Print("Введите названия application (через запятую): ")
		applicationsInput, _ := reader.ReadString('\n')
		applications = splitAndTrim(applicationsInput)
	}

	fmt.Print("Будем указывать названия repositories? (да/нет): ")
	repositoriesInput, _ := reader.ReadString('\n')
	repositoriesInput = strings.TrimSpace(repositoriesInput)
	var repositories []string
	if strings.ToLower(repositoriesInput) == "да" {
		fmt.Print("Введите названия repositories (через запятую): ")
		repositoriesInput, _ := reader.ReadString('\n')
		repositories = splitAndTrim(repositoriesInput)
	}

	fmt.Print("Создать файл go.mod? (да/нет): ")
	goModInput, _ := reader.ReadString('\n')
	goModInput = strings.TrimSpace(goModInput)
	var goVersion string
	if strings.ToLower(goModInput) == "да" {
		fmt.Print("Выберите версию Go (например, 1.18): ")
		goVersionInput, _ := reader.ReadString('\n')
		goVersion = strings.TrimSpace(goVersionInput)
	}

	fmt.Print("Создать файл .gitignore? (да/нет): ")
	gitignoreInput, _ := reader.ReadString('\n')
	gitignoreInput = strings.TrimSpace(gitignoreInput)

	fmt.Print("Создать файл .env? (да/нет): ")
	envInput, _ := reader.ReadString('\n')
	envInput = strings.TrimSpace(envInput)

	fmt.Print("Создать Dockerfile? (да/нет): ")
	dockerfileInput, _ := reader.ReadString('\n')
	dockerfileInput = strings.TrimSpace(dockerfileInput)

	fmt.Print("Создать Makefile? (да/нет): ")
	makefileInput, _ := reader.ReadString('\n')
	makefileInput = strings.TrimSpace(makefileInput)

	fmt.Print("Создать README.md? (да/нет): ")
	readmeInput, _ := reader.ReadString('\n')
	readmeInput = strings.TrimSpace(readmeInput)

	createProjectStructure(projectName, entities, services, usecases, interfaces, infrastructures, applications, repositories, goModInput, goVersion, gitignoreInput, envInput, gitInitInput, dockerfileInput, makefileInput, readmeInput)
	fmt.Println("Структура проекта создана успешно!")
}

func splitAndTrim(input string) []string {
	input = strings.TrimSpace(input)
	return strings.Split(input, ",")
}

func createProjectStructure(projectName string, entities, services, usecases, interfaces, infrastructures, applications, repositories []string, goModInput, goVersion, gitignoreInput, envInput, gitInitInput, dockerfileInput, makefileInput, readmeInput string) {
	generatedPath := filepath.Join(".", "generated", projectName)
	os.MkdirAll(generatedPath, 0755)

	// Выполняем git init, если выбрано
	if strings.ToLower(gitInitInput) == "да" {
		cmd := exec.Command("git", "init")
		cmd.Dir = generatedPath
		err := cmd.Run()
		if err != nil {
			fmt.Println("Ошибка при выполнении git init:", err)
		}
	}

	// Создаем основные директории
	createDir(filepath.Join(generatedPath, "cmd"))
	createDir(filepath.Join(generatedPath, "internal"))
	createDir(filepath.Join(generatedPath, "pkg"))
	createDir(filepath.Join(generatedPath, "api"))
	createDir(filepath.Join(generatedPath, "config"))
	createDir(filepath.Join(generatedPath, "internal", "domain"))
	createDir(filepath.Join(generatedPath, "internal", "usecase"))
	createDir(filepath.Join(generatedPath, "internal", "interface"))
	createDir(filepath.Join(generatedPath, "internal", "infrastructure"))
	createDir(filepath.Join(generatedPath, "internal", "application"))
	createDir(filepath.Join(generatedPath, "internal", "infrastructure", "repositories"))

	// Создаем файл main.go в директории cmd
	createMainFile(filepath.Join(generatedPath, "cmd", "main.go"), projectName)

	// Создаем директории и файлы для сущностей
	for _, entity := range entities {
		entityPath := filepath.Join(generatedPath, "internal", "domain", entity)
		createDir(entityPath)
		createFileWithPackage(filepath.Join(entityPath, entity+".go"), entity)
	}

	// Создаем директории и файлы для сервисов
	for _, service := range services {
		servicePath := filepath.Join(generatedPath, "internal", "usecase", service)
		createDir(servicePath)
		createFileWithPackage(filepath.Join(servicePath, service+".go"), service)
	}

	// Создаем директории и файлы для usecase
	if len(usecases) > 0 {
		for _, usecase := range usecases {
			usecasePath := filepath.Join(generatedPath, "internal", "usecase", usecase)
			createDir(usecasePath)
			createFileWithPackage(filepath.Join(usecasePath, usecase+".go"), usecase)
		}
	}

	// Создаем директории и файлы для interface
	if len(interfaces) > 0 {
		for _, iface := range interfaces {
			ifacePath := filepath.Join(generatedPath, "internal", "interface", iface)
			createDir(ifacePath)
			createFileWithPackage(filepath.Join(ifacePath, iface+".go"), iface)
		}
	}

	// Создаем директории и файлы для infrastructure
	if len(infrastructures) > 0 {
		for _, infrastructure := range infrastructures {
			infrastructurePath := filepath.Join(generatedPath, "internal", "infrastructure", infrastructure)
			createDir(infrastructurePath)
			createFileWithPackage(filepath.Join(infrastructurePath, infrastructure+".go"), infrastructure)
		}
	}

	// Создаем директории и файлы для application
	if len(applications) > 0 {
		for _, application := range applications {
			applicationPath := filepath.Join(generatedPath, "internal", "application", application)
			createDir(applicationPath)
			createFileWithPackage(filepath.Join(applicationPath, application+".go"), application)
		}
	}

	// Создаем директории и файлы для repositories
	if len(repositories) > 0 {
		for _, repository := range repositories {
			repositoryPath := filepath.Join(generatedPath, "internal", "infrastructure", "repositories", repository)
			createDir(repositoryPath)
			createFileWithPackage(filepath.Join(repositoryPath, repository+".go"), repository)
		}
	}

	// Создаем файл go.mod
	if strings.ToLower(goModInput) == "да" {
		createGoModFile(filepath.Join(generatedPath, "go.mod"), projectName, goVersion)
	}

	// Создаем файл .gitignore
	if strings.ToLower(gitignoreInput) == "да" {
		createGitignoreFile(filepath.Join(generatedPath, ".gitignore"))
	}

	// Создаем файл .env
	if strings.ToLower(envInput) == "да" {
		createFile(filepath.Join(generatedPath, ".env"))
		// Добавляем .env в .gitignore
		appendToGitignore(filepath.Join(generatedPath, ".gitignore"), ".env")
	}

	// Создаем Dockerfile
	if strings.ToLower(dockerfileInput) == "да" {
		createDockerfile(filepath.Join(generatedPath, "Dockerfile"))
	}

	// Создаем Makefile
	if strings.ToLower(makefileInput) == "да" {
		createMakefile(filepath.Join(generatedPath, "Makefile"))
	}

	// Создаем README.md
	if strings.ToLower(readmeInput) == "да" {
		createReadmeFile(filepath.Join(generatedPath, "README.md"), projectName)
	}
}

func createDir(dirPath string) {
	os.MkdirAll(dirPath, 0755)
}

func createFileWithPackage(filePath, packageName string) {
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Ошибка при создании файла:", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(fmt.Sprintf("package %s\n", packageName))
	if err != nil {
		fmt.Println("Ошибка при записи в файл:", err)
	}
}

func createFile(filePath string) {
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Ошибка при создании файла:", err)
		return
	}
	file.Close()
}

func createGoModFile(filePath, projectName, goVersion string) {
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Ошибка при создании файла:", err)
		return
	}
	defer file.Close()

	content := fmt.Sprintf(`module %s

go %s
`, projectName, goVersion)

	_, err = file.WriteString(content)
	if err != nil {
		fmt.Println("Ошибка при записи в файл:", err)
	}
}

func createMainFile(filePath, projectName string) {
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Ошибка при создании файла:", err)
		return
	}
	defer file.Close()

	content := fmt.Sprintf(`package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Добро пожаловать в проект %s!")
}
`, projectName)

	_, err = file.WriteString(content)
	if err != nil {
		fmt.Println("Ошибка при записи в файл:", err)
	}
}

func createGitignoreFile(filePath string) {
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Ошибка при создании файла:", err)
		return
	}
	defer file.Close()

	content := `# Бинарные файлы
*.exe
*.exe~
*.dll
*.so
*.dylib

# Пакеты и объектные файлы
*.a
*.o
*.so
*.dylib

# Тестовые файлы
*.out

# Локальные конфигурационные файлы
.env

# Временные файлы
*.tmp
*.swp

# Файлы, созданные IDE
*.iml
.idea/
.vscode/
`

	_, err = file.WriteString(content)
	if err != nil {
		fmt.Println("Ошибка при записи в файл:", err)
	}
}

func appendToGitignore(filePath, content string) {
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Ошибка при открытии файла:", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(content + "\n")
	if err != nil {
		fmt.Println("Ошибка при записи в файл:", err)
	}
}

func createDockerfile(filePath string) {
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Ошибка при создании файла:", err)
		return
	}
	defer file.Close()

	content := `# Используем официальный образ Golang для создания артефакта сборки.
FROM golang:1.18 as builder

# Устанавливаем текущую рабочую директорию внутри контейнера
WORKDIR /app

# Копируем файлы go.mod и go.sum
COPY go.mod go.sum ./

# Скачиваем все зависимости. Зависимости будут кэшированы, если файлы go.mod и go.sum не изменятся
RUN go mod download

# Копируем исходный код из текущей директории в рабочую директорию внутри контейнера
COPY . .

# Собираем приложение Go
RUN go build -o main .

# Начинаем новый этап с нуля
FROM alpine:latest

# Копируем предварительно собранный бинарный файл из предыдущего этапа в текущий этап
COPY --from=builder /app/main .

# Команда для запуска исполняемого файла
CMD ["./main"]
`

	_, err = file.WriteString(content)
	if err != nil {
		fmt.Println("Ошибка при записи в файл:", err)
	}
}

func createMakefile(filePath string) {
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Ошибка при создании файла:", err)
		return
	}
	defer file.Close()

	content := `# Makefile для проекта Go

# Цель по умолчанию
all: build

# Собираем проект
build:
	go build -o main .

# Запускаем проект
run: build
	./main

# Очищаем проект
clean:
	rm -f main

# Тестируем проект
test:
	go test ./...
`

	_, err = file.WriteString(content)
	if err != nil {
		fmt.Println("Ошибка при записи в файл:", err)
	}
}

func createReadmeFile(filePath, projectName string) {
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Ошибка при создании файла:", err)
		return
	}
	defer file.Close()

	content := fmt.Sprintf(`# %s

Добро пожаловать в проект %s!

## Начало работы

Эти инструкции помогут вам запустить копию проекта на вашем локальном компьютере для разработки и тестирования.

### Предварительные условия

Что нужно установить для работы программного обеспечения и как это сделать.

### Установка

Пошаговая серия примеров, которые помогут вам запустить среду разработки.

### Запуск тестов

Объясните, как запустить автоматические тесты для этой системы.

### Развертывание

Добавьте дополнительные заметки о том, как развернуть это на живой системе.
`, projectName, projectName)

	_, err = file.WriteString(content)
	if err != nil {
		fmt.Println("Ошибка при записи в файл:", err)
	}
}
