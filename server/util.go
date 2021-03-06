package main

import (
	argoversioned "github.com/argoproj/argo-workflows/v3/pkg/client/clientset/versioned"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/klog"
	"volcano.sh/apis/pkg/client/clientset/versioned"
)

// getKubeClient Get a clientset with restConfig.
func getKubeClient(restConfig *rest.Config) *kubernetes.Clientset {
	clientset, err := kubernetes.NewForConfig(restConfig)
	if err != nil {
		klog.Fatal(err)
	}
	return clientset
}

// GetVolcanoClient get a clientset for volcano.
func getVolcanoClient(restConfig *rest.Config) *versioned.Clientset {
	clientset, err := versioned.NewForConfig(restConfig)
	klog.Info(clientset.ServerVersion())
	if err != nil {
		klog.Fatal(err)
	}
	return clientset
}

// GetArgoClient get a clientset for argo
func getArgoClient(restConfig *rest.Config) *argoversioned.Clientset {
	clientset, err := argoversioned.NewForConfig(restConfig)
	klog.Info(clientset.ServerVersion())
	if err != nil {
		klog.Fatal(err)
	}
	return clientset
}
