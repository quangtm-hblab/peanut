package usecase_test

import (
	"peanut/usecase"
	"testing"

	"peanut/repository/mock"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var bookRepo *mock.MockBookRepo
var contentRepo *mock.MockContentRepo
var bookUc usecase.BookUsecase
var contentUc usecase.ContentUsecase

func TestUsecase(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Usecase Suite")
}

var _ = BeforeSuite(func() {
	sqlDB, smock, _ := sqlmock.New()
	defer sqlDB.Close()

	// dsn := "root:9899@tcp(localhost:3307)/peanut_db?charset=utf8mb4&parseTime=True&loc=Local"
	// sqlDB, _ := sql.Open("mysql", dsn)
	db, err := gorm.Open(mysql.New(mysql.Config{
		Conn:                      sqlDB,
		SkipInitializeWithVersion: true, // auto configure based on currently MySQL version
	}), &gorm.Config{})
	Expect(err).To(BeNil())
	Expect(smock).NotTo(BeNil())
	Expect(db).NotTo(BeNil())

	ctrl := gomock.NewController(GinkgoT())
	defer ctrl.Finish()
	bookRepo = mock.NewMockBookRepo(ctrl)
	contentRepo = mock.NewMockContentRepo(ctrl)
	bookUc = usecase.NewBookUsecase(bookRepo)
	contentUc = usecase.ContentUsecase(contentRepo)
})
