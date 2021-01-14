# SMI Controller SDK

This project is a scaffold Kubernetes controller for the SMI specification.

Projects that would like to build a SMI Spec complient controller only need to 
define a plugin that implement the extension points defined by the SDK.  

Core Kubernetes controller methods that handle the lifecylcle are implemented by the SDK and 
the methods in your API are called accordingly.
