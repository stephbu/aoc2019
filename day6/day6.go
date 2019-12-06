package day6

import (
	"log"
	"strings"
)

type Orbit struct {
	CenterOfMass string
	Body         string
	Hops         int
}

func expandOrbits(orbitExpressions []string) (direct []*Orbit, indirect []*Orbit) {

	direct = []*Orbit{}
	indirect = []*Orbit{}

	orbitIndex := make(map[string]*Orbit, 0)

	for index, orbitExpression := range orbitExpressions {

		segments := strings.Split(orbitExpression, ")")
		if len(segments) != 2 {
			log.Fatalf("invalid input line:%v", index)
		}

		orbit := &Orbit{CenterOfMass: segments[0], Body: segments[1], Hops: 1}
		orbitIndex[orbit.Body] = orbit
		direct = append(direct, orbit)

	}

	for _, directOrbit := range direct {

		ok := true
		body := directOrbit

		hops := 1

		for ok {

			body, ok = orbitIndex[body.CenterOfMass]

			if ok {
				hops++
				indirectOrbit := &Orbit{CenterOfMass: body.CenterOfMass, Body: directOrbit.Body, Hops: hops}
				indirect = append(indirect, indirectOrbit)
			}
		}
	}

	return
}

func findShortestPath(orbits []*Orbit, body1 string, body2 string) (int, string) {

	path1 := []*Orbit{}
	path2 := []*Orbit{}

	for _, orbit := range orbits {
		if orbit.Body == body1 {
			path1 = append(path1, orbit)
		}
		if orbit.Body == body2 {
			path2 = append(path2, orbit)
		}
	}

	shortestPath := 0
	shortestCommonCOM := ""

	for _, orbitPath1 := range path1 {
		for _, orbitPath2 := range path2 {
			if orbitPath1.CenterOfMass == orbitPath2.CenterOfMass {

				// shortest path starts one away from each arm on a common center of mass
				totalHops := orbitPath1.Hops + orbitPath2.Hops - 2
				if shortestPath == 0 || (totalHops < shortestPath) {
					shortestPath = totalHops
					shortestCommonCOM = orbitPath1.CenterOfMass
				}

			} else if orbitPath1.CenterOfMass == orbitPath2.Body {

				// shortest path starts one away from each arm on a common center of mass
				shortestPath = orbitPath1.Hops - 1
				shortestCommonCOM = orbitPath1.CenterOfMass

			} else if orbitPath2.CenterOfMass == orbitPath1.Body {

				// shortest path starts one away from each arm on a common center of mass
				shortestPath = orbitPath2.Hops - 1
				shortestCommonCOM = orbitPath2.CenterOfMass

			}
		}
	}

	return shortestPath, shortestCommonCOM

}
