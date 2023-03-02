package config

import "gohub/pkg/config"

func init() {
	config.Add("jwt", func() map[string]interface{} {
		return map[string]interface{}{
			// "signing_key": config.GetString("app.key"),
			"expire_time":       config.Env("JWT_EXPIRE_TIME", 120),
			"max_refresh_time":  config.Env("JWT_MAX_REFRESH_TIME", 86400),
			"debug_expire_time": 86400,
		}
	})
}
