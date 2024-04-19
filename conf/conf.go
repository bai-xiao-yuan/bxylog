package conf

type LogLevel int8

const (
	LInfo  LogLevel = 1
	LDebug LogLevel = 2
	LWarn  LogLevel = 3
	LError LogLevel = 4
	LPanic LogLevel = 5
)

type OutTarget int8

const (
	Std  OutTarget = 1
	File OutTarget = 2
)

type SliceType int8

const (
	None     SliceType = 0
	Day      SliceType = 1 //天
	FileSize SliceType = 2 //kb
	TimeDate SliceType = 3 //hours
)

type Config struct {
	Level         LogLevel
	OutTarget     OutTarget
	FileName      string
	FileSliceType SliceType
	SliceProp     int64
	Prefix        string
	Color         bool
	TimeFlag      bool
}
