package replicaset

import "fmt"

// RolesAddCtx describes a context for adding roles.
type RolesAddCtx struct {
	// InstName is an instance name in which add role.
	InstName string
	// GroupName is an instance name in which add role.
	GroupName string
	// ReplicasetName is an instance name in which add role.
	ReplicasetName string
	// IsGlobal is an boolean value if role needs to add in global scope.
	IsGlobal bool
	// RoleName is a role name which needs to add into config.
	RoleName string
	// Force is true when promoting can skip
	// some non-critical checks.
	Force bool
	// Timeout is a timeout for promoting waitings in seconds.
	// Keep int, because it can be passed to the target instance.
	Timeout int
}

// RolesAdder is an interface for adding roles to a replicaset.
type RolesAdder interface {
	// RolesAdder adds role to a replicasets by its name.
	RolesAdd(ctx RolesAddCtx) error
}

// newErrRolesAddByInstanceNotSupported creates a new error that 'roles add' is not
// supported by the orchestrator for a single instance.
func newErrRolesAddByInstanceNotSupported(orchestrator Orchestrator) error {
	return fmt.Errorf("roles add is not supported for a single instance by %q orchestrator",
		orchestrator)
}

// newErrRolesAddByAppNotSupported creates a new error that 'roles add' by URI is not
// supported by the orchestrator for an application.
func newErrRolesAddByAppNotSupported(orchestrator Orchestrator) error {
	return fmt.Errorf("roles add is not supported for an application by %q orchestrator",
		orchestrator)
}

// parseRoles is a function to convert roles type 'any'
// from yaml config. Returns slice of roles and error.
func parseRoles(value any) ([]string, error) {
	sliceVal, ok := value.([]interface{})
	if !ok {
		return []string{}, fmt.Errorf("%v is not a slice", value)
	}
	existingRoles := make([]string, 0, len(sliceVal)+1)
	for _, v := range sliceVal {
		vStr, ok := v.(string)
		if !ok {
			return []string{}, fmt.Errorf("%v is not a string", v)
		}
		existingRoles = append(existingRoles, vStr)
	}
	return existingRoles, nil
}
