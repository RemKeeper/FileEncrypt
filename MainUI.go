package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"os"
	"strconv"
)

var (
	FileBrowserDia      *dialog.FileDialog
	PrivateKeySelectDia *dialog.FileDialog
	FilePath            string
	PrivateKeyPath      string
	EncryptMode         int
)

const (
	DesMode       = 1
	TripleDESMode = 2
	AesCBCMode    = 3
	AesCTRMode    = 4
	RsaMode       = 5
)

func MainUI() *fyne.Container {
	SelectFileButton := widget.NewButton("选择文件", func() {
		FileBrowserDia.Show()
	})
	FileBrowserDia = dialog.NewFileOpen(func(closer fyne.URIReadCloser, err error) {
		if closer == nil {
			return
		}
		FilePath = closer.URI().Path()
		SelectFileButton.Text = "文件: " + FilePath
		SelectFileButton.Refresh()
	}, Windows)

	SelectPrivateKeyButton := widget.NewButton("选择私钥", func() {
		PrivateKeySelectDia.Show()
	})
	PrivateKeySelectDia = dialog.NewFileOpen(func(closer fyne.URIReadCloser, err error) {
		if closer == nil {
			return
		}
		PrivateKeyPath = closer.URI().Path()
		SelectPrivateKeyButton.Text = "私钥: " + PrivateKeyPath
		SelectPrivateKeyButton.Refresh()
	}, Windows)

	PasswordReminderLabel := widget.NewLabel("密码提示")

	PasswordInput := widget.NewPasswordEntry()
	PasswordInput.OnChanged = func(s string) {
		switch EncryptMode {

		case DesMode:
			if len(PasswordInput.Text) > 8 {
				PasswordInput.Text = PasswordInput.Text[:8]
			}
			PasswordReminderLabel.SetText("DES模式请输入8位密钥,当前位数" + strconv.Itoa(len(PasswordInput.Text)))
		case TripleDESMode:
			if len(PasswordInput.Text) > 24 {
				PasswordInput.Text = PasswordInput.Text[:24]
			}
			PasswordReminderLabel.SetText("三重DES模式请输入24位密钥,当前位数" + strconv.Itoa(len(PasswordInput.Text)))
		case AesCBCMode:

			if len(PasswordInput.Text) > 16 {
				PasswordInput.Text = PasswordInput.Text[:16]
			}
			PasswordReminderLabel.SetText("AES-CBC模式请输入16位密钥,当前位数" + strconv.Itoa(len(PasswordInput.Text)))
		case AesCTRMode:

			if len(PasswordInput.Text) > 16 {
				PasswordInput.Text = PasswordInput.Text[:16]
			}
			PasswordReminderLabel.SetText("AES-CTR模式请输入16位密钥,当前位数" + strconv.Itoa(len(PasswordInput.Text)))
		}
		PasswordReminderLabel.Refresh()
	}

	EncryptModeSelect := widget.NewRadioGroup([]string{"DES", "三重DES", "AES-CBC", "AES-CTR", "RSA"}, func(s string) {
		switch s {
		case "DES":
			EncryptMode = DesMode
			SelectPrivateKeyButton.Hide()
			PasswordInput.Show()
			PasswordReminderLabel.SetText("DES模式请输入8位密钥")
		case "三重DES":
			EncryptMode = TripleDESMode
			SelectPrivateKeyButton.Hide()
			PasswordInput.Show()
			PasswordReminderLabel.SetText("三重DES模式请输入24位密钥")
		case "AES-CBC":
			EncryptMode = AesCBCMode
			SelectPrivateKeyButton.Hide()
			PasswordInput.Show()
			PasswordReminderLabel.SetText("AES模式请输入16位密钥")
		case "AES-CTR":
			EncryptMode = AesCTRMode
			SelectPrivateKeyButton.Hide()
			PasswordInput.Show()
			PasswordReminderLabel.SetText("AES模式请输入16位密钥")
		case "RSA":
			EncryptMode = RsaMode
			SelectPrivateKeyButton.Show()
			PasswordInput.Hide()
			PasswordReminderLabel.SetText("RSA模式将自动生成公私钥")
		}
	})

	StartEncryptButton := widget.NewButton("开始加密", func() {
		if FilePath == "" || EncryptMode == 0 {
			return
		}
		switch EncryptMode {
		case DesMode:
			err := DesEncrypt(FilePath, PasswordInput.Text)
			if err != nil {
				DisplayErrorDialog(err, "加密错误")
				return
			}
		case TripleDESMode:
			err := TripleDesEncrypt(FilePath, PasswordInput.Text)
			if err != nil {
				DisplayErrorDialog(err, "加密错误")
				return
			}
		case AesCBCMode:
			err := AesCBCEncrypt(FilePath, PasswordInput.Text)
			if err != nil {
				DisplayErrorDialog(err, "加密错误")
				return
			}
		case AesCTRMode:
			err := AesCTREncrypt(FilePath, PasswordInput.Text)
			if err != nil {
				DisplayErrorDialog(err, "加密错误")
				return
			}
		case RsaMode:
			err := RsaEncrypt(FilePath)
			if err != nil {
				return
			}
		}
		dialog.NewCustom("加密成功", "确认", widget.NewLabel("加密成功"), Windows).Show()
	})

	StartDecryptButton := widget.NewButton("开始解密", func() {
		switch EncryptMode {
		case DesMode:
			err := DesDecrypt(FilePath, PasswordInput.Text)
			if err != nil {
				DisplayErrorDialog(err, "解密错误")
				return
			}
		case TripleDESMode:
			err := TripleDesDecrypt(FilePath, PasswordInput.Text)
			if err != nil {
				DisplayErrorDialog(err, "解密错误")
				return
			}
		case AesCBCMode:
			err := AesCBCDecrypt(FilePath, PasswordInput.Text)
			if err != nil {
				DisplayErrorDialog(err, "解密错误")
				return
			}
		case AesCTRMode:
			err := AesCTRDecrypt(FilePath, PasswordInput.Text)
			if err != nil {
				DisplayErrorDialog(err, "解密错误")
				return
			}
		case RsaMode:
			key, err := os.ReadFile(PrivateKeyPath)
			if err != nil {
				DisplayErrorDialog(err, "解密错误")
				return
			}
			err = RsaDecrypt(FilePath, key)
			if err != nil {
				DisplayErrorDialog(err, "解密错误")
				return
			}
		}
		dialog.NewCustom("解密成功", "确认", widget.NewLabel("解密成功"), Windows).Show()
	})

	return container.NewVBox(
		SelectFileButton,
		SelectPrivateKeyButton,
		EncryptModeSelect,
		PasswordReminderLabel,
		PasswordInput,
		container.NewHBox(StartEncryptButton, StartDecryptButton),
	)

}
