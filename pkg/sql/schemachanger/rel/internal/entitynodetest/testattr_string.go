// Code generated by "stringer --type testAttr --tags test"; DO NOT EDIT.

package entitynodetest

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[i8-0]
	_ = x[pi8-1]
	_ = x[i16-2]
	_ = x[value-3]
	_ = x[left-4]
	_ = x[right-5]
}

const _testAttr_name = "i8pi8i16valueleftright"

var _testAttr_index = [...]uint8{0, 2, 5, 8, 13, 17, 22}

func (i testAttr) String() string {
	if i < 0 || i >= testAttr(len(_testAttr_index)-1) {
		return "testAttr(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _testAttr_name[_testAttr_index[i]:_testAttr_index[i+1]]
}
