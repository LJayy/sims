// Code generated by "stringer -type=EnvType"; DO NOT EDIT.

package main

import (
	"errors"
	"strconv"
)

var _ = errors.New("dummy error")

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Probe-0]
	_ = x[Besner-1]
	_ = x[Glushko-2]
	_ = x[Taraban-3]
	_ = x[EnvTypeN-4]
}

const _EnvType_name = "ProbeBesnerGlushkoTarabanEnvTypeN"

var _EnvType_index = [...]uint8{0, 5, 11, 18, 25, 33}

func (i EnvType) String() string {
	if i < 0 || i >= EnvType(len(_EnvType_index)-1) {
		return "EnvType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _EnvType_name[_EnvType_index[i]:_EnvType_index[i+1]]
}

func (i *EnvType) FromString(s string) error {
	for j := 0; j < len(_EnvType_index)-1; j++ {
		if s == _EnvType_name[_EnvType_index[j]:_EnvType_index[j+1]] {
			*i = EnvType(j)
			return nil
		}
	}
	return errors.New("String: " + s + " is not a valid option for type: EnvType")
}
