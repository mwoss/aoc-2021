package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"regexp"
	"strconv"
	"strings"
)

type TargetArea struct {
	x1, x2, y1, y2 int
}

type Velocity struct {
	x, y int
}

func convertStrToInt(str string) int {
	converted, err := strconv.Atoi(str)
	if err != nil {
		log.Fatal(err)
	}
	return converted
}

func parseFileContent(content string) TargetArea {
	reg, err := regexp.Compile("([-]?\\d+..[-]?\\d+)")
	if err != nil {
		log.Fatal(err)
	}

	matches := reg.FindAllString(content, -1)
	if len(matches) != 2 {
		log.Fatal("Input file should only contains x and y ranges")
	}

	x := strings.Split(matches[0], "..")
	y := strings.Split(matches[1], "..")

	return TargetArea{
		x1: convertStrToInt(x[0]),
		x2: convertStrToInt(x[1]),
		y1: convertStrToInt(y[0]),
		y2: convertStrToInt(y[1]),
	}
}

func findHighestPositionToReachArea(area TargetArea) int {
	// height decrease by 1 with each step, so maybe we could just calculate the sequence of n..1 numbers?
	n := -area.y1 - 1
	return n * (n + 1) / 2
}

func findNumberOfInitVelocityToReachArea(area TargetArea) int {
	minXVelocity := getMinimalXVelocity(area.x1)
	maxXVelocity := area.x2
	minYVelocity := area.y1
	maxYVelocity := findHighestPositionToReachArea(area)

	numOfInitVelocities := 0

	for x := minXVelocity; x <= maxXVelocity; x++ {
		for y := minYVelocity; y <= maxYVelocity; y++ {
			v := Velocity{x, y}
			if canReachArea(v, area) {
				numOfInitVelocities++
			}
		}
	}

	return numOfInitVelocities
}

func canReachArea(v Velocity, area TargetArea) bool {
    // unoptimized, but it works :3
	currX, currY := 0, 0
	for {
		currX += v.x
		currY += v.y

        if v.x > 0 {
            v.x -= 1
        } else if (v.x < 0) {
            v.x += 1
        }

		v.y -= 1

		if currX >= area.x1 && currX <= area.x2 && currY >= area.y1 && currY <= area.y2 {
			return true
		}

		if currX > area.x2 && currY > area.y2 {
            return false
        }

        if v.x == 0 && !(currX >= area.x1 && currX <= area.x2) {
            return false
        }

        if v.y < 0 && currY < area.y1 {
            return false
        }
	}

	return false
}

func getMinimalXVelocity(x int) int {
	// a(a+1)/2 = x
	// a^2 + a = 2x
	// a^2 + a - 2x = 0
	delta := 1 - 4*(-2*x)
	if delta < 0 {
		log.Fatal("Delta cannot be a negative number")
	}
	if delta == 0 {
		return 0 // roundUp(-b/2a)
	}

	sqrDelta := math.Sqrt(float64(delta))
	x1 := (-1.0 - sqrDelta) / 2.0
	x2 := (-1.0 + sqrDelta) / 2.0

	if x1 > x2 {
		return int(x1)
	}
	return int(x2)
}

func main() {
	rawContent, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	area := parseFileContent(string(rawContent))

	fmt.Println(findHighestPositionToReachArea(area))
	fmt.Println(findNumberOfInitVelocityToReachArea(area))
}
