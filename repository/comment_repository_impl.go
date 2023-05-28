package repository

import (
	"belajar-golang-db/entity"
	"context"
	"database/sql"
	"errors"
	"strconv"
)

type CommentRepositoryImplement struct {
	DB *sql.DB
}

func newCommentRepository(db *sql.DB) CommentRepository {
	return &CommentRepositoryImplement{DB: db}
}
func (repo *CommentRepositoryImplement) Insert(ctx context.Context, comment entity.Comments) (entity.Comments, error) {
	querySql := "INSERT INTO comments(email,comment) VALUES(?,?)"
	result, err := repo.DB.ExecContext(ctx, querySql, comment.Email, comment.Comment)
	if err != nil {
		return comment, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return comment, err
	}
	comment.Id = int32(id)

	return comment, nil
}

func (repo *CommentRepositoryImplement) FindById(ctx context.Context, id int32) (entity.Comments, error) {
	querySql := "SELECT email FROM comments WHERE id = ?  LIMIT 1"
	rows, err := repo.DB.QueryContext(ctx, querySql, id)
	comment := entity.Comments{}
	if err != nil {
		return comment, err
	}
	defer rows.Close()
	if rows.Next() {
		var email string
		err = rows.Scan(&email)
		if err != nil {
			return comment, err
		}
		return comment, nil

	} else {
		return comment, errors.New("id dengan" + strconv.Itoa(int(id)) + "not found")
	}

}

func (repo *CommentRepositoryImplement) FindAll(ctx context.Context) ([]entity.Comments, error) {
	querySql := "SELECT id,email,comment FROM comments"
	rows, err := repo.DB.QueryContext(ctx, querySql)
	if err != nil {
		return nil, err

	}
	defer rows.Close()
	var comments []entity.Comments
	if rows.Next() {
		comment := entity.Comments{}
		rows.Scan(&comment.Id, &comment.Email, &comment.Comment)
		comments = append(comments, comment)

	}
	return comments, nil
}
