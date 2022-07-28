//go:build !windows

package main

func ffmpeg_name() string {
	return "./ffmpeg"
}