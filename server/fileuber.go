package server

type ShareConfig struct {
	InitialPath string
	Path        string
	Username    string
	Password    string
	Port        string
	PID         int
}

func NewShareConfig() *ShareConfig {

	return &ShareConfig{}

}

func (scfg *ShareConfig) ResetShareConfig() {

	scfg.Path = scfg.InitialPath

}

func (scfg *ShareConfig) SetShareConfigPID(pid int) {

	scfg.PID = pid

}

func (scfg *ShareConfig) UpdateShareConfig(path string) {

	scfg.Path = path

}

func (scfg *ShareConfig) SetShareConfig(path string, username string, password string, port string) {

	scfg.InitialPath = path

	scfg.Path = path

	scfg.Username = username

	scfg.Password = password

	scfg.Port = port
}

var ShareCfg = NewShareConfig()