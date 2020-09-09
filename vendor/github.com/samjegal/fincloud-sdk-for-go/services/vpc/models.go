package vpc

// FINCLOUD_APACHE_NO_VERSION

import (
	"github.com/Azure/go-autorest/autorest"
)

// The package's fully qualified name.
const fqdn = "github.com/samjegal/fincloud-sdk-for-go/services/vpc"

// CreateResponse ...
type CreateResponse struct {
	autorest.Response `json:"-"`
	CreateVpcResponse *CreateResponseCreateVpcResponse `json:"createVpcResponse,omitempty"`
}

// CreateResponseCreateVpcResponse ...
type CreateResponseCreateVpcResponse struct {
	// ReturnCode - 리턴 코드
	ReturnCode *string `json:"returnCode,omitempty"`
	// ReturnMessage - 리턴 메시지
	ReturnMessage *string `json:"returnMessage,omitempty"`
	// TotalRows - 총 행 개수
	TotalRows *int32 `json:"totalRows,omitempty"`
	// VpcList - VPC 리스트
	VpcList *[]List `json:"vpcList,omitempty"`
}

// DeleteResponse ...
type DeleteResponse struct {
	autorest.Response `json:"-"`
	DeleteVpcResponse *DeleteResponseDeleteVpcResponse `json:"deleteVpcResponse,omitempty"`
}

// DeleteResponseDeleteVpcResponse ...
type DeleteResponseDeleteVpcResponse struct {
	// ReturnCode - 리턴 코드
	ReturnCode *string `json:"returnCode,omitempty"`
	// ReturnMessage - 리턴 메시지
	ReturnMessage *string `json:"returnMessage,omitempty"`
	// TotalRows - 총 행 개수
	TotalRows *int32 `json:"totalRows,omitempty"`
	// VpcList - VPC 리스트
	VpcList *[]List `json:"vpcList,omitempty"`
}

// DetailResponse ...
type DetailResponse struct {
	autorest.Response    `json:"-"`
	GetVpcDetailResponse *DetailResponseGetVpcDetailResponse `json:"getVpcDetailResponse,omitempty"`
}

// DetailResponseGetVpcDetailResponse ...
type DetailResponseGetVpcDetailResponse struct {
	// ReturnCode - 리턴 코드
	ReturnCode *string `json:"returnCode,omitempty"`
	// ReturnMessage - 리턴 메시지
	ReturnMessage *string `json:"returnMessage,omitempty"`
	// TotalRows - 총 행 개수
	TotalRows *int32 `json:"totalRows,omitempty"`
	// VpcList - VPC 리스트
	VpcList *[]List `json:"vpcList,omitempty"`
}

// List ...
type List struct {
	// VpcNo - VPC 번호
	VpcNo *string `json:"vpcNo,omitempty"`
	// VpcName - VPC 이름
	VpcName *string `json:"vpcName,omitempty"`
	// Ipv4CidrBlock - IP 주소 CIDR
	Ipv4CidrBlock *string `json:"ipv4CidrBlock,omitempty"`
	// VpcStatus - VPC 상태
	VpcStatus *Status `json:"vpcStatus,omitempty"`
	// RegionCode - 리전 코드
	RegionCode *string `json:"regionCode,omitempty"`
	// CreateDate - 생성 일자
	CreateDate *string `json:"createDate,omitempty"`
}

// ListResponse ...
type ListResponse struct {
	autorest.Response  `json:"-"`
	GetVpcListResponse *ListResponseGetVpcListResponse `json:"getVpcListResponse,omitempty"`
}

// ListResponseGetVpcListResponse ...
type ListResponseGetVpcListResponse struct {
	// ReturnCode - 리턴 코드
	ReturnCode *string `json:"returnCode,omitempty"`
	// ReturnMessage - 리턴 메시지
	ReturnMessage *string `json:"returnMessage,omitempty"`
	// TotalRows - 총 행 개수
	TotalRows *int32 `json:"totalRows,omitempty"`
	// VpcList - VPC 리스트
	VpcList *[]List `json:"vpcList,omitempty"`
}

// NatGatewayInstanceCreateResponse ...
type NatGatewayInstanceCreateResponse struct {
	autorest.Response                `json:"-"`
	CreateNatGatewayInstanceResponse *NatGatewayInstanceCreateResponseCreateNatGatewayInstanceResponse `json:"createNatGatewayInstanceResponse,omitempty"`
}

// NatGatewayInstanceCreateResponseCreateNatGatewayInstanceResponse ...
type NatGatewayInstanceCreateResponseCreateNatGatewayInstanceResponse struct {
	// ReturnCode - 리턴 코드
	ReturnCode *string `json:"returnCode,omitempty"`
	// ReturnMessage - 리턴 메시지
	ReturnMessage *string `json:"returnMessage,omitempty"`
	// TotalRows - 총 행 개수
	TotalRows *int32 `json:"totalRows,omitempty"`
	// NatGatewayInstanceList - NAT 게이트웨이 인스턴스 리스트
	NatGatewayInstanceList *[]NatGatewayInstanceList `json:"natGatewayInstanceList,omitempty"`
}

// NatGatewayInstanceDeleteResponse ...
type NatGatewayInstanceDeleteResponse struct {
	autorest.Response                `json:"-"`
	DeleteNatGatewayInstanceResponse *NatGatewayInstanceDeleteResponseDeleteNatGatewayInstanceResponse `json:"deleteNatGatewayInstanceResponse,omitempty"`
}

// NatGatewayInstanceDeleteResponseDeleteNatGatewayInstanceResponse ...
type NatGatewayInstanceDeleteResponseDeleteNatGatewayInstanceResponse struct {
	// ReturnCode - 리턴 코드
	ReturnCode *string `json:"returnCode,omitempty"`
	// ReturnMessage - 리턴 메시지
	ReturnMessage *string `json:"returnMessage,omitempty"`
	// TotalRows - 총 행 개수
	TotalRows *int32 `json:"totalRows,omitempty"`
	// NatGatewayInstanceList - NAT 게이트웨이 인스턴스 리스트
	NatGatewayInstanceList *[]NatGatewayInstanceList `json:"natGatewayInstanceList,omitempty"`
}

// NatGatewayInstanceDetailResponse ...
type NatGatewayInstanceDetailResponse struct {
	autorest.Response                   `json:"-"`
	GetNatGatewayInstanceDetailResponse *NatGatewayInstanceDetailResponseGetNatGatewayInstanceDetailResponse `json:"getNatGatewayInstanceDetailResponse,omitempty"`
}

// NatGatewayInstanceDetailResponseGetNatGatewayInstanceDetailResponse ...
type NatGatewayInstanceDetailResponseGetNatGatewayInstanceDetailResponse struct {
	// ReturnCode - 리턴 코드
	ReturnCode *string `json:"returnCode,omitempty"`
	// ReturnMessage - 리턴 메시지
	ReturnMessage *string `json:"returnMessage,omitempty"`
	// TotalRows - 총 행 개수
	TotalRows *int32 `json:"totalRows,omitempty"`
	// NatGatewayInstanceList - NAT 게이트웨이 인스턴스 리스트
	NatGatewayInstanceList *[]NatGatewayInstanceList `json:"natGatewayInstanceList,omitempty"`
}

// NatGatewayInstanceList ...
type NatGatewayInstanceList struct {
	// VpcName - VPC 이름
	VpcName *string `json:"vpcName,omitempty"`
	// NatGatewayInstanceNo - NAT 게이트웨이 인스턴스 번호
	NatGatewayInstanceNo *string `json:"natGatewayInstanceNo,omitempty"`
	// NatGatewayName - NAT 게이트웨이 이름
	NatGatewayName *string `json:"natGatewayName,omitempty"`
	// PublicIP - 공인 IP 주소
	PublicIP *string `json:"publicIp,omitempty"`
	// NatGatewayInstanceStatus - NAT 게이트웨이 인스턴스 상태
	NatGatewayInstanceStatus *NatGatewayInstanceStatus `json:"natGatewayInstanceStatus,omitempty"`
	// NatGatewayInstanceStatusName - NAT 게이트웨이 인스턴스 상태 이름
	NatGatewayInstanceStatusName *string     `json:"natGatewayInstanceStatusName,omitempty"`
	NatGatewayInstanceOperation  interface{} `json:"natGatewayInstanceOperation,omitempty"`
	// CreateDate - 생성 일자
	CreateDate *string `json:"createDate,omitempty"`
	// NatGatewayDescription - NAT 게이트웨이 설명
	NatGatewayDescription *string `json:"natGatewayDescription,omitempty"`
}

// NatGatewayInstanceListResponse ...
type NatGatewayInstanceListResponse struct {
	autorest.Response                 `json:"-"`
	GetNatGatewayInstanceListResponse *NatGatewayInstanceListResponseGetNatGatewayInstanceListResponse `json:"getNatGatewayInstanceListResponse,omitempty"`
}

// NatGatewayInstanceListResponseGetNatGatewayInstanceListResponse ...
type NatGatewayInstanceListResponseGetNatGatewayInstanceListResponse struct {
	// ReturnCode - 리턴 코드
	ReturnCode *string `json:"returnCode,omitempty"`
	// ReturnMessage - 리턴 메시지
	ReturnMessage *string `json:"returnMessage,omitempty"`
	// TotalRows - 총 행 개수
	TotalRows *int32 `json:"totalRows,omitempty"`
	// NatGatewayInstanceList - NAT 게이트웨이 인스턴스 리스트
	NatGatewayInstanceList *[]NatGatewayInstanceList `json:"natGatewayInstanceList,omitempty"`
}

// NatGatewayInstanceStatus ...
type NatGatewayInstanceStatus struct {
	// Code - 상태 코드
	Code *string `json:"code,omitempty"`
	// CodeName - 상태 코드 이름
	CodeName *string `json:"codeName,omitempty"`
}

// NetworkACLCreateResponse ...
type NetworkACLCreateResponse struct {
	autorest.Response        `json:"-"`
	CreateNetworkACLResponse *NetworkACLCreateResponseCreateNetworkACLResponse `json:"createNetworkAclResponse,omitempty"`
}

// NetworkACLCreateResponseCreateNetworkACLResponse ...
type NetworkACLCreateResponseCreateNetworkACLResponse struct {
	// ReturnCode - 리턴 코드
	ReturnCode *string `json:"returnCode,omitempty"`
	// ReturnMessage - 리턴 메시지
	ReturnMessage *string `json:"returnMessage,omitempty"`
	// TotalRows - 총 행 개수
	TotalRows *int32 `json:"totalRows,omitempty"`
	// NetworkACLList - 네트워크 ACL 리스트
	NetworkACLList *[]NetworkACLList `json:"networkAclList,omitempty"`
}

// NetworkACLDeleteResponse ...
type NetworkACLDeleteResponse struct {
	autorest.Response        `json:"-"`
	DeleteNetworkACLResponse *NetworkACLDeleteResponseDeleteNetworkACLResponse `json:"deleteNetworkAclResponse,omitempty"`
}

// NetworkACLDeleteResponseDeleteNetworkACLResponse ...
type NetworkACLDeleteResponseDeleteNetworkACLResponse struct {
	// ReturnCode - 리턴 코드
	ReturnCode *string `json:"returnCode,omitempty"`
	// ReturnMessage - 리턴 메시지
	ReturnMessage *string `json:"returnMessage,omitempty"`
	// TotalRows - 총 행 개수
	TotalRows *int32 `json:"totalRows,omitempty"`
	// NetworkACLList - 네트워크 ACL 리스트
	NetworkACLList *[]NetworkACLList `json:"networkAclList,omitempty"`
}

// NetworkACLDetailResponse ...
type NetworkACLDetailResponse struct {
	autorest.Response           `json:"-"`
	GetNetworkACLDetailResponse *NetworkACLDetailResponseGetNetworkACLDetailResponse `json:"getNetworkAclDetailResponse,omitempty"`
}

// NetworkACLDetailResponseGetNetworkACLDetailResponse ...
type NetworkACLDetailResponseGetNetworkACLDetailResponse struct {
	// ReturnCode - 리턴 코드
	ReturnCode *string `json:"returnCode,omitempty"`
	// ReturnMessage - 리턴 메시지
	ReturnMessage *string `json:"returnMessage,omitempty"`
	// TotalRows - 총 행 개수
	TotalRows *int32 `json:"totalRows,omitempty"`
	// NetworkACLList - 네트워크 ACL 리스트
	NetworkACLList *[]NetworkACLList `json:"networkAclList,omitempty"`
}

// NetworkACLInboundRuleAddResponse ...
type NetworkACLInboundRuleAddResponse struct {
	autorest.Response                `json:"-"`
	AddNetworkACLInboundRuleResponse *NetworkACLInboundRuleAddResponseAddNetworkACLInboundRuleResponse `json:"addNetworkAclInboundRuleResponse,omitempty"`
}

// NetworkACLInboundRuleAddResponseAddNetworkACLInboundRuleResponse ...
type NetworkACLInboundRuleAddResponseAddNetworkACLInboundRuleResponse struct {
	// ReturnCode - 리턴 코드
	ReturnCode *string `json:"returnCode,omitempty"`
	// ReturnMessage - 리턴 메시지
	ReturnMessage *string `json:"returnMessage,omitempty"`
	// TotalRows - 총 행 개수
	TotalRows *int32 `json:"totalRows,omitempty"`
	// NetworkACLRuleList - 네트워크 ACL 룰 리스트
	NetworkACLRuleList *[]NetworkACLRuleList `json:"networkAclRuleList,omitempty"`
}

// NetworkACLInboundRuleRemoveResponse ...
type NetworkACLInboundRuleRemoveResponse struct {
	autorest.Response                   `json:"-"`
	RemoveNetworkACLInboundRuleResponse *NetworkACLInboundRuleRemoveResponseRemoveNetworkACLInboundRuleResponse `json:"removeNetworkAclInboundRuleResponse,omitempty"`
}

// NetworkACLInboundRuleRemoveResponseRemoveNetworkACLInboundRuleResponse ...
type NetworkACLInboundRuleRemoveResponseRemoveNetworkACLInboundRuleResponse struct {
	// ReturnCode - 리턴 코드
	ReturnCode *string `json:"returnCode,omitempty"`
	// ReturnMessage - 리턴 메시지
	ReturnMessage *string `json:"returnMessage,omitempty"`
	// TotalRows - 총 행 개수
	TotalRows *int32 `json:"totalRows,omitempty"`
	// NetworkACLRuleList - 네트워크 ACL 룰 리스트
	NetworkACLRuleList *[]NetworkACLRuleList `json:"networkAclRuleList,omitempty"`
}

// NetworkACLList ...
type NetworkACLList struct {
	// NetworkACLNo - 네트워크 ACL 번호
	NetworkACLNo *string `json:"networkAclNo,omitempty"`
	// NetworkACLName - 네트워크 ACL 이름
	NetworkACLName *string `json:"networkAclName,omitempty"`
	// VpcNo - VPC 번호
	VpcNo *string `json:"vpcNo,omitempty"`
	// NetworkACLStatus - 네트워크 ACL 상태
	NetworkACLStatus *NetworkACLStatus `json:"networkAclStatus,omitempty"`
	// NetworkACLDescription - 네트워크 ACL 설명
	NetworkACLDescription *string `json:"networkAclDescription,omitempty"`
	// CreateDate - 생성 일자
	CreateDate *string `json:"createDate,omitempty"`
	// IsDefault - 네트워크 ACL 기본 여부
	IsDefault *bool `json:"isDefault,omitempty"`
}

// NetworkACLListResponse ...
type NetworkACLListResponse struct {
	autorest.Response         `json:"-"`
	GetNetworkACLListResponse *NetworkACLListResponseGetNetworkACLListResponse `json:"getNetworkAclListResponse,omitempty"`
}

// NetworkACLListResponseGetNetworkACLListResponse ...
type NetworkACLListResponseGetNetworkACLListResponse struct {
	// ReturnCode - 리턴 코드
	ReturnCode *string `json:"returnCode,omitempty"`
	// ReturnMessage - 리턴 메시지
	ReturnMessage *string `json:"returnMessage,omitempty"`
	// TotalRows - 총 행 개수
	TotalRows *int32 `json:"totalRows,omitempty"`
	// NetworkACLList - 네트워크 ACL 리스트
	NetworkACLList *[]NetworkACLList `json:"networkAclList,omitempty"`
}

// NetworkACLOutboundRuleAddResponse ...
type NetworkACLOutboundRuleAddResponse struct {
	autorest.Response                 `json:"-"`
	AddNetworkACLOutboundRuleResponse *NetworkACLOutboundRuleAddResponseAddNetworkACLOutboundRuleResponse `json:"addNetworkAclOutboundRuleResponse,omitempty"`
}

// NetworkACLOutboundRuleAddResponseAddNetworkACLOutboundRuleResponse ...
type NetworkACLOutboundRuleAddResponseAddNetworkACLOutboundRuleResponse struct {
	// ReturnCode - 리턴 코드
	ReturnCode *string `json:"returnCode,omitempty"`
	// ReturnMessage - 리턴 메시지
	ReturnMessage *string `json:"returnMessage,omitempty"`
	// TotalRows - 총 행 개수
	TotalRows *int32 `json:"totalRows,omitempty"`
	// NetworkACLRuleList - 네트워크 ACL 룰 리스트
	NetworkACLRuleList *[]NetworkACLRuleList `json:"networkAclRuleList,omitempty"`
}

// NetworkACLOutboundRuleRemoveResponse ...
type NetworkACLOutboundRuleRemoveResponse struct {
	autorest.Response                    `json:"-"`
	RemoveNetworkACLOutboundRuleResponse *NetworkACLOutboundRuleRemoveResponseRemoveNetworkACLOutboundRuleResponse `json:"removeNetworkAclOutboundRuleResponse,omitempty"`
}

// NetworkACLOutboundRuleRemoveResponseRemoveNetworkACLOutboundRuleResponse ...
type NetworkACLOutboundRuleRemoveResponseRemoveNetworkACLOutboundRuleResponse struct {
	// ReturnCode - 리턴 코드
	ReturnCode *string `json:"returnCode,omitempty"`
	// ReturnMessage - 리턴 메시지
	ReturnMessage *string `json:"returnMessage,omitempty"`
	// TotalRows - 총 행 개수
	TotalRows *int32 `json:"totalRows,omitempty"`
	// NetworkACLRuleList - 네트워크 ACL 룰 리스트
	NetworkACLRuleList *[]NetworkACLRuleList `json:"networkAclRuleList,omitempty"`
}

// NetworkACLRuleList ...
type NetworkACLRuleList struct {
	// NetworkACLNo - 네트워크 ACL 번호
	NetworkACLNo *string `json:"networkAclNo,omitempty"`
	// Priority - 우선순위
	Priority *int32 `json:"priority,omitempty"`
	// ProtocolType - 프로토콜 타입
	ProtocolType *ProtocolType `json:"protocolType,omitempty"`
	// PortRange - 포트 범위
	PortRange *string `json:"portRange,omitempty"`
	// RuleAction - 룰 액션
	RuleAction *RuleAction `json:"ruleAction,omitempty"`
	// CreateDate - 생성 일자
	CreateDate *string `json:"createDate,omitempty"`
	// IPBlock - 블럭 여부
	IPBlock *string `json:"ipBlock,omitempty"`
	// NetworkACLRuleType - 네트워크 ACL 룰 타입
	NetworkACLRuleType *NetworkACLRuleType `json:"networkAclRuleType,omitempty"`
	// NetworkACLRuleDescription - 네트워크 ACL 룰 설명
	NetworkACLRuleDescription *string `json:"networkAclRuleDescription,omitempty"`
}

// NetworkACLRuleListResponse ...
type NetworkACLRuleListResponse struct {
	autorest.Response             `json:"-"`
	GetNetworkACLRuleListResponse *NetworkACLRuleListResponseGetNetworkACLRuleListResponse `json:"getNetworkAclRuleListResponse,omitempty"`
}

// NetworkACLRuleListResponseGetNetworkACLRuleListResponse ...
type NetworkACLRuleListResponseGetNetworkACLRuleListResponse struct {
	// ReturnCode - 리턴 코드
	ReturnCode *string `json:"returnCode,omitempty"`
	// ReturnMessage - 리턴 메시지
	ReturnMessage *string `json:"returnMessage,omitempty"`
	// TotalRows - 총 행 개수
	TotalRows *int32 `json:"totalRows,omitempty"`
	// NetworkACLRuleList - 네트워크 ACL 룰 리스트
	NetworkACLRuleList *[]NetworkACLRuleList `json:"networkAclRuleList,omitempty"`
}

// NetworkACLRuleType ...
type NetworkACLRuleType struct {
	// Code - 상태 코드
	Code *string `json:"code,omitempty"`
	// CodeName - 상태 코드 이름
	CodeName *string `json:"codeName,omitempty"`
}

// NetworkACLStatus ...
type NetworkACLStatus struct {
	// Code - 상태 코드
	Code *string `json:"code,omitempty"`
	// CodeName - 상태 코드 이름
	CodeName *string `json:"codeName,omitempty"`
}

// PeeringInstanceAcceptOrRejectResponse ...
type PeeringInstanceAcceptOrRejectResponse struct {
	autorest.Response                `json:"-"`
	AcceptOrRejectVpcPeeringResponse *PeeringInstanceAcceptOrRejectResponseAcceptOrRejectVpcPeeringResponse `json:"acceptOrRejectVpcPeeringResponse,omitempty"`
}

// PeeringInstanceAcceptOrRejectResponseAcceptOrRejectVpcPeeringResponse ...
type PeeringInstanceAcceptOrRejectResponseAcceptOrRejectVpcPeeringResponse struct {
	// ReturnCode - 리턴 코드
	ReturnCode *string `json:"returnCode,omitempty"`
	// ReturnMessage - 리턴 메시지
	ReturnMessage *string `json:"returnMessage,omitempty"`
	// TotalRows - 총 행 개수
	TotalRows *int32 `json:"totalRows,omitempty"`
	// VpcPeeringInstanceList - VPC 피어링 인스턴스 리스트
	VpcPeeringInstanceList *[]PeeringInstanceList `json:"VpcPeeringInstanceList,omitempty"`
}

// PeeringInstanceCreateResponse ...
type PeeringInstanceCreateResponse struct {
	autorest.Response                `json:"-"`
	CreateVpcPeeringInstanceResponse *PeeringInstanceCreateResponseCreateVpcPeeringInstanceResponse `json:"createVpcPeeringInstanceResponse,omitempty"`
}

// PeeringInstanceCreateResponseCreateVpcPeeringInstanceResponse ...
type PeeringInstanceCreateResponseCreateVpcPeeringInstanceResponse struct {
	// ReturnCode - 리턴 코드
	ReturnCode *string `json:"returnCode,omitempty"`
	// ReturnMessage - 리턴 메시지
	ReturnMessage *string `json:"returnMessage,omitempty"`
	// TotalRows - 총 행 개수
	TotalRows *int32 `json:"totalRows,omitempty"`
	// VpcPeeringInstanceList - VPC 피어링 인스턴스 리스트
	VpcPeeringInstanceList *[]PeeringInstanceList `json:"VpcPeeringInstanceList,omitempty"`
}

// PeeringInstanceDeleteResponse ...
type PeeringInstanceDeleteResponse struct {
	autorest.Response                `json:"-"`
	DeleteVpcPeeringInstanceResponse *PeeringInstanceDeleteResponseDeleteVpcPeeringInstanceResponse `json:"deleteVpcPeeringInstanceResponse,omitempty"`
}

// PeeringInstanceDeleteResponseDeleteVpcPeeringInstanceResponse ...
type PeeringInstanceDeleteResponseDeleteVpcPeeringInstanceResponse struct {
	// ReturnCode - 리턴 코드
	ReturnCode *string `json:"returnCode,omitempty"`
	// ReturnMessage - 리턴 메시지
	ReturnMessage *string `json:"returnMessage,omitempty"`
	// TotalRows - 총 행 개수
	TotalRows *int32 `json:"totalRows,omitempty"`
	// VpcPeeringInstanceList - VPC 피어링 인스턴스 리스트
	VpcPeeringInstanceList *[]PeeringInstanceList `json:"VpcPeeringInstanceList,omitempty"`
}

// PeeringInstanceDetailResponse ...
type PeeringInstanceDetailResponse struct {
	autorest.Response                   `json:"-"`
	GetVpcPeeringInstanceDetailResponse *PeeringInstanceDetailResponseGetVpcPeeringInstanceDetailResponse `json:"getVpcPeeringInstanceDetailResponse,omitempty"`
}

// PeeringInstanceDetailResponseGetVpcPeeringInstanceDetailResponse ...
type PeeringInstanceDetailResponseGetVpcPeeringInstanceDetailResponse struct {
	// ReturnCode - 리턴 코드
	ReturnCode *string `json:"returnCode,omitempty"`
	// ReturnMessage - 리턴 메시지
	ReturnMessage *string `json:"returnMessage,omitempty"`
	// TotalRows - 총 행 개수
	TotalRows *int32 `json:"totalRows,omitempty"`
	// VpcPeeringInstanceList - VPC 피어링 인스턴스 리스트
	VpcPeeringInstanceList *[]PeeringInstanceList `json:"VpcPeeringInstanceList,omitempty"`
}

// PeeringInstanceList ...
type PeeringInstanceList struct {
	// VpcPeeringInstanceNo - VPC 피어링 인스턴스 번호
	VpcPeeringInstanceNo *string `json:"vpcPeeringInstanceNo,omitempty"`
	// VpcPeeringName - VPC 피어링 이름
	VpcPeeringName *string `json:"vpcPeeringName,omitempty"`
	// RegionCode - 리전 코드
	RegionCode *string `json:"regionCode,omitempty"`
	// CreateDate - 생성 일자
	CreateDate *string `json:"createDate,omitempty"`
	// VpcPeeringInstanceStatus - VPC 피어링 인스턴스 상태
	VpcPeeringInstanceStatus *PeeringInstanceStatus `json:"vpcPeeringInstanceStatus,omitempty"`
	// VpcPeeringInstanceStatusName - VPC 피어링 인스턴스 상태 이름
	VpcPeeringInstanceStatusName *string     `json:"vpcPeeringInstanceStatusName,omitempty"`
	VpcPeeringInstanceOperation  interface{} `json:"vpcPeeringInstanceOperation,omitempty"`
	// SourceVpcNo - 소스 VPC 번호
	SourceVpcNo *string `json:"sourceVpcNo,omitempty"`
	// SourceVpcName - 소스 VPC 이름
	SourceVpcName *string `json:"sourceVpcName,omitempty"`
	// SourceVpcIpv4CidrBlock - 소스 VPC IP 주소 CIDR
	SourceVpcIpv4CidrBlock *string `json:"sourceVpcIpv4CidrBlock,omitempty"`
	// SourceVpcLoginID - 소스 VPC 로그인 아이디
	SourceVpcLoginID *string `json:"sourceVpcLoginId,omitempty"`
	// TargetVpcNo - 타겟 VPC 번호
	TargetVpcNo *string `json:"targetVpcNo,omitempty"`
	// TargetVpcName - 타겟 VPC 이름
	TargetVpcName *string `json:"targetVpcName,omitempty"`
	// TargetVpcIpv4CidrBlock - 타겟 VPC IP 주소 CIDR
	TargetVpcIpv4CidrBlock *string `json:"targetVpcIpv4CidrBlock,omitempty"`
	// TargetVpcLoginID - 타겟 VPC 로그인 ID
	TargetVpcLoginID *string `json:"targetVpcLoginId,omitempty"`
	// VpcPeeringDescription - VPC 피어링 설명
	VpcPeeringDescription *string `json:"vpcPeeringDescription,omitempty"`
	// HasReverseVpcPeering - 리버스 VPC 피어링 여부
	HasReverseVpcPeering *bool `json:"hasReverseVpcPeering,omitempty"`
	// IsBetweenAccounts - 계정간 피어링 여부
	IsBetweenAccounts *bool `json:"isBetweenAccounts,omitempty"`
	// ReverseVpcPeeringInstanceNo - 리버스 VPC 피어링 인스턴스 번호
	ReverseVpcPeeringInstanceNo *string `json:"reverseVpcPeeringInstanceNo,omitempty"`
}

// PeeringInstanceListResponse ...
type PeeringInstanceListResponse struct {
	autorest.Response                 `json:"-"`
	GetVpcPeeringInstanceListResponse *PeeringInstanceListResponseGetVpcPeeringInstanceListResponse `json:"getVpcPeeringInstanceListResponse,omitempty"`
}

// PeeringInstanceListResponseGetVpcPeeringInstanceListResponse ...
type PeeringInstanceListResponseGetVpcPeeringInstanceListResponse struct {
	// ReturnCode - 리턴 코드
	ReturnCode *string `json:"returnCode,omitempty"`
	// ReturnMessage - 리턴 메시지
	ReturnMessage *string `json:"returnMessage,omitempty"`
	// TotalRows - 총 행 개수
	TotalRows *int32 `json:"totalRows,omitempty"`
	// VpcPeeringInstanceList - VPC 피어링 인스턴스 리스트
	VpcPeeringInstanceList *[]PeeringInstanceList `json:"VpcPeeringInstanceList,omitempty"`
}

// PeeringInstanceStatus ...
type PeeringInstanceStatus struct {
	// Code - 상태 코드
	Code *string `json:"code,omitempty"`
	// CodeName - 상태 코드 이름
	CodeName *string `json:"codeName,omitempty"`
}

// PeeringResponse ...
type PeeringResponse struct {
	// ReturnCode - 리턴 코드
	ReturnCode *string `json:"returnCode,omitempty"`
	// ReturnMessage - 리턴 메시지
	ReturnMessage *string `json:"returnMessage,omitempty"`
	// TotalRows - 총 행 개수
	TotalRows *int32 `json:"totalRows,omitempty"`
	// VpcPeeringInstanceList - VPC 피어링 인스턴스 리스트
	VpcPeeringInstanceList *[]PeeringInstanceList `json:"VpcPeeringInstanceList,omitempty"`
}

// ProtocolType ...
type ProtocolType struct {
	// Code - 상태 코드
	Code *string `json:"code,omitempty"`
	// CodeName - 상태 코드 이름
	CodeName *string `json:"codeName,omitempty"`
}

// RouteAddResponse ...
type RouteAddResponse struct {
	autorest.Response `json:"-"`
	AddRouteResponse  *RouteAddResponseAddRouteResponse `json:"addRouteResponse,omitempty"`
}

// RouteAddResponseAddRouteResponse ...
type RouteAddResponseAddRouteResponse struct {
	// ReturnCode - 리턴 코드
	ReturnCode *string `json:"returnCode,omitempty"`
	// ReturnMessage - 리턴 메시지
	ReturnMessage *string `json:"returnMessage,omitempty"`
	// TotalRows - 총 행 개수
	TotalRows *int32 `json:"totalRows,omitempty"`
	// RouteList - 라우트 리스트
	RouteList *[]RouteList `json:"routeList,omitempty"`
}

// RouteList ...
type RouteList struct {
	// DestinationCidrBlock - 목적이 IP 주소 CIDR
	DestinationCidrBlock *string `json:"destinationCidrBlock,omitempty"`
	// TargetName - 타겟 이름
	TargetName *string `json:"targetName,omitempty"`
	// RouteTableNo - 라우트 테이블 번호
	RouteTableNo *string `json:"routeTableNo,omitempty"`
	// TargetType - 타겟 타입
	TargetType *TargetType `json:"targetType,omitempty"`
	// TargetNo - 타겟 번호
	TargetNo *string `json:"targetNo,omitempty"`
	// IsDefault - 기본 라우트 여부
	IsDefault *bool `json:"isDefault,omitempty"`
}

// RouteListResponse ...
type RouteListResponse struct {
	autorest.Response    `json:"-"`
	GetRouteListResponse *RouteListResponseGetRouteListResponse `json:"getRouteListResponse,omitempty"`
}

// RouteListResponseGetRouteListResponse ...
type RouteListResponseGetRouteListResponse struct {
	// ReturnCode - 리턴 코드
	ReturnCode *string `json:"returnCode,omitempty"`
	// ReturnMessage - 리턴 메시지
	ReturnMessage *string `json:"returnMessage,omitempty"`
	// TotalRows - 총 행 개수
	TotalRows *int32 `json:"totalRows,omitempty"`
	// RouteList - 라우트 리스트
	RouteList *[]RouteList `json:"routeList,omitempty"`
}

// RouteRemoveResponse ...
type RouteRemoveResponse struct {
	autorest.Response   `json:"-"`
	RemoveRouteResponse *RouteRemoveResponseRemoveRouteResponse `json:"removeRouteResponse,omitempty"`
}

// RouteRemoveResponseRemoveRouteResponse ...
type RouteRemoveResponseRemoveRouteResponse struct {
	// ReturnCode - 리턴 코드
	ReturnCode *string `json:"returnCode,omitempty"`
	// ReturnMessage - 리턴 메시지
	ReturnMessage *string `json:"returnMessage,omitempty"`
	// TotalRows - 총 행 개수
	TotalRows *int32 `json:"totalRows,omitempty"`
	// RouteList - 라우트 리스트
	RouteList *[]RouteList `json:"routeList,omitempty"`
}

// RouteTableCreateResponse ...
type RouteTableCreateResponse struct {
	autorest.Response        `json:"-"`
	CreateRouteTableResponse *RouteTableCreateResponseCreateRouteTableResponse `json:"createRouteTableResponse,omitempty"`
}

// RouteTableCreateResponseCreateRouteTableResponse ...
type RouteTableCreateResponseCreateRouteTableResponse struct {
	// ReturnCode - 리턴 코드
	ReturnCode *string `json:"returnCode,omitempty"`
	// ReturnMessage - 리턴 메시지
	ReturnMessage *string `json:"returnMessage,omitempty"`
	// TotalRows - 총 행 개수
	TotalRows *int32 `json:"totalRows,omitempty"`
	// RouteTableList - 라우트 테이블 리스트
	RouteTableList *[]RouteTableList `json:"routeTableList,omitempty"`
}

// RouteTableDeleteResponse ...
type RouteTableDeleteResponse struct {
	autorest.Response        `json:"-"`
	DeleteRouteTableResponse *RouteTableDeleteResponseDeleteRouteTableResponse `json:"deleteRouteTableResponse,omitempty"`
}

// RouteTableDeleteResponseDeleteRouteTableResponse ...
type RouteTableDeleteResponseDeleteRouteTableResponse struct {
	// ReturnCode - 리턴 코드
	ReturnCode *string `json:"returnCode,omitempty"`
	// ReturnMessage - 리턴 메시지
	ReturnMessage *string `json:"returnMessage,omitempty"`
	// TotalRows - 총 행 개수
	TotalRows *int32 `json:"totalRows,omitempty"`
	// RouteTableList - 라우트 테이블 리스트
	RouteTableList *[]RouteTableList `json:"routeTableList,omitempty"`
}

// RouteTableDetailResponse ...
type RouteTableDetailResponse struct {
	autorest.Response           `json:"-"`
	GetRouteTableDetailResponse *RouteTableDetailResponseGetRouteTableDetailResponse `json:"getRouteTableDetailResponse,omitempty"`
}

// RouteTableDetailResponseGetRouteTableDetailResponse ...
type RouteTableDetailResponseGetRouteTableDetailResponse struct {
	// ReturnCode - 리턴 코드
	ReturnCode *string `json:"returnCode,omitempty"`
	// ReturnMessage - 리턴 메시지
	ReturnMessage *string `json:"returnMessage,omitempty"`
	// TotalRows - 총 행 개수
	TotalRows *int32 `json:"totalRows,omitempty"`
	// RouteTableList - 라우트 테이블 리스트
	RouteTableList *[]RouteTableList `json:"routeTableList,omitempty"`
}

// RouteTableList ...
type RouteTableList struct {
	// RouteTableNo - 라우트 테이블 번호
	RouteTableNo *string `json:"routeTableNo,omitempty"`
	// RouteTableName - 라우트 테이블 이름
	RouteTableName *string `json:"routeTableName,omitempty"`
	// RegionCode - 리전 코드
	RegionCode *string `json:"regionCode,omitempty"`
	// VpcNo - VPC 번호
	VpcNo *string `json:"vpcNo,omitempty"`
	// SupportedSubnetType - 지원 서브넷 타입
	SupportedSubnetType *SupportedSubnetType `json:"supportedSubnetType,omitempty"`
	// IsDefault - 기본 라우트 여부
	IsDefault *bool `json:"isDefault,omitempty"`
	// RouteTableStatus - 라우트 테이블 상태
	RouteTableStatus *RouteTableStatus `json:"routeTableStatus,omitempty"`
	// RouteTableDescription - 라우트 테이블 설명
	RouteTableDescription *string `json:"routeTableDescription,omitempty"`
}

// RouteTableListResponse ...
type RouteTableListResponse struct {
	autorest.Response         `json:"-"`
	GetRouteTableListResponse *RouteTableListResponseGetRouteTableListResponse `json:"getRouteTableListResponse,omitempty"`
}

// RouteTableListResponseGetRouteTableListResponse ...
type RouteTableListResponseGetRouteTableListResponse struct {
	// ReturnCode - 리턴 코드
	ReturnCode *string `json:"returnCode,omitempty"`
	// ReturnMessage - 리턴 메시지
	ReturnMessage *string `json:"returnMessage,omitempty"`
	// TotalRows - 총 행 개수
	TotalRows *int32 `json:"totalRows,omitempty"`
	// RouteTableList - 라우트 테이블 리스트
	RouteTableList *[]RouteTableList `json:"routeTableList,omitempty"`
}

// RouteTableStatus ...
type RouteTableStatus struct {
	// Code - 상태 코드
	Code *string `json:"code,omitempty"`
	// CodeName - 상태 코드 이름
	CodeName *string `json:"codeName,omitempty"`
}

// RouteTableSubnetAddResponse ...
type RouteTableSubnetAddResponse struct {
	autorest.Response           `json:"-"`
	AddRouteTableSubnetResponse *RouteTableSubnetAddResponseAddRouteTableSubnetResponse `json:"addRouteTableSubnetResponse,omitempty"`
}

// RouteTableSubnetAddResponseAddRouteTableSubnetResponse ...
type RouteTableSubnetAddResponseAddRouteTableSubnetResponse struct {
	// ReturnCode - 리턴 코드
	ReturnCode *string `json:"returnCode,omitempty"`
	// ReturnMessage - 리턴 메시지
	ReturnMessage *string `json:"returnMessage,omitempty"`
	// TotalRows - 총 행 개수
	TotalRows *int32 `json:"totalRows,omitempty"`
	// SubnetList - 서브넷 리스트
	SubnetList *[]SubnetList `json:"subnetList,omitempty"`
}

// RouteTableSubnetListResponse ...
type RouteTableSubnetListResponse struct {
	autorest.Response               `json:"-"`
	GetRouteTableSubnetListResponse *RouteTableSubnetListResponseGetRouteTableSubnetListResponse `json:"getRouteTableSubnetListResponse,omitempty"`
}

// RouteTableSubnetListResponseGetRouteTableSubnetListResponse ...
type RouteTableSubnetListResponseGetRouteTableSubnetListResponse struct {
	// ReturnCode - 리턴 코드
	ReturnCode *string `json:"returnCode,omitempty"`
	// ReturnMessage - 리턴 메시지
	ReturnMessage *string `json:"returnMessage,omitempty"`
	// TotalRows - 총 행 개수
	TotalRows *int32 `json:"totalRows,omitempty"`
	// SubnetList - 서브넷 리스트
	SubnetList *[]SubnetList `json:"subnetList,omitempty"`
}

// RouteTableSubnetRemoveResponse ...
type RouteTableSubnetRemoveResponse struct {
	autorest.Response              `json:"-"`
	RemoveRouteTableSubnetResponse *RouteTableSubnetRemoveResponseRemoveRouteTableSubnetResponse `json:"removeRouteTableSubnetResponse,omitempty"`
}

// RouteTableSubnetRemoveResponseRemoveRouteTableSubnetResponse ...
type RouteTableSubnetRemoveResponseRemoveRouteTableSubnetResponse struct {
	// ReturnCode - 리턴 코드
	ReturnCode *string `json:"returnCode,omitempty"`
	// ReturnMessage - 리턴 메시지
	ReturnMessage *string `json:"returnMessage,omitempty"`
	// TotalRows - 총 행 개수
	TotalRows *int32 `json:"totalRows,omitempty"`
	// SubnetList - 서브넷 리스트
	SubnetList *[]SubnetList `json:"subnetList,omitempty"`
}

// RuleAction ...
type RuleAction struct {
	// Code - 상태 코드
	Code *string `json:"code,omitempty"`
	// CodeName - 상태 코드 이름
	CodeName *string `json:"codeName,omitempty"`
}

// Status ...
type Status struct {
	// Code - 상태 코드
	Code *string `json:"code,omitempty"`
	// CodeName - 상태 코드 이름
	CodeName *string `json:"codeName,omitempty"`
}

// SubnetCreateResponse ...
type SubnetCreateResponse struct {
	autorest.Response    `json:"-"`
	CreateSubnetResponse *SubnetCreateResponseCreateSubnetResponse `json:"createSubnetResponse,omitempty"`
}

// SubnetCreateResponseCreateSubnetResponse ...
type SubnetCreateResponseCreateSubnetResponse struct {
	// ReturnCode - 리턴 코드
	ReturnCode *string `json:"returnCode,omitempty"`
	// ReturnMessage - 리턴 메시지
	ReturnMessage *string `json:"returnMessage,omitempty"`
	// TotalRows - 총 행 개수
	TotalRows *int32 `json:"totalRows,omitempty"`
	// SubnetList - 서브넷 리스트
	SubnetList *[]SubnetList `json:"subnetList,omitempty"`
}

// SubnetDeleteResponse ...
type SubnetDeleteResponse struct {
	autorest.Response    `json:"-"`
	DeleteSubnetResponse *SubnetDeleteResponseDeleteSubnetResponse `json:"deleteSubnetResponse,omitempty"`
}

// SubnetDeleteResponseDeleteSubnetResponse ...
type SubnetDeleteResponseDeleteSubnetResponse struct {
	// ReturnCode - 리턴 코드
	ReturnCode *string `json:"returnCode,omitempty"`
	// ReturnMessage - 리턴 메시지
	ReturnMessage *string `json:"returnMessage,omitempty"`
	// TotalRows - 총 행 개수
	TotalRows *int32 `json:"totalRows,omitempty"`
	// SubnetList - 서브넷 리스트
	SubnetList *[]SubnetList `json:"subnetList,omitempty"`
}

// SubnetDetailResponse ...
type SubnetDetailResponse struct {
	autorest.Response       `json:"-"`
	GetSubnetDetailResponse *SubnetDetailResponseGetSubnetDetailResponse `json:"getSubnetDetailResponse,omitempty"`
}

// SubnetDetailResponseGetSubnetDetailResponse ...
type SubnetDetailResponseGetSubnetDetailResponse struct {
	// ReturnCode - 리턴 코드
	ReturnCode *string `json:"returnCode,omitempty"`
	// ReturnMessage - 리턴 메시지
	ReturnMessage *string `json:"returnMessage,omitempty"`
	// TotalRows - 총 행 개수
	TotalRows *int32 `json:"totalRows,omitempty"`
	// SubnetList - 서브넷 리스트
	SubnetList *[]SubnetList `json:"subnetList,omitempty"`
}

// SubnetList ...
type SubnetList struct {
	// SubnetNo - 서브넷 번호
	SubnetNo *string `json:"subnetNo,omitempty"`
	// VpcNo - VPC 번호
	VpcNo *string `json:"vpcNo,omitempty"`
	// ZoneCode - ZONE 코드
	ZoneCode *string `json:"zoneCode,omitempty"`
	// SubnetName - 서브넷 이름
	SubnetName *string `json:"subnetName,omitempty"`
	// Subnet - 서브넷 CIDR
	Subnet *string `json:"subnet,omitempty"`
	// SubnetStatus - 서브넷 상태
	SubnetStatus *SubnetStatus `json:"subnetStatus,omitempty"`
	// SubnetType - 서브넷 타입
	SubnetType *SubnetType `json:"subnetType,omitempty"`
}

// SubnetListResponse ...
type SubnetListResponse struct {
	autorest.Response     `json:"-"`
	GetSubnetListResponse *SubnetListResponseGetSubnetListResponse `json:"getSubnetListResponse,omitempty"`
}

// SubnetListResponseGetSubnetListResponse ...
type SubnetListResponseGetSubnetListResponse struct {
	// ReturnCode - 리턴 코드
	ReturnCode *string `json:"returnCode,omitempty"`
	// ReturnMessage - 리턴 메시지
	ReturnMessage *string `json:"returnMessage,omitempty"`
	// TotalRows - 총 행 개수
	TotalRows *int32 `json:"totalRows,omitempty"`
	// SubnetList - 서브넷 리스트
	SubnetList *[]SubnetList `json:"subnetList,omitempty"`
}

// SubnetNetworkACLResponse ...
type SubnetNetworkACLResponse struct {
	autorest.Response           `json:"-"`
	SetSubnetNetworkACLResponse *SubnetNetworkACLResponseSetSubnetNetworkACLResponse `json:"setSubnetNetworkAclResponse,omitempty"`
}

// SubnetNetworkACLResponseSetSubnetNetworkACLResponse ...
type SubnetNetworkACLResponseSetSubnetNetworkACLResponse struct {
	// ReturnCode - 리턴 코드
	ReturnCode *string `json:"returnCode,omitempty"`
	// ReturnMessage - 리턴 메시지
	ReturnMessage *string `json:"returnMessage,omitempty"`
	// TotalRows - 총 행 개수
	TotalRows *int32 `json:"totalRows,omitempty"`
	// NetworkACLList - 네트워크 ACL 리스트
	NetworkACLList *[]NetworkACLList `json:"networkAclList,omitempty"`
}

// SubnetStatus ...
type SubnetStatus struct {
	// Code - 상태 코드
	Code *string `json:"code,omitempty"`
	// CodeName - 상태 코드 이름
	CodeName *string `json:"codeName,omitempty"`
}

// SubnetType ...
type SubnetType struct {
	// Code - 상태 코드
	Code *string `json:"code,omitempty"`
	// CodeName - 상태 코드 이름
	CodeName *string `json:"codeName,omitempty"`
}

// SupportedSubnetType ...
type SupportedSubnetType struct {
	// Code - 상태 코드
	Code *string `json:"code,omitempty"`
	// CodeName - 상태 코드 이름
	CodeName *string `json:"codeName,omitempty"`
}

// TargetType ...
type TargetType struct {
	// Code - 상태 코드
	Code *string `json:"code,omitempty"`
	// CodeName - 상태 코드 이름
	CodeName *string `json:"codeName,omitempty"`
}

// UsageType ...
type UsageType struct {
	// Code - 상태 코드
	Code *string `json:"code,omitempty"`
	// CodeName - 상태 코드 이름
	CodeName *string `json:"codeName,omitempty"`
}
