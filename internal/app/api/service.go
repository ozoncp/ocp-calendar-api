package api

import (
	"encoding/json"
	"github.com/Shopify/sarama"
	"github.com/ozoncp/ocp-calendar-api/internal/app/models"
	"github.com/ozoncp/ocp-calendar-api/internal/repo"
	desc "github.com/ozoncp/ocp-calendar-api/pkg/ocp-calendar-api"
	"github.com/prometheus/client_golang/prometheus"
)

const NotFoundMessage = "Calendar not found"
const SuccessLabelValue = "success"

var brokers = []string{"127.0.0.1:9094"}

var CreateCalendarCounter = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "http_request_post_create_calendar_counter",
		Help: "Number of create calendar request",
	},
	[]string{"status"},
)

var CreateCalendarsCounter = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "http_request_post_multi_create_calendars_counter",
		Help: "Number of multi create calendars request",
	},
	[]string{"status"},
)

var RemoveCalendarCounter = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "http_request_delete_remove_calendar_counter",
		Help: "Number of remove calendar request",
	},
	[]string{"status"},
)

var UpdateCalendarCounter = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "http_request_put_update_calendar_counter",
		Help: "Number of update calendar request",
	},
	[]string{"status"},
)

type OcpCalendarApi struct {
	desc.UnimplementedOcpCalendarApiServer
	repo     repo.Repo
	producer sarama.SyncProducer
}

func NewOcpCalendarApi(repo repo.Repo) (desc.OcpCalendarApiServer, error) {
	producer, err := newProducer()
	prometheus.MustRegister(
		CreateCalendarCounter,
		CreateCalendarsCounter,
		RemoveCalendarCounter,
		UpdateCalendarCounter,
	)

	if err != nil {
		return nil, err
	}

	return &OcpCalendarApi{
		repo:     repo,
		producer: producer,
	}, nil
}

func newProducer() (sarama.SyncProducer, error) {
	config := sarama.NewConfig()
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true
	producer, err := sarama.NewSyncProducer(brokers, config)

	return producer, err
}

func PrepareMessage(topic string, calendar []models.Calendar) (*sarama.ProducerMessage, error) {
	b, err := json.Marshal(calendar)

	if err != nil {
		return nil, err
	}

	msg := &sarama.ProducerMessage{
		Topic:     topic,
		Partition: -1,
		Value:     sarama.StringEncoder(b),
	}

	return msg, nil
}
