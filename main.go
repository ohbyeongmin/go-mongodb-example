package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type InfoPerson struct {
	ID     primitive.ObjectID `bson:"_id"`
	Name   string             `bson:"name,omitempty"`
	Age    int                `bson:"age,omitempty"`
	Gender string             `bson:"gender,omitempty"`
}

type myStructBson struct {
	MyStruct InfoPerson `bson:",inine"`
}

func (m *InfoPerson) MarshalBSON() ([]byte, error) {
	return bson.Marshal(myStructBson{
		MyStruct: *m,
	})
}

func mongoConn() (client *mongo.Client) {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	return client
}

func addPersonInfo(m *mongo.Collection, info InfoPerson) (*mongo.InsertOneResult, error) {
	insertResult, err := m.InsertOne(context.Background(), info)
	return insertResult, err
}

func main() {
	conn := mongoConn()
	mongo := conn.Database("data_person").Collection("info")

	person := InfoPerson{
		ID:     primitive.NewObjectID(),
		Name:   "sol4",
		Age:    2412,
		Gender: "female",
	}

	insertResult, err := addPersonInfo(mongo, person)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(insertResult.InsertedID)

	var infos []InfoPerson
	cursor, err := mongo.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	if err = cursor.All(context.Background(), &infos); err != nil {
		log.Fatal(err)
	}
	fmt.Println(infos)
}
