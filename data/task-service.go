package data

import (
	"Task-Manager/models"
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const mongoURI = "mongodb://localhost:27017"

var (
	// client     *mongo.Client
	taskCol    *mongo.Collection
	dbName     = "taskdb"
	colName    = "tasks"
	ctxTimeout = 10 * time.Second * 10
)

func InitDB() {
	ctx, cancel := context.WithTimeout(context.Background(), ctxTimeout)
	defer cancel()
	clientOption := options.Client().ApplyURI(mongoURI)
	client, err := mongo.Connect(ctx, clientOption)
	if err != nil {
		log.Fatal(err)
		return
	}

	// Test the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
		return
	}

	taskCol = client.Database(dbName).Collection(colName)
	fmt.Println("Connected to Mongodb and ready")
}

// create task

func CreateTask(task models.Task) error {
	ctx, cancel := context.WithTimeout(context.Background(), ctxTimeout)
	defer cancel()
	_, err := taskCol.InsertOne(ctx, task)
	return err
}

func GetAllTasks() ([]models.Task, error) {
	ctx, cancel := context.WithTimeout(context.Background(), ctxTimeout)
	defer cancel()
	cursor, err := taskCol.Find(ctx, bson.M{})
	if err != nil {
		return nil, errors.New("There is no task")
	}
	defer cursor.Close(ctx)
	var tasks []models.Task
	for cursor.Next(ctx) {
		var task models.Task
		err = cursor.Decode(&task)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func GetTaskById(id string) (*models.Task, error) {
	ctx, cancel := context.WithTimeout(context.Background(), ctxTimeout)
	defer cancel()
	var task models.Task
	err := taskCol.FindOne(ctx, bson.M{"_id": id}).Decode(&task)
	if err != nil {
		return nil, errors.New("Task not Found")
	}
	return &task, nil
}

func UpdateTask(id string, updated models.Task) error {
	ctx, cancel := context.WithTimeout(context.Background(), ctxTimeout)
	defer cancel()
	update := bson.M{
		"$set": bson.M{
			"name":        updated.Name,
			"description": updated.Description,
			"status":      updated.Status,
			"due_date":    updated.DueDate,
		},
	}
	result, err := taskCol.UpdateOne(ctx, bson.M{"_id": id}, update)
	if err != nil {
		return err
	}
	if result.MatchedCount == 0 {
		return errors.New("No task found with the given ID")
	}
	return nil
}

func DeleteTask(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), ctxTimeout)
	defer cancel()
	result, err := taskCol.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return errors.New("Nothing matched with the given id")
	}
	return nil
}
