module github.com/oam-dev/kubevela-core-api

go 1.16

require (
	cuelang.org/go v0.5.0-alpha.1
	github.com/crossplane/crossplane-runtime v0.14.1-0.20210722005935-0b469fcc77cd
	github.com/davecgh/go-spew v1.1.1
	github.com/google/go-cmp v0.5.9
	github.com/kubevela/pkg v0.0.0-20230103022800-b6bdefd2e1de
	github.com/kubevela/workflow v0.3.5
	github.com/oam-dev/cluster-gateway v1.7.0-alpha.1
	github.com/oam-dev/terraform-controller v0.7.0
	github.com/onsi/ginkgo/v2 v2.1.6
	github.com/onsi/gomega v1.20.2
	github.com/pkg/errors v0.9.1
	github.com/spf13/pflag v1.0.5
	github.com/stretchr/testify v1.8.0
	k8s.io/api v0.25.3
	k8s.io/apimachinery v0.25.3
	k8s.io/client-go v0.25.3
	k8s.io/klog/v2 v2.70.1
	k8s.io/utils v0.0.0-20220728103510-ee6ede2d64ed
	sigs.k8s.io/controller-runtime v0.12.3
	sigs.k8s.io/controller-tools v0.6.2
	sigs.k8s.io/yaml v1.3.0
)

replace sigs.k8s.io/apiserver-network-proxy/konnectivity-client => sigs.k8s.io/apiserver-network-proxy/konnectivity-client v0.0.24
