package controllers

import (
	"context"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	splitv1alpha1 "github.com/servicemeshinterface/smi-sdk-go/pkg/apis/split/v1alpha1"
	"k8s.io/apimachinery/pkg/api/resource"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var _ = Describe("Traffic Split Controller", func() {

	// Define utility constants for object names and testing timeouts/durations and intervals.
	const (
		SplitName      = "test-trafficsplit"
		SplitNamespace = "default"
		ServiceName    = "test-service"

		timeout  = time.Second * 10
		duration = time.Second * 10
		interval = time.Millisecond * 250
	)

	Context("When adding a TrafficSplit", func() {
		It("Should register correctly", func() {
			By("By creating a new TrafficSplit")
			ctx := context.Background()
			split := &splitv1alpha1.TrafficSplit{
				TypeMeta: metav1.TypeMeta{
					APIVersion: "split.smi-spec.io/v1alpha",
					Kind:       "TrafficSplit",
				},
				ObjectMeta: metav1.ObjectMeta{
					Name:      SplitName,
					Namespace: SplitNamespace,
				},
				Spec: splitv1alpha1.TrafficSplitSpec{
					Service: ServiceName,
					Backends: []splitv1alpha1.TrafficSplitBackend{
						splitv1alpha1.TrafficSplitBackend{
							Service: ServiceName,
							Weight:  resource.NewQuantity(100, resource.BinarySI),
						},
					},
				},
			}
			Expect(k8sClient.Create(ctx, split)).Should(Succeed())
		})
	})
})
