package config

import (
	"io/ioutil"
	"os"
	"strconv"

	"gopkg.in/yaml.v3"
)

var cfg *Config

type Config struct {
	JWTSecret             string         `yaml:"jwt_secret"`
	JWTRefreshTokenSecret string         `yaml:"jwt_refresh_token_secret"`
	SessionLength         int            `yaml:"session_length"`
	Port                  int            `yaml:"port"`
	FrontendBaseURL       string         `yaml:"frontend_base_url"`
	Database              DatabaseConfig `yaml:"database"`
	PublicRoutes          []string       `yaml:"public_routes"`
	Origins               []string       `yaml:"origins"`
	Mail                  MailConfig     `yaml:"mail"`
	Health                HealthConfig   `yaml:"health"`
	Log                   LogConfig      `yaml:"log"`
	Limiter               LimiterConfig  `yaml:"limiter"`
}

type DatabaseConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbname"`
	SSLMode  string `yaml:"sslmode"`
}

type MailConfig struct {
	SESUsername  string `yaml:"ses_username"`
	SESPassword  string `yaml:"ses_password"`
	SESHost      string `yaml:"ses_host"`
	SESPort      string `yaml:"ses_port"`
	FromName     string `yaml:"from_name"`
	FromEmail    string `yaml:"from_email"`
	TemplatesDir string `yaml:"templates_dir"`
}

type HealthConfig struct {
	EmailTo     string `yaml:"email_to"`
	EmailToName string `yaml:"email_to_name"`
}

type LogConfig struct {
	Path    string `yaml:"path"`
	MaxSize int    `yaml:"max_size"`
	Rotate  int    `yaml:"rotate"`
}

type LimiterConfig struct {
	Rate  int `yaml:"rate"`
	Burst int `yaml:"burst"`
}

func Load(path string) error {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	c := &Config{}
	if err := yaml.Unmarshal(b, c); err != nil {
		return err
	}
	overrideConfigWithEnv(c)
	cfg = c
	return nil
}

func overrideConfigWithEnv(c *Config) {
	if v := os.Getenv("JWT_SECRET"); v != "" {
		c.JWTSecret = v
	}
	if v := os.Getenv("JWT_REFRESH"); v != "" {
		c.JWTRefreshTokenSecret = v
	}
	if v := os.Getenv("SESSION_LENGTH"); v != "" {
		if val, err := strconv.Atoi(v); err == nil {
			c.SessionLength = val
		}
	}
	if v := os.Getenv("PORT"); v != "" {
		if val, err := strconv.Atoi(v); err == nil {
			c.Port = val
		}
	}
	if v := os.Getenv("FRONTEND_BASE_URL"); v != "" {
		c.FrontendBaseURL = v
	}
	if v := os.Getenv("DATABASE_HOST"); v != "" {
		c.Database.Host = v
	}
	if v := os.Getenv("DATABASE_PORT"); v != "" {
		if val, err := strconv.Atoi(v); err == nil {
			c.Database.Port = val
		}
	}
	if v := os.Getenv("DATABASE_USERNAME"); v != "" {
		c.Database.Username = v
	}
	if v := os.Getenv("DATABASE_PASSWORD"); v != "" {
		c.Database.Password = v
	}
	if v := os.Getenv("DATABASE_NAME"); v != "" {
		c.Database.DBName = v
	}
	if v := os.Getenv("DATABASE_SSLMODE"); v != "" {
		c.Database.SSLMode = v
	}
	if v := os.Getenv("ORIGIN"); v != "" {
		c.Origins = []string{v}
	}
	if v := os.Getenv("SMTP_USER_NAME"); v != "" {
		c.Mail.SESUsername = v
	}
	if v := os.Getenv("SMTP_PASSWORD"); v != "" {
		c.Mail.SESPassword = v
	}
	if v := os.Getenv("SMTP_HOST"); v != "" {
		c.Mail.SESHost = v
	}
	if v := os.Getenv("SMTP_PORT"); v != "" {
		c.Mail.SESPort = v
	}
	if v := os.Getenv("SMTP_FROM_NAME"); v != "" {
		c.Mail.FromName = v
	}
	if v := os.Getenv("SMTP_FROM_EMAIL"); v != "" {
		c.Mail.FromEmail = v
	}
	if v := os.Getenv("HEALTH_TO_EMAIL"); v != "" {
		c.Health.EmailTo = v
	}
	if v := os.Getenv("HEALTH_TO_NAME"); v != "" {
		c.Health.EmailToName = v
	}
}

func Get() *Config {
	return cfg
}
