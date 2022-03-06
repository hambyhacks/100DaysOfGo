package robotname

import (
	"errors"
	"fmt"
	"math/rand"
)

type Robot struct {
	RobotName string
}

const Names = "ABCDEFGHIJKLMNOPQRSTUVWYXZ"

var UsedNames = make(map[string]bool)

func (r *Robot) Name() (string, error) {
	// Check if robot name is not equal to empty string.
	if r.RobotName != "" {
		return r.RobotName, nil
	}
	// Check if all the names are used.
	if len(UsedNames) == 676000 {
		return "Error!", errors.New("all names used")
	}
	// Check if the generated name is unique
	if _, ok := UsedNames[r.Generate()]; !ok {
		UsedNames[r.RobotName] = true
	} else {
		r.RobotName, _ = r.Name()
	}
	return r.RobotName, nil

}

func (r *Robot) Reset() {
	// Uses random seed to generate different results.
	//rand.Seed(time.Now().UnixNano())
	delete(UsedNames, r.RobotName)
	r.RobotName = ""
}

func (r *Robot) Generate() string {
	r.RobotName = string(Names[rand.Intn(26)]) + string(Names[rand.Intn(26)]) + fmt.Sprintf("%03d", rand.Intn(1000))
	return string(r.RobotName)
}
