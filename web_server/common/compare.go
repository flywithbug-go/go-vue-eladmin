package common

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

type CompareState int

const (
	VersionCompareStateFailed CompareState = iota
	VersionCompareStateGreater
	VersionCompareStateEqual
	VersionCompareStateLess
)

var versionRe = regexp.MustCompile(`[\d.]`)

//版本号规则xx.xx.xx 只能有数字和点
func VersionCompare(version1, version2 string) (CompareState, error) {
	if !checkVersionOK(version1) || !checkVersionOK(version2) {
		return VersionCompareStateFailed, errors.New("version string not ok")
	}
	length1 := len(version1)
	length2 := len(version2)
	min := length1
	if min > length2 {
		min = length2
	}
	versionN1, _ := strconv.Atoi(strings.Replace(version1[0:min], ".", "", -1))
	versionN2, _ := strconv.Atoi(strings.Replace(version2[0:min], ".", "", -1))

	if versionN1 > versionN2 {
		return VersionCompareStateGreater, nil
	} else if versionN1 < versionN2 {
		return VersionCompareStateLess, nil
	}
	if length1 == length2 {
		return VersionCompareStateEqual, nil
	} else if length1 > length2 {
		return VersionCompareStateGreater, nil
	} else {
		return VersionCompareStateLess, nil
	}
	return VersionCompareStateFailed, nil
}

func checkVersionOK(version string) bool {
	if strings.HasPrefix(version, ".") || strings.HasSuffix(version, ".") {
		return false
	}
	if strings.Contains(version, "..") {
		return false
	}
	str := strings.Replace(version, ".", "", -1)
	_, err := strconv.Atoi(str)
	if err != nil {
		return false
	}
	return true
}
