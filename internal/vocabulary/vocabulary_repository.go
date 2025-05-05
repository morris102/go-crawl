package vocabulary

import (
	"context"
	"log"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type VocalbularyRepository interface {
	CreateMany(items []Word) error
}

type VocalbularyRepositoryImpl struct {
	mongoClient *mongo.Client
	tableName   string
}

func NewVocalbularyRepository(mongoClient *mongo.Client) (VocalbularyRepository, error) {
	return &VocalbularyRepositoryImpl{
		mongoClient: mongoClient,
		tableName:   "vocabulary",
	}, nil
}

func (inst *VocalbularyRepositoryImpl) CreateMany(items []Word) error {
	var db = inst.mongoClient.Database("crawl-japanese")
	ctx := context.Background()

	var operations []mongo.WriteModel
	for _, item := range items {
		filter := bson.M{"_id": uuid.New()}
		update := bson.M{"$set": item}
		upsert := true
		model := mongo.NewUpdateOneModel().
			SetFilter(filter).
			SetUpdate(update).
			SetUpsert(upsert)
		operations = append(operations, model)
	}

	bulkOpts := options.BulkWrite().SetOrdered(false)
	_, err := db.Collection("vocalbulary").BulkWrite(ctx, operations, bulkOpts)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}
