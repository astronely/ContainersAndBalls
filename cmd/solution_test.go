package main

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/stretchr/testify/require"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestSolveWithSort(t *testing.T) {
	testPaths := []string{"yes", "no"}

	for _, dir := range testPaths {
		files, err := filepath.Glob(filepath.Join("../tests", dir, "input_*.txt"))
		if err != nil {
			t.Fatalf("Failed to read test files in %s: %v", dir, err)
		}

		for _, file := range files {
			t.Run(fmt.Sprintf("%s/%s", dir, filepath.Base(file)), func(t *testing.T) {
				data, err := os.ReadFile(file)
				if err != nil {
					t.Fatalf("Failed to read file %s: %v", file, err)
				}

				in = bufio.NewReader(bytes.NewReader(data))

				expected := strings.ToLower(dir)
				have := solveWithSort()

				require.Equal(t, expected, have, "Test %s failed", file)
			})
		}
	}
}

func TestSolveWithMap(t *testing.T) {
	testPaths := []string{"yes", "no"}

	for _, dir := range testPaths {
		files, err := filepath.Glob(filepath.Join("../tests", dir, "input_*.txt"))
		if err != nil {
			t.Fatalf("Failed to read test files in %s: %v", dir, err)
		}

		for _, file := range files {
			t.Run(fmt.Sprintf("%s/%s", dir, filepath.Base(file)), func(t *testing.T) {
				data, err := os.ReadFile(file)
				if err != nil {
					t.Fatalf("Failed to read file %s: %v", file, err)
				}

				in = bufio.NewReader(bytes.NewReader(data))

				expected := strings.ToLower(dir)
				have := solveWithMap()

				require.Equal(t, expected, have, "Test %s failed", file)
			})
		}
	}
}
