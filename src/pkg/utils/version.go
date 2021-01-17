package utils

import (
	"fmt"
	"strconv"
	"strings"
)

// Version ...
type Version struct {
	major  uint
	minor  uint
	patch  uint
	suffix string
}

// NewVersion creates a new Version instance
func NewVersion(major, minor, patch uint) Version {
	return Version{
		major:  major,
		minor:  minor,
		patch:  patch,
		suffix: "",
	}
}

// SetSuffix ...
// TODO: suffix が ascii 以外の文字を含む時は error を返すようにしたい
func (v Version) SetSuffix(suffix string) (Version, error) {
	return Version{
		major:  v.major,
		minor:  v.minor,
		patch:  v.patch,
		suffix: suffix,
	}, nil
}

// String ...
func (v Version) String() string {
	s := fmt.Sprintf("%d.%d.%d", v.major, v.minor, v.patch)
	if v.suffix != "" {
		s += fmt.Sprintf("-%s", v.suffix)
	}
	return s
}

// MarshalJSON ...
func (v Version) MarshalJSON() ([]byte, error) {
	return []byte("\"" + v.String() + "\""), nil
}

// UnmarshalJSON ...
func (v *Version) UnmarshalJSON(rawbyte []byte) error {
	raw := string(rawbyte)
	if !strings.HasPrefix(raw, `"`) || !strings.HasSuffix(raw, `"`) || len(raw) < 2 {
		return fmt.Errorf("must be string")
	}
	naked := strings.Trim(raw, `"`)
	components := strings.SplitN(naked, ".", 3)
	if len(components) > 0 {
		major, err := strconv.ParseUint(components[0], 10, 32)
		if err != nil {
			return err
		}
		v.major = uint(major)
	}
	if len(components) > 1 {
		minor, err := strconv.ParseUint(components[1], 10, 32)
		if err != nil {
			return err
		}
		v.minor = uint(minor)
	}
	if len(components) > 2 {
		patchAndSuffix := strings.SplitN(components[2], "-", 2)
		patch, err := strconv.ParseUint(patchAndSuffix[0], 10, 32)
		if err != nil {
			return err
		}
		v.patch = uint(patch)

		if len(patchAndSuffix) == 2 {
			v.suffix = patchAndSuffix[1]
		}
	}
	return nil
}
