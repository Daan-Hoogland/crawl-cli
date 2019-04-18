package pkg

import "testing"

func Test_GenerateUrl(t *testing.T) {
	res := GenerateURL("127.0.0.1", 9000, nil, false, "connect")
	exp := "http://127.0.0.1:9000/connect"
	if res != exp {
		t.Errorf("get generate url = %s want %s", res, exp)
	}
}

func Test_GenerateUrlSSL(t *testing.T) {
	res := GenerateURL("127.0.0.1", 9000, nil, true, "connect")
	exp := "https://127.0.0.1:9000/connect"
	if res != exp {
		t.Errorf("get generate url = %s want %s", res, exp)
	}
}

func Test_GenerateUrlParametersSingle(t *testing.T) {
	res := GenerateURL("127.0.0.1", 9000, map[string]string{"test": "testing"}, true, "connect")
	exp := "https://127.0.0.1:9000/connect?test=testing"
	if res != exp {
		t.Errorf("get generate url = %s want %s", res, exp)
	}
}

func Test_GenerateUrlParametersMultiple(t *testing.T) {
	res := GenerateURL("127.0.0.1", 9000, map[string]string{"test": "testing", "test2": "testing2"}, true, "connect")
	exp := "https://127.0.0.1:9000/connect?test=testing&test2=testing2"
	if res != exp {
		t.Errorf("get generate url = %s want %s", res, exp)
	}
}
