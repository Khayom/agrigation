package main

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// InsertToDB()
	Average()
}

func Average() {
	clientOptions := options.Client().ApplyURI("mongodb://127.0.0.1:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	connect := client.Database("agrigation-hw").Collection("item")

	Cursor, _ := connect.Aggregate(context.TODO(),bson.A{
		bson.M{
			"$group": bson.M{
				"_id": "$store",
				"average_price": bson.M{
					"$avg":"$price",
				},
				"type_total": bson.M{"$sum": 1},
			},
		},
	})
	
	var Slice = []bson.M{}

	for Cursor.Next(context.TODO()) {
		var Shablon bson.M
		Cursor.Decode(&Shablon)

		Slice = append(Slice, Shablon)
	}
	fmt.Printf("Slice: %v\n", Slice)
}

func InsertToDB() {

	clientOptions := options.Client().ApplyURI("mongodb://127.0.0.1:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	connect := client.Database("agrigation-hw").Collection("item")

	connect.InsertMany(context.TODO(), bson.A{
		bson.M{ "_id": 1, "order_id": "ORD1001", "customer_id": "C001", "item": "Laptop", "quantity": 1, "price": 1000 },
		bson.M{ "_id": 2, "order_id": "ORD1002", "customer_id": "C002", "item": "Mouse", "quantity": 2, "price": 50 },
		bson.M{ "_id": 3, "order_id": "ORD1003", "customer_id": "C001", "item": "Keyboard", "quantity": 1, "price": 100 },
		bson.M{ "_id": 4, "order_id": "ORD1004", "customer_id": "C003", "item": "Monitor", "quantity": 1, "price": 300 },
	})
}
