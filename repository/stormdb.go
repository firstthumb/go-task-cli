package repository

import (
	"TaskApp/entity"
	"github.com/asdine/storm"
	"fmt"
	"github.com/asdine/storm/codec/json"
)

type StormDB struct {
	c *storm.DB
}

func newStormDB(dbPath string) Repository {
	db, _ := Init(dbPath)
	return db
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

func (db *StormDB) FindAll() ([]*entity.Task, error) {
	defer db.Close()

	var tasks []*entity.Task
	err := db.c.All(&tasks)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func (db *StormDB) Store(task *entity.Task) (int, error) {
	defer db.Close()

	err := db.c.Save(task)
	if err != nil {
		fmt.Print("Error : ", err.Error())
		return -1, err
	}
	return task.ID, nil
}

func (db *StormDB) Close() {
	db.c.Close()
}
