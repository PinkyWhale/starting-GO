package main

import "fmt"

type Curso struct {
	Titulo string
	Videos []Video
}

type Video struct {
	Titulo string
	Curso  Curso
}

func main() {
	video1 := Video{Titulo: "1-Video Intro"}
	video2 := Video{Titulo: "2-Instalacion"}

	//arreglo
	videos := []Video{video1, video2}

	//objeto
	curso := Curso{Titulo: "Curso Prof de GO!", Videos: videos} //relacion 1 a muchos

	video1.Curso = curso
	video2.Curso = curso

	fmt.Println(video1.Curso.Titulo)

	for _, video := range curso.Videos {
		fmt.Println(video.Titulo)
	}

}
