package service

import (
	"flag"
	"fmt"
	"math/rand"
	"os"

	"github.com/Deadrafa/Multithreaded-file-handling/internal/config"
	"github.com/Deadrafa/Multithreaded-file-handling/internal/models"
	"gopkg.in/yaml.v3"
)

func ParseFlags() models.Config {
	numFiles := flag.Int("files", 10, "number of files")
	iterations := flag.Int("iter", 100, "number of iterations")
	workers := flag.Int("workers", 20, "number of workers")
	timeout := flag.Duration("timeout", 0, "execution timeout")
	flag.Parse()

	return models.Config{
		NumFiles:   *numFiles,
		Iterations: *iterations,
		Workers:    *workers,
		Timeout:    *timeout,
	}
}

func writeFile(filename string, data *models.Data) error {
	yamlData, err := yaml.Marshal(data)
	if err != nil {
		return err
	}
	return os.WriteFile(filename, yamlData, 0644)
}

func generateID() string {
	b := make([]byte, 16)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}

func generateText(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = config.Letters[rand.Intn(len(config.Letters))]
	}
	return string(b)
}
