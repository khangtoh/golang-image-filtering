package main

import(
  "image"
  "image/jpeg"
  "fmt"
  "os"
  "github.com/disintegration/gift"
)

func main(){

  if (len(os.Args) != 2){
    fmt.Println("Usage:\tgoimger <file>")
    os.Exit(1)
  }

  srcFileName := os.Args[1]
  srcFile, _ := os.Open(srcFileName)
  src, _, _ := image.Decode(srcFile)

  // let's make a new gift
  g := gift.New(
    gift.Grayscale(),
    gift.UnsharpMask(1.0, 0.5, 0.0),
  )

  // dest - output image
  dest := image.NewRGBA(g.Bounds(src.Bounds()))
  // draw result
  g.Draw(dest, src)

  outFileName := srcFileName + "_goimger.jpg"
  toimg, _ := os.Create(outFileName)
  defer toimg.Close()

  jpeg.Encode(toimg, dest, &jpeg.Options{jpeg.DefaultQuality})
}



