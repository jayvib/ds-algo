package stock_buy_sell

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestMaxProfit(t *testing.T) {
	data := []int{100, 180, 260, 310, 40, 535, 695}
	want := 865
	got := MaxProfit(data, 0, len(data)-1)
	assert.Equal(t, want, got)
}
