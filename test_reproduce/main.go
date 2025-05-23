package main

import (
	"fmt"
	"image"

	"github.com/disintegration/imaging"
)

func Test() {
	yStride := 288
	cStride := 144
	width := 288
	height := 147

	im := &image.YCbCr{
		Y:              make([]uint8, yStride*height),
		Cb:             make([]uint8, (cStride*height)/2),
		Cr:             make([]uint8, (cStride*height)/2),
		YStride:        yStride,
		CStride:        yStride / 2,
		SubsampleRatio: image.YCbCrSubsampleRatio420,
		Rect: image.Rectangle{
			Min: image.Point{
				X: 0,
				Y: 0,
			},
			Max: image.Point{
				X: width,
				Y: height,
			},
		},
	}
	imaging.Rotate90(im)
}

func TestOddDimensions() {
	fmt.Println("Testing YCbCr with odd dimensions...")

	// Test case 1: Original issue - height=147 (odd)
	fmt.Print("Test 1 (height=147): ")
	testCase1 := &image.YCbCr{
		Y:              make([]uint8, 288*147),  // yStride * height
		Cb:             make([]uint8, (144*147)/2), // (cStride * height) / 2
		Cr:             make([]uint8, (144*147)/2),
		YStride:        288,
		CStride:        144,
		SubsampleRatio: image.YCbCrSubsampleRatio420,
		Rect: image.Rectangle{
			Min: image.Point{X: 0, Y: 0},
			Max: image.Point{X: 288, Y: 147},
		},
	}
	imaging.Rotate90(testCase1)
	fmt.Println("PASS")

	// Test case 2: Width odd, height even
	fmt.Print("Test 2 (width=289, height=146): ")
	testCase2 := &image.YCbCr{
		Y:              make([]uint8, 289*146),
		Cb:             make([]uint8, (145*146)/2),
		Cr:             make([]uint8, (145*146)/2),
		YStride:        289,
		CStride:        145,
		SubsampleRatio: image.YCbCrSubsampleRatio420,
		Rect: image.Rectangle{
			Min: image.Point{X: 0, Y: 0},
			Max: image.Point{X: 289, Y: 146},
		},
	}
	imaging.Rotate90(testCase2)
	fmt.Println("PASS")

	// Test case 3: Both dimensions odd
	fmt.Print("Test 3 (width=289, height=147): ")
	testCase3 := &image.YCbCr{
		Y:              make([]uint8, 289*147),
		Cb:             make([]uint8, (145*147)/2),
		Cr:             make([]uint8, (145*147)/2),
		YStride:        289,
		CStride:        145,
		SubsampleRatio: image.YCbCrSubsampleRatio420,
		Rect: image.Rectangle{
			Min: image.Point{X: 0, Y: 0},
			Max: image.Point{X: 289, Y: 147},
		},
	}
	imaging.Rotate90(testCase3)
	fmt.Println("PASS")

	// Test case 4: Small odd dimensions
	fmt.Print("Test 4 (width=3, height=5): ")
	testCase4 := &image.YCbCr{
		Y:              make([]uint8, 3*5),
		Cb:             make([]uint8, (2*5)/2),
		Cr:             make([]uint8, (2*5)/2),
		YStride:        3,
		CStride:        2,
		SubsampleRatio: image.YCbCrSubsampleRatio420,
		Rect: image.Rectangle{
			Min: image.Point{X: 0, Y: 0},
			Max: image.Point{X: 3, Y: 5},
		},
	}
	imaging.Rotate90(testCase4)
	fmt.Println("PASS")

	// Test other transformations too
	fmt.Print("Test 5 (Rotate180): ")
	imaging.Rotate180(testCase1)
	fmt.Println("PASS")

	fmt.Print("Test 6 (Rotate270): ")
	imaging.Rotate270(testCase1)
	fmt.Println("PASS")

	fmt.Print("Test 7 (FlipH): ")
	imaging.FlipH(testCase1)
	fmt.Println("PASS")

	fmt.Print("Test 8 (FlipV): ")
	imaging.FlipV(testCase1)
	fmt.Println("PASS")

	fmt.Println("All tests passed!")
}

func main() {
	Test()
	TestOddDimensions()
} 