package pkg

import (
	"context"
	"strings"

	"github.com/aws/aws-sdk-go-v2/service/eks"
)

type EKSClusterName string

func (e EKSClusterName) ToKey() string {
	return string(e)
}
func DeletePublicEKSClusterActionByClient(client *eks.Client) DeletePublicResourceAction[EKSClusterName] {
	return func(resources []PublicResourceReference[EKSClusterName]) error {
		for _, resource := range resources {
			name := resource.Id.ToKey()
			_, err := client.DeleteCluster(context.Background(), &eks.DeleteClusterInput{
				Name: &name,
			})
			if err != nil {
				if strings.Contains(err.Error(), "NotFound") {
					continue
				}
				return err
			}
		}
		return nil
	}
}

func SearchPublicSubnetsByClient(client *eks.Client) SearchPublicSubnets[EKSClusterName] {
	return func(subnets []PublicSubnetId) ([]PublicResourceReference[EKSClusterName], error) {
		var resources []PublicResourceReference[EKSClusterName]
		output, err := client.ListClusters(context.Background(), &eks.ListClustersInput{})
		if err != nil {
			return nil, err
		}
		for _, cluster := range output.Clusters {
			clusterInfo, err := client.DescribeCluster(context.Background(), &eks.DescribeClusterInput{
				Name: &cluster,
			})
			if err != nil {
				return nil, err
			}
			publicSubnets := FilterPublicSubnets(subnets, clusterInfo.Cluster.ResourcesVpcConfig.SubnetIds)
			if len(publicSubnets) == 0 {
				continue
			}
			resources = append(resources, PublicResourceReference[EKSClusterName]{
				ResourceType:    EKSCluster,
				VpcId:           VpcId(*clusterInfo.Cluster.ResourcesVpcConfig.VpcId),
				PublicSubnetIds: publicSubnets,
				Id:              EKSClusterName(cluster),
			})
		}
		return resources, nil
	}
}

