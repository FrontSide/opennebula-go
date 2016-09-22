package vnet

import (
  "github.com/megamsys/opennebula-go/api"
  "encoding/xml"
)

type VNETemplate struct {
  Template  Vnet `xml:"TEMPLATE"`
  T   *api.Rpc
}

type Vnet struct {
  Id             int    `json:"id" xml:"ID"`
  Name           string `json:"name" xml:"NAME"`
  Type           string  `json:"type" xml:"TYPE"`
  Description    string `json:"description" xml:"DESCRIPTION"`
  Bridge         string `json:"bridge" xml:"BRIDGE"`
  Network_addr   string `json:"network_addr" xml:"NETWORK_ADDRESS"`
  Network_mask   string `json:"network_mask" xml:"NETWORK_MASK"`
  Dns            string `json:"dns" xml:"DNS"`
  Gateway        string `json:"gateway" xml:"GATEWAY"`
  Vn_mad         string  `json:"vn_mad" xml:"VN_MAD"`
  Addrs          []*Address  `json:"addrs" xml:"AR"`
}

type Address struct {
  Type        string `json:"type" xml:"TYPE"`
  StartIP     string `json:"ip" xml:"IP"`
  Size        string `json:"size" xml:"SIZE"`
}

func (v *VNETemplate) CreateVnet(cluster_id int) (interface{} ,error)  {

  finalXML := VNETemplate{}
	finalXML.Template = v.Template
	finalData, _ := xml.Marshal(finalXML.Template)
	data := string(finalData)
	args := []interface{}{v.T.Key,data, cluster_id}
	res, err := v.T.Call(api.VNET_CREATE, args)
	if err != nil {
		return nil,err
	}
  return res, nil
}

func (v *VNETemplate) VnetAddIps() (interface{} ,error)  {
  finalXML := VNETemplate{}
	finalXML.Template.Addrs = v.Template.Addrs
	finalData, _ := xml.Marshal(finalXML.Template)
	data := string(finalData)
	args := []interface{}{v.T.Key,data, v.Template.Id}
	res, err := v.T.Call(api.VNET_ADDIP, args)
	if err != nil {
		return nil,err
	}
  return res, nil
}

func (v *VNETemplate) VnetInfos(vnet_id int) (interface{}, error) {
	args := []interface{}{v.T.Key, vnet_id}
	res, err := v.T.Call(api.VNET_SHOW, args)
	if err != nil {
		return nil,err
	}
  return res, nil
}


func (v *VNETemplate) VnetsInfos(filter_id int) (interface{}, error) {
  start_id := -1  //-1 for smaller values this is the offset used for pagination.
  end_id := -1  //-1 for get until the last ID
	args := []interface{}{v.T.Key, filter_id, start_id, end_id}
	res, err := v.T.Call(api.VNET_LIST, args)
	if err != nil {
		return nil,err
	}
  return res, nil
}
