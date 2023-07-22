package main

import (
	"os"

	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/metrics/filters"
)

func main() {
	ctrlOptions := ctrl.Options{
		MetricsFilterProvider: filters.WithAuthenticationAndAuthorization,
		MetricsSecureServing:  true,
	}

	mgr, err := ctrl.NewManager(ctrl.GetConfigOrDie(), ctrlOptions)
	if err != nil {
		os.Exit(1)
	}

	if err := mgr.Start(ctrl.SetupSignalHandler()); err != nil {
		os.Exit(1)
	}
}
