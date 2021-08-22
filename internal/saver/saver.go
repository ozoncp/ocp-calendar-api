package saver

import (
	"github.com/ozoncp/ocp-calendar-api/internal/app/models"
	"github.com/ozoncp/ocp-calendar-api/internal/flusher"
	"time"
)

const SaveDelay = 2 * time.Second

type Saver interface {
	Save(calendar models.Calendar)
	Init()
	Close()
}

type saver struct {
	buffer  []models.Calendar
	done    chan bool
	flusher flusher.Flusher
}

func (s *saver) Save(calendar models.Calendar) {
	s.buffer = append(s.buffer, calendar)
}

func (s *saver) Close() {
	select {
	case _, ok := <-s.done:
		if !ok {
			return
		}
	default:
		s.done <- true
		close(s.done)
	}
}

func (s *saver) Init() {
	ticker := time.NewTicker(SaveDelay)
	defer ticker.Stop()

	flush := func() {
		if len(s.buffer) > 0 {
			s.flusher.Flush(s.buffer)
			s.buffer = nil
		}
	}

	for {
		select {
		case <-ticker.C:
			flush()
		case <-s.done:
			flush()
			return
		}
	}
}

func NewSaver(capacity uint, flusher flusher.Flusher) Saver {
	saverInstance := saver{
		buffer:  make([]models.Calendar, 0, capacity),
		done:    make(chan bool),
		flusher: flusher,
	}

	go saverInstance.Init()

	return &saverInstance
}
