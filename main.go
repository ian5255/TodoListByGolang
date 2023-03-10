package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"TodoListByGolang/model"

	"net/http"

	"github.com/gin-gonic/gin"
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

	// InsertMany
	var datas []interface{}
	for i := 0; i < 5; i++ {
		el := new(model.LogRecord)
		el.JobName = fmt.Sprint("撿破爛", i)
		el.Command = fmt.Sprint("推車撿破爛", i)
		el.Err = fmt.Sprint("撿不到破爛", i)
		el.Content = fmt.Sprint("剪了好多破爛", i)
		datas = append(datas, el)
	}
	_, err := collection.InsertMany(context.TODO(), datas)
	if err != nil {
		log.Fatal(err)
	}

	// GET
	findOptions := options.Find()
	findOptions.SetLimit(3)
	cur, err := collection.Find(ctx, bson.D{{Key: "jobName", Value: "撿破爛5566"}}, findOptions)
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

	// UpdateOne
	filter := bson.D{{Key: "jobName", Value: "撿破爛8"}}
	opts := options.Update().SetUpsert(true) // 如果資料不存在則新增
	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "command", Value: "推車撿破爛7777888800001"}},
		}}
	result, err := collection.UpdateOne(context.TODO(), filter, update, opts)
	if err != nil {
		log.Fatal(err)
	}
	if result.MatchedCount != 0 {
		fmt.Printf("Matched %v documents and updated %v documents.\n", result.MatchedCount, result.ModifiedCount)
	}
	if result.UpsertedCount != 0 {
		fmt.Printf("inserted a new document with ID %v\n", result.UpsertedID)
	}

	// UpdateMany
	result2, err2 := collection.UpdateMany(context.TODO(), filter, update, opts)
	if err2 != nil {
		log.Fatal(err)
	}
	if result2.MatchedCount != 0 {
		fmt.Printf("Matched %v documents and updated %v documents.\n", result2.MatchedCount, result2.ModifiedCount)
	}
	if result2.UpsertedCount != 0 {
		fmt.Printf("inserted a new document with ID %v\n", result2.UpsertedID)
	}

	// DeleteMany
	filterD := bson.D{{Key: "jobName", Value: "撿破爛1"}}
	deleteResult, err := collection.DeleteMany(context.TODO(), filterD)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Deleted %v documents in the trainers collection\n", deleteResult.DeletedCount)

	// =================================================================================================================
	// Gin
	r := gin.Default()
	r.GET("/todolist", func(c *gin.Context) {
		c.JSON(http.StatusOK, results)
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
