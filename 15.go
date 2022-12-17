package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type P struct {
	X, Y int
}

func (p *P) String() string {
	return fmt.Sprintf("%d,%d", p.X, p.Y)
}

func main() {
	bytes, err := os.ReadFile("15.in")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(bytes), "\n")

	inputRegexp := regexp.MustCompile(`Sensor at x=(.+), y=(.+): closest beacon is at x=(.+), y=(.+)`)

	sensors := make([]*P, len(lines))
	beacons := make([]*P, len(lines))
	dists := make([]int, len(lines))
	maxDist := 0
	minSensorX, maxSensorX := -1, -1
	for i, line := range lines {
		result := inputRegexp.FindStringSubmatch(line)

		sensors[i] = &P{num(result[1]), num(result[2])}
		beacons[i] = &P{num(result[3]), num(result[4])}
		dists[i] = dist(sensors[i], beacons[i])

		d := dist(sensors[i], beacons[i])
		if d > maxDist {
			maxDist = d
		}

		if minSensorX == -1 || sensors[i].X < minSensorX {
			minSensorX = sensors[i].X
		}
		if maxSensorX == -1 || sensors[i].X > maxSensorX {
			maxSensorX = sensors[i].X
		}
	}

	beaconMap := make(map[string]interface{})
	for _, beacon := range beacons {
		beaconMap[beacon.String()] = nil
	}

	noBeaconPositions := part1(minSensorX, maxSensorX, maxDist, beacons, sensors, beaconMap)
	fmt.Printf("part 1, noBeaconPositions: %d\n", noBeaconPositions)

	freq := part2(beacons, sensors, dists)
	fmt.Printf("part 2, tuning frequency: %d\n", freq)
}

func num(s string) int {
	num, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return num
}

func dist(a *P, b *P) int {
	return int(math.Abs(float64(a.X-b.X)) + math.Abs(float64(a.Y-b.Y)))
}

func part1(
	minSensorX int,
	maxSensorX int,
	maxDist int,
	beacons []*P,
	sensors []*P,
	beaconMap map[string]interface{},
) int {
	noBeaconPositions := 0
	for x := minSensorX - maxDist; x <= maxSensorX+maxDist; x++ {
		p := &P{x, 2000000}
		if tooCloseToASensor(p, beacons, sensors, beaconMap) {
			noBeaconPositions++
		}
	}

	return noBeaconPositions

}

func part2(beacons []*P, sensors []*P, dists []int) int {
	beaconMap := make(map[string]interface{})
	for _, beacon := range beacons {
		beaconMap[beacon.String()] = nil
	}

	maxCoord := 4000000

	for sensorIdx, sensor := range sensors {
		fmt.Printf("traversing sensor %d/%d\n", sensorIdx, len(sensors))
		d := dists[sensorIdx]

		for i := 0; i <= d+1; i++ {
			sides := []*P{
				{sensor.X - i, sensor.Y - (d + 1 - i)},
				{sensor.X - i, sensor.Y + (d + 1 - i)},
				{sensor.X + i, sensor.Y - (d + 1 - i)},
				{sensor.X + i, sensor.Y + (d + 1 - i)},
			}
			for _, p := range sides {
				if p.X >= 0 && p.X <= maxCoord &&
					p.Y >= 0 && p.Y <= maxCoord &&
					!tooCloseToASensor(p, beacons, sensors, beaconMap) {
					return p.X*maxCoord + p.Y
				}
			}
		}
	}

	return -1
}

func tooCloseToASensor(p *P, beacons []*P, sensors []*P, beaconMap map[string]interface{}) bool {
	if _, ok := beaconMap[p.String()]; ok {
		return false
	}

	for i := range sensors {
		if dist(p, sensors[i]) <= dist(beacons[i], sensors[i]) {
			// fmt.Printf("point %s too close (%d) to sensor %s (its closest beacon is %s (dist %d))\n",
			// 	p, dist(p, sensors[i]), sensors[i], beacons[i], dist(beacons[i], sensors[i]))
			return true
		}
	}

	return false
}
