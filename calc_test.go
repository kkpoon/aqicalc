package aqicalc

import (
	"fmt"
	"testing"
)

func Test_so2_iaqi(t *testing.T) {
	cases := []struct {
		in, expected float64
	}{
		{0, 0},
		{50, 50},
		{150, 100},
		{475, 150},
		{800, 200},
		{1600, 300},
		{2100, 400},
		{2620, 500},
		{12620, 500},
	}
	for _, c := range cases {
		actual := toIAQI(c.in, so2_RANGE)
		if actual != c.expected {
			t.Errorf("convert SO2 conc to IAQI: input=%v, expected=%v, actual=%v",
				c.in, c.expected, actual)
		}
	}
}

func Test_no2_iaqi(t *testing.T) {
	cases := []struct {
		in, expected float64
	}{
		{0, 0},
		{40, 50},
		{80, 100},
		{180, 150},
		{280, 200},
		{565, 300},
		{750, 400},
		{940, 500},
		{1940, 500},
	}
	for _, c := range cases {
		actual := toIAQI(c.in, no2_RANGE)
		if actual != c.expected {
			t.Errorf("convert NO2 conc to IAQI: input=%v, expected=%v, actual=%v",
				c.in, c.expected, actual)
		}
	}
}

func Test_pm10_iaqi(t *testing.T) {
	cases := []struct {
		in, expected float64
	}{
		{0, 0},
		{50, 50},
		{150, 100},
		{250, 150},
		{350, 200},
		{420, 300},
		{500, 400},
		{600, 500},
		{1600, 500},
	}
	for _, c := range cases {
		actual := toIAQI(c.in, pm10_RANGE)
		if actual != c.expected {
			t.Errorf("convert PM10 conc to IAQI: input=%v, expected=%v, actual=%v",
				c.in, c.expected, actual)
		}
	}
}

func Test_co_iaqi(t *testing.T) {
	cases := []struct {
		in, expected float64
	}{
		{0, 0},
		{2, 50},
		{4, 100},
		{14, 150},
		{24, 200},
		{36, 300},
		{48, 400},
		{60, 500},
		{160, 500},
	}
	for _, c := range cases {
		actual := toIAQI(c.in, co_RANGE)
		if actual != c.expected {
			t.Errorf("convert CO conc to IAQI: input=%v, expected=%v, actual=%v",
				c.in, c.expected, actual)
		}
	}
}

func Test_o3_iaqi(t *testing.T) {
	cases := []struct {
		in, expected float64
	}{
		{0, 0},
		{160, 50},
		{200, 100},
		{300, 150},
		{400, 200},
		{800, 300},
		{1000, 400},
		{1200, 500},
		{11200, 500},
	}
	for _, c := range cases {
		actual := toIAQI(c.in, o3_RANGE)
		if actual != c.expected {
			t.Errorf("convert O3 conc to IAQI: input=%v, expected=%v, actual=%v",
				c.in, c.expected, actual)
		}
	}
}

func Test_pm25_iaqi(t *testing.T) {
	cases := []struct {
		in, expected float64
	}{
		{0, 0},
		{35, 50},
		{75, 100},
		{115, 150},
		{150, 200},
		{250, 300},
		{350, 400},
		{500, 500},
		{1500, 500},
	}
	for _, c := range cases {
		actual := toIAQI(c.in, pm25_RANGE)
		if actual != c.expected {
			t.Errorf("convert PM2.5 conc to IAQI: input=%v, expected=%v, actual=%v",
				c.in, c.expected, actual)
		}
	}
}

func ExampleCalculateAQI() {
	fmt.Println(CalculateAQI(Conc{0, 0, 0, 0, 0, 0}))
	fmt.Println(CalculateAQI(Conc{300, 0, 0, 0, 0, 0}))
	fmt.Println(CalculateAQI(Conc{0, 300, 0, 0, 0, 0}))
	fmt.Println(CalculateAQI(Conc{0, 0, 460, 0, 0, 0}))
	fmt.Println(CalculateAQI(Conc{0, 0, 0, 23, 0, 0}))
	fmt.Println(CalculateAQI(Conc{0, 0, 0, 0, 216, 0}))
	fmt.Println(CalculateAQI(Conc{0, 0, 0, 0, 0, 233}))
	// Output:
	// {0 PM2.5}
	// {123 SO2}
	// {207 NO2}
	// {350 PM10}
	// {195 CO}
	// {108 O3}
	// {283 PM2.5}
}
