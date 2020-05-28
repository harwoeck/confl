package confl

import "time"

// Confl defines a common interface for loading configuration values. It
// represents all functions for loading something that are present in viper
// by default.
type Confl interface {
	InConfig(key string) bool
	IsSet(key string) bool
	Sub(key string) Confl

	Get(key string) interface{}
	GetBool(key string) bool
	GetDuration(key string) time.Duration
	GetFloat64(key string) float64
	GetInt(key string) int
	GetInt32(key string) int32
	GetInt64(key string) int64
	GetIntSlice(key string) []int
	GetSizeInBytes(key string) uint
	GetString(key string) string
	GetStringMap(key string) map[string]interface{}
	GetStringMapString(key string) map[string]string
	GetStringMapStringSlice(key string) map[string][]string
	GetStringSlice(key string) []string
	GetTime(key string) time.Time
	GetUint(key string) uint
	GetUint32(key string) uint32
	GetUint64(key string) uint64
}

// WithDefaults extends Confl be adding *Default functions that return the
// default value, when the key is not set (e.g. IsSet would return false).
type WithDefaults interface {
	Confl
	GetDefault(key string, val interface{}) interface{}
	GetBoolDefault(key string, val bool) bool
	GetDurationDefault(key string, val time.Duration) time.Duration
	GetFloat64Default(key string, val float64) float64
	GetIntDefault(key string, val int) int
	GetInt32Default(key string, val int32) int32
	GetInt64Default(key string, val int64) int64
	GetIntSliceDefault(key string, val []int) []int
	GetSizeInBytesDefault(key string, val uint) uint
	GetStringDefault(key string, val string) string
	GetStringMapDefault(key string, val map[string]interface{}) map[string]interface{}
	GetStringMapStringDefault(key string, val map[string]string) map[string]string
	GetStringMapStringSliceDefault(key string, val map[string][]string) map[string][]string
	GetStringSliceDefault(key string, val []string) []string
	GetTimeDefault(key string, val time.Time) time.Time
	GetUintDefault(key string, val uint) uint
	GetUint32Default(key string, val uint32) uint32
	GetUint64Default(key string, val uint64) uint64
}
