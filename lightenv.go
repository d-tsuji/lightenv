package lightenv

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

// Process receives and sets a type
// that defines a list that gets its value from an environment variable.
//
// Types that can be received as environment variables support string, int, float32, and float64.
func Process(input interface{}) error {
	infos, err := gatherInfo(input)
	if err != nil {
		return err
	}
	for _, info := range infos {
		value, ok := os.LookupEnv(info.Key)

		req := info.Tags.Get("required")
		if !ok {
			if isTrue(req) {
				return fmt.Errorf("required key %s missing value", info.Key)
			}
		}

		def := info.Tags.Get("default")
		if def != "" && !ok {
			value = def
		}

		if err := setParameter(value, info.Field); err != nil {
			return err
		}
	}
	return nil
}

// varInfo maintains information about the configuration variable
type varInfo struct {
	Name  string
	Key   string
	Field reflect.Value
	Tags  reflect.StructTag
}

// GatherInfo gathers information about the specified struct
func gatherInfo(spec interface{}) ([]varInfo, error) {
	s := reflect.ValueOf(spec)

	if s.Kind() != reflect.Ptr {
		return nil, fmt.Errorf("ErrInvalidSpecification")
	}
	s = s.Elem()
	if s.Kind() != reflect.Struct {
		return nil, fmt.Errorf("ErrInvalidSpecification")
	}
	typeOfSpec := s.Type()

	// over allocate an info array, we will extend if needed later
	infos := make([]varInfo, 0, s.NumField())
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		ftype := typeOfSpec.Field(i)
		if !f.CanSet() {
			continue
		}

		// Capture information about the config variable
		info := varInfo{
			Name:  ftype.Name,
			Field: f,
			Tags:  ftype.Tag,
		}

		// Default to the field name as the env var name (will be upcased)
		info.Key = info.Name
		info.Key = strings.ToUpper(info.Key)

		// Overwrite from struct tag parameter
		tname := ftype.Tag.Get("name")
		if tname != "" {
			info.Key = ftype.Tag.Get("name")
		}

		infos = append(infos, info)
	}
	return infos, nil
}

func setParameter(value string, field reflect.Value) error {
	// Set key parameter
	typ := field.Type()
	switch typ.Kind() {
	case reflect.String:
		field.SetString(value)
	case reflect.Int:
		val, err := strconv.ParseInt(value, 0, typ.Bits())
		if err != nil {
			return err
		}
		field.SetInt(val)
	case reflect.Float32, reflect.Float64:
		val, err := strconv.ParseFloat(value, typ.Bits())
		if err != nil {
			return err
		}
		field.SetFloat(val)
	default:
		return fmt.Errorf("type %v is not supported", typ.Kind())
	}
	return nil
}

func isTrue(s string) bool {
	b, _ := strconv.ParseBool(s)
	return b
}
