package coins

import "log/slog"

type CoinInterface interface {
	InitializeCoinTable() error
	AddTrackedCoin(symbol string) error
}

type CoinService struct {
	logger *slog.Logger
}

func NewCoinService(logger *slog.Logger) *CoinService {
	return &CoinService{
		logger: logger,
	}
}

func (c *CoinService) InitializeCoinTable() error {
	c.logger.Info("Initializing coin table")
	return nil
}

func (c *CoinService) AddTrackedCoin(symbol string) error {
	c.logger.Info("Adding coin to table", "symbol", symbol)
	return nil
}
