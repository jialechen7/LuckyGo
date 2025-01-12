package wxnotice

const (
	TemplateIdMessageWishSign = "y1amPc00bPVyfgljYSjFuL2TLE6DspUMVqVJqcJNaJc"
	_                         = iota
	TypeWishSign
)

type (
	Message interface {
		Type() int
		TemplateId() string
	}

	Item struct {
		Value string `json:"value"`
		Color string `json:"color"`
	}

	// MessageWishCheckin 签到开始提醒的消息内容
	MessageWishCheckin struct {
		ActivityName          Item `json:"thing1"`
		ContinuousCheckinDays Item `json:"number3"`
		Time                  Item `json:"time4"`
		CheckinAward          Item `json:"thing10"`
		RemindText            Item `json:"thing9"`
	}
)

func (m *MessageWishCheckin) TemplateId() string {
	return TemplateIdMessageWishSign
}
func (m *MessageWishCheckin) Type() int {
	return TypeWishSign
}
