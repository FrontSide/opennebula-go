package host
//
// import (
//   "fmt"
// 	"github.com/megamsys/opennebula-go/api"
// 	"gopkg.in/check.v1"
// 	"testing"
// )
//
// func Test(t *testing.T) {
// 	check.TestingT(t)
// }
//
// type S struct {
// 	cm map[string]string
// }
//
// var _ = check.Suite(&S{})
//
// func (s *S) SetUpSuite(c *check.C) {
// 	cm := make(map[string]string)
// 	cm[api.ENDPOINT] = "http://192.168.0.117:2633/RPC2"
// 	cm[api.USERID] = "oneadmin"
// 	cm[api.PASSWORD] = "dyovAupAuck9"
// 	s.cm = cm
// }
//
//
// func (s *S) TestGetVMs(c *check.C) {
// 	client, _ := api.NewClient(s.cm)
// 	vm := HQuery{T: client}
//   fmt.Printf("%#v",vm)
// 	_, err := vm.GetVMs(2)
// 	c.Assert(err, check.IsNil)
// }
//
// // func (s *S) TestAllocateHost(c *check.C) {
// //   client, _ := api.NewClient(s.cm)
// //   host := HQuery{T: client}
// //   hostname := "192.168.1.103"
// //   _, err := host.AllocateHost(hostname,"kvm","kvm", -1)
// //   c.Assert(err, check.NotNil)
// //
// // }
// /*
// func (s *S) TestDelHost(c *check.C) {
//   client, _ := api.NewClient(s.cm)
//   host := HQuery{T: client}
//   _, err := host.DelHost(4)
//   c.Assert(err, check.NotNil)
// }
// */
