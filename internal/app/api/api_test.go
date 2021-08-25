package api_test

import (
	context "context"
	"database/sql"
	"encoding/json"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
	"github.com/ozoncp/ocp-calendar-api/internal/app/api"
	"github.com/ozoncp/ocp-calendar-api/internal/app/models"
	"github.com/ozoncp/ocp-calendar-api/internal/mocks"
	desc "github.com/ozoncp/ocp-calendar-api/pkg/ocp-calendar-api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/test/bufconn"
	"log"
	"net"
)

var _ = Describe("Api", func() {
	var (
		ctrl     *gomock.Controller
		mockRepo *mocks.MockRepo
		bufSize  = 1024 * 1024
		listener *bufconn.Listener
		conn     *grpc.ClientConn
		ctx      context.Context
		client   desc.OcpCalendarApiClient
	)

	ctxDialer := func(context.Context, string) (net.Conn, error) {
		return listener.Dial()
	}

	BeforeSuite(func() {
		ctrl = gomock.NewController(GinkgoT())
		mockRepo = mocks.NewMockRepo(ctrl)
		listener = bufconn.Listen(bufSize)
		server := grpc.NewServer()
		desc.RegisterOcpCalendarApiServer(server, api.NewOcpCalendarApi(mockRepo))

		go func() {
			if err := server.Serve(listener); err != nil {
				log.Fatalf("Server exited with error: %v", err)
			}
		}()
	})

	AfterSuite(func() {
		ctrl.Finish()
	})

	BeforeEach(func() {
		ctx = context.Background()
		conn, _ = grpc.DialContext(
			ctx,
			"",
			grpc.WithInsecure(),
			grpc.WithContextDialer(ctxDialer),
		)
		client = desc.NewOcpCalendarApiClient(conn)
	})

	AfterEach(func() {
		if err := conn.Close(); err != nil {
			log.Fatalf("can't close connection %v", err)
		}
	})

	Context("Create Calendar", func() {
		It("should not create calendar with invalid data", func() {
			mockRepo.EXPECT().
				AddCalendars(gomock.Any(), gomock.Any()).
				Times(0)

			res, err := client.CreateCalendarV1(ctx, &desc.CreateCalendarRequestV1{
				Link: "Ha-Ha",
			})

			gomega.Expect(res).Should(gomega.BeNil())
			gomega.Expect(status.Code(err)).Should(gomega.BeEquivalentTo(codes.InvalidArgument))
		})

		It("should success create calendar", func() {
			mockRepo.EXPECT().
				AddCalendars(gomock.Any(), gomock.Any()).
				Times(1)

			res, err := client.CreateCalendarV1(ctx, &desc.CreateCalendarRequestV1{
				UserId: 1,
				Type:   1,
				Link:   "Ha-Ha",
			})

			gomega.Expect(json.Marshal(res)).Should(gomega.BeEquivalentTo("{}"))
			gomega.Expect(err).Should(gomega.BeNil())
		})
	})

	Context("Remove Calendar", func() {
		It("should remove calendar by email", func() {
			mockRepo.EXPECT().
				RemoveCalendar(gomock.Any(), gomock.Any()).
				Times(1)

			res, err := client.RemoveCalendarV1(ctx, &desc.RemoveCalendarRequestV1{
				Id: 1,
			})

			gomega.Expect(json.Marshal(res)).Should(gomega.BeEquivalentTo("{}"))
			gomega.Expect(err).Should(gomega.BeNil())
		})
	})

	Context("Describe Calendar", func() {
		It("should return calendar data", func() {
			mockRepo.EXPECT().
				DescribeCalendar(gomock.Any(), gomock.Any()).
				Times(1)

			_, err := client.DescribeCalendarV1(ctx, &desc.DescribeCalendarRequestV1{
				Id: 1,
			})

			gomega.Expect(err).Should(gomega.BeNil())
		})

		It("should return not found error", func() {
			mockRepo.EXPECT().
				DescribeCalendar(gomock.Any(), gomock.Any()).
				Times(1).
				Return(models.Calendar{}, sql.ErrNoRows)
			_, err := client.DescribeCalendarV1(ctx, &desc.DescribeCalendarRequestV1{
				Id: 2,
			})

			gomega.Expect(status.Code(err)).Should(gomega.BeEquivalentTo(codes.NotFound))
		})
	})

	Context("List Calendar", func() {
		It("should return list of calendars", func() {
			mockRepo.EXPECT().
				ListCalendars(
					gomock.Any(),
					gomock.Any(),
					gomock.Any(),
					gomock.Any(),
					gomock.Any(),
				).Times(1)

			_, err := client.ListCalendarsV1(ctx, &desc.ListCalendarRequestV1{
				Limit:  0,
				Offset: 0,
				UserId: 0,
				Type:   0,
			})

			gomega.Expect(err).Should(gomega.BeNil())
		})
	})
})
