package repo

// Для связывания id клиента и его публичного ключа
type HostKey struct {
	Id  string
	Key string
}

// Для связывания id клиента и токена, который выдал ему сервер
type TokenHost struct {
	Id    string
	Token string
}

// Для сохранение и взымании ключей хостов
type IKeys interface {
	GetHostKey(string) (HostKey, error)
	SaveHostKey(HostKey) error
	GetTokenHost(string) (TokenHost, error)
	SaveTokenHost(TokenHost) error
}

// Сохранение и взывание ключей сервера
type IServerKeys interface {
	GetServerPublicKey() (string, error)
	SaveServerPublicKey(string) error
	GetServerPrivateKey() (string, error)
	SaveServerPrivateKey(string) error
}
