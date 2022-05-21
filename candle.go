package techan

import (
	"fmt"
	"strings"

	"github.com/sdcoffey/big"
)

type Candle interface {
	Period() TimePeriod
	OpenPrice() big.Decimal
	ClosePrice() big.Decimal
	MaxPrice() big.Decimal
	MinPrice() big.Decimal
	Volume() int64
	String() string
}

// Candle represents basic market information for a security over a given time period
type candle struct {
	period     TimePeriod
	openPrice  big.Decimal
	closePrice big.Decimal
	maxPrice   big.Decimal
	minPrice   big.Decimal
	volume     int64
}

type candleBuildOption = func(*candle)

// NewCandle returns a new *Candle for a given time period
func NewCandle(period TimePeriod, options ...candleBuildOption) Candle {
	candle := &candle{
		period:     period,
		openPrice:  big.ZERO,
		closePrice: big.ZERO,
		maxPrice:   big.ZERO,
		minPrice:   big.ZERO,
		volume:     0,
	}

	for i := range options {
		options[i](candle)
	}

	return candle
}

func WithOpenPrice(openPrice big.Decimal) candleBuildOption {
	return func(i *candle) {
		i.openPrice = openPrice
	}
}

func WithClosePrice(closePrice big.Decimal) candleBuildOption {
	return func(i *candle) {
		i.closePrice = closePrice
	}
}

func WithMinPrice(minPrice big.Decimal) candleBuildOption {
	return func(i *candle) {
		i.minPrice = minPrice
	}
}

func WithMaxPrice(maxPrice big.Decimal) candleBuildOption {
	return func(i *candle) {
		i.maxPrice = maxPrice
	}
}

func WithVolumePrice(volume int64) candleBuildOption {
	return func(i *candle) {
		i.volume = volume
	}
}

func (c *candle) Period() TimePeriod {
	return c.period
}

func (c *candle) OpenPrice() big.Decimal {
	return c.openPrice
}

func (c *candle) ClosePrice() big.Decimal {
	return c.closePrice
}

func (c *candle) MaxPrice() big.Decimal {
	return c.maxPrice
}

func (c *candle) MinPrice() big.Decimal {
	return c.minPrice
}

func (c *candle) Volume() int64 {
	return c.volume
}

func (c *candle) String() string {
	return strings.TrimSpace(fmt.Sprintf(
		`
Time:	%s
Open:	%s
Close:	%s
High:	%s
Low:	%s
Volume:	%d
	`,
		c.period,
		c.openPrice.FormattedString(2),
		c.closePrice.FormattedString(2),
		c.maxPrice.FormattedString(2),
		c.minPrice.FormattedString(2),
		c.volume,
	))
}
