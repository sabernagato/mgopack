package mgopack

import mgo "gopkg.in/mgo.v2"

// MgoConnect ...
type MgoConnect struct {
	URL        string
	DataBase   string
	MgoSession *mgo.Session
}

// GetCloneSession ...
func (mgoc *MgoConnect) GetCloneSession() (*mgo.Session, error) {
	if mgoc.MgoSession == nil {
		var err error
		mgoc.MgoSession, err = mgo.Dial(mgoc.URL)
		if err != nil {
			//panic(err) //直接终止程序运行
			return nil, err
		}
	}
	//最大连接池默认为4096
	return mgoc.MgoSession.Clone(), nil
}

// GetCopySession ...
func (mgoc *MgoConnect) GetCopySession() (*mgo.Session, error) {
	if mgoc.MgoSession == nil {
		var err error
		mgoc.MgoSession, err = mgo.Dial(mgoc.URL)
		if err != nil {
			//panic(err) //直接终止程序运行
			return nil, err
		}
	}
	//最大连接池默认为4096
	return mgoc.MgoSession.Copy(), nil
}
