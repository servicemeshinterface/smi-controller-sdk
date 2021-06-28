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
	"github.com/servicemeshinterface/smi-controller-sdk/controllers/helpers"
	"github.com/servicemeshinterface/smi-controller-sdk/sdk"
	"github.com/servicemeshinterface/smi-controller-sdk/sdk/controller"
	"github.com/stretchr/testify/mock"
	"k8s.io/apimachinery/pkg/api/resource"
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

var mockAPI *helpers.MockAPI
var logger logr.Logger
var k8sClient client.Client
var config controller.Config

// store a reference to any objects submitted to the controller for later cleanup
var alpha4TrafficSplits []*splitv1alpha4.TrafficSplit
var alpha3TrafficSplits []*splitv1alpha3.TrafficSplit
var alpha2TrafficSplits []*splitv1alpha2.TrafficSplit
var alpha1TrafficSplits []*splitv1alpha1.TrafficSplit

func main() {
	godog.BindFlags("godog.", flag.CommandLine, opts)
	flag.Parse()

	status := godog.TestSuite{
		Name:                 "SDK Functional Tests",
		ScenarioInitializer:  initializeScenario,
		TestSuiteInitializer: initializeSuite,
		Options:              opts,
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

func initializeScenario(ctx *godog.ScenarioContext) {
	alpha4TrafficSplits = []*splitv1alpha4.TrafficSplit{}
	alpha3TrafficSplits = []*splitv1alpha3.TrafficSplit{}
	alpha2TrafficSplits = []*splitv1alpha2.TrafficSplit{}
	alpha1TrafficSplits = []*splitv1alpha1.TrafficSplit{}

	ctx.Step(`^the server is running$`, theServerIsRunning)
	ctx.Step(`^I create an "(.*)" TrafficSplitter$`, iCreateATrafficSplitter)
	ctx.Step(`^I expect the controller to have received the details$`, iExpectTheControllerToHaveRecievedTheDetails)

	ctx.AfterScenario(func(s *messages.Pickle, err error) {
		cleanupTrafficSplit()

		if err != nil {
			fmt.Println("Error occurred running the tests", err)
			fmt.Println(logger.(StringLogger).String())
		}

		// wait for server to have cleaned up objects and exit
		// as deleting an object is not immediate.
		// We should probably handle this eventual consistency in the cleanup
		// function however this 5 delay should be fine.
		// If you are raising a PR to fix this "should be fine" be sure to shame me in the comments
		time.Sleep(5 * time.Second)
	})
}

func initializeSuite(ctx *godog.TestSuiteContext) {
	logger = Log()

	err := setupClient()
	if err != nil {
		panic(err)
	}

	// create and start the controller
	setupMockAPI()

	sdk.API().RegisterV1Alpha(mockAPI)

	config = controller.DefaultConfig()
	config.WebhooksEnabled = true
	config.Logger = logger

	go controller.Start(config)
}

func cleanupTrafficSplit() {
	ctx := context.Background()
	for _, ts := range alpha4TrafficSplits {
		err := k8sClient.Delete(ctx, ts, &client.DeleteOptions{})
		if err != nil {
			fmt.Println("Error removing TrafficSplit object", err)
		}
	}

	for _, ts := range alpha3TrafficSplits {
		err := k8sClient.Delete(ctx, ts, &client.DeleteOptions{})
		if err != nil {
			fmt.Println("Error removing TrafficSplit object", err)
		}
	}

	for _, ts := range alpha2TrafficSplits {
		err := k8sClient.Delete(ctx, ts, &client.DeleteOptions{})
		if err != nil {
			fmt.Println("Error removing TrafficSplit object", err)
		}
	}

	for _, ts := range alpha1TrafficSplits {
		err := k8sClient.Delete(ctx, ts, &client.DeleteOptions{})
		if err != nil {
			fmt.Println("Error removing TrafficSplit object", err)
		}
	}
}

func setupMockAPI() {
	mockAPI = &helpers.MockAPI{}
	mockAPI.On("UpsertTrafficSplit", mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything).Return(ctrl.Result{}, nil)

	mockAPI.On("DeleteTrafficSplit", mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything).Return(ctrl.Result{}, nil)
}

func theServerIsRunning() error {
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

func iCreateATrafficSplitter(version string) error {
	switch version {
	case "alphav1":
		return iCreateAnAlpha1TrafficSplitter()
	case "alphav2":
		return iCreateAnAlpha2TrafficSplitter()
	case "alphav3":
		return iCreateAnAlpha3TrafficSplitter()
	case "alphav4":
		return iCreateAnAlpha4TrafficSplitter()
	default:
		return fmt.Errorf("Version %s not supported", version)
	}

	return nil
}

func iCreateAnAlpha4TrafficSplitter() error {
	ts := &splitv1alpha4.TrafficSplit{
		ObjectMeta: v1.ObjectMeta{Name: "splitalphav4", Namespace: "default"},
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
	if err == nil {
		alpha4TrafficSplits = append(alpha4TrafficSplits, ts)
	}

	return err
}

func iCreateAnAlpha3TrafficSplitter() error {
	ts := &splitv1alpha3.TrafficSplit{
		ObjectMeta: v1.ObjectMeta{Name: "splitalphav3", Namespace: "default"},
		Spec: splitv1alpha3.TrafficSplitSpec{
			Service: "myService",
			Backends: []splitv1alpha3.TrafficSplitBackend{
				splitv1alpha3.TrafficSplitBackend{
					Service: "v1",
					Weight:  100,
				},
			},
		},
	}

	ctx := context.Background()
	err := k8sClient.Create(ctx, ts, &client.CreateOptions{})
	if err == nil {
		alpha3TrafficSplits = append(alpha3TrafficSplits, ts)
	}

	return err
}

func iCreateAnAlpha2TrafficSplitter() error {
	ts := &splitv1alpha2.TrafficSplit{
		ObjectMeta: v1.ObjectMeta{Name: "splitalphav2", Namespace: "default"},
		Spec: splitv1alpha2.TrafficSplitSpec{
			Service: "myService",
			Backends: []splitv1alpha2.TrafficSplitBackend{
				splitv1alpha2.TrafficSplitBackend{
					Service: "v1",
					Weight:  100,
				},
			},
		},
	}

	ctx := context.Background()
	err := k8sClient.Create(ctx, ts, &client.CreateOptions{})
	if err == nil {
		alpha2TrafficSplits = append(alpha2TrafficSplits, ts)
	}

	return err
}

func iCreateAnAlpha1TrafficSplitter() error {
	ts := &splitv1alpha1.TrafficSplit{
		ObjectMeta: v1.ObjectMeta{Name: "splitalphav1", Namespace: "default"},
		Spec: splitv1alpha1.TrafficSplitSpec{
			Service: "myService",
			Backends: []splitv1alpha1.TrafficSplitBackend{
				splitv1alpha1.TrafficSplitBackend{
					Service: "v1",
					Weight:  resource.NewQuantity(100, resource.DecimalSI),
				},
			},
		},
	}

	ctx := context.Background()
	err := k8sClient.Create(ctx, ts, &client.CreateOptions{})
	if err == nil {
		alpha1TrafficSplits = append(alpha1TrafficSplits, ts)
	}

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
	timeout := time.After(duration)

	var err error

	go func() {
		for {
			err = f()
			if err == nil {
				done <- struct{}{}
				break
			}

			// retry after 1s
			time.Sleep(1 * time.Second)
		}
	}()

	select {
	case <-timeout:
		return err
	case <-done:
		return nil
	}
}
