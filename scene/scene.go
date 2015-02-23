package scene

import (
	"fmt"
	"image"
	"image/color"
	"time"

	"github.com/creack/goray/objects"
	// Load all Object modules
	_ "github.com/creack/goray/objects/all"
)

// Scene represent the Scene.
// Hold the bondaries as well as the processed image buffer.
// Also holds various extra metadata.
type Scene struct {
	Img    *image.RGBA
	Width  int
	Height int

	MaxGoRoutines int
	Verbose       bool
}

// Eye represent the scene's Camera
// It is a point with a vector.
type Eye struct {
	Position objects.Point
	Rotation objects.Vector
}

// Config represent the configuration for a Scene.
// This contains Scene sizes, the Camera and the Object list.
type Config struct {
	Height  int
	Width   int
	Eye     Eye
	Objects []objects.Object
}

// NewScene instantiates a new Scene.
func NewScene(w, h int, maxGoRoutines int) *Scene {
	return &Scene{
		Img:           image.NewRGBA(image.Rect(0, 0, w, h)),
		Width:         w,
		Height:        h,
		MaxGoRoutines: maxGoRoutines,
	}
}

// calc calculates the color of a single point
// relative the the given camera (eye) and object list.
// To find the color, we first need to find the closest object
// to the eye crossing the line Point / Eye, then fetch the Color
// of the found object.
func (s *Scene) calc(x, y int, eye objects.Point, objs []objects.Object) color.Color {
	var (
		k   float64     = -1
		col color.Color = color.Black
		v               = objects.Vector{
			X: float64(1000 - eye.X),
			Y: float64(s.Width/2 - x - eye.Y),
			Z: float64(s.Height/2 - y - eye.Z),
		}
	)
	for _, obj := range objs {
		// If k == -1, it is our first pass, so if we have a solution, keep it.
		// After that, we check that the solution is smaller than the one we have.
		if tmp := obj.Intersect(v, eye); tmp > 0 && (k == -1 || tmp < k) {
			k = tmp
			col = obj.Color()
		}
	}
	return col
}

// Compute process the Scene with the given Camera (Eye)
// and the given Object list.
func (s *Scene) Compute(eye objects.Point, objs []objects.Object) {
	if true || s.Verbose {
		start := time.Now().UTC()
		defer func() { fmt.Printf("compute time: %s\n", time.Now().UTC().Sub(start)) }()
	}
	sem := make(chan struct{}, s.MaxGoRoutines)
	for i := 0; i < s.MaxGoRoutines; i++ {
		sem <- struct{}{}
	}
	for x := 0; x < s.Width; x++ {
		<-sem
		go func(x int) {
			for y := 0; y < s.Height; y++ {
				if s.Verbose && x == 0 && y%10 == 0 {
					fmt.Printf("\rProcessing: %d%%", int((float64(y)/float64(s.Height))*100+1))
				}
				s.Img.Set(x, y, s.calc(x, y, eye, objs))
			}
			sem <- struct{}{}
		}(x)
	}
	// Make sure all routines are finished
	for i := 0; i < s.MaxGoRoutines; i++ {
		<-sem
	}
	close(sem)
	if s.Verbose {
		fmt.Printf("\rProcessing: 100%%\n")
	}
}
