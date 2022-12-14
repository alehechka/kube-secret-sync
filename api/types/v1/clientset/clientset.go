package clientset

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
)

// KubeSecretSyncClient represents the REST client for kube-secret-sync
type KubeSecretSyncClientset struct {
	client rest.Interface
}

// NewForConfig creates a REST Client for the kube-secret-sync CustomResourceDefinitions
func NewForConfig(c *rest.Config) (*KubeSecretSyncClientset, error) {
	AddToScheme(scheme.Scheme)

	config := *c
	config.ContentConfig.GroupVersion = &schema.GroupVersion{Group: GroupName, Version: GroupVersion}
	config.APIPath = "/apis"
	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()
	config.UserAgent = rest.DefaultKubernetesUserAgent()

	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}

	return &KubeSecretSyncClientset{client: client}, nil
}

func (c *KubeSecretSyncClientset) SecretSyncRules() SecretSyncRuleInterface {
	return newSecretSyncRules(c)
}
