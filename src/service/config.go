package service

type Config struct {
	Storage struct {
		RedisURL string `json:"redis_url"`
	} `json:"storage"`
	ServeOn struct {
		Addr string `json:"addr"`
		Port string `json:"port"`
	} `json:"serve_on"`
}

func ConfigFromFile(path string) *Config {
	var c Config
	MustUnmarshal(MustReadFile(path), &c)
	return &c
}
