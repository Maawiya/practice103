package main

import (
	"fmt"
	"image"
	"image/color"
	"log"
	"os"
	"path/filepath"

	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
	"github.com/disintegration/imaging"
)

// Function to add a text watermark to an image
func addTextWatermark(img image.Image, watermark string) (image.Image, error) {
	// Clone the original image into an RGBA image for drawing
	rgbaImg := imaging.Clone(img)

	// Get image bounds (width and height)
	bounds := img.Bounds()
	width, height := bounds.Dx(), bounds.Dy()

	// Create the font face (use a basic font, or change to a custom one)
	c := basicfont.Face7x13

	// Measure the width of the watermark text
	textWidth := font.MeasureString(c, watermark)

	// Set the position for the watermark (bottom-right corner with some margin)
	x := width - int(textWidth) - 10 // 10px margin from the right
	y := height - 10                // 10px margin from the bottom

	// Set up the color for the watermark (white with transparency)
	watermarkColor := color.RGBA{R: 255, G: 255, B: 255, A: 200}

	// Create a drawer to draw the watermark text
	drawer := font.Drawer{
		Dst:  rgbaImg,
		Src:  image.NewUniform(watermarkColor),
		Face: c,
	}
	drawer.Dot = fixed.P(x, y)
	drawer.DrawString(watermark)

	return rgbaImg, nil
}

// Create a watermarked directory if it doesn't exist
func createWatermarkedDirectory() string {
	watermarkedDir := filepath.Join("images", "watermarked")
	err := os.MkdirAll(watermarkedDir, os.ModePerm)
	if err != nil {
		log.Fatalf("Error creating watermarked directory: %v\n", err)
	}
	return watermarkedDir
}

func main() {
	// Get the current working directory
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Current working directory:", dir)

	// Define the input image path and the watermark text
	inputImagePath := filepath.Join(dir, "images", "watermarked", "image2.png")
	watermarkText := "My Watermark"

	// Open the input image
	img, err := imaging.Open(inputImagePath)
	if err != nil {
		log.Fatalf("Error opening image: %v\n", err)
	}

	// Add the text watermark to the image
	imgWithWatermark, err := addTextWatermark(img, watermarkText)
	if err != nil {
		log.Fatalf("Error adding watermark: %v\n", err)
	}

	// Create the watermarked directory
	watermarkedDir := createWatermarkedDirectory()

	// Define the output path for the watermarked image
	outputPath := filepath.Join(watermarkedDir, "watermarked_image.png")
	fmt.Println("Saving watermarked image to:", outputPath)

	// Save the watermarked image
	err = imaging.Save(imgWithWatermark, outputPath)
	if err != nil {
		log.Fatalf("Error saving watermarked image: %v\n", err)
	}

	fmt.Println("Watermark added and image saved successfully!")
}
