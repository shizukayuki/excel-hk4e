// Code generated by "enumer --json --type=PlayerProp playerprop.go"; DO NOT EDIT.

package excel

import (
	"encoding/json"
	"fmt"
	"strings"
)

const (
	_PlayerPropName_0      = "PROP_NONE"
	_PlayerPropLowerName_0 = "prop_none"
	_PlayerPropName_1      = "PROP_EXPPROP_BREAK_LEVELPROP_SATIATION_VALPROP_SATIATION_PENALTY_TIME"
	_PlayerPropLowerName_1 = "prop_expprop_break_levelprop_satiation_valprop_satiation_penalty_time"
	_PlayerPropName_2      = "PROP_LEVEL"
	_PlayerPropLowerName_2 = "prop_level"
	_PlayerPropName_3      = "PROP_LAST_CHANGE_AVATAR_TIMEPROP_MAX_SPRING_VOLUMEPROP_CUR_SPRING_VOLUMEPROP_IS_SPRING_AUTO_USEPROP_SPRING_AUTO_USE_PERCENTPROP_IS_FLYABLEPROP_IS_WEATHER_LOCKEDPROP_IS_GAME_TIME_LOCKEDPROP_IS_TRANSFERABLEPROP_MAX_STAMINAPROP_CUR_PERSIST_STAMINAPROP_CUR_TEMPORARY_STAMINAPROP_PLAYER_LEVELPROP_PLAYER_EXPPROP_PLAYER_HCOINPROP_PLAYER_SCOINPROP_PLAYER_MP_SETTING_TYPEPROP_IS_MP_MODE_AVAILABLEPROP_PLAYER_WORLD_LEVELPROP_PLAYER_RESIN"
	_PlayerPropLowerName_3 = "prop_last_change_avatar_timeprop_max_spring_volumeprop_cur_spring_volumeprop_is_spring_auto_useprop_spring_auto_use_percentprop_is_flyableprop_is_weather_lockedprop_is_game_time_lockedprop_is_transferableprop_max_staminaprop_cur_persist_staminaprop_cur_temporary_staminaprop_player_levelprop_player_expprop_player_hcoinprop_player_scoinprop_player_mp_setting_typeprop_is_mp_mode_availableprop_player_world_levelprop_player_resin"
	_PlayerPropName_4      = "PROP_PLAYER_WAIT_SUB_HCOINPROP_PLAYER_WAIT_SUB_SCOINPROP_IS_ONLY_MP_WITH_PS_PLAYERPROP_PLAYER_MCOINPROP_PLAYER_WAIT_SUB_MCOINPROP_PLAYER_LEGENDARY_KEYPROP_IS_HAS_FIRST_SHAREPROP_PLAYER_FORGE_POINT"
	_PlayerPropLowerName_4 = "prop_player_wait_sub_hcoinprop_player_wait_sub_scoinprop_is_only_mp_with_ps_playerprop_player_mcoinprop_player_wait_sub_mcoinprop_player_legendary_keyprop_is_has_first_shareprop_player_forge_point"
	_PlayerPropName_5      = "PROP_CUR_CLIMATE_METERPROP_CUR_CLIMATE_TYPEPROP_CUR_CLIMATE_AREA_IDPROP_CUR_CLIMATE_AREA_CLIMATE_TYPEPROP_PLAYER_WORLD_LEVEL_LIMITPROP_PLAYER_WORLD_LEVEL_ADJUST_CDPROP_PLAYER_LEGENDARY_DAILY_TASK_NUMPROP_PLAYER_HOME_COINPROP_PLAYER_WAIT_SUB_HOME_COINPROP_IS_AUTO_UNLOCK_SPECIFIC_EQUIPPROP_PLAYER_GCG_COINPROP_PLAYER_WAIT_SUB_GCG_COINPROP_PLAYER_ONLINE_TIMEPROP_PLAYER_CAN_DIVEPROP_DIVE_MAX_STAMINAPROP_DIVE_CUR_STAMINA"
	_PlayerPropLowerName_5 = "prop_cur_climate_meterprop_cur_climate_typeprop_cur_climate_area_idprop_cur_climate_area_climate_typeprop_player_world_level_limitprop_player_world_level_adjust_cdprop_player_legendary_daily_task_numprop_player_home_coinprop_player_wait_sub_home_coinprop_is_auto_unlock_specific_equipprop_player_gcg_coinprop_player_wait_sub_gcg_coinprop_player_online_timeprop_player_can_diveprop_dive_max_staminaprop_dive_cur_stamina"
)

var (
	_PlayerPropIndex_0 = [...]uint8{0, 9}
	_PlayerPropIndex_1 = [...]uint8{0, 8, 24, 42, 69}
	_PlayerPropIndex_2 = [...]uint8{0, 10}
	_PlayerPropIndex_3 = [...]uint16{0, 28, 50, 72, 95, 123, 138, 160, 184, 204, 220, 244, 270, 287, 302, 319, 336, 363, 388, 411, 428}
	_PlayerPropIndex_4 = [...]uint8{0, 26, 52, 82, 99, 125, 150, 173, 196}
	_PlayerPropIndex_5 = [...]uint16{0, 22, 43, 67, 101, 130, 163, 199, 220, 250, 284, 304, 333, 356, 376, 397, 418}
)

func (i PlayerProp) String() string {
	switch {
	case i == 0:
		return _PlayerPropName_0
	case 1001 <= i && i <= 1004:
		i -= 1001
		return _PlayerPropName_1[_PlayerPropIndex_1[i]:_PlayerPropIndex_1[i+1]]
	case i == 4001:
		return _PlayerPropName_2
	case 10001 <= i && i <= 10020:
		i -= 10001
		return _PlayerPropName_3[_PlayerPropIndex_3[i]:_PlayerPropIndex_3[i+1]]
	case 10022 <= i && i <= 10029:
		i -= 10022
		return _PlayerPropName_4[_PlayerPropIndex_4[i]:_PlayerPropIndex_4[i+1]]
	case 10035 <= i && i <= 10050:
		i -= 10035
		return _PlayerPropName_5[_PlayerPropIndex_5[i]:_PlayerPropIndex_5[i+1]]
	default:
		return fmt.Sprintf("PlayerProp(%d)", i)
	}
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _PlayerPropNoOp() {
	var x [1]struct{}
	_ = x[PROP_NONE-(0)]
	_ = x[PROP_EXP-(1001)]
	_ = x[PROP_BREAK_LEVEL-(1002)]
	_ = x[PROP_SATIATION_VAL-(1003)]
	_ = x[PROP_SATIATION_PENALTY_TIME-(1004)]
	_ = x[PROP_LEVEL-(4001)]
	_ = x[PROP_LAST_CHANGE_AVATAR_TIME-(10001)]
	_ = x[PROP_MAX_SPRING_VOLUME-(10002)]
	_ = x[PROP_CUR_SPRING_VOLUME-(10003)]
	_ = x[PROP_IS_SPRING_AUTO_USE-(10004)]
	_ = x[PROP_SPRING_AUTO_USE_PERCENT-(10005)]
	_ = x[PROP_IS_FLYABLE-(10006)]
	_ = x[PROP_IS_WEATHER_LOCKED-(10007)]
	_ = x[PROP_IS_GAME_TIME_LOCKED-(10008)]
	_ = x[PROP_IS_TRANSFERABLE-(10009)]
	_ = x[PROP_MAX_STAMINA-(10010)]
	_ = x[PROP_CUR_PERSIST_STAMINA-(10011)]
	_ = x[PROP_CUR_TEMPORARY_STAMINA-(10012)]
	_ = x[PROP_PLAYER_LEVEL-(10013)]
	_ = x[PROP_PLAYER_EXP-(10014)]
	_ = x[PROP_PLAYER_HCOIN-(10015)]
	_ = x[PROP_PLAYER_SCOIN-(10016)]
	_ = x[PROP_PLAYER_MP_SETTING_TYPE-(10017)]
	_ = x[PROP_IS_MP_MODE_AVAILABLE-(10018)]
	_ = x[PROP_PLAYER_WORLD_LEVEL-(10019)]
	_ = x[PROP_PLAYER_RESIN-(10020)]
	_ = x[PROP_PLAYER_WAIT_SUB_HCOIN-(10022)]
	_ = x[PROP_PLAYER_WAIT_SUB_SCOIN-(10023)]
	_ = x[PROP_IS_ONLY_MP_WITH_PS_PLAYER-(10024)]
	_ = x[PROP_PLAYER_MCOIN-(10025)]
	_ = x[PROP_PLAYER_WAIT_SUB_MCOIN-(10026)]
	_ = x[PROP_PLAYER_LEGENDARY_KEY-(10027)]
	_ = x[PROP_IS_HAS_FIRST_SHARE-(10028)]
	_ = x[PROP_PLAYER_FORGE_POINT-(10029)]
	_ = x[PROP_CUR_CLIMATE_METER-(10035)]
	_ = x[PROP_CUR_CLIMATE_TYPE-(10036)]
	_ = x[PROP_CUR_CLIMATE_AREA_ID-(10037)]
	_ = x[PROP_CUR_CLIMATE_AREA_CLIMATE_TYPE-(10038)]
	_ = x[PROP_PLAYER_WORLD_LEVEL_LIMIT-(10039)]
	_ = x[PROP_PLAYER_WORLD_LEVEL_ADJUST_CD-(10040)]
	_ = x[PROP_PLAYER_LEGENDARY_DAILY_TASK_NUM-(10041)]
	_ = x[PROP_PLAYER_HOME_COIN-(10042)]
	_ = x[PROP_PLAYER_WAIT_SUB_HOME_COIN-(10043)]
	_ = x[PROP_IS_AUTO_UNLOCK_SPECIFIC_EQUIP-(10044)]
	_ = x[PROP_PLAYER_GCG_COIN-(10045)]
	_ = x[PROP_PLAYER_WAIT_SUB_GCG_COIN-(10046)]
	_ = x[PROP_PLAYER_ONLINE_TIME-(10047)]
	_ = x[PROP_PLAYER_CAN_DIVE-(10048)]
	_ = x[PROP_DIVE_MAX_STAMINA-(10049)]
	_ = x[PROP_DIVE_CUR_STAMINA-(10050)]
}

var _PlayerPropValues = []PlayerProp{PROP_NONE, PROP_EXP, PROP_BREAK_LEVEL, PROP_SATIATION_VAL, PROP_SATIATION_PENALTY_TIME, PROP_LEVEL, PROP_LAST_CHANGE_AVATAR_TIME, PROP_MAX_SPRING_VOLUME, PROP_CUR_SPRING_VOLUME, PROP_IS_SPRING_AUTO_USE, PROP_SPRING_AUTO_USE_PERCENT, PROP_IS_FLYABLE, PROP_IS_WEATHER_LOCKED, PROP_IS_GAME_TIME_LOCKED, PROP_IS_TRANSFERABLE, PROP_MAX_STAMINA, PROP_CUR_PERSIST_STAMINA, PROP_CUR_TEMPORARY_STAMINA, PROP_PLAYER_LEVEL, PROP_PLAYER_EXP, PROP_PLAYER_HCOIN, PROP_PLAYER_SCOIN, PROP_PLAYER_MP_SETTING_TYPE, PROP_IS_MP_MODE_AVAILABLE, PROP_PLAYER_WORLD_LEVEL, PROP_PLAYER_RESIN, PROP_PLAYER_WAIT_SUB_HCOIN, PROP_PLAYER_WAIT_SUB_SCOIN, PROP_IS_ONLY_MP_WITH_PS_PLAYER, PROP_PLAYER_MCOIN, PROP_PLAYER_WAIT_SUB_MCOIN, PROP_PLAYER_LEGENDARY_KEY, PROP_IS_HAS_FIRST_SHARE, PROP_PLAYER_FORGE_POINT, PROP_CUR_CLIMATE_METER, PROP_CUR_CLIMATE_TYPE, PROP_CUR_CLIMATE_AREA_ID, PROP_CUR_CLIMATE_AREA_CLIMATE_TYPE, PROP_PLAYER_WORLD_LEVEL_LIMIT, PROP_PLAYER_WORLD_LEVEL_ADJUST_CD, PROP_PLAYER_LEGENDARY_DAILY_TASK_NUM, PROP_PLAYER_HOME_COIN, PROP_PLAYER_WAIT_SUB_HOME_COIN, PROP_IS_AUTO_UNLOCK_SPECIFIC_EQUIP, PROP_PLAYER_GCG_COIN, PROP_PLAYER_WAIT_SUB_GCG_COIN, PROP_PLAYER_ONLINE_TIME, PROP_PLAYER_CAN_DIVE, PROP_DIVE_MAX_STAMINA, PROP_DIVE_CUR_STAMINA}

var _PlayerPropNameToValueMap = map[string]PlayerProp{
	_PlayerPropName_0[0:9]:          PROP_NONE,
	_PlayerPropLowerName_0[0:9]:     PROP_NONE,
	_PlayerPropName_1[0:8]:          PROP_EXP,
	_PlayerPropLowerName_1[0:8]:     PROP_EXP,
	_PlayerPropName_1[8:24]:         PROP_BREAK_LEVEL,
	_PlayerPropLowerName_1[8:24]:    PROP_BREAK_LEVEL,
	_PlayerPropName_1[24:42]:        PROP_SATIATION_VAL,
	_PlayerPropLowerName_1[24:42]:   PROP_SATIATION_VAL,
	_PlayerPropName_1[42:69]:        PROP_SATIATION_PENALTY_TIME,
	_PlayerPropLowerName_1[42:69]:   PROP_SATIATION_PENALTY_TIME,
	_PlayerPropName_2[0:10]:         PROP_LEVEL,
	_PlayerPropLowerName_2[0:10]:    PROP_LEVEL,
	_PlayerPropName_3[0:28]:         PROP_LAST_CHANGE_AVATAR_TIME,
	_PlayerPropLowerName_3[0:28]:    PROP_LAST_CHANGE_AVATAR_TIME,
	_PlayerPropName_3[28:50]:        PROP_MAX_SPRING_VOLUME,
	_PlayerPropLowerName_3[28:50]:   PROP_MAX_SPRING_VOLUME,
	_PlayerPropName_3[50:72]:        PROP_CUR_SPRING_VOLUME,
	_PlayerPropLowerName_3[50:72]:   PROP_CUR_SPRING_VOLUME,
	_PlayerPropName_3[72:95]:        PROP_IS_SPRING_AUTO_USE,
	_PlayerPropLowerName_3[72:95]:   PROP_IS_SPRING_AUTO_USE,
	_PlayerPropName_3[95:123]:       PROP_SPRING_AUTO_USE_PERCENT,
	_PlayerPropLowerName_3[95:123]:  PROP_SPRING_AUTO_USE_PERCENT,
	_PlayerPropName_3[123:138]:      PROP_IS_FLYABLE,
	_PlayerPropLowerName_3[123:138]: PROP_IS_FLYABLE,
	_PlayerPropName_3[138:160]:      PROP_IS_WEATHER_LOCKED,
	_PlayerPropLowerName_3[138:160]: PROP_IS_WEATHER_LOCKED,
	_PlayerPropName_3[160:184]:      PROP_IS_GAME_TIME_LOCKED,
	_PlayerPropLowerName_3[160:184]: PROP_IS_GAME_TIME_LOCKED,
	_PlayerPropName_3[184:204]:      PROP_IS_TRANSFERABLE,
	_PlayerPropLowerName_3[184:204]: PROP_IS_TRANSFERABLE,
	_PlayerPropName_3[204:220]:      PROP_MAX_STAMINA,
	_PlayerPropLowerName_3[204:220]: PROP_MAX_STAMINA,
	_PlayerPropName_3[220:244]:      PROP_CUR_PERSIST_STAMINA,
	_PlayerPropLowerName_3[220:244]: PROP_CUR_PERSIST_STAMINA,
	_PlayerPropName_3[244:270]:      PROP_CUR_TEMPORARY_STAMINA,
	_PlayerPropLowerName_3[244:270]: PROP_CUR_TEMPORARY_STAMINA,
	_PlayerPropName_3[270:287]:      PROP_PLAYER_LEVEL,
	_PlayerPropLowerName_3[270:287]: PROP_PLAYER_LEVEL,
	_PlayerPropName_3[287:302]:      PROP_PLAYER_EXP,
	_PlayerPropLowerName_3[287:302]: PROP_PLAYER_EXP,
	_PlayerPropName_3[302:319]:      PROP_PLAYER_HCOIN,
	_PlayerPropLowerName_3[302:319]: PROP_PLAYER_HCOIN,
	_PlayerPropName_3[319:336]:      PROP_PLAYER_SCOIN,
	_PlayerPropLowerName_3[319:336]: PROP_PLAYER_SCOIN,
	_PlayerPropName_3[336:363]:      PROP_PLAYER_MP_SETTING_TYPE,
	_PlayerPropLowerName_3[336:363]: PROP_PLAYER_MP_SETTING_TYPE,
	_PlayerPropName_3[363:388]:      PROP_IS_MP_MODE_AVAILABLE,
	_PlayerPropLowerName_3[363:388]: PROP_IS_MP_MODE_AVAILABLE,
	_PlayerPropName_3[388:411]:      PROP_PLAYER_WORLD_LEVEL,
	_PlayerPropLowerName_3[388:411]: PROP_PLAYER_WORLD_LEVEL,
	_PlayerPropName_3[411:428]:      PROP_PLAYER_RESIN,
	_PlayerPropLowerName_3[411:428]: PROP_PLAYER_RESIN,
	_PlayerPropName_4[0:26]:         PROP_PLAYER_WAIT_SUB_HCOIN,
	_PlayerPropLowerName_4[0:26]:    PROP_PLAYER_WAIT_SUB_HCOIN,
	_PlayerPropName_4[26:52]:        PROP_PLAYER_WAIT_SUB_SCOIN,
	_PlayerPropLowerName_4[26:52]:   PROP_PLAYER_WAIT_SUB_SCOIN,
	_PlayerPropName_4[52:82]:        PROP_IS_ONLY_MP_WITH_PS_PLAYER,
	_PlayerPropLowerName_4[52:82]:   PROP_IS_ONLY_MP_WITH_PS_PLAYER,
	_PlayerPropName_4[82:99]:        PROP_PLAYER_MCOIN,
	_PlayerPropLowerName_4[82:99]:   PROP_PLAYER_MCOIN,
	_PlayerPropName_4[99:125]:       PROP_PLAYER_WAIT_SUB_MCOIN,
	_PlayerPropLowerName_4[99:125]:  PROP_PLAYER_WAIT_SUB_MCOIN,
	_PlayerPropName_4[125:150]:      PROP_PLAYER_LEGENDARY_KEY,
	_PlayerPropLowerName_4[125:150]: PROP_PLAYER_LEGENDARY_KEY,
	_PlayerPropName_4[150:173]:      PROP_IS_HAS_FIRST_SHARE,
	_PlayerPropLowerName_4[150:173]: PROP_IS_HAS_FIRST_SHARE,
	_PlayerPropName_4[173:196]:      PROP_PLAYER_FORGE_POINT,
	_PlayerPropLowerName_4[173:196]: PROP_PLAYER_FORGE_POINT,
	_PlayerPropName_5[0:22]:         PROP_CUR_CLIMATE_METER,
	_PlayerPropLowerName_5[0:22]:    PROP_CUR_CLIMATE_METER,
	_PlayerPropName_5[22:43]:        PROP_CUR_CLIMATE_TYPE,
	_PlayerPropLowerName_5[22:43]:   PROP_CUR_CLIMATE_TYPE,
	_PlayerPropName_5[43:67]:        PROP_CUR_CLIMATE_AREA_ID,
	_PlayerPropLowerName_5[43:67]:   PROP_CUR_CLIMATE_AREA_ID,
	_PlayerPropName_5[67:101]:       PROP_CUR_CLIMATE_AREA_CLIMATE_TYPE,
	_PlayerPropLowerName_5[67:101]:  PROP_CUR_CLIMATE_AREA_CLIMATE_TYPE,
	_PlayerPropName_5[101:130]:      PROP_PLAYER_WORLD_LEVEL_LIMIT,
	_PlayerPropLowerName_5[101:130]: PROP_PLAYER_WORLD_LEVEL_LIMIT,
	_PlayerPropName_5[130:163]:      PROP_PLAYER_WORLD_LEVEL_ADJUST_CD,
	_PlayerPropLowerName_5[130:163]: PROP_PLAYER_WORLD_LEVEL_ADJUST_CD,
	_PlayerPropName_5[163:199]:      PROP_PLAYER_LEGENDARY_DAILY_TASK_NUM,
	_PlayerPropLowerName_5[163:199]: PROP_PLAYER_LEGENDARY_DAILY_TASK_NUM,
	_PlayerPropName_5[199:220]:      PROP_PLAYER_HOME_COIN,
	_PlayerPropLowerName_5[199:220]: PROP_PLAYER_HOME_COIN,
	_PlayerPropName_5[220:250]:      PROP_PLAYER_WAIT_SUB_HOME_COIN,
	_PlayerPropLowerName_5[220:250]: PROP_PLAYER_WAIT_SUB_HOME_COIN,
	_PlayerPropName_5[250:284]:      PROP_IS_AUTO_UNLOCK_SPECIFIC_EQUIP,
	_PlayerPropLowerName_5[250:284]: PROP_IS_AUTO_UNLOCK_SPECIFIC_EQUIP,
	_PlayerPropName_5[284:304]:      PROP_PLAYER_GCG_COIN,
	_PlayerPropLowerName_5[284:304]: PROP_PLAYER_GCG_COIN,
	_PlayerPropName_5[304:333]:      PROP_PLAYER_WAIT_SUB_GCG_COIN,
	_PlayerPropLowerName_5[304:333]: PROP_PLAYER_WAIT_SUB_GCG_COIN,
	_PlayerPropName_5[333:356]:      PROP_PLAYER_ONLINE_TIME,
	_PlayerPropLowerName_5[333:356]: PROP_PLAYER_ONLINE_TIME,
	_PlayerPropName_5[356:376]:      PROP_PLAYER_CAN_DIVE,
	_PlayerPropLowerName_5[356:376]: PROP_PLAYER_CAN_DIVE,
	_PlayerPropName_5[376:397]:      PROP_DIVE_MAX_STAMINA,
	_PlayerPropLowerName_5[376:397]: PROP_DIVE_MAX_STAMINA,
	_PlayerPropName_5[397:418]:      PROP_DIVE_CUR_STAMINA,
	_PlayerPropLowerName_5[397:418]: PROP_DIVE_CUR_STAMINA,
}

var _PlayerPropNames = []string{
	_PlayerPropName_0[0:9],
	_PlayerPropName_1[0:8],
	_PlayerPropName_1[8:24],
	_PlayerPropName_1[24:42],
	_PlayerPropName_1[42:69],
	_PlayerPropName_2[0:10],
	_PlayerPropName_3[0:28],
	_PlayerPropName_3[28:50],
	_PlayerPropName_3[50:72],
	_PlayerPropName_3[72:95],
	_PlayerPropName_3[95:123],
	_PlayerPropName_3[123:138],
	_PlayerPropName_3[138:160],
	_PlayerPropName_3[160:184],
	_PlayerPropName_3[184:204],
	_PlayerPropName_3[204:220],
	_PlayerPropName_3[220:244],
	_PlayerPropName_3[244:270],
	_PlayerPropName_3[270:287],
	_PlayerPropName_3[287:302],
	_PlayerPropName_3[302:319],
	_PlayerPropName_3[319:336],
	_PlayerPropName_3[336:363],
	_PlayerPropName_3[363:388],
	_PlayerPropName_3[388:411],
	_PlayerPropName_3[411:428],
	_PlayerPropName_4[0:26],
	_PlayerPropName_4[26:52],
	_PlayerPropName_4[52:82],
	_PlayerPropName_4[82:99],
	_PlayerPropName_4[99:125],
	_PlayerPropName_4[125:150],
	_PlayerPropName_4[150:173],
	_PlayerPropName_4[173:196],
	_PlayerPropName_5[0:22],
	_PlayerPropName_5[22:43],
	_PlayerPropName_5[43:67],
	_PlayerPropName_5[67:101],
	_PlayerPropName_5[101:130],
	_PlayerPropName_5[130:163],
	_PlayerPropName_5[163:199],
	_PlayerPropName_5[199:220],
	_PlayerPropName_5[220:250],
	_PlayerPropName_5[250:284],
	_PlayerPropName_5[284:304],
	_PlayerPropName_5[304:333],
	_PlayerPropName_5[333:356],
	_PlayerPropName_5[356:376],
	_PlayerPropName_5[376:397],
	_PlayerPropName_5[397:418],
}

// PlayerPropString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func PlayerPropString(s string) (PlayerProp, error) {
	if val, ok := _PlayerPropNameToValueMap[s]; ok {
		return val, nil
	}

	if val, ok := _PlayerPropNameToValueMap[strings.ToLower(s)]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to PlayerProp values", s)
}

// PlayerPropValues returns all values of the enum
func PlayerPropValues() []PlayerProp {
	return _PlayerPropValues
}

// PlayerPropStrings returns a slice of all String values of the enum
func PlayerPropStrings() []string {
	strs := make([]string, len(_PlayerPropNames))
	copy(strs, _PlayerPropNames)
	return strs
}

// IsAPlayerProp returns "true" if the value is listed in the enum definition. "false" otherwise
func (i PlayerProp) IsAPlayerProp() bool {
	for _, v := range _PlayerPropValues {
		if i == v {
			return true
		}
	}
	return false
}

// MarshalJSON implements the json.Marshaler interface for PlayerProp
func (i PlayerProp) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for PlayerProp
func (i *PlayerProp) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("PlayerProp should be a string, got %s", data)
	}

	var err error
	*i, err = PlayerPropString(s)
	return err
}
