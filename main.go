package main

import (
	"fmt"
	"time"

	"github.com/lookuptable/kubernetes-go-client-samples/client"

	"k8s.io/client-go/pkg/api/v1"
)

func main() {
	clientset, err := client.NewClientSet()
	if err != nil {
		panic(err.Error())
	}

	for {
		pods, err := clientset.CoreV1().Pods("").List(v1.ListOptions{})
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))
		time.Sleep(10 * time.Second)
	}
}
