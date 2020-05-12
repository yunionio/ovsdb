// DO NOT EDIT: automatically generated code

package ovn_sb

import (
	"fmt"

	"github.com/pkg/errors"

	"yunion.io/x/ovsdb/types"
)

type OVNSouthbound struct {
	AddressSet      AddressSetTable
	Chassis         ChassisTable
	Connection      ConnectionTable
	DHCPOptions     DHCPOptionsTable
	DHCPv6Options   DHCPv6OptionsTable
	DNS             DNSTable
	DatapathBinding DatapathBindingTable
	Encap           EncapTable
	GatewayChassis  GatewayChassisTable
	LogicalFlow     LogicalFlowTable
	MACBinding      MACBindingTable
	MulticastGroup  MulticastGroupTable
	PortBinding     PortBindingTable
	RBACPermission  RBACPermissionTable
	RBACRole        RBACRoleTable
	SBGlobal        SBGlobalTable
	SSL             SSLTable
}

func (db OVNSouthbound) FindOneMatchNonZeros(irow types.IRow) types.IRow {
	switch row := irow.(type) {
	case *AddressSet:
		if r := db.AddressSet.FindOneMatchNonZeros(row); r != nil {
			return r
		}
		return nil
	case *Chassis:
		if r := db.Chassis.FindOneMatchNonZeros(row); r != nil {
			return r
		}
		return nil
	case *Connection:
		if r := db.Connection.FindOneMatchNonZeros(row); r != nil {
			return r
		}
		return nil
	case *DHCPOptions:
		if r := db.DHCPOptions.FindOneMatchNonZeros(row); r != nil {
			return r
		}
		return nil
	case *DHCPv6Options:
		if r := db.DHCPv6Options.FindOneMatchNonZeros(row); r != nil {
			return r
		}
		return nil
	case *DNS:
		if r := db.DNS.FindOneMatchNonZeros(row); r != nil {
			return r
		}
		return nil
	case *DatapathBinding:
		if r := db.DatapathBinding.FindOneMatchNonZeros(row); r != nil {
			return r
		}
		return nil
	case *Encap:
		if r := db.Encap.FindOneMatchNonZeros(row); r != nil {
			return r
		}
		return nil
	case *GatewayChassis:
		if r := db.GatewayChassis.FindOneMatchNonZeros(row); r != nil {
			return r
		}
		return nil
	case *LogicalFlow:
		if r := db.LogicalFlow.FindOneMatchNonZeros(row); r != nil {
			return r
		}
		return nil
	case *MACBinding:
		if r := db.MACBinding.FindOneMatchNonZeros(row); r != nil {
			return r
		}
		return nil
	case *MulticastGroup:
		if r := db.MulticastGroup.FindOneMatchNonZeros(row); r != nil {
			return r
		}
		return nil
	case *PortBinding:
		if r := db.PortBinding.FindOneMatchNonZeros(row); r != nil {
			return r
		}
		return nil
	case *RBACPermission:
		if r := db.RBACPermission.FindOneMatchNonZeros(row); r != nil {
			return r
		}
		return nil
	case *RBACRole:
		if r := db.RBACRole.FindOneMatchNonZeros(row); r != nil {
			return r
		}
		return nil
	case *SBGlobal:
		if r := db.SBGlobal.FindOneMatchNonZeros(row); r != nil {
			return r
		}
		return nil
	case *SSL:
		if r := db.SSL.FindOneMatchNonZeros(row); r != nil {
			return r
		}
		return nil
	}
	panic(types.ErrBadType)
}

func (db OVNSouthbound) FindOneMatchByAnyIndex(irow types.IRow) types.IRow {
	switch row := irow.(type) {
	case *AddressSet:
		if r := db.AddressSet.OvsdbGetByAnyIndex(row); r != nil {
			return r
		}
		return nil
	case *Chassis:
		if r := db.Chassis.OvsdbGetByAnyIndex(row); r != nil {
			return r
		}
		return nil
	case *Connection:
		if r := db.Connection.OvsdbGetByAnyIndex(row); r != nil {
			return r
		}
		return nil
	case *DHCPOptions:
		if r := db.DHCPOptions.OvsdbGetByAnyIndex(row); r != nil {
			return r
		}
		return nil
	case *DHCPv6Options:
		if r := db.DHCPv6Options.OvsdbGetByAnyIndex(row); r != nil {
			return r
		}
		return nil
	case *DNS:
		if r := db.DNS.OvsdbGetByAnyIndex(row); r != nil {
			return r
		}
		return nil
	case *DatapathBinding:
		if r := db.DatapathBinding.OvsdbGetByAnyIndex(row); r != nil {
			return r
		}
		return nil
	case *Encap:
		if r := db.Encap.OvsdbGetByAnyIndex(row); r != nil {
			return r
		}
		return nil
	case *GatewayChassis:
		if r := db.GatewayChassis.OvsdbGetByAnyIndex(row); r != nil {
			return r
		}
		return nil
	case *LogicalFlow:
		if r := db.LogicalFlow.OvsdbGetByAnyIndex(row); r != nil {
			return r
		}
		return nil
	case *MACBinding:
		if r := db.MACBinding.OvsdbGetByAnyIndex(row); r != nil {
			return r
		}
		return nil
	case *MulticastGroup:
		if r := db.MulticastGroup.OvsdbGetByAnyIndex(row); r != nil {
			return r
		}
		return nil
	case *PortBinding:
		if r := db.PortBinding.OvsdbGetByAnyIndex(row); r != nil {
			return r
		}
		return nil
	case *RBACPermission:
		if r := db.RBACPermission.OvsdbGetByAnyIndex(row); r != nil {
			return r
		}
		return nil
	case *RBACRole:
		if r := db.RBACRole.OvsdbGetByAnyIndex(row); r != nil {
			return r
		}
		return nil
	case *SBGlobal:
		if r := db.SBGlobal.OvsdbGetByAnyIndex(row); r != nil {
			return r
		}
		return nil
	case *SSL:
		if r := db.SSL.OvsdbGetByAnyIndex(row); r != nil {
			return r
		}
		return nil
	}
	panic(types.ErrBadType)
}

type AddressSetTable []AddressSet

var _ types.ITable = &AddressSetTable{}

func (tbl AddressSetTable) OvsdbTableName() string {
	return "Address_Set"
}

func (tbl AddressSetTable) OvsdbIsRoot() bool {
	return true
}

func (tbl AddressSetTable) Rows() []types.IRow {
	r := make([]types.IRow, len(tbl))
	for i := range tbl {
		r[i] = &tbl[i]
	}
	return r
}

func (tbl AddressSetTable) NewRow() types.IRow {
	return &AddressSet{}
}

func (tbl *AddressSetTable) AppendRow(irow types.IRow) {
	row := irow.(*AddressSet)
	*tbl = append(*tbl, *row)
}

func (tbl AddressSetTable) OvsdbHasIndex() bool {
	return true
}

func (row *AddressSet) MatchByName(row1 *AddressSet) bool {
	if !types.MatchString(row.Name, row1.Name) {
		return false
	}
	return true
}

func (tbl AddressSetTable) GetByName(row1 *AddressSet) *AddressSet {
	for i := range tbl {
		row := &tbl[i]
		if row.MatchByName(row1) {
			return row
		}
	}
	return nil
}

func (tbl AddressSetTable) OvsdbGetByAnyIndex(irow1 types.IRow) types.IRow {
	row1 := irow1.(*AddressSet)
	if !(types.IsZeroString(row1.Name)) {
		if row := tbl.GetByName(row1); row != nil {
			return row
		}
	}
	return nil
}

func (tbl AddressSetTable) FindOneMatchNonZeros(row1 *AddressSet) *AddressSet {
	for i := range tbl {
		row := &tbl[i]
		if row.MatchNonZeros(row1) {
			return row
		}
	}
	return nil
}

type AddressSet struct {
	Uuid      string   `json:"_uuid"`
	Version   string   `json:"_version"`
	Addresses []string `json:"addresses"`
	Name      string   `json:"name"`
}

var _ types.IRow = &AddressSet{}

func (row *AddressSet) OvsdbTableName() string {
	return "Address_Set"
}

func (row *AddressSet) OvsdbIsRoot() bool {
	return true
}

func (row *AddressSet) OvsdbUuid() string {
	return row.Uuid
}

func (row *AddressSet) OvsdbCmdArgs() []string {
	r := []string{}
	r = append(r, types.OvsdbCmdArgsStringMultiples("addresses", row.Addresses)...)
	r = append(r, types.OvsdbCmdArgsString("name", row.Name)...)
	return r
}

func (row *AddressSet) SetColumn(name string, val interface{}) (err error) {
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
	case "addresses":
		row.Addresses = types.EnsureStringMultiples(val)
	case "name":
		row.Name = types.EnsureString(val)
	default:
		panic(types.ErrUnknownColumn)
	}
	return
}

func (row *AddressSet) MatchNonZeros(row1 *AddressSet) bool {
	if !types.MatchUuidIfNonZero(row.Uuid, row1.Uuid) {
		return false
	}
	if !types.MatchUuidIfNonZero(row.Version, row1.Version) {
		return false
	}
	if !types.MatchStringMultiplesIfNonZero(row.Addresses, row1.Addresses) {
		return false
	}
	if !types.MatchStringIfNonZero(row.Name, row1.Name) {
		return false
	}
	return true
}

func (row *AddressSet) HasExternalIds() bool {
	return false
}

func (row *AddressSet) SetExternalId(k, v string) {
	panic(errors.Wrap(types.ErrUnknownColumn, "external_ids"))
}

func (row *AddressSet) GetExternalId(k string) (string, bool) {
	panic(errors.Wrap(types.ErrUnknownColumn, "external_ids"))
}

func (row *AddressSet) RemoveExternalId(k string) (string, bool) {
	panic(errors.Wrap(types.ErrUnknownColumn, "external_ids"))
}

type ChassisTable []Chassis

var _ types.ITable = &ChassisTable{}

func (tbl ChassisTable) OvsdbTableName() string {
	return "Chassis"
}

func (tbl ChassisTable) OvsdbIsRoot() bool {
	return true
}

func (tbl ChassisTable) Rows() []types.IRow {
	r := make([]types.IRow, len(tbl))
	for i := range tbl {
		r[i] = &tbl[i]
	}
	return r
}

func (tbl ChassisTable) NewRow() types.IRow {
	return &Chassis{}
}

func (tbl *ChassisTable) AppendRow(irow types.IRow) {
	row := irow.(*Chassis)
	*tbl = append(*tbl, *row)
}

func (tbl ChassisTable) OvsdbHasIndex() bool {
	return true
}

func (row *Chassis) MatchByName(row1 *Chassis) bool {
	if !types.MatchString(row.Name, row1.Name) {
		return false
	}
	return true
}

func (tbl ChassisTable) GetByName(row1 *Chassis) *Chassis {
	for i := range tbl {
		row := &tbl[i]
		if row.MatchByName(row1) {
			return row
		}
	}
	return nil
}

func (tbl ChassisTable) OvsdbGetByAnyIndex(irow1 types.IRow) types.IRow {
	row1 := irow1.(*Chassis)
	if !(types.IsZeroString(row1.Name)) {
		if row := tbl.GetByName(row1); row != nil {
			return row
		}
	}
	return nil
}

func (tbl ChassisTable) FindOneMatchNonZeros(row1 *Chassis) *Chassis {
	for i := range tbl {
		row := &tbl[i]
		if row.MatchNonZeros(row1) {
			return row
		}
	}
	return nil
}

type Chassis struct {
	Uuid                string            `json:"_uuid"`
	Version             string            `json:"_version"`
	Encaps              []string          `json:"encaps"`
	ExternalIds         map[string]string `json:"external_ids"`
	Hostname            string            `json:"hostname"`
	Name                string            `json:"name"`
	NbCfg               int64             `json:"nb_cfg"`
	VtepLogicalSwitches []string          `json:"vtep_logical_switches"`
}

var _ types.IRow = &Chassis{}

func (row *Chassis) OvsdbTableName() string {
	return "Chassis"
}

func (row *Chassis) OvsdbIsRoot() bool {
	return true
}

func (row *Chassis) OvsdbUuid() string {
	return row.Uuid
}

func (row *Chassis) OvsdbCmdArgs() []string {
	r := []string{}
	r = append(r, types.OvsdbCmdArgsUuidMultiples("encaps", row.Encaps)...)
	r = append(r, types.OvsdbCmdArgsMapStringString("external_ids", row.ExternalIds)...)
	r = append(r, types.OvsdbCmdArgsString("hostname", row.Hostname)...)
	r = append(r, types.OvsdbCmdArgsString("name", row.Name)...)
	r = append(r, types.OvsdbCmdArgsInteger("nb_cfg", row.NbCfg)...)
	r = append(r, types.OvsdbCmdArgsStringMultiples("vtep_logical_switches", row.VtepLogicalSwitches)...)
	return r
}

func (row *Chassis) SetColumn(name string, val interface{}) (err error) {
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
	case "encaps":
		row.Encaps = types.EnsureUuidMultiples(val)
	case "external_ids":
		row.ExternalIds = types.EnsureMapStringString(val)
	case "hostname":
		row.Hostname = types.EnsureString(val)
	case "name":
		row.Name = types.EnsureString(val)
	case "nb_cfg":
		row.NbCfg = types.EnsureInteger(val)
	case "vtep_logical_switches":
		row.VtepLogicalSwitches = types.EnsureStringMultiples(val)
	default:
		panic(types.ErrUnknownColumn)
	}
	return
}

func (row *Chassis) MatchNonZeros(row1 *Chassis) bool {
	if !types.MatchUuidIfNonZero(row.Uuid, row1.Uuid) {
		return false
	}
	if !types.MatchUuidIfNonZero(row.Version, row1.Version) {
		return false
	}
	if !types.MatchUuidMultiplesIfNonZero(row.Encaps, row1.Encaps) {
		return false
	}
	if !types.MatchMapStringStringIfNonZero(row.ExternalIds, row1.ExternalIds) {
		return false
	}
	if !types.MatchStringIfNonZero(row.Hostname, row1.Hostname) {
		return false
	}
	if !types.MatchStringIfNonZero(row.Name, row1.Name) {
		return false
	}
	if !types.MatchIntegerIfNonZero(row.NbCfg, row1.NbCfg) {
		return false
	}
	if !types.MatchStringMultiplesIfNonZero(row.VtepLogicalSwitches, row1.VtepLogicalSwitches) {
		return false
	}
	return true
}

func (row *Chassis) HasExternalIds() bool {
	return true
}

func (row *Chassis) SetExternalId(k, v string) {
	if row.ExternalIds == nil {
		row.ExternalIds = map[string]string{}
	}
	row.ExternalIds[k] = v
}

func (row *Chassis) GetExternalId(k string) (string, bool) {
	if row.ExternalIds == nil {
		return "", false
	}
	r, ok := row.ExternalIds[k]
	return r, ok
}

func (row *Chassis) RemoveExternalId(k string) (string, bool) {
	if row.ExternalIds == nil {
		return "", false
	}
	r, ok := row.ExternalIds[k]
	if ok {
		delete(row.ExternalIds, k)
	}
	return r, ok
}

type ConnectionTable []Connection

var _ types.ITable = &ConnectionTable{}

func (tbl ConnectionTable) OvsdbTableName() string {
	return "Connection"
}

func (tbl ConnectionTable) OvsdbIsRoot() bool {
	return false
}

func (tbl ConnectionTable) Rows() []types.IRow {
	r := make([]types.IRow, len(tbl))
	for i := range tbl {
		r[i] = &tbl[i]
	}
	return r
}

func (tbl ConnectionTable) NewRow() types.IRow {
	return &Connection{}
}

func (tbl *ConnectionTable) AppendRow(irow types.IRow) {
	row := irow.(*Connection)
	*tbl = append(*tbl, *row)
}

func (tbl ConnectionTable) OvsdbHasIndex() bool {
	return true
}

func (row *Connection) MatchByTarget(row1 *Connection) bool {
	if !types.MatchString(row.Target, row1.Target) {
		return false
	}
	return true
}

func (tbl ConnectionTable) GetByTarget(row1 *Connection) *Connection {
	for i := range tbl {
		row := &tbl[i]
		if row.MatchByTarget(row1) {
			return row
		}
	}
	return nil
}

func (tbl ConnectionTable) OvsdbGetByAnyIndex(irow1 types.IRow) types.IRow {
	row1 := irow1.(*Connection)
	if !(types.IsZeroString(row1.Target)) {
		if row := tbl.GetByTarget(row1); row != nil {
			return row
		}
	}
	return nil
}

func (tbl ConnectionTable) FindOneMatchNonZeros(row1 *Connection) *Connection {
	for i := range tbl {
		row := &tbl[i]
		if row.MatchNonZeros(row1) {
			return row
		}
	}
	return nil
}

type Connection struct {
	Uuid            string            `json:"_uuid"`
	Version         string            `json:"_version"`
	ExternalIds     map[string]string `json:"external_ids"`
	InactivityProbe *int64            `json:"inactivity_probe"`
	IsConnected     bool              `json:"is_connected"`
	MaxBackoff      *int64            `json:"max_backoff"`
	OtherConfig     map[string]string `json:"other_config"`
	ReadOnly        bool              `json:"read_only"`
	Role            string            `json:"role"`
	Status          map[string]string `json:"status"`
	Target          string            `json:"target"`
}

var _ types.IRow = &Connection{}

func (row *Connection) OvsdbTableName() string {
	return "Connection"
}

func (row *Connection) OvsdbIsRoot() bool {
	return false
}

func (row *Connection) OvsdbUuid() string {
	return row.Uuid
}

func (row *Connection) OvsdbCmdArgs() []string {
	r := []string{}
	r = append(r, types.OvsdbCmdArgsMapStringString("external_ids", row.ExternalIds)...)
	r = append(r, types.OvsdbCmdArgsIntegerOptional("inactivity_probe", row.InactivityProbe)...)
	r = append(r, types.OvsdbCmdArgsBoolean("is_connected", row.IsConnected)...)
	r = append(r, types.OvsdbCmdArgsIntegerOptional("max_backoff", row.MaxBackoff)...)
	r = append(r, types.OvsdbCmdArgsMapStringString("other_config", row.OtherConfig)...)
	r = append(r, types.OvsdbCmdArgsBoolean("read_only", row.ReadOnly)...)
	r = append(r, types.OvsdbCmdArgsString("role", row.Role)...)
	r = append(r, types.OvsdbCmdArgsMapStringString("status", row.Status)...)
	r = append(r, types.OvsdbCmdArgsString("target", row.Target)...)
	return r
}

func (row *Connection) SetColumn(name string, val interface{}) (err error) {
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
	case "inactivity_probe":
		row.InactivityProbe = types.EnsureIntegerOptional(val)
	case "is_connected":
		row.IsConnected = types.EnsureBoolean(val)
	case "max_backoff":
		row.MaxBackoff = types.EnsureIntegerOptional(val)
	case "other_config":
		row.OtherConfig = types.EnsureMapStringString(val)
	case "read_only":
		row.ReadOnly = types.EnsureBoolean(val)
	case "role":
		row.Role = types.EnsureString(val)
	case "status":
		row.Status = types.EnsureMapStringString(val)
	case "target":
		row.Target = types.EnsureString(val)
	default:
		panic(types.ErrUnknownColumn)
	}
	return
}

func (row *Connection) MatchNonZeros(row1 *Connection) bool {
	if !types.MatchUuidIfNonZero(row.Uuid, row1.Uuid) {
		return false
	}
	if !types.MatchUuidIfNonZero(row.Version, row1.Version) {
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
	if !types.MatchBooleanIfNonZero(row.ReadOnly, row1.ReadOnly) {
		return false
	}
	if !types.MatchStringIfNonZero(row.Role, row1.Role) {
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

func (row *Connection) HasExternalIds() bool {
	return true
}

func (row *Connection) SetExternalId(k, v string) {
	if row.ExternalIds == nil {
		row.ExternalIds = map[string]string{}
	}
	row.ExternalIds[k] = v
}

func (row *Connection) GetExternalId(k string) (string, bool) {
	if row.ExternalIds == nil {
		return "", false
	}
	r, ok := row.ExternalIds[k]
	return r, ok
}

func (row *Connection) RemoveExternalId(k string) (string, bool) {
	if row.ExternalIds == nil {
		return "", false
	}
	r, ok := row.ExternalIds[k]
	if ok {
		delete(row.ExternalIds, k)
	}
	return r, ok
}

type DHCPOptionsTable []DHCPOptions

var _ types.ITable = &DHCPOptionsTable{}

func (tbl DHCPOptionsTable) OvsdbTableName() string {
	return "DHCP_Options"
}

func (tbl DHCPOptionsTable) OvsdbIsRoot() bool {
	return true
}

func (tbl DHCPOptionsTable) Rows() []types.IRow {
	r := make([]types.IRow, len(tbl))
	for i := range tbl {
		r[i] = &tbl[i]
	}
	return r
}

func (tbl DHCPOptionsTable) NewRow() types.IRow {
	return &DHCPOptions{}
}

func (tbl *DHCPOptionsTable) AppendRow(irow types.IRow) {
	row := irow.(*DHCPOptions)
	*tbl = append(*tbl, *row)
}

func (tbl DHCPOptionsTable) OvsdbHasIndex() bool {
	return false
}

func (tbl DHCPOptionsTable) OvsdbGetByAnyIndex(irow1 types.IRow) types.IRow {
	return nil
}

func (tbl DHCPOptionsTable) FindOneMatchNonZeros(row1 *DHCPOptions) *DHCPOptions {
	for i := range tbl {
		row := &tbl[i]
		if row.MatchNonZeros(row1) {
			return row
		}
	}
	return nil
}

type DHCPOptions struct {
	Uuid    string `json:"_uuid"`
	Version string `json:"_version"`
	Code    int64  `json:"code"`
	Name    string `json:"name"`
	Type    string `json:"type"`
}

var _ types.IRow = &DHCPOptions{}

func (row *DHCPOptions) OvsdbTableName() string {
	return "DHCP_Options"
}

func (row *DHCPOptions) OvsdbIsRoot() bool {
	return true
}

func (row *DHCPOptions) OvsdbUuid() string {
	return row.Uuid
}

func (row *DHCPOptions) OvsdbCmdArgs() []string {
	r := []string{}
	r = append(r, types.OvsdbCmdArgsInteger("code", row.Code)...)
	r = append(r, types.OvsdbCmdArgsString("name", row.Name)...)
	r = append(r, types.OvsdbCmdArgsString("type", row.Type)...)
	return r
}

func (row *DHCPOptions) SetColumn(name string, val interface{}) (err error) {
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
	case "code":
		row.Code = types.EnsureInteger(val)
	case "name":
		row.Name = types.EnsureString(val)
	case "type":
		row.Type = types.EnsureString(val)
	default:
		panic(types.ErrUnknownColumn)
	}
	return
}

func (row *DHCPOptions) MatchNonZeros(row1 *DHCPOptions) bool {
	if !types.MatchUuidIfNonZero(row.Uuid, row1.Uuid) {
		return false
	}
	if !types.MatchUuidIfNonZero(row.Version, row1.Version) {
		return false
	}
	if !types.MatchIntegerIfNonZero(row.Code, row1.Code) {
		return false
	}
	if !types.MatchStringIfNonZero(row.Name, row1.Name) {
		return false
	}
	if !types.MatchStringIfNonZero(row.Type, row1.Type) {
		return false
	}
	return true
}

func (row *DHCPOptions) HasExternalIds() bool {
	return false
}

func (row *DHCPOptions) SetExternalId(k, v string) {
	panic(errors.Wrap(types.ErrUnknownColumn, "external_ids"))
}

func (row *DHCPOptions) GetExternalId(k string) (string, bool) {
	panic(errors.Wrap(types.ErrUnknownColumn, "external_ids"))
}

func (row *DHCPOptions) RemoveExternalId(k string) (string, bool) {
	panic(errors.Wrap(types.ErrUnknownColumn, "external_ids"))
}

type DHCPv6OptionsTable []DHCPv6Options

var _ types.ITable = &DHCPv6OptionsTable{}

func (tbl DHCPv6OptionsTable) OvsdbTableName() string {
	return "DHCPv6_Options"
}

func (tbl DHCPv6OptionsTable) OvsdbIsRoot() bool {
	return true
}

func (tbl DHCPv6OptionsTable) Rows() []types.IRow {
	r := make([]types.IRow, len(tbl))
	for i := range tbl {
		r[i] = &tbl[i]
	}
	return r
}

func (tbl DHCPv6OptionsTable) NewRow() types.IRow {
	return &DHCPv6Options{}
}

func (tbl *DHCPv6OptionsTable) AppendRow(irow types.IRow) {
	row := irow.(*DHCPv6Options)
	*tbl = append(*tbl, *row)
}

func (tbl DHCPv6OptionsTable) OvsdbHasIndex() bool {
	return false
}

func (tbl DHCPv6OptionsTable) OvsdbGetByAnyIndex(irow1 types.IRow) types.IRow {
	return nil
}

func (tbl DHCPv6OptionsTable) FindOneMatchNonZeros(row1 *DHCPv6Options) *DHCPv6Options {
	for i := range tbl {
		row := &tbl[i]
		if row.MatchNonZeros(row1) {
			return row
		}
	}
	return nil
}

type DHCPv6Options struct {
	Uuid    string `json:"_uuid"`
	Version string `json:"_version"`
	Code    int64  `json:"code"`
	Name    string `json:"name"`
	Type    string `json:"type"`
}

var _ types.IRow = &DHCPv6Options{}

func (row *DHCPv6Options) OvsdbTableName() string {
	return "DHCPv6_Options"
}

func (row *DHCPv6Options) OvsdbIsRoot() bool {
	return true
}

func (row *DHCPv6Options) OvsdbUuid() string {
	return row.Uuid
}

func (row *DHCPv6Options) OvsdbCmdArgs() []string {
	r := []string{}
	r = append(r, types.OvsdbCmdArgsInteger("code", row.Code)...)
	r = append(r, types.OvsdbCmdArgsString("name", row.Name)...)
	r = append(r, types.OvsdbCmdArgsString("type", row.Type)...)
	return r
}

func (row *DHCPv6Options) SetColumn(name string, val interface{}) (err error) {
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
	case "code":
		row.Code = types.EnsureInteger(val)
	case "name":
		row.Name = types.EnsureString(val)
	case "type":
		row.Type = types.EnsureString(val)
	default:
		panic(types.ErrUnknownColumn)
	}
	return
}

func (row *DHCPv6Options) MatchNonZeros(row1 *DHCPv6Options) bool {
	if !types.MatchUuidIfNonZero(row.Uuid, row1.Uuid) {
		return false
	}
	if !types.MatchUuidIfNonZero(row.Version, row1.Version) {
		return false
	}
	if !types.MatchIntegerIfNonZero(row.Code, row1.Code) {
		return false
	}
	if !types.MatchStringIfNonZero(row.Name, row1.Name) {
		return false
	}
	if !types.MatchStringIfNonZero(row.Type, row1.Type) {
		return false
	}
	return true
}

func (row *DHCPv6Options) HasExternalIds() bool {
	return false
}

func (row *DHCPv6Options) SetExternalId(k, v string) {
	panic(errors.Wrap(types.ErrUnknownColumn, "external_ids"))
}

func (row *DHCPv6Options) GetExternalId(k string) (string, bool) {
	panic(errors.Wrap(types.ErrUnknownColumn, "external_ids"))
}

func (row *DHCPv6Options) RemoveExternalId(k string) (string, bool) {
	panic(errors.Wrap(types.ErrUnknownColumn, "external_ids"))
}

type DNSTable []DNS

var _ types.ITable = &DNSTable{}

func (tbl DNSTable) OvsdbTableName() string {
	return "DNS"
}

func (tbl DNSTable) OvsdbIsRoot() bool {
	return true
}

func (tbl DNSTable) Rows() []types.IRow {
	r := make([]types.IRow, len(tbl))
	for i := range tbl {
		r[i] = &tbl[i]
	}
	return r
}

func (tbl DNSTable) NewRow() types.IRow {
	return &DNS{}
}

func (tbl *DNSTable) AppendRow(irow types.IRow) {
	row := irow.(*DNS)
	*tbl = append(*tbl, *row)
}

func (tbl DNSTable) OvsdbHasIndex() bool {
	return false
}

func (tbl DNSTable) OvsdbGetByAnyIndex(irow1 types.IRow) types.IRow {
	return nil
}

func (tbl DNSTable) FindOneMatchNonZeros(row1 *DNS) *DNS {
	for i := range tbl {
		row := &tbl[i]
		if row.MatchNonZeros(row1) {
			return row
		}
	}
	return nil
}

type DNS struct {
	Uuid        string            `json:"_uuid"`
	Version     string            `json:"_version"`
	Datapaths   []string          `json:"datapaths"`
	ExternalIds map[string]string `json:"external_ids"`
	Records     map[string]string `json:"records"`
}

var _ types.IRow = &DNS{}

func (row *DNS) OvsdbTableName() string {
	return "DNS"
}

func (row *DNS) OvsdbIsRoot() bool {
	return true
}

func (row *DNS) OvsdbUuid() string {
	return row.Uuid
}

func (row *DNS) OvsdbCmdArgs() []string {
	r := []string{}
	r = append(r, types.OvsdbCmdArgsUuidMultiples("datapaths", row.Datapaths)...)
	r = append(r, types.OvsdbCmdArgsMapStringString("external_ids", row.ExternalIds)...)
	r = append(r, types.OvsdbCmdArgsMapStringString("records", row.Records)...)
	return r
}

func (row *DNS) SetColumn(name string, val interface{}) (err error) {
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
	case "datapaths":
		row.Datapaths = types.EnsureUuidMultiples(val)
	case "external_ids":
		row.ExternalIds = types.EnsureMapStringString(val)
	case "records":
		row.Records = types.EnsureMapStringString(val)
	default:
		panic(types.ErrUnknownColumn)
	}
	return
}

func (row *DNS) MatchNonZeros(row1 *DNS) bool {
	if !types.MatchUuidIfNonZero(row.Uuid, row1.Uuid) {
		return false
	}
	if !types.MatchUuidIfNonZero(row.Version, row1.Version) {
		return false
	}
	if !types.MatchUuidMultiplesIfNonZero(row.Datapaths, row1.Datapaths) {
		return false
	}
	if !types.MatchMapStringStringIfNonZero(row.ExternalIds, row1.ExternalIds) {
		return false
	}
	if !types.MatchMapStringStringIfNonZero(row.Records, row1.Records) {
		return false
	}
	return true
}

func (row *DNS) HasExternalIds() bool {
	return true
}

func (row *DNS) SetExternalId(k, v string) {
	if row.ExternalIds == nil {
		row.ExternalIds = map[string]string{}
	}
	row.ExternalIds[k] = v
}

func (row *DNS) GetExternalId(k string) (string, bool) {
	if row.ExternalIds == nil {
		return "", false
	}
	r, ok := row.ExternalIds[k]
	return r, ok
}

func (row *DNS) RemoveExternalId(k string) (string, bool) {
	if row.ExternalIds == nil {
		return "", false
	}
	r, ok := row.ExternalIds[k]
	if ok {
		delete(row.ExternalIds, k)
	}
	return r, ok
}

type DatapathBindingTable []DatapathBinding

var _ types.ITable = &DatapathBindingTable{}

func (tbl DatapathBindingTable) OvsdbTableName() string {
	return "Datapath_Binding"
}

func (tbl DatapathBindingTable) OvsdbIsRoot() bool {
	return true
}

func (tbl DatapathBindingTable) Rows() []types.IRow {
	r := make([]types.IRow, len(tbl))
	for i := range tbl {
		r[i] = &tbl[i]
	}
	return r
}

func (tbl DatapathBindingTable) NewRow() types.IRow {
	return &DatapathBinding{}
}

func (tbl *DatapathBindingTable) AppendRow(irow types.IRow) {
	row := irow.(*DatapathBinding)
	*tbl = append(*tbl, *row)
}

func (tbl DatapathBindingTable) OvsdbHasIndex() bool {
	return true
}

func (row *DatapathBinding) MatchByTunnelKey(row1 *DatapathBinding) bool {
	if !types.MatchInteger(row.TunnelKey, row1.TunnelKey) {
		return false
	}
	return true
}

func (tbl DatapathBindingTable) GetByTunnelKey(row1 *DatapathBinding) *DatapathBinding {
	for i := range tbl {
		row := &tbl[i]
		if row.MatchByTunnelKey(row1) {
			return row
		}
	}
	return nil
}

func (tbl DatapathBindingTable) OvsdbGetByAnyIndex(irow1 types.IRow) types.IRow {
	row1 := irow1.(*DatapathBinding)
	if !(types.IsZeroInteger(row1.TunnelKey)) {
		if row := tbl.GetByTunnelKey(row1); row != nil {
			return row
		}
	}
	return nil
}

func (tbl DatapathBindingTable) FindOneMatchNonZeros(row1 *DatapathBinding) *DatapathBinding {
	for i := range tbl {
		row := &tbl[i]
		if row.MatchNonZeros(row1) {
			return row
		}
	}
	return nil
}

type DatapathBinding struct {
	Uuid        string            `json:"_uuid"`
	Version     string            `json:"_version"`
	ExternalIds map[string]string `json:"external_ids"`
	TunnelKey   int64             `json:"tunnel_key"`
}

var _ types.IRow = &DatapathBinding{}

func (row *DatapathBinding) OvsdbTableName() string {
	return "Datapath_Binding"
}

func (row *DatapathBinding) OvsdbIsRoot() bool {
	return true
}

func (row *DatapathBinding) OvsdbUuid() string {
	return row.Uuid
}

func (row *DatapathBinding) OvsdbCmdArgs() []string {
	r := []string{}
	r = append(r, types.OvsdbCmdArgsMapStringString("external_ids", row.ExternalIds)...)
	r = append(r, types.OvsdbCmdArgsInteger("tunnel_key", row.TunnelKey)...)
	return r
}

func (row *DatapathBinding) SetColumn(name string, val interface{}) (err error) {
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
	case "tunnel_key":
		row.TunnelKey = types.EnsureInteger(val)
	default:
		panic(types.ErrUnknownColumn)
	}
	return
}

func (row *DatapathBinding) MatchNonZeros(row1 *DatapathBinding) bool {
	if !types.MatchUuidIfNonZero(row.Uuid, row1.Uuid) {
		return false
	}
	if !types.MatchUuidIfNonZero(row.Version, row1.Version) {
		return false
	}
	if !types.MatchMapStringStringIfNonZero(row.ExternalIds, row1.ExternalIds) {
		return false
	}
	if !types.MatchIntegerIfNonZero(row.TunnelKey, row1.TunnelKey) {
		return false
	}
	return true
}

func (row *DatapathBinding) HasExternalIds() bool {
	return true
}

func (row *DatapathBinding) SetExternalId(k, v string) {
	if row.ExternalIds == nil {
		row.ExternalIds = map[string]string{}
	}
	row.ExternalIds[k] = v
}

func (row *DatapathBinding) GetExternalId(k string) (string, bool) {
	if row.ExternalIds == nil {
		return "", false
	}
	r, ok := row.ExternalIds[k]
	return r, ok
}

func (row *DatapathBinding) RemoveExternalId(k string) (string, bool) {
	if row.ExternalIds == nil {
		return "", false
	}
	r, ok := row.ExternalIds[k]
	if ok {
		delete(row.ExternalIds, k)
	}
	return r, ok
}

type EncapTable []Encap

var _ types.ITable = &EncapTable{}

func (tbl EncapTable) OvsdbTableName() string {
	return "Encap"
}

func (tbl EncapTable) OvsdbIsRoot() bool {
	return false
}

func (tbl EncapTable) Rows() []types.IRow {
	r := make([]types.IRow, len(tbl))
	for i := range tbl {
		r[i] = &tbl[i]
	}
	return r
}

func (tbl EncapTable) NewRow() types.IRow {
	return &Encap{}
}

func (tbl *EncapTable) AppendRow(irow types.IRow) {
	row := irow.(*Encap)
	*tbl = append(*tbl, *row)
}

func (tbl EncapTable) OvsdbHasIndex() bool {
	return false
}

func (tbl EncapTable) OvsdbGetByAnyIndex(irow1 types.IRow) types.IRow {
	return nil
}

func (tbl EncapTable) FindOneMatchNonZeros(row1 *Encap) *Encap {
	for i := range tbl {
		row := &tbl[i]
		if row.MatchNonZeros(row1) {
			return row
		}
	}
	return nil
}

type Encap struct {
	Uuid        string            `json:"_uuid"`
	Version     string            `json:"_version"`
	ChassisName string            `json:"chassis_name"`
	Ip          string            `json:"ip"`
	Options     map[string]string `json:"options"`
	Type        string            `json:"type"`
}

var _ types.IRow = &Encap{}

func (row *Encap) OvsdbTableName() string {
	return "Encap"
}

func (row *Encap) OvsdbIsRoot() bool {
	return false
}

func (row *Encap) OvsdbUuid() string {
	return row.Uuid
}

func (row *Encap) OvsdbCmdArgs() []string {
	r := []string{}
	r = append(r, types.OvsdbCmdArgsString("chassis_name", row.ChassisName)...)
	r = append(r, types.OvsdbCmdArgsString("ip", row.Ip)...)
	r = append(r, types.OvsdbCmdArgsMapStringString("options", row.Options)...)
	r = append(r, types.OvsdbCmdArgsString("type", row.Type)...)
	return r
}

func (row *Encap) SetColumn(name string, val interface{}) (err error) {
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
	case "chassis_name":
		row.ChassisName = types.EnsureString(val)
	case "ip":
		row.Ip = types.EnsureString(val)
	case "options":
		row.Options = types.EnsureMapStringString(val)
	case "type":
		row.Type = types.EnsureString(val)
	default:
		panic(types.ErrUnknownColumn)
	}
	return
}

func (row *Encap) MatchNonZeros(row1 *Encap) bool {
	if !types.MatchUuidIfNonZero(row.Uuid, row1.Uuid) {
		return false
	}
	if !types.MatchUuidIfNonZero(row.Version, row1.Version) {
		return false
	}
	if !types.MatchStringIfNonZero(row.ChassisName, row1.ChassisName) {
		return false
	}
	if !types.MatchStringIfNonZero(row.Ip, row1.Ip) {
		return false
	}
	if !types.MatchMapStringStringIfNonZero(row.Options, row1.Options) {
		return false
	}
	if !types.MatchStringIfNonZero(row.Type, row1.Type) {
		return false
	}
	return true
}

func (row *Encap) HasExternalIds() bool {
	return false
}

func (row *Encap) SetExternalId(k, v string) {
	panic(errors.Wrap(types.ErrUnknownColumn, "external_ids"))
}

func (row *Encap) GetExternalId(k string) (string, bool) {
	panic(errors.Wrap(types.ErrUnknownColumn, "external_ids"))
}

func (row *Encap) RemoveExternalId(k string) (string, bool) {
	panic(errors.Wrap(types.ErrUnknownColumn, "external_ids"))
}

type GatewayChassisTable []GatewayChassis

var _ types.ITable = &GatewayChassisTable{}

func (tbl GatewayChassisTable) OvsdbTableName() string {
	return "Gateway_Chassis"
}

func (tbl GatewayChassisTable) OvsdbIsRoot() bool {
	return false
}

func (tbl GatewayChassisTable) Rows() []types.IRow {
	r := make([]types.IRow, len(tbl))
	for i := range tbl {
		r[i] = &tbl[i]
	}
	return r
}

func (tbl GatewayChassisTable) NewRow() types.IRow {
	return &GatewayChassis{}
}

func (tbl *GatewayChassisTable) AppendRow(irow types.IRow) {
	row := irow.(*GatewayChassis)
	*tbl = append(*tbl, *row)
}

func (tbl GatewayChassisTable) OvsdbHasIndex() bool {
	return true
}

func (row *GatewayChassis) MatchByName(row1 *GatewayChassis) bool {
	if !types.MatchString(row.Name, row1.Name) {
		return false
	}
	return true
}

func (tbl GatewayChassisTable) GetByName(row1 *GatewayChassis) *GatewayChassis {
	for i := range tbl {
		row := &tbl[i]
		if row.MatchByName(row1) {
			return row
		}
	}
	return nil
}

func (tbl GatewayChassisTable) OvsdbGetByAnyIndex(irow1 types.IRow) types.IRow {
	row1 := irow1.(*GatewayChassis)
	if !(types.IsZeroString(row1.Name)) {
		if row := tbl.GetByName(row1); row != nil {
			return row
		}
	}
	return nil
}

func (tbl GatewayChassisTable) FindOneMatchNonZeros(row1 *GatewayChassis) *GatewayChassis {
	for i := range tbl {
		row := &tbl[i]
		if row.MatchNonZeros(row1) {
			return row
		}
	}
	return nil
}

type GatewayChassis struct {
	Uuid        string            `json:"_uuid"`
	Version     string            `json:"_version"`
	Chassis     *string           `json:"chassis"`
	ExternalIds map[string]string `json:"external_ids"`
	Name        string            `json:"name"`
	Options     map[string]string `json:"options"`
	Priority    int64             `json:"priority"`
}

var _ types.IRow = &GatewayChassis{}

func (row *GatewayChassis) OvsdbTableName() string {
	return "Gateway_Chassis"
}

func (row *GatewayChassis) OvsdbIsRoot() bool {
	return false
}

func (row *GatewayChassis) OvsdbUuid() string {
	return row.Uuid
}

func (row *GatewayChassis) OvsdbCmdArgs() []string {
	r := []string{}
	r = append(r, types.OvsdbCmdArgsUuidOptional("chassis", row.Chassis)...)
	r = append(r, types.OvsdbCmdArgsMapStringString("external_ids", row.ExternalIds)...)
	r = append(r, types.OvsdbCmdArgsString("name", row.Name)...)
	r = append(r, types.OvsdbCmdArgsMapStringString("options", row.Options)...)
	r = append(r, types.OvsdbCmdArgsInteger("priority", row.Priority)...)
	return r
}

func (row *GatewayChassis) SetColumn(name string, val interface{}) (err error) {
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
	case "chassis":
		row.Chassis = types.EnsureUuidOptional(val)
	case "external_ids":
		row.ExternalIds = types.EnsureMapStringString(val)
	case "name":
		row.Name = types.EnsureString(val)
	case "options":
		row.Options = types.EnsureMapStringString(val)
	case "priority":
		row.Priority = types.EnsureInteger(val)
	default:
		panic(types.ErrUnknownColumn)
	}
	return
}

func (row *GatewayChassis) MatchNonZeros(row1 *GatewayChassis) bool {
	if !types.MatchUuidIfNonZero(row.Uuid, row1.Uuid) {
		return false
	}
	if !types.MatchUuidIfNonZero(row.Version, row1.Version) {
		return false
	}
	if !types.MatchUuidOptionalIfNonZero(row.Chassis, row1.Chassis) {
		return false
	}
	if !types.MatchMapStringStringIfNonZero(row.ExternalIds, row1.ExternalIds) {
		return false
	}
	if !types.MatchStringIfNonZero(row.Name, row1.Name) {
		return false
	}
	if !types.MatchMapStringStringIfNonZero(row.Options, row1.Options) {
		return false
	}
	if !types.MatchIntegerIfNonZero(row.Priority, row1.Priority) {
		return false
	}
	return true
}

func (row *GatewayChassis) HasExternalIds() bool {
	return true
}

func (row *GatewayChassis) SetExternalId(k, v string) {
	if row.ExternalIds == nil {
		row.ExternalIds = map[string]string{}
	}
	row.ExternalIds[k] = v
}

func (row *GatewayChassis) GetExternalId(k string) (string, bool) {
	if row.ExternalIds == nil {
		return "", false
	}
	r, ok := row.ExternalIds[k]
	return r, ok
}

func (row *GatewayChassis) RemoveExternalId(k string) (string, bool) {
	if row.ExternalIds == nil {
		return "", false
	}
	r, ok := row.ExternalIds[k]
	if ok {
		delete(row.ExternalIds, k)
	}
	return r, ok
}

type LogicalFlowTable []LogicalFlow

var _ types.ITable = &LogicalFlowTable{}

func (tbl LogicalFlowTable) OvsdbTableName() string {
	return "Logical_Flow"
}

func (tbl LogicalFlowTable) OvsdbIsRoot() bool {
	return true
}

func (tbl LogicalFlowTable) Rows() []types.IRow {
	r := make([]types.IRow, len(tbl))
	for i := range tbl {
		r[i] = &tbl[i]
	}
	return r
}

func (tbl LogicalFlowTable) NewRow() types.IRow {
	return &LogicalFlow{}
}

func (tbl *LogicalFlowTable) AppendRow(irow types.IRow) {
	row := irow.(*LogicalFlow)
	*tbl = append(*tbl, *row)
}

func (tbl LogicalFlowTable) OvsdbHasIndex() bool {
	return false
}

func (tbl LogicalFlowTable) OvsdbGetByAnyIndex(irow1 types.IRow) types.IRow {
	return nil
}

func (tbl LogicalFlowTable) FindOneMatchNonZeros(row1 *LogicalFlow) *LogicalFlow {
	for i := range tbl {
		row := &tbl[i]
		if row.MatchNonZeros(row1) {
			return row
		}
	}
	return nil
}

type LogicalFlow struct {
	Uuid            string            `json:"_uuid"`
	Version         string            `json:"_version"`
	Actions         string            `json:"actions"`
	ExternalIds     map[string]string `json:"external_ids"`
	LogicalDatapath string            `json:"logical_datapath"`
	Match           string            `json:"match"`
	Pipeline        string            `json:"pipeline"`
	Priority        int64             `json:"priority"`
	TableId         int64             `json:"table_id"`
}

var _ types.IRow = &LogicalFlow{}

func (row *LogicalFlow) OvsdbTableName() string {
	return "Logical_Flow"
}

func (row *LogicalFlow) OvsdbIsRoot() bool {
	return true
}

func (row *LogicalFlow) OvsdbUuid() string {
	return row.Uuid
}

func (row *LogicalFlow) OvsdbCmdArgs() []string {
	r := []string{}
	r = append(r, types.OvsdbCmdArgsString("actions", row.Actions)...)
	r = append(r, types.OvsdbCmdArgsMapStringString("external_ids", row.ExternalIds)...)
	r = append(r, types.OvsdbCmdArgsUuid("logical_datapath", row.LogicalDatapath)...)
	r = append(r, types.OvsdbCmdArgsString("match", row.Match)...)
	r = append(r, types.OvsdbCmdArgsString("pipeline", row.Pipeline)...)
	r = append(r, types.OvsdbCmdArgsInteger("priority", row.Priority)...)
	r = append(r, types.OvsdbCmdArgsInteger("table_id", row.TableId)...)
	return r
}

func (row *LogicalFlow) SetColumn(name string, val interface{}) (err error) {
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
	case "actions":
		row.Actions = types.EnsureString(val)
	case "external_ids":
		row.ExternalIds = types.EnsureMapStringString(val)
	case "logical_datapath":
		row.LogicalDatapath = types.EnsureUuid(val)
	case "match":
		row.Match = types.EnsureString(val)
	case "pipeline":
		row.Pipeline = types.EnsureString(val)
	case "priority":
		row.Priority = types.EnsureInteger(val)
	case "table_id":
		row.TableId = types.EnsureInteger(val)
	default:
		panic(types.ErrUnknownColumn)
	}
	return
}

func (row *LogicalFlow) MatchNonZeros(row1 *LogicalFlow) bool {
	if !types.MatchUuidIfNonZero(row.Uuid, row1.Uuid) {
		return false
	}
	if !types.MatchUuidIfNonZero(row.Version, row1.Version) {
		return false
	}
	if !types.MatchStringIfNonZero(row.Actions, row1.Actions) {
		return false
	}
	if !types.MatchMapStringStringIfNonZero(row.ExternalIds, row1.ExternalIds) {
		return false
	}
	if !types.MatchUuidIfNonZero(row.LogicalDatapath, row1.LogicalDatapath) {
		return false
	}
	if !types.MatchStringIfNonZero(row.Match, row1.Match) {
		return false
	}
	if !types.MatchStringIfNonZero(row.Pipeline, row1.Pipeline) {
		return false
	}
	if !types.MatchIntegerIfNonZero(row.Priority, row1.Priority) {
		return false
	}
	if !types.MatchIntegerIfNonZero(row.TableId, row1.TableId) {
		return false
	}
	return true
}

func (row *LogicalFlow) HasExternalIds() bool {
	return true
}

func (row *LogicalFlow) SetExternalId(k, v string) {
	if row.ExternalIds == nil {
		row.ExternalIds = map[string]string{}
	}
	row.ExternalIds[k] = v
}

func (row *LogicalFlow) GetExternalId(k string) (string, bool) {
	if row.ExternalIds == nil {
		return "", false
	}
	r, ok := row.ExternalIds[k]
	return r, ok
}

func (row *LogicalFlow) RemoveExternalId(k string) (string, bool) {
	if row.ExternalIds == nil {
		return "", false
	}
	r, ok := row.ExternalIds[k]
	if ok {
		delete(row.ExternalIds, k)
	}
	return r, ok
}

type MACBindingTable []MACBinding

var _ types.ITable = &MACBindingTable{}

func (tbl MACBindingTable) OvsdbTableName() string {
	return "MAC_Binding"
}

func (tbl MACBindingTable) OvsdbIsRoot() bool {
	return true
}

func (tbl MACBindingTable) Rows() []types.IRow {
	r := make([]types.IRow, len(tbl))
	for i := range tbl {
		r[i] = &tbl[i]
	}
	return r
}

func (tbl MACBindingTable) NewRow() types.IRow {
	return &MACBinding{}
}

func (tbl *MACBindingTable) AppendRow(irow types.IRow) {
	row := irow.(*MACBinding)
	*tbl = append(*tbl, *row)
}

func (tbl MACBindingTable) OvsdbHasIndex() bool {
	return true
}

func (row *MACBinding) MatchByLogicalPortIp(row1 *MACBinding) bool {
	if !types.MatchString(row.LogicalPort, row1.LogicalPort) {
		return false
	}
	if !types.MatchString(row.Ip, row1.Ip) {
		return false
	}
	return true
}

func (tbl MACBindingTable) GetByLogicalPortIp(row1 *MACBinding) *MACBinding {
	for i := range tbl {
		row := &tbl[i]
		if row.MatchByLogicalPortIp(row1) {
			return row
		}
	}
	return nil
}

func (tbl MACBindingTable) OvsdbGetByAnyIndex(irow1 types.IRow) types.IRow {
	row1 := irow1.(*MACBinding)
	if !(types.IsZeroString(row1.LogicalPort) || types.IsZeroString(row1.Ip)) {
		if row := tbl.GetByLogicalPortIp(row1); row != nil {
			return row
		}
	}
	return nil
}

func (tbl MACBindingTable) FindOneMatchNonZeros(row1 *MACBinding) *MACBinding {
	for i := range tbl {
		row := &tbl[i]
		if row.MatchNonZeros(row1) {
			return row
		}
	}
	return nil
}

type MACBinding struct {
	Uuid        string `json:"_uuid"`
	Version     string `json:"_version"`
	Datapath    string `json:"datapath"`
	Ip          string `json:"ip"`
	LogicalPort string `json:"logical_port"`
	Mac         string `json:"mac"`
}

var _ types.IRow = &MACBinding{}

func (row *MACBinding) OvsdbTableName() string {
	return "MAC_Binding"
}

func (row *MACBinding) OvsdbIsRoot() bool {
	return true
}

func (row *MACBinding) OvsdbUuid() string {
	return row.Uuid
}

func (row *MACBinding) OvsdbCmdArgs() []string {
	r := []string{}
	r = append(r, types.OvsdbCmdArgsUuid("datapath", row.Datapath)...)
	r = append(r, types.OvsdbCmdArgsString("ip", row.Ip)...)
	r = append(r, types.OvsdbCmdArgsString("logical_port", row.LogicalPort)...)
	r = append(r, types.OvsdbCmdArgsString("mac", row.Mac)...)
	return r
}

func (row *MACBinding) SetColumn(name string, val interface{}) (err error) {
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
	case "datapath":
		row.Datapath = types.EnsureUuid(val)
	case "ip":
		row.Ip = types.EnsureString(val)
	case "logical_port":
		row.LogicalPort = types.EnsureString(val)
	case "mac":
		row.Mac = types.EnsureString(val)
	default:
		panic(types.ErrUnknownColumn)
	}
	return
}

func (row *MACBinding) MatchNonZeros(row1 *MACBinding) bool {
	if !types.MatchUuidIfNonZero(row.Uuid, row1.Uuid) {
		return false
	}
	if !types.MatchUuidIfNonZero(row.Version, row1.Version) {
		return false
	}
	if !types.MatchUuidIfNonZero(row.Datapath, row1.Datapath) {
		return false
	}
	if !types.MatchStringIfNonZero(row.Ip, row1.Ip) {
		return false
	}
	if !types.MatchStringIfNonZero(row.LogicalPort, row1.LogicalPort) {
		return false
	}
	if !types.MatchStringIfNonZero(row.Mac, row1.Mac) {
		return false
	}
	return true
}

func (row *MACBinding) HasExternalIds() bool {
	return false
}

func (row *MACBinding) SetExternalId(k, v string) {
	panic(errors.Wrap(types.ErrUnknownColumn, "external_ids"))
}

func (row *MACBinding) GetExternalId(k string) (string, bool) {
	panic(errors.Wrap(types.ErrUnknownColumn, "external_ids"))
}

func (row *MACBinding) RemoveExternalId(k string) (string, bool) {
	panic(errors.Wrap(types.ErrUnknownColumn, "external_ids"))
}

type MulticastGroupTable []MulticastGroup

var _ types.ITable = &MulticastGroupTable{}

func (tbl MulticastGroupTable) OvsdbTableName() string {
	return "Multicast_Group"
}

func (tbl MulticastGroupTable) OvsdbIsRoot() bool {
	return true
}

func (tbl MulticastGroupTable) Rows() []types.IRow {
	r := make([]types.IRow, len(tbl))
	for i := range tbl {
		r[i] = &tbl[i]
	}
	return r
}

func (tbl MulticastGroupTable) NewRow() types.IRow {
	return &MulticastGroup{}
}

func (tbl *MulticastGroupTable) AppendRow(irow types.IRow) {
	row := irow.(*MulticastGroup)
	*tbl = append(*tbl, *row)
}

func (tbl MulticastGroupTable) OvsdbHasIndex() bool {
	return true
}

func (row *MulticastGroup) MatchByDatapathTunnelKey(row1 *MulticastGroup) bool {
	if !types.MatchUuid(row.Datapath, row1.Datapath) {
		return false
	}
	if !types.MatchInteger(row.TunnelKey, row1.TunnelKey) {
		return false
	}
	return true
}

func (tbl MulticastGroupTable) GetByDatapathTunnelKey(row1 *MulticastGroup) *MulticastGroup {
	for i := range tbl {
		row := &tbl[i]
		if row.MatchByDatapathTunnelKey(row1) {
			return row
		}
	}
	return nil
}

func (row *MulticastGroup) MatchByDatapathName(row1 *MulticastGroup) bool {
	if !types.MatchUuid(row.Datapath, row1.Datapath) {
		return false
	}
	if !types.MatchString(row.Name, row1.Name) {
		return false
	}
	return true
}

func (tbl MulticastGroupTable) GetByDatapathName(row1 *MulticastGroup) *MulticastGroup {
	for i := range tbl {
		row := &tbl[i]
		if row.MatchByDatapathName(row1) {
			return row
		}
	}
	return nil
}

func (tbl MulticastGroupTable) OvsdbGetByAnyIndex(irow1 types.IRow) types.IRow {
	row1 := irow1.(*MulticastGroup)
	if !(types.IsZeroUuid(row1.Datapath) || types.IsZeroInteger(row1.TunnelKey)) {
		if row := tbl.GetByDatapathTunnelKey(row1); row != nil {
			return row
		}
	}
	if !(types.IsZeroUuid(row1.Datapath) || types.IsZeroString(row1.Name)) {
		if row := tbl.GetByDatapathName(row1); row != nil {
			return row
		}
	}
	return nil
}

func (tbl MulticastGroupTable) FindOneMatchNonZeros(row1 *MulticastGroup) *MulticastGroup {
	for i := range tbl {
		row := &tbl[i]
		if row.MatchNonZeros(row1) {
			return row
		}
	}
	return nil
}

type MulticastGroup struct {
	Uuid      string   `json:"_uuid"`
	Version   string   `json:"_version"`
	Datapath  string   `json:"datapath"`
	Name      string   `json:"name"`
	Ports     []string `json:"ports"`
	TunnelKey int64    `json:"tunnel_key"`
}

var _ types.IRow = &MulticastGroup{}

func (row *MulticastGroup) OvsdbTableName() string {
	return "Multicast_Group"
}

func (row *MulticastGroup) OvsdbIsRoot() bool {
	return true
}

func (row *MulticastGroup) OvsdbUuid() string {
	return row.Uuid
}

func (row *MulticastGroup) OvsdbCmdArgs() []string {
	r := []string{}
	r = append(r, types.OvsdbCmdArgsUuid("datapath", row.Datapath)...)
	r = append(r, types.OvsdbCmdArgsString("name", row.Name)...)
	r = append(r, types.OvsdbCmdArgsUuidMultiples("ports", row.Ports)...)
	r = append(r, types.OvsdbCmdArgsInteger("tunnel_key", row.TunnelKey)...)
	return r
}

func (row *MulticastGroup) SetColumn(name string, val interface{}) (err error) {
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
	case "datapath":
		row.Datapath = types.EnsureUuid(val)
	case "name":
		row.Name = types.EnsureString(val)
	case "ports":
		row.Ports = types.EnsureUuidMultiples(val)
	case "tunnel_key":
		row.TunnelKey = types.EnsureInteger(val)
	default:
		panic(types.ErrUnknownColumn)
	}
	return
}

func (row *MulticastGroup) MatchNonZeros(row1 *MulticastGroup) bool {
	if !types.MatchUuidIfNonZero(row.Uuid, row1.Uuid) {
		return false
	}
	if !types.MatchUuidIfNonZero(row.Version, row1.Version) {
		return false
	}
	if !types.MatchUuidIfNonZero(row.Datapath, row1.Datapath) {
		return false
	}
	if !types.MatchStringIfNonZero(row.Name, row1.Name) {
		return false
	}
	if !types.MatchUuidMultiplesIfNonZero(row.Ports, row1.Ports) {
		return false
	}
	if !types.MatchIntegerIfNonZero(row.TunnelKey, row1.TunnelKey) {
		return false
	}
	return true
}

func (row *MulticastGroup) HasExternalIds() bool {
	return false
}

func (row *MulticastGroup) SetExternalId(k, v string) {
	panic(errors.Wrap(types.ErrUnknownColumn, "external_ids"))
}

func (row *MulticastGroup) GetExternalId(k string) (string, bool) {
	panic(errors.Wrap(types.ErrUnknownColumn, "external_ids"))
}

func (row *MulticastGroup) RemoveExternalId(k string) (string, bool) {
	panic(errors.Wrap(types.ErrUnknownColumn, "external_ids"))
}

type PortBindingTable []PortBinding

var _ types.ITable = &PortBindingTable{}

func (tbl PortBindingTable) OvsdbTableName() string {
	return "Port_Binding"
}

func (tbl PortBindingTable) OvsdbIsRoot() bool {
	return true
}

func (tbl PortBindingTable) Rows() []types.IRow {
	r := make([]types.IRow, len(tbl))
	for i := range tbl {
		r[i] = &tbl[i]
	}
	return r
}

func (tbl PortBindingTable) NewRow() types.IRow {
	return &PortBinding{}
}

func (tbl *PortBindingTable) AppendRow(irow types.IRow) {
	row := irow.(*PortBinding)
	*tbl = append(*tbl, *row)
}

func (tbl PortBindingTable) OvsdbHasIndex() bool {
	return true
}

func (row *PortBinding) MatchByDatapathTunnelKey(row1 *PortBinding) bool {
	if !types.MatchUuid(row.Datapath, row1.Datapath) {
		return false
	}
	if !types.MatchInteger(row.TunnelKey, row1.TunnelKey) {
		return false
	}
	return true
}

func (tbl PortBindingTable) GetByDatapathTunnelKey(row1 *PortBinding) *PortBinding {
	for i := range tbl {
		row := &tbl[i]
		if row.MatchByDatapathTunnelKey(row1) {
			return row
		}
	}
	return nil
}

func (row *PortBinding) MatchByLogicalPort(row1 *PortBinding) bool {
	if !types.MatchString(row.LogicalPort, row1.LogicalPort) {
		return false
	}
	return true
}

func (tbl PortBindingTable) GetByLogicalPort(row1 *PortBinding) *PortBinding {
	for i := range tbl {
		row := &tbl[i]
		if row.MatchByLogicalPort(row1) {
			return row
		}
	}
	return nil
}

func (tbl PortBindingTable) OvsdbGetByAnyIndex(irow1 types.IRow) types.IRow {
	row1 := irow1.(*PortBinding)
	if !(types.IsZeroUuid(row1.Datapath) || types.IsZeroInteger(row1.TunnelKey)) {
		if row := tbl.GetByDatapathTunnelKey(row1); row != nil {
			return row
		}
	}
	if !(types.IsZeroString(row1.LogicalPort)) {
		if row := tbl.GetByLogicalPort(row1); row != nil {
			return row
		}
	}
	return nil
}

func (tbl PortBindingTable) FindOneMatchNonZeros(row1 *PortBinding) *PortBinding {
	for i := range tbl {
		row := &tbl[i]
		if row.MatchNonZeros(row1) {
			return row
		}
	}
	return nil
}

type PortBinding struct {
	Uuid           string            `json:"_uuid"`
	Version        string            `json:"_version"`
	Chassis        *string           `json:"chassis"`
	Datapath       string            `json:"datapath"`
	ExternalIds    map[string]string `json:"external_ids"`
	GatewayChassis []string          `json:"gateway_chassis"`
	LogicalPort    string            `json:"logical_port"`
	Mac            []string          `json:"mac"`
	NatAddresses   []string          `json:"nat_addresses"`
	Options        map[string]string `json:"options"`
	ParentPort     *string           `json:"parent_port"`
	Tag            *int64            `json:"tag"`
	TunnelKey      int64             `json:"tunnel_key"`
	Type           string            `json:"type"`
}

var _ types.IRow = &PortBinding{}

func (row *PortBinding) OvsdbTableName() string {
	return "Port_Binding"
}

func (row *PortBinding) OvsdbIsRoot() bool {
	return true
}

func (row *PortBinding) OvsdbUuid() string {
	return row.Uuid
}

func (row *PortBinding) OvsdbCmdArgs() []string {
	r := []string{}
	r = append(r, types.OvsdbCmdArgsUuidOptional("chassis", row.Chassis)...)
	r = append(r, types.OvsdbCmdArgsUuid("datapath", row.Datapath)...)
	r = append(r, types.OvsdbCmdArgsMapStringString("external_ids", row.ExternalIds)...)
	r = append(r, types.OvsdbCmdArgsUuidMultiples("gateway_chassis", row.GatewayChassis)...)
	r = append(r, types.OvsdbCmdArgsString("logical_port", row.LogicalPort)...)
	r = append(r, types.OvsdbCmdArgsStringMultiples("mac", row.Mac)...)
	r = append(r, types.OvsdbCmdArgsStringMultiples("nat_addresses", row.NatAddresses)...)
	r = append(r, types.OvsdbCmdArgsMapStringString("options", row.Options)...)
	r = append(r, types.OvsdbCmdArgsStringOptional("parent_port", row.ParentPort)...)
	r = append(r, types.OvsdbCmdArgsIntegerOptional("tag", row.Tag)...)
	r = append(r, types.OvsdbCmdArgsInteger("tunnel_key", row.TunnelKey)...)
	r = append(r, types.OvsdbCmdArgsString("type", row.Type)...)
	return r
}

func (row *PortBinding) SetColumn(name string, val interface{}) (err error) {
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
	case "chassis":
		row.Chassis = types.EnsureUuidOptional(val)
	case "datapath":
		row.Datapath = types.EnsureUuid(val)
	case "external_ids":
		row.ExternalIds = types.EnsureMapStringString(val)
	case "gateway_chassis":
		row.GatewayChassis = types.EnsureUuidMultiples(val)
	case "logical_port":
		row.LogicalPort = types.EnsureString(val)
	case "mac":
		row.Mac = types.EnsureStringMultiples(val)
	case "nat_addresses":
		row.NatAddresses = types.EnsureStringMultiples(val)
	case "options":
		row.Options = types.EnsureMapStringString(val)
	case "parent_port":
		row.ParentPort = types.EnsureStringOptional(val)
	case "tag":
		row.Tag = types.EnsureIntegerOptional(val)
	case "tunnel_key":
		row.TunnelKey = types.EnsureInteger(val)
	case "type":
		row.Type = types.EnsureString(val)
	default:
		panic(types.ErrUnknownColumn)
	}
	return
}

func (row *PortBinding) MatchNonZeros(row1 *PortBinding) bool {
	if !types.MatchUuidIfNonZero(row.Uuid, row1.Uuid) {
		return false
	}
	if !types.MatchUuidIfNonZero(row.Version, row1.Version) {
		return false
	}
	if !types.MatchUuidOptionalIfNonZero(row.Chassis, row1.Chassis) {
		return false
	}
	if !types.MatchUuidIfNonZero(row.Datapath, row1.Datapath) {
		return false
	}
	if !types.MatchMapStringStringIfNonZero(row.ExternalIds, row1.ExternalIds) {
		return false
	}
	if !types.MatchUuidMultiplesIfNonZero(row.GatewayChassis, row1.GatewayChassis) {
		return false
	}
	if !types.MatchStringIfNonZero(row.LogicalPort, row1.LogicalPort) {
		return false
	}
	if !types.MatchStringMultiplesIfNonZero(row.Mac, row1.Mac) {
		return false
	}
	if !types.MatchStringMultiplesIfNonZero(row.NatAddresses, row1.NatAddresses) {
		return false
	}
	if !types.MatchMapStringStringIfNonZero(row.Options, row1.Options) {
		return false
	}
	if !types.MatchStringOptionalIfNonZero(row.ParentPort, row1.ParentPort) {
		return false
	}
	if !types.MatchIntegerOptionalIfNonZero(row.Tag, row1.Tag) {
		return false
	}
	if !types.MatchIntegerIfNonZero(row.TunnelKey, row1.TunnelKey) {
		return false
	}
	if !types.MatchStringIfNonZero(row.Type, row1.Type) {
		return false
	}
	return true
}

func (row *PortBinding) HasExternalIds() bool {
	return true
}

func (row *PortBinding) SetExternalId(k, v string) {
	if row.ExternalIds == nil {
		row.ExternalIds = map[string]string{}
	}
	row.ExternalIds[k] = v
}

func (row *PortBinding) GetExternalId(k string) (string, bool) {
	if row.ExternalIds == nil {
		return "", false
	}
	r, ok := row.ExternalIds[k]
	return r, ok
}

func (row *PortBinding) RemoveExternalId(k string) (string, bool) {
	if row.ExternalIds == nil {
		return "", false
	}
	r, ok := row.ExternalIds[k]
	if ok {
		delete(row.ExternalIds, k)
	}
	return r, ok
}

type RBACPermissionTable []RBACPermission

var _ types.ITable = &RBACPermissionTable{}

func (tbl RBACPermissionTable) OvsdbTableName() string {
	return "RBAC_Permission"
}

func (tbl RBACPermissionTable) OvsdbIsRoot() bool {
	return true
}

func (tbl RBACPermissionTable) Rows() []types.IRow {
	r := make([]types.IRow, len(tbl))
	for i := range tbl {
		r[i] = &tbl[i]
	}
	return r
}

func (tbl RBACPermissionTable) NewRow() types.IRow {
	return &RBACPermission{}
}

func (tbl *RBACPermissionTable) AppendRow(irow types.IRow) {
	row := irow.(*RBACPermission)
	*tbl = append(*tbl, *row)
}

func (tbl RBACPermissionTable) OvsdbHasIndex() bool {
	return false
}

func (tbl RBACPermissionTable) OvsdbGetByAnyIndex(irow1 types.IRow) types.IRow {
	return nil
}

func (tbl RBACPermissionTable) FindOneMatchNonZeros(row1 *RBACPermission) *RBACPermission {
	for i := range tbl {
		row := &tbl[i]
		if row.MatchNonZeros(row1) {
			return row
		}
	}
	return nil
}

type RBACPermission struct {
	Uuid          string   `json:"_uuid"`
	Version       string   `json:"_version"`
	Authorization []string `json:"authorization"`
	InsertDelete  bool     `json:"insert_delete"`
	Table         string   `json:"table"`
	Update        []string `json:"update"`
}

var _ types.IRow = &RBACPermission{}

func (row *RBACPermission) OvsdbTableName() string {
	return "RBAC_Permission"
}

func (row *RBACPermission) OvsdbIsRoot() bool {
	return true
}

func (row *RBACPermission) OvsdbUuid() string {
	return row.Uuid
}

func (row *RBACPermission) OvsdbCmdArgs() []string {
	r := []string{}
	r = append(r, types.OvsdbCmdArgsStringMultiples("authorization", row.Authorization)...)
	r = append(r, types.OvsdbCmdArgsBoolean("insert_delete", row.InsertDelete)...)
	r = append(r, types.OvsdbCmdArgsString("table", row.Table)...)
	r = append(r, types.OvsdbCmdArgsStringMultiples("update", row.Update)...)
	return r
}

func (row *RBACPermission) SetColumn(name string, val interface{}) (err error) {
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
	case "authorization":
		row.Authorization = types.EnsureStringMultiples(val)
	case "insert_delete":
		row.InsertDelete = types.EnsureBoolean(val)
	case "table":
		row.Table = types.EnsureString(val)
	case "update":
		row.Update = types.EnsureStringMultiples(val)
	default:
		panic(types.ErrUnknownColumn)
	}
	return
}

func (row *RBACPermission) MatchNonZeros(row1 *RBACPermission) bool {
	if !types.MatchUuidIfNonZero(row.Uuid, row1.Uuid) {
		return false
	}
	if !types.MatchUuidIfNonZero(row.Version, row1.Version) {
		return false
	}
	if !types.MatchStringMultiplesIfNonZero(row.Authorization, row1.Authorization) {
		return false
	}
	if !types.MatchBooleanIfNonZero(row.InsertDelete, row1.InsertDelete) {
		return false
	}
	if !types.MatchStringIfNonZero(row.Table, row1.Table) {
		return false
	}
	if !types.MatchStringMultiplesIfNonZero(row.Update, row1.Update) {
		return false
	}
	return true
}

func (row *RBACPermission) HasExternalIds() bool {
	return false
}

func (row *RBACPermission) SetExternalId(k, v string) {
	panic(errors.Wrap(types.ErrUnknownColumn, "external_ids"))
}

func (row *RBACPermission) GetExternalId(k string) (string, bool) {
	panic(errors.Wrap(types.ErrUnknownColumn, "external_ids"))
}

func (row *RBACPermission) RemoveExternalId(k string) (string, bool) {
	panic(errors.Wrap(types.ErrUnknownColumn, "external_ids"))
}

type RBACRoleTable []RBACRole

var _ types.ITable = &RBACRoleTable{}

func (tbl RBACRoleTable) OvsdbTableName() string {
	return "RBAC_Role"
}

func (tbl RBACRoleTable) OvsdbIsRoot() bool {
	return true
}

func (tbl RBACRoleTable) Rows() []types.IRow {
	r := make([]types.IRow, len(tbl))
	for i := range tbl {
		r[i] = &tbl[i]
	}
	return r
}

func (tbl RBACRoleTable) NewRow() types.IRow {
	return &RBACRole{}
}

func (tbl *RBACRoleTable) AppendRow(irow types.IRow) {
	row := irow.(*RBACRole)
	*tbl = append(*tbl, *row)
}

func (tbl RBACRoleTable) OvsdbHasIndex() bool {
	return false
}

func (tbl RBACRoleTable) OvsdbGetByAnyIndex(irow1 types.IRow) types.IRow {
	return nil
}

func (tbl RBACRoleTable) FindOneMatchNonZeros(row1 *RBACRole) *RBACRole {
	for i := range tbl {
		row := &tbl[i]
		if row.MatchNonZeros(row1) {
			return row
		}
	}
	return nil
}

type RBACRole struct {
	Uuid        string            `json:"_uuid"`
	Version     string            `json:"_version"`
	Name        string            `json:"name"`
	Permissions map[string]string `json:"permissions"`
}

var _ types.IRow = &RBACRole{}

func (row *RBACRole) OvsdbTableName() string {
	return "RBAC_Role"
}

func (row *RBACRole) OvsdbIsRoot() bool {
	return true
}

func (row *RBACRole) OvsdbUuid() string {
	return row.Uuid
}

func (row *RBACRole) OvsdbCmdArgs() []string {
	r := []string{}
	r = append(r, types.OvsdbCmdArgsString("name", row.Name)...)
	r = append(r, types.OvsdbCmdArgsMapStringUuid("permissions", row.Permissions)...)
	return r
}

func (row *RBACRole) SetColumn(name string, val interface{}) (err error) {
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
	case "name":
		row.Name = types.EnsureString(val)
	case "permissions":
		row.Permissions = types.EnsureMapStringUuid(val)
	default:
		panic(types.ErrUnknownColumn)
	}
	return
}

func (row *RBACRole) MatchNonZeros(row1 *RBACRole) bool {
	if !types.MatchUuidIfNonZero(row.Uuid, row1.Uuid) {
		return false
	}
	if !types.MatchUuidIfNonZero(row.Version, row1.Version) {
		return false
	}
	if !types.MatchStringIfNonZero(row.Name, row1.Name) {
		return false
	}
	if !types.MatchMapStringUuidIfNonZero(row.Permissions, row1.Permissions) {
		return false
	}
	return true
}

func (row *RBACRole) HasExternalIds() bool {
	return false
}

func (row *RBACRole) SetExternalId(k, v string) {
	panic(errors.Wrap(types.ErrUnknownColumn, "external_ids"))
}

func (row *RBACRole) GetExternalId(k string) (string, bool) {
	panic(errors.Wrap(types.ErrUnknownColumn, "external_ids"))
}

func (row *RBACRole) RemoveExternalId(k string) (string, bool) {
	panic(errors.Wrap(types.ErrUnknownColumn, "external_ids"))
}

type SBGlobalTable []SBGlobal

var _ types.ITable = &SBGlobalTable{}

func (tbl SBGlobalTable) OvsdbTableName() string {
	return "SB_Global"
}

func (tbl SBGlobalTable) OvsdbIsRoot() bool {
	return true
}

func (tbl SBGlobalTable) Rows() []types.IRow {
	r := make([]types.IRow, len(tbl))
	for i := range tbl {
		r[i] = &tbl[i]
	}
	return r
}

func (tbl SBGlobalTable) NewRow() types.IRow {
	return &SBGlobal{}
}

func (tbl *SBGlobalTable) AppendRow(irow types.IRow) {
	row := irow.(*SBGlobal)
	*tbl = append(*tbl, *row)
}

func (tbl SBGlobalTable) OvsdbHasIndex() bool {
	return false
}

func (tbl SBGlobalTable) OvsdbGetByAnyIndex(irow1 types.IRow) types.IRow {
	return nil
}

func (tbl SBGlobalTable) FindOneMatchNonZeros(row1 *SBGlobal) *SBGlobal {
	for i := range tbl {
		row := &tbl[i]
		if row.MatchNonZeros(row1) {
			return row
		}
	}
	return nil
}

type SBGlobal struct {
	Uuid        string            `json:"_uuid"`
	Version     string            `json:"_version"`
	Connections []string          `json:"connections"`
	ExternalIds map[string]string `json:"external_ids"`
	NbCfg       int64             `json:"nb_cfg"`
	Ssl         *string           `json:"ssl"`
}

var _ types.IRow = &SBGlobal{}

func (row *SBGlobal) OvsdbTableName() string {
	return "SB_Global"
}

func (row *SBGlobal) OvsdbIsRoot() bool {
	return true
}

func (row *SBGlobal) OvsdbUuid() string {
	return row.Uuid
}

func (row *SBGlobal) OvsdbCmdArgs() []string {
	r := []string{}
	r = append(r, types.OvsdbCmdArgsUuidMultiples("connections", row.Connections)...)
	r = append(r, types.OvsdbCmdArgsMapStringString("external_ids", row.ExternalIds)...)
	r = append(r, types.OvsdbCmdArgsInteger("nb_cfg", row.NbCfg)...)
	r = append(r, types.OvsdbCmdArgsUuidOptional("ssl", row.Ssl)...)
	return r
}

func (row *SBGlobal) SetColumn(name string, val interface{}) (err error) {
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
	case "connections":
		row.Connections = types.EnsureUuidMultiples(val)
	case "external_ids":
		row.ExternalIds = types.EnsureMapStringString(val)
	case "nb_cfg":
		row.NbCfg = types.EnsureInteger(val)
	case "ssl":
		row.Ssl = types.EnsureUuidOptional(val)
	default:
		panic(types.ErrUnknownColumn)
	}
	return
}

func (row *SBGlobal) MatchNonZeros(row1 *SBGlobal) bool {
	if !types.MatchUuidIfNonZero(row.Uuid, row1.Uuid) {
		return false
	}
	if !types.MatchUuidIfNonZero(row.Version, row1.Version) {
		return false
	}
	if !types.MatchUuidMultiplesIfNonZero(row.Connections, row1.Connections) {
		return false
	}
	if !types.MatchMapStringStringIfNonZero(row.ExternalIds, row1.ExternalIds) {
		return false
	}
	if !types.MatchIntegerIfNonZero(row.NbCfg, row1.NbCfg) {
		return false
	}
	if !types.MatchUuidOptionalIfNonZero(row.Ssl, row1.Ssl) {
		return false
	}
	return true
}

func (row *SBGlobal) HasExternalIds() bool {
	return true
}

func (row *SBGlobal) SetExternalId(k, v string) {
	if row.ExternalIds == nil {
		row.ExternalIds = map[string]string{}
	}
	row.ExternalIds[k] = v
}

func (row *SBGlobal) GetExternalId(k string) (string, bool) {
	if row.ExternalIds == nil {
		return "", false
	}
	r, ok := row.ExternalIds[k]
	return r, ok
}

func (row *SBGlobal) RemoveExternalId(k string) (string, bool) {
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
	SslCiphers      string            `json:"ssl_ciphers"`
	SslProtocols    string            `json:"ssl_protocols"`
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
	r = append(r, types.OvsdbCmdArgsString("ssl_ciphers", row.SslCiphers)...)
	r = append(r, types.OvsdbCmdArgsString("ssl_protocols", row.SslProtocols)...)
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
	case "ssl_ciphers":
		row.SslCiphers = types.EnsureString(val)
	case "ssl_protocols":
		row.SslProtocols = types.EnsureString(val)
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
	if !types.MatchStringIfNonZero(row.SslCiphers, row1.SslCiphers) {
		return false
	}
	if !types.MatchStringIfNonZero(row.SslProtocols, row1.SslProtocols) {
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
