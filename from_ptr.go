package ptr

import (
	"time"
)

// ToBool returns bool value dereferenced if the passed
// in pointer was not nil. Returns a bool zero value if the
// pointer was nil.
func ToBool(p *bool) (v bool) {
	if p == nil {
		return v
	}

	return *p
}

// ToBoolSlice returns a slice of bool values, that are
// dereferenced if the passed in pointer was not nil. Returns a bool
// zero value if the pointer was nil.
func ToBoolSlice(vs []*bool) []bool {
	ps := make([]bool, len(vs))
	for i, v := range vs {
		ps[i] = ToBool(v)
	}

	return ps
}

// ToBoolMap returns a map of bool values, that are
// dereferenced if the passed in pointer was not nil. The bool
// zero value is used if the pointer was nil.
func ToBoolMap(vs map[string]*bool) map[string]bool {
	ps := make(map[string]bool, len(vs))
	for k, v := range vs {
		ps[k] = ToBool(v)
	}

	return ps
}

// ToByte returns byte value dereferenced if the passed
// in pointer was not nil. Returns a byte zero value if the
// pointer was nil.
func ToByte(p *byte) (v byte) {
	if p == nil {
		return v
	}

	return *p
}

// ToByteSlice returns a slice of byte values, that are
// dereferenced if the passed in pointer was not nil. Returns a byte
// zero value if the pointer was nil.
func ToByteSlice(vs []*byte) []byte {
	ps := make([]byte, len(vs))
	for i, v := range vs {
		ps[i] = ToByte(v)
	}

	return ps
}

// ToByteMap returns a map of byte values, that are
// dereferenced if the passed in pointer was not nil. The byte
// zero value is used if the pointer was nil.
func ToByteMap(vs map[string]*byte) map[string]byte {
	ps := make(map[string]byte, len(vs))
	for k, v := range vs {
		ps[k] = ToByte(v)
	}

	return ps
}

// ToString returns string value dereferenced if the passed
// in pointer was not nil. Returns a string zero value if the
// pointer was nil.
func ToString(p *string) (v string) {
	if p == nil {
		return v
	}

	return *p
}

// ToStringSlice returns a slice of string values, that are
// dereferenced if the passed in pointer was not nil. Returns a string
// zero value if the pointer was nil.
func ToStringSlice(vs []*string) []string {
	ps := make([]string, len(vs))
	for i, v := range vs {
		ps[i] = ToString(v)
	}

	return ps
}

// ToStringMap returns a map of string values, that are
// dereferenced if the passed in pointer was not nil. The string
// zero value is used if the pointer was nil.
func ToStringMap(vs map[string]*string) map[string]string {
	ps := make(map[string]string, len(vs))
	for k, v := range vs {
		ps[k] = ToString(v)
	}

	return ps
}

// ToInt returns int value dereferenced if the passed
// in pointer was not nil. Returns a int zero value if the
// pointer was nil.
func ToInt(p *int) (v int) {
	if p == nil {
		return v
	}

	return *p
}

// ToIntSlice returns a slice of int values, that are
// dereferenced if the passed in pointer was not nil. Returns a int
// zero value if the pointer was nil.
func ToIntSlice(vs []*int) []int {
	ps := make([]int, len(vs))
	for i, v := range vs {
		ps[i] = ToInt(v)
	}

	return ps
}

// ToIntMap returns a map of int values, that are
// dereferenced if the passed in pointer was not nil. The int
// zero value is used if the pointer was nil.
func ToIntMap(vs map[string]*int) map[string]int {
	ps := make(map[string]int, len(vs))
	for k, v := range vs {
		ps[k] = ToInt(v)
	}

	return ps
}

// ToInt8 returns int8 value dereferenced if the passed
// in pointer was not nil. Returns a int8 zero value if the
// pointer was nil.
func ToInt8(p *int8) (v int8) {
	if p == nil {
		return v
	}

	return *p
}

// ToInt8Slice returns a slice of int8 values, that are
// dereferenced if the passed in pointer was not nil. Returns a int8
// zero value if the pointer was nil.
func ToInt8Slice(vs []*int8) []int8 {
	ps := make([]int8, len(vs))
	for i, v := range vs {
		ps[i] = ToInt8(v)
	}

	return ps
}

// ToInt8Map returns a map of int8 values, that are
// dereferenced if the passed in pointer was not nil. The int8
// zero value is used if the pointer was nil.
func ToInt8Map(vs map[string]*int8) map[string]int8 {
	ps := make(map[string]int8, len(vs))
	for k, v := range vs {
		ps[k] = ToInt8(v)
	}

	return ps
}

// ToInt16 returns int16 value dereferenced if the passed
// in pointer was not nil. Returns a int16 zero value if the
// pointer was nil.
func ToInt16(p *int16) (v int16) {
	if p == nil {
		return v
	}

	return *p
}

// ToInt16Slice returns a slice of int16 values, that are
// dereferenced if the passed in pointer was not nil. Returns a int16
// zero value if the pointer was nil.
func ToInt16Slice(vs []*int16) []int16 {
	ps := make([]int16, len(vs))
	for i, v := range vs {
		ps[i] = ToInt16(v)
	}

	return ps
}

// ToInt16Map returns a map of int16 values, that are
// dereferenced if the passed in pointer was not nil. The int16
// zero value is used if the pointer was nil.
func ToInt16Map(vs map[string]*int16) map[string]int16 {
	ps := make(map[string]int16, len(vs))
	for k, v := range vs {
		ps[k] = ToInt16(v)
	}

	return ps
}

// ToInt32 returns int32 value dereferenced if the passed
// in pointer was not nil. Returns a int32 zero value if the
// pointer was nil.
func ToInt32(p *int32) (v int32) {
	if p == nil {
		return v
	}

	return *p
}

// ToInt32Slice returns a slice of int32 values, that are
// dereferenced if the passed in pointer was not nil. Returns a int32
// zero value if the pointer was nil.
func ToInt32Slice(vs []*int32) []int32 {
	ps := make([]int32, len(vs))
	for i, v := range vs {
		ps[i] = ToInt32(v)
	}

	return ps
}

// ToInt32Map returns a map of int32 values, that are
// dereferenced if the passed in pointer was not nil. The int32
// zero value is used if the pointer was nil.
func ToInt32Map(vs map[string]*int32) map[string]int32 {
	ps := make(map[string]int32, len(vs))
	for k, v := range vs {
		ps[k] = ToInt32(v)
	}

	return ps
}

// ToInt64 returns int64 value dereferenced if the passed
// in pointer was not nil. Returns a int64 zero value if the
// pointer was nil.
func ToInt64(p *int64) (v int64) {
	if p == nil {
		return v
	}

	return *p
}

// ToInt64Slice returns a slice of int64 values, that are
// dereferenced if the passed in pointer was not nil. Returns a int64
// zero value if the pointer was nil.
func ToInt64Slice(vs []*int64) []int64 {
	ps := make([]int64, len(vs))
	for i, v := range vs {
		ps[i] = ToInt64(v)
	}

	return ps
}

// ToInt64Map returns a map of int64 values, that are
// dereferenced if the passed in pointer was not nil. The int64
// zero value is used if the pointer was nil.
func ToInt64Map(vs map[string]*int64) map[string]int64 {
	ps := make(map[string]int64, len(vs))
	for k, v := range vs {
		ps[k] = ToInt64(v)
	}

	return ps
}

// ToUint returns uint value dereferenced if the passed
// in pointer was not nil. Returns a uint zero value if the
// pointer was nil.
func ToUint(p *uint) (v uint) {
	if p == nil {
		return v
	}

	return *p
}

// ToUintSlice returns a slice of uint values, that are
// dereferenced if the passed in pointer was not nil. Returns a uint
// zero value if the pointer was nil.
func ToUintSlice(vs []*uint) []uint {
	ps := make([]uint, len(vs))
	for i, v := range vs {
		ps[i] = ToUint(v)
	}

	return ps
}

// ToUintMap returns a map of uint values, that are
// dereferenced if the passed in pointer was not nil. The uint
// zero value is used if the pointer was nil.
func ToUintMap(vs map[string]*uint) map[string]uint {
	ps := make(map[string]uint, len(vs))
	for k, v := range vs {
		ps[k] = ToUint(v)
	}

	return ps
}

// ToUint8 returns uint8 value dereferenced if the passed
// in pointer was not nil. Returns a uint8 zero value if the
// pointer was nil.
func ToUint8(p *uint8) (v uint8) {
	if p == nil {
		return v
	}

	return *p
}

// ToUint8Slice returns a slice of uint8 values, that are
// dereferenced if the passed in pointer was not nil. Returns a uint8
// zero value if the pointer was nil.
func ToUint8Slice(vs []*uint8) []uint8 {
	ps := make([]uint8, len(vs))
	for i, v := range vs {
		ps[i] = ToUint8(v)
	}

	return ps
}

// ToUint8Map returns a map of uint8 values, that are
// dereferenced if the passed in pointer was not nil. The uint8
// zero value is used if the pointer was nil.
func ToUint8Map(vs map[string]*uint8) map[string]uint8 {
	ps := make(map[string]uint8, len(vs))
	for k, v := range vs {
		ps[k] = ToUint8(v)
	}

	return ps
}

// ToUint16 returns uint16 value dereferenced if the passed
// in pointer was not nil. Returns a uint16 zero value if the
// pointer was nil.
func ToUint16(p *uint16) (v uint16) {
	if p == nil {
		return v
	}

	return *p
}

// ToUint16Slice returns a slice of uint16 values, that are
// dereferenced if the passed in pointer was not nil. Returns a uint16
// zero value if the pointer was nil.
func ToUint16Slice(vs []*uint16) []uint16 {
	ps := make([]uint16, len(vs))
	for i, v := range vs {
		ps[i] = ToUint16(v)
	}

	return ps
}

// ToUint16Map returns a map of uint16 values, that are
// dereferenced if the passed in pointer was not nil. The uint16
// zero value is used if the pointer was nil.
func ToUint16Map(vs map[string]*uint16) map[string]uint16 {
	ps := make(map[string]uint16, len(vs))
	for k, v := range vs {
		ps[k] = ToUint16(v)
	}

	return ps
}

// ToUint32 returns uint32 value dereferenced if the passed
// in pointer was not nil. Returns a uint32 zero value if the
// pointer was nil.
func ToUint32(p *uint32) (v uint32) {
	if p == nil {
		return v
	}

	return *p
}

// ToUint32Slice returns a slice of uint32 values, that are
// dereferenced if the passed in pointer was not nil. Returns a uint32
// zero value if the pointer was nil.
func ToUint32Slice(vs []*uint32) []uint32 {
	ps := make([]uint32, len(vs))
	for i, v := range vs {
		ps[i] = ToUint32(v)
	}

	return ps
}

// ToUint32Map returns a map of uint32 values, that are
// dereferenced if the passed in pointer was not nil. The uint32
// zero value is used if the pointer was nil.
func ToUint32Map(vs map[string]*uint32) map[string]uint32 {
	ps := make(map[string]uint32, len(vs))
	for k, v := range vs {
		ps[k] = ToUint32(v)
	}

	return ps
}

// ToUint64 returns uint64 value dereferenced if the passed
// in pointer was not nil. Returns a uint64 zero value if the
// pointer was nil.
func ToUint64(p *uint64) (v uint64) {
	if p == nil {
		return v
	}

	return *p
}

// ToUint64Slice returns a slice of uint64 values, that are
// dereferenced if the passed in pointer was not nil. Returns a uint64
// zero value if the pointer was nil.
func ToUint64Slice(vs []*uint64) []uint64 {
	ps := make([]uint64, len(vs))
	for i, v := range vs {
		ps[i] = ToUint64(v)
	}

	return ps
}

// ToUint64Map returns a map of uint64 values, that are
// dereferenced if the passed in pointer was not nil. The uint64
// zero value is used if the pointer was nil.
func ToUint64Map(vs map[string]*uint64) map[string]uint64 {
	ps := make(map[string]uint64, len(vs))
	for k, v := range vs {
		ps[k] = ToUint64(v)
	}

	return ps
}

// ToFloat32 returns float32 value dereferenced if the passed
// in pointer was not nil. Returns a float32 zero value if the
// pointer was nil.
func ToFloat32(p *float32) (v float32) {
	if p == nil {
		return v
	}

	return *p
}

// ToFloat32Slice returns a slice of float32 values, that are
// dereferenced if the passed in pointer was not nil. Returns a float32
// zero value if the pointer was nil.
func ToFloat32Slice(vs []*float32) []float32 {
	ps := make([]float32, len(vs))
	for i, v := range vs {
		ps[i] = ToFloat32(v)
	}

	return ps
}

// ToFloat32Map returns a map of float32 values, that are
// dereferenced if the passed in pointer was not nil. The float32
// zero value is used if the pointer was nil.
func ToFloat32Map(vs map[string]*float32) map[string]float32 {
	ps := make(map[string]float32, len(vs))
	for k, v := range vs {
		ps[k] = ToFloat32(v)
	}

	return ps
}

// ToFloat64 returns float64 value dereferenced if the passed
// in pointer was not nil. Returns a float64 zero value if the
// pointer was nil.
func ToFloat64(p *float64) (v float64) {
	if p == nil {
		return v
	}

	return *p
}

// ToFloat64Slice returns a slice of float64 values, that are
// dereferenced if the passed in pointer was not nil. Returns a float64
// zero value if the pointer was nil.
func ToFloat64Slice(vs []*float64) []float64 {
	ps := make([]float64, len(vs))
	for i, v := range vs {
		ps[i] = ToFloat64(v)
	}

	return ps
}

// ToFloat64Map returns a map of float64 values, that are
// dereferenced if the passed in pointer was not nil. The float64
// zero value is used if the pointer was nil.
func ToFloat64Map(vs map[string]*float64) map[string]float64 {
	ps := make(map[string]float64, len(vs))
	for k, v := range vs {
		ps[k] = ToFloat64(v)
	}

	return ps
}

// ToTime returns time.Time value dereferenced if the passed
// in pointer was not nil. Returns a time.Time zero value if the
// pointer was nil.
func ToTime(p *time.Time) (v time.Time) {
	if p == nil {
		return v
	}

	return *p
}

// ToTimeSlice returns a slice of time.Time values, that are
// dereferenced if the passed in pointer was not nil. Returns a time.Time
// zero value if the pointer was nil.
func ToTimeSlice(vs []*time.Time) []time.Time {
	ps := make([]time.Time, len(vs))
	for i, v := range vs {
		ps[i] = ToTime(v)
	}

	return ps
}

// ToTimeMap returns a map of time.Time values, that are
// dereferenced if the passed in pointer was not nil. The time.Time
// zero value is used if the pointer was nil.
func ToTimeMap(vs map[string]*time.Time) map[string]time.Time {
	ps := make(map[string]time.Time, len(vs))
	for k, v := range vs {
		ps[k] = ToTime(v)
	}

	return ps
}

// ToDuration returns time.Duration value dereferenced if the passed
// in pointer was not nil. Returns a time.Duration zero value if the
// pointer was nil.
func ToDuration(p *time.Duration) (v time.Duration) {
	if p == nil {
		return v
	}

	return *p
}

// ToDurationSlice returns a slice of time.Duration values, that are
// dereferenced if the passed in pointer was not nil. Returns a time.Duration
// zero value if the pointer was nil.
func ToDurationSlice(vs []*time.Duration) []time.Duration {
	ps := make([]time.Duration, len(vs))
	for i, v := range vs {
		ps[i] = ToDuration(v)
	}

	return ps
}

// ToDurationMap returns a map of time.Duration values, that are
// dereferenced if the passed in pointer was not nil. The time.Duration
// zero value is used if the pointer was nil.
func ToDurationMap(vs map[string]*time.Duration) map[string]time.Duration {
	ps := make(map[string]time.Duration, len(vs))
	for k, v := range vs {
		ps[k] = ToDuration(v)
	}

	return ps
}
