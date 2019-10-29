// Code generated by "string2enum -linecomment -type SelectionKind ../../ir/enum"; DO NOT EDIT.

package enum

import (
	"fmt"

	"github.com/llir/llvm/ir/enum"
)

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the string2enum command to generate them again.
	var x [1]struct{}
	_ = x[enum.SelectionKindAny-0]
	_ = x[enum.SelectionKindExactMatch-1]
	_ = x[enum.SelectionKindLargest-2]
	_ = x[enum.SelectionKindNoDuplicates-3]
	_ = x[enum.SelectionKindSameSize-4]
}

const _SelectionKind_name = "anyexactmatchlargestnoduplicatessamesize"

var _SelectionKind_index = [...]uint8{0, 3, 13, 20, 32, 40}

func SelectionKindFromString(s string) enum.SelectionKind {
	if len(s) == 0 {
		return 0
	}
	for i := range _SelectionKind_index[:len(_SelectionKind_index)-1] {
		if s == _SelectionKind_name[_SelectionKind_index[i]:_SelectionKind_index[i+1]] {
			return enum.SelectionKind(i)
		}
	}
	panic(fmt.Errorf("unable to locate SelectionKind enum corresponding to %q", s))
}
