package usecase_test

import (
	"errors"
	"peanut/domain"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Content", func() {
	var content domain.Content
	var allContent []domain.Content
	BeforeEach(func() {
		content = domain.Content{
			Thumbnail:   "public/thumbnail/20230108233109WIN_20221010_10_33_29_Pro.jpg",
			Media:       "public/media/20230108233109WIN_20221010_10_33_36_Pro.jpg",
			Name:        "Bia Ha noi",
			Description: "Quang cao bia ha noi",
			Playtime:    100,
			Resolution:  1080,
			ARwidth:     3,
			ARheight:    2,
			Fever:       true,
			Ondemand:    true,
		}
		allContent = []domain.Content{content}
	})
	Describe("API create content", func() {
		Context("Create content successfully", func() {
			It("Should be success", func() {
				//prepare
				contentRepo.EXPECT().CreateContent(content).Return(&content, nil)
				//do
				newContent, err := contentUc.CreateContent(content)
				//check
				Expect(err).To(BeNil())
				Expect(newContent).NotTo(BeNil())
			})
		})
		Context("Create content failed", func() {
			It("Should be failed", func() {
				//prepare
				contentRepo.EXPECT().CreateContent(content).Return(nil, errors.New("some error"))
				//do
				_, err := contentUc.CreateContent(content)
				//check
				Expect(err).Should(HaveOccurred())
			})
		})
	})

	Describe("API get content", func() {
		Context("Get content successfully", func() {
			It("Should be success", func() {
				//prepare
				contentRepo.EXPECT().GetContents().Return(allContent, nil)
				//do
				all, err := contentUc.GetContents()
				//check
				Expect(err).To(BeNil())
				Expect(all).NotTo(BeNil())
			})
		})

		Context("Get content failed", func() {
			It("Should be failed", func() {
				//prepare
				contentRepo.EXPECT().GetContents().Return(nil, errors.New("some error"))
				//do
				_, err := contentUc.GetContents()
				//check
				Expect(err).Should(HaveOccurred())
			})
		})
	})

})
