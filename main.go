package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func replaceBytesInFile(filePath string, oldBytes, newBytes []byte) error {
	// Read the file in binary mode
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to read file: %v", err)
	}

	// Search for the byte sequence to replace
	index := bytes.Index(data, oldBytes)
	if index == -1 {
		return fmt.Errorf("sequence not found in the file")
	}

	// Replace the found sequence with the new sequence
	copy(data[index:index+len(newBytes)], newBytes)

	// Write the modified file
	err = ioutil.WriteFile(filePath, data, os.ModePerm)
	if err != nil {
		return fmt.Errorf("failed to write file: %v", err)
	}

	return nil
}

func main() {
	// Path to the .pak file
	filePath := "C:\\Program Files\\Epic Games\\Fortnite\\FortniteGame\\Content\\Paks\\pakchunk10-WindowsClient_s2.ucas"

	// Byte sequence to replace
	oldBytes := []byte{
		0x41, 0x00, 0x11, 0x98, 0x90, 0x02, 0x33, 0x00, 0x00, 0x09, 0x52, 0x00, 0xF1, 0xFF, 0x99, 0x40,
		0x50, 0x00, 0x40, 0x10, 0x00, 0x40, 0x58, 0x00, 0x49, 0xEF, 0xB8, 0xDF, 0x76, 0x08, 0x83, 0x3C,
		0x33, 0x0B, 0x41, 0x00, 0xF0, 0xB9, 0xF1, 0x01, 0x00, 0x14, 0x00, 0x80, 0x13, 0x00, 0xF6, 0x01,
		0x02, 0x04, 0x02, 0x08, 0x06, 0x86, 0x4F, 0x01, 0x02, 0x02, 0x05, 0x02, 0x08, 0x04, 0x03, 0x02,
		0x32, 0x02, 0x16, 0x04, 0x02, 0x06, 0x01, 0x03, 0x39, 0x00, 0x91, 0x05, 0x42, 0x00, 0x12, 0x74,
		0xC2, 0x0F, 0x00, 0x02, 0x01, 0x01, 0xA4, 0x00, 0x40, 0x5D, 0x00, 0xF2, 0x06, 0x00, 0xB1, 0xF0,
		0x05, 0x18, 0x00, 0x40, 0x50, 0x00, 0x61, 0x04, 0x2B, 0x00, 0x10, 0x2C, 0x00, 0x31, 0x07, 0x10,
		0x00, 0x81, 0x01, 0x08, 0x00, 0x2F, 0x21, 0x21, 0x00, 0x00, 0x00, 0x41, 0x45, 0x39, 0x35, 0x35,
		0x38, 0x37, 0x33, 0x34, 0x42, 0x38, 0x44, 0x39, 0x44, 0x46, 0x33, 0x37, 0x45, 0x46, 0x42, 0x32,
		0x35, 0x39, 0x31, 0x42, 0x37, 0x35, 0x35, 0x45, 0x41, 0x34, 0x41, 0x00, 0x08, 0x00, 0x00, 0x00,
	}

	// New byte sequence
	newBytes := []byte{
		0x41, 0x00, 0x11, 0x98, 0x90, 0x02, 0x33, 0x00, 0x00, 0x09, 0x52, 0x00, 0xF1, 0xFF, 0x99, 0x40,
		0x50, 0x00, 0x40, 0x10, 0x00, 0x40, 0x58, 0x00, 0x49, 0xEF, 0xB8, 0xDF, 0x76, 0x08, 0x83, 0x3C,
		0x33, 0x0B, 0x41, 0x00, 0xF0, 0xB9, 0xF1, 0x01, 0x00, 0x14, 0x00, 0x80, 0x13, 0x00, 0xF6, 0x01,
		0x02, 0x04, 0x02, 0x08, 0x06, 0x86, 0x4F, 0x01, 0x02, 0x02, 0x05, 0x02, 0x08, 0x04, 0x03, 0x02,
		0x32, 0x02, 0x16, 0x04, 0x02, 0x06, 0x01, 0x03, 0x39, 0x00, 0x91, 0x05, 0x42, 0x00, 0x12, 0xAF,
		0xC3, 0x0F, 0x00, 0x02, 0x01, 0x01, 0xA4, 0x00, 0x40, 0x5D, 0x00, 0xF2, 0x06, 0x00, 0xB1, 0xF0,
		0x05, 0x18, 0x00, 0x40, 0x50, 0x00, 0x61, 0x04, 0x2B, 0x00, 0x10, 0x2C, 0x00, 0x31, 0x07, 0x10,
		0x00, 0x81, 0x01, 0x08, 0x00, 0x2F, 0x21, 0x21, 0x00, 0x00, 0x00, 0x41, 0x45, 0x39, 0x35, 0x35,
		0x38, 0x37, 0x33, 0x34, 0x42, 0x38, 0x44, 0x39, 0x44, 0x46, 0x33, 0x37, 0x45, 0x46, 0x42, 0x32,
		0x35, 0x39, 0x31, 0x42, 0x37, 0x35, 0x35, 0x45, 0x41, 0x34, 0x41, 0x00, 0x08, 0x00, 0x00, 0x00,
	}

	err := replaceBytesInFile(filePath, oldBytes, newBytes)
	if err != nil {
		log.Fatalf("Error replacing bytes: %v", err)
	}

	fmt.Println("Bytes replaced successfully!")
}
