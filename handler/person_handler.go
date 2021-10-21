package handler

import (
	"context"
	"crud/model"
	"encoding/json"

	"crud/db"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreatePerson(c *fiber.Ctx) error {
	collection := db.PersonCollection()

	var person model.Person
	json.Unmarshal([]byte(c.Body()), &person)

	res, err := collection.InsertOne(context.Background(), person)
	if err != nil {
		c.Status(500)
		return err
	}

	response, _ := json.Marshal(res)
	return c.Send(response)
}

func GetPerson(c *fiber.Ctx) error {
	collection := db.PersonCollection()

	var filter bson.M = bson.M{}

	if c.Params("id") != "" {
		id := c.Params("id")
		objID, _ := primitive.ObjectIDFromHex(id)
		filter = bson.M{"_id": objID}
	}

	var results []bson.M
	cur, err := collection.Find(context.Background(), filter)
	defer cur.Close(context.Background())

	if err != nil {
		c.Status(500)
		return err
	}

	cur.All(context.Background(), &results)

	if results == nil {
		c.SendStatus(404)
		return err
	}

	json, _ := json.Marshal(results)
	return c.Send(json)
}

func GetAllPerson(c *fiber.Ctx) error {
	collection := db.PersonCollection()

	var filter bson.M = bson.M{}

	var results []bson.M
	cur, err := collection.Find(context.Background(), filter)
	defer cur.Close(context.Background())

	if err != nil {
		c.Status(500)
		return err
	}

	cur.All(context.Background(), &results)

	if results == nil {
		c.SendStatus(404)
		return err
	}

	json, _ := json.Marshal(results)
	return c.Send(json)
}

func UpdatePerson(c *fiber.Ctx) error {
	collection := db.PersonCollection()

	var person model.Person
	json.Unmarshal([]byte(c.Body()), &person)

	update := bson.M{
		"$set": person,
	}

	objID, _ := primitive.ObjectIDFromHex(c.Params("id"))
	res, err := collection.UpdateOne(context.Background(), bson.M{"_id": objID}, update)

	if err != nil {
		c.Status(500)
		return err
	}

	response, _ := json.Marshal(res)
	return c.Send(response)
}

func DeletePerson(c *fiber.Ctx) error {
	collection := db.PersonCollection()

	objID, _ := primitive.ObjectIDFromHex(c.Params("id"))
	res, err := collection.DeleteOne(context.Background(), bson.M{"_id": objID})

	if err != nil {
		c.Status(500)
		return err
	}

	jsonResponse, _ := json.Marshal(res)
	return c.Send(jsonResponse)
}
