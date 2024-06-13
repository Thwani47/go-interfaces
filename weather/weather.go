package weather

import "log"

type Celsius float64

func (c Celsius) Unit() string {
	return "°C"
}

func (c Celsius) Amount() float64 {
	return float64(c)
}

func Current() (Celsius, error) {
	c := Celsius(24.1)
	unit := c.Unit()
	log.Printf("INFO: the temperature is %.1f %s\n", c, unit)
	return c, nil
}
