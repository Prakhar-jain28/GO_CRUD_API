package datastore

import (
	"os"
	"strconv"

	"github.com/Prakhar-jain28/GO_CRUD_API/model"

	"context"
	"fmt"
	"log"

	"gofr.dev/pkg/errors"
	"gofr.dev/pkg/gofr"

	"go.mongodb.org/mongo-driver/bson"
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

func (s *Blog) GetByID(ctx *gofr.Context, ID string) (*model.Blog, error) {
	coll := s.connectMongoDB(ctx)
	var result model.Blog
	i, errr := strconv.Atoi(ID)
	if errr != nil {
		return nil, errr
	}
	fmt.Println("ID:", i)

	err := coll.FindOne(context.Background(), bson.M{"id": i}).Decode(&result)
	if err != nil {
		fmt.Println("FindOne ERROR:", err)
		return nil, err
	}

	fmt.Println(coll.FindOne(context.Background(), bson.M{"id": i}))
	return &result, nil
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

func (s *Blog) Delete(ctx *gofr.Context, id int) error {
	coll := s.connectMongoDB(ctx)
	
	filter := bson.M{"id": id}
	result, err := coll.DeleteMany(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Number of documents deleted: %d\n", result.DeletedCount)
	return nil
}

func (s *Blog) Update(ctx *gofr.Context, blog *model.Blog) (*model.Blog, error) {
	coll := s.connectMongoDB(ctx)

	existingBlog, err := s.GetByID(ctx, strconv.Itoa(blog.ID))
	if err != nil {
		fmt.Println("Error fetching existing Blog data:", err)
		return nil, err
	}

	updatedBlog := merge(existingBlog, blog)

	filter := bson.M{"id": blog.ID}
	update := bson.M{"$set": updatedBlog}

	_, err = coll.UpdateOne(context.Background(), filter, update)
	if err != nil {
		fmt.Println("Failed to update student")
		return nil, errors.DB{Err: err}
	}

	fmt.Println("Blog updated successfully", blog.ID)
	return updatedBlog, nil
}

func merge(existing *model.Blog, update *model.Blog) *model.Blog {
	if update.Name != "" {
		existing.Name = update.Name
	}
	if update.Title != "" {
		existing.Title = update.Title
	}
	if update.Content != "" {
		existing.Content = update.Content
	}
	return existing
}
