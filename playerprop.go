package excel

//go:generate enumer --json --type=PlayerProp $GOFILE
type PlayerProp uint32

const (
	PROP_NONE                            PlayerProp = 0
	PROP_EXP                             PlayerProp = 1001
	PROP_BREAK_LEVEL                     PlayerProp = 1002
	PROP_SATIATION_VAL                   PlayerProp = 1003
	PROP_SATIATION_PENALTY_TIME          PlayerProp = 1004
	PROP_LEVEL                           PlayerProp = 4001
	PROP_LAST_CHANGE_AVATAR_TIME         PlayerProp = 10001
	PROP_MAX_SPRING_VOLUME               PlayerProp = 10002
	PROP_CUR_SPRING_VOLUME               PlayerProp = 10003
	PROP_IS_SPRING_AUTO_USE              PlayerProp = 10004
	PROP_SPRING_AUTO_USE_PERCENT         PlayerProp = 10005
	PROP_IS_FLYABLE                      PlayerProp = 10006
	PROP_IS_WEATHER_LOCKED               PlayerProp = 10007
	PROP_IS_GAME_TIME_LOCKED             PlayerProp = 10008
	PROP_IS_TRANSFERABLE                 PlayerProp = 10009
	PROP_MAX_STAMINA                     PlayerProp = 10010
	PROP_CUR_PERSIST_STAMINA             PlayerProp = 10011
	PROP_CUR_TEMPORARY_STAMINA           PlayerProp = 10012
	PROP_PLAYER_LEVEL                    PlayerProp = 10013
	PROP_PLAYER_EXP                      PlayerProp = 10014
	PROP_PLAYER_HCOIN                    PlayerProp = 10015
	PROP_PLAYER_SCOIN                    PlayerProp = 10016
	PROP_PLAYER_MP_SETTING_TYPE          PlayerProp = 10017
	PROP_IS_MP_MODE_AVAILABLE            PlayerProp = 10018
	PROP_PLAYER_WORLD_LEVEL              PlayerProp = 10019
	PROP_PLAYER_RESIN                    PlayerProp = 10020
	PROP_PLAYER_WAIT_SUB_HCOIN           PlayerProp = 10022
	PROP_PLAYER_WAIT_SUB_SCOIN           PlayerProp = 10023
	PROP_IS_ONLY_MP_WITH_PS_PLAYER       PlayerProp = 10024
	PROP_PLAYER_MCOIN                    PlayerProp = 10025
	PROP_PLAYER_WAIT_SUB_MCOIN           PlayerProp = 10026
	PROP_PLAYER_LEGENDARY_KEY            PlayerProp = 10027
	PROP_IS_HAS_FIRST_SHARE              PlayerProp = 10028
	PROP_PLAYER_FORGE_POINT              PlayerProp = 10029
	PROP_CUR_CLIMATE_METER               PlayerProp = 10035
	PROP_CUR_CLIMATE_TYPE                PlayerProp = 10036
	PROP_CUR_CLIMATE_AREA_ID             PlayerProp = 10037
	PROP_CUR_CLIMATE_AREA_CLIMATE_TYPE   PlayerProp = 10038
	PROP_PLAYER_WORLD_LEVEL_LIMIT        PlayerProp = 10039
	PROP_PLAYER_WORLD_LEVEL_ADJUST_CD    PlayerProp = 10040
	PROP_PLAYER_LEGENDARY_DAILY_TASK_NUM PlayerProp = 10041
	PROP_PLAYER_HOME_COIN                PlayerProp = 10042
	PROP_PLAYER_WAIT_SUB_HOME_COIN       PlayerProp = 10043
	PROP_IS_AUTO_UNLOCK_SPECIFIC_EQUIP   PlayerProp = 10044
	PROP_PLAYER_GCG_COIN                 PlayerProp = 10045
	PROP_PLAYER_WAIT_SUB_GCG_COIN        PlayerProp = 10046
	PROP_PLAYER_ONLINE_TIME              PlayerProp = 10047
	PROP_PLAYER_CAN_DIVE                 PlayerProp = 10048
	PROP_DIVE_MAX_STAMINA                PlayerProp = 10049
	PROP_DIVE_CUR_STAMINA                PlayerProp = 10050
)
