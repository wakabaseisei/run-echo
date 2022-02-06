package usecase_test

import (
	"log"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	"github.com/wakabaseisei/runapp/domain"
	"github.com/wakabaseisei/runapp/usecase"
)

type dbRepositoryMock struct {
	usecase.DBRepository
	FakeConnect func() *gorm.DB
}

func (db *dbRepositoryMock) Connect() *gorm.DB {
	return db.FakeConnect()
}

type postRepositoryMock struct {
	usecase.PostRepository
	FakeFindByID func(db *gorm.DB, id int) (post domain.Post, err error)
}

func (repo *postRepositoryMock) FindByID(db *gorm.DB, id int) (post domain.Post, err error) {
	return repo.FakeFindByID(db, id)
}

func TestPostInteractor(t *testing.T) {
	// データベースをモック化する
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	if mock == nil {
		t.Error("sqlmock is null")
	}
	defer db.Close()

	mockDB, err := gorm.Open("mysql", db)
	if err != nil {
		t.Fatalf("")
	}

	postRepo := &postRepositoryMock{
		FakeFindByID: func(db *gorm.DB, id int) (post domain.Post, err error) {
			post = domain.Post{
				Id:          1,
				Title:       "ガラガラへびがやってくる",
				Content:     "お腹を空かせてやってくる",
				PublishDate: 1640270612610,
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			}
			return post, nil
		},
	}
	dbRepo := &dbRepositoryMock{
		FakeConnect: func() *gorm.DB {
			return mockDB
		},
	}

	postInteractor := usecase.PostInteractor{DB: dbRepo, PostRepository: postRepo}
	post, resultStatus := postInteractor.Get(1)
	log.Print(resultStatus)
	log.Print(post)
	if resultStatus.Error != nil {
		t.Fatalf("ttt")
	}
	if post.Id != 1 {
		t.Fatalf("postInteractor.Get() should return Post.Id = 1, but got = %d", post.Id)
	}
	if post.Title != "ガラガラへびがやってくる" {
		t.Fatalf("postInteractor.Get() should return Post.Title = %s, but got = %s", "ガラガラへびがやってくる", post.Title)
	}

	if post.Content != "お腹を空かせてやってくる" {
		t.Fatalf("postInteractor.Get() should return Post.Content = %s, but got = %s", "お腹を空かせてやってくる", post.Content)
	}

	if post.PublishDate != 1640270612610 {
		t.Fatalf("postInteractor.Get() should return Post.PublishDate = 1640270612610, but got = %d", post.PublishDate)
	}
}
