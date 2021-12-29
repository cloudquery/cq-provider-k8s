package rbac

import (
	"testing"

	"github.com/cloudquery/faker/v3"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func fakeManagedFields(t *testing.T) metav1.ManagedFieldsEntry {
	m := metav1.ManagedFieldsEntry{}
	if err := faker.FakeData(&m); err != nil {
		t.Fatal(err)
	}
	m.FieldsV1 = &metav1.FieldsV1{
		Raw: []byte("{\"test\":1}"),
	}
	return m
}
