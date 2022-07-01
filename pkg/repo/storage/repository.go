package repo

// Для хранения информации о хосте
type Host struct {
	ID       string
	IP       string
	GoOS     string
	Kernel   string
	Core     string
	Platform string
	OS       string
	Hostname string
	CPUs     string
}

// Для сохранение и взымании информации о хосте
type IInfoHost interface {
	GetInfoHost(string) (Host, error)
	SaveInfoHost(Host) error
}
