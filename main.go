package main

import (
	"context"
	"fmt"
	"time"

	"TodoListByGolang/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var lr = &model.LogRecord{
	JobName: "撿破爛1",
	Command: "推車撿破爛1",
	Err:     "撿不到破爛1",
	Content: "剪了好多破爛1",
}

func main() {
	// lr := new(model.LogRecord)
	// lr.JobName = "撿破爛"
	// lr.Command = "推車撿破爛"
	// lr.Err = "撿不到破爛"
	// lr.Content = "剪了好多破爛"

	// lr := &model.LogRecord{
	// 	JobName: "撿破爛",
	// 	Command: "推車撿破爛",
	// 	Err:     "撿不到破爛",
	// 	Content: "剪了好多破爛",
	// }
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, _ := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	collection := client.Database("first-project").Collection("todoList")
	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	collection.InsertOne(context.TODO(), lr)
	res, _ := collection.InsertOne(ctx, bson.D{{"name", "pi"}, {"value", 3.14159}})
	id := res.InsertedID
	fmt.Println(id)
}
