package config

type Config struct {
	LogConfig Log `json:"log_config"`
	Db Db `json:"db"`
}

type Log struct {
	Count int `json:"count"`
	Level string `json:"level"`
	Uint string `json:"uint"`
	RollType string `json:"roll_type"`
	File string `json:"file"`
	Dir string `json:"dir"`
	Size int `json:"size"`
	Compress int `json:"compress"`
}

type Db struct {
	UserName string `json:"user_name"`
	Host string `json:"host"`
	Password string `json:"password"`
	Port int `json:"port"`
	DataBase string `json:"data_base"`
}