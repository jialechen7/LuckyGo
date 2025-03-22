package xerr

var message map[uint32]string

func init() {
	message = make(map[uint32]string)
	message[OK] = "SUCCESS"
	message[SERVER_COMMON_ERROR] = "服务器开小差啦,稍后再来试一试"
	message[REUQEST_PARAM_ERROR] = "参数错误"
	message[TOKEN_EXPIRE_ERROR] = "token失效，请重新登陆"
	message[TOKEN_GENERATE_ERROR] = "生成token失败"
	message[DB_ERROR] = "数据库繁忙,请稍后再试"
	message[DB_UPDATE_AFFECTED_ZERO_ERROR] = "更新数据影响行数为0"
	message[DB_ERROR_NOT_FOUND] = "数据未找到"
	message[DB_TRANSACTION_ERROR] = "数据库事务错误"

	// 抽奖模块
	message[DB_GETLASTID_ERROR] = "获取最后一个抽奖id失败"
	message[DB_GET_LOTTERY_LIST_ERROR] = "获取抽奖列表失败"
	message[DB_GET_WIN_LOTTERY_LIST_ERROR] = "获取中奖抽奖列表失败"
	message[DB_GET_PRIZE_LIST_ERROR] = "获取奖品列表失败"
	message[DB_GET_USER_WIN_LOTTERY_LIST_ERROR] = "获取用户中奖抽奖列表失败"
	message[DB_GET_ALL_USER_LOTTERY_LIST_ERROR] = "获取用户所有抽奖列表失败"
	message[DB_GET_USER_CREATED_LOTTERY_LIST_ERROR] = "获取用户创建抽奖列表失败"
	message[DB_GET_PRIZE_ERROR] = "获取奖品失败"
	message[DB_GET_PARTICIPATED_BY_LOTTERY_ID] = "获取用户参与抽奖信息失败"
	message[DB_GET_PARTICIPATIONS_COUNT] = "获取用户参与抽奖数量失败"
	message[DB_GET_CREATED_COUNT_BY_USER_ID] = "获取用户创建抽奖数量失败"
	message[DB_GET_WON_COUNT_BY_USER_ID] = "获取用户中奖数量失败"
	message[DB_GET_PARTICIPATION_COUNT_BY_USER_ID] = "获取用户参与抽奖数量失败"
	message[DB_GET_PRIZE_BY_BY_LOTTERY_ID] = "获取抽奖奖品失败"
	message[DB_CHECK_IS_PARTICIPATED] = "检查用户是否参与抽奖失败"
	message[DB_PARTICIPATE_LOTTERY] = "参与抽奖失败"
	message[DB_GET_PENDING_LOTTERY_LIST_FAIL] = "获取待开奖抽奖列表失败"
	message[DB_GET_PARTICIPATORS_COUNT_FAIL] = "获取参与抽奖人数失败"
	message[DB_GET_PARTICIPATORS_ERROR] = "获取参与抽奖人员失败"
	message[DB_ANNOUNCE_LOTTERY_FAIL] = "开奖失败"
	message[ANNOUNCE_LOTTERY_FAIL] = "开奖失败"
	message[DB_UPDATE_STATUS_ERROR] = "更新抽奖状态失败"
	message[DB_UPDATE_WINNER_ERROR] = "更新中奖者失败"
	message[DB_FIND_PARTICIPATION_BY_LOTTERY_ID_AND_USER_ID_ERR] = "查找抽奖参与者失败"
	message[DB_FIND_LOTTERY_BY_ID_ERR] = "查找抽奖失败"
	message[DB_FIND_PRIZES_BY_LOTTERY_ID_ERR] = "查找抽奖奖品失败"
	message[DB_DECR_PRIZE_STOCK_ERR] = "减少奖品库存失败"

	// 用户模块
	message[DB_INSERT_USER_SPONSOR_ERROR] = "插入用户赞助商失败"
	message[DB_UPDATE_USER_SPONSOR_ERROR] = "更新用户赞助商失败"
	message[DB_DELETE_USER_SPONSOR_ERROR] = "删除用户赞助商失败"
	message[DB_GET_USER_SPONSOR_LIST_ERROR] = "获取用户赞助商列表失败"
	message[DB_GET_USERS_INFO_ERROR] = "获取用户信息失败"
	message[DB_USER_AUTH_FIND_ONE_BY_USER_ID_ERROR] = "查找用户授权失败"

	// 评论模块
	message[DB_INSERT_COMMENT_ERROR] = "插入评论失败"
	message[DB_INSERT_PRAISE_ERROR] = "插入点赞失败"
	message[DB_DELETE_COMMENT_ERROR] = "删除评论失败"
	message[DB_DELETE_PRAISE_ERROR] = "删除点赞失败"
	message[DB_UPDATE_COMMENT_ERROR] = "更新评论失败"
	message[DB_UPDATE_PRAISE_ERROR] = "更新点赞失败"
	message[DB_FIND_COMMENT_ERROR] = "查找评论失败"
	message[DB_FIND_PRAISE_ERROR] = "查找点赞失败"
	message[DB_GET_COMMENT_LAST_ID_ERROR] = "获取评论最后一个id失败"
	message[DB_GET_COMMENT_LIST_ERROR] = "获取评论列表失败"
	message[DB_GET_COMMENT_LIST_BY_USER_ID_ERROR] = "获取用户评论列表失败"
	message[DB_PRAISE_COMMENT_ERROR] = "点赞评论失败"
	message[DB_GET_PRAISE_LIST_ERROR] = "获取点赞列表失败"

	// 签到模块
	message[DB_FIND_CHECKIN_RECORD_ERROR] = "查找签到记录失败"
	message[DB_CHECKIN_RECORD_INSERT_ERROR] = "插入签到记录失败"
	message[DB_FIND_INTEGRAL_ERROR] = "查找积分失败"
	message[DB_INTEGRAL_INSERT_ERROR] = "插入积分失败"
	message[DB_CHECKIN_TRANSACT_ERROR] = "签到事务失败"
	message[CHECKIN_REPEAT_ERROR] = "重复签到"
	message[DB_INTEGRAL_UPDATE_ERROR] = "更新积分失败"
	message[DB_CHECKIN_RECORD_UPDATE_ERROR] = "更新签到记录失败"
	message[DB_INTEGRAL_RECORD_INSERT_ERROR] = "插入积分记录失败"
	message[DB_GET_CHECKIN_RECORD_BY_USER_ID_ERROR] = "获取签到记录失败"
	message[DB_TASK_PROGRESS_FIND_ONE_BY_USER_ID] = "获取积分失败"
	message[DB_TASK_PROGRESS_INSERT_BY_USER_ID] = "插入积分失败"
	message[CHECKIN_TASK_NEWBIE_RUN] = "新手任务失败"
	message[DB_TASK_RECORD_FIND_BY_USER_ID_AND_TASK_ID] = "获取任务记录失败"
	message[DB_GET_TASK_PROGRESS_TRANSACT_ERROR] = "获取任务进度事务失败"
	message[CHECKIN_TASK_NEWBIE_CHECK_LOTTERY_PARTICIPATED] = "检查抽奖参与失败"
	message[DB_ADD_TASK_RECORD_ERROR] = "添加任务记录失败"
	message[DB_TASK_PROGRESS_UPDATE_BY_USER_ID] = "更新任务进度失败"
	message[DB_TASK_RECORD_FIND_BY_USER_ID] = "获取任务进度失败"
	message[DB_TASKS_FIND_ALL] = "获取任务列表失败"
	message[DB_TASK_FIND_ONE_ERROR] = "获取任务失败"
	message[DB_TASK_RECORD_INSERT_ERROR] = "插入任务记录失败"
	message[CHECKIN_TASK_HAS_AWARDED] = "任务已领取"
	message[DB_TASK_RECORD_UPDATE_ERROR] = "更新任务记录失败"
	message[DB_GET_LOTTERY_PARTICIPATED_COUNT] = "获取抽奖参与次数失败"
	message[DB_GET_LOTTERY_CREATED_COUNT] = "获取抽奖创建次数失败"
	message[CHECKIN_TASK_NEWBIE_CHECK_LOTTERY_CREATED] = "检查抽奖创建失败"
	message[DB_INTEGRAL_RECORD_DELETE_ERROR] = "删除积分记录失败"
	message[DB_TASK_RECORD_DELETE_ERROR] = "删除任务记录失败"
	message[DB_TASKS_DELETE_ERROR] = "删除任务失败"
	message[DB_TASKS_INSERT_ERROR] = "插入任务失败"
	message[DB_TASK_PROGRESS_FIND_ALL_SUBSCRIBE_USER_IDS_ERROR] = "获取所有订阅用户id失败"
}

func MapErrMsg(errcode uint32) string {
	if msg, ok := message[errcode]; ok {
		return msg
	} else {
		return "服务器开小差啦,稍后再来试一试"
	}
}

func IsCodeErr(errcode uint32) bool {
	if _, ok := message[errcode]; ok {
		return true
	} else {
		return false
	}
}
