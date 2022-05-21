package techan

import (
	"testing"
	"time"

	"fmt"
	"strings"

	"github.com/sdcoffey/big"
	"github.com/stretchr/testify/assert"
)

func TestCandle_String(t *testing.T) {
	now := time.Now()
	candle := NewCandle(TimePeriod{
		Start: now,
		End:   now.Add(time.Minute),
	},
		WithClosePrice(big.NewFromString("1")),
		WithOpenPrice(big.NewFromString("2")),
		WithMaxPrice(big.NewFromString("3")),
		WithMinPrice(big.NewFromString("0")),
		WithVolumePrice(10),
	)

	expected := strings.TrimSpace(fmt.Sprintf(`
Time:	%s
Open:	2.00
Close:	1.00
High:	3.00
Low:	0.00
Volume:	10
`, candle.Period()))

	assert.EqualValues(t, expected, candle.String())
}
