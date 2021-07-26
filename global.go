package env

import "time"

var defaultEnv Env

func GetString(key string) string {
	return defaultEnv.GetString(key)
}

func GetStringDefault(key, defaultValue string) string {
	return defaultEnv.GetStringDefault(key, defaultValue)
}

func GetInt(key string) (int, error) {
	return defaultEnv.GetInt(key)
}

func GetIntDefault(key string, defaultValue int) (int, error) {
	return defaultEnv.GetIntDefault(key, defaultValue)
}

func GetInt64(key string) (int64, error) {
	return defaultEnv.GetInt64(key)
}

func GetInt64Default(key string, defaultValue int64) (int64, error) {
	return defaultEnv.GetInt64Default(key, defaultValue)
}

func GetUint(key string) (uint, error) {
	return defaultEnv.GetUint(key)
}

func GetUintDefault(key string, defaultValue uint) (uint, error) {
	return defaultEnv.GetUintDefault(key, defaultValue)
}

func GetUint64(key string) (uint64, error) {
	return defaultEnv.GetUint64(key)
}

func GetUint64Default(key string, defaultValue uint64) (uint64, error) {
	return defaultEnv.GetUint64Default(key, defaultValue)
}

func GetFloat(key string) (float64, error) {
	return defaultEnv.GetFloat(key)
}

func GetFloatDefault(key string, defaultValue float64) (float64, error) {
	return defaultEnv.GetFloatDefault(key, defaultValue)
}

func GetBool(key string) (bool, error) {
	return defaultEnv.GetBool(key)
}

func GetBoolDefault(key string, defaultValue bool) (bool, error) {
	return defaultEnv.GetBoolDefault(key, defaultValue)
}

func GetDuration(key string) (time.Duration, error) {
	return defaultEnv.GetDuration(key)
}

func GetDurationDefault(key string, defaultValue time.Duration) (time.Duration, error) {
	return defaultEnv.GetDurationDefault(key, defaultValue)
}
