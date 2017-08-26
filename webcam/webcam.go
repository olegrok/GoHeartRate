package webcam

import (
	"fmt"
	"github.com/lazywei/go-opencv/opencv"
	"image/color"
	"os"
	"path"
	"reflect"
	"time"
)

func sum(src *opencv.IplImage) uint32 {
	var sum uint32
	for i := 0; i < src.Width(); i++ {
		for j := 0; j < src.Height(); j++ {
			channels := src.Get2D(i, j).Val()
			sum += uint32(channels[1])
		}
	}

	return sum
}

func img_rgb2ycrcb(img *opencv.IplImage) *opencv.IplImage {
	resultMask := opencv.Crop(img, 0, 0, img.Width(), img.Height())
	for i := 0; i < img.Width(); i++ {
		for j := 0; j < img.Height(); j++ {
			img_pix := img.Get2D(i, j).Val()
			y, cb, cr := color.RGBToYCbCr(uint8(img_pix[0]), uint8(img_pix[1]), uint8(img_pix[2]))
			yCrCb := opencv.NewScalar(float64(y), float64(cb), float64(cr), float64(255))
			resultMask.Set2D(i, j, yCrCb)
		}
	}
	return resultMask
}

func image_pixel_and(img *opencv.IplImage, mask *opencv.IplImage) *opencv.IplImage {
	resultMask := opencv.Crop(img, 0, 0, img.Width(), img.Height())
	for i := 0; i < img.Width(); i++ {
		for j := 0; j < img.Height(); j++ {
			img_pix := img.Get2D(i, j).Val()
			mask_pix := mask.Get2D(i, j).Val()
			s := opencv.NewScalar(img_pix[0]*mask_pix[0], img_pix[1]*mask_pix[1], img_pix[2]*mask_pix[2], 0)
			resultMask.Set2D(i, j, s)
		}
	}
	return resultMask
}

func in_range(img *opencv.IplImage, low_thresh [3]float64, high_thresh [3]float64) *opencv.IplImage {
	resultMask := opencv.Crop(img, 0, 0, img.Width(), img.Height())
	for i := 0; i < img.Width(); i++ {
		for j := 0; j < img.Height(); j++ {
			channels := img.Get2D(i, j).Val()

			if channels[0] < high_thresh[0] && channels[1] < high_thresh[1] && channels[2] < high_thresh[2] {
				if channels[0] > low_thresh[0] && channels[1] > low_thresh[1] && channels[2] > low_thresh[2] {
					s := opencv.NewScalar(1., 1., 1., 0.0)
					resultMask.Set2D(i, j, s)
				} else {
					s := opencv.NewScalar(0., 0., 0., 0.0)
					resultMask.Set2D(i, j, s)
				}
			} else {
				s := opencv.NewScalar(0., 0., 0., 0.0)
				resultMask.Set2D(i, j, s)
			}
		}
	}
	return resultMask

}

func detectSkin(face *opencv.IplImage) *opencv.IplImage {

	ycrcbFace := img_rgb2ycrcb(face)
	low := [3]float64{0.0, 133.0, 77.0}
	high := [3]float64{255.0, 173.0, 127.0}
	mask := in_range(ycrcbFace, low, high)
	result := image_pixel_and(face, mask)
	// offset := opencv.Point{0, 0}
	//
	// contours := mask.FindContours(0, 1, offset)
	// for i := 0; i < contours.Total(); i++ {
	// 	area := opencv.ContourArea(contours, opencv.Slice{}, 0)
	// 	if area >2000 {
	// 		opencv.DrawContours(face, contours, externalColor, holeColor, maxLevel, thickness, lineType, offset)
	// 	}
	// }
	winResult := opencv.NewWindow("Some shit")
	winResult.ShowImage(result)
	return result
}

func Detectface(img *opencv.IplImage, cascade *opencv.HaarCascade) *opencv.IplImage {
	faces := cascade.DetectObjects(img)
	//opencv.Circle(img,
	//	opencv.Point{
	//		faces[0].X() + (faces[0].Width() / 2),
	//		faces[0].Y() + (faces[0].Height() / 2),
	//	},
	//	faces[0].Width()/2,
	//	opencv.ScalarAll(255.0), 1, 1, 0)
	if len(faces) > 0 {
		only_face := opencv.Crop(img, faces[0].X(), faces[0].Y(), faces[0].Width(), faces[0].Height())
		return only_face
	}
	return img
}

func Start() ([]float64, []float64) {
	const n int = 128
	var signal []float64
	var sampleTime []float64
	win := opencv.NewWindow("Go-OpenCV Webcam Face Detection")
	//winMask := opencv.NewWindow("Mask")
	//winResult := opencv.NewWindow("Result")
	defer win.Destroy()

	cap := opencv.NewCameraCapture(0)
	fmt.Println(reflect.TypeOf(*cap).Kind())
	if cap == nil {
		panic("cannot open camera")
	}
	defer cap.Release()

	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	cascade := opencv.LoadHaarClassifierCascade(path.Join(cwd, "../haarcascade_frontalface_alt.xml"))

	fmt.Println("Press ESC to quit")

	frame_num := 0

	for frame_num < n {
		startTime := time.Now()
		if cap.GrabFrame() {
			img := cap.RetrieveFrame(1)

			if img != nil {
				only_face := Detectface(img, cascade)
				win.ShowImage(only_face)

				if only_face.Width() != 0 && only_face.Height() != 0 {
					skinRegion := detectSkin(only_face)
					signal = append(signal, float64(sum(skinRegion)))
					sampleTime = append(sampleTime, time.Since(startTime).Seconds())
					//sampleTime[frame_num] = timeDelta.Seconds()
					frame_num += 1
				}

			} else {
				fmt.Println("nil image")
			}
		}
		key := opencv.WaitKey(1)

		if key == 27 {
			os.Exit(0)
		}
	}
	fmt.Println(signal)
	return signal, sampleTime
}
