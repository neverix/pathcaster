package surfaces

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/neverix/pathcaster/shaders"
	"github.com/neverix/pathcaster/util"
)

// ParseOBJFile reads a model from an OBJ file
func ParseOBJFile(path string) (model *Model, err error) {
	file, err := os.Open(path)
	if err != nil {
		return
	}
	reader := bufio.NewReader(file)
	model, err = ParseOBJ(reader)
	return
}

// ParseOBJ reads a model from an OBJ file reader
func ParseOBJ(reader *bufio.Reader) (model *Model, err error) {
	model = new(Model)
	model.Shader = &shaders.DiffuseShader{} // TODO
	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSuffix(line, "\n")
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
		if strings.HasPrefix(line, "v ") {
			parseVertex(line, model)
		}
		if strings.HasPrefix(line, "f ") {
			parseFace(line, model)
		}
	}
	return model, nil
}

func parseVertex(line string, model *Model) {
	nums := strings.Split(strings.TrimPrefix(line, "v "), " ")
	model.Vertices = append(model.Vertices, util.Vec{
		X: parseFloat(nums[0]),
		Y: parseFloat(nums[1]),
		Z: parseFloat(nums[2])})
}

func parseFace(line string, model *Model) {
	nums := strings.Split(strings.TrimPrefix(line, "f "), " ")
	model.Faces = append(model.Faces, Face{
		A: faceElem(nums[0], 0),
		B: faceElem(nums[1], 0),
		C: faceElem(nums[2], 0)})
}

func faceElem(faceNum string, index int64) int64 {
	return parseInt(strings.Split(faceNum, "/")[index])
}

func parseFloat(num string) float64 {
	float, err := strconv.ParseFloat(num, 64)
	if err != nil {
		handleConversionError(err.(*strconv.NumError))
	}
	return float
}

func parseInt(num string) int64 {
	integer, err := strconv.ParseInt(num, 10, 64)
	if err != nil {
		handleConversionError(err.(*strconv.NumError))
	}
	return integer
}

func handleConversionError(err *strconv.NumError) {
	fmt.Printf("OBJ conversion error: %s", err)
	os.Exit(1)
}
