package output

import (
	"encoding/json"
	"fmt"
	"io"
	"reflect"
	"sort"
	"strconv"
	"text/tabwriter"
)

type Format string

const (
	JSON  Format = "json"
	Table Format = "table"
)

func Write(w io.Writer, format Format, value any) error {
	if format == Table {
		return writeTable(w, value)
	}
	data, err := json.MarshalIndent(value, "", "  ")
	if err != nil {
		return err
	}
	_, err = fmt.Fprintln(w, string(data))
	return err
}

func writeTable(w io.Writer, value any) error {
	tw := tabwriter.NewWriter(w, 0, 4, 2, ' ', 0)
	v := reflect.ValueOf(value)
	if v.Kind() == reflect.Slice || v.Kind() == reflect.Array {
		if v.Len() == 0 {
			_, err := fmt.Fprintln(tw, "(empty)")
			if err != nil {
				return err
			}
			return tw.Flush()
		}
		first := v.Index(0)
		if first.Kind() == reflect.Interface {
			first = first.Elem()
		}
		if first.Kind() == reflect.Map {
			keys := mapKeys(first)
			if _, err := fmt.Fprintln(tw, stringsJoin(keys, "\t")); err != nil {
				return err
			}
			for i := 0; i < v.Len(); i++ {
				row := v.Index(i)
				if row.Kind() == reflect.Interface {
					row = row.Elem()
				}
				values := make([]string, len(keys))
				for j, key := range keys {
					values[j] = formatValue(row.MapIndex(reflect.ValueOf(key)))
				}
				if _, err := fmt.Fprintln(tw, stringsJoin(values, "\t")); err != nil {
					return err
				}
			}
			return tw.Flush()
		}
	}
	if v.Kind() == reflect.Map {
		keys := mapKeys(v)
		if _, err := fmt.Fprintln(tw, "KEY\tVALUE"); err != nil {
			return err
		}
		for _, key := range keys {
			if _, err := fmt.Fprintf(tw, "%s\t%s\n", key, formatValue(v.MapIndex(reflect.ValueOf(key)))); err != nil {
				return err
			}
		}
		return tw.Flush()
	}
	if _, err := fmt.Fprintln(tw, formatValue(v)); err != nil {
		return err
	}
	return tw.Flush()
}

func mapKeys(v reflect.Value) []string {
	keys := make([]string, 0, v.Len())
	for _, key := range v.MapKeys() {
		keys = append(keys, fmt.Sprint(key.Interface()))
	}
	sort.Strings(keys)
	return keys
}
func formatValue(v reflect.Value) string {
	if !v.IsValid() {
		return ""
	}
	if v.Kind() == reflect.Interface {
		return formatValue(v.Elem())
	}
	if v.Kind() == reflect.String {
		return v.String()
	}
	if v.Kind() == reflect.Bool {
		return strconv.FormatBool(v.Bool())
	}
	data, err := json.Marshal(v.Interface())
	if err != nil {
		return fmt.Sprint(v.Interface())
	}
	return string(data)
}
func stringsJoin(values []string, sep string) string {
	result := ""
	for i, value := range values {
		if i > 0 {
			result += sep
		}
		result += value
	}
	return result
}

type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func WriteError(w io.Writer, err error) error {
	return Write(w, JSON, Error{Code: "command_failed", Message: err.Error()})
}
