package asset_manager

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"regexp"

	"GoTenancy/libs/admin"
	"GoTenancy/libs/media/oss"
	"GoTenancy/libs/qor/resource"
	"github.com/jinzhu/gorm"
)

// AssetManager defined a asset manager that could be used to manage assets in qor admin
type AssetManager struct {
	gorm.Model
	File oss.OSS `media_library:"URL:/system/assets/{{primary_key}}/{{filename_with_hash}}"`
}

// ConfigureQorResource configure qor locale for Qor Admin
func (*AssetManager) ConfigureQorResource(res resource.Resourcer) {
	if res, ok := res.(*admin.Resource); ok {
		router := res.GetAdmin().GetRouter()
		router.Post(fmt.Sprintf("/%v/upload", res.ToParam()), func(context *admin.Context) {
			result := AssetManager{}
			result.File.Scan(context.Request.MultipartForm.File["file[]"])
			context.GetDB().Save(&result)
			bytes, _ := json.Marshal(map[string]interface{}{"file-0": map[string]string{"url": result.File.URL(), "id": result.File.GetFileName()}})
			context.Writer.Write(bytes)
		})

		assetURL := regexp.MustCompile(`^/system/assets/(\d+)/`)
		router.Post(fmt.Sprintf("/%v/crop", res.ToParam()), func(context *admin.Context) {
			defer context.Request.Body.Close()
			var (
				err error
				url struct{ URL string }
				buf bytes.Buffer
			)

			io.Copy(&buf, context.Request.Body)
			if err = json.Unmarshal(buf.Bytes(), &url); err == nil {
				if matches := assetURL.FindStringSubmatch(url.URL); len(matches) > 1 {
					result := &AssetManager{}
					if err = context.GetDB().Find(result, matches[1]).Error; err == nil {
						if err = result.File.Scan(buf.Bytes()); err == nil {
							if err = context.GetDB().Save(result).Error; err == nil {
								bytes, _ := json.Marshal(map[string]interface{}{"file-0": map[string]string{"url": result.File.URL(), "id": result.File.GetFileName()}})
								context.Writer.Write(bytes)
								return
							}
						}
					}
				}
			}

			bytes, _ := json.Marshal(map[string]string{"err": err.Error()})
			context.Writer.Write(bytes)
		})
	}
}
