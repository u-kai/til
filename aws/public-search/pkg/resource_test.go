package pkg_test

import (
	"public-search/pkg"
	"reflect"
	"testing"
)

func TestFilterPublicSubnets(t *testing.T) {
	for _, tt := range []struct {
		name            string
		publicSubnetIds []pkg.PublicSubnetId
		subnets         []string
		toBeSubnets     []pkg.PublicSubnetId
	}{
		{
			name:            "all public subnets",
			publicSubnetIds: []pkg.PublicSubnetId{"subnet-1", "subnet-2"},
			subnets:         []string{"subnet-1", "subnet-2"},
			toBeSubnets:     []pkg.PublicSubnetId{"subnet-1", "subnet-2"},
		},
		{
			name:            "no public subnets",
			publicSubnetIds: []pkg.PublicSubnetId{},
			subnets:         []string{"subnet-1", "subnet-2"},
			toBeSubnets:     nil,
		},
		{
			name:            "not contain",
			publicSubnetIds: []pkg.PublicSubnetId{"subnet-1", "subnet-2"},
			subnets:         []string{"subnet-3", "subnet-4"},
			toBeSubnets:     nil,
		},
		{
			name:            "contain one",
			publicSubnetIds: []pkg.PublicSubnetId{"subnet-1", "subnet-2"},
			subnets:         []string{"subnet-1", "subnet-3"},
			toBeSubnets:     []pkg.PublicSubnetId{"subnet-1"},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			got := pkg.FilterPublicSubnets(tt.publicSubnetIds, tt.subnets)
			if !reflect.DeepEqual(got, tt.toBeSubnets) {
				t.Errorf("FilterPublicSubnets() = %v, want %v", got, tt.toBeSubnets)
			}
		})
	}
}
