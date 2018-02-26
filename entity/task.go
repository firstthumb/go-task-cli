package entity

type Task struct {
	ID    int    `storm:"id,increment"`
	Value string `storm:"index"`
}
