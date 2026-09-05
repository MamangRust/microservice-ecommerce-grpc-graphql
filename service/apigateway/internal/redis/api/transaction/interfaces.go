package transaction_cache

import (
	"context"

	"github.com/MamangRust/monolith-graphql-ecommerce-apigateway/internal/model"
)

type TransactionStatsCache interface {
	GetCachedMonthAmountSuccessCached(ctx context.Context, req *model.FindMonthlyTransactionStatus) (*model.APIResponseTransactionMonthAmountSuccess, bool)
	SetCachedMonthAmountSuccessCached(ctx context.Context, req *model.FindMonthlyTransactionStatus, data *model.APIResponseTransactionMonthAmountSuccess)

	GetCachedYearAmountSuccessCached(ctx context.Context, year int) (*model.APIResponseTransactionYearAmountSuccess, bool)
	SetCachedYearAmountSuccessCached(ctx context.Context, year int, data *model.APIResponseTransactionYearAmountSuccess)

	GetCachedMonthAmountFailedCached(ctx context.Context, req *model.FindMonthlyTransactionStatus) (*model.APIResponseTransactionMonthAmountFailed, bool)
	SetCachedMonthAmountFailedCached(ctx context.Context, req *model.FindMonthlyTransactionStatus, data *model.APIResponseTransactionMonthAmountFailed)

	GetCachedYearAmountFailedCached(ctx context.Context, year int) (*model.APIResponseTransactionYearAmountFailed, bool)
	SetCachedYearAmountFailedCached(ctx context.Context, year int, data *model.APIResponseTransactionYearAmountFailed)

	GetCachedMonthMethodSuccessCached(ctx context.Context, req *model.MonthTransactionMethod) (*model.APIResponseTransactionMonthPaymentMethod, bool)
	SetCachedMonthMethodSuccessCached(ctx context.Context, req *model.MonthTransactionMethod, data *model.APIResponseTransactionMonthPaymentMethod)

	GetCachedYearMethodSuccessCached(ctx context.Context, year int) (*model.APIResponseTransactionYearPaymentMethod, bool)
	SetCachedYearMethodSuccessCached(ctx context.Context, year int, data *model.APIResponseTransactionYearPaymentMethod)

	GetCachedMonthMethodFailedCached(ctx context.Context, req *model.MonthTransactionMethod) (*model.APIResponseTransactionMonthPaymentMethod, bool)
	SetCachedMonthMethodFailedCached(ctx context.Context, req *model.MonthTransactionMethod, data *model.APIResponseTransactionMonthPaymentMethod)

	GetCachedYearMethodFailedCached(ctx context.Context, year int) (*model.APIResponseTransactionYearPaymentMethod, bool)
	SetCachedYearMethodFailedCached(ctx context.Context, year int, data *model.APIResponseTransactionYearPaymentMethod)
}

type TransactionStatsByMerchantCache interface {
	GetCachedMonthAmountSuccessByMerchant(ctx context.Context, req *model.FindMonthlyTransactionStatusByMerchant) (*model.APIResponseTransactionMonthAmountSuccess, bool)
	SetCachedMonthAmountSuccessByMerchant(ctx context.Context, req *model.FindMonthlyTransactionStatusByMerchant, data *model.APIResponseTransactionMonthAmountSuccess)

	GetCachedYearAmountSuccessByMerchant(ctx context.Context, req *model.FindYearlyTransactionStatusByMerchant) (*model.APIResponseTransactionYearAmountSuccess, bool)
	SetCachedYearAmountSuccessByMerchant(ctx context.Context, req *model.FindYearlyTransactionStatusByMerchant, data *model.APIResponseTransactionYearAmountSuccess)

	GetCachedMonthAmountFailedByMerchant(ctx context.Context, req *model.FindMonthlyTransactionStatusByMerchant) (*model.APIResponseTransactionMonthAmountFailed, bool)
	SetCachedMonthAmountFailedByMerchant(ctx context.Context, req *model.FindMonthlyTransactionStatusByMerchant, data *model.APIResponseTransactionMonthAmountFailed)

	GetCachedYearAmountFailedByMerchant(ctx context.Context, req *model.FindYearlyTransactionStatusByMerchant) (*model.APIResponseTransactionYearAmountFailed, bool)
	SetCachedYearAmountFailedByMerchant(ctx context.Context, req *model.FindYearlyTransactionStatusByMerchant, data *model.APIResponseTransactionYearAmountFailed)

	GetCachedMonthMethodSuccessByMerchant(ctx context.Context, req *model.MonthTransactionMethodByMerchant) (*model.APIResponseTransactionMonthPaymentMethod, bool)
	SetCachedMonthMethodSuccessByMerchant(ctx context.Context, req *model.MonthTransactionMethodByMerchant, data *model.APIResponseTransactionMonthPaymentMethod)

	GetCachedYearMethodSuccessByMerchant(ctx context.Context, req *model.YearTransactionMethodByMerchant) (*model.APIResponseTransactionYearPaymentMethod, bool)
	SetCachedYearMethodSuccessByMerchant(ctx context.Context, req *model.YearTransactionMethodByMerchant, data *model.APIResponseTransactionYearPaymentMethod)

	GetCachedMonthMethodFailedByMerchant(ctx context.Context, req *model.MonthTransactionMethodByMerchant) (*model.APIResponseTransactionMonthPaymentMethod, bool)
	SetCachedMonthMethodFailedByMerchant(ctx context.Context, req *model.MonthTransactionMethodByMerchant, data *model.APIResponseTransactionMonthPaymentMethod)

	GetCachedYearMethodFailedByMerchant(ctx context.Context, req *model.YearTransactionMethodByMerchant) (*model.APIResponseTransactionYearPaymentMethod, bool)
	SetCachedYearMethodFailedByMerchant(ctx context.Context, req *model.YearTransactionMethodByMerchant, data *model.APIResponseTransactionYearPaymentMethod)
}

type TransactionQueryCache interface {
	GetCachedTransactionsCache(ctx context.Context, req *model.FindAllTransactionRequest) (*model.APIResponsePaginationTransaction, bool)
	SetCachedTransactionsCache(ctx context.Context, req *model.FindAllTransactionRequest, data *model.APIResponsePaginationTransaction)

	GetCachedTransactionByMerchant(ctx context.Context, req *model.FindAllTransactionMerchantRequest) (*model.APIResponsePaginationTransaction, bool)
	SetCachedTransactionByMerchant(ctx context.Context, req *model.FindAllTransactionMerchantRequest, data *model.APIResponsePaginationTransaction)

	GetCachedTransactionActiveCache(ctx context.Context, req *model.FindAllTransactionRequest) (*model.APIResponsePaginationTransaction, bool)
	SetCachedTransactionActiveCache(ctx context.Context, req *model.FindAllTransactionRequest, data *model.APIResponsePaginationTransaction)

	GetCachedTransactionTrashedCache(ctx context.Context, req *model.FindAllTransactionRequest) (*model.APIResponsePaginationTransactionDeleteAt, bool)
	SetCachedTransactionTrashedCache(ctx context.Context, req *model.FindAllTransactionRequest, data *model.APIResponsePaginationTransactionDeleteAt)

	GetCachedTransactionCache(ctx context.Context, id int) (*model.APIResponseTransaction, bool)
	SetCachedTransactionCache(ctx context.Context, data *model.APIResponseTransaction)

	GetCachedTransactionByOrderId(ctx context.Context, orderID int) (*model.APIResponseTransaction, bool)
	SetCachedTransactionByOrderId(ctx context.Context, orderID int, data *model.APIResponseTransaction)
}

type TransactionCommandCache interface {
	DeleteTransactionCache(ctx context.Context, transactionID int)
}
