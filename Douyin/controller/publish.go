package controller

import (
	"Douyin/common"
	"Douyin/config"
	"Douyin/dao"
	"Douyin/service"
	"Douyin/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	//"github.com/tencentyun/cos-go-sdk-v5"
	//"golang.org/x/net/context"
	"log"
	"net/http"
	//"net/url"
	"path/filepath"
	"strconv"
)

type PublishedResponse struct {
	common.Response
	VideoLists []dao.Video `json:"video_list,omitempty"`
}

// PublishList 已发布的视频list
func PublishList(c *gin.Context) {
	//返回该用户所有视频信息
	userid, _ := strconv.Atoi(c.Query("user_id"))
	videoLists, err := service.PublishedVideoLists(int64(userid))
	//错误判断
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusOK, FeedResponse{
			Response: common.Response{StatusCode: 1, StatusMsg: err.Error()},
		})
		return
	} else {
		c.JSON(http.StatusOK, FeedResponse{
			Response:   common.Response{StatusCode: 0, StatusMsg: "successful"},
			VideoLists: videoLists,
		})
		return
	}
}

// Publish 发布视频
func Publish(c *gin.Context) {
	//返回所有视频信息
	title := c.PostForm("title")
	//得到token 获取userid
	userid := c.MustGet("userid").(int64)
	//获取文件
	file, err := c.FormFile("data")

	if err != nil {
		//得到的文件错误
		log.Println(err.Error())
		c.JSON(http.StatusOK, FeedResponse{
			Response: common.Response{StatusCode: 1, StatusMsg: err.Error()},
		})
		return
	}
	//保证了文件名的随机性并且得到文件
	filename := filepath.Base(file.Filename)
	//得到文件的后缀名
	finalName := fmt.Sprintf("%d_%s", userid, filename)
	//选择路径映射为对外的static文件
	//nginx代理静态资源
	saveFilePath := filepath.Join("./public/", finalName)
	//保存文件到对应的路径
	err = c.SaveUploadedFile(file, saveFilePath)
	if err != nil {
		//得到的文件错误
		log.Println(err.Error())
		c.JSON(http.StatusOK, FeedResponse{
			Response: common.Response{StatusCode: 1, StatusMsg: err.Error()},
		})
		return
	}

	//生成对应的快照
	savePagePath := "./public/" + finalName
	err = utils.GenerateSnapshot(saveFilePath, savePagePath, 1)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusOK, FeedResponse{
			Response: common.Response{StatusCode: 1, StatusMsg: err.Error()},
		})
		return
	}

	/*//fixme:使用对象存储
	// 初始化腾讯云COS客户端
	client := createCosClient()
	// 指定上传到COS的路径和文件名
	objectKey := "public/" + finalName
	// 创建上传请求
	_, err = client.Object.Put(context.TODO(), objectKey, file, nil)
	if err != nil {
		http.Error(w, "Failed to upload file to COS", http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "File uploaded successfully")*/

	//静态资源的地址
	playUrl := "https://" + config.C.Resource.Ipaddress + ":" + config.C.Resource.Port + "/" + "public/" + finalName
	coverUrl := "https://" + config.C.Resource.Ipaddress + ":" + config.C.Resource.Port + "/" + "public/" + finalName + ".png"
	//发布视频
	err = service.PublishVideo(userid, playUrl, coverUrl, title)

	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusOK, FeedResponse{
			Response: common.Response{StatusCode: 1, StatusMsg: err.Error()},
		})
		return
	}

	c.JSON(http.StatusOK, FeedResponse{
		Response: common.Response{StatusCode: 0, StatusMsg: "upload successful"},
	})
	return
}

/*func createCosClient() *cos.Client {
	u, _ := url.Parse("https://tencentobject-1320078852.cos-website.ap-guangzhou.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  "**************************",
			SecretKey: "**************************",
		},
	})
	return client
}
*/
