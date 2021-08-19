package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"reflect"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func regRoutes() *gin.Engine {

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	fmt.Println("ClientOptopm TYPE:", reflect.TypeOf(clientOptions), "\n")

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		fmt.Println("Mongo.connect() ERROR: ", err)
		os.Exit(1)
	}

	r := gin.Default()
	r.LoadHTMLGlob("D:/goworkspace/src/response/templates/*.html")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	sample := r.Group("/user")

	sample.GET("/adduser", func(c *gin.Context) {
		c.HTML(http.StatusOK, "user-form.html", nil)
	})

	sample.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})

	sample.GET("/home", func(c *gin.Context) {
		c.HTML(http.StatusOK, "html-sample2.html", nil)
	})

	sample.POST("/", func(c *gin.Context) {
		var user User

		c.Bind(&user)

		user.ID = 3

		users["3"] = user
		c.IndentedJSON(http.StatusOK, user)

		ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)

		col := client.Database("ajinkya_db").Collection("goTraining")

		result, _ := col.InsertOne(ctx, user)

		fmt.Println("result ", result)

		newID := result.InsertedID
		fmt.Println("InsertedOne(), newID", result)
		fmt.Println("InsertedOne(), newID type:", reflect.TypeOf(newID))

	})

	sample.POST("/home", func(c *gin.Context) {
		var cred Creds

		c.Bind(&cred)

		c.IndentedJSON(http.StatusOK, cred)

		// ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)

		// col := client.Database("ajinkya_db").Collection("goTraining")

		// result, _ := col.InsertOne(ctx, user)

		// fmt.Println("result ", result)

		// newID := result.InsertedID
		fmt.Println("Cred", cred)
		// fmt.Println("InsertedOne(), newID type:", reflect.TypeOf(newID))

	})

	r.Static("/public", "D:goworkspace/src/response/public")

	return r
}
