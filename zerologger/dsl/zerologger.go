package dsl

import (
	"github.com/hirosassa/goaplugin/zerologger/expr"

	// Register code generators for the zerologger plugin
	_ "github.com/hirosassa/goaplugin/zerologger"
)

func HealthCheckPaths(paths ...string) {
	hexpr := &expr.HealthCheckExpr{Paths: paths}
	expr.Root.HealthChecks = append(expr.Root.HealthChecks, hexpr)
}
