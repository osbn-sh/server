package Activityconstants

type ActivityTriggersName string

//اضافه کردن دروس پاس شده هر کدام ۱ امتیاز ( یعنی هرچه ترم بالاتر امتیاز بیشتر )
//اضافه کردن دانشگاه ۱۳ امتیاز در صورت قبولی نهایی
//اضافه کردن استاد ۱۳ امتیاز در صورت قبولی نهایی
//اضافه کردن درس ۱۱ امتیاز در صورت قبولی نهایی
//اضافه کردن رشته ۹ امتیاز در صورت قبولی نهایی

const (
	TriggerRegisterAccount ActivityTriggersName = "register-account"
	TriggerMakeAdmin       ActivityTriggersName = "make-admin"
	TriggerUnmakeAdmin     ActivityTriggersName = "unmake-admin"

	TriggerAddPassedLesson ActivityTriggersName = "add-passed-lesson"

	TriggerStabilizeLesson     ActivityTriggersName = "stabilize-lesson"
	TriggerStabilizeProfessor  ActivityTriggersName = "stabilize-professor"
	TriggerStabilizeUniversity ActivityTriggersName = "stabilize-university"
	TriggerStabilizeMajor      ActivityTriggersName = "stabilize-major"
)
