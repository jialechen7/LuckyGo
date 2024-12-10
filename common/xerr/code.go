package xerr

// 成功返回
const OK uint32 = 200

/**(前3位代表业务,后三位代表具体功能)**/

// 全局错误码
const (
	SERVER_COMMON_ERROR uint32 = 100001 + iota
	REUQEST_PARAM_ERROR
	TOKEN_EXPIRE_ERROR
	TOKEN_GENERATE_ERROR
	DB_ERROR
	DB_UPDATE_AFFECTED_ZERO_ERROR
)

//用户模块

// 抽奖模块
const (
	// 抽奖列表
	DB_GETLASTID_ERROR uint32 = 500001 + iota
	DB_GET_LOTTERY_LIST_ERROR
)
