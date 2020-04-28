package compute

// FINCLOUD_APACHE_NO_VERSION

import (
	"github.com/Azure/go-autorest/autorest"
	"io"
)

// The package's fully qualified name.
const fqdn = "github.com/samjegal/fincloud-sdk-for-go/services/compute"

// BlockStorageInstanceStatusCode enumerates the values for block storage instance status code.
type BlockStorageInstanceStatusCode string

const (
	// ATTAC ...
	ATTAC BlockStorageInstanceStatusCode = "ATTAC"
)

// PossibleBlockStorageInstanceStatusCodeValues returns an array of possible values for the BlockStorageInstanceStatusCode const type.
func PossibleBlockStorageInstanceStatusCodeValues() []BlockStorageInstanceStatusCode {
	return []BlockStorageInstanceStatusCode{ATTAC}
}

// DiskType2DetailCode enumerates the values for disk type 2 detail code.
type DiskType2DetailCode string

const (
	// HDD ...
	HDD DiskType2DetailCode = "HDD"
	// SSD ...
	SSD DiskType2DetailCode = "SSD"
)

// PossibleDiskType2DetailCodeValues returns an array of possible values for the DiskType2DetailCode const type.
func PossibleDiskType2DetailCodeValues() []DiskType2DetailCode {
	return []DiskType2DetailCode{HDD, SSD}
}

// NetworkInterfaceStatusCode enumerates the values for network interface status code.
type NetworkInterfaceStatusCode string

const (
	// NOTUSED ...
	NOTUSED NetworkInterfaceStatusCode = "NOTUSED"
	// USED ...
	USED NetworkInterfaceStatusCode = "USED"
)

// PossibleNetworkInterfaceStatusCodeValues returns an array of possible values for the NetworkInterfaceStatusCode const type.
func PossibleNetworkInterfaceStatusCodeValues() []NetworkInterfaceStatusCode {
	return []NetworkInterfaceStatusCode{NOTUSED, USED}
}

// ProtocolTypeCode enumerates the values for protocol type code.
type ProtocolTypeCode string

const (
	// ICMP ...
	ICMP ProtocolTypeCode = "ICMP"
	// TCP ...
	TCP ProtocolTypeCode = "TCP"
	// UDP ...
	UDP ProtocolTypeCode = "UDP"
)

// PossibleProtocolTypeCodeValues returns an array of possible values for the ProtocolTypeCode const type.
func PossibleProtocolTypeCodeValues() []ProtocolTypeCode {
	return []ProtocolTypeCode{ICMP, TCP, UDP}
}

// SecurityRuleStatusCode enumerates the values for security rule status code.
type SecurityRuleStatusCode string

const (
	// INIT ...
	INIT SecurityRuleStatusCode = "INIT"
	// RUN ...
	RUN SecurityRuleStatusCode = "RUN"
	// SET ...
	SET SecurityRuleStatusCode = "SET"
)

// PossibleSecurityRuleStatusCodeValues returns an array of possible values for the SecurityRuleStatusCode const type.
func PossibleSecurityRuleStatusCodeValues() []SecurityRuleStatusCode {
	return []SecurityRuleStatusCode{INIT, RUN, SET}
}

// ServerInstanceStatus enumerates the values for server instance status.
type ServerInstanceStatus string

const (
	// 반납중 ...
	반납중 ServerInstanceStatus = "반납중"
	// 부팅중 ...
	부팅중 ServerInstanceStatus = "부팅중"
	// 생성중 ...
	생성중 ServerInstanceStatus = "생성중"
	// 설정중 ...
	설정중 ServerInstanceStatus = "설정중"
	// 운영중 ...
	운영중 ServerInstanceStatus = "운영중"
	// 종료중 ...
	종료중 ServerInstanceStatus = "종료중"
)

// PossibleServerInstanceStatusValues returns an array of possible values for the ServerInstanceStatus const type.
func PossibleServerInstanceStatusValues() []ServerInstanceStatus {
	return []ServerInstanceStatus{반납중, 부팅중, 생성중, 설정중, 운영중, 종료중}
}

// ServerInstanceStatusCode enumerates the values for server instance status code.
type ServerInstanceStatusCode string

const (
	// ServerInstanceStatusCodeCREAT ...
	ServerInstanceStatusCodeCREAT ServerInstanceStatusCode = "CREAT"
	// ServerInstanceStatusCodeINIT ...
	ServerInstanceStatusCodeINIT ServerInstanceStatusCode = "INIT"
	// ServerInstanceStatusCodeNSTOP ...
	ServerInstanceStatusCodeNSTOP ServerInstanceStatusCode = "NSTOP"
	// ServerInstanceStatusCodeRUN ...
	ServerInstanceStatusCodeRUN ServerInstanceStatusCode = "RUN"
)

// PossibleServerInstanceStatusCodeValues returns an array of possible values for the ServerInstanceStatusCode const type.
func PossibleServerInstanceStatusCodeValues() []ServerInstanceStatusCode {
	return []ServerInstanceStatusCode{ServerInstanceStatusCodeCREAT, ServerInstanceStatusCodeINIT, ServerInstanceStatusCodeNSTOP, ServerInstanceStatusCodeRUN}
}

// ServerInstanceStatusName enumerates the values for server instance status name.
type ServerInstanceStatusName string

const (
	// Booting ...
	Booting ServerInstanceStatusName = "booting"
	// Creating ...
	Creating ServerInstanceStatusName = "creating"
	// Init ...
	Init ServerInstanceStatusName = "init"
	// Running ...
	Running ServerInstanceStatusName = "running"
	// Settingup ...
	Settingup ServerInstanceStatusName = "setting up"
	// Shuttingdown ...
	Shuttingdown ServerInstanceStatusName = "shutting down"
	// Stopped ...
	Stopped ServerInstanceStatusName = "stopped"
	// Terminating ...
	Terminating ServerInstanceStatusName = "terminating"
)

// PossibleServerInstanceStatusNameValues returns an array of possible values for the ServerInstanceStatusName const type.
func PossibleServerInstanceStatusNameValues() []ServerInstanceStatusName {
	return []ServerInstanceStatusName{Booting, Creating, Init, Running, Settingup, Shuttingdown, Stopped, Terminating}
}

// ServerOperationCode enumerates the values for server operation code.
type ServerOperationCode string

const (
	// NULL ...
	NULL ServerOperationCode = "NULL"
	// RESTA ...
	RESTA ServerOperationCode = "RESTA"
	// SHTDN ...
	SHTDN ServerOperationCode = "SHTDN"
	// START ...
	START ServerOperationCode = "START"
	// TERMT ...
	TERMT ServerOperationCode = "TERMT"
)

// PossibleServerOperationCodeValues returns an array of possible values for the ServerOperationCode const type.
func PossibleServerOperationCodeValues() []ServerOperationCode {
	return []ServerOperationCode{NULL, RESTA, SHTDN, START, TERMT}
}

// ServerStatusCode enumerates the values for server status code.
type ServerStatusCode string

const (
	// ServerStatusCodeNOTUSED ...
	ServerStatusCodeNOTUSED ServerStatusCode = "NOTUSED"
	// ServerStatusCodeSET ...
	ServerStatusCodeSET ServerStatusCode = "SET"
	// ServerStatusCodeUNSET ...
	ServerStatusCodeUNSET ServerStatusCode = "UNSET"
)

// PossibleServerStatusCodeValues returns an array of possible values for the ServerStatusCode const type.
func PossibleServerStatusCodeValues() []ServerStatusCode {
	return []ServerStatusCode{ServerStatusCodeNOTUSED, ServerStatusCodeSET, ServerStatusCodeUNSET}
}

// StatusCode enumerates the values for status code.
type StatusCode string

const (
	// StatusCodeCREATING ...
	StatusCodeCREATING StatusCode = "CREATING"
	// StatusCodeRUN ...
	StatusCodeRUN StatusCode = "RUN"
	// StatusCodeTERMTING ...
	StatusCodeTERMTING StatusCode = "TERMTING"
)

// PossibleStatusCodeValues returns an array of possible values for the StatusCode const type.
func PossibleStatusCodeValues() []StatusCode {
	return []StatusCode{StatusCodeCREATING, StatusCodeRUN, StatusCodeTERMTING}
}

// StorageStatusCode enumerates the values for storage status code.
type StorageStatusCode string

const (
	// StorageStatusCodeATTAC ...
	StorageStatusCodeATTAC StorageStatusCode = "ATTAC"
	// StorageStatusCodeCREAT ...
	StorageStatusCodeCREAT StorageStatusCode = "CREAT"
)

// PossibleStorageStatusCodeValues returns an array of possible values for the StorageStatusCode const type.
func PossibleStorageStatusCodeValues() []StorageStatusCode {
	return []StorageStatusCode{StorageStatusCodeATTAC, StorageStatusCodeCREAT}
}

// StorageStatusName enumerates the values for storage status name.
type StorageStatusName string

const (
	// Attached ...
	Attached StorageStatusName = "attached"
	// Attaching ...
	Attaching StorageStatusName = "attaching"
	// Detaching ...
	Detaching StorageStatusName = "detaching"
	// Initialized ...
	Initialized StorageStatusName = "initialized"
)

// PossibleStorageStatusNameValues returns an array of possible values for the StorageStatusName const type.
func PossibleStorageStatusNameValues() []StorageStatusName {
	return []StorageStatusName{Attached, Attaching, Detaching, Initialized}
}

// BlockStorageContentParameter server block storage 전체 검색 결과 리스트
type BlockStorageContentParameter struct {
	autorest.Response `json:"-"`
	// Content - Server block storage 컨텐츠 리스트
	Content *[]ServerBlockStorageContentParameterProperties `json:"content,omitempty"`
	// Total - 전체 server의 block storage 개수
	Total *int32 `json:"total,omitempty"`
}

// ErrorMessageParameter ...
type ErrorMessageParameter struct {
	autorest.Response `json:"-"`
	// Error - 에러 상세정보
	Error *ErrorMessageProperties `json:"error,omitempty"`
}

// ErrorMessageProperties ...
type ErrorMessageProperties struct {
	// ErrorCode - 에러 코드
	ErrorCode *string `json:"errorCode,omitempty"`
	// Message - 에러 메시지
	Message *string `json:"message,omitempty"`
	// OriginCode - 원본 코드
	OriginCode *string `json:"originCode,omitempty"`
	// OriginMessage - 원본 메시지
	OriginMessage *string `json:"originMessage,omitempty"`
}

// InitScriptContentProperties ...
type InitScriptContentProperties struct {
	// InitConfigurationScriptNo - Init Script 번호
	InitConfigurationScriptNo *string `json:"initConfigurationScriptNo,omitempty"`
	// InitConfigurationScriptType - Init Script 타입
	InitConfigurationScriptType *string `json:"initConfigurationScriptType,omitempty"`
	// InitConfigurationScriptName - Init Script 이름
	InitConfigurationScriptName *string `json:"initConfigurationScriptName,omitempty"`
	// OsTypeCode - OS 타입 코드
	OsTypeCode *string `json:"osTypeCode,omitempty"`
	// CreateYmdt - Init Script 생성 시간
	CreateYmdt *float64 `json:"createYmdt,omitempty"`
	// LastModifyYmdt - Init Script 마지막 수정 시간
	LastModifyYmdt *float64 `json:"lastModifyYmdt,omitempty"`
	// UseYn - Init Script 사용 유무
	UseYn *string `json:"useYn,omitempty"`
	// Disabled - Init Script 비활성화 유무
	Disabled *bool `json:"disabled,omitempty"`
}

// InitScriptDetailParameters init Script 상세정보 결과
type InitScriptDetailParameters struct {
	autorest.Response `json:"-"`
	// Content - Init Script 컨텐츠 정보
	Content *InitScriptContentProperties `json:"content,omitempty"`
}

// InitScriptListParameter init Script 전체 검색 결과 리스트
type InitScriptListParameter struct {
	autorest.Response `json:"-"`
	// Content - Init Script 컨텐츠 리스트
	Content *[]InitScriptContentProperties `json:"content,omitempty"`
	// Total - 전체 Init Script의 개수
	Total *int32 `json:"total,omitempty"`
}

// InitScriptNumberListParameter ...
type InitScriptNumberListParameter struct {
	// InitConfigurationScriptNoList - Init Script 번호 리스트
	InitConfigurationScriptNoList *[]string `json:"initConfigurationScriptNoList,omitempty"`
}

// InitScriptNumberParameter ...
type InitScriptNumberParameter struct {
	// InitConfigurationScriptNo - Init Script 번호
	InitConfigurationScriptNo *string `json:"initConfigurationScriptNo,omitempty"`
	// InitConfigurationScriptType - Init Script 타입
	InitConfigurationScriptType *string `json:"initConfigurationScriptType,omitempty"`
	// InitConfigurationScriptContent - Init Script 내용
	InitConfigurationScriptContent *string `json:"initConfigurationScriptContent,omitempty"`
	// OsTypeCode - OS 타입
	OsTypeCode *string `json:"osTypeCode,omitempty"`
	// UseYn - Init Script 사용 유무
	UseYn *string `json:"useYn,omitempty"`
}

// InitScriptParameter ...
type InitScriptParameter struct {
	// InitConfigurationScriptName - Init Script 이름
	InitConfigurationScriptName *string `json:"initConfigurationScriptName,omitempty"`
	// InitConfigurationScriptType - Init Script 타입
	InitConfigurationScriptType *string `json:"initConfigurationScriptType,omitempty"`
	// InitConfigurationScriptContent - Init Script 내용
	InitConfigurationScriptContent *string `json:"initConfigurationScriptContent,omitempty"`
	// OsTypeCode - OS 타입
	OsTypeCode *string `json:"osTypeCode,omitempty"`
	// UseYn - Init Script 사용 유무
	UseYn *string `json:"useYn,omitempty"`
	// InitConfigurationScriptDescription - Init Script 설명
	InitConfigurationScriptDescription *string `json:"initConfigurationScriptDescription,omitempty"`
}

// LoginKeyContentParameter login key 전체 검색 결과 리스트
type LoginKeyContentParameter struct {
	autorest.Response `json:"-"`
	// Content - Login key 컨텐츠 리스트
	Content *[]LoginKeyContentProperties `json:"content,omitempty"`
}

// LoginKeyContentProperties ...
type LoginKeyContentProperties struct {
	// KeyName - Login key 이름
	KeyName *string `json:"keyName,omitempty"`
	// Fingerprint - Login fingerprint 정보
	Fingerprint *string `json:"fingerprint,omitempty"`
	// RegistYmdt - Login 등록 일자
	RegistYmdt *float64 `json:"registYmdt,omitempty"`
}

// LoginKeyListParameter login key 전체 리스트
type LoginKeyListParameter struct {
	// KeyNameList - Login key 리스트
	KeyNameList *[]string `json:"keyNameList,omitempty"`
}

// NetworkInterfaceAttachableContentParameter server attachable network interface 전체 검색 결과 리스트
type NetworkInterfaceAttachableContentParameter struct {
	autorest.Response `json:"-"`
	// Content - Server attachable network interface 컨텐츠 리스트
	Content *[]ServerAttachableNetworkInterfaceContentParameterProperties `json:"content,omitempty"`
	// Total - 전체 server의 attachable network interface 개수
	Total *int32 `json:"total,omitempty"`
}

// NetworkInterfaceAttachableListParameter network interface attachable 전체 검색 결과 리스트
type NetworkInterfaceAttachableListParameter struct {
	autorest.Response `json:"-"`
	// Content - Network interface attachable 컨텐츠 리스트
	Content *[]NetworkInterfaceAttachableListProperties `json:"content,omitempty"`
	// Total - 전체 network interface attachable의 개수
	Total *int32 `json:"total,omitempty"`
}

// NetworkInterfaceAttachableListProperties ...
type NetworkInterfaceAttachableListProperties struct {
	// NetworkInterfaceNo - Network interface 번호
	NetworkInterfaceNo *int32 `json:"networkInterfaceNo,omitempty"`
	// OverlayIP - Overlay IP 주소
	OverlayIP *string `json:"overlayIp,omitempty"`
	// NetworkInterfaceName - Network interface 이름
	NetworkInterfaceName *string `json:"networkInterfaceName,omitempty"`
}

// NetworkInterfaceContentParameter server network interface 전체 검색 결과 리스트
type NetworkInterfaceContentParameter struct {
	autorest.Response `json:"-"`
	// Content - Server network interface 컨텐츠 리스트
	Content *[]ServerNetworkInterfaceContentParameterProperties `json:"content,omitempty"`
	// Total - 전체 server의 network interface 개수
	Total *int32 `json:"total,omitempty"`
}

// NetworkInterfaceContentProperties ...
type NetworkInterfaceContentProperties struct {
	// DefaultYn - 기본 network interface 여부
	DefaultYn *string `json:"defaultYn,omitempty"`
	// Order - Network interface 순서
	Order *int32 `json:"order,omitempty"`
	// NetworkInterfaceNo - Network interface 번호
	NetworkInterfaceNo *int32 `json:"networkInterfaceNo,omitempty"`
	// SubnetNo - Subnet 번호
	SubnetNo *string `json:"subnetNo,omitempty"`
	// IgwYn - Internet gateway 여부
	IgwYn *string `json:"igwYn,omitempty"`
	// OverlayIP - Overlay IP 주소
	OverlayIP *string `json:"overlayIp,omitempty"`
	// AttachedDeviceName - Network interface 디바이스 이름
	AttachedDeviceName *string `json:"attachedDeviceName,omitempty"`
	// VpcName - VPC 이름
	VpcName *string `json:"vpcName,omitempty"`
	// MacAddress - MAC 주소
	MacAddress *string `json:"macAddress,omitempty"`
	// VpcNo - VPC 번호
	VpcNo *string `json:"vpcNo,omitempty"`
	// SubnetName - Subnet 이름
	SubnetName *string `json:"subnetName,omitempty"`
	// ZoneNo - Zone 번호
	ZoneNo *int32 `json:"zoneNo,omitempty"`
	// Subnet - Subnet CIDR 주소
	Subnet *string `json:"subnet,omitempty"`
	// NetworkInterfaceName - Network interface 이름
	NetworkInterfaceName *string `json:"networkInterfaceName,omitempty"`
	// InstanceNo - Instance 번호
	InstanceNo *int32 `json:"instanceNo,omitempty"`
	// ServerName - Server 이름
	ServerName *string `json:"serverName,omitempty"`
	// InstanceTypeCode - Instance 타입
	InstanceTypeCode *string `json:"instanceTypeCode,omitempty"`
	// IPAllotementBlockNo - IP allotement block 번호
	IPAllotementBlockNo *int32 `json:"ipAllotementBlockNo,omitempty"`
	// StatusCode - Network interface 상태 코드. Possible values include: 'NOTUSED', 'USED'
	StatusCode NetworkInterfaceStatusCode `json:"statusCode,omitempty"`
	// AllocatedYmdt - Network interface 할당일자
	AllocatedYmdt *float64 `json:"allocatedYmdt,omitempty"`
	// DeleteOnTerminationYn - 반납시 삭제여부
	DeleteOnTerminationYn *string `json:"deleteOnTerminationYn,omitempty"`
	// BmYn - Bare Metal 사용 여부
	BmYn *string `json:"bmYn,omitempty"`
	// AccessControlGroups - Network interface에 연결된 ACG 정보
	AccessControlGroups *[]NetworkInterfaceSecurityGroupsProperties `json:"accessControlGroups,omitempty"`
}

// NetworkInterfaceListParameter network interface 전체 검색 결과 리스트
type NetworkInterfaceListParameter struct {
	autorest.Response `json:"-"`
	// Content - Network interface 컨텐츠 리스트
	Content *[]NetworkInterfaceContentProperties `json:"content,omitempty"`
	// Total - 전체 network interface의 개수
	Total *int32 `json:"total,omitempty"`
}

// NetworkInterfaceParameter network interface 생성 파라미터
type NetworkInterfaceParameter struct {
	// AccessControlGroups - Network interface에 연결된 ACG 정보
	AccessControlGroups *[]NetworkInterfaceSecurityGroupsProperties `json:"accessControlGroups,omitempty"`
	// AccessControlGroupNoList - Network interface에 연결된 ACG 정보 리스트
	AccessControlGroupNoList *[]int32 `json:"accessControlGroupNoList,omitempty"`
	// Description - Network interface 설명
	Description *string `json:"description,omitempty"`
	// NetworkInterfaceName - Network interface 이름
	NetworkInterfaceName *string `json:"networkInterfaceName,omitempty"`
	// SubnetNo - Subnet 이름
	SubnetNo *string `json:"subnetNo,omitempty"`
	// VpcNo - VPC 이름
	VpcNo *string `json:"vpcNo,omitempty"`
	// BmYn - Bare Metal 사용 여부
	BmYn *string `json:"bmYn,omitempty"`
	// OverlayIP - Network interface에 할당할 IP 주소
	OverlayIP *string `json:"overlayIp,omitempty"`
}

// NetworkInterfaceSearchFilterProperties ...
type NetworkInterfaceSearchFilterProperties struct {
	// Field - 필터에 적용할 필드 이름
	Field *string `json:"field,omitempty"`
	// Test - 테스트 영역 (TBD)
	Test *string `json:"test,omitempty"`
}

// NetworkInterfaceSearchParameter ...
type NetworkInterfaceSearchParameter struct {
	// PageNo - 검색할 network interface 페이지 번호
	PageNo *int32 `json:"pageNo,omitempty"`
	// PageSizeNo - 한 페이지에 나올 network interface 개수
	PageSizeNo *int32 `json:"pageSizeNo,omitempty"`
	// Filter - Network interface 검색 필터 리스트
	Filter *[]NetworkInterfaceSearchFilterProperties `json:"filter,omitempty"`
}

// NetworkInterfaceSecurityGroupsProperties ...
type NetworkInterfaceSecurityGroupsProperties struct {
	// AccessControlGroupName - ACG 이름
	AccessControlGroupName *string `json:"accessControlGroupName,omitempty"`
	// AccessControlGroupNo - ACG 번호
	AccessControlGroupNo *int32 `json:"accessControlGroupNo,omitempty"`
	// StatusCode - ACG 상태 코드. Possible values include: 'INIT', 'SET', 'RUN'
	StatusCode SecurityRuleStatusCode `json:"statusCode,omitempty"`
}

// NetworkInterfaceSubnetParameter ...
type NetworkInterfaceSubnetParameter struct {
	// SubnetNo - Subnet 번호
	SubnetNo *string `json:"subnetNo,omitempty"`
}

// NetworkInterfaceValidationParameter network interface에 할당할 IP 주소 검증 결과
type NetworkInterfaceValidationParameter struct {
	autorest.Response `json:"-"`
	// Content - Network interface IP 주소 검증 결과
	Content *bool `json:"content,omitempty"`
}

// PublicIPAddressSearchFilterParameter ...
type PublicIPAddressSearchFilterParameter struct {
	// PageNo - 검색할 Public IP 페이지 번호
	PageNo *int32 `json:"pageNo,omitempty"`
	// PageSizeNo - 한 페이지에 나올 Public IP 개수
	PageSizeNo *int32 `json:"pageSizeNo,omitempty"`
	// Filter - Public IP 검색 필터 리스트
	Filter *[]PublicIPSearchFilterProperties `json:"filter,omitempty"`
}

// PublicIPAddressSearchParameter public IP 전체 검색 결과 리스트
type PublicIPAddressSearchParameter struct {
	autorest.Response `json:"-"`
	// Content - Public IP 컨텐츠 리스트
	Content *[]PublicIPAddressSearchProperties `json:"content,omitempty"`
	// Total - 전체 Public IP의 개수
	Total *int32 `json:"total,omitempty"`
	// UserRequestID - Public IP의 Request 고유 아이디
	UserRequestID *string `json:"userRequestId,omitempty"`
}

// PublicIPAddressSearchProperties ...
type PublicIPAddressSearchProperties struct {
	// Disabled - 사용 가능여부
	Disabled *bool `json:"disabled,omitempty"`
	// InstanceNo - Public IP 인스턴스 번호
	InstanceNo *string `json:"instanceNo,omitempty"`
	// PublicIP - Public IP 주소
	PublicIP *string `json:"publicIp,omitempty"`
	// StatusCode - Public IP 상태코드. Possible values include: 'StatusCodeCREATING', 'StatusCodeRUN', 'StatusCodeTERMTING'
	StatusCode StatusCode `json:"statusCode,omitempty"`
	// VpcNo - VPC 번호
	VpcNo *string `json:"vpcNo,omitempty"`
	// VpcName - VPC 이름
	VpcName *string `json:"vpcName,omitempty"`
	// PrivateIP - Private Ip 주소
	PrivateIP *string `json:"privateIp,omitempty"`
	// ServerInstanceNo - 서버 인스턴스 번호
	ServerInstanceNo *int32 `json:"serverInstanceNo,omitempty"`
	// ServerInstanceName - 서버 인스턴스 이름
	ServerInstanceName *string `json:"serverInstanceName,omitempty"`
	// InstanceTypeCode - Public IP 인스턴스 타입 코드
	InstanceTypeCode *string `json:"instanceTypeCode,omitempty"`
	// StatusName - Public IP 상태 이름
	StatusName *string `json:"statusName,omitempty"`
	// CreatedYmdt - Public IP 생성일자
	CreatedYmdt *float64 `json:"createdYmdt,omitempty"`
}

// PublicIPAddressServerInstanceParameter ...
type PublicIPAddressServerInstanceParameter struct {
	// ServerInstanceNo - 서버 인스턴스 번호
	ServerInstanceNo *string `json:"serverInstanceNo,omitempty"`
	// Description - 공인 IP 주소 정보
	Description *string `json:"description,omitempty"`
}

// PublicIPAddressServerListParameter public IP 전체 서버 리스트
type PublicIPAddressServerListParameter struct {
	autorest.Response `json:"-"`
	// Content - Public IP 서버 컨텐츠 리스트
	Content *[]PublicIPAddressServerListProperties `json:"content,omitempty"`
	// Total - 전체 Public IP의 개수
	Total *int32 `json:"total,omitempty"`
	// UserRequestID - Public IP의 Request 고유 아이디
	UserRequestID *string `json:"userRequestId,omitempty"`
}

// PublicIPAddressServerListProperties ...
type PublicIPAddressServerListProperties struct {
	// InstanceNo - Public IP 인스턴스 번호
	InstanceNo *string `json:"instanceNo,omitempty"`
	// ServerName - 서버 이름
	ServerName *string `json:"serverName,omitempty"`
	// Disabled - Public IP 사용 가능여부
	Disabled *bool `json:"disabled,omitempty"`
	// ActionName - Action 이름
	ActionName *string `json:"actionName,omitempty"`
	// Permission - 권한 상태정보
	Permission *string `json:"permission,omitempty"`
}

// PublicIPAddressSummaryParameter public IP 사용가능한 서버 요약정보
type PublicIPAddressSummaryParameter struct {
	autorest.Response `json:"-"`
	// Content - 공인 IP 주소 요약정보
	Content *PublicIPAddressSummaryProperties `json:"content,omitempty"`
}

// PublicIPAddressSummaryProperties ...
type PublicIPAddressSummaryProperties struct {
	// CountOfServersWithPublicSubnet - Public 서브넷의 서버 수
	CountOfServersWithPublicSubnet *int32 `json:"countOfServersWithPublicSubnet,omitempty"`
}

// PublicIPSearchFilterProperties ...
type PublicIPSearchFilterProperties struct {
	// Field - 필터에 적용할 필드 이름
	Field *string `json:"field,omitempty"`
	// Text - 텍스트 영역 (TBD)
	Text *string `json:"text,omitempty"`
}

// ReadCloser ...
type ReadCloser struct {
	autorest.Response `json:"-"`
	Value             *io.ReadCloser `json:"value,omitempty"`
}

// RootPasswordContentParameter server root password 정보
type RootPasswordContentParameter struct {
	autorest.Response `json:"-"`
	// Content - Server root password 컨텐츠
	Content *ServerRootPasswordContentParameterProperties `json:"content,omitempty"`
}

// SecurityGroupContentListParameter ACG name 검색 리스트
type SecurityGroupContentListParameter struct {
	autorest.Response `json:"-"`
	// Content - ACG name 검색 컨텐츠 리스트
	Content *[]SecurityGroupContentNameParameter `json:"content,omitempty"`
	// Total - ACG name 전체 개수
	Total *int32 `json:"total,omitempty"`
}

// SecurityGroupContentNameParameter ...
type SecurityGroupContentNameParameter struct {
	// AccessControlGroupName - ACG 이름
	AccessControlGroupName *string `json:"accessControlGroupName,omitempty"`
	// AccessControlGroupNo - ACG 번호
	AccessControlGroupNo *int32 `json:"accessControlGroupNo,omitempty"`
	// VpcNo - VPC 번호
	VpcNo *string `json:"vpcNo,omitempty"`
}

// SecurityGroupListContentParameter ...
type SecurityGroupListContentParameter struct {
	// AccessControlGroupName - ACG 이름
	AccessControlGroupName *string `json:"accessControlGroupName,omitempty"`
	// AccessControlGroupNo - ACG 번호
	AccessControlGroupNo *int32 `json:"accessControlGroupNo,omitempty"`
	// CreatedYmdt - ACG 생성일자
	CreatedYmdt *float64 `json:"createdYmdt,omitempty"`
	// DefaultYn - 기본 ACG 여부
	DefaultYn *string `json:"defaultYn,omitempty"`
	// ModifiedYmdt - ACG 수정일자
	ModifiedYmdt *float64 `json:"modifiedYmdt,omitempty"`
	// NetworkInterfaceCount - ACG에 연결된 네트워크 인터페이스 개수
	NetworkInterfaceCount *int32 `json:"networkInterfaceCount,omitempty"`
	// StatusCode - ACG 상태 코드. Possible values include: 'INIT', 'SET', 'RUN'
	StatusCode SecurityRuleStatusCode `json:"statusCode,omitempty"`
	// VpcName - VPC 이름
	VpcName *string `json:"vpcName,omitempty"`
	// VpcNo - VPC 번호
	VpcNo *string `json:"vpcNo,omitempty"`
}

// SecurityGroupListParameter ACG 전체 리스트
type SecurityGroupListParameter struct {
	autorest.Response `json:"-"`
	// Content - ACG 전체 컨텐츠 리스트
	Content *[]SecurityGroupListContentParameter `json:"content,omitempty"`
	// Total - ACG 전체 개수
	Total *int32 `json:"total,omitempty"`
}

// SecurityGroupNameParameter ...
type SecurityGroupNameParameter struct {
	// AccessControlGroupName - ACG 이름
	AccessControlGroupName *string `json:"accessControlGroupName,omitempty"`
	// VpcNo - VPC 번호
	VpcNo *string `json:"vpcNo,omitempty"`
	// Description - ACG 설명
	Description *string `json:"description,omitempty"`
}

// SecurityGroupNumberParameter ...
type SecurityGroupNumberParameter struct {
	// AccessControlGroupNo - ACG 번호
	AccessControlGroupNo *string `json:"accessControlGroupNo,omitempty"`
	// VpcNo - VPC 번호
	VpcNo *string `json:"vpcNo,omitempty"`
}

// SecurityGroupParameter ...
type SecurityGroupParameter struct {
	// AccessControlGroupNo - ACG 번호
	AccessControlGroupNo *int32 `json:"accessControlGroupNo,omitempty"`
	// AccessControlGroupRuleNo - ACG rule 번호
	AccessControlGroupRuleNo *int32 `json:"accessControlGroupRuleNo,omitempty"`
	// AccessControlGroupName - ACG 이름
	AccessControlGroupName *string `json:"accessControlGroupName,omitempty"`
	// AccessControlGroupSequence - ACG 번호
	AccessControlGroupSequence *int32 `json:"accessControlGroupSequence,omitempty"`
	// CreatedYmdt - ACG rule 생성일자
	CreatedYmdt *float64 `json:"createdYmdt,omitempty"`
	// DefaultYn - 기본 ACG 여부
	DefaultYn *string `json:"defaultYn,omitempty"`
	// NetworkInterfaceCount - ACG에 연결된 네트워크 인터페이스 개수
	NetworkInterfaceCount *int32 `json:"networkInterfaceCount,omitempty"`
	// VpcName - VPC 이름
	VpcName *string `json:"vpcName,omitempty"`
	// VpcNo - VPC 번호
	VpcNo *string `json:"vpcNo,omitempty"`
	// Description - ACG rule 설명
	Description *string `json:"description,omitempty"`
	// IPBlock - ACG IP 차단 CIDR 주소값
	IPBlock *string `json:"ipBlock,omitempty"`
	// IsInboundRule - ACG inbound rule 여부
	IsInboundRule *bool `json:"isInboundRule,omitempty"`
	// ModifiedYmdt - ACG rule 수정일자
	ModifiedYmdt *float64 `json:"modifiedYmdt,omitempty"`
	// PortRange - 포트 범위
	PortRange *string `json:"portRange,omitempty"`
	// ProtocolTypeCode - 프로토콜 타입 코드. Possible values include: 'ICMP', 'UDP', 'TCP'
	ProtocolTypeCode ProtocolTypeCode `json:"protocolTypeCode,omitempty"`
	// StatusCode - ACG 상태 코드. Possible values include: 'INIT', 'SET', 'RUN'
	StatusCode SecurityRuleStatusCode `json:"statusCode,omitempty"`
}

// SecurityGroupRuleContentParameter ACG rule 전체 리스트
type SecurityGroupRuleContentParameter struct {
	autorest.Response `json:"-"`
	// Content - ACG rule 컨텐츠 리스트
	Content *[]SecurityGroupParameter `json:"content,omitempty"`
	// Total - ACG rule 전체 개수
	Total *int32 `json:"total,omitempty"`
}

// SecurityGroupRuleProperties ...
type SecurityGroupRuleProperties struct {
	// ProtocolTypeCode - 프로토콜 타입 코드. Possible values include: 'ICMP', 'UDP', 'TCP'
	ProtocolTypeCode ProtocolTypeCode `json:"protocolTypeCode,omitempty"`
	// PortRange - 포트 범위
	PortRange *string `json:"portRange,omitempty"`
	// Description - ACG rule 설명
	Description *string `json:"description,omitempty"`
	// AccessControlGroupSequence - ACG 순번
	AccessControlGroupSequence *int32 `json:"accessControlGroupSequence,omitempty"`
	// AccessControlGroupName - ACG 이름
	AccessControlGroupName *string `json:"accessControlGroupName,omitempty"`
	// IPBlock - ACG IP 차단 CIDR 주소값
	IPBlock *string `json:"ipBlock,omitempty"`
	// AccessControlGroupNo - ACG 번호
	AccessControlGroupNo *int32 `json:"accessControlGroupNo,omitempty"`
	// AccessControlGroupRuleNo - ACG rule 번호
	AccessControlGroupRuleNo *int32 `json:"accessControlGroupRuleNo,omitempty"`
	// IsInboundRule - ACG inbound rule 여부
	IsInboundRule *bool `json:"isInboundRule,omitempty"`
	// StatusCode - ACG 상태 코드. Possible values include: 'INIT', 'SET', 'RUN'
	StatusCode SecurityRuleStatusCode `json:"statusCode,omitempty"`
}

// SecurityGroupRulesProperties ...
type SecurityGroupRulesProperties struct {
	// VpcNo - VPC 번호
	VpcNo *string `json:"vpcNo,omitempty"`
	// AccessControlGroupInboundRules - ACG inbound rule 컨텐츠 리스트
	AccessControlGroupInboundRules *[]SecurityGroupRuleProperties `json:"accessControlGroupInboundRules,omitempty"`
	// AccessControlGroupOutboundRules - ACG inbound rule 컨텐츠 리스트
	AccessControlGroupOutboundRules *[]SecurityGroupRuleProperties `json:"accessControlGroupOutboundRules,omitempty"`
}

// ServerAttachableNetworkInterfaceContentParameterProperties ...
type ServerAttachableNetworkInterfaceContentParameterProperties struct {
	// NetworkInterfaceNo - Network interface 번호
	NetworkInterfaceNo *int32 `json:"networkInterfaceNo,omitempty"`
	// SubnetNo - Subnet 번호
	SubnetNo *string `json:"subnetNo,omitempty"`
	// InstanceNo - Server instance 번호
	InstanceNo *int32 `json:"instanceNo,omitempty"`
}

// ServerBlockStorageContentParameterProperties ...
type ServerBlockStorageContentParameterProperties struct {
	// BlockStorageTypeCode - Block storage 타입 코드
	BlockStorageTypeCode *string `json:"blockStorageTypeCode,omitempty"`
	// CncNo - CNC 번호
	CncNo *string `json:"cncNo,omitempty"`
	// ComputeInstanceName - Compute instance 이름
	ComputeInstanceName *string `json:"computeInstanceName,omitempty"`
	// ComputeInstanceNo - Compute instance 번호
	ComputeInstanceNo *int32 `json:"computeInstanceNo,omitempty"`
	// CreateYmdt - Block storage 생성 일자
	CreateYmdt *float64 `json:"createYmdt,omitempty"`
	// Deletable - 삭제 가능 여부
	Deletable *bool `json:"deletable,omitempty"`
	// Detachable - 할당 해제 가능 여부
	Detachable *bool `json:"detachable,omitempty"`
	// DeviceName - 장치 이름
	DeviceName *string `json:"deviceName,omitempty"`
	// DiskType2Code - Disk 타입 2 코드
	DiskType2Code *string `json:"diskType2Code,omitempty"`
	// DiskType2DetailCode - Disk 타입 2 상세 코드
	DiskType2DetailCode *string `json:"diskType2DetailCode,omitempty"`
	// InstanceDesc - Instance 설명
	InstanceDesc *string `json:"instanceDesc,omitempty"`
	// InstanceNo - Instance 번호
	InstanceNo *string `json:"instanceNo,omitempty"`
	// InstanceStatusCode - Instance 상태 코드. Possible values include: 'ATTAC'
	InstanceStatusCode BlockStorageInstanceStatusCode `json:"instanceStatusCode,omitempty"`
	// InstanceStatusName - Instance 상태 이름. Possible values include: 'Init', 'Booting', 'Creating', 'Settingup', 'Running', 'Stopped', 'Shuttingdown', 'Terminating'
	InstanceStatusName ServerInstanceStatusName `json:"instanceStatusName,omitempty"`
	// InstanceUUID - Instance UUID 번호
	InstanceUUID *string `json:"instanceUuid,omitempty"`
	// MaxIopsThroughput - IOPS 최대 성능값
	MaxIopsThroughput *int32 `json:"maxIopsThroughput,omitempty"`
	// OperationCode - 운영 코드. Possible values include: 'START', 'RESTA', 'NULL', 'SHTDN', 'TERMT'
	OperationCode ServerOperationCode `json:"operationCode,omitempty"`
	// ProductCode - 제품 코드
	ProductCode *string `json:"productCode,omitempty"`
	// Resizeable - 사이즈 변환 가능 여부
	Resizeable *bool `json:"resizeable,omitempty"`
	// SrNo - SR 번호
	SrNo *int32 `json:"srNo,omitempty"`
	// VbdUUID - VBD UUID 번호
	VbdUUID *string `json:"vbdUuid,omitempty"`
	// VolumeName - Volume 이름
	VolumeName *string `json:"volumeName,omitempty"`
	// VolumeSize - Volume 크기
	VolumeSize *float64 `json:"volumeSize,omitempty"`
	// ZoneName - Zone 이름
	ZoneName *string `json:"zoneName,omitempty"`
	// ZoneNo - Zone 번호
	ZoneNo *string `json:"zoneNo,omitempty"`
	// DeviceInfo - Device 정보
	DeviceInfo *string `json:"deviceInfo,omitempty"`
	// VolumeSizeGb - Volume 크기 (GB)
	VolumeSizeGb *string `json:"volumeSizeGb,omitempty"`
	// InstanceStatus - Instance 상태. Possible values include: '생성중', '부팅중', '설정중', '운영중', '종료중', '반납중'
	InstanceStatus ServerInstanceStatus `json:"instanceStatus,omitempty"`
	// InstanceStatusIcon - Instance 상태 아이콘
	InstanceStatusIcon *string `json:"instanceStatusIcon,omitempty"`
}

// ServerContentContractRestrictionParameter server 계약 정보
type ServerContentContractRestrictionParameter struct {
	autorest.Response `json:"-"`
	// Content - Server 계약 컨텐츠 정보
	Content *ServerContentContractRestrictionParameterProperties `json:"content,omitempty"`
}

// ServerContentContractRestrictionParameterProperties ...
type ServerContentContractRestrictionParameterProperties struct {
	// ContractRestrictionCount - 계약 제한 개수
	ContractRestrictionCount *int32 `json:"contractRestrictionCount,omitempty"`
	// IsPossible - 가능 여부
	IsPossible *bool `json:"isPossible,omitempty"`
	// ExistingContractCount - 현재 계약 개수
	ExistingContractCount *int32 `json:"existingContractCount,omitempty"`
	// ExistingContractProductCount - 제품 계약 상태 개수
	ExistingContractProductCount *int32 `json:"existingContractProductCount,omitempty"`
	// ContractProductRestrictionCount - 제품 계약 제한 개수
	ContractProductRestrictionCount *int32 `json:"contractProductRestrictionCount,omitempty"`
}

// ServerContentParameter ...
type ServerContentParameter struct {
	autorest.Response `json:"-"`
	// Content - Server 상태 체크 여부
	Content *bool `json:"content,omitempty"`
}

// ServerDetailParameter server 상세 정보
type ServerDetailParameter struct {
	autorest.Response `json:"-"`
	// Content - Server 컨텐츠 정보
	Content *ServerListParametersProperties `json:"content,omitempty"`
}

// ServerImageParameter ...
type ServerImageParameter struct {
	// OriginalCopyServerInstanceNo - 원본 복사할 server instance 번호
	OriginalCopyServerInstanceNo *string `json:"originalCopyServerInstanceNo,omitempty"`
	// PromotionNo - 프로모션 번호
	PromotionNo *string `json:"promotionNo,omitempty"`
	// MemberNsiName - Member NSI(Server) 이름
	MemberNsiName *string `json:"memberNsiName,omitempty"`
	// MemberNsiDesc - Member NSI(Server) 설명
	MemberNsiDesc *string `json:"memberNsiDesc,omitempty"`
}

// ServerInstanceListParameter ...
type ServerInstanceListParameter struct {
	// InstanceNoList - Server instance 리스트
	InstanceNoList *[]string `json:"instanceNoList,omitempty"`
	// IsToReleaseAssociatedPublicIPTogether - Public IP와 같이 해제할지 여부
	IsToReleaseAssociatedPublicIPTogether *bool `json:"isToReleaseAssociatedPublicIpTogether,omitempty"`
}

// ServerListAdditionalParameterMapProperties ...
type ServerListAdditionalParameterMapProperties struct {
	// NetworkInterfaceList - Network interface 리스트
	NetworkInterfaceList *[]ServerListNetworkInterfaceParameterProperties `json:"networkInterfaceList,omitempty"`
	// VpcNo - VPC 번호
	VpcNo *string `json:"vpcNo,omitempty"`
	// VpcName - VPC 이름
	VpcName *string `json:"vpcName,omitempty"`
	// SubnetNo - Subnet 번호
	SubnetNo *string `json:"subnetNo,omitempty"`
	// SubnetName - Subnet 이름
	SubnetName *string `json:"subnetName,omitempty"`
}

// ServerListNetworkInterfaceParameterProperties ...
type ServerListNetworkInterfaceParameterProperties struct {
	// NetworkInterfaceNo - Network interface 번호
	NetworkInterfaceNo *int32 `json:"networkInterfaceNo,omitempty"`
	// SubnetNo - Subnet 번호
	SubnetNo *string `json:"subnetNo,omitempty"`
	// IgwYn - Internet gateway 여부
	IgwYn *string `json:"igwYn,omitempty"`
	// AttachedDeviceName - 할당된 network interface 장비 이름
	AttachedDeviceName *string `json:"attachedDeviceName,omitempty"`
	// VpcName - VPC 이름
	VpcName *string `json:"vpcName,omitempty"`
	// MacAddress - MAC 주소
	MacAddress *string `json:"macAddress,omitempty"`
	// VpcNo - VPC 번호
	VpcNo *string `json:"vpcNo,omitempty"`
	// InternalIP - 내부 IP 주소
	InternalIP *string `json:"internalIp,omitempty"`
	// SubnetName - Subnet 이름
	SubnetName *string `json:"subnetName,omitempty"`
	// Subnet - Subnet CIDR 주소
	Subnet *string `json:"subnet,omitempty"`
	// NetworkInterfaceName - Network interface 이름
	NetworkInterfaceName *string `json:"networkInterfaceName,omitempty"`
}

// ServerListParameter server 전체 검색 결과 리스트
type ServerListParameter struct {
	autorest.Response `json:"-"`
	// Content - Server 컨텐츠 리스트
	Content *[]ServerListParametersProperties `json:"content,omitempty"`
	// Total - 전체 server의 개수
	Total *int32 `json:"total,omitempty"`
}

// ServerListParametersProperties ...
type ServerListParametersProperties struct {
	// InstanceNo - Server instance 번호
	InstanceNo *string `json:"instanceNo,omitempty"`
	// InstanceTypeCode - Server instance 타입 코드
	InstanceTypeCode *string `json:"instanceTypeCode,omitempty"`
	// ProductCode - Server 제품 코드
	ProductCode *string `json:"productCode,omitempty"`
	// ContractNo - Server 계약 번호
	ContractNo *string `json:"contractNo,omitempty"`
	// InstanceStatusCode - Server instance 상태 코드. Possible values include: 'ServerInstanceStatusCodeINIT', 'ServerInstanceStatusCodeCREAT', 'ServerInstanceStatusCodeRUN', 'ServerInstanceStatusCodeNSTOP'
	InstanceStatusCode ServerInstanceStatusCode `json:"instanceStatusCode,omitempty"`
	// OperationCode - Server 운영 코드. Possible values include: 'START', 'RESTA', 'NULL', 'SHTDN', 'TERMT'
	OperationCode ServerOperationCode `json:"operationCode,omitempty"`
	// InstanceStatus - Server instance 상태. Possible values include: '생성중', '부팅중', '설정중', '운영중', '종료중', '반납중'
	InstanceStatus ServerInstanceStatus `json:"instanceStatus,omitempty"`
	// CreateYmdt - Server 생성일자
	CreateYmdt *float64 `json:"createYmdt,omitempty"`
	// Uptime - Server 가동시간
	Uptime *float64 `json:"uptime,omitempty"`
	// OperationYmdt - Server 운영시간
	OperationYmdt *float64 `json:"operationYmdt,omitempty"`
	// InstanceStatusName - Server instance 상태 이름. Possible values include: 'Init', 'Booting', 'Creating', 'Settingup', 'Running', 'Stopped', 'Shuttingdown', 'Terminating'
	InstanceStatusName ServerInstanceStatusName `json:"instanceStatusName,omitempty"`
	// ServerName - Server 이름
	ServerName *string `json:"serverName,omitempty"`
	// CPUCount - Server CPU 갯수
	CPUCount *string `json:"cpuCount,omitempty"`
	// MemorySize - Server 메모리 크기
	MemorySize *float64 `json:"memorySize,omitempty"`
	// ServerInstanceTypeCode - Server instance 타입 코드
	ServerInstanceTypeCode *string `json:"serverInstanceTypeCode,omitempty"`
	// ServerSpec - Server 사양
	ServerSpec *string `json:"serverSpec,omitempty"`
	// NsiName - OS 이름
	NsiName *string `json:"nsiName,omitempty"`
	// OsInformation - OS 정보
	OsInformation *string `json:"osInformation,omitempty"`
	// AdditionalParameterMap - 추가적인 파라미터 맵
	AdditionalParameterMap *ServerListAdditionalParameterMapProperties `json:"additionalParameterMap,omitempty"`
	// PlatformTypeCode - Server platform 타입 코드
	PlatformTypeCode *string `json:"platformTypeCode,omitempty"`
	// LoginKeyName - Server 로그인 키 이름
	LoginKeyName *string `json:"loginKeyName,omitempty"`
	// RootPassword - Server 관리자 비밀번호
	RootPassword *string `json:"rootPassword,omitempty"`
	// ZoneName - Zone 이름
	ZoneName *string `json:"zoneName,omitempty"`
	// ZoneNo - Zone 번호
	ZoneNo *string `json:"zoneNo,omitempty"`
	// RegionName - Region 이름
	RegionName *string `json:"regionName,omitempty"`
	// CncNo - CNC 번호
	CncNo *string `json:"cncNo,omitempty"`
	// InstanceProductType2Code - Server instance 타입 2 코드
	InstanceProductType2Code *string `json:"instanceProductType2Code,omitempty"`
	// VpcNo - VPC 번호
	VpcNo *string `json:"vpcNo,omitempty"`
	// IsHddStorageAddable - HDD storage 추가 가능 여부
	IsHddStorageAddable *bool `json:"isHddStorageAddable,omitempty"`
	// IsSsdStorageAddable - SSD storage 추가 가능 여부
	IsSsdStorageAddable *bool `json:"isSsdStorageAddable,omitempty"`
	// SoftwareProductCode - Software 제품 코드
	SoftwareProductCode *string `json:"softwareProductCode,omitempty"`
	// OsDiskTypeCode - OS 디스크 타입 코드
	OsDiskTypeCode *string `json:"osDiskTypeCode,omitempty"`
	// OsDiskTypeDetailCode - OS 디스크 타입 상세 코드
	OsDiskTypeDetailCode *string `json:"osDiskTypeDetailCode,omitempty"`
	// ReturnProtectionYn - 반환 보호 여부
	ReturnProtectionYn *string `json:"returnProtectionYn,omitempty"`
	// SubAccountLoginID - Subaccount 로그인 아이디
	SubAccountLoginID *string `json:"subAccountLoginId,omitempty"`
	// MemberNo - Member 번호
	MemberNo *string `json:"memberNo,omitempty"`
	// RegionCode - Region 코드
	RegionCode *string `json:"regionCode,omitempty"`
	// RegionNo - Region 번호
	RegionNo *string `json:"regionNo,omitempty"`
	// IsInVpc - VPC 실장 여부
	IsInVpc *bool `json:"isInVpc,omitempty"`
}

// ServerNetworkInterfaceContentParameterProperties ...
type ServerNetworkInterfaceContentParameterProperties struct {
	// DefaultYn - Network interface 기본 여부
	DefaultYn *string `json:"defaultYn,omitempty"`
	// Order - Network interface 순서
	Order *int32 `json:"order,omitempty"`
	// NetworkInterfaceNo - Network interface 번호
	NetworkInterfaceNo *int32 `json:"networkInterfaceNo,omitempty"`
	// SubnetNo - Subnet 번호
	SubnetNo *string `json:"subnetNo,omitempty"`
	// OverlayIP - Overlay IP 주소
	OverlayIP *string `json:"overlayIp,omitempty"`
	// AttachedDeviceName - Attached 장치 이름
	AttachedDeviceName *string `json:"attachedDeviceName,omitempty"`
	// VpcName - VPC 이름
	VpcName *string `json:"vpcName,omitempty"`
	// MacAddress - MAC 주소
	MacAddress *string `json:"macAddress,omitempty"`
	// VpcNo - VPC 번호
	VpcNo *string `json:"vpcNo,omitempty"`
	// SubnetName - Subnet 이름
	SubnetName *string `json:"subnetName,omitempty"`
	// ZoneNo - Zone 번호
	ZoneNo *int32 `json:"zoneNo,omitempty"`
	// Subnet - Subnet CIDR 주소
	Subnet *string `json:"subnet,omitempty"`
	// NetworkInterfaceName - Network interface 이름
	NetworkInterfaceName *string `json:"networkInterfaceName,omitempty"`
	// InstanceNo - Server instance 번호
	InstanceNo *int32 `json:"instanceNo,omitempty"`
	// ServerName - Server 이름
	ServerName *string `json:"serverName,omitempty"`
	// InstanceTypeCode - Server instance 타입 코드
	InstanceTypeCode *string `json:"instanceTypeCode,omitempty"`
	// IPAllotementBlockNo - IP allotement block 번호
	IPAllotementBlockNo *int32 `json:"ipAllotementBlockNo,omitempty"`
	// StatusCode - Network interface 상태 코드. Possible values include: 'ServerStatusCodeSET', 'ServerStatusCodeUNSET', 'ServerStatusCodeNOTUSED'
	StatusCode ServerStatusCode `json:"statusCode,omitempty"`
	// AllocatedYmdt - Network interface 할당 일자
	AllocatedYmdt *float64 `json:"allocatedYmdt,omitempty"`
	// DeleteOnTerminationYn - Server 반납시 삭제 여부
	DeleteOnTerminationYn *string `json:"deleteOnTerminationYn,omitempty"`
}

// ServerNetworkInterfaceProperties ...
type ServerNetworkInterfaceProperties struct {
	// DefaultYn - Network interface 기본 여부
	DefaultYn *string `json:"defaultYn,omitempty"`
	// Order - Network interface 순서
	Order *int32 `json:"order,omitempty"`
	// NetworkInterfaceNo - Network interface 번호
	NetworkInterfaceNo *int32 `json:"networkInterfaceNo,omitempty"`
	// SubnetNo - Subnet 번호
	SubnetNo *string `json:"subnetNo,omitempty"`
	// OverlayIP - Overlay IP 주소
	OverlayIP *string `json:"overlayIp,omitempty"`
	// AccessControlGroups - ACG 리스트
	AccessControlGroups *[]ServerNetworkInterfaceSecurityGroupProperties `json:"accessControlGroups,omitempty"`
}

// ServerNetworkInterfaceSecurityGroupProperties ...
type ServerNetworkInterfaceSecurityGroupProperties struct {
	// AccessControlGroupNo - ACG 번호
	AccessControlGroupNo *int32 `json:"accessControlGroupNo,omitempty"`
}

// ServerParameter ...
type ServerParameter struct {
	// ReturnProtectionYn - Server 반환보호 여부
	ReturnProtectionYn *string `json:"returnProtectionYn,omitempty"`
	// ServerName - Server 이름
	ServerName *string `json:"serverName,omitempty"`
	// LoginKeyName - 로그인 키 이름
	LoginKeyName *string `json:"loginKeyName,omitempty"`
	// ServerInstanceProductCode - Server instance 제품 코드
	ServerInstanceProductCode *string `json:"serverInstanceProductCode,omitempty"`
	// SoftwareProductCode - Server 제품 코드
	SoftwareProductCode *string `json:"softwareProductCode,omitempty"`
	// RegionNo - Region 번호
	RegionNo *string `json:"regionNo,omitempty"`
	// ServerCreateCount - Server 생성 횟수
	ServerCreateCount *string `json:"serverCreateCount,omitempty"`
	// ServerDesc - Server 설명
	ServerDesc *string `json:"serverDesc,omitempty"`
	// VpcNo - VPC 번호
	VpcNo *string `json:"vpcNo,omitempty"`
	// SubnetNo - Subnet 번호
	SubnetNo *string `json:"subnetNo,omitempty"`
	// HostServerNameSameYn - Host server 이름 비교
	HostServerNameSameYn *string `json:"hostServerNameSameYn,omitempty"`
	// NetworkInterfaces - Server network interface 리스트
	NetworkInterfaces *[]ServerNetworkInterfaceProperties `json:"networkInterfaces,omitempty"`
	// InitConfigurationScriptNo - Init Script 번호
	InitConfigurationScriptNo *string `json:"initConfigurationScriptNo,omitempty"`
	// MetaTypeCode - Meta 타입 코드
	MetaTypeCode *string `json:"metaTypeCode,omitempty"`
	// ZoneNo - Zone 번호
	ZoneNo *int32 `json:"zoneNo,omitempty"`
	// FeeSystemTypeCode - Server 요금 타입 코드
	FeeSystemTypeCode *string `json:"feeSystemTypeCode,omitempty"`
	// AccessControlGroupNoList - Server에 연결된 ACG 정보 리스트
	AccessControlGroupNoList *[]int32 `json:"accessControlGroupNoList,omitempty"`
}

// ServerRootPasswordContentParameterProperties ...
type ServerRootPasswordContentParameterProperties struct {
	// RootPassword - Root password 정보
	RootPassword *string `json:"rootPassword,omitempty"`
	// DecryptRootPassword - Root password 복호화 정보
	DecryptRootPassword *string `json:"decryptRootPassword,omitempty"`
}

// ServerSearchParameter ...
type ServerSearchParameter struct {
	// PageNo - 검색할 서버 페이지 번호
	PageNo *int32 `json:"pageNo,omitempty"`
	// PageSizeNo - 한 페이지에 나올 서버 개수
	PageSizeNo *int32 `json:"pageSizeNo,omitempty"`
	// SortedBy - 서버 페이지 정렬 기준
	SortedBy *string `json:"sortedBy,omitempty"`
	// SortingOrder - 정렬 순서
	SortingOrder *string `json:"sortingOrder,omitempty"`
	// SubAccountNo - 부계정 번호
	SubAccountNo *string `json:"subAccountNo,omitempty"`
	// Owner - 서버 소유주
	Owner *bool `json:"owner,omitempty"`
	// Filter - 필터 값
	Filter interface{} `json:"filter,omitempty"`
}

// StorageAddableParameter 전체 storage의 적용 가능 서버 리스트
type StorageAddableParameter struct {
	autorest.Response `json:"-"`
	// Content - Storage 적용 가능 서버 리스트
	Content *[]StorageAddableProperties `json:"content,omitempty"`
}

// StorageAddableProperties ...
type StorageAddableProperties struct {
	// ServerInstanceNo - Server 인스턴스 번호
	ServerInstanceNo *string `json:"serverInstanceNo,omitempty"`
	// ServerName - Server 이름
	ServerName *string `json:"serverName,omitempty"`
	// RedhatYn - 레드햇 OS 사용 유무
	RedhatYn *string `json:"redhatYn,omitempty"`
	// OsInformation - 운영체제 정보
	OsInformation *string `json:"osInformation,omitempty"`
	// Disabled - 사용 불가 여부
	Disabled *bool `json:"disabled,omitempty"`
	// ActionName - 동작 이름
	ActionName *string `json:"actionName,omitempty"`
}

// StorageDetachAndDeleteParameter ...
type StorageDetachAndDeleteParameter struct {
	// InstanceNoList - Storage 인스턴스 번호 리스트
	InstanceNoList *[]string `json:"instanceNoList,omitempty"`
}

// StorageListParameter storage 전체 검색 결과 리스트
type StorageListParameter struct {
	autorest.Response `json:"-"`
	// Content - Storage 컨텐츠 리스트
	Content *[]StorageListProperties `json:"content,omitempty"`
	// Total - 전체 storage의 개수
	Total *int32 `json:"total,omitempty"`
}

// StorageListProperties ...
type StorageListProperties struct {
	// BlockStorageTypeCode - Storage 타입 코드
	BlockStorageTypeCode *string `json:"blockStorageTypeCode,omitempty"`
	// CncNo - CNC 번호
	CncNo *string `json:"cncNo,omitempty"`
	// ComputeInstanceName - 서버 인스턴스 이름
	ComputeInstanceName *string `json:"computeInstanceName,omitempty"`
	// ComputeInstanceNo - 서버 인스턴스 번호
	ComputeInstanceNo *int32 `json:"computeInstanceNo,omitempty"`
	// CreateYmdt - Storage 생성일자
	CreateYmdt *float64 `json:"createYmdt,omitempty"`
	// Deletable - Storage 삭제 가능 여부
	Deletable *bool `json:"deletable,omitempty"`
	// Detachable - Storage 해제 가능 여부
	Detachable *bool `json:"detachable,omitempty"`
	// DeviceName - 장치 이름
	DeviceName *string `json:"deviceName,omitempty"`
	// DiskType2Code - Storage 디스크 타입 코드
	DiskType2Code *string `json:"diskType2Code,omitempty"`
	// DiskType2DetailCode - Storage 디스크 타입 상세 코드. Possible values include: 'SSD', 'HDD'
	DiskType2DetailCode DiskType2DetailCode `json:"diskType2DetailCode,omitempty"`
	// InstanceDesc - Storage 인스턴스 설명
	InstanceDesc *string `json:"instanceDesc,omitempty"`
	// InstanceNo - Storage 인스턴스 번호
	InstanceNo *string `json:"instanceNo,omitempty"`
	// InstanceStatusCode - Storage 인스턴스 상태 코드. Possible values include: 'StorageStatusCodeCREAT', 'StorageStatusCodeATTAC'
	InstanceStatusCode StorageStatusCode `json:"instanceStatusCode,omitempty"`
	// InstanceStatusName - Storage 인스턴스 상태 이름. Possible values include: 'Attaching', 'Attached', 'Detaching', 'Initialized'
	InstanceStatusName StorageStatusName `json:"instanceStatusName,omitempty"`
	// InstanceUUID - Storage 인스턴스 UUID
	InstanceUUID *string `json:"instanceUuid,omitempty"`
	// MaxIopsThroughput - Storage 최대 IOPS 성능값
	MaxIopsThroughput *int32 `json:"maxIopsThroughput,omitempty"`
	// MemberNo - 멤버 번호
	MemberNo *int32 `json:"memberNo,omitempty"`
	// OperationCode - 운영 코드
	OperationCode *string `json:"operationCode,omitempty"`
	// ProductCode - Product 코드
	ProductCode *string `json:"productCode,omitempty"`
	// Resizeable - Storage 크기 조절 가능 여부
	Resizeable *bool `json:"resizeable,omitempty"`
	// SrNo - SR 번호
	SrNo *int32 `json:"srNo,omitempty"`
	// VbdUUID - VBD UUID
	VbdUUID *string `json:"vbdUuid,omitempty"`
	// VolumeName - 볼륨 이름
	VolumeName *string `json:"volumeName,omitempty"`
	// VolumeSize - 볼륨 크기
	VolumeSize *float64 `json:"volumeSize,omitempty"`
	// ZoneName - Zone 이름
	ZoneName *string `json:"zoneName,omitempty"`
	// ZoneNo - Zone 번호
	ZoneNo *string `json:"zoneNo,omitempty"`
	// DeviceInfo - 장치 정보
	DeviceInfo *string `json:"deviceInfo,omitempty"`
	// VolumeSizeGb - 볼륨 크기 (Gb)
	VolumeSizeGb *string `json:"volumeSizeGb,omitempty"`
	// InstanceStatus - Storage 인스턴스 상태
	InstanceStatus *string `json:"instanceStatus,omitempty"`
	// InstanceStatusIcon - Storage 인스턴스 아이콘
	InstanceStatusIcon *string `json:"instanceStatusIcon,omitempty"`
	// RedhatYn - 레드햇 OS 사용 유무
	RedhatYn *string `json:"redhatYn,omitempty"`
}

// StorageParameter ...
type StorageParameter struct {
	// SnapshotInstanceNo - Snapshot 인스턴스 번호
	SnapshotInstanceNo *string `json:"snapshotInstanceNo,omitempty"`
	// ServerInstanceNo - Server 인스턴스 번호
	ServerInstanceNo *string `json:"serverInstanceNo,omitempty"`
	// PromotionNo - 프로모션 번호
	PromotionNo *string `json:"promotionNo,omitempty"`
	// DiskType2DetailCode - Storage 디스크 타입 상세 코드. Possible values include: 'SSD', 'HDD'
	DiskType2DetailCode DiskType2DetailCode `json:"diskType2DetailCode,omitempty"`
	// VolumeName - Storage 볼륨 이름
	VolumeName *string `json:"volumeName,omitempty"`
	// VolumeSize - Storage 볼륨 크기
	VolumeSize *int32 `json:"volumeSize,omitempty"`
	// InstanceDesc - Storage 인스턴스 설명
	InstanceDesc *string `json:"instanceDesc,omitempty"`
}
