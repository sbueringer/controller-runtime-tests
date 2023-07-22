package main

import (
	"os"

	ctrl "sigs.k8s.io/controller-runtime"
)

func main() {
	ctrlOptions := ctrl.Options{}

	mgr, err := ctrl.NewManager(ctrl.GetConfigOrDie(), ctrlOptions)
	if err != nil {
		os.Exit(1)
	}

	if err := mgr.Start(ctrl.SetupSignalHandler()); err != nil {
		os.Exit(1)
	}
}
