package v1beta1

import (
	"context"
	"fmt"

	"k8s.io/apimachinery/pkg/runtime"
)

type ExternalSecretValidator struct{}

func (esv *ExternalSecretValidator) ValidateCreate(ctx context.Context, obj runtime.Object) error {
	es, ok := obj.(*ExternalSecret)
	if !ok {
		return fmt.Errorf("unexpected type")
	}

	if (es.Spec.Target.DeletionPolicy == DeletionPolicyDelete && es.Spec.Target.CreationPolicy == CreatePolicyMerge) ||
		(es.Spec.Target.DeletionPolicy == DeletionPolicyDelete && es.Spec.Target.CreationPolicy == CreatePolicyNone) {
		return fmt.Errorf("deletionPolicy=Delete must not be used when the controller doesn't own the secret. Please set creationPolcy=Owner")
	}

	if es.Spec.Target.DeletionPolicy == DeletionPolicyMerge && es.Spec.Target.CreationPolicy == CreatePolicyNone {
		return fmt.Errorf("deletionPolicy=Merge must not be used with creationPolcy=None. There is no Secret to merge with")
	}

	return nil
}

func (esv *ExternalSecretValidator) ValidateUpdate(ctx context.Context, oldObj, newObj runtime.Object) error {
	return nil
}

func (esv *ExternalSecretValidator) ValidateDelete(ctx context.Context, obj runtime.Object) error {
	return nil
}
