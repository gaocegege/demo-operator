package controllers

import (
	"context"
	"fmt"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	batchv1 "k8s.io/api/batch/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"

	demov1alpha1 "github.com/gaocegege/demo-operator/api/v1alpha1"
)

var _ = Describe("job Reconciler", func() {
	const timeout = time.Second * 5
	const interval = time.Second * 1

	demoJobName := "test"
	key := types.NamespacedName{
		Name:      demoJobName,
		Namespace: "default",
	}
	demoJob := &demov1alpha1.DemoJob{}

	BeforeEach(func() {
	})

	AfterEach(func() {
		// Add any teardown steps that needs to be executed after each test
	})

	Context("With the correct DemoJob", func() {
		It("Should create the job correctly", func() {
			toCreated := &demov1alpha1.DemoJob{
				ObjectMeta: metav1.ObjectMeta{
					Name:      demoJobName,
					Namespace: key.Namespace,
				},
				Spec: demov1alpha1.DemoJobSpec{
					Image: "test",
				},
			}
			Expect(k8sClient.Create(context.Background(), toCreated)).Should(Succeed())

			By("Expecting to get the new job from k8s client")
			job := &batchv1.Job{}
			Expect(k8sClient.Get(context.Background(), types.NamespacedName{
				Namespace: key.Namespace,
				Name:      demoJobName,
			}, job)).Should(Succeed())
			fmt.Printf("%v", job)

			Expect(len(job.OwnerReferences)).Should(Equal(1))

			By("Expecting to delete successfully")
			Eventually(func() error {
				m := &demov1alpha1.DemoJob{}
				k8sClient.Get(context.Background(), key, m)
				return k8sClient.Delete(context.Background(), demoJob)
			}, timeout, interval).Should(Succeed())

			By("Expecting not to get the new job from k8s client")
			Expect(k8sClient.Get(context.Background(), types.NamespacedName{
				Namespace: key.Namespace,
				Name:      demoJob.Name,
			}, job)).ShouldNot(BeNil())
			fmt.Printf("%v", job)
		})
	})
})
