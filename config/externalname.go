/*
Copyright 2022 Upbound Inc.
*/

package config

import (
	"context"
	"fmt"
	"strings"

	"github.com/crossplane/crossplane-runtime/pkg/errors"

	"github.com/upbound/upjet/pkg/config"

	"github.com/upbound/provider-aws/config/common"
)

// ExternalNameConfigs contains all external name configurations for this
// provider.
var ExternalNameConfigs = map[string]config.ExternalName{

	// ACM
	// Imported using ARN that has a random substring:
	// arn:aws:acm:eu-central-1:123456789012:certificate/7e7a28d2-163f-4b8f-b9cd-822f96c08d6a
	"aws_acm_certificate": config.IdentifierFromProvider,
	// No import documented, but https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/acm_certificate_validation#id
	"aws_acm_certificate_validation": config.IdentifierFromProvider,

	// ACM PCA
	// aws_acmpca_certificate can not be imported at this time.
	"aws_acmpca_certificate": config.IdentifierFromProvider,
	// Imported using ARN that has a random substring:
	//	// arn:aws:acm-pca:us-east-1:123456789012:certificate-authority/12345678-1234-1234-1234-123456789012
	"aws_acmpca_certificate_authority": config.IdentifierFromProvider,
	// No doc on import, but resource is getting CA ARN:
	// arn:aws:acm-pca:eu-central-1:609897127049:certificate-authority/ba0c7989-9641-4f36-a033-dee60121d595
	"aws_acmpca_certificate_authority_certificate": config.IdentifierFromProvider,

	// amp
	//
	// ID is a random UUID.
	"aws_prometheus_workspace":            config.IdentifierFromProvider,
	"aws_prometheus_rule_group_namespace": config.TemplatedStringAsIdentifier("name", "arn:aws:aps:{{ .parameters.region }}:{{ .client_metadata.account_id }}:rulegroupsnamespace/IDstring/{{ .external_name }}"),
	// Uses the ID of workspace, workspace_id parameter.
	"aws_prometheus_alert_manager_definition": config.IdentifierFromProvider,

	// apigatewayv2
	//
	"aws_apigatewayv2_api": config.IdentifierFromProvider,
	// Case4: Imported by using the API mapping identifier and domain name.
	"aws_apigatewayv2_api_mapping": TemplatedStringAsIdentifierWithNoName("{{ .external_name }}/{{ .parameters.domain_name }}"),
	// Case4: Imported by using the API identifier and authorizer identifier.
	"aws_apigatewayv2_authorizer": TemplatedStringAsIdentifierWithNoName("{{ .parameters.api_id }}/{{ .external_name }}"),
	// Case4: Imported by using the API identifier and deployment identifier.
	"aws_apigatewayv2_deployment":  TemplatedStringAsIdentifierWithNoName("{{ .parameters.api_id }}/{{ .external_name }}"),
	"aws_apigatewayv2_domain_name": config.ParameterAsIdentifier("domain_name"),
	// Case4: Imported by using the API identifier and integration identifier.
	"aws_apigatewayv2_integration": TemplatedStringAsIdentifierWithNoName("{{ .parameters.api_id }}/{{ .external_name }}"),
	// Case4: Imported by using the API identifier, integration identifier and
	// integration response identifier.
	"aws_apigatewayv2_integration_response": TemplatedStringAsIdentifierWithNoName("{{ .parameters.api_id }}/{{ .parameters.integration_id }}/{{ .external_name }}"),
	// Case4: Imported by using the API identifier and model identifier.
	"aws_apigatewayv2_model": TemplatedStringAsIdentifierWithNoName("{{ .parameters.api_id }}/{{ .external_name }}"),
	// Case4: Imported by using the API identifier and route identifier.
	"aws_apigatewayv2_route": TemplatedStringAsIdentifierWithNoName("{{ .parameters.api_id }}/{{ .external_name }}"),
	// Case4: Imported by using the API identifier, route identifier and route
	// response identifier.
	"aws_apigatewayv2_route_response": TemplatedStringAsIdentifierWithNoName("{{ .parameters.api_id }}/{{ .parameters.route_id }}/{{ .external_name }}"),
	// Imported by using the API identifier and stage name.
	"aws_apigatewayv2_stage": config.TemplatedStringAsIdentifier("name", "{{ .parameters.api_id }}/{{ .external_name }}"),
	// aws_apigatewayv2_vpc_link can be imported by using the VPC Link id
	"aws_apigatewayv2_vpc_link": config.IdentifierFromProvider,

	// autoscaling
	//
	"aws_autoscaling_group": config.NameAsIdentifier,
	// No terraform import.
	"aws_autoscaling_attachment": config.IdentifierFromProvider,

	// DynamoDB Table Items can be imported using the name
	"aws_dynamodb_table_item": config.IdentifierFromProvider,
	// DynamoDB contributor insights
	"aws_dynamodb_contributor_insights": config.IdentifierFromProvider,
	// Dynamodb Kinesis streaming destinations are imported using "table_name,stream_arn"
	"aws_dynamodb_kinesis_streaming_destination": config.IdentifierFromProvider,

	// cognitoidentity
	//
	// us-west-2_abc123
	"aws_cognito_identity_pool": config.IdentifierFromProvider,
	// us-west-2:b64805ad-cb56-40ba-9ffc-f5d8207e6d42
	"aws_cognito_identity_pool_roles_attachment": config.IdentifierFromProvider,
	// us-west-2_abc123:CorpAD
	"aws_cognito_identity_pool_provider_principal_tag": config.IdentifierFromProvider,

	// cognitoidp
	//
	// us-west-2_abc123
	"aws_cognito_user_pool": config.IdentifierFromProvider,
	// us-west-2_abc123/3ho4ek12345678909nh3fmhpko
	"aws_cognito_user_pool_client": config.IdentifierFromProvider,
	// auth.example.org
	"aws_cognito_user_pool_domain": config.IdentifierFromProvider,
	// us-west-2_ZCTarbt5C,12bu4fuk3mlgqa2rtrujgp6egq
	"aws_cognito_user_pool_ui_customization": config.IdentifierFromProvider,
	// aws_cognito_user_group.group us-east-1_vG78M4goG/user-group
	// disabled until the fix of https://github.com/upbound/official-providers/issues/531
	// "aws_cognito_user_group": config.IdentifierFromProvider,
	// us-west-2_abc123|https://example.com
	"aws_cognito_resource_server": config.IdentifierFromProvider,
	// us-west-2_abc123:CorpAD
	"aws_cognito_identity_provider": config.IdentifierFromProvider,
	// user_pool_id/name: us-east-1_vG78M4goG/user
	"aws_cognito_user": config.TemplatedStringAsIdentifier("username", "{{ .parameters.user_pool_id }}/{{ .external_name }}"),
	// no doc
	// disabled until the fix of https://github.com/upbound/official-providers/issues/531
	// "aws_cognito_user_in_group": config.IdentifierFromProvider,

	// ebs
	//
	// EBS Volumes can be imported using the id: vol-049df61146c4d7901
	"aws_ebs_volume": config.IdentifierFromProvider,
	// EBS Snapshot can be imported using the id
	"aws_ebs_snapshot": config.IdentifierFromProvider,
	// No import
	"aws_ebs_snapshot_copy": config.IdentifierFromProvider,
	// No import
	"aws_ebs_snapshot_import": config.IdentifierFromProvider,

	// ec2
	//
	// Instances can be imported using the id: i-12345678
	"aws_instance": config.IdentifierFromProvider,
	// No terraform import.
	"aws_eip": config.IdentifierFromProvider,
	// Imported by using the EC2 Transit Gateway identifier: tgw-12345678
	"aws_ec2_transit_gateway": config.IdentifierFromProvider,
	// Imported by using the EC2 Transit Gateway Route Table, an underscore,
	// and the destination CIDR: tgw-rtb-12345678_0.0.0.0/0
	"aws_ec2_transit_gateway_route": FormattedIdentifierFromProvider("_", "transit_gateway_route_table_id", "destination_cidr_block"),
	// Imported by using the EC2 Transit Gateway Route Table identifier:
	// tgw-rtb-12345678
	"aws_ec2_transit_gateway_route_table": config.IdentifierFromProvider,
	// Imported by using the EC2 Transit Gateway Route Table identifier, an
	// underscore, and the EC2 Transit Gateway Attachment identifier, e.g.,
	// tgw-rtb-12345678_tgw-attach-87654321
	"aws_ec2_transit_gateway_route_table_association": FormattedIdentifierFromProvider("_", "transit_gateway_route_table_id", "transit_gateway_attachment_id"),
	// Imported by using the EC2 Transit Gateway Attachment identifier:
	// tgw-attach-12345678
	"aws_ec2_transit_gateway_vpc_attachment": config.IdentifierFromProvider,
	// Imported by using the EC2 Transit Gateway Attachment identifier: tgw-attach-12345678
	"aws_ec2_transit_gateway_vpc_attachment_accepter": FormattedIdentifierFromProvider("", "transit_gateway_attachment_id"),
	// Imported using the id: lt-12345678
	"aws_launch_template": config.IdentifierFromProvider,
	// Launch configurations can be imported using the name
	"aws_launch_configuration": config.NameAsIdentifier,
	// Imported using the id: vpc-23123
	"aws_vpc": config.IdentifierFromProvider,
	// Imported using the vpc endpoint id: vpce-3ecf2a57
	"aws_vpc_endpoint": config.IdentifierFromProvider,
	// Imported using the subnet id: subnet-9d4a7b6c
	"aws_subnet": config.IdentifierFromProvider,
	// Imported using the id: eni-e5aa89a3
	"aws_network_interface": config.IdentifierFromProvider,
	// Imported using the id: sg-903004f8
	"aws_security_group": config.IdentifierFromProvider,
	// Imported using a very complex format:
	// https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/security_group_rule
	"aws_security_group_rule": config.IdentifierFromProvider,
	// Imported by using the VPC CIDR Association ID: vpc-cidr-assoc-xxxxxxxx
	"aws_vpc_ipv4_cidr_block_association": config.IdentifierFromProvider,
	// Imported using the vpc peering id: pcx-111aaa111
	"aws_vpc_peering_connection": config.IdentifierFromProvider,
	// Imported using the following format: ROUTETABLEID_DESTINATION
	"aws_route": route(),
	// Imported using id: rtb-4e616f6d69
	"aws_route_table": config.IdentifierFromProvider,
	// Imported using the associated resource ID and Route Table ID separated
	// by a forward slash (/)
	"aws_route_table_association": routeTableAssociation(),
	// No import.
	"aws_main_route_table_association": config.IdentifierFromProvider,
	// No import
	"aws_ec2_transit_gateway_multicast_group_member": config.IdentifierFromProvider,
	// No import
	"aws_ec2_transit_gateway_multicast_group_source": config.IdentifierFromProvider,
	// Imported by using the EC2 Transit Gateway Route Table identifier, an
	// underscore, and the EC2 Transit Gateway Attachment identifier:
	// tgw-rtb-12345678_tgw-attach-87654321
	"aws_ec2_transit_gateway_route_table_propagation": FormattedIdentifierFromProvider("_", "transit_gateway_attachment_id", "transit_gateway_route_table_id"),
	// Imported using the id: igw-c0a643a9
	"aws_internet_gateway": config.IdentifierFromProvider,
	// NAT Gateways can be imported using the id
	"aws_nat_gateway": config.IdentifierFromProvider,
	// aws_ec2_transit_gateway_multicast_domain can be imported by using the EC2 Transit Gateway Multicast Domain identifier
	"aws_ec2_transit_gateway_multicast_domain": config.IdentifierFromProvider,
	// No import
	"aws_ec2_transit_gateway_multicast_domain_association": config.IdentifierFromProvider,
	// aws_ec2_transit_gateway_peering_attachment can be imported by using the EC2 Transit Gateway Attachment identifier
	"aws_ec2_transit_gateway_peering_attachment": config.IdentifierFromProvider,
	// Prefix List Entries can be imported using the prefix_list_id and cidr separated by a ,
	"aws_ec2_managed_prefix_list_entry": FormattedIdentifierFromProvider(",", "prefix_list_id", "cidr"),
	// Prefix Lists can be imported using the id
	"aws_ec2_managed_prefix_list": config.IdentifierFromProvider,
	// aws_ec2_transit_gateway_prefix_list_reference can be imported by using the EC2 Transit Gateway Route Table identifier and EC2 Prefix List identifier, separated by an underscore (_
	"aws_ec2_transit_gateway_prefix_list_reference": FormattedIdentifierFromProvider("_", "transit_gateway_route_table_id", "prefix_list_id"),
	// Egress-only Internet gateways can be imported using the id
	"aws_egress_only_internet_gateway": config.IdentifierFromProvider,
	// EIP Assocations can be imported using their association ID.
	"aws_eip_association": config.IdentifierFromProvider,
	// Flow Logs can be imported using the id
	"aws_flow_log": config.IdentifierFromProvider,
	// Key Pairs can be imported using the key_name
	"aws_key_pair": config.ParameterAsIdentifier("key_name"),
	// Network ACLs can be imported using the id
	"aws_network_acl": config.IdentifierFromProvider,
	// No import
	"aws_network_interface_attachment": config.IdentifierFromProvider,
	// No import
	"aws_network_interface_sg_attachment": config.IdentifierFromProvider,
	// Individual rules can be imported using NETWORK_ACL_ID:RULE_NUMBER:PROTOCOL:EGRESS
	"aws_network_acl_rule": config.IdentifierFromProvider,
	// No import
	"aws_spot_instance_request": config.IdentifierFromProvider,
	// EBS Volume Attachments can be imported using DEVICE_NAME:VOLUME_ID:INSTANCE_ID
	"aws_volume_attachment": config.IdentifierFromProvider,
	// VPC DHCP Options can be imported using the dhcp options id
	"aws_vpc_dhcp_options": config.IdentifierFromProvider,
	// DHCP associations can be imported by providing the VPC ID associated with the options
	// terraform import aws_vpc_dhcp_options_association.imported vpc-0f001273ec18911b1
	"aws_vpc_dhcp_options_association": config.IdentifierFromProvider,
	// VPC Endpoint Services can be imported using the VPC endpoint service id
	"aws_vpc_endpoint_service": config.IdentifierFromProvider,
	// VPC Endpoint connection notifications can be imported using the VPC endpoint connection notification id
	"aws_vpc_endpoint_connection_notification": config.IdentifierFromProvider,
	// VPC Endpoint Route Table Associations can be imported using vpc_endpoint_id together with route_table_id
	"aws_vpc_endpoint_route_table_association": FormattedIdentifierFromProvider("/", "vpc_endpoint_id", "route_table_id"),
	// Placement groups can be imported using the name
	"aws_placement_group": config.NameAsIdentifier,
	// A Spot Datafeed Subscription can be imported using the word spot-datafeed-subscription
	"aws_spot_datafeed_subscription": config.IdentifierFromProvider,
	// No import
	"aws_vpc_endpoint_service_allowed_principal": config.IdentifierFromProvider,
	// VPC Endpoint Subnet Associations can be imported using vpc_endpoint_id together with subnet_id
	"aws_vpc_endpoint_subnet_association": FormattedIdentifierFromProvider("/", "vpc_endpoint_id", "subnet_id"),
	// Default VPC route tables can be imported using the vpc_id
	"aws_default_route_table": config.IdentifierFromProvider,
	// Hosts can be imported using the host id
	"aws_ec2_host": config.IdentifierFromProvider,
	// Default VPCs can be imported using the vpc id
	"aws_default_vpc": config.IdentifierFromProvider,
	// Subnets can be imported using the subnet id
	"aws_default_subnet": config.IdentifierFromProvider,
	// VPC DHCP Options can be imported using the dhcp options id
	"aws_default_vpc_dhcp_options": config.IdentifierFromProvider,
	// The EBS default KMS CMK can be imported with the KMS key ARN
	"aws_ebs_default_kms_key": config.IdentifierFromProvider,
	// Default EBS encryption state can be imported
	"aws_ebs_encryption_by_default": config.IdentifierFromProvider,
	// EC2 Availability Zone Groups can be imported using the group name
	"aws_ec2_availability_zone_group": config.ParameterAsIdentifier("group_name"),
	// Capacity Reservations can be imported using the id
	"aws_ec2_capacity_reservation": config.IdentifierFromProvider,
	// aws_ec2_carrier_gateway can be imported using the carrier gateway's ID
	"aws_ec2_carrier_gateway": config.IdentifierFromProvider,
	// Serial console access state can be imported
	"aws_ec2_serial_console_access": config.IdentifierFromProvider,
	// Existing CIDR reservations can be imported using SUBNET_ID:RESERVATION_ID
	"aws_ec2_subnet_cidr_reservation": config.IdentifierFromProvider,
	// Traffic mirror filter can be imported using the id
	"aws_ec2_traffic_mirror_filter": config.IdentifierFromProvider,
	// Traffic mirror rules can be imported using the traffic_mirror_filter_id and id separated by :
	"aws_ec2_traffic_mirror_filter_rule": config.IdentifierFromProvider,
	// Traffic mirror targets can be imported using the id
	"aws_ec2_transit_gateway_connect": config.IdentifierFromProvider,
	// Network Insights Paths can be imported using the id
	"aws_ec2_network_insights_path": config.IdentifierFromProvider,
	// aws_ec2_transit_gateway_peering_attachment_accepter can be imported by using the EC2 Transit Gateway Attachment identifier
	"aws_ec2_transit_gateway_peering_attachment_accepter": config.IdentifierFromProvider,
	// No import
	"aws_snapshot_create_volume_permission": config.IdentifierFromProvider,

	// ecr
	//
	"aws_ecr_repository": config.NameAsIdentifier,
	// Imported using the name of the repository.
	"aws_ecr_lifecycle_policy": config.IdentifierFromProvider,
	// Use the ecr_repository_prefix to import a Pull Through Cache Rule.
	"aws_ecr_pull_through_cache_rule": config.IdentifierFromProvider,
	// Imported using the registry id, which is not a parameter at all.
	"aws_ecr_registry_policy": config.IdentifierFromProvider,
	// Imported using the registry id, which is not a parameter at all.
	"aws_ecr_registry_scanning_configuration": config.IdentifierFromProvider,
	// Imported using the registry id, which is not a parameter at all.
	"aws_ecr_replication_configuration": config.IdentifierFromProvider,
	// Imported using the parameter called repository but this is not the name
	// of the resource, only a configuration/reference.
	"aws_ecr_repository_policy": config.IdentifierFromProvider,

	// ecrpublic
	//
	"aws_ecrpublic_repository": config.ParameterAsIdentifier("repository_name"),
	// Imported using the repository name.
	"aws_ecrpublic_repository_policy": config.IdentifierFromProvider,

	// ecs
	//
	"aws_ecs_cluster":           config.NameAsIdentifier,
	"aws_ecs_service":           config.NameAsIdentifier,
	"aws_ecs_capacity_provider": config.TemplatedStringAsIdentifier("name", "arn:aws:ecs:{{ .setup.configuration.region }}:{{ .setup.client_metadata.account_id }}:capacity-provider/{{ .external_name }}"),
	// Imported using ARN that has a random substring, revision at the end:
	// arn:aws:ecs:us-east-1:012345678910:task-definition/mytaskfamily:123
	"aws_ecs_task_definition": config.IdentifierFromProvider,
	// ECS Account Setting defaults can be imported using the name
	"aws_ecs_account_setting_default": config.IdentifierFromProvider,
	// ECS cluster capacity providers can be imported using the cluster_name attribute
	"aws_ecs_cluster_capacity_providers": config.IdentifierFromProvider,

	// eks
	//
	"aws_eks_cluster": config.NameAsIdentifier,
	// Imported using the cluster_name and node_group_name separated by a
	// colon (:): my_cluster:my_node_group
	"aws_eks_node_group": config.TemplatedStringAsIdentifier("node_group_name", "{{ .parameters.cluster_name }}:{{ .external_name }}"),
	// my_cluster:my_eks_addon
	"aws_eks_addon": FormattedIdentifierUserDefinedNameLast("addon_name", ":", "cluster_name"),
	// my_cluster:my_fargate_profile
	"aws_eks_fargate_profile": FormattedIdentifierUserDefinedNameLast("fargate_profile_name", ":", "cluster_name"),
	// It has a complex config, adding empty entry here just to enable it.
	"aws_eks_identity_provider_config": eksOIDCIdentityProvider(),

	// elasticache
	//
	"aws_elasticache_parameter_group":   config.IdentifierFromProvider,
	"aws_elasticache_subnet_group":      config.NameAsIdentifier,
	"aws_elasticache_cluster":           config.ParameterAsIdentifier("cluster_id"),
	"aws_elasticache_replication_group": config.ParameterAsIdentifier("replication_group_id"),
	"aws_elasticache_user":              config.ParameterAsIdentifier("user_id"),
	"aws_elasticache_user_group":        config.ParameterAsIdentifier("user_group_id"),

	// elasticloadbalancing
	//
	// arn:aws:elasticloadbalancing:us-west-2:123456789012:loadbalancer/app/my-load-balancer/50dc6c495c0c9188
	"aws_lb": config.IdentifierFromProvider,
	// arn:aws:elasticloadbalancing:us-west-2:187416307283:listener/app/front-end-alb/8e4497da625e2d8a/9ab28ade35828f96
	"aws_lb_listener": config.IdentifierFromProvider,
	// arn:aws:elasticloadbalancing:us-west-2:187416307283:targetgroup/app-front-end/20cfe21448b66314
	"aws_lb_target_group": config.IdentifierFromProvider,
	// No import.
	"aws_lb_target_group_attachment": config.IdentifierFromProvider,

	// globalaccelerator
	//
	// arn:aws:globalaccelerator::111111111111:accelerator/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
	"aws_globalaccelerator_accelerator": config.IdentifierFromProvider,
	// arn:aws:globalaccelerator::111111111111:accelerator/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx/listener/xxxxxxx/endpoint-group/xxxxxxxx
	"aws_globalaccelerator_endpoint_group": config.IdentifierFromProvider,
	// arn:aws:globalaccelerator::111111111111:accelerator/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx/listener/xxxxxxxx
	"aws_globalaccelerator_listener": config.IdentifierFromProvider,

	// glue
	//
	// Imported using "name".
	"aws_glue_workflow": config.NameAsIdentifier,
	// Imported using arn: arn:aws:glue:us-west-2:123456789012:schema/example/example
	// "aws_glue_schema": config.IdentifierFromProvider,
	// Imported using "name".
	"aws_glue_trigger":               config.NameAsIdentifier,
	"aws_glue_user_defined_function": config.TemplatedStringAsIdentifier("name", "{{ .parameters.catalog_id }}:{{ .parameters.database_name }}:{{ .external_name }}"),
	// "aws_glue_security_configuration": config.NameAsIdentifier,
	// Imported using the account ID: 12356789012
	"aws_glue_resource_policy":  config.IdentifierFromProvider,
	"aws_glue_catalog_database": config.TemplatedStringAsIdentifier("name", "{{ .parameters.catalog_id }}:{{ .external_name }}"),
	"aws_glue_catalog_table":    config.TemplatedStringAsIdentifier("name", "{{ .parameters.catalog_id }}:{{ .parameters.database_name }}:{{ .external_name }}"),
	"aws_glue_classifier":       config.NameAsIdentifier,
	// Imported as CATALOG_ID:name 123456789012:MyConnection
	"aws_glue_connection": config.TemplatedStringAsIdentifier("name", "{{ .parameters.catalog_id }}:{{ .external_name }}"),
	"aws_glue_crawler":    config.NameAsIdentifier,
	// Imported using CATALOG-ID (AWS account ID if not custom), e.g., 123456789012
	"aws_glue_data_catalog_encryption_settings": config.IdentifierFromProvider,
	// "aws_glue_dev_endpoint":                     config.NameAsIdentifier,
	"aws_glue_job": config.NameAsIdentifier,
	// Imported using id, e.g., tfm-c2cafbe83b1c575f49eaca9939220e2fcd58e2d5
	// "aws_glue_ml_transform": config.IdentifierFromProvider,
	// It has no naming argument, imported with their catalog ID (usually
	// AWS account ID), database name, table name and partition values e.g.,
	// 123456789012:MyDatabase:MyTable:val1#val2
	// "aws_glue_partition": config.IdentifierFromProvider,
	// Documentation does not match schema where there are multiple indexes
	// each with their own name.
	// "aws_glue_partition_index": config.IdentifierFromProvider,
	// Imported using ARN: arn:aws:glue:us-west-2:123456789012:registry/example
	"aws_glue_registry": config.TemplatedStringAsIdentifier("registry_name", "arn:aws:glue:{{ .setup.configuration.region }}:{{ .setup.client_metadata.account_id }}:registry/{{ .external_name }}"),

	// Imported using "name".
	"aws_glue_security_configuration": config.NameAsIdentifier,

	// iam
	//
	// AKIA1234567890
	"aws_iam_access_key":       config.IdentifierFromProvider,
	"aws_iam_instance_profile": config.NameAsIdentifier,
	// arn:aws:iam::123456789012:policy/UsersManageOwnCredentials
	"aws_iam_policy": config.TemplatedStringAsIdentifier("name", "arn:aws:iam::{{ .setup.client_metadata.account_id }}:policy/{{ .external_name }}"),
	"aws_iam_user":   config.NameAsIdentifier,
	"aws_iam_group":  config.NameAsIdentifier,
	"aws_iam_role":   config.NameAsIdentifier,
	// Imported using the role name and policy arn separated by /
	// test-role/arn:aws:iam::xxxxxxxxxxxx:policy/test-policy
	"aws_iam_role_policy_attachment": config.IdentifierFromProvider,
	// Imported using the user name and policy arn separated by /
	// test-user/arn:aws:iam::xxxxxxxxxxxx:policy/test-policy
	"aws_iam_user_policy_attachment": config.IdentifierFromProvider,
	// Imported using the group name and policy arn separated by /
	// test-group/arn:aws:iam::xxxxxxxxxxxx:policy/test-policy
	"aws_iam_group_policy_attachment": config.IdentifierFromProvider,
	// Imported using the user name and group names separated by /
	// user1/group1/group2
	"aws_iam_user_group_membership": iamUserGroupMembership(),
	// arn:aws:iam::123456789012:oidc-provider/accounts.google.com
	"aws_iam_openid_connect_provider": config.IdentifierFromProvider,
	// The current Account Alias can be imported using the account_alias
	"aws_iam_account_alias": config.ParameterAsIdentifier("account_alias"),
	// IAM Account Password Policy can be imported using the word iam-account-password-policy
	"aws_iam_account_password_policy": config.IdentifierFromProvider,
	// No import
	"aws_iam_group_membership": config.IdentifierFromProvider,
	// IAM SAML Providers can be imported using the arn
	"aws_iam_saml_provider": config.TemplatedStringAsIdentifier("name", "arn:aws:iam::{{ .setup.client_metadata.account_id }}:saml-provider/{{ .external_name }}"),
	// IAM Server Certificates can be imported using the name
	"aws_iam_server_certificate": config.NameAsIdentifier,
	// IAM service-linked roles can be imported using role ARN that contains the
	// service name.
	"aws_iam_service_linked_role": config.IdentifierFromProvider,
	// IAM Service Specific Credentials can be imported using the service_name:user_name:service_specific_credential_id
	"aws_iam_service_specific_credential": config.IdentifierFromProvider,
	// IAM Signing Certificates can be imported using the id
	"aws_iam_signing_certificate": config.IdentifierFromProvider,
	// IAM User Login Profiles can be imported without password information support via the IAM User name
	"aws_iam_user_login_profile": config.IdentifierFromProvider,
	// SSH public keys can be imported using the username, ssh_public_key_id, and encoding
	"aws_iam_user_ssh_key": config.IdentifierFromProvider,
	// IAM Virtual MFA Devices can be imported using the arn
	"aws_iam_virtual_mfa_device": config.IdentifierFromProvider,

	// kms
	//
	// 1234abcd-12ab-34cd-56ef-1234567890ab
	"aws_kms_key": config.IdentifierFromProvider,
	// KMS aliases can be imported using the name
	"aws_kms_alias": config.NameAsIdentifier,
	// No import
	"aws_kms_ciphertext": config.IdentifierFromProvider,
	// KMS External Keys can be imported using the id
	"aws_kms_external_key": config.IdentifierFromProvider,
	// KMS Grants can be imported using the Key ID and Grant ID separated by a colon (:)
	"aws_kms_grant": config.IdentifierFromProvider,
	// KMS multi-Region replica keys can be imported using the id
	"aws_kms_replica_external_key": config.IdentifierFromProvider,
	// KMS multi-Region replica keys can be imported using the id
	"aws_kms_replica_key": config.IdentifierFromProvider,

	// mq
	//
	// a1b2c3d4-d5f6-7777-8888-9999aaaabbbbcccc
	"aws_mq_broker": config.IdentifierFromProvider,
	// c-0187d1eb-88c8-475a-9b79-16ef5a10c94f
	"aws_mq_configuration": config.IdentifierFromProvider,

	// neptune
	//
	"aws_neptune_cluster": config.ParameterAsIdentifier("cluster_identifier"),
	// my_cluster:my_cluster_endpoint
	"aws_neptune_cluster_endpoint":        FormattedIdentifierUserDefinedNameLast("cluster_endpoint_identifier", ":", "cluster_identifier"),
	"aws_neptune_cluster_instance":        config.ParameterAsIdentifier("identifier"),
	"aws_neptune_cluster_parameter_group": config.NameAsIdentifier,
	"aws_neptune_cluster_snapshot":        config.ParameterAsIdentifier("db_cluster_snapshot_identifier"),
	"aws_neptune_event_subscription":      config.NameAsIdentifier,
	"aws_neptune_parameter_group":         config.NameAsIdentifier,
	"aws_neptune_subnet_group":            config.NameAsIdentifier,

	// rds
	//
	"aws_rds_cluster":        config.ParameterAsIdentifier("cluster_identifier"),
	"aws_db_instance":        config.ParameterAsIdentifier("identifier"),
	"aws_db_parameter_group": config.NameAsIdentifier,
	"aws_db_subnet_group":    config.NameAsIdentifier,
	// aws_db_instance_role_association can be imported using the DB Instance Identifier and IAM Role ARN separated by a comma
	// $ terraform import aws_db_instance_role_association.example my-db-instance,arn:aws:iam::123456789012:role/my-role
	"aws_db_instance_role_association": config.IdentifierFromProvider,
	// DB Option groups can be imported using the name
	"aws_db_option_group": config.NameAsIdentifier,
	// DB proxies can be imported using the name
	"aws_db_proxy": config.NameAsIdentifier,
	// DB proxy default target groups can be imported using the db_proxy_name
	"aws_db_proxy_default_target_group": config.IdentifierFromProvider,
	// DB proxy endpoints can be imported using the DB-PROXY-NAME/DB-PROXY-ENDPOINT-NAME
	"aws_db_proxy_endpoint": config.TemplatedStringAsIdentifier("db_proxy_endpoint_name", "{{ .external_name }}/{{ .parameters.db_proxy_name }}"),
	// RDS DB Proxy Targets can be imported using the db_proxy_name, target_group_name, target type (e.g., RDS_INSTANCE or TRACKED_CLUSTER), and resource identifier separated by forward slashes (/)
	"aws_db_proxy_target": config.IdentifierFromProvider,
	// DB Security groups can be imported using the name
	"aws_db_security_group": config.NameAsIdentifier,
	// aws_db_snapshot can be imported by using the snapshot identifier
	"aws_db_snapshot": config.ParameterAsIdentifier("db_snapshot_identifier"),
	// RDS Aurora Cluster Database Activity Streams can be imported using the resource_arn
	"aws_rds_cluster_activity_stream": config.IdentifierFromProvider,
	// RDS Clusters Endpoint can be imported using the cluster_endpoint_identifier
	"aws_rds_cluster_endpoint": config.ParameterAsIdentifier("cluster_endpoint_identifier"),
	// RDS Cluster Instances can be imported using the identifier
	"aws_rds_cluster_instance": config.ParameterAsIdentifier("identifier"),
	// RDS Cluster Parameter Groups can be imported using the name
	"aws_rds_cluster_parameter_group": config.NameAsIdentifier,
	// aws_rds_cluster_role_association can be imported using the DB Cluster Identifier and IAM Role ARN separated by a comma (,)
	// $ terraform import aws_rds_cluster_role_association.example my-db-cluster,arn:aws:iam::123456789012:role/my-role
	"aws_rds_cluster_role_association": FormattedIdentifierFromProvider(",", "db_cluster_identifier", "role_arn"),
	// aws_rds_global_cluster can be imported by using the RDS Global Cluster identifie
	"aws_rds_global_cluster": config.ParameterAsIdentifier("global_cluster_identifier"),

	// route53
	//
	// N1PA6795SAMPLE
	"aws_route53_delegation_set": config.IdentifierFromProvider,
	// abcdef11-2222-3333-4444-555555fedcba
	"aws_route53_health_check": config.IdentifierFromProvider,
	// Z1D633PJN98FT9
	"aws_route53_hosted_zone_dnssec": config.IdentifierFromProvider,
	// Imported by using the Route 53 Hosted Zone identifier and KMS Key
	// identifier, separated by a comma (,), e.g., Z1D633PJN98FT9,example
	// disabled until it's successfully tested
	// "aws_route53_key_signing_key": FormattedIdentifierUserDefinedNameLast("name", ",", "hosted_zone_id"),
	// xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
	// disabled until it's successfully tested
	// "aws_route53_query_log": config.IdentifierFromProvider,
	// Imported using ID of the record, which is the zone identifier, record
	// name, and record type, separated by underscores (_)
	// Z4KAPRWWNC7JR_dev.example.com_NS
	"aws_route53_record": config.IdentifierFromProvider,
	// Z123456ABCDEFG:vpc-12345678
	"aws_route53_vpc_association_authorization": FormattedIdentifierFromProvider(":", "zone_id", "vpc_id"),
	// Z1D633PJN98FT9
	"aws_route53_zone": config.IdentifierFromProvider,
	// Z123456ABCDEFG:vpc-12345678
	// disabled until it's successfully tested
	// "aws_route53_zone_association": FormattedIdentifierFromProvider(":", "zone_id", "vpc_id"),
	// Imported using the id and version, e.g.,
	// 01a52019-d16f-422a-ae72-c306d2b6df7e/1
	"aws_route53_traffic_policy": config.IdentifierFromProvider,
	// df579d9a-6396-410e-ac22-e7ad60cf9e7e
	"aws_route53_traffic_policy_instance": config.IdentifierFromProvider,

	// route53resolver
	//
	// rdsc-be1866ecc1683e95
	// disabled until it's successfully tested
	// "aws_route53_resolver_dnssec_config": config.IdentifierFromProvider,
	// rslvr-in-abcdef01234567890
	"aws_route53_resolver_endpoint": config.IdentifierFromProvider,
	// rdsc-be1866ecc1683e95
	// disabled until it's successfully tested
	// "aws_route53_resolver_firewall_config": config.IdentifierFromProvider,
	// rslvr-fdl-0123456789abcdef
	// disabled until it's successfully tested
	// "aws_route53_resolver_firewall_domain_list": config.IdentifierFromProvider,
	// Imported using the Route 53 Resolver DNS Firewall rule group ID and
	// domain list ID separated by ':', e.g.,
	// rslvr-frg-0123456789abcdef:rslvr-fdl-0123456789abcdef
	// disabled until it's successfully tested
	// "aws_route53_resolver_firewall_rule": config.IdentifierFromProvider,
	// rslvr-frg-0123456789abcdef
	// disabled until it's successfully tested
	// "aws_route53_resolver_firewall_rule_group": config.IdentifierFromProvider,
	// rslvr-frgassoc-0123456789abcdef
	// disabled until it's successfully tested
	// "aws_route53_resolver_firewall_rule_group_association": config.IdentifierFromProvider,
	// rqlc-92edc3b1838248bf
	// disabled until it's successfully tested
	// "aws_route53_resolver_query_log_config": config.IdentifierFromProvider,
	// rqlca-b320624fef3c4d70
	// disabled until it's successfully tested
	// "aws_route53_resolver_query_log_config_association": config.IdentifierFromProvider,
	// rslvr-rr-0123456789abcdef0
	"aws_route53_resolver_rule": config.IdentifierFromProvider,
	// rslvr-rrassoc-97242eaf88example
	"aws_route53_resolver_rule_association": config.IdentifierFromProvider,

	// s3
	//
	// S3 bucket can be imported using the bucket
	"aws_s3_bucket": config.ParameterAsIdentifier("bucket"),
	// the S3 bucket accelerate configuration resource should be imported using the bucket
	"aws_s3_bucket_object_lock_configuration": config.IdentifierFromProvider,
	// the S3 bucket accelerate configuration resource should be imported using the bucket
	"aws_s3_bucket_accelerate_configuration": config.IdentifierFromProvider,
	// the S3 bucket ACL resource should be imported using the bucket
	"aws_s3_bucket_acl": config.IdentifierFromProvider,
	// S3 bucket analytics configurations can be imported using bucket:analytics
	"aws_s3_bucket_analytics_configuration": FormattedIdentifierFromProvider(":", "bucket", "name"),
	// The S3 bucket CORS configuration resource should be imported using the bucket
	"aws_s3_bucket_cors_configuration": config.IdentifierFromProvider,
	// S3 bucket intelligent tiering configurations can be imported using bucket:name
	// $ terraform import aws_s3_bucket_intelligent_tiering_configuration.my-bucket-entire-bucket my-bucket:EntireBucket
	"aws_s3_bucket_intelligent_tiering_configuration": FormattedIdentifierFromProvider(":", "bucket", "name"),
	// S3 bucket inventory configurations can be imported using bucket:inventory
	// $ terraform import aws_s3_bucket_inventory.my-bucket-entire-bucket my-bucket:EntireBucket
	"aws_s3_bucket_inventory": FormattedIdentifierFromProvider(":", "bucket", "name"),
	// The S3 bucket lifecycle configuration resource should be imported using the bucket
	"aws_s3_bucket_lifecycle_configuration": config.IdentifierFromProvider,
	// The S3 bucket logging resource should be imported using the bucket
	"aws_s3_bucket_logging": config.IdentifierFromProvider,
	// S3 bucket metric configurations can be imported using bucket:metric
	"aws_s3_bucket_metric": FormattedIdentifierFromProvider(":", "bucket", "name"),
	// S3 bucket notification can be imported using the bucket
	"aws_s3_bucket_notification": config.IdentifierFromProvider,
	// Objects can be imported using the id. The id is the bucket name and the key together
	"aws_s3_bucket_object": config.IdentifierFromProvider,
	// S3 Bucket Ownership Controls can be imported using S3 Bucket name
	"aws_s3_bucket_ownership_controls": config.IdentifierFromProvider,
	// S3 bucket policies can be imported using the bucket name
	"aws_s3_bucket_policy": config.IdentifierFromProvider,
	// aws_s3_bucket_public_access_block can be imported by using the bucket name
	"aws_s3_bucket_public_access_block": config.IdentifierFromProvider,
	// S3 bucket replication configuration can be imported using the bucket
	"aws_s3_bucket_replication_configuration": config.IdentifierFromProvider,
	// The S3 bucket request payment configuration resource should be imported using the bucket
	"aws_s3_bucket_request_payment_configuration": config.IdentifierFromProvider,
	// The S3 server-side encryption configuration resource should be imported using the bucket
	"aws_s3_bucket_server_side_encryption_configuration": config.IdentifierFromProvider,
	// The S3 bucket versioning resource should be imported using the bucket
	"aws_s3_bucket_versioning": config.IdentifierFromProvider,
	// The S3 bucket website configuration resource should be imported using the bucket
	"aws_s3_bucket_website_configuration": config.IdentifierFromProvider,
	// Objects can be imported using the id. The id is the bucket name and the key together
	// $ terraform import aws_s3_object.object some-bucket-name/some/key.txt
	"aws_s3_object": FormattedIdentifierFromProvider("/", "bucket", "key"),

	// cloudfront
	//
	// Cloudfront Cache Policies can be imported using the id
	"aws_cloudfront_cache_policy": config.IdentifierFromProvider,
	// Cloudfront Distributions can be imported using the id
	"aws_cloudfront_distribution": config.IdentifierFromProvider,
	// Cloudfront Field Level Encryption Config can be imported using the id
	"aws_cloudfront_field_level_encryption_config": config.IdentifierFromProvider,
	// Cloudfront Field Level Encryption Profile can be imported using the id
	"aws_cloudfront_field_level_encryption_profile": config.IdentifierFromProvider,
	// CloudFront Functions can be imported using the name
	"aws_cloudfront_function": config.NameAsIdentifier,
	// CloudFront Key Group can be imported using the id
	"aws_cloudfront_key_group": config.IdentifierFromProvider,
	// CloudFront monitoring subscription can be imported using the id
	"aws_cloudfront_monitoring_subscription": config.IdentifierFromProvider,
	// Cloudfront Origin Access Identities can be imported using the id
	"aws_cloudfront_origin_access_identity": config.IdentifierFromProvider,
	// No import documented, but https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/cloudfront_origin_request_policy#name
	"aws_cloudfront_origin_request_policy": config.NameAsIdentifier,
	// CloudFront Public Key can be imported using the id
	"aws_cloudfront_public_key": config.IdentifierFromProvider,
	// CloudFront real-time log configurations can be imported using the ARN,
	// $ terraform import aws_cloudfront_realtime_log_config.example arn:aws:cloudfront::111122223333:realtime-log-config/ExampleNameForRealtimeLogConfig
	"aws_cloudfront_realtime_log_config": config.IdentifierFromProvider,
	// Cloudfront Response Headers Policies can be imported using the id
	"aws_cloudfront_response_headers_policy": config.IdentifierFromProvider,

	// resource groups

	// Resource groups can be imported using the name
	"aws_resourcegroups_group": config.NameAsIdentifier,

	// docdb
	//
	// DocDB Clusters can be imported using the cluster_identifier
	"aws_docdb_cluster": config.ParameterAsIdentifier("cluster_identifier"),
	// aws_docdb_global_cluster can be imported by using the Global Cluster id
	"aws_docdb_global_cluster": config.IdentifierFromProvider,
	// DocDB Cluster Instances can be imported using the identifier
	"aws_docdb_cluster_instance": config.ParameterAsIdentifier("identifier"),
	// DocumentDB Subnet groups can be imported using the name
	"aws_docdb_subnet_group": config.NameAsIdentifier,

	// efs
	//
	// The EFS file systems can be imported using the id
	"aws_efs_file_system": config.IdentifierFromProvider,
	// The EFS mount targets can be imported using the id
	"aws_efs_mount_target": config.IdentifierFromProvider,
	// The EFS access points can be imported using the id
	"aws_efs_access_point": config.IdentifierFromProvider,
	// The EFS backup policies can be imported using the id
	"aws_efs_backup_policy": config.IdentifierFromProvider,
	// The EFS file system policies can be imported using the id
	"aws_efs_file_system_policy": config.IdentifierFromProvider,

	// servicediscovery
	//
	// Service Discovery Private DNS Namespace can be imported using the namespace ID and VPC ID: 0123456789:vpc-123345
	"aws_service_discovery_private_dns_namespace": config.IdentifierFromProvider,
	// Service Discovery Public DNS Namespace can be imported using the namespace ID
	"aws_service_discovery_public_dns_namespace": config.IdentifierFromProvider,
	// Service Discovery HTTP Namespace can be imported using the namespace ID,
	"aws_service_discovery_http_namespace": config.IdentifierFromProvider,

	// sqs
	//
	// SQS Queues can be imported using the queue url / id
	"aws_sqs_queue": config.IdentifierFromProvider,
	// SQS Queue Policies can be imported using the queue URL
	// e.g. https://queue.amazonaws.com/0123456789012/myqueue
	"aws_sqs_queue_policy": config.IdentifierFromProvider,

	// secretsmanager
	//
	// It be imported by using the secret Amazon Resource Name (ARN)
	// However, the real ID of the Secret has an Amazon-assigned random suffix,
	// i.e. if you name it with `example`, the real ID is
	// arn:aws:secretsmanager:us-west-1:609897127049:secret:example-VaznFM
	"aws_secretsmanager_secret": config.IdentifierFromProvider,
	// It uses ARN of secret and a randomly assigned ID.
	"aws_secretsmanager_secret_version": config.IdentifierFromProvider,
	// It uses its own secret_id parameter.
	"aws_secretsmanager_secret_rotation": config.IdentifierFromProvider,
	// It uses its own secert_arn parameter.
	"aws_secretsmanager_secret_policy": config.IdentifierFromProvider,

	// transfer
	//
	// Transfer Servers can be imported using the id
	"aws_transfer_server": config.IdentifierFromProvider,
	// Transfer Users can be imported using the server_id and user_name separated by /
	"aws_transfer_user": FormattedIdentifierUserDefinedNameLast("user_name", "/", "server_id"),

	// dynamodb
	//
	// DynamoDB tables can be imported using the name
	"aws_dynamodb_table": config.NameAsIdentifier,
	// DynamoDB Global Tables can be imported using the global table name
	"aws_dynamodb_global_table": config.NameAsIdentifier,

	// sns
	//
	// SNS Topics can be imported using the topic arn
	"aws_sns_topic": config.TemplatedStringAsIdentifier("name", "arn:aws:sns:{{ .setup.configuration.region }}:{{ .setup.client_metadata.account_id }}:{{ .external_name }}"),
	// SNS Topic Subscriptions can be imported using the subscription arn that
	// contains a random substring in the end.
	"aws_sns_topic_subscription": config.IdentifierFromProvider,

	// backup
	//
	// Backup Framework can be imported using the id which corresponds to the name of the Backup Framework
	"aws_backup_framework": config.IdentifierFromProvider,
	// Backup Global Settings can be imported using the id
	"aws_backup_global_settings": config.IdentifierFromProvider,
	// Backup Plan can be imported using the id
	"aws_backup_plan": config.IdentifierFromProvider,
	// Backup vault can be imported using the name
	"aws_backup_vault": config.NameAsIdentifier,
	// Backup Region Settings can be imported using the region
	"aws_backup_region_settings": config.IdentifierFromProvider,
	// Backup Report Plan can be imported using the id which corresponds to the name of the Backup Report Plan
	"aws_backup_report_plan": config.IdentifierFromProvider,
	// Backup selection can be imported using the role plan_id and id separated by | plan-id|selection-id
	"aws_backup_selection": config.IdentifierFromProvider,
	// Backup vault lock configuration can be imported using the name of the backup vault
	"aws_backup_vault_lock_configuration": config.IdentifierFromProvider,
	// Backup vault notifications can be imported using the name of the backup vault
	"aws_backup_vault_notifications": config.IdentifierFromProvider,
	// Backup vault policy can be imported using the name of the backup vault
	"aws_backup_vault_policy": config.IdentifierFromProvider,

	// grafana
	//
	// Grafana Workspace can be imported using the workspace's id
	"aws_grafana_workspace": config.IdentifierFromProvider,
	// No import
	"aws_grafana_role_association": config.IdentifierFromProvider,
	// Grafana Workspace SAML configuration can be imported using the workspace's id
	"aws_grafana_workspace_saml_configuration": FormattedIdentifierFromProvider("", "workspace_id"),

	// gamelift
	//
	// GameLift Aliases can be imported using the ID
	"aws_gamelift_alias": config.IdentifierFromProvider,
	// GameLift Builds can be imported using the ID
	"aws_gamelift_build": config.IdentifierFromProvider,
	// GameLift Fleets can be imported using the ID
	"aws_gamelift_fleet": config.IdentifierFromProvider,
	// GameLift Game Session Queues can be imported by their name
	"aws_gamelift_game_session_queue": config.NameAsIdentifier,
	// GameLift Scripts can be imported using the ID
	"aws_gamelift_script": config.IdentifierFromProvider,

	// kinesis
	//
	// Even though the documentation says the ID is name, it uses ARN..
	"aws_kinesis_stream": config.TemplatedStringAsIdentifier("name", " arn:aws:kinesis:{{ .setup.configuration.region }}:{{ .setup.client_metadata.account_id }}:stream/{{ .external_name }}"),
	// Kinesis Stream Consumers can be imported using the Amazon Resource Name (ARN)
	// that has a random substring.
	"aws_kinesis_stream_consumer": config.IdentifierFromProvider,

	// kinesisanalytics
	//
	"aws_kinesis_analytics_application": config.TemplatedStringAsIdentifier("name", "arn:aws:kinesisanalytics:{{ .setup.configuration.region }}:{{ .setup.client_metadata.account_id }}:application/{{ .external_name }}"),

	// kinesisanalyticsv2
	//
	"aws_kinesisanalyticsv2_application": config.TemplatedStringAsIdentifier("name", "arn:aws:kinesisanalytics:{{ .setup.configuration.region }}:{{ .setup.client_metadata.account_id }}:application/{{ .external_name }}"),
	// aws_kinesisanalyticsv2_application can be imported by using application_name together with snapshot_name
	// e.g. example-application/example-snapshot
	"aws_kinesisanalyticsv2_application_snapshot": FormattedIdentifierUserDefinedNameLast("snapshot_name", "/", "application_name"),

	// kinesisvideo
	//
	// Kinesis Streams can be imported using the arn that has a random substring
	// in the end.
	// arn:aws:kinesisvideo:us-west-2:123456789012:stream/terraform-kinesis-test/1554978910975
	"aws_kinesis_video_stream": config.IdentifierFromProvider,

	// firehose
	//
	"aws_kinesis_firehose_delivery_stream": config.TemplatedStringAsIdentifier("name", "arn:aws:firehose:{{ .setup.configuration.region }}:{{ .setup.client_metadata.account_id }}:deliverystream/{{ .external_name }}"),

	// lakeformation
	//
	// No import
	"aws_lakeformation_data_lake_settings": config.IdentifierFromProvider,
	// No import
	"aws_lakeformation_permissions": config.IdentifierFromProvider,
	// No import
	"aws_lakeformation_resource": config.IdentifierFromProvider,

	// lexmodels
	//
	// Bots can be imported using their name.
	"aws_lex_bot": config.NameAsIdentifier,
	// Bot aliases can be imported using an ID with the format bot_name:bot_alias_name
	"aws_lex_bot_alias": FormattedIdentifierUserDefinedNameLast("name", ":", "bot_name"),
	// Intents can be imported using their name.
	"aws_lex_intent": config.NameAsIdentifier,
	// Slot types can be imported using their name.
	"aws_lex_slot_type": config.NameAsIdentifier,

	// licensemanager
	//
	// License configurations can be imported in the form resource_arn,license_configuration_arn
	"aws_licensemanager_association": FormattedIdentifierFromProvider(",", "resource_arn", "license_configuration_arn"),
	// License configurations can be imported using the id
	"aws_licensemanager_license_configuration": config.IdentifierFromProvider,

	// lambda
	//
	// Lambda Function Aliases can be imported using the function_name/alias
	"aws_lambda_alias": config.TemplatedStringAsIdentifier("name", "{{ .parameters.function_name }}/{{ .external_name }}"),
	// Code Signing Configs can be imported using their ARN that has a random
	// substring in the end.
	// arn:aws:lambda:us-west-2:123456789012:code-signing-config:csc-0f6c334abcdea4d8b
	"aws_lambda_code_signing_config": config.IdentifierFromProvider,
	// Lambda event source mappings can be imported using the UUID (event source mapping identifier)
	"aws_lambda_event_source_mapping": config.IdentifierFromProvider,
	// Lambda Functions can be imported using the function_name
	"aws_lambda_function": config.ParameterAsIdentifier("function_name"),
	// Lambda Function Event Invoke Configs can be imported using the
	// fully qualified Function name or Amazon Resource Name (ARN) of the function.
	"aws_lambda_function_event_invoke_config": config.IdentifierFromProvider,
	// Lambda function URLs can be imported using the function_name or function_name/qualifier
	"aws_lambda_function_url": lambdaFunctionURL(),
	// No import"
	"aws_lambda_invocation": config.IdentifierFromProvider,
	// Lambda Layers can be imported using arn that has an assigned version in the
	// end
	"aws_lambda_layer_version": config.IdentifierFromProvider,
	// Lambda Layer Permissions can be imported using layer_name and version_number, separated by a comma (,)
	"aws_lambda_layer_version_permission": config.IdentifierFromProvider,
	// Lambda permission statements can be imported using function_name/statement_id, with an optional qualifier
	"aws_lambda_permission": config.IdentifierFromProvider,
	// Lambda Provisioned Concurrency Configs can be imported using the function_name and qualifier separated by a colon (:)
	"aws_lambda_provisioned_concurrency_config": config.IdentifierFromProvider,

	// signer
	//
	// Signer signing profiles can be imported using the name
	"aws_signer_signing_profile": config.NameAsIdentifier,

	// athena
	//
	// Athena Workgroups can be imported using their name
	"aws_athena_workgroup": config.NameAsIdentifier,
	// Data catalogs can be imported using their name
	"aws_athena_data_catalog": config.NameAsIdentifier,
	// Athena Databases can be imported using their name
	"aws_athena_database": config.NameAsIdentifier,
	// Athena Named Query can be imported using the query ID
	"aws_athena_named_query": config.IdentifierFromProvider,

	// cloudwatchlogs
	//
	// Cloudwatch Log Groups can be imported using the name
	"aws_cloudwatch_log_group": config.NameAsIdentifier,
	// CloudWatch Log Metric Filter can be imported using the log_group_name:name
	"aws_cloudwatch_log_metric_filter": config.TemplatedStringAsIdentifier("name", "{{ .parameters.log_group_name }}:{{ .external_name }}"),
	// CloudWatch query definitions can be imported using the query definition ARN.
	"aws_cloudwatch_query_definition": config.IdentifierFromProvider,
	// Cloudwatch Log Stream can be imported using the stream's log_group_name and name
	"aws_cloudwatch_log_stream": config.IdentifierFromProvider,
	// CloudWatch log resource policies can be imported using the policy name
	"aws_cloudwatch_log_resource_policy": config.ParameterAsIdentifier("policy_name"),

	// elb
	//
	// ELBs can be imported using the name
	"aws_elb": config.NameAsIdentifier,
	// No import
	"aws_elb_attachment": config.IdentifierFromProvider,

	// iot
	//
	// IoT policies can be imported using the name
	"aws_iot_policy": config.NameAsIdentifier,
	// IOT Things can be imported using the name
	"aws_iot_thing": config.NameAsIdentifier,

	// kafka
	//
	// MSK configurations can be imported using the configuration ARN that has
	// a random substring in the end.
	"aws_msk_configuration": config.IdentifierFromProvider,
	// MSK clusters can be imported using the cluster arn that has a random substring
	// in the end.
	"aws_msk_cluster": config.IdentifierFromProvider,

	// ram
	//
	// Resource shares can be imported using the id
	"aws_ram_resource_share": config.IdentifierFromProvider,

	// redshift
	//
	// Redshift Clusters can be imported using the cluster_identifier
	"aws_redshift_cluster": config.ParameterAsIdentifier("cluster_identifier"),

	// sfn
	//
	"aws_sfn_activity":      config.TemplatedStringAsIdentifier("name", "arn:aws:states:{{ .setup.configuration.region }}:{{ .setup.client_metadata.account_id }}:activity/{{ .external_name }}"),
	"aws_sfn_state_machine": config.TemplatedStringAsIdentifier("name", "arn:aws:states:{{ .setup.configuration.region }}:{{ .setup.client_metadata.account_id }}:stateMachine/{{ .external_name }}"),

	// dax
	//
	// DAX Clusters can be imported using the cluster_name
	"aws_dax_cluster": config.ParameterAsIdentifier("cluster_name"),
	// DAX Parameter Group can be imported using the name
	"aws_dax_parameter_group": config.NameAsIdentifier,
	// DAX Subnet Group can be imported using the name
	"aws_dax_subnet_group": config.NameAsIdentifier,

	// cloudsearch
	//
	// CloudSearch Domains can be imported using the name
	"aws_cloudsearch_domain": config.NameAsIdentifier,
	// CloudSearch domain service access policies can be imported using the domain name
	"aws_cloudsearch_domain_service_access_policy": config.IdentifierFromProvider,

	// apigateway
	//
	// API Gateway Keys can be imported using the id
	"aws_api_gateway_api_key": config.IdentifierFromProvider,
	// API Gateway Client Certificates can be imported using the id
	"aws_api_gateway_client_certificate": config.IdentifierFromProvider,
	// aws_api_gateway_rest_api can be imported by using the REST API ID
	"aws_api_gateway_rest_api": config.IdentifierFromProvider,
	// API Gateway documentation_parts can be imported using REST-API-ID/DOC-PART-ID
	"aws_api_gateway_documentation_part": config.IdentifierFromProvider,
	// API Gateway documentation versions can be imported using REST-API-ID/VERSION
	"aws_api_gateway_documentation_version": FormattedIdentifierFromProvider("/", "rest_api_id", "version"),
	// aws_api_gateway_gateway_response can be imported using REST-API-ID/RESPONSE-TYPE
	"aws_api_gateway_gateway_response": FormattedIdentifierFromProvider("/", "rest_api_id", "response_type"),
	// aws_api_gateway_resource can be imported using REST-API-ID/RESOURCE-ID
	"aws_api_gateway_resource": config.IdentifierFromProvider,
	// aws_api_gateway_method can be imported using REST-API-ID/RESOURCE-ID/HTTP-METHOD
	"aws_api_gateway_method": FormattedIdentifierFromProvider("/", "rest_api_id", "resource_id", "http_method"),
	// aws_api_gateway_method_response can be imported using REST-API-ID/RESOURCE-ID/HTTP-METHOD/STATUS-CODE
	"aws_api_gateway_method_response": FormattedIdentifierFromProvider("/", "rest_api_id", "resource_id", "http_method", "status_code"),
	// No import
	"aws_api_gateway_deployment": config.IdentifierFromProvider,
	// API Gateway Accounts can be imported using the word api-gateway-account
	"aws_api_gateway_account": config.IdentifierFromProvider,
	// aws_api_gateway_stage can be imported using REST-API-ID/STAGE-NAME
	"aws_api_gateway_stage": FormattedIdentifierFromProvider("/", "rest_api_id", "stage_name"),
	// aws_api_gateway_integration can be imported using REST-API-ID/RESOURCE-ID/HTTP-METHOD
	"aws_api_gateway_integration": FormattedIdentifierFromProvider("/", "rest_api_id", "resource_id", "http_method"),
	// aws_api_gateway_integration_response can be imported using REST-API-ID/RESOURCE-ID/HTTP-METHOD/STATUS-CODE
	"aws_api_gateway_integration_response": FormattedIdentifierFromProvider("/", "rest_api_id", "resource_id", "http_method", "status_code"),
	// aws_api_gateway_method_settings can be imported using REST-API-ID/STAGE-NAME/METHOD-PATH
	"aws_api_gateway_method_settings": FormattedIdentifierFromProvider("/", "rest_api_id", "stage_name", "method_path"),
	// aws_api_gateway_model can be imported using REST-API-ID/NAME
	"aws_api_gateway_model": config.IdentifierFromProvider,
	// aws_api_gateway_request_validator can be imported using REST-API-ID/REQUEST-VALIDATOR-ID
	"aws_api_gateway_request_validator": config.IdentifierFromProvider,
	// aws_api_gateway_rest_api_policy can be imported by using the REST API ID
	"aws_api_gateway_rest_api_policy": FormattedIdentifierFromProvider("", "rest_api_id"),
	// AWS API Gateway Authorizer can be imported using the REST-API-ID/AUTHORIZER-ID
	"aws_api_gateway_authorizer": config.IdentifierFromProvider,
	// aws_api_gateway_base_path_mapping can be imported by using the domain name and base path.
	// For empty base_path (e.g., root path (/)): example.com/
	// Otherwise: example.com/base-path
	"aws_api_gateway_base_path_mapping": config.IdentifierFromProvider,
	// API Gateway domain names can be imported using their name
	"aws_api_gateway_domain_name": config.IdentifierFromProvider,
	// AWS API Gateway Usage Plan can be imported using the id
	"aws_api_gateway_usage_plan": config.IdentifierFromProvider,
	// AWS API Gateway Usage Plan Key can be imported using the USAGE-PLAN-ID/USAGE-PLAN-KEY-ID
	"aws_api_gateway_usage_plan_key": config.IdentifierFromProvider,
	// API Gateway VPC Link can be imported using the id
	"aws_api_gateway_vpc_link": config.IdentifierFromProvider,

	// opensearch
	//
	// NOTE(sergen): Parameter as identifier cannot be used, because terraform
	// overrides the id after terraform calls.
	// Please see the following issue in upjet: https://github.com/upbound/upjet/issues/32
	// OpenSearch domains can be imported using the domain_name
	"aws_opensearch_domain": config.IdentifierFromProvider,
	// No imports
	"aws_opensearch_domain_policy": config.IdentifierFromProvider,
	// NOTE(sergen): Parameter as identifier cannot be used, because terraform
	// overrides the id after terraform calls.
	// Please see the following issue in upjet: https://github.com/upbound/upjet/issues/32
	// OpenSearch domains can be imported using the domain_name
	"aws_opensearch_domain_saml_options": config.IdentifierFromProvider,

	// cloudwatch
	//
	// Use the alarm_name to import a CloudWatch Composite Alarm.
	"aws_cloudwatch_composite_alarm": config.ParameterAsIdentifier("alarm_name"),
	// CloudWatch dashboards can be imported using the dashboard_name
	"aws_cloudwatch_dashboard": config.ParameterAsIdentifier("dashboard_name"),
	// CloudWatch Metric Alarm can be imported using the alarm_name
	"aws_cloudwatch_metric_alarm": config.ParameterAsIdentifier("alarm_name"),
	// CloudWatch metric streams can be imported using the name
	"aws_cloudwatch_metric_stream": config.IdentifierFromProvider,

	// appautoscaling
	//
	// No import
	"aws_appautoscaling_scheduled_action": config.IdentifierFromProvider,
	// Application AutoScaling Policy can be imported using the service-namespace, resource-id, scalable-dimension and policy-name separated by /
	"aws_appautoscaling_policy": config.TemplatedStringAsIdentifier("name", "{{ .parameters.service_namespace }}/{{ .parameters.resource_id }}/{{ .parameters.scalable_dimension }}/{{ .external_name }}"),
	// Application AutoScaling Target can be imported using the service-namespace , resource-id and scalable-dimension separated by /
	"aws_appautoscaling_target": TemplatedStringAsIdentifierWithNoName("{{ .parameters.service_namespace }}/{{ .parameters.resource_id }}/{{ .parameters.scalable_dimension }}"),

	// codecommit
	//
	// Codecommit repository can be imported using repository name
	"aws_codecommit_repository": config.ParameterAsIdentifier("repository_name"),
	// CodeCommit approval rule templates can be imported using the name
	"aws_codecommit_approval_rule_template": config.NameAsIdentifier,
	// CodeCommit approval rule template associations can be imported using the approval_rule_template_name and repository_name separated by a comma (,)
	"aws_codecommit_approval_rule_template_association": config.TemplatedStringAsIdentifier("", "{{ .parameters.approval_rule_template_name }},{{ .parameters.repository_name }}"),
	// No import
	"aws_codecommit_trigger": config.IdentifierFromProvider,

	// deploy
	//
	// CodeDeploy Applications can be imported using the name
	"aws_codedeploy_app": config.TemplatedStringAsIdentifier("name", "{{ .parameters.application_id }}:{{ .external_name }}"),
	// CodeDeploy Deployment Configurations can be imported using the deployment_config_name
	"aws_codedeploy_deployment_config": config.ParameterAsIdentifier("deployment_config_name"),
	// CodeDeploy Deployment Groups can be imported by their app_name, a colon, and deployment_group_name
	"aws_codedeploy_deployment_group": config.TemplatedStringAsIdentifier("deployment_group_name", "{{ .parameters.app_name }}:{{ .external_name }}"),

	// codepipeline
	//
	// CodePipelines can be imported using the name
	"aws_codepipeline": config.NameAsIdentifier,
	// CodePipeline Webhooks can be imported by their ARN: arn:aws:codepipeline:us-west-2:123456789012:webhook:example
	"aws_codepipeline_webhook": config.TemplatedStringAsIdentifier("name", "arn:aws:codepipeline:{{ .setup.configuration.region }}:{{ .setup.client_metadata.account_id }}:webhook:{{ .external_name }}"),

	// codestarconnections
	//
	// CodeStar connections can be imported using the ARN
	"aws_codestarconnections_connection": config.IdentifierFromProvider,
	// CodeStar Host can be imported using the ARN
	"aws_codestarconnections_host": config.IdentifierFromProvider,

	// codestarnotifications
	//
	// CodeStar notification rule can be imported using the ARN
	"aws_codestarnotifications_notification_rule": config.IdentifierFromProvider,

	// connect
	//
	// aws_connect_bot_association can be imported by using the Amazon Connect instance ID, Lex (V1) bot name, and Lex (V1) bot region separated by colons (:)
	// TODO: lex_bot.lex_region parameter is not `Required` in TF schema. But we use this field in id construction. So, please mark as required this field while configuration
	"aws_connect_bot_association": config.TemplatedStringAsIdentifier("", "{{ .parameters.instance_id }}:{{ (index .parameters.lex_bot 0).name }}:{{ (index .parameters.lex_bot 0).lex_region }}"),
	// Amazon Connect Contact Flows can be imported using the instance_id and contact_flow_id separated by a colon (:)
	"aws_connect_contact_flow": config.IdentifierFromProvider,
	// Amazon Connect Contact Flow Modules can be imported using the instance_id and contact_flow_module_id separated by a colon (:)
	"aws_connect_contact_flow_module": config.IdentifierFromProvider,
	// Amazon Connect Hours of Operations can be imported using the instance_id and hours_of_operation_id separated by a colon (:)
	"aws_connect_hours_of_operation": config.IdentifierFromProvider,
	// Connect instances can be imported using the id
	"aws_connect_instance": config.IdentifierFromProvider,
	// aws_connect_lambda_function_association can be imported using the instance_id and function_arn separated by a comma (,)
	"aws_connect_lambda_function_association": config.TemplatedStringAsIdentifier("", "{{ .parameters.instance_id }},{{ .parameters.function_arn }}"),
	// Amazon Connect Queues can be imported using the instance_id and queue_id separated by a colon (:)
	"aws_connect_queue": config.IdentifierFromProvider,
	// Amazon Connect Quick Connects can be imported using the instance_id and quick_connect_id separated by a colon (:)
	"aws_connect_quick_connect": config.IdentifierFromProvider,
	// Amazon Connect Routing Profiles can be imported using the instance_id and routing_profile_id separated by a colon (:)
	"aws_connect_routing_profile": config.IdentifierFromProvider,
	// Amazon Connect Security Profiles can be imported using the instance_id and security_profile_id separated by a colon (:)
	"aws_connect_security_profile": config.IdentifierFromProvider,
	// Amazon Connect User Hierarchy Structures can be imported using the instance_id
	"aws_connect_user_hierarchy_structure": config.IdentifierFromProvider,

	// apprunner
	//
	// App Runner AutoScaling Configuration Versions can be imported by using the arn
	"aws_apprunner_auto_scaling_configuration_version": config.IdentifierFromProvider,
	// App Runner Connections can be imported by using the connection_name
	"aws_apprunner_connection": config.ParameterAsIdentifier("connection_name"),
	// App Runner Services can be imported by using the arn
	"aws_apprunner_service": config.IdentifierFromProvider,
	// App Runner vpc connector can be imported by using the arn
	"aws_apprunner_vpc_connector": config.IdentifierFromProvider,

	// appstream
	//
	// aws_appstream_directory_config can be imported using the id
	"aws_appstream_directory_config": config.IdentifierFromProvider,
	// aws_appstream_fleet can be imported using the id
	"aws_appstream_fleet": config.IdentifierFromProvider,
	// AppStream Stack Fleet Association can be imported by using the fleet_name and stack_name separated by a slash (/)
	"aws_appstream_fleet_stack_association": config.TemplatedStringAsIdentifier("", "{{ .parameters.fleet_name }}/{{ .parameters.stack_name}}"),
	// aws_appstream_image_builder can be imported using the name
	"aws_appstream_image_builder": config.NameAsIdentifier,
	// aws_appstream_stack can be imported using the id
	"aws_appstream_stack": config.IdentifierFromProvider,
	// aws_appstream_user can be imported using the user_name and authentication_type separated by a slash (/)
	"aws_appstream_user": config.TemplatedStringAsIdentifier("user_name", "{{ .external_name }}/{{ .parameters.authentication_type }}"),
	// AppStream User Stack Association can be imported by using the user_name, authentication_type, and stack_name, separated by a slash (/)
	"aws_appstream_user_stack_association": config.TemplatedStringAsIdentifier("", "{{ .parameters.user_name }}/{{ .parameters.authentication_type }}/{{ .parameters.stack_name }}/"),

	// appmesh
	//
	// App Mesh service meshes can be imported using the name
	"aws_appmesh_mesh": config.NameAsIdentifier,
	// App Mesh virtual nodes can be imported using mesh_name together with the virtual node's name: simpleapp/serviceBv1
	"aws_appmesh_virtual_node": config.IdentifierFromProvider,
	// App Mesh virtual routers can be imported using mesh_name together with the virtual router's name: simpleapp/serviceB
	"aws_appmesh_virtual_router": config.IdentifierFromProvider,
	// App Mesh virtual gateway can be imported using mesh_name together with the virtual gateway's name: mesh/gw1
	"aws_appmesh_virtual_gateway": config.IdentifierFromProvider,
	// App Mesh virtual services can be imported using mesh_name together with the virtual service's name: simpleapp/servicea.simpleapp.local
	"aws_appmesh_virtual_service": config.IdentifierFromProvider,
	// mesh/gw1/example-gateway-route
	"aws_appmesh_gateway_route": config.IdentifierFromProvider,
	// App Mesh virtual routes can be imported using mesh_name and virtual_router_name together with the route's name, e.g.,
	// simpleapp/serviceB/serviceB-route
	"aws_appmesh_route": config.IdentifierFromProvider,

	// configservice
	//
	// Config Rule can be imported using the name
	"aws_config_config_rule": config.NameAsIdentifier,
	// Configuration Aggregators can be imported using the name
	"aws_config_configuration_aggregator": config.NameAsIdentifier,
	// Configuration Recorder can be imported using the name
	"aws_config_configuration_recorder": config.NameAsIdentifier,
	// Configuration Recorder Status can be imported using the name of the Configuration Recorder
	"aws_config_configuration_recorder_status": config.NameAsIdentifier,
	// Config Conformance Packs can be imported using the name
	"aws_config_conformance_pack": config.NameAsIdentifier,
	// Delivery Channel can be imported using the name
	"aws_config_delivery_channel": config.NameAsIdentifier,
	// Remediation Configurations can be imported using the name config_rule_name
	"aws_config_remediation_configuration": config.ParameterAsIdentifier("config_rule_name"),

	// appsync
	//
	// aws_appsync_api_cache can be imported using the AppSync API ID
	"aws_appsync_api_cache": config.IdentifierFromProvider,
	// aws_appsync_api_key can be imported using the AppSync API ID and key separated by :
	"aws_appsync_api_key": TemplatedStringAsIdentifierWithNoName("{{ .parameters.api_id }}:{{ .external_name }}"),
	// aws_appsync_datasource can be imported with their api_id, a hyphen, and name
	"aws_appsync_datasource": config.TemplatedStringAsIdentifier("name", "{{ .parameters.api_id }}-{{ .external_name }}"),
	// aws_appsync_function can be imported using the AppSync API ID and Function ID separated by -
	"aws_appsync_function": config.IdentifierFromProvider,
	// AppSync GraphQL API can be imported using the GraphQL API ID
	"aws_appsync_graphql_api": config.IdentifierFromProvider,
	// aws_appsync_resolver can be imported with their api_id, a hyphen, type, a hypen and field
	"aws_appsync_resolver": config.TemplatedStringAsIdentifier("", "{{ .parameters.api_id }}-{{ .parameters.type }}-{{ .parameters.field }}"),

	// accessanalyzer
	//
	// Access Analyzer Analyzers can be imported using the analyzer_name
	"aws_accessanalyzer_analyzer": config.ParameterAsIdentifier("analyzer_name"),

	// account
	//
	// The Alternate Contact for the current account can be imported using the alternate_contact_type
	"aws_account_alternate_contact": config.TemplatedStringAsIdentifier("", "{{ .parameters.alternate_contact_type }}"),

	// amplify
	//
	// Amplify App can be imported using Amplify App ID (appId)
	"aws_amplify_app": config.IdentifierFromProvider,
	// Amplify branch can be imported using app_id and branch_name: d2ypk4k47z8u6/master
	"aws_amplify_branch": config.TemplatedStringAsIdentifier("branch_name", "{{ .parameters.app_id }}/{{ .external_name }}"),
	// Amplify backend environment can be imported using app_id and environment_name: d2ypk4k47z8u6/example
	"aws_amplify_backend_environment": config.TemplatedStringAsIdentifier("environment_name", "{{ .parameters.app_id }}/{{ .external_name }}"),
	// Amplify webhook can be imported using a webhook ID
	"aws_amplify_webhook": config.IdentifierFromProvider,

	// cur
	//
	// Report Definitions can be imported using the report_name
	"aws_cur_report_definition": config.ParameterAsIdentifier("report_name"),

	// dataexchange
	//
	// DataExchange DataSets can be imported by their arn
	"aws_dataexchange_data_set": config.IdentifierFromProvider,
	// DataExchange Revisions can be imported by their data-set-id:revision-id
	"aws_dataexchange_revision": config.IdentifierFromProvider,

	// datapipeline
	//
	// aws_datapipeline_pipeline can be imported by using the id (Pipeline ID)
	"aws_datapipeline_pipeline": config.IdentifierFromProvider,

	// detective
	//
	// aws_detective_graph can be imported using the ARN
	"aws_detective_graph": config.IdentifierFromProvider,
	// aws_detective_member can be imported using the ARN of the graph followed by the account ID of the member account
	"aws_detective_member": config.IdentifierFromProvider,
	// aws_detective_invitation_accepter can be imported using the graph ARN
	"aws_detective_invitation_accepter": config.IdentifierFromProvider,

	// devicefarm
	//
	// DeviceFarm Projects can be imported by their arn
	"aws_devicefarm_project": config.IdentifierFromProvider,
	// DeviceFarm Instance Profiles can be imported by their arn
	"aws_devicefarm_instance_profile": config.IdentifierFromProvider,
	// DeviceFarm Device Pools can be imported by their arn
	"aws_devicefarm_device_pool": config.IdentifierFromProvider,
	// DeviceFarm Network Profiles can be imported by their arn
	"aws_devicefarm_network_profile": config.IdentifierFromProvider,
	// DeviceFarm Uploads can be imported by their arn
	"aws_devicefarm_upload": config.IdentifierFromProvider,
	// DeviceFarm Test Grid Projects can be imported by their arn
	"aws_devicefarm_test_grid_project": config.IdentifierFromProvider,

	// organization
	//
	// imported by using the account id, which is provider-generated
	"aws_organizations_account": config.IdentifierFromProvider,
	// imported by using the account ID and its service principal:
	// 123456789012/config.amazonaws.com
	"aws_organizations_delegated_administrator": FormattedIdentifierFromProvider("/", "account_id", "service_principal"),
	//  imported by using the id, which is a Cloud provider-generated string:
	// o-1234567
	"aws_organizations_organization": config.IdentifierFromProvider,
	// imported by using the id, which is a Cloud provider-generated string:
	// ou-1234567
	"aws_organizations_organizational_unit": config.IdentifierFromProvider,
	// imported by using the policy ID,
	// which is a Cloud provider-generated string:
	// p-12345678
	"aws_organizations_policy": config.IdentifierFromProvider,
	// imported by using the target ID and policy ID
	// 123456789012:p-12345678
	"aws_organizations_policy_attachment": FormattedIdentifierFromProvider(":", "target_id", "policy_id"),

	// batch
	//
	// Batch Scheduling Policy can be imported using the arn: arn:aws:batch:us-east-1:123456789012:scheduling-policy/sample
	"aws_batch_scheduling_policy": config.TemplatedStringAsIdentifier("name", "arn:aws:batch:{{ .setup.configuration.region }}:{{ .setup.client_metadata.account_id }}:scheduling-policy/{{ .external_name }}"),

	// budgets
	//
	// Budgets can be imported using AccountID:BudgetName
	"aws_budgets_budget": config.TemplatedStringAsIdentifier("name", "{{ .setup.client_metadata.account_id }}:{{ .external_name }}"),
	// Budgets can be imported using AccountID:ActionID:BudgetName
	"aws_budgets_budget_action": config.IdentifierFromProvider,

	// chime
	//
	// Configuration Recorder can be imported using the name
	"aws_chime_voice_connector": config.NameAsIdentifier,
	// Configuration Recorder can be imported using the name
	"aws_chime_voice_connector_group": config.NameAsIdentifier,
	// Chime Voice Connector Logging can be imported using the voice_connector_id
	"aws_chime_voice_connector_logging": config.IdentifierFromProvider,
	// Chime Voice Connector Origination can be imported using the voice_connector_id
	"aws_chime_voice_connector_origination": config.IdentifierFromProvider,
	// Chime Voice Connector Streaming can be imported using the voice_connector_id
	"aws_chime_voice_connector_streaming": config.IdentifierFromProvider,
	// Chime Voice Connector Termination can be imported using the voice_connector_id
	"aws_chime_voice_connector_termination": config.IdentifierFromProvider,
	// Chime Voice Connector Termination Credentials can be imported using the voice_connector_id
	"aws_chime_voice_connector_termination_credentials": config.IdentifierFromProvider,

	// cloud9
	//
	// No import
	"aws_cloud9_environment_ec2": config.IdentifierFromProvider,
	// Cloud9 environment membership can be imported using the environment-id#user-arn
	"aws_cloud9_environment_membership": config.TemplatedStringAsIdentifier("", "{{ .parameters.environment_id }}#{{ .parameters.user_arn }}"),

	// cloudcontrol
	//
	// No import
	"aws_cloudcontrolapi_resource": config.IdentifierFromProvider,
}

func lambdaFunctionURL() config.ExternalName {
	e := config.IdentifierFromProvider
	e.GetIDFn = func(ctx context.Context, externalName string, parameters map[string]interface{}, terraformProviderConfig map[string]interface{}) (string, error) {
		functionName, ok := parameters["function_name"]
		if !ok {
			return "", errors.New("function_name cannot be empty")
		}

		qualifier := parameters["qualifier"]
		if qualifier == nil || qualifier == "" {
			return functionName.(string), nil
		}
		return fmt.Sprintf("%s/%s", functionName.(string), qualifier.(string)), nil
	}
	return e
}

func iamUserGroupMembership() config.ExternalName {
	e := config.IdentifierFromProvider
	e.GetIDFn = func(_ context.Context, _ string, parameters map[string]interface{}, _ map[string]interface{}) (string, error) {
		u, ok := parameters["user"]
		if !ok {
			return "", errors.New("user cannot be empty")
		}
		gs, ok := parameters["groups"]
		if !ok {
			return "", errors.New("groups cannot be empty")
		}
		var groups []string
		for _, g := range gs.([]interface{}) {
			groups = append(groups, g.(string))
		}
		return strings.Join(append([]string{u.(string)}, groups...), "/"), nil
	}
	return e
}

func route() config.ExternalName {
	e := config.IdentifierFromProvider
	e.GetIDFn = func(_ context.Context, _ string, parameters map[string]interface{}, _ map[string]interface{}) (string, error) {
		rtb, ok := parameters["route_table_id"]
		if !ok {
			return "", errors.New("route_table_id cannot be empty")
		}
		switch {
		case parameters["destination_cidr_block"] != nil:
			return fmt.Sprintf("%s_%s", rtb.(string), parameters["destination_cidr_block"].(string)), nil
		case parameters["destination_ipv6_cidr_block"] != nil:
			return fmt.Sprintf("%s_%s", rtb.(string), parameters["destination_ipv6_cidr_block"].(string)), nil
		case parameters["destination_prefix_list_id"] != nil:
			return fmt.Sprintf("%s_%s", rtb.(string), parameters["destination_prefix_list_id"].(string)), nil
		}
		return "", errors.New("destination_cidr_block or destination_ipv6_cidr_block or destination_prefix_list_id has to be given")
	}
	return e
}

func routeTableAssociation() config.ExternalName {
	e := config.IdentifierFromProvider
	e.GetIDFn = func(_ context.Context, _ string, parameters map[string]interface{}, _ map[string]interface{}) (string, error) {
		rtb, ok := parameters["route_table_id"]
		if !ok {
			return "", errors.New("route_table_id cannot be empty")
		}
		switch {
		case parameters["subnet_id"] != nil:
			return fmt.Sprintf("%s/%s", parameters["subnet_id"].(string), rtb.(string)), nil
		case parameters["gateway_id"] != nil:
			return fmt.Sprintf("%s/%s", parameters["gateway_id"].(string), rtb.(string)), nil
		}
		return "", errors.New("gateway_id or subnet_id has to be given")
	}
	return e
}

func eksOIDCIdentityProvider() config.ExternalName {
	// OmittedFields in config.ExternalName works only for the top-level fields.
	// Hence, omitting is done in individual config override in `eks/config.go`
	return config.ExternalName{
		SetIdentifierArgumentFn: func(base map[string]interface{}, externalName string) {
			if _, ok := base["oidc"]; !ok {
				base["oidc"] = map[string]interface{}{}
			}
			// max length is 1:
			// https://github.com/hashicorp/terraform-provider-aws/blob/7ff39c5b11aafe812e3a4b414aa6d345286b95ec/internal/service/eks/identity_provider_config.go#L58
			if arr, ok := base["oidc"].([]interface{}); ok && len(arr) == 1 {
				if m, ok := arr[0].(map[string]interface{}); ok {
					m["identity_provider_config_name"] = externalName
				}
			}
		},
		GetExternalNameFn: func(tfstate map[string]interface{}) (string, error) {
			if id, ok := tfstate["id"]; ok {
				return strings.Split(id.(string), ":")[1], nil
			}
			return "", errors.New("there is no id in tfstate")
		},
		GetIDFn: func(_ context.Context, externalName string, parameters map[string]interface{}, _ map[string]interface{}) (string, error) {
			cl, ok := parameters["cluster_name"]
			if !ok {
				return "", errors.New("cluster_name cannot be empty")
			}
			return fmt.Sprintf("%s:%s", cl.(string), externalName), nil
		},
	}
}

// FormattedIdentifierFromProvider is a helper function to construct Terraform
// IDs that use elements from the parameters in a certain string format.
// It should be used in cases where all information in the ID is gathered from
// the spec and not user defined like name. For example, zone_id:vpc_id.
func FormattedIdentifierFromProvider(separator string, keys ...string) config.ExternalName {
	e := config.IdentifierFromProvider
	e.GetIDFn = func(_ context.Context, _ string, parameters map[string]interface{}, _ map[string]interface{}) (string, error) {
		vals := make([]string, len(keys))
		for i, key := range keys {
			val, ok := parameters[key]
			if !ok {
				return "", errors.Errorf("%s cannot be empty", key)
			}
			s, ok := val.(string)
			if !ok {
				return "", errors.Errorf("%s needs to be string", key)
			}
			vals[i] = s
		}
		return strings.Join(vals, separator), nil
	}
	return e
}

// FormattedIdentifierUserDefinedNameLast is used in cases where the ID is constructed
// using some of the spec fields as well as a field that users use to name the
// resource. For example, vpc_id:cluster_name where vpc_id comes from spec
// but cluster_name is a naming field we can use external name for.
// This function assumes that the naming field is the LAST component
// in the constructed identifier, which may not always hold
// (e.g., aws_servicecatalog_budget_resource_association).
func FormattedIdentifierUserDefinedNameLast(param, separator string, keys ...string) config.ExternalName {
	e := config.ParameterAsIdentifier(param)
	e.GetIDFn = func(_ context.Context, externalName string, parameters map[string]interface{}, _ map[string]interface{}) (string, error) {
		vals := make([]string, len(keys)+1)
		for i, k := range keys {
			v, ok := parameters[k]
			if !ok {
				return "", errors.Errorf("%s cannot be empty", k)
			}
			s, ok := v.(string)
			if !ok {
				return "", errors.Errorf("%s needs to be a string", k)
			}
			vals[i] = s
		}
		vals[len(vals)-1] = externalName
		return strings.Join(vals, separator), nil
	}
	e.GetExternalNameFn = func(tfstate map[string]interface{}) (string, error) {
		id, ok := tfstate["id"]
		if !ok {
			return "", errors.New("id in tfstate cannot be empty")
		}
		s, ok := id.(string)
		if !ok {
			return "", errors.New("value of id needs to be string")
		}
		w := strings.Split(s, separator)
		return w[len(w)-1], nil
	}
	return e
}

// FormattedIdentifierUserDefinedNameFirst is used in cases where the ID is constructed
// using some of the spec fields as well as a field that users use to name the
// resource. For example, budget_name:product_id where product_id comes from spec
// but budget_name is a naming field we can use external name for.
// This function assumes that the naming field is the FIRST component
// in the constructed identifier, which may not always hold
// (e.g., aws_eks_addon).
func FormattedIdentifierUserDefinedNameFirst(param, separator string, keys ...string) config.ExternalName {
	e := config.ParameterAsIdentifier(param)
	e.GetIDFn = func(_ context.Context, externalName string, parameters map[string]interface{}, _ map[string]interface{}) (string, error) {
		vals := make([]string, len(keys)+1)
		for i, k := range keys {
			v, ok := parameters[k]
			if !ok {
				return "", errors.Errorf("%s cannot be empty", k)
			}
			s, ok := v.(string)
			if !ok {
				return "", errors.Errorf("%s needs to be a string", k)
			}
			vals[i+1] = s
		}
		vals[0] = externalName
		return strings.Join(vals, separator), nil
	}
	e.GetExternalNameFn = func(tfstate map[string]interface{}) (string, error) {
		id, ok := tfstate["id"]
		if !ok {
			return "", errors.New("id in tfstate cannot be empty")
		}
		s, ok := id.(string)
		if !ok {
			return "", errors.New("value of id needs to be string")
		}
		w := strings.Split(s, separator)
		return w[0], nil
	}
	return e
}

// TemplatedStringAsIdentifierWithNoName uses TemplatedStringAsIdentifier but
// without the name initializer. This allows it to be used in cases where the ID
// is constructed with parameters and a provider-defined value, meaning no
// user-defined input. Since the external name is not user-defined, the name
// initializer has to be disabled.
func TemplatedStringAsIdentifierWithNoName(tmpl string) config.ExternalName {
	e := config.TemplatedStringAsIdentifier("", tmpl)
	e.DisableNameInitializer = true
	return e
}

// ExternalNameConfigurations applies all external name configs listed in the
// table ExternalNameConfigs and sets the version of those resources to v1beta1
// assuming they will be tested.
func ExternalNameConfigurations() config.ResourceOption {
	return func(r *config.Resource) {
		if e, ok := ExternalNameConfigs[r.Name]; ok {
			r.Version = common.VersionV1Beta1
			r.ExternalName = e
		}
	}
}
