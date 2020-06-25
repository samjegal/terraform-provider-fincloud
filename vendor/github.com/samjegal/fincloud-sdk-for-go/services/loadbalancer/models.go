package loadbalancer

// FINCLOUD_APACHE_NO_VERSION

import (
	"github.com/Azure/go-autorest/autorest"
)

// The package's fully qualified name.
const fqdn = "github.com/samjegal/fincloud-sdk-for-go/services/loadbalancer"

// AlgorithmTypeCode enumerates the values for algorithm type code.
type AlgorithmTypeCode string

const (
	// MH ...
	MH AlgorithmTypeCode = "MH"
	// RR ...
	RR AlgorithmTypeCode = "RR"
)

// PossibleAlgorithmTypeCodeValues returns an array of possible values for the AlgorithmTypeCode const type.
func PossibleAlgorithmTypeCodeValues() []AlgorithmTypeCode {
	return []AlgorithmTypeCode{MH, RR}
}

// HealthCheckStatusCode enumerates the values for health check status code.
type HealthCheckStatusCode string

const (
	// AVAIL ...
	AVAIL HealthCheckStatusCode = "AVAIL"
)

// PossibleHealthCheckStatusCodeValues returns an array of possible values for the HealthCheckStatusCode const type.
func PossibleHealthCheckStatusCodeValues() []HealthCheckStatusCode {
	return []HealthCheckStatusCode{AVAIL}
}

// IPTypeCode enumerates the values for ip type code.
type IPTypeCode string

const (
	// PRIVATE ...
	PRIVATE IPTypeCode = "PRIVATE"
	// PUBLIC ...
	PUBLIC IPTypeCode = "PUBLIC"
)

// PossibleIPTypeCodeValues returns an array of possible values for the IPTypeCode const type.
func PossibleIPTypeCodeValues() []IPTypeCode {
	return []IPTypeCode{PRIVATE, PUBLIC}
}

// LayerTypeCode enumerates the values for layer type code.
type LayerTypeCode string

const (
	// APPLICATION ...
	APPLICATION LayerTypeCode = "APPLICATION"
	// NETWORK ...
	NETWORK LayerTypeCode = "NETWORK"
	// NETWORKPROXY ...
	NETWORKPROXY LayerTypeCode = "NETWORK_PROXY"
)

// PossibleLayerTypeCodeValues returns an array of possible values for the LayerTypeCode const type.
func PossibleLayerTypeCodeValues() []LayerTypeCode {
	return []LayerTypeCode{APPLICATION, NETWORK, NETWORKPROXY}
}

// OperationCode enumerates the values for operation code.
type OperationCode string

const (
	// CHANG ...
	CHANG OperationCode = "CHANG"
	// CREAT ...
	CREAT OperationCode = "CREAT"
	// NULL ...
	NULL OperationCode = "NULL"
	// TERMT ...
	TERMT OperationCode = "TERMT"
)

// PossibleOperationCodeValues returns an array of possible values for the OperationCode const type.
func PossibleOperationCodeValues() []OperationCode {
	return []OperationCode{CHANG, CREAT, NULL, TERMT}
}

// ProtocolCode enumerates the values for protocol code.
type ProtocolCode string

const (
	// ICMP ...
	ICMP ProtocolCode = "ICMP"
	// TCP ...
	TCP ProtocolCode = "TCP"
	// UDP ...
	UDP ProtocolCode = "UDP"
)

// PossibleProtocolCodeValues returns an array of possible values for the ProtocolCode const type.
func PossibleProtocolCodeValues() []ProtocolCode {
	return []ProtocolCode{ICMP, TCP, UDP}
}

// ServerInstanceStatusCode enumerates the values for server instance status code.
type ServerInstanceStatusCode string

const (
	// ServerInstanceStatusCodeCREAT ...
	ServerInstanceStatusCodeCREAT ServerInstanceStatusCode = "CREAT"
	// ServerInstanceStatusCodeINIT ...
	ServerInstanceStatusCodeINIT ServerInstanceStatusCode = "INIT"
	// ServerInstanceStatusCodeRUN ...
	ServerInstanceStatusCodeRUN ServerInstanceStatusCode = "RUN"
	// ServerInstanceStatusCodeSTOP ...
	ServerInstanceStatusCodeSTOP ServerInstanceStatusCode = "STOP"
)

// PossibleServerInstanceStatusCodeValues returns an array of possible values for the ServerInstanceStatusCode const type.
func PossibleServerInstanceStatusCodeValues() []ServerInstanceStatusCode {
	return []ServerInstanceStatusCode{ServerInstanceStatusCodeCREAT, ServerInstanceStatusCodeINIT, ServerInstanceStatusCodeRUN, ServerInstanceStatusCodeSTOP}
}

// ServerProtocolCode enumerates the values for server protocol code.
type ServerProtocolCode string

const (
	// HTTP ...
	HTTP ServerProtocolCode = "HTTP"
	// HTTPS ...
	HTTPS ServerProtocolCode = "HTTPS"
)

// PossibleServerProtocolCodeValues returns an array of possible values for the ServerProtocolCode const type.
func PossibleServerProtocolCodeValues() []ServerProtocolCode {
	return []ServerProtocolCode{HTTP, HTTPS}
}

// StatusCode enumerates the values for status code.
type StatusCode string

const (
	// INIT ...
	INIT StatusCode = "INIT"
	// USED ...
	USED StatusCode = "USED"
)

// PossibleStatusCodeValues returns an array of possible values for the StatusCode const type.
func PossibleStatusCodeValues() []StatusCode {
	return []StatusCode{INIT, USED}
}

// StatusName enumerates the values for status name.
type StatusName string

const (
	// 변경중 ...
	변경중 StatusName = "변경중"
	// 삭제중 ...
	삭제중 StatusName = "삭제중"
	// 생성중 ...
	생성중 StatusName = "생성중"
	// 운영중 ...
	운영중 StatusName = "운영중"
)

// PossibleStatusNameValues returns an array of possible values for the StatusName const type.
func PossibleStatusNameValues() []StatusName {
	return []StatusName{변경중, 삭제중, 생성중, 운영중}
}

// Throughput enumerates the values for throughput.
type Throughput string

const (
	// LARGE ...
	LARGE Throughput = "LARGE"
	// MEDIUM ...
	MEDIUM Throughput = "MEDIUM"
	// SMALL ...
	SMALL Throughput = "SMALL"
)

// PossibleThroughputValues returns an array of possible values for the Throughput const type.
func PossibleThroughputValues() []Throughput {
	return []Throughput{LARGE, MEDIUM, SMALL}
}

// CheckNameParameter ...
type CheckNameParameter struct {
	autorest.Response `json:"-"`
	// Content - 로드밸런서 이름 적합성 검사 결과값
	Content *bool `json:"content,omitempty"`
}

// InstanceListParameter ...
type InstanceListParameter struct {
	// InstanceNoList - 로드밸런서 인스턴스 번호
	InstanceNoList *[]int32 `json:"instanceNoList,omitempty"`
}

// InstanceParameter ...
type InstanceParameter struct {
	// ServerInstanceNoList - 서버 인스턴스 번호 리스트
	ServerInstanceNoList *[]int32 `json:"serverInstanceNoList,omitempty"`
	// IPTypeCode - 로드밸런서 IP 타입 코드. Possible values include: 'PUBLIC', 'PRIVATE'
	IPTypeCode IPTypeCode `json:"ipTypeCode,omitempty"`
	// LoadBalancerName - 로드밸런서 이름
	LoadBalancerName *string `json:"loadBalancerName,omitempty"`
	// LoadBalancerRuleList - 로드밸런서 룰 리스트
	LoadBalancerRuleList *[]RuleListParameter `json:"loadBalancerRuleList,omitempty"`
	// ZoneList - 금융존 리스트
	ZoneList *[]ZoneListParameter `json:"zoneList,omitempty"`
	// InstanceDescription - 로드밸런서 설명
	InstanceDescription *string `json:"instanceDescription,omitempty"`
	// Throughput - 로드밸런서 처리량. Possible values include: 'SMALL', 'MEDIUM', 'LARGE'
	Throughput Throughput `json:"throughput,omitempty"`
	// VpcNo - VPC 번호
	VpcNo *string `json:"vpcNo,omitempty"`
	// LayerTypeCode - 로드밸런서 레이어 타입 코드. Possible values include: 'NETWORK', 'APPLICATION', 'NETWORKPROXY'
	LayerTypeCode LayerTypeCode `json:"layerTypeCode,omitempty"`
	// AlgorithmTypeCode - 로드밸랜서 알고리즘 타입 코드. Possible values include: 'MH', 'RR'
	AlgorithmTypeCode AlgorithmTypeCode `json:"algorithmTypeCode,omitempty"`
	// HTTPKeepAliveTimeout - HTTP Keep Alive 타임아웃
	HTTPKeepAliveTimeout *int32 `json:"httpKeepAliveTimeout,omitempty"`
}

// ListenerParameter ...
type ListenerParameter struct {
	// LoadBalancerRuleList - 로드밸런서 룰 리스트
	LoadBalancerRuleList *[]RuleListParameter `json:"loadBalancerRuleList,omitempty"`
	// InstanceNo - 서버 인스턴스 번호
	InstanceNo *int32 `json:"instanceNo,omitempty"`
	// AlgorithmTypeCode - 로드밸랜서 알고리즘 타입 코드. Possible values include: 'MH', 'RR'
	AlgorithmTypeCode AlgorithmTypeCode `json:"algorithmTypeCode,omitempty"`
	// HTTPKeepAliveTimeout - HTTP Keep Alive 타임아웃
	HTTPKeepAliveTimeout *int32 `json:"httpKeepAliveTimeout,omitempty"`
}

// RuleListParameter ...
type RuleListParameter struct {
	// LoadBalancerPort - 로드밸런서 포트
	LoadBalancerPort *int32 `json:"loadBalancerPort,omitempty"`
	// ServerPort - 서버 리스닝 포트
	ServerPort *int32 `json:"serverPort,omitempty"`
	// ProtocolCode - 서버 프로토콜 코드. Possible values include: 'ICMP', 'UDP', 'TCP'
	ProtocolCode ProtocolCode `json:"protocolCode,omitempty"`
	// ServerProtocolCode - 서버 프로토콜 코드. Possible values include: 'HTTP', 'HTTPS'
	ServerProtocolCode ServerProtocolCode `json:"serverProtocolCode,omitempty"`
	// HealthCheckPath - Health Check 경로
	HealthCheckPath *string `json:"healthCheckPath,omitempty"`
	// HTTP2UseYn - HTTP/2.0 사용 유무
	HTTP2UseYn *string `json:"http2UseYn,omitempty"`
	// StickySessionUseYn - Sticky 세션 사용 유무
	StickySessionUseYn *string `json:"stickySessionUseYn,omitempty"`
}

// SearchContentParameter ...
type SearchContentParameter struct {
	// InstanceNo - 로드밸런서 번호
	InstanceNo *int32 `json:"instanceNo,omitempty"`
	// InstanceUUID - 로드밸런서 UUID
	InstanceUUID *string `json:"instanceUuid,omitempty"`
	// InstanceStatusCode - 로드밸런서 상태 코드. Possible values include: 'INIT', 'USED'
	InstanceStatusCode StatusCode `json:"instanceStatusCode,omitempty"`
	// OperationCode - 로드밸런서 동작 코드. Possible values include: 'NULL', 'CREAT', 'CHANG', 'TERMT'
	OperationCode OperationCode `json:"operationCode,omitempty"`
	// InstanceStatusName - 로드밸런서 상태 이름. Possible values include: '운영중', '생성중', '변경중', '삭제중'
	InstanceStatusName StatusName `json:"instanceStatusName,omitempty"`
	// MemberNo - 로드밸런서 멤버 번호
	MemberNo *int32 `json:"memberNo,omitempty"`
	// CreateYmdt - 로드밸런서 생성 일자
	CreateYmdt *float64 `json:"createYmdt,omitempty"`
	// OperationYmdt - 로드밸런서 운영 일자
	OperationYmdt *float64 `json:"operationYmdt,omitempty"`
	// LoadBalancerName - 로드밸런서 이름
	LoadBalancerName *string `json:"loadBalancerName,omitempty"`
	// LoadBalancerIP - 로드밸런서 IP 주소
	LoadBalancerIP *string `json:"loadBalancerIp,omitempty"`
	// LayerTypeCode - 로드밸런서 레이어 타입 코드. Possible values include: 'NETWORK', 'APPLICATION', 'NETWORKPROXY'
	LayerTypeCode LayerTypeCode `json:"layerTypeCode,omitempty"`
	// IPTypeCode - 로드밸런서 IP 타입 코드. Possible values include: 'PUBLIC', 'PRIVATE'
	IPTypeCode IPTypeCode `json:"ipTypeCode,omitempty"`
	// AlgorithmTypeCode - 로드밸런서 알고리즘 타입 코드. Possible values include: 'MH', 'RR'
	AlgorithmTypeCode AlgorithmTypeCode `json:"algorithmTypeCode,omitempty"`
	// Throughput - 로드밸런서 처리량. Possible values include: 'SMALL', 'MEDIUM', 'LARGE'
	Throughput Throughput `json:"throughput,omitempty"`
	// VpcNo - VPC 번호
	VpcNo *string `json:"vpcNo,omitempty"`
	// VpcName - VPC 이름
	VpcName *string `json:"vpcName,omitempty"`
	// VpcIpv4Cidr - VPC IP 주소 CIDR
	VpcIpv4Cidr *string `json:"vpcIpv4Cidr,omitempty"`
	// RegionNo - 리전 번호
	RegionNo *int32 `json:"regionNo,omitempty"`
	// RegionName - 리전 이름
	RegionName *string `json:"regionName,omitempty"`
	// RegionCode - 리전 코드
	RegionCode *string `json:"regionCode,omitempty"`
	// ZoneList - 금융존 리스트
	ZoneList *[]ZoneListParameter `json:"zoneList,omitempty"`
	// LoadBalancerRuleList - 로드밸런서 룰 리스트
	LoadBalancerRuleList *[]RuleListParameter `json:"loadBalancerRuleList,omitempty"`
	// ServerInstanceList - 로드밸런서에 적용된 서버 인스턴스 리스트
	ServerInstanceList *[]ServerInstanceParameter `json:"serverInstanceList,omitempty"`
}

// SearchListParameter ...
type SearchListParameter struct {
	autorest.Response `json:"-"`
	// Content - 로드밸런서 컨텐츠 리스트
	Content *[]SearchContentParameter `json:"content,omitempty"`
	// Total - 로드밸런서 전체 개수
	Total *int32 `json:"total,omitempty"`
	// UserRequestID - 로드밸런서 UUID
	UserRequestID *string `json:"userRequestId,omitempty"`
}

// SearchParameter ...
type SearchParameter struct {
	// PageNo - 검색할 로드밸런서 페이지 번호
	PageNo *int32 `json:"pageNo,omitempty"`
	// PageSizeNo - 한 페이지에 나올 로드밸런서 개수
	PageSizeNo *int32 `json:"pageSizeNo,omitempty"`
	// Sort - 페이지 정렬 방법
	Sort *[]string `json:"sort,omitempty"`
}

// ServerInstanceListParameter ...
type ServerInstanceListParameter struct {
	autorest.Response `json:"-"`
	// Content - 로드밸런서 타겟 서버 리스트
	Content *[]TargetServerInstanceListParameter `json:"content,omitempty"`
	// Total - 로드밸런서 타겟 서버 전체 개수
	Total *int32 `json:"total,omitempty"`
	// UserRequestID - 로드밸런서 타겟 서버 UUID
	UserRequestID *string `json:"userRequestId,omitempty"`
}

// ServerInstanceParameter ...
type ServerInstanceParameter struct {
	// InstanceNo - 서버 번호
	InstanceNo *int32 `json:"instanceNo,omitempty"`
	// InstanceStatusCode - 서버 상태 코드. Possible values include: 'ServerInstanceStatusCodeINIT', 'ServerInstanceStatusCodeCREAT', 'ServerInstanceStatusCodeRUN', 'ServerInstanceStatusCodeSTOP'
	InstanceStatusCode ServerInstanceStatusCode `json:"instanceStatusCode,omitempty"`
	// OperationCode - 서버 운영 코드
	OperationCode *string `json:"operationCode,omitempty"`
	// ServerName - 서버 이름
	ServerName *string `json:"serverName,omitempty"`
	// ServerIP - 서버 IP 주소
	ServerIP *string `json:"serverIp,omitempty"`
	// SubnetNo - 서브넷 번호
	SubnetNo *string `json:"subnetNo,omitempty"`
	// Subnet - 서브넷 IP 주소 CIDR
	Subnet *string `json:"subnet,omitempty"`
	// SubnetName - 서브넷 이름
	SubnetName *string `json:"subnetName,omitempty"`
	// ServerInstanceLoadBalancerRuleList - 서버 로드밸런서 리스너 룰 리스트
	ServerInstanceLoadBalancerRuleList *[]ServerInstanceRuleList `json:"serverInstanceLoadBalancerRuleList,omitempty"`
}

// ServerInstanceRuleList ...
type ServerInstanceRuleList struct {
	// LoadBalancerPort - 로드밸런서 포트
	LoadBalancerPort *int32 `json:"loadBalancerPort,omitempty"`
	// ServerPort - 서버 포트
	ServerPort *int32 `json:"serverPort,omitempty"`
	// HealthCheckStatusCode - 상태 체크 코드. Possible values include: 'AVAIL'
	HealthCheckStatusCode HealthCheckStatusCode `json:"healthCheckStatusCode,omitempty"`
	// ProtocolCode - 프로토콜 코드. Possible values include: 'ICMP', 'UDP', 'TCP'
	ProtocolCode ProtocolCode `json:"protocolCode,omitempty"`
}

// ServerParameter ...
type ServerParameter struct {
	// LbInstanceNo - 로드밸런서 인스턴스 번호
	LbInstanceNo *int32 `json:"lbInstanceNo,omitempty"`
	// ServerInstanceNoList - 로드밸런서 서버 인스턴스 번호 리스트
	ServerInstanceNoList *[]int32 `json:"serverInstanceNoList,omitempty"`
}

// SettingParameter ...
type SettingParameter struct {
	// InstanceNo - 로드밸런서 인스턴스 번호
	InstanceNo *int32 `json:"instanceNo,omitempty"`
	// Throughput - 로드밸런서 처리량. Possible values include: 'SMALL', 'MEDIUM', 'LARGE'
	Throughput Throughput `json:"throughput,omitempty"`
	// InstanceDescription - 로드밸런서 인스턴스 설명
	InstanceDescription *string `json:"instanceDescription,omitempty"`
}

// TargetServerInstanceListParameter ...
type TargetServerInstanceListParameter struct {
	// TargetServerInstanceList - 로드밸런서 타겟 서버
	TargetServerInstanceList *[]TargetServerInstanceParameter `json:"targetServerInstanceList,omitempty"`
}

// TargetServerInstanceParameter ...
type TargetServerInstanceParameter struct {
	// Disabled - 서버 활성화 여부
	Disabled *bool `json:"disabled,omitempty"`
	// ActionName - 서버 액션 이름
	ActionName *string `json:"actionName,omitempty"`
	// Permission - 서버 권한
	Permission *string `json:"permission,omitempty"`
	// InstanceNo - 서버 인스턴스 번호
	InstanceNo *int32 `json:"instanceNo,omitempty"`
	// InstanceStatusCode - 서버 인스턴스 상태 코드
	InstanceStatusCode *string `json:"instanceStatusCode,omitempty"`
	// OperationCode - 서버 인스턴스 운영 코드
	OperationCode *string `json:"operationCode,omitempty"`
	// InstanceStatusName - 서버 인스턴스 상태 이름
	InstanceStatusName *string `json:"instanceStatusName,omitempty"`
	// ServerName - 서버 이름
	ServerName *string `json:"serverName,omitempty"`
	// ServerIP - 서버 IP 주소
	ServerIP *string `json:"serverIp,omitempty"`
	// SubnetNo - 서버 서브넷 번호
	SubnetNo *string `json:"subnetNo,omitempty"`
	// Subnet - 서브넷 IP 주소 CIDR
	Subnet *string `json:"subnet,omitempty"`
	// SubnetName - 서브넷 이름
	SubnetName *string `json:"subnetName,omitempty"`
	// ZoneName - 금융존 이름
	ZoneName *string `json:"zoneName,omitempty"`
	// ZoneNo - 금융존 번호
	ZoneNo *string `json:"zoneNo,omitempty"`
	// ServerInstanceTypeName - 서버 인스턴스 타입 이름
	ServerInstanceTypeName *string `json:"serverInstanceTypeName,omitempty"`
	// NodeRoleName - 서버 노드 롤 이름
	NodeRoleName *string `json:"nodeRoleName,omitempty"`
}

// ZoneContentParameter ...
type ZoneContentParameter struct {
	// ZoneNo - 금융존 번호
	ZoneNo *int32 `json:"zoneNo,omitempty"`
	// ZoneName - 금융존 이름
	ZoneName *string `json:"zoneName,omitempty"`
}

// ZoneListParameter ...
type ZoneListParameter struct {
	// ZoneNo - 금융존 번호
	ZoneNo *int32 `json:"zoneNo,omitempty"`
	// SubnetNo - 서브넷 번호
	SubnetNo *string `json:"subnetNo,omitempty"`
	// ZoneName - 금융존 이름
	ZoneName *string `json:"zoneName,omitempty"`
	// SubnetName - 서브넷 이름
	SubnetName *string `json:"subnetName,omitempty"`
	// SubnetIpv4Cidr - 서브넷 IP 주소 CIDR
	SubnetIpv4Cidr *string `json:"subnetIpv4Cidr,omitempty"`
}

// ZoneSubnetParameter ...
type ZoneSubnetParameter struct {
	autorest.Response `json:"-"`
	// Content - 로드밸런서 금융존 콘텐츠
	Content *[]ZoneContentParameter `json:"content,omitempty"`
	// Total - 로드밸런서 금융존 개수
	Total *int32 `json:"total,omitempty"`
	// UserRequestID - 로드밸런서 금융존 UUID
	UserRequestID *string `json:"userRequestId,omitempty"`
}
