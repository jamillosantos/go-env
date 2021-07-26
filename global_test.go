package env

import (
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetString(t *testing.T) {
	wantValue := "value"
	os.Setenv("KEY", wantValue)

	assert.Equal(t, wantValue, GetString("KEY"))
}

func TestGetStringDefault(t *testing.T) {
	t.Run("key exists", func(t *testing.T) {
		wantValue := "value"
		os.Setenv("KEY", wantValue)
		assert.Equal(t, wantValue, GetStringDefault("KEY", "defaultValue"))
	})

	t.Run("key does not exists", func(t *testing.T) {
		wantValue := "value"
		os.Setenv("KEY", "")
		assert.Equal(t, wantValue, GetStringDefault("KEY", wantValue))
	})
}

func TestGetInt(t *testing.T) {
	t.Run("valid int", func(t *testing.T) {
		wantValue := 1
		os.Setenv("KEY", strconv.Itoa(wantValue))
		gotValue, err := GetInt("KEY")
		assert.NoError(t, err)
		assert.Equal(t, wantValue, gotValue)
	})

	t.Run("invalid int", func(t *testing.T) {
		os.Setenv("KEY", "invalid int")
		gotValue, err := GetInt("KEY")
		assert.Error(t, err)
		assert.Zero(t, gotValue)
	})
}

func TestGetIntDefault(t *testing.T) {
	t.Run("key exists", func(t *testing.T) {
		t.Run("valid", func(t *testing.T) {
			wantValue := 1
			os.Setenv("KEY", strconv.Itoa(wantValue))
			gotValue, err := GetIntDefault("KEY", 2)
			assert.NoError(t, err)
			assert.Equal(t, wantValue, gotValue)
		})

		t.Run("invalid", func(t *testing.T) {
			wantValue := 2
			os.Setenv("KEY", "invalid int")
			gotValue, err := GetIntDefault("KEY", wantValue)
			assert.Error(t, err)
			assert.Equal(t, gotValue, wantValue)
		})
	})

	t.Run("key does not exists", func(t *testing.T) {
		wantValue := 1
		os.Setenv("KEY", "")
		gotValue, err := GetIntDefault("KEY", wantValue)
		assert.NoError(t, err)
		assert.Equal(t, gotValue, wantValue)
	})
}

func TestGetUint(t *testing.T) {
	t.Run("valid uint", func(t *testing.T) {
		wantValue := uint(1)
		os.Setenv("KEY", strconv.FormatUint(uint64(wantValue), 10))
		gotValue, err := GetUint("KEY")
		assert.NoError(t, err)
		assert.Equal(t, wantValue, gotValue)
	})

	t.Run("invalid uint", func(t *testing.T) {
		os.Setenv("KEY", "invalid uint")
		gotValue, err := GetUint("KEY")
		assert.Error(t, err)
		assert.Zero(t, gotValue)
	})
}

func TestGetUintDefault(t *testing.T) {
	t.Run("key exists", func(t *testing.T) {
		t.Run("valid", func(t *testing.T) {
			wantValue := uint(1)
			os.Setenv("KEY", strconv.FormatUint(uint64(wantValue), 10))
			gotValue, err := GetUintDefault("KEY", 2)
			assert.NoError(t, err)
			assert.Equal(t, wantValue, gotValue)
		})

		t.Run("invalid", func(t *testing.T) {
			wantValue := uint(2)
			os.Setenv("KEY", "invalid uint")
			gotValue, err := GetUintDefault("KEY", wantValue)
			assert.Error(t, err)
			assert.Equal(t, gotValue, wantValue)
		})
	})

	t.Run("key does not exists", func(t *testing.T) {
		wantValue := uint(1)
		os.Setenv("KEY", "")
		gotValue, err := GetUintDefault("KEY", wantValue)
		assert.NoError(t, err)
		assert.Equal(t, gotValue, wantValue)
	})
}

func TestGetInt64(t *testing.T) {
	t.Run("valid int64", func(t *testing.T) {
		wantValue := int64(1)
		os.Setenv("KEY", strconv.FormatInt(wantValue, 10))
		gotValue, err := GetInt64("KEY")
		assert.NoError(t, err)
		assert.Equal(t, wantValue, gotValue)
	})

	t.Run("invalid int64", func(t *testing.T) {
		os.Setenv("KEY", "invalid int64")
		gotValue, err := GetInt64("KEY")
		assert.Error(t, err)
		assert.Zero(t, gotValue)
	})
}

func TestGetInt64Default(t *testing.T) {
	t.Run("key exists", func(t *testing.T) {
		t.Run("valid", func(t *testing.T) {
			wantValue := int64(1)
			os.Setenv("KEY", strconv.FormatInt(wantValue, 10))
			gotValue, err := GetInt64Default("KEY", 2)
			assert.NoError(t, err)
			assert.Equal(t, wantValue, gotValue)
		})

		t.Run("invalid", func(t *testing.T) {
			wantValue := int64(2)
			os.Setenv("KEY", "invalid int64")
			gotValue, err := GetInt64Default("KEY", wantValue)
			assert.Error(t, err)
			assert.Equal(t, gotValue, wantValue)
		})
	})

	t.Run("key does not exists", func(t *testing.T) {
		wantValue := int64(1)
		os.Setenv("KEY", "")
		gotValue, err := GetInt64Default("KEY", wantValue)
		assert.NoError(t, err)
		assert.Equal(t, gotValue, wantValue)
	})
}

func TestGetUint64(t *testing.T) {
	t.Run("valid uint64", func(t *testing.T) {
		wantValue := uint64(1)
		os.Setenv("KEY", strconv.FormatUint(wantValue, 10))
		gotValue, err := GetUint64("KEY")
		assert.NoError(t, err)
		assert.Equal(t, wantValue, gotValue)
	})

	t.Run("invalid uint64", func(t *testing.T) {
		os.Setenv("KEY", "invalid uint64")
		gotValue, err := GetUint64("KEY")
		assert.Error(t, err)
		assert.Zero(t, gotValue)
	})
}

func TestGetUint64Default(t *testing.T) {
	t.Run("key exists", func(t *testing.T) {
		t.Run("valid", func(t *testing.T) {
			wantValue := uint64(1)
			os.Setenv("KEY", strconv.FormatUint(wantValue, 10))
			gotValue, err := GetUint64Default("KEY", 2)
			assert.NoError(t, err)
			assert.Equal(t, wantValue, gotValue)
		})

		t.Run("invalid", func(t *testing.T) {
			wantValue := uint64(2)
			os.Setenv("KEY", "invalid uint64")
			gotValue, err := GetUint64Default("KEY", wantValue)
			assert.Error(t, err)
			assert.Equal(t, gotValue, wantValue)
		})
	})

	t.Run("key does not exists", func(t *testing.T) {
		wantValue := uint64(1)
		os.Setenv("KEY", "")
		gotValue, err := GetUint64Default("KEY", wantValue)
		assert.NoError(t, err)
		assert.Equal(t, gotValue, wantValue)
	})
}

func TestGetFloat(t *testing.T) {
	t.Run("valid float64", func(t *testing.T) {
		wantValue := float64(1)
		os.Setenv("KEY", strconv.FormatFloat(wantValue, 'f', -1, 64))
		gotValue, err := GetFloat("KEY")
		assert.NoError(t, err)
		assert.Equal(t, wantValue, gotValue)
	})

	t.Run("invalid float64", func(t *testing.T) {
		os.Setenv("KEY", "invalid float64")
		gotValue, err := GetFloat("KEY")
		assert.Error(t, err)
		assert.Zero(t, gotValue)
	})
}

func TestGetFloatDefault(t *testing.T) {
	t.Run("key exists", func(t *testing.T) {
		t.Run("valid", func(t *testing.T) {
			wantValue := float64(1)
			os.Setenv("KEY", strconv.FormatFloat(wantValue, 'f', -1, 64))
			gotValue, err := GetFloatDefault("KEY", 2)
			assert.NoError(t, err)
			assert.Equal(t, wantValue, gotValue)
		})

		t.Run("invalid", func(t *testing.T) {
			wantValue := float64(2)
			os.Setenv("KEY", "invalid float64")
			gotValue, err := GetFloatDefault("KEY", wantValue)
			assert.Error(t, err)
			assert.Equal(t, gotValue, wantValue)
		})
	})

	t.Run("key does not exists", func(t *testing.T) {
		wantValue := float64(1)
		os.Setenv("KEY", "")
		gotValue, err := GetFloatDefault("KEY", wantValue)
		assert.NoError(t, err)
		assert.Equal(t, gotValue, wantValue)
	})
}

func TestGetBool(t *testing.T) {
	t.Run("valid bool", func(t *testing.T) {
		wantValue := true
		os.Setenv("KEY", strconv.FormatBool(wantValue))
		gotValue, err := GetBool("KEY")
		assert.NoError(t, err)
		assert.Equal(t, wantValue, gotValue)
	})

	t.Run("invalid bool", func(t *testing.T) {
		os.Setenv("KEY", "invalid bool")
		gotValue, err := GetBool("KEY")
		assert.Error(t, err)
		assert.False(t, gotValue)
	})

	t.Run("kkey does not exists", func(t *testing.T) {
		os.Setenv("KEY", "")
		gotValue, err := GetBool("KEY")
		assert.NoError(t, err)
		assert.False(t, gotValue)
	})

}

func TestGetBoolDefault(t *testing.T) {
	t.Run("key exists", func(t *testing.T) {
		t.Run("valid", func(t *testing.T) {
			wantValue := true
			os.Setenv("KEY", strconv.FormatBool(wantValue))
			gotValue, err := GetBoolDefault("KEY", false)
			assert.NoError(t, err)
			assert.Equal(t, wantValue, gotValue)
		})

		t.Run("invalid", func(t *testing.T) {
			wantValue := false
			os.Setenv("KEY", "invalid bool")
			gotValue, err := GetBoolDefault("KEY", wantValue)
			assert.Error(t, err)
			assert.Equal(t, gotValue, wantValue)
		})
	})

	t.Run("key does not exists", func(t *testing.T) {
		wantValue := true
		os.Setenv("KEY", "")
		gotValue, err := GetBoolDefault("KEY", wantValue)
		assert.NoError(t, err)
		assert.Equal(t, gotValue, wantValue)
	})
}

func TestGetDuration(t *testing.T) {
	t.Run("valid duration", func(t *testing.T) {
		wantValue := time.Second
		os.Setenv("KEY", wantValue.String())
		gotValue, err := GetDuration("KEY")
		assert.NoError(t, err)
		assert.Equal(t, wantValue, gotValue)
	})

	t.Run("invalid duration", func(t *testing.T) {
		os.Setenv("KEY", "invalid duration")
		gotValue, err := GetDuration("KEY")
		assert.Error(t, err)
		assert.Zero(t, gotValue)
	})
}

func TestGetDurationDefault(t *testing.T) {
	t.Run("key exists", func(t *testing.T) {
		t.Run("valid", func(t *testing.T) {
			wantValue := time.Second
			os.Setenv("KEY", "1s")
			gotValue, err := GetDurationDefault("KEY", time.Second*2)
			assert.NoError(t, err)
			assert.Equal(t, wantValue, gotValue)
		})

		t.Run("invalid", func(t *testing.T) {
			wantValue := time.Second
			os.Setenv("KEY", "invalid duration")
			gotValue, err := GetDurationDefault("KEY", wantValue)
			assert.Error(t, err)
			assert.Equal(t, gotValue, wantValue)
		})
	})

	t.Run("key does not exists", func(t *testing.T) {
		wantValue := time.Duration(1)
		os.Setenv("KEY", "")
		gotValue, err := GetDurationDefault("KEY", wantValue)
		assert.NoError(t, err)
		assert.Equal(t, gotValue, wantValue)
	})
}
