package test

import (
	"../../drive/mongo"
	"gopkg.in/mgo.v2/bson"
	"fmt"
	"sync"
)

var testDb *TestDB
var once sync.Once

type Test struct {
	Name string
	Version int8
}

type TestDB struct {
	collection string
}

func NewTest() *TestDB {
	once.Do(func() {
		testDb = &TestDB{
			collection : "test",
		}
	})
	return testDb
}

func FindAll() {
	mon := mongo.NewMongo()
	var numbers []Test
	mon.Collection("test").Find(bson.M{}).All(&numbers)
	fmt.Println(numbers)
}
