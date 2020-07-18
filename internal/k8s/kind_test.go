package k8s

import (
	projectcontour "github.com/projectcontour/contour/apis/projectcontour/v1"
	v1 "k8s.io/api/core/v1"
	"k8s.io/api/networking/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"testing"
)

func TestKindOf(t *testing.T) {
	cases := []struct {
		Kind string
		Obj  interface{}
	}{
		{"Secret", &v1.Secret{}},
		{"Service", &v1.Service{}},
		{"Endpoints", &v1.Endpoints{}},
		{"", &v1.Pod{}},
		{"Ingress", &v1beta1.Ingress{}},
		{"HTTPProxy", &projectcontour.HTTPProxy{}},
		{"TLSCertificateDelegation", &projectcontour.TLSCertificateDelegation{}},
		{"Foo", &unstructured.Unstructured{
			Object: map[string]interface{}{
				"apiVersion": "test.projectcontour.io/v1",
				"kind":       "Foo",
			}},
		},
	}

	for _, c := range cases {
		kindOf := KindOf(c.Obj)
		if kindOf != c.Kind {
			t.Errorf("got %q for KindOf(%T), wanted %q",
				kindOf, c.Obj, c.Kind)
		}
	}
}

func TestApiVersionOf(t *testing.T) {
	cases := []struct {
		apiVersion string
		Obj        interface{}
	}{
		{"apps/v1", &v1.Service{
			TypeMeta: metav1.TypeMeta{
				APIVersion: "apps/v1",
			},
		}},
		{"apps/v1", &v1.Endpoints{
			TypeMeta: metav1.TypeMeta{
				APIVersion: "apps/v1",
			},
		}},
		{"apps/v1", &v1.Pod{
			TypeMeta: metav1.TypeMeta{
				APIVersion: "apps/v1",
			}}},
		{"apps/v1", &v1beta1.Ingress{
			TypeMeta: metav1.TypeMeta{
				APIVersion: "apps/v1",
			}}},
		{"apps/v1", &projectcontour.HTTPProxy{
			TypeMeta: metav1.TypeMeta{
				APIVersion: "apps/v1",
			}}},
		{"apps/v1", &projectcontour.TLSCertificateDelegation{
			TypeMeta: metav1.TypeMeta{
				APIVersion: "apps/v1",
			}}},
		{"test.projectcontour.io/v1", &unstructured.Unstructured{
			Object: map[string]interface{}{
				"apiVersion": "test.projectcontour.io/v1",
				"kind":       "Foo",
			}},
		},
	}

	for _, c := range cases {
		apiVersionOf := ApiVersionOf(c.Obj)
		if apiVersionOf != c.apiVersion {
			t.Errorf("got %q for KindOf(%T), wanted %q",
				apiVersionOf, c.Obj, c.apiVersion)
		}
	}
}
