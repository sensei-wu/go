package odds

import (
  "math"
  "fmt"
)

type DecimalOdds float64
type AmericanOdds int64

func DecimalToAmericanWithRoundingDown(d DecimalOdds) AmericanOdds {
	if d < 2.00 {
		fmt.Printf("Value before rounding: %.2f\n", -100/(float64(d)-1))
		return AmericanOdds(math.Floor(-100 / (float64(d) - 1)))
	} else {
		fmt.Printf("Value before rounding: %.2f\n", (float64(d)-1)*100)
		return AmericanOdds(math.Floor((float64(d) - 1) * 100))
	}
}

func DecimalToAmericanWithNoRoundingDown(d DecimalOdds) AmericanOdds {
	if d < 2.00 {
		return AmericanOdds(math.Round(-100 / (float64(d) - 1)))
	} else {
		return AmericanOdds(math.Round((float64(d) - 1) * 100))
	}
}

func AmericanToDecimal(a AmericanOdds) DecimalOdds {
	if a < 0 {
		return DecimalOdds((a / 100) + 1)
	} else {
		return DecimalOdds((100 / a) + 1)
	}
}
