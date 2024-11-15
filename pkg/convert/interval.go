package convert

import "math"

// 根据传入的fps，转成对应的dropInterval值
func ConvertFPSToDropInterval(fps float64, unit string) float64 {
	if fps == 0 {
		return 1.0
	}

	var unitVal float64
	switch unit {
	case "s":
		unitVal = 1.0
	case "m":
		unitVal = 60.0
	case "h":
		unitVal = 3600.0
	default:
		unitVal = 1.0
	}

	var frame float64 = 25 // 默认摄像头帧数25
	baseVal := (frame / fps) * unitVal
	return float64(math.Ceil(baseVal)) //向上取整
}

// 入参为多少帧抽一帧，比如25为25帧抽一帧，即一秒一帧。返回值是每秒多少帧
func ConvertDropFrameIntervalToFPS(interval int) float64 {
	if interval == 0 {
		return 25
	}

	return 25.0 / float64(interval)
}
