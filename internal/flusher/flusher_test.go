package flusher_test

import (
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	"github.com/ozoncp/ocp-calendar-api/app/models"
	. "github.com/ozoncp/ocp-calendar-api/internal/flusher"
	"github.com/ozoncp/ocp-calendar-api/internal/mocks"
)

var _ = Describe("Flusher", func() {
	var (
		ctrl     *gomock.Controller
		mockRepo *mocks.MockRepo
		flusher  Flusher
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		mockRepo = mocks.NewMockRepo(ctrl)
		flusher = NewFlusher(1, mockRepo)
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	Context("", func() {
		It("should split to chunks and flush calendars entities collection", func() {
			mockRepo.EXPECT().
				AddCalendars(gomock.Any()).
				Times(2).
				Return(nil)

			flusher.Flush([]models.Calendar{{
				Id:     0,
				UserId: 0,
				Type:   0,
				Link:   "",
			}, {
				Id:     0,
				UserId: 0,
				Type:   0,
				Link:   "",
			}})
		})
	})
})
