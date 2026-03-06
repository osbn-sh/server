package manipulationRepository

import (
	"context"
	"errors"
	"ostadbun/pkg/richerror"
	"time"
)

var (
	ErrMajorNotFound      = errors.New("major not found")
	ErrInvalidMajorStatus = errors.New("invalid status for major: must be 'approved' or 'rejected'")
)

// ApproveMajor approves a pending major
func (d DB) ApproveMajor(ctx context.Context, pendingMajorID int64, approvedBy int64) error {
	return d.updateMajorStatus(ctx, pendingMajorID, "approved", approvedBy, nil)
}

// RejectMajor rejects a pending major with optional rejection reason
func (d DB) RejectMajor(ctx context.Context, pendingMajorID int64, rejectedBy int64, rejectionReason *string) error {
	return d.updateMajorStatus(ctx, pendingMajorID, "rejected", rejectedBy, rejectionReason)
}

// updateMajorStatus is a helper method for updating major status
func (d DB) updateMajorStatus(
	ctx context.Context,
	pendingMajorID int64,
	status string,
	approvedBy int64,
	rejectionReason *string,
) error {
	if status != "approved" && status != "rejected" {
		return richerror.New("manipulationRepository-updateMajorStatus").WithErr(ErrInvalidMajorStatus).WithKind(richerror.KindUnexpected).WithMessage("error on update major status and not found")
	}

	query := `
		UPDATE pending_major 
		SET 
			status = $1,
			approved_by = $2,
			approved_at = $3,
			rejection_reason = $4
		WHERE id = $5 AND status = 'pending'
	`

	result, err := d.conn.Conn().ExecContext(ctx, query,
		status,
		approvedBy,
		time.Now().UTC(),
		rejectionReason,
		pendingMajorID,
	)
	if err != nil {
		return richerror.New("manipulationRepository-updateMajorStatus").WithErr(err).WithKind(richerror.KindUnexpected).WithMessage("error on update major status")
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return richerror.New("manipulationRepository-updateMajorStatus").WithErr(err).WithKind(richerror.KindUnexpected).WithMessage("error on update major status")
	}
	if rowsAffected == 0 {
		return richerror.New("manipulationRepository-updateMajorStatus").WithErr(ErrMajorNotFound).WithKind(richerror.KindUnexpected).WithMessage("error on update major status and not found")
	}

	return nil
}
