package creationalPatterns

import (
	"bytes"
	"fmt"
	"io"
	"reflect"
	"testing"
)

// 原型模式
// 用于创建重复的对象，同时又能保证性能。
// 是实现了一个原型接口，该接口用于创建当前对象的克隆

// example:

// Paper 报纸接口，包含读取和克隆方法，作为原型模式的接口
type Paper interface {
	io.Reader     // 读取
	Clone() Paper // 克隆
}

// NewsPaper 报纸，实现原型接口
type NewsPaper struct {
	headline string
	content  string
}

func (np *NewsPaper) Read(p []byte) (n int, err error) {
	buf := bytes.NewBufferString(fmt.Sprintf("headline:%s,content:%s", np.headline, np.content))
	return buf.Read(p)
}

func (np *NewsPaper) Clone() Paper {
	return &NewsPaper{
		headline: np.headline + "_clone",
		content:  np.content,
	}
}

func NewNewsPaper(headline, content string) *NewsPaper {
	return &NewsPaper{headline: headline, content: content}
}

var _ Paper = (*NewsPaper)(nil)

// Resume 简历，实现原型接口
type Resume struct {
	name       string
	age        int
	experience string
}

func (r *Resume) Read(p []byte) (n int, err error) {
	buf := bytes.NewBufferString(fmt.Sprintf("name:%s,age:%d,experience:%s", r.name, r.age, r.experience))
	return buf.Read(p)
}

func (r *Resume) Clone() Paper {
	return &Resume{
		name:       r.name + "_clone",
		age:        r.age,
		experience: r.experience,
	}
}

func NewResume(name string, age int, experience string) *Resume {
	return &Resume{name: name, age: age, experience: experience}
}

var _ Paper = (*Resume)(nil)

// 使用

func TestPrototype(t *testing.T) {
	copier := NewCopier("云打印机")
	oneNewspaper := NewNewsPaper("Go是最好的编程语言", "Go语言十大优势")
	oneResume := NewResume("小明", 29, "5年码农")

	otherNewspaper := copier.Copy(oneNewspaper)
	copyNewspaperMsg := make([]byte, 100)
	byteSize, _ := otherNewspaper.Read(copyNewspaperMsg)
	fmt.Println("copyNewspaperMsg:" + string(copyNewspaperMsg[:byteSize]))

	otherResume := copier.Copy(oneResume)
	copyResumeMsg := make([]byte, 100)
	byteSize, _ = otherResume.Read(copyResumeMsg)
	fmt.Println("copyResumeMsg:" + string(copyResumeMsg[:byteSize]))
}

// Copier 复印机
type Copier struct {
	name string
}

func NewCopier(name string) *Copier {
	return &Copier{name: name}
}

func (c *Copier) Copy(p Paper) Paper {
	fmt.Printf("copy name:%s is copying: %v\n", c.name, reflect.TypeOf(p).String())
	return p.Clone()
}
