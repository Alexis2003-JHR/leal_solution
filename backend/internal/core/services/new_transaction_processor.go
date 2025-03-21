package services

import (
	"leal/internal/core/domain/models/request"
	"leal/internal/core/domain/ports"
	"log"
	"sync"
)

type TransactionService struct {
	transactionQueue chan request.ProcessTransaction
	wg               sync.WaitGroup
	repo             ports.Repository
}

func NewTransactionService(repo ports.Repository, workCount int) *TransactionService {
	tp := &TransactionService{
		transactionQueue: make(chan request.ProcessTransaction, 100),
		repo:             repo,
	}

	for i := 0; i < workCount; i++ {
		go tp.worker(i)
	}

	return tp
}

func (tp *TransactionService) AddTransaction(tx request.ProcessTransaction) error {
	tp.transactionQueue <- tx
	return nil
}

func (tp *TransactionService) worker(id int) {
	for tx := range tp.transactionQueue {
		log.Printf("[Worker %d] Procesando transacción: %+v\n", id, tx)
		err := tp.ProcessTransaction(tx)
		if err != nil {
			log.Printf("Error en id %d: %v\n", id, err)
		}
	}
}

func (tp *TransactionService) Shutdown() {
	close(tp.transactionQueue)
	tp.wg.Wait()
}
