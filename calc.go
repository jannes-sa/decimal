package decimal

func calcDecimal(val []float64, operator string) (fl float64) {
	fl = 0
	if len(val) > 0 {
		calc := NewFromFloat(val[0])
		for _, v := range val[1:] {
			if operator == `+` {
				calc = calc.Add(NewFromFloat(v))
			} else if operator == `-` {
				calc = calc.Sub(NewFromFloat(v))
			} else if operator == `*` {
				calc = calc.Mul(NewFromFloat(v))
			} else if operator == `/` {
				calc = calc.Div(NewFromFloat(v))
			}
		}

		fl, _ = calc.Float64()
	}
	return
}
