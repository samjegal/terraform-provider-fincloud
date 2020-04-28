package sslvpn

// FINCLOUD_APACHE_NO_VERSION

import (
	"github.com/Azure/go-autorest/autorest"
)

// The package's fully qualified name.
const fqdn = "github.com/samjegal/fincloud-sdk-for-go/services/sslvpn"

// StatusCode enumerates the values for status code.
type StatusCode string

const (
	// RUN ...
	RUN StatusCode = "RUN"
	// SET ...
	SET StatusCode = "SET"
)

// PossibleStatusCodeValues returns an array of possible values for the StatusCode const type.
func PossibleStatusCodeValues() []StatusCode {
	return []StatusCode{RUN, SET}
}

// StatusName enumerates the values for status name.
type StatusName string

const (
	// 설정중 ...
	설정중 StatusName = "설정중"
	// 운영중 ...
	운영중 StatusName = "운영중"
)

// PossibleStatusNameValues returns an array of possible values for the StatusName const type.
func PossibleStatusNameValues() []StatusName {
	return []StatusName{설정중, 운영중}
}

// ContentParameter ...
type ContentParameter struct {
	// VpcNo - VPC 번호
	VpcNo *int32 `json:"vpcNo,omitempty"`
	// VpcName - VPC 이름
	VpcName *string `json:"vpcName,omitempty"`
	// InstanceNo - SSL VPN 번호
	InstanceNo *int32 `json:"instanceNo,omitempty"`
	// MemberNo - SSL VPN 멤버 번호
	MemberNo *int32 `json:"memberNo,omitempty"`
	// SslVpnName - SSL VPN 이름
	SslVpnName *string `json:"sslVpnName,omitempty"`
	// StatusCode - SSL VPN 상태 코드. Possible values include: 'RUN', 'SET'
	StatusCode StatusCode `json:"statusCode,omitempty"`
	// StatusName - SSL VPN 상태 이름. Possible values include: '운영중', '설정중'
	StatusName StatusName `json:"statusName,omitempty"`
	// UserCount - SSL VPN 등록된 유저수
	UserCount *int32 `json:"userCount,omitempty"`
	// SslVpnUser - SSL VPN 유저 리스트
	SslVpnUser *[]UserContentParameter `json:"sslVpnUser,omitempty"`
	// UserCountLimitation - SSL VPN 최대 등록 가능한 유저수
	UserCountLimitation *int32 `json:"userCountLimitation,omitempty"`
	// CreatedYmdt - SSL VPN 생성일자
	CreatedYmdt *float64 `json:"createdYmdt,omitempty"`
}

// LimitUserCountParameter ...
type LimitUserCountParameter struct {
	// UserCountLimitation - 사용자 등록 가능 수
	UserCountLimitation *int32 `json:"userCountLimitation,omitempty"`
}

// Parameter ...
type Parameter struct {
	autorest.Response `json:"-"`
	// Content - SSL VPN 컨텐츠 리스트
	Content *[]ContentParameter `json:"content,omitempty"`
	// Total - 전체 SSL VPN 개수
	Total *int32 `json:"total,omitempty"`
	// UserRequestID - SSL VPN UUID
	UserRequestID *string `json:"userRequestId,omitempty"`
}

// UserContentParameter ...
type UserContentParameter struct {
	// UserName - 사용자 이름
	UserName *string `json:"userName,omitempty"`
	// Password - 비밀번호
	Password *string `json:"password,omitempty"`
	// UserSeq - 사용자 번호
	UserSeq interface{} `json:"userSeq,omitempty"`
	// RegionNo - 리전 번호
	RegionNo *int32 `json:"regionNo,omitempty"`
	// InstanceNo - SSL VPN 번호
	InstanceNo *int32 `json:"instanceNo,omitempty"`
	// CountryNo - 국가 번호
	CountryNo *string `json:"countryNo,omitempty"`
	// CellphoneNo - 전화번호
	CellphoneNo *string `json:"cellphoneNo,omitempty"`
	// Email - 이메일 주소
	Email *string `json:"email,omitempty"`
	// OperationCode - 운영코드
	OperationCode *string `json:"operationCode,omitempty"`
}
