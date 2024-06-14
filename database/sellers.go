package database

import (
	"database/sql"
	"time"
)

type Seller struct {
	ID        int     `db:"id" json:"ID"`
	Name      string  `db:"name" json:"Name"`
	Phone     string  `db:"phone" json:"Phone"`
	CreatedAt string  `db:"created_at" json:"CreatedAt"`
	UpdatedAt string  `db:"updated_at" json:"UpdatedAt,omitempty"`
	DeletedAt *string `db:"deleted_at" json:"DeletedAt,omitempty"`
}

type SellerDatabaseInterface interface {
	SelectAll() (*[]Seller, error)
	SelectByID(id int) (*Seller, error)
	Insert(seller Seller) (int, error)
	Update(seller Seller) error
	Delete(id int) error
}

type SellerDatabase struct {
	Connection *sql.DB
}

func (sb *SellerDatabase) SelectAll() (*[]Seller, error) {
	rows, err := sb.Connection.Query("SELECT id, name, phone, created_at, updated_at FROM sellers WHERE deleted_at IS NULL")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var createdAt, updatedAt time.Time
	var sellers []Seller
	for rows.Next() {
		var seller Seller
		if err := rows.Scan(&seller.ID, &seller.Name, &seller.Phone, &createdAt, &updatedAt); err != nil {
			return nil, err
		}

		seller.CreatedAt = createdAt.Format("15:04:05 02:01:06")
		seller.UpdatedAt = updatedAt.Format("15:04:05 02:01:06")
		sellers = append(sellers, seller)
	}

	return &sellers, nil
}

func (sb *SellerDatabase) SelectByID(id int) (*Seller, error) {
	var seller Seller
	var createdAt, updatedAt time.Time
	err := sb.Connection.QueryRow("SELECT id, name, phone, created_at, updated_at, deleted_at FROM sellers WHERE id = $1 AND deleted_at IS NULL", id).
		Scan(&seller.ID, &seller.Name, &seller.Phone, &createdAt, &updatedAt, &seller.DeletedAt)
	if err != nil {
		return nil, err
	}
	seller.CreatedAt = createdAt.Format("15:04:05 02:01:06")
	seller.UpdatedAt = updatedAt.Format("15:04:05 02:01:06")

	return &seller, nil
}

func (sb *SellerDatabase) Insert(seller Seller) (int, error) {
	var id int
	err := sb.Connection.QueryRow("INSERT INTO sellers (name, phone) VALUES ($1, $2) RETURNING id", seller.Name, seller.Phone).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (sb *SellerDatabase) Update(seller Seller) error {
	result, err := sb.Connection.Exec("UPDATE sellers SET name = $1, phone = $2, updated_at = $3 WHERE id = $4 AND deleted_at IS NULL", seller.Name, seller.Phone, time.Now(), seller.ID)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (sb *SellerDatabase) Delete(id int) error {
	result, err := sb.Connection.Exec("UPDATE sellers SET deleted_at = $1 WHERE id = $2 AND deleted_at IS NULL", time.Now(), id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}
