package slog

import "log/slog"

// adapter to add wrap error and add it to slog Logger message
func Err(err error) slog.Attr {
	// original code
	// return slog.Attr{
	// 	Key:   "error",
	// 	Value: slog.StringValue(err.Error()),
	// }

	if err == nil {
		return slog.Any("error", nil)
	}
	return slog.Any("error", err)
}
