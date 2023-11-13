package facade

type IWalletFacade interface {
	addMoneyToWallet()
	deductMoneyFromWallet()
}

type WalletFacade struct {
}

func (w *WalletFacade) addMoneyToWallet() {
	//检查账户
	//检查code
	//检查余额
	//发送通知
	//更新余额
}

func (w *WalletFacade) deductMoneyToWallet() {
	//检查账户
	//检查code
	//检查余额
	//发送通知
	//更新余额
}
