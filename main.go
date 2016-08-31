package main

import (
	"flag"
	"fmt"
	"os"

	client "k8s.io/kubernetes/pkg/client/unversioned"
	"k8s.io/kubernetes/pkg/client/unversioned/clientcmd"
)

func newKubeClient() (*client.Client, error) {
	loadingRules := clientcmd.NewDefaultClientConfigLoadingRules()
	loadingRules.ExplicitPath = clientcmd.RecommendedHomeFile

	loader := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(loadingRules, &clientcmd.ConfigOverrides{})

	clientConfig, err := loader.ClientConfig()

	if err != nil {
		return nil, err
	}

	kubeClient, err := client.New(clientConfig)

	if err != nil {
		return nil, err
	}

	return kubeClient, nil
}

func main() {

	fs := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	var (
		pod   = fs.String("p", "", "help message for long")
		image = fs.String("i", "", "help message for long")
	)
	fs.Parse(os.Args[2:])

	var params = map[string]string{
		"subCommand": os.Args[1],
		"image":      *image,
		"pod":        *pod,
	}

	kubeClient, err := newKubeClient()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	cli(kubeClient, params)

}
