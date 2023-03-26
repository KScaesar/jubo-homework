package database

import (
	"fmt"
	"reflect"

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
		where = append(where, func(db *gorm.DB) *gorm.DB {
			return db.Where(sql, value)
		})
	}

	return where
}

func GormSortDsl(sort interface{}) []func(db *gorm.DB) *gorm.DB {
	if sort == nil {
		return nil
	}

	fields := structs.New(sort).Fields()
	orderBy := make([]func(db *gorm.DB) *gorm.DB, 0, len(fields))

	for _, field := range fields {
		if field.IsZero() {
			continue
		}

		var value any
		if field.Kind() == reflect.Pointer {
			value = reflect.ValueOf(field.Value()).Elem().Interface()
		} else {
			value = field.Value()
		}

		if field.IsEmbedded() {
			embed := GormSortDsl(value)
			orderBy = append(orderBy, embed...)
			continue
		}

		sqlName := field.Tag("rdb")
		orderBy = append(orderBy, func(db *gorm.DB) *gorm.DB {
			return db.Order(fmt.Sprintf("%v %s", sqlName, value))
		})
	}

	return orderBy
}

func GormPageDsl(
	originDB *gorm.DB,
	pageOption util.DtoPageParam,
) (
	modifiedDB *gorm.DB,
	listBufferSize int64,
	pageResponse *util.DtoPageResponse,
	err error,
) {
	modifiedDB = originDB
	const defaultBufferSize = 10
	listBufferSize = defaultBufferSize

	if !pageOption.RestrictedSize() {
		return
	}

	err = pageOption.Validate()
	if err != nil {
		return
	}

	var totalSize int64
	err = originDB.Count(&totalSize).Error
	if err != nil {
		err = GormError(err)
		return
	}

	listBufferSize = *pageOption.Size
	pageResponse, _ = util.NewDtoPageResponse(pageOption, totalSize)

	modifiedDB = originDB.
		Offset(int(pageOption.OffsetOrSkip())).
		Limit(int(*pageOption.Size))

	return
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

func GormQueryListFromSingleTable[R any](db *gorm.DB, tableName string, dto util.DtoQryParam) (util.DtoListResponse[R], error) {
	if dto == nil {
		return util.DtoListResponse[R]{}, errors.WrapWithMessage(errors.ErrDeveloper, "dto should not be nil")
	}

	// filter
	where := GormFilterDsl(dto.FilterParam())
	db = db.Table(tableName).Scopes(where...)

	// page
	db, listBufferSize, pageResponse, err := GormPageDsl(db, dto.PageParam())
	if err != nil {
		return util.DtoListResponse[R]{}, err
	}

	// sort
	orderBy := GormSortDsl(dto.SortParam())
	db = db.Scopes(orderBy...)

	// result
	list := make([]R, 0, listBufferSize)
	err = db.Find(&list).Error
	if err != nil {
		return util.DtoListResponse[R]{}, GormError(err)
	}

	return util.NewDtoListResponse(list, pageResponse), nil
}
