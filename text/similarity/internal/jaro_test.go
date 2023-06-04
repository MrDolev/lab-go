package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/** Distance
  distance.apply(null, null)          = IllegalArgumentException
  distance.apply("foo", null)         = IllegalArgumentException
  distance.apply(null, "foo")         = IllegalArgumentException
  distance.apply("", "")              = 0.0
  distance.apply("foo", "foo")        = 0.0
  distance.apply("foo", "foo ")       = 0.06
  distance.apply("foo", "foo  ")      = 0.09
  distance.apply("foo", " foo ")      = 0.13
  distance.apply("foo", "  foo")      = 0.49
  distance.apply("", "a")             = 1.0
  distance.apply("aaapppp", "")       = 1.0
  distance.apply("frog", "fog")       = 0.07
  distance.apply("fly", "ant")        = 1.0
  distance.apply("elephant", "hippo") = 0.56
  distance.apply("hippo", "elephant") = 0.56
  distance.apply("hippo", "zzzzzzzz") = 1.0
  distance.apply("hello", "hallo")    = 0.12
  distance.apply("ABC Corporation", "ABC Corp") = 0.09
  distance.apply("D N H Enterprises Inc", "D & H Enterprises, Inc.") = 0.05
  distance.apply("My Gym Children's Fitness Center", "My Gym. Childrens Fitness") = 0.08
  distance.apply("PENNSYLVANIA", "PENNCISYLVNIA") = 0.12
*/

type TestMockedDistance struct {
	str1 string
	str2 string
	dist float64
}

func TestDistance(t *testing.T) {
	mockedDistiance := getMockedDistances()

	for _, mockedDst := range mockedDistiance {
		res := Distance(mockedDst.str1, mockedDst.str2)
		if res != mockedDst.dist {
			t.Errorf("JaroWinkler('%s', '%s') = %v, want %v", mockedDst.str1, mockedDst.str2, res, mockedDst.dist)
		}
		assert.Equal(t, mockedDst.dist, res, "Distance is equals")
	}
}

func getMockedDistances() []TestMockedDistance {
	var mockedDistiance []TestMockedDistance
	mockedDistiance = append(
		mockedDistiance,
		TestMockedDistance{"", "", 0.0},
	)
	mockedDistiance = append(
		mockedDistiance,
		TestMockedDistance{"foo", "foo", 0.0},
	)
	mockedDistiance = append(
		mockedDistiance,
		TestMockedDistance{"foo", "foo ", 0.05833333333333335},
	)
	mockedDistiance = append(
		mockedDistiance,
		TestMockedDistance{"foo", "foo  ", 0.09333333333333327},
	)
	mockedDistiance = append(
		mockedDistiance,
		TestMockedDistance{"foo", " foo ", 0.12},
	)
	mockedDistiance = append(
		mockedDistiance,
		TestMockedDistance{"foo", "  foo", 0.48888888888888893},
	)
	mockedDistiance = append(
		mockedDistiance,
		TestMockedDistance{"", "a", 1.0},
	)
	mockedDistiance = append(
		mockedDistiance,
		TestMockedDistance{"aaapppp", "", 1.0},
	)
	mockedDistiance = append(
		mockedDistiance,
		TestMockedDistance{"frog", "fog", 0.06666666666666665},
	)
	mockedDistiance = append(
		mockedDistiance,
		TestMockedDistance{"fly", "ant", 1.0},
	)
	mockedDistiance = append(
		mockedDistiance,
		TestMockedDistance{"elephant", "hippo", 0.5583333333333333},
	)
	mockedDistiance = append(
		mockedDistiance,
		TestMockedDistance{"hippo", "elephant", 0.5583333333333333},
	)
	mockedDistiance = append(
		mockedDistiance,
		TestMockedDistance{"hippo", "zzzzzzzz", 1.0},
	)
	mockedDistiance = append(
		mockedDistiance,
		TestMockedDistance{"hello", "hallo", 0.10666666666666669},
	)

	mockedDistiance = append(
		mockedDistiance,
		TestMockedDistance{"ABC Corporation", "ABC Corp", 0.09333333333333338},
	)

	mockedDistiance = append(
		mockedDistiance,
		TestMockedDistance{"D N H Enterprises Inc", "D & H Enterprises, Inc.", 0.04154589371980677},
	)

	mockedDistiance = append(
		mockedDistiance,
		TestMockedDistance{"My Gym Children's Fitness Center", "My Gym. Childrens Fitness", 0.05800000000000005},
	)
	return mockedDistiance
}

/** Similarity
  sim.apply(null, null)          = IllegalArgumentException
  sim.apply("foo", null)         = IllegalArgumentException
  sim.apply(null, "foo")         = IllegalArgumentException
  sim.apply("", "")              = 1.0
  sim.apply("foo", "foo")        = 1.0
  sim.apply("foo", "foo ")       = 0.94
  sim.apply("foo", "foo  ")      = 0.91
  sim.apply("foo", " foo ")      = 0.87
  sim.apply("foo", "  foo")      = 0.51
  sim.apply("", "a")             = 0.0
  sim.apply("aaapppp", "")       = 0.0
  sim.apply("frog", "fog")       = 0.93
  sim.apply("fly", "ant")        = 0.0
  sim.apply("elephant", "hippo") = 0.44
  sim.apply("hippo", "elephant") = 0.44
  sim.apply("hippo", "zzzzzzzz") = 0.0
  sim.apply("hello", "hallo")    = 0.88
  sim.apply("ABC Corporation", "ABC Corp") = 0.91
  sim.apply("D N H Enterprises Inc", "D & H Enterprises, Inc.") = 0.95
  sim.apply("My Gym Children's Fitness Center", "My Gym. Childrens Fitness") = 0.942
  sim.apply("PENNSYLVANIA", "PENNCISYLVNIA") = 0.8980186480186481
*/

type TestMockedSimilarity struct {
	str1 string
	str2 string
	sim  float64
}

func TestSimilarity(t *testing.T) {

	mockedSimilarities := getMockedSimilarities()

	for _, mockedSim := range mockedSimilarities {
		res := Similarity(mockedSim.str1, mockedSim.str2)
		if res != mockedSim.sim {
			t.Errorf("JaroWinkler('%s', '%s') = %v, want %v", mockedSim.str1, mockedSim.str2, res, mockedSim.sim)
		}
		assert.Equal(t, mockedSim.sim, res, "Distance is equals")
	}
}

func getMockedSimilarities() []TestMockedSimilarity {
	var mockedSimilarities []TestMockedSimilarity
	mockedSimilarities = append(
		mockedSimilarities,
		TestMockedSimilarity{"", "", 1.0},
	)
	mockedSimilarities = append(
		mockedSimilarities,
		TestMockedSimilarity{"foo", "foo", 1.0},
	)
	mockedSimilarities = append(
		mockedSimilarities,
		TestMockedSimilarity{"foo", "foo ", 0.9416666666666667},
	)
	mockedSimilarities = append(
		mockedSimilarities,
		TestMockedSimilarity{"foo", "foo  ", 0.9066666666666667},
	)
	mockedSimilarities = append(
		mockedSimilarities,
		TestMockedSimilarity{"foo", " foo ", 0.88},
	)
	mockedSimilarities = append(
		mockedSimilarities,
		TestMockedSimilarity{"foo", "  foo", 0.5111111111111111},
	)
	mockedSimilarities = append(
		mockedSimilarities,
		TestMockedSimilarity{"", "a", 0.0},
	)
	mockedSimilarities = append(
		mockedSimilarities,
		TestMockedSimilarity{"aaapppp", "", 0.0},
	)
	mockedSimilarities = append(
		mockedSimilarities,
		TestMockedSimilarity{"frog", "fog", 0.9333333333333333},
	)
	mockedSimilarities = append(
		mockedSimilarities,
		TestMockedSimilarity{"fly", "ant", 0.0},
	)
	mockedSimilarities = append(
		mockedSimilarities,
		TestMockedSimilarity{"elephant", "hippo", 0.44166666666666665},
	)
	mockedSimilarities = append(
		mockedSimilarities,
		TestMockedSimilarity{"hippo", "elephant", 0.44166666666666665},
	)
	mockedSimilarities = append(
		mockedSimilarities,
		TestMockedSimilarity{"hippo", "zzzzzzzz", 0.0},
	)
	mockedSimilarities = append(
		mockedSimilarities,
		TestMockedSimilarity{"hello", "hallo", 0.8933333333333333},
	)

	mockedSimilarities = append(
		mockedSimilarities,
		TestMockedSimilarity{"ABC Corporation", "ABC Corp", 0.9066666666666666},
	)

	mockedSimilarities = append(
		mockedSimilarities,
		TestMockedSimilarity{"D N H Enterprises Inc", "D & H Enterprises, Inc.", 0.9584541062801932},
	)

	mockedSimilarities = append(
		mockedSimilarities,
		TestMockedSimilarity{"My Gym Children's Fitness Center", "My Gym. Childrens Fitness", 0.942},
	)

	mockedSimilarities = append(
		mockedSimilarities,
		TestMockedSimilarity{"PENNSYLVANIA", "PENNCISYLVNIA", 0.8980186480186481},
	)
	return mockedSimilarities
}
