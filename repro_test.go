package imaging

import (
	"image"
	"testing"

	"github.com/disintegration/imaging"
)

func TestRotate90_OddHeightYCbCrPanic(t *testing.T) {
	yStride := 288
	cStride := 144
	width := 288
	height := 147

	im := &image.YCbCr{
		Y:              make([]uint8, yStride*height),
		Cb:             make([]uint8, (cStride*height)/2),
		Cr:             make([]uint8, (cStride*height)/2),
		YStride:        yStride,
		CStride:        cStride, // Correctly use cStride here
		SubsampleRatio: image.YCbCrSubsampleRatio420,
		Rect:           image.Rect(0, 0, width, height),
	}

	_ = imaging.Rotate90(im) // Call Rotate90, result is not used yet
}
