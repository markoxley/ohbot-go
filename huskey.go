package ohbot

const (
	CommandRequest byte = 0x20 + iota
	CommandRequestBlocks
	CommandRRequestArrows
	CommandRequestLearned
	CommandRequestBlocksLearner
	CommandRArrowsLearned
	CommandRequestByID
	CommandRequestBlocksByID
	CommandRequestArrowsByID
	CommandReturnInfo
	CommandReturnBlock
	CommandReturnArrow
	CommandRequestKnock
	CommandRequestAlgorithm
	CommandReturnOK
	CommandRequestCustomNames
	CommandRequestPhoto
	CommandRequestSendPhoto
	CommandRequestSendKnowledges
	CommandRequestReceiveKnowledges
	CommandRequestCustomText
	CommandRequestClearText
	CommandRequestLearn
	CommandRequestForget
	CommandRequestSendScreenshot
	CommandRequestSaveScreenshot
	CommandRequestLoadAIFrameFromUSB
	CommandRequestIsPro

	CommandRequestFirmwareVersion = 0x3C + iota
	CommandRequestSensor

	HuskyI2C = 0x32
)

func SetI2CPort() {

}

func I2CWrite() {}

func I2CRead() {}
