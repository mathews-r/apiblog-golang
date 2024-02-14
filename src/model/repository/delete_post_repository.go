package repository

import (
	"context"
	"os"

	"github.com/mathews-r/golang/src/configs/logger"
	"github.com/mathews-r/golang/src/configs/rest_err"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (pr *postRepository) DeletePost(postId string) *rest_err.RestErr {
	logger.Info("Init DeletePost repository")

	collectionName := os.Getenv(DB_POST_COLLECTION)

	collection := pr.databaseConnection.Collection(collectionName)

	postIdHex, _ := primitive.ObjectIDFromHex(postId)

	filter := bson.D{{Key: "_id", Value: postIdHex}}

	_, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return rest_err.NewInternalServerErr(err.Error())
	}

	return nil
}