package pb

import sharedpb "github.com/MamangRust/microservice-ecommerce-shared/pb"

// Re-export all types from shared/pb for backward compatibility
type PaginationMeta = sharedpb.PaginationMeta
type Empty = sharedpb.Empty

// Auth
type AuthServiceClient = sharedpb.AuthServiceClient

func NewAuthServiceClient(cc interface{ ClientConn interface{} }) AuthServiceClient {
	return sharedpb.NewAuthServiceClient(nil)
}
