package main

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"syscall/js"
)

func gzipCompress(this js.Value, args []js.Value) interface{} {
	if len(args) == 0 {
		return "Error: No input provided"
	}

	input := args[0].String()
	var buf bytes.Buffer
	gzipWriter := gzip.NewWriter(&buf)

	_, err := gzipWriter.Write([]byte(input))
	if err != nil {
		return "Error: " + err.Error()
	}

	err = gzipWriter.Close()
	if err != nil {
		return "Error: " + err.Error()
	}

	compressed := base64.StdEncoding.EncodeToString(buf.Bytes())
	return compressed
}

func gunzipCompress(this js.Value, args []js.Value) interface{} {
	if len(args) == 0 {
		return "Error: No input provided"
	}

	input := args[0].String()
	decoded, err := base64.StdEncoding.DecodeString(input)
	if err != nil {
		return "Error: " + err.Error()
	}

	buf := bytes.NewBuffer(decoded)
	gzipReader, err := gzip.NewReader(buf)
	if err != nil {
		return "Error: " + err.Error()
	}

	var out bytes.Buffer
	_, err = out.ReadFrom(gzipReader)
	if err != nil {
		return "Error: " + err.Error()
	}

	return out.String()
}

func main() {
	c := make(chan struct{})
	js.Global().Set("gzipCompress", js.FuncOf(gzipCompress))
	js.Global().Set("gunzipCompress", js.FuncOf(gunzipCompress))
	<-c

	fmt.Println("WASM Go Initialized")
}
