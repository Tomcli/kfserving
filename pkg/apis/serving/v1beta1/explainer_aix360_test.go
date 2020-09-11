package v1beta1

import (
	"testing"

	"github.com/golang/protobuf/proto"
	"github.com/kubeflow/kfserving/pkg/constants"
	"github.com/onsi/gomega"
	"github.com/onsi/gomega/types"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
)

func TestAIXExplainer(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	config := InferenceServicesConfig{
		Explainers: ExplainersConfig{
			AIXExplainer: ExplainerConfig{
				ContainerImage:      "aipipeline/aixexplainer",
				DefaultImageVersion: "latest",
			},
		},
	}

	scenarios := map[string]struct {
		spec    AIXExplainerSpec
		matcher types.GomegaMatcher
	}{
		"AcceptGoodRuntimeVersion": {
			spec: AIXExplainerSpec{
				RuntimeVersion: proto.String("latest"),
			},
			matcher: gomega.Succeed(),
		},
	}
	for name, scenario := range scenarios {
		t.Run(name, func(t *testing.T) {
			scenario.spec.Default(&config)
			res := scenario.spec.Validate()
			if !g.Expect(res).To(scenario.matcher) {
				t.Errorf("got %q, want %q", res, scenario.matcher)
			}
		})
	}
}

func TestCreateAIXExplainerContainer(t *testing.T) {

	var requestedResource = v1.ResourceRequirements{
		Limits: v1.ResourceList{
			"cpu": resource.Quantity{
				Format: "100",
			},
		},
		Requests: v1.ResourceList{
			"cpu": resource.Quantity{
				Format: "90",
			},
		},
	}
	config := &InferenceServicesConfig{
		Explainers: ExplainersConfig{
			AIXExplainer: ExplainerConfig{
				ContainerImage:      "aipipeline/aixexplainer",
				DefaultImageVersion: "latest",
			},
		},
	}
	var spec = AIXExplainerSpec{
		Type:       "LimeImages",
		StorageURI: "gs://someUri",
		Container: v1.Container{
			Resources: requestedResource,
		},
		RuntimeVersion: proto.String("0.2.2"),
	}
	g := gomega.NewGomegaWithT(t)

	expectedContainer := &v1.Container{
		Image:     "aipipeline/aixexplainer:0.2.2",
		Name:      constants.InferenceServiceContainerName,
		Resources: requestedResource,
		Args: []string{
			constants.ArgumentModelName,
			"someName",
			constants.ArgumentPredictorHost,
			"predictor.svc.cluster.local",
			constants.ArgumentHttpPort,
			constants.InferenceServiceDefaultHttpPort,
			"--storage_uri",
			"/mnt/models",
			"--explainer_type",
			"LimeImages",
		},
	}

	// Test Create with config
	container := spec.CreateExplainerContainer("someName", 0, "predictor.svc.cluster.local", config)
	g.Expect(container).To(gomega.Equal(expectedContainer))
}

func TestCreateAIXExplainerContainerWithConfig(t *testing.T) {

	var requestedResource = v1.ResourceRequirements{
		Limits: v1.ResourceList{
			"cpu": resource.Quantity{
				Format: "100",
			},
		},
		Requests: v1.ResourceList{
			"cpu": resource.Quantity{
				Format: "90",
			},
		},
	}
	config := &InferenceServicesConfig{
		Explainers: ExplainersConfig{
			AIXExplainer: ExplainerConfig{
				ContainerImage:      "aipipeline/aixexplainer",
				DefaultImageVersion: "latest",
			},
		},
	}
	var spec = AIXExplainerSpec{
		Type:       "LimeImages",
		StorageURI: "gs://someUri",
		Container: v1.Container{
			Resources: requestedResource,
		},
		RuntimeVersion: proto.String("0.2.2"),
		Config: map[string]string{
			"num_classes": "10",
			"num_samples": "20",
			"min_weight":  "0",
		},
	}
	g := gomega.NewGomegaWithT(t)

	expectedContainer := &v1.Container{
		Image:     "aipipeline/aixexplainer:0.2.2",
		Name:      constants.InferenceServiceContainerName,
		Resources: requestedResource,
		Args: []string{
			"--model_name",
			"someName",
			"--predictor_host",
			"predictor.svc.cluster.local",
			"--http_port",
			"8080",
			"--storage_uri",
			"/mnt/models",
			"--explainer_type",
			"LimeImages",
			"--min_weight",
			"0",
			"--num_classes",
			"10",
			"--num_samples",
			"20",
		},
	}

	// Test Create with config
	container := spec.CreateExplainerContainer("someName", 0, "predictor.svc.cluster.local", config)
	g.Expect(container).To(gomega.Equal(expectedContainer))
}
