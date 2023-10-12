package internal

import (
	"github.com/starter-go/libgorm"
	"github.com/starter-go/vlog"
)

// GroupManagerImpl ...
type GroupManagerImpl struct {

	//starter:component
	_as func(libgorm.GroupManager) //starter:as("#")

	Groups []libgorm.GroupRegistry //starter:inject(".")

	glist []*libgorm.Group
}

func (inst *GroupManagerImpl) _impl() libgorm.GroupManager {
	return inst
}

// ListGroups ...
func (inst *GroupManagerImpl) ListGroups() []*libgorm.Group {
	dst := inst.glist
	if dst == nil {
		dst = inst.load()
		inst.glist = dst
	}
	return dst
}

func (inst *GroupManagerImpl) load() []*libgorm.Group {
	src := inst.Groups
	dst := make([]*libgorm.Group, 0)
	for _, r1 := range src {
		g := r1.Group()
		dst = append(dst, g)
		vlog.Info("load table group: '%s'", g.Name)
	}
	return dst
}
