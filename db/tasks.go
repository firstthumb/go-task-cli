package db

import (
	"github.com/asdine/storm"
	"github.com/asdine/storm/codec/json"
	"fmt"
)

type StormDB struct {
	c *storm.DB
}

type Task struct {
	ID    int    `storm:"id,increment"`
	Value string `storm:"index"`
}

func Init(dbPath string) (*StormDB, error) {
	var err error
	db, err := storm.Open(dbPath, storm.Codec(json.Codec))
	if err != nil {
		fmt.Println("Could not open database", err.Error())
		return nil, err
	}
	s := &StormDB{c:db}
	return s, nil
}

func (db *StormDB) CreateTask(task string) (int, error) {
	defer db.Close()

	entity := Task{Value: task}
	err := db.c.Save(&entity)
	if err != nil {
		fmt.Print("Error : ", err.Error())
		return -1, err
	}
	return entity.ID, nil
}

func (db *StormDB) AllTasks() ([]Task, error) {
	defer db.Close()

	var tasks []Task
	err := db.c.All(&tasks)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func (db *StormDB) Close() {
	db.c.Close()
}

