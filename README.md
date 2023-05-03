## About Apps

This apps is processor with go native for compress and resize images, before using this apps your image must be in `base64 format`.

### Install Apps

```sh
  $ go get -u github.com/riskifeb/goImageProcessor
```

### Running Gin

First you need to import goImageProcessor package for using goImageProcessor, one simplest example likes the follow `main.go`:

```go
package main

import (
  "fmt"
  "encoding/base64"

  "github.com/riskifeb/goImageProcessor"
)

func main() {

  // this image must be base64 formated
  var image := "/9j/4AAQSkZJRgABAQEASABIAAD/4gIcSUNDX1BST0ZJTEUAAQEAAAIMbGNtcwIQAABtbnRyU..."

  // ImageProcessor({imageBase64}, {width}, {quality})
  imageBytes, err := goImageProcessor.ImageProcessor(image, 500, 50)
  if(err != nil) {
    panic(err)
  }
  //convert to string
  imageResultString := base64.StdEncoding.EncodeToString(imageBytes.Bytes())

  fmt.Println("data:image/jpeg;base64,"+imageResultString)
}
```
