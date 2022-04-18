package g

type GlobalConfig struct {
	Debug         bool              `json:"debug"`
	LogLevel   		string			`json:"loglevel"`
	LogFile 		string			`json:"logfile"`
	LogPath 		string 			`json:"logpath"`
        LogMaxAge               int                     `json:"logmaxage"`
        LogRotateAge  		int                     `json:"logrotateage"`

}

type Fields map[string]fieldFn

type fieldFn func() interface{}
