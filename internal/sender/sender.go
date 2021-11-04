package sender

/**
 *
 * @author hushengdong
 */
func SendMqReqByCode(f string, code int, targetType int) {

	initConfig(f)
	mq := NewMqReq(reqCode(code))
	target := getTargetByCode(targetType)
	sendTcp(mq, target)
}

func getTargetByCode(code int) string {

	if code == 0 {
		return getNameserver()
	} else if code == 1 {
		return getBroker()
	}
	return getNameserver()
}
