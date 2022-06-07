package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

type parsedIMEI struct {
	//define parsedIMEI struct
	typeAllocationCode string
	serialNumber       string
	checksum           int
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter IMEI")
	scanner.Scan()

	IMEI := scanner.Text()

	parsed, err1 := (parseIMEI(IMEI))

	validity, err2 := (validateChecksum(IMEI))

	// fmt.Println(err1)
	// fmt.Println(err2)

	if err1 == nil && err2 == nil {
		fmt.Println("TAC:", parsed.typeAllocationCode)
		fmt.Println("Serial Number:", parsed.serialNumber)
		fmt.Println("Checksum:", parsed.checksum)
	}

	if validity == false && err1 != nil {
		fmt.Println(err1)
	} else {
		fmt.Println(err2)
	}

}

func parseIMEI(IMEI string) (parsed parsedIMEI, err error) {
	// length check
	length := len(IMEI)
	// fmt.Print(length)
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

	intChecksum, err := strconv.Atoi(checksumString)

	parsed.checksum = intChecksum

	return parsed, err
}

func validateChecksum(IMEI string) (bool, error) {

	var sum int
	var err error

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

	if sum%10 == 0 {
		return true, err
	}
	err = errors.New("invalid checksum")
	return false, err

}
