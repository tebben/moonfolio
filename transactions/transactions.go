package transactions

//ToDo: Channels for sending created, deleted, editted transactions

func GetTransactions() ([]*Transaction, error) {
	return mockTransactions(), nil
}

func mockTransactions() []*Transaction {
	t1 := &Transaction{ID: 0, Type: TransactionBuy, CoinID: "BTC", CoinAmount: 0.084, DateTime: 1515940982305, PriceUSD: 243.93}
	t2 := &Transaction{ID: 1, Type: TransactionBuy, CoinID: "VSX", CoinAmount: 11071, DateTime: 1515940982305, PriceUSD: 0}
	t3 := &Transaction{ID: 2, Type: TransactionBuy, CoinID: "HUSH", CoinAmount: 100, DateTime: 1515940982305, PriceUSD: 0}
	t4 := &Transaction{ID: 3, Type: TransactionBuy, CoinID: "XRP", CoinAmount: 49.95, DateTime: 1515940982305, PriceUSD: 160}
	t5 := &Transaction{ID: 4, Type: TransactionBuy, CoinID: "ADA", CoinAmount: 132.87, DateTime: 1515940982305, PriceUSD: 160}
	t6 := &Transaction{ID: 5, Type: TransactionBuy, CoinID: "IOT", CoinAmount: 160.91, DateTime: 1515940982305, PriceUSD: 400}
	t7 := &Transaction{ID: 6, Type: TransactionBuy, CoinID: "NLG", CoinAmount: 1016.98, DateTime: 1515940982305, PriceUSD: 80}
	t8 := &Transaction{ID: 7, Type: TransactionBuy, CoinID: "XVG", CoinAmount: 694.30, DateTime: 1515940982305, PriceUSD: 70}
	t9 := &Transaction{ID: 8, Type: TransactionBuy, CoinID: "ECC", CoinAmount: 46850, DateTime: 1515940982305, PriceUSD: 70}
	t10 := &Transaction{ID: 9, Type: TransactionBuy, CoinID: "SHND", CoinAmount: 6917600, DateTime: 1515940982305, PriceUSD: 40}
	t11 := &Transaction{ID: 10, Type: TransactionBuy, CoinID: "HUSH", CoinAmount: 1, DateTime: 1515940982305, PriceUSD: 0}
	t12 := &Transaction{ID: 11, Type: TransactionBuy, CoinID: "XRB", CoinAmount: 11.39, DateTime: 1515940982305, PriceUSD: 0}
	t13 := &Transaction{ID: 12, Type: TransactionBuy, CoinID: "VEN", CoinAmount: 20, DateTime: 1515940982305, PriceUSD: 0}
	t14 := &Transaction{ID: 13, Type: TransactionBuy, CoinID: "ENJ", CoinAmount: 450, DateTime: 1515940982305, PriceUSD: 0}
	t15 := &Transaction{ID: 14, Type: TransactionBuy, CoinID: "BAT", CoinAmount: 185, DateTime: 1515940982305, PriceUSD: 0}
	t16 := &Transaction{ID: 15, Type: TransactionBuy, CoinID: "QASH", CoinAmount: 56.14, DateTime: 1515940982305, PriceUSD: 0}
	t17 := &Transaction{ID: 16, Type: TransactionBuy, CoinID: "ETHOS", CoinAmount: 15, DateTime: 1515940982305, PriceUSD: 0}
	t18 := &Transaction{ID: 17, Type: TransactionBuy, CoinID: "LUX", CoinAmount: 1, DateTime: 1515940982305, PriceUSD: 0}

	return []*Transaction{
		t1, t2, t3, t4, t5, t6, t7, t8, t9, t10, t11, t12, t13, t14, t15, t16, t17, t18,
	}
}
