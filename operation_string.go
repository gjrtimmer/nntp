// Code generated by "stringer -type Operation"; DO NOT EDIT

package nntp

import "fmt"

const _Operation_name = "CHECKFETCH"

var _Operation_index = [...]uint8{0, 5, 10}

func (i Operation) String() string {
	if i >= Operation(len(_Operation_index)-1) {
		return fmt.Sprintf("Operation(%d)", i)
	}
	return _Operation_name[_Operation_index[i]:_Operation_index[i+1]]
}
