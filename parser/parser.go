package parser

import (
	"fmt"
	"github.com/Waffle-osu/osu-parser/osu_parser"
)

func main() {
	fmt.Println("Initialisation of the map")
	file, err := osu_parser.ParseFile("map.osu")
	if err != nil {
		return
	}
	fmt.Println(file)
	for i := 0; i < min(len(file.HitObjects.List), 20); i++ {
		fmt.Println(file.HitObjects.List[i])
	}
	fmt.Println(file.HitObjects.List)
}
