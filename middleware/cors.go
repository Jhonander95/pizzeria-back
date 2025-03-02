/*
	package middleware

import (

	//"net/http"

	"github.com/gin-gonic/gin"

)

	func CORSMiddleware() gin.HandlerFunc {
		return func(c *gin.Context) {
			// Permite cualquier origen en desarrollo
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

			// Añade más cabeceras permitidas
			c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With, Accept")
			c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
			c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")

			// Manejo de pre-flight OPTIONS
			if c.Request.Method == "OPTIONS" {
				c.AbortWithStatus(204)
				return
			}

			c.Next()
		}
	}
*/
package middleware

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// CORSConfig contiene la configuración para el middleware CORS
type CORSConfig struct {
	AllowOrigins     []string
	AllowMethods     []string
	AllowHeaders     []string
	ExposeHeaders    []string
	AllowCredentials bool
	MaxAge           time.Duration
}

// DefaultCORSConfig proporciona una configuración predeterminada
func DefaultCORSConfig() *CORSConfig {
	return &CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization", "X-Requested-With", "Accept", "Origin"},
		ExposeHeaders:    []string{"Content-Length", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}
}

// DevelopmentCORSConfig proporciona una configuración para desarrollo
func DevelopmentCORSConfig() *CORSConfig {
	config := DefaultCORSConfig()
	config.MaxAge = 5 * time.Second // Tiempo de caché corto para desarrollo
	return config
}

// ProductionCORSConfig proporciona una configuración para producción
func ProductionCORSConfig(allowedOrigins []string) *CORSConfig {
	config := DefaultCORSConfig()
	config.AllowOrigins = allowedOrigins // En producción, especifica los orígenes permitidos
	return config
}

// CORSMiddleware crea un middleware CORS configurable
func CORSMiddleware(config *CORSConfig) gin.HandlerFunc {
	if config == nil {
		config = DefaultCORSConfig()
	}

	return func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin")

		// Establecer Allow-Origin basado en la configuración
		if len(config.AllowOrigins) == 1 && config.AllowOrigins[0] == "*" {
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		} else {
			// Comprueba si el origen está permitido
			allowOrigin := false
			for _, allowedOrigin := range config.AllowOrigins {
				if origin == allowedOrigin {
					allowOrigin = true
					break
				}
			}

			if allowOrigin {
				c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
			} else {
				// Si el origen no está permitido, continúa sin configurar CORS
				c.Next()
				return
			}
		}

		// Configurar otras cabeceras CORS
		c.Writer.Header().Set("Access-Control-Allow-Methods", joinStrings(config.AllowMethods))
		c.Writer.Header().Set("Access-Control-Allow-Headers", joinStrings(config.AllowHeaders))
		c.Writer.Header().Set("Access-Control-Expose-Headers", joinStrings(config.ExposeHeaders))

		if config.AllowCredentials {
			c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		}

		if config.MaxAge > 0 {
			c.Writer.Header().Set("Access-Control-Max-Age", secondsToString(config.MaxAge.Seconds()))
		}

		// Manejo de pre-flight OPTIONS
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}

// Función auxiliar para unir strings con comas
func joinStrings(strings []string) string {
	if len(strings) == 0 {
		return ""
	}

	joined := strings[0]
	for i := 1; i < len(strings); i++ {
		joined += ", " + strings[i]
	}
	return joined
}

// Convierte segundos a string
func secondsToString(seconds float64) string {
	return string(int(seconds))
}
