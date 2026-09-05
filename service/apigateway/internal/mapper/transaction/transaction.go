package transactiongraphqlmapper

import (
	"github.com/MamangRust/microservice-ecommerce-shared/pb"
	graphqlmapper "github.com/MamangRust/monolith-graphql-ecommerce-apigateway/internal/mapper/pagination"
	"github.com/MamangRust/monolith-graphql-ecommerce-apigateway/internal/model"
)

type transactionResponseMapper struct{}

func NewTransactionResponseMapper() *transactionResponseMapper {
	return &transactionResponseMapper{}
}

func (t *transactionResponseMapper) ToGraphqlResponseTransaction(res *pb.ApiResponseTransaction) *model.APIResponseTransaction {
	return &model.APIResponseTransaction{
		Status:  res.Status,
		Message: res.Message,
		Data:    t.mapResponseTransaction(res.Data),
	}
}

func (t *transactionResponseMapper) ToGraphqlResponseTransactionDeleteAt(res *pb.ApiResponseTransactionDeleteAt) *model.APIResponseTransactionDeleteAt {
	return &model.APIResponseTransactionDeleteAt{
		Status:  res.Status,
		Message: res.Message,
		Data:    t.mapResponseTransactionDeleteAt(res.Data),
	}
}

func (t *transactionResponseMapper) ToGraphqlResponseTransactionDelete(res *pb.ApiResponseTransactionDelete) *model.APIResponseTransactionDelete {
	return &model.APIResponseTransactionDelete{
		Status:  res.Status,
		Message: res.Message,
	}
}

func (t *transactionResponseMapper) ToGraphqlResponseTransactionAll(res *pb.ApiResponseTransactionAll) *model.APIResponseTransactionAll {
	return &model.APIResponseTransactionAll{
		Status:  res.Status,
		Message: res.Message,
	}
}

func (t *transactionResponseMapper) ToGraphqlResponsePaginationTransactionDeleteAt(
	res *pb.ApiResponsePaginationTransactionDeleteAt,
) *model.APIResponsePaginationTransactionDeleteAt {
	return &model.APIResponsePaginationTransactionDeleteAt{
		Status:     res.Status,
		Message:    res.Message,
		Data:       t.mapResponsesTransactionDeleteAt(res.Data),
		Pagination: graphqlmapper.MapPaginationMeta(res.Pagination),
	}
}

func (t *transactionResponseMapper) ToGraphqlResponsePaginationTransaction(
	res *pb.ApiResponsePaginationTransaction,
) *model.APIResponsePaginationTransaction {
	return &model.APIResponsePaginationTransaction{
		Status:     res.Status,
		Message:    res.Message,
		Data:       t.mapResponsesTransaction(res.Data),
		Pagination: graphqlmapper.MapPaginationMeta(res.Pagination),
	}
}

func (t *transactionResponseMapper) ToGraphqlResponseMonthAmountSuccess(res *pb.ApiResponseTransactionMonthAmountSuccess) *model.APIResponseTransactionMonthAmountSuccess {
	return &model.APIResponseTransactionMonthAmountSuccess{
		Status:  res.Status,
		Message: res.Message,
		Data:    t.mapResponsesTransactionMonthlyAmountSuccess(res.Data),
	}
}

func (t *transactionResponseMapper) ToGraphqlResponseYearAmountSuccess(res *pb.ApiResponseTransactionYearAmountSuccess) *model.APIResponseTransactionYearAmountSuccess {
	return &model.APIResponseTransactionYearAmountSuccess{
		Status:  res.Status,
		Message: res.Message,
		Data:    t.mapResponsesTransactionYearlyAmountSuccess(res.Data),
	}
}

func (t *transactionResponseMapper) ToGraphqlResponseMonthAmountFailed(res *pb.ApiResponseTransactionMonthAmountFailed) *model.APIResponseTransactionMonthAmountFailed {
	return &model.APIResponseTransactionMonthAmountFailed{
		Status:  res.Status,
		Message: res.Message,
		Data:    t.mapResponsesTransactionMonthlyAmountFailed(res.Data),
	}
}

func (t *transactionResponseMapper) ToGraphqlResponseYearAmountFailed(res *pb.ApiResponseTransactionYearAmountFailed) *model.APIResponseTransactionYearAmountFailed {
	return &model.APIResponseTransactionYearAmountFailed{
		Status:  res.Status,
		Message: res.Message,
		Data:    t.mapResponsesTransactionYearlyAmountFailed(res.Data),
	}
}

func (t *transactionResponseMapper) ToGraphqlResponseMonthMethod(res *pb.ApiResponseTransactionMonthPaymentMethod) *model.APIResponseTransactionMonthPaymentMethod {
	return &model.APIResponseTransactionMonthPaymentMethod{
		Status:  res.Status,
		Message: res.Message,
		Data:    t.mapResponsesTransactionMonthlyMethod(res.Data),
	}
}

func (t *transactionResponseMapper) ToGraphqlResponseYearMethod(res *pb.ApiResponseTransactionYearPaymentmethod) *model.APIResponseTransactionYearPaymentMethod {
	return &model.APIResponseTransactionYearPaymentMethod{
		Status:  res.Status,
		Message: res.Message,
		Data:    t.mapResponsesTransactionYearlyMethod(res.Data),
	}
}

func (t *transactionResponseMapper) mapResponseTransaction(transaction *pb.TransactionResponse) *model.TransactionResponse {
	if transaction == nil {
		return nil
	}
	return &model.TransactionResponse{
		ID:            int32(transaction.Id),
		OrderID:       int32(transaction.OrderId),
		MerchantID:    int32(transaction.MerchantId),
		PaymentMethod: transaction.PaymentMethod,
		Amount:        int32(transaction.Amount),
		PaymentStatus: transaction.PaymentStatus,
		CreatedAt:     transaction.CreatedAt,
		UpdatedAt:     transaction.UpdatedAt,
	}
}

func (t *transactionResponseMapper) mapResponsesTransaction(transactions []*pb.TransactionResponse) []*model.TransactionResponse {
	mapped := make([]*model.TransactionResponse, 0, len(transactions))
	for _, transaction := range transactions {
		mapped = append(mapped, t.mapResponseTransaction(transaction))
	}
	return mapped
}

func (t *transactionResponseMapper) mapResponseTransactionDeleteAt(transaction *pb.TransactionResponseDeleteAt) *model.TransactionResponseDeleteAt {
	var deletedAt string

	if transaction.DeletedAt != nil {
		deletedAt = transaction.DeletedAt.Value
	}

	return &model.TransactionResponseDeleteAt{
		ID:            int32(transaction.Id),
		OrderID:       int32(transaction.OrderId),
		MerchantID:    int32(transaction.MerchantId),
		PaymentMethod: transaction.PaymentMethod,
		Amount:        int32(transaction.Amount),
		PaymentStatus: transaction.PaymentStatus,
		CreatedAt:     transaction.CreatedAt,
		UpdatedAt:     transaction.UpdatedAt,
		DeletedAt:     deletedAt,
	}
}

func (t *transactionResponseMapper) mapResponsesTransactionDeleteAt(transactions []*pb.TransactionResponseDeleteAt) []*model.TransactionResponseDeleteAt {
	mapped := make([]*model.TransactionResponseDeleteAt, 0, len(transactions))
	for _, transaction := range transactions {
		mapped = append(mapped, t.mapResponseTransactionDeleteAt(transaction))
	}
	return mapped
}

func (t *transactionResponseMapper) mapResponseTransactionMonthAmountSuccess(row *pb.TransactionMonthlyAmountSuccess) *model.TransactionMonthlyAmountSuccess {
	return &model.TransactionMonthlyAmountSuccess{
		Year:         row.Year,
		Month:        row.Month,
		TotalSuccess: int32(row.TotalSuccess),
		TotalAmount:  int32(row.TotalAmount),
	}
}

func (t *transactionResponseMapper) mapResponsesTransactionMonthlyAmountSuccess(rows []*pb.TransactionMonthlyAmountSuccess) []*model.TransactionMonthlyAmountSuccess {
	var res []*model.TransactionMonthlyAmountSuccess
	for _, row := range rows {
		res = append(res, t.mapResponseTransactionMonthAmountSuccess(row))
	}
	return res
}

func (t *transactionResponseMapper) mapResponseTransactionYearAmountSuccess(row *pb.TransactionYearlyAmountSuccess) *model.TransactionYearlyAmountSuccess {
	return &model.TransactionYearlyAmountSuccess{
		Year:         row.Year,
		TotalSuccess: int32(row.TotalSuccess),
		TotalAmount:  int32(row.TotalAmount),
	}
}

func (t *transactionResponseMapper) mapResponsesTransactionYearlyAmountSuccess(rows []*pb.TransactionYearlyAmountSuccess) []*model.TransactionYearlyAmountSuccess {
	var res []*model.TransactionYearlyAmountSuccess
	for _, row := range rows {
		res = append(res, t.mapResponseTransactionYearAmountSuccess(row))
	}
	return res
}

func (t *transactionResponseMapper) mapResponseTransactionMonthAmountFailed(row *pb.TransactionMonthlyAmountFailed) *model.TransactionMonthlyAmountFailed {
	return &model.TransactionMonthlyAmountFailed{
		Year:        row.Year,
		Month:       row.Month,
		TotalFailed: int32(row.TotalFailed),
		TotalAmount: int32(row.TotalAmount),
	}
}

func (t *transactionResponseMapper) mapResponsesTransactionMonthlyAmountFailed(rows []*pb.TransactionMonthlyAmountFailed) []*model.TransactionMonthlyAmountFailed {
	var res []*model.TransactionMonthlyAmountFailed
	for _, row := range rows {
		res = append(res, t.mapResponseTransactionMonthAmountFailed(row))
	}
	return res
}

func (t *transactionResponseMapper) mapResponseTransactionYearAmountFailed(row *pb.TransactionYearlyAmountFailed) *model.TransactionYearlyAmountFailed {
	return &model.TransactionYearlyAmountFailed{
		Year:        row.Year,
		TotalFailed: int32(row.TotalFailed),
		TotalAmount: int32(row.TotalAmount),
	}
}

func (t *transactionResponseMapper) mapResponsesTransactionYearlyAmountFailed(rows []*pb.TransactionYearlyAmountFailed) []*model.TransactionYearlyAmountFailed {
	var res []*model.TransactionYearlyAmountFailed
	for _, row := range rows {
		res = append(res, t.mapResponseTransactionYearAmountFailed(row))
	}
	return res
}

func (t *transactionResponseMapper) mapResponseTransactionMonthMethod(row *pb.TransactionMonthlyMethod) *model.TransactionMonthlyMethod {
	return &model.TransactionMonthlyMethod{
		Month:             row.Month,
		PaymentMethod:     row.PaymentMethod,
		TotalTransactions: int32(row.TotalTransactions),
		TotalAmount:       int32(row.TotalAmount),
	}
}

func (t *transactionResponseMapper) mapResponsesTransactionMonthlyMethod(rows []*pb.TransactionMonthlyMethod) []*model.TransactionMonthlyMethod {
	var res []*model.TransactionMonthlyMethod
	for _, row := range rows {
		res = append(res, t.mapResponseTransactionMonthMethod(row))
	}
	return res
}

func (t *transactionResponseMapper) mapResponseTransactionYearMethod(row *pb.TransactionYearlyMethod) *model.TransactionYearlyMethod {
	return &model.TransactionYearlyMethod{
		Year:              row.Year,
		PaymentMethod:     row.PaymentMethod,
		TotalTransactions: int32(row.TotalTransactions),
		TotalAmount:       int32(row.TotalAmount),
	}
}

func (t *transactionResponseMapper) mapResponsesTransactionYearlyMethod(rows []*pb.TransactionYearlyMethod) []*model.TransactionYearlyMethod {
	var res []*model.TransactionYearlyMethod
	for _, row := range rows {
		res = append(res, t.mapResponseTransactionYearMethod(row))
	}
	return res
}
