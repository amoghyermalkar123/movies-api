package tests

import (
	"log"
	"movieserver/db"
	"movieserver/types"
	"os"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Database Tests", func() {
	SetTestEnv()

	Describe("CheckAdminDetails", func() {
		It("Should return nil if user is admin and has correct details", func() {
			err := db.CheckAdminDetails(int64(4), "adminuser@gmail.com", "admin@123")
			Expect(err).To(BeNil())
		})

		It("Should return error if auth fails", func() {
			err := db.CheckAdminDetails(int64(1), "different@gmail.com", "aaa@123")
			Expect(err).ToNot(BeNil())
		})
	})

	Describe("CreateContent", func() {
		It("should create content", func() {
			err := db.CreateContent(&types.Movie{
				Name:        "test",
				PublishedAt: time.Now(),
			})
			Expect(err).To(BeNil())

			db.DeleteTestMovieData("test")
		})
	})

	Describe("SearchMovieByName", func() {
		It("Should return result movie data from search query", func() {
			searchQuery := "AndhaDhun"
			data, err := db.SearchMovieByName(searchQuery)

			Expect(err).To(BeNil())
			Expect(data).ToNot(BeNil())
			Expect(data.Name).To(Equal(searchQuery))
		})
	})

	Describe("GetMovieSearchDetails", func() {
		It("Should return aggregated result for movie data from search query", func() {
			searchQuery := "AndhaDhun"
			data, err := db.GetMovieSearchDetails(searchQuery)

			Expect(err).To(BeNil())
			Expect(data).ToNot(BeNil())
		})
	})

	Describe("GetUserInteractedMovies", func() {
		It("Should return result for users interaction data", func() {
			data, err := db.GetUserInteractedMovies(int64(1))

			Expect(err).To(BeNil())
			Expect(data).ToNot(BeNil())
		})
	})

	Describe("RateMovie", func() {
		It("Should return result for users interaction data", func() {
			err := db.RateMovie(&types.UserInteraction{
				UserID:  1,
				MovieID: 4,
				Rating:  4,
			})

			Expect(err).To(BeNil())
		})
	})

	Describe("CommentOnMovie", func() {
		It("Should return result for users interaction data", func() {
			err := db.CommentOnMovie(&types.UserInteraction{
				UserID:  1,
				MovieID: 4,
				Comment: "Greatest Movie Ever!!",
			})

			Expect(err).To(BeNil())
		})
	})
})

func SetTestEnv() {
	workingdir, err := os.Getwd()

	if err != nil {
		log.Println(err)
	}

	os.Setenv("configPath", workingdir+"../../../config/")

}
