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
	// convert CRLF to LF
	IMEI := scanner.Text()
	// IMEI := "358692051234567"

	parsed, err1 := (parseIMEI(IMEI))

	validity, err2 := (validateChecksum(IMEI))

	// if err1 != nil {
	// 	fmt.Errorf("invalid length")
	// }
	// if validity == false || err2 != nil {
	// 	fmt.Errorf("invalid imei")
	// }
	fmt.Print(err1, validity, err2)
	fmt.Println("TAC:", parsed.typeAllocationCode)
	fmt.Println("Serial Number:", parsed.serialNumber)
	fmt.Println("Checksum:", parsed.checksum)

}

func parseIMEI(IMEI string) (parsed parsedIMEI, err error) {
	// length check
	length := len(IMEI)
	fmt.Print(length)
	if length != 15 {

		fmt.Print("length is bad")

		return parsed, errors.New("invalid length")

	}

	// parse for TAC as first 8 digits
	parsed.typeAllocationCode = IMEI[:8]
	// parse for serial number as next 6 digts
	parsed.serialNumber = IMEI[8:14]

	// parse checksum by converting last digit to int
	var checksumString string = IMEI[14:15]

	intChecksum, err := strconv.Atoi(checksumString)
	if err != nil {
		fmt.Print("wtf2")
		return parsed, errors.New("invalid imei")
	}
	parsed.checksum = intChecksum

	return parsed, err
}

func validateChecksum(IMEI string) (bool, error) {

	var sum int
	var err error

	for i := len(IMEI) - 1; i >= 0; i-- {
		current, err := strconv.Atoi(string(IMEI[i]))
		if err != nil {
			fmt.Print("wtf3")
			return false, fmt.Errorf("invalid IMEI")

		}

		if i%2 == 1 {
			current *= 2
			if current > 9 {
				current -= 9
			}
		}
		sum += current
	}

	return sum%10 == 0, err
}
