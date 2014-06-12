package simplejson

import (
	"math"
	"fmt"
)

func float64ToInt(value float64) int {
	if intValue := math.Trunc(value); intValue != value {
		panic(fmt.Sprintf("float64 could not be converted to int. Value was: %d", value))
	}else {
		return int(intValue)
	}
}
