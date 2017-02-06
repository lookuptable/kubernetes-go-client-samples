package main

import (
	"flag"
	"fmt"

	"github.com/lookuptable/kubernetes-go-client-samples/client"

	"k8s.io/client-go/pkg/api/v1"
)

func main() {
	flag.Parse()

	clientset, err := client.NewClientSet()
	if err != nil {
		panic(err.Error())
	}

	w, err := clientset.CoreV1().Pods("").Watch(v1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}

	c := w.ResultChan()
	for {
		event := <-c
		fmt.Println(event)
	}
}
