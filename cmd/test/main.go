package main

import (
	"context"
	"flag"
	"fmt"
	"path/filepath"

	v1 "github.com/sakiib/crd/pkg/apis/book.com/v1alpha1"
	crd "github.com/sakiib/crd/pkg/client/clientset/versioned"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func main() {
	config, err := getRestConfig()
	if err != nil {
		panic(err)
	}

	crdClient := crd.NewForConfigOrDie(config)
	bookApi, err := crdClient.BookV1alpha1().BookAPIs("default").Create(context.Background(), &v1.BookAPI{}, metav1.CreateOptions{})
	if err != nil {
		panic(err)
	}

	fmt.Println(bookApi.Spec)
}

func getRestConfig() (*rest.Config, error) {
	var kubeConfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeConfig = flag.String("kubeConfig", filepath.Join(home, ".kube", "config"), "absolute path to the kubeConfig file")
	} else {
		kubeConfig = flag.String("kubeConfig", "", "absolute path to the kubeConfig file")
	}

	flag.Parse()
	config, err := clientcmd.BuildConfigFromFlags("", *kubeConfig)
	if err != nil {
		return nil, err
	}
	return config, nil
}
