package omnibus

import "time"

func UintPtr(i uint) *uint {
	return &i
}

func UintDeref(ptr *uint, def uint) uint {
	if ptr != nil {
		return *ptr
	}
	return def
}

func UintPtrEqual(a, b *uint) bool {
	if (a == nil) != (b == nil) {
		return false
	}
	if a == nil {
		return true
	}
	return *a == *b
}

func Uint8Ptr(i uint8) *uint8 {
	return &i
}

func Uint8Deref(ptr *uint8, def uint8) uint8 {
	if ptr != nil {
		return *ptr
	}
	return def
}

func Uint8PtrEqual(a, b *uint8) bool {
	if (a == nil) != (b == nil) {
		return false
	}
	if a == nil {
		return true
	}
	return *a == *b
}

func Uint16Ptr(i uint16) *uint16 {
	return &i
}

func Uint16Deref(ptr *uint16, def uint16) uint16 {
	if ptr != nil {
		return *ptr
	}
	return def
}

func Uint16PtrEqual(a, b *uint16) bool {
	if (a == nil) != (b == nil) {
		return false
	}
	if a == nil {
		return true
	}
	return *a == *b
}

func Uint32Ptr(i uint32) *uint32 {
	return &i
}

func Uint32Deref(ptr *uint32, def uint32) uint32 {
	if ptr != nil {
		return *ptr
	}
	return def
}

func Uint32PtrEqual(a, b *uint32) bool {
	if (a == nil) != (b == nil) {
		return false
	}
	if a == nil {
		return true
	}
	return *a == *b
}

func Uint64Ptr(i uint64) *uint64 {
	return &i
}

func Uint64Deref(ptr *uint64, def uint64) uint64 {
	if ptr != nil {
		return *ptr
	}
	return def
}

func Uint64PtrEqual(a, b *uint64) bool {
	if (a == nil) != (b == nil) {
		return false
	}
	if a == nil {
		return true
	}
	return *a == *b
}

func IntPtr(i int) *int {
	return &i
}

func IntDeref(ptr *int, def int) int {
	if ptr != nil {
		return *ptr
	}
	return def
}

func IntPtrEqual(a, b *int) bool {
	if (a == nil) != (b == nil) {
		return false
	}
	if a == nil {
		return true
	}
	return *a == *b
}

func Int8Ptr(i int8) *int8 {
	return &i
}

func Int8Deref(ptr *int8, def int8) int8 {
	if ptr != nil {
		return *ptr
	}
	return def
}

func Int8PtrEqual(a, b *int8) bool {
	if (a == nil) != (b == nil) {
		return false
	}
	if a == nil {
		return true
	}
	return *a == *b
}

func Int16Ptr(i int16) *int16 {
	return &i
}

func Int16Deref(ptr *int16, def int16) int16 {
	if ptr != nil {
		return *ptr
	}
	return def
}

func Int16PtrEqual(a, b *int16) bool {
	if (a == nil) != (b == nil) {
		return false
	}
	if a == nil {
		return true
	}
	return *a == *b
}

func Int32Ptr(i int32) *int32 {
	return &i
}

func Int32Deref(ptr *int32, def int32) int32 {
	if ptr != nil {
		return *ptr
	}
	return def
}

func Int32PtrEqual(a, b *int32) bool {
	if (a == nil) != (b == nil) {
		return false
	}
	if a == nil {
		return true
	}
	return *a == *b
}

func Int64Ptr(i int64) *int64 {
	return &i
}

func Int64Deref(ptr *int64, def int64) int64 {
	if ptr != nil {
		return *ptr
	}
	return def
}

func Int64PtrEqual(a, b *int64) bool {
	if (a == nil) != (b == nil) {
		return false
	}
	if a == nil {
		return true
	}
	return *a == *b
}

func StringPtr(s string) *string {
	return &s
}

func StringDeref(ptr *string, def string) string {
	if ptr != nil {
		return *ptr
	}
	return def
}

func StringPtrEqual(a, b *string) bool {
	if (a == nil) != (b == nil) {
		return false
	}
	if a == nil {
		return true
	}
	return *a == *b
}

func Float32Ptr(f float32) *float32 {
	return &f
}

func Float32Deref(ptr *float32, def float32) float32 {
	if ptr != nil {
		return *ptr
	}
	return def
}

func Float32PtrEqual(a, b *float32) bool {
	if (a == nil) != (b == nil) {
		return false
	}
	if a == nil {
		return true
	}
	return *a == *b
}

func Float64Ptr(f float64) *float64 {
	return &f
}

func Float64Deref(ptr *float64, def float64) float64 {
	if ptr != nil {
		return *ptr
	}
	return def
}

func Float64PtrEqual(a, b *float64) bool {
	if (a == nil) != (b == nil) {
		return false
	}
	if a == nil {
		return true
	}
	return *a == *b
}

func BoolPtr(b bool) *bool {
	return &b
}

func BoolDeref(ptr *bool, def bool) bool {
	if ptr != nil {
		return *ptr
	}
	return def
}

func BoolPtrEqual(a, b *bool) bool {
	if (a == nil) != (b == nil) {
		return false
	}
	if a == nil {
		return true
	}
	return *a == *b
}

func Duration(d time.Duration) *time.Duration {
	return &d
}

func DurationDeref(ptr *time.Duration, def time.Duration) time.Duration {
	if ptr != nil {
		return *ptr
	}
	return def
}

func DurationPtrEqual(a, b *time.Duration) bool {
	if (a == nil) != (b == nil) {
		return false
	}
	if a == nil {
		return true
	}
	return *a == *b
}
