package compute

import (
	"testing"
        
	"github.com/megamsys/opennebula-go/api"

	"gopkg.in/check.v1"
)

func Test(t *testing.T) {
	check.TestingT(t)
}

type S struct{}

var _ = check.Suite(&S{})

func (s *S) TestCreate(c *check.C) {
	//oneadmin:yourWuOtHij3
	client, _ := api.NewRPCClient("http://103.56.92.4:2633/RPC2", "oneadmin", "yourWuOtHij3")
     
	vmObj := VirtualMachine{Name: "test", TemplateName: "megam", Cpu: "1", Memory: "1024", Image: "megam", Client: client, ContextMap: map[string]string{"assembly_id": "ASM-007", "assemblies_id": "AMS-007"}} //memory in terms of MB! duh!
	_, error := vmObj.Create()
	c.Assert(error, check.IsNil)
}

func (s *S) TestDelete(c *check.C) {
	client, _ := api.NewRPCClient("http://103.56.92.4:2633/RPC2", "oneadmin", "yourWuOtHij3")
	vmObj := VirtualMachine{Name: "test", Client: client}

	_, error := vmObj.Delete()
	c.Assert(error, check.IsNil)
}

func (s *S) TestResume(c *check.C) {
	client, _ := api.NewRPCClient("http://103.56.92.4:2633/RPC2", "oneadmin", "yourWuOtHij3")
	vmObj := VirtualMachine{Name: "test", Client: client}

	_, error := vmObj.Resume()
	c.Assert(error, check.IsNil)
}

func (s *S) TestReboot(c *check.C) {
	client, _ := api.NewRPCClient("http://103.56.92.4:2633/RPC2", "oneadmin", "yourWuOtHij3")
	vmObj := VirtualMachine{Name: "test", Client: client}

	_, error := vmObj.Reboot()
	c.Assert(error, check.IsNil)
}

func (s *S) TestPoweroff(c *check.C) {
	client, _ := api.NewRPCClient("http://103.56.92.4:2633/RPC2", "oneadmin", "yourWuOtHij3")
	vmObj := VirtualMachine{Name: "test", Client: client}

	_, error := vmObj.Poweroff()
	c.Assert(error, check.IsNil)
}


