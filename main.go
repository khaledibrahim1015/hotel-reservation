package main

import (
	"flag"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/khaledibrahim1015/hotel-reservation/api"
	"github.com/khaledibrahim1015/hotel-reservation/config"
	"github.com/khaledibrahim1015/hotel-reservation/db"
)

const (
	dbName   = "hotel-reservation"
	userColl = "user"
)

func main() {

	// // test connect to mongodb
	// mongoclient, err := config.ConnectToMongoDB()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("Connected to MongoDB!")

	// coll := mongoclient.Database(dbName).Collection(userColl)

	// //  create user
	// var user types.User = types.User{
	// 	FirstName: "jooooo",
	// 	LastName:  "omaaaaar",
	// }
	// res, err := coll.InsertOne(context.Background(), user)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// insertedId := res.InsertedID
	// fmt.Println(insertedId)

	// // Get all users
	// cursor, err := coll.Find(context.Background(), bson.D{})
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// var users []types.User
	// for cursor.Next(context.Background()) {
	// 	var user types.User
	// 	if err := cursor.Decode(&user); err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	users = append(users, user)

	// }
	// fmt.Println("users data :", users)

	// // Get a user by ID
	// var usr types.User

	// var id string = "6738a45b5f03aebb66e1b9e9"
	// objectid, err := primitive.ObjectIDFromHex(id)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// filter := bson.M{"_id": objectid}
	// err = coll.FindOne(context.Background(), filter).Decode(&usr)

	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(usr)
	// fmt.Println("%T", usr.ID)

	// //  or
	// var result bson.M
	// err = coll.FindOne(context.Background(), filter).Decode(&result)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(result)

	// // update user
	// id = "6738a3fc70a287b9053ac1ee"
	// objID, err := primitive.ObjectIDFromHex(id)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// filter = bson.M{"_id": objID}
	// var updatedUser types.User = types.User{
	// 	FirstName: "khaled ibrahim",
	// 	LastName:  "ahmed ali ",
	// }
	// update := bson.M{"$set": updatedUser}
	// updatedres, err := coll.UpdateOne(context.Background(), filter, update)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("updated result ", updatedres)
	// fmt.Println(updatedres.UpsertedID)

	// command line
	var listenAddr *string
	listenAddr = flag.String("listenAddr", ":5000", "the listen address of api server ")
	flag.Parse()

	// Handlers Intialization
	mongoclient, err := config.ConnectToMongoDB()
	if err != nil {
		log.Fatal(err)
	}
	userHandler := api.NewUserHandler(db.NewMongoUserStore(mongoclient))

	app := fiber.New()
	apiv1 := app.Group("/api/v1")
	apiv1.Get("/user", userHandler.HandleGetUsers)
	apiv1.Get("/user/:id", userHandler.HandleGetUser)
	apiv1.Post("/user", userHandler.HandlePostUser)
	apiv1.Put("user/:id", userHandler.HandlePutUser)
	app.Listen(*listenAddr)
}
