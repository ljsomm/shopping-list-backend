package datasource

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDatasource struct {
	connection *mongo.Database
}

func NewMongoConnection() (MongoDatasource, error) {
	datasource := MongoDatasource{}
	connection, err := mongo.Connect(context.TODO(), options.Client().SetAuth(options.Credential{Username: "root", Password: "root"}).ApplyURI("mongodb://localhost:27017"))
	datasource.connection = connection.Database("shopping_list")
	return datasource, err
}

func (mongo MongoDatasource) GetRepository(collection string) *mongo.Collection {
	return mongo.connection.Collection(collection)
}
