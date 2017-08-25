/*
MIT License

Copyright (c) 2017 Simon Schmidt

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/


package fastnntp

import "sync"

type Group struct{
	Group []byte
	Number int64
	Low int64
	High int64
}
var pool_Group = sync.Pool{New: func()interface{} { return new(Group) }}
func pool_Group_put(g *Group) {
	g.Group = nil
	pool_Group.Put(g)
}

type GroupCaps interface {
	GetGroup(g *Group) bool
	ListGroup(g *Group,w *DotWriter,first,last int64)
	CursorMoveGroup(g *Group,i int64,backward bool,id_buf []byte) (ni int64,id []byte,ok bool)
}

type Article struct{
	MessageId []byte
	Group []byte
	Number int64
	HasId bool
	HasNum bool
}
type ArticleCaps interface {
	// Every method must set the message-id on success, if it is not given.
	
	StatArticle(a *Article) bool
	GetArticle(a *Article,head, body bool) func(w *DotWriter)
}

type Handler struct {
	GroupCaps
	ArticleCaps
}
func (h *Handler) fill() {
	if h.GroupCaps==nil { h.GroupCaps = DefaultCaps }
	if h.ArticleCaps==nil { h.ArticleCaps = DefaultCaps }
}

var DefaultCaps = new(defCaps)

type defCaps struct {}
// GroupCaps
func (d *defCaps) GetGroup(g *Group) bool { return false }
func (d *defCaps) ListGroup(g *Group,w *DotWriter,first,last int64) { }
func (d *defCaps) CursorMoveGroup(g *Group,i int64,backward bool, id_buf []byte) (ni int64,id []byte,ok bool) { ok = false; return }

// ArticleCaps
func (d *defCaps) StatArticle(a *Article) bool { return false }
func (d *defCaps) GetArticle(a *Article,head, body bool) func(w *DotWriter) { return nil }

