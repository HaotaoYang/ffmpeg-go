package ffmpeg_go

import (
	"bytes"
	"context"
	"os/exec"
	"time"
)

// Run ffprobe on the specified file and return a JSON representation of the output.
//
//  Raises:
//        :class:`ffmpeg.Error`: if ffprobe returns a non-zero exit code,
//            an :class:`Error` is returned with a generic error message.
//            The stderr output can be retrieved by accessing the
//            ``stderr`` property of the exception.

func Probe(fileName string, kwargs KwArgs) (string, error) {
	return ProbeWithTimeout(fileName, 0, kwargs)
}

func ProbeWithTimeout(fileName string, timeOut time.Duration, kwargs KwArgs) (string, error) {
	args := []string{"-show_format", "-show_streams", "-of", "json"}
	args = append(args, ConvertKwargsToCmdLineArgs(kwargs)...)
	args = append(args, fileName)
	ctx := context.Background()
	if timeOut > 0 {
		var cancel func()
		ctx, cancel = context.WithTimeout(context.Background(), timeOut)
		defer cancel()
	}
	cmd := exec.CommandContext(ctx, "ffprobe", args...)
	buf := bytes.NewBuffer(nil)
	cmd.Stdout = buf
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	return string(buf.Bytes()), nil
}
