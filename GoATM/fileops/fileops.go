package fileops

import (
	"fmt"
	"os"
	"strconv"
)

// function to write VALUE into a text file.
func WriteFloatToFile(value float64, filename string) {
	valuetext := fmt.Sprint(value)                  // converts balance float tpe into string.
	os.WriteFile(filename, []byte(valuetext), 0644) //takes the string and converts into byte code.
}

// function to read VALUE from a text file.
func GetFloatFromFile(filename string) float64 {
	data, _ := os.ReadFile(filename)              // data in byte code since its reading from file.
	valueText := string(data)                     // convert byte code data into string as byte code to float64 conversion does not work.
	value, _ := strconv.ParseFloat(valueText, 64) // to convert string into float 64.
	return value
}
