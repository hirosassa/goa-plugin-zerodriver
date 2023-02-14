package testdata

import (
	. "goa.design/goa/v3/dsl"

	zerologger "github.com/hirosassa/goaplugin/zerologger/dsl"
)

var SimpleServiceDSL = func() {
	zerologger.HealthCheckPaths("/liveness", "/readiness", "/healthz")

	Service("SimpleService", func() {
		Method("SimpleMethod", func() {
			HTTP(func() {
				GET("/")
			})
		})
	})
}
