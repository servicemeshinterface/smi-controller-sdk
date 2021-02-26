package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/cucumber/godog"
	"github.com/cucumber/godog/colors"
	"github.com/cucumber/messages-go/v10"
	"github.com/go-logr/logr"
	"github.com/nicholasjackson/smi-controller-sdk/sdk"
	"github.com/nicholasjackson/smi-controller-sdk/sdk/controller"
	"github.com/stretchr/testify/mock"
	"k8s.io/apimachinery/pkg/api/resource"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	ctrl "sigs.k8s.io/controller-runtime"

	splitv1alpha1 "github.com/servicemeshinterface/smi-sdk-go/pkg/apis/split/v1alpha1"
	splitClientSet "github.com/servicemeshinterface/smi-sdk-go/pkg/gen/client/split/clientset/versioned"
)

var opts = &godog.Options{
	Format: "pretty",
	Output: colors.Colored(os.Stdout),
}

var mockAPI *MockAPI
var logger logr.Logger

// store a reference to any objects submitted to the controller for later cleanup
var trafficSplits []*splitv1alpha1.TrafficSplit

func main() {
	godog.BindFlags("godog.", flag.CommandLine, opts)
	flag.Parse()

	status := godog.TestSuite{
		Name:                "SDK Functional Tests",
		ScenarioInitializer: initializeSuite,
		Options:             opts,
	}.Run()

	os.Exit(status)
}

func initializeSuite(ctx *godog.ScenarioContext) {
	trafficSplits = []*splitv1alpha1.TrafficSplit{}
	logger = Log()

	ctx.Step(`^the server is running$`, theServerIsRunning)
	ctx.Step(`^I create a TrafficSplitter$`, iCreateATrafficSplitter)
	ctx.Step(`^I expect the controller to have received the details$`, iExpectTheControllerToHaveRecievedTheDetails)

	ctx.AfterScenario(func(s *messages.Pickle, err error) {
		cleanupTrafficSplit()

		if err != nil {
			fmt.Println(logger.(*StringLogger).String())
		}
	})
}

func cleanupTrafficSplit() {
	c := getK8sConfig()
	sl, err := splitClientSet.NewForConfig(c)
	if err != nil {
		panic(err.Error())
	}

	for _, ts := range trafficSplits {
		sl.SplitV1alpha1().TrafficSplits("default").Delete(context.Background(), ts.Name, v1.DeleteOptions{})
	}
}

func theServerIsRunning() error {
	mockAPI = &MockAPI{}
	mockAPI.On("UpsertTrafficSplit", mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything).Return(ctrl.Result{}, nil)

	sdk.API().RegisterV1Alpha(mockAPI)

	// create and start the controller
	config := controller.DefaultConfig()
	config.WebhooksEnabled = false
	config.Logger = logger

	go controller.Start(config)

	return waitForComplete(
		30*time.Second,
		func() error {
			resp, err := http.Get(fmt.Sprintf("http://%s/readyz", config.HealthProbeBindAddress))
			if err == nil {
				if resp != nil && resp.StatusCode == http.StatusOK {
					return nil
				}
			}

			return fmt.Errorf("Timeout waiting for service to become ready")
		},
	)
}

func iCreateATrafficSplitter() error {
	c := getK8sConfig()
	sl, err := splitClientSet.NewForConfig(c)
	if err != nil {
		return err
	}

	ts := &splitv1alpha1.TrafficSplit{
		ObjectMeta: v1.ObjectMeta{Name: "testing"},
		Spec: splitv1alpha1.TrafficSplitSpec{
			Service: "myService",
			Backends: []splitv1alpha1.TrafficSplitBackend{
				splitv1alpha1.TrafficSplitBackend{
					Service: "v1",
					Weight:  resource.NewQuantity(100, resource.BinarySI),
				},
			},
		},
	}

	// add to our collection so we can cleanup later
	trafficSplits = append(trafficSplits, ts)

	ts, err = sl.SplitV1alpha1().TrafficSplits("default").Create(context.Background(), ts, v1.CreateOptions{})

	return err
}

// The controller is eventually consistent so we need to check this in a loop
func iExpectTheControllerToHaveRecievedTheDetails() error {
	return waitForComplete(
		30*time.Second,
		func() error {
			if len(mockAPI.Calls) < 1 {
				return fmt.Errorf("Expected UpsertTrafficSplit to have been called")
			}

			return nil
		},
	)
}

func getK8sConfig() *rest.Config {
	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", os.Getenv("KUBECONFIG"))
	if err != nil {
		panic(err.Error())
	}

	return config
}

// helper function to loop until a condition is met
func waitForComplete(duration time.Duration, f func() error) error {
	// wait for the server to mark it is ready
	done := make(chan struct{})
	timeout := time.After(30 * time.Second)

	var err error

	go func() {
		for {
			err = f()
			if err == nil {
				done <- struct{}{}
			}

			time.Sleep(2 * time.Second)
		}
	}()

	select {
	case <-timeout:
		return err
	case <-done:
		return nil
	}
}
