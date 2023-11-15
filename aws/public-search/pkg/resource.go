package pkg

type PublicSubnetId string
type VpcId string
type ResourceType string

const (
	EKSCluster ResourceType = "EKSCluster"
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

type DeletePublicResourceAction[Id ResourceIdentifier] func([]PublicSubnetId) error
type SearchPublicSubnets[Id ResourceIdentifier] func([]PublicSubnetId) ([]PublicResourceReference[Id], error)
