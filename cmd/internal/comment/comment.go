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
// structure for our Service
type Comment struct {
	ID     string
	Slug   string
	Body   string
	Author string
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
}

// Service - is the struct on which all our
type Service struct {
	Store Store
}

// NewService - its like a constructor
// returns a pointer to a new [Service] struct
func NewService(store Store) *Service {
	return &Service{
		Store: store,
	}
}

// Implementing the declared methods
func (s *Service) GetComment(ctx context.Context, id string) (Comment, error) {

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

func (s *Service) DeleteComment(ctx context.Context, id string) error {

	err := s.Store.DeleteComment(ctx, id)
	if err != nil {
		fmt.Println(err)
		return err

	}

	fmt.Println("deleting a comment")

	return nil
}

func (s *Service) PostComment(ctx context.Context, cmt Comment) (Comment, error) {

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
