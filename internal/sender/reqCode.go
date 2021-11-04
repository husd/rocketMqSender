package sender

/**
 *
 * @author hushengdong
 */
type reqCode int32

const (
	// Broker 发送消息
	send_message reqCode = 10
	// broker 订阅消息
	pull_message reqCode = 11
	// broker 查询消息
	query_message reqCode = 12
	// broker 查询broker offset
	query_broker_offset reqCode = 13
	// broker 查询consumer offset
	query_consumer_offset reqCode = 14
	// broker 更新consumer offset
	update_consumer_offset reqCode = 15
	// broker 更新或者增加一个topic
	update_and_create_topic reqCode = 17
	// broker 获取所有topic的配置（slave和namesrv都会向master请求此配置）
	get_all_topic_config reqCode = 21
	// broker 获取所有topic配置（slave和namesrv都会向master请求此配置）
	get_topic_config_list reqCode = 22
	// broker 获取所有topic名称列表
	get_topic_name_list reqCode = 23
	// broker 更新broker上的配置
	update_broker_config reqCode = 25
	// broker 获取broker上的配置
	get_broker_config reqCode = 26
	// broker 触发broker删除文件
	trigger_delete_files reqCode = 27
	// broker 获取broker运行时信息
	get_broker_runtime_info reqCode = 28
	// broker 根据时间查询队列的offset
	search_offset_by_timestamp reqCode = 29
	// broker 查询队列最大offset
	get_max_offset reqCode = 30
	// broker 查询队列最小offset
	get_min_offset reqCode = 31
	// broker 查询队列最早消息对应时间
	get_earliest_msg_storetime reqCode = 32
	// broker 根据消息id来查询消息
	view_message_by_id reqCode = 33
	// broker client向client发送心跳，并注册自身
	heart_beat reqCode = 34
	// broker client注销
	unregister_client reqCode = 35
	// broker consumer将处理不了的消息发回服务器
	consumer_send_msg_back reqCode = 36
	// broker commit或者rollback事务
	end_transaction reqCode = 37
	// broker 获取consumerid列表通过groupnamrebalanceservicee
	get_consumer_list_by_group reqCode = 38
	// broker 主动向producer回查事务状态
	check_transaction_state reqCode = 39
	// broker broker通知consumer列表变化
	notify_consumer_ids_changed reqCode = 40
	// broker consumer向master锁定队列
	lock_batch_mq reqCode = 41
	// broker consumer向master解锁队列
	unlock_batch_mq reqCode = 42
	// broker 获取所有consumer offset
	get_all_consumer_offset reqCode = 43
	// broker 获取所有定时进度
	get_all_delay_offset reqCode = 45
	// namesrv 向namesrv追加kv配置
	put_kv_config reqCode = 100
	// namesrv 从namesrv获取kv配置
	get_kv_config reqCode = 101
	// namesrv 从namesrv获取kv配置
	delete_kv_config reqCode = 102
	// namesrv 注册一个broker，数据都是持久化的，如果存在则覆盖配置
	register_broker reqCode = 103
	// namesrv 卸载一个broker，数据都是持久化的
	unregister_broker reqCode = 104
	// namesrv 根据topic获取broker name、队列数(包含读队列与写队列)
	get_routeinto_by_topic reqCode = 105
	// namesrv 获取注册到name server的所有broker集群信息
	get_broker_cluster_info             reqCode = 106
	update_and_create_subscriptiongroup reqCode = 200
	get_all_subscriptiongroup_config    reqCode = 201
	get_topic_stats_info                reqCode = 202
	get_consumer_connection_list        reqCode = 203
	get_producer_connection_list        reqCode = 204
	wipe_write_perm_of_broker           reqCode = 205

	// 从name server获取完整topic列表
	get_all_topic_list_from_nameserver reqCode = 206
	// 从broker删除订阅组
	delete_subscriptiongroup reqCode = 207
	// 从broker获取消费状态（进度）
	get_consume_stats reqCode = 208
	// suspend consumer消费过程
	suspend_consumer reqCode = 209
	// resume consumer消费过程
	resume_consumer reqCode = 210
	// 重置consumer offset
	reset_consumer_offset_in_consumer reqCode = 211
	// 重置consumer offset
	reset_consumer_offset_in_broker reqCode = 212
	// 调整consumer线程池数量
	adjust_consumer_thread_pool reqCode = 213
	// 查询消息被哪些消费组消费
	who_consume_the_message reqCode = 214

	// 从broker删除topic配置
	delete_topic_in_broker reqCode = 215
	// 从namesrv删除topic配置
	delete_topic_in_namesrv reqCode = 216
	// namesrv 通过 project 获取所有的 server ip 信息
	get_kv_config_by_value reqCode = 217
	// namesrv 删除指定 project group 下的所有 server ip 信息
	delete_kv_config_by_value reqCode = 218
	// 通过namespace获取所有的kv list
	get_kvlist_by_namespace reqCode = 219

	// offset 重置
	reset_consumer_client_offset reqCode = 220
	// 客户端订阅消息
	get_consumer_status_from_client reqCode = 221
	// 通知 broker 调用 offset 重置处理
	invoke_broker_to_reset_offset reqCode = 222
	// 通知 broker 调用客户端订阅消息处理
	invoke_broker_to_get_consumer_status reqCode = 223

	// broker 查询topic被谁消费
	// 2014-03-21 add by shijia
	query_topic_consume_by_who reqCode = 300

	// 获取指定集群下的所有 topic
	// 2014-03-26
	get_topics_by_cluster reqCode = 224

	// 向broker注册filter server
	// 2014-04-06 add by shijia
	register_filter_server reqCode = 301
	// 向filter server注册class
	// 2014-04-06 add by shijia
	register_message_filter_class reqCode = 302
	// 根据 topic 和 group 获取消息的时间跨度
	query_consume_time_span reqCode = 303
	// 获取所有系统内置 topic 列表
	get_system_topic_list_from_ns     reqCode = 304
	get_system_topic_list_from_broker reqCode = 305

	// 清理失效队列
	clean_expired_consumequeue reqCode = 306

	// 通过broker查询consumer内存数据
	// 2014-07-19 add by shijia
	get_consumer_running_info reqCode = 307

	// 查找被修正 offset (转发组件）
	query_correction_offset reqCode = 308

	// 通过broker直接向某个consumer发送一条消息，并立刻消费，返回结果给broker，再返回给调用方
	// 2014-08-11 add by shijia
	consume_message_directly reqCode = 309

	// broker 发送消息，优化网络数据包
	send_message_v2 reqCode = 310

	// 单元化相关 topic
	get_unit_topic_list reqCode = 311
	// 获取含有单元化订阅组的 topic 列表
	get_has_unit_sub_topic_list reqCode = 312
	// 获取含有单元化订阅组的非单元化 topic 列表
	get_has_unit_sub_ununit_topic_list reqCode = 313
	// 克隆某一个组的消费进度到新的组
	clone_group_offset reqCode = 314

	// 查看broker上的各种统计信息
	view_broker_stats_data reqCode = 315
	/** 发送死信队列 */
	send_dlq_message reqCode = 316
)
