// +build !windows

package env

import (
	log "github.com/newrelic/newrelic-diagnostics-cli/logger"
	"github.com/newrelic/newrelic-diagnostics-cli/tasks"
)

// RegisterWith - will register any plugins in this package
func RegisterWith(registrationFunc func(tasks.Task, bool)) {
	log.Debug("Registering Java/Env/*")

	registrationFunc(JavaEnvVersion{}, true)
	registrationFunc(JavaEnvProcess{
		findProcByName: tasks.FindProcessByName,
		getCmdLineArgs: getCmdLineArgs,
		getCurrentDir:  getCurrentDir}, true)
}
