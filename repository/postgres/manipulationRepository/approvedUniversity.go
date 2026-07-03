package manipulationRepository

import (
	"context"
	"errors"
	"ostadbun/pkg/richerror"
	"time"
)

var (
	ErrUniversityNotFound      = errors.New("university not found")
	ErrInvalidUniversityStatus = errors.New("invalid status for university: must be 'approved' or 'rejected'")
)

// ApproveUniversity approves a pending university
func (d DB) ApproveUniversity(ctx context.Context, pendingUniversityID int64, approvedBy int64) error {
	return d.updateUniversityStatus(ctx, pendingUniversityID, "approved", approvedBy, nil)
}

// RejectUniversity rejects a pending university with optional rejection reason
func (d DB) RejectUniversity(ctx context.Context, pendingUniversityID int64, rejectedBy int64, rejectionReason *string) error {
	return d.updateUniversityStatus(ctx, pendingUniversityID, "rejected", rejectedBy, rejectionReason)
}

// updateUniversityStatus is a helper method for updating university status
func (d DB) updateUniversityStatus(
	ctx context.Context,
	pendingUniversityID int64,
	status string,
	approvedBy int64,
	rejectionReason *string,
) error {
	if status != "approved" && status != "rejected" {
		return richerror.New("manipulationRepository-updateUniversityStatus").WithErr(ErrInvalidUniversityStatus).WithKind(richerror.KindUnexpected).WithMessage("error on update university status and invalid")
	}

	query := `
		UPDATE pending_university 
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
		pendingUniversityID,
	)
	if err != nil {
		return richerror.New("manipulationRepository-updateUniversityStatus").WithErr(err).WithKind(richerror.KindUnexpected).WithMessage("error on update university status ")
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return richerror.New("manipulationRepository-updateUniversityStatus").WithErr(err).WithKind(richerror.KindUnexpected).WithMessage("error on update university status ")
	}
	if rowsAffected == 0 {
		return richerror.New("manipulationRepository-updateUniversityStatus").WithErr(ErrUniversityNotFound).WithKind(richerror.KindUnexpected).WithMessage("error on update university status and not found")
	}

	return nil
}
