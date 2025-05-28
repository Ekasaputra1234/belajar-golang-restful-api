package helper

import (
	"strings"

	goHelper "gitlab.com/vneu/go-helper/helper"
	"gorm.io/gorm"
)

func PathApplyFilters(filters *map[string]string, tx *gorm.DB) {
	for filterKey, value := range *filters {
		switch filterKey {
		case "KandangStockOut__kandang_project_id.eq":
			tx = tx.Where("KandangStockOut.kandang_project_id = ?", value)
		default:
			continue
		}
		// Remove the filter key after processing
		delete(*filters, filterKey)
	}
}

func PathApplyFilterRecordings(filters *map[string]string, tx *gorm.DB) {
	for key, value := range *filters {
		switch key {
		case "date.eq":
			tx = tx.Where("DATE_FORMAT(recordings.date, \"%Y-%m-%d\") = ?", value)
		case "date.gte":
			tx = tx.Where("DATE_FORMAT(recordings.date, \"%Y-%m-%d\") >= ?", value)
		case "date.lte":
			tx = tx.Where("DATE_FORMAT(recordings.date, \"%Y-%m-%d\") <= ?", value)
		case "kandang_id.eq":
			tx = tx.Where("recordings.kandang_id = ?", value)
		case "kandang_project_id.eq":
			tx = tx.Where("recordings.kandang_project_id = ?", value)
		case "KandangSub__nama_kandang.like":
			tx = tx.Where("LOWER(KandangSub.nama_kandang) LIKE LOWER(?)", "%"+value+"%")
		case "kandang_sub_id.eq":
			tx = tx.Where("recordings.kandang_sub_id = ?", value)
		case "kandang_sub_id.in":
			kandangSubIdIn := strings.Split(value, ",")
			tx = tx.Where("recordings.kandang_sub_id IN (?)", kandangSubIdIn)
		}

		// Remove the filter after applying it
		delete(*filters, key)
	}
}

func FilterPaginationDataSort(paginationData *goHelper.PaginationData, tx *gorm.DB) {
	switch paginationData.OrderBy {
	case "title_project":
		tx.Order("KandangProject." + paginationData.OrderBy)
	case "nama_kandang":
		tx.Order("KandangSub." + paginationData.OrderBy)
	default:
		if paginationData.OrderBy != "" {
			tx.Order("panens." + paginationData.OrderBy)
		}
	}
}

func PathApplyFilterCustomerTransactions(filters *map[string]string, tx *gorm.DB) {
	for key, value := range *filters {
		switch key {
		case "customer_id.eq":
			tx = tx.Where("customer_transactions.customer_id = ?", value)
		case "period.eq":
			tx = tx.Where("customer_transactions.period = ?", value)
		case "panen_id.eq":
			tx = tx.Where("customer_transactions.panen_id = ?", value)
		case "status.eq":
			tx = tx.Where("customer_transactions.status = ?", value)
		case "unit_id.eq":
			tx = tx.Where("customer_transactions.unit_id = ?", value)
		case "date_time.eq":
			tx = tx.Where("DATE_FORMAT(customer_transactions.date_time,\"%Y-%m-%d\") = ?", value)
		case "date_time.gte":
			tx = tx.Where("DATE_FORMAT(customer_transactions.date_time,\"%Y-%m-%d\") >= ?", value)
		case "date_time.lte":
			tx = tx.Where("DATE_FORMAT(customer_transactions.date_time,\"%Y-%m-%d\") <= ?", value)
		case "date_time.sort":
			tx = tx.Order("customer_transactions.date_time " + value + ", customer_transactions.created_at " + value)
		case "unit_account_id.eq":
			tx = tx.Where("customer_transactions.unit_account_id = ?", value)
		case "Panen__customer_id.eq":
			tx = tx.Where("Panen.customer_id = ?", value)
		case "Panen__kandang_project_id.eq":
			tx = tx.Where("Panen.kandang_project_id = ?", value)
		case "Panen__date_time.eq":
			tx = tx.Where("Panen.customer_id = ?", value)
		case "Panen__date_time.gte":
			tx = tx.Where("DATE_FORMAT(Panen.date_time,\"%Y-%m-%d\") >= ?", value)
		case "Panen__date_time.lte":
			tx = tx.Where("DATE_FORMAT(Panen.date_time,\"%Y-%m-%d\") <= ?", value)
		}

		// Remove the filter after applying it
		delete(*filters, key)
	}
}
