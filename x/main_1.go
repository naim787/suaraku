package main

import (
	"fmt"
	"os"

	"github.com/tendermint/tendermint/abci/server"
	"github.com/tendermint/tendermint/abci/types"
)

// MySimpleApp adalah implementasi ABCI yang sangat sederhana.
type MySimpleApp struct {
	types.BaseApplication
}

// CheckTx: Logika untuk memeriksa validitas transaksi sebelum masuk ke mempool.
func (app *MySimpleApp) CheckTx(req types.RequestCheckTx) types.ResponseCheckTx {
	fmt.Printf("Menerima CheckTx dengan string: %s\n", string(req.Tx))
	return types.ResponseCheckTx{Code: 0, Data: req.Tx}
}

// DeliverTx: Logika untuk memproses transaksi yang sudah divalidasi dan diterima.
func (app *MySimpleApp) DeliverTx(req types.RequestDeliverTx) types.ResponseDeliverTx {
	fmt.Printf("Menerima transaksi: %s\n", string(req.Tx))
	return types.ResponseDeliverTx{Code: 0}
}

// Commit: Logika untuk menyimpan state aplikasi ke disk (jika ada).
func (app *MySimpleApp) Commit() types.ResponseCommit {
	return types.ResponseCommit{Data: nil}
}

// Main function untuk memulai server ABCI.
func main() {
	// Buat instance aplikasi
	app := &MySimpleApp{}

	// Buat server ABCI
	srv, err := server.NewServer("tcp://0.0.0.0:26658", "socket", app)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	// Mulai server
	if err := srv.Start(); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	// Tunggu sampai program dihentikan
	select {}
}
