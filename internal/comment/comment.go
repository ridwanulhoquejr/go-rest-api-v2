package comment

import (
	"context"
	"errors"
	"fmt"
)

var (
	ErrFetchingComment = errors.New("failed to fetch comment by id")
	ErrNotImplemented  = errors.New("not implemented")
)

// Comment - a representation of the comment
// ? structure for our Service
type Comment struct {
	ID     string `json:"id"`
	Slug   string `json:"slug"`
	Body   string `json:"body"`
	Author string `json:"author"`
}

// Store interface: its a contract
// it ensure that whoeve satisfied this interface decalation
// they will be directly communicating with these decalared func
// for example:
// we defince Store struct in Service layer, so whoever had an instance of this Service
// they will access the interface decalaration func form anywhere
type Store interface {
	GetComment(context.Context, string) (Comment, error)
	PostComment(context.Context, Comment) (Comment, error)
	DeleteComment(context.Context, string) error
	UpdateComment(context.Context, string, Comment) (Comment, error)
	GetMultipleComment(context.Context) ([]Comment, error)
}

// Service - is the struct on which all our
// logic will be built on top of
type Service struct {
	// here, we ultimatly have a db connection
	Store Store //? db connection

	//? why a struct  field as an interface?
	//ans:
	/*
		the Service struct has a field Store of type Store, which is an interface.
		This means that any value assigned to the Store field
		must implement the methods defined by the Store interface.
	*/
	//* We must have to implement all of the methods decalred inside Store interface, in Repository layer ie, comment_db.go
	//* since our Store is a DB instance which is [`a reciver of *Database struct`]
	//* and our repository layer is also ['a reciver of *Database struct']
	//* so we can access the methods of the repository layer from the Service layer
}

// NewService - its like a constructor
// returns a pointer to a new [Service] struct,
// where Service.Store is a db connection
// so, every method in this interface can access the db connection
func NewService(store Store) *Service {
	return &Service{
		Store: store,
	}
}

// Implementing the declared methods

// GetMultipleComment - get all the comments
func (s *Service) GetMultipleComment(
	ctx context.Context,
) ([]Comment, error) {

	///? bcz of repository methods are a reciver of (*Database) struct
	//? and our Store is also a reciver of (*Database) struct
	//? eventually, we can access the repo method bcz
	//? that method can be accessibale by an instance of *Database struct
	//! Long story short:
	//* We make these methods a reciver of *Database struct, so that we can access these methods accross the layers

	cmts, err := s.Store.GetMultipleComment(ctx)

	if err != nil {
		fmt.Println(err)
		return []Comment{}, err
	}

	fmt.Println(cmts)
	return cmts, nil
}

func (s *Service) GetComment(
	ctx context.Context,
	id string,
) (Comment, error) {

	fmt.Println("retreiving a comment")
	cmt, err := s.Store.GetComment(ctx, id)
	if err != nil {
		fmt.Println(err)
		return Comment{}, ErrFetchingComment
	}
	return cmt, nil
}

func (s *Service) UpdateComment(
	ctx context.Context,
	id string,
	cmt Comment,
) (Comment, error) {

	updatedCmt, err := s.Store.UpdateComment(ctx, id, cmt)

	if err != nil {
		fmt.Println(err)
		return Comment{}, err
	}

	return updatedCmt, nil
}

func (s *Service) DeleteComment(
	ctx context.Context,
	id string,
) error {

	err := s.Store.DeleteComment(ctx, id)
	if err != nil {
		fmt.Println(err)
		return err

	}

	fmt.Println("deleting a comment")

	return nil
}

func (s *Service) PostComment(
	ctx context.Context,
	cmt Comment,
) (Comment, error) {

	// communitcate with the Repository layer: Store.PostComment
	// which is a member of the Service struct
	// beacause, when we instantiate the Service struct, we pass the Store struct i,e a db connection
	// which eventually store in a Service struct property called Store
	// and the interface decalred in the Store struct, so we can access the PostComment method
	// from the implementation of the Store interface
	// so that we can call the Method form reppo layer by calling the Store.PostComment;
	// which also takes a reciver of the Store struct i,e a db connection

	insertedCmt, err := s.Store.PostComment(ctx, cmt)

	if err != nil {
		fmt.Println(err)
		return Comment{}, err
	}

	return insertedCmt, nil
}
