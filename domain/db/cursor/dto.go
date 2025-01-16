package cursor

import (
	"encoding/base64"
	"errors"
	"strconv"
	"strings"
	"time"
)

type (
	Cursor struct {
		ID        int
		Timestamp time.Time
	}
)

func NewCursor(id int, timestamp time.Time) *Cursor { return &Cursor{ID: id, Timestamp: timestamp} }

func Decode(cursor *string) (*Cursor, error) {
	if cursor != nil {
		decodedCursor, err := base64.StdEncoding.DecodeString(*cursor)
		if err != nil {
			return nil, errors.New("invalid cursor: " + err.Error())
		}
		parts := strings.SplitN(string(decodedCursor), "|", 2)
		if len(parts) != 2 {
			return nil, errors.New("invalid cursor: wrong number of parts")
		}
		id, err := strconv.Atoi(parts[0])
		if err != nil {
			return nil, errors.New("invalid cursor: " + err.Error())
		}
		timestamp, err := time.Parse(time.RFC3339, parts[1])
		if err != nil {
			return nil, errors.New("invalid cursor: " + err.Error())
		}
		return &Cursor{ID: id, Timestamp: timestamp}, nil
	}
	return nil, nil
}

func Encode(cursor Cursor) string {
	return base64.StdEncoding.EncodeToString([]byte(strconv.Itoa(cursor.ID) + "|" + cursor.Timestamp.Format(time.RFC3339)))
}
