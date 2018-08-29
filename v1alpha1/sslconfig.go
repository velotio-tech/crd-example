package v1alpha1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
)

func (c *SslConfigV1Alpha1Client) SslConfigs(namespace string) SslConfigInterface {
	return &sslConfigclient{
		client: c.restClient,
		ns:     namespace,
	}
}

type SslConfigV1Alpha1Client struct {
	restClient rest.Interface
}

type SslConfigInterface interface {
	Create(obj *SslConfig) (*SslConfig, error)
	Update(obj *SslConfig) (*SslConfig, error)
	Delete(name string, options *meta_v1.DeleteOptions) error
	Get(name string) (*SslConfig, error)
}

type sslConfigclient struct {
	client rest.Interface
	ns     string
}

func (c *sslConfigclient) Create(obj *SslConfig) (*SslConfig, error) {
	result := &SslConfig{}
	err := c.client.Post().
		Namespace(c.ns).Resource("sslconfigs").
		Body(obj).Do().Into(result)
	return result, err
}

func (c *sslConfigclient) Update(obj *SslConfig) (*SslConfig, error) {
	result := &SslConfig{}
	err := c.client.Put().
		Namespace(c.ns).Resource("sslconfigs").
		Body(obj).Do().Into(result)
	return result, err
}

func (c *sslConfigclient) Delete(name string, options *meta_v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).Resource("sslconfigs").
		Name(name).Body(options).Do().
		Error()
}

func (c *sslConfigclient) Get(name string) (*SslConfig, error) {
	result := &SslConfig{}
	err := c.client.Get().
		Namespace(c.ns).Resource("sslconfigs").
		Name(name).Do().Into(result)
	return result, err
}
