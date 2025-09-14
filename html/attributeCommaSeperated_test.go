package html

import (
	"fmt"
)

func ExampleAccept() {
	el := Input(Attributes(TypeFile(), Accept("image/png", "image/jpeg")))
	fmt.Println(el.String())
	// Output: <input type="file" accept="image/png,image/jpeg">
}

func ExampleImageSrcSet() {
	el := Link(Attributes(Rel("preload"), ImageSrcSet("img.png", "img-2x.png 2x")))
	fmt.Println(el.String())
	// Output: <link rel="preload" imagesrcset="img.png,img-2x.png 2x">
}

func ExampleSrcSet() {
	el := Img(Attributes(SrcSet("img.png", "img-2x.png 2x"), Src("fallback.png")))
	fmt.Println(el.String())
	// Output: <img srcset="img.png,img-2x.png 2x" src="fallback.png">
}
