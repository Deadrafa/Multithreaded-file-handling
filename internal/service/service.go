package service

import (
	"fmt"
	"math/rand"
	"os"
	"sync"
	"time"

	"github.com/Deadrafa/Multithreaded-file-handling/internal/models"

	"gopkg.in/yaml.v3"
)

func Worker(id int, cfg models.Config, tasks <-chan int, mutexes []sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done()

	for range tasks {
		fileIdx := rand.Intn(cfg.NumFiles)
		filename := fmt.Sprintf("data/file_%d.yaml", fileIdx)

		start := time.Now()
		mutexes[fileIdx].Lock()

		data, err := readOrCreateFile(filename)
		if err != nil {
			fmt.Printf("Worker %d error: %v\n", id, err)
			mutexes[fileIdx].Unlock()
			continue
		}

		// Обновление данных
		data.UpdatedAt = time.Now()
		data.Chain = append(data.Chain, id)
		data.Text = generateText(20)

		// Расчет времени выполнения
		latency := time.Since(start)
		if data.Latencies == nil {
			data.Latencies = make(map[int]time.Duration)
		}
		data.Latencies[id] = latency

		// Запись файла
		if err := writeFile(filename, data); err != nil {
			fmt.Printf("Worker %d write error: %v\n", id, err)
		}

		mutexes[fileIdx].Unlock()
	}
}

func readOrCreateFile(filename string) (*models.Data, error) {
	data := &models.Data{}

	file, err := os.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			// Создаем новые данные
			return &models.Data{
				ID:        generateID(),
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
				Chain:     []int{},
				Latencies: make(map[int]time.Duration),
				Text:      generateText(30),
			}, nil
		}
		return nil, err
	}

	if err := yaml.Unmarshal(file, data); err != nil {
		return nil, err
	}
	return data, nil
}
