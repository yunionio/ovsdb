// DO NOT EDIT: automatically generated code

package vswitch

import (
	"fmt"

	"github.com/pkg/errors"

	"yunion.io/x/ovsdb/types"
)

type OvsdbOpenVSwitch struct {
	AutoAttach             AutoAttachTable
	Bridge                 BridgeTable
	Controller             ControllerTable
	FlowSampleCollectorSet FlowSampleCollectorSetTable
	FlowTable              FlowTableTable
	IPFIX                  IPFIXTable
	Interface              InterfaceTable
	Manager                ManagerTable
	Mirror                 MirrorTable
	NetFlow                NetFlowTable
	OpenVSwitch            OpenVSwitchTable
	Port                   PortTable
	QoS                    QoSTable
	Queue                  QueueTable
	SSL                    SSLTable
	SFlow                  SFlowTable
}

func (db OvsdbOpenVSwitch) FindOneMatchNonZeros(irow types.IRow) types.IRow {
	switch row := irow.(type) {
	case *AutoAttach:
		if r := db.AutoAttach.FindOneMatchNonZeros(row); r != nil {
			return r
		}
		return nil
	case *Bridge:
		if r := db.Bridge.FindOneMatchNonZeros(row); r != nil {
			return r
		}
		return nil
	case *Controller:
		if r := db.Controller.FindOneMatchNonZeros(row); r != nil {
			return r
		}
		return nil
	case *FlowSampleCollectorSet:
		if r := db.FlowSampleCollectorSet.FindOneMatchNonZeros(row); r != nil {
			return r
		}
		return nil
	case *FlowTable:
		if r := db.FlowTable.FindOneMatchNonZeros(row); r != nil {
			return r
		}
		return nil
	case *IPFIX:
		if r := db.IPFIX.FindOneMatchNonZeros(row); r != nil {
			return r
		}
		return nil
	case *Interface:
		if r := db.Interface.FindOneMatchNonZeros(row); r != nil {
			return r
		}
		return nil
	case *Manager:
		if r := db.Manager.FindOneMatchNonZeros(row); r != nil {
			return r
		}
		return nil
	case *Mirror:
		if r := db.Mirror.FindOneMatchNonZeros(row); r != nil {
			return r
		}
		return nil
	case *NetFlow:
		if r := db.NetFlow.FindOneMatchNonZeros(row); r != nil {
			return r
		}
		return nil
	case *OpenVSwitch:
		if r := db.OpenVSwitch.FindOneMatchNonZeros(row); r != nil {
			return r
		}
		return nil
	case *Port:
		if r := db.Port.FindOneMatchNonZeros(row); r != nil {
			return r
		}
		return nil
	case *QoS:
		if r := db.QoS.FindOneMatchNonZeros(row); r != nil {
			return r
		}
		return nil
	case *Queue:
		if r := db.Queue.FindOneMatchNonZeros(row); r != nil {
			return r
		}
		return nil
	case *SSL:
		if r := db.SSL.FindOneMatchNonZeros(row); r != nil {
			return r
		}
		return nil
	case *SFlow:
		if r := db.SFlow.FindOneMatchNonZeros(row); r != nil {
			return r
		}
		return nil
	}
	panic(types.ErrBadType)
}

func (db OvsdbOpenVSwitch) FindOneMatchByAnyIndex(irow types.IRow) types.IRow {
	switch row := irow.(type) {
	case *AutoAttach:
		if r := db.AutoAttach.OvsdbGetByAnyIndex(row); r != nil {
			return r
		}
		return nil
	case *Bridge:
		if r := db.Bridge.OvsdbGetByAnyIndex(row); r != nil {
			return r
		}
		return nil
	case *Controller:
		if r := db.Controller.OvsdbGetByAnyIndex(row); r != nil {
			return r
		}
		return nil
	case *FlowSampleCollectorSet:
		if r := db.FlowSampleCollectorSet.OvsdbGetByAnyIndex(row); r != nil {
			return r
		}
		return nil
	case *FlowTable:
		if r := db.FlowTable.OvsdbGetByAnyIndex(row); r != nil {
			return r
		}
		return nil
	case *IPFIX:
		if r := db.IPFIX.OvsdbGetByAnyIndex(row); r != nil {
			return r
		}
		return nil
	case *Interface:
		if r := db.Interface.OvsdbGetByAnyIndex(row); r != nil {
			return r
		}
		return nil
	case *Manager:
		if r := db.Manager.OvsdbGetByAnyIndex(row); r != nil {
			return r
		}
		return nil
	case *Mirror:
		if r := db.Mirror.OvsdbGetByAnyIndex(row); r != nil {
			return r
		}
		return nil
	case *NetFlow:
		if r := db.NetFlow.OvsdbGetByAnyIndex(row); r != nil {
			return r
		}
		return nil
	case *OpenVSwitch:
		if r := db.OpenVSwitch.OvsdbGetByAnyIndex(row); r != nil {
			return r
		}
		return nil
	case *Port:
		if r := db.Port.OvsdbGetByAnyIndex(row); r != nil {
			return r
		}
		return nil
	case *QoS:
		if r := db.QoS.OvsdbGetByAnyIndex(row); r != nil {
			return r
		}
		return nil
	case *Queue:
		if r := db.Queue.OvsdbGetByAnyIndex(row); r != nil {
			return r
		}
		return nil
	case *SSL:
		if r := db.SSL.OvsdbGetByAnyIndex(row); r != nil {
			return r
		}
		return nil
	case *SFlow:
		if r := db.SFlow.OvsdbGetByAnyIndex(row); r != nil {
			return r
		}
		return nil
	}
	panic(types.ErrBadType)
}

type AutoAttachTable []AutoAttach

var _ types.ITable = &AutoAttachTable{}

func (tbl AutoAttachTable) OvsdbTableName() string {
	return "AutoAttach"
}

func (tbl AutoAttachTable) OvsdbIsRoot() bool {
	return false
}

func (tbl AutoAttachTable) Rows() []types.IRow {
	r := make([]types.IRow, len(tbl))
	for i := range tbl {
		r[i] = &tbl[i]
	}
	return r
}

func (tbl AutoAttachTable) NewRow() types.IRow {
	return &AutoAttach{}
}

func (tbl *AutoAttachTable) AppendRow(irow types.IRow) {
	row := irow.(*AutoAttach)
	*tbl = append(*tbl, *row)
}

func (tbl AutoAttachTable) OvsdbHasIndex() bool {
	return false
}

func (tbl AutoAttachTable) OvsdbGetByAnyIndex(irow1 types.IRow) types.IRow {
	return nil
}

func (tbl AutoAttachTable) FindOneMatchNonZeros(row1 *AutoAttach) *AutoAttach {
	for i := range tbl {
		row := &tbl[i]
		if row.MatchNonZeros(row1) {
			return row
		}
	}
	return nil
}

type AutoAttach struct {
	Uuid              string          `json:"_uuid"`
	Version           string          `json:"_version"`
	Mappings          map[int64]int64 `json:"mappings"`
	SystemDescription string          `json:"system_description"`
	SystemName        string          `json:"system_name"`
}

var _ types.IRow = &AutoAttach{}

func (row *AutoAttach) OvsdbTableName() string {
	return "AutoAttach"
}

func (row *AutoAttach) OvsdbIsRoot() bool {
	return false
}

func (row *AutoAttach) OvsdbUuid() string {
	return row.Uuid
}

func (row *AutoAttach) OvsdbCmdArgs() []string {
	r := []string{}
	r = append(r, types.OvsdbCmdArgsMapIntegerInteger("mappings", row.Mappings)...)
	r = append(r, types.OvsdbCmdArgsString("system_description", row.SystemDescription)...)
	r = append(r, types.OvsdbCmdArgsString("system_name", row.SystemName)...)
	return r
}

func (row *AutoAttach) SetColumn(name string, val interface{}) (err error) {
	defer func() {
		if panicErr := recover(); panicErr != nil {
			err = errors.Wrapf(panicErr.(error), "%s: %#v", name, fmt.Sprintf("%#v", val))
		}
	}()
	switch name {
	case "_uuid":
		row.Uuid = types.EnsureUuid(val)
	case "_version":
		row.Version = types.EnsureUuid(val)
	case "mappings":
		row.Mappings = types.EnsureMapIntegerInteger(val)
	case "system_description":
		row.SystemDescription = types.EnsureString(val)
	case "system_name":
		row.SystemName = types.EnsureString(val)
	default:
		panic(types.ErrUnknownColumn)
	}
	return
}

func (row *AutoAttach) MatchNonZeros(row1 *AutoAttach) bool {
	if !types.MatchUuidIfNonZero(row.Uuid, row1.Uuid) {
		return false
	}
	if !types.MatchUuidIfNonZero(row.Version, row1.Version) {
		return false
	}
	if !types.MatchMapIntegerIntegerIfNonZero(row.Mappings, row1.Mappings) {
		return false
	}
	if !types.MatchStringIfNonZero(row.SystemDescription, row1.SystemDescription) {
		return false
	}
	if !types.MatchStringIfNonZero(row.SystemName, row1.SystemName) {
		return false
	}
	return true
}

func (row *AutoAttach) HasExternalIds() bool {
	return false
}

func (row *AutoAttach) SetExternalId(k, v string) {
	panic(errors.Wrap(types.ErrUnknownColumn, "external_ids"))
}

func (row *AutoAttach) GetExternalId(k string) (string, bool) {
	panic(errors.Wrap(types.ErrUnknownColumn, "external_ids"))
}

func (row *AutoAttach) RemoveExternalId(k string) (string, bool) {
	panic(errors.Wrap(types.ErrUnknownColumn, "external_ids"))
}

type BridgeTable []Bridge

var _ types.ITable = &BridgeTable{}

func (tbl BridgeTable) OvsdbTableName() string {
	return "Bridge"
}

func (tbl BridgeTable) OvsdbIsRoot() bool {
	return false
}

func (tbl BridgeTable) Rows() []types.IRow {
	r := make([]types.IRow, len(tbl))
	for i := range tbl {
		r[i] = &tbl[i]
	}
	return r
}

func (tbl BridgeTable) NewRow() types.IRow {
	return &Bridge{}
}

func (tbl *BridgeTable) AppendRow(irow types.IRow) {
	row := irow.(*Bridge)
	*tbl = append(*tbl, *row)
}

func (tbl BridgeTable) OvsdbHasIndex() bool {
	return true
}

func (row *Bridge) MatchByName(row1 *Bridge) bool {
	if !types.MatchString(row.Name, row1.Name) {
		return false
	}
	return true
}

func (tbl BridgeTable) GetByName(row1 *Bridge) *Bridge {
	for i := range tbl {
		row := &tbl[i]
		if row.MatchByName(row1) {
			return row
		}
	}
	return nil
}

func (tbl BridgeTable) OvsdbGetByAnyIndex(irow1 types.IRow) types.IRow {
	row1 := irow1.(*Bridge)
	if !(types.IsZeroString(row1.Name)) {
		if row := tbl.GetByName(row1); row != nil {
			return row
		}
	}
	return nil
}

func (tbl BridgeTable) FindOneMatchNonZeros(row1 *Bridge) *Bridge {
	for i := range tbl {
		row := &tbl[i]
		if row.MatchNonZeros(row1) {
			return row
		}
	}
	return nil
}

type Bridge struct {
	Uuid                string            `json:"_uuid"`
	Version             string            `json:"_version"`
	AutoAttach          *string           `json:"auto_attach"`
	Controller          []string          `json:"controller"`
	DatapathId          *string           `json:"datapath_id"`
	DatapathType        string            `json:"datapath_type"`
	DatapathVersion     string            `json:"datapath_version"`
	ExternalIds         map[string]string `json:"external_ids"`
	FailMode            *string           `json:"fail_mode"`
	FloodVlans          []int64           `json:"flood_vlans"`
	FlowTables          map[int64]string  `json:"flow_tables"`
	Ipfix               *string           `json:"ipfix"`
	McastSnoopingEnable bool              `json:"mcast_snooping_enable"`
	Mirrors             []string          `json:"mirrors"`
	Name                string            `json:"name"`
	Netflow             *string           `json:"netflow"`
	OtherConfig         map[string]string `json:"other_config"`
	Ports               []string          `json:"ports"`
	Protocols           []string          `json:"protocols"`
	RstpEnable          bool              `json:"rstp_enable"`
	RstpStatus          map[string]string `json:"rstp_status"`
	Sflow               *string           `json:"sflow"`
	Status              map[string]string `json:"status"`
	StpEnable           bool              `json:"stp_enable"`
}

var _ types.IRow = &Bridge{}

func (row *Bridge) OvsdbTableName() string {
	return "Bridge"
}

func (row *Bridge) OvsdbIsRoot() bool {
	return false
}

func (row *Bridge) OvsdbUuid() string {
	return row.Uuid
}

func (row *Bridge) OvsdbCmdArgs() []string {
	r := []string{}
	r = append(r, types.OvsdbCmdArgsUuidOptional("auto_attach", row.AutoAttach)...)
	r = append(r, types.OvsdbCmdArgsUuidMultiples("controller", row.Controller)...)
	r = append(r, types.OvsdbCmdArgsStringOptional("datapath_id", row.DatapathId)...)
	r = append(r, types.OvsdbCmdArgsString("datapath_type", row.DatapathType)...)
	r = append(r, types.OvsdbCmdArgsString("datapath_version", row.DatapathVersion)...)
	r = append(r, types.OvsdbCmdArgsMapStringString("external_ids", row.ExternalIds)...)
	r = append(r, types.OvsdbCmdArgsStringOptional("fail_mode", row.FailMode)...)
	r = append(r, types.OvsdbCmdArgsIntegerMultiples("flood_vlans", row.FloodVlans)...)
	r = append(r, types.OvsdbCmdArgsMapIntegerUuid("flow_tables", row.FlowTables)...)
	r = append(r, types.OvsdbCmdArgsUuidOptional("ipfix", row.Ipfix)...)
	r = append(r, types.OvsdbCmdArgsBoolean("mcast_snooping_enable", row.McastSnoopingEnable)...)
	r = append(r, types.OvsdbCmdArgsUuidMultiples("mirrors", row.Mirrors)...)
	r = append(r, types.OvsdbCmdArgsString("name", row.Name)...)
	r = append(r, types.OvsdbCmdArgsUuidOptional("netflow", row.Netflow)...)
	r = append(r, types.OvsdbCmdArgsMapStringString("other_config", row.OtherConfig)...)
	r = append(r, types.OvsdbCmdArgsUuidMultiples("ports", row.Ports)...)
	r = append(r, types.OvsdbCmdArgsStringMultiples("protocols", row.Protocols)...)
	r = append(r, types.OvsdbCmdArgsBoolean("rstp_enable", row.RstpEnable)...)
	r = append(r, types.OvsdbCmdArgsMapStringString("rstp_status", row.RstpStatus)...)
	r = append(r, types.OvsdbCmdArgsUuidOptional("sflow", row.Sflow)...)
	r = append(r, types.OvsdbCmdArgsMapStringString("status", row.Status)...)
	r = append(r, types.OvsdbCmdArgsBoolean("stp_enable", row.StpEnable)...)
	return r
}

func (row *Bridge) SetColumn(name string, val interface{}) (err error) {
	defer func() {
		if panicErr := recover(); panicErr != nil {
			err = errors.Wrapf(panicErr.(error), "%s: %#v", name, fmt.Sprintf("%#v", val))
		}
	}()
	switch name {
	case "_uuid":
		row.Uuid = types.EnsureUuid(val)
	case "_version":
		row.Version = types.EnsureUuid(val)
	case "auto_attach":
		row.AutoAttach = types.EnsureUuidOptional(val)
	case "controller":
		row.Controller = types.EnsureUuidMultiples(val)
	case "datapath_id":
		row.DatapathId = types.EnsureStringOptional(val)
	case "datapath_type":
		row.DatapathType = types.EnsureString(val)
	case "datapath_version":
		row.DatapathVersion = types.EnsureString(val)
	case "external_ids":
		row.ExternalIds = types.EnsureMapStringString(val)
	case "fail_mode":
		row.FailMode = types.EnsureStringOptional(val)
	case "flood_vlans":
		row.FloodVlans = types.EnsureIntegerMultiples(val)
	case "flow_tables":
		row.FlowTables = types.EnsureMapIntegerUuid(val)
	case "ipfix":
		row.Ipfix = types.EnsureUuidOptional(val)
	case "mcast_snooping_enable":
		row.McastSnoopingEnable = types.EnsureBoolean(val)
	case "mirrors":
		row.Mirrors = types.EnsureUuidMultiples(val)
	case "name":
		row.Name = types.EnsureString(val)
	case "netflow":
		row.Netflow = types.EnsureUuidOptional(val)
	case "other_config":
		row.OtherConfig = types.EnsureMapStringString(val)
	case "ports":
		row.Ports = types.EnsureUuidMultiples(val)
	case "protocols":
		row.Protocols = types.EnsureStringMultiples(val)
	case "rstp_enable":
		row.RstpEnable = types.EnsureBoolean(val)
	case "rstp_status":
		row.RstpStatus = types.EnsureMapStringString(val)
	case "sflow":
		row.Sflow = types.EnsureUuidOptional(val)
	case "status":
		row.Status = types.EnsureMapStringString(val)
	case "stp_enable":
		row.StpEnable = types.EnsureBoolean(val)
	default:
		panic(types.ErrUnknownColumn)
	}
	return
}

func (row *Bridge) MatchNonZeros(row1 *Bridge) bool {
	if !types.MatchUuidIfNonZero(row.Uuid, row1.Uuid) {
		return false
	}
	if !types.MatchUuidIfNonZero(row.Version, row1.Version) {
		return false
	}
	if !types.MatchUuidOptionalIfNonZero(row.AutoAttach, row1.AutoAttach) {
		return false
	}
	if !types.MatchUuidMultiplesIfNonZero(row.Controller, row1.Controller) {
		return false
	}
	if !types.MatchStringOptionalIfNonZero(row.DatapathId, row1.DatapathId) {
		return false
	}
	if !types.MatchStringIfNonZero(row.DatapathType, row1.DatapathType) {
		return false
	}
	if !types.MatchStringIfNonZero(row.DatapathVersion, row1.DatapathVersion) {
		return false
	}
	if !types.MatchMapStringStringIfNonZero(row.ExternalIds, row1.ExternalIds) {
		return false
	}
	if !types.MatchStringOptionalIfNonZero(row.FailMode, row1.FailMode) {
		return false
	}
	if !types.MatchIntegerMultiplesIfNonZero(row.FloodVlans, row1.FloodVlans) {
		return false
	}
	if !types.MatchMapIntegerUuidIfNonZero(row.FlowTables, row1.FlowTables) {
		return false
	}
	if !types.MatchUuidOptionalIfNonZero(row.Ipfix, row1.Ipfix) {
		return false
	}
	if !types.MatchBooleanIfNonZero(row.McastSnoopingEnable, row1.McastSnoopingEnable) {
		return false
	}
	if !types.MatchUuidMultiplesIfNonZero(row.Mirrors, row1.Mirrors) {
		return false
	}
	if !types.MatchStringIfNonZero(row.Name, row1.Name) {
		return false
	}
	if !types.MatchUuidOptionalIfNonZero(row.Netflow, row1.Netflow) {
		return false
	}
	if !types.MatchMapStringStringIfNonZero(row.OtherConfig, row1.OtherConfig) {
		return false
	}
	if !types.MatchUuidMultiplesIfNonZero(row.Ports, row1.Ports) {
		return false
	}
	if !types.MatchStringMultiplesIfNonZero(row.Protocols, row1.Protocols) {
		return false
	}
	if !types.MatchBooleanIfNonZero(row.RstpEnable, row1.RstpEnable) {
		return false
	}
	if !types.MatchMapStringStringIfNonZero(row.RstpStatus, row1.RstpStatus) {
		return false
	}
	if !types.MatchUuidOptionalIfNonZero(row.Sflow, row1.Sflow) {
		return false
	}
	if !types.MatchMapStringStringIfNonZero(row.Status, row1.Status) {
		return false
	}
	if !types.MatchBooleanIfNonZero(row.StpEnable, row1.StpEnable) {
		return false
	}
	return true
}

func (row *Bridge) HasExternalIds() bool {
	return true
}

func (row *Bridge) SetExternalId(k, v string) {
	if row.ExternalIds == nil {
		row.ExternalIds = map[string]string{}
	}
	row.ExternalIds[k] = v
}

func (row *Bridge) GetExternalId(k string) (string, bool) {
	if row.ExternalIds == nil {
		return "", false
	}
	r, ok := row.ExternalIds[k]
	return r, ok
}

func (row *Bridge) RemoveExternalId(k string) (string, bool) {
	if row.ExternalIds == nil {
		return "", false
	}
	r, ok := row.ExternalIds[k]
	if ok {
		delete(row.ExternalIds, k)
	}
	return r, ok
}

type ControllerTable []Controller

var _ types.ITable = &ControllerTable{}

func (tbl ControllerTable) OvsdbTableName() string {
	return "Controller"
}

func (tbl ControllerTable) OvsdbIsRoot() bool {
	return false
}

func (tbl ControllerTable) Rows() []types.IRow {
	r := make([]types.IRow, len(tbl))
	for i := range tbl {
		r[i] = &tbl[i]
	}
	return r
}

func (tbl ControllerTable) NewRow() types.IRow {
	return &Controller{}
}

func (tbl *ControllerTable) AppendRow(irow types.IRow) {
	row := irow.(*Controller)
	*tbl = append(*tbl, *row)
}

func (tbl ControllerTable) OvsdbHasIndex() bool {
	return false
}

func (tbl ControllerTable) OvsdbGetByAnyIndex(irow1 types.IRow) types.IRow {
	return nil
}

func (tbl ControllerTable) FindOneMatchNonZeros(row1 *Controller) *Controller {
	for i := range tbl {
		row := &tbl[i]
		if row.MatchNonZeros(row1) {
			return row
		}
	}
	return nil
}

type Controller struct {
	Uuid                 string            `json:"_uuid"`
	Version              string            `json:"_version"`
	ConnectionMode       *string           `json:"connection_mode"`
	ControllerBurstLimit *int64            `json:"controller_burst_limit"`
	ControllerRateLimit  *int64            `json:"controller_rate_limit"`
	EnableAsyncMessages  *bool             `json:"enable_async_messages"`
	ExternalIds          map[string]string `json:"external_ids"`
	InactivityProbe      *int64            `json:"inactivity_probe"`
	IsConnected          bool              `json:"is_connected"`
	LocalGateway         *string           `json:"local_gateway"`
	LocalIp              *string           `json:"local_ip"`
	LocalNetmask         *string           `json:"local_netmask"`
	MaxBackoff           *int64            `json:"max_backoff"`
	OtherConfig          map[string]string `json:"other_config"`
	Role                 *string           `json:"role"`
	Status               map[string]string `json:"status"`
	Target               string            `json:"target"`
}

var _ types.IRow = &Controller{}

func (row *Controller) OvsdbTableName() string {
	return "Controller"
}

func (row *Controller) OvsdbIsRoot() bool {
	return false
}

func (row *Controller) OvsdbUuid() string {
	return row.Uuid
}

func (row *Controller) OvsdbCmdArgs() []string {
	r := []string{}
	r = append(r, types.OvsdbCmdArgsStringOptional("connection_mode", row.ConnectionMode)...)
	r = append(r, types.OvsdbCmdArgsIntegerOptional("controller_burst_limit", row.ControllerBurstLimit)...)
	r = append(r, types.OvsdbCmdArgsIntegerOptional("controller_rate_limit", row.ControllerRateLimit)...)
	r = append(r, types.OvsdbCmdArgsBooleanOptional("enable_async_messages", row.EnableAsyncMessages)...)
	r = append(r, types.OvsdbCmdArgsMapStringString("external_ids", row.ExternalIds)...)
	r = append(r, types.OvsdbCmdArgsIntegerOptional("inactivity_probe", row.InactivityProbe)...)
	r = append(r, types.OvsdbCmdArgsBoolean("is_connected", row.IsConnected)...)
	r = append(r, types.OvsdbCmdArgsStringOptional("local_gateway", row.LocalGateway)...)
	r = append(r, types.OvsdbCmdArgsStringOptional("local_ip", row.LocalIp)...)
	r = append(r, types.OvsdbCmdArgsStringOptional("local_netmask", row.LocalNetmask)...)
	r = append(r, types.OvsdbCmdArgsIntegerOptional("max_backoff", row.MaxBackoff)...)
	r = append(r, types.OvsdbCmdArgsMapStringString("other_config", row.OtherConfig)...)
	r = append(r, types.OvsdbCmdArgsStringOptional("role", row.Role)...)
	r = append(r, types.OvsdbCmdArgsMapStringString("status", row.Status)...)
	r = append(r, types.OvsdbCmdArgsString("target", row.Target)...)
	return r
}

func (row *Controller) SetColumn(name string, val interface{}) (err error) {
	defer func() {
		if panicErr := recover(); panicErr != nil {
			err = errors.Wrapf(panicErr.(error), "%s: %#v", name, fmt.Sprintf("%#v", val))
		}
	}()
	switch name {
	case "_uuid":
		row.Uuid = types.EnsureUuid(val)
	case "_version":
		row.Version = types.EnsureUuid(val)
	case "connection_mode":
		row.ConnectionMode = types.EnsureStringOptional(val)
	case "controller_burst_limit":
		row.ControllerBurstLimit = types.EnsureIntegerOptional(val)
	case "controller_rate_limit":
		row.ControllerRateLimit = types.EnsureIntegerOptional(val)
	case "enable_async_messages":
		row.EnableAsyncMessages = types.EnsureBooleanOptional(val)
	case "external_ids":
		row.ExternalIds = types.EnsureMapStringString(val)
	case "inactivity_probe":
		row.InactivityProbe = types.EnsureIntegerOptional(val)
	case "is_connected":
		row.IsConnected = types.EnsureBoolean(val)
	case "local_gateway":
		row.LocalGateway = types.EnsureStringOptional(val)
	case "local_ip":
		row.LocalIp = types.EnsureStringOptional(val)
	case "local_netmask":
		row.LocalNetmask = types.EnsureStringOptional(val)
	case "max_backoff":
		row.MaxBackoff = types.EnsureIntegerOptional(val)
	case "other_config":
		row.OtherConfig = types.EnsureMapStringString(val)
	case "role":
		row.Role = types.EnsureStringOptional(val)
	case "status":
		row.Status = types.EnsureMapStringString(val)
	case "target":
		row.Target = types.EnsureString(val)
	default:
		panic(types.ErrUnknownColumn)
	}
	return
}

func (row *Controller) MatchNonZeros(row1 *Controller) bool {
	if !types.MatchUuidIfNonZero(row.Uuid, row1.Uuid) {
		return false
	}
	if !types.MatchUuidIfNonZero(row.Version, row1.Version) {
		return false
	}
	if !types.MatchStringOptionalIfNonZero(row.ConnectionMode, row1.ConnectionMode) {
		return false
	}
	if !types.MatchIntegerOptionalIfNonZero(row.ControllerBurstLimit, row1.ControllerBurstLimit) {
		return false
	}
	if !types.MatchIntegerOptionalIfNonZero(row.ControllerRateLimit, row1.ControllerRateLimit) {
		return false
	}
	if !types.MatchBooleanOptionalIfNonZero(row.EnableAsyncMessages, row1.EnableAsyncMessages) {
		return false
	}
	if !types.MatchMapStringStringIfNonZero(row.ExternalIds, row1.ExternalIds) {
		return false
	}
	if !types.MatchIntegerOptionalIfNonZero(row.InactivityProbe, row1.InactivityProbe) {
		return false
	}
	if !types.MatchBooleanIfNonZero(row.IsConnected, row1.IsConnected) {
		return false
	}
	if !types.MatchStringOptionalIfNonZero(row.LocalGateway, row1.LocalGateway) {
		return false
	}
	if !types.MatchStringOptionalIfNonZero(row.LocalIp, row1.LocalIp) {
		return false
	}
	if !types.MatchStringOptionalIfNonZero(row.LocalNetmask, row1.LocalNetmask) {
		return false
	}
	if !types.MatchIntegerOptionalIfNonZero(row.MaxBackoff, row1.MaxBackoff) {
		return false
	}
	if !types.MatchMapStringStringIfNonZero(row.OtherConfig, row1.OtherConfig) {
		return false
	}
	if !types.MatchStringOptionalIfNonZero(row.Role, row1.Role) {
		return false
	}
	if !types.MatchMapStringStringIfNonZero(row.Status, row1.Status) {
		return false
	}
	if !types.MatchStringIfNonZero(row.Target, row1.Target) {
		return false
	}
	return true
}

func (row *Controller) HasExternalIds() bool {
	return true
}

func (row *Controller) SetExternalId(k, v string) {
	if row.ExternalIds == nil {
		row.ExternalIds = map[string]string{}
	}
	row.ExternalIds[k] = v
}

func (row *Controller) GetExternalId(k string) (string, bool) {
	if row.ExternalIds == nil {
		return "", false
	}
	r, ok := row.ExternalIds[k]
	return r, ok
}

func (row *Controller) RemoveExternalId(k string) (string, bool) {
	if row.ExternalIds == nil {
		return "", false
	}
	r, ok := row.ExternalIds[k]
	if ok {
		delete(row.ExternalIds, k)
	}
	return r, ok
}

type FlowSampleCollectorSetTable []FlowSampleCollectorSet

var _ types.ITable = &FlowSampleCollectorSetTable{}

func (tbl FlowSampleCollectorSetTable) OvsdbTableName() string {
	return "Flow_Sample_Collector_Set"
}

func (tbl FlowSampleCollectorSetTable) OvsdbIsRoot() bool {
	return true
}

func (tbl FlowSampleCollectorSetTable) Rows() []types.IRow {
	r := make([]types.IRow, len(tbl))
	for i := range tbl {
		r[i] = &tbl[i]
	}
	return r
}

func (tbl FlowSampleCollectorSetTable) NewRow() types.IRow {
	return &FlowSampleCollectorSet{}
}

func (tbl *FlowSampleCollectorSetTable) AppendRow(irow types.IRow) {
	row := irow.(*FlowSampleCollectorSet)
	*tbl = append(*tbl, *row)
}

func (tbl FlowSampleCollectorSetTable) OvsdbHasIndex() bool {
	return true
}

func (row *FlowSampleCollectorSet) MatchByIdBridge(row1 *FlowSampleCollectorSet) bool {
	if !types.MatchInteger(row.Id, row1.Id) {
		return false
	}
	if !types.MatchUuid(row.Bridge, row1.Bridge) {
		return false
	}
	return true
}

func (tbl FlowSampleCollectorSetTable) GetByIdBridge(row1 *FlowSampleCollectorSet) *FlowSampleCollectorSet {
	for i := range tbl {
		row := &tbl[i]
		if row.MatchByIdBridge(row1) {
			return row
		}
	}
	return nil
}

func (tbl FlowSampleCollectorSetTable) OvsdbGetByAnyIndex(irow1 types.IRow) types.IRow {
	row1 := irow1.(*FlowSampleCollectorSet)
	if !(types.IsZeroInteger(row1.Id) || types.IsZeroUuid(row1.Bridge)) {
		if row := tbl.GetByIdBridge(row1); row != nil {
			return row
		}
	}
	return nil
}

func (tbl FlowSampleCollectorSetTable) FindOneMatchNonZeros(row1 *FlowSampleCollectorSet) *FlowSampleCollectorSet {
	for i := range tbl {
		row := &tbl[i]
		if row.MatchNonZeros(row1) {
			return row
		}
	}
	return nil
}

type FlowSampleCollectorSet struct {
	Uuid        string            `json:"_uuid"`
	Version     string            `json:"_version"`
	Bridge      string            `json:"bridge"`
	ExternalIds map[string]string `json:"external_ids"`
	Id          int64             `json:"id"`
	Ipfix       *string           `json:"ipfix"`
}

var _ types.IRow = &FlowSampleCollectorSet{}

func (row *FlowSampleCollectorSet) OvsdbTableName() string {
	return "Flow_Sample_Collector_Set"
}

func (row *FlowSampleCollectorSet) OvsdbIsRoot() bool {
	return true
}

func (row *FlowSampleCollectorSet) OvsdbUuid() string {
	return row.Uuid
}

func (row *FlowSampleCollectorSet) OvsdbCmdArgs() []string {
	r := []string{}
	r = append(r, types.OvsdbCmdArgsUuid("bridge", row.Bridge)...)
	r = append(r, types.OvsdbCmdArgsMapStringString("external_ids", row.ExternalIds)...)
	r = append(r, types.OvsdbCmdArgsInteger("id", row.Id)...)
	r = append(r, types.OvsdbCmdArgsUuidOptional("ipfix", row.Ipfix)...)
	return r
}

func (row *FlowSampleCollectorSet) SetColumn(name string, val interface{}) (err error) {
	defer func() {
		if panicErr := recover(); panicErr != nil {
			err = errors.Wrapf(panicErr.(error), "%s: %#v", name, fmt.Sprintf("%#v", val))
		}
	}()
	switch name {
	case "_uuid":
		row.Uuid = types.EnsureUuid(val)
	case "_version":
		row.Version = types.EnsureUuid(val)
	case "bridge":
		row.Bridge = types.EnsureUuid(val)
	case "external_ids":
		row.ExternalIds = types.EnsureMapStringString(val)
	case "id":
		row.Id = types.EnsureInteger(val)
	case "ipfix":
		row.Ipfix = types.EnsureUuidOptional(val)
	default:
		panic(types.ErrUnknownColumn)
	}
	return
}

func (row *FlowSampleCollectorSet) MatchNonZeros(row1 *FlowSampleCollectorSet) bool {
	if !types.MatchUuidIfNonZero(row.Uuid, row1.Uuid) {
		return false
	}
	if !types.MatchUuidIfNonZero(row.Version, row1.Version) {
		return false
	}
	if !types.MatchUuidIfNonZero(row.Bridge, row1.Bridge) {
		return false
	}
	if !types.MatchMapStringStringIfNonZero(row.ExternalIds, row1.ExternalIds) {
		return false
	}
	if !types.MatchIntegerIfNonZero(row.Id, row1.Id) {
		return false
	}
	if !types.MatchUuidOptionalIfNonZero(row.Ipfix, row1.Ipfix) {
		return false
	}
	return true
}

func (row *FlowSampleCollectorSet) HasExternalIds() bool {
	return true
}

func (row *FlowSampleCollectorSet) SetExternalId(k, v string) {
	if row.ExternalIds == nil {
		row.ExternalIds = map[string]string{}
	}
	row.ExternalIds[k] = v
}

func (row *FlowSampleCollectorSet) GetExternalId(k string) (string, bool) {
	if row.ExternalIds == nil {
		return "", false
	}
	r, ok := row.ExternalIds[k]
	return r, ok
}

func (row *FlowSampleCollectorSet) RemoveExternalId(k string) (string, bool) {
	if row.ExternalIds == nil {
		return "", false
	}
	r, ok := row.ExternalIds[k]
	if ok {
		delete(row.ExternalIds, k)
	}
	return r, ok
}

type FlowTableTable []FlowTable

var _ types.ITable = &FlowTableTable{}

func (tbl FlowTableTable) OvsdbTableName() string {
	return "Flow_Table"
}

func (tbl FlowTableTable) OvsdbIsRoot() bool {
	return false
}

func (tbl FlowTableTable) Rows() []types.IRow {
	r := make([]types.IRow, len(tbl))
	for i := range tbl {
		r[i] = &tbl[i]
	}
	return r
}

func (tbl FlowTableTable) NewRow() types.IRow {
	return &FlowTable{}
}

func (tbl *FlowTableTable) AppendRow(irow types.IRow) {
	row := irow.(*FlowTable)
	*tbl = append(*tbl, *row)
}

func (tbl FlowTableTable) OvsdbHasIndex() bool {
	return false
}

func (tbl FlowTableTable) OvsdbGetByAnyIndex(irow1 types.IRow) types.IRow {
	return nil
}

func (tbl FlowTableTable) FindOneMatchNonZeros(row1 *FlowTable) *FlowTable {
	for i := range tbl {
		row := &tbl[i]
		if row.MatchNonZeros(row1) {
			return row
		}
	}
	return nil
}

type FlowTable struct {
	Uuid           string            `json:"_uuid"`
	Version        string            `json:"_version"`
	ExternalIds    map[string]string `json:"external_ids"`
	FlowLimit      *int64            `json:"flow_limit"`
	Groups         []string          `json:"groups"`
	Name           *string           `json:"name"`
	OverflowPolicy *string           `json:"overflow_policy"`
	Prefixes       []string          `json:"prefixes"`
}

var _ types.IRow = &FlowTable{}

func (row *FlowTable) OvsdbTableName() string {
	return "Flow_Table"
}

func (row *FlowTable) OvsdbIsRoot() bool {
	return false
}

func (row *FlowTable) OvsdbUuid() string {
	return row.Uuid
}

func (row *FlowTable) OvsdbCmdArgs() []string {
	r := []string{}
	r = append(r, types.OvsdbCmdArgsMapStringString("external_ids", row.ExternalIds)...)
	r = append(r, types.OvsdbCmdArgsIntegerOptional("flow_limit", row.FlowLimit)...)
	r = append(r, types.OvsdbCmdArgsStringMultiples("groups", row.Groups)...)
	r = append(r, types.OvsdbCmdArgsStringOptional("name", row.Name)...)
	r = append(r, types.OvsdbCmdArgsStringOptional("overflow_policy", row.OverflowPolicy)...)
	r = append(r, types.OvsdbCmdArgsStringMultiples("prefixes", row.Prefixes)...)
	return r
}

func (row *FlowTable) SetColumn(name string, val interface{}) (err error) {
	defer func() {
		if panicErr := recover(); panicErr != nil {
			err = errors.Wrapf(panicErr.(error), "%s: %#v", name, fmt.Sprintf("%#v", val))
		}
	}()
	switch name {
	case "_uuid":
		row.Uuid = types.EnsureUuid(val)
	case "_version":
		row.Version = types.EnsureUuid(val)
	case "external_ids":
		row.ExternalIds = types.EnsureMapStringString(val)
	case "flow_limit":
		row.FlowLimit = types.EnsureIntegerOptional(val)
	case "groups":
		row.Groups = types.EnsureStringMultiples(val)
	case "name":
		row.Name = types.EnsureStringOptional(val)
	case "overflow_policy":
		row.OverflowPolicy = types.EnsureStringOptional(val)
	case "prefixes":
		row.Prefixes = types.EnsureStringMultiples(val)
	default:
		panic(types.ErrUnknownColumn)
	}
	return
}

func (row *FlowTable) MatchNonZeros(row1 *FlowTable) bool {
	if !types.MatchUuidIfNonZero(row.Uuid, row1.Uuid) {
		return false
	}
	if !types.MatchUuidIfNonZero(row.Version, row1.Version) {
		return false
	}
	if !types.MatchMapStringStringIfNonZero(row.ExternalIds, row1.ExternalIds) {
		return false
	}
	if !types.MatchIntegerOptionalIfNonZero(row.FlowLimit, row1.FlowLimit) {
		return false
	}
	if !types.MatchStringMultiplesIfNonZero(row.Groups, row1.Groups) {
		return false
	}
	if !types.MatchStringOptionalIfNonZero(row.Name, row1.Name) {
		return false
	}
	if !types.MatchStringOptionalIfNonZero(row.OverflowPolicy, row1.OverflowPolicy) {
		return false
	}
	if !types.MatchStringMultiplesIfNonZero(row.Prefixes, row1.Prefixes) {
		return false
	}
	return true
}

func (row *FlowTable) HasExternalIds() bool {
	return true
}

func (row *FlowTable) SetExternalId(k, v string) {
	if row.ExternalIds == nil {
		row.ExternalIds = map[string]string{}
	}
	row.ExternalIds[k] = v
}

func (row *FlowTable) GetExternalId(k string) (string, bool) {
	if row.ExternalIds == nil {
		return "", false
	}
	r, ok := row.ExternalIds[k]
	return r, ok
}

func (row *FlowTable) RemoveExternalId(k string) (string, bool) {
	if row.ExternalIds == nil {
		return "", false
	}
	r, ok := row.ExternalIds[k]
	if ok {
		delete(row.ExternalIds, k)
	}
	return r, ok
}

type IPFIXTable []IPFIX

var _ types.ITable = &IPFIXTable{}

func (tbl IPFIXTable) OvsdbTableName() string {
	return "IPFIX"
}

func (tbl IPFIXTable) OvsdbIsRoot() bool {
	return false
}

func (tbl IPFIXTable) Rows() []types.IRow {
	r := make([]types.IRow, len(tbl))
	for i := range tbl {
		r[i] = &tbl[i]
	}
	return r
}

func (tbl IPFIXTable) NewRow() types.IRow {
	return &IPFIX{}
}

func (tbl *IPFIXTable) AppendRow(irow types.IRow) {
	row := irow.(*IPFIX)
	*tbl = append(*tbl, *row)
}

func (tbl IPFIXTable) OvsdbHasIndex() bool {
	return false
}

func (tbl IPFIXTable) OvsdbGetByAnyIndex(irow1 types.IRow) types.IRow {
	return nil
}

func (tbl IPFIXTable) FindOneMatchNonZeros(row1 *IPFIX) *IPFIX {
	for i := range tbl {
		row := &tbl[i]
		if row.MatchNonZeros(row1) {
			return row
		}
	}
	return nil
}

type IPFIX struct {
	Uuid               string            `json:"_uuid"`
	Version            string            `json:"_version"`
	CacheActiveTimeout *int64            `json:"cache_active_timeout"`
	CacheMaxFlows      *int64            `json:"cache_max_flows"`
	ExternalIds        map[string]string `json:"external_ids"`
	ObsDomainId        *int64            `json:"obs_domain_id"`
	ObsPointId         *int64            `json:"obs_point_id"`
	OtherConfig        map[string]string `json:"other_config"`
	Sampling           *int64            `json:"sampling"`
	Targets            []string          `json:"targets"`
}

var _ types.IRow = &IPFIX{}

func (row *IPFIX) OvsdbTableName() string {
	return "IPFIX"
}

func (row *IPFIX) OvsdbIsRoot() bool {
	return false
}

func (row *IPFIX) OvsdbUuid() string {
	return row.Uuid
}

func (row *IPFIX) OvsdbCmdArgs() []string {
	r := []string{}
	r = append(r, types.OvsdbCmdArgsIntegerOptional("cache_active_timeout", row.CacheActiveTimeout)...)
	r = append(r, types.OvsdbCmdArgsIntegerOptional("cache_max_flows", row.CacheMaxFlows)...)
	r = append(r, types.OvsdbCmdArgsMapStringString("external_ids", row.ExternalIds)...)
	r = append(r, types.OvsdbCmdArgsIntegerOptional("obs_domain_id", row.ObsDomainId)...)
	r = append(r, types.OvsdbCmdArgsIntegerOptional("obs_point_id", row.ObsPointId)...)
	r = append(r, types.OvsdbCmdArgsMapStringString("other_config", row.OtherConfig)...)
	r = append(r, types.OvsdbCmdArgsIntegerOptional("sampling", row.Sampling)...)
	r = append(r, types.OvsdbCmdArgsStringMultiples("targets", row.Targets)...)
	return r
}

func (row *IPFIX) SetColumn(name string, val interface{}) (err error) {
	defer func() {
		if panicErr := recover(); panicErr != nil {
			err = errors.Wrapf(panicErr.(error), "%s: %#v", name, fmt.Sprintf("%#v", val))
		}
	}()
	switch name {
	case "_uuid":
		row.Uuid = types.EnsureUuid(val)
	case "_version":
		row.Version = types.EnsureUuid(val)
	case "cache_active_timeout":
		row.CacheActiveTimeout = types.EnsureIntegerOptional(val)
	case "cache_max_flows":
		row.CacheMaxFlows = types.EnsureIntegerOptional(val)
	case "external_ids":
		row.ExternalIds = types.EnsureMapStringString(val)
	case "obs_domain_id":
		row.ObsDomainId = types.EnsureIntegerOptional(val)
	case "obs_point_id":
		row.ObsPointId = types.EnsureIntegerOptional(val)
	case "other_config":
		row.OtherConfig = types.EnsureMapStringString(val)
	case "sampling":
		row.Sampling = types.EnsureIntegerOptional(val)
	case "targets":
		row.Targets = types.EnsureStringMultiples(val)
	default:
		panic(types.ErrUnknownColumn)
	}
	return
}

func (row *IPFIX) MatchNonZeros(row1 *IPFIX) bool {
	if !types.MatchUuidIfNonZero(row.Uuid, row1.Uuid) {
		return false
	}
	if !types.MatchUuidIfNonZero(row.Version, row1.Version) {
		return false
	}
	if !types.MatchIntegerOptionalIfNonZero(row.CacheActiveTimeout, row1.CacheActiveTimeout) {
		return false
	}
	if !types.MatchIntegerOptionalIfNonZero(row.CacheMaxFlows, row1.CacheMaxFlows) {
		return false
	}
	if !types.MatchMapStringStringIfNonZero(row.ExternalIds, row1.ExternalIds) {
		return false
	}
	if !types.MatchIntegerOptionalIfNonZero(row.ObsDomainId, row1.ObsDomainId) {
		return false
	}
	if !types.MatchIntegerOptionalIfNonZero(row.ObsPointId, row1.ObsPointId) {
		return false
	}
	if !types.MatchMapStringStringIfNonZero(row.OtherConfig, row1.OtherConfig) {
		return false
	}
	if !types.MatchIntegerOptionalIfNonZero(row.Sampling, row1.Sampling) {
		return false
	}
	if !types.MatchStringMultiplesIfNonZero(row.Targets, row1.Targets) {
		return false
	}
	return true
}

func (row *IPFIX) HasExternalIds() bool {
	return true
}

func (row *IPFIX) SetExternalId(k, v string) {
	if row.ExternalIds == nil {
		row.ExternalIds = map[string]string{}
	}
	row.ExternalIds[k] = v
}

func (row *IPFIX) GetExternalId(k string) (string, bool) {
	if row.ExternalIds == nil {
		return "", false
	}
	r, ok := row.ExternalIds[k]
	return r, ok
}

func (row *IPFIX) RemoveExternalId(k string) (string, bool) {
	if row.ExternalIds == nil {
		return "", false
	}
	r, ok := row.ExternalIds[k]
	if ok {
		delete(row.ExternalIds, k)
	}
	return r, ok
}

type InterfaceTable []Interface

var _ types.ITable = &InterfaceTable{}

func (tbl InterfaceTable) OvsdbTableName() string {
	return "Interface"
}

func (tbl InterfaceTable) OvsdbIsRoot() bool {
	return false
}

func (tbl InterfaceTable) Rows() []types.IRow {
	r := make([]types.IRow, len(tbl))
	for i := range tbl {
		r[i] = &tbl[i]
	}
	return r
}

func (tbl InterfaceTable) NewRow() types.IRow {
	return &Interface{}
}

func (tbl *InterfaceTable) AppendRow(irow types.IRow) {
	row := irow.(*Interface)
	*tbl = append(*tbl, *row)
}

func (tbl InterfaceTable) OvsdbHasIndex() bool {
	return true
}

func (row *Interface) MatchByName(row1 *Interface) bool {
	if !types.MatchString(row.Name, row1.Name) {
		return false
	}
	return true
}

func (tbl InterfaceTable) GetByName(row1 *Interface) *Interface {
	for i := range tbl {
		row := &tbl[i]
		if row.MatchByName(row1) {
			return row
		}
	}
	return nil
}

func (tbl InterfaceTable) OvsdbGetByAnyIndex(irow1 types.IRow) types.IRow {
	row1 := irow1.(*Interface)
	if !(types.IsZeroString(row1.Name)) {
		if row := tbl.GetByName(row1); row != nil {
			return row
		}
	}
	return nil
}

func (tbl InterfaceTable) FindOneMatchNonZeros(row1 *Interface) *Interface {
	for i := range tbl {
		row := &tbl[i]
		if row.MatchNonZeros(row1) {
			return row
		}
	}
	return nil
}

type Interface struct {
	Uuid                 string            `json:"_uuid"`
	Version              string            `json:"_version"`
	AdminState           *string           `json:"admin_state"`
	Bfd                  map[string]string `json:"bfd"`
	BfdStatus            map[string]string `json:"bfd_status"`
	CfmFault             *bool             `json:"cfm_fault"`
	CfmFaultStatus       []string          `json:"cfm_fault_status"`
	CfmFlapCount         *int64            `json:"cfm_flap_count"`
	CfmHealth            *int64            `json:"cfm_health"`
	CfmMpid              *int64            `json:"cfm_mpid"`
	CfmRemoteMpids       []int64           `json:"cfm_remote_mpids"`
	CfmRemoteOpstate     *string           `json:"cfm_remote_opstate"`
	Duplex               *string           `json:"duplex"`
	Error                *string           `json:"error"`
	ExternalIds          map[string]string `json:"external_ids"`
	Ifindex              *int64            `json:"ifindex"`
	IngressPolicingBurst int64             `json:"ingress_policing_burst"`
	IngressPolicingRate  int64             `json:"ingress_policing_rate"`
	LacpCurrent          *bool             `json:"lacp_current"`
	LinkResets           *int64            `json:"link_resets"`
	LinkSpeed            *int64            `json:"link_speed"`
	LinkState            *string           `json:"link_state"`
	Lldp                 map[string]string `json:"lldp"`
	Mac                  *string           `json:"mac"`
	MacInUse             *string           `json:"mac_in_use"`
	Mtu                  *int64            `json:"mtu"`
	MtuRequest           *int64            `json:"mtu_request"`
	Name                 string            `json:"name"`
	Ofport               *int64            `json:"ofport"`
	OfportRequest        *int64            `json:"ofport_request"`
	Options              map[string]string `json:"options"`
	OtherConfig          map[string]string `json:"other_config"`
	Statistics           map[string]int64  `json:"statistics"`
	Status               map[string]string `json:"status"`
	Type                 string            `json:"type"`
}

var _ types.IRow = &Interface{}

func (row *Interface) OvsdbTableName() string {
	return "Interface"
}

func (row *Interface) OvsdbIsRoot() bool {
	return false
}

func (row *Interface) OvsdbUuid() string {
	return row.Uuid
}

func (row *Interface) OvsdbCmdArgs() []string {
	r := []string{}
	r = append(r, types.OvsdbCmdArgsStringOptional("admin_state", row.AdminState)...)
	r = append(r, types.OvsdbCmdArgsMapStringString("bfd", row.Bfd)...)
	r = append(r, types.OvsdbCmdArgsMapStringString("bfd_status", row.BfdStatus)...)
	r = append(r, types.OvsdbCmdArgsBooleanOptional("cfm_fault", row.CfmFault)...)
	r = append(r, types.OvsdbCmdArgsStringMultiples("cfm_fault_status", row.CfmFaultStatus)...)
	r = append(r, types.OvsdbCmdArgsIntegerOptional("cfm_flap_count", row.CfmFlapCount)...)
	r = append(r, types.OvsdbCmdArgsIntegerOptional("cfm_health", row.CfmHealth)...)
	r = append(r, types.OvsdbCmdArgsIntegerOptional("cfm_mpid", row.CfmMpid)...)
	r = append(r, types.OvsdbCmdArgsIntegerMultiples("cfm_remote_mpids", row.CfmRemoteMpids)...)
	r = append(r, types.OvsdbCmdArgsStringOptional("cfm_remote_opstate", row.CfmRemoteOpstate)...)
	r = append(r, types.OvsdbCmdArgsStringOptional("duplex", row.Duplex)...)
	r = append(r, types.OvsdbCmdArgsStringOptional("error", row.Error)...)
	r = append(r, types.OvsdbCmdArgsMapStringString("external_ids", row.ExternalIds)...)
	r = append(r, types.OvsdbCmdArgsIntegerOptional("ifindex", row.Ifindex)...)
	r = append(r, types.OvsdbCmdArgsInteger("ingress_policing_burst", row.IngressPolicingBurst)...)
	r = append(r, types.OvsdbCmdArgsInteger("ingress_policing_rate", row.IngressPolicingRate)...)
	r = append(r, types.OvsdbCmdArgsBooleanOptional("lacp_current", row.LacpCurrent)...)
	r = append(r, types.OvsdbCmdArgsIntegerOptional("link_resets", row.LinkResets)...)
	r = append(r, types.OvsdbCmdArgsIntegerOptional("link_speed", row.LinkSpeed)...)
	r = append(r, types.OvsdbCmdArgsStringOptional("link_state", row.LinkState)...)
	r = append(r, types.OvsdbCmdArgsMapStringString("lldp", row.Lldp)...)
	r = append(r, types.OvsdbCmdArgsStringOptional("mac", row.Mac)...)
	r = append(r, types.OvsdbCmdArgsStringOptional("mac_in_use", row.MacInUse)...)
	r = append(r, types.OvsdbCmdArgsIntegerOptional("mtu", row.Mtu)...)
	r = append(r, types.OvsdbCmdArgsIntegerOptional("mtu_request", row.MtuRequest)...)
	r = append(r, types.OvsdbCmdArgsString("name", row.Name)...)
	r = append(r, types.OvsdbCmdArgsIntegerOptional("ofport", row.Ofport)...)
	r = append(r, types.OvsdbCmdArgsIntegerOptional("ofport_request", row.OfportRequest)...)
	r = append(r, types.OvsdbCmdArgsMapStringString("options", row.Options)...)
	r = append(r, types.OvsdbCmdArgsMapStringString("other_config", row.OtherConfig)...)
	r = append(r, types.OvsdbCmdArgsMapStringInteger("statistics", row.Statistics)...)
	r = append(r, types.OvsdbCmdArgsMapStringString("status", row.Status)...)
	r = append(r, types.OvsdbCmdArgsString("type", row.Type)...)
	return r
}

func (row *Interface) SetColumn(name string, val interface{}) (err error) {
	defer func() {
		if panicErr := recover(); panicErr != nil {
			err = errors.Wrapf(panicErr.(error), "%s: %#v", name, fmt.Sprintf("%#v", val))
		}
	}()
	switch name {
	case "_uuid":
		row.Uuid = types.EnsureUuid(val)
	case "_version":
		row.Version = types.EnsureUuid(val)
	case "admin_state":
		row.AdminState = types.EnsureStringOptional(val)
	case "bfd":
		row.Bfd = types.EnsureMapStringString(val)
	case "bfd_status":
		row.BfdStatus = types.EnsureMapStringString(val)
	case "cfm_fault":
		row.CfmFault = types.EnsureBooleanOptional(val)
	case "cfm_fault_status":
		row.CfmFaultStatus = types.EnsureStringMultiples(val)
	case "cfm_flap_count":
		row.CfmFlapCount = types.EnsureIntegerOptional(val)
	case "cfm_health":
		row.CfmHealth = types.EnsureIntegerOptional(val)
	case "cfm_mpid":
		row.CfmMpid = types.EnsureIntegerOptional(val)
	case "cfm_remote_mpids":
		row.CfmRemoteMpids = types.EnsureIntegerMultiples(val)
	case "cfm_remote_opstate":
		row.CfmRemoteOpstate = types.EnsureStringOptional(val)
	case "duplex":
		row.Duplex = types.EnsureStringOptional(val)
	case "error":
		row.Error = types.EnsureStringOptional(val)
	case "external_ids":
		row.ExternalIds = types.EnsureMapStringString(val)
	case "ifindex":
		row.Ifindex = types.EnsureIntegerOptional(val)
	case "ingress_policing_burst":
		row.IngressPolicingBurst = types.EnsureInteger(val)
	case "ingress_policing_rate":
		row.IngressPolicingRate = types.EnsureInteger(val)
	case "lacp_current":
		row.LacpCurrent = types.EnsureBooleanOptional(val)
	case "link_resets":
		row.LinkResets = types.EnsureIntegerOptional(val)
	case "link_speed":
		row.LinkSpeed = types.EnsureIntegerOptional(val)
	case "link_state":
		row.LinkState = types.EnsureStringOptional(val)
	case "lldp":
		row.Lldp = types.EnsureMapStringString(val)
	case "mac":
		row.Mac = types.EnsureStringOptional(val)
	case "mac_in_use":
		row.MacInUse = types.EnsureStringOptional(val)
	case "mtu":
		row.Mtu = types.EnsureIntegerOptional(val)
	case "mtu_request":
		row.MtuRequest = types.EnsureIntegerOptional(val)
	case "name":
		row.Name = types.EnsureString(val)
	case "ofport":
		row.Ofport = types.EnsureIntegerOptional(val)
	case "ofport_request":
		row.OfportRequest = types.EnsureIntegerOptional(val)
	case "options":
		row.Options = types.EnsureMapStringString(val)
	case "other_config":
		row.OtherConfig = types.EnsureMapStringString(val)
	case "statistics":
		row.Statistics = types.EnsureMapStringInteger(val)
	case "status":
		row.Status = types.EnsureMapStringString(val)
	case "type":
		row.Type = types.EnsureString(val)
	default:
		panic(types.ErrUnknownColumn)
	}
	return
}

func (row *Interface) MatchNonZeros(row1 *Interface) bool {
	if !types.MatchUuidIfNonZero(row.Uuid, row1.Uuid) {
		return false
	}
	if !types.MatchUuidIfNonZero(row.Version, row1.Version) {
		return false
	}
	if !types.MatchStringOptionalIfNonZero(row.AdminState, row1.AdminState) {
		return false
	}
	if !types.MatchMapStringStringIfNonZero(row.Bfd, row1.Bfd) {
		return false
	}
	if !types.MatchMapStringStringIfNonZero(row.BfdStatus, row1.BfdStatus) {
		return false
	}
	if !types.MatchBooleanOptionalIfNonZero(row.CfmFault, row1.CfmFault) {
		return false
	}
	if !types.MatchStringMultiplesIfNonZero(row.CfmFaultStatus, row1.CfmFaultStatus) {
		return false
	}
	if !types.MatchIntegerOptionalIfNonZero(row.CfmFlapCount, row1.CfmFlapCount) {
		return false
	}
	if !types.MatchIntegerOptionalIfNonZero(row.CfmHealth, row1.CfmHealth) {
		return false
	}
	if !types.MatchIntegerOptionalIfNonZero(row.CfmMpid, row1.CfmMpid) {
		return false
	}
	if !types.MatchIntegerMultiplesIfNonZero(row.CfmRemoteMpids, row1.CfmRemoteMpids) {
		return false
	}
	if !types.MatchStringOptionalIfNonZero(row.CfmRemoteOpstate, row1.CfmRemoteOpstate) {
		return false
	}
	if !types.MatchStringOptionalIfNonZero(row.Duplex, row1.Duplex) {
		return false
	}
	if !types.MatchStringOptionalIfNonZero(row.Error, row1.Error) {
		return false
	}
	if !types.MatchMapStringStringIfNonZero(row.ExternalIds, row1.ExternalIds) {
		return false
	}
	if !types.MatchIntegerOptionalIfNonZero(row.Ifindex, row1.Ifindex) {
		return false
	}
	if !types.MatchIntegerIfNonZero(row.IngressPolicingBurst, row1.IngressPolicingBurst) {
		return false
	}
	if !types.MatchIntegerIfNonZero(row.IngressPolicingRate, row1.IngressPolicingRate) {
		return false
	}
	if !types.MatchBooleanOptionalIfNonZero(row.LacpCurrent, row1.LacpCurrent) {
		return false
	}
	if !types.MatchIntegerOptionalIfNonZero(row.LinkResets, row1.LinkResets) {
		return false
	}
	if !types.MatchIntegerOptionalIfNonZero(row.LinkSpeed, row1.LinkSpeed) {
		return false
	}
	if !types.MatchStringOptionalIfNonZero(row.LinkState, row1.LinkState) {
		return false
	}
	if !types.MatchMapStringStringIfNonZero(row.Lldp, row1.Lldp) {
		return false
	}
	if !types.MatchStringOptionalIfNonZero(row.Mac, row1.Mac) {
		return false
	}
	if !types.MatchStringOptionalIfNonZero(row.MacInUse, row1.MacInUse) {
		return false
	}
	if !types.MatchIntegerOptionalIfNonZero(row.Mtu, row1.Mtu) {
		return false
	}
	if !types.MatchIntegerOptionalIfNonZero(row.MtuRequest, row1.MtuRequest) {
		return false
	}
	if !types.MatchStringIfNonZero(row.Name, row1.Name) {
		return false
	}
	if !types.MatchIntegerOptionalIfNonZero(row.Ofport, row1.Ofport) {
		return false
	}
	if !types.MatchIntegerOptionalIfNonZero(row.OfportRequest, row1.OfportRequest) {
		return false
	}
	if !types.MatchMapStringStringIfNonZero(row.Options, row1.Options) {
		return false
	}
	if !types.MatchMapStringStringIfNonZero(row.OtherConfig, row1.OtherConfig) {
		return false
	}
	if !types.MatchMapStringIntegerIfNonZero(row.Statistics, row1.Statistics) {
		return false
	}
	if !types.MatchMapStringStringIfNonZero(row.Status, row1.Status) {
		return false
	}
	if !types.MatchStringIfNonZero(row.Type, row1.Type) {
		return false
	}
	return true
}

func (row *Interface) HasExternalIds() bool {
	return true
}

func (row *Interface) SetExternalId(k, v string) {
	if row.ExternalIds == nil {
		row.ExternalIds = map[string]string{}
	}
	row.ExternalIds[k] = v
}

func (row *Interface) GetExternalId(k string) (string, bool) {
	if row.ExternalIds == nil {
		return "", false
	}
	r, ok := row.ExternalIds[k]
	return r, ok
}

func (row *Interface) RemoveExternalId(k string) (string, bool) {
	if row.ExternalIds == nil {
		return "", false
	}
	r, ok := row.ExternalIds[k]
	if ok {
		delete(row.ExternalIds, k)
	}
	return r, ok
}

type ManagerTable []Manager

var _ types.ITable = &ManagerTable{}

func (tbl ManagerTable) OvsdbTableName() string {
	return "Manager"
}

func (tbl ManagerTable) OvsdbIsRoot() bool {
	return false
}

func (tbl ManagerTable) Rows() []types.IRow {
	r := make([]types.IRow, len(tbl))
	for i := range tbl {
		r[i] = &tbl[i]
	}
	return r
}

func (tbl ManagerTable) NewRow() types.IRow {
	return &Manager{}
}

func (tbl *ManagerTable) AppendRow(irow types.IRow) {
	row := irow.(*Manager)
	*tbl = append(*tbl, *row)
}

func (tbl ManagerTable) OvsdbHasIndex() bool {
	return true
}

func (row *Manager) MatchByTarget(row1 *Manager) bool {
	if !types.MatchString(row.Target, row1.Target) {
		return false
	}
	return true
}

func (tbl ManagerTable) GetByTarget(row1 *Manager) *Manager {
	for i := range tbl {
		row := &tbl[i]
		if row.MatchByTarget(row1) {
			return row
		}
	}
	return nil
}

func (tbl ManagerTable) OvsdbGetByAnyIndex(irow1 types.IRow) types.IRow {
	row1 := irow1.(*Manager)
	if !(types.IsZeroString(row1.Target)) {
		if row := tbl.GetByTarget(row1); row != nil {
			return row
		}
	}
	return nil
}

func (tbl ManagerTable) FindOneMatchNonZeros(row1 *Manager) *Manager {
	for i := range tbl {
		row := &tbl[i]
		if row.MatchNonZeros(row1) {
			return row
		}
	}
	return nil
}

type Manager struct {
	Uuid            string            `json:"_uuid"`
	Version         string            `json:"_version"`
	ConnectionMode  *string           `json:"connection_mode"`
	ExternalIds     map[string]string `json:"external_ids"`
	InactivityProbe *int64            `json:"inactivity_probe"`
	IsConnected     bool              `json:"is_connected"`
	MaxBackoff      *int64            `json:"max_backoff"`
	OtherConfig     map[string]string `json:"other_config"`
	Status          map[string]string `json:"status"`
	Target          string            `json:"target"`
}

var _ types.IRow = &Manager{}

func (row *Manager) OvsdbTableName() string {
	return "Manager"
}

func (row *Manager) OvsdbIsRoot() bool {
	return false
}

func (row *Manager) OvsdbUuid() string {
	return row.Uuid
}

func (row *Manager) OvsdbCmdArgs() []string {
	r := []string{}
	r = append(r, types.OvsdbCmdArgsStringOptional("connection_mode", row.ConnectionMode)...)
	r = append(r, types.OvsdbCmdArgsMapStringString("external_ids", row.ExternalIds)...)
	r = append(r, types.OvsdbCmdArgsIntegerOptional("inactivity_probe", row.InactivityProbe)...)
	r = append(r, types.OvsdbCmdArgsBoolean("is_connected", row.IsConnected)...)
	r = append(r, types.OvsdbCmdArgsIntegerOptional("max_backoff", row.MaxBackoff)...)
	r = append(r, types.OvsdbCmdArgsMapStringString("other_config", row.OtherConfig)...)
	r = append(r, types.OvsdbCmdArgsMapStringString("status", row.Status)...)
	r = append(r, types.OvsdbCmdArgsString("target", row.Target)...)
	return r
}

func (row *Manager) SetColumn(name string, val interface{}) (err error) {
	defer func() {
		if panicErr := recover(); panicErr != nil {
			err = errors.Wrapf(panicErr.(error), "%s: %#v", name, fmt.Sprintf("%#v", val))
		}
	}()
	switch name {
	case "_uuid":
		row.Uuid = types.EnsureUuid(val)
	case "_version":
		row.Version = types.EnsureUuid(val)
	case "connection_mode":
		row.ConnectionMode = types.EnsureStringOptional(val)
	case "external_ids":
		row.ExternalIds = types.EnsureMapStringString(val)
	case "inactivity_probe":
		row.InactivityProbe = types.EnsureIntegerOptional(val)
	case "is_connected":
		row.IsConnected = types.EnsureBoolean(val)
	case "max_backoff":
		row.MaxBackoff = types.EnsureIntegerOptional(val)
	case "other_config":
		row.OtherConfig = types.EnsureMapStringString(val)
	case "status":
		row.Status = types.EnsureMapStringString(val)
	case "target":
		row.Target = types.EnsureString(val)
	default:
		panic(types.ErrUnknownColumn)
	}
	return
}

func (row *Manager) MatchNonZeros(row1 *Manager) bool {
	if !types.MatchUuidIfNonZero(row.Uuid, row1.Uuid) {
		return false
	}
	if !types.MatchUuidIfNonZero(row.Version, row1.Version) {
		return false
	}
	if !types.MatchStringOptionalIfNonZero(row.ConnectionMode, row1.ConnectionMode) {
		return false
	}
	if !types.MatchMapStringStringIfNonZero(row.ExternalIds, row1.ExternalIds) {
		return false
	}
	if !types.MatchIntegerOptionalIfNonZero(row.InactivityProbe, row1.InactivityProbe) {
		return false
	}
	if !types.MatchBooleanIfNonZero(row.IsConnected, row1.IsConnected) {
		return false
	}
	if !types.MatchIntegerOptionalIfNonZero(row.MaxBackoff, row1.MaxBackoff) {
		return false
	}
	if !types.MatchMapStringStringIfNonZero(row.OtherConfig, row1.OtherConfig) {
		return false
	}
	if !types.MatchMapStringStringIfNonZero(row.Status, row1.Status) {
		return false
	}
	if !types.MatchStringIfNonZero(row.Target, row1.Target) {
		return false
	}
	return true
}

func (row *Manager) HasExternalIds() bool {
	return true
}

func (row *Manager) SetExternalId(k, v string) {
	if row.ExternalIds == nil {
		row.ExternalIds = map[string]string{}
	}
	row.ExternalIds[k] = v
}

func (row *Manager) GetExternalId(k string) (string, bool) {
	if row.ExternalIds == nil {
		return "", false
	}
	r, ok := row.ExternalIds[k]
	return r, ok
}

func (row *Manager) RemoveExternalId(k string) (string, bool) {
	if row.ExternalIds == nil {
		return "", false
	}
	r, ok := row.ExternalIds[k]
	if ok {
		delete(row.ExternalIds, k)
	}
	return r, ok
}

type MirrorTable []Mirror

var _ types.ITable = &MirrorTable{}

func (tbl MirrorTable) OvsdbTableName() string {
	return "Mirror"
}

func (tbl MirrorTable) OvsdbIsRoot() bool {
	return false
}

func (tbl MirrorTable) Rows() []types.IRow {
	r := make([]types.IRow, len(tbl))
	for i := range tbl {
		r[i] = &tbl[i]
	}
	return r
}

func (tbl MirrorTable) NewRow() types.IRow {
	return &Mirror{}
}

func (tbl *MirrorTable) AppendRow(irow types.IRow) {
	row := irow.(*Mirror)
	*tbl = append(*tbl, *row)
}

func (tbl MirrorTable) OvsdbHasIndex() bool {
	return false
}

func (tbl MirrorTable) OvsdbGetByAnyIndex(irow1 types.IRow) types.IRow {
	return nil
}

func (tbl MirrorTable) FindOneMatchNonZeros(row1 *Mirror) *Mirror {
	for i := range tbl {
		row := &tbl[i]
		if row.MatchNonZeros(row1) {
			return row
		}
	}
	return nil
}

type Mirror struct {
	Uuid          string            `json:"_uuid"`
	Version       string            `json:"_version"`
	ExternalIds   map[string]string `json:"external_ids"`
	Name          string            `json:"name"`
	OutputPort    *string           `json:"output_port"`
	OutputVlan    *int64            `json:"output_vlan"`
	SelectAll     bool              `json:"select_all"`
	SelectDstPort []string          `json:"select_dst_port"`
	SelectSrcPort []string          `json:"select_src_port"`
	SelectVlan    []int64           `json:"select_vlan"`
	Snaplen       *int64            `json:"snaplen"`
	Statistics    map[string]int64  `json:"statistics"`
}

var _ types.IRow = &Mirror{}

func (row *Mirror) OvsdbTableName() string {
	return "Mirror"
}

func (row *Mirror) OvsdbIsRoot() bool {
	return false
}

func (row *Mirror) OvsdbUuid() string {
	return row.Uuid
}

func (row *Mirror) OvsdbCmdArgs() []string {
	r := []string{}
	r = append(r, types.OvsdbCmdArgsMapStringString("external_ids", row.ExternalIds)...)
	r = append(r, types.OvsdbCmdArgsString("name", row.Name)...)
	r = append(r, types.OvsdbCmdArgsUuidOptional("output_port", row.OutputPort)...)
	r = append(r, types.OvsdbCmdArgsIntegerOptional("output_vlan", row.OutputVlan)...)
	r = append(r, types.OvsdbCmdArgsBoolean("select_all", row.SelectAll)...)
	r = append(r, types.OvsdbCmdArgsUuidMultiples("select_dst_port", row.SelectDstPort)...)
	r = append(r, types.OvsdbCmdArgsUuidMultiples("select_src_port", row.SelectSrcPort)...)
	r = append(r, types.OvsdbCmdArgsIntegerMultiples("select_vlan", row.SelectVlan)...)
	r = append(r, types.OvsdbCmdArgsIntegerOptional("snaplen", row.Snaplen)...)
	r = append(r, types.OvsdbCmdArgsMapStringInteger("statistics", row.Statistics)...)
	return r
}

func (row *Mirror) SetColumn(name string, val interface{}) (err error) {
	defer func() {
		if panicErr := recover(); panicErr != nil {
			err = errors.Wrapf(panicErr.(error), "%s: %#v", name, fmt.Sprintf("%#v", val))
		}
	}()
	switch name {
	case "_uuid":
		row.Uuid = types.EnsureUuid(val)
	case "_version":
		row.Version = types.EnsureUuid(val)
	case "external_ids":
		row.ExternalIds = types.EnsureMapStringString(val)
	case "name":
		row.Name = types.EnsureString(val)
	case "output_port":
		row.OutputPort = types.EnsureUuidOptional(val)
	case "output_vlan":
		row.OutputVlan = types.EnsureIntegerOptional(val)
	case "select_all":
		row.SelectAll = types.EnsureBoolean(val)
	case "select_dst_port":
		row.SelectDstPort = types.EnsureUuidMultiples(val)
	case "select_src_port":
		row.SelectSrcPort = types.EnsureUuidMultiples(val)
	case "select_vlan":
		row.SelectVlan = types.EnsureIntegerMultiples(val)
	case "snaplen":
		row.Snaplen = types.EnsureIntegerOptional(val)
	case "statistics":
		row.Statistics = types.EnsureMapStringInteger(val)
	default:
		panic(types.ErrUnknownColumn)
	}
	return
}

func (row *Mirror) MatchNonZeros(row1 *Mirror) bool {
	if !types.MatchUuidIfNonZero(row.Uuid, row1.Uuid) {
		return false
	}
	if !types.MatchUuidIfNonZero(row.Version, row1.Version) {
		return false
	}
	if !types.MatchMapStringStringIfNonZero(row.ExternalIds, row1.ExternalIds) {
		return false
	}
	if !types.MatchStringIfNonZero(row.Name, row1.Name) {
		return false
	}
	if !types.MatchUuidOptionalIfNonZero(row.OutputPort, row1.OutputPort) {
		return false
	}
	if !types.MatchIntegerOptionalIfNonZero(row.OutputVlan, row1.OutputVlan) {
		return false
	}
	if !types.MatchBooleanIfNonZero(row.SelectAll, row1.SelectAll) {
		return false
	}
	if !types.MatchUuidMultiplesIfNonZero(row.SelectDstPort, row1.SelectDstPort) {
		return false
	}
	if !types.MatchUuidMultiplesIfNonZero(row.SelectSrcPort, row1.SelectSrcPort) {
		return false
	}
	if !types.MatchIntegerMultiplesIfNonZero(row.SelectVlan, row1.SelectVlan) {
		return false
	}
	if !types.MatchIntegerOptionalIfNonZero(row.Snaplen, row1.Snaplen) {
		return false
	}
	if !types.MatchMapStringIntegerIfNonZero(row.Statistics, row1.Statistics) {
		return false
	}
	return true
}

func (row *Mirror) HasExternalIds() bool {
	return true
}

func (row *Mirror) SetExternalId(k, v string) {
	if row.ExternalIds == nil {
		row.ExternalIds = map[string]string{}
	}
	row.ExternalIds[k] = v
}

func (row *Mirror) GetExternalId(k string) (string, bool) {
	if row.ExternalIds == nil {
		return "", false
	}
	r, ok := row.ExternalIds[k]
	return r, ok
}

func (row *Mirror) RemoveExternalId(k string) (string, bool) {
	if row.ExternalIds == nil {
		return "", false
	}
	r, ok := row.ExternalIds[k]
	if ok {
		delete(row.ExternalIds, k)
	}
	return r, ok
}

type NetFlowTable []NetFlow

var _ types.ITable = &NetFlowTable{}

func (tbl NetFlowTable) OvsdbTableName() string {
	return "NetFlow"
}

func (tbl NetFlowTable) OvsdbIsRoot() bool {
	return false
}

func (tbl NetFlowTable) Rows() []types.IRow {
	r := make([]types.IRow, len(tbl))
	for i := range tbl {
		r[i] = &tbl[i]
	}
	return r
}

func (tbl NetFlowTable) NewRow() types.IRow {
	return &NetFlow{}
}

func (tbl *NetFlowTable) AppendRow(irow types.IRow) {
	row := irow.(*NetFlow)
	*tbl = append(*tbl, *row)
}

func (tbl NetFlowTable) OvsdbHasIndex() bool {
	return false
}

func (tbl NetFlowTable) OvsdbGetByAnyIndex(irow1 types.IRow) types.IRow {
	return nil
}

func (tbl NetFlowTable) FindOneMatchNonZeros(row1 *NetFlow) *NetFlow {
	for i := range tbl {
		row := &tbl[i]
		if row.MatchNonZeros(row1) {
			return row
		}
	}
	return nil
}

type NetFlow struct {
	Uuid             string            `json:"_uuid"`
	Version          string            `json:"_version"`
	ActiveTimeout    int64             `json:"active_timeout"`
	AddIdToInterface bool              `json:"add_id_to_interface"`
	EngineId         *int64            `json:"engine_id"`
	EngineType       *int64            `json:"engine_type"`
	ExternalIds      map[string]string `json:"external_ids"`
	Targets          []string          `json:"targets"`
}

var _ types.IRow = &NetFlow{}

func (row *NetFlow) OvsdbTableName() string {
	return "NetFlow"
}

func (row *NetFlow) OvsdbIsRoot() bool {
	return false
}

func (row *NetFlow) OvsdbUuid() string {
	return row.Uuid
}

func (row *NetFlow) OvsdbCmdArgs() []string {
	r := []string{}
	r = append(r, types.OvsdbCmdArgsInteger("active_timeout", row.ActiveTimeout)...)
	r = append(r, types.OvsdbCmdArgsBoolean("add_id_to_interface", row.AddIdToInterface)...)
	r = append(r, types.OvsdbCmdArgsIntegerOptional("engine_id", row.EngineId)...)
	r = append(r, types.OvsdbCmdArgsIntegerOptional("engine_type", row.EngineType)...)
	r = append(r, types.OvsdbCmdArgsMapStringString("external_ids", row.ExternalIds)...)
	r = append(r, types.OvsdbCmdArgsStringMultiples("targets", row.Targets)...)
	return r
}

func (row *NetFlow) SetColumn(name string, val interface{}) (err error) {
	defer func() {
		if panicErr := recover(); panicErr != nil {
			err = errors.Wrapf(panicErr.(error), "%s: %#v", name, fmt.Sprintf("%#v", val))
		}
	}()
	switch name {
	case "_uuid":
		row.Uuid = types.EnsureUuid(val)
	case "_version":
		row.Version = types.EnsureUuid(val)
	case "active_timeout":
		row.ActiveTimeout = types.EnsureInteger(val)
	case "add_id_to_interface":
		row.AddIdToInterface = types.EnsureBoolean(val)
	case "engine_id":
		row.EngineId = types.EnsureIntegerOptional(val)
	case "engine_type":
		row.EngineType = types.EnsureIntegerOptional(val)
	case "external_ids":
		row.ExternalIds = types.EnsureMapStringString(val)
	case "targets":
		row.Targets = types.EnsureStringMultiples(val)
	default:
		panic(types.ErrUnknownColumn)
	}
	return
}

func (row *NetFlow) MatchNonZeros(row1 *NetFlow) bool {
	if !types.MatchUuidIfNonZero(row.Uuid, row1.Uuid) {
		return false
	}
	if !types.MatchUuidIfNonZero(row.Version, row1.Version) {
		return false
	}
	if !types.MatchIntegerIfNonZero(row.ActiveTimeout, row1.ActiveTimeout) {
		return false
	}
	if !types.MatchBooleanIfNonZero(row.AddIdToInterface, row1.AddIdToInterface) {
		return false
	}
	if !types.MatchIntegerOptionalIfNonZero(row.EngineId, row1.EngineId) {
		return false
	}
	if !types.MatchIntegerOptionalIfNonZero(row.EngineType, row1.EngineType) {
		return false
	}
	if !types.MatchMapStringStringIfNonZero(row.ExternalIds, row1.ExternalIds) {
		return false
	}
	if !types.MatchStringMultiplesIfNonZero(row.Targets, row1.Targets) {
		return false
	}
	return true
}

func (row *NetFlow) HasExternalIds() bool {
	return true
}

func (row *NetFlow) SetExternalId(k, v string) {
	if row.ExternalIds == nil {
		row.ExternalIds = map[string]string{}
	}
	row.ExternalIds[k] = v
}

func (row *NetFlow) GetExternalId(k string) (string, bool) {
	if row.ExternalIds == nil {
		return "", false
	}
	r, ok := row.ExternalIds[k]
	return r, ok
}

func (row *NetFlow) RemoveExternalId(k string) (string, bool) {
	if row.ExternalIds == nil {
		return "", false
	}
	r, ok := row.ExternalIds[k]
	if ok {
		delete(row.ExternalIds, k)
	}
	return r, ok
}

type OpenVSwitchTable []OpenVSwitch

var _ types.ITable = &OpenVSwitchTable{}

func (tbl OpenVSwitchTable) OvsdbTableName() string {
	return "Open_vSwitch"
}

func (tbl OpenVSwitchTable) OvsdbIsRoot() bool {
	return true
}

func (tbl OpenVSwitchTable) Rows() []types.IRow {
	r := make([]types.IRow, len(tbl))
	for i := range tbl {
		r[i] = &tbl[i]
	}
	return r
}

func (tbl OpenVSwitchTable) NewRow() types.IRow {
	return &OpenVSwitch{}
}

func (tbl *OpenVSwitchTable) AppendRow(irow types.IRow) {
	row := irow.(*OpenVSwitch)
	*tbl = append(*tbl, *row)
}

func (tbl OpenVSwitchTable) OvsdbHasIndex() bool {
	return false
}

func (tbl OpenVSwitchTable) OvsdbGetByAnyIndex(irow1 types.IRow) types.IRow {
	return nil
}

func (tbl OpenVSwitchTable) FindOneMatchNonZeros(row1 *OpenVSwitch) *OpenVSwitch {
	for i := range tbl {
		row := &tbl[i]
		if row.MatchNonZeros(row1) {
			return row
		}
	}
	return nil
}

type OpenVSwitch struct {
	Uuid           string            `json:"_uuid"`
	Version        string            `json:"_version"`
	Bridges        []string          `json:"bridges"`
	CurCfg         int64             `json:"cur_cfg"`
	DatapathTypes  []string          `json:"datapath_types"`
	DbVersion      *string           `json:"db_version"`
	ExternalIds    map[string]string `json:"external_ids"`
	IfaceTypes     []string          `json:"iface_types"`
	ManagerOptions []string          `json:"manager_options"`
	NextCfg        int64             `json:"next_cfg"`
	OtherConfig    map[string]string `json:"other_config"`
	OvsVersion     *string           `json:"ovs_version"`
	Ssl            *string           `json:"ssl"`
	Statistics     map[string]string `json:"statistics"`
	SystemType     *string           `json:"system_type"`
	SystemVersion  *string           `json:"system_version"`
}

var _ types.IRow = &OpenVSwitch{}

func (row *OpenVSwitch) OvsdbTableName() string {
	return "Open_vSwitch"
}

func (row *OpenVSwitch) OvsdbIsRoot() bool {
	return true
}

func (row *OpenVSwitch) OvsdbUuid() string {
	return row.Uuid
}

func (row *OpenVSwitch) OvsdbCmdArgs() []string {
	r := []string{}
	r = append(r, types.OvsdbCmdArgsUuidMultiples("bridges", row.Bridges)...)
	r = append(r, types.OvsdbCmdArgsInteger("cur_cfg", row.CurCfg)...)
	r = append(r, types.OvsdbCmdArgsStringMultiples("datapath_types", row.DatapathTypes)...)
	r = append(r, types.OvsdbCmdArgsStringOptional("db_version", row.DbVersion)...)
	r = append(r, types.OvsdbCmdArgsMapStringString("external_ids", row.ExternalIds)...)
	r = append(r, types.OvsdbCmdArgsStringMultiples("iface_types", row.IfaceTypes)...)
	r = append(r, types.OvsdbCmdArgsUuidMultiples("manager_options", row.ManagerOptions)...)
	r = append(r, types.OvsdbCmdArgsInteger("next_cfg", row.NextCfg)...)
	r = append(r, types.OvsdbCmdArgsMapStringString("other_config", row.OtherConfig)...)
	r = append(r, types.OvsdbCmdArgsStringOptional("ovs_version", row.OvsVersion)...)
	r = append(r, types.OvsdbCmdArgsUuidOptional("ssl", row.Ssl)...)
	r = append(r, types.OvsdbCmdArgsMapStringString("statistics", row.Statistics)...)
	r = append(r, types.OvsdbCmdArgsStringOptional("system_type", row.SystemType)...)
	r = append(r, types.OvsdbCmdArgsStringOptional("system_version", row.SystemVersion)...)
	return r
}

func (row *OpenVSwitch) SetColumn(name string, val interface{}) (err error) {
	defer func() {
		if panicErr := recover(); panicErr != nil {
			err = errors.Wrapf(panicErr.(error), "%s: %#v", name, fmt.Sprintf("%#v", val))
		}
	}()
	switch name {
	case "_uuid":
		row.Uuid = types.EnsureUuid(val)
	case "_version":
		row.Version = types.EnsureUuid(val)
	case "bridges":
		row.Bridges = types.EnsureUuidMultiples(val)
	case "cur_cfg":
		row.CurCfg = types.EnsureInteger(val)
	case "datapath_types":
		row.DatapathTypes = types.EnsureStringMultiples(val)
	case "db_version":
		row.DbVersion = types.EnsureStringOptional(val)
	case "external_ids":
		row.ExternalIds = types.EnsureMapStringString(val)
	case "iface_types":
		row.IfaceTypes = types.EnsureStringMultiples(val)
	case "manager_options":
		row.ManagerOptions = types.EnsureUuidMultiples(val)
	case "next_cfg":
		row.NextCfg = types.EnsureInteger(val)
	case "other_config":
		row.OtherConfig = types.EnsureMapStringString(val)
	case "ovs_version":
		row.OvsVersion = types.EnsureStringOptional(val)
	case "ssl":
		row.Ssl = types.EnsureUuidOptional(val)
	case "statistics":
		row.Statistics = types.EnsureMapStringString(val)
	case "system_type":
		row.SystemType = types.EnsureStringOptional(val)
	case "system_version":
		row.SystemVersion = types.EnsureStringOptional(val)
	default:
		panic(types.ErrUnknownColumn)
	}
	return
}

func (row *OpenVSwitch) MatchNonZeros(row1 *OpenVSwitch) bool {
	if !types.MatchUuidIfNonZero(row.Uuid, row1.Uuid) {
		return false
	}
	if !types.MatchUuidIfNonZero(row.Version, row1.Version) {
		return false
	}
	if !types.MatchUuidMultiplesIfNonZero(row.Bridges, row1.Bridges) {
		return false
	}
	if !types.MatchIntegerIfNonZero(row.CurCfg, row1.CurCfg) {
		return false
	}
	if !types.MatchStringMultiplesIfNonZero(row.DatapathTypes, row1.DatapathTypes) {
		return false
	}
	if !types.MatchStringOptionalIfNonZero(row.DbVersion, row1.DbVersion) {
		return false
	}
	if !types.MatchMapStringStringIfNonZero(row.ExternalIds, row1.ExternalIds) {
		return false
	}
	if !types.MatchStringMultiplesIfNonZero(row.IfaceTypes, row1.IfaceTypes) {
		return false
	}
	if !types.MatchUuidMultiplesIfNonZero(row.ManagerOptions, row1.ManagerOptions) {
		return false
	}
	if !types.MatchIntegerIfNonZero(row.NextCfg, row1.NextCfg) {
		return false
	}
	if !types.MatchMapStringStringIfNonZero(row.OtherConfig, row1.OtherConfig) {
		return false
	}
	if !types.MatchStringOptionalIfNonZero(row.OvsVersion, row1.OvsVersion) {
		return false
	}
	if !types.MatchUuidOptionalIfNonZero(row.Ssl, row1.Ssl) {
		return false
	}
	if !types.MatchMapStringStringIfNonZero(row.Statistics, row1.Statistics) {
		return false
	}
	if !types.MatchStringOptionalIfNonZero(row.SystemType, row1.SystemType) {
		return false
	}
	if !types.MatchStringOptionalIfNonZero(row.SystemVersion, row1.SystemVersion) {
		return false
	}
	return true
}

func (row *OpenVSwitch) HasExternalIds() bool {
	return true
}

func (row *OpenVSwitch) SetExternalId(k, v string) {
	if row.ExternalIds == nil {
		row.ExternalIds = map[string]string{}
	}
	row.ExternalIds[k] = v
}

func (row *OpenVSwitch) GetExternalId(k string) (string, bool) {
	if row.ExternalIds == nil {
		return "", false
	}
	r, ok := row.ExternalIds[k]
	return r, ok
}

func (row *OpenVSwitch) RemoveExternalId(k string) (string, bool) {
	if row.ExternalIds == nil {
		return "", false
	}
	r, ok := row.ExternalIds[k]
	if ok {
		delete(row.ExternalIds, k)
	}
	return r, ok
}

type PortTable []Port

var _ types.ITable = &PortTable{}

func (tbl PortTable) OvsdbTableName() string {
	return "Port"
}

func (tbl PortTable) OvsdbIsRoot() bool {
	return false
}

func (tbl PortTable) Rows() []types.IRow {
	r := make([]types.IRow, len(tbl))
	for i := range tbl {
		r[i] = &tbl[i]
	}
	return r
}

func (tbl PortTable) NewRow() types.IRow {
	return &Port{}
}

func (tbl *PortTable) AppendRow(irow types.IRow) {
	row := irow.(*Port)
	*tbl = append(*tbl, *row)
}

func (tbl PortTable) OvsdbHasIndex() bool {
	return true
}

func (row *Port) MatchByName(row1 *Port) bool {
	if !types.MatchString(row.Name, row1.Name) {
		return false
	}
	return true
}

func (tbl PortTable) GetByName(row1 *Port) *Port {
	for i := range tbl {
		row := &tbl[i]
		if row.MatchByName(row1) {
			return row
		}
	}
	return nil
}

func (tbl PortTable) OvsdbGetByAnyIndex(irow1 types.IRow) types.IRow {
	row1 := irow1.(*Port)
	if !(types.IsZeroString(row1.Name)) {
		if row := tbl.GetByName(row1); row != nil {
			return row
		}
	}
	return nil
}

func (tbl PortTable) FindOneMatchNonZeros(row1 *Port) *Port {
	for i := range tbl {
		row := &tbl[i]
		if row.MatchNonZeros(row1) {
			return row
		}
	}
	return nil
}

type Port struct {
	Uuid            string            `json:"_uuid"`
	Version         string            `json:"_version"`
	BondActiveSlave *string           `json:"bond_active_slave"`
	BondDowndelay   int64             `json:"bond_downdelay"`
	BondFakeIface   bool              `json:"bond_fake_iface"`
	BondMode        *string           `json:"bond_mode"`
	BondUpdelay     int64             `json:"bond_updelay"`
	Cvlans          []int64           `json:"cvlans"`
	ExternalIds     map[string]string `json:"external_ids"`
	FakeBridge      bool              `json:"fake_bridge"`
	Interfaces      []string          `json:"interfaces"`
	Lacp            *string           `json:"lacp"`
	Mac             *string           `json:"mac"`
	Name            string            `json:"name"`
	OtherConfig     map[string]string `json:"other_config"`
	Protected       bool              `json:"protected"`
	Qos             *string           `json:"qos"`
	RstpStatistics  map[string]int64  `json:"rstp_statistics"`
	RstpStatus      map[string]string `json:"rstp_status"`
	Statistics      map[string]int64  `json:"statistics"`
	Status          map[string]string `json:"status"`
	Tag             *int64            `json:"tag"`
	Trunks          []int64           `json:"trunks"`
	VlanMode        *string           `json:"vlan_mode"`
}

var _ types.IRow = &Port{}

func (row *Port) OvsdbTableName() string {
	return "Port"
}

func (row *Port) OvsdbIsRoot() bool {
	return false
}

func (row *Port) OvsdbUuid() string {
	return row.Uuid
}

func (row *Port) OvsdbCmdArgs() []string {
	r := []string{}
	r = append(r, types.OvsdbCmdArgsStringOptional("bond_active_slave", row.BondActiveSlave)...)
	r = append(r, types.OvsdbCmdArgsInteger("bond_downdelay", row.BondDowndelay)...)
	r = append(r, types.OvsdbCmdArgsBoolean("bond_fake_iface", row.BondFakeIface)...)
	r = append(r, types.OvsdbCmdArgsStringOptional("bond_mode", row.BondMode)...)
	r = append(r, types.OvsdbCmdArgsInteger("bond_updelay", row.BondUpdelay)...)
	r = append(r, types.OvsdbCmdArgsIntegerMultiples("cvlans", row.Cvlans)...)
	r = append(r, types.OvsdbCmdArgsMapStringString("external_ids", row.ExternalIds)...)
	r = append(r, types.OvsdbCmdArgsBoolean("fake_bridge", row.FakeBridge)...)
	r = append(r, types.OvsdbCmdArgsUuidMultiples("interfaces", row.Interfaces)...)
	r = append(r, types.OvsdbCmdArgsStringOptional("lacp", row.Lacp)...)
	r = append(r, types.OvsdbCmdArgsStringOptional("mac", row.Mac)...)
	r = append(r, types.OvsdbCmdArgsString("name", row.Name)...)
	r = append(r, types.OvsdbCmdArgsMapStringString("other_config", row.OtherConfig)...)
	r = append(r, types.OvsdbCmdArgsBoolean("protected", row.Protected)...)
	r = append(r, types.OvsdbCmdArgsUuidOptional("qos", row.Qos)...)
	r = append(r, types.OvsdbCmdArgsMapStringInteger("rstp_statistics", row.RstpStatistics)...)
	r = append(r, types.OvsdbCmdArgsMapStringString("rstp_status", row.RstpStatus)...)
	r = append(r, types.OvsdbCmdArgsMapStringInteger("statistics", row.Statistics)...)
	r = append(r, types.OvsdbCmdArgsMapStringString("status", row.Status)...)
	r = append(r, types.OvsdbCmdArgsIntegerOptional("tag", row.Tag)...)
	r = append(r, types.OvsdbCmdArgsIntegerMultiples("trunks", row.Trunks)...)
	r = append(r, types.OvsdbCmdArgsStringOptional("vlan_mode", row.VlanMode)...)
	return r
}

func (row *Port) SetColumn(name string, val interface{}) (err error) {
	defer func() {
		if panicErr := recover(); panicErr != nil {
			err = errors.Wrapf(panicErr.(error), "%s: %#v", name, fmt.Sprintf("%#v", val))
		}
	}()
	switch name {
	case "_uuid":
		row.Uuid = types.EnsureUuid(val)
	case "_version":
		row.Version = types.EnsureUuid(val)
	case "bond_active_slave":
		row.BondActiveSlave = types.EnsureStringOptional(val)
	case "bond_downdelay":
		row.BondDowndelay = types.EnsureInteger(val)
	case "bond_fake_iface":
		row.BondFakeIface = types.EnsureBoolean(val)
	case "bond_mode":
		row.BondMode = types.EnsureStringOptional(val)
	case "bond_updelay":
		row.BondUpdelay = types.EnsureInteger(val)
	case "cvlans":
		row.Cvlans = types.EnsureIntegerMultiples(val)
	case "external_ids":
		row.ExternalIds = types.EnsureMapStringString(val)
	case "fake_bridge":
		row.FakeBridge = types.EnsureBoolean(val)
	case "interfaces":
		row.Interfaces = types.EnsureUuidMultiples(val)
	case "lacp":
		row.Lacp = types.EnsureStringOptional(val)
	case "mac":
		row.Mac = types.EnsureStringOptional(val)
	case "name":
		row.Name = types.EnsureString(val)
	case "other_config":
		row.OtherConfig = types.EnsureMapStringString(val)
	case "protected":
		row.Protected = types.EnsureBoolean(val)
	case "qos":
		row.Qos = types.EnsureUuidOptional(val)
	case "rstp_statistics":
		row.RstpStatistics = types.EnsureMapStringInteger(val)
	case "rstp_status":
		row.RstpStatus = types.EnsureMapStringString(val)
	case "statistics":
		row.Statistics = types.EnsureMapStringInteger(val)
	case "status":
		row.Status = types.EnsureMapStringString(val)
	case "tag":
		row.Tag = types.EnsureIntegerOptional(val)
	case "trunks":
		row.Trunks = types.EnsureIntegerMultiples(val)
	case "vlan_mode":
		row.VlanMode = types.EnsureStringOptional(val)
	default:
		panic(types.ErrUnknownColumn)
	}
	return
}

func (row *Port) MatchNonZeros(row1 *Port) bool {
	if !types.MatchUuidIfNonZero(row.Uuid, row1.Uuid) {
		return false
	}
	if !types.MatchUuidIfNonZero(row.Version, row1.Version) {
		return false
	}
	if !types.MatchStringOptionalIfNonZero(row.BondActiveSlave, row1.BondActiveSlave) {
		return false
	}
	if !types.MatchIntegerIfNonZero(row.BondDowndelay, row1.BondDowndelay) {
		return false
	}
	if !types.MatchBooleanIfNonZero(row.BondFakeIface, row1.BondFakeIface) {
		return false
	}
	if !types.MatchStringOptionalIfNonZero(row.BondMode, row1.BondMode) {
		return false
	}
	if !types.MatchIntegerIfNonZero(row.BondUpdelay, row1.BondUpdelay) {
		return false
	}
	if !types.MatchIntegerMultiplesIfNonZero(row.Cvlans, row1.Cvlans) {
		return false
	}
	if !types.MatchMapStringStringIfNonZero(row.ExternalIds, row1.ExternalIds) {
		return false
	}
	if !types.MatchBooleanIfNonZero(row.FakeBridge, row1.FakeBridge) {
		return false
	}
	if !types.MatchUuidMultiplesIfNonZero(row.Interfaces, row1.Interfaces) {
		return false
	}
	if !types.MatchStringOptionalIfNonZero(row.Lacp, row1.Lacp) {
		return false
	}
	if !types.MatchStringOptionalIfNonZero(row.Mac, row1.Mac) {
		return false
	}
	if !types.MatchStringIfNonZero(row.Name, row1.Name) {
		return false
	}
	if !types.MatchMapStringStringIfNonZero(row.OtherConfig, row1.OtherConfig) {
		return false
	}
	if !types.MatchBooleanIfNonZero(row.Protected, row1.Protected) {
		return false
	}
	if !types.MatchUuidOptionalIfNonZero(row.Qos, row1.Qos) {
		return false
	}
	if !types.MatchMapStringIntegerIfNonZero(row.RstpStatistics, row1.RstpStatistics) {
		return false
	}
	if !types.MatchMapStringStringIfNonZero(row.RstpStatus, row1.RstpStatus) {
		return false
	}
	if !types.MatchMapStringIntegerIfNonZero(row.Statistics, row1.Statistics) {
		return false
	}
	if !types.MatchMapStringStringIfNonZero(row.Status, row1.Status) {
		return false
	}
	if !types.MatchIntegerOptionalIfNonZero(row.Tag, row1.Tag) {
		return false
	}
	if !types.MatchIntegerMultiplesIfNonZero(row.Trunks, row1.Trunks) {
		return false
	}
	if !types.MatchStringOptionalIfNonZero(row.VlanMode, row1.VlanMode) {
		return false
	}
	return true
}

func (row *Port) HasExternalIds() bool {
	return true
}

func (row *Port) SetExternalId(k, v string) {
	if row.ExternalIds == nil {
		row.ExternalIds = map[string]string{}
	}
	row.ExternalIds[k] = v
}

func (row *Port) GetExternalId(k string) (string, bool) {
	if row.ExternalIds == nil {
		return "", false
	}
	r, ok := row.ExternalIds[k]
	return r, ok
}

func (row *Port) RemoveExternalId(k string) (string, bool) {
	if row.ExternalIds == nil {
		return "", false
	}
	r, ok := row.ExternalIds[k]
	if ok {
		delete(row.ExternalIds, k)
	}
	return r, ok
}

type QoSTable []QoS

var _ types.ITable = &QoSTable{}

func (tbl QoSTable) OvsdbTableName() string {
	return "QoS"
}

func (tbl QoSTable) OvsdbIsRoot() bool {
	return true
}

func (tbl QoSTable) Rows() []types.IRow {
	r := make([]types.IRow, len(tbl))
	for i := range tbl {
		r[i] = &tbl[i]
	}
	return r
}

func (tbl QoSTable) NewRow() types.IRow {
	return &QoS{}
}

func (tbl *QoSTable) AppendRow(irow types.IRow) {
	row := irow.(*QoS)
	*tbl = append(*tbl, *row)
}

func (tbl QoSTable) OvsdbHasIndex() bool {
	return false
}

func (tbl QoSTable) OvsdbGetByAnyIndex(irow1 types.IRow) types.IRow {
	return nil
}

func (tbl QoSTable) FindOneMatchNonZeros(row1 *QoS) *QoS {
	for i := range tbl {
		row := &tbl[i]
		if row.MatchNonZeros(row1) {
			return row
		}
	}
	return nil
}

type QoS struct {
	Uuid        string            `json:"_uuid"`
	Version     string            `json:"_version"`
	ExternalIds map[string]string `json:"external_ids"`
	OtherConfig map[string]string `json:"other_config"`
	Queues      map[int64]string  `json:"queues"`
	Type        string            `json:"type"`
}

var _ types.IRow = &QoS{}

func (row *QoS) OvsdbTableName() string {
	return "QoS"
}

func (row *QoS) OvsdbIsRoot() bool {
	return true
}

func (row *QoS) OvsdbUuid() string {
	return row.Uuid
}

func (row *QoS) OvsdbCmdArgs() []string {
	r := []string{}
	r = append(r, types.OvsdbCmdArgsMapStringString("external_ids", row.ExternalIds)...)
	r = append(r, types.OvsdbCmdArgsMapStringString("other_config", row.OtherConfig)...)
	r = append(r, types.OvsdbCmdArgsMapIntegerUuid("queues", row.Queues)...)
	r = append(r, types.OvsdbCmdArgsString("type", row.Type)...)
	return r
}

func (row *QoS) SetColumn(name string, val interface{}) (err error) {
	defer func() {
		if panicErr := recover(); panicErr != nil {
			err = errors.Wrapf(panicErr.(error), "%s: %#v", name, fmt.Sprintf("%#v", val))
		}
	}()
	switch name {
	case "_uuid":
		row.Uuid = types.EnsureUuid(val)
	case "_version":
		row.Version = types.EnsureUuid(val)
	case "external_ids":
		row.ExternalIds = types.EnsureMapStringString(val)
	case "other_config":
		row.OtherConfig = types.EnsureMapStringString(val)
	case "queues":
		row.Queues = types.EnsureMapIntegerUuid(val)
	case "type":
		row.Type = types.EnsureString(val)
	default:
		panic(types.ErrUnknownColumn)
	}
	return
}

func (row *QoS) MatchNonZeros(row1 *QoS) bool {
	if !types.MatchUuidIfNonZero(row.Uuid, row1.Uuid) {
		return false
	}
	if !types.MatchUuidIfNonZero(row.Version, row1.Version) {
		return false
	}
	if !types.MatchMapStringStringIfNonZero(row.ExternalIds, row1.ExternalIds) {
		return false
	}
	if !types.MatchMapStringStringIfNonZero(row.OtherConfig, row1.OtherConfig) {
		return false
	}
	if !types.MatchMapIntegerUuidIfNonZero(row.Queues, row1.Queues) {
		return false
	}
	if !types.MatchStringIfNonZero(row.Type, row1.Type) {
		return false
	}
	return true
}

func (row *QoS) HasExternalIds() bool {
	return true
}

func (row *QoS) SetExternalId(k, v string) {
	if row.ExternalIds == nil {
		row.ExternalIds = map[string]string{}
	}
	row.ExternalIds[k] = v
}

func (row *QoS) GetExternalId(k string) (string, bool) {
	if row.ExternalIds == nil {
		return "", false
	}
	r, ok := row.ExternalIds[k]
	return r, ok
}

func (row *QoS) RemoveExternalId(k string) (string, bool) {
	if row.ExternalIds == nil {
		return "", false
	}
	r, ok := row.ExternalIds[k]
	if ok {
		delete(row.ExternalIds, k)
	}
	return r, ok
}

type QueueTable []Queue

var _ types.ITable = &QueueTable{}

func (tbl QueueTable) OvsdbTableName() string {
	return "Queue"
}

func (tbl QueueTable) OvsdbIsRoot() bool {
	return true
}

func (tbl QueueTable) Rows() []types.IRow {
	r := make([]types.IRow, len(tbl))
	for i := range tbl {
		r[i] = &tbl[i]
	}
	return r
}

func (tbl QueueTable) NewRow() types.IRow {
	return &Queue{}
}

func (tbl *QueueTable) AppendRow(irow types.IRow) {
	row := irow.(*Queue)
	*tbl = append(*tbl, *row)
}

func (tbl QueueTable) OvsdbHasIndex() bool {
	return false
}

func (tbl QueueTable) OvsdbGetByAnyIndex(irow1 types.IRow) types.IRow {
	return nil
}

func (tbl QueueTable) FindOneMatchNonZeros(row1 *Queue) *Queue {
	for i := range tbl {
		row := &tbl[i]
		if row.MatchNonZeros(row1) {
			return row
		}
	}
	return nil
}

type Queue struct {
	Uuid        string            `json:"_uuid"`
	Version     string            `json:"_version"`
	Dscp        *int64            `json:"dscp"`
	ExternalIds map[string]string `json:"external_ids"`
	OtherConfig map[string]string `json:"other_config"`
}

var _ types.IRow = &Queue{}

func (row *Queue) OvsdbTableName() string {
	return "Queue"
}

func (row *Queue) OvsdbIsRoot() bool {
	return true
}

func (row *Queue) OvsdbUuid() string {
	return row.Uuid
}

func (row *Queue) OvsdbCmdArgs() []string {
	r := []string{}
	r = append(r, types.OvsdbCmdArgsIntegerOptional("dscp", row.Dscp)...)
	r = append(r, types.OvsdbCmdArgsMapStringString("external_ids", row.ExternalIds)...)
	r = append(r, types.OvsdbCmdArgsMapStringString("other_config", row.OtherConfig)...)
	return r
}

func (row *Queue) SetColumn(name string, val interface{}) (err error) {
	defer func() {
		if panicErr := recover(); panicErr != nil {
			err = errors.Wrapf(panicErr.(error), "%s: %#v", name, fmt.Sprintf("%#v", val))
		}
	}()
	switch name {
	case "_uuid":
		row.Uuid = types.EnsureUuid(val)
	case "_version":
		row.Version = types.EnsureUuid(val)
	case "dscp":
		row.Dscp = types.EnsureIntegerOptional(val)
	case "external_ids":
		row.ExternalIds = types.EnsureMapStringString(val)
	case "other_config":
		row.OtherConfig = types.EnsureMapStringString(val)
	default:
		panic(types.ErrUnknownColumn)
	}
	return
}

func (row *Queue) MatchNonZeros(row1 *Queue) bool {
	if !types.MatchUuidIfNonZero(row.Uuid, row1.Uuid) {
		return false
	}
	if !types.MatchUuidIfNonZero(row.Version, row1.Version) {
		return false
	}
	if !types.MatchIntegerOptionalIfNonZero(row.Dscp, row1.Dscp) {
		return false
	}
	if !types.MatchMapStringStringIfNonZero(row.ExternalIds, row1.ExternalIds) {
		return false
	}
	if !types.MatchMapStringStringIfNonZero(row.OtherConfig, row1.OtherConfig) {
		return false
	}
	return true
}

func (row *Queue) HasExternalIds() bool {
	return true
}

func (row *Queue) SetExternalId(k, v string) {
	if row.ExternalIds == nil {
		row.ExternalIds = map[string]string{}
	}
	row.ExternalIds[k] = v
}

func (row *Queue) GetExternalId(k string) (string, bool) {
	if row.ExternalIds == nil {
		return "", false
	}
	r, ok := row.ExternalIds[k]
	return r, ok
}

func (row *Queue) RemoveExternalId(k string) (string, bool) {
	if row.ExternalIds == nil {
		return "", false
	}
	r, ok := row.ExternalIds[k]
	if ok {
		delete(row.ExternalIds, k)
	}
	return r, ok
}

type SSLTable []SSL

var _ types.ITable = &SSLTable{}

func (tbl SSLTable) OvsdbTableName() string {
	return "SSL"
}

func (tbl SSLTable) OvsdbIsRoot() bool {
	return false
}

func (tbl SSLTable) Rows() []types.IRow {
	r := make([]types.IRow, len(tbl))
	for i := range tbl {
		r[i] = &tbl[i]
	}
	return r
}

func (tbl SSLTable) NewRow() types.IRow {
	return &SSL{}
}

func (tbl *SSLTable) AppendRow(irow types.IRow) {
	row := irow.(*SSL)
	*tbl = append(*tbl, *row)
}

func (tbl SSLTable) OvsdbHasIndex() bool {
	return false
}

func (tbl SSLTable) OvsdbGetByAnyIndex(irow1 types.IRow) types.IRow {
	return nil
}

func (tbl SSLTable) FindOneMatchNonZeros(row1 *SSL) *SSL {
	for i := range tbl {
		row := &tbl[i]
		if row.MatchNonZeros(row1) {
			return row
		}
	}
	return nil
}

type SSL struct {
	Uuid            string            `json:"_uuid"`
	Version         string            `json:"_version"`
	BootstrapCaCert bool              `json:"bootstrap_ca_cert"`
	CaCert          string            `json:"ca_cert"`
	Certificate     string            `json:"certificate"`
	ExternalIds     map[string]string `json:"external_ids"`
	PrivateKey      string            `json:"private_key"`
}

var _ types.IRow = &SSL{}

func (row *SSL) OvsdbTableName() string {
	return "SSL"
}

func (row *SSL) OvsdbIsRoot() bool {
	return false
}

func (row *SSL) OvsdbUuid() string {
	return row.Uuid
}

func (row *SSL) OvsdbCmdArgs() []string {
	r := []string{}
	r = append(r, types.OvsdbCmdArgsBoolean("bootstrap_ca_cert", row.BootstrapCaCert)...)
	r = append(r, types.OvsdbCmdArgsString("ca_cert", row.CaCert)...)
	r = append(r, types.OvsdbCmdArgsString("certificate", row.Certificate)...)
	r = append(r, types.OvsdbCmdArgsMapStringString("external_ids", row.ExternalIds)...)
	r = append(r, types.OvsdbCmdArgsString("private_key", row.PrivateKey)...)
	return r
}

func (row *SSL) SetColumn(name string, val interface{}) (err error) {
	defer func() {
		if panicErr := recover(); panicErr != nil {
			err = errors.Wrapf(panicErr.(error), "%s: %#v", name, fmt.Sprintf("%#v", val))
		}
	}()
	switch name {
	case "_uuid":
		row.Uuid = types.EnsureUuid(val)
	case "_version":
		row.Version = types.EnsureUuid(val)
	case "bootstrap_ca_cert":
		row.BootstrapCaCert = types.EnsureBoolean(val)
	case "ca_cert":
		row.CaCert = types.EnsureString(val)
	case "certificate":
		row.Certificate = types.EnsureString(val)
	case "external_ids":
		row.ExternalIds = types.EnsureMapStringString(val)
	case "private_key":
		row.PrivateKey = types.EnsureString(val)
	default:
		panic(types.ErrUnknownColumn)
	}
	return
}

func (row *SSL) MatchNonZeros(row1 *SSL) bool {
	if !types.MatchUuidIfNonZero(row.Uuid, row1.Uuid) {
		return false
	}
	if !types.MatchUuidIfNonZero(row.Version, row1.Version) {
		return false
	}
	if !types.MatchBooleanIfNonZero(row.BootstrapCaCert, row1.BootstrapCaCert) {
		return false
	}
	if !types.MatchStringIfNonZero(row.CaCert, row1.CaCert) {
		return false
	}
	if !types.MatchStringIfNonZero(row.Certificate, row1.Certificate) {
		return false
	}
	if !types.MatchMapStringStringIfNonZero(row.ExternalIds, row1.ExternalIds) {
		return false
	}
	if !types.MatchStringIfNonZero(row.PrivateKey, row1.PrivateKey) {
		return false
	}
	return true
}

func (row *SSL) HasExternalIds() bool {
	return true
}

func (row *SSL) SetExternalId(k, v string) {
	if row.ExternalIds == nil {
		row.ExternalIds = map[string]string{}
	}
	row.ExternalIds[k] = v
}

func (row *SSL) GetExternalId(k string) (string, bool) {
	if row.ExternalIds == nil {
		return "", false
	}
	r, ok := row.ExternalIds[k]
	return r, ok
}

func (row *SSL) RemoveExternalId(k string) (string, bool) {
	if row.ExternalIds == nil {
		return "", false
	}
	r, ok := row.ExternalIds[k]
	if ok {
		delete(row.ExternalIds, k)
	}
	return r, ok
}

type SFlowTable []SFlow

var _ types.ITable = &SFlowTable{}

func (tbl SFlowTable) OvsdbTableName() string {
	return "sFlow"
}

func (tbl SFlowTable) OvsdbIsRoot() bool {
	return false
}

func (tbl SFlowTable) Rows() []types.IRow {
	r := make([]types.IRow, len(tbl))
	for i := range tbl {
		r[i] = &tbl[i]
	}
	return r
}

func (tbl SFlowTable) NewRow() types.IRow {
	return &SFlow{}
}

func (tbl *SFlowTable) AppendRow(irow types.IRow) {
	row := irow.(*SFlow)
	*tbl = append(*tbl, *row)
}

func (tbl SFlowTable) OvsdbHasIndex() bool {
	return false
}

func (tbl SFlowTable) OvsdbGetByAnyIndex(irow1 types.IRow) types.IRow {
	return nil
}

func (tbl SFlowTable) FindOneMatchNonZeros(row1 *SFlow) *SFlow {
	for i := range tbl {
		row := &tbl[i]
		if row.MatchNonZeros(row1) {
			return row
		}
	}
	return nil
}

type SFlow struct {
	Uuid        string            `json:"_uuid"`
	Version     string            `json:"_version"`
	Agent       *string           `json:"agent"`
	ExternalIds map[string]string `json:"external_ids"`
	Header      *int64            `json:"header"`
	Polling     *int64            `json:"polling"`
	Sampling    *int64            `json:"sampling"`
	Targets     []string          `json:"targets"`
}

var _ types.IRow = &SFlow{}

func (row *SFlow) OvsdbTableName() string {
	return "sFlow"
}

func (row *SFlow) OvsdbIsRoot() bool {
	return false
}

func (row *SFlow) OvsdbUuid() string {
	return row.Uuid
}

func (row *SFlow) OvsdbCmdArgs() []string {
	r := []string{}
	r = append(r, types.OvsdbCmdArgsStringOptional("agent", row.Agent)...)
	r = append(r, types.OvsdbCmdArgsMapStringString("external_ids", row.ExternalIds)...)
	r = append(r, types.OvsdbCmdArgsIntegerOptional("header", row.Header)...)
	r = append(r, types.OvsdbCmdArgsIntegerOptional("polling", row.Polling)...)
	r = append(r, types.OvsdbCmdArgsIntegerOptional("sampling", row.Sampling)...)
	r = append(r, types.OvsdbCmdArgsStringMultiples("targets", row.Targets)...)
	return r
}

func (row *SFlow) SetColumn(name string, val interface{}) (err error) {
	defer func() {
		if panicErr := recover(); panicErr != nil {
			err = errors.Wrapf(panicErr.(error), "%s: %#v", name, fmt.Sprintf("%#v", val))
		}
	}()
	switch name {
	case "_uuid":
		row.Uuid = types.EnsureUuid(val)
	case "_version":
		row.Version = types.EnsureUuid(val)
	case "agent":
		row.Agent = types.EnsureStringOptional(val)
	case "external_ids":
		row.ExternalIds = types.EnsureMapStringString(val)
	case "header":
		row.Header = types.EnsureIntegerOptional(val)
	case "polling":
		row.Polling = types.EnsureIntegerOptional(val)
	case "sampling":
		row.Sampling = types.EnsureIntegerOptional(val)
	case "targets":
		row.Targets = types.EnsureStringMultiples(val)
	default:
		panic(types.ErrUnknownColumn)
	}
	return
}

func (row *SFlow) MatchNonZeros(row1 *SFlow) bool {
	if !types.MatchUuidIfNonZero(row.Uuid, row1.Uuid) {
		return false
	}
	if !types.MatchUuidIfNonZero(row.Version, row1.Version) {
		return false
	}
	if !types.MatchStringOptionalIfNonZero(row.Agent, row1.Agent) {
		return false
	}
	if !types.MatchMapStringStringIfNonZero(row.ExternalIds, row1.ExternalIds) {
		return false
	}
	if !types.MatchIntegerOptionalIfNonZero(row.Header, row1.Header) {
		return false
	}
	if !types.MatchIntegerOptionalIfNonZero(row.Polling, row1.Polling) {
		return false
	}
	if !types.MatchIntegerOptionalIfNonZero(row.Sampling, row1.Sampling) {
		return false
	}
	if !types.MatchStringMultiplesIfNonZero(row.Targets, row1.Targets) {
		return false
	}
	return true
}

func (row *SFlow) HasExternalIds() bool {
	return true
}

func (row *SFlow) SetExternalId(k, v string) {
	if row.ExternalIds == nil {
		row.ExternalIds = map[string]string{}
	}
	row.ExternalIds[k] = v
}

func (row *SFlow) GetExternalId(k string) (string, bool) {
	if row.ExternalIds == nil {
		return "", false
	}
	r, ok := row.ExternalIds[k]
	return r, ok
}

func (row *SFlow) RemoveExternalId(k string) (string, bool) {
	if row.ExternalIds == nil {
		return "", false
	}
	r, ok := row.ExternalIds[k]
	if ok {
		delete(row.ExternalIds, k)
	}
	return r, ok
}
