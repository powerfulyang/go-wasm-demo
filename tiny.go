package main

import (
	"bytes"
	"fmt"
	"github.com/klauspost/compress/gzip"
	"io"
	"log"
)

func main() {
	// 原始数据
	data := []byte("Hello, TinyGo Gzip!")

	// 压缩数据
	var compressedData bytes.Buffer
	gzWriter := gzip.NewWriter(&compressedData)
	_, err := gzWriter.Write(data)
	if err != nil {
		log.Fatalf("Failed to write compressed data: %v", err)
	}
	err = gzWriter.Close()
	if err != nil {
		log.Fatalf("Failed to close gzip writer: %v", err)
	}

	fmt.Println("Compressed Data:", compressedData.Bytes())

	// 解压缩数据
	gzReader, err := gzip.NewReader(&compressedData)
	if err != nil {
		log.Fatalf("Failed to create gzip reader: %v", err)
	}
	defer func() {
		if err := gzReader.Close(); err != nil {
			log.Printf("Warning: Failed to close gzip reader: %v", err)
		}
	}()

	decompressedData, err := io.ReadAll(gzReader)
	if err != nil {
		log.Fatalf("Failed to read decompressed data: %v", err)
	}

	fmt.Println("Decompressed Data:", string(decompressedData))
}
