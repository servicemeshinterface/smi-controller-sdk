package sdk

// API defines an object containing functions that the
// specific service meshes must implement to use this controller
type api struct {
	v1alpha2 *v1Alpha2Impl
}

var apiInstance *api

// Returns a sinlgeton instance of the API
func API() *api {
	if apiInstance == nil {
		apiInstance = &api{
			v1alpha2: &v1Alpha2Impl{},
		}
	}

	return apiInstance
}

// register an interface which contains methods from the V1Alpha2 interface
// we register an interface to allow optional methods
func (a *api) RegisterV1Alpha2(i interface{}) {
	a.v1alpha2.RegisterV1Alpha2(i)
}

func (a *api) V1Alpha2() V1Alpha2 {
	return a.v1alpha2
}
