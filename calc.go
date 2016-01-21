// Package aqicalc provides a calculator to calculate AQI from the
// concentration value.
package aqicalc

type Conc struct {
	SO2_24hr, NO2_24hr, PM10_24hr, CO_24hr, O3_1hr, PM25_24hr float64
}

type AirQuality struct {
	AQI     int
	Primary string
}

var so2_RANGE = []float64{0, 50, 150, 475, 800, 1600, 2100, 2620}
var no2_RANGE = []float64{0, 40, 80, 180, 280, 565, 750, 940}
var pm10_RANGE = []float64{0, 50, 150, 250, 350, 420, 500, 600}
var co_RANGE = []float64{0, 2, 4, 14, 24, 36, 48, 60}
var o3_RANGE = []float64{0, 160, 200, 300, 400, 800, 1000, 1200}
var pm25_RANGE = []float64{0, 35, 75, 115, 150, 250, 350, 500}
var iaqi_RANGE = []float64{0, 50, 100, 150, 200, 300, 400, 500}

func linear(value, fromLow, fromHigh, toLow, toHigh float64) float64 {
	return ((value-fromLow)/(fromHigh-fromLow))*(toHigh-toLow) + toLow
}

func constrain(x, a, b float64) float64 {
	if x < a {
		return a
	} else if x > b {
		return b
	} else {
		return x
	}
}

func findDomainRange(d, r []float64, v float64) (d_lo, d_hi, r_lo, r_hi float64) {
	var idx int
	for i, _ := range d[:len(d)-1] {
		idx = i
		if d[i] <= v && v <= d[i+1] {
			break
		}
	}
	d_lo, d_hi, r_lo, r_hi = d[idx], d[idx+1], r[idx], r[idx+1]
	return
}

func toIAQI(conc float64, d []float64) float64 {
	var d_lo, d_hi, r_lo, r_hi float64 = findDomainRange(d, iaqi_RANGE, conc)
	return constrain(linear(conc, d_lo, d_hi, r_lo, r_hi), 0, 500)
}

func CalculateAQI(conc Conc) AirQuality {
	var iso2, ino2, ipm10, ico, io3, ipm25, max float64
	var primary string

	iso2 = toIAQI(conc.SO2_24hr, so2_RANGE)
	ino2 = toIAQI(conc.NO2_24hr, no2_RANGE)
	ipm10 = toIAQI(conc.PM10_24hr, pm10_RANGE)
	ico = toIAQI(conc.CO_24hr, co_RANGE)
	io3 = toIAQI(conc.O3_1hr, o3_RANGE)
	ipm25 = toIAQI(conc.PM25_24hr, pm25_RANGE)

	max, primary = iso2, "SO2"

	if ino2 >= max {
		max, primary = ino2, "NO2"
	}
	if ipm10 >= max {
		max, primary = ipm10, "PM10"
	}
	if ico >= max {
		max, primary = ico, "CO"
	}
	if io3 >= max {
		max, primary = io3, "O3"
	}
	if ipm25 >= max {
		max, primary = ipm25, "PM2.5"
	}

	return AirQuality{int(max + 0.5), primary}
}
