package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/tarm/serial"
)

/*
*
Author: Wang Fei (Kaelan)
institute: VIK
Date: 2023-3-28
Description: This program is used to read the data from the serial port and storage the data to a csv file.
*
*/
func getdata() {
	// open the serial port,the baud rate is 9600
	config := &serial.Config{Name: "/dev/virtual-device", Baud: 9600}
	// open the serial port
	port, err := serial.OpenPort(config)
	//return the error if the serial port is not open
	if err != nil {
		panic(err)
	}
	//when finish the program,close the serial port
	defer port.Close()

	// open the output file,storage the data on a csv file
	f, err := os.Create("temperature-data-sensor.csv")
	//return the error if the file is not open
	if err != nil {
		panic(err)
	}
	//when finish the program,close the file
	defer f.Close()
	//create a csv writer(stream)
	writer := csv.NewWriter(f)
	//when finish the program,flush the data to the file
	defer writer.Flush()

	// write the header,the first column of the csv file is time,the second column is the temperature data
	writer.Write([]string{"time", "data"})

	// read the data from the serial port
	for {
		buf := make([]byte, 128)
		n, err := port.Read(buf)
		if err != nil {
			fmt.Println("failed read data from the sensor device", err)
			break
		}

		// convert the data to string
		data := string(buf[:n])
		data = data[:len(data)-2] // remove the "\r of each line
		// get the current time
		t := time.Now().Format(time.RFC3339)
		// if the data is "END",break the loop
		if data == "END" {
			break
		}
		if _, err := strconv.Atoi(data); err == nil {
			writer.Write([]string{t, data})
		}
	}
}
