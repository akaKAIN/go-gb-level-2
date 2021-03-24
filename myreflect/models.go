package myreflect

const (
	Female = iota
	Male
)

type EmptyUser struct {}

type User struct {
	Age    uint8
	Gender uint8
	Email  string
	Name   string
}

