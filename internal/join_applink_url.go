package internal

import (
	"net/url"
	"reflect"
	"sort"
	"strings"
)

func JoinAppLinkURL(uri string, data interface{}) string {
	x, _ := url.Parse(uri)
	arr := toURLValues(data)
	sort.Slice(arr, func(i, j int) bool {
		return arr[i][0] < arr[j][0]
	})
	query := url.Values{}
	for _, v := range arr {
		query.Set(v[0], v[1])
	}
	x.RawQuery = query.Encode()
	return x.String()
}

func toURLValues(data interface{}) [][2]string {
	vv := reflect.ValueOf(data)
	vt := reflect.TypeOf(data)
	res := [][2]string{}

	if vv.Kind() == reflect.Ptr {
		vv = vv.Elem()
		vt = vt.Elem()
	}
	for i := 0; i < vv.NumField(); i++ {
		fv := vv.Field(i)
		ft := vt.Field(i)
		if !fv.IsValid() {
			continue
		}
		jsonTag := ft.Tag.Get("json")
		if jsonTag == "" {
			continue
		}
		if strings.HasSuffix(jsonTag, ",omitempty") {
			jsonTag = jsonTag[:len(jsonTag)-len(",omitempty")]
		}
		if fv.Kind() == reflect.Ptr {
			if fv.IsNil() {
				continue
			}
			fv = fv.Elem()
		}
		if fvs := fv.String(); fvs != "" {
			res = append(res, [2]string{jsonTag, fvs})
		}
	}
	return res
}
