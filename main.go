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

// var lr = &model.LogRecord{
// 	JobName: "撿破爛1",
// 	Command: "推車撿破爛1",
// 	Err:     "撿不到破爛1",
// 	Content: "剪了好多破爛1",
// }

func main() {
	lr := new(model.LogRecord)
	lr.JobName = "撿破爛5566"
	lr.Command = "推車撿破爛5566"
	lr.Err = "撿不到破爛5566"
	lr.Content = "剪了好多破爛5566"

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

	// GET
	cur, err := collection.Find(ctx, bson.D{{"jobName", "撿破爛5566"}})
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(ctx)
	var results []model.LogRecord
	// 使用All函数获取所有查询结果，并将结果保存至results变量。
	if err = cur.All(context.TODO(), &results); err != nil {
		log.Fatal(err)
	}
	jsondata, _ := json.Marshal(results)
	fmt.Println(string(jsondata))
	// fmt.Println(results)
	// 遍历结果数组
	// for _, result := range results {
	// 	fmt.Println(*result)
	// }

	// for cur.Next(ctx) {
	// 	var document bson.M
	// 	err = cur.Decode(&document)
	// 	if err != nil {
	// 		log.Println(err)
	// 	}
	// 	results = append(results, document)
	// }
	// fmt.Println(results)
}
