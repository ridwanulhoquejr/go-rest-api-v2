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

func (s *Service) UpdateComment(ctx context.Context, cmt Comment) error {
	return ErrNotImplemented
}

func (s *Service) DeleteComment(ctx context.Context, id string) error {
	return ErrNotImplemented
}

func (s *Service) CreateComment(ctx context.Context, cmt Comment) (Comment, error) {
	return Comment{}, ErrNotImplemented
}
