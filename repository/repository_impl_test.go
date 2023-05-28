package repository

import (
	belajar_golang_db "belajar-golang-db"
	"belajar-golang-db/entity"
	"context"
	"fmt"
	"testing"
)

func TestRepository(t *testing.T) {
	connection := belajar_golang_db.GetConnectionDb()
	commentRepository := newCommentRepository(connection)

	ctx := context.Background()
	comment := entity.Comments{
		Email:   "Test@gmai.com",
		Comment: "Anda Terbaik Lur",
	}
	result, err := commentRepository.Insert(ctx, comment)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(result)
}

func TestById(t *testing.T) {
	commentRepo := newCommentRepository(belajar_golang_db.GetConnectionDb())
	ctx := context.Background()
	comment, err := commentRepo.FindById(ctx, 37)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(comment)

}

func TestGetAllComments(t *testing.T) {
	commentRepo := newCommentRepository(belajar_golang_db.GetConnectionDb())
	ctx := context.Background()
	comments, err := commentRepo.FindAll(ctx)
	if err != nil {
		panic(err.Error())
	}
	for _, comment := range comments {
		fmt.Println(comment)
	}

}
