package main

import "fmt"

//	模板方法模式是一种行为设计模式，
//	它在超类中定义了一个算法的框架， 允许子类在不修改结构的情况下重写算法的特定步骤。

//	让我们来考虑一个一次性密码功能 （OTP） 的例子。
//	将 OTP 传递给用户的方式多种多样 （短信、 邮件等）。
//	但无论是短信还是邮件， 整个 OTP 流程都是相同的：
//
//	生成随机的 n 位数字。
//	在缓存中保存这组数字以便进行后续验证。
//	准备内容。
//	发送通知。
//	后续引入的任何新 OTP 类型都很有可能需要进行相同的上述步骤。
//
//	因此， 我们会有这样的一个场景， 其中某个特定操作的步骤是相同的，
//	但实现方式却可能有所不同。 这正是适合考虑使用模板方法模式的情况。
//
//	首先， 我们定义一个由固定数量的方法组成的基础模板算法。
//	这就是我们的模板方法。 然后我们将实现每一个步骤方法， 但不会改变模板方法。

func main() {
	// otp := otp{}

	// smsOTP := &sms{
	//  otp: otp,
	// }

	// smsOTP.genAndSendOTP(smsOTP, 4)

	// emailOTP := &email{
	//  otp: otp,
	// }
	// emailOTP.genAndSendOTP(emailOTP, 4)
	// fmt.Scanln()
	smsOTP := &Sms{}
	o := Otp{
		iOtp: smsOTP,
	}
	o.genAndSendOTP(4)

	fmt.Println("")
	emailOTP := &Email{}
	o = Otp{
		iOtp: emailOTP,
	}
	o.genAndSendOTP(4)

}

type IOtp interface {
	genRandomOTP(int) string
	saveOTPCache(string)
	getMessage(string) string
	sendNotification(string) error
}

// type otp struct {
// }

// func (o *otp) genAndSendOTP(iOtp iOtp, otpLength int) error {
//  otp := iOtp.genRandomOTP(otpLength)
//  iOtp.saveOTPCache(otp)
//  message := iOtp.getMessage(otp)
//  err := iOtp.sendNotification(message)
//  if err != nil {
//      return err
//  }
//  return nil
// }

type Otp struct {
	iOtp IOtp
}

func (o *Otp) genAndSendOTP(otpLength int) error {
	otp := o.iOtp.genRandomOTP(otpLength)
	o.iOtp.saveOTPCache(otp)
	message := o.iOtp.getMessage(otp)
	err := o.iOtp.sendNotification(message)
	if err != nil {
		return err
	}
	return nil
}

type Sms struct {
	Otp
}

func (s *Sms) genRandomOTP(len int) string {
	randomOTP := "1234"
	fmt.Printf("SMS: generating random otp %s\n", randomOTP)
	return randomOTP
}

func (s *Sms) saveOTPCache(otp string) {
	fmt.Printf("SMS: saving otp: %s to cache\n", otp)
}

func (s *Sms) getMessage(otp string) string {
	return "SMS OTP for login is " + otp
}

func (s *Sms) sendNotification(message string) error {
	fmt.Printf("SMS: sending sms: %s\n", message)
	return nil
}

type Email struct {
	Otp
}

func (s *Email) genRandomOTP(len int) string {
	randomOTP := "1234"
	fmt.Printf("EMAIL: generating random otp %s\n", randomOTP)
	return randomOTP
}

func (s *Email) saveOTPCache(otp string) {
	fmt.Printf("EMAIL: saving otp: %s to cache\n", otp)
}

func (s *Email) getMessage(otp string) string {
	return "EMAIL OTP for login is " + otp
}

func (s *Email) sendNotification(message string) error {
	fmt.Printf("EMAIL: sending email: %s\n", message)
	return nil
}
