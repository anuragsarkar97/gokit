package observer

import (
	"context"
	"strings"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

const (
	Operation = "operation"
	Layer     = "layer"
	RequestId = "requestId"
	INFO      = "INFO"
	WARN      = "WARN"
	ERROR     = "ERROR"
	EventTag  = "event"
	Code      = "code"
	Namespace = "location_platform"
	Success   = "success"
	Failure   = "failure"
	Total     = "total"
)

type Observer struct {
	layer     string
	operation string
	requestId string
}

var genericLabels = []string{Layer, Operation, EventTag, Code}

var counter = promauto.NewCounterVec(prometheus.CounterOpts{
	Name: Namespace + "_requests_total",
	Help: "total requests processed",
}, genericLabels,
)

var histogram = promauto.NewHistogramVec(prometheus.HistogramOpts{
	Name:    Namespace + "_response_time",
	Help:    "Histogram of response time for handler",
	Buckets: prometheus.ExponentialBuckets(0.0001, 2, 18),
}, []string{Layer, Operation})

var gauge = promauto.NewGaugeVec(prometheus.GaugeOpts{
	Name: Namespace + "_guage",
	Help: "guage",
}, genericLabels)

func New(ctx context.Context) *Observer {
	ob := Observer{}
	ob.layer, _ = ctx.Value(Layer).(string)
	ob.operation, _ = ctx.Value(Operation).(string)
	ob.requestId, _ = ctx.Value(RequestId).(string)
	return &ob
}

type Event struct {
	Name     string
	Err      error
	Code     string
	Fields   map[string]interface{}
	Msg      string
	LogLevel string
}

func (as *Observer) Instrument(eventCtx Event) {
	if eventCtx.Err != nil {
		log.Error().Str(RequestId, as.requestId).Fields(eventCtx.Fields).Err(eventCtx.Err).Str(Layer, as.layer).Str(Operation, as.operation).
			Str(EventTag, eventCtx.Name).Str(Code, eventCtx.Code).Msg(eventCtx.Msg)
	} else {
		var event *zerolog.Event
		event = as.getLogLevel(eventCtx, event)
		event.Str(RequestId, as.requestId).Fields(eventCtx.Fields).Str(Layer, as.layer).Str(Operation, as.operation).
			Str(EventTag, eventCtx.Name).Str(Code, eventCtx.Code).Msg(eventCtx.Msg)
	}
	labels := []string{as.layer, as.operation, eventCtx.Name, eventCtx.Code}
	counter.WithLabelValues(labels...).Inc()
}

func (as *Observer) getLogLevel(eventCtx Event, event *zerolog.Event) *zerolog.Event {
	switch strings.ToUpper(eventCtx.LogLevel) {
	case INFO:
		event = log.Info()
	case WARN:
		event = log.Warn()
	case ERROR:
		event = log.Error()
	default:
		event = log.Debug()
	}
	return event
}

func (as *Observer) RecordTime(startTime time.Time) {
	d := time.Since(startTime)
	histogram.WithLabelValues(as.layer, as.operation).Observe(d.Seconds())
}

func Instrument(layer, operation, eventName, code string) {
	labels := []string{layer, operation, eventName, code}
	counter.WithLabelValues(labels...).Inc()
}

func InstrumentLatency(startTime time.Time, layer, operation string) {
	d := time.Since(startTime)
	histogram.WithLabelValues(layer, operation).Observe(d.Seconds())
}

func (as *Observer) InstrumentGuage(eventName, code string, value float64) {
	labels := []string{as.layer, as.operation, eventName, code}
	gauge.WithLabelValues(labels...).Set(value)
}

func MeasureDelay(delay float64, lvs ...string) {
	histogram.WithLabelValues(lvs...).Observe(delay)
}
