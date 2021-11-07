package sender

/**
 *
 * @author hushengdong
 */
func SendMqReqByCode(f string, code int, targetType int) {

	initConfig(f)
	mq := NewMqReq(reqCode(code))
	if mustBeNameserver(code) {
		sendTcp(mq, getNameserver())
	} else {
		sendTcp(mq, getTargetByType(targetType))
	}
}

func mustBeNameserver(code int) bool {

	switch code {
	case 100, 101, 102, 103, 104,
		105, 106, 205, 206, 216, 217,
		218, 219, 224, 304, 311, 312, 313:
		return true
	default:
		return false
	}
}

func getTargetByType(targetType int) string {

	switch targetType {
	case 0:
		return getNameserver()
	case 1:
		return getBroker()
	case 2:
		return getConsumer()
	case 3:
		return getProducer()
	default:
		return getNameserver()
	}
}

func getTargetByCode(code int) string {

	switch code {
	case 34:
		return getBroker()
	case 105, 103, 101:
		return getNameserver()
	default:
		return getNameserver()
	}
}
