/**
 * 功能描述: md5加密工具
 * @Date: 2019-07-25
 * @author: lixiaoming
 */
package uuid

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
)

const len = 12
const (
	clusterID = "CID-"
)

func EncodeMD5(str string) string {
	h := md5.New()
	io.WriteString(h, str)

	strSum := h.Sum(nil)
	res := hex.EncodeToString(strSum)
	return res
}

// newShortID generate short uuid and add type prefix
func generateFromString(str string, length int) string {
	md5id := EncodeMD5(str)
	return md5id[:length]
}

// newLongID generate long uuid and add type prefix
func combinedId(prefix, content string) string {
	id := generateFromString(content, len)
	return fmt.Sprintf("%s%s", prefix, id)
}

//集群id做Md5
func NewMd5ClusterId(content string) string {
	return combinedId(clusterID, content)
}

//组合集群UUId
//apiurl 接口地址
//apitoken 接口token
func CombineClusterUUId(apiurl, apitoken string) string {
	word := fmt.Sprintf("%s-%s", apiurl, apitoken)
	return NewMd5ClusterId(word)
}
