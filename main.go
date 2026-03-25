package main

import (
	"fmt"
)

type User struct {
	name string
	age  int
}

type Task struct {
	name string
	date string
}

func selection() int {
	var a int
	fmt.Println("\nПривет! Это мелкий CRM.")
	fmt.Println("1 - Создать сотрудника")
	fmt.Println("2 - Убрать сотрудника")
	fmt.Println("3 - Создать задачу")
	fmt.Println("4 - Удалить задачу")
	fmt.Println("0 - Выход")
	fmt.Print("Выбор: ")
	fmt.Scanln(&a)
	return a
}

func main() {
	for { // 🔥 ВОТ ЭТО ГЛАВНОЕ
		a := selection()

		if a == 1 {
			var name string
			var age int

			fmt.Println("Имя:")
			fmt.Scanln(&name)

			fmt.Println("Возраст:")
			fmt.Scanln(&age)

			user := User{name: name, age: age}
			fmt.Println("✅ Создан:", user.name, user.age)

		} else if a == 2 {
			fmt.Println("🗑 Удаление пока фейк 😄")

		} else if a == 3 {
			var name string
			var date string

			fmt.Println("Имя задачи:")
			fmt.Scanln(&name)

			fmt.Println("Дата:")
			fmt.Scanln(&date)

			task := Task{name: name, date: date}
			fmt.Println("✅ Создана задача:", task.name, task.date)

		} else if a == 4 {
			fmt.Println("🗑 Удаление задачи пока фейк 😄")

		} else if a == 0 {
			fmt.Println("Пока 👋")
			break // 🔥 выход из программы

		} else {
			fmt.Println("❌ Неверный ввод")
		}
	}
}