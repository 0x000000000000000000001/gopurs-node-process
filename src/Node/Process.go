package Process

import (
	"os"
	"strings"
	"runtime"
	"unsafe"
	"gopurs/output/gopurs_runtime"
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
	arr := make([]gopurs_runtime.Value, len(os.Args))
	for i, arg := range os.Args {
		arr[i] = gopurs_runtime.Str(arg)
	}
	// Return the array directly as a Value
	return gopurs_runtime.Value{Type: gopurs_runtime.TypeArray, UnsafePtr: unsafe.Pointer(&arr)}
}

func Argv0(_ interface{}) interface{} {
	if len(os.Args) > 0 {
		return gopurs_runtime.Str(os.Args[0])
	}
	return gopurs_runtime.Str("")
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

func Cwd(_ interface{}) string {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return dir
}

func DebugPort(_ interface{}) interface{} {
	panic("Not implemented: debugPort")
}

func DisconnectImpl(_ interface{}) interface{} {
	panic("Not implemented: disconnectImpl")
}

func GetEnv(_ interface{}) interface{} {
	env := make(map[string]interface{})
	for _, e := range os.Environ() {
		parts := strings.SplitN(e, "=", 2)
		if len(parts) == 2 {
			env[parts[0]] = parts[1]
		}
	}
	return env
}

func UnsafeGetEnv(_ interface{}) interface{} {
	return GetEnv(nil)
}

func SetEnvImpl(keyVal interface{}, valVal interface{}) interface{} {
	key := gopurs_runtime.Unbox[string](keyVal)
	val := gopurs_runtime.Unbox[string](valVal)
	os.Setenv(key, val)
	return nil
}

func UnsetEnvImpl(keyVal interface{}) interface{} {
	key := gopurs_runtime.Unbox[string](keyVal)
	os.Unsetenv(key)
	return nil
}

func ExecArgv(_ interface{}) interface{} {
	panic("Not implemented: execArgv")
}

func ExecPath(_ interface{}) string {
	exe, err := os.Executable()
	if err != nil {
		panic(err)
	}
	return exe
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

var Pid = int64(os.Getpid())

var PlatformStr = runtime.GOOS

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

var Version = runtime.Version()
