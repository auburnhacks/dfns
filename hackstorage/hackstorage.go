package hackstorage

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
)

// Storage is an interface that abstracts all the things that the
// hackstorage package can perform
type Storage interface {
	Ping() error
	LinkUser(userID string, linkID int) error
	AddAttendee(eventID string, userID int) error
	CanAttend(eventID string, userID int) (bool, error)
}

var (
	// ErrCannotAttend is an error that says when a hacker cannot attend an
	// event because they've already been scanned
	ErrCannotAttend = errors.New("error: hacker cannot attend event")
)

// New returns a type that implement the Storage interface
func New(url, collection string) (Storage, error) {
	store, err := newMongoStore(url)
	if err != nil {
		return nil, err
	}
	return store, nil
}

type mongoStore struct {
	client *mongo.Client
}

func newMongoStore(url string) (*mongoStore, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, url)
	if err != nil {
		return nil, err
	}
	return &mongoStore{client: client}, nil
}

func (st *mongoStore) Ping() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := st.client.Ping(ctx, nil)
	if err != nil {
		return err
	}
	return nil
}

func (st *mongoStore) LinkUser(userID string, randNum int) error {
	// Create a user and save it to the database
	collection := st.client.Database("dayof").Collection("hackers")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err := collection.InsertOne(ctx,
		bson.M{"uid": userID, "rand_num": randNum})
	if err != nil {
		return err
	}
	return nil
}

func (st *mongoStore) AddAttendee(eventID string, userID int) error {
	att := bson.M{"event_id": eventID, "user_id": userID}
	collection := st.client.Database("dayof").Collection("events")
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	_, err := collection.InsertOne(ctx, att)
	if err != nil {
		return err
	}
	return nil
}

func (st *mongoStore) CanAttend(eventID string, userID int) (bool, error) {
	filter := bson.M{"event_id": eventID, "user_id": userID}
	collection := st.client.Database("dayof").Collection("events")
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	res := collection.FindOne(ctx, filter)
	// no records found so hacker can attend event
	var result bson.M
	if err := res.Decode(&result); err != nil {
		if err == mongo.ErrNoDocuments {
			return true, nil
		}
		log.Printf("error while decoding from result: %v", err)
		return false, err
	}
	return false, ErrCannotAttend
}
