package agents

import (
	"github.com/appscode/kube-mon/agents/coreosprometheusoperator"
	"github.com/appscode/kube-mon/agents/prometheusbuiltin"
	"github.com/appscode/kube-mon/api"
	prom "github.com/coreos/prometheus-operator/pkg/client/monitoring/v1"
	ecs "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset/typed/apiextensions/v1beta1"
	"k8s.io/client-go/kubernetes"
)

func New(at api.AgentType, k8sClient kubernetes.Interface, extClient ecs.ApiextensionsV1beta1Interface, promClient prom.MonitoringV1Interface) api.Agent {
	switch at {
	case api.AgentCoreOSPrometheus, api.DeprecatedAgentCoreOSPrometheus:
		return coreosprometheusoperator.New(at, k8sClient, extClient, promClient)
	case api.AgentPrometheusBuiltin:
		return prometheusbuiltin.New(k8sClient)
	}
	return nil
}
