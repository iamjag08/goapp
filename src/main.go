package main

import (
	"context"
	"fmt"
	"log"

	"github.com/astaxie/beego"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	if err != nil {
		log.Fatal(err)
	}
	beego.Router("states/:name/:capital", &mainController{})
	beego.Run()
}

type mainController struct {
	beego.Controller
}

func insertData(name string, capital string) {
	clientOptions := options.Client().ApplyURI("mongodb://mongodb:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	collection := client.Database("test").Collection("locations")
	loc := Location{name: name, capital: capital}
	locs := []interface{}{loc}
	insertManyResult, err := collection.InsertMany(context.TODO(), locs)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted documents: ", insertManyResult.InsertedIDs)
}


func fetchData() []interface {
	clientOptions := options.Client().ApplyURI("mongodb://mongodb:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	collection := client.Database("test").Collection("locations")
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
    	log.Fatal(err)
	}
	defer cursor.Close(ctx)
	locs := []interface{}
	for cursor.Next(ctx) {
    	var episode bson.M
    	if err = cursor.Decode(&episode); err != nil {
        	log.Fatal(err)
    	}
    	locs = append(locs, episode)
	}
	fmt.Println("Fetched documents: ", locs)
	return locs
}
func (c *mainController) Get() {

	//Obtain the values of the route parameters defined in the route above
	state := c.Ctx.Input.Param(":name")
	capital := c.Ctx.Input.Param(":capital")
	insertData(state, capital)
	c.Data["msg"] = "Added location!!"
	c.Data["result"] = fetchData()
}
