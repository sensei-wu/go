package odds

import (
  "testing"
)

func TestDecimalToAmerican(t *testing.T) {
  var dOdds DecimalOdds = 1.952

  amOdds := DecimalToAmericanWithNoRoundingDown(dOdds)

  if amOdds != -105 {
    t.Logf("American Odds found %v", amOdds)
    t.Error("Expected -105 as American odds")
  }
}
