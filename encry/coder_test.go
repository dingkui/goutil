package encry_test

import (
	"gitee.com/dk83/goutils/dlog"
	"gitee.com/dk83/goutils/encry"
	"testing"
)

var data = []byte("local")

func TestHashs(t *testing.T) {
	bytes, err := encry.Md5.SumBytes(data)
	dlog.Info(bytes, err)
	bytes, err = encry.Sha1.SumBytes(data)
	dlog.Info(bytes, err)
	bytes, err = encry.Sha256.SumBytes(data)
	dlog.Info(bytes, err)
	bytes, err = encry.Sha512.SumBytes(data)
	dlog.Info(bytes, err)
	bytes, err = encry.HSha512.SumBytes(data)
	dlog.Info(bytes, err)
}

func TestCoders(t *testing.T) {

	bytes, _ := encry.Xor.EncodeBytes(data)
	dlog.Info("encry.Xor.EncodeBytes:", bytes.String())
	bytes, _ = encry.Xor.DecodeBytes(bytes.Bytes())
	dlog.Info("encry.Xor.DecodeBytes:", bytes.String())

	bytes, _ = encry.Base64.EncodeBytes(data)
	dlog.Info("encry.Base64.EncodeBytes:", bytes.String())
	bytes, _ = encry.Base64.DecodeBytes(bytes.Bytes())
	dlog.Info("encry.Base64.DecodeBytes:", bytes.String())

	bytes, _ = encry.Xbase.EncodeBytes(data)
	dlog.Info("encry.Xbase.EncodeBytes:", bytes.String())
	bytes, _ = encry.Xbase.DecodeBytes(bytes.Bytes())
	dlog.Info("encry.Xbase.DecodeBytes:", bytes.String())

	bytes, _ = encry.Xbase.EncodeBytesByPass(data, []byte("DEMO"))
	dlog.Info("encry.Xbase.EncodeBytes:", bytes.String())
	bytes, _ = encry.Xbase.DecodeBytesByPass(bytes.Bytes(), []byte("DEMO"))
	dlog.Info("encry.Xbase.DecodeBytes:", bytes.String())

	bytes, _ = encry.PXor.EncodeBytesByPass(data, []byte("DEMO"))
	dlog.Info("encry.PXor.EncodeBytes:", bytes.String())
	bytes, e := encry.PXor.DecodeBytesByPass(bytes.Bytes(), []byte("DEMO"))
	dlog.Info("encry.PXor.DecodeBytes:", bytes.String(), e)

	bytes, _ = encry.PXor2.EncodeBytesByPass(data, []byte("DEMO"))
	dlog.Info("encry.PXor2.EncodeBytes:", bytes.String())
	bytes, e = encry.PXor2.DecodeBytesByPass(bytes.Bytes(), []byte("DEMO"))
	dlog.Info("encry.PXor2.DecodeBytes:", bytes.String(), e)
}
func TestCoders1(t *testing.T) {

	bytes, _ := encry.Base64.EncodeBytes(data)
	dlog.Info("encry.Base64.EncodeBytes:", bytes.String())
	bytes, _ = encry.Base64.DecodeBytes(bytes.Bytes())
	dlog.Info("encry.Base64.DecodeBytes:", bytes.String())

	bytes, _ = encry.Xbase.EncodeBytes(data)
	dlog.Info("encry.Xbase.EncodeBytes:", bytes.String())
	bytes, _ = encry.Xbase.DecodeBytes(bytes.Bytes())
	dlog.Info("encry.Xbase.DecodeBytes:", bytes.String())

	bytes, _ = encry.Xbase.EncodeBytesByPass(data, []byte("DEMO"))
	dlog.Info("encry.Xbase.EncodeBytes:", bytes.String())
	bytes, _ = encry.Xbase.DecodeBytesByPass(bytes.Bytes(), []byte("DEMO"))
	dlog.Info("encry.Xbase.DecodeBytes:", bytes.String())

	bytes, _ = encry.PXor.EncodeBytesByPass(data, []byte("DEMO"))
	dlog.Info("encry.PXor.EncodeBytes:", bytes.String())
	bytes, e := encry.PXor.DecodeBytesByPass(bytes.Bytes(), []byte("DEMO"))
	dlog.Info("encry.PXor.DecodeBytes:", bytes.String(), e)

	bytes, _ = encry.PXor2.EncodeBytesByPass(data, []byte("DEMO"))
	dlog.Info("encry.PXor2.EncodeBytes:", bytes.String())
	bytes, e = encry.PXor2.DecodeBytesByPass(bytes.Bytes(), []byte("DEMO"))
	dlog.Info("encry.PXor2.DecodeBytes:", bytes.String(), e)
}
func TestRsa(t *testing.T) {
	puKey, prKey, _ := encry.RsaUtil.GenerateKeys(2048)
	dlog.Info("GenerateKeys:", string(puKey), string(prKey))

	bytes, e := encry.RsaCoder.EncodeBytesByPass(data, puKey)
	dlog.Info("encry.RsaCoder.EncodeBytes:", bytes.String(), e)
	bytes2, e := encry.RsaCoder.DecodeBytesByPass(bytes.Bytes(), prKey)
	dlog.Info("encry.RsaCoder.DecodeBytes:", bytes2.String(), e)

	bytes, e = encry.RsaCoder.EncodeBytesByPass(data, puKey)
	dlog.Info("encry.RsaCoder.EncodeBytes:", bytes.String(), e)
	bytes, e = encry.RsaCoder.DecodeBytesByPass(bytes.Bytes(), prKey)
	dlog.Info("encry.RsaCoder.DecodeBytes:", bytes.String(), e)
}

func TestMd5(t *testing.T) {

	dlog.Info("md5:%s", encry.Md5.SumFileN("D:\\code\\go\\chan3dgoserver\\utils\\appconf\\appconf_test.go"))
}
