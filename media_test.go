package wxm_test

import (
	"github.com/smartwalle/wxm"
	"testing"
)

func TestMiniProgram_UploadTempMedia(t *testing.T) {
	rsp, err := miniProgram.UploadTempMedia(wxm.MediaTypeOfImage, "./a.png")
	if err != nil {
		t.Error(err)
	}

	t.Log(rsp.MediaId)
}
