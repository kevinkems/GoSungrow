package AppService

import (
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/acceptPsSharing"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/activateEmail"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/addConfig"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/addDeviceRepair"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/addDeviceToStructureForHousehold"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/addDeviceToStructureForHouseholdByPsIdS"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/addFault"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/addFaultOrder"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/addFaultPlan"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/addFaultRepairSteps"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/addHouseholdEvaluation"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/addHouseholdLeaveMessage"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/addHouseholdOpinionFeedback"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/addHouseholdWorkOrder"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/addOnDutyInfo"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/addOperRule"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/addOrDelPsStructure"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/addOrderStep"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/addPowerStationForHousehold"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/addPowerStationInfo"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/addReportConfigEmail"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/addSysAdvancedParam"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/addSysOrgNew"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/aliPayAppTest"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/auditOperRule"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/batchAddStationBySn"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/batchImportSN"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/batchInsertUserAndOrg"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/batchModifyDevicesInfoAndPropertis"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/batchProcessPlantReport"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/batchUpdateDeviceSim"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/batchUpdateUserIsAgreeGdpr"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/boundMobilePhone"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/boundUserMail"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/caculateDeviceInputDiscrete"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/calculateDeviceDiscrete"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/calculateInitialCompensationData"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/cancelDeliverMail"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/cancelOrderScan"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/cancelParamSetTask"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/cancelPsSharing"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/cancelRechargeOrder"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/changRechargeOrderToCancel"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/changeHouseholdUser2Installer"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/changeRemoteParam"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/checkDealerOrgCode"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/checkDevSnIsBelongsToUser"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/checkInverterResult"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/checkIsCanDoParamSet"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/checkIsIvScan"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/checkOssObjectExist"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/checkServiceIsConnect"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/checkTechnicalParameters"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/checkUnitStatus"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/checkUpRechargeDevicePaying"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/checkUserAccountUnique"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/checkUserAccountUniqueAll"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/checkUserInfoUnique"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/checkUserIsExist"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/checkUserListIsExist"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/checkUserPassword"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/cloudDeploymentRecord"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/comfirmParamModel"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/communicationModuleDetail"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/compareValidateCode"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/componentInfo2Cloud"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/confirmFault"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/confirmIvFault"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/confirmReportConfig"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/createAppkeyInfo"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/createRenewInvoice"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/dealCommandReply"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/dealDeletePsFailPsDelete"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/dealFailRemoteUpgradeSubTasks"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/dealFailRemoteUpgradeTasks"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/dealFaultOrder"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/dealGroupStringDisableOrEnable"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/dealNumberOfServiceCalls2Mysql"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/dealParamSettingAfterComplete"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/dealPsDataSupplement"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/dealPsReportEmailSend"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/dealRemoteUpgrade"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/dealSnElectrifyCheck"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/dealSysDeviceSimFlowInfo"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/dealSysDeviceSimInfo"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/definiteTimeDealSnExpRemind"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/definiteTimeDealSnStatus"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/delDeviceRepair"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/delOperRule"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/delayCallApiResidueTimes"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/deleteComponent"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/deleteCustomerEmployee"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/deleteDeviceAccount"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/deleteDeviceSimById"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/deleteElectricitySettlementData"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/deleteFaultPlan"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/deleteFirmwareFiles"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/deleteHouseholdEvaluation"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/deleteHouseholdLeaveMessage"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/deleteHouseholdWorkOrder"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/deleteInverterSnInChnnl"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/deleteModuleLog"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/deleteOnDutyInfo"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/deleteOperateBillFile"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/deleteOssObject"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/deletePowerDevicePointById"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/deletePowerPicture"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/deletePowerRobotInfoBySnAndPsId"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/deletePowerRobotSweepStrategy"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/deleteProductionData"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/deletePs"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/deleteRechargeOrder"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/deleteRegularlyConnectionInfo"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/deleteReportConfigEmailAddr"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/deleteSysAdvancedParam"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/deleteSysOrgNew"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/deleteTemplate"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/deleteUserInfoAllByUserId"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/deviceInputDiscreteDeleteTime"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/deviceInputDiscreteGetTime"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/deviceInputDiscreteInsertTime"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/deviceInputDiscreteUpdateTime"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/devicePointsDataFromMySql"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/deviceReplace"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/editDeviceRepair"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/editOperRule"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/energyPovertyAlleviation"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/energyTrend"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/exportParamSettingValPDF"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/exportPlantReportPDF"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/faultAutoClose"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/faultCloseRemindOrderHandler"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/findCodeValueList"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/findEmgOrgInfo"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/findEnvironmentInfo"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/findFromHbaseAndRedis"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/findInfoByuuid"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/findLossAnalysisList"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/findOnDutyInfo"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/findPsType"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/findSingleStationPR"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/findUserPassword"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/genTLSUserSigByUserAccount"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/generateRandomPassword"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getAPIServiceInfo"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getAccessedPermission"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getAllDeviceByPsId"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getAllPowerDeviceSetName"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getAllPowerRobotViewInfoByPsId"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getAllPsIdByOrgIds"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getAllUserRemindCount"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getAndOutletsAndUnit"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getApiCallsForAppkeys"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getAreaInfoCodeByCounty"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getAreaList"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getAutoCreatePowerStation"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getBackReadValue"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getBatchNewestPointData"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getCallApiResidueTimes"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getChangedPsListByTime"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getChnnlListByPsId"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getCloudList"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getCloudServiceMappingConfig"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getCommunicationDeviceConfigInfo"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getCommunicationModuleMonitorData"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getComponentModelFactory"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getConfigList"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getConnectionInfoBySnAndLocalPort"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getCountDown"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getCountryServiceInfo"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getCounty"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getCustomerEmployee"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getCustomerList"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getDataFromHBase"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getDataFromHbaseByRowKey"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getDevInstalledPowerByPsId"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getDevRecord"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getDevRunRecordList"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getDevSimByList"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getDevSimList"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getDeviceAccountById"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getDeviceFaultStatisticsData"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getDeviceInfo"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getDeviceList"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getDeviceModelInfoList"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getDevicePointMinuteDataList"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getDevicePoints"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getDevicePropertys"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getDeviceRepairDetail"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getDeviceTechBranchCount"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getDeviceTypeInfoList"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getDeviceTypeList"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getDstInfo"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getElectricitySettlementData"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getElectricitySettlementDetailData"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getEncryptPublicKey"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getFaultCount"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getFaultDetail"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getFaultMsgByFaultCode"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getFaultMsgListByPageWithYYYYMM"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getFaultMsgListWithYYYYMM"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getFaultPlanList"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getFileOperationRecordOne"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getFormulaFaultAnalyseList"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getGroupStringCheckResult"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getGroupStringCheckRule"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getHisData"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getHistoryInfo"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getHouseholdEvaluation"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getHouseholdLeaveMessage"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getHouseholdOpinionFeedback"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getHouseholdPsInstallerByUserId"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getHouseholdStoragePsReport"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getHouseholdUserInfo"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getHouseholdWorkOrderInfo"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getHouseholdWorkOrderList"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getI18nConfigByType"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getI18nFileInfo"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getI18nInfoByKey"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getI18nVersion"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getIncomeSettingInfos"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getInfoFromAMap"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getInfomationFromRedis"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getInstallInfoList"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getInstallerInfoByDealerOrgCodeOrId"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getInvertDataList"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getInverterDataCount"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getInverterProcess"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getInverterUuidBytotalId"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getIvEchartsData"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getIvEchartsDataById"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getKpiInfo"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getListMiFromHBase"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getMapInfo"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getMapMiFromHBase"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getMenuAndPrivileges"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getMicrogridEStoragePsReport"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getModuleLogInfo"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getModuleLogTaskList"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getNationProvJSON"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getNeedOpAsynOpRecordList"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getNoticeInfo"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getNumberOfServiceCalls"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getOSSConfig"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getOperRuleDetail"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getOperateBillFileId"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getOperateTicketForDetail"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getOrCreateNetEaseUserToken"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getOrderDataList"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getOrderDataSql2"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getOrderDatas"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getOrderDetail"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getOrderStatistics"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getOrgIdNameByUserId"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getOrgInfoByDealerOrgCode"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getOrgListByName"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getOrgListByUserId"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getOrgListForUser"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getOssObjectStream"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getOwnerFaultConfigList"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getPListinfoFromMysql"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getParamSetTemplate4NewProtocol"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getParamSetTemplatePointInfo"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getParamterSettingBase"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getPhotoInfo"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getPlanedOrNotPsList"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getPlantReportPDFList"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getPowerChargeSettingInfo"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getPowerDeviceModelTechList"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getPowerDeviceModelTree"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getPowerDevicePointInfo"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getPowerDevicePointNames"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getPowerDeviceSetTaskDetailList"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getPowerDeviceSetTaskList"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getPowerFormulaFaultAnalyse"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getPowerPictureList"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getPowerRobotInfoByRobotSn"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getPowerRobotSweepAttrByPsId"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getPowerRobotSweepStrategy"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getPowerRobotSweepStrategyList"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getPowerSettingCharges"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getPowerSettingHistoryRecords"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getPowerStationBasicInfo"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getPowerStationData"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getPowerStationForHousehold"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getPowerStationInfo"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getPowerStationPR"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getPowerStationTableDataSql"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getPowerStationTableDataSqlCount"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getPowerStatistics"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getPowerTrendDayData"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getPrivateCloudValidityPeriod"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getProvInfoListByNationCode"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getPsAuthKey"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getPsCurveInfo"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getPsDataSupplementTaskList"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getPsDetail"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getPsDetailByUserTokens"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getPsDetailForSinglePage"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getPsDetailWithPsType"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getPsHealthState"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getPsInstallerByPsId"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getPsInstallerOrgInfoByPsId"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getPsList"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getPsListByName"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getPsListForPsDataByPsId"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getPsListStaticData"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getPsReport"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getPsUser"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getPsWeatherList"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getRechargeOrderDetail"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getRechargeOrderItemDeviceList"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getRechargeOrderList"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getRegionalTree"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getRemoteParamSettingList"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getRemoteUpgradeDeviceList"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getRemoteUpgradeScheduleDetails"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getRemoteUpgradeSubTasksList"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getRemoteUpgradeTaskList"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getReportData"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getReportEmailConfigInfo"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getReportExportColumns"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getReportListByUserId"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getRobotDynamicCleaningView"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getRobotNumAndSweepCapacity"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getRuleUnit"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getSendReportConfigCron"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getSerialNum"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getShieldMapConditionList"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getSimIdBySnList"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getSingleIVData"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getSnChangeRecord"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getSnConnectionInfo"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getStationInfoSql"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getSungwsConfigCache"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getSungwsGlobalConfigCache"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getSweepDevParamSetTemplate"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getSweepRobotDevList"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getSysMsg"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getSysOrgNewList"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getSysOrgNewOne"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getSysUserById"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getTableDataSql"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getTableDataSqlCount"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getTemplateByInfoType"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getTemplateList"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getUUIDByUpuuid"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getUpTimePoint"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getUserById"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getUserByInstaller"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getUserDevOnlineOffineCount"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getUserGDPRAttrs"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getUserHavePowerStationCount"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getUserInfoByUserAccounts"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getUserList"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getUserPsOrderList"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getValFromHBase"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getValidateCode"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getValidateCodeAtRegister"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getWeatherInfo"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getWechatPushConfig"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/getWorkInfo"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/groupStringCheck"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/handleDevByCommunicationSN"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/householdResetPassBySN"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/immediatePayment"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/importExcelData"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/incomeStatistics"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/informPush"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/insertEmgOrgInfo"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/insightSynDeviceStructure2Cloud"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/intoDataToHbase"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/ipLocationQuery"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/isHave2GSn"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/judgeDevIsHasInitSetTemplate"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/judgeIsSettingMan"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/listOssFiles"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/loadAreaInfo"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/loadPowerStation"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/login"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/loginByToken"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/logout"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/mobilePhoneHasBound"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/modifiedDeviceInfo"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/modifyEmgOrgStruc"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/modifyFaultPlan"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/modifyOnDutyInfo"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/modifyPassword"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/modifyPersonalUnitList"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/modifyPsUser"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/moduleLogParamSet"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/operateOssFile"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/operationPowerRobotSweepStrategy"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/orgPowerReport"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/paramSetTryAgain"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/paramSetting"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/planPower"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/powerDevicePointList"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/powerTrendChartData"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/psForcastInfo"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/psHourPointsValue"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/queryAllPsIdAndName"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/queryBatchCreatePsTaskList"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/queryBatchSpeedyAddPowerStationResult"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/queryCardStatusCTCC"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/queryChildAccountList"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/queryCompensationRecordData"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/queryCompensationRecordList"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/queryComponent"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/queryComponentTechnicalParam"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/queryCountryGridAndRelation"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/queryCountryList"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/queryCtrlTaskById"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/queryDeviceInfo"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/queryDeviceInfoForApp"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/queryDeviceList"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/queryDeviceListByUserId"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/queryDeviceListForApp"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/queryDeviceModelTechnical"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/queryDevicePointDayMonthYearDataList"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/queryDevicePointMinuteDataList"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/queryDevicePointsDayMonthYearDataList"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/queryDeviceRealTimeDataByPsKeys"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/queryDeviceRepairList"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/queryDeviceTypeInfoList"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/queryEnvironmentList"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/queryFaultList"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/queryFaultPlanDetail"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/queryFaultRepairSteps"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/queryFaultTypeAndLevelByCode"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/queryFaultTypeByDevice"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/queryFaultTypeByDevicePage"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/queryFirmwareFilesPage"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/queryInfotoAlert"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/queryInverterModelList"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/queryInverterVersionList"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/queryM2MCardInfoCMCC"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/queryM2MCardTermInfoCMCC"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/queryModelInfoByModelId"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/queryMutiPointDataList"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/queryNoticeList"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/queryNumberOfRenewalReminders"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/queryOperRules"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/queryOrderList"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/queryOrderStep"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/queryOrgGenerationReport"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/queryOrgInfoList"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/queryOrgPowerElecPercent"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/queryOrgPsCompensationRecordList"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/queryParamSettingTask"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/queryPersonalUnitList"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/queryPointDataTopOne"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/queryPowerStationInfo"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/queryPsAreaByUserIdAndAreaCode"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/queryPsCompensationRecordList"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/queryPsDataByDate"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/queryPsIdList"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/queryPsListByUserIdAndAreaCode"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/queryPsNameByPsId"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/queryPsPrByDate"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/queryPsProfit"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/queryPsReportComparativeAnalysisOfPowerGeneration"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/queryPsStructureList"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/queryPuuidsByCommandTotalId"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/queryPuuidsByCommandTotalId2"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/queryRepairRuleList"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/queryReportListForManagementPage"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/queryReportMsg"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/querySharingPs"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/querySysAdvancedParam"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/queryTimeBySN"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/queryTrafficByDateCTCC"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/queryTrafficCTCC"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/queryUnitList"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/queryUnitUuidBytotalId"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/queryUserBtnPri"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/queryUserByUserIds"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/queryUserExtensionAttribute"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/queryUserForStep"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/queryUserList"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/queryUserProcessPri"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/queryUserWechatBindRel"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/queryUuidByTotalIdAndUuid"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/rechargeOrderSetMeal"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/renewSendReportConfirmEmail"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/reportList"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/saveCustomerEmployee"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/saveDevSimList"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/saveDeviceAccountBatchData"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/saveEnviromentIncomeInfos"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/saveEnvironmentCurve"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/saveFirmwareFile"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/saveIncomeSettingInfos"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/saveOrUpdateGroupStringCheckRule"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/saveParamModel"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/savePowerCharges"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/savePowerDevicePoint"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/savePowerRobotInfo"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/savePowerRobotSweepAttr"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/savePowerSettingCharges"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/savePowerSettingInfo"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/saveProductionBatchData"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/saveRechargeOrderObj"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/saveRechargeOrderOtherInfo"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/saveRepair"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/saveReportExportColumns"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/saveSetParam"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/saveSysUserMsg"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/saveTemplate"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/searchM2MMonthFlowCMCC"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/selectSysTranslationNames"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/sendPsTimeZoneInstruction"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/setUpFormulaFaultAnalyse"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/setUserGDPRAttrs"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/settingNotice"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/shareMyPs"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/sharePsBySN"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/showInverterByUnit"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/showOnlineUsers"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/showWarning"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/snIsExist"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/snsIsExist"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/speedyAddPowerStation"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/stationDeviceHistoryDataList"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/stationUnitsList"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/stationsDiscreteData"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/stationsIncomeList"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/stationsPointReport"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/stationsYearPlanReport"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/sureAndImportSelettlementData"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/sweepDevParamSet"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/sweepDevRunControl"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/sweepDevStrategyIssue"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/sysTimeZoneList"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/unLockUser"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/unlockChildAccount"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/updateCommunicationModuleState"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/updateDevInstalledPower"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/updateFault"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/updateFaultData"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/updateFaultMsgByFaultCode"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/updateFaultStatus"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/updateHouseholdWorkOrder"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/updateInverterSn2ModuleSn"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/updateOperateTicketAttachmentId"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/updateOrderDeviceByCustomerService"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/updateOwnerFaultConfig"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/updateParamSettingSysMsg"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/updatePlatformLevelFaultLevel"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/updatePowerDevicePoint"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/updatePowerRobotInfo"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/updatePowerRobotSweepAttr"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/updatePowerStationForHousehold"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/updatePowerStationInfo"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/updatePowerUserInfo"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/updateReportConfigByEmailAddr"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/updateShareAttr"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/updateSnIsSureFlag"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/updateStationPics"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/updateSysAdvancedParam"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/updateSysOrgNew"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/updateTemplate"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/updateUinfoNetEaseUser"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/updateUserExtensionAttribute"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/updateUserLanguage"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/updateUserPosition"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/updateUserUpOrg"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/upgrade"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/upgrate"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/uploadFileToOss"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/userAgreeGdprProtocol"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/userInfoUniqueCheck"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/userMailHasBound"
	"github.com/kevinkems/GoSungrow/iSolarCloud/AppService/userRegister"
	"github.com/kevinkems/GoSungrow/iSolarCloud/api"
	"github.com/kevinkems/GoSungrow/iSolarCloud/api/GoStruct/output"
	"fmt"
)

var _ api.Area = (*Area)(nil)

type Area api.AreaStruct

func init() {
	// name := api.GetArea(Area{})
	// fmt.Printf("Name: %s\n", name)
}

func Init(apiRoot api.Web) Area {
	area := Area{
		ApiRoot: apiRoot,
		Name:    api.GetArea(Area{}),
		EndPoints: api.TypeEndPoints{
			api.GetName(login.EndPoint{}):                       login.Init(apiRoot),
			api.GetName(findPsType.EndPoint{}):                  findPsType.Init(apiRoot),
			api.GetName(getPowerDevicePointNames.EndPoint{}):    getPowerDevicePointNames.Init(apiRoot),
			api.GetName(getPowerStatistics.EndPoint{}):          getPowerStatistics.Init(apiRoot),
			api.GetName(getPsDetailWithPsType.EndPoint{}):       getPsDetailWithPsType.Init(apiRoot),
			api.GetName(getPsList.EndPoint{}):                   getPsList.Init(apiRoot),
			api.GetName(getHouseholdStoragePsReport.EndPoint{}): getHouseholdStoragePsReport.Init(apiRoot),
			api.GetName(queryUnitList.EndPoint{}):               queryUnitList.Init(apiRoot),
			api.GetName(queryMutiPointDataList.EndPoint{}):      queryMutiPointDataList.Init(apiRoot),
			api.GetName(communicationModuleDetail.EndPoint{}):   communicationModuleDetail.Init(apiRoot),
			api.GetName(getHistoryInfo.EndPoint{}):              getHistoryInfo.Init(apiRoot), // @TODO - Not working.
			api.GetName(energyTrend.EndPoint{}):                 energyTrend.Init(apiRoot),
			api.GetName(getTemplateList.EndPoint{}):             getTemplateList.Init(apiRoot),
			api.GetName(getDevicePoints.EndPoint{}):             getDevicePoints.Init(apiRoot), // @TODO - Doesn't seem to generate anything.
			api.GetName(getPsDetail.EndPoint{}):                 getPsDetail.Init(apiRoot),
			api.GetName(getPsHealthState.EndPoint{}):            getPsHealthState.Init(apiRoot),
			api.GetName(queryAllPsIdAndName.EndPoint{}):         queryAllPsIdAndName.Init(apiRoot),
			api.GetName(getAllPowerDeviceSetName.EndPoint{}):    getAllPowerDeviceSetName.Init(apiRoot),
			api.GetName(getOrgListByName.EndPoint{}):            getOrgListByName.Init(apiRoot),
			api.GetName(getPsListByName.EndPoint{}):             getPsListByName.Init(apiRoot),
			api.GetName(checkUnitStatus.EndPoint{}):             checkUnitStatus.Init(apiRoot), // @TODO - Returns 404 error. Disabled.
			api.GetName(getDeviceInfo.EndPoint{}):               getDeviceInfo.Init(apiRoot),   // @TODO - Returns null.
			api.GetName(getDeviceList.EndPoint{}):               getDeviceList.Init(apiRoot),
			api.GetName(powerDevicePointList.EndPoint{}):        powerDevicePointList.Init(apiRoot),  // @TODO - CRITICAL - GetByJson all point_id definitions.
			api.GetName(queryDeviceListForApp.EndPoint{}):       queryDeviceListForApp.Init(apiRoot), // @TODO - CRITICAL - Show more detail info on devices.
			api.GetName(queryDeviceListByUserId.EndPoint{}):     queryDeviceListByUserId.Init(apiRoot),
			api.GetName(getPsListStaticData.EndPoint{}):         getPsListStaticData.Init(apiRoot),
			api.GetName(getPsReport.EndPoint{}):                 getPsReport.Init(apiRoot),
			api.GetName(getPsUser.EndPoint{}):                   getPsUser.Init(apiRoot),
			api.GetName(getPsWeatherList.EndPoint{}):            getPsWeatherList.Init(apiRoot),
			api.GetName(getAreaList.EndPoint{}):                 getAreaList.Init(apiRoot),
			api.GetName(getCloudList.EndPoint{}):                getCloudList.Init(apiRoot),
			api.GetName(getConfigList.EndPoint{}):               getConfigList.Init(apiRoot),          // @TODO - Returns 404 error. Disabled.
			api.GetName(getDeviceModelInfoList.EndPoint{}):      getDeviceModelInfoList.Init(apiRoot), // @TODO - CRITICAL - Show all device model info.
			api.GetName(getDeviceTypeList.EndPoint{}):           getDeviceTypeList.Init(apiRoot),
			api.GetName(getDeviceTypeInfoList.EndPoint{}):       getDeviceTypeInfoList.Init(apiRoot), // @TODO - CRITICAL - Show all device types.
			api.GetName(getInvertDataList.EndPoint{}):           getInvertDataList.Init(apiRoot),
			api.GetName(queryUserList.EndPoint{}):               queryUserList.Init(apiRoot),
			api.GetName(getPowerPictureList.EndPoint{}):         getPowerPictureList.Init(apiRoot),
			api.GetName(getUserList.EndPoint{}):                 getUserList.Init(apiRoot),
			api.GetName(getUserPsOrderList.EndPoint{}):          getUserPsOrderList.Init(apiRoot),
			api.GetName(reportList.EndPoint{}):                  reportList.Init(apiRoot),

			api.GetName(getKpiInfo.EndPoint{}): getKpiInfo.Init(apiRoot),
			api.GetName(getMapInfo.EndPoint{}): getMapInfo.Init(apiRoot), // @TODO - No data.

			// Tasks
			api.GetName(getPowerDeviceSetTaskDetailList.EndPoint{}): getPowerDeviceSetTaskDetailList.Init(apiRoot), // @TODO - Needs to be checked.
			api.GetName(getPowerDeviceSetTaskList.EndPoint{}):       getPowerDeviceSetTaskList.Init(apiRoot),
			api.GetName(getPsDataSupplementTaskList.EndPoint{}):     getPsDataSupplementTaskList.Init(apiRoot),
			api.GetName(queryBatchCreatePsTaskList.EndPoint{}):      queryBatchCreatePsTaskList.Init(apiRoot),
			api.GetName(getModuleLogTaskList.EndPoint{}):            getModuleLogTaskList.Init(apiRoot),
			api.GetName(getRemoteUpgradeTaskList.EndPoint{}):        getRemoteUpgradeTaskList.Init(apiRoot),     // @TODO - No data.
			api.GetName(getRemoteUpgradeSubTasksList.EndPoint{}):    getRemoteUpgradeSubTasksList.Init(apiRoot), // @TODO - No data.
			api.GetName(queryParamSettingTask.EndPoint{}):           queryParamSettingTask.Init(apiRoot),
			api.GetName(queryCtrlTaskById.EndPoint{}):               queryCtrlTaskById.Init(apiRoot), // @TODO - Returns 404 error. Disabled.

			// Use the following for all critical data every 5 minutes.
			api.GetName(queryDeviceList.EndPoint{}): queryDeviceList.Init(apiRoot), // @TODO - This gives you ALL info at 5 min intervals.

			// Disabled from here on.

			api.GetName(cancelParamSetTask.EndPoint{}):            cancelParamSetTask.Init(apiRoot),
			api.GetName(dealFailRemoteUpgradeSubTasks.EndPoint{}): dealFailRemoteUpgradeSubTasks.Init(apiRoot),
			api.GetName(dealFailRemoteUpgradeTasks.EndPoint{}):    dealFailRemoteUpgradeTasks.Init(apiRoot),

			api.GetName(updateTemplate.EndPoint{}):                  updateTemplate.Init(apiRoot),
			api.GetName(getTemplateByInfoType.EndPoint{}):           getTemplateByInfoType.Init(apiRoot),
			api.GetName(deleteTemplate.EndPoint{}):                  deleteTemplate.Init(apiRoot),
			api.GetName(getParamSetTemplate4NewProtocol.EndPoint{}): getParamSetTemplate4NewProtocol.Init(apiRoot),
			api.GetName(getParamSetTemplatePointInfo.EndPoint{}):    getParamSetTemplatePointInfo.Init(apiRoot),
			api.GetName(getSweepDevParamSetTemplate.EndPoint{}):     getSweepDevParamSetTemplate.Init(apiRoot),
			api.GetName(judgeDevIsHasInitSetTemplate.EndPoint{}):    judgeDevIsHasInitSetTemplate.Init(apiRoot),
			api.GetName(saveTemplate.EndPoint{}):                    saveTemplate.Init(apiRoot),

			api.GetName(queryDevicePointMinuteDataList.EndPoint{}):          queryDevicePointMinuteDataList.Init(apiRoot),
			api.GetName(getDevicePointMinuteDataList.EndPoint{}):            getDevicePointMinuteDataList.Init(apiRoot),
			api.GetName(queryDeviceInfo.EndPoint{}):                         queryDeviceInfo.Init(apiRoot),
			api.GetName(queryDeviceInfoForApp.EndPoint{}):                   queryDeviceInfoForApp.Init(apiRoot),
			api.GetName(queryDeviceModelTechnical.EndPoint{}):               queryDeviceModelTechnical.Init(apiRoot),
			api.GetName(queryDevicePointDayMonthYearDataList.EndPoint{}):    queryDevicePointDayMonthYearDataList.Init(apiRoot),
			api.GetName(queryDevicePointsDayMonthYearDataList.EndPoint{}):   queryDevicePointsDayMonthYearDataList.Init(apiRoot),
			api.GetName(queryDeviceRealTimeDataByPsKeys.EndPoint{}):         queryDeviceRealTimeDataByPsKeys.Init(apiRoot),
			api.GetName(addDeviceRepair.EndPoint{}):                         addDeviceRepair.Init(apiRoot),
			api.GetName(addDeviceToStructureForHousehold.EndPoint{}):        addDeviceToStructureForHousehold.Init(apiRoot),
			api.GetName(addDeviceToStructureForHouseholdByPsIdS.EndPoint{}): addDeviceToStructureForHouseholdByPsIdS.Init(apiRoot),
			api.GetName(batchModifyDevicesInfoAndPropertis.EndPoint{}):      batchModifyDevicesInfoAndPropertis.Init(apiRoot),
			api.GetName(batchUpdateDeviceSim.EndPoint{}):                    batchUpdateDeviceSim.Init(apiRoot),
			api.GetName(caculateDeviceInputDiscrete.EndPoint{}):             caculateDeviceInputDiscrete.Init(apiRoot),
			api.GetName(calculateDeviceDiscrete.EndPoint{}):                 calculateDeviceDiscrete.Init(apiRoot),
			api.GetName(checkUpRechargeDevicePaying.EndPoint{}):             checkUpRechargeDevicePaying.Init(apiRoot),
			api.GetName(dealSysDeviceSimFlowInfo.EndPoint{}):                dealSysDeviceSimFlowInfo.Init(apiRoot),
			api.GetName(dealSysDeviceSimInfo.EndPoint{}):                    dealSysDeviceSimInfo.Init(apiRoot),
			api.GetName(delDeviceRepair.EndPoint{}):                         delDeviceRepair.Init(apiRoot),
			api.GetName(deleteDeviceAccount.EndPoint{}):                     deleteDeviceAccount.Init(apiRoot),
			api.GetName(deleteDeviceSimById.EndPoint{}):                     deleteDeviceSimById.Init(apiRoot),
			api.GetName(deletePowerDevicePointById.EndPoint{}):              deletePowerDevicePointById.Init(apiRoot),
			api.GetName(editDeviceRepair.EndPoint{}):                        editDeviceRepair.Init(apiRoot),
			api.GetName(getAllDeviceByPsId.EndPoint{}):                      getAllDeviceByPsId.Init(apiRoot),
			api.GetName(getCommunicationDeviceConfigInfo.EndPoint{}):        getCommunicationDeviceConfigInfo.Init(apiRoot),
			api.GetName(getDeviceAccountById.EndPoint{}):                    getDeviceAccountById.Init(apiRoot),
			api.GetName(getDeviceFaultStatisticsData.EndPoint{}):            getDeviceFaultStatisticsData.Init(apiRoot),
			api.GetName(getDevicePropertys.EndPoint{}):                      getDevicePropertys.Init(apiRoot),
			api.GetName(getDeviceRepairDetail.EndPoint{}):                   getDeviceRepairDetail.Init(apiRoot),
			api.GetName(getDeviceTechBranchCount.EndPoint{}):                getDeviceTechBranchCount.Init(apiRoot),
			api.GetName(getPowerDeviceModelTechList.EndPoint{}):             getPowerDeviceModelTechList.Init(apiRoot),
			api.GetName(getPowerDeviceModelTree.EndPoint{}):                 getPowerDeviceModelTree.Init(apiRoot),
			api.GetName(getPowerDevicePointInfo.EndPoint{}):                 getPowerDevicePointInfo.Init(apiRoot),
			api.GetName(getRechargeOrderItemDeviceList.EndPoint{}):          getRechargeOrderItemDeviceList.Init(apiRoot),
			api.GetName(getRemoteUpgradeDeviceList.EndPoint{}):              getRemoteUpgradeDeviceList.Init(apiRoot),
			api.GetName(insightSynDeviceStructure2Cloud.EndPoint{}):         insightSynDeviceStructure2Cloud.Init(apiRoot),
			api.GetName(modifiedDeviceInfo.EndPoint{}):                      modifiedDeviceInfo.Init(apiRoot),
			api.GetName(queryDeviceRepairList.EndPoint{}):                   queryDeviceRepairList.Init(apiRoot),
			api.GetName(queryDeviceTypeInfoList.EndPoint{}):                 queryDeviceTypeInfoList.Init(apiRoot),
			api.GetName(queryFaultTypeByDevice.EndPoint{}):                  queryFaultTypeByDevice.Init(apiRoot),
			api.GetName(queryFaultTypeByDevicePage.EndPoint{}):              queryFaultTypeByDevicePage.Init(apiRoot),
			api.GetName(saveDeviceAccountBatchData.EndPoint{}):              saveDeviceAccountBatchData.Init(apiRoot),
			api.GetName(savePowerDevicePoint.EndPoint{}):                    savePowerDevicePoint.Init(apiRoot),
			api.GetName(stationDeviceHistoryDataList.EndPoint{}):            stationDeviceHistoryDataList.Init(apiRoot),
			api.GetName(updateOrderDeviceByCustomerService.EndPoint{}):      updateOrderDeviceByCustomerService.Init(apiRoot),
			api.GetName(updatePowerDevicePoint.EndPoint{}):                  updatePowerDevicePoint.Init(apiRoot),

			api.GetName(psHourPointsValue.EndPoint{}):                 psHourPointsValue.Init(apiRoot),
			api.GetName(getPsCurveInfo.EndPoint{}):                    getPsCurveInfo.Init(apiRoot),
			api.GetName(exportPlantReportPDF.EndPoint{}):              exportPlantReportPDF.Init(apiRoot),
			api.GetName(getTableDataSql.EndPoint{}):                   getTableDataSql.Init(apiRoot),
			api.GetName(getTableDataSqlCount.EndPoint{}):              getTableDataSqlCount.Init(apiRoot),
			api.GetName(getPListinfoFromMysql.EndPoint{}):             getPListinfoFromMysql.Init(apiRoot),
			api.GetName(getDataFromHBase.EndPoint{}):                  getDataFromHBase.Init(apiRoot),
			api.GetName(getListMiFromHBase.EndPoint{}):                getListMiFromHBase.Init(apiRoot),
			api.GetName(getMapMiFromHBase.EndPoint{}):                 getMapMiFromHBase.Init(apiRoot),
			api.GetName(getValFromHBase.EndPoint{}):                   getValFromHBase.Init(apiRoot),
			api.GetName(getPowerStationTableDataSql.EndPoint{}):       getPowerStationTableDataSql.Init(apiRoot),
			api.GetName(getPowerStationTableDataSqlCount.EndPoint{}):  getPowerStationTableDataSqlCount.Init(apiRoot),
			api.GetName(getDataFromHbaseByRowKey.EndPoint{}):          getDataFromHbaseByRowKey.Init(apiRoot),
			api.GetName(findInfoByuuid.EndPoint{}):                    findInfoByuuid.Init(apiRoot),
			api.GetName(getStationInfoSql.EndPoint{}):                 getStationInfoSql.Init(apiRoot),
			api.GetName(getPowerTrendDayData.EndPoint{}):              getPowerTrendDayData.Init(apiRoot),
			api.GetName(powerTrendChartData.EndPoint{}):               powerTrendChartData.Init(apiRoot),
			api.GetName(getUpTimePoint.EndPoint{}):                    getUpTimePoint.Init(apiRoot),
			api.GetName(getSerialNum.EndPoint{}):                      getSerialNum.Init(apiRoot),
			api.GetName(getPsAuthKey.EndPoint{}):                      getPsAuthKey.Init(apiRoot),
			api.GetName(getApiCallsForAppkeys.EndPoint{}):             getApiCallsForAppkeys.Init(apiRoot),
			api.GetName(exportParamSettingValPDF.EndPoint{}):          exportParamSettingValPDF.Init(apiRoot),
			api.GetName(devicePointsDataFromMySql.EndPoint{}):         devicePointsDataFromMySql.Init(apiRoot),
			api.GetName(acceptPsSharing.EndPoint{}):                   acceptPsSharing.Init(apiRoot),
			api.GetName(activateEmail.EndPoint{}):                     activateEmail.Init(apiRoot),
			api.GetName(addConfig.EndPoint{}):                         addConfig.Init(apiRoot),
			api.GetName(addFault.EndPoint{}):                          addFault.Init(apiRoot),
			api.GetName(addFaultOrder.EndPoint{}):                     addFaultOrder.Init(apiRoot),
			api.GetName(addFaultPlan.EndPoint{}):                      addFaultPlan.Init(apiRoot),
			api.GetName(addFaultRepairSteps.EndPoint{}):               addFaultRepairSteps.Init(apiRoot),
			api.GetName(addHouseholdEvaluation.EndPoint{}):            addHouseholdEvaluation.Init(apiRoot),
			api.GetName(addHouseholdLeaveMessage.EndPoint{}):          addHouseholdLeaveMessage.Init(apiRoot),
			api.GetName(addHouseholdOpinionFeedback.EndPoint{}):       addHouseholdOpinionFeedback.Init(apiRoot),
			api.GetName(addHouseholdWorkOrder.EndPoint{}):             addHouseholdWorkOrder.Init(apiRoot),
			api.GetName(addOnDutyInfo.EndPoint{}):                     addOnDutyInfo.Init(apiRoot),
			api.GetName(addOperRule.EndPoint{}):                       addOperRule.Init(apiRoot),
			api.GetName(addOrDelPsStructure.EndPoint{}):               addOrDelPsStructure.Init(apiRoot),
			api.GetName(addOrderStep.EndPoint{}):                      addOrderStep.Init(apiRoot),
			api.GetName(addPowerStationForHousehold.EndPoint{}):       addPowerStationForHousehold.Init(apiRoot),
			api.GetName(addPowerStationInfo.EndPoint{}):               addPowerStationInfo.Init(apiRoot),
			api.GetName(addReportConfigEmail.EndPoint{}):              addReportConfigEmail.Init(apiRoot),
			api.GetName(addSysAdvancedParam.EndPoint{}):               addSysAdvancedParam.Init(apiRoot),
			api.GetName(addSysOrgNew.EndPoint{}):                      addSysOrgNew.Init(apiRoot),
			api.GetName(aliPayAppTest.EndPoint{}):                     aliPayAppTest.Init(apiRoot),
			api.GetName(auditOperRule.EndPoint{}):                     auditOperRule.Init(apiRoot),
			api.GetName(batchAddStationBySn.EndPoint{}):               batchAddStationBySn.Init(apiRoot),
			api.GetName(batchImportSN.EndPoint{}):                     batchImportSN.Init(apiRoot),
			api.GetName(batchInsertUserAndOrg.EndPoint{}):             batchInsertUserAndOrg.Init(apiRoot),
			api.GetName(batchProcessPlantReport.EndPoint{}):           batchProcessPlantReport.Init(apiRoot),
			api.GetName(batchUpdateUserIsAgreeGdpr.EndPoint{}):        batchUpdateUserIsAgreeGdpr.Init(apiRoot),
			api.GetName(boundMobilePhone.EndPoint{}):                  boundMobilePhone.Init(apiRoot),
			api.GetName(boundUserMail.EndPoint{}):                     boundUserMail.Init(apiRoot),
			api.GetName(calculateInitialCompensationData.EndPoint{}):  calculateInitialCompensationData.Init(apiRoot),
			api.GetName(cancelDeliverMail.EndPoint{}):                 cancelDeliverMail.Init(apiRoot),
			api.GetName(cancelOrderScan.EndPoint{}):                   cancelOrderScan.Init(apiRoot),
			api.GetName(cancelPsSharing.EndPoint{}):                   cancelPsSharing.Init(apiRoot),
			api.GetName(cancelRechargeOrder.EndPoint{}):               cancelRechargeOrder.Init(apiRoot),
			api.GetName(changRechargeOrderToCancel.EndPoint{}):        changRechargeOrderToCancel.Init(apiRoot),
			api.GetName(changeHouseholdUser2Installer.EndPoint{}):     changeHouseholdUser2Installer.Init(apiRoot),
			api.GetName(changeRemoteParam.EndPoint{}):                 changeRemoteParam.Init(apiRoot),
			api.GetName(checkDealerOrgCode.EndPoint{}):                checkDealerOrgCode.Init(apiRoot),
			api.GetName(checkDevSnIsBelongsToUser.EndPoint{}):         checkDevSnIsBelongsToUser.Init(apiRoot),
			api.GetName(checkInverterResult.EndPoint{}):               checkInverterResult.Init(apiRoot),
			api.GetName(checkIsCanDoParamSet.EndPoint{}):              checkIsCanDoParamSet.Init(apiRoot),
			api.GetName(checkIsIvScan.EndPoint{}):                     checkIsIvScan.Init(apiRoot),
			api.GetName(checkOssObjectExist.EndPoint{}):               checkOssObjectExist.Init(apiRoot),
			api.GetName(checkServiceIsConnect.EndPoint{}):             checkServiceIsConnect.Init(apiRoot),
			api.GetName(checkTechnicalParameters.EndPoint{}):          checkTechnicalParameters.Init(apiRoot),
			api.GetName(checkUserAccountUnique.EndPoint{}):            checkUserAccountUnique.Init(apiRoot),
			api.GetName(checkUserAccountUniqueAll.EndPoint{}):         checkUserAccountUniqueAll.Init(apiRoot),
			api.GetName(checkUserInfoUnique.EndPoint{}):               checkUserInfoUnique.Init(apiRoot),
			api.GetName(checkUserIsExist.EndPoint{}):                  checkUserIsExist.Init(apiRoot),
			api.GetName(checkUserListIsExist.EndPoint{}):              checkUserListIsExist.Init(apiRoot),
			api.GetName(checkUserPassword.EndPoint{}):                 checkUserPassword.Init(apiRoot),
			api.GetName(cloudDeploymentRecord.EndPoint{}):             cloudDeploymentRecord.Init(apiRoot),
			api.GetName(comfirmParamModel.EndPoint{}):                 comfirmParamModel.Init(apiRoot),
			api.GetName(compareValidateCode.EndPoint{}):               compareValidateCode.Init(apiRoot),
			api.GetName(componentInfo2Cloud.EndPoint{}):               componentInfo2Cloud.Init(apiRoot),
			api.GetName(confirmFault.EndPoint{}):                      confirmFault.Init(apiRoot),
			api.GetName(confirmIvFault.EndPoint{}):                    confirmIvFault.Init(apiRoot),
			api.GetName(confirmReportConfig.EndPoint{}):               confirmReportConfig.Init(apiRoot),
			api.GetName(createAppkeyInfo.EndPoint{}):                  createAppkeyInfo.Init(apiRoot),
			api.GetName(createRenewInvoice.EndPoint{}):                createRenewInvoice.Init(apiRoot),
			api.GetName(dealCommandReply.EndPoint{}):                  dealCommandReply.Init(apiRoot),
			api.GetName(dealDeletePsFailPsDelete.EndPoint{}):          dealDeletePsFailPsDelete.Init(apiRoot),
			api.GetName(dealFaultOrder.EndPoint{}):                    dealFaultOrder.Init(apiRoot),
			api.GetName(dealGroupStringDisableOrEnable.EndPoint{}):    dealGroupStringDisableOrEnable.Init(apiRoot),
			api.GetName(dealNumberOfServiceCalls2Mysql.EndPoint{}):    dealNumberOfServiceCalls2Mysql.Init(apiRoot),
			api.GetName(dealParamSettingAfterComplete.EndPoint{}):     dealParamSettingAfterComplete.Init(apiRoot),
			api.GetName(dealPsDataSupplement.EndPoint{}):              dealPsDataSupplement.Init(apiRoot),
			api.GetName(dealPsReportEmailSend.EndPoint{}):             dealPsReportEmailSend.Init(apiRoot),
			api.GetName(dealRemoteUpgrade.EndPoint{}):                 dealRemoteUpgrade.Init(apiRoot),
			api.GetName(dealSnElectrifyCheck.EndPoint{}):              dealSnElectrifyCheck.Init(apiRoot),
			api.GetName(definiteTimeDealSnExpRemind.EndPoint{}):       definiteTimeDealSnExpRemind.Init(apiRoot),
			api.GetName(definiteTimeDealSnStatus.EndPoint{}):          definiteTimeDealSnStatus.Init(apiRoot),
			api.GetName(delOperRule.EndPoint{}):                       delOperRule.Init(apiRoot),
			api.GetName(delayCallApiResidueTimes.EndPoint{}):          delayCallApiResidueTimes.Init(apiRoot),
			api.GetName(deleteComponent.EndPoint{}):                   deleteComponent.Init(apiRoot),
			api.GetName(deleteCustomerEmployee.EndPoint{}):            deleteCustomerEmployee.Init(apiRoot),
			api.GetName(deleteElectricitySettlementData.EndPoint{}):   deleteElectricitySettlementData.Init(apiRoot),
			api.GetName(deleteFaultPlan.EndPoint{}):                   deleteFaultPlan.Init(apiRoot),
			api.GetName(deleteFirmwareFiles.EndPoint{}):               deleteFirmwareFiles.Init(apiRoot),
			api.GetName(deleteHouseholdEvaluation.EndPoint{}):         deleteHouseholdEvaluation.Init(apiRoot),
			api.GetName(deleteHouseholdLeaveMessage.EndPoint{}):       deleteHouseholdLeaveMessage.Init(apiRoot),
			api.GetName(deleteHouseholdWorkOrder.EndPoint{}):          deleteHouseholdWorkOrder.Init(apiRoot),
			api.GetName(deleteInverterSnInChnnl.EndPoint{}):           deleteInverterSnInChnnl.Init(apiRoot),
			api.GetName(deleteModuleLog.EndPoint{}):                   deleteModuleLog.Init(apiRoot),
			api.GetName(deleteOnDutyInfo.EndPoint{}):                  deleteOnDutyInfo.Init(apiRoot),
			api.GetName(deleteOperateBillFile.EndPoint{}):             deleteOperateBillFile.Init(apiRoot),
			api.GetName(deleteOssObject.EndPoint{}):                   deleteOssObject.Init(apiRoot),
			api.GetName(deletePowerPicture.EndPoint{}):                deletePowerPicture.Init(apiRoot),
			api.GetName(deletePowerRobotInfoBySnAndPsId.EndPoint{}):   deletePowerRobotInfoBySnAndPsId.Init(apiRoot),
			api.GetName(deletePowerRobotSweepStrategy.EndPoint{}):     deletePowerRobotSweepStrategy.Init(apiRoot),
			api.GetName(deleteProductionData.EndPoint{}):              deleteProductionData.Init(apiRoot),
			api.GetName(deletePs.EndPoint{}):                          deletePs.Init(apiRoot),
			api.GetName(deleteRechargeOrder.EndPoint{}):               deleteRechargeOrder.Init(apiRoot),
			api.GetName(deleteRegularlyConnectionInfo.EndPoint{}):     deleteRegularlyConnectionInfo.Init(apiRoot),
			api.GetName(deleteReportConfigEmailAddr.EndPoint{}):       deleteReportConfigEmailAddr.Init(apiRoot),
			api.GetName(deleteSysAdvancedParam.EndPoint{}):            deleteSysAdvancedParam.Init(apiRoot),
			api.GetName(deleteSysOrgNew.EndPoint{}):                   deleteSysOrgNew.Init(apiRoot),
			api.GetName(deleteUserInfoAllByUserId.EndPoint{}):         deleteUserInfoAllByUserId.Init(apiRoot),
			api.GetName(deviceInputDiscreteDeleteTime.EndPoint{}):     deviceInputDiscreteDeleteTime.Init(apiRoot),
			api.GetName(deviceInputDiscreteGetTime.EndPoint{}):        deviceInputDiscreteGetTime.Init(apiRoot),
			api.GetName(deviceInputDiscreteInsertTime.EndPoint{}):     deviceInputDiscreteInsertTime.Init(apiRoot),
			api.GetName(deviceInputDiscreteUpdateTime.EndPoint{}):     deviceInputDiscreteUpdateTime.Init(apiRoot),
			api.GetName(deviceReplace.EndPoint{}):                     deviceReplace.Init(apiRoot),
			api.GetName(editOperRule.EndPoint{}):                      editOperRule.Init(apiRoot),
			api.GetName(energyPovertyAlleviation.EndPoint{}):          energyPovertyAlleviation.Init(apiRoot),
			api.GetName(faultAutoClose.EndPoint{}):                    faultAutoClose.Init(apiRoot),
			api.GetName(faultCloseRemindOrderHandler.EndPoint{}):      faultCloseRemindOrderHandler.Init(apiRoot),
			api.GetName(findCodeValueList.EndPoint{}):                 findCodeValueList.Init(apiRoot),
			api.GetName(findEmgOrgInfo.EndPoint{}):                    findEmgOrgInfo.Init(apiRoot),
			api.GetName(findEnvironmentInfo.EndPoint{}):               findEnvironmentInfo.Init(apiRoot),
			api.GetName(findFromHbaseAndRedis.EndPoint{}):             findFromHbaseAndRedis.Init(apiRoot),
			api.GetName(findLossAnalysisList.EndPoint{}):              findLossAnalysisList.Init(apiRoot),
			api.GetName(findOnDutyInfo.EndPoint{}):                    findOnDutyInfo.Init(apiRoot),
			api.GetName(findSingleStationPR.EndPoint{}):               findSingleStationPR.Init(apiRoot),
			api.GetName(findUserPassword.EndPoint{}):                  findUserPassword.Init(apiRoot),
			api.GetName(genTLSUserSigByUserAccount.EndPoint{}):        genTLSUserSigByUserAccount.Init(apiRoot),
			api.GetName(generateRandomPassword.EndPoint{}):            generateRandomPassword.Init(apiRoot),
			api.GetName(getAPIServiceInfo.EndPoint{}):                 getAPIServiceInfo.Init(apiRoot),
			api.GetName(getAccessedPermission.EndPoint{}):             getAccessedPermission.Init(apiRoot),
			api.GetName(getAllPowerRobotViewInfoByPsId.EndPoint{}):    getAllPowerRobotViewInfoByPsId.Init(apiRoot),
			api.GetName(getAllPsIdByOrgIds.EndPoint{}):                getAllPsIdByOrgIds.Init(apiRoot),
			api.GetName(getAllUserRemindCount.EndPoint{}):             getAllUserRemindCount.Init(apiRoot),
			api.GetName(getAndOutletsAndUnit.EndPoint{}):              getAndOutletsAndUnit.Init(apiRoot),
			api.GetName(getAreaInfoCodeByCounty.EndPoint{}):           getAreaInfoCodeByCounty.Init(apiRoot),
			api.GetName(getAutoCreatePowerStation.EndPoint{}):         getAutoCreatePowerStation.Init(apiRoot),
			api.GetName(getBackReadValue.EndPoint{}):                  getBackReadValue.Init(apiRoot),
			api.GetName(getBatchNewestPointData.EndPoint{}):           getBatchNewestPointData.Init(apiRoot),
			api.GetName(getCallApiResidueTimes.EndPoint{}):            getCallApiResidueTimes.Init(apiRoot),
			api.GetName(getChangedPsListByTime.EndPoint{}):            getChangedPsListByTime.Init(apiRoot),
			api.GetName(getChnnlListByPsId.EndPoint{}):                getChnnlListByPsId.Init(apiRoot),
			api.GetName(getCloudServiceMappingConfig.EndPoint{}):      getCloudServiceMappingConfig.Init(apiRoot),
			api.GetName(getCommunicationModuleMonitorData.EndPoint{}): getCommunicationModuleMonitorData.Init(apiRoot),
			api.GetName(getComponentModelFactory.EndPoint{}):          getComponentModelFactory.Init(apiRoot),

			api.GetName(getConnectionInfoBySnAndLocalPort.EndPoint{}):                 getConnectionInfoBySnAndLocalPort.Init(apiRoot),
			api.GetName(getCountDown.EndPoint{}):                                      getCountDown.Init(apiRoot),
			api.GetName(getCountryServiceInfo.EndPoint{}):                             getCountryServiceInfo.Init(apiRoot),
			api.GetName(getCounty.EndPoint{}):                                         getCounty.Init(apiRoot),
			api.GetName(getCustomerEmployee.EndPoint{}):                               getCustomerEmployee.Init(apiRoot),
			api.GetName(getCustomerList.EndPoint{}):                                   getCustomerList.Init(apiRoot),
			api.GetName(getDevInstalledPowerByPsId.EndPoint{}):                        getDevInstalledPowerByPsId.Init(apiRoot),
			api.GetName(getDevRecord.EndPoint{}):                                      getDevRecord.Init(apiRoot),
			api.GetName(getDevRunRecordList.EndPoint{}):                               getDevRunRecordList.Init(apiRoot),
			api.GetName(getDevSimByList.EndPoint{}):                                   getDevSimByList.Init(apiRoot),
			api.GetName(getDevSimList.EndPoint{}):                                     getDevSimList.Init(apiRoot),
			api.GetName(getDstInfo.EndPoint{}):                                        getDstInfo.Init(apiRoot),
			api.GetName(getElectricitySettlementData.EndPoint{}):                      getElectricitySettlementData.Init(apiRoot),
			api.GetName(getElectricitySettlementDetailData.EndPoint{}):                getElectricitySettlementDetailData.Init(apiRoot),
			api.GetName(getEncryptPublicKey.EndPoint{}):                               getEncryptPublicKey.Init(apiRoot),
			api.GetName(getFaultCount.EndPoint{}):                                     getFaultCount.Init(apiRoot),
			api.GetName(getFaultDetail.EndPoint{}):                                    getFaultDetail.Init(apiRoot),
			api.GetName(getFaultMsgByFaultCode.EndPoint{}):                            getFaultMsgByFaultCode.Init(apiRoot),
			api.GetName(getFaultMsgListByPageWithYYYYMM.EndPoint{}):                   getFaultMsgListByPageWithYYYYMM.Init(apiRoot),
			api.GetName(getFaultMsgListWithYYYYMM.EndPoint{}):                         getFaultMsgListWithYYYYMM.Init(apiRoot),
			api.GetName(getFaultPlanList.EndPoint{}):                                  getFaultPlanList.Init(apiRoot),
			api.GetName(getFileOperationRecordOne.EndPoint{}):                         getFileOperationRecordOne.Init(apiRoot),
			api.GetName(getFormulaFaultAnalyseList.EndPoint{}):                        getFormulaFaultAnalyseList.Init(apiRoot),
			api.GetName(getGroupStringCheckResult.EndPoint{}):                         getGroupStringCheckResult.Init(apiRoot),
			api.GetName(getGroupStringCheckRule.EndPoint{}):                           getGroupStringCheckRule.Init(apiRoot),
			api.GetName(getHisData.EndPoint{}):                                        getHisData.Init(apiRoot),
			api.GetName(getHouseholdEvaluation.EndPoint{}):                            getHouseholdEvaluation.Init(apiRoot),
			api.GetName(getHouseholdLeaveMessage.EndPoint{}):                          getHouseholdLeaveMessage.Init(apiRoot),
			api.GetName(getHouseholdOpinionFeedback.EndPoint{}):                       getHouseholdOpinionFeedback.Init(apiRoot),
			api.GetName(getHouseholdPsInstallerByUserId.EndPoint{}):                   getHouseholdPsInstallerByUserId.Init(apiRoot),
			api.GetName(getHouseholdUserInfo.EndPoint{}):                              getHouseholdUserInfo.Init(apiRoot),
			api.GetName(getHouseholdWorkOrderInfo.EndPoint{}):                         getHouseholdWorkOrderInfo.Init(apiRoot),
			api.GetName(getHouseholdWorkOrderList.EndPoint{}):                         getHouseholdWorkOrderList.Init(apiRoot),
			api.GetName(getI18nConfigByType.EndPoint{}):                               getI18nConfigByType.Init(apiRoot),
			api.GetName(getI18nFileInfo.EndPoint{}):                                   getI18nFileInfo.Init(apiRoot),
			api.GetName(getI18nInfoByKey.EndPoint{}):                                  getI18nInfoByKey.Init(apiRoot),
			api.GetName(getI18nVersion.EndPoint{}):                                    getI18nVersion.Init(apiRoot),
			api.GetName(getIncomeSettingInfos.EndPoint{}):                             getIncomeSettingInfos.Init(apiRoot),
			api.GetName(getInfoFromAMap.EndPoint{}):                                   getInfoFromAMap.Init(apiRoot),
			api.GetName(getInfomationFromRedis.EndPoint{}):                            getInfomationFromRedis.Init(apiRoot),
			api.GetName(getInstallInfoList.EndPoint{}):                                getInstallInfoList.Init(apiRoot),
			api.GetName(getInstallerInfoByDealerOrgCodeOrId.EndPoint{}):               getInstallerInfoByDealerOrgCodeOrId.Init(apiRoot),
			api.GetName(getInverterDataCount.EndPoint{}):                              getInverterDataCount.Init(apiRoot),
			api.GetName(getInverterProcess.EndPoint{}):                                getInverterProcess.Init(apiRoot),
			api.GetName(getInverterUuidBytotalId.EndPoint{}):                          getInverterUuidBytotalId.Init(apiRoot),
			api.GetName(getIvEchartsData.EndPoint{}):                                  getIvEchartsData.Init(apiRoot),
			api.GetName(getIvEchartsDataById.EndPoint{}):                              getIvEchartsDataById.Init(apiRoot),
			api.GetName(getMenuAndPrivileges.EndPoint{}):                              getMenuAndPrivileges.Init(apiRoot),
			api.GetName(getMicrogridEStoragePsReport.EndPoint{}):                      getMicrogridEStoragePsReport.Init(apiRoot),
			api.GetName(getModuleLogInfo.EndPoint{}):                                  getModuleLogInfo.Init(apiRoot),
			api.GetName(getNationProvJSON.EndPoint{}):                                 getNationProvJSON.Init(apiRoot),
			api.GetName(getNeedOpAsynOpRecordList.EndPoint{}):                         getNeedOpAsynOpRecordList.Init(apiRoot),
			api.GetName(getNoticeInfo.EndPoint{}):                                     getNoticeInfo.Init(apiRoot),
			api.GetName(getNumberOfServiceCalls.EndPoint{}):                           getNumberOfServiceCalls.Init(apiRoot),
			api.GetName(getOSSConfig.EndPoint{}):                                      getOSSConfig.Init(apiRoot),
			api.GetName(getOperRuleDetail.EndPoint{}):                                 getOperRuleDetail.Init(apiRoot),
			api.GetName(getOperateBillFileId.EndPoint{}):                              getOperateBillFileId.Init(apiRoot),
			api.GetName(getOperateTicketForDetail.EndPoint{}):                         getOperateTicketForDetail.Init(apiRoot),
			api.GetName(getOrCreateNetEaseUserToken.EndPoint{}):                       getOrCreateNetEaseUserToken.Init(apiRoot),
			api.GetName(getOrderDataList.EndPoint{}):                                  getOrderDataList.Init(apiRoot),
			api.GetName(getOrderDataSql2.EndPoint{}):                                  getOrderDataSql2.Init(apiRoot),
			api.GetName(getOrderDatas.EndPoint{}):                                     getOrderDatas.Init(apiRoot),
			api.GetName(getOrderDetail.EndPoint{}):                                    getOrderDetail.Init(apiRoot),
			api.GetName(getOrderStatistics.EndPoint{}):                                getOrderStatistics.Init(apiRoot),
			api.GetName(getOrgIdNameByUserId.EndPoint{}):                              getOrgIdNameByUserId.Init(apiRoot),
			api.GetName(getOrgInfoByDealerOrgCode.EndPoint{}):                         getOrgInfoByDealerOrgCode.Init(apiRoot),
			api.GetName(getOrgListByUserId.EndPoint{}):                                getOrgListByUserId.Init(apiRoot),
			api.GetName(getOrgListForUser.EndPoint{}):                                 getOrgListForUser.Init(apiRoot),
			api.GetName(getOssObjectStream.EndPoint{}):                                getOssObjectStream.Init(apiRoot),
			api.GetName(getOwnerFaultConfigList.EndPoint{}):                           getOwnerFaultConfigList.Init(apiRoot),
			api.GetName(getParamterSettingBase.EndPoint{}):                            getParamterSettingBase.Init(apiRoot),
			api.GetName(getPhotoInfo.EndPoint{}):                                      getPhotoInfo.Init(apiRoot),
			api.GetName(getPlanedOrNotPsList.EndPoint{}):                              getPlanedOrNotPsList.Init(apiRoot),
			api.GetName(getPlantReportPDFList.EndPoint{}):                             getPlantReportPDFList.Init(apiRoot),
			api.GetName(getPowerChargeSettingInfo.EndPoint{}):                         getPowerChargeSettingInfo.Init(apiRoot),
			api.GetName(getPowerFormulaFaultAnalyse.EndPoint{}):                       getPowerFormulaFaultAnalyse.Init(apiRoot),
			api.GetName(getPowerRobotInfoByRobotSn.EndPoint{}):                        getPowerRobotInfoByRobotSn.Init(apiRoot),
			api.GetName(getPowerRobotSweepAttrByPsId.EndPoint{}):                      getPowerRobotSweepAttrByPsId.Init(apiRoot),
			api.GetName(getPowerRobotSweepStrategy.EndPoint{}):                        getPowerRobotSweepStrategy.Init(apiRoot),
			api.GetName(getPowerRobotSweepStrategyList.EndPoint{}):                    getPowerRobotSweepStrategyList.Init(apiRoot),
			api.GetName(getPowerSettingCharges.EndPoint{}):                            getPowerSettingCharges.Init(apiRoot),
			api.GetName(getPowerSettingHistoryRecords.EndPoint{}):                     getPowerSettingHistoryRecords.Init(apiRoot),
			api.GetName(getPowerStationBasicInfo.EndPoint{}):                          getPowerStationBasicInfo.Init(apiRoot),
			api.GetName(getPowerStationData.EndPoint{}):                               getPowerStationData.Init(apiRoot),
			api.GetName(getPowerStationForHousehold.EndPoint{}):                       getPowerStationForHousehold.Init(apiRoot),
			api.GetName(getPowerStationInfo.EndPoint{}):                               getPowerStationInfo.Init(apiRoot),
			api.GetName(getPowerStationPR.EndPoint{}):                                 getPowerStationPR.Init(apiRoot),
			api.GetName(getPrivateCloudValidityPeriod.EndPoint{}):                     getPrivateCloudValidityPeriod.Init(apiRoot),
			api.GetName(getProvInfoListByNationCode.EndPoint{}):                       getProvInfoListByNationCode.Init(apiRoot),
			api.GetName(getPsDetailByUserTokens.EndPoint{}):                           getPsDetailByUserTokens.Init(apiRoot),
			api.GetName(getPsDetailForSinglePage.EndPoint{}):                          getPsDetailForSinglePage.Init(apiRoot),
			api.GetName(getPsInstallerByPsId.EndPoint{}):                              getPsInstallerByPsId.Init(apiRoot),
			api.GetName(getPsInstallerOrgInfoByPsId.EndPoint{}):                       getPsInstallerOrgInfoByPsId.Init(apiRoot),
			api.GetName(getPsListForPsDataByPsId.EndPoint{}):                          getPsListForPsDataByPsId.Init(apiRoot),
			api.GetName(getRechargeOrderDetail.EndPoint{}):                            getRechargeOrderDetail.Init(apiRoot),
			api.GetName(getRechargeOrderList.EndPoint{}):                              getRechargeOrderList.Init(apiRoot),
			api.GetName(getRegionalTree.EndPoint{}):                                   getRegionalTree.Init(apiRoot),
			api.GetName(getRemoteParamSettingList.EndPoint{}):                         getRemoteParamSettingList.Init(apiRoot),
			api.GetName(getRemoteUpgradeScheduleDetails.EndPoint{}):                   getRemoteUpgradeScheduleDetails.Init(apiRoot),
			api.GetName(getReportData.EndPoint{}):                                     getReportData.Init(apiRoot),
			api.GetName(getReportEmailConfigInfo.EndPoint{}):                          getReportEmailConfigInfo.Init(apiRoot),
			api.GetName(getReportExportColumns.EndPoint{}):                            getReportExportColumns.Init(apiRoot),
			api.GetName(getReportListByUserId.EndPoint{}):                             getReportListByUserId.Init(apiRoot),
			api.GetName(getRobotDynamicCleaningView.EndPoint{}):                       getRobotDynamicCleaningView.Init(apiRoot),
			api.GetName(getRobotNumAndSweepCapacity.EndPoint{}):                       getRobotNumAndSweepCapacity.Init(apiRoot),
			api.GetName(getRuleUnit.EndPoint{}):                                       getRuleUnit.Init(apiRoot),
			api.GetName(getSendReportConfigCron.EndPoint{}):                           getSendReportConfigCron.Init(apiRoot),
			api.GetName(getShieldMapConditionList.EndPoint{}):                         getShieldMapConditionList.Init(apiRoot),
			api.GetName(getSimIdBySnList.EndPoint{}):                                  getSimIdBySnList.Init(apiRoot),
			api.GetName(getSingleIVData.EndPoint{}):                                   getSingleIVData.Init(apiRoot),
			api.GetName(getSnChangeRecord.EndPoint{}):                                 getSnChangeRecord.Init(apiRoot),
			api.GetName(getSnConnectionInfo.EndPoint{}):                               getSnConnectionInfo.Init(apiRoot),
			api.GetName(getSungwsConfigCache.EndPoint{}):                              getSungwsConfigCache.Init(apiRoot),
			api.GetName(getSungwsGlobalConfigCache.EndPoint{}):                        getSungwsGlobalConfigCache.Init(apiRoot),
			api.GetName(getSweepRobotDevList.EndPoint{}):                              getSweepRobotDevList.Init(apiRoot),
			api.GetName(getSysMsg.EndPoint{}):                                         getSysMsg.Init(apiRoot),
			api.GetName(getSysOrgNewList.EndPoint{}):                                  getSysOrgNewList.Init(apiRoot),
			api.GetName(getSysOrgNewOne.EndPoint{}):                                   getSysOrgNewOne.Init(apiRoot),
			api.GetName(getSysUserById.EndPoint{}):                                    getSysUserById.Init(apiRoot),
			api.GetName(getUUIDByUpuuid.EndPoint{}):                                   getUUIDByUpuuid.Init(apiRoot),
			api.GetName(getUserById.EndPoint{}):                                       getUserById.Init(apiRoot),
			api.GetName(getUserByInstaller.EndPoint{}):                                getUserByInstaller.Init(apiRoot),
			api.GetName(getUserDevOnlineOffineCount.EndPoint{}):                       getUserDevOnlineOffineCount.Init(apiRoot),
			api.GetName(getUserGDPRAttrs.EndPoint{}):                                  getUserGDPRAttrs.Init(apiRoot),
			api.GetName(getUserHavePowerStationCount.EndPoint{}):                      getUserHavePowerStationCount.Init(apiRoot),
			api.GetName(getUserInfoByUserAccounts.EndPoint{}):                         getUserInfoByUserAccounts.Init(apiRoot),
			api.GetName(getValidateCode.EndPoint{}):                                   getValidateCode.Init(apiRoot),
			api.GetName(getValidateCodeAtRegister.EndPoint{}):                         getValidateCodeAtRegister.Init(apiRoot),
			api.GetName(getWeatherInfo.EndPoint{}):                                    getWeatherInfo.Init(apiRoot),
			api.GetName(getWechatPushConfig.EndPoint{}):                               getWechatPushConfig.Init(apiRoot),
			api.GetName(getWorkInfo.EndPoint{}):                                       getWorkInfo.Init(apiRoot),
			api.GetName(groupStringCheck.EndPoint{}):                                  groupStringCheck.Init(apiRoot),
			api.GetName(handleDevByCommunicationSN.EndPoint{}):                        handleDevByCommunicationSN.Init(apiRoot),
			api.GetName(householdResetPassBySN.EndPoint{}):                            householdResetPassBySN.Init(apiRoot),
			api.GetName(immediatePayment.EndPoint{}):                                  immediatePayment.Init(apiRoot),
			api.GetName(importExcelData.EndPoint{}):                                   importExcelData.Init(apiRoot),
			api.GetName(incomeStatistics.EndPoint{}):                                  incomeStatistics.Init(apiRoot),
			api.GetName(informPush.EndPoint{}):                                        informPush.Init(apiRoot),
			api.GetName(insertEmgOrgInfo.EndPoint{}):                                  insertEmgOrgInfo.Init(apiRoot),
			api.GetName(intoDataToHbase.EndPoint{}):                                   intoDataToHbase.Init(apiRoot),
			api.GetName(ipLocationQuery.EndPoint{}):                                   ipLocationQuery.Init(apiRoot),
			api.GetName(isHave2GSn.EndPoint{}):                                        isHave2GSn.Init(apiRoot),
			api.GetName(judgeIsSettingMan.EndPoint{}):                                 judgeIsSettingMan.Init(apiRoot),
			api.GetName(listOssFiles.EndPoint{}):                                      listOssFiles.Init(apiRoot),
			api.GetName(loadAreaInfo.EndPoint{}):                                      loadAreaInfo.Init(apiRoot),
			api.GetName(loadPowerStation.EndPoint{}):                                  loadPowerStation.Init(apiRoot),
			api.GetName(loginByToken.EndPoint{}):                                      loginByToken.Init(apiRoot),
			api.GetName(logout.EndPoint{}):                                            logout.Init(apiRoot),
			api.GetName(mobilePhoneHasBound.EndPoint{}):                               mobilePhoneHasBound.Init(apiRoot),
			api.GetName(modifyEmgOrgStruc.EndPoint{}):                                 modifyEmgOrgStruc.Init(apiRoot),
			api.GetName(modifyFaultPlan.EndPoint{}):                                   modifyFaultPlan.Init(apiRoot),
			api.GetName(modifyOnDutyInfo.EndPoint{}):                                  modifyOnDutyInfo.Init(apiRoot),
			api.GetName(modifyPassword.EndPoint{}):                                    modifyPassword.Init(apiRoot),
			api.GetName(modifyPersonalUnitList.EndPoint{}):                            modifyPersonalUnitList.Init(apiRoot),
			api.GetName(modifyPsUser.EndPoint{}):                                      modifyPsUser.Init(apiRoot),
			api.GetName(moduleLogParamSet.EndPoint{}):                                 moduleLogParamSet.Init(apiRoot),
			api.GetName(operateOssFile.EndPoint{}):                                    operateOssFile.Init(apiRoot),
			api.GetName(operationPowerRobotSweepStrategy.EndPoint{}):                  operationPowerRobotSweepStrategy.Init(apiRoot),
			api.GetName(orgPowerReport.EndPoint{}):                                    orgPowerReport.Init(apiRoot),
			api.GetName(paramSetTryAgain.EndPoint{}):                                  paramSetTryAgain.Init(apiRoot),
			api.GetName(paramSetting.EndPoint{}):                                      paramSetting.Init(apiRoot),
			api.GetName(planPower.EndPoint{}):                                         planPower.Init(apiRoot),
			api.GetName(psForcastInfo.EndPoint{}):                                     psForcastInfo.Init(apiRoot),
			api.GetName(queryBatchSpeedyAddPowerStationResult.EndPoint{}):             queryBatchSpeedyAddPowerStationResult.Init(apiRoot),
			api.GetName(queryCardStatusCTCC.EndPoint{}):                               queryCardStatusCTCC.Init(apiRoot),
			api.GetName(queryChildAccountList.EndPoint{}):                             queryChildAccountList.Init(apiRoot),
			api.GetName(queryCompensationRecordData.EndPoint{}):                       queryCompensationRecordData.Init(apiRoot),
			api.GetName(queryCompensationRecordList.EndPoint{}):                       queryCompensationRecordList.Init(apiRoot),
			api.GetName(queryComponent.EndPoint{}):                                    queryComponent.Init(apiRoot),
			api.GetName(queryComponentTechnicalParam.EndPoint{}):                      queryComponentTechnicalParam.Init(apiRoot),
			api.GetName(queryCountryGridAndRelation.EndPoint{}):                       queryCountryGridAndRelation.Init(apiRoot),
			api.GetName(queryCountryList.EndPoint{}):                                  queryCountryList.Init(apiRoot),
			api.GetName(queryEnvironmentList.EndPoint{}):                              queryEnvironmentList.Init(apiRoot),
			api.GetName(queryFaultList.EndPoint{}):                                    queryFaultList.Init(apiRoot),
			api.GetName(queryFaultPlanDetail.EndPoint{}):                              queryFaultPlanDetail.Init(apiRoot),
			api.GetName(queryFaultRepairSteps.EndPoint{}):                             queryFaultRepairSteps.Init(apiRoot),
			api.GetName(queryFaultTypeAndLevelByCode.EndPoint{}):                      queryFaultTypeAndLevelByCode.Init(apiRoot),
			api.GetName(queryFirmwareFilesPage.EndPoint{}):                            queryFirmwareFilesPage.Init(apiRoot),
			api.GetName(queryInfotoAlert.EndPoint{}):                                  queryInfotoAlert.Init(apiRoot),
			api.GetName(queryInverterModelList.EndPoint{}):                            queryInverterModelList.Init(apiRoot),
			api.GetName(queryInverterVersionList.EndPoint{}):                          queryInverterVersionList.Init(apiRoot),
			api.GetName(queryM2MCardInfoCMCC.EndPoint{}):                              queryM2MCardInfoCMCC.Init(apiRoot),
			api.GetName(queryM2MCardTermInfoCMCC.EndPoint{}):                          queryM2MCardTermInfoCMCC.Init(apiRoot),
			api.GetName(queryModelInfoByModelId.EndPoint{}):                           queryModelInfoByModelId.Init(apiRoot),
			api.GetName(queryNoticeList.EndPoint{}):                                   queryNoticeList.Init(apiRoot),
			api.GetName(queryNumberOfRenewalReminders.EndPoint{}):                     queryNumberOfRenewalReminders.Init(apiRoot),
			api.GetName(queryOperRules.EndPoint{}):                                    queryOperRules.Init(apiRoot),
			api.GetName(queryOrderList.EndPoint{}):                                    queryOrderList.Init(apiRoot),
			api.GetName(queryOrderStep.EndPoint{}):                                    queryOrderStep.Init(apiRoot),
			api.GetName(queryOrgGenerationReport.EndPoint{}):                          queryOrgGenerationReport.Init(apiRoot),
			api.GetName(queryOrgInfoList.EndPoint{}):                                  queryOrgInfoList.Init(apiRoot),
			api.GetName(queryOrgPowerElecPercent.EndPoint{}):                          queryOrgPowerElecPercent.Init(apiRoot),
			api.GetName(queryOrgPsCompensationRecordList.EndPoint{}):                  queryOrgPsCompensationRecordList.Init(apiRoot),
			api.GetName(queryPersonalUnitList.EndPoint{}):                             queryPersonalUnitList.Init(apiRoot),
			api.GetName(queryPointDataTopOne.EndPoint{}):                              queryPointDataTopOne.Init(apiRoot),
			api.GetName(queryPowerStationInfo.EndPoint{}):                             queryPowerStationInfo.Init(apiRoot),
			api.GetName(queryPsAreaByUserIdAndAreaCode.EndPoint{}):                    queryPsAreaByUserIdAndAreaCode.Init(apiRoot),
			api.GetName(queryPsCompensationRecordList.EndPoint{}):                     queryPsCompensationRecordList.Init(apiRoot),
			api.GetName(queryPsDataByDate.EndPoint{}):                                 queryPsDataByDate.Init(apiRoot),
			api.GetName(queryPsIdList.EndPoint{}):                                     queryPsIdList.Init(apiRoot),
			api.GetName(queryPsListByUserIdAndAreaCode.EndPoint{}):                    queryPsListByUserIdAndAreaCode.Init(apiRoot),
			api.GetName(queryPsNameByPsId.EndPoint{}):                                 queryPsNameByPsId.Init(apiRoot),
			api.GetName(queryPsPrByDate.EndPoint{}):                                   queryPsPrByDate.Init(apiRoot),
			api.GetName(queryPsProfit.EndPoint{}):                                     queryPsProfit.Init(apiRoot),
			api.GetName(queryPsReportComparativeAnalysisOfPowerGeneration.EndPoint{}): queryPsReportComparativeAnalysisOfPowerGeneration.Init(apiRoot),
			api.GetName(queryPsStructureList.EndPoint{}):                              queryPsStructureList.Init(apiRoot),
			api.GetName(queryPuuidsByCommandTotalId.EndPoint{}):                       queryPuuidsByCommandTotalId.Init(apiRoot),
			api.GetName(queryPuuidsByCommandTotalId2.EndPoint{}):                      queryPuuidsByCommandTotalId2.Init(apiRoot),
			api.GetName(queryRepairRuleList.EndPoint{}):                               queryRepairRuleList.Init(apiRoot),
			api.GetName(queryReportListForManagementPage.EndPoint{}):                  queryReportListForManagementPage.Init(apiRoot),
			api.GetName(queryReportMsg.EndPoint{}):                                    queryReportMsg.Init(apiRoot),
			api.GetName(querySharingPs.EndPoint{}):                                    querySharingPs.Init(apiRoot),
			api.GetName(querySysAdvancedParam.EndPoint{}):                             querySysAdvancedParam.Init(apiRoot),
			api.GetName(queryTimeBySN.EndPoint{}):                                     queryTimeBySN.Init(apiRoot),
			api.GetName(queryTrafficByDateCTCC.EndPoint{}):                            queryTrafficByDateCTCC.Init(apiRoot),
			api.GetName(queryTrafficCTCC.EndPoint{}):                                  queryTrafficCTCC.Init(apiRoot),
			api.GetName(queryUnitUuidBytotalId.EndPoint{}):                            queryUnitUuidBytotalId.Init(apiRoot),
			api.GetName(queryUserBtnPri.EndPoint{}):                                   queryUserBtnPri.Init(apiRoot),
			api.GetName(queryUserByUserIds.EndPoint{}):                                queryUserByUserIds.Init(apiRoot),
			api.GetName(queryUserExtensionAttribute.EndPoint{}):                       queryUserExtensionAttribute.Init(apiRoot),
			api.GetName(queryUserForStep.EndPoint{}):                                  queryUserForStep.Init(apiRoot),
			api.GetName(queryUserProcessPri.EndPoint{}):                               queryUserProcessPri.Init(apiRoot),
			api.GetName(queryUserWechatBindRel.EndPoint{}):                            queryUserWechatBindRel.Init(apiRoot),
			api.GetName(queryUuidByTotalIdAndUuid.EndPoint{}):                         queryUuidByTotalIdAndUuid.Init(apiRoot),
			api.GetName(rechargeOrderSetMeal.EndPoint{}):                              rechargeOrderSetMeal.Init(apiRoot),
			api.GetName(renewSendReportConfirmEmail.EndPoint{}):                       renewSendReportConfirmEmail.Init(apiRoot),
			api.GetName(saveCustomerEmployee.EndPoint{}):                              saveCustomerEmployee.Init(apiRoot),
			api.GetName(saveDevSimList.EndPoint{}):                                    saveDevSimList.Init(apiRoot),
			api.GetName(saveEnviromentIncomeInfos.EndPoint{}):                         saveEnviromentIncomeInfos.Init(apiRoot),
			api.GetName(saveEnvironmentCurve.EndPoint{}):                              saveEnvironmentCurve.Init(apiRoot),
			api.GetName(saveFirmwareFile.EndPoint{}):                                  saveFirmwareFile.Init(apiRoot),
			api.GetName(saveIncomeSettingInfos.EndPoint{}):                            saveIncomeSettingInfos.Init(apiRoot),
			api.GetName(saveOrUpdateGroupStringCheckRule.EndPoint{}):                  saveOrUpdateGroupStringCheckRule.Init(apiRoot),
			api.GetName(saveParamModel.EndPoint{}):                                    saveParamModel.Init(apiRoot),
			api.GetName(savePowerCharges.EndPoint{}):                                  savePowerCharges.Init(apiRoot),
			api.GetName(savePowerRobotInfo.EndPoint{}):                                savePowerRobotInfo.Init(apiRoot),
			api.GetName(savePowerRobotSweepAttr.EndPoint{}):                           savePowerRobotSweepAttr.Init(apiRoot),
			api.GetName(savePowerSettingCharges.EndPoint{}):                           savePowerSettingCharges.Init(apiRoot),
			api.GetName(savePowerSettingInfo.EndPoint{}):                              savePowerSettingInfo.Init(apiRoot),
			api.GetName(saveProductionBatchData.EndPoint{}):                           saveProductionBatchData.Init(apiRoot),
			api.GetName(saveRechargeOrderObj.EndPoint{}):                              saveRechargeOrderObj.Init(apiRoot),
			api.GetName(saveRechargeOrderOtherInfo.EndPoint{}):                        saveRechargeOrderOtherInfo.Init(apiRoot),
			api.GetName(saveRepair.EndPoint{}):                                        saveRepair.Init(apiRoot),
			api.GetName(saveReportExportColumns.EndPoint{}):                           saveReportExportColumns.Init(apiRoot),
			api.GetName(saveSetParam.EndPoint{}):                                      saveSetParam.Init(apiRoot),
			api.GetName(saveSysUserMsg.EndPoint{}):                                    saveSysUserMsg.Init(apiRoot),
			api.GetName(searchM2MMonthFlowCMCC.EndPoint{}):                            searchM2MMonthFlowCMCC.Init(apiRoot),
			api.GetName(selectSysTranslationNames.EndPoint{}):                         selectSysTranslationNames.Init(apiRoot),
			api.GetName(sendPsTimeZoneInstruction.EndPoint{}):                         sendPsTimeZoneInstruction.Init(apiRoot),
			api.GetName(setUpFormulaFaultAnalyse.EndPoint{}):                          setUpFormulaFaultAnalyse.Init(apiRoot),
			api.GetName(setUserGDPRAttrs.EndPoint{}):                                  setUserGDPRAttrs.Init(apiRoot),
			api.GetName(settingNotice.EndPoint{}):                                     settingNotice.Init(apiRoot),
			api.GetName(shareMyPs.EndPoint{}):                                         shareMyPs.Init(apiRoot),
			api.GetName(sharePsBySN.EndPoint{}):                                       sharePsBySN.Init(apiRoot),
			api.GetName(showInverterByUnit.EndPoint{}):                                showInverterByUnit.Init(apiRoot),
			api.GetName(showOnlineUsers.EndPoint{}):                                   showOnlineUsers.Init(apiRoot),
			api.GetName(showWarning.EndPoint{}):                                       showWarning.Init(apiRoot),
			api.GetName(snIsExist.EndPoint{}):                                         snIsExist.Init(apiRoot),
			api.GetName(snsIsExist.EndPoint{}):                                        snsIsExist.Init(apiRoot),
			api.GetName(speedyAddPowerStation.EndPoint{}):                             speedyAddPowerStation.Init(apiRoot),
			api.GetName(stationUnitsList.EndPoint{}):                                  stationUnitsList.Init(apiRoot),
			api.GetName(stationsDiscreteData.EndPoint{}):                              stationsDiscreteData.Init(apiRoot),
			api.GetName(stationsIncomeList.EndPoint{}):                                stationsIncomeList.Init(apiRoot),
			api.GetName(stationsPointReport.EndPoint{}):                               stationsPointReport.Init(apiRoot),
			api.GetName(stationsYearPlanReport.EndPoint{}):                            stationsYearPlanReport.Init(apiRoot),
			api.GetName(sureAndImportSelettlementData.EndPoint{}):                     sureAndImportSelettlementData.Init(apiRoot),
			api.GetName(sweepDevParamSet.EndPoint{}):                                  sweepDevParamSet.Init(apiRoot),
			api.GetName(sweepDevRunControl.EndPoint{}):                                sweepDevRunControl.Init(apiRoot),
			api.GetName(sweepDevStrategyIssue.EndPoint{}):                             sweepDevStrategyIssue.Init(apiRoot),
			api.GetName(sysTimeZoneList.EndPoint{}):                                   sysTimeZoneList.Init(apiRoot),
			api.GetName(unLockUser.EndPoint{}):                                        unLockUser.Init(apiRoot),
			api.GetName(unlockChildAccount.EndPoint{}):                                unlockChildAccount.Init(apiRoot),
			api.GetName(updateCommunicationModuleState.EndPoint{}):                    updateCommunicationModuleState.Init(apiRoot),
			api.GetName(updateDevInstalledPower.EndPoint{}):                           updateDevInstalledPower.Init(apiRoot),
			api.GetName(updateFault.EndPoint{}):                                       updateFault.Init(apiRoot),
			api.GetName(updateFaultData.EndPoint{}):                                   updateFaultData.Init(apiRoot),
			api.GetName(updateFaultMsgByFaultCode.EndPoint{}):                         updateFaultMsgByFaultCode.Init(apiRoot),
			api.GetName(updateFaultStatus.EndPoint{}):                                 updateFaultStatus.Init(apiRoot),
			api.GetName(updateHouseholdWorkOrder.EndPoint{}):                          updateHouseholdWorkOrder.Init(apiRoot),
			api.GetName(updateInverterSn2ModuleSn.EndPoint{}):                         updateInverterSn2ModuleSn.Init(apiRoot),
			api.GetName(updateOperateTicketAttachmentId.EndPoint{}):                   updateOperateTicketAttachmentId.Init(apiRoot),
			api.GetName(updateOwnerFaultConfig.EndPoint{}):                            updateOwnerFaultConfig.Init(apiRoot),
			api.GetName(updateParamSettingSysMsg.EndPoint{}):                          updateParamSettingSysMsg.Init(apiRoot),
			api.GetName(updatePlatformLevelFaultLevel.EndPoint{}):                     updatePlatformLevelFaultLevel.Init(apiRoot),
			api.GetName(updatePowerRobotInfo.EndPoint{}):                              updatePowerRobotInfo.Init(apiRoot),
			api.GetName(updatePowerRobotSweepAttr.EndPoint{}):                         updatePowerRobotSweepAttr.Init(apiRoot),
			api.GetName(updatePowerStationForHousehold.EndPoint{}):                    updatePowerStationForHousehold.Init(apiRoot),
			api.GetName(updatePowerStationInfo.EndPoint{}):                            updatePowerStationInfo.Init(apiRoot),
			api.GetName(updatePowerUserInfo.EndPoint{}):                               updatePowerUserInfo.Init(apiRoot),
			api.GetName(updateReportConfigByEmailAddr.EndPoint{}):                     updateReportConfigByEmailAddr.Init(apiRoot),
			api.GetName(updateShareAttr.EndPoint{}):                                   updateShareAttr.Init(apiRoot),
			api.GetName(updateSnIsSureFlag.EndPoint{}):                                updateSnIsSureFlag.Init(apiRoot),
			api.GetName(updateStationPics.EndPoint{}):                                 updateStationPics.Init(apiRoot),
			api.GetName(updateSysAdvancedParam.EndPoint{}):                            updateSysAdvancedParam.Init(apiRoot),
			api.GetName(updateSysOrgNew.EndPoint{}):                                   updateSysOrgNew.Init(apiRoot),
			api.GetName(updateUinfoNetEaseUser.EndPoint{}):                            updateUinfoNetEaseUser.Init(apiRoot),
			api.GetName(updateUserExtensionAttribute.EndPoint{}):                      updateUserExtensionAttribute.Init(apiRoot),
			api.GetName(updateUserLanguage.EndPoint{}):                                updateUserLanguage.Init(apiRoot),
			api.GetName(updateUserPosition.EndPoint{}):                                updateUserPosition.Init(apiRoot),
			api.GetName(updateUserUpOrg.EndPoint{}):                                   updateUserUpOrg.Init(apiRoot),
			api.GetName(upgrade.EndPoint{}):                                           upgrade.Init(apiRoot),
			api.GetName(upgrate.EndPoint{}):                                           upgrate.Init(apiRoot),
			api.GetName(uploadFileToOss.EndPoint{}):                                   uploadFileToOss.Init(apiRoot),
			api.GetName(userAgreeGdprProtocol.EndPoint{}):                             userAgreeGdprProtocol.Init(apiRoot),
			api.GetName(userInfoUniqueCheck.EndPoint{}):                               userInfoUniqueCheck.Init(apiRoot),
			api.GetName(userMailHasBound.EndPoint{}):                                  userMailHasBound.Init(apiRoot),
			api.GetName(userRegister.EndPoint{}):                                      userRegister.Init(apiRoot),
		},
	}

	return area
}

// ****************************************
// Methods not scoped by api.EndPoint interface type

func GetAreaName() string {
	return string(api.GetArea(Area{}))
}

func (a Area) GetEndPoint(name api.EndPointName) api.EndPoint {
	var ret api.EndPoint
	for _, e := range a.EndPoints {
		// fmt.Printf("endpoint: %v\n", e)
		if e.GetName() == name {
			ret = e
			break
		}
	}
	return ret
}

// ****************************************
// Methods scoped by api.Area interface type

func (a Area) Init(apiRoot *api.Web) api.AreaStruct {
	panic("implement me")
}

func (a Area) GetAreaName() api.AreaName {
	return a.Name
}

func (a Area) GetEndPoints() api.TypeEndPoints {
	for _, endpoint := range a.EndPoints {
		fmt.Printf("endpoint: %v\n", endpoint)
	}
	return a.EndPoints
}

func (a Area) Call(name api.EndPointName) output.Json {
	panic("implement me")
}

func (a Area) SetRequest(name api.EndPointName, ref interface{}) error {
	panic("implement me")
}

func (a Area) GetRequest(name api.EndPointName) output.Json {
	panic("implement me")
}

func (a Area) GetResponse(name api.EndPointName) output.Json {
	panic("implement me")
}

func (a Area) GetData(name api.EndPointName) output.Json {
	panic("implement me")
}

func (a Area) IsValid(name api.EndPointName) error {
	panic("implement me")
}

func (a Area) GetError(name api.EndPointName) error {
	panic("implement me")
}
