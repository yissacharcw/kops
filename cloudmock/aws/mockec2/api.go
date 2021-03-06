/*
Copyright 2016 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package mockec2

import (
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/ec2/ec2iface"
)

type MockEC2 struct {
	addressNumber int
	Addresses     []*ec2.Address

	RouteTables []*ec2.RouteTable

	Images []*ec2.Image

	subnetNumber int
	Subnets      []*ec2.Subnet

	Tags []*ec2.TagDescription

	vpcNumber int
	Vpcs      map[string]*vpcInfo
}

var _ ec2iface.EC2API = &MockEC2{}
