package saver_test

import (
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	"github.com/ozoncp/ocp-calendar-api/internal/app/models"
	"github.com/ozoncp/ocp-calendar-api/internal/mocks"
	"github.com/ozoncp/ocp-calendar-api/internal/saver"
	"time"
)

var _ = Describe("Saver", func() {
	var (
		ctrl        *gomock.Controller
		mockFlusher *mocks.MockFlusher
		s           saver.Saver
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		mockFlusher = mocks.NewMockFlusher(ctrl)
		s = saver.NewSaver(3, mockFlusher)
	})

	AfterEach(func() {
		ctrl.Finish()
		s.Close()
	})

	Context("Saver logic", func() {
		It("Should save changes in timeout", func() {
			mockFlusher.EXPECT().
				Flush(gomock.Any()).
				Times(2)
			s.Save(models.Calendar{
				Id:     0,
				UserId: 0,
				Type:   0,
				Link:   "",
			})
			s.Save(models.Calendar{
				Id:     0,
				UserId: 0,
				Type:   0,
				Link:   "",
			})
			time.Sleep(saver.SaveDelay + time.Second)
			s.Save(models.Calendar{
				Id:     0,
				UserId: 0,
				Type:   0,
				Link:   "",
			})
			s.Close()
			time.Sleep(time.Millisecond * 200)
		})

		It("Should save changes on close", func() {
			mockFlusher.EXPECT().
				Flush(gomock.Any()).
				Times(1)
			s.Save(models.Calendar{
				Id:     0,
				UserId: 0,
				Type:   0,
				Link:   "",
			})
			s.Close()
		})

		It("Should not save changes before timeout", func() {
			mockFlusher.EXPECT().
				Flush(gomock.Any()).
				Times(0)
			s.Save(models.Calendar{
				Id:     0,
				UserId: 0,
				Type:   0,
				Link:   "",
			})
		})
	})
})
