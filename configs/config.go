package configs

import (
	"github.com/BurntSushi/toml"
	"os"
)

type Config struct {
}

func loadConfig() {
	var conf = new(Config)
	files := []string{"http.toml", "db.toml", "redis.toml"}
	for _, v := range files {
		data, err := os.ReadFile(v)
		if err != nil {
			// 处理文件读取错误
		}
		if err = toml.Unmarshal(data, &conf); err != nil {
			// 处理解析错误
		}
	}

	// 使用 conf 对象
	// ...
}
