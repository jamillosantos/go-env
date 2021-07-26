package env

import (
	"os"
	"strconv"
	"time"
)

type Environment interface {
	WithPrefix(prefix string) Environment

	GetString(key string) string
	GetStringDefault(key, defaultValue string) string

	GetInt(key string) (int, error)
	GetIntDefault(key string, defaultValue int) (int, error)

	GetInt64(key string) (int64, error)
	GetInt64Default(key string, defaultValue int64) (int64, error)

	GetUint(key string) (uint, error)
	GetUintDefault(key string, defaultValue uint) (uint, error)

	GetUint64(key string) (uint64, error)
	GetUint64Default(key string, defaultValue uint64) (uint64, error)

	GetFloat(key string) (float64, error)
	GetFloatDefault(key string, defaultValue float64) (float64, error)

	GetBool(key string) (bool, error)
	GetBoolDefault(key string, defaultValue bool) (bool, error)

	GetDuration(key string) (time.Duration, error)
	GetDurationDefault(key string, defaultValue time.Duration) (time.Duration, error)
}

type Env struct {
	prefix    string
	overrides map[string]string
}

func NewEnv() Env {
	return Env{}
}

func (env *Env) key(key string) string {
	return key
}

func (env Env) WithPrefix(prefix string) Env {
	return Env{
		prefix:    prefix,
		overrides: env.overrides,
	}
}

func (env Env) WithOverrides(overrides map[string]string) Env {
	return Env{
		prefix:    env.prefix,
		overrides: overrides,
	}
}

func (env *Env) getValue(key string) string {
	return os.Getenv(env.key(key))
}

func (env *Env) GetString(key string) string {
	return env.getValue(key)
}

func (env *Env) GetStringDefault(key string, defaultValue string) string {
	value := env.getValue(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func (env *Env) GetInt(key string) (int, error) {
	value := env.getValue(key)
	n, err := strconv.ParseInt(value, 10, 32)
	if err != nil {
		return 0, err
	}
	return int(n), nil
}

func (env *Env) GetIntDefault(key string, defaultValue int) (int, error) {
	value := env.getValue(key)
	if value == "" {
		return defaultValue, nil
	}
	n, err := strconv.ParseInt(value, 10, 32)
	if err != nil {
		return defaultValue, err
	}
	return int(n), nil
}

func (env *Env) GetInt64(key string) (int64, error) {
	value := env.getValue(key)
	return strconv.ParseInt(value, 10, 64)
}

func (env *Env) GetInt64Default(key string, defaultValue int64) (int64, error) {
	value := env.getValue(key)
	if value == "" {
		return defaultValue, nil
	}
	n, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return defaultValue, err
	}
	return n, nil
}

func (env *Env) GetUint(key string) (uint, error) {
	value := env.getValue(key)
	n, err := strconv.ParseUint(value, 10, 32)
	if err != nil {
		return 0, err
	}
	return uint(n), nil
}

func (env *Env) GetUintDefault(key string, defaultValue uint) (uint, error) {
	value := env.getValue(key)
	if value == "" {
		return defaultValue, nil
	}
	n, err := strconv.ParseUint(value, 10, 32)
	if err != nil {
		return defaultValue, err
	}
	return uint(n), nil
}

func (env *Env) GetUint64(key string) (uint64, error) {
	value := env.getValue(key)
	return strconv.ParseUint(value, 10, 64)
}

func (env *Env) GetUint64Default(key string, defaultValue uint64) (uint64, error) {
	value := env.getValue(key)
	if value == "" {
		return defaultValue, nil
	}
	n, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		return defaultValue, err
	}
	return n, nil
}

func (env *Env) GetFloat(key string) (float64, error) {
	value := env.getValue(key)
	return strconv.ParseFloat(value, 64)
}

func (env *Env) GetFloatDefault(key string, defaultValue float64) (float64, error) {
	value := env.getValue(key)
	if value == "" {
		return defaultValue, nil
	}
	n, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return defaultValue, err
	}
	return n, nil
}

func (env *Env) GetBool(key string) (bool, error) {
	value := env.getValue(key)
	if value == "" {
		return false, nil
	}
	return strconv.ParseBool(value)
}

func (env *Env) GetBoolDefault(key string, defaultValue bool) (bool, error) {
	value := env.getValue(key)
	if value == "" {
		return defaultValue, nil
	}
	n, err := strconv.ParseBool(value)
	if err != nil {
		return defaultValue, err
	}
	return n, nil
}

func (env *Env) GetDuration(key string) (time.Duration, error) {
	value := env.getValue(key)
	return time.ParseDuration(value)
}

func (env *Env) GetDurationDefault(key string, defaultValue time.Duration) (time.Duration, error) {
	value := env.getValue(key)
	if value == "" {
		return defaultValue, nil
	}
	n, err := time.ParseDuration(value)
	if err != nil {
		return defaultValue, err
	}
	return n, nil
}
