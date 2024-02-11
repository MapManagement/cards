package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.43

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"log"

	"github.com/FachschaftMathPhysInfo/cards/server/graph/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// CreateDeck is the resolver for the createDeck field.
func (r *mutationResolver) CreateDeck(ctx context.Context, input model.NewDeck) (*model.Deck, error) {
	cardDecks := r.DB.Collection("cardDecks")

	if input.Semester != nil && !(*input.Semester == "SoSe" || *input.Semester == "WiSe") {
		return nil, fmt.Errorf("Input \"%s\" is not a valid input for field input.Semester", *input.Semester)
	}

	// hash and copy to the buffer at the same time
	fileBuf := &bytes.Buffer{}
	tee := io.TeeReader(input.File.File, fileBuf)

	// generate the hash of the input file
	fileHash := sha256.New()
	if _, err := io.Copy(fileHash, tee); err != nil {
		log.Fatal(err)
	}
	encodedHash := hex.EncodeToString(fileHash.Sum(nil))

	dbDeck := model.Deck{}

	// check if the deck already exists
	filter := bson.D{{Key: "hash", Value: encodedHash}}
	searchRes := cardDecks.FindOne(context.Background(), filter).Decode(&dbDeck)
	if searchRes != mongo.ErrNoDocuments {
		return nil, fmt.Errorf("Deck already exists")
	}

	// map the GraphQL input to the model
	deck := model.Deck{
		ID:        dbDeck.ID,
		Subject:   input.Subject,
		Module:    input.Module,
		ModuleAlt: input.ModuleAlt,
		Examiners: input.Examiners,
		Semester:  input.Semester,
		Year:      input.Year,
		Hash:      encodedHash,
	}

	// insert deck
	_, insertErr := cardDecks.InsertOne(context.Background(), &deck)
	if insertErr != nil {
		return nil, insertErr
	}

	return &deck, nil
}

// Decks is the resolver for the decks field.
func (r *queryResolver) Decks(ctx context.Context) ([]*model.Deck, error) {
	cardDecks := r.DB.Collection("cardDecks")

	filter := bson.D{{}}
	cursor, searchErr := cardDecks.Find(context.Background(), filter)
	if searchErr != nil {
		return nil, searchErr
	}
	defer cursor.Close(context.Background())

	var decks []*model.Deck
	for cursor.Next(context.Background()) {
		var deck *model.Deck
		if err := cursor.Decode(&deck); err != nil {
			return nil, err
		}
		decks = append(decks, deck)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return decks, nil
}

// GetDeck is the resolver for the getDeck field.
func (r *queryResolver) GetDeck(ctx context.Context, id string) (*model.Deck, error) {
	panic("unimplemented")
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//   - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//     it when you're done.
//   - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *deckResolver) UUID(ctx context.Context, obj *model.Deck) (string, error) {
	return obj.ID, nil
}

type deckResolver struct{ *Resolver }
