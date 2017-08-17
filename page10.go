<<<<<<< HEAD
package main

import (
	"image/color"
	"image"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	"time"
	"net/http"
	"log"
)

var palette = []color.Color{color.White,color.Black}

const (
	whiteIndex = 0
	blackIndex = 1
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	if len(os.Args)>1 && os.Args[1]=="web"{
		handler:= func(w http.ResponseWriter,r *http.Request){
			lissajous(w)
		}



		http.HandleFunc("/",handler)
		log.Fatal(http.ListenAndServe("localhost:8080",nil))
		return
	}
	lissajous(os.Stdout)
}

func lissajous(out io.Writer){
	const (
		cycle =5
		res = 0.001
		size = 100
		nframes = 64
		delay = 8
	)

	freq:=rand.Float64()*3.0
	anim:=gif.GIF{LoopCount:nframes}
	phase:=0.0
	for i := 0; i < nframes; i++ {
		rect :=image.Rect(0,0,2*size+1,2*size+1)
		img:=image.NewPaletted(rect,palette)
		for t:=0.0;t<cycle*2*math.Pi;t+=res{
			x:=math.Sin(t)
			y:=math.Sin(t*freq+phase)
			img.SetColorIndex(size+int(x*size+0.5),size+int(y*size+0.5),blackIndex)
		}
		phase+=0.1
		anim.Delay=append(anim.Delay,delay)
		anim.Image = append(anim.Image,img)

	}
	gif.EncodeAll(out,&anim)

}

//运行命令：go build page10.go
=======
package main

import (
	"image/color"
	"image"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	"time"
	"net/http"
	"log"
)

var palette = []color.Color{color.White,color.Black}

const (
	whiteIndex = 0
	blackIndex = 1
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	if len(os.Args)>1 && os.Args[1]=="web"{
		handler:= func(w http.ResponseWriter,r *http.Request){
			lissajous(w)
		}



		http.HandleFunc("/",handler)
		log.Fatal(http.ListenAndServe("localhost:8080",nil))
		return
	}
	lissajous(os.Stdout)
}

func lissajous(out io.Writer){
	const (
		cycle =5
		res = 0.001
		size = 100
		nframes = 64
		delay = 8
	)

	freq:=rand.Float64()*3.0
	anim:=gif.GIF{LoopCount:nframes}
	phase:=0.0
	for i := 0; i < nframes; i++ {
		rect :=image.Rect(0,0,2*size+1,2*size+1)
		img:=image.NewPaletted(rect,palette)
		for t:=0.0;t<cycle*2*math.Pi;t+=res{
			x:=math.Sin(t)
			y:=math.Sin(t*freq+phase)
			img.SetColorIndex(size+int(x*size+0.5),size+int(y*size+0.5),blackIndex)
		}
		phase+=0.1
		anim.Delay=append(anim.Delay,delay)
		anim.Image = append(anim.Image,img)

	}
	gif.EncodeAll(out,&anim)

}

//运行命令：go build page10.go
>>>>>>> 24181879451b8db8a6db4d439894d3ade71ca404
//go run page10.go >out.gif