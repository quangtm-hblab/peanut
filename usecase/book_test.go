package usecase_test

import (
	"errors"
	"peanut/domain"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Book", func() {

	var book domain.Book
	var updatingBookForm domain.UpdateBookForm
	var bookID int
	BeforeEach(func() {
		book = domain.Book{
			Name:  "test",
			Year:  1999,
			Price: 19.99,
		}
		//updateBookForm
		updatingBookForm.Name = &book.Name
		updatingBookForm.Year = &book.Year
		updatingBookForm.Price = nil
		bookID = 1
	})

	Describe("API get book by id", func() {
		Context("Found book", func() {
			It("Should be success", func() {
				//prepare
				bookRepo.EXPECT().GetBookById(bookID).Return(&book, nil)
				//do
				b, err := bookUc.GetBook(bookID)
				//check
				Expect(err).To(BeNil())
				Expect(b).NotTo(BeNil())
			})
		})
		Context("Not found book", func() {
			It("Should be failed", func() {
				//prepare
				bookRepo.EXPECT().GetBookById(bookID).Return(nil, errors.New("Book not found"))
				//do
				b, err := bookUc.GetBook(bookID)
				//check
				Expect(err).Should(HaveOccurred())
				Expect(b).To(BeNil())
			})
		})
	})

	Describe("API create book", func() {
		Context("Create book success", func() {
			It("Should be success", func() {
				//prepare
				bookRepo.EXPECT().CreateBook(book).Return(&book, nil)
				//do
				_, err := bookUc.CreateBook(book)
				//check
				Expect(err).To(BeNil())
			})
		})
		Context("Create book failed", func() {
			It("Should be failed", func() {
				//prepare
				bookRepo.EXPECT().CreateBook(book).Return(nil, errors.New("username exist"))
				//do
				_, err := bookUc.CreateBook(book)
				//check
				Expect(err).Should(HaveOccurred())
			})
		})
	})

	Describe("Api update book", func() {
		Context("Found book", func() {
			It("Should be update success", func() {
				//prepare
				//Find book
				bookRepo.EXPECT().GetBookById(bookID).Return(&book, nil)
				//Update book
				bookRepo.EXPECT().UpdateBook(book, bookID).Return(&book, nil)
				//do
				b, err := bookUc.UpdateBook(updatingBookForm, bookID)
				//check
				Expect(err).To(BeNil())
				Expect(b.Name).To(Equal(*updatingBookForm.Name))
				Expect(b.Year).To(Equal(*updatingBookForm.Year))
				Expect(b.Price).NotTo(BeNil())
			})
			It("Should be update fail", func() {
				//prepare
				//Find book <Repo>
				bookRepo.EXPECT().GetBookById(bookID).Return(&book, nil)
				//Update book <Repo>
				bookRepo.EXPECT().UpdateBook(book, bookID).Return(nil, errors.New("Internal error"))
				//do
				b, err := bookUc.UpdateBook(updatingBookForm, bookID)
				//check
				Expect(err).Should(HaveOccurred())
				Expect(b).To(BeNil())
			})
		})

		Context("Not found book", func() {
			It("Should be failed", func() {
				//prepare
				bookRepo.EXPECT().GetBookById(bookID).Return(nil, errors.New("Book not exist"))
				//do
				b, err := bookUc.UpdateBook(updatingBookForm, bookID)
				//check
				Expect(err).Should(HaveOccurred())
				Expect(b).To(BeNil())
			})
		})
	})

	Describe("API delete book", func() {
		Context("Delete success", func() {
			It("Should be success", func() {
				//prepare
				bookRepo.EXPECT().DeleteBook(bookID).Return(nil)
				//do
				err := bookUc.DeleteBook(bookID)
				//check
				Expect(err).To(BeNil())
			})
		})
		Context("Delete failed", func() {
			It("Should be failed", func() {
				//prepare
				bookRepo.EXPECT().DeleteBook(bookID).Return(errors.New("Delete failed"))
				//do
				err := bookUc.DeleteBook(bookID)
				//check
				Expect(err).Should(HaveOccurred())
			})
		})
	})

})
