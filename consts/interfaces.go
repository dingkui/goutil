package consts

type IfToStr interface{ ToStr() (string, error) }
type IfToBool interface{ ToBool() (bool, error) }
type IfToBytes interface{ ToBytes() ([]byte, error) }
type IfToFloat64 interface{ ToFloat64() (float64, error) }
type IfToInt interface{ ToInt() (int, error) }
type IfToInt64 interface{ ToInt64() (int64, error) }
