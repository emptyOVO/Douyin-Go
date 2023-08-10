package utils

import "testing"

func TestFfmpeg(t *testing.T) {
	err := GenerateSnapshot("/public/1_video_20230728_090622.mp4", "1_video_20230728_090622.mp4", 1)
	if err != nil {
		return
	}
}
