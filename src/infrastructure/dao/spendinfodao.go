package dao

import (
	"collectionneur-cli/src/domain/entities"
	"database/sql"
	"regexp"
	"strconv"
	"time"
)

const costMessageReString = `^Priorbank\. Karta (?P<card>[0-9*]+)\. (?P<date>[0-9\- :]+)\. Oplata (?P<sum>[0-9\.]+) (?P<currency>[\w]{3})\. (?P<mname>[\w 0-9]+)`

type SpendInfoDAO struct {
	chatId int

	db            *sql.DB
	costMessageRe *regexp.Regexp
}

// FIXME: Sort all the things in memory because there is not so much data in that db
// First place for optimisation
func (d *SpendInfoDAO) GetSpendInfoBefore(dt time.Time) ([]*entities.SpendInfo, error) {
	infos := []*entities.SpendInfo{}

	rows, err := d.db.Query("SELECT text FROM message WHERE handle_id=? ORDER BY -date", d.chatId)
	if err != nil {
		return nil, err
	}

	var messageText string
	for rows.Next() {
		err = rows.Scan(&messageText)
		if err != nil {
			return nil, err
		}

		info := d.convertToEntity(messageText)
		if info == nil {
			continue
		}

		if !info.Date.After(dt) {
			break
		}

		infos = append(infos, info)
	}

	return infos, nil
}

// TODO: maybe here I need to process error, but I don't know what to do with this error
func (d *SpendInfoDAO) convertToEntity(raw string) *entities.SpendInfo {
	items := d.costMessageRe.FindStringSubmatch(raw)
	if len(items) == 0 {
		return nil
	}

	layout := "02-01-06 15:04:05"
	dt, err := time.Parse(layout, items[2])
	if err != nil {
		return nil
	}
	sum, err := strconv.ParseFloat(items[3], 64)
	if err != nil {
		return nil
	}

	return &entities.SpendInfo{
		Card:        items[1],
		Date:        dt,
		Sum:         float32(sum),
		Currency:    items[4],
		Description: items[5],
	}
}

func NewSpendInfoDAO(db *sql.DB, chatId int) *SpendInfoDAO {
	dao := &SpendInfoDAO{}
	dao.chatId = chatId
	dao.db = db
	dao.costMessageRe = regexp.MustCompile(costMessageReString)
	return dao
}
