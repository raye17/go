package main

import (
	"fmt"
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"log"
	"math"
	"strings"
)

const (
	width = 500
	height
	VertexShaderSource = `
		     #version 410
		     in vec3 vp;
		     void main() {
		         gl_position = vec4(vp,1.0);
		     }
		` + "\x00"
	FragmentShaderSource = `
		     #version 410
			 out vec4 frag_colour;
			 uniform vec4 FragColor;
			 void main() {
				 frag_colour = FragColor;
			 }
		`
)

var (
	vertices = []float32{
		-0.5, -0.5, 0.0,
		-0.5, 0.5, 0.0,
		0.5, 0.5, 0.0,
		0.5, -0.5, 0.0,
	}
	indices = []uint32{
		0, 1, 2,
		2, 3, 0,
	}
)

func main() {
	//runtime.LockOSThread()
	//window := initGlfw()
	//defer glfw.Terminate()
	//program := initOpenGl()
	//vao := makeVao(vertices, indices)
	//for !window.ShouldClose() {
	//	draw(vao, window, program)
	//}
	//glfw.Terminate()

}
func initGlfw() *glfw.Window {
	//初始化，并设置属性
	if err := glfw.Init(); err != nil {
		panic(err)
	}
	glfw.WindowHint(glfw.Resizable, glfw.False)
	window, err := glfw.CreateWindow(width, height, "Conway's game of life", nil, nil)
	if err != nil {
		panic(err)
	}
	window.MakeContextCurrent()
	return window
}
func initOpenGl() uint32 {
	if err := glfw.Init(); err != nil {
		panic(err)
	}
	version := gl.GoStr(gl.GetString(gl.VERSION))
	log.Println("OpenGl version", version)
	vertexShader, err := CompileShader(VertexShaderSource, gl.VERTEX_SHADER)
	if err != nil {
		panic(err)
	}
	fragmentShader, err := CompileShader(FragmentShaderSource, gl.FRAGMENT_SHADER)
	if err != nil {
		panic(err)
	}
	prog := gl.CreateProgram()
	gl.AttachShader(prog, vertexShader)
	gl.AttachShader(prog, fragmentShader)
	gl.Viewport(0, 0, width, height)
	gl.LinkProgram(prog)
	return prog
}
func makeVao(points []float32, indices []uint32) uint32 {
	var vbo uint32
	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, 4*len(points), gl.Ptr(points), gl.STATIC_DRAW)

	var vao uint32
	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)
	gl.EnableVertexAttribArray(0)
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 0, gl.Ptr(nil))

	if indices != nil {
		var ebo uint32
		gl.GenBuffers(2, &ebo)
		gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, ebo)
		gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, 4*len(indices), gl.Ptr(indices), gl.STATIC_DRAW)

	}
	return vao
}
func draw(vao uint32, window *glfw.Window, program uint32) {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	timeValue := glfw.GetTime()
	greenValue := float32(math.Sin(timeValue)/2.0 + 0.5)
	vertexColorLocation := gl.GetUniformLocation(program, gl.Str("FragColor\x00"))
	gl.UseProgram(program)
	gl.BindVertexArray(vao)
	gl.Uniform4f(vertexColorLocation, 0, greenValue, 0, 1)

	//gl.DrawArrays(gl.TRIANGLES, 0, 4)
	gl.DrawElements(gl.TRIANGLES, 6, gl.UNSIGNED_INT, gl.PtrOffset(0))
	glfw.PollEvents()
	window.SwapBuffers()
}

// CompileShader 编写着色器
func CompileShader(source string, shaderType uint32) (uint32, error) {
	shader := gl.CreateShader(shaderType)
	csources, free := gl.Strs(source)
	gl.ShaderSource(shader, 1, csources, nil)
	free()
	gl.CompileShader(shader)
	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)
		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(log))
		return 0, fmt.Errorf("failed to compile %v: %v", source, log)
	}
	return shader, nil
}
