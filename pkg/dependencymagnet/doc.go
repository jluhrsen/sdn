// +build tools

// go mod won't pull in code that isn't depended upon, but we have some code we don't depend on from code that must be included
// for our build to work.
package dependencymagnet

import (
	_ "github.com/containernetworking/plugins/plugins/ipam/host-local"
	_ "github.com/openshift/library-go/alpha-build-machinery"
	_ "k8s.io/kubernetes/cmd/kube-proxy"
)