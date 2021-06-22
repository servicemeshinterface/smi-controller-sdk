package split

import (
	"context"
	"testing"
	"time"

	splitv1alpha4 "github.com/servicemeshinterface/smi-controller-sdk/apis/split/v1alpha4"
	"github.com/stretchr/testify/require"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Define utility constants for object names and testing timeouts/durations and intervals.
const (
	SplitName      = "test-trafficsplit"
	SplitNamespace = "default"
	ServiceName    = "test-service"

	timeout  = time.Second * 10
	duration = time.Second * 10
	interval = time.Millisecond * 250
)

func testCreateTrafficSplit(t *testing.T) {
	split := &splitv1alpha4.TrafficSplit{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "split.smi-spec.io/v1alpha",
			Kind:       "TrafficSplit",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      SplitName,
			Namespace: SplitNamespace,
		},
		Spec: splitv1alpha4.TrafficSplitSpec{
			Service: ServiceName,
			Backends: []splitv1alpha4.TrafficSplitBackend{
				splitv1alpha4.TrafficSplitBackend{
					Service: ServiceName,
					Weight:  100,
				},
			},
		},
	}

	ctx := context.Background()
	err := k8sClient.Create(ctx, split)
	require.NoError(t, err)
}

func testDeleteTrafficSplit(t *testing.T) {
	split := &splitv1alpha4.TrafficSplit{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "split.smi-spec.io/v1alpha",
			Kind:       "TrafficSplit",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      SplitName,
			Namespace: SplitNamespace,
		},
	}

	ctx := context.Background()
	err := k8sClient.Delete(ctx, split)
	require.NoError(t, err)
}
