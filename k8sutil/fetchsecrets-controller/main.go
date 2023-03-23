package main

import (
	"context"
	"fmt"
	"log"

	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
)

var c client.Client

func main() {
	c, err := client.New(config.GetConfigOrDie(), client.Options{})
	if err != nil {
		log.Fatalf("Failed to create client")
	}

	// Using a typed object.
	secret := &corev1.Secret{}

	// c is a created client.
	// c.List(context.Background(), secret, &client.ListOptions{
	// 	Namespace: "tenant-purewebash-ash",
	// })
	// err = c.Get(context.TODO(), client.ObjectKey{
	// 	Namespace: "tenant-purewebash-ash",
	// 	Name:      "app-server-login-test",
	// }, secret)
	// if err != nil {
	// 	fmt.Printf(err.Error())
	// }
	c.Get(context.TODO(), client.ObjectKey{
		Namespace: "tenant-purewebash-ash",
		Name:      "samba-credentials",
	}, secret)
	// if err != nil {
	// 	fmt.Printf(err.Error())
	// }

	fmt.Println(string(secret.Data["samba_password"]))

	// // Using a unstructured object.
	// u := &unstructured.Unstructured{}
	// u.SetGroupVersionKind(schema.GroupVersionKind{
	// 	Group:   "apps",
	// 	Kind:    "Deployment",
	// 	Version: "v1",
	// })
	// _ = c.Get(context.Background(), client.ObjectKey{
	// 	Namespace: "namespace",
	// 	Name:      "name",
	// }, u)
}
