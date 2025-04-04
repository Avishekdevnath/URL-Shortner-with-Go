// S:\SDE\Hard Core\Learn\Golang\Projects\URL-Shortner-with-Go\Backend\internal\store\store.go
package store

import (
	"Backend/internal/models"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	cacheHits = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "cache_hits_total",
		Help: "Number of cache hits by store",
	}, []string{"store"})
)

func init() {
	prometheus.MustRegister(cacheHits)
}

type Store interface {
	Save(request *models.ShortURLRequest) (*models.ShortURLResponse, error)
	Get(shortCode string) (*models.ShortURLResponse, error)
	GetAll() ([]models.URL, error)
	Health() error
}

type HybridStore struct {
	Postgres Store
	Redis    Store
	Memory   Store
}

func (h *HybridStore) Save(request *models.ShortURLRequest) (*models.ShortURLResponse, error) {
	resp, err := h.Postgres.Save(request)
	if err != nil {
		return nil, err
	}
	_, _ = h.Redis.Save(request)
	_, _ = h.Memory.Save(request)
	return resp, nil
}

func (h *HybridStore) Get(shortCode string) (*models.ShortURLResponse, error) {
	if resp, err := h.Memory.Get(shortCode); err == nil {
		cacheHits.WithLabelValues("memory").Inc()
		return resp, nil
	}
	if resp, err := h.Redis.Get(shortCode); err == nil {
		cacheHits.WithLabelValues("redis").Inc()
		h.Memory.Save(&models.ShortURLRequest{URL: resp.OriginalURL})
		return resp, nil
	}
	resp, err := h.Postgres.Get(shortCode)
	if err != nil {
		return nil, err
	}
	cacheHits.WithLabelValues("postgres").Inc()
	_, _ = h.Redis.Save(&models.ShortURLRequest{URL: resp.OriginalURL})
	_, _ = h.Memory.Save(&models.ShortURLRequest{URL: resp.OriginalURL})
	return resp, nil
}

func (h *HybridStore) GetAll() ([]models.URL, error) {
	if urls, err := h.Redis.GetAll(); err == nil && len(urls) > 0 {
		cacheHits.WithLabelValues("redis").Inc()
		for _, url := range urls {
			_, _ = h.Memory.Save(&models.ShortURLRequest{URL: url.OriginalURL})
		}
		return urls, nil
	}
	urls, err := h.Postgres.GetAll()
	if err != nil {
		return nil, err
	}
	cacheHits.WithLabelValues("postgres").Inc()
	for _, url := range urls {
		_, _ = h.Redis.Save(&models.ShortURLRequest{URL: url.OriginalURL})
		_, _ = h.Memory.Save(&models.ShortURLRequest{URL: url.OriginalURL})
	}
	return urls, nil
}

func (h *HybridStore) Health() error {
	if err := h.Postgres.Health(); err != nil {
		return err
	}
	if err := h.Redis.Health(); err != nil {
		return err
	}
	return nil // Memory is always "healthy" in-process
}