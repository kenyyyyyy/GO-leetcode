package router

import (
	"errors"
	"strings"
)

type node struct {
	path     string           // 路由路径
	part     string           // 由'/'分隔的部分
	children map[string]*node // 子节点
	isWild   bool             // 通配符节点
}

type Router struct {
	root map[string]*node // 路由树根节点
}

var router = new(Router)

func NewRouter() *Router {
	router = &Router{make(map[string]*node)}
	return router
}

// parsePath 分隔路由为part
// 比如路由/hello/:name将被分隔为["hello", ":name"]
func parsePath(path string) (parts []string) {
	// 将path以"/"分隔为parts
	par := strings.Split(path, "/")
	for _, p := range par {
		if p != "" {
			parts = append(parts, p)
			// 如果part是以通配符*开头的
			if p[0] == '*' {
				break
			}
		}
	}
	return
}

// AddRoute:增加路由
func (r *Router) AddRoute(method, path string) {
	parts := parsePath(path)
	// 如果该方法的路由树没有就初始化一个
	if _, ok := r.root[method]; !ok {
		r.root[method] = &node{children: make(map[string]*node)}
	}
	root := r.root[method]
	// 将parts插入到路由树
	for _, part := range parts {
		// 如果part没有在路由树中，初始化一个node
		if root.children[part] == nil {
			root.children[part] = &node{
				part:     part,
				children: make(map[string]*node),
				isWild:   part[0] == ':' || part[0] == '*'}
		}
		// root往下遍历
		root = root.children[part]
	}
	// 遍历完成后注册路由path
	root.path = path
}

// GetRoute:获取路由节点
func (r *Router) GetRoute(method, path string) (node *node, err error) {
	searchParts := parsePath(path)
	// 获取方法的路由树
	var ok bool
	if node, ok = r.root[method]; !ok {
		return nil, errors.New("没有找到该路径")
	}
	// 在该方法的路由树上查找该路径
	for _, part := range searchParts {
		// 查找child是否等于part
		var temp string
		for _, child := range node.children {
			if child.part == part || child.isWild {
				temp = child.part
				break
			}
		}
		// 没有找到part
		if temp == "" {
			return nil, errors.New("没有找到该路径")
		}
		// 遇到通配符*，直接返回
		if temp[0] == '*' {
			return node.children[temp], nil
		}
		// 查找下一个part
		node = node.children[temp]
	}
	// 最后判断添加路由时是否注册path
	if node.path == "" {
		return nil, errors.New("没有找到该路径")
	}
	return
}
