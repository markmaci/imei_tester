package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

type parsedIMEI struct {
	// define parsedIMEI struct

	typeAllocationCode string
	serialNumber       string
	checksum           int
}

func main() {
	// read in from terminal
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter IMEI")
	scanner.Scan()

	IMEI := scanner.Text()
	// run parse and validate functions on input
	parsed, err1 := (parseIMEI(IMEI))

	validity, err2 := (validateChecksum(IMEI))

	// if there are no errors, print the parsed IMEI
	if err1 == nil && err2 == nil {
		fmt.Println("TAC:", parsed.typeAllocationCode)
		fmt.Println("Serial Number:", parsed.serialNumber)
		fmt.Println("Checksum:", parsed.checksum)
	}

	// print the relavent error message
	if validity == false && err1 != nil {
		fmt.Println(err1)
	} else if err2 != nil {
		fmt.Println(err2)
	}

}

func parseIMEI(IMEI string) (parsed parsedIMEI, err error) {
	// length check
	length := len(IMEI)
	if length != 15 {
		err = errors.New("invalid length")
		return parsed, err

	}

	// parse for TAC as first 8 digits
	parsed.typeAllocationCode = IMEI[:8]
	// parse for serial number as next 6 digts
	parsed.serialNumber = IMEI[8:14]

	// parse checksum by converting last digit to int
	var checksumString string = IMEI[14:15]

	// convert last character to int for checksum calculations
	intChecksum, err := strconv.Atoi(checksumString)

	parsed.checksum = intChecksum

	return parsed, err
}

func validateChecksum(IMEI string) (bool, error) {

	var sum int
	var err error

	// luhn algorithm implementation for checksum validation
	for i := len(IMEI) - 1; i >= 0; i-- {
		current, err := strconv.Atoi(string(IMEI[i]))
		if err != nil {

			return false, errors.New("invalid IMEI")

		}

		if i%2 == 1 {
			current *= 2
			if current > 9 {
				current -= 9
			}
		}
		sum += current
	}

	// return boolean and error based on validity
	if sum%10 == 0 {
		return true, err
	}
	err = errors.New("invalid checksum")
	return false, err

}
