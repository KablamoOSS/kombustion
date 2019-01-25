package parsers

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"github.com/KablamoOSS/kombustion/pkg/parsers/resources"
	"github.com/KablamoOSS/kombustion/types"
)

// GetParsersResources returns parser functions
func GetParsersResources() map[string]types.ParserFunc {
	return map[string]types.ParserFunc{
		"AWS::AmazonMQ::Broker":                                 resources.ParseAmazonMQBroker,
		"AWS::AmazonMQ::Configuration":                          resources.ParseAmazonMQConfiguration,
		"AWS::AmazonMQ::ConfigurationAssociation":               resources.ParseAmazonMQConfigurationAssociation,
		"AWS::ApiGateway::Account":                              resources.ParseApiGatewayAccount,
		"AWS::ApiGateway::ApiKey":                               resources.ParseApiGatewayApiKey,
		"AWS::ApiGateway::Authorizer":                           resources.ParseApiGatewayAuthorizer,
		"AWS::ApiGateway::BasePathMapping":                      resources.ParseApiGatewayBasePathMapping,
		"AWS::ApiGateway::ClientCertificate":                    resources.ParseApiGatewayClientCertificate,
		"AWS::ApiGateway::Deployment":                           resources.ParseApiGatewayDeployment,
		"AWS::ApiGateway::DocumentationPart":                    resources.ParseApiGatewayDocumentationPart,
		"AWS::ApiGateway::DocumentationVersion":                 resources.ParseApiGatewayDocumentationVersion,
		"AWS::ApiGateway::DomainName":                           resources.ParseApiGatewayDomainName,
		"AWS::ApiGateway::GatewayResponse":                      resources.ParseApiGatewayGatewayResponse,
		"AWS::ApiGateway::Method":                               resources.ParseApiGatewayMethod,
		"AWS::ApiGateway::Model":                                resources.ParseApiGatewayModel,
		"AWS::ApiGateway::RequestValidator":                     resources.ParseApiGatewayRequestValidator,
		"AWS::ApiGateway::Resource":                             resources.ParseApiGatewayResource,
		"AWS::ApiGateway::RestApi":                              resources.ParseApiGatewayRestApi,
		"AWS::ApiGateway::Stage":                                resources.ParseApiGatewayStage,
		"AWS::ApiGateway::UsagePlan":                            resources.ParseApiGatewayUsagePlan,
		"AWS::ApiGateway::UsagePlanKey":                         resources.ParseApiGatewayUsagePlanKey,
		"AWS::ApiGateway::VpcLink":                              resources.ParseApiGatewayVpcLink,
		"AWS::AppStream::DirectoryConfig":                       resources.ParseAppStreamDirectoryConfig,
		"AWS::AppStream::Fleet":                                 resources.ParseAppStreamFleet,
		"AWS::AppStream::ImageBuilder":                          resources.ParseAppStreamImageBuilder,
		"AWS::AppStream::Stack":                                 resources.ParseAppStreamStack,
		"AWS::AppStream::StackFleetAssociation":                 resources.ParseAppStreamStackFleetAssociation,
		"AWS::AppStream::StackUserAssociation":                  resources.ParseAppStreamStackUserAssociation,
		"AWS::AppStream::User":                                  resources.ParseAppStreamUser,
		"AWS::AppSync::ApiKey":                                  resources.ParseAppSyncApiKey,
		"AWS::AppSync::DataSource":                              resources.ParseAppSyncDataSource,
		"AWS::AppSync::FunctionConfiguration":                   resources.ParseAppSyncFunctionConfiguration,
		"AWS::AppSync::GraphQLApi":                              resources.ParseAppSyncGraphQLApi,
		"AWS::AppSync::GraphQLSchema":                           resources.ParseAppSyncGraphQLSchema,
		"AWS::AppSync::Resolver":                                resources.ParseAppSyncResolver,
		"AWS::ApplicationAutoScaling::ScalableTarget":           resources.ParseApplicationAutoScalingScalableTarget,
		"AWS::ApplicationAutoScaling::ScalingPolicy":            resources.ParseApplicationAutoScalingScalingPolicy,
		"AWS::Athena::NamedQuery":                               resources.ParseAthenaNamedQuery,
		"AWS::AutoScaling::AutoScalingGroup":                    resources.ParseAutoScalingAutoScalingGroup,
		"AWS::AutoScaling::LaunchConfiguration":                 resources.ParseAutoScalingLaunchConfiguration,
		"AWS::AutoScaling::LifecycleHook":                       resources.ParseAutoScalingLifecycleHook,
		"AWS::AutoScaling::ScalingPolicy":                       resources.ParseAutoScalingScalingPolicy,
		"AWS::AutoScaling::ScheduledAction":                     resources.ParseAutoScalingScheduledAction,
		"AWS::AutoScalingPlans::ScalingPlan":                    resources.ParseAutoScalingPlansScalingPlan,
		"AWS::Batch::ComputeEnvironment":                        resources.ParseBatchComputeEnvironment,
		"AWS::Batch::JobDefinition":                             resources.ParseBatchJobDefinition,
		"AWS::Batch::JobQueue":                                  resources.ParseBatchJobQueue,
		"AWS::Budgets::Budget":                                  resources.ParseBudgetsBudget,
		"AWS::CertificateManager::Certificate":                  resources.ParseCertificateManagerCertificate,
		"AWS::Cloud9::EnvironmentEC2":                           resources.ParseCloud9EnvironmentEC2,
		"AWS::CloudFormation::CustomResource":                   resources.ParseCloudFormationCustomResource,
		"AWS::CloudFormation::Macro":                            resources.ParseCloudFormationMacro,
		"AWS::CloudFormation::Stack":                            resources.ParseCloudFormationStack,
		"AWS::CloudFormation::WaitCondition":                    resources.ParseCloudFormationWaitCondition,
		"AWS::CloudFormation::WaitConditionHandle":              resources.ParseCloudFormationWaitConditionHandle,
		"AWS::CloudFront::CloudFrontOriginAccessIdentity":       resources.ParseCloudFrontCloudFrontOriginAccessIdentity,
		"AWS::CloudFront::Distribution":                         resources.ParseCloudFrontDistribution,
		"AWS::CloudFront::StreamingDistribution":                resources.ParseCloudFrontStreamingDistribution,
		"AWS::CloudTrail::Trail":                                resources.ParseCloudTrailTrail,
		"AWS::CloudWatch::Alarm":                                resources.ParseCloudWatchAlarm,
		"AWS::CloudWatch::Dashboard":                            resources.ParseCloudWatchDashboard,
		"AWS::CodeBuild::Project":                               resources.ParseCodeBuildProject,
		"AWS::CodeCommit::Repository":                           resources.ParseCodeCommitRepository,
		"AWS::CodeDeploy::Application":                          resources.ParseCodeDeployApplication,
		"AWS::CodeDeploy::DeploymentConfig":                     resources.ParseCodeDeployDeploymentConfig,
		"AWS::CodeDeploy::DeploymentGroup":                      resources.ParseCodeDeployDeploymentGroup,
		"AWS::CodePipeline::CustomActionType":                   resources.ParseCodePipelineCustomActionType,
		"AWS::CodePipeline::Pipeline":                           resources.ParseCodePipelinePipeline,
		"AWS::CodePipeline::Webhook":                            resources.ParseCodePipelineWebhook,
		"AWS::Cognito::IdentityPool":                            resources.ParseCognitoIdentityPool,
		"AWS::Cognito::IdentityPoolRoleAttachment":              resources.ParseCognitoIdentityPoolRoleAttachment,
		"AWS::Cognito::UserPool":                                resources.ParseCognitoUserPool,
		"AWS::Cognito::UserPoolClient":                          resources.ParseCognitoUserPoolClient,
		"AWS::Cognito::UserPoolGroup":                           resources.ParseCognitoUserPoolGroup,
		"AWS::Cognito::UserPoolUser":                            resources.ParseCognitoUserPoolUser,
		"AWS::Cognito::UserPoolUserToGroupAttachment":           resources.ParseCognitoUserPoolUserToGroupAttachment,
		"AWS::Config::AggregationAuthorization":                 resources.ParseConfigAggregationAuthorization,
		"AWS::Config::ConfigRule":                               resources.ParseConfigConfigRule,
		"AWS::Config::ConfigurationAggregator":                  resources.ParseConfigConfigurationAggregator,
		"AWS::Config::ConfigurationRecorder":                    resources.ParseConfigConfigurationRecorder,
		"AWS::Config::DeliveryChannel":                          resources.ParseConfigDeliveryChannel,
		"AWS::DAX::Cluster":                                     resources.ParseDAXCluster,
		"AWS::DAX::ParameterGroup":                              resources.ParseDAXParameterGroup,
		"AWS::DAX::SubnetGroup":                                 resources.ParseDAXSubnetGroup,
		"AWS::DLM::LifecyclePolicy":                             resources.ParseDLMLifecyclePolicy,
		"AWS::DMS::Certificate":                                 resources.ParseDMSCertificate,
		"AWS::DMS::Endpoint":                                    resources.ParseDMSEndpoint,
		"AWS::DMS::EventSubscription":                           resources.ParseDMSEventSubscription,
		"AWS::DMS::ReplicationInstance":                         resources.ParseDMSReplicationInstance,
		"AWS::DMS::ReplicationSubnetGroup":                      resources.ParseDMSReplicationSubnetGroup,
		"AWS::DMS::ReplicationTask":                             resources.ParseDMSReplicationTask,
		"AWS::DataPipeline::Pipeline":                           resources.ParseDataPipelinePipeline,
		"AWS::DirectoryService::MicrosoftAD":                    resources.ParseDirectoryServiceMicrosoftAD,
		"AWS::DirectoryService::SimpleAD":                       resources.ParseDirectoryServiceSimpleAD,
		"AWS::DocDB::DBCluster":                                 resources.ParseDocDBDBCluster,
		"AWS::DocDB::DBClusterParameterGroup":                   resources.ParseDocDBDBClusterParameterGroup,
		"AWS::DocDB::DBInstance":                                resources.ParseDocDBDBInstance,
		"AWS::DocDB::DBSubnetGroup":                             resources.ParseDocDBDBSubnetGroup,
		"AWS::DynamoDB::Table":                                  resources.ParseDynamoDBTable,
		"AWS::EC2::CustomerGateway":                             resources.ParseEC2CustomerGateway,
		"AWS::EC2::DHCPOptions":                                 resources.ParseEC2DHCPOptions,
		"AWS::EC2::EC2Fleet":                                    resources.ParseEC2EC2Fleet,
		"AWS::EC2::EIP":                                         resources.ParseEC2EIP,
		"AWS::EC2::EIPAssociation":                              resources.ParseEC2EIPAssociation,
		"AWS::EC2::EgressOnlyInternetGateway":                   resources.ParseEC2EgressOnlyInternetGateway,
		"AWS::EC2::FlowLog":                                     resources.ParseEC2FlowLog,
		"AWS::EC2::Host":                                        resources.ParseEC2Host,
		"AWS::EC2::Instance":                                    resources.ParseEC2Instance,
		"AWS::EC2::InternetGateway":                             resources.ParseEC2InternetGateway,
		"AWS::EC2::LaunchTemplate":                              resources.ParseEC2LaunchTemplate,
		"AWS::EC2::NatGateway":                                  resources.ParseEC2NatGateway,
		"AWS::EC2::NetworkAcl":                                  resources.ParseEC2NetworkAcl,
		"AWS::EC2::NetworkAclEntry":                             resources.ParseEC2NetworkAclEntry,
		"AWS::EC2::NetworkInterface":                            resources.ParseEC2NetworkInterface,
		"AWS::EC2::NetworkInterfaceAttachment":                  resources.ParseEC2NetworkInterfaceAttachment,
		"AWS::EC2::NetworkInterfacePermission":                  resources.ParseEC2NetworkInterfacePermission,
		"AWS::EC2::PlacementGroup":                              resources.ParseEC2PlacementGroup,
		"AWS::EC2::Route":                                       resources.ParseEC2Route,
		"AWS::EC2::RouteTable":                                  resources.ParseEC2RouteTable,
		"AWS::EC2::SecurityGroup":                               resources.ParseEC2SecurityGroup,
		"AWS::EC2::SecurityGroupEgress":                         resources.ParseEC2SecurityGroupEgress,
		"AWS::EC2::SecurityGroupIngress":                        resources.ParseEC2SecurityGroupIngress,
		"AWS::EC2::SpotFleet":                                   resources.ParseEC2SpotFleet,
		"AWS::EC2::Subnet":                                      resources.ParseEC2Subnet,
		"AWS::EC2::SubnetCidrBlock":                             resources.ParseEC2SubnetCidrBlock,
		"AWS::EC2::SubnetNetworkAclAssociation":                 resources.ParseEC2SubnetNetworkAclAssociation,
		"AWS::EC2::SubnetRouteTableAssociation":                 resources.ParseEC2SubnetRouteTableAssociation,
		"AWS::EC2::TransitGateway":                              resources.ParseEC2TransitGateway,
		"AWS::EC2::TransitGatewayAttachment":                    resources.ParseEC2TransitGatewayAttachment,
		"AWS::EC2::TransitGatewayRoute":                         resources.ParseEC2TransitGatewayRoute,
		"AWS::EC2::TransitGatewayRouteTable":                    resources.ParseEC2TransitGatewayRouteTable,
		"AWS::EC2::TransitGatewayRouteTableAssociation":         resources.ParseEC2TransitGatewayRouteTableAssociation,
		"AWS::EC2::TransitGatewayRouteTablePropagation":         resources.ParseEC2TransitGatewayRouteTablePropagation,
		"AWS::EC2::TrunkInterfaceAssociation":                   resources.ParseEC2TrunkInterfaceAssociation,
		"AWS::EC2::VPC":                                         resources.ParseEC2VPC,
		"AWS::EC2::VPCCidrBlock":                                resources.ParseEC2VPCCidrBlock,
		"AWS::EC2::VPCDHCPOptionsAssociation":                   resources.ParseEC2VPCDHCPOptionsAssociation,
		"AWS::EC2::VPCEndpoint":                                 resources.ParseEC2VPCEndpoint,
		"AWS::EC2::VPCEndpointConnectionNotification":           resources.ParseEC2VPCEndpointConnectionNotification,
		"AWS::EC2::VPCEndpointServicePermissions":               resources.ParseEC2VPCEndpointServicePermissions,
		"AWS::EC2::VPCGatewayAttachment":                        resources.ParseEC2VPCGatewayAttachment,
		"AWS::EC2::VPCPeeringConnection":                        resources.ParseEC2VPCPeeringConnection,
		"AWS::EC2::VPNConnection":                               resources.ParseEC2VPNConnection,
		"AWS::EC2::VPNConnectionRoute":                          resources.ParseEC2VPNConnectionRoute,
		"AWS::EC2::VPNGateway":                                  resources.ParseEC2VPNGateway,
		"AWS::EC2::VPNGatewayRoutePropagation":                  resources.ParseEC2VPNGatewayRoutePropagation,
		"AWS::EC2::Volume":                                      resources.ParseEC2Volume,
		"AWS::EC2::VolumeAttachment":                            resources.ParseEC2VolumeAttachment,
		"AWS::ECR::Repository":                                  resources.ParseECRRepository,
		"AWS::ECS::Cluster":                                     resources.ParseECSCluster,
		"AWS::ECS::Service":                                     resources.ParseECSService,
		"AWS::ECS::TaskDefinition":                              resources.ParseECSTaskDefinition,
		"AWS::EFS::FileSystem":                                  resources.ParseEFSFileSystem,
		"AWS::EFS::MountTarget":                                 resources.ParseEFSMountTarget,
		"AWS::EKS::Cluster":                                     resources.ParseEKSCluster,
		"AWS::EMR::Cluster":                                     resources.ParseEMRCluster,
		"AWS::EMR::InstanceFleetConfig":                         resources.ParseEMRInstanceFleetConfig,
		"AWS::EMR::InstanceGroupConfig":                         resources.ParseEMRInstanceGroupConfig,
		"AWS::EMR::SecurityConfiguration":                       resources.ParseEMRSecurityConfiguration,
		"AWS::EMR::Step":                                        resources.ParseEMRStep,
		"AWS::ElastiCache::CacheCluster":                        resources.ParseElastiCacheCacheCluster,
		"AWS::ElastiCache::ParameterGroup":                      resources.ParseElastiCacheParameterGroup,
		"AWS::ElastiCache::ReplicationGroup":                    resources.ParseElastiCacheReplicationGroup,
		"AWS::ElastiCache::SecurityGroup":                       resources.ParseElastiCacheSecurityGroup,
		"AWS::ElastiCache::SecurityGroupIngress":                resources.ParseElastiCacheSecurityGroupIngress,
		"AWS::ElastiCache::SubnetGroup":                         resources.ParseElastiCacheSubnetGroup,
		"AWS::ElasticBeanstalk::Application":                    resources.ParseElasticBeanstalkApplication,
		"AWS::ElasticBeanstalk::ApplicationVersion":             resources.ParseElasticBeanstalkApplicationVersion,
		"AWS::ElasticBeanstalk::ConfigurationTemplate":          resources.ParseElasticBeanstalkConfigurationTemplate,
		"AWS::ElasticBeanstalk::Environment":                    resources.ParseElasticBeanstalkEnvironment,
		"AWS::ElasticLoadBalancing::LoadBalancer":               resources.ParseElasticLoadBalancingLoadBalancer,
		"AWS::ElasticLoadBalancingV2::Listener":                 resources.ParseElasticLoadBalancingV2Listener,
		"AWS::ElasticLoadBalancingV2::ListenerCertificate":      resources.ParseElasticLoadBalancingV2ListenerCertificate,
		"AWS::ElasticLoadBalancingV2::ListenerRule":             resources.ParseElasticLoadBalancingV2ListenerRule,
		"AWS::ElasticLoadBalancingV2::LoadBalancer":             resources.ParseElasticLoadBalancingV2LoadBalancer,
		"AWS::ElasticLoadBalancingV2::TargetGroup":              resources.ParseElasticLoadBalancingV2TargetGroup,
		"AWS::Elasticsearch::Domain":                            resources.ParseElasticsearchDomain,
		"AWS::Events::EventBusPolicy":                           resources.ParseEventsEventBusPolicy,
		"AWS::Events::Rule":                                     resources.ParseEventsRule,
		"AWS::GameLift::Alias":                                  resources.ParseGameLiftAlias,
		"AWS::GameLift::Build":                                  resources.ParseGameLiftBuild,
		"AWS::GameLift::Fleet":                                  resources.ParseGameLiftFleet,
		"AWS::Glue::Classifier":                                 resources.ParseGlueClassifier,
		"AWS::Glue::Connection":                                 resources.ParseGlueConnection,
		"AWS::Glue::Crawler":                                    resources.ParseGlueCrawler,
		"AWS::Glue::Database":                                   resources.ParseGlueDatabase,
		"AWS::Glue::DevEndpoint":                                resources.ParseGlueDevEndpoint,
		"AWS::Glue::Job":                                        resources.ParseGlueJob,
		"AWS::Glue::Partition":                                  resources.ParseGluePartition,
		"AWS::Glue::Table":                                      resources.ParseGlueTable,
		"AWS::Glue::Trigger":                                    resources.ParseGlueTrigger,
		"AWS::GuardDuty::Detector":                              resources.ParseGuardDutyDetector,
		"AWS::GuardDuty::Filter":                                resources.ParseGuardDutyFilter,
		"AWS::GuardDuty::IPSet":                                 resources.ParseGuardDutyIPSet,
		"AWS::GuardDuty::Master":                                resources.ParseGuardDutyMaster,
		"AWS::GuardDuty::Member":                                resources.ParseGuardDutyMember,
		"AWS::GuardDuty::ThreatIntelSet":                        resources.ParseGuardDutyThreatIntelSet,
		"AWS::IAM::AccessKey":                                   resources.ParseIAMAccessKey,
		"AWS::IAM::Group":                                       resources.ParseIAMGroup,
		"AWS::IAM::InstanceProfile":                             resources.ParseIAMInstanceProfile,
		"AWS::IAM::ManagedPolicy":                               resources.ParseIAMManagedPolicy,
		"AWS::IAM::Policy":                                      resources.ParseIAMPolicy,
		"AWS::IAM::Role":                                        resources.ParseIAMRole,
		"AWS::IAM::ServiceLinkedRole":                           resources.ParseIAMServiceLinkedRole,
		"AWS::IAM::User":                                        resources.ParseIAMUser,
		"AWS::IAM::UserToGroupAddition":                         resources.ParseIAMUserToGroupAddition,
		"AWS::Inspector::AssessmentTarget":                      resources.ParseInspectorAssessmentTarget,
		"AWS::Inspector::AssessmentTemplate":                    resources.ParseInspectorAssessmentTemplate,
		"AWS::Inspector::ResourceGroup":                         resources.ParseInspectorResourceGroup,
		"AWS::IoT1Click::Device":                                resources.ParseIoT1ClickDevice,
		"AWS::IoT1Click::Placement":                             resources.ParseIoT1ClickPlacement,
		"AWS::IoT1Click::Project":                               resources.ParseIoT1ClickProject,
		"AWS::IoT::Certificate":                                 resources.ParseIoTCertificate,
		"AWS::IoT::Policy":                                      resources.ParseIoTPolicy,
		"AWS::IoT::PolicyPrincipalAttachment":                   resources.ParseIoTPolicyPrincipalAttachment,
		"AWS::IoT::Thing":                                       resources.ParseIoTThing,
		"AWS::IoT::ThingPrincipalAttachment":                    resources.ParseIoTThingPrincipalAttachment,
		"AWS::IoT::TopicRule":                                   resources.ParseIoTTopicRule,
		"AWS::IoTAnalytics::Channel":                            resources.ParseIoTAnalyticsChannel,
		"AWS::IoTAnalytics::Dataset":                            resources.ParseIoTAnalyticsDataset,
		"AWS::IoTAnalytics::Datastore":                          resources.ParseIoTAnalyticsDatastore,
		"AWS::IoTAnalytics::Pipeline":                           resources.ParseIoTAnalyticsPipeline,
		"AWS::KMS::Alias":                                       resources.ParseKMSAlias,
		"AWS::KMS::Key":                                         resources.ParseKMSKey,
		"AWS::Kinesis::Stream":                                  resources.ParseKinesisStream,
		"AWS::Kinesis::StreamConsumer":                          resources.ParseKinesisStreamConsumer,
		"AWS::KinesisAnalytics::Application":                    resources.ParseKinesisAnalyticsApplication,
		"AWS::KinesisAnalytics::ApplicationOutput":              resources.ParseKinesisAnalyticsApplicationOutput,
		"AWS::KinesisAnalytics::ApplicationReferenceDataSource": resources.ParseKinesisAnalyticsApplicationReferenceDataSource,
		"AWS::KinesisFirehose::DeliveryStream":                  resources.ParseKinesisFirehoseDeliveryStream,
		"AWS::Lambda::Alias":                                    resources.ParseLambdaAlias,
		"AWS::Lambda::EventSourceMapping":                       resources.ParseLambdaEventSourceMapping,
		"AWS::Lambda::Function":                                 resources.ParseLambdaFunction,
		"AWS::Lambda::LayerVersion":                             resources.ParseLambdaLayerVersion,
		"AWS::Lambda::LayerVersionPermission":                   resources.ParseLambdaLayerVersionPermission,
		"AWS::Lambda::Permission":                               resources.ParseLambdaPermission,
		"AWS::Lambda::Version":                                  resources.ParseLambdaVersion,
		"AWS::Logs::Destination":                                resources.ParseLogsDestination,
		"AWS::Logs::LogGroup":                                   resources.ParseLogsLogGroup,
		"AWS::Logs::LogStream":                                  resources.ParseLogsLogStream,
		"AWS::Logs::MetricFilter":                               resources.ParseLogsMetricFilter,
		"AWS::Logs::SubscriptionFilter":                         resources.ParseLogsSubscriptionFilter,
		"AWS::Neptune::DBCluster":                               resources.ParseNeptuneDBCluster,
		"AWS::Neptune::DBClusterParameterGroup":                 resources.ParseNeptuneDBClusterParameterGroup,
		"AWS::Neptune::DBInstance":                              resources.ParseNeptuneDBInstance,
		"AWS::Neptune::DBParameterGroup":                        resources.ParseNeptuneDBParameterGroup,
		"AWS::Neptune::DBSubnetGroup":                           resources.ParseNeptuneDBSubnetGroup,
		"AWS::OpsWorks::App":                                    resources.ParseOpsWorksApp,
		"AWS::OpsWorks::ElasticLoadBalancerAttachment":          resources.ParseOpsWorksElasticLoadBalancerAttachment,
		"AWS::OpsWorks::Instance":                               resources.ParseOpsWorksInstance,
		"AWS::OpsWorks::Layer":                                  resources.ParseOpsWorksLayer,
		"AWS::OpsWorks::Stack":                                  resources.ParseOpsWorksStack,
		"AWS::OpsWorks::UserProfile":                            resources.ParseOpsWorksUserProfile,
		"AWS::OpsWorks::Volume":                                 resources.ParseOpsWorksVolume,
		"AWS::OpsWorksCM::Server":                               resources.ParseOpsWorksCMServer,
		"AWS::RDS::DBCluster":                                   resources.ParseRDSDBCluster,
		"AWS::RDS::DBClusterParameterGroup":                     resources.ParseRDSDBClusterParameterGroup,
		"AWS::RDS::DBInstance":                                  resources.ParseRDSDBInstance,
		"AWS::RDS::DBParameterGroup":                            resources.ParseRDSDBParameterGroup,
		"AWS::RDS::DBSecurityGroup":                             resources.ParseRDSDBSecurityGroup,
		"AWS::RDS::DBSecurityGroupIngress":                      resources.ParseRDSDBSecurityGroupIngress,
		"AWS::RDS::DBSubnetGroup":                               resources.ParseRDSDBSubnetGroup,
		"AWS::RDS::EventSubscription":                           resources.ParseRDSEventSubscription,
		"AWS::RDS::OptionGroup":                                 resources.ParseRDSOptionGroup,
		"AWS::Redshift::Cluster":                                resources.ParseRedshiftCluster,
		"AWS::Redshift::ClusterParameterGroup":                  resources.ParseRedshiftClusterParameterGroup,
		"AWS::Redshift::ClusterSecurityGroup":                   resources.ParseRedshiftClusterSecurityGroup,
		"AWS::Redshift::ClusterSecurityGroupIngress":            resources.ParseRedshiftClusterSecurityGroupIngress,
		"AWS::Redshift::ClusterSubnetGroup":                     resources.ParseRedshiftClusterSubnetGroup,
		"AWS::Route53::HealthCheck":                             resources.ParseRoute53HealthCheck,
		"AWS::Route53::HostedZone":                              resources.ParseRoute53HostedZone,
		"AWS::Route53::RecordSet":                               resources.ParseRoute53RecordSet,
		"AWS::Route53::RecordSetGroup":                          resources.ParseRoute53RecordSetGroup,
		"AWS::Route53Resolver::ResolverEndpoint":                resources.ParseRoute53ResolverResolverEndpoint,
		"AWS::Route53Resolver::ResolverRule":                    resources.ParseRoute53ResolverResolverRule,
		"AWS::Route53Resolver::ResolverRuleAssociation":         resources.ParseRoute53ResolverResolverRuleAssociation,
		"AWS::S3::Bucket":                                       resources.ParseS3Bucket,
		"AWS::S3::BucketPolicy":                                 resources.ParseS3BucketPolicy,
		"AWS::SDB::Domain":                                      resources.ParseSDBDomain,
		"AWS::SES::ConfigurationSet":                            resources.ParseSESConfigurationSet,
		"AWS::SES::ConfigurationSetEventDestination":            resources.ParseSESConfigurationSetEventDestination,
		"AWS::SES::ReceiptFilter":                               resources.ParseSESReceiptFilter,
		"AWS::SES::ReceiptRule":                                 resources.ParseSESReceiptRule,
		"AWS::SES::ReceiptRuleSet":                              resources.ParseSESReceiptRuleSet,
		"AWS::SES::Template":                                    resources.ParseSESTemplate,
		"AWS::SNS::Subscription":                                resources.ParseSNSSubscription,
		"AWS::SNS::Topic":                                       resources.ParseSNSTopic,
		"AWS::SNS::TopicPolicy":                                 resources.ParseSNSTopicPolicy,
		"AWS::SQS::Queue":                                       resources.ParseSQSQueue,
		"AWS::SQS::QueuePolicy":                                 resources.ParseSQSQueuePolicy,
		"AWS::SSM::Association":                                 resources.ParseSSMAssociation,
		"AWS::SSM::Document":                                    resources.ParseSSMDocument,
		"AWS::SSM::MaintenanceWindow":                           resources.ParseSSMMaintenanceWindow,
		"AWS::SSM::MaintenanceWindowTarget":                     resources.ParseSSMMaintenanceWindowTarget,
		"AWS::SSM::MaintenanceWindowTask":                       resources.ParseSSMMaintenanceWindowTask,
		"AWS::SSM::Parameter":                                   resources.ParseSSMParameter,
		"AWS::SSM::PatchBaseline":                               resources.ParseSSMPatchBaseline,
		"AWS::SSM::ResourceDataSync":                            resources.ParseSSMResourceDataSync,
		"AWS::SageMaker::Endpoint":                              resources.ParseSageMakerEndpoint,
		"AWS::SageMaker::EndpointConfig":                        resources.ParseSageMakerEndpointConfig,
		"AWS::SageMaker::Model":                                 resources.ParseSageMakerModel,
		"AWS::SageMaker::NotebookInstance":                      resources.ParseSageMakerNotebookInstance,
		"AWS::SageMaker::NotebookInstanceLifecycleConfig":       resources.ParseSageMakerNotebookInstanceLifecycleConfig,
		"AWS::SecretsManager::ResourcePolicy":                   resources.ParseSecretsManagerResourcePolicy,
		"AWS::SecretsManager::RotationSchedule":                 resources.ParseSecretsManagerRotationSchedule,
		"AWS::SecretsManager::Secret":                           resources.ParseSecretsManagerSecret,
		"AWS::SecretsManager::SecretTargetAttachment":           resources.ParseSecretsManagerSecretTargetAttachment,
		"AWS::ServiceCatalog::AcceptedPortfolioShare":           resources.ParseServiceCatalogAcceptedPortfolioShare,
		"AWS::ServiceCatalog::CloudFormationProduct":            resources.ParseServiceCatalogCloudFormationProduct,
		"AWS::ServiceCatalog::CloudFormationProvisionedProduct": resources.ParseServiceCatalogCloudFormationProvisionedProduct,
		"AWS::ServiceCatalog::LaunchNotificationConstraint":     resources.ParseServiceCatalogLaunchNotificationConstraint,
		"AWS::ServiceCatalog::LaunchRoleConstraint":             resources.ParseServiceCatalogLaunchRoleConstraint,
		"AWS::ServiceCatalog::LaunchTemplateConstraint":         resources.ParseServiceCatalogLaunchTemplateConstraint,
		"AWS::ServiceCatalog::Portfolio":                        resources.ParseServiceCatalogPortfolio,
		"AWS::ServiceCatalog::PortfolioPrincipalAssociation":    resources.ParseServiceCatalogPortfolioPrincipalAssociation,
		"AWS::ServiceCatalog::PortfolioProductAssociation":      resources.ParseServiceCatalogPortfolioProductAssociation,
		"AWS::ServiceCatalog::PortfolioShare":                   resources.ParseServiceCatalogPortfolioShare,
		"AWS::ServiceCatalog::TagOption":                        resources.ParseServiceCatalogTagOption,
		"AWS::ServiceCatalog::TagOptionAssociation":             resources.ParseServiceCatalogTagOptionAssociation,
		"AWS::ServiceDiscovery::HttpNamespace":                  resources.ParseServiceDiscoveryHttpNamespace,
		"AWS::ServiceDiscovery::Instance":                       resources.ParseServiceDiscoveryInstance,
		"AWS::ServiceDiscovery::PrivateDnsNamespace":            resources.ParseServiceDiscoveryPrivateDnsNamespace,
		"AWS::ServiceDiscovery::PublicDnsNamespace":             resources.ParseServiceDiscoveryPublicDnsNamespace,
		"AWS::ServiceDiscovery::Service":                        resources.ParseServiceDiscoveryService,
		"AWS::StepFunctions::Activity":                          resources.ParseStepFunctionsActivity,
		"AWS::StepFunctions::StateMachine":                      resources.ParseStepFunctionsStateMachine,
		"AWS::WAF::ByteMatchSet":                                resources.ParseWAFByteMatchSet,
		"AWS::WAF::IPSet":                                       resources.ParseWAFIPSet,
		"AWS::WAF::Rule":                                        resources.ParseWAFRule,
		"AWS::WAF::SizeConstraintSet":                           resources.ParseWAFSizeConstraintSet,
		"AWS::WAF::SqlInjectionMatchSet":                        resources.ParseWAFSqlInjectionMatchSet,
		"AWS::WAF::WebACL":                                      resources.ParseWAFWebACL,
		"AWS::WAF::XssMatchSet":                                 resources.ParseWAFXssMatchSet,
		"AWS::WAFRegional::ByteMatchSet":                        resources.ParseWAFRegionalByteMatchSet,
		"AWS::WAFRegional::IPSet":                               resources.ParseWAFRegionalIPSet,
		"AWS::WAFRegional::Rule":                                resources.ParseWAFRegionalRule,
		"AWS::WAFRegional::SizeConstraintSet":                   resources.ParseWAFRegionalSizeConstraintSet,
		"AWS::WAFRegional::SqlInjectionMatchSet":                resources.ParseWAFRegionalSqlInjectionMatchSet,
		"AWS::WAFRegional::WebACL":                              resources.ParseWAFRegionalWebACL,
		"AWS::WAFRegional::WebACLAssociation":                   resources.ParseWAFRegionalWebACLAssociation,
		"AWS::WAFRegional::XssMatchSet":                         resources.ParseWAFRegionalXssMatchSet,
		"AWS::WorkSpaces::Workspace":                            resources.ParseWorkSpacesWorkspace,
		"Alexa::ASK::Skill":                                     resources.ParseASKSkill,
	}
}
