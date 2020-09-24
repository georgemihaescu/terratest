package azure

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2019-09-01/network"
)

// LoadBalancerExistsE returns true if the load balancer exists, else returns false with err
func LoadBalancerExistsE(loadBalancerName string, resourceGroupName string, subscriptionID string) (bool, error) {
	subscriptionID, err := getTargetAzureSubscription(subscriptionID)
	if err != nil {
		return false, err
	}
	client, err := GetLoadBalancerClientE(subscriptionID)
	if err != nil {
		return false, err
	}
	lb, err := client.Get(context.Background(), resourceGroupName, loadBalancerName, "")
	if err != nil {
		return false, err
	}

	return *lb.Name == loadBalancerName, nil
}

// GetLoadBalancerE returns a load balancer resource as specified by name, else returns nil with err
func GetLoadBalancerE(loadBalancerName string, resourceGroupName string, subscriptionID string) (*network.LoadBalancer, error) {
	subscriptionID, err := getTargetAzureSubscription(subscriptionID)
	if err != nil {
		return nil, err
	}
	client, err := GetLoadBalancerClientE(subscriptionID)
	if err != nil {
		return nil, err
	}
	lb, err := client.Get(context.Background(), resourceGroupName, loadBalancerName, "")
	if err != nil {
		return nil, err
	}

	return &lb, nil
}

// GetLoadBalancerClientE creates a load balancer client.
func GetLoadBalancerClientE(subscriptionID string) (*network.LoadBalancersClient, error) {
	loadBalancerClient := network.NewLoadBalancersClient(subscriptionID)
	authorizer, err := NewAuthorizer()
	if err != nil {
		return nil, err
	}
	loadBalancerClient.Authorizer = *authorizer
	return &loadBalancerClient, nil
}

// GetLoadBalancerFrontendConfig returns an IP address and specifies public or private
func GetLoadBalancerFrontendConfig(pipResource string, resourceGroupName string, subscriptionID string) (ipAddress string, publicOrPrivate string, err1 error) {
	// TODO: pipResource is non-nil for public, nil for private
	// TODO: refactor to check private IP first, to determine how to get IP value

	subscriptionID, err := getTargetAzureSubscription(subscriptionID)
	if err != nil {
		return "", "", err
	}
	client, err := GetPublicIPAddressClientE(subscriptionID)
	if err != nil {
		return "", "", err
	}
	publicIPAddress, err := client.Get(context.Background(), resourceGroupName, pipResource, "")
	if err != nil {
		return "", "", err
	}

	pipProps := *publicIPAddress.PublicIPAddressPropertiesFormat
	ipValue := (pipProps.IPAddress)

	// TODO: return public or private after determination
	return *ipValue, "public", nil
}

// GetPublicIPAddressE returns a Public IP Address resource, else returns nil with err
func GetPublicIPAddressE(publicIPAddressName string, resourceGroupName string, subscriptionID string) (*network.PublicIPAddress, error) {
	subscriptionID, err := getTargetAzureSubscription(subscriptionID)
	if err != nil {
		return nil, err
	}
	client, err := GetPublicIPAddressClientE(subscriptionID)
	if err != nil {
		return nil, err
	}
	publicIPAddress, err := client.Get(context.Background(), resourceGroupName, publicIPAddressName, "")
	if err != nil {
		return nil, err
	}
	return &publicIPAddress, nil
}

// GetPublicIPAddressClientE creates a PublicIPAddresses client
func GetPublicIPAddressClientE(subscriptionID string) (*network.PublicIPAddressesClient, error) {
	client := network.NewPublicIPAddressesClient(subscriptionID)
	authorizer, err := NewAuthorizer()
	if err != nil {
		return nil, err
	}
	client.Authorizer = *authorizer
	return &client, nil
}
