package util

import (
	"errors"
	"fmt"

	"gocv.io/x/gocv"
)

type RTSPValidationResult struct {
	Resolution string
	Codec      string
	FramePath  string
}

func ValidateRTSP(rtspURL string) (RTSPValidationResult, error) {
	var result RTSPValidationResult

	webcam, err := gocv.VideoCaptureDevice(rtspURL)
	if err != nil {
		return result, errors.New("failed to connect to RTSP stream")
	}
	defer webcam.Close()

	// Check if the stream is open
	if !webcam.IsOpened() {
		return result, errors.New("RTSP stream is offline")
	}

	// Get properties: Resolution
	width := webcam.Get(gocv.VideoCaptureFrameWidth)
	height := webcam.Get(gocv.VideoCaptureFrameHeight)
	result.Resolution = fmt.Sprintf("%dx%d", int(width), int(height))

	// Capture a frame
	frame := gocv.NewMat()
	defer frame.Close()
	if ok := webcam.Read(&frame); !ok || frame.Empty() {
		return result, errors.New("failed to capture frame")
	}

	// Save the frame as an image
	framePath := "frame.jpg"
	if ok := gocv.IMWrite(framePath, frame); !ok {
		return result, errors.New("failed to save frame")
	}
	result.FramePath = framePath

	// Get codec info (for example purposes, use a placeholder)
	result.Codec = "H.264 (example codec)"

	return result, nil
}
