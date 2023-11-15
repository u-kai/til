package pkg

type PublicSubnetId string
type VpcId string
type ResourceType string

const (
	EKSCluster   ResourceType = "EKSCluster"
	EKSNodeGroup ResourceType = "EKSNodeGroup"
)

type PublicResourceReference[Id ResourceIdentifier] struct {
	ResourceType
	VpcId
	PublicSubnetIds []PublicSubnetId
	Id              Id
}
type ResourceIdentifier interface {
	ToKey() string
}

type DeletePublicResourceAction[Id ResourceIdentifier] func([]PublicResourceReference[Id]) error
type SearchPublicSubnets[Id ResourceIdentifier] func([]PublicSubnetId) ([]PublicResourceReference[Id], error)

func checkDelete[Id ResourceIdentifier](f DeletePublicResourceAction[Id]) error {
	return f([]PublicResourceReference[Id]{})
}
func checkSearch[Id ResourceIdentifier](f SearchPublicSubnets[Id]) ([]PublicResourceReference[Id], error) {
	return f([]PublicSubnetId{})
}

func FilterPublicSubnets(publicSubnets []PublicSubnetId, subnets []string) []PublicSubnetId {
	var result []PublicSubnetId
	for _, subnet := range subnets {
		for _, publicSubnet := range publicSubnets {
			if subnet == string(publicSubnet) {
				result = append(result, publicSubnet)
			}
		}
	}
	return result
}
