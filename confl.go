package confl

import "time"

// Confl defines a common interface for loading ("getting") configuration
// values in various go types. The available functions and naming schema is
// heavily inspired by viper (https://github.com/spf13/viper), though
// additional *Default functions are included for every type. These return the
// default value, when the key is not set in the config source (e.g. IsSet
// returns false).
//
// An implementation that uses viper as a config source can be found under
//
type Confl interface {
	Source() interface{}

	InConfig(key string) bool
	IsSet(key string) bool
	Sub(key string) Confl

	Get(key string) interface{}
	GetDefault(key string, val interface{}) interface{}
	GetBool(key string) bool
	GetBoolDefault(key string, val bool) bool
	GetDuration(key string) time.Duration
	GetDurationDefault(key string, val time.Duration) time.Duration
	GetFloat64(key string) float64
	GetFloat64Default(key string, val float64) float64
	GetInt(key string) int
	GetIntDefault(key string, val int) int
	GetInt32(key string) int32
	GetInt32Default(key string, val int32) int32
	GetInt64(key string) int64
	GetInt64Default(key string, val int64) int64
	GetIntSlice(key string) []int
	GetIntSliceDefault(key string, val []int) []int
	GetSizeInBytes(key string) uint
	GetSizeInBytesDefault(key string, val uint) uint
	GetString(key string) string
	GetStringDefault(key string, val string) string
	GetStringMap(key string) map[string]interface{}
	GetStringMapDefault(key string, val map[string]interface{}) map[string]interface{}
	GetStringMapString(key string) map[string]string
	GetStringMapStringDefault(key string, val map[string]string) map[string]string
	GetStringMapStringSlice(key string) map[string][]string
	GetStringMapStringSliceDefault(key string, val map[string][]string) map[string][]string
	GetStringSlice(key string) []string
	GetStringSliceDefault(key string, val []string) []string
	GetTime(key string) time.Time
	GetTimeDefault(key string, val time.Time) time.Time
	GetUint(key string) uint
	GetUintDefault(key string, val uint) uint
	GetUint32(key string) uint32
	GetUint32Default(key string, val uint32) uint32
	GetUint64(key string) uint64
	GetUint64Default(key string, val uint64) uint64
}
