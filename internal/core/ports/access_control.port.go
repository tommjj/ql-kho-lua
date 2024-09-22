package ports

import "context"

type IAccessControlRepository interface {
	// HasAccess check if user has access
	HasAccess(ctx context.Context, storeHouseID int, userID int) error
	// SetAccess set access for user
	SetAccess(ctx context.Context, storeHouseID int, userID int) error
	// DelAccess remove user access
	DelAccess(ctx context.Context, storeHouseID int, userID int) error
}

type IAccessControlService interface {
	// HasAccess check if user has access
	HasAccess(ctx context.Context, storeHouseID int, userID int) error
	// SetAccess set access for user
	SetAccess(ctx context.Context, storeHouseID int, userID int) error
	// DelAccess remove user access
	DelAccess(ctx context.Context, storeHouseID int, userID int) error
}
