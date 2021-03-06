package network_test

import (
	"github.com/terraform-providers/terraform-provider-hcloud/internal/network"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hetznercloud/hcloud-go/hcloud"
	"github.com/terraform-providers/terraform-provider-hcloud/internal/testsupport"
	"github.com/terraform-providers/terraform-provider-hcloud/internal/testtemplate"
)

func TestNetworkSubnetResource_Basic(t *testing.T) {
	var nw hcloud.Network

	resNetwork := &network.RData{
		Name:    "network-test-subnet",
		IPRange: "10.0.0.0/16",
		Labels:  nil,
	}
	resNetwork.SetRName("network-subnet")
	res := &network.RDataSubnet{
		Type:        "cloud",
		NetworkID:   resNetwork.TFID() + ".id",
		NetworkZone: "eu-central",
		IPRange:     "10.0.0.0/24",
	}
	res.SetRName("network-subnet-test")
	tmplMan := testtemplate.Manager{}
	resource.Test(t, resource.TestCase{
		PreCheck:     testsupport.AccTestPreCheck(t),
		Providers:    testsupport.AccTestProviders(),
		CheckDestroy: testsupport.CheckResourcesDestroyed(network.ResourceType, network.ByID(t, &nw)),
		Steps: []resource.TestStep{
			{
				// Create a new Network using the required values
				// only.
				Config: tmplMan.Render(t,
					"testdata/r/hcloud_network", resNetwork,
					"testdata/r/hcloud_network_subnet", res,
				),
				Check: resource.ComposeTestCheckFunc(
					testsupport.CheckResourceExists(resNetwork.TFID(), network.ByID(t, &nw)),
					resource.TestCheckResourceAttr(res.TFID(), "type", res.Type),
					resource.TestCheckResourceAttr(res.TFID(), "ip_range", res.IPRange),
					resource.TestCheckResourceAttr(res.TFID(), "network_zone", res.NetworkZone),
				),
			},
		},
	})
}
