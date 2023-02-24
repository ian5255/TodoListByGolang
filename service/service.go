package service

import (
	"log"

	"TodoListByGolang/model"
	db "TodoListByGolang/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetList() (list []model.LogRecord, err error) {
	list = make([]model.LogRecord, 0)
	ctx, collection, err := db.ConnectDB()
	if err != nil {
		return list, err
	}

	// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// defer cancel()

	// // Connect
	// client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	// collection := client.Database("firstProject").Collection("todoList")

	// GET
	findOptions := options.Find()
	findOptions.SetLimit(3)
	cur, err := collection.Find(ctx, bson.D{{Key: "jobName", Value: "撿破爛5566"}}, findOptions)
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(ctx)
	var results []model.LogRecord
	if err = cur.All(ctx, &results); err != nil {
		log.Fatal(err)
	}
	// jsondata, _ := json.Marshal(results)
	// fmt.Println(string(jsondata))
	return results, nil
}
