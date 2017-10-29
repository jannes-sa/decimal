package decimal

import "testing"

func TestDecimalV1(t *testing.T) {
	a := 0.07
	b := 0.07
	c := 0.07

	varCheck := map[string]float64{
		"add": 0.21,
		"sub": 0.11,
		"mul": 0.0231,
		"div": 0.11,
	}

	capCheck := map[string]float64{
		"add": calcDecimal([]float64{a, b, c}, `+`),
		"sub": calcDecimal([]float64{varCheck["add"], 0.1}, `-`),
		"mul": calcDecimal([]float64{varCheck["sub"], varCheck["add"]}, `*`),
		"div": calcDecimal([]float64{varCheck["mul"], varCheck["add"]}, `/`),
	}

	i := 0
	for k, v := range capCheck {
		if v != varCheck[k] {
			t.Fatalf("%v Value Incorrect %v Should Be %v", k, v, varCheck[k])
		}

		i++
	}

}

func TestCalcNested(t *testing.T) {
	a := 0.07
	b := 0.07
	c := 0.07
	shouldVal := 0.0105

	val := calcDecimal([]float64{
		calcDecimal([]float64{a, b, c}, `+`),
		0.05,
	}, `*`)

	if val != shouldVal {
		t.Fatalf("Value is %v Should Be %v", val, shouldVal)
	}
}
