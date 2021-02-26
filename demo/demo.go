package demo

import "fmt"

// Модель пользователя
type User struct {
	Name string // Имя пользователя
	Age  int    // Возраст пользователя
}

// Метод вывода строкового представления модели пользователя
func (u User) String() string {
	return fmt.Sprintf("user: %s, %d years old", u.Name, u.Age)
}
