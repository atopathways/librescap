// Code generated by https://github.com/gocomply/xsd2go; DO NOT EDIT.
// Models for http://oval.mitre.org/XMLSchema/oval-system-characteristics-5#linux
package linux_sc

import (
	"encoding/xml"
	"github.com/complianceascode/librescap/pkg/scap/models/oval"
	"github.com/complianceascode/librescap/pkg/scap/models/oval_sc"
)

// Element
type ApparmorstatusItem struct {
	XMLName xml.Name `xml:apparmorstatus_item`

	Id string `xml:"id,attr"`

	Status string `xml:"status,attr"`

	LoadedProfilesCount *oval_sc.EntityItemIntType `xml:"loaded_profiles_count"`

	EnforceModeProfilesCount *oval_sc.EntityItemIntType `xml:"enforce_mode_profiles_count"`

	ComplainModeProfilesCount *oval_sc.EntityItemIntType `xml:"complain_mode_profiles_count"`

	ProcessesWithProfilesCount *oval_sc.EntityItemIntType `xml:"processes_with_profiles_count"`

	EnforceModeProcessesCount *oval_sc.EntityItemIntType `xml:"enforce_mode_processes_count"`

	ComplainModeProcessesCount *oval_sc.EntityItemIntType `xml:"complain_mode_processes_count"`

	UnconfinedProcessesWithProfilesCount *oval_sc.EntityItemIntType `xml:"unconfined_processes_with_profiles_count"`

	Message []oval.MessageType `xml:"message"`
}

// Element
type DpkginfoItem struct {
	XMLName xml.Name `xml:dpkginfo_item`

	Id string `xml:"id,attr"`

	Status string `xml:"status,attr"`

	Name *oval_sc.EntityItemStringType `xml:"name"`

	Arch *oval_sc.EntityItemStringType `xml:"arch"`

	Epoch *DpkginfoItemEpoch `xml:"epoch"`

	Release *DpkginfoItemRelease `xml:"release"`

	Version *DpkginfoItemVersion `xml:"version"`

	Evr *DpkginfoItemEvr `xml:"evr"`

	Message []oval.MessageType `xml:"message"`
}

// Element
type IflistenersItem struct {
	XMLName xml.Name `xml:iflisteners_item`

	Id string `xml:"id,attr"`

	Status string `xml:"status,attr"`

	InterfaceName *oval_sc.EntityItemStringType `xml:"interface_name"`

	Protocol *EntityItemProtocolType `xml:"protocol"`

	HwAddress *oval_sc.EntityItemStringType `xml:"hw_address"`

	ProgramName *oval_sc.EntityItemStringType `xml:"program_name"`

	Pid *oval_sc.EntityItemIntType `xml:"pid"`

	UserId *oval_sc.EntityItemIntType `xml:"user_id"`

	Message []oval.MessageType `xml:"message"`
}

// Element
type InetlisteningserverItem struct {
	XMLName xml.Name `xml:inetlisteningserver_item`

	Id string `xml:"id,attr"`

	Status string `xml:"status,attr"`

	Protocol *oval_sc.EntityItemStringType `xml:"protocol"`

	LocalAddress *oval_sc.EntityItemIPAddressStringType `xml:"local_address"`

	LocalPort *oval_sc.EntityItemIntType `xml:"local_port"`

	LocalFullAddress *oval_sc.EntityItemStringType `xml:"local_full_address"`

	ProgramName *oval_sc.EntityItemStringType `xml:"program_name"`

	ForeignAddress *oval_sc.EntityItemIPAddressStringType `xml:"foreign_address"`

	ForeignPort *oval_sc.EntityItemIntType `xml:"foreign_port"`

	ForeignFullAddress *oval_sc.EntityItemStringType `xml:"foreign_full_address"`

	Pid *oval_sc.EntityItemIntType `xml:"pid"`

	UserId *oval_sc.EntityItemIntType `xml:"user_id"`

	Message []oval.MessageType `xml:"message"`
}

// Element
type PartitionItem struct {
	XMLName xml.Name `xml:partition_item`

	Id string `xml:"id,attr"`

	Status string `xml:"status,attr"`

	MountPoint *oval_sc.EntityItemStringType `xml:"mount_point"`

	Device *oval_sc.EntityItemStringType `xml:"device"`

	Uuid *oval_sc.EntityItemStringType `xml:"uuid"`

	FsType *oval_sc.EntityItemStringType `xml:"fs_type"`

	MountOptions []oval_sc.EntityItemStringType `xml:"mount_options"`

	TotalSpace *oval_sc.EntityItemIntType `xml:"total_space"`

	SpaceUsed *oval_sc.EntityItemIntType `xml:"space_used"`

	SpaceLeft *oval_sc.EntityItemIntType `xml:"space_left"`

	SpaceLeftForUnprivilegedUsers *oval_sc.EntityItemIntType `xml:"space_left_for_unprivileged_users"`

	BlockSize *oval_sc.EntityItemIntType `xml:"block_size"`

	Message []oval.MessageType `xml:"message"`
}

// Element
type RpminfoItem struct {
	XMLName xml.Name `xml:rpminfo_item`

	Id string `xml:"id,attr"`

	Status string `xml:"status,attr"`

	Name *oval_sc.EntityItemStringType `xml:"name"`

	Arch *oval_sc.EntityItemStringType `xml:"arch"`

	Epoch *RpminfoItemEpoch `xml:"epoch"`

	Release *RpminfoItemRelease `xml:"release"`

	Version *RpminfoItemVersion `xml:"version"`

	Evr *oval_sc.EntityItemEVRStringType `xml:"evr"`

	SignatureKeyid *oval_sc.EntityItemStringType `xml:"signature_keyid"`

	ExtendedName *oval_sc.EntityItemStringType `xml:"extended_name"`

	Filepath []oval_sc.EntityItemStringType `xml:"filepath"`

	Message []oval.MessageType `xml:"message"`
}

// Element
type RpmverifyItem struct {
	XMLName xml.Name `xml:rpmverify_item`

	Id string `xml:"id,attr"`

	Status string `xml:"status,attr"`

	Name *oval_sc.EntityItemStringType `xml:"name"`

	Filepath *oval_sc.EntityItemStringType `xml:"filepath"`

	SizeDiffers *EntityItemRpmVerifyResultType `xml:"size_differs"`

	ModeDiffers *EntityItemRpmVerifyResultType `xml:"mode_differs"`

	Md5Differs *EntityItemRpmVerifyResultType `xml:"md5_differs"`

	DeviceDiffers *EntityItemRpmVerifyResultType `xml:"device_differs"`

	LinkMismatch *EntityItemRpmVerifyResultType `xml:"link_mismatch"`

	OwnershipDiffers *EntityItemRpmVerifyResultType `xml:"ownership_differs"`

	GroupDiffers *EntityItemRpmVerifyResultType `xml:"group_differs"`

	MtimeDiffers *EntityItemRpmVerifyResultType `xml:"mtime_differs"`

	CapabilitiesDiffer *EntityItemRpmVerifyResultType `xml:"capabilities_differ"`

	ConfigurationFile *oval_sc.EntityItemBoolType `xml:"configuration_file"`

	DocumentationFile *oval_sc.EntityItemBoolType `xml:"documentation_file"`

	GhostFile *oval_sc.EntityItemBoolType `xml:"ghost_file"`

	LicenseFile *oval_sc.EntityItemBoolType `xml:"license_file"`

	ReadmeFile *oval_sc.EntityItemBoolType `xml:"readme_file"`

	Message []oval.MessageType `xml:"message"`
}

// Element
type RpmverifyfileItem struct {
	XMLName xml.Name `xml:rpmverifyfile_item`

	Id string `xml:"id,attr"`

	Status string `xml:"status,attr"`

	Name *oval_sc.EntityItemStringType `xml:"name"`

	Epoch *RpmverifyfileItemEpoch `xml:"epoch"`

	Version *RpmverifyfileItemVersion `xml:"version"`

	Release *RpmverifyfileItemRelease `xml:"release"`

	Arch *oval_sc.EntityItemStringType `xml:"arch"`

	Filepath *oval_sc.EntityItemStringType `xml:"filepath"`

	ExtendedName *oval_sc.EntityItemStringType `xml:"extended_name"`

	SizeDiffers *EntityItemRpmVerifyResultType `xml:"size_differs"`

	ModeDiffers *EntityItemRpmVerifyResultType `xml:"mode_differs"`

	Md5Differs *EntityItemRpmVerifyResultType `xml:"md5_differs"`

	FiledigestDiffers *EntityItemRpmVerifyResultType `xml:"filedigest_differs"`

	DeviceDiffers *EntityItemRpmVerifyResultType `xml:"device_differs"`

	LinkMismatch *EntityItemRpmVerifyResultType `xml:"link_mismatch"`

	OwnershipDiffers *EntityItemRpmVerifyResultType `xml:"ownership_differs"`

	GroupDiffers *EntityItemRpmVerifyResultType `xml:"group_differs"`

	MtimeDiffers *EntityItemRpmVerifyResultType `xml:"mtime_differs"`

	CapabilitiesDiffer *EntityItemRpmVerifyResultType `xml:"capabilities_differ"`

	ConfigurationFile *oval_sc.EntityItemBoolType `xml:"configuration_file"`

	DocumentationFile *oval_sc.EntityItemBoolType `xml:"documentation_file"`

	GhostFile *oval_sc.EntityItemBoolType `xml:"ghost_file"`

	LicenseFile *oval_sc.EntityItemBoolType `xml:"license_file"`

	ReadmeFile *oval_sc.EntityItemBoolType `xml:"readme_file"`

	Message []oval.MessageType `xml:"message"`
}

// Element
type RpmverifypackageItem struct {
	XMLName xml.Name `xml:rpmverifypackage_item`

	Id string `xml:"id,attr"`

	Status string `xml:"status,attr"`

	Name *oval_sc.EntityItemStringType `xml:"name"`

	Epoch *RpmverifypackageItemEpoch `xml:"epoch"`

	Version *RpmverifypackageItemVersion `xml:"version"`

	Release *RpmverifypackageItemRelease `xml:"release"`

	Arch *oval_sc.EntityItemStringType `xml:"arch"`

	ExtendedName *oval_sc.EntityItemStringType `xml:"extended_name"`

	DependencyCheckPassed *oval_sc.EntityItemBoolType `xml:"dependency_check_passed"`

	DigestCheckPassed *oval_sc.EntityItemBoolType `xml:"digest_check_passed"`

	VerificationScriptSuccessful *oval_sc.EntityItemBoolType `xml:"verification_script_successful"`

	SignatureCheckPassed *oval_sc.EntityItemBoolType `xml:"signature_check_passed"`

	Message []oval.MessageType `xml:"message"`
}

// Element
type SelinuxbooleanItem struct {
	XMLName xml.Name `xml:selinuxboolean_item`

	Id string `xml:"id,attr"`

	Status string `xml:"status,attr"`

	Name *oval_sc.EntityItemStringType `xml:"name"`

	CurrentStatus *oval_sc.EntityItemBoolType `xml:"current_status"`

	PendingStatus *oval_sc.EntityItemBoolType `xml:"pending_status"`

	Message []oval.MessageType `xml:"message"`
}

// Element
type SelinuxsecuritycontextItem struct {
	XMLName xml.Name `xml:selinuxsecuritycontext_item`

	Id string `xml:"id,attr"`

	Status string `xml:"status,attr"`

	Filepath *oval_sc.EntityItemStringType `xml:"filepath"`

	Path *oval_sc.EntityItemStringType `xml:"path"`

	Filename *oval_sc.EntityItemStringType `xml:"filename"`

	Pid *oval_sc.EntityItemIntType `xml:"pid"`

	User *oval_sc.EntityItemStringType `xml:"user"`

	Role *oval_sc.EntityItemStringType `xml:"role"`

	Type *oval_sc.EntityItemStringType `xml:"type"`

	LowSensitivity *oval_sc.EntityItemStringType `xml:"low_sensitivity"`

	LowCategory *oval_sc.EntityItemStringType `xml:"low_category"`

	HighSensitivity *oval_sc.EntityItemStringType `xml:"high_sensitivity"`

	HighCategory *oval_sc.EntityItemStringType `xml:"high_category"`

	RawlowSensitivity *oval_sc.EntityItemStringType `xml:"rawlow_sensitivity"`

	RawlowCategory *oval_sc.EntityItemStringType `xml:"rawlow_category"`

	RawhighSensitivity *oval_sc.EntityItemStringType `xml:"rawhigh_sensitivity"`

	RawhighCategory *oval_sc.EntityItemStringType `xml:"rawhigh_category"`

	Message []oval.MessageType `xml:"message"`
}

// Element
type SlackwarepkginfoItem struct {
	XMLName xml.Name `xml:slackwarepkginfo_item`

	Id string `xml:"id,attr"`

	Status string `xml:"status,attr"`

	Name *oval_sc.EntityItemStringType `xml:"name"`

	Version *SlackwarepkginfoItemVersion `xml:"version"`

	Architecture *oval_sc.EntityItemStringType `xml:"architecture"`

	Revision *oval_sc.EntityItemStringType `xml:"revision"`

	Message []oval.MessageType `xml:"message"`
}

// Element
type SystemdunitdependencyItem struct {
	XMLName xml.Name `xml:systemdunitdependency_item`

	Id string `xml:"id,attr"`

	Status string `xml:"status,attr"`

	Unit *oval_sc.EntityItemStringType `xml:"unit"`

	Dependency []oval_sc.EntityItemStringType `xml:"dependency"`

	Message []oval.MessageType `xml:"message"`
}

// Element
type SystemdunitpropertyItem struct {
	XMLName xml.Name `xml:systemdunitproperty_item`

	Id string `xml:"id,attr"`

	Status string `xml:"status,attr"`

	Unit *oval_sc.EntityItemStringType `xml:"unit"`

	Property *oval_sc.EntityItemStringType `xml:"property"`

	Value []oval_sc.EntityItemAnySimpleType `xml:"value"`

	Message []oval.MessageType `xml:"message"`
}

// Element
type DpkginfoItemEpoch struct {
	XMLName xml.Name `xml:epoch`
}

// Element
type DpkginfoItemRelease struct {
	XMLName xml.Name `xml:release`
}

// Element
type DpkginfoItemVersion struct {
	XMLName xml.Name `xml:version`
}

// Element
type DpkginfoItemEvr struct {
	XMLName xml.Name `xml:evr`
}

// Element
type RpminfoItemEpoch struct {
	XMLName xml.Name `xml:epoch`
}

// Element
type RpminfoItemRelease struct {
	XMLName xml.Name `xml:release`
}

// Element
type RpminfoItemVersion struct {
	XMLName xml.Name `xml:version`
}

// Element
type RpmverifyfileItemEpoch struct {
	XMLName xml.Name `xml:epoch`
}

// Element
type RpmverifyfileItemVersion struct {
	XMLName xml.Name `xml:version`
}

// Element
type RpmverifyfileItemRelease struct {
	XMLName xml.Name `xml:release`
}

// Element
type RpmverifypackageItemEpoch struct {
	XMLName xml.Name `xml:epoch`
}

// Element
type RpmverifypackageItemVersion struct {
	XMLName xml.Name `xml:version`
}

// Element
type RpmverifypackageItemRelease struct {
	XMLName xml.Name `xml:release`
}

// Element
type SlackwarepkginfoItemVersion struct {
	XMLName xml.Name `xml:version`
}

// XSD ComplexType declarations

type EntityItemRpmVerifyResultType struct {
}

type EntityItemProtocolType struct {
}