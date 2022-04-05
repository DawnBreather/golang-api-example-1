package utils

func MustParseInterfaceToFloat64(object interface{}) (res float64){
  if object == nil {
    res = -1.0
  } else {
    res = object.(float64)
  }

  return
}