package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/ridwanulhoquejr/go-rest-api-v2/cmd/internal/comment"
	uuid "github.com/satori/go.uuid"
)

type CommentRow struct {
	ID     string
	Slug   sql.NullString
	Body   sql.NullString
	Author sql.NullString
}

func convertCommentRowToComment(c CommentRow) comment.Comment {
	return comment.Comment{
		ID:     c.ID,
		Slug:   c.Slug.String,
		Body:   c.Body.String,
		Author: c.Author.String,
	}
}

func (d *Database) GetComment(ctx context.Context, uuid string) (comment.Comment, error) {
	var cmtRow CommentRow

	row := d.Client.QueryRowContext(ctx,
		`SELECT * FROM comments
		 WHERE id = $1`,
		uuid,
	)

	err := row.Scan(
		&cmtRow.ID,
		&cmtRow.Slug,
		&cmtRow.Body,
		&cmtRow.Author,
	)
	if err != nil {
		return comment.Comment{}, fmt.Errorf("error featching comment by uuid: %w", err)
	}

	return convertCommentRowToComment(cmtRow), nil
}

func (d *Database) PostComment(ctx context.Context, c comment.Comment) (comment.Comment, error) {

	c.ID = uuid.NewV4().String()

	postRow := CommentRow{
		ID:     c.ID,
		Slug:   sql.NullString{String: c.Slug, Valid: true},
		Author: sql.NullString{String: c.Author, Valid: true},
		Body:   sql.NullString{String: c.Body, Valid: true},
	}

	rows, err := d.Client.NamedQueryContext(ctx,
		`INSERT INTO comments
		 (id, slug,  author, body)
		 VALUES (:id, :slug, :author, :body)`,
		postRow,
	)

	if err != nil {
		return comment.Comment{}, fmt.Errorf("error creating comment: %w", err)
	}
	if err := rows.Close(); err != nil {
		return comment.Comment{}, fmt.Errorf("error closing rows: %w", err)
	}

	return c, nil
}

func (d *Database) DeleteComment(ctx context.Context, uuid string) error {

	_, err := d.Client.ExecContext(ctx,
		`DELETE FROM comments
		 WHERE id = $1`,
		uuid,
	)

	if err != nil {
		return fmt.Errorf("error deleting comment by uuid: %w", err)
	}

	return nil
}

func (d *Database) UpdateComment(
	ctx context.Context,
	id string,
	c comment.Comment,
) (comment.Comment, error) {

	updateRow := CommentRow{
		ID:     id,
		Slug:   sql.NullString{String: c.Slug, Valid: true},
		Author: sql.NullString{String: c.Author, Valid: true},
		Body:   sql.NullString{String: c.Body, Valid: true},
	}

	row, err := d.Client.NamedQueryContext(
		ctx,
		`UPDATE comments SET
		slug = :slug,
		author = :author,
		body = :body
		WHERE id = :id`,
		updateRow,
	)

	if err != nil {
		return comment.Comment{}, fmt.Errorf("error updating comment: %w", err)
	}

	if err := row.Close(); err != nil {
		return comment.Comment{}, fmt.Errorf("error closing rows: %w", err)
	}

	return convertCommentRowToComment(updateRow), nil

}
