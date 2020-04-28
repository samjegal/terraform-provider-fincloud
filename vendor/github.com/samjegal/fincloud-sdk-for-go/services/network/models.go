package network

// FINCLOUD_APACHE_NO_VERSION

import (
	"github.com/Azure/go-autorest/autorest"
)

// The package's fully qualified name.
const fqdn = "github.com/samjegal/fincloud-sdk-for-go/services/network"

// NatGatewayStatusCode enumerates the values for nat gateway status code.
type NatGatewayStatusCode string

const (
	// INIT ...
	INIT NatGatewayStatusCode = "INIT"
	// RUN ...
	RUN NatGatewayStatusCode = "RUN"
	// TERMTING ...
	TERMTING NatGatewayStatusCode = "TERMTING"
)

// PossibleNatGatewayStatusCodeValues returns an array of possible values for the NatGatewayStatusCode const type.
func PossibleNatGatewayStatusCodeValues() []NatGatewayStatusCode {
	return []NatGatewayStatusCode{INIT, RUN, TERMTING}
}

// NatGatewayStatusName enumerates the values for nat gateway status name.
type NatGatewayStatusName string

const (
	// 삭제중 ...
	삭제중 NatGatewayStatusName = "삭제중"
	// 운영중 ...
	운영중 NatGatewayStatusName = "운영중"
	// 준비중 ...
	준비중 NatGatewayStatusName = "준비중"
)

// PossibleNatGatewayStatusNameValues returns an array of possible values for the NatGatewayStatusName const type.
func PossibleNatGatewayStatusNameValues() []NatGatewayStatusName {
	return []NatGatewayStatusName{삭제중, 운영중, 준비중}
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

// ProtocolTypeCode1 enumerates the values for protocol type code 1.
type ProtocolTypeCode1 string

const (
	// ProtocolTypeCode1ICMP ...
	ProtocolTypeCode1ICMP ProtocolTypeCode1 = "ICMP"
	// ProtocolTypeCode1TCP ...
	ProtocolTypeCode1TCP ProtocolTypeCode1 = "TCP"
	// ProtocolTypeCode1UDP ...
	ProtocolTypeCode1UDP ProtocolTypeCode1 = "UDP"
)

// PossibleProtocolTypeCode1Values returns an array of possible values for the ProtocolTypeCode1 const type.
func PossibleProtocolTypeCode1Values() []ProtocolTypeCode1 {
	return []ProtocolTypeCode1{ProtocolTypeCode1ICMP, ProtocolTypeCode1TCP, ProtocolTypeCode1UDP}
}

// RouteTableStatusCode enumerates the values for route table status code.
type RouteTableStatusCode string

const (
	// RouteTableStatusCodeRUN ...
	RouteTableStatusCodeRUN RouteTableStatusCode = "RUN"
	// RouteTableStatusCodeSET ...
	RouteTableStatusCodeSET RouteTableStatusCode = "SET"
)

// PossibleRouteTableStatusCodeValues returns an array of possible values for the RouteTableStatusCode const type.
func PossibleRouteTableStatusCodeValues() []RouteTableStatusCode {
	return []RouteTableStatusCode{RouteTableStatusCodeRUN, RouteTableStatusCodeSET}
}

// RouteTableStatusName enumerates the values for route table status name.
type RouteTableStatusName string

const (
	// RouteTableStatusName설정중 ...
	RouteTableStatusName설정중 RouteTableStatusName = "설정중"
	// RouteTableStatusName운영중 ...
	RouteTableStatusName운영중 RouteTableStatusName = "운영중"
)

// PossibleRouteTableStatusNameValues returns an array of possible values for the RouteTableStatusName const type.
func PossibleRouteTableStatusNameValues() []RouteTableStatusName {
	return []RouteTableStatusName{RouteTableStatusName설정중, RouteTableStatusName운영중}
}

// SubnetStatusCode enumerates the values for subnet status code.
type SubnetStatusCode string

const (
	// SubnetStatusCodeCREATING ...
	SubnetStatusCodeCREATING SubnetStatusCode = "CREATING"
	// SubnetStatusCodeRUN ...
	SubnetStatusCodeRUN SubnetStatusCode = "RUN"
	// SubnetStatusCodeTERMTING ...
	SubnetStatusCodeTERMTING SubnetStatusCode = "TERMTING"
)

// PossibleSubnetStatusCodeValues returns an array of possible values for the SubnetStatusCode const type.
func PossibleSubnetStatusCodeValues() []SubnetStatusCode {
	return []SubnetStatusCode{SubnetStatusCodeCREATING, SubnetStatusCodeRUN, SubnetStatusCodeTERMTING}
}

// VirtualPrivateCloudStatusCode enumerates the values for virtual private cloud status code.
type VirtualPrivateCloudStatusCode string

const (
	// VirtualPrivateCloudStatusCodeCREATING ...
	VirtualPrivateCloudStatusCodeCREATING VirtualPrivateCloudStatusCode = "CREATING"
	// VirtualPrivateCloudStatusCodeINIT ...
	VirtualPrivateCloudStatusCodeINIT VirtualPrivateCloudStatusCode = "INIT"
	// VirtualPrivateCloudStatusCodeRUN ...
	VirtualPrivateCloudStatusCodeRUN VirtualPrivateCloudStatusCode = "RUN"
	// VirtualPrivateCloudStatusCodeTERMTING ...
	VirtualPrivateCloudStatusCodeTERMTING VirtualPrivateCloudStatusCode = "TERMTING"
)

// PossibleVirtualPrivateCloudStatusCodeValues returns an array of possible values for the VirtualPrivateCloudStatusCode const type.
func PossibleVirtualPrivateCloudStatusCodeValues() []VirtualPrivateCloudStatusCode {
	return []VirtualPrivateCloudStatusCode{VirtualPrivateCloudStatusCodeCREATING, VirtualPrivateCloudStatusCodeINIT, VirtualPrivateCloudStatusCodeRUN, VirtualPrivateCloudStatusCodeTERMTING}
}

// VirtualPrivateCloudStatusName enumerates the values for virtual private cloud status name.
type VirtualPrivateCloudStatusName string

const (
	// VirtualPrivateCloudStatusName삭제중 ...
	VirtualPrivateCloudStatusName삭제중 VirtualPrivateCloudStatusName = "삭제중"
	// VirtualPrivateCloudStatusName생성중 ...
	VirtualPrivateCloudStatusName생성중 VirtualPrivateCloudStatusName = "생성중"
	// VirtualPrivateCloudStatusName운영중 ...
	VirtualPrivateCloudStatusName운영중 VirtualPrivateCloudStatusName = "운영중"
	// VirtualPrivateCloudStatusName준비중 ...
	VirtualPrivateCloudStatusName준비중 VirtualPrivateCloudStatusName = "준비중"
)

// PossibleVirtualPrivateCloudStatusNameValues returns an array of possible values for the VirtualPrivateCloudStatusName const type.
func PossibleVirtualPrivateCloudStatusNameValues() []VirtualPrivateCloudStatusName {
	return []VirtualPrivateCloudStatusName{VirtualPrivateCloudStatusName삭제중, VirtualPrivateCloudStatusName생성중, VirtualPrivateCloudStatusName운영중, VirtualPrivateCloudStatusName준비중}
}

// ACLDetailParameter ACL 상세정보
type ACLDetailParameter struct {
	autorest.Response `json:"-"`
	// Content - ACL 컨텐츠 정보
	Content *ACLRuleDetailProperties `json:"content,omitempty"`
}

// ACLRequestParameter ...
type ACLRequestParameter struct {
	// NetworkACLName - ACL 이름
	NetworkACLName *string `json:"networkAclName,omitempty"`
	// VpcNo - VPC 번호
	VpcNo *string `json:"vpcNo,omitempty"`
	// Description - ACL 설명
	Description *string `json:"description,omitempty"`
}

// ACLResponseParameter ACL 전체 리스트
type ACLResponseParameter struct {
	autorest.Response `json:"-"`
	// Content - ACL 전체 컨텐츠 리스트
	Content *[]ACLRuleListProperties `json:"content,omitempty"`
	// Total - ACL 전체 개수
	Total *int32 `json:"total,omitempty"`
}

// ACLRuleContentParameter ACL rule 상세정보
type ACLRuleContentParameter struct {
	autorest.Response `json:"-"`
	// Content - ACL rule 컨텐츠 정보
	Content *[]ACLRuleContentProperties `json:"content,omitempty"`
	// Total - ACL 전체 개수
	Total *int32 `json:"total,omitempty"`
}

// ACLRuleContentProperties ...
type ACLRuleContentProperties struct {
	// RuleNo - ACL rule 번호
	RuleNo *int32 `json:"ruleNo,omitempty"`
	// NetworkACLNo - ACL 번호
	NetworkACLNo *int32 `json:"networkAclNo,omitempty"`
	// Priority - 우선순위 번호
	Priority *int32 `json:"priority,omitempty"`
	// ProtocolTypeCode - 프로토콜 타입 코드. Possible values include: 'ProtocolTypeCode1ICMP', 'ProtocolTypeCode1UDP', 'ProtocolTypeCode1TCP'
	ProtocolTypeCode ProtocolTypeCode1 `json:"protocolTypeCode,omitempty"`
	// IPBlock - ACL rule 정책 정보 (CIDR)
	IPBlock *string `json:"ipBlock,omitempty"`
	// PortRange - 포트 범위
	PortRange *string `json:"portRange,omitempty"`
	// IsAllowRule - Rule 허용여부
	IsAllowRule *bool `json:"isAllowRule,omitempty"`
	// IsInboundRule - ACL inbound rule 여부
	IsInboundRule *bool `json:"isInboundRule,omitempty"`
	// CreatedYmdt - ACL rule 생성일자
	CreatedYmdt *float64 `json:"createdYmdt,omitempty"`
	// ModifiedYmdt - ACL 수정일자
	ModifiedYmdt *float64 `json:"modifiedYmdt,omitempty"`
	// Description - ACL rule 설명
	Description *string `json:"description,omitempty"`
}

// ACLRuleDetailProperties ...
type ACLRuleDetailProperties struct {
	// NetworkACLNo - ACL 번호
	NetworkACLNo *int32 `json:"networkAclNo,omitempty"`
	// VpcNo - VPC 번호
	VpcNo *string `json:"vpcNo,omitempty"`
	// InboundRuleCount - Inbound rule 개수
	InboundRuleCount *int32 `json:"inboundRuleCount,omitempty"`
	// OutboundRuleCount - Outbound rule 개수
	OutboundRuleCount *int32 `json:"outboundRuleCount,omitempty"`
	// NetworkACLName - ALC 이름
	NetworkACLName *string `json:"networkAclName,omitempty"`
	// StatusCode - READ-ONLY; ACL 상태 코드
	StatusCode *string `json:"statusCode,omitempty"`
	// VpcName - VPC 이름
	VpcName *string `json:"vpcName,omitempty"`
	// CreatedYmdt - ACL 생성일자
	CreatedYmdt *float64 `json:"createdYmdt,omitempty"`
	// ModifiedYmdt - ACL 수정일자
	ModifiedYmdt *float64 `json:"modifiedYmdt,omitempty"`
	// DefaultYn - 기본 ACL 여부
	DefaultYn *string `json:"defaultYn,omitempty"`
	// SubnetCount - 연결된 subnet 개수
	SubnetCount *int32 `json:"subnetCount,omitempty"`
}

// ACLRuleListParameter ...
type ACLRuleListParameter struct {
	// NetworkACLRules - ACL rule 리스트 정보
	NetworkACLRules *[]ACLRuleProperties `json:"networkAclRules,omitempty"`
}

// ACLRuleListProperties ...
type ACLRuleListProperties struct {
	// NetworkACLNo - ACL 번호
	NetworkACLNo *int32 `json:"networkAclNo,omitempty"`
	// VpcNo - VPC 번호
	VpcNo *string `json:"vpcNo,omitempty"`
	// NetworkACLName - ALC 이름
	NetworkACLName *string `json:"networkAclName,omitempty"`
	// StatusCode - READ-ONLY; ACL 상태 코드
	StatusCode *string `json:"statusCode,omitempty"`
	// Description - ACL 설명
	Description *string `json:"description,omitempty"`
	// VpcName - VPC 이름
	VpcName *string `json:"vpcName,omitempty"`
	// CreatedYmdt - ACL 생성일자
	CreatedYmdt *float64 `json:"createdYmdt,omitempty"`
	// ModifiedYmdt - ACL 수정일자
	ModifiedYmdt *float64 `json:"modifiedYmdt,omitempty"`
	// DefaultYn - 기본 ACL 여부
	DefaultYn *string `json:"defaultYn,omitempty"`
	// SubnetCount - 연결된 subnet 개수
	SubnetCount *int32 `json:"subnetCount,omitempty"`
}

// ACLRuleProperties ...
type ACLRuleProperties struct {
	// IsInboundRule - ACL inbound rule 여부
	IsInboundRule *bool `json:"isInboundRule,omitempty"`
	// IsAllowRule - Rule 허용여부
	IsAllowRule *bool `json:"isAllowRule,omitempty"`
	// PortRange - 포트 범위
	PortRange *string `json:"portRange,omitempty"`
	// Priority - 우선순위 번호
	Priority *int32 `json:"priority,omitempty"`
	// ProtocolTypeCode - 프로토콜 타입 코드. Possible values include: 'ICMP', 'UDP', 'TCP'
	ProtocolTypeCode ProtocolTypeCode `json:"protocolTypeCode,omitempty"`
	// IPBlock - ACL rule 정책 정보 (CIDR)
	IPBlock *string `json:"ipBlock,omitempty"`
	// Description - ACL rule 설명
	Description *string `json:"description,omitempty"`
}

// ErrorMessage ...
type ErrorMessage struct {
	// ErrorCode - 에러 코드
	ErrorCode *string `json:"errorCode,omitempty"`
	// Message - 에러 메시지
	Message *string `json:"message,omitempty"`
	// OriginCode - 원본 코드
	OriginCode *string `json:"originCode,omitempty"`
	// OriginMessage - 원본 메시지
	OriginMessage *string `json:"originMessage,omitempty"`
}

// ErrorMessageParameter ...
type ErrorMessageParameter struct {
	autorest.Response `json:"-"`
	// Error - 에러 상세정보
	Error *ErrorMessage `json:"error,omitempty"`
}

// NatGatewayParameter ...
type NatGatewayParameter struct {
	// Description - NAT Gateway 설명
	Description *string `json:"description,omitempty"`
	// NatGatewayName - NAT Gateway 이름
	NatGatewayName *string `json:"natGatewayName,omitempty"`
	// VpcNo - VPC 번호
	VpcNo *string `json:"vpcNo,omitempty"`
	// ZoneNo - Zone 번호
	ZoneNo *string `json:"zoneNo,omitempty"`
}

// NatGatewaySearchContentParameter ...
type NatGatewaySearchContentParameter struct {
	// InstanceNo - NAT Gateway 번호
	InstanceNo *int32 `json:"instanceNo,omitempty"`
	// NatGatewayName - NAT Gateway 이름
	NatGatewayName *string `json:"natGatewayName,omitempty"`
	// EndpointNo - NAT Gateway 엔드포인트 번호
	EndpointNo *int32 `json:"endpointNo,omitempty"`
	// VpcName - VPC 이름
	VpcName *string `json:"vpcName,omitempty"`
	// PublicIP - 공용 IP 주소 CIDR
	PublicIP *string `json:"publicIp,omitempty"`
	// StatusCode - NAT Gateway 상태 코드. Possible values include: 'INIT', 'RUN', 'TERMTING'
	StatusCode NatGatewayStatusCode `json:"statusCode,omitempty"`
	// StatusName - NAT Gateway 상태 이름. Possible values include: '준비중', '운영중', '삭제중'
	StatusName NatGatewayStatusName `json:"statusName,omitempty"`
	// CreatedYmdt - NAT Gateway 생성일자
	CreatedYmdt *float64 `json:"createdYmdt,omitempty"`
}

// NatGatewaySearchFilterParameter ...
type NatGatewaySearchFilterParameter struct {
	// Field - 필터에 적용할 필드 이름
	Field *string `json:"field,omitempty"`
	// Test - 테스트 영역 (TBD)
	Test *string `json:"test,omitempty"`
}

// NatGatewaySearchListParameter ...
type NatGatewaySearchListParameter struct {
	autorest.Response `json:"-"`
	// Content - NAT Gateway 컨텐츠 리스트
	Content *[]NatGatewaySearchContentParameter `json:"content,omitempty"`
	// Total - 전체 NAT Gateway 개수
	Total *int32 `json:"total,omitempty"`
	// UserRequestID - NAT Gateway UUID
	UserRequestID *string `json:"userRequestId,omitempty"`
}

// NatGatewaySearchParameter ...
type NatGatewaySearchParameter struct {
	// Filter - NAT Gateway 검색 필터 리스트
	Filter *[]NatGatewaySearchFilterParameter `json:"filter,omitempty"`
	// PageNo - 검색할 NAT Gateway 페이지 번호
	PageNo *int32 `json:"pageNo,omitempty"`
	// PageSizeNo - 한 페이지에 나올 NAT Gateway 개수
	PageSizeNo *int32 `json:"pageSizeNo,omitempty"`
}

// RouteTableEndpointsContentParameter ...
type RouteTableEndpointsContentParameter struct {
	// EndpointNo - Endpoint 번호
	EndpointNo *int32 `json:"endpointNo,omitempty"`
	// EndpointInstanceName - Endpoint 인스턴스 이름
	EndpointInstanceName *string `json:"endpointInstanceName,omitempty"`
	// RegionNo - Region 번호
	RegionNo *int32 `json:"regionNo,omitempty"`
	// VpcNo - VPC 번호
	VpcNo *string `json:"vpcNo,omitempty"`
}

// RouteTableEndpointsParameter ...
type RouteTableEndpointsParameter struct {
	autorest.Response `json:"-"`
	// Content - Route Table 컨텐츠 리스트
	Content *[]RouteTableEndpointsContentParameter `json:"content,omitempty"`
	// Total - 전체 Route Table 개수
	Total *int32 `json:"total,omitempty"`
	// UserRequestID - Route Table UUID
	UserRequestID *string `json:"userRequestId,omitempty"`
}

// RouteTableEndpointTypesContentParameter ...
type RouteTableEndpointTypesContentParameter struct {
	// VpcNo - VPC 번호
	VpcNo *string `json:"vpcNo,omitempty"`
	// EndpointTypeCode - 엔드포인트 타입 코드
	EndpointTypeCode *string `json:"endpointTypeCode,omitempty"`
	// EndpointTypeName - 엔드포인트 타입 이름
	EndpointTypeName *string `json:"endpointTypeName,omitempty"`
}

// RouteTableEndpointTypesParameter ...
type RouteTableEndpointTypesParameter struct {
	autorest.Response `json:"-"`
	// Content - Route Table 컨텐츠 리스트
	Content *[]RouteTableEndpointTypesContentParameter `json:"content,omitempty"`
	// Total - 전체 Route Table 개수
	Total *int32 `json:"total,omitempty"`
	// UserRequestID - Route Table UUID
	UserRequestID *string `json:"userRequestId,omitempty"`
}

// RouteTableIsSettedContentParameter ...
type RouteTableIsSettedContentParameter struct {
	// VpcNo - VPC 번호
	VpcNo *string `json:"vpcNo,omitempty"`
	// SubnetNo - Subnet 번호
	SubnetNo *int32 `json:"subnetNo,omitempty"`
	// IsSetOtherRouteTable - Route Table 세팅 여부
	IsSetOtherRouteTable *bool `json:"isSetOtherRouteTable,omitempty"`
}

// RouteTableIsSettedParameter ...
type RouteTableIsSettedParameter struct {
	autorest.Response `json:"-"`
	// Content - Route Table 컨텐츠 리스트
	Content *[]RouteTableIsSettedContentParameter `json:"content,omitempty"`
	// Total - 전체 Route Table 개수
	Total *int32 `json:"total,omitempty"`
	// UserRequestID - Route Table UUID
	UserRequestID *string `json:"userRequestId,omitempty"`
}

// RouteTableParameter ...
type RouteTableParameter struct {
	// RouteTableName - Route Table 이름
	RouteTableName *string `json:"routeTableName,omitempty"`
	// Description - Route Table 설명
	Description *string `json:"description,omitempty"`
	// IgwYn - Subnet 지원 유형
	IgwYn *string `json:"igwYn,omitempty"`
	// VpcNo - VPC 번호
	VpcNo *string `json:"vpcNo,omitempty"`
}

// RouteTableRuleContentParameter ...
type RouteTableRuleContentParameter struct {
	// DestinationSubnetCidr - 목적지 서브넷 CIDR
	DestinationSubnetCidr *string `json:"destinationSubnetCidr,omitempty"`
	// EndpointNo - Endpoint 번호
	EndpointNo *int32 `json:"endpointNo,omitempty"`
	// RouteNo - Route Table 번호
	RouteNo *int32 `json:"routeNo,omitempty"`
}

// RouteTableRuleParameter ...
type RouteTableRuleParameter struct {
	// RouteTableNo - Route Table 번호
	RouteTableNo *int32 `json:"routeTableNo,omitempty"`
	// VpcNo - VPC 번호
	VpcNo *string `json:"vpcNo,omitempty"`
	// Route - Route Table 리스트
	Route *[]RouteTableRuleContentParameter `json:"route,omitempty"`
}

// RouteTableRuleSearchContentParameter ...
type RouteTableRuleSearchContentParameter struct {
	// DestinationSubnetCidr - 목적지 서브넷 CIDR
	DestinationSubnetCidr *string `json:"destinationSubnetCidr,omitempty"`
	// RouteNo - Route 번호
	RouteNo *int32 `json:"routeNo,omitempty"`
	// EndpointNo - Endpoint 번호
	EndpointNo *int32 `json:"endpointNo,omitempty"`
	// EndpointInstanceName - Endpoint 인스턴스 이름
	EndpointInstanceName *string `json:"endpointInstanceName,omitempty"`
	// EndpointTypeCode - Endpoint 타입 코드
	EndpointTypeCode *string `json:"endpointTypeCode,omitempty"`
	// RouteTableNo - Route Table 번호
	RouteTableNo *int32 `json:"routeTableNo,omitempty"`
	// DefaultYn - Route Table 기본 설정
	DefaultYn *string `json:"defaultYn,omitempty"`
}

// RouteTableRuleSearchParameter ...
type RouteTableRuleSearchParameter struct {
	autorest.Response `json:"-"`
	// Content - Route Table 컨텐츠 리스트
	Content *[]RouteTableRuleSearchContentParameter `json:"content,omitempty"`
	// Total - 전체 Route Table 개수
	Total *int32 `json:"total,omitempty"`
	// UserRequestID - Route Table UUID
	UserRequestID *string `json:"userRequestId,omitempty"`
}

// RouteTableSearchContentParameter ...
type RouteTableSearchContentParameter struct {
	// RouteTableNo - Route Table 번호
	RouteTableNo *int32 `json:"routeTableNo,omitempty"`
	// RouteTableName - Route Table 이름
	RouteTableName *string `json:"routeTableName,omitempty"`
	// VpcNo - VPC 번호
	VpcNo *string `json:"vpcNo,omitempty"`
	// VpcName - VPC 이름
	VpcName *string `json:"vpcName,omitempty"`
	// IgwYn - Subnet 지원 유형
	IgwYn *string `json:"igwYn,omitempty"`
	// SubnetCount - Subnet 갯수
	SubnetCount *int32 `json:"subnetCount,omitempty"`
	// DefaultYn - Route Table 기본 상태
	DefaultYn *string `json:"defaultYn,omitempty"`
	// StatusCode - Route Table 상태 코드. Possible values include: 'RouteTableStatusCodeRUN', 'RouteTableStatusCodeSET'
	StatusCode RouteTableStatusCode `json:"statusCode,omitempty"`
	// StatusName - Route Table 상태 이름. Possible values include: 'RouteTableStatusName운영중', 'RouteTableStatusName설정중'
	StatusName RouteTableStatusName `json:"statusName,omitempty"`
}

// RouteTableSearchFilterParameter ...
type RouteTableSearchFilterParameter struct {
	// Field - 필터에 적용할 필드 이름
	Field *string `json:"field,omitempty"`
	// Test - 테스트 영역 (TBD)
	Test *string `json:"test,omitempty"`
}

// RouteTableSearchListParameter ...
type RouteTableSearchListParameter struct {
	autorest.Response `json:"-"`
	// Content - Route Table 컨텐츠 리스트
	Content *[]RouteTableSearchContentParameter `json:"content,omitempty"`
	// Total - 전체 Route Table 개수
	Total *int32 `json:"total,omitempty"`
	// UserRequestID - Route Table UUID
	UserRequestID *string `json:"userRequestId,omitempty"`
}

// RouteTableSearchParameter ...
type RouteTableSearchParameter struct {
	// Filter - Route Table 검색 필터 리스트
	Filter *[]RouteTableSearchFilterParameter `json:"filter,omitempty"`
	// PageNo - 검색할 Route Table 페이지 번호
	PageNo *int32 `json:"pageNo,omitempty"`
	// PageSizeNo - 한 페이지에 나올 Route Table 개수
	PageSizeNo *int32 `json:"pageSizeNo,omitempty"`
}

// RouteTableSubnetListContentParameter ...
type RouteTableSubnetListContentParameter struct {
	// Disabled - 서브넷 해제 여부
	Disabled *bool `json:"disabled,omitempty"`
	// VpcNo - VPC 번호
	VpcNo *string `json:"vpcNo,omitempty"`
	// ZoneNo - Zone 번호
	ZoneNo *int32 `json:"zoneNo,omitempty"`
	// SubnetNo - Subnet 번호
	SubnetNo *string `json:"subnetNo,omitempty"`
	// SubnetName - Subnet 이름
	SubnetName *string `json:"subnetName,omitempty"`
	// Subnet - Subnet CIDR
	Subnet *string `json:"subnet,omitempty"`
	// Gateway - Subnet Gateway IP 주소
	Gateway *string `json:"gateway,omitempty"`
	// DNS1 - DNS1 IP 주소
	DNS1 *string `json:"dns1,omitempty"`
	// DNS2 - DNS2 IP 주소
	DNS2 *string `json:"dns2,omitempty"`
	// IgwYn - 인터넷 사용여부
	IgwYn *string `json:"igwYn,omitempty"`
	// StatusCode - Subnet 상태 코드. Possible values include: 'SubnetStatusCodeCREATING', 'SubnetStatusCodeRUN', 'SubnetStatusCodeTERMTING'
	StatusCode SubnetStatusCode `json:"statusCode,omitempty"`
	// StatusName - Subnet 상태 이름
	StatusName *string `json:"statusName,omitempty"`
}

// RouteTableSubnetListParameter ...
type RouteTableSubnetListParameter struct {
	autorest.Response `json:"-"`
	// Content - Route Table 연관 서브넷 컨텐츠 리스트
	Content *[]RouteTableSubnetListContentParameter `json:"content,omitempty"`
	// Total - 전체 Route Table 개수
	Total *int32 `json:"total,omitempty"`
	// UserRequestID - Route Table UUID
	UserRequestID *string `json:"userRequestId,omitempty"`
}

// RouteTableSubnetParameter ...
type RouteTableSubnetParameter struct {
	// RouteTableNo - Route Table 번호
	RouteTableNo *int32 `json:"routeTableNo,omitempty"`
	// VpcNo - VPC 번호
	VpcNo *string `json:"vpcNo,omitempty"`
	// SubnetNos - Subnet 리스트
	SubnetNos *[]string `json:"subnetNos,omitempty"`
}

// SubnetParameter ...
type SubnetParameter struct {
	// Subnet - Subnet CIDR 정보
	Subnet *string `json:"subnet,omitempty"`
	// SubnetName - Subnet 이름
	SubnetName *string `json:"subnetName,omitempty"`
	// VpcNo - VPC 번호
	VpcNo *string `json:"vpcNo,omitempty"`
	// ZoneNo - 금융 클라우드 zone 번호
	ZoneNo *int32 `json:"zoneNo,omitempty"`
	// IgwYn - Internet gateway 사용 여부
	IgwYn *string `json:"igwYn,omitempty"`
	// NetworkACLNo - Subnet에 적용된 ACL 번호
	NetworkACLNo *int32 `json:"networkAclNo,omitempty"`
	// LbYn - Load balancer 사용 여부
	LbYn *string `json:"lbYn,omitempty"`
	// BmYn - Bare Metal 사용 여부
	BmYn *string `json:"bmYn,omitempty"`
}

// SubnetSearchContentParameter ...
type SubnetSearchContentParameter struct {
	// VpcNo - VPC 번호
	VpcNo *string `json:"vpcNo,omitempty"`
	// VpcName - VPC 이름
	VpcName *string `json:"vpcName,omitempty"`
	// ZoneNo - 금융 클라우드 zone 번호
	ZoneNo *int32 `json:"zoneNo,omitempty"`
	// ZoneName - 금융 클라우드 zone 이름
	ZoneName *string `json:"zoneName,omitempty"`
	// SubnetNo - Subnet 번호
	SubnetNo *string `json:"subnetNo,omitempty"`
	// SubnetName - Subnet 이름
	SubnetName *string `json:"subnetName,omitempty"`
	// Subnet - Subnet CIDR 정보
	Subnet *string `json:"subnet,omitempty"`
	// Gateway - Gateway IP 주소
	Gateway *string `json:"gateway,omitempty"`
	// DNS1 - 1차 DNS IP 주소
	DNS1 *string `json:"dns1,omitempty"`
	// DNS2 - 2차 DNS IP 주소
	DNS2 *string `json:"dns2,omitempty"`
	// CreatedYmdt - Subnet 생성일자
	CreatedYmdt *float64 `json:"createdYmdt,omitempty"`
	// StatusCode - Subnet 상태 코드. Possible values include: 'SubnetStatusCodeCREATING', 'SubnetStatusCodeRUN', 'SubnetStatusCodeTERMTING'
	StatusCode SubnetStatusCode `json:"statusCode,omitempty"`
	// StatusName - Subnet 상태이름
	StatusName *string `json:"statusName,omitempty"`
	// NetworkACLNo - Subnet에 적용된 ACL 번호
	NetworkACLNo *int32 `json:"networkAclNo,omitempty"`
	// NetworkACLName - Subnet에 적용된 ACL 이름
	NetworkACLName *string `json:"networkAclName,omitempty"`
	// IgwYn - Internet gateway 사용 여부
	IgwYn *string `json:"igwYn,omitempty"`
	// LbYn - Load balancer 사용 여부
	LbYn *string `json:"lbYn,omitempty"`
	// BmYn - Bear Metal 사용 여부
	BmYn *string `json:"bmYn,omitempty"`
	// Disabled - 사용 여부
	Disabled *bool `json:"disabled,omitempty"`
}

// SubnetSearchFilterParameter ...
type SubnetSearchFilterParameter struct {
	// Field - 필터에 적용할 필드 이름
	Field *string `json:"field,omitempty"`
	// Test - 테스트 영역 (TBD)
	Test *string `json:"test,omitempty"`
}

// SubnetSearchListParameter subnet 전체 검색 결과 리스트
type SubnetSearchListParameter struct {
	autorest.Response `json:"-"`
	// Content - Subnet 컨텐츠 리스트
	Content *[]SubnetSearchContentParameter `json:"content,omitempty"`
	// Total - 전체 subnet의 개수
	Total *int32 `json:"total,omitempty"`
}

// SubnetSearchParameter ...
type SubnetSearchParameter struct {
	// PageNo - 검색할 subnet 페이지 번호
	PageNo *int32 `json:"pageNo,omitempty"`
	// PageSizeNo - 한 페이지에 나올 subnet 개수
	PageSizeNo *int32 `json:"pageSizeNo,omitempty"`
	// Filter - Subnet 검색 필터 리스트
	Filter *[]SubnetSearchFilterParameter `json:"filter,omitempty"`
}

// VirtualPrivateCloudContentParameter VPC 정보
type VirtualPrivateCloudContentParameter struct {
	// VpcNo - VPC 번호
	VpcNo *string `json:"vpcNo,omitempty"`
	// VxlanNo - VXLAN 번호
	VxlanNo *int32 `json:"vxlanNo,omitempty"`
	// VpcName - VPC 이름
	VpcName *string `json:"vpcName,omitempty"`
	// Ipv4Cidr - IPv4 CIDR 정보
	Ipv4Cidr *string `json:"ipv4Cidr,omitempty"`
	// StatusCode - VPC 상태 코드. Possible values include: 'VirtualPrivateCloudStatusCodeINIT', 'VirtualPrivateCloudStatusCodeCREATING', 'VirtualPrivateCloudStatusCodeRUN', 'VirtualPrivateCloudStatusCodeTERMTING'
	StatusCode VirtualPrivateCloudStatusCode `json:"statusCode,omitempty"`
	// StatusName - VPC 상태설명. Possible values include: 'VirtualPrivateCloudStatusName준비중', 'VirtualPrivateCloudStatusName생성중', 'VirtualPrivateCloudStatusName운영중', 'VirtualPrivateCloudStatusName삭제중'
	StatusName VirtualPrivateCloudStatusName `json:"statusName,omitempty"`
	// CreateDate - VPC 생성일자
	CreateDate *float64 `json:"createDate,omitempty"`
}

// VirtualPrivateCloudListParameter 전체 VPC 정보
type VirtualPrivateCloudListParameter struct {
	autorest.Response `json:"-"`
	// Content - VPC 정보 리스트
	Content *[]VirtualPrivateCloudContentParameter `json:"content,omitempty"`
	// Total - VPC 총 개수
	Total *int32 `json:"total,omitempty"`
}

// VirtualPrivateCloudParameter VPC 정보
type VirtualPrivateCloudParameter struct {
	// VpcName - VPC 이름
	VpcName *string `json:"vpcName,omitempty"`
	// Ipv4Cidr - IPv4 CIDR 정보
	Ipv4Cidr *string `json:"ipv4Cidr,omitempty"`
}
