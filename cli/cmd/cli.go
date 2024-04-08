package main

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"os"
	"path"
	"regexp"
	"strings"

	"github.com/andrescosta/goico/pkg/runtimes/wasm"
)

func main() {
	var dump bool
	if os.Args[2] == "dump" {
		dump = true
	}
	input := os.Args[3]
	dirTest := os.Args[1]
	ctx := context.Background()
	runtime, err := wasm.NewRuntimeWithCompilationCache("./cache")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error initializing Wazero: %v\n", err)
		os.Exit(1)
	}
	defer runtime.Close(ctx)
	wasmf, err := os.ReadFile(path.Join(dirTest, "/js.wasm"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading wasm binary: %v\n", err)
		os.Exit(1)
	}
	mounts := []string{
		dirTest + ":/hello",
		dirTest + "/sdk:/hello/sdk",
	}
	args := []string{
		path.Join(dirTest, "/js.wasm"),
		"--module=/hello/test.js",
	}
	buffIn := &bytes.Buffer{}
	buffOut := &bytes.Buffer{}
	buffErr := &bytes.Buffer{}
	r, err := regexp.Compile("\r\n|\n|\r")
	if err != nil {
		print(err)
		os.Exit(1)
	}
	v := r.ReplaceAllString(input, "")
	buffIn.WriteString(v)
	modi, err := wasm.NewIntModule(ctx, runtime, wasmf, Log, mounts, args, []wasm.EnvVar{}, buffIn, buffOut, buffErr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error instantiating the module: %v\n", err)
		fmt.Printf("Std Error: %s\n", buffErr.String())
		fmt.Printf("Std out: %s\n", buffOut.String())
		os.Exit(1)
	}
	defer modi.Close(ctx)
	err = modi.Run(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error instantiating the module: %v\n", err)
		fmt.Printf("Std Error: %s\n", buffErr.String())
		fmt.Printf("Std out: %s\n", buffOut.String())
		os.Exit(1)
	}
	if dump {
		fmt.Printf("Std Error: %s\n", buffErr.String())
		fmt.Printf("Std out: %s\n", buffOut.String())
		os.Exit(1)
	}
	scErr := bufio.NewScanner(buffErr)
	for scErr.Scan() {
		msgErr := scErr.Text()
		fmt.Printf("Level:%s\n", string(msgErr[0]))
		fmt.Printf("Text:%s\n", msgErr[1:])
	}
	scOut := bufio.NewScanner(buffOut)
	if scOut.Scan() {
		msgOut := scOut.Text()
		fmt.Printf("Result:%s\n", strings.TrimSpace(msgOut[0:11]))
		fmt.Printf("Text:%s\n", msgOut[11:])
	}
}

func Log(context.Context, uint32, string) error {
	return nil
}
