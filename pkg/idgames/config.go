package idgames

const DefaultApiUrl string = "https://www.doomworld.com/idgames/api/api.php"

type Config struct {
	ApiUrl string
}

func (c *Config) apiUrl() string {
	if c == nil || c.ApiUrl != "" {
		return c.ApiUrl
	}
	return DefaultApiUrl
}
