package internal

// IntPtr int to *int
func IntPtr(i int) *int {
	return &i
}

// Int32Ptr int32 to *int32
func Int32Ptr(i int32) *int32 {
	return &i
}

// Int64Ptr int64 to *int64
func Int64Ptr(i int64) *int64 {
	return &i
}

// UintPtr uint to *uint
func UintPtr(u uint) *uint {
	return &u
}

// Uint32Ptr uint32 to *uint32
func Uint32Ptr(u uint32) *uint32 {
	return &u
}

// Uint64Ptr uint64 to *uint64
func Uint64Ptr(u uint64) *uint64 {
	return &u
}

// Float32Ptr float32 to *float32
func Float32Ptr(f float32) *float32 {
	return &f
}

// Float64Ptr float64 to *float64
func Float64Ptr(f float64) *float64 {
	return &f
}

// StrPtr string to *string
func StrPtr(s string) *string {
	return &s
}
