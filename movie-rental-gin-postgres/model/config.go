package model

type Config struct {
	InitFile
	DbConfig
	ServerConfig
}

type DbConfig struct {
	Driver  string
	ConnStr string
}

type InitFile struct {
	Path string
}

type ServerConfig struct {
	Address string
}
