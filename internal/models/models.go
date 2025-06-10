package models

import "time"

type Data struct {
	ID        string                `yaml:"id"`
	CreatedAt time.Time             `yaml:"created_at"`
	UpdatedAt time.Time             `yaml:"updated_at"`
	Chain     []int                 `yaml:"chain"`
	Latencies map[int]time.Duration `yaml:"latencies"`
	Text      string                `yaml:"text"`
}
