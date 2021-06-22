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
	"github.com/servicemeshinterface/smi-controller-sdk/sdk"
	"github.com/servicemeshinterface/smi-controller-sdk/sdk/controller"
	"github.com/stretchr/testify/mock"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/kubectl/pkg/scheme"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	accessv1alpha1 "github.com/servicemeshinterface/smi-controller-sdk/apis/access/v1alpha1"
	accessv1alpha2 "github.com/servicemeshinterface/smi-controller-sdk/apis/access/v1alpha2"

	splitv1alpha1 "github.com/servicemeshinterface/smi-controller-sdk/apis/split/v1alpha1"
	splitv1alpha2 "github.com/servicemeshinterface/smi-controller-sdk/apis/split/v1alpha2"
	splitv1alpha3 "github.com/servicemeshinterface/smi-controller-sdk/apis/split/v1alpha3"
	splitv1alpha4 "github.com/servicemeshinterface/smi-controller-sdk/apis/split/v1alpha4"
)

var opts = &godog.Options{
	Format: "pretty",
	Output: colors.Colored(os.Stdout),
}

var mockAPI *MockAPI
var logger logr.Logger
var k8sClient client.Client

// store a reference to any objects submitted to the controller for later cleanup
var trafficSplits []*splitv1alpha4.TrafficSplit

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

func setupClient() error {
	err := accessv1alpha1.AddToScheme(scheme.Scheme)
	if err != nil {
		return err
	}

	err = accessv1alpha2.AddToScheme(scheme.Scheme)
	if err != nil {
		return err
	}

	err = splitv1alpha1.AddToScheme(scheme.Scheme)
	if err != nil {
		return err
	}

	err = splitv1alpha2.AddToScheme(scheme.Scheme)
	if err != nil {
		return err
	}

	err = splitv1alpha3.AddToScheme(scheme.Scheme)
	if err != nil {
		return err
	}

	err = splitv1alpha4.AddToScheme(scheme.Scheme)
	if err != nil {
		return err
	}

	c := getK8sConfig()
	k8sClient, err = client.New(c, client.Options{Scheme: scheme.Scheme})
	if err != nil {
		return err
	}

	return nil
}

func initializeSuite(ctx *godog.ScenarioContext) {
	trafficSplits = []*splitv1alpha4.TrafficSplit{}
	logger = Log()

	err := setupClient()
	if err != nil {
		panic(err)
	}

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
	ctx := context.Background()
	k8sClient.DeleteAllOf(ctx, &splitv1alpha4.TrafficSplit{}, &client.DeleteAllOfOptions{})
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
	ts := &splitv1alpha4.TrafficSplit{
		ObjectMeta: v1.ObjectMeta{Name: "testing"},
		Spec: splitv1alpha4.TrafficSplitSpec{
			Service: "myService",
			Backends: []splitv1alpha4.TrafficSplitBackend{
				splitv1alpha4.TrafficSplitBackend{
					Service: "v1",
					Weight:  100,
				},
			},
		},
	}

	ctx := context.Background()
	err := k8sClient.Create(ctx, ts, &client.CreateOptions{})

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
