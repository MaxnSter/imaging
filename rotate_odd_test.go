package imaging

import (
	"image"
	"testing"
)

func TestRotate90_YCbCrOddHeight(t *testing.T) {
	yStride := 288
	cStride := 144
	width := 288
	height := 147
	im := &image.YCbCr{
		Y:              make([]uint8, yStride*height),
		Cb:             make([]uint8, (cStride*height)/2),
		Cr:             make([]uint8, (cStride*height)/2),
		YStride:        yStride,
		CStride:        cStride,
		SubsampleRatio: image.YCbCrSubsampleRatio420,
		Rect:           image.Rect(0, 0, width, height),
	}
	defer func() {
		if r := recover(); r != nil {
			t.Fatalf("Rotate90 panicked: %v", r)
		}
	}()
	_ = Rotate90(im)
}
