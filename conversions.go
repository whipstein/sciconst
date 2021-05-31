package sciconst

type TemperatureScale int

const (
	Kelvin TemperatureScale = iota
	Celsius
	Fahrenheit
)

func (t TemperatureScale) IsValid() bool {
	if t == Kelvin || t == Celsius || t == Fahrenheit {
		return true
	}
	return false
}

func ConvertTemperature(t float64, from, to TemperatureScale) (tnew float64) {
	if !from.IsValid() || !to.IsValid() {
		panic("TemperatureScale is invalid!")
	}

	if from == Celsius {
		tnew = t + ZeroCelsius
	} else if from == Kelvin {
		tnew = t
	} else if from == Fahrenheit {
		tnew = (t-32)*5/9 + ZeroCelsius
	}

	if to == Celsius {
		tnew -= ZeroCelsius
	} else if to == Kelvin {
		return
	} else if to == Fahrenheit {
		tnew = (tnew-ZeroCelsius)*9/5 + 32
	}

	return
}

func Lambda2Nu(lambda float64) float64 {
	return C / lambda
}

func Nu2Lambda(nu float64) float64 {
	return C / nu
}
