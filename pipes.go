package mgopipes

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	"github.com/globalsign/mgo/bson"
)

// ErrUnmarshalJSON is returned when there is an error unmarshalling the json
var ErrUnmarshalJSON = fmt.Errorf("Error unmarshalling")

// GetPipeline returns the corresponding []bson.M from the traditional JSON string
func GetPipeline(jsonString string, args ...interface{}) ([]bson.M, error) {
	var p []bson.M
	ps := fmt.Sprintf(jsonString, args...)
	err := json.Unmarshal([]byte(ps), &p)
	if err != nil {
		return []bson.M{}, ErrUnmarshalJSON
	}
	return p, nil
}

func getFormattedPipelineString(jsonString string, args ...interface{}) string {
	normalArgs := []interface{}{}
	for _, arg := range args {
		if !sliceOrArray(arg) {
			normalArgs = append(normalArgs, arg)
			continue
		}
		elemType := reflect.TypeOf(arg).Elem().Kind()
		stringArgs := []string{}
		switch elemType {
		case reflect.Int:
			for _, a := range arg.([]int) {
				stringArgs = append(stringArgs, fmt.Sprintf("%d", a))
			}
		case reflect.String:
			for _, a := range arg.([]string) {
				stringArgs = append(stringArgs, fmt.Sprintf(`"%s"`, a))
			}
		case reflect.Float32:
			for _, a := range arg.([]float32) {
				stringArgs = append(stringArgs, fmt.Sprintf("%f", a))
			}
		default:
			for _, a := range arg.([]interface{}) {
				stringArgs = append(stringArgs, fmt.Sprintf("%v", a))
			}
		}
		replaceString := strings.Join(stringArgs, ",")
		jsonString = strings.Replace(jsonString, `"%l"`, fmt.Sprintf("[%s]", replaceString), 1)
	}
	return fmt.Sprintf(jsonString, normalArgs...)
}

func sliceOrArray(i interface{}) bool {
	switch reflect.TypeOf(i).Kind() {
	case reflect.Slice:
		return true
	case reflect.Array:
		return true
	}
	return false
}
