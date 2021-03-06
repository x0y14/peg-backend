package db

import (
	typesv1 "backend/gen/types/v1"
	userv1 "backend/gen/user/v1"
	"backend/util"
	"database/sql"
	"fmt"
	"github.com/bufbuild/connect-go"
	"strings"
)

func CreateProfile(db *sql.DB, userId string, displayName string, iconPath string) (*typesv1.Profile, error) {
	prep, err := db.Prepare("INSERT INTO profiles (user_id, display_name, icon_path) values (?, ?, ?)")
	if err != nil {
		return nil, err
	}
	defer prep.Close()

	_, err = prep.Exec(userId, displayName, iconPath)
	if err != nil {
		return nil, err
	}

	return GetProfile(db, userId)
}

func GetProfile(db *sql.DB, userId string) (*typesv1.Profile, error) {
	prep, err := db.Prepare("SELECT display_name, icon_path, status_message, metadata from peg.profiles where user_id = ?")
	if err != nil {
		return nil, err
	}
	defer prep.Close()

	var displayName string
	var iconPath string
	var statusMessage string
	var metadata string
	err = prep.QueryRow(userId).Scan(&displayName, &iconPath, &statusMessage, &metadata)

	return &typesv1.Profile{
		UserId:        userId,
		DisplayName:   displayName,
		IconPath:      iconPath,
		StatusMessage: statusMessage,
		Metadata:      metadata,
	}, err
}

func UpdateProfile(db *sql.DB, userId string, request *userv1.UpdateProfileRequest) (*typesv1.Profile, error) {
	var values []interface{}
	var wheres []string

	if request.DisplayName != nil {
		wheres = append(wheres, "display_name = ?")
		values = append(values, request.GetDisplayName())
	}

	if request.IconPath != nil {
		wheres = append(wheres, "icon_path = ?")
		values = append(values, request.GetIconPath())
	}

	if request.StatusMessage != nil {
		wheres = append(wheres, "status_message = ?")
		values = append(values, request.GetStatusMessage())
	}

	if request.Metadata != nil {
		// 簡易jsonチェック.
		if !util.IsDbJSON(request.GetMetadata()) {
			return nil, connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("invalid argument: metadata"))
		}
		wheres = append(wheres, "metadata = ?")
		values = append(values, request.GetMetadata())
	}

	// 式を用意
	upd, err := db.Prepare("UPDATE profiles set " + strings.Join(wheres, ", ") + " where user_id = ?")
	if err != nil {
		return nil, err
	}
	defer upd.Close()

	// (x, y...)はできるけど(x..., y)はできないぽいので、最後に加えてあげる
	values = append(values, userId)

	/// ...で配列の中身を展開
	_, err = upd.Exec(values...)
	if err != nil {
		return nil, err
	}

	// 新しくなった物を返してあげる.
	// あんまりよろしくないかもしれないが。
	// でもexecの戻り値は、何個変更されたかなので、正当なのかもしれない。
	return GetProfile(db, userId)
}
