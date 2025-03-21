package helpers

import (
	"context"
	"encoding/json"
	"fmt"
	"math"
	"net"
	"reflect"
	"strconv"
	"strings"

	"google.golang.org/grpc/peer"
)

func IsEmptyValue(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.String:
		return v.Len() == 0
	case reflect.Array, reflect.Slice, reflect.Map, reflect.Chan:
		return v.Len() == 0
	case reflect.Ptr, reflect.Interface:
		return v.IsNil()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Float32, reflect.Float64:
		return v.IsZero()
	case reflect.Bool:
		return !v.Bool()
	}
	return false
}

func ReflactTagsScylla(data interface{}) ([]string, string, []string) {
	var (
		filledTags []string
		pk         string
		allTags    []string
	)

	t := reflect.TypeOf(data)
	v := reflect.ValueOf(data)
	if t.Kind() != reflect.Struct {
		panic("GetscyllaTagsCached: provided data is not a struct")
	}

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		scyllaTag := field.Tag.Get("scylla")
		fieldValue := v.Field(i)

		if strings.Contains(scyllaTag, "primaryKey") {
			splitStringPk := strings.Split(scyllaTag, ";")
			if len(splitStringPk) > 0 {
				pk = splitStringPk[0]
				fmt.Println("Primary Key Field:", pk)
				allTags = append(allTags, pk)
			}
			continue
		}

		if IsEmptyValue(fieldValue) {
			allTags = append(allTags, scyllaTag)
			continue
		}

		if scyllaTag != "" {
			filledTags = append(filledTags, scyllaTag)
			allTags = append(allTags, scyllaTag)
		}

	}

	return filledTags, pk, allTags
}

func ConvertWhereClauseToCQL(primaryKeys []string) string {
	var whereClause []string

	for _, data := range primaryKeys {
		keyWhereClause := data + " IS NOT NULL AND"
		whereClause = append(whereClause, keyWhereClause)
	}
	whereClause[len(whereClause)-1] = strings.TrimSuffix(whereClause[len(whereClause)-1], " AND")
	return strings.Join(whereClause, " ")
}

func ChangeStrcutToMapInterfaceFilters(req interface{}) (map[string]interface{}, error) {
	raw, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	filters := make(map[string]interface{})
	err = json.Unmarshal(raw, &filters)
	if err != nil {
		return nil, err
	}

	// Filter field kosong
	for key, value := range filters {
		if value == "" || value == 0 {
			delete(filters, key)
		}
	}
	return filters, nil
}

func PaginationHelpers(countData string, pageNumber string, size string) map[string]interface{} {
	floTotalRow, _ := strconv.ParseFloat(countData, 64)
	floRecordPerPage, _ := strconv.ParseFloat(size, 64)
	floCurrentPage, _ := strconv.ParseFloat(pageNumber, 64)
	totalPage := math.Ceil(floTotalRow / floRecordPerPage)
	currentPage := pageNumber
	firstRecord := (floCurrentPage - 1) * floRecordPerPage
	startRecord := firstRecord + 1
	pagination := map[string]interface{}{
		"total_record":    countData,
		"total_page":      strconv.FormatFloat(float64(totalPage), 'f', 0, 64),
		"record_per_page": size,
		"current_page":    currentPage,
		"start_record":    strconv.FormatFloat(startRecord, 'f', 0, 64),
		"first_record":    strconv.FormatFloat(firstRecord, 'f', 0, 64),
	}
	return pagination
}

func GetPeerIP(ctx context.Context) string {
	p, ok := peer.FromContext(ctx)
	if !ok {
		return "unknown"
	}

	addr, ok := p.Addr.(*net.TCPAddr)
	if !ok {
		return "unknown"
	}

	return addr.IP.String()
}
