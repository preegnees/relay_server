package repo

import (
	"context"

	"relay_server/pkg/logger"
)

// Для хранения информации о хосте
type Host struct {
	ID       string `json:"id" bson:"id"`
	GoOS     string `json:"go_os" bson:"go_os"`
	Kernel   string `json:"kernel" bson:"kernel"`
	Core     string `json:"core" bson:"core"`
	Platform string `json:"platform" bson:"platform"`
	OS       string `json:"os" bson:"os"`
	Hostname string `json:"hostanme" bson:"hostanme"`
	CPUs     string `json:"cpus" bson:"cpus"`
	IP       string `json:"ip" bson:"ip"`
}

type CreateHostDTO struct {
	ID       string `json:"id"`
	GoOS     string `json:"go_os"`
	Kernel   string `json:"kernel"`
	Core     string `json:"core"`
	Platform string `json:"platform"`
	OS       string `json:"os"`
	Hostname string `json:"hostanme"`
	CPUs     string `json:"cpus"`
}

// Для сохранение и взымании информации о хосте
type IInfoHost interface {
	GetIdInfoHost(context.Context, logger.ILogger, string) (*Host, error)
	GetAllInfoHost(context.Context, logger.ILogger) (*[]Host, error)
	SaveInfoHost(context.Context, *Host) error
}
