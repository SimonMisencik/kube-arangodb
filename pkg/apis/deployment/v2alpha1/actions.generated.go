//
// Copyright 2016-2022 ArangoDB GmbH, Cologne, Germany
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Copyright holder is ArangoDB GmbH, Cologne, Germany
//

package v2alpha1

import "time"

const (
	// Timeouts

	// ActionsDefaultTimeout define default timeout
	ActionsDefaultTimeout time.Duration = 600 * time.Second // 10m0s
	// ActionAddMemberDefaultTimeout define default timeout for action ActionAddMember
	ActionAddMemberDefaultTimeout time.Duration = 600 * time.Second // 10m0s
	// ActionAppendTLSCACertificateDefaultTimeout define default timeout for action ActionAppendTLSCACertificate
	ActionAppendTLSCACertificateDefaultTimeout time.Duration = 1800 * time.Second // 30m0s
	// ActionArangoMemberUpdatePodSpecDefaultTimeout define default timeout for action ActionArangoMemberUpdatePodSpec
	ActionArangoMemberUpdatePodSpecDefaultTimeout time.Duration = ActionsDefaultTimeout
	// ActionArangoMemberUpdatePodStatusDefaultTimeout define default timeout for action ActionArangoMemberUpdatePodStatus
	ActionArangoMemberUpdatePodStatusDefaultTimeout time.Duration = ActionsDefaultTimeout
	// ActionBackupRestoreDefaultTimeout define default timeout for action ActionBackupRestore
	ActionBackupRestoreDefaultTimeout time.Duration = 900 * time.Second // 15m0s
	// ActionBackupRestoreCleanDefaultTimeout define default timeout for action ActionBackupRestoreClean
	ActionBackupRestoreCleanDefaultTimeout time.Duration = 900 * time.Second // 15m0s
	// ActionBootstrapSetPasswordDefaultTimeout define default timeout for action ActionBootstrapSetPassword
	ActionBootstrapSetPasswordDefaultTimeout time.Duration = ActionsDefaultTimeout
	// ActionBootstrapUpdateDefaultTimeout define default timeout for action ActionBootstrapUpdate
	ActionBootstrapUpdateDefaultTimeout time.Duration = ActionsDefaultTimeout
	// ActionCleanOutMemberDefaultTimeout define default timeout for action ActionCleanOutMember
	ActionCleanOutMemberDefaultTimeout time.Duration = 172800 * time.Second // 48h0m0s
	// ActionCleanTLSCACertificateDefaultTimeout define default timeout for action ActionCleanTLSCACertificate
	ActionCleanTLSCACertificateDefaultTimeout time.Duration = 1800 * time.Second // 30m0s
	// ActionCleanTLSKeyfileCertificateDefaultTimeout define default timeout for action ActionCleanTLSKeyfileCertificate
	ActionCleanTLSKeyfileCertificateDefaultTimeout time.Duration = 1800 * time.Second // 30m0s
	// ActionClusterMemberCleanupDefaultTimeout define default timeout for action ActionClusterMemberCleanup
	ActionClusterMemberCleanupDefaultTimeout time.Duration = ActionsDefaultTimeout
	// ActionDisableClusterScalingDefaultTimeout define default timeout for action ActionDisableClusterScaling
	ActionDisableClusterScalingDefaultTimeout time.Duration = ActionsDefaultTimeout
	// ActionDisableMaintenanceDefaultTimeout define default timeout for action ActionDisableMaintenance
	ActionDisableMaintenanceDefaultTimeout time.Duration = ActionsDefaultTimeout
	// ActionDisableMemberMaintenanceDefaultTimeout define default timeout for action ActionDisableMemberMaintenance
	ActionDisableMemberMaintenanceDefaultTimeout time.Duration = ActionsDefaultTimeout
	// ActionEnableClusterScalingDefaultTimeout define default timeout for action ActionEnableClusterScaling
	ActionEnableClusterScalingDefaultTimeout time.Duration = ActionsDefaultTimeout
	// ActionEnableMaintenanceDefaultTimeout define default timeout for action ActionEnableMaintenance
	ActionEnableMaintenanceDefaultTimeout time.Duration = ActionsDefaultTimeout
	// ActionEnableMemberMaintenanceDefaultTimeout define default timeout for action ActionEnableMemberMaintenance
	ActionEnableMemberMaintenanceDefaultTimeout time.Duration = ActionsDefaultTimeout
	// ActionEncryptionKeyAddDefaultTimeout define default timeout for action ActionEncryptionKeyAdd
	ActionEncryptionKeyAddDefaultTimeout time.Duration = ActionsDefaultTimeout
	// ActionEncryptionKeyPropagatedDefaultTimeout define default timeout for action ActionEncryptionKeyPropagated
	ActionEncryptionKeyPropagatedDefaultTimeout time.Duration = ActionsDefaultTimeout
	// ActionEncryptionKeyRefreshDefaultTimeout define default timeout for action ActionEncryptionKeyRefresh
	ActionEncryptionKeyRefreshDefaultTimeout time.Duration = ActionsDefaultTimeout
	// ActionEncryptionKeyRemoveDefaultTimeout define default timeout for action ActionEncryptionKeyRemove
	ActionEncryptionKeyRemoveDefaultTimeout time.Duration = ActionsDefaultTimeout
	// ActionEncryptionKeyStatusUpdateDefaultTimeout define default timeout for action ActionEncryptionKeyStatusUpdate
	ActionEncryptionKeyStatusUpdateDefaultTimeout time.Duration = ActionsDefaultTimeout
	// ActionIdleDefaultTimeout define default timeout for action ActionIdle
	ActionIdleDefaultTimeout time.Duration = ActionsDefaultTimeout
	// ActionJWTAddDefaultTimeout define default timeout for action ActionJWTAdd
	ActionJWTAddDefaultTimeout time.Duration = ActionsDefaultTimeout
	// ActionJWTCleanDefaultTimeout define default timeout for action ActionJWTClean
	ActionJWTCleanDefaultTimeout time.Duration = ActionsDefaultTimeout
	// ActionJWTPropagatedDefaultTimeout define default timeout for action ActionJWTPropagated
	ActionJWTPropagatedDefaultTimeout time.Duration = ActionsDefaultTimeout
	// ActionJWTRefreshDefaultTimeout define default timeout for action ActionJWTRefresh
	ActionJWTRefreshDefaultTimeout time.Duration = ActionsDefaultTimeout
	// ActionJWTSetActiveDefaultTimeout define default timeout for action ActionJWTSetActive
	ActionJWTSetActiveDefaultTimeout time.Duration = ActionsDefaultTimeout
	// ActionJWTStatusUpdateDefaultTimeout define default timeout for action ActionJWTStatusUpdate
	ActionJWTStatusUpdateDefaultTimeout time.Duration = ActionsDefaultTimeout
	// ActionKillMemberPodDefaultTimeout define default timeout for action ActionKillMemberPod
	ActionKillMemberPodDefaultTimeout time.Duration = ActionsDefaultTimeout
	// ActionLicenseSetDefaultTimeout define default timeout for action ActionLicenseSet
	ActionLicenseSetDefaultTimeout time.Duration = ActionsDefaultTimeout
	// ActionMarkToRemoveMemberDefaultTimeout define default timeout for action ActionMarkToRemoveMember
	ActionMarkToRemoveMemberDefaultTimeout time.Duration = ActionsDefaultTimeout
	// ActionMemberPhaseUpdateDefaultTimeout define default timeout for action ActionMemberPhaseUpdate
	ActionMemberPhaseUpdateDefaultTimeout time.Duration = ActionsDefaultTimeout
	// ActionMemberRIDUpdateDefaultTimeout define default timeout for action ActionMemberRIDUpdate
	ActionMemberRIDUpdateDefaultTimeout time.Duration = ActionsDefaultTimeout
	// ActionPVCResizeDefaultTimeout define default timeout for action ActionPVCResize
	ActionPVCResizeDefaultTimeout time.Duration = 1800 * time.Second // 30m0s
	// ActionPVCResizedDefaultTimeout define default timeout for action ActionPVCResized
	ActionPVCResizedDefaultTimeout time.Duration = 900 * time.Second // 15m0s
	// ActionPlaceHolderDefaultTimeout define default timeout for action ActionPlaceHolder
	ActionPlaceHolderDefaultTimeout time.Duration = ActionsDefaultTimeout
	// ActionRebalancerCheckDefaultTimeout define default timeout for action ActionRebalancerCheck
	ActionRebalancerCheckDefaultTimeout time.Duration = ActionsDefaultTimeout
	// ActionRebalancerCleanDefaultTimeout define default timeout for action ActionRebalancerClean
	ActionRebalancerCleanDefaultTimeout time.Duration = ActionsDefaultTimeout
	// ActionRebalancerGenerateDefaultTimeout define default timeout for action ActionRebalancerGenerate
	ActionRebalancerGenerateDefaultTimeout time.Duration = ActionsDefaultTimeout
	// ActionRecreateMemberDefaultTimeout define default timeout for action ActionRecreateMember
	ActionRecreateMemberDefaultTimeout time.Duration = 900 * time.Second // 15m0s
	// ActionRefreshTLSKeyfileCertificateDefaultTimeout define default timeout for action ActionRefreshTLSKeyfileCertificate
	ActionRefreshTLSKeyfileCertificateDefaultTimeout time.Duration = 1800 * time.Second // 30m0s
	// ActionRemoveMemberDefaultTimeout define default timeout for action ActionRemoveMember
	ActionRemoveMemberDefaultTimeout time.Duration = 900 * time.Second // 15m0s
	// ActionRenewTLSCACertificateDefaultTimeout define default timeout for action ActionRenewTLSCACertificate
	ActionRenewTLSCACertificateDefaultTimeout time.Duration = 1800 * time.Second // 30m0s
	// ActionRenewTLSCertificateDefaultTimeout define default timeout for action ActionRenewTLSCertificate
	ActionRenewTLSCertificateDefaultTimeout time.Duration = 1800 * time.Second // 30m0s
	// ActionResignLeadershipDefaultTimeout define default timeout for action ActionResignLeadership
	ActionResignLeadershipDefaultTimeout time.Duration = 1800 * time.Second // 30m0s
	// ActionResourceSyncDefaultTimeout define default timeout for action ActionResourceSync
	ActionResourceSyncDefaultTimeout time.Duration = ActionsDefaultTimeout
	// ActionRotateMemberDefaultTimeout define default timeout for action ActionRotateMember
	ActionRotateMemberDefaultTimeout time.Duration = 900 * time.Second // 15m0s
	// ActionRotateStartMemberDefaultTimeout define default timeout for action ActionRotateStartMember
	ActionRotateStartMemberDefaultTimeout time.Duration = 900 * time.Second // 15m0s
	// ActionRotateStopMemberDefaultTimeout define default timeout for action ActionRotateStopMember
	ActionRotateStopMemberDefaultTimeout time.Duration = 900 * time.Second // 15m0s
	// ActionRuntimeContainerArgsLogLevelUpdateDefaultTimeout define default timeout for action ActionRuntimeContainerArgsLogLevelUpdate
	ActionRuntimeContainerArgsLogLevelUpdateDefaultTimeout time.Duration = ActionsDefaultTimeout
	// ActionRuntimeContainerImageUpdateDefaultTimeout define default timeout for action ActionRuntimeContainerImageUpdate
	ActionRuntimeContainerImageUpdateDefaultTimeout time.Duration = ActionsDefaultTimeout
	// ActionSetConditionDefaultTimeout define default timeout for action ActionSetCondition
	ActionSetConditionDefaultTimeout time.Duration = ActionsDefaultTimeout
	// ActionSetConditionV2DefaultTimeout define default timeout for action ActionSetConditionV2
	ActionSetConditionV2DefaultTimeout time.Duration = ActionsDefaultTimeout
	// ActionSetCurrentImageDefaultTimeout define default timeout for action ActionSetCurrentImage
	ActionSetCurrentImageDefaultTimeout time.Duration = 21600 * time.Second // 6h0m0s
	// ActionSetMaintenanceConditionDefaultTimeout define default timeout for action ActionSetMaintenanceCondition
	ActionSetMaintenanceConditionDefaultTimeout time.Duration = ActionsDefaultTimeout
	// ActionSetMemberConditionDefaultTimeout define default timeout for action ActionSetMemberCondition
	ActionSetMemberConditionDefaultTimeout time.Duration = ActionsDefaultTimeout
	// ActionSetMemberConditionV2DefaultTimeout define default timeout for action ActionSetMemberConditionV2
	ActionSetMemberConditionV2DefaultTimeout time.Duration = ActionsDefaultTimeout
	// ActionSetMemberCurrentImageDefaultTimeout define default timeout for action ActionSetMemberCurrentImage
	ActionSetMemberCurrentImageDefaultTimeout time.Duration = ActionsDefaultTimeout
	// ActionShutdownMemberDefaultTimeout define default timeout for action ActionShutdownMember
	ActionShutdownMemberDefaultTimeout time.Duration = 1800 * time.Second // 30m0s
	// ActionTLSKeyStatusUpdateDefaultTimeout define default timeout for action ActionTLSKeyStatusUpdate
	ActionTLSKeyStatusUpdateDefaultTimeout time.Duration = ActionsDefaultTimeout
	// ActionTLSPropagatedDefaultTimeout define default timeout for action ActionTLSPropagated
	ActionTLSPropagatedDefaultTimeout time.Duration = ActionsDefaultTimeout
	// ActionTimezoneSecretSetDefaultTimeout define default timeout for action ActionTimezoneSecretSet
	ActionTimezoneSecretSetDefaultTimeout time.Duration = 1800 * time.Second // 30m0s
	// ActionTopologyDisableDefaultTimeout define default timeout for action ActionTopologyDisable
	ActionTopologyDisableDefaultTimeout time.Duration = ActionsDefaultTimeout
	// ActionTopologyEnableDefaultTimeout define default timeout for action ActionTopologyEnable
	ActionTopologyEnableDefaultTimeout time.Duration = ActionsDefaultTimeout
	// ActionTopologyMemberAssignmentDefaultTimeout define default timeout for action ActionTopologyMemberAssignment
	ActionTopologyMemberAssignmentDefaultTimeout time.Duration = ActionsDefaultTimeout
	// ActionTopologyZonesUpdateDefaultTimeout define default timeout for action ActionTopologyZonesUpdate
	ActionTopologyZonesUpdateDefaultTimeout time.Duration = ActionsDefaultTimeout
	// ActionUpToDateUpdateDefaultTimeout define default timeout for action ActionUpToDateUpdate
	ActionUpToDateUpdateDefaultTimeout time.Duration = ActionsDefaultTimeout
	// ActionUpdateTLSSNIDefaultTimeout define default timeout for action ActionUpdateTLSSNI
	ActionUpdateTLSSNIDefaultTimeout time.Duration = 600 * time.Second // 10m0s
	// ActionUpgradeMemberDefaultTimeout define default timeout for action ActionUpgradeMember
	ActionUpgradeMemberDefaultTimeout time.Duration = 21600 * time.Second // 6h0m0s
	// ActionWaitForMemberInSyncDefaultTimeout define default timeout for action ActionWaitForMemberInSync
	ActionWaitForMemberInSyncDefaultTimeout time.Duration = 1800 * time.Second // 30m0s
	// ActionWaitForMemberUpDefaultTimeout define default timeout for action ActionWaitForMemberUp
	ActionWaitForMemberUpDefaultTimeout time.Duration = 1800 * time.Second // 30m0s

	// Actions

	// ActionTypeAddMember in scopes Normal. Adds new member to the Member list
	ActionTypeAddMember ActionType = "AddMember"
	// ActionTypeAppendTLSCACertificate in scopes Normal. Append Certificate into CA TrustStore
	ActionTypeAppendTLSCACertificate ActionType = "AppendTLSCACertificate"
	// ActionTypeArangoMemberUpdatePodSpec in scopes High. Propagate Member Pod spec (requested)
	ActionTypeArangoMemberUpdatePodSpec ActionType = "ArangoMemberUpdatePodSpec"
	// ActionTypeArangoMemberUpdatePodStatus in scopes High. Propagate Member Pod status (current)
	ActionTypeArangoMemberUpdatePodStatus ActionType = "ArangoMemberUpdatePodStatus"
	// ActionTypeBackupRestore in scopes Normal. Restore selected Backup
	ActionTypeBackupRestore ActionType = "BackupRestore"
	// ActionTypeBackupRestoreClean in scopes Normal. Clean restore status in case of restore spec change
	ActionTypeBackupRestoreClean ActionType = "BackupRestoreClean"
	// ActionTypeBootstrapSetPassword in scopes Normal. Change password during bootstrap procedure
	ActionTypeBootstrapSetPassword ActionType = "BootstrapSetPassword"
	// ActionTypeBootstrapUpdate in scopes Normal. Update bootstrap status
	ActionTypeBootstrapUpdate ActionType = "BootstrapUpdate"
	// ActionTypeCleanOutMember in scopes Normal. Run the CleanOut job on member
	ActionTypeCleanOutMember ActionType = "CleanOutMember"
	// ActionTypeCleanTLSCACertificate in scopes Normal. Remove Certificate from CA TrustStore
	ActionTypeCleanTLSCACertificate ActionType = "CleanTLSCACertificate"
	// ActionTypeCleanTLSKeyfileCertificate in scopes Normal. Remove old TLS certificate from server
	ActionTypeCleanTLSKeyfileCertificate ActionType = "CleanTLSKeyfileCertificate"
	// ActionTypeClusterMemberCleanup in scopes Normal. Remove member from Cluster if it is gone already (Coordinators)
	ActionTypeClusterMemberCleanup ActionType = "ClusterMemberCleanup"
	// ActionTypeDisableClusterScaling in scopes Normal. (Deprecated) Disable Cluster Scaling integration
	ActionTypeDisableClusterScaling ActionType = "DisableClusterScaling"
	// ActionTypeDisableMaintenance in scopes Normal. Disable ArangoDB maintenance mode
	ActionTypeDisableMaintenance ActionType = "DisableMaintenance"
	// ActionTypeDisableMemberMaintenance in scopes Normal. Disable ArangoDB DBServer maintenance mode
	ActionTypeDisableMemberMaintenance ActionType = "DisableMemberMaintenance"
	// ActionTypeEnableClusterScaling in scopes Normal. (Deprecated) Enable Cluster Scaling integration
	ActionTypeEnableClusterScaling ActionType = "EnableClusterScaling"
	// ActionTypeEnableMaintenance in scopes Normal. Enable ArangoDB maintenance mode
	ActionTypeEnableMaintenance ActionType = "EnableMaintenance"
	// ActionTypeEnableMemberMaintenance in scopes Normal. Enable ArangoDB DBServer maintenance mode
	ActionTypeEnableMemberMaintenance ActionType = "EnableMemberMaintenance"
	// ActionTypeEncryptionKeyAdd in scopes Normal. Add the encryption key to the pool
	ActionTypeEncryptionKeyAdd ActionType = "EncryptionKeyAdd"
	// ActionTypeEncryptionKeyPropagated in scopes Normal. Update condition of encryption propagation
	ActionTypeEncryptionKeyPropagated ActionType = "EncryptionKeyPropagated"
	// ActionTypeEncryptionKeyRefresh in scopes Normal. Refresh the encryption keys on member
	ActionTypeEncryptionKeyRefresh ActionType = "EncryptionKeyRefresh"
	// ActionTypeEncryptionKeyRemove in scopes Normal. Remove the encryption key to the pool
	ActionTypeEncryptionKeyRemove ActionType = "EncryptionKeyRemove"
	// ActionTypeEncryptionKeyStatusUpdate in scopes Normal. Update status of encryption propagation
	ActionTypeEncryptionKeyStatusUpdate ActionType = "EncryptionKeyStatusUpdate"
	// ActionTypeIdle in scopes Normal. Define idle operation in case if preconditions are not meet
	ActionTypeIdle ActionType = "Idle"
	// ActionTypeJWTAdd in scopes Normal. Adds new JWT to the pool
	ActionTypeJWTAdd ActionType = "JWTAdd"
	// ActionTypeJWTClean in scopes Normal. Remove JWT key from the pool
	ActionTypeJWTClean ActionType = "JWTClean"
	// ActionTypeJWTPropagated in scopes Normal. Update condition of JWT propagation
	ActionTypeJWTPropagated ActionType = "JWTPropagated"
	// ActionTypeJWTRefresh in scopes Normal. Refresh current JWT secrets on the member
	ActionTypeJWTRefresh ActionType = "JWTRefresh"
	// ActionTypeJWTSetActive in scopes Normal. Change active JWT key on the cluster
	ActionTypeJWTSetActive ActionType = "JWTSetActive"
	// ActionTypeJWTStatusUpdate in scopes Normal. Update status of JWT propagation
	ActionTypeJWTStatusUpdate ActionType = "JWTStatusUpdate"
	// ActionTypeKillMemberPod in scopes Normal. Execute Delete on Pod 9put pod in Terminating state)
	ActionTypeKillMemberPod ActionType = "KillMemberPod"
	// ActionTypeLicenseSet in scopes Normal. Update Cluster license (3.9+)
	ActionTypeLicenseSet ActionType = "LicenseSet"
	// ActionTypeMarkToRemoveMember in scopes Normal. Marks member to be removed. Used when member Pod is annotated with replace annotation
	ActionTypeMarkToRemoveMember ActionType = "MarkToRemoveMember"
	// ActionTypeMemberPhaseUpdate in scopes High. Change member phase
	ActionTypeMemberPhaseUpdate ActionType = "MemberPhaseUpdate"
	// ActionTypeMemberRIDUpdate in scopes High. Update Run ID of member
	ActionTypeMemberRIDUpdate ActionType = "MemberRIDUpdate"
	// ActionTypePVCResize in scopes Normal. Start the resize procedure. Updates PVC Requests field
	ActionTypePVCResize ActionType = "PVCResize"
	// ActionTypePVCResized in scopes Normal. Waits for PVC resize to be completed
	ActionTypePVCResized ActionType = "PVCResized"
	// ActionTypePlaceHolder in scopes Normal. Empty placeholder action
	ActionTypePlaceHolder ActionType = "PlaceHolder"
	// ActionTypeRebalancerCheck in scopes Normal. Check Rebalancer job progress
	ActionTypeRebalancerCheck ActionType = "RebalancerCheck"
	// ActionTypeRebalancerClean in scopes Normal. Cleans Rebalancer jobs
	ActionTypeRebalancerClean ActionType = "RebalancerClean"
	// ActionTypeRebalancerGenerate in scopes Normal. Generates the Rebalancer plan
	ActionTypeRebalancerGenerate ActionType = "RebalancerGenerate"
	// ActionTypeRecreateMember in scopes Normal. Recreate member with same ID and Data
	ActionTypeRecreateMember ActionType = "RecreateMember"
	// ActionTypeRefreshTLSKeyfileCertificate in scopes Normal. Recreate Server TLS Certificate secret
	ActionTypeRefreshTLSKeyfileCertificate ActionType = "RefreshTLSKeyfileCertificate"
	// ActionTypeRemoveMember in scopes Normal. Removes member from the Cluster and Status
	ActionTypeRemoveMember ActionType = "RemoveMember"
	// ActionTypeRenewTLSCACertificate in scopes Normal. Recreate Managed CA secret
	ActionTypeRenewTLSCACertificate ActionType = "RenewTLSCACertificate"
	// ActionTypeRenewTLSCertificate in scopes Normal. Recreate Server TLS Certificate secret
	ActionTypeRenewTLSCertificate ActionType = "RenewTLSCertificate"
	// ActionTypeResignLeadership in scopes Normal. Run the ResignLeadership job on DBServer
	ActionTypeResignLeadership ActionType = "ResignLeadership"
	// ActionTypeResourceSync in scopes Normal. Runs the Resource sync
	ActionTypeResourceSync ActionType = "ResourceSync"
	// ActionTypeRotateMember in scopes Normal. Waits for Pod restart and recreation
	ActionTypeRotateMember ActionType = "RotateMember"
	// ActionTypeRotateStartMember in scopes Normal. Start member rotation. After this action member is down
	ActionTypeRotateStartMember ActionType = "RotateStartMember"
	// ActionTypeRotateStopMember in scopes Normal. Finalize member rotation. After this action member is started back
	ActionTypeRotateStopMember ActionType = "RotateStopMember"
	// ActionTypeRuntimeContainerArgsLogLevelUpdate in scopes Normal. Change ArangoDB Member log levels in runtime
	ActionTypeRuntimeContainerArgsLogLevelUpdate ActionType = "RuntimeContainerArgsLogLevelUpdate"
	// ActionTypeRuntimeContainerImageUpdate in scopes Normal. Update Container Image in runtime
	ActionTypeRuntimeContainerImageUpdate ActionType = "RuntimeContainerImageUpdate"
	// ActionTypeSetCondition in scopes High. (Deprecated) Set deployment condition
	ActionTypeSetCondition ActionType = "SetCondition"
	// ActionTypeSetConditionV2 in scopes High. Set deployment condition
	ActionTypeSetConditionV2 ActionType = "SetConditionV2"
	// ActionTypeSetCurrentImage in scopes Normal. Update deployment current image after image discovery
	ActionTypeSetCurrentImage ActionType = "SetCurrentImage"
	// ActionTypeSetMaintenanceCondition in scopes Normal. Update ArangoDB maintenance condition
	ActionTypeSetMaintenanceCondition ActionType = "SetMaintenanceCondition"
	// ActionTypeSetMemberCondition in scopes High. (Deprecated) Set member condition
	ActionTypeSetMemberCondition ActionType = "SetMemberCondition"
	// ActionTypeSetMemberConditionV2 in scopes High. Set member condition
	ActionTypeSetMemberConditionV2 ActionType = "SetMemberConditionV2"
	// ActionTypeSetMemberCurrentImage in scopes Normal. Update Member current image
	ActionTypeSetMemberCurrentImage ActionType = "SetMemberCurrentImage"
	// ActionTypeShutdownMember in scopes Normal. Sends Shutdown requests and waits for container to be stopped
	ActionTypeShutdownMember ActionType = "ShutdownMember"
	// ActionTypeTLSKeyStatusUpdate in scopes Normal. Update Status of TLS propagation process
	ActionTypeTLSKeyStatusUpdate ActionType = "TLSKeyStatusUpdate"
	// ActionTypeTLSPropagated in scopes Normal. Update TLS propagation condition
	ActionTypeTLSPropagated ActionType = "TLSPropagated"
	// ActionTypeTimezoneSecretSet in scopes Normal. Set timezone details in cluster
	ActionTypeTimezoneSecretSet ActionType = "TimezoneSecretSet"
	// ActionTypeTopologyDisable in scopes Normal. Disable TopologyAwareness
	ActionTypeTopologyDisable ActionType = "TopologyDisable"
	// ActionTypeTopologyEnable in scopes Normal. Enable TopologyAwareness
	ActionTypeTopologyEnable ActionType = "TopologyEnable"
	// ActionTypeTopologyMemberAssignment in scopes Normal. Update TopologyAwareness Members assignments
	ActionTypeTopologyMemberAssignment ActionType = "TopologyMemberAssignment"
	// ActionTypeTopologyZonesUpdate in scopes Normal. Update TopologyAwareness Zones info
	ActionTypeTopologyZonesUpdate ActionType = "TopologyZonesUpdate"
	// ActionTypeUpToDateUpdate in scopes Normal. Update UpToDate condition
	ActionTypeUpToDateUpdate ActionType = "UpToDateUpdate"
	// ActionTypeUpdateTLSSNI in scopes Normal. Update certificate in SNI
	ActionTypeUpdateTLSSNI ActionType = "UpdateTLSSNI"
	// ActionTypeUpgradeMember in scopes Normal. Run the Upgrade procedure on member
	ActionTypeUpgradeMember ActionType = "UpgradeMember"
	// ActionTypeWaitForMemberInSync in scopes Normal. Wait for member to be in sync. In case of DBServer waits for shards. In case of Agents to catch-up on Agency index
	ActionTypeWaitForMemberInSync ActionType = "WaitForMemberInSync"
	// ActionTypeWaitForMemberUp in scopes Normal. Wait for member to be responsive
	ActionTypeWaitForMemberUp ActionType = "WaitForMemberUp"
)

func ActionDefaultTimeout(in ActionType) time.Duration {
	switch in {
	case ActionTypeAddMember:
		return ActionAddMemberDefaultTimeout
	case ActionTypeAppendTLSCACertificate:
		return ActionAppendTLSCACertificateDefaultTimeout
	case ActionTypeArangoMemberUpdatePodSpec:
		return ActionArangoMemberUpdatePodSpecDefaultTimeout
	case ActionTypeArangoMemberUpdatePodStatus:
		return ActionArangoMemberUpdatePodStatusDefaultTimeout
	case ActionTypeBackupRestore:
		return ActionBackupRestoreDefaultTimeout
	case ActionTypeBackupRestoreClean:
		return ActionBackupRestoreCleanDefaultTimeout
	case ActionTypeBootstrapSetPassword:
		return ActionBootstrapSetPasswordDefaultTimeout
	case ActionTypeBootstrapUpdate:
		return ActionBootstrapUpdateDefaultTimeout
	case ActionTypeCleanOutMember:
		return ActionCleanOutMemberDefaultTimeout
	case ActionTypeCleanTLSCACertificate:
		return ActionCleanTLSCACertificateDefaultTimeout
	case ActionTypeCleanTLSKeyfileCertificate:
		return ActionCleanTLSKeyfileCertificateDefaultTimeout
	case ActionTypeClusterMemberCleanup:
		return ActionClusterMemberCleanupDefaultTimeout
	case ActionTypeDisableClusterScaling:
		return ActionDisableClusterScalingDefaultTimeout
	case ActionTypeDisableMaintenance:
		return ActionDisableMaintenanceDefaultTimeout
	case ActionTypeDisableMemberMaintenance:
		return ActionDisableMemberMaintenanceDefaultTimeout
	case ActionTypeEnableClusterScaling:
		return ActionEnableClusterScalingDefaultTimeout
	case ActionTypeEnableMaintenance:
		return ActionEnableMaintenanceDefaultTimeout
	case ActionTypeEnableMemberMaintenance:
		return ActionEnableMemberMaintenanceDefaultTimeout
	case ActionTypeEncryptionKeyAdd:
		return ActionEncryptionKeyAddDefaultTimeout
	case ActionTypeEncryptionKeyPropagated:
		return ActionEncryptionKeyPropagatedDefaultTimeout
	case ActionTypeEncryptionKeyRefresh:
		return ActionEncryptionKeyRefreshDefaultTimeout
	case ActionTypeEncryptionKeyRemove:
		return ActionEncryptionKeyRemoveDefaultTimeout
	case ActionTypeEncryptionKeyStatusUpdate:
		return ActionEncryptionKeyStatusUpdateDefaultTimeout
	case ActionTypeIdle:
		return ActionIdleDefaultTimeout
	case ActionTypeJWTAdd:
		return ActionJWTAddDefaultTimeout
	case ActionTypeJWTClean:
		return ActionJWTCleanDefaultTimeout
	case ActionTypeJWTPropagated:
		return ActionJWTPropagatedDefaultTimeout
	case ActionTypeJWTRefresh:
		return ActionJWTRefreshDefaultTimeout
	case ActionTypeJWTSetActive:
		return ActionJWTSetActiveDefaultTimeout
	case ActionTypeJWTStatusUpdate:
		return ActionJWTStatusUpdateDefaultTimeout
	case ActionTypeKillMemberPod:
		return ActionKillMemberPodDefaultTimeout
	case ActionTypeLicenseSet:
		return ActionLicenseSetDefaultTimeout
	case ActionTypeMarkToRemoveMember:
		return ActionMarkToRemoveMemberDefaultTimeout
	case ActionTypeMemberPhaseUpdate:
		return ActionMemberPhaseUpdateDefaultTimeout
	case ActionTypeMemberRIDUpdate:
		return ActionMemberRIDUpdateDefaultTimeout
	case ActionTypePVCResize:
		return ActionPVCResizeDefaultTimeout
	case ActionTypePVCResized:
		return ActionPVCResizedDefaultTimeout
	case ActionTypePlaceHolder:
		return ActionPlaceHolderDefaultTimeout
	case ActionTypeRebalancerCheck:
		return ActionRebalancerCheckDefaultTimeout
	case ActionTypeRebalancerClean:
		return ActionRebalancerCleanDefaultTimeout
	case ActionTypeRebalancerGenerate:
		return ActionRebalancerGenerateDefaultTimeout
	case ActionTypeRecreateMember:
		return ActionRecreateMemberDefaultTimeout
	case ActionTypeRefreshTLSKeyfileCertificate:
		return ActionRefreshTLSKeyfileCertificateDefaultTimeout
	case ActionTypeRemoveMember:
		return ActionRemoveMemberDefaultTimeout
	case ActionTypeRenewTLSCACertificate:
		return ActionRenewTLSCACertificateDefaultTimeout
	case ActionTypeRenewTLSCertificate:
		return ActionRenewTLSCertificateDefaultTimeout
	case ActionTypeResignLeadership:
		return ActionResignLeadershipDefaultTimeout
	case ActionTypeResourceSync:
		return ActionResourceSyncDefaultTimeout
	case ActionTypeRotateMember:
		return ActionRotateMemberDefaultTimeout
	case ActionTypeRotateStartMember:
		return ActionRotateStartMemberDefaultTimeout
	case ActionTypeRotateStopMember:
		return ActionRotateStopMemberDefaultTimeout
	case ActionTypeRuntimeContainerArgsLogLevelUpdate:
		return ActionRuntimeContainerArgsLogLevelUpdateDefaultTimeout
	case ActionTypeRuntimeContainerImageUpdate:
		return ActionRuntimeContainerImageUpdateDefaultTimeout
	case ActionTypeSetCondition:
		return ActionSetConditionDefaultTimeout
	case ActionTypeSetConditionV2:
		return ActionSetConditionV2DefaultTimeout
	case ActionTypeSetCurrentImage:
		return ActionSetCurrentImageDefaultTimeout
	case ActionTypeSetMaintenanceCondition:
		return ActionSetMaintenanceConditionDefaultTimeout
	case ActionTypeSetMemberCondition:
		return ActionSetMemberConditionDefaultTimeout
	case ActionTypeSetMemberConditionV2:
		return ActionSetMemberConditionV2DefaultTimeout
	case ActionTypeSetMemberCurrentImage:
		return ActionSetMemberCurrentImageDefaultTimeout
	case ActionTypeShutdownMember:
		return ActionShutdownMemberDefaultTimeout
	case ActionTypeTLSKeyStatusUpdate:
		return ActionTLSKeyStatusUpdateDefaultTimeout
	case ActionTypeTLSPropagated:
		return ActionTLSPropagatedDefaultTimeout
	case ActionTypeTimezoneSecretSet:
		return ActionTimezoneSecretSetDefaultTimeout
	case ActionTypeTopologyDisable:
		return ActionTopologyDisableDefaultTimeout
	case ActionTypeTopologyEnable:
		return ActionTopologyEnableDefaultTimeout
	case ActionTypeTopologyMemberAssignment:
		return ActionTopologyMemberAssignmentDefaultTimeout
	case ActionTypeTopologyZonesUpdate:
		return ActionTopologyZonesUpdateDefaultTimeout
	case ActionTypeUpToDateUpdate:
		return ActionUpToDateUpdateDefaultTimeout
	case ActionTypeUpdateTLSSNI:
		return ActionUpdateTLSSNIDefaultTimeout
	case ActionTypeUpgradeMember:
		return ActionUpgradeMemberDefaultTimeout
	case ActionTypeWaitForMemberInSync:
		return ActionWaitForMemberInSyncDefaultTimeout
	case ActionTypeWaitForMemberUp:
		return ActionWaitForMemberUpDefaultTimeout
	default:
		return ActionsDefaultTimeout
	}
}

func GetActionPriority(in ActionType) ActionPriority {
	switch in {
	case ActionTypeAddMember:
		return ActionPriorityNormal
	case ActionTypeAppendTLSCACertificate:
		return ActionPriorityNormal
	case ActionTypeArangoMemberUpdatePodSpec:
		return ActionPriorityHigh
	case ActionTypeArangoMemberUpdatePodStatus:
		return ActionPriorityHigh
	case ActionTypeBackupRestore:
		return ActionPriorityNormal
	case ActionTypeBackupRestoreClean:
		return ActionPriorityNormal
	case ActionTypeBootstrapSetPassword:
		return ActionPriorityNormal
	case ActionTypeBootstrapUpdate:
		return ActionPriorityNormal
	case ActionTypeCleanOutMember:
		return ActionPriorityNormal
	case ActionTypeCleanTLSCACertificate:
		return ActionPriorityNormal
	case ActionTypeCleanTLSKeyfileCertificate:
		return ActionPriorityNormal
	case ActionTypeClusterMemberCleanup:
		return ActionPriorityNormal
	case ActionTypeDisableClusterScaling:
		return ActionPriorityNormal
	case ActionTypeDisableMaintenance:
		return ActionPriorityNormal
	case ActionTypeDisableMemberMaintenance:
		return ActionPriorityNormal
	case ActionTypeEnableClusterScaling:
		return ActionPriorityNormal
	case ActionTypeEnableMaintenance:
		return ActionPriorityNormal
	case ActionTypeEnableMemberMaintenance:
		return ActionPriorityNormal
	case ActionTypeEncryptionKeyAdd:
		return ActionPriorityNormal
	case ActionTypeEncryptionKeyPropagated:
		return ActionPriorityNormal
	case ActionTypeEncryptionKeyRefresh:
		return ActionPriorityNormal
	case ActionTypeEncryptionKeyRemove:
		return ActionPriorityNormal
	case ActionTypeEncryptionKeyStatusUpdate:
		return ActionPriorityNormal
	case ActionTypeIdle:
		return ActionPriorityNormal
	case ActionTypeJWTAdd:
		return ActionPriorityNormal
	case ActionTypeJWTClean:
		return ActionPriorityNormal
	case ActionTypeJWTPropagated:
		return ActionPriorityNormal
	case ActionTypeJWTRefresh:
		return ActionPriorityNormal
	case ActionTypeJWTSetActive:
		return ActionPriorityNormal
	case ActionTypeJWTStatusUpdate:
		return ActionPriorityNormal
	case ActionTypeKillMemberPod:
		return ActionPriorityNormal
	case ActionTypeLicenseSet:
		return ActionPriorityNormal
	case ActionTypeMarkToRemoveMember:
		return ActionPriorityNormal
	case ActionTypeMemberPhaseUpdate:
		return ActionPriorityHigh
	case ActionTypeMemberRIDUpdate:
		return ActionPriorityHigh
	case ActionTypePVCResize:
		return ActionPriorityNormal
	case ActionTypePVCResized:
		return ActionPriorityNormal
	case ActionTypePlaceHolder:
		return ActionPriorityNormal
	case ActionTypeRebalancerCheck:
		return ActionPriorityNormal
	case ActionTypeRebalancerClean:
		return ActionPriorityNormal
	case ActionTypeRebalancerGenerate:
		return ActionPriorityNormal
	case ActionTypeRecreateMember:
		return ActionPriorityNormal
	case ActionTypeRefreshTLSKeyfileCertificate:
		return ActionPriorityNormal
	case ActionTypeRemoveMember:
		return ActionPriorityNormal
	case ActionTypeRenewTLSCACertificate:
		return ActionPriorityNormal
	case ActionTypeRenewTLSCertificate:
		return ActionPriorityNormal
	case ActionTypeResignLeadership:
		return ActionPriorityNormal
	case ActionTypeResourceSync:
		return ActionPriorityNormal
	case ActionTypeRotateMember:
		return ActionPriorityNormal
	case ActionTypeRotateStartMember:
		return ActionPriorityNormal
	case ActionTypeRotateStopMember:
		return ActionPriorityNormal
	case ActionTypeRuntimeContainerArgsLogLevelUpdate:
		return ActionPriorityNormal
	case ActionTypeRuntimeContainerImageUpdate:
		return ActionPriorityNormal
	case ActionTypeSetCondition:
		return ActionPriorityHigh
	case ActionTypeSetConditionV2:
		return ActionPriorityHigh
	case ActionTypeSetCurrentImage:
		return ActionPriorityNormal
	case ActionTypeSetMaintenanceCondition:
		return ActionPriorityNormal
	case ActionTypeSetMemberCondition:
		return ActionPriorityHigh
	case ActionTypeSetMemberConditionV2:
		return ActionPriorityHigh
	case ActionTypeSetMemberCurrentImage:
		return ActionPriorityNormal
	case ActionTypeShutdownMember:
		return ActionPriorityNormal
	case ActionTypeTLSKeyStatusUpdate:
		return ActionPriorityNormal
	case ActionTypeTLSPropagated:
		return ActionPriorityNormal
	case ActionTypeTimezoneSecretSet:
		return ActionPriorityNormal
	case ActionTypeTopologyDisable:
		return ActionPriorityNormal
	case ActionTypeTopologyEnable:
		return ActionPriorityNormal
	case ActionTypeTopologyMemberAssignment:
		return ActionPriorityNormal
	case ActionTypeTopologyZonesUpdate:
		return ActionPriorityNormal
	case ActionTypeUpToDateUpdate:
		return ActionPriorityNormal
	case ActionTypeUpdateTLSSNI:
		return ActionPriorityNormal
	case ActionTypeUpgradeMember:
		return ActionPriorityNormal
	case ActionTypeWaitForMemberInSync:
		return ActionPriorityNormal
	case ActionTypeWaitForMemberUp:
		return ActionPriorityNormal
	default:
		return ActionPriorityUnknown
	}
}
