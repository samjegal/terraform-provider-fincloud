package vpc

// FINCLOUD_APACHE_NO_VERSION

// NatGatewayInstanceStatusCode enumerates the values for nat gateway instance status code.
type NatGatewayInstanceStatusCode string

const (
	// INIT ...
	INIT NatGatewayInstanceStatusCode = "INIT"
	// RUN ...
	RUN NatGatewayInstanceStatusCode = "RUN"
	// TERMTING ...
	TERMTING NatGatewayInstanceStatusCode = "TERMTING"
)

// PossibleNatGatewayInstanceStatusCodeValues returns an array of possible values for the NatGatewayInstanceStatusCode const type.
func PossibleNatGatewayInstanceStatusCodeValues() []NatGatewayInstanceStatusCode {
	return []NatGatewayInstanceStatusCode{INIT, RUN, TERMTING}
}

// NetworkACLRuleTypeCode enumerates the values for network acl rule type code.
type NetworkACLRuleTypeCode string

const (
	// INBND ...
	INBND NetworkACLRuleTypeCode = "INBND"
	// OTBND ...
	OTBND NetworkACLRuleTypeCode = "OTBND"
)

// PossibleNetworkACLRuleTypeCodeValues returns an array of possible values for the NetworkACLRuleTypeCode const type.
func PossibleNetworkACLRuleTypeCodeValues() []NetworkACLRuleTypeCode {
	return []NetworkACLRuleTypeCode{INBND, OTBND}
}

// NetworkACLStatusCode enumerates the values for network acl status code.
type NetworkACLStatusCode string

const (
	// NetworkACLStatusCodeINIT ...
	NetworkACLStatusCodeINIT NetworkACLStatusCode = "INIT"
	// NetworkACLStatusCodeRUN ...
	NetworkACLStatusCodeRUN NetworkACLStatusCode = "RUN"
	// NetworkACLStatusCodeSET ...
	NetworkACLStatusCodeSET NetworkACLStatusCode = "SET"
	// NetworkACLStatusCodeTERMTING ...
	NetworkACLStatusCodeTERMTING NetworkACLStatusCode = "TERMTING"
)

// PossibleNetworkACLStatusCodeValues returns an array of possible values for the NetworkACLStatusCode const type.
func PossibleNetworkACLStatusCodeValues() []NetworkACLStatusCode {
	return []NetworkACLStatusCode{NetworkACLStatusCodeINIT, NetworkACLStatusCodeRUN, NetworkACLStatusCodeSET, NetworkACLStatusCodeTERMTING}
}

// PeeringInstanceStatusCode enumerates the values for peering instance status code.
type PeeringInstanceStatusCode string

const (
	// PeeringInstanceStatusCodeINIT ...
	PeeringInstanceStatusCodeINIT PeeringInstanceStatusCode = "INIT"
	// PeeringInstanceStatusCodeRUN ...
	PeeringInstanceStatusCodeRUN PeeringInstanceStatusCode = "RUN"
	// PeeringInstanceStatusCodeTERMTING ...
	PeeringInstanceStatusCodeTERMTING PeeringInstanceStatusCode = "TERMTING"
)

// PossiblePeeringInstanceStatusCodeValues returns an array of possible values for the PeeringInstanceStatusCode const type.
func PossiblePeeringInstanceStatusCodeValues() []PeeringInstanceStatusCode {
	return []PeeringInstanceStatusCode{PeeringInstanceStatusCodeINIT, PeeringInstanceStatusCodeRUN, PeeringInstanceStatusCodeTERMTING}
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

// RuleActionCode enumerates the values for rule action code.
type RuleActionCode string

const (
	// ALLOW ...
	ALLOW RuleActionCode = "ALLOW"
	// DROP ...
	DROP RuleActionCode = "DROP"
)

// PossibleRuleActionCodeValues returns an array of possible values for the RuleActionCode const type.
func PossibleRuleActionCodeValues() []RuleActionCode {
	return []RuleActionCode{ALLOW, DROP}
}

// SortedBy enumerates the values for sorted by.
type SortedBy string

const (
	// SourceVpcName ...
	SourceVpcName SortedBy = "sourceVpcName"
	// TargetVpcName ...
	TargetVpcName SortedBy = "targetVpcName"
	// VpcPeeringName ...
	VpcPeeringName SortedBy = "vpcPeeringName"
)

// PossibleSortedByValues returns an array of possible values for the SortedBy const type.
func PossibleSortedByValues() []SortedBy {
	return []SortedBy{SourceVpcName, TargetVpcName, VpcPeeringName}
}

// SortingOrder enumerates the values for sorting order.
type SortingOrder string

const (
	// ASC ...
	ASC SortingOrder = "ASC"
	// DESC ...
	DESC SortingOrder = "DESC"
)

// PossibleSortingOrderValues returns an array of possible values for the SortingOrder const type.
func PossibleSortingOrderValues() []SortingOrder {
	return []SortingOrder{ASC, DESC}
}

// StatusCode enumerates the values for status code.
type StatusCode string

const (
	// StatusCodeCREATING ...
	StatusCodeCREATING StatusCode = "CREATING"
	// StatusCodeINIT ...
	StatusCodeINIT StatusCode = "INIT"
	// StatusCodeRUN ...
	StatusCodeRUN StatusCode = "RUN"
	// StatusCodeTERMTING ...
	StatusCodeTERMTING StatusCode = "TERMTING"
)

// PossibleStatusCodeValues returns an array of possible values for the StatusCode const type.
func PossibleStatusCodeValues() []StatusCode {
	return []StatusCode{StatusCodeCREATING, StatusCodeINIT, StatusCodeRUN, StatusCodeTERMTING}
}

// SubnetStatusCode enumerates the values for subnet status code.
type SubnetStatusCode string

const (
	// SubnetStatusCodeCREATING ...
	SubnetStatusCodeCREATING SubnetStatusCode = "CREATING"
	// SubnetStatusCodeINIT ...
	SubnetStatusCodeINIT SubnetStatusCode = "INIT"
	// SubnetStatusCodeRUN ...
	SubnetStatusCodeRUN SubnetStatusCode = "RUN"
	// SubnetStatusCodeTERMTING ...
	SubnetStatusCodeTERMTING SubnetStatusCode = "TERMTING"
)

// PossibleSubnetStatusCodeValues returns an array of possible values for the SubnetStatusCode const type.
func PossibleSubnetStatusCodeValues() []SubnetStatusCode {
	return []SubnetStatusCode{SubnetStatusCodeCREATING, SubnetStatusCodeINIT, SubnetStatusCodeRUN, SubnetStatusCodeTERMTING}
}

// SubnetTypeCode enumerates the values for subnet type code.
type SubnetTypeCode string

const (
	// PRIVATE ...
	PRIVATE SubnetTypeCode = "PRIVATE"
	// PUBLIC ...
	PUBLIC SubnetTypeCode = "PUBLIC"
)

// PossibleSubnetTypeCodeValues returns an array of possible values for the SubnetTypeCode const type.
func PossibleSubnetTypeCodeValues() []SubnetTypeCode {
	return []SubnetTypeCode{PRIVATE, PUBLIC}
}

// SupportedSubnetTypeCode enumerates the values for supported subnet type code.
type SupportedSubnetTypeCode string

const (
	// SupportedSubnetTypeCodePRIVATE ...
	SupportedSubnetTypeCodePRIVATE SupportedSubnetTypeCode = "PRIVATE"
	// SupportedSubnetTypeCodePUBLIC ...
	SupportedSubnetTypeCodePUBLIC SupportedSubnetTypeCode = "PUBLIC"
)

// PossibleSupportedSubnetTypeCodeValues returns an array of possible values for the SupportedSubnetTypeCode const type.
func PossibleSupportedSubnetTypeCodeValues() []SupportedSubnetTypeCode {
	return []SupportedSubnetTypeCode{SupportedSubnetTypeCodePRIVATE, SupportedSubnetTypeCodePUBLIC}
}

// TargetTypeCode enumerates the values for target type code.
type TargetTypeCode string

const (
	// NATGW ...
	NATGW TargetTypeCode = "NATGW"
	// VGW ...
	VGW TargetTypeCode = "VGW"
	// VPCPEERING ...
	VPCPEERING TargetTypeCode = "VPCPEERING"
)

// PossibleTargetTypeCodeValues returns an array of possible values for the TargetTypeCode const type.
func PossibleTargetTypeCodeValues() []TargetTypeCode {
	return []TargetTypeCode{NATGW, VGW, VPCPEERING}
}

// UsageTypeCode enumerates the values for usage type code.
type UsageTypeCode string

const (
	// BM ...
	BM UsageTypeCode = "BM"
	// GEN ...
	GEN UsageTypeCode = "GEN"
	// LOADB ...
	LOADB UsageTypeCode = "LOADB"
)

// PossibleUsageTypeCodeValues returns an array of possible values for the UsageTypeCode const type.
func PossibleUsageTypeCodeValues() []UsageTypeCode {
	return []UsageTypeCode{BM, GEN, LOADB}
}
