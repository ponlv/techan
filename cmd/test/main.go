package main

import (
	"fmt"
	"time"

	"github.com/sdcoffey/big"
	"github.com/konrin/techan"
)

func main() {
	closePrices := []float64{
		1.2, 1.21, 1.23, 1.22, 6.7, 1.21,
		1.24, 1.243, 1.245, 1.246,
	}
	series := techan.NewTimeSeries()

	lastTime := time.Now()
	for i := range closePrices {
		candle := techan.NewCandle(
			techan.NewTimePeriod(lastTime, time.Minute),
			techan.WithClosePrice(big.NewDecimal(closePrices[i])),
		)

		series.AddCandle(candle)
		lastTime = lastTime.Add(time.Minute)
	}

	ema := techan.NewEMAIndicator(
		techan.NewClosePriceIndicator(series),
		4,
	)

	for i := range closePrices {
		fmt.Printf("%d -> %f\n", i+1, ema.Calculate(i).Float())
	}

	fmt.Println("----")

	newPrices := []float64{
		1.246, 2.4, 2.45, 2.4, 2.57,
	}

	for j := range newPrices {
		fmt.Printf("10 | %f -> ", series.LastCandle().ClosePrice().Float())

		candle := techan.NewCandle(
			techan.NewTimePeriod(lastTime, time.Minute),
			techan.WithClosePrice(big.NewDecimal(newPrices[j])),
		)
		series.AddCandle(candle)

		fmt.Printf("%f | ", series.LastCandle().ClosePrice().Float())

		result := ema.(techan.ResetCachedIndicator).ResetCacheByIndex(series.LastIndex()).Calculate(series.LastIndex())
		fmt.Printf("=> %f\n", result.Float())
	}
}
