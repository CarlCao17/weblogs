package mongo

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type DB struct {
	c      *mongo.Client
	dbName string
}

func NewMongoDB(client *mongo.Client, dbName string) *DB {
	return &DB{
		c:      client,
		dbName: dbName,
	}
}

func (d *DB) Raw() *mongo.Client {
	return d.c
}

func (d *DB) DBName() string {
	return d.dbName
}

func (d *DB) Database() *mongo.Database {
	return d.c.Database(d.dbName)
}

func (d *DB) Collection(collectionName string) *mongo.Collection {
	return d.Database().Collection(collectionName)
}
