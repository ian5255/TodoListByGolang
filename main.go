package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"TodoListByGolang/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	lr := new(model.LogRecord)
	lr.JobName = "撿破爛5566"
	lr.Command = "推車撿破爛5566"
	lr.Err = "撿不到破爛5566"
	lr.Content = "剪了好多破爛5566"

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Connect
	client, _ := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	collection := client.Database("first-project").Collection("todoList")
	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Insert
	collection.InsertOne(context.TODO(), lr)

	// GET
	cur, err := collection.Find(ctx, bson.D{{"jobName", "撿破爛5566"}})
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(ctx)
	var results []model.LogRecord
	if err = cur.All(context.TODO(), &results); err != nil {
		log.Fatal(err)
	}
	jsondata, _ := json.Marshal(results)
	fmt.Println(string(jsondata))
}
