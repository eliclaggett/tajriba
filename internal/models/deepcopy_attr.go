// generated by deep-copy --pointer-receiver --type Attribute --skip Versions -o deepcopy_attr.go .; DO NOT EDIT.

package models

import (
	"time"
)

// DeepCopy generates a deep copy of *Attribute
func (o *Attribute) DeepCopy() *Attribute {
	var cp Attribute = *o
	if o.DeletedAt != nil {
		cp.DeletedAt = new(time.Time)
		*cp.DeletedAt = *o.DeletedAt
	}
	if o.Val != nil {
		cp.Val = new(string)
		*cp.Val = *o.Val
	}
	if o.Index != nil {
		cp.Index = new(int)
		*cp.Index = *o.Index
	}
	return &cp
}
