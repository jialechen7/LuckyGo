package jobtype

// WxMiniProgramNotifyUserPayload mini program notify user
type WxMiniProgramNotifyUserPayload struct {
	MsgType  int
	OpenId   string
	PageAddr string
	Data     string
}
