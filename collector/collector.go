package collector

import (
	"context"
	"fmt"

	"github.com/lidofinance/terra-monitors/collector/config"
	"github.com/lidofinance/terra-monitors/collector/monitors"
	"github.com/lidofinance/terra-monitors/openapi/client"
	"github.com/sirupsen/logrus"
)

type Collector interface {
	Get(metric monitors.Metric) (float64, error)
	GetVector(metric monitors.Metric) (map[string]float64, error)
	ProvidedMetrics() []monitors.Metric
	ProvidedMetricVectors() []monitors.Metric
	UpdateData(ctx context.Context) []error
}

func NewLCDCollector(cfg config.CollectorConfig) LCDCollector {
	return LCDCollector{
		Metrics:       make(map[monitors.Metric]monitors.Monitor),
		MetricVectors: make(map[monitors.Metric]monitors.Monitor),
		logger:        cfg.Logger,
		apiClient:     cfg.GetTerraClient(),
	}
}

type LCDCollector struct {
	Metrics       map[monitors.Metric]monitors.Monitor
	MetricVectors map[monitors.Metric]monitors.Monitor
	Monitors      []monitors.Monitor
	logger        *logrus.Logger
	apiClient     *client.TerraLiteForTerra
}

func (c LCDCollector) GetApiClient() *client.TerraLiteForTerra {
	return c.apiClient
}

func (c LCDCollector) GetLogger() *logrus.Logger {
	return c.logger
}

func (c LCDCollector) ProvidedMetrics() []monitors.Metric {
	var metrics []monitors.Metric
	for m := range c.Metrics {
		metrics = append(metrics, m)
	}
	return metrics
}

func (c LCDCollector) ProvidedMetricVectors() []monitors.Metric {
	var metrics []monitors.Metric
	for m := range c.MetricVectors {
		metrics = append(metrics, m)
	}
	return metrics
}

func (c LCDCollector) Get(metric monitors.Metric) (float64, error) {
	monitor, found := c.Metrics[metric]
	if !found {
		return 0, fmt.Errorf("monitor for metric \"%s\" not found", metric)
	}
	return monitor.GetMetrics()[metric], nil
}

func (c LCDCollector) GetVector(metric monitors.Metric) (map[string]float64, error) {
	monitor, found := c.MetricVectors[metric]
	if !found {
		return nil, fmt.Errorf("monitor for metric vector \"%s\" not found", metric)
	}
	return monitor.GetMetricVectors()[metric], nil
}

func (c *LCDCollector) UpdateData(ctx context.Context) []error {
	var errors []error
	for _, monitor := range c.Monitors {
		err := monitor.Handler(ctx)
		if err != nil {
			// one error is not a reason to stop collecting data
			// collecting monitors errors to return them to calling code
			errors = append(errors, fmt.Errorf("failed to update data: %w", err))
		}
	}
	return errors
}

func (c *LCDCollector) RegisterMonitor(m monitors.Monitor) {
	m.InitMetrics()
	for metric := range m.GetMetrics() {
		if wantedMonitor, found := c.Metrics[metric]; found {
			panic(fmt.Sprintf("register monitor %s failed. metrics collision. Monitor %s has declared metric %s", m.Name(), wantedMonitor.Name(), metric))
		}

		c.Metrics[metric] = m
	}
	for metric := range m.GetMetricVectors() {
		if wantedMonitor, found := c.MetricVectors[metric]; found {
			panic(fmt.Sprintf("register monitor %s failed. metrics collision. Monitor %s has declared metric %s", m.Name(), wantedMonitor.Name(), metric))
		}

		c.MetricVectors[metric] = m
	}
	c.Monitors = append(c.Monitors, m)
}
