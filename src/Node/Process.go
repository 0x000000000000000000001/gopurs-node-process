package Process

import (
	"os"
)

func ExitImpl(code int64) interface{} {
	os.Exit(int(code))
	return nil
}

func Process(_ interface{}) interface{} {
	panic("Not implemented: process")
}

func AbortImpl(_ interface{}) interface{} {
	panic("Not implemented: abortImpl")
}

func Argv(_ interface{}) interface{} {
	panic("Not implemented: argv")
}

func Argv0(_ interface{}) interface{} {
	panic("Not implemented: argv0")
}

func ChannelRefImpl(_ interface{}) interface{} {
	panic("Not implemented: channelRefImpl")
}

func ChannelUnrefImpl(_ interface{}) interface{} {
	panic("Not implemented: channelUnrefImpl")
}

func ChdirImpl(_ interface{}) interface{} {
	panic("Not implemented: chdirImpl")
}

func Config(_ interface{}) interface{} {
	panic("Not implemented: config")
}

func Connected(_ interface{}) interface{} {
	panic("Not implemented: connected")
}

func CpuUsage(_ interface{}) interface{} {
	panic("Not implemented: cpuUsage")
}

func CpuUsageDiffImpl(_ interface{}) interface{} {
	panic("Not implemented: cpuUsageDiffImpl")
}

func Cwd(_ interface{}) interface{} {
	panic("Not implemented: cwd")
}

func DebugPort(_ interface{}) interface{} {
	panic("Not implemented: debugPort")
}

func DisconnectImpl(_ interface{}) interface{} {
	panic("Not implemented: disconnectImpl")
}

func GetEnv(_ interface{}) interface{} {
	panic("Not implemented: getEnv")
}

func UnsafeGetEnv(_ interface{}) interface{} {
	panic("Not implemented: unsafeGetEnv")
}

func SetEnvImpl(_ interface{}, _ interface{}) interface{} {
	panic("Not implemented: setEnvImpl")
}

func UnsetEnvImpl(_ interface{}) interface{} {
	panic("Not implemented: unsetEnvImpl")
}

func ExecArgv(_ interface{}) interface{} {
	panic("Not implemented: execArgv")
}

func ExecPath(_ interface{}) interface{} {
	panic("Not implemented: execPath")
}

func Exit(_ interface{}) interface{} {
	panic("Not implemented: exit")
}

func SetExitCodeImpl(_ interface{}) interface{} {
	panic("Not implemented: setExitCodeImpl")
}

func GetExitCodeImpl(_ interface{}) interface{} {
	panic("Not implemented: getExitCodeImpl")
}

func GetGidImpl(_ interface{}) interface{} {
	panic("Not implemented: getGidImpl")
}

func GetUidImpl(_ interface{}) interface{} {
	panic("Not implemented: getUidImpl")
}

func HasUncaughtExceptionCaptureCallback(_ interface{}) interface{} {
	panic("Not implemented: hasUncaughtExceptionCaptureCallback")
}

func KillImpl(_ interface{}) interface{} {
	panic("Not implemented: killImpl")
}

func KillStrImpl(_ interface{}, _ interface{}) interface{} {
	panic("Not implemented: killStrImpl")
}

func KillIntImpl(_ interface{}, _ interface{}) interface{} {
	panic("Not implemented: killIntImpl")
}

func MemoryUsage(_ interface{}) interface{} {
	panic("Not implemented: memoryUsage")
}

func MemoryUsageRss(_ interface{}) interface{} {
	panic("Not implemented: memoryUsageRss")
}

func NextTickImpl(_ interface{}) interface{} {
	panic("Not implemented: nextTickImpl")
}

func NextTickCbImpl(_ interface{}, _ interface{}) interface{} {
	panic("Not implemented: nextTickCbImpl")
}

func Pid(_ interface{}) interface{} {
	panic("Not implemented: pid")
}

func PlatformStr(_ interface{}) interface{} {
	panic("Not implemented: platformStr")
}

func Ppid(_ interface{}) interface{} {
	panic("Not implemented: ppid")
}

func ResourceUsage(_ interface{}) interface{} {
	panic("Not implemented: resourceUsage")
}

func SendImpl(_ interface{}, _ interface{}) interface{} {
	panic("Not implemented: sendImpl")
}

func SendOptsImpl(_ interface{}, _ interface{}, _ interface{}) interface{} {
	panic("Not implemented: sendOptsImpl")
}

func SendCbImpl(_ interface{}, _ interface{}, _ interface{}) interface{} {
	panic("Not implemented: sendCbImpl")
}

func SendOptsCbImpl(_ interface{}, _ interface{}, _ interface{}, _ interface{}) interface{} {
	panic("Not implemented: sendOptsCbImpl")
}

func SetUncaughtExceptionCaptureCallbackImpl(_ interface{}) interface{} {
	panic("Not implemented: setUncaughtExceptionCaptureCallbackImpl")
}

func ClearUncaughtExceptionCaptureCallback(_ interface{}) interface{} {
	panic("Not implemented: clearUncaughtExceptionCaptureCallback")
}

func Stdin(_ interface{}) interface{} {
	panic("Not implemented: stdin")
}

func Stdout(_ interface{}) interface{} {
	panic("Not implemented: stdout")
}

func Stderr(_ interface{}) interface{} {
	panic("Not implemented: stderr")
}

func StdinIsTTY(_ interface{}) interface{} {
	panic("Not implemented: stdinIsTTY")
}

func StdoutIsTTY(_ interface{}) interface{} {
	panic("Not implemented: stdoutIsTTY")
}

func StderrIsTTY(_ interface{}) interface{} {
	panic("Not implemented: stderrIsTTY")
}

func GetTitle(_ interface{}) interface{} {
	panic("Not implemented: getTitle")
}

func SetTitleImpl(_ interface{}) interface{} {
	panic("Not implemented: setTitleImpl")
}

func Uptime(_ interface{}) interface{} {
	panic("Not implemented: uptime")
}

func Version(_ interface{}) interface{} {
	panic("Not implemented: version")
}

