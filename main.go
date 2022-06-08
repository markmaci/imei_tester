package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

type parsedIMEI struct {
	// define parsedIMEI struct

	typeAllocationCode string
	serialNumber       string
	checksum           int
}

func check(e error) {
	// error handling
	if e != nil {
		panic(e)
	}
}

func main() {
	// read in from file
	pwd, err := os.Getwd()
	filepath := filepath.Join(pwd + "/testdata.txt")
	file, _ := os.Open(filepath)

	if err != nil {
		fmt.Print("error with data file")
		fmt.Print(err)
	}

	defer file.Close()
	// create array of imeis on each line using scanner
	var entries []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		entries = append(entries, scanner.Text())
	}

	// create output file
	newfile, err := os.Create("parsed.txt")

	for i := 0; i < len(entries); i++ {

		// run parse and validate functions on input
		parsed, err1 := (parseIMEI(entries[i]))

		validity, err2 := (validateChecksum(entries[i]))

		iString := strconv.Itoa(i)
		checksumString := strconv.Itoa(parsed.checksum)

		newfile.WriteString("Parsing IMEI on line " + iString + ", " + entries[i] + "\n")

		// if there are no errors, write the parsed imei to file
		if err1 == nil && err2 == nil {

			newfile.WriteString("TAC:" + parsed.typeAllocationCode + "\n")

			newfile.WriteString("Serial Number:" + parsed.serialNumber + "\n")

			newfile.WriteString("Checksum:" + checksumString + "\n")

			newfile.WriteString("\n")
		}

		// write the relavent error message
		if validity == false && err1 != nil {
			newfile.WriteString(err1.Error() + "\n")
			newfile.WriteString("\n")
		} else if err2 != nil {
			newfile.WriteString(err2.Error() + "\n")
			newfile.WriteString("\n")
		}
	}

}

func parseIMEI(IMEI string) (parsed parsedIMEI, err error) {
	// length check
	length := len(IMEI)
	if length != 15 {
		err = errors.New("error = invalid length")
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

			return false, errors.New("error = invalid IMEI")

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
	err = errors.New("error = invalid checksum")
	return false, err

}
