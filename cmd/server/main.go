package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/k0kubun/pp"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	// Use cli
	kubeconfig := filepath.Join(os.Getenv("HOME"), ".kube", "config")
	config, err := clientcmd.BuildConfigFromFlags("https://ops.sudojoss.com:6443", kubeconfig)
	if err != nil {
		log.Fatal(err)
	}
	// Use rest api
	// config := &rest.Config{}
	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}

	pods, err := clientSet.CoreV1().Pods("default").List(metav1.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}

	pp.Println(pods)
}
