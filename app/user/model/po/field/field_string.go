// Code generated by "stringer -type=Field -linecomment"; DO NOT EDIT.

package field

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Null-0]
	_ = x[UserID-1]
	_ = x[UserName-2]
	_ = x[Password-3]
	_ = x[UserRole-4]
	_ = x[Email-5]
	_ = x[FirstName-6]
	_ = x[LastName-7]
	_ = x[IsDel-8]
	_ = x[End-9]
}

const _Field_name = "nulluser_iduser_namepassworduser_roleemailfirst_namelast_nameis_delend"

var _Field_index = [...]uint8{0, 4, 11, 20, 28, 37, 42, 52, 61, 67, 70}

func (i Field) String() string {
	if i < 0 || i >= Field(len(_Field_index)-1) {
		return "Field(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Field_name[_Field_index[i]:_Field_index[i+1]]
}
