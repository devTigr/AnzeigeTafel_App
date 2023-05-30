package models

import (
	"encoding/csv"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

// FileName declares how the File is called, that is used to get the data from
const FileName = "data.csv"

// DelimiterInDataStorage defines how the values are seperated
const DelimiterInDataStorage = ","

// Data stored in File or in Program
var filePersistence bool = false

// EnableFilePersistence enables the file persistence
func EnableFilePersistence() {
	filePersistence = true
}

// DisableFilePersistence disables the file persistence
func DisableFilePersistence() {
	filePersistence = false
}

// define a type for data handling
var samples []Sample

// Sample as type
type Sample struct {
	Charge      int
	Tank        int
	Arbeitsgang string
	Status      int
}

// Initialize does the initialization of the repository
func Initialize() error {
	if filePersistence {
		var err error
		samples, err = getDataFromFile()
		if err != nil {
			return err
		}
	}

	return nil
}

func getDataFromFile() ([]Sample, error) {
	file, err := os.Open(FileName)
	if err != nil {
		return nil, err
	}

	csvReader := csv.NewReader(file)
	numberOfFieldsInSampleStruct := reflect.TypeOf(Sample{}).NumField()
	csvReader.FieldsPerRecord = numberOfFieldsInSampleStruct
	records, err := csvReader.ReadAll()

	if err != nil {
		file.Close()
		return nil, err
	}

	var readSamples []Sample
	for _, record := range records {
		parsedSample, err := parseSampleFromStringList(record)

		if err != nil {
			file.Close()
			return nil, err
		}

		readSamples = append(readSamples, parsedSample) // Add parsedSample to slice
	}

	file.Close()

	return readSamples, nil
}

func parseSampleFromStringList(record []string) (Sample, error) {
	parsedSample := Sample{}
	numberOfReceivedElements := len(record)
	numberOfFieldsInSampleStruct := reflect.TypeOf(Sample{}).NumField()

	if numberOfReceivedElements != numberOfFieldsInSampleStruct {
		err := fmt.Errorf("data record does not contain enough elements for parsing: received %d, expected: %d",
			numberOfReceivedElements, numberOfFieldsInSampleStruct)
		return parsedSample, err
	}

	// Create new Sample based on parsed values
	//
	parsedSample = Sample{
		Charge:      StringToInt(record[0]),
		Tank:        StringToInt(record[1]),
		Arbeitsgang: record[2],
		Status:      StringToInt(record[3])}

	return parsedSample, nil
}

func GetAllSamples() []Sample {
	allSamples := make([]Sample, len(samples))
	copy(allSamples, samples)
	return allSamples
}

func getSampleAsStringSlice(sample Sample) []string {
	sampleSerialized := []string{
		string(sample.Charge),
		string(sample.Tank),
		sample.Arbeitsgang,
		string(sample.Status)}

	return sampleSerialized
}

// StringToInt converts a string to an integer value
func StringToInt(info string) int {
	infoTrimmed := strings.TrimSpace(info)
	aInt, _ := strconv.Atoi(infoTrimmed)
	return aInt
}

// IntToString converts an int to a string value
func IntToString(info int) string {
	aString := strconv.Itoa(info)
	return aString
}
