package datastore

import (
	"GO-LANG/model"
	"os"

	"context"
	"fmt"
	"log"

	"gofr.dev/pkg/gofr"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Blog struct {
	client *mongo.Client
}

func New() *Blog {
	return &Blog{}
}

func (s *Blog) connectMongoDB(ctx *gofr.Context) *mongo.Collection {
	if s.client == nil {
		uri := "mongodb+srv://prakharjain496:golangdev@golangdb.v2pxtfw.mongodb.net/?retryWrites=true&w=majority"
		if uri == "" {
			log.Fatal("You must set your 'MONGODB_URI' environment variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
		}
		client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
		if err != nil {
			panic(err)
		}
		s.client = client
	}
	fmt.Println("Connected to MongoDB")
	return s.client.Database("sample_mflix").Collection("Blogs")
}

func (s *Blog) Create(ctx *gofr.Context, blog *model.Blog) (*model.Blog, error) {

	coll := s.connectMongoDB(ctx)

	data := model.Blog{ID: blog.ID, Name: blog.Name, Title: blog.Title, Content: blog.Content}
	result, insertErr := coll.InsertOne(ctx, data)
	if insertErr != nil {
		fmt.Println("InsertOne ERROR:", insertErr)
		os.Exit(1)
	} else {
		fmt.Println("Data inserted with objectID: ", result.InsertedID)
	}

	return &data, nil
}