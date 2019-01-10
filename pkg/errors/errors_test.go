package errors_test

import (
	. "manigandand-golang-test/pkg/errors"
	"net/http"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Errors", func() {
	It("Should create new app error", func() {
		err := NewAppError(http.StatusBadRequest, "invalid recipe id")
		Expect(err.Status).To(Equal(http.StatusBadRequest))
		Expect(err.Error()).To(Equal("invalid recipe id"))

		err = NewAppError(http.StatusInternalServerError, "client time out")
		Expect(err.Status).To(Equal(http.StatusInternalServerError))
		Expect(err.Error()).To(Equal("client time out"))

		err = NewAppError(http.StatusUnprocessableEntity, "max recipes ids")
		Expect(err.Status).To(Equal(http.StatusUnprocessableEntity))
		Expect(err.Error()).To(Equal("max recipes ids"))
	})
	It("Should create new bad request error", func() {
		err := BadRequest("invalid recipe id")
		Expect(err.Status).To(Equal(http.StatusBadRequest))
		Expect(err.Error()).To(Equal("invalid recipe id"))
	})
	It("Should create new not found error", func() {
		err := NotFound("recipe not found for the given key")
		Expect(err.Status).To(Equal(http.StatusNotFound))
		Expect(err.Error()).To(Equal("recipe not found for the given key"))
	})
	It("Should create new UnprocessableEntity app error", func() {
		err := UnprocessableEntity("max recipes ids")
		Expect(err.Status).To(Equal(http.StatusUnprocessableEntity))
		Expect(err.Error()).To(Equal("max recipes ids"))
	})
	It("Should create new internal server error", func() {
		err := InternalServer("client time out")
		Expect(err.Status).To(Equal(http.StatusInternalServerError))
		Expect(err.Error()).To(Equal("client time out"))
	})
	It("Should check for not found error", func() {
		err := NotFound("recipe not found for the given key")
		Expect(err.Status).To(Equal(http.StatusNotFound))
		Expect(err.Error()).To(Equal("recipe not found for the given key"))
		Ω(err.IsStatusNotFound()).Should(BeTrue())

		err = InternalServer("client time out")
		Ω(err.IsStatusNotFound()).Should(BeFalse())
	})
})
