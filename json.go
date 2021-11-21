package ky3k

import (
	"encoding/json"
	"errors"

	"gorm.io/datatypes"
)

func JsonToUnmarshal(v datatypes.JSON, items interface{}) error {
	if v != nil {
		if d, err := v.Value(); err == nil {
			if err = json.Unmarshal([]byte(d.(string)), items); err == nil {
				return nil
			} else {
				return err
			}
		} else {
			return err
		}
	} else {
		return errors.New("传入为空值")
	}
}

func StringToJson(s string, data interface{}) error {

	if err := json.Unmarshal([]byte(s), &data); err != nil {
		return err
	}

	return nil
}

func JsonToString(items interface{}) string {

	if src, err := json.Marshal(items); err == nil {
		return string(src)
	}

	return ""
}

func JsonToJsonb(items interface{}) datatypes.JSON {
	if src, err := json.Marshal(items); err == nil {
		return src
	}
	return nil
}
