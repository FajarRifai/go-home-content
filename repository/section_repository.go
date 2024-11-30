package repository

import (
	"database/sql"
	"fmt"
	"go-home-content/models"
	"log"
)

type SectionRepository struct {
	DB *sql.DB
}

func (repo *SectionRepository) CreateSection(section models.CreateSection) (int64, error) {
	tx, err := repo.DB.Begin()
	if err != nil {
		return 0, err
	}

	result, err := tx.Exec("INSERT INTO section (title, description, active, deleted) VALUES (?, ?, ?, ?)",
		section.Title, section.Description, section.Active, section.Deleted)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	sectionID, err := result.LastInsertId()
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	for _, detail := range section.ProductDetail {
		_, err := tx.Exec("INSERT INTO section_detail (code, rank, section_id) VALUES (?, ?, ?)",
			detail.Code, detail.Rank, sectionID)
		if err != nil {
			tx.Rollback()
			return 0, err
		}
	}

	err = tx.Commit()
	if err != nil {
		return 0, err
	}

	return sectionID, nil
}

func (repo *SectionRepository) GetSections() ([]models.ListSection, error) {
	query := "SELECT id, title, description, active, deleted FROM section WHERE deleted = false"
	rows, err := repo.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var sections []models.ListSection
	for rows.Next() {
		var section models.ListSection
		if err := rows.Scan(&section.ID, &section.Title, &section.Description, &section.Active, &section.Deleted); err != nil {
			return nil, err
		}
		sections = append(sections, section)
	}
	return sections, nil
}

// GetSectionById retrieves a section record by its ID.
func (repo *SectionRepository) GetSectionById(id int) (models.Section, error) {
	var section models.Section
	query := "SELECT title, description, active, deleted FROM section WHERE id = ? AND deleted = false"
	err := repo.DB.QueryRow(query, id).Scan(&section.Title, &section.Description, &section.Active, &section.Deleted)
	if err != nil {
		if err == sql.ErrNoRows {
			return section, fmt.Errorf("section not found")
		}
		return section, fmt.Errorf("error fetching section: %v", err)
	}
	return section, nil
}

type CodeProduct struct {
	Code string
}

// GetSectionById retrieves a section record by its ID.
func (repo *SectionRepository) GetSectionDetailById(id int) ([]string, error) {
	// SQL query to get codes for a specific section_id
	query := "SELECT code FROM section_detail WHERE section_id = ?"

	// Execute the query
	rows, err := repo.DB.Query(query, id)
	if err != nil {
		log.Fatal("Error executing query:", err)
		return nil, err // return nil and error if query fails
	}
	defer rows.Close()

	// Slice to store the result codes
	var codes []string

	// Iterate through the rows and scan the data into the codes slice
	for rows.Next() {
		var code string
		if err := rows.Scan(&code); err != nil {
			log.Fatal("Error scanning row:", err)
			return nil, err // return nil and error if scanning fails
		}
		// Append the code to the string slice
		codes = append(codes, code)
	}

	// Check for any errors from iterating over rows
	if err := rows.Err(); err != nil {
		log.Fatal("Error iterating rows:", err)
		return nil, err // return nil and error if iterating fails
	}

	// Return the list of codes
	return codes, nil
}
