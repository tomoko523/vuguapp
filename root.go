package main

// DO NOT EDIT: This file was generated by vugu. Please regenerate instead of editing or add additional code in a separate file.

import "fmt"
import "reflect"
import "github.com/vugu/vugu"

type RootData struct {
	Url string
}

func (comp *Root) NewData(props vugu.Props) (interface{}, error) {
	url := "https://random.dog/woof.json"
	return &RootData{Url: url}, nil
}

var _ vugu.ComponentType = (*Root)(nil)

func (comp *Root) BuildVDOM(dataI interface{}) (vdom *vugu.VGNode, css *vugu.VGNode, reterr error) {
	data := dataI.(*RootData)
	_ = data
	_ = fmt.Sprint
	_ = reflect.Value{}
	event := vugu.DOMEventStub
	_ = event
	css = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "style", DataAtom: vugu.VGAtom(458501), Namespace: "", Attr: []vugu.VGAttribute(nil)}
	css.AppendChild(&vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n", DataAtom: vugu.VGAtom(0), Namespace: "", Attr: []vugu.VGAttribute(nil)})
	var n *vugu.VGNode
	n = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "div", DataAtom: vugu.VGAtom(92931), Namespace: "", Attr: []vugu.VGAttribute{vugu.VGAttribute{Namespace: "", Key: "class", Val: "root"}}}
	vdom = n
	{
		parent := n
		n = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n    ", DataAtom: vugu.VGAtom(0), Namespace: "", Attr: []vugu.VGAttribute(nil)}
		parent.AppendChild(n)
		n = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "random-animal", DataAtom: vugu.VGAtom(0), Namespace: "", Attr: []vugu.VGAttribute{vugu.VGAttribute{Namespace: "", Key: "target", Val: "わんこ"}}}
		parent.AppendChild(n)
		n.Props = vugu.Props{
			"url": data.Url,
		}
		{
			parent := n
			n = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n", DataAtom: vugu.VGAtom(0), Namespace: "", Attr: []vugu.VGAttribute(nil)}
			parent.AppendChild(n)
		}
	}
	return
}

type Root struct {}

func init() { vugu.RegisterComponentType("root", &Root{}) }
