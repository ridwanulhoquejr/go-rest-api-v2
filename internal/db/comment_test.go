//go:build integration

package db

import (
	"context"
	"testing"

	"github.com/ridwanulhoquejr/go-rest-api-v2/internal/comment"
	"github.com/stretchr/testify/assert"
)

func TestCommentDatabase(t *testing.T) {
	t.Run("test create comment", func(t *testing.T) {
		db, err := NewDatabase()
		assert.NoError(t, err)

		cmt, err := db.PostComment(context.Background(), comment.Comment{
			Slug:   "go test",
			Author: "testuser",
			Body:   "body of test user",
		})
		assert.NoError(t, err)

		newCmt, err := db.GetComment(context.Background(), cmt.ID)
		assert.NoError(t, err)
		assert.Equal(t, "go test", newCmt.Slug)
	})

	t.Run("test delete comment", func(t *testing.T) {
		db, err := NewDatabase()
		assert.NoError(t, err)

		cmt, err := db.PostComment(context.Background(), comment.Comment{
			Slug:   "delete test",
			Author: "deletetestuser",
			Body:   "body of delete test user",
		})
		assert.NoError(t, err)

		// delete comment test
		err = db.DeleteComment(context.Background(), cmt.ID)
		assert.NoError(t, err)

		// try to get the deleted comment
		_, err = db.GetComment(context.Background(), cmt.ID)
		assert.Error(t, err)
	})
}
