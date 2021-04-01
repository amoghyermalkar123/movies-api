package tests

import (
	"net/http"
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/ghttp"
)

var _ = Describe("Handler Connection Tests", func() {
	var server *ghttp.Server

	BeforeEach(func() {
		// start a test http server
		server = ghttp.NewServer()
	})

	AfterEach(func() {
		//shut down the server between tests
		server.Close()
	})

	Describe("Admin Handler Verification", func() {
		BeforeEach(func() {
			server.AppendHandlers(
				ghttp.VerifyRequest("POST", "/create_content/"),
			)
		})

		It("Should verify endpoint", func() {
			r := strings.NewReader("Test")
			resp, err := http.Post(server.URL()+"/create_content/", "application/json", r)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(resp.StatusCode).Should(Equal(http.StatusOK))
		})
	})

	Describe("SearchForMovies Verification", func() {
		BeforeEach(func() {
			server.AppendHandlers(
				ghttp.VerifyRequest("GET", "/search_movie/:AndhaDhun"),
			)
		})

		It("Should verify endpoint", func() {
			resp, err := http.Get(server.URL() + "/search_movie/:AndhaDhun")
			Expect(err).ShouldNot(HaveOccurred())
			Expect(resp.StatusCode).Should(Equal(http.StatusOK))
		})
	})

	Describe("SearchUserInteractedMovies Verification", func() {
		BeforeEach(func() {
			server.AppendHandlers(
				ghttp.VerifyRequest("GET", "/search_user_movies/:1"),
			)
		})

		It("Should verify endpoint", func() {
			resp, err := http.Get(server.URL() + "/search_user_movies/:1")
			Expect(err).ShouldNot(HaveOccurred())
			Expect(resp.StatusCode).Should(Equal(http.StatusOK))
		})
	})

	Describe("RateMovie Verification", func() {
		BeforeEach(func() {
			server.AppendHandlers(
				ghttp.VerifyRequest("POST", "/rate/"),
			)
		})

		It("Should verify endpoint", func() {
			r := strings.NewReader("Test Data")
			resp, err := http.Post(server.URL()+"/rate/", "application/json", r)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(resp.StatusCode).Should(Equal(http.StatusOK))
		})
	})

	Describe("CommentMovie Verification", func() {
		BeforeEach(func() {
			server.AppendHandlers(
				ghttp.VerifyRequest("POST", "/comment/"),
			)
		})

		It("Should verify endpoint", func() {
			r := strings.NewReader("Test Data")
			resp, err := http.Post(server.URL()+"/comment/", "application/json", r)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(resp.StatusCode).Should(Equal(http.StatusOK))
		})
	})

	Describe("Auth Verification", func() {
		BeforeEach(func() {
			server.AppendHandlers(
				ghttp.VerifyRequest("POST", "/auth/"),
			)
		})

		It("Should verify endpoint", func() {
			r := strings.NewReader("Test Data")
			resp, err := http.Post(server.URL()+"/auth/", "application/json", r)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(resp.StatusCode).Should(Equal(http.StatusOK))
		})
	})
})
