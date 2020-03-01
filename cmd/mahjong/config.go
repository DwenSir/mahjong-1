package main

type Config struct {
	LogConfig Log `json:"log_config"`
	DbConfig  Db  `json:"db_config"`
	WebConfig Web `json:"web_config"`
}

type Log struct {
	Count    int32  `json:"count"`
	Level    string `json:"level"`
	Unit     string `json:"unit"`
	RollType string `json:"roll_type"`
	File     string `json:"file"`
	Dir      string `json:"dir"`
	Size     int64  `json:"size"`
	Compress int64  `json:"compress"`
}

type Db struct {
	UserName string `json:"user_name"`
	Host     string `json:"host"`
	Password string `json:"password"`
	Port     int    `json:"port"`
	DataBase string `json:"data_base"`
}

type Web struct {
	HTTPS        bool   `json:"https"`
	IP           string `json:"ip"`
	Port         int    `json:"port"`
	Mode         string `json:"mode"`
	ReadTimeout  uint64 `json:"read_timeout"`
	WriteTimeout uint64 `json:"write_timeout"`
}
