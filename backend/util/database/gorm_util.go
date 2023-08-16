package database

import (
	"github.com/fatih/structs"
	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/gorm"

	"github.com/KScaesar/jubo-homework/backend/util"
	"github.com/KScaesar/jubo-homework/backend/util/errors"
)

// GormFilterDsl
// filter 可傳入 struct, 然後會解析 tag key = rdb 的內容,
// 其 field 數值 若為 go 的初始值 會忽略此條件,
// 需要在 struct tag 填入 適當的 sql where 敘述.
//
// 若想查 xxx=false, 將型別定義為 pointer *bool;
// 若想查 xxx is null, 將型別定義為 string;
// 理論上全部數值可以都用 string 來表示.
//
// example:
//
//	type DtoFilterUserOption struct {
//	    UserIDs     []int     `rdb:"user_id in (?)"`
//	    IsAdmin     *bool     `rdb:"is_admin = ?"`
//	    BlockedTime time.Time `rdb:"blocked_time is ?"`
//	    IsBlocked   string    `rdb:"is_blocked = ?"`
//	}
//
//	filter := DtoFilterUserOption{
//	    UserIDs:     []int{2, 4},
//	    IsAdmin:     &BoolFalse,
//	    BlockedTime: time.Now(),
//	    IsBlocked:   "false",
//	}
func GormFilterDsl(filter interface{}) []func(db *gorm.DB) *gorm.DB {
	if filter == nil {
		return nil
	}

	fields := structs.New(filter).Fields()
	where := make([]func(db *gorm.DB) *gorm.DB, 0, len(fields))

	for _, field := range fields {
		if field.IsZero() {
			continue
		}

		value := field.Value()
		if field.IsEmbedded() {
			embed := GormFilterDsl(value)
			where = append(where, embed...)
			continue
		}

		sql := field.Tag("rdb")
		where = append(
			where, func(db *gorm.DB) *gorm.DB {
				return db.Where(sql, value)
			},
		)
	}

	return where
}

func GormError(err error) error {
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		return PgsqlError(pgErr)
	}

	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		return errors.ErrNotFound
	default:
		return errors.Join3rdParty(errors.ErrSystem, err)
	}
}

func GormQueryListWithPagination[R any](
	db *gorm.DB, dtoFilter any, dtoPage util.PageParam, dtoSort util.SortParam,
) (
	resp util.ListResponse[R], err error,
) {
	db = db.Scopes(GormFilterDsl(dtoFilter)...)
	var totalSize int64
	err = db.Count(&totalSize).Error
	if err != nil {
		err = GormError(err)
		return
	}

	db = db.
		Offset(int(dtoPage.OffsetOrSkip())).
		Limit(int(dtoPage.Size)).
		Order(dtoSort.KeyValueString())

	list := make([]R, 0, dtoPage.Size)
	err = db.Find(&list).Error
	if err != nil {
		err = GormError(err)
		return
	}

	pageResponse := util.NewPageResponse(dtoPage, uint64(totalSize))
	return util.NewListResponseWithPagination(list, &pageResponse), nil
}

func GormQueryListWithoutPagination[R any](
	db *gorm.DB, dtoFilter any, dtoSort util.SortParam,
) (
	resp util.ListResponse[R], err error,
) {
	const defaultSize = 50
	list := make([]R, 0, defaultSize)

	err = db.
		Scopes(GormFilterDsl(dtoFilter)...).
		Order(dtoSort.KeyValueString()).
		Find(&list).Error
	if err != nil {
		err = GormError(err)
		return
	}

	return util.NewListResponseWithoutPagination(list), nil
}
