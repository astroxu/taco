/**
 * @Author: Sean
 * @Date: 2022/5/31 15:31
 */

package facade_pattern

/*
设计思想:
	通过组合模式来实现外观模式, 为子系统实现统一的访问api
	1. Facade struct: 属性为子系统的结构体
	2. 子系统结构体
*/

// 示例： 微服务框架： 音乐服务、视频服务
// 创建子服务struct
// music struct
type Music struct {
	Name string
}

func (m *Music) GetMusic() string {
	return m.Name
}

// Video struct
type Video struct {
	Id int64
}

func (v *Video) GetVideoId() int64 {
	return v.Id
}

// count struct
type Count struct {
	Comment int64
	Praise  int64
	Collect int64
}

func (c *Count) GetCount() int64 {
	return c.Comment
}

// Facade 外观结构
type Facade struct {
	music Music
	count Count
	video Video
}

// 对外访问接口
func (f *Facade) PrintServerInfo() {
	f.music.GetMusic()
	f.video.GetVideoId()
	f.count.GetCount()
}

func NewFacade(music Music, count Count, video Video) *Facade {
	return &Facade{
		music: music,
		count: count,
		video: video,
	}
}
