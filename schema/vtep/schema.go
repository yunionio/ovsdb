// DO NOT EDIT: automatically generated code

package vtep

import (
	"fmt"

	"github.com/pkg/errors"

	"yunion.io/x/ovsdb/types"
)

type HardwareVtep struct {
	ACL                 ACLTable
	ACLEntry            ACLEntryTable
	ArpSourcesLocal     ArpSourcesLocalTable
	ArpSourcesRemote    ArpSourcesRemoteTable
	Global              GlobalTable
	LogicalBindingStats LogicalBindingStatsTable
	LogicalRouter       LogicalRouterTable
	LogicalSwitch       LogicalSwitchTable
	Manager             ManagerTable
	McastMacsLocal      McastMacsLocalTable
	McastMacsRemote     McastMacsRemoteTable
	PhysicalLocator     PhysicalLocatorTable
	PhysicalLocatorSet  PhysicalLocatorSetTable
	PhysicalPort        PhysicalPortTable
	PhysicalSwitch      PhysicalSwitchTable
	Tunnel              TunnelTable
	UcastMacsLocal      UcastMacsLocalTable
	UcastMacsRemote     UcastMacsRemoteTable
}

var _ types.IDatabase = &HardwareVtep{}

func (db HardwareVtep) FindOneMatchNonZeros(irow types.IRow) types.IRow {
	switch row := irow.(type) {
	case *ACL:
		if r := db.ACL.FindOneMatchNonZeros(row); r != nil {
			return r
		}
		return nil
	case *ACLEntry:
		if r := db.ACLEntry.FindOneMatchNonZeros(row); r != nil {
			return r
		}
		return nil
	case *ArpSourcesLocal:
		if r := db.ArpSourcesLocal.FindOneMatchNonZeros(row); r != nil {
			return r
		}
		return nil
	case *ArpSourcesRemote:
		if r := db.ArpSourcesRemote.FindOneMatchNonZeros(row); r != nil {
			return r
		}
		return nil
	case *Global:
		if r := db.Global.FindOneMatchNonZeros(row); r != nil {
			return r
		}
		return nil
	case *LogicalBindingStats:
		if r := db.LogicalBindingStats.FindOneMatchNonZeros(row); r != nil {
			return r
		}
		return nil
	case *LogicalRouter:
		if r := db.LogicalRouter.FindOneMatchNonZeros(row); r != nil {
			return r
		}
		return nil
	case *LogicalSwitch:
		if r := db.LogicalSwitch.FindOneMatchNonZeros(row); r != nil {
			return r
		}
		return nil
	case *Manager:
		if r := db.Manager.FindOneMatchNonZeros(row); r != nil {
			return r
		}
		return nil
	case *McastMacsLocal:
		if r := db.McastMacsLocal.FindOneMatchNonZeros(row); r != nil {
			return r
		}
		return nil
	case *McastMacsRemote:
		if r := db.McastMacsRemote.FindOneMatchNonZeros(row); r != nil {
			return r
		}
		return nil
	case *PhysicalLocator:
		if r := db.PhysicalLocator.FindOneMatchNonZeros(row); r != nil {
			return r
		}
		return nil
	case *PhysicalLocatorSet:
		if r := db.PhysicalLocatorSet.FindOneMatchNonZeros(row); r != nil {
			return r
		}
		return nil
	case *PhysicalPort:
		if r := db.PhysicalPort.FindOneMatchNonZeros(row); r != nil {
			return r
		}
		return nil
	case *PhysicalSwitch:
		if r := db.PhysicalSwitch.FindOneMatchNonZeros(row); r != nil {
			return r
		}
		return nil
	case *Tunnel:
		if r := db.Tunnel.FindOneMatchNonZeros(row); r != nil {
			return r
		}
		return nil
	case *UcastMacsLocal:
		if r := db.UcastMacsLocal.FindOneMatchNonZeros(row); r != nil {
			return r
		}
		return nil
	case *UcastMacsRemote:
		if r := db.UcastMacsRemote.FindOneMatchNonZeros(row); r != nil {
			return r
		}
		return nil
	}
	panic(types.ErrBadType)
}

func (db HardwareVtep) FindOneMatchByAnyIndex(irow types.IRow) types.IRow {
	switch row := irow.(type) {
	case *ACL:
		if r := db.ACL.OvsdbGetByAnyIndex(row); r != nil {
			return r
		}
		return nil
	case *ACLEntry:
		if r := db.ACLEntry.OvsdbGetByAnyIndex(row); r != nil {
			return r
		}
		return nil
	case *ArpSourcesLocal:
		if r := db.ArpSourcesLocal.OvsdbGetByAnyIndex(row); r != nil {
			return r
		}
		return nil
	case *ArpSourcesRemote:
		if r := db.ArpSourcesRemote.OvsdbGetByAnyIndex(row); r != nil {
			return r
		}
		return nil
	case *Global:
		if r := db.Global.OvsdbGetByAnyIndex(row); r != nil {
			return r
		}
		return nil
	case *LogicalBindingStats:
		if r := db.LogicalBindingStats.OvsdbGetByAnyIndex(row); r != nil {
			return r
		}
		return nil
	case *LogicalRouter:
		if r := db.LogicalRouter.OvsdbGetByAnyIndex(row); r != nil {
			return r
		}
		return nil
	case *LogicalSwitch:
		if r := db.LogicalSwitch.OvsdbGetByAnyIndex(row); r != nil {
			return r
		}
		return nil
	case *Manager:
		if r := db.Manager.OvsdbGetByAnyIndex(row); r != nil {
			return r
		}
		return nil
	case *McastMacsLocal:
		if r := db.McastMacsLocal.OvsdbGetByAnyIndex(row); r != nil {
			return r
		}
		return nil
	case *McastMacsRemote:
		if r := db.McastMacsRemote.OvsdbGetByAnyIndex(row); r != nil {
			return r
		}
		return nil
	case *PhysicalLocator:
		if r := db.PhysicalLocator.OvsdbGetByAnyIndex(row); r != nil {
			return r
		}
		return nil
	case *PhysicalLocatorSet:
		if r := db.PhysicalLocatorSet.OvsdbGetByAnyIndex(row); r != nil {
			return r
		}
		return nil
	case *PhysicalPort:
		if r := db.PhysicalPort.OvsdbGetByAnyIndex(row); r != nil {
			return r
		}
		return nil
	case *PhysicalSwitch:
		if r := db.PhysicalSwitch.OvsdbGetByAnyIndex(row); r != nil {
			return r
		}
		return nil
	case *Tunnel:
		if r := db.Tunnel.OvsdbGetByAnyIndex(row); r != nil {
			return r
		}
		return nil
	case *UcastMacsLocal:
		if r := db.UcastMacsLocal.OvsdbGetByAnyIndex(row); r != nil {
			return r
		}
		return nil
	case *UcastMacsRemote:
		if r := db.UcastMacsRemote.OvsdbGetByAnyIndex(row); r != nil {
			return r
		}
		return nil
	}
	panic(types.ErrBadType)
}

type ACLTable []ACL

var _ types.ITable = &ACLTable{}

func (tbl ACLTable) OvsdbTableName() string {
	return "ACL"
}

func (tbl ACLTable) OvsdbIsRoot() bool {
	return true
}

func (tbl ACLTable) Rows() []types.IRow {
	r := make([]types.IRow, len(tbl))
	for i := range tbl {
		r[i] = &tbl[i]
	}
	return r
}

func (tbl ACLTable) NewRow() types.IRow {
	return &ACL{}
}

func (tbl *ACLTable) AppendRow(irow types.IRow) {
	row := irow.(*ACL)
	*tbl = append(*tbl, *row)
}

func (tbl ACLTable) OvsdbHasIndex() bool {
	return true
}

func (row *ACL) MatchByAclName(row1 *ACL) bool {
	if !types.MatchString(row.AclName, row1.AclName) {
		return false
	}
	return true
}

func (tbl ACLTable) GetByAclName(row1 *ACL) *ACL {
	for i := range tbl {
		row := &tbl[i]
		if row.MatchByAclName(row1) {
			return row
		}
	}
	return nil
}

func (tbl ACLTable) OvsdbGetByAnyIndex(irow1 types.IRow) types.IRow {
	row1 := irow1.(*ACL)
	if !(types.IsZeroString(row1.AclName)) {
		if row := tbl.GetByAclName(row1); row != nil {
			return row
		}
	}
	return nil
}

func (tbl ACLTable) FindOneMatchNonZeros(row1 *ACL) *ACL {
	for i := range tbl {
		row := &tbl[i]
		if row.MatchNonZeros(row1) {
			return row
		}
	}
	return nil
}

type ACL struct {
	Uuid           string   `json:"_uuid"`
	Version        string   `json:"_version"`
	AclEntries     []string `json:"acl_entries"`
	AclFaultStatus []string `json:"acl_fault_status"`
	AclName        string   `json:"acl_name"`
}

var _ types.IRow = &ACL{}

func (row *ACL) OvsdbTableName() string {
	return "ACL"
}

func (row *ACL) OvsdbIsRoot() bool {
	return true
}

func (row *ACL) OvsdbUuid() string {
	return row.Uuid
}

func (row *ACL) OvsdbCmdArgs() []string {
	r := []string{}
	r = append(r, types.OvsdbCmdArgsUuidMultiples("acl_entries", row.AclEntries)...)
	r = append(r, types.OvsdbCmdArgsStringMultiples("acl_fault_status", row.AclFaultStatus)...)
	r = append(r, types.OvsdbCmdArgsString("acl_name", row.AclName)...)
	return r
}

func (row *ACL) SetColumn(name string, val interface{}) (err error) {
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
	case "acl_entries":
		row.AclEntries = types.EnsureUuidMultiples(val)
	case "acl_fault_status":
		row.AclFaultStatus = types.EnsureStringMultiples(val)
	case "acl_name":
		row.AclName = types.EnsureString(val)
	default:
		panic(types.ErrUnknownColumn)
	}
	return
}

func (row *ACL) MatchNonZeros(row1 *ACL) bool {
	if !types.MatchUuidIfNonZero(row.Uuid, row1.Uuid) {
		return false
	}
	if !types.MatchUuidIfNonZero(row.Version, row1.Version) {
		return false
	}
	if !types.MatchUuidMultiplesIfNonZero(row.AclEntries, row1.AclEntries) {
		return false
	}
	if !types.MatchStringMultiplesIfNonZero(row.AclFaultStatus, row1.AclFaultStatus) {
		return false
	}
	if !types.MatchStringIfNonZero(row.AclName, row1.AclName) {
		return false
	}
	return true
}

func (tbl ACLTable) FindACLEntryReferrer_acl_entries(refUuid string) (r []*ACL) {
	for i := range tbl {
		row := &tbl[i]
		for _, val := range row.AclEntries {
			if val == refUuid {
				r = append(r, row)
			}
		}
	}
	return r
}

func (row *ACL) HasExternalIds() bool {
	return false
}

func (row *ACL) SetExternalId(k, v string) {
	panic(errors.Wrap(types.ErrUnknownColumn, "external_ids"))
}

func (row *ACL) GetExternalId(k string) (string, bool) {
	panic(errors.Wrap(types.ErrUnknownColumn, "external_ids"))
}

func (row *ACL) RemoveExternalId(k string) (string, bool) {
	panic(errors.Wrap(types.ErrUnknownColumn, "external_ids"))
}

type ACLEntryTable []ACLEntry

var _ types.ITable = &ACLEntryTable{}

func (tbl ACLEntryTable) OvsdbTableName() string {
	return "ACL_entry"
}

func (tbl ACLEntryTable) OvsdbIsRoot() bool {
	return true
}

func (tbl ACLEntryTable) Rows() []types.IRow {
	r := make([]types.IRow, len(tbl))
	for i := range tbl {
		r[i] = &tbl[i]
	}
	return r
}

func (tbl ACLEntryTable) NewRow() types.IRow {
	return &ACLEntry{}
}

func (tbl *ACLEntryTable) AppendRow(irow types.IRow) {
	row := irow.(*ACLEntry)
	*tbl = append(*tbl, *row)
}

func (tbl ACLEntryTable) OvsdbHasIndex() bool {
	return false
}

func (tbl ACLEntryTable) OvsdbGetByAnyIndex(irow1 types.IRow) types.IRow {
	return nil
}

func (tbl ACLEntryTable) FindOneMatchNonZeros(row1 *ACLEntry) *ACLEntry {
	for i := range tbl {
		row := &tbl[i]
		if row.MatchNonZeros(row1) {
			return row
		}
	}
	return nil
}

type ACLEntry struct {
	Uuid            string   `json:"_uuid"`
	Version         string   `json:"_version"`
	AcleFaultStatus []string `json:"acle_fault_status"`
	Action          string   `json:"action"`
	DestIp          *string  `json:"dest_ip"`
	DestMac         *string  `json:"dest_mac"`
	DestMask        *string  `json:"dest_mask"`
	DestPortMax     *int64   `json:"dest_port_max"`
	DestPortMin     *int64   `json:"dest_port_min"`
	Direction       string   `json:"direction"`
	Ethertype       *string  `json:"ethertype"`
	IcmpCode        *int64   `json:"icmp_code"`
	IcmpType        *int64   `json:"icmp_type"`
	Protocol        *int64   `json:"protocol"`
	Sequence        int64    `json:"sequence"`
	SourceIp        *string  `json:"source_ip"`
	SourceMac       *string  `json:"source_mac"`
	SourceMask      *string  `json:"source_mask"`
	SourcePortMax   *int64   `json:"source_port_max"`
	SourcePortMin   *int64   `json:"source_port_min"`
	TcpFlags        *int64   `json:"tcp_flags"`
	TcpFlagsMask    *int64   `json:"tcp_flags_mask"`
}

var _ types.IRow = &ACLEntry{}

func (row *ACLEntry) OvsdbTableName() string {
	return "ACL_entry"
}

func (row *ACLEntry) OvsdbIsRoot() bool {
	return true
}

func (row *ACLEntry) OvsdbUuid() string {
	return row.Uuid
}

func (row *ACLEntry) OvsdbCmdArgs() []string {
	r := []string{}
	r = append(r, types.OvsdbCmdArgsStringMultiples("acle_fault_status", row.AcleFaultStatus)...)
	r = append(r, types.OvsdbCmdArgsString("action", row.Action)...)
	r = append(r, types.OvsdbCmdArgsStringOptional("dest_ip", row.DestIp)...)
	r = append(r, types.OvsdbCmdArgsStringOptional("dest_mac", row.DestMac)...)
	r = append(r, types.OvsdbCmdArgsStringOptional("dest_mask", row.DestMask)...)
	r = append(r, types.OvsdbCmdArgsIntegerOptional("dest_port_max", row.DestPortMax)...)
	r = append(r, types.OvsdbCmdArgsIntegerOptional("dest_port_min", row.DestPortMin)...)
	r = append(r, types.OvsdbCmdArgsString("direction", row.Direction)...)
	r = append(r, types.OvsdbCmdArgsStringOptional("ethertype", row.Ethertype)...)
	r = append(r, types.OvsdbCmdArgsIntegerOptional("icmp_code", row.IcmpCode)...)
	r = append(r, types.OvsdbCmdArgsIntegerOptional("icmp_type", row.IcmpType)...)
	r = append(r, types.OvsdbCmdArgsIntegerOptional("protocol", row.Protocol)...)
	r = append(r, types.OvsdbCmdArgsInteger("sequence", row.Sequence)...)
	r = append(r, types.OvsdbCmdArgsStringOptional("source_ip", row.SourceIp)...)
	r = append(r, types.OvsdbCmdArgsStringOptional("source_mac", row.SourceMac)...)
	r = append(r, types.OvsdbCmdArgsStringOptional("source_mask", row.SourceMask)...)
	r = append(r, types.OvsdbCmdArgsIntegerOptional("source_port_max", row.SourcePortMax)...)
	r = append(r, types.OvsdbCmdArgsIntegerOptional("source_port_min", row.SourcePortMin)...)
	r = append(r, types.OvsdbCmdArgsIntegerOptional("tcp_flags", row.TcpFlags)...)
	r = append(r, types.OvsdbCmdArgsIntegerOptional("tcp_flags_mask", row.TcpFlagsMask)...)
	return r
}

func (row *ACLEntry) SetColumn(name string, val interface{}) (err error) {
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
	case "acle_fault_status":
		row.AcleFaultStatus = types.EnsureStringMultiples(val)
	case "action":
		row.Action = types.EnsureString(val)
	case "dest_ip":
		row.DestIp = types.EnsureStringOptional(val)
	case "dest_mac":
		row.DestMac = types.EnsureStringOptional(val)
	case "dest_mask":
		row.DestMask = types.EnsureStringOptional(val)
	case "dest_port_max":
		row.DestPortMax = types.EnsureIntegerOptional(val)
	case "dest_port_min":
		row.DestPortMin = types.EnsureIntegerOptional(val)
	case "direction":
		row.Direction = types.EnsureString(val)
	case "ethertype":
		row.Ethertype = types.EnsureStringOptional(val)
	case "icmp_code":
		row.IcmpCode = types.EnsureIntegerOptional(val)
	case "icmp_type":
		row.IcmpType = types.EnsureIntegerOptional(val)
	case "protocol":
		row.Protocol = types.EnsureIntegerOptional(val)
	case "sequence":
		row.Sequence = types.EnsureInteger(val)
	case "source_ip":
		row.SourceIp = types.EnsureStringOptional(val)
	case "source_mac":
		row.SourceMac = types.EnsureStringOptional(val)
	case "source_mask":
		row.SourceMask = types.EnsureStringOptional(val)
	case "source_port_max":
		row.SourcePortMax = types.EnsureIntegerOptional(val)
	case "source_port_min":
		row.SourcePortMin = types.EnsureIntegerOptional(val)
	case "tcp_flags":
		row.TcpFlags = types.EnsureIntegerOptional(val)
	case "tcp_flags_mask":
		row.TcpFlagsMask = types.EnsureIntegerOptional(val)
	default:
		panic(types.ErrUnknownColumn)
	}
	return
}

func (row *ACLEntry) MatchNonZeros(row1 *ACLEntry) bool {
	if !types.MatchUuidIfNonZero(row.Uuid, row1.Uuid) {
		return false
	}
	if !types.MatchUuidIfNonZero(row.Version, row1.Version) {
		return false
	}
	if !types.MatchStringMultiplesIfNonZero(row.AcleFaultStatus, row1.AcleFaultStatus) {
		return false
	}
	if !types.MatchStringIfNonZero(row.Action, row1.Action) {
		return false
	}
	if !types.MatchStringOptionalIfNonZero(row.DestIp, row1.DestIp) {
		return false
	}
	if !types.MatchStringOptionalIfNonZero(row.DestMac, row1.DestMac) {
		return false
	}
	if !types.MatchStringOptionalIfNonZero(row.DestMask, row1.DestMask) {
		return false
	}
	if !types.MatchIntegerOptionalIfNonZero(row.DestPortMax, row1.DestPortMax) {
		return false
	}
	if !types.MatchIntegerOptionalIfNonZero(row.DestPortMin, row1.DestPortMin) {
		return false
	}
	if !types.MatchStringIfNonZero(row.Direction, row1.Direction) {
		return false
	}
	if !types.MatchStringOptionalIfNonZero(row.Ethertype, row1.Ethertype) {
		return false
	}
	if !types.MatchIntegerOptionalIfNonZero(row.IcmpCode, row1.IcmpCode) {
		return false
	}
	if !types.MatchIntegerOptionalIfNonZero(row.IcmpType, row1.IcmpType) {
		return false
	}
	if !types.MatchIntegerOptionalIfNonZero(row.Protocol, row1.Protocol) {
		return false
	}
	if !types.MatchIntegerIfNonZero(row.Sequence, row1.Sequence) {
		return false
	}
	if !types.MatchStringOptionalIfNonZero(row.SourceIp, row1.SourceIp) {
		return false
	}
	if !types.MatchStringOptionalIfNonZero(row.SourceMac, row1.SourceMac) {
		return false
	}
	if !types.MatchStringOptionalIfNonZero(row.SourceMask, row1.SourceMask) {
		return false
	}
	if !types.MatchIntegerOptionalIfNonZero(row.SourcePortMax, row1.SourcePortMax) {
		return false
	}
	if !types.MatchIntegerOptionalIfNonZero(row.SourcePortMin, row1.SourcePortMin) {
		return false
	}
	if !types.MatchIntegerOptionalIfNonZero(row.TcpFlags, row1.TcpFlags) {
		return false
	}
	if !types.MatchIntegerOptionalIfNonZero(row.TcpFlagsMask, row1.TcpFlagsMask) {
		return false
	}
	return true
}

func (row *ACLEntry) HasExternalIds() bool {
	return false
}

func (row *ACLEntry) SetExternalId(k, v string) {
	panic(errors.Wrap(types.ErrUnknownColumn, "external_ids"))
}

func (row *ACLEntry) GetExternalId(k string) (string, bool) {
	panic(errors.Wrap(types.ErrUnknownColumn, "external_ids"))
}

func (row *ACLEntry) RemoveExternalId(k string) (string, bool) {
	panic(errors.Wrap(types.ErrUnknownColumn, "external_ids"))
}

type ArpSourcesLocalTable []ArpSourcesLocal

var _ types.ITable = &ArpSourcesLocalTable{}

func (tbl ArpSourcesLocalTable) OvsdbTableName() string {
	return "Arp_Sources_Local"
}

func (tbl ArpSourcesLocalTable) OvsdbIsRoot() bool {
	return true
}

func (tbl ArpSourcesLocalTable) Rows() []types.IRow {
	r := make([]types.IRow, len(tbl))
	for i := range tbl {
		r[i] = &tbl[i]
	}
	return r
}

func (tbl ArpSourcesLocalTable) NewRow() types.IRow {
	return &ArpSourcesLocal{}
}

func (tbl *ArpSourcesLocalTable) AppendRow(irow types.IRow) {
	row := irow.(*ArpSourcesLocal)
	*tbl = append(*tbl, *row)
}

func (tbl ArpSourcesLocalTable) OvsdbHasIndex() bool {
	return false
}

func (tbl ArpSourcesLocalTable) OvsdbGetByAnyIndex(irow1 types.IRow) types.IRow {
	return nil
}

func (tbl ArpSourcesLocalTable) FindOneMatchNonZeros(row1 *ArpSourcesLocal) *ArpSourcesLocal {
	for i := range tbl {
		row := &tbl[i]
		if row.MatchNonZeros(row1) {
			return row
		}
	}
	return nil
}

type ArpSourcesLocal struct {
	Uuid    string `json:"_uuid"`
	Version string `json:"_version"`
	Locator string `json:"locator"`
	SrcMac  string `json:"src_mac"`
}

var _ types.IRow = &ArpSourcesLocal{}

func (row *ArpSourcesLocal) OvsdbTableName() string {
	return "Arp_Sources_Local"
}

func (row *ArpSourcesLocal) OvsdbIsRoot() bool {
	return true
}

func (row *ArpSourcesLocal) OvsdbUuid() string {
	return row.Uuid
}

func (row *ArpSourcesLocal) OvsdbCmdArgs() []string {
	r := []string{}
	r = append(r, types.OvsdbCmdArgsUuid("locator", row.Locator)...)
	r = append(r, types.OvsdbCmdArgsString("src_mac", row.SrcMac)...)
	return r
}

func (row *ArpSourcesLocal) SetColumn(name string, val interface{}) (err error) {
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
	case "locator":
		row.Locator = types.EnsureUuid(val)
	case "src_mac":
		row.SrcMac = types.EnsureString(val)
	default:
		panic(types.ErrUnknownColumn)
	}
	return
}

func (row *ArpSourcesLocal) MatchNonZeros(row1 *ArpSourcesLocal) bool {
	if !types.MatchUuidIfNonZero(row.Uuid, row1.Uuid) {
		return false
	}
	if !types.MatchUuidIfNonZero(row.Version, row1.Version) {
		return false
	}
	if !types.MatchUuidIfNonZero(row.Locator, row1.Locator) {
		return false
	}
	if !types.MatchStringIfNonZero(row.SrcMac, row1.SrcMac) {
		return false
	}
	return true
}

func (tbl ArpSourcesLocalTable) FindPhysicalLocatorReferrer_locator(refUuid string) (r []*ArpSourcesLocal) {
	for i := range tbl {
		row := &tbl[i]
		if row.Locator == refUuid {
			r = append(r, row)
		}
	}
	return r
}

func (row *ArpSourcesLocal) HasExternalIds() bool {
	return false
}

func (row *ArpSourcesLocal) SetExternalId(k, v string) {
	panic(errors.Wrap(types.ErrUnknownColumn, "external_ids"))
}

func (row *ArpSourcesLocal) GetExternalId(k string) (string, bool) {
	panic(errors.Wrap(types.ErrUnknownColumn, "external_ids"))
}

func (row *ArpSourcesLocal) RemoveExternalId(k string) (string, bool) {
	panic(errors.Wrap(types.ErrUnknownColumn, "external_ids"))
}

type ArpSourcesRemoteTable []ArpSourcesRemote

var _ types.ITable = &ArpSourcesRemoteTable{}

func (tbl ArpSourcesRemoteTable) OvsdbTableName() string {
	return "Arp_Sources_Remote"
}

func (tbl ArpSourcesRemoteTable) OvsdbIsRoot() bool {
	return true
}

func (tbl ArpSourcesRemoteTable) Rows() []types.IRow {
	r := make([]types.IRow, len(tbl))
	for i := range tbl {
		r[i] = &tbl[i]
	}
	return r
}

func (tbl ArpSourcesRemoteTable) NewRow() types.IRow {
	return &ArpSourcesRemote{}
}

func (tbl *ArpSourcesRemoteTable) AppendRow(irow types.IRow) {
	row := irow.(*ArpSourcesRemote)
	*tbl = append(*tbl, *row)
}

func (tbl ArpSourcesRemoteTable) OvsdbHasIndex() bool {
	return false
}

func (tbl ArpSourcesRemoteTable) OvsdbGetByAnyIndex(irow1 types.IRow) types.IRow {
	return nil
}

func (tbl ArpSourcesRemoteTable) FindOneMatchNonZeros(row1 *ArpSourcesRemote) *ArpSourcesRemote {
	for i := range tbl {
		row := &tbl[i]
		if row.MatchNonZeros(row1) {
			return row
		}
	}
	return nil
}

type ArpSourcesRemote struct {
	Uuid    string `json:"_uuid"`
	Version string `json:"_version"`
	Locator string `json:"locator"`
	SrcMac  string `json:"src_mac"`
}

var _ types.IRow = &ArpSourcesRemote{}

func (row *ArpSourcesRemote) OvsdbTableName() string {
	return "Arp_Sources_Remote"
}

func (row *ArpSourcesRemote) OvsdbIsRoot() bool {
	return true
}

func (row *ArpSourcesRemote) OvsdbUuid() string {
	return row.Uuid
}

func (row *ArpSourcesRemote) OvsdbCmdArgs() []string {
	r := []string{}
	r = append(r, types.OvsdbCmdArgsUuid("locator", row.Locator)...)
	r = append(r, types.OvsdbCmdArgsString("src_mac", row.SrcMac)...)
	return r
}

func (row *ArpSourcesRemote) SetColumn(name string, val interface{}) (err error) {
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
	case "locator":
		row.Locator = types.EnsureUuid(val)
	case "src_mac":
		row.SrcMac = types.EnsureString(val)
	default:
		panic(types.ErrUnknownColumn)
	}
	return
}

func (row *ArpSourcesRemote) MatchNonZeros(row1 *ArpSourcesRemote) bool {
	if !types.MatchUuidIfNonZero(row.Uuid, row1.Uuid) {
		return false
	}
	if !types.MatchUuidIfNonZero(row.Version, row1.Version) {
		return false
	}
	if !types.MatchUuidIfNonZero(row.Locator, row1.Locator) {
		return false
	}
	if !types.MatchStringIfNonZero(row.SrcMac, row1.SrcMac) {
		return false
	}
	return true
}

func (tbl ArpSourcesRemoteTable) FindPhysicalLocatorReferrer_locator(refUuid string) (r []*ArpSourcesRemote) {
	for i := range tbl {
		row := &tbl[i]
		if row.Locator == refUuid {
			r = append(r, row)
		}
	}
	return r
}

func (row *ArpSourcesRemote) HasExternalIds() bool {
	return false
}

func (row *ArpSourcesRemote) SetExternalId(k, v string) {
	panic(errors.Wrap(types.ErrUnknownColumn, "external_ids"))
}

func (row *ArpSourcesRemote) GetExternalId(k string) (string, bool) {
	panic(errors.Wrap(types.ErrUnknownColumn, "external_ids"))
}

func (row *ArpSourcesRemote) RemoveExternalId(k string) (string, bool) {
	panic(errors.Wrap(types.ErrUnknownColumn, "external_ids"))
}

type GlobalTable []Global

var _ types.ITable = &GlobalTable{}

func (tbl GlobalTable) OvsdbTableName() string {
	return "Global"
}

func (tbl GlobalTable) OvsdbIsRoot() bool {
	return true
}

func (tbl GlobalTable) Rows() []types.IRow {
	r := make([]types.IRow, len(tbl))
	for i := range tbl {
		r[i] = &tbl[i]
	}
	return r
}

func (tbl GlobalTable) NewRow() types.IRow {
	return &Global{}
}

func (tbl *GlobalTable) AppendRow(irow types.IRow) {
	row := irow.(*Global)
	*tbl = append(*tbl, *row)
}

func (tbl GlobalTable) OvsdbHasIndex() bool {
	return false
}

func (tbl GlobalTable) OvsdbGetByAnyIndex(irow1 types.IRow) types.IRow {
	return nil
}

func (tbl GlobalTable) FindOneMatchNonZeros(row1 *Global) *Global {
	for i := range tbl {
		row := &tbl[i]
		if row.MatchNonZeros(row1) {
			return row
		}
	}
	return nil
}

type Global struct {
	Uuid        string            `json:"_uuid"`
	Version     string            `json:"_version"`
	Managers    []string          `json:"managers"`
	OtherConfig map[string]string `json:"other_config"`
	Switches    []string          `json:"switches"`
}

var _ types.IRow = &Global{}

func (row *Global) OvsdbTableName() string {
	return "Global"
}

func (row *Global) OvsdbIsRoot() bool {
	return true
}

func (row *Global) OvsdbUuid() string {
	return row.Uuid
}

func (row *Global) OvsdbCmdArgs() []string {
	r := []string{}
	r = append(r, types.OvsdbCmdArgsUuidMultiples("managers", row.Managers)...)
	r = append(r, types.OvsdbCmdArgsMapStringString("other_config", row.OtherConfig)...)
	r = append(r, types.OvsdbCmdArgsUuidMultiples("switches", row.Switches)...)
	return r
}

func (row *Global) SetColumn(name string, val interface{}) (err error) {
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
	case "managers":
		row.Managers = types.EnsureUuidMultiples(val)
	case "other_config":
		row.OtherConfig = types.EnsureMapStringString(val)
	case "switches":
		row.Switches = types.EnsureUuidMultiples(val)
	default:
		panic(types.ErrUnknownColumn)
	}
	return
}

func (row *Global) MatchNonZeros(row1 *Global) bool {
	if !types.MatchUuidIfNonZero(row.Uuid, row1.Uuid) {
		return false
	}
	if !types.MatchUuidIfNonZero(row.Version, row1.Version) {
		return false
	}
	if !types.MatchUuidMultiplesIfNonZero(row.Managers, row1.Managers) {
		return false
	}
	if !types.MatchMapStringStringIfNonZero(row.OtherConfig, row1.OtherConfig) {
		return false
	}
	if !types.MatchUuidMultiplesIfNonZero(row.Switches, row1.Switches) {
		return false
	}
	return true
}

func (tbl GlobalTable) FindManagerReferrer_managers(refUuid string) (r []*Global) {
	for i := range tbl {
		row := &tbl[i]
		for _, val := range row.Managers {
			if val == refUuid {
				r = append(r, row)
			}
		}
	}
	return r
}

func (tbl GlobalTable) FindPhysicalSwitchReferrer_switches(refUuid string) (r []*Global) {
	for i := range tbl {
		row := &tbl[i]
		for _, val := range row.Switches {
			if val == refUuid {
				r = append(r, row)
			}
		}
	}
	return r
}

func (row *Global) HasExternalIds() bool {
	return false
}

func (row *Global) SetExternalId(k, v string) {
	panic(errors.Wrap(types.ErrUnknownColumn, "external_ids"))
}

func (row *Global) GetExternalId(k string) (string, bool) {
	panic(errors.Wrap(types.ErrUnknownColumn, "external_ids"))
}

func (row *Global) RemoveExternalId(k string) (string, bool) {
	panic(errors.Wrap(types.ErrUnknownColumn, "external_ids"))
}

type LogicalBindingStatsTable []LogicalBindingStats

var _ types.ITable = &LogicalBindingStatsTable{}

func (tbl LogicalBindingStatsTable) OvsdbTableName() string {
	return "Logical_Binding_Stats"
}

func (tbl LogicalBindingStatsTable) OvsdbIsRoot() bool {
	return false
}

func (tbl LogicalBindingStatsTable) Rows() []types.IRow {
	r := make([]types.IRow, len(tbl))
	for i := range tbl {
		r[i] = &tbl[i]
	}
	return r
}

func (tbl LogicalBindingStatsTable) NewRow() types.IRow {
	return &LogicalBindingStats{}
}

func (tbl *LogicalBindingStatsTable) AppendRow(irow types.IRow) {
	row := irow.(*LogicalBindingStats)
	*tbl = append(*tbl, *row)
}

func (tbl LogicalBindingStatsTable) OvsdbHasIndex() bool {
	return false
}

func (tbl LogicalBindingStatsTable) OvsdbGetByAnyIndex(irow1 types.IRow) types.IRow {
	return nil
}

func (tbl LogicalBindingStatsTable) FindOneMatchNonZeros(row1 *LogicalBindingStats) *LogicalBindingStats {
	for i := range tbl {
		row := &tbl[i]
		if row.MatchNonZeros(row1) {
			return row
		}
	}
	return nil
}

type LogicalBindingStats struct {
	Uuid             string `json:"_uuid"`
	Version          string `json:"_version"`
	BytesFromLocal   int64  `json:"bytes_from_local"`
	BytesToLocal     int64  `json:"bytes_to_local"`
	PacketsFromLocal int64  `json:"packets_from_local"`
	PacketsToLocal   int64  `json:"packets_to_local"`
}

var _ types.IRow = &LogicalBindingStats{}

func (row *LogicalBindingStats) OvsdbTableName() string {
	return "Logical_Binding_Stats"
}

func (row *LogicalBindingStats) OvsdbIsRoot() bool {
	return false
}

func (row *LogicalBindingStats) OvsdbUuid() string {
	return row.Uuid
}

func (row *LogicalBindingStats) OvsdbCmdArgs() []string {
	r := []string{}
	r = append(r, types.OvsdbCmdArgsInteger("bytes_from_local", row.BytesFromLocal)...)
	r = append(r, types.OvsdbCmdArgsInteger("bytes_to_local", row.BytesToLocal)...)
	r = append(r, types.OvsdbCmdArgsInteger("packets_from_local", row.PacketsFromLocal)...)
	r = append(r, types.OvsdbCmdArgsInteger("packets_to_local", row.PacketsToLocal)...)
	return r
}

func (row *LogicalBindingStats) SetColumn(name string, val interface{}) (err error) {
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
	case "bytes_from_local":
		row.BytesFromLocal = types.EnsureInteger(val)
	case "bytes_to_local":
		row.BytesToLocal = types.EnsureInteger(val)
	case "packets_from_local":
		row.PacketsFromLocal = types.EnsureInteger(val)
	case "packets_to_local":
		row.PacketsToLocal = types.EnsureInteger(val)
	default:
		panic(types.ErrUnknownColumn)
	}
	return
}

func (row *LogicalBindingStats) MatchNonZeros(row1 *LogicalBindingStats) bool {
	if !types.MatchUuidIfNonZero(row.Uuid, row1.Uuid) {
		return false
	}
	if !types.MatchUuidIfNonZero(row.Version, row1.Version) {
		return false
	}
	if !types.MatchIntegerIfNonZero(row.BytesFromLocal, row1.BytesFromLocal) {
		return false
	}
	if !types.MatchIntegerIfNonZero(row.BytesToLocal, row1.BytesToLocal) {
		return false
	}
	if !types.MatchIntegerIfNonZero(row.PacketsFromLocal, row1.PacketsFromLocal) {
		return false
	}
	if !types.MatchIntegerIfNonZero(row.PacketsToLocal, row1.PacketsToLocal) {
		return false
	}
	return true
}

func (row *LogicalBindingStats) HasExternalIds() bool {
	return false
}

func (row *LogicalBindingStats) SetExternalId(k, v string) {
	panic(errors.Wrap(types.ErrUnknownColumn, "external_ids"))
}

func (row *LogicalBindingStats) GetExternalId(k string) (string, bool) {
	panic(errors.Wrap(types.ErrUnknownColumn, "external_ids"))
}

func (row *LogicalBindingStats) RemoveExternalId(k string) (string, bool) {
	panic(errors.Wrap(types.ErrUnknownColumn, "external_ids"))
}

type LogicalRouterTable []LogicalRouter

var _ types.ITable = &LogicalRouterTable{}

func (tbl LogicalRouterTable) OvsdbTableName() string {
	return "Logical_Router"
}

func (tbl LogicalRouterTable) OvsdbIsRoot() bool {
	return true
}

func (tbl LogicalRouterTable) Rows() []types.IRow {
	r := make([]types.IRow, len(tbl))
	for i := range tbl {
		r[i] = &tbl[i]
	}
	return r
}

func (tbl LogicalRouterTable) NewRow() types.IRow {
	return &LogicalRouter{}
}

func (tbl *LogicalRouterTable) AppendRow(irow types.IRow) {
	row := irow.(*LogicalRouter)
	*tbl = append(*tbl, *row)
}

func (tbl LogicalRouterTable) OvsdbHasIndex() bool {
	return true
}

func (row *LogicalRouter) MatchByName(row1 *LogicalRouter) bool {
	if !types.MatchString(row.Name, row1.Name) {
		return false
	}
	return true
}

func (tbl LogicalRouterTable) GetByName(row1 *LogicalRouter) *LogicalRouter {
	for i := range tbl {
		row := &tbl[i]
		if row.MatchByName(row1) {
			return row
		}
	}
	return nil
}

func (tbl LogicalRouterTable) OvsdbGetByAnyIndex(irow1 types.IRow) types.IRow {
	row1 := irow1.(*LogicalRouter)
	if !(types.IsZeroString(row1.Name)) {
		if row := tbl.GetByName(row1); row != nil {
			return row
		}
	}
	return nil
}

func (tbl LogicalRouterTable) FindOneMatchNonZeros(row1 *LogicalRouter) *LogicalRouter {
	for i := range tbl {
		row := &tbl[i]
		if row.MatchNonZeros(row1) {
			return row
		}
	}
	return nil
}

type LogicalRouter struct {
	LRFaultStatus []string          `json:"LR_fault_status"`
	Uuid          string            `json:"_uuid"`
	Version       string            `json:"_version"`
	AclBinding    map[string]string `json:"acl_binding"`
	Description   string            `json:"description"`
	Name          string            `json:"name"`
	OtherConfig   map[string]string `json:"other_config"`
	StaticRoutes  map[string]string `json:"static_routes"`
	SwitchBinding map[string]string `json:"switch_binding"`
}

var _ types.IRow = &LogicalRouter{}

func (row *LogicalRouter) OvsdbTableName() string {
	return "Logical_Router"
}

func (row *LogicalRouter) OvsdbIsRoot() bool {
	return true
}

func (row *LogicalRouter) OvsdbUuid() string {
	return row.Uuid
}

func (row *LogicalRouter) OvsdbCmdArgs() []string {
	r := []string{}
	r = append(r, types.OvsdbCmdArgsStringMultiples("LR_fault_status", row.LRFaultStatus)...)
	r = append(r, types.OvsdbCmdArgsMapStringUuid("acl_binding", row.AclBinding)...)
	r = append(r, types.OvsdbCmdArgsString("description", row.Description)...)
	r = append(r, types.OvsdbCmdArgsString("name", row.Name)...)
	r = append(r, types.OvsdbCmdArgsMapStringString("other_config", row.OtherConfig)...)
	r = append(r, types.OvsdbCmdArgsMapStringString("static_routes", row.StaticRoutes)...)
	r = append(r, types.OvsdbCmdArgsMapStringUuid("switch_binding", row.SwitchBinding)...)
	return r
}

func (row *LogicalRouter) SetColumn(name string, val interface{}) (err error) {
	defer func() {
		if panicErr := recover(); panicErr != nil {
			err = errors.Wrapf(panicErr.(error), "%s: %#v", name, fmt.Sprintf("%#v", val))
		}
	}()
	switch name {
	case "LR_fault_status":
		row.LRFaultStatus = types.EnsureStringMultiples(val)
	case "_uuid":
		row.Uuid = types.EnsureUuid(val)
	case "_version":
		row.Version = types.EnsureUuid(val)
	case "acl_binding":
		row.AclBinding = types.EnsureMapStringUuid(val)
	case "description":
		row.Description = types.EnsureString(val)
	case "name":
		row.Name = types.EnsureString(val)
	case "other_config":
		row.OtherConfig = types.EnsureMapStringString(val)
	case "static_routes":
		row.StaticRoutes = types.EnsureMapStringString(val)
	case "switch_binding":
		row.SwitchBinding = types.EnsureMapStringUuid(val)
	default:
		panic(types.ErrUnknownColumn)
	}
	return
}

func (row *LogicalRouter) MatchNonZeros(row1 *LogicalRouter) bool {
	if !types.MatchStringMultiplesIfNonZero(row.LRFaultStatus, row1.LRFaultStatus) {
		return false
	}
	if !types.MatchUuidIfNonZero(row.Uuid, row1.Uuid) {
		return false
	}
	if !types.MatchUuidIfNonZero(row.Version, row1.Version) {
		return false
	}
	if !types.MatchMapStringUuidIfNonZero(row.AclBinding, row1.AclBinding) {
		return false
	}
	if !types.MatchStringIfNonZero(row.Description, row1.Description) {
		return false
	}
	if !types.MatchStringIfNonZero(row.Name, row1.Name) {
		return false
	}
	if !types.MatchMapStringStringIfNonZero(row.OtherConfig, row1.OtherConfig) {
		return false
	}
	if !types.MatchMapStringStringIfNonZero(row.StaticRoutes, row1.StaticRoutes) {
		return false
	}
	if !types.MatchMapStringUuidIfNonZero(row.SwitchBinding, row1.SwitchBinding) {
		return false
	}
	return true
}

func (tbl LogicalRouterTable) FindACLReferrer2_acl_binding(refUuid string) (r []*LogicalRouter) {
	for i := range tbl {
		row := &tbl[i]
		for _, val := range row.AclBinding {
			if val == refUuid {
				r = append(r, row)
			}
		}
	}
	return r
}

func (tbl LogicalRouterTable) FindLogicalSwitchReferrer2_switch_binding(refUuid string) (r []*LogicalRouter) {
	for i := range tbl {
		row := &tbl[i]
		for _, val := range row.SwitchBinding {
			if val == refUuid {
				r = append(r, row)
			}
		}
	}
	return r
}

func (row *LogicalRouter) HasExternalIds() bool {
	return false
}

func (row *LogicalRouter) SetExternalId(k, v string) {
	panic(errors.Wrap(types.ErrUnknownColumn, "external_ids"))
}

func (row *LogicalRouter) GetExternalId(k string) (string, bool) {
	panic(errors.Wrap(types.ErrUnknownColumn, "external_ids"))
}

func (row *LogicalRouter) RemoveExternalId(k string) (string, bool) {
	panic(errors.Wrap(types.ErrUnknownColumn, "external_ids"))
}

type LogicalSwitchTable []LogicalSwitch

var _ types.ITable = &LogicalSwitchTable{}

func (tbl LogicalSwitchTable) OvsdbTableName() string {
	return "Logical_Switch"
}

func (tbl LogicalSwitchTable) OvsdbIsRoot() bool {
	return true
}

func (tbl LogicalSwitchTable) Rows() []types.IRow {
	r := make([]types.IRow, len(tbl))
	for i := range tbl {
		r[i] = &tbl[i]
	}
	return r
}

func (tbl LogicalSwitchTable) NewRow() types.IRow {
	return &LogicalSwitch{}
}

func (tbl *LogicalSwitchTable) AppendRow(irow types.IRow) {
	row := irow.(*LogicalSwitch)
	*tbl = append(*tbl, *row)
}

func (tbl LogicalSwitchTable) OvsdbHasIndex() bool {
	return true
}

func (row *LogicalSwitch) MatchByName(row1 *LogicalSwitch) bool {
	if !types.MatchString(row.Name, row1.Name) {
		return false
	}
	return true
}

func (tbl LogicalSwitchTable) GetByName(row1 *LogicalSwitch) *LogicalSwitch {
	for i := range tbl {
		row := &tbl[i]
		if row.MatchByName(row1) {
			return row
		}
	}
	return nil
}

func (tbl LogicalSwitchTable) OvsdbGetByAnyIndex(irow1 types.IRow) types.IRow {
	row1 := irow1.(*LogicalSwitch)
	if !(types.IsZeroString(row1.Name)) {
		if row := tbl.GetByName(row1); row != nil {
			return row
		}
	}
	return nil
}

func (tbl LogicalSwitchTable) FindOneMatchNonZeros(row1 *LogicalSwitch) *LogicalSwitch {
	for i := range tbl {
		row := &tbl[i]
		if row.MatchNonZeros(row1) {
			return row
		}
	}
	return nil
}

type LogicalSwitch struct {
	Uuid            string            `json:"_uuid"`
	Version         string            `json:"_version"`
	Description     string            `json:"description"`
	Name            string            `json:"name"`
	OtherConfig     map[string]string `json:"other_config"`
	ReplicationMode *string           `json:"replication_mode"`
	TunnelKey       *int64            `json:"tunnel_key"`
}

var _ types.IRow = &LogicalSwitch{}

func (row *LogicalSwitch) OvsdbTableName() string {
	return "Logical_Switch"
}

func (row *LogicalSwitch) OvsdbIsRoot() bool {
	return true
}

func (row *LogicalSwitch) OvsdbUuid() string {
	return row.Uuid
}

func (row *LogicalSwitch) OvsdbCmdArgs() []string {
	r := []string{}
	r = append(r, types.OvsdbCmdArgsString("description", row.Description)...)
	r = append(r, types.OvsdbCmdArgsString("name", row.Name)...)
	r = append(r, types.OvsdbCmdArgsMapStringString("other_config", row.OtherConfig)...)
	r = append(r, types.OvsdbCmdArgsStringOptional("replication_mode", row.ReplicationMode)...)
	r = append(r, types.OvsdbCmdArgsIntegerOptional("tunnel_key", row.TunnelKey)...)
	return r
}

func (row *LogicalSwitch) SetColumn(name string, val interface{}) (err error) {
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
	case "description":
		row.Description = types.EnsureString(val)
	case "name":
		row.Name = types.EnsureString(val)
	case "other_config":
		row.OtherConfig = types.EnsureMapStringString(val)
	case "replication_mode":
		row.ReplicationMode = types.EnsureStringOptional(val)
	case "tunnel_key":
		row.TunnelKey = types.EnsureIntegerOptional(val)
	default:
		panic(types.ErrUnknownColumn)
	}
	return
}

func (row *LogicalSwitch) MatchNonZeros(row1 *LogicalSwitch) bool {
	if !types.MatchUuidIfNonZero(row.Uuid, row1.Uuid) {
		return false
	}
	if !types.MatchUuidIfNonZero(row.Version, row1.Version) {
		return false
	}
	if !types.MatchStringIfNonZero(row.Description, row1.Description) {
		return false
	}
	if !types.MatchStringIfNonZero(row.Name, row1.Name) {
		return false
	}
	if !types.MatchMapStringStringIfNonZero(row.OtherConfig, row1.OtherConfig) {
		return false
	}
	if !types.MatchStringOptionalIfNonZero(row.ReplicationMode, row1.ReplicationMode) {
		return false
	}
	if !types.MatchIntegerOptionalIfNonZero(row.TunnelKey, row1.TunnelKey) {
		return false
	}
	return true
}

func (row *LogicalSwitch) HasExternalIds() bool {
	return false
}

func (row *LogicalSwitch) SetExternalId(k, v string) {
	panic(errors.Wrap(types.ErrUnknownColumn, "external_ids"))
}

func (row *LogicalSwitch) GetExternalId(k string) (string, bool) {
	panic(errors.Wrap(types.ErrUnknownColumn, "external_ids"))
}

func (row *LogicalSwitch) RemoveExternalId(k string) (string, bool) {
	panic(errors.Wrap(types.ErrUnknownColumn, "external_ids"))
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
	return false
}

func (row *Manager) SetExternalId(k, v string) {
	panic(errors.Wrap(types.ErrUnknownColumn, "external_ids"))
}

func (row *Manager) GetExternalId(k string) (string, bool) {
	panic(errors.Wrap(types.ErrUnknownColumn, "external_ids"))
}

func (row *Manager) RemoveExternalId(k string) (string, bool) {
	panic(errors.Wrap(types.ErrUnknownColumn, "external_ids"))
}

type McastMacsLocalTable []McastMacsLocal

var _ types.ITable = &McastMacsLocalTable{}

func (tbl McastMacsLocalTable) OvsdbTableName() string {
	return "Mcast_Macs_Local"
}

func (tbl McastMacsLocalTable) OvsdbIsRoot() bool {
	return true
}

func (tbl McastMacsLocalTable) Rows() []types.IRow {
	r := make([]types.IRow, len(tbl))
	for i := range tbl {
		r[i] = &tbl[i]
	}
	return r
}

func (tbl McastMacsLocalTable) NewRow() types.IRow {
	return &McastMacsLocal{}
}

func (tbl *McastMacsLocalTable) AppendRow(irow types.IRow) {
	row := irow.(*McastMacsLocal)
	*tbl = append(*tbl, *row)
}

func (tbl McastMacsLocalTable) OvsdbHasIndex() bool {
	return false
}

func (tbl McastMacsLocalTable) OvsdbGetByAnyIndex(irow1 types.IRow) types.IRow {
	return nil
}

func (tbl McastMacsLocalTable) FindOneMatchNonZeros(row1 *McastMacsLocal) *McastMacsLocal {
	for i := range tbl {
		row := &tbl[i]
		if row.MatchNonZeros(row1) {
			return row
		}
	}
	return nil
}

type McastMacsLocal struct {
	MAC           string `json:"MAC"`
	Uuid          string `json:"_uuid"`
	Version       string `json:"_version"`
	Ipaddr        string `json:"ipaddr"`
	LocatorSet    string `json:"locator_set"`
	LogicalSwitch string `json:"logical_switch"`
}

var _ types.IRow = &McastMacsLocal{}

func (row *McastMacsLocal) OvsdbTableName() string {
	return "Mcast_Macs_Local"
}

func (row *McastMacsLocal) OvsdbIsRoot() bool {
	return true
}

func (row *McastMacsLocal) OvsdbUuid() string {
	return row.Uuid
}

func (row *McastMacsLocal) OvsdbCmdArgs() []string {
	r := []string{}
	r = append(r, types.OvsdbCmdArgsString("MAC", row.MAC)...)
	r = append(r, types.OvsdbCmdArgsString("ipaddr", row.Ipaddr)...)
	r = append(r, types.OvsdbCmdArgsUuid("locator_set", row.LocatorSet)...)
	r = append(r, types.OvsdbCmdArgsUuid("logical_switch", row.LogicalSwitch)...)
	return r
}

func (row *McastMacsLocal) SetColumn(name string, val interface{}) (err error) {
	defer func() {
		if panicErr := recover(); panicErr != nil {
			err = errors.Wrapf(panicErr.(error), "%s: %#v", name, fmt.Sprintf("%#v", val))
		}
	}()
	switch name {
	case "MAC":
		row.MAC = types.EnsureString(val)
	case "_uuid":
		row.Uuid = types.EnsureUuid(val)
	case "_version":
		row.Version = types.EnsureUuid(val)
	case "ipaddr":
		row.Ipaddr = types.EnsureString(val)
	case "locator_set":
		row.LocatorSet = types.EnsureUuid(val)
	case "logical_switch":
		row.LogicalSwitch = types.EnsureUuid(val)
	default:
		panic(types.ErrUnknownColumn)
	}
	return
}

func (row *McastMacsLocal) MatchNonZeros(row1 *McastMacsLocal) bool {
	if !types.MatchStringIfNonZero(row.MAC, row1.MAC) {
		return false
	}
	if !types.MatchUuidIfNonZero(row.Uuid, row1.Uuid) {
		return false
	}
	if !types.MatchUuidIfNonZero(row.Version, row1.Version) {
		return false
	}
	if !types.MatchStringIfNonZero(row.Ipaddr, row1.Ipaddr) {
		return false
	}
	if !types.MatchUuidIfNonZero(row.LocatorSet, row1.LocatorSet) {
		return false
	}
	if !types.MatchUuidIfNonZero(row.LogicalSwitch, row1.LogicalSwitch) {
		return false
	}
	return true
}

func (tbl McastMacsLocalTable) FindPhysicalLocatorSetReferrer_locator_set(refUuid string) (r []*McastMacsLocal) {
	for i := range tbl {
		row := &tbl[i]
		if row.LocatorSet == refUuid {
			r = append(r, row)
		}
	}
	return r
}

func (tbl McastMacsLocalTable) FindLogicalSwitchReferrer_logical_switch(refUuid string) (r []*McastMacsLocal) {
	for i := range tbl {
		row := &tbl[i]
		if row.LogicalSwitch == refUuid {
			r = append(r, row)
		}
	}
	return r
}

func (row *McastMacsLocal) HasExternalIds() bool {
	return false
}

func (row *McastMacsLocal) SetExternalId(k, v string) {
	panic(errors.Wrap(types.ErrUnknownColumn, "external_ids"))
}

func (row *McastMacsLocal) GetExternalId(k string) (string, bool) {
	panic(errors.Wrap(types.ErrUnknownColumn, "external_ids"))
}

func (row *McastMacsLocal) RemoveExternalId(k string) (string, bool) {
	panic(errors.Wrap(types.ErrUnknownColumn, "external_ids"))
}

type McastMacsRemoteTable []McastMacsRemote

var _ types.ITable = &McastMacsRemoteTable{}

func (tbl McastMacsRemoteTable) OvsdbTableName() string {
	return "Mcast_Macs_Remote"
}

func (tbl McastMacsRemoteTable) OvsdbIsRoot() bool {
	return true
}

func (tbl McastMacsRemoteTable) Rows() []types.IRow {
	r := make([]types.IRow, len(tbl))
	for i := range tbl {
		r[i] = &tbl[i]
	}
	return r
}

func (tbl McastMacsRemoteTable) NewRow() types.IRow {
	return &McastMacsRemote{}
}

func (tbl *McastMacsRemoteTable) AppendRow(irow types.IRow) {
	row := irow.(*McastMacsRemote)
	*tbl = append(*tbl, *row)
}

func (tbl McastMacsRemoteTable) OvsdbHasIndex() bool {
	return false
}

func (tbl McastMacsRemoteTable) OvsdbGetByAnyIndex(irow1 types.IRow) types.IRow {
	return nil
}

func (tbl McastMacsRemoteTable) FindOneMatchNonZeros(row1 *McastMacsRemote) *McastMacsRemote {
	for i := range tbl {
		row := &tbl[i]
		if row.MatchNonZeros(row1) {
			return row
		}
	}
	return nil
}

type McastMacsRemote struct {
	MAC           string `json:"MAC"`
	Uuid          string `json:"_uuid"`
	Version       string `json:"_version"`
	Ipaddr        string `json:"ipaddr"`
	LocatorSet    string `json:"locator_set"`
	LogicalSwitch string `json:"logical_switch"`
}

var _ types.IRow = &McastMacsRemote{}

func (row *McastMacsRemote) OvsdbTableName() string {
	return "Mcast_Macs_Remote"
}

func (row *McastMacsRemote) OvsdbIsRoot() bool {
	return true
}

func (row *McastMacsRemote) OvsdbUuid() string {
	return row.Uuid
}

func (row *McastMacsRemote) OvsdbCmdArgs() []string {
	r := []string{}
	r = append(r, types.OvsdbCmdArgsString("MAC", row.MAC)...)
	r = append(r, types.OvsdbCmdArgsString("ipaddr", row.Ipaddr)...)
	r = append(r, types.OvsdbCmdArgsUuid("locator_set", row.LocatorSet)...)
	r = append(r, types.OvsdbCmdArgsUuid("logical_switch", row.LogicalSwitch)...)
	return r
}

func (row *McastMacsRemote) SetColumn(name string, val interface{}) (err error) {
	defer func() {
		if panicErr := recover(); panicErr != nil {
			err = errors.Wrapf(panicErr.(error), "%s: %#v", name, fmt.Sprintf("%#v", val))
		}
	}()
	switch name {
	case "MAC":
		row.MAC = types.EnsureString(val)
	case "_uuid":
		row.Uuid = types.EnsureUuid(val)
	case "_version":
		row.Version = types.EnsureUuid(val)
	case "ipaddr":
		row.Ipaddr = types.EnsureString(val)
	case "locator_set":
		row.LocatorSet = types.EnsureUuid(val)
	case "logical_switch":
		row.LogicalSwitch = types.EnsureUuid(val)
	default:
		panic(types.ErrUnknownColumn)
	}
	return
}

func (row *McastMacsRemote) MatchNonZeros(row1 *McastMacsRemote) bool {
	if !types.MatchStringIfNonZero(row.MAC, row1.MAC) {
		return false
	}
	if !types.MatchUuidIfNonZero(row.Uuid, row1.Uuid) {
		return false
	}
	if !types.MatchUuidIfNonZero(row.Version, row1.Version) {
		return false
	}
	if !types.MatchStringIfNonZero(row.Ipaddr, row1.Ipaddr) {
		return false
	}
	if !types.MatchUuidIfNonZero(row.LocatorSet, row1.LocatorSet) {
		return false
	}
	if !types.MatchUuidIfNonZero(row.LogicalSwitch, row1.LogicalSwitch) {
		return false
	}
	return true
}

func (tbl McastMacsRemoteTable) FindPhysicalLocatorSetReferrer_locator_set(refUuid string) (r []*McastMacsRemote) {
	for i := range tbl {
		row := &tbl[i]
		if row.LocatorSet == refUuid {
			r = append(r, row)
		}
	}
	return r
}

func (tbl McastMacsRemoteTable) FindLogicalSwitchReferrer_logical_switch(refUuid string) (r []*McastMacsRemote) {
	for i := range tbl {
		row := &tbl[i]
		if row.LogicalSwitch == refUuid {
			r = append(r, row)
		}
	}
	return r
}

func (row *McastMacsRemote) HasExternalIds() bool {
	return false
}

func (row *McastMacsRemote) SetExternalId(k, v string) {
	panic(errors.Wrap(types.ErrUnknownColumn, "external_ids"))
}

func (row *McastMacsRemote) GetExternalId(k string) (string, bool) {
	panic(errors.Wrap(types.ErrUnknownColumn, "external_ids"))
}

func (row *McastMacsRemote) RemoveExternalId(k string) (string, bool) {
	panic(errors.Wrap(types.ErrUnknownColumn, "external_ids"))
}

type PhysicalLocatorTable []PhysicalLocator

var _ types.ITable = &PhysicalLocatorTable{}

func (tbl PhysicalLocatorTable) OvsdbTableName() string {
	return "Physical_Locator"
}

func (tbl PhysicalLocatorTable) OvsdbIsRoot() bool {
	return false
}

func (tbl PhysicalLocatorTable) Rows() []types.IRow {
	r := make([]types.IRow, len(tbl))
	for i := range tbl {
		r[i] = &tbl[i]
	}
	return r
}

func (tbl PhysicalLocatorTable) NewRow() types.IRow {
	return &PhysicalLocator{}
}

func (tbl *PhysicalLocatorTable) AppendRow(irow types.IRow) {
	row := irow.(*PhysicalLocator)
	*tbl = append(*tbl, *row)
}

func (tbl PhysicalLocatorTable) OvsdbHasIndex() bool {
	return true
}

func (row *PhysicalLocator) MatchByEncapsulationTypeDstIpTunnelKey(row1 *PhysicalLocator) bool {
	if !types.MatchString(row.EncapsulationType, row1.EncapsulationType) {
		return false
	}
	if !types.MatchString(row.DstIp, row1.DstIp) {
		return false
	}
	if !types.MatchIntegerOptional(row.TunnelKey, row1.TunnelKey) {
		return false
	}
	return true
}

func (tbl PhysicalLocatorTable) GetByEncapsulationTypeDstIpTunnelKey(row1 *PhysicalLocator) *PhysicalLocator {
	for i := range tbl {
		row := &tbl[i]
		if row.MatchByEncapsulationTypeDstIpTunnelKey(row1) {
			return row
		}
	}
	return nil
}

func (tbl PhysicalLocatorTable) OvsdbGetByAnyIndex(irow1 types.IRow) types.IRow {
	row1 := irow1.(*PhysicalLocator)
	if !(types.IsZeroString(row1.EncapsulationType) || types.IsZeroString(row1.DstIp) || types.IsZeroIntegerOptional(row1.TunnelKey)) {
		if row := tbl.GetByEncapsulationTypeDstIpTunnelKey(row1); row != nil {
			return row
		}
	}
	return nil
}

func (tbl PhysicalLocatorTable) FindOneMatchNonZeros(row1 *PhysicalLocator) *PhysicalLocator {
	for i := range tbl {
		row := &tbl[i]
		if row.MatchNonZeros(row1) {
			return row
		}
	}
	return nil
}

type PhysicalLocator struct {
	Uuid              string `json:"_uuid"`
	Version           string `json:"_version"`
	DstIp             string `json:"dst_ip"`
	EncapsulationType string `json:"encapsulation_type"`
	TunnelKey         *int64 `json:"tunnel_key"`
}

var _ types.IRow = &PhysicalLocator{}

func (row *PhysicalLocator) OvsdbTableName() string {
	return "Physical_Locator"
}

func (row *PhysicalLocator) OvsdbIsRoot() bool {
	return false
}

func (row *PhysicalLocator) OvsdbUuid() string {
	return row.Uuid
}

func (row *PhysicalLocator) OvsdbCmdArgs() []string {
	r := []string{}
	r = append(r, types.OvsdbCmdArgsString("dst_ip", row.DstIp)...)
	r = append(r, types.OvsdbCmdArgsString("encapsulation_type", row.EncapsulationType)...)
	r = append(r, types.OvsdbCmdArgsIntegerOptional("tunnel_key", row.TunnelKey)...)
	return r
}

func (row *PhysicalLocator) SetColumn(name string, val interface{}) (err error) {
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
	case "dst_ip":
		row.DstIp = types.EnsureString(val)
	case "encapsulation_type":
		row.EncapsulationType = types.EnsureString(val)
	case "tunnel_key":
		row.TunnelKey = types.EnsureIntegerOptional(val)
	default:
		panic(types.ErrUnknownColumn)
	}
	return
}

func (row *PhysicalLocator) MatchNonZeros(row1 *PhysicalLocator) bool {
	if !types.MatchUuidIfNonZero(row.Uuid, row1.Uuid) {
		return false
	}
	if !types.MatchUuidIfNonZero(row.Version, row1.Version) {
		return false
	}
	if !types.MatchStringIfNonZero(row.DstIp, row1.DstIp) {
		return false
	}
	if !types.MatchStringIfNonZero(row.EncapsulationType, row1.EncapsulationType) {
		return false
	}
	if !types.MatchIntegerOptionalIfNonZero(row.TunnelKey, row1.TunnelKey) {
		return false
	}
	return true
}

func (row *PhysicalLocator) HasExternalIds() bool {
	return false
}

func (row *PhysicalLocator) SetExternalId(k, v string) {
	panic(errors.Wrap(types.ErrUnknownColumn, "external_ids"))
}

func (row *PhysicalLocator) GetExternalId(k string) (string, bool) {
	panic(errors.Wrap(types.ErrUnknownColumn, "external_ids"))
}

func (row *PhysicalLocator) RemoveExternalId(k string) (string, bool) {
	panic(errors.Wrap(types.ErrUnknownColumn, "external_ids"))
}

type PhysicalLocatorSetTable []PhysicalLocatorSet

var _ types.ITable = &PhysicalLocatorSetTable{}

func (tbl PhysicalLocatorSetTable) OvsdbTableName() string {
	return "Physical_Locator_Set"
}

func (tbl PhysicalLocatorSetTable) OvsdbIsRoot() bool {
	return false
}

func (tbl PhysicalLocatorSetTable) Rows() []types.IRow {
	r := make([]types.IRow, len(tbl))
	for i := range tbl {
		r[i] = &tbl[i]
	}
	return r
}

func (tbl PhysicalLocatorSetTable) NewRow() types.IRow {
	return &PhysicalLocatorSet{}
}

func (tbl *PhysicalLocatorSetTable) AppendRow(irow types.IRow) {
	row := irow.(*PhysicalLocatorSet)
	*tbl = append(*tbl, *row)
}

func (tbl PhysicalLocatorSetTable) OvsdbHasIndex() bool {
	return false
}

func (tbl PhysicalLocatorSetTable) OvsdbGetByAnyIndex(irow1 types.IRow) types.IRow {
	return nil
}

func (tbl PhysicalLocatorSetTable) FindOneMatchNonZeros(row1 *PhysicalLocatorSet) *PhysicalLocatorSet {
	for i := range tbl {
		row := &tbl[i]
		if row.MatchNonZeros(row1) {
			return row
		}
	}
	return nil
}

type PhysicalLocatorSet struct {
	Uuid     string   `json:"_uuid"`
	Version  string   `json:"_version"`
	Locators []string `json:"locators"`
}

var _ types.IRow = &PhysicalLocatorSet{}

func (row *PhysicalLocatorSet) OvsdbTableName() string {
	return "Physical_Locator_Set"
}

func (row *PhysicalLocatorSet) OvsdbIsRoot() bool {
	return false
}

func (row *PhysicalLocatorSet) OvsdbUuid() string {
	return row.Uuid
}

func (row *PhysicalLocatorSet) OvsdbCmdArgs() []string {
	r := []string{}
	r = append(r, types.OvsdbCmdArgsUuidMultiples("locators", row.Locators)...)
	return r
}

func (row *PhysicalLocatorSet) SetColumn(name string, val interface{}) (err error) {
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
	case "locators":
		row.Locators = types.EnsureUuidMultiples(val)
	default:
		panic(types.ErrUnknownColumn)
	}
	return
}

func (row *PhysicalLocatorSet) MatchNonZeros(row1 *PhysicalLocatorSet) bool {
	if !types.MatchUuidIfNonZero(row.Uuid, row1.Uuid) {
		return false
	}
	if !types.MatchUuidIfNonZero(row.Version, row1.Version) {
		return false
	}
	if !types.MatchUuidMultiplesIfNonZero(row.Locators, row1.Locators) {
		return false
	}
	return true
}

func (tbl PhysicalLocatorSetTable) FindPhysicalLocatorReferrer_locators(refUuid string) (r []*PhysicalLocatorSet) {
	for i := range tbl {
		row := &tbl[i]
		for _, val := range row.Locators {
			if val == refUuid {
				r = append(r, row)
			}
		}
	}
	return r
}

func (row *PhysicalLocatorSet) HasExternalIds() bool {
	return false
}

func (row *PhysicalLocatorSet) SetExternalId(k, v string) {
	panic(errors.Wrap(types.ErrUnknownColumn, "external_ids"))
}

func (row *PhysicalLocatorSet) GetExternalId(k string) (string, bool) {
	panic(errors.Wrap(types.ErrUnknownColumn, "external_ids"))
}

func (row *PhysicalLocatorSet) RemoveExternalId(k string) (string, bool) {
	panic(errors.Wrap(types.ErrUnknownColumn, "external_ids"))
}

type PhysicalPortTable []PhysicalPort

var _ types.ITable = &PhysicalPortTable{}

func (tbl PhysicalPortTable) OvsdbTableName() string {
	return "Physical_Port"
}

func (tbl PhysicalPortTable) OvsdbIsRoot() bool {
	return false
}

func (tbl PhysicalPortTable) Rows() []types.IRow {
	r := make([]types.IRow, len(tbl))
	for i := range tbl {
		r[i] = &tbl[i]
	}
	return r
}

func (tbl PhysicalPortTable) NewRow() types.IRow {
	return &PhysicalPort{}
}

func (tbl *PhysicalPortTable) AppendRow(irow types.IRow) {
	row := irow.(*PhysicalPort)
	*tbl = append(*tbl, *row)
}

func (tbl PhysicalPortTable) OvsdbHasIndex() bool {
	return false
}

func (tbl PhysicalPortTable) OvsdbGetByAnyIndex(irow1 types.IRow) types.IRow {
	return nil
}

func (tbl PhysicalPortTable) FindOneMatchNonZeros(row1 *PhysicalPort) *PhysicalPort {
	for i := range tbl {
		row := &tbl[i]
		if row.MatchNonZeros(row1) {
			return row
		}
	}
	return nil
}

type PhysicalPort struct {
	Uuid            string            `json:"_uuid"`
	Version         string            `json:"_version"`
	AclBindings     map[int64]string  `json:"acl_bindings"`
	Description     string            `json:"description"`
	Name            string            `json:"name"`
	OtherConfig     map[string]string `json:"other_config"`
	PortFaultStatus []string          `json:"port_fault_status"`
	VlanBindings    map[int64]string  `json:"vlan_bindings"`
	VlanStats       map[int64]string  `json:"vlan_stats"`
}

var _ types.IRow = &PhysicalPort{}

func (row *PhysicalPort) OvsdbTableName() string {
	return "Physical_Port"
}

func (row *PhysicalPort) OvsdbIsRoot() bool {
	return false
}

func (row *PhysicalPort) OvsdbUuid() string {
	return row.Uuid
}

func (row *PhysicalPort) OvsdbCmdArgs() []string {
	r := []string{}
	r = append(r, types.OvsdbCmdArgsMapIntegerUuid("acl_bindings", row.AclBindings)...)
	r = append(r, types.OvsdbCmdArgsString("description", row.Description)...)
	r = append(r, types.OvsdbCmdArgsString("name", row.Name)...)
	r = append(r, types.OvsdbCmdArgsMapStringString("other_config", row.OtherConfig)...)
	r = append(r, types.OvsdbCmdArgsStringMultiples("port_fault_status", row.PortFaultStatus)...)
	r = append(r, types.OvsdbCmdArgsMapIntegerUuid("vlan_bindings", row.VlanBindings)...)
	r = append(r, types.OvsdbCmdArgsMapIntegerUuid("vlan_stats", row.VlanStats)...)
	return r
}

func (row *PhysicalPort) SetColumn(name string, val interface{}) (err error) {
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
	case "acl_bindings":
		row.AclBindings = types.EnsureMapIntegerUuid(val)
	case "description":
		row.Description = types.EnsureString(val)
	case "name":
		row.Name = types.EnsureString(val)
	case "other_config":
		row.OtherConfig = types.EnsureMapStringString(val)
	case "port_fault_status":
		row.PortFaultStatus = types.EnsureStringMultiples(val)
	case "vlan_bindings":
		row.VlanBindings = types.EnsureMapIntegerUuid(val)
	case "vlan_stats":
		row.VlanStats = types.EnsureMapIntegerUuid(val)
	default:
		panic(types.ErrUnknownColumn)
	}
	return
}

func (row *PhysicalPort) MatchNonZeros(row1 *PhysicalPort) bool {
	if !types.MatchUuidIfNonZero(row.Uuid, row1.Uuid) {
		return false
	}
	if !types.MatchUuidIfNonZero(row.Version, row1.Version) {
		return false
	}
	if !types.MatchMapIntegerUuidIfNonZero(row.AclBindings, row1.AclBindings) {
		return false
	}
	if !types.MatchStringIfNonZero(row.Description, row1.Description) {
		return false
	}
	if !types.MatchStringIfNonZero(row.Name, row1.Name) {
		return false
	}
	if !types.MatchMapStringStringIfNonZero(row.OtherConfig, row1.OtherConfig) {
		return false
	}
	if !types.MatchStringMultiplesIfNonZero(row.PortFaultStatus, row1.PortFaultStatus) {
		return false
	}
	if !types.MatchMapIntegerUuidIfNonZero(row.VlanBindings, row1.VlanBindings) {
		return false
	}
	if !types.MatchMapIntegerUuidIfNonZero(row.VlanStats, row1.VlanStats) {
		return false
	}
	return true
}

func (tbl PhysicalPortTable) FindACLReferrer2_acl_bindings(refUuid string) (r []*PhysicalPort) {
	for i := range tbl {
		row := &tbl[i]
		for _, val := range row.AclBindings {
			if val == refUuid {
				r = append(r, row)
			}
		}
	}
	return r
}

func (tbl PhysicalPortTable) FindLogicalSwitchReferrer2_vlan_bindings(refUuid string) (r []*PhysicalPort) {
	for i := range tbl {
		row := &tbl[i]
		for _, val := range row.VlanBindings {
			if val == refUuid {
				r = append(r, row)
			}
		}
	}
	return r
}

func (tbl PhysicalPortTable) FindLogicalBindingStatsReferrer2_vlan_stats(refUuid string) (r []*PhysicalPort) {
	for i := range tbl {
		row := &tbl[i]
		for _, val := range row.VlanStats {
			if val == refUuid {
				r = append(r, row)
			}
		}
	}
	return r
}

func (row *PhysicalPort) HasExternalIds() bool {
	return false
}

func (row *PhysicalPort) SetExternalId(k, v string) {
	panic(errors.Wrap(types.ErrUnknownColumn, "external_ids"))
}

func (row *PhysicalPort) GetExternalId(k string) (string, bool) {
	panic(errors.Wrap(types.ErrUnknownColumn, "external_ids"))
}

func (row *PhysicalPort) RemoveExternalId(k string) (string, bool) {
	panic(errors.Wrap(types.ErrUnknownColumn, "external_ids"))
}

type PhysicalSwitchTable []PhysicalSwitch

var _ types.ITable = &PhysicalSwitchTable{}

func (tbl PhysicalSwitchTable) OvsdbTableName() string {
	return "Physical_Switch"
}

func (tbl PhysicalSwitchTable) OvsdbIsRoot() bool {
	return false
}

func (tbl PhysicalSwitchTable) Rows() []types.IRow {
	r := make([]types.IRow, len(tbl))
	for i := range tbl {
		r[i] = &tbl[i]
	}
	return r
}

func (tbl PhysicalSwitchTable) NewRow() types.IRow {
	return &PhysicalSwitch{}
}

func (tbl *PhysicalSwitchTable) AppendRow(irow types.IRow) {
	row := irow.(*PhysicalSwitch)
	*tbl = append(*tbl, *row)
}

func (tbl PhysicalSwitchTable) OvsdbHasIndex() bool {
	return true
}

func (row *PhysicalSwitch) MatchByName(row1 *PhysicalSwitch) bool {
	if !types.MatchString(row.Name, row1.Name) {
		return false
	}
	return true
}

func (tbl PhysicalSwitchTable) GetByName(row1 *PhysicalSwitch) *PhysicalSwitch {
	for i := range tbl {
		row := &tbl[i]
		if row.MatchByName(row1) {
			return row
		}
	}
	return nil
}

func (tbl PhysicalSwitchTable) OvsdbGetByAnyIndex(irow1 types.IRow) types.IRow {
	row1 := irow1.(*PhysicalSwitch)
	if !(types.IsZeroString(row1.Name)) {
		if row := tbl.GetByName(row1); row != nil {
			return row
		}
	}
	return nil
}

func (tbl PhysicalSwitchTable) FindOneMatchNonZeros(row1 *PhysicalSwitch) *PhysicalSwitch {
	for i := range tbl {
		row := &tbl[i]
		if row.MatchNonZeros(row1) {
			return row
		}
	}
	return nil
}

type PhysicalSwitch struct {
	Uuid              string            `json:"_uuid"`
	Version           string            `json:"_version"`
	Description       string            `json:"description"`
	ManagementIps     []string          `json:"management_ips"`
	Name              string            `json:"name"`
	OtherConfig       map[string]string `json:"other_config"`
	Ports             []string          `json:"ports"`
	SwitchFaultStatus []string          `json:"switch_fault_status"`
	TunnelIps         []string          `json:"tunnel_ips"`
	Tunnels           []string          `json:"tunnels"`
}

var _ types.IRow = &PhysicalSwitch{}

func (row *PhysicalSwitch) OvsdbTableName() string {
	return "Physical_Switch"
}

func (row *PhysicalSwitch) OvsdbIsRoot() bool {
	return false
}

func (row *PhysicalSwitch) OvsdbUuid() string {
	return row.Uuid
}

func (row *PhysicalSwitch) OvsdbCmdArgs() []string {
	r := []string{}
	r = append(r, types.OvsdbCmdArgsString("description", row.Description)...)
	r = append(r, types.OvsdbCmdArgsStringMultiples("management_ips", row.ManagementIps)...)
	r = append(r, types.OvsdbCmdArgsString("name", row.Name)...)
	r = append(r, types.OvsdbCmdArgsMapStringString("other_config", row.OtherConfig)...)
	r = append(r, types.OvsdbCmdArgsUuidMultiples("ports", row.Ports)...)
	r = append(r, types.OvsdbCmdArgsStringMultiples("switch_fault_status", row.SwitchFaultStatus)...)
	r = append(r, types.OvsdbCmdArgsStringMultiples("tunnel_ips", row.TunnelIps)...)
	r = append(r, types.OvsdbCmdArgsUuidMultiples("tunnels", row.Tunnels)...)
	return r
}

func (row *PhysicalSwitch) SetColumn(name string, val interface{}) (err error) {
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
	case "description":
		row.Description = types.EnsureString(val)
	case "management_ips":
		row.ManagementIps = types.EnsureStringMultiples(val)
	case "name":
		row.Name = types.EnsureString(val)
	case "other_config":
		row.OtherConfig = types.EnsureMapStringString(val)
	case "ports":
		row.Ports = types.EnsureUuidMultiples(val)
	case "switch_fault_status":
		row.SwitchFaultStatus = types.EnsureStringMultiples(val)
	case "tunnel_ips":
		row.TunnelIps = types.EnsureStringMultiples(val)
	case "tunnels":
		row.Tunnels = types.EnsureUuidMultiples(val)
	default:
		panic(types.ErrUnknownColumn)
	}
	return
}

func (row *PhysicalSwitch) MatchNonZeros(row1 *PhysicalSwitch) bool {
	if !types.MatchUuidIfNonZero(row.Uuid, row1.Uuid) {
		return false
	}
	if !types.MatchUuidIfNonZero(row.Version, row1.Version) {
		return false
	}
	if !types.MatchStringIfNonZero(row.Description, row1.Description) {
		return false
	}
	if !types.MatchStringMultiplesIfNonZero(row.ManagementIps, row1.ManagementIps) {
		return false
	}
	if !types.MatchStringIfNonZero(row.Name, row1.Name) {
		return false
	}
	if !types.MatchMapStringStringIfNonZero(row.OtherConfig, row1.OtherConfig) {
		return false
	}
	if !types.MatchUuidMultiplesIfNonZero(row.Ports, row1.Ports) {
		return false
	}
	if !types.MatchStringMultiplesIfNonZero(row.SwitchFaultStatus, row1.SwitchFaultStatus) {
		return false
	}
	if !types.MatchStringMultiplesIfNonZero(row.TunnelIps, row1.TunnelIps) {
		return false
	}
	if !types.MatchUuidMultiplesIfNonZero(row.Tunnels, row1.Tunnels) {
		return false
	}
	return true
}

func (tbl PhysicalSwitchTable) FindPhysicalPortReferrer_ports(refUuid string) (r []*PhysicalSwitch) {
	for i := range tbl {
		row := &tbl[i]
		for _, val := range row.Ports {
			if val == refUuid {
				r = append(r, row)
			}
		}
	}
	return r
}

func (tbl PhysicalSwitchTable) FindTunnelReferrer_tunnels(refUuid string) (r []*PhysicalSwitch) {
	for i := range tbl {
		row := &tbl[i]
		for _, val := range row.Tunnels {
			if val == refUuid {
				r = append(r, row)
			}
		}
	}
	return r
}

func (row *PhysicalSwitch) HasExternalIds() bool {
	return false
}

func (row *PhysicalSwitch) SetExternalId(k, v string) {
	panic(errors.Wrap(types.ErrUnknownColumn, "external_ids"))
}

func (row *PhysicalSwitch) GetExternalId(k string) (string, bool) {
	panic(errors.Wrap(types.ErrUnknownColumn, "external_ids"))
}

func (row *PhysicalSwitch) RemoveExternalId(k string) (string, bool) {
	panic(errors.Wrap(types.ErrUnknownColumn, "external_ids"))
}

type TunnelTable []Tunnel

var _ types.ITable = &TunnelTable{}

func (tbl TunnelTable) OvsdbTableName() string {
	return "Tunnel"
}

func (tbl TunnelTable) OvsdbIsRoot() bool {
	return false
}

func (tbl TunnelTable) Rows() []types.IRow {
	r := make([]types.IRow, len(tbl))
	for i := range tbl {
		r[i] = &tbl[i]
	}
	return r
}

func (tbl TunnelTable) NewRow() types.IRow {
	return &Tunnel{}
}

func (tbl *TunnelTable) AppendRow(irow types.IRow) {
	row := irow.(*Tunnel)
	*tbl = append(*tbl, *row)
}

func (tbl TunnelTable) OvsdbHasIndex() bool {
	return false
}

func (tbl TunnelTable) OvsdbGetByAnyIndex(irow1 types.IRow) types.IRow {
	return nil
}

func (tbl TunnelTable) FindOneMatchNonZeros(row1 *Tunnel) *Tunnel {
	for i := range tbl {
		row := &tbl[i]
		if row.MatchNonZeros(row1) {
			return row
		}
	}
	return nil
}

type Tunnel struct {
	Uuid            string            `json:"_uuid"`
	Version         string            `json:"_version"`
	BfdConfigLocal  map[string]string `json:"bfd_config_local"`
	BfdConfigRemote map[string]string `json:"bfd_config_remote"`
	BfdParams       map[string]string `json:"bfd_params"`
	BfdStatus       map[string]string `json:"bfd_status"`
	Local           string            `json:"local"`
	Remote          string            `json:"remote"`
}

var _ types.IRow = &Tunnel{}

func (row *Tunnel) OvsdbTableName() string {
	return "Tunnel"
}

func (row *Tunnel) OvsdbIsRoot() bool {
	return false
}

func (row *Tunnel) OvsdbUuid() string {
	return row.Uuid
}

func (row *Tunnel) OvsdbCmdArgs() []string {
	r := []string{}
	r = append(r, types.OvsdbCmdArgsMapStringString("bfd_config_local", row.BfdConfigLocal)...)
	r = append(r, types.OvsdbCmdArgsMapStringString("bfd_config_remote", row.BfdConfigRemote)...)
	r = append(r, types.OvsdbCmdArgsMapStringString("bfd_params", row.BfdParams)...)
	r = append(r, types.OvsdbCmdArgsMapStringString("bfd_status", row.BfdStatus)...)
	r = append(r, types.OvsdbCmdArgsUuid("local", row.Local)...)
	r = append(r, types.OvsdbCmdArgsUuid("remote", row.Remote)...)
	return r
}

func (row *Tunnel) SetColumn(name string, val interface{}) (err error) {
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
	case "bfd_config_local":
		row.BfdConfigLocal = types.EnsureMapStringString(val)
	case "bfd_config_remote":
		row.BfdConfigRemote = types.EnsureMapStringString(val)
	case "bfd_params":
		row.BfdParams = types.EnsureMapStringString(val)
	case "bfd_status":
		row.BfdStatus = types.EnsureMapStringString(val)
	case "local":
		row.Local = types.EnsureUuid(val)
	case "remote":
		row.Remote = types.EnsureUuid(val)
	default:
		panic(types.ErrUnknownColumn)
	}
	return
}

func (row *Tunnel) MatchNonZeros(row1 *Tunnel) bool {
	if !types.MatchUuidIfNonZero(row.Uuid, row1.Uuid) {
		return false
	}
	if !types.MatchUuidIfNonZero(row.Version, row1.Version) {
		return false
	}
	if !types.MatchMapStringStringIfNonZero(row.BfdConfigLocal, row1.BfdConfigLocal) {
		return false
	}
	if !types.MatchMapStringStringIfNonZero(row.BfdConfigRemote, row1.BfdConfigRemote) {
		return false
	}
	if !types.MatchMapStringStringIfNonZero(row.BfdParams, row1.BfdParams) {
		return false
	}
	if !types.MatchMapStringStringIfNonZero(row.BfdStatus, row1.BfdStatus) {
		return false
	}
	if !types.MatchUuidIfNonZero(row.Local, row1.Local) {
		return false
	}
	if !types.MatchUuidIfNonZero(row.Remote, row1.Remote) {
		return false
	}
	return true
}

func (tbl TunnelTable) FindPhysicalLocatorReferrer_local(refUuid string) (r []*Tunnel) {
	for i := range tbl {
		row := &tbl[i]
		if row.Local == refUuid {
			r = append(r, row)
		}
	}
	return r
}

func (tbl TunnelTable) FindPhysicalLocatorReferrer_remote(refUuid string) (r []*Tunnel) {
	for i := range tbl {
		row := &tbl[i]
		if row.Remote == refUuid {
			r = append(r, row)
		}
	}
	return r
}

func (row *Tunnel) HasExternalIds() bool {
	return false
}

func (row *Tunnel) SetExternalId(k, v string) {
	panic(errors.Wrap(types.ErrUnknownColumn, "external_ids"))
}

func (row *Tunnel) GetExternalId(k string) (string, bool) {
	panic(errors.Wrap(types.ErrUnknownColumn, "external_ids"))
}

func (row *Tunnel) RemoveExternalId(k string) (string, bool) {
	panic(errors.Wrap(types.ErrUnknownColumn, "external_ids"))
}

type UcastMacsLocalTable []UcastMacsLocal

var _ types.ITable = &UcastMacsLocalTable{}

func (tbl UcastMacsLocalTable) OvsdbTableName() string {
	return "Ucast_Macs_Local"
}

func (tbl UcastMacsLocalTable) OvsdbIsRoot() bool {
	return true
}

func (tbl UcastMacsLocalTable) Rows() []types.IRow {
	r := make([]types.IRow, len(tbl))
	for i := range tbl {
		r[i] = &tbl[i]
	}
	return r
}

func (tbl UcastMacsLocalTable) NewRow() types.IRow {
	return &UcastMacsLocal{}
}

func (tbl *UcastMacsLocalTable) AppendRow(irow types.IRow) {
	row := irow.(*UcastMacsLocal)
	*tbl = append(*tbl, *row)
}

func (tbl UcastMacsLocalTable) OvsdbHasIndex() bool {
	return false
}

func (tbl UcastMacsLocalTable) OvsdbGetByAnyIndex(irow1 types.IRow) types.IRow {
	return nil
}

func (tbl UcastMacsLocalTable) FindOneMatchNonZeros(row1 *UcastMacsLocal) *UcastMacsLocal {
	for i := range tbl {
		row := &tbl[i]
		if row.MatchNonZeros(row1) {
			return row
		}
	}
	return nil
}

type UcastMacsLocal struct {
	MAC           string `json:"MAC"`
	Uuid          string `json:"_uuid"`
	Version       string `json:"_version"`
	Ipaddr        string `json:"ipaddr"`
	Locator       string `json:"locator"`
	LogicalSwitch string `json:"logical_switch"`
}

var _ types.IRow = &UcastMacsLocal{}

func (row *UcastMacsLocal) OvsdbTableName() string {
	return "Ucast_Macs_Local"
}

func (row *UcastMacsLocal) OvsdbIsRoot() bool {
	return true
}

func (row *UcastMacsLocal) OvsdbUuid() string {
	return row.Uuid
}

func (row *UcastMacsLocal) OvsdbCmdArgs() []string {
	r := []string{}
	r = append(r, types.OvsdbCmdArgsString("MAC", row.MAC)...)
	r = append(r, types.OvsdbCmdArgsString("ipaddr", row.Ipaddr)...)
	r = append(r, types.OvsdbCmdArgsUuid("locator", row.Locator)...)
	r = append(r, types.OvsdbCmdArgsUuid("logical_switch", row.LogicalSwitch)...)
	return r
}

func (row *UcastMacsLocal) SetColumn(name string, val interface{}) (err error) {
	defer func() {
		if panicErr := recover(); panicErr != nil {
			err = errors.Wrapf(panicErr.(error), "%s: %#v", name, fmt.Sprintf("%#v", val))
		}
	}()
	switch name {
	case "MAC":
		row.MAC = types.EnsureString(val)
	case "_uuid":
		row.Uuid = types.EnsureUuid(val)
	case "_version":
		row.Version = types.EnsureUuid(val)
	case "ipaddr":
		row.Ipaddr = types.EnsureString(val)
	case "locator":
		row.Locator = types.EnsureUuid(val)
	case "logical_switch":
		row.LogicalSwitch = types.EnsureUuid(val)
	default:
		panic(types.ErrUnknownColumn)
	}
	return
}

func (row *UcastMacsLocal) MatchNonZeros(row1 *UcastMacsLocal) bool {
	if !types.MatchStringIfNonZero(row.MAC, row1.MAC) {
		return false
	}
	if !types.MatchUuidIfNonZero(row.Uuid, row1.Uuid) {
		return false
	}
	if !types.MatchUuidIfNonZero(row.Version, row1.Version) {
		return false
	}
	if !types.MatchStringIfNonZero(row.Ipaddr, row1.Ipaddr) {
		return false
	}
	if !types.MatchUuidIfNonZero(row.Locator, row1.Locator) {
		return false
	}
	if !types.MatchUuidIfNonZero(row.LogicalSwitch, row1.LogicalSwitch) {
		return false
	}
	return true
}

func (tbl UcastMacsLocalTable) FindPhysicalLocatorReferrer_locator(refUuid string) (r []*UcastMacsLocal) {
	for i := range tbl {
		row := &tbl[i]
		if row.Locator == refUuid {
			r = append(r, row)
		}
	}
	return r
}

func (tbl UcastMacsLocalTable) FindLogicalSwitchReferrer_logical_switch(refUuid string) (r []*UcastMacsLocal) {
	for i := range tbl {
		row := &tbl[i]
		if row.LogicalSwitch == refUuid {
			r = append(r, row)
		}
	}
	return r
}

func (row *UcastMacsLocal) HasExternalIds() bool {
	return false
}

func (row *UcastMacsLocal) SetExternalId(k, v string) {
	panic(errors.Wrap(types.ErrUnknownColumn, "external_ids"))
}

func (row *UcastMacsLocal) GetExternalId(k string) (string, bool) {
	panic(errors.Wrap(types.ErrUnknownColumn, "external_ids"))
}

func (row *UcastMacsLocal) RemoveExternalId(k string) (string, bool) {
	panic(errors.Wrap(types.ErrUnknownColumn, "external_ids"))
}

type UcastMacsRemoteTable []UcastMacsRemote

var _ types.ITable = &UcastMacsRemoteTable{}

func (tbl UcastMacsRemoteTable) OvsdbTableName() string {
	return "Ucast_Macs_Remote"
}

func (tbl UcastMacsRemoteTable) OvsdbIsRoot() bool {
	return true
}

func (tbl UcastMacsRemoteTable) Rows() []types.IRow {
	r := make([]types.IRow, len(tbl))
	for i := range tbl {
		r[i] = &tbl[i]
	}
	return r
}

func (tbl UcastMacsRemoteTable) NewRow() types.IRow {
	return &UcastMacsRemote{}
}

func (tbl *UcastMacsRemoteTable) AppendRow(irow types.IRow) {
	row := irow.(*UcastMacsRemote)
	*tbl = append(*tbl, *row)
}

func (tbl UcastMacsRemoteTable) OvsdbHasIndex() bool {
	return false
}

func (tbl UcastMacsRemoteTable) OvsdbGetByAnyIndex(irow1 types.IRow) types.IRow {
	return nil
}

func (tbl UcastMacsRemoteTable) FindOneMatchNonZeros(row1 *UcastMacsRemote) *UcastMacsRemote {
	for i := range tbl {
		row := &tbl[i]
		if row.MatchNonZeros(row1) {
			return row
		}
	}
	return nil
}

type UcastMacsRemote struct {
	MAC           string `json:"MAC"`
	Uuid          string `json:"_uuid"`
	Version       string `json:"_version"`
	Ipaddr        string `json:"ipaddr"`
	Locator       string `json:"locator"`
	LogicalSwitch string `json:"logical_switch"`
}

var _ types.IRow = &UcastMacsRemote{}

func (row *UcastMacsRemote) OvsdbTableName() string {
	return "Ucast_Macs_Remote"
}

func (row *UcastMacsRemote) OvsdbIsRoot() bool {
	return true
}

func (row *UcastMacsRemote) OvsdbUuid() string {
	return row.Uuid
}

func (row *UcastMacsRemote) OvsdbCmdArgs() []string {
	r := []string{}
	r = append(r, types.OvsdbCmdArgsString("MAC", row.MAC)...)
	r = append(r, types.OvsdbCmdArgsString("ipaddr", row.Ipaddr)...)
	r = append(r, types.OvsdbCmdArgsUuid("locator", row.Locator)...)
	r = append(r, types.OvsdbCmdArgsUuid("logical_switch", row.LogicalSwitch)...)
	return r
}

func (row *UcastMacsRemote) SetColumn(name string, val interface{}) (err error) {
	defer func() {
		if panicErr := recover(); panicErr != nil {
			err = errors.Wrapf(panicErr.(error), "%s: %#v", name, fmt.Sprintf("%#v", val))
		}
	}()
	switch name {
	case "MAC":
		row.MAC = types.EnsureString(val)
	case "_uuid":
		row.Uuid = types.EnsureUuid(val)
	case "_version":
		row.Version = types.EnsureUuid(val)
	case "ipaddr":
		row.Ipaddr = types.EnsureString(val)
	case "locator":
		row.Locator = types.EnsureUuid(val)
	case "logical_switch":
		row.LogicalSwitch = types.EnsureUuid(val)
	default:
		panic(types.ErrUnknownColumn)
	}
	return
}

func (row *UcastMacsRemote) MatchNonZeros(row1 *UcastMacsRemote) bool {
	if !types.MatchStringIfNonZero(row.MAC, row1.MAC) {
		return false
	}
	if !types.MatchUuidIfNonZero(row.Uuid, row1.Uuid) {
		return false
	}
	if !types.MatchUuidIfNonZero(row.Version, row1.Version) {
		return false
	}
	if !types.MatchStringIfNonZero(row.Ipaddr, row1.Ipaddr) {
		return false
	}
	if !types.MatchUuidIfNonZero(row.Locator, row1.Locator) {
		return false
	}
	if !types.MatchUuidIfNonZero(row.LogicalSwitch, row1.LogicalSwitch) {
		return false
	}
	return true
}

func (tbl UcastMacsRemoteTable) FindPhysicalLocatorReferrer_locator(refUuid string) (r []*UcastMacsRemote) {
	for i := range tbl {
		row := &tbl[i]
		if row.Locator == refUuid {
			r = append(r, row)
		}
	}
	return r
}

func (tbl UcastMacsRemoteTable) FindLogicalSwitchReferrer_logical_switch(refUuid string) (r []*UcastMacsRemote) {
	for i := range tbl {
		row := &tbl[i]
		if row.LogicalSwitch == refUuid {
			r = append(r, row)
		}
	}
	return r
}

func (row *UcastMacsRemote) HasExternalIds() bool {
	return false
}

func (row *UcastMacsRemote) SetExternalId(k, v string) {
	panic(errors.Wrap(types.ErrUnknownColumn, "external_ids"))
}

func (row *UcastMacsRemote) GetExternalId(k string) (string, bool) {
	panic(errors.Wrap(types.ErrUnknownColumn, "external_ids"))
}

func (row *UcastMacsRemote) RemoveExternalId(k string) (string, bool) {
	panic(errors.Wrap(types.ErrUnknownColumn, "external_ids"))
}
