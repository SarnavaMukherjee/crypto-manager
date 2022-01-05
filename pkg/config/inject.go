/**
 * File: inject.go
 * Author: Kedar Kekan
 * Contact: (support@airavana.ai)
 * Copyright (c) 2020 - 2021 Airavana Inc.
 */

package config

import "github.com/gin-gonic/gin"

func Inject(cfg *Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("cfg", cfg)
		c.Next()
	}
}
