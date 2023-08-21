package database

import (
	"TaskAPP/models"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client     *mongo.Client
	db         *mongo.Database
	collection *mongo.Collection
)

func InitDatabase() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	db = client.Database("dbgolang")
	collection = db.Collection("tasks")
}

func GetTasks() ([]models.Task, error) {
	var tasks []models.Task

	cursor, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var task models.Task
		err := cursor.Decode(&task)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}
func GetTaskByID(taskID primitive.ObjectID) (models.Task, error) {
	var task models.Task

	filter := bson.M{"_id": taskID}
	err := collection.FindOne(context.Background(), filter).Decode(&task)
	if err != nil {
		return models.Task{}, err
	}

	return task, nil
}

func InsertTask(task models.Task) (primitive.ObjectID, error) {
	task.Subtasks = []models.Subtask{}       // Inisialisasi subtasks sebagai array kosong
	task.Attachments = []models.Attachment{} // Inisialisasi attachments sebagai array kosong

	result, err := collection.InsertOne(context.Background(), task)
	if err != nil {
		return primitive.NilObjectID, err
	}

	insertedID, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return primitive.NilObjectID, err
	}

	return insertedID, nil
}

func UpdateTask(updatedTask models.Task) error {
	filter := bson.D{{"_id", updatedTask.ID}}
	update := bson.D{{"$set", bson.D{
		{"title", updatedTask.Title},
		{"description", updatedTask.Description},
	}}}

	_, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}

	return nil
}

func DeleteTask(taskID primitive.ObjectID) error {
	filter := bson.D{{"_id", taskID}}

	_, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}

	return nil
}

func AddSubtaskToTask(taskID primitive.ObjectID, subtask models.Subtask) error {
	filter := bson.M{"_id": taskID}
	update := bson.M{"$push": bson.M{"subtasks": subtask}}

	_, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Printf("Error menambah subtask: %v\n", err)
		return err
	}
	return nil
}

func UpdateSubtaskInTask(taskID, subtaskID primitive.ObjectID, updatedSubtask models.Subtask) error {
	filter := bson.M{"_id": taskID, "subtasks._id": subtaskID}
	update := bson.M{"$set": bson.M{"subtasks.$.title": updatedSubtask.Title, "subtasks.$.done": updatedSubtask.Done}}

	_, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Printf("Error updating subtask: %v\n", err)
		return err
	}

	return nil
}

func DeleteSubtask(subtaskID primitive.ObjectID) error {
	_, err := collection.UpdateOne(context.Background(), bson.M{"subtasks._id": subtaskID},
		bson.M{"$pull": bson.M{"subtasks": bson.M{"_id": subtaskID}}})
	if err != nil {
		log.Printf("Error deleting subtask: %v\n", err)
		return err
	}

	return nil
}
func AddAttachmentToTask(taskID primitive.ObjectID, attachment models.Attachment) error {
	filter := bson.M{"_id": taskID}
	update := bson.M{"$push": bson.M{"attachments": attachment}}

	_, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Printf("Error adding attachment: %v\n", err)
		return err
	}

	return nil
}

func UpdateAttachmentInTask(taskID, attachmentID primitive.ObjectID, updatedAttachment models.Attachment) error {
	filter := bson.M{"_id": taskID, "attachments._id": attachmentID}
	update := bson.M{"$set": bson.M{"attachments.$": updatedAttachment}}

	_, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Printf("Error updating attachment: %v\n", err)
		return err
	}

	return nil
}

func DeleteAttachmentFromTask(taskID, attachmentID primitive.ObjectID) error {
	filter := bson.M{"_id": taskID}
	update := bson.M{"$pull": bson.M{"attachments": bson.M{"_id": attachmentID}}}

	_, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Printf("Error deleting attachment: %v\n", err)
		return err
	}

	return nil
}
