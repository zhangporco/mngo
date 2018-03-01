package mongo

import (
	"gopkg.in/mgo.v2"
	"../../config"
	"sync"
)

var mongo *Mongo
var once sync.Once

type Mongo struct {
	db *mgo.Database
	collection *mgo.Collection
}

func NewMongo() *Mongo {
	once.Do(func() {
		mongo = &Mongo{}
		mongo.setDB()
	})
	return mongo
}

func (mongo *Mongo) setDB() {
	con := config.NewConfig()
	session, err := mgo.Dial(con.Mongo.Url)
	if err != nil {
		panic(err)
	}
	//defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	db := session.DB(con.Mongo.DbName)
	mongo.db = db
}

func (mongo *Mongo) Collection(name string) *Mongo {
	mongo.collection = mongo.db.C(name)
	return mongo
}

func (mongo *Mongo) Find(query interface{}) *mgo.Query {
	return mongo.collection.Find(query)
}
