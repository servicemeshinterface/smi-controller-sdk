package v1alpha1

import (
	"github.com/servicemeshinterface/smi-controller-sdk/apis/access/v1alpha4"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
)

/*
Our "spoke" versions need to implement the
[`Convertible`](https://pkg.go.dev/sigs.k8s.io/controller-runtime/pkg/conversion?tab=doc#Convertible)
interface.  Namely, they'll need `ConvertTo` and `ConvertFrom` methods to convert to/from
the hub version.
*/

/*
ConvertTo is expected to modify its argument to contain the converted object.
Most of the conversion is straightforward copying, except for converting our changed field.
*/
// ConvertTo converts this TrafficTarget to the Hub version (v1alpha4).
func (src *TrafficTarget) ConvertTo(dstRaw conversion.Hub) error {
	traffictargetlog.Info("ConvertTo v1alpha4 from v1alpha1")

	dst := dstRaw.(*v1alpha4.TrafficTarget)
	dst.ObjectMeta = src.ObjectMeta

	dst.TypeMeta = src.TypeMeta
	dst.APIVersion = v1alpha4.GroupVersion.Identifier()

	dst.Spec.Destination = v1alpha4.IdentityBindingSubject{
		Kind:      src.Destination.Kind,
		Name:      src.Destination.Name,
		Namespace: src.Destination.Namespace,
	}

	dst.Spec.Sources = []v1alpha4.IdentityBindingSubject{}
	for _, ibs := range src.Sources {
		s := v1alpha4.IdentityBindingSubject{
			Kind:      ibs.Kind,
			Name:      ibs.Name,
			Namespace: ibs.Namespace,
		}

		dst.Spec.Sources = append(dst.Spec.Sources, s)
	}

	dst.Spec.Rules = []v1alpha4.TrafficTargetRule{}
	for _, ibs := range src.Specs {
		s := v1alpha4.TrafficTargetRule{
			Kind:    ibs.Kind,
			Name:    ibs.Name,
			Matches: ibs.Matches,
		}

		dst.Spec.Rules = append(dst.Spec.Rules, s)
	}

	return nil
}

/*
ConvertFrom is expected to modify its receiver to contain the converted object.
Most of the conversion is straightforward copying, except for converting our changed field.
*/

// ConvertFrom converts from the Hub version (v1alpha4) to this version.
func (dst *TrafficTarget) ConvertFrom(srcRaw conversion.Hub) error {
	traffictargetlog.Info("ConvertFrom v1alpha4 to v1alpha1")

	src := srcRaw.(*v1alpha4.TrafficTarget)
	dst.ObjectMeta = src.ObjectMeta

	dst.TypeMeta = src.TypeMeta
	dst.APIVersion = GroupVersion.Identifier()

	dst.Destination = IdentityBindingSubject{
		Kind:      src.Spec.Destination.Kind,
		Name:      src.Spec.Destination.Name,
		Namespace: src.Spec.Destination.Namespace,
	}

	dst.Sources = []IdentityBindingSubject{}
	for _, ibs := range src.Spec.Sources {
		s := IdentityBindingSubject{
			Kind:      ibs.Kind,
			Name:      ibs.Name,
			Namespace: ibs.Namespace,
		}

		dst.Sources = append(dst.Sources, s)
	}

	dst.Specs = []TrafficTargetSpec{}
	for _, ibs := range src.Spec.Rules {
		s := TrafficTargetSpec{
			Kind:    ibs.Kind,
			Name:    ibs.Name,
			Matches: ibs.Matches,
		}

		dst.Specs = append(dst.Specs, s)
	}

	return nil
}
