package techan

import (
	"testing"
	"time"

	"github.com/sdcoffey/big"
	"github.com/stretchr/testify/assert"
)

func TestNewVolumeIndicator(t *testing.T) {
	assert.NotNil(t, NewVolumeIndicator(NewTimeSeries()))
}

func TestVolumeIndicator_Calculate(t *testing.T) {
	series := NewTimeSeries()

	candle := NewCandle(TimePeriod{
		Start: time.Now(),
		End:   time.Now().Add(time.Minute),
	}, WithVolumePrice(2))

	series.AddCandle(candle)

	indicator := NewVolumeIndicator(series)
	assert.EqualValues(t, "2.000", indicator.Calculate(0).FormattedString(3))
}

func TestTypicalPriceIndicator_Calculate(t *testing.T) {
	series := NewTimeSeries()

	candle := NewCandle(TimePeriod{
		Start: time.Now(),
		End:   time.Now().Add(time.Minute),
	},
		WithMinPrice(big.NewFromString("1.2080")),
		WithMaxPrice(big.NewFromString("1.22")),
		WithClosePrice(big.NewFromString("1.215")),
	)

	series.AddCandle(candle)

	typicalPrice := NewTypicalPriceIndicator(series).Calculate(0)

	assert.EqualValues(t, "1.2143", typicalPrice.FormattedString(4))
}
