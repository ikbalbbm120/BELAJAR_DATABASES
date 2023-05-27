package repository

import (
	"belajar_database"
	"context"
	"golang_database/entity"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"testing"
)

func TestInsertComment(t *testing.T){
	commentRepository := NewCommentRepository(belajar_database.GetConnection())

	ctx := context.Background()
	comment := entity.Comment{
		Email : "repository@test.com",
		Comment : "Test repository",
	}

	result, err := commentRepository.Insert(ctx, comment)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func TestFindById(t *testing.T){
	commentRepository := NewCommentRepository(golang_database.GetConnection())

	comment, err := commentRepository.FindById(context.Background(), 38)
	if err != nil {
		panic(err)
	}

	fmt.Println(comment)
}

func TestFindAll(t *testing.T){
	commentRepository := NewCommentRepository(golang_database.GetConnection())

	comments, err := commentRepository.FindAll(context.Background())
	if err != nil {
		panic(err)
	}

	for _, comment := range comments {
			fmt.Println(comment)
		}	
}